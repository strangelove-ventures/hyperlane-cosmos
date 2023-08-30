package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	igptypes "github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"

	ismtypes "github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types/legacy_multisig"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types/merkle_root_multisig"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types/message_id_multisig"
	mailboxtypes "github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"

	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/icza/dyno"

	interchaintest "github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/testreporter"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

var (
	pathHyperlaneGaia   = "hyperlane-gaia" // Replace with 2nd cosmos chain supporting hyperlane
	genesisWalletAmount = int64(10_000_000)
)

// hyperlaneEncoding registers the Hyperlane specific module codecs so that the associated types and msgs
// will be supported when writing to the blocksdb sqlite database.
func hyperlaneEncoding() *testutil.TestEncodingConfig {
	cfg := cosmos.DefaultEncoding()

	// register custom types
	wasmtypes.RegisterInterfaces(cfg.InterfaceRegistry)
	mailboxtypes.RegisterInterfaces(cfg.InterfaceRegistry)
	ismtypes.RegisterInterfaces(cfg.InterfaceRegistry)
	merkle_root_multisig.RegisterInterfaces(cfg.InterfaceRegistry)
	message_id_multisig.RegisterInterfaces(cfg.InterfaceRegistry)
	legacy_multisig.RegisterInterfaces(cfg.InterfaceRegistry)
	igptypes.RegisterInterfaces(cfg.InterfaceRegistry)

	return &cfg
}

func CreateSingleHyperlaneSimd(t *testing.T) []ibc.Chain {
	// Create chain factory with hyperlane-simd

	votingPeriod := "10s"
	maxDepositPeriod := "10s"

	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			ChainName: "hsimd",
			ChainConfig: ibc.ChainConfig{
				Type:    "cosmos",
				Name:    "hsimd",
				ChainID: "hsimd-1",
				Images: []ibc.DockerImage{
					{
						Repository: "hyperlane-simd",
						Version:    "local",
						UidGid:     "1025:1025",
					},
				},
				Bin:            "simd",
				Bech32Prefix:   "cosmos",
				Denom:          "stake",
				CoinType:       "118",
				GasPrices:      "0.00stake",
				GasAdjustment:  1.8,
				TrustingPeriod: "112h",
				NoHostMount:    false,
				// ConfigFileOverrides: nil,
				EncodingConfig:         hyperlaneEncoding(),
				ModifyGenesis:          ModifyGenesisProposalTime(votingPeriod, maxDepositPeriod),
				UsingNewGenesisCommand: true,
			},
		},
	})

	// Get chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	return chains
}

func CreateDoubleHyperlaneSimd(t *testing.T, image ibc.DockerImage) []ibc.Chain {
	// Create chain factory with hyperlane-simd

	votingPeriod := "10s"
	maxDepositPeriod := "10s"

	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			ChainName: "simd-1",
			ChainConfig: ibc.ChainConfig{
				Type:    "cosmos",
				Name:    "simd",
				ChainID: "simd-1",
				Images: []ibc.DockerImage{
					image,
				},
				Bin:            "simd",
				Bech32Prefix:   "cosmos",
				Denom:          "stake",
				CoinType:       "118",
				GasPrices:      "0.00stake",
				GasAdjustment:  1.8,
				TrustingPeriod: "112h",
				NoHostMount:    false,
				// ConfigFileOverrides: nil,
				EncodingConfig:         hyperlaneEncoding(),
				ModifyGenesis:          ModifyGenesisProposalTime(votingPeriod, maxDepositPeriod),
				UsingNewGenesisCommand: true,
			},
		},
		{
			ChainName: "simd-2",
			ChainConfig: ibc.ChainConfig{
				Type:    "cosmos",
				Name:    "simd",
				ChainID: "simd-2",
				Images: []ibc.DockerImage{
					image,
				},
				Bin:            "simd",
				Bech32Prefix:   "cosmos",
				Denom:          "stake",
				CoinType:       "118",
				GasPrices:      "0.00stake",
				GasAdjustment:  1.8,
				TrustingPeriod: "112h",
				NoHostMount:    false,
				// ConfigFileOverrides: nil,
				EncodingConfig:         hyperlaneEncoding(),
				ModifyGenesis:          ModifyGenesisProposalTime(votingPeriod, maxDepositPeriod),
				UsingNewGenesisCommand: true,
			},
		},
	})

	// Get chains from the chain factory
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	return chains
}

func BuildInitialChain(t *testing.T, chains []ibc.Chain) context.Context {
	// Create a new Interchain object which describes the chains, relayers, and IBC connections we want to use
	ic := interchaintest.NewInterchain()

	for _, chain := range chains {
		ic.AddChain(chain)
	}

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)

	ctx := context.Background()
	client, network := interchaintest.DockerSetup(t)

	err := ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:          t.Name(),
		Client:            client,
		NetworkID:         network,
		SkipPathCreation:  true,
		BlockDatabaseFile: interchaintest.DefaultBlockDatabaseFilepath(),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = ic.Close()
	})

	return ctx
}

func ModifyGenesisProposalTime(votingPeriod string, maxDepositPeriod string) func(ibc.ChainConfig, []byte) ([]byte, error) {
	return func(chainConfig ibc.ChainConfig, genbz []byte) ([]byte, error) {
		g := make(map[string]interface{})
		if err := json.Unmarshal(genbz, &g); err != nil {
			return nil, fmt.Errorf("failed to unmarshal genesis file: %w", err)
		}
		if err := dyno.Set(g, votingPeriod, "app_state", "gov", "params", "voting_period"); err != nil {
			return nil, fmt.Errorf("failed to set voting period in genesis json: %w", err)
		}
		if err := dyno.Set(g, maxDepositPeriod, "app_state", "gov", "params", "max_deposit_period"); err != nil {
			return nil, fmt.Errorf("failed to set max deposit period in genesis json: %w", err)
		}
		if err := dyno.Set(g, chainConfig.Denom, "app_state", "gov", "params", "min_deposit", 0, "denom"); err != nil {
			return nil, fmt.Errorf("failed to set min deposit in genesis json: %w", err)
		}
		out, err := json.Marshal(g)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal genesis bytes to json: %w", err)
		}
		return out, nil
	}
}
