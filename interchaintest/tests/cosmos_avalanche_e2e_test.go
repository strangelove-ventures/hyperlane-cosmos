package tests

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/ava-labs/coreth/accounts/abi/bind"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/counterchain"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/docker"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/helpers"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/tests/ethcontracts/announce"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/tests/ethcontracts/ism/legacy_multisig"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/tests/ethcontracts/mailbox"
	icv7 "github.com/strangelove-ventures/interchaintest/v7"
	interchaintest "github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/chain/avalanche"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	hyperlane "github.com/strangelove-ventures/interchaintest/v7/chain/hyperlane"
	subnetevm "github.com/strangelove-ventures/interchaintest/v7/examples/avalanche/subnet-evm"
	"golang.org/x/sync/errgroup"

	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/testreporter"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

const (
	COSMOS_AVA_E2E      = "hyperlane-cosmos-avalanche.yaml"
	AVA_FUNDED_TEST_KEY = "56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027" // the Avalanche genesis funded private key
)

// Set up the solidity smart contracts on our Avalanche node
func TestConfigureAvalanche(t *testing.T) {
	avaPrivateKey, _ := crypto.HexToECDSA(AVA_FUNDED_TEST_KEY)
	avalancheAddr := crypto.PubkeyToAddress(avaPrivateKey.PublicKey)
	logger := NewLogger(t)

	// Start setup for Docker network
	// Mailbox address - will be used in the hyperlane validator config
	mailboxHex, expectedMailbox := helpers.GetMailboxAddress()
	prefixedMailboxHex := "0x" + mailboxHex

	// Directories where files related to this test will be stored
	val1TmpDir := t.TempDir()
	val2TmpDir := t.TempDir()
	rlyTmpDir := t.TempDir()

	// Get the hyperlane agent raw configs (before variable replacements)
	valSimd1, valAvalanche, rly := readHyperlaneConfig(t, COSMOS_AVA_E2E, "hyperlane-validator-simd1", "hyperlane-avalanche-validator", logger)

	// Get the validator key for the agents. We also need this key to configure the chain ISM.
	valSimd1PrivKey, err := getHyperlaneBaseValidatorKey(valSimd1)
	require.NoError(t, err)
	valSimd2PrivKey, err := getHyperlaneBaseValidatorKey(valAvalanche)
	require.NoError(t, err)

	// Build the chain docker image from the local repo
	optionalBuildChainImage()

	simd1Domain := uint32(23456)
	avalancheDomain := uint32(34567)

	// Set up one Cosmos chain (with our hyperlane modules) for the test.
	// The image must be in our local registry so we skip image pull.
	chains := CreateHyperlaneSimds(t, ibc.DockerImage{
		Repository: docker.HyperlaneImageName,
		Version:    "local",
		UidGid:     "1025:1025",
	}, []uint32{simd1Domain})
	simd1 := chains[0].(*cosmos.CosmosChain)
	simd1.SkipImagePull = true

	// Set up the Avalanche chain
	nv := 5
	nf := 0

	avaChain, err := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			Name:    "avalanche",
			Version: "v1.10.1",
			ChainConfig: ibc.ChainConfig{
				ChainID: "neto-123123",
				Images: []ibc.DockerImage{
					ibc.DockerImage{
						Repository: "avalanche",
						Version:    "v1.10.1",
						UidGid:     "1025:1025",
					},
				},
				AvalancheSubnets: []ibc.AvalancheSubnetConfig{
					{
						Name:                "subnetevm",
						VM:                  subnetevm.VM,
						Genesis:             subnetevm.Genesis,
						SubnetClientFactory: subnetevm.NewSubnetEvmClient,
					},
				},
			},
			NumFullNodes:  &nf,
			NumValidators: &nv,
		},
	},
	).Chains(t.Name())

	avalancheChain := avaChain[0].(*avalanche.AvalancheChain)
	chains = append(chains, avalancheChain)
	avalancheNode := avalancheChain.Node()

	// Create a new Interchain object which describes the chains, relayers, and IBC connections we want to use
	ic := interchaintest.NewInterchain()

	for _, chain := range chains {
		ic.AddChain(chain)
	}

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)
	ctx := context.Background()

	// Note: make sure that both the 'ic' interchain AND the hyperlane network share this client/network
	client, network := interchaintest.DockerSetup(t)
	opts := interchaintest.InterchainBuildOptions{
		TestName:         t.Name(),
		Client:           client,
		NetworkID:        network,
		SkipPathCreation: true,
	}

	err = ic.Build(ctx, eRep, opts)
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = ic.Close()
	})

	// End setup for Docker network

	// Query the cosmos chains and ensure the hyperlane domain configuration is as expected
	verifyDomain(t, ctx, simd1, uint64(simd1Domain))

	// Now we need to configure the Hyperlane modules and setup some test users...
	userSimd1 := icv7.GetAndFundTestUsers(t, ctx, "default", int64(10_000_000_000), simd1)[0]

	// Fund this wallet since the validator will announce its storage location with this mnemonic's private key.
	_, err = icv7.GetAndFundTestUserWithMnemonic(ctx, "valannounce", mnemonic, int64(10_000_000_000), simd1)
	require.NoError(t, err)

	// Set up the contract that will receive a message on simd1
	msg := `{}`
	_, contract := helpers.SetupContract(t, ctx, simd1, userSimd1.KeyName(), "../contracts/hyperlane.wasm", msg)
	logger.Info("simd1 contract", zap.String("address", contract))
	verifyContractEntryPoints(t, ctx, simd1, userSimd1, contract)
	// TODO: set up the contract that will receive a message on avalanche

	// Create counter chain 1 with val set signing legacy multisig.
	// The private key used here MUST be the one from the validator config file.
	simd1IsmValidator := counterchain.CreateEmperorValidator(t, simd1Domain, counterchain.LEGACY_MULTISIG, valSimd1PrivKey)
	// Create counter chain 2 with val set signing legacy multisig
	avaIsmValidator := counterchain.CreateEmperorValidator(t, avalancheDomain, counterchain.LEGACY_MULTISIG, valSimd2PrivKey)

	// Set default isms for counter chains
	helpers.SetDefaultIsm(t, ctx, simd1, userSimd1.KeyName(), avaIsmValidator)
	// TODO: isms for Avalanche
	// helpers.SetDefaultIsm(t, ctx, simd2, userSimd2.KeyName(), simd1IsmValidator)
	var avaRecipientContract common.Address
	recipientDispatch := hexutil.Encode(avaRecipientContract.Bytes())
	fmt.Printf("Recipient dispatch addr hex: %s", recipientDispatch)
	logger.Info("Preconfiguring Hyperlane (getting configs)")

	valJson, err := preconfigureHyperlaneValidator(t, valSimd1, val1TmpDir, mnemonicPrivKey, chains[0].Config().ChainID, chains[0].Config().Name, chains[0].GetRPCAddress(), "http://"+chains[0].GetGRPCAddress(), prefixedMailboxHex, 23456)
	require.NoError(t, err)

	simd1MailboxHex, err := getMailbox(valJson, chains[0].Config().Name)
	require.NoError(t, err)
	_, simd1MailboxUnprefixed, found := strings.Cut(simd1MailboxHex, "0x")
	require.True(t, found)
	simd1Mailbox, err := hex.DecodeString(simd1MailboxUnprefixed)
	require.NoError(t, err)
	originMailboxB := []byte(simd1Mailbox)
	require.Equal(t, expectedMailbox, originMailboxB)
	avalancheRpcEndpoint := avalancheChain.GetChainRPCAddress(0)

	valJson, err = preconfigureAvalancheValidator(t, valAvalanche, val2TmpDir, AVA_FUNDED_TEST_KEY, avalancheChain.Config().ChainID,
		avalancheChain.Config().Name, avalancheRpcEndpoint, prefixedMailboxHex, avalancheDomain)
	require.NoError(t, err)
	simd1ValidatorSignaturesDir := filepath.Join(val1TmpDir, "signatures-"+chains[0].Config().Name)          //${val_dir}/signatures-${chainName}
	avalancheValidatorSignaturesDir := filepath.Join(val2TmpDir, "signatures-"+avalancheChain.Config().Name) //${val_dir}/signatures-${chainName}

	// Our images are currently local. You must build locally in monorepo, e.g. "cd rust && docker build .".
	// Also make sure that the tags in hyperlane.yaml match the local docker image repo and version.
	hyperlaneNetwork := hyperlane.NewHyperlaneNetwork(true, true)
	err = hyperlaneNetwork.Build(ctx, logger, eRep, opts, *valSimd1, *valAvalanche, *rly)
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	ec, err := avalancheNode.GetSubnetEvmClient(ctx, 0)
	require.NoError(t, err)

	networkId, err := ec.ChainID(ctx)
	require.NoError(t, err)
	auth, err := bind.NewKeyedTransactorWithChainID(avaPrivateKey, networkId)
	require.NoError(t, err)

	// Deploy the mailbox contract and wait for the deployment to succeed
	mailboxAddr, mailboxTx, mailboxContract, err := mailbox.DeployMailbox(auth, ec, avalancheDomain)
	require.NoError(t, err)
	mailboxTxReceipt, err := awaitAvalancheTx(ctx, ec, mailboxTx)
	require.NoError(t, err)
	require.Equal(t, mailboxTxReceipt.Status, 0)

	// Deploy the legacy multisig ISM contract and wait for the deployment to succeed
	legacyMultisigAddr, legacyMultisigTx, legacyMultisigContract, err := legacy_multisig.DeployLegacyMultisig(auth, ec)
	require.NoError(t, err)
	legacyMultisigTxReceipt, err := awaitAvalancheTx(ctx, ec, legacyMultisigTx)
	require.NoError(t, err)
	require.Equal(t, legacyMultisigTxReceipt.Status, 0)

	avaValidatorPrivateKey, _ := crypto.HexToECDSA(valSimd1PrivKey)
	avaValAddr := crypto.PubkeyToAddress(avaValidatorPrivateKey.PublicKey)

	// Set the validator address for the remote domain (simd1)
	legacyMultisigContract.EnrollValidator(auth, simd1Domain, avaValAddr)
	// Set a threshold of 1 so a single validator can sign
	legacyMultisigContract.SetThreshold(auth, simd1Domain, 1)

	// Set the owner and default ISM
	mailboxContract.Initialize(auth, avalancheAddr, legacyMultisigAddr)

	// Deploy the validator announce contract and wait for the deployment to succeed
	announceAddr, announceTx, _, err := announce.DeployAnnounce(auth, ec, mailboxAddr)
	require.NoError(t, err)
	announceTxReceipt, err := awaitAvalancheTx(ctx, ec, announceTx)
	require.NoError(t, err)
	require.Equal(t, announceTxReceipt.Status, 0)

	avalancheMailboxAddrHex := hexutil.Encode(mailboxAddr.Bytes())
	avalancheAnnounceAddrHex := hexutil.Encode(announceAddr.Bytes())

	rlyCfgs := []RelayerChainConfig{
		&cosmosRelayerChainCfg{
			privKey:                mnemonicPrivKey,
			chainID:                simd1.Config().ChainID,
			chainName:              simd1.Config().Name,
			rpcUrl:                 simd1.GetRPCAddress(),
			grpcUrl:                "http://" + simd1.GetGRPCAddress(),
			originMailboxHex:       prefixedMailboxHex,
			domain:                 23456,
			validatorSignaturePath: simd1ValidatorSignaturesDir,
		},
		&avalancheRelayerChainCfg{
			privKey:                AVA_FUNDED_TEST_KEY,
			chainID:                avalancheChain.Config().ChainID,
			chainName:              avalancheChain.Config().Name,
			rpcUrl:                 avalancheRpcEndpoint,
			originMailboxHex:       avalancheMailboxAddrHex,
			domain:                 avalancheDomain,
			validatorSignaturePath: avalancheValidatorSignaturesDir,
			validatorAnnounceAddr:  avalancheAnnounceAddrHex,
		},
	}
	_, err = preconfigureHyperlaneRelayer(t, rly, rlyTmpDir, rlyCfgs)
	require.NoError(t, err)

	// Give the hyperlane validators and relayer time to start up and start watching the mailbox for the chain
	time.Sleep(10 * time.Second)

	var eg errgroup.Group
	numMsgs := 1

	for i := 0; i < numMsgs; i++ {
		i := i
		// Dispatch on SIMD1
		eg.Go(func() (err error) {
			dispatchedRecipientAddrHex, dispatchedMsgBody, dispatchSender, dispatchedMsgId := dispatchMsg(t, i, ctx, avalancheDomain, recipientDispatch, userSimd1.KeyName(), contract, simd1, logger)
			return avalancheProcessMsg(t, ctx, simd1IsmValidator, avalancheChain, simd1Domain, avalancheDomain, dispatchedRecipientAddrHex, dispatchedMsgBody, dispatchSender, dispatchedMsgId, mailboxContract)
		})

		// Dispatch on Avalanche
		// eg.Go(func() (err error) {
		// 	dispatchedRecipientAddrHex, dispatchedMsgBody, dispatchSender, dispatchedMsgId := dispatchMsg(t, i, ctx, simd1Domain, recipientDispatch, userSimd2.KeyName(), contract2, simd2, logger)
		// 	return processMsg(t, ctx, simd2IsmValidator, simd1, simd2Domain, simd1Domain, dispatchedRecipientAddrHex, dispatchedMsgBody, dispatchSender, dispatchedMsgId)
		// })

		// Here we sleep for 5s to avoid sequence related issues with TX signatures
		time.Sleep(5 * time.Second)
	}

	err = eg.Wait()
	require.NoError(t, err)
}

// This API call is deprecated according to https://docs.avax.network/reference/avalanchego/p-chain/api.
// There is no explanation of what the new API is, though, even in the release notes. So for now we use this.
// I believe we may need this to look up the blockchain ID for subnet-evm so we can query it later.
// But for now we are only using C-chain queries which are simpler, "http://127.0.0.1:9650/ext/bc/C/rpc"
func getBlockchains(t *testing.T) {
	jsonBody := []byte(`{"jsonrpc": "2.0","id": 1,"method": "platform.getBlockchains", "params" :[]}`)
	bodyReader := bytes.NewReader(jsonBody)
	subnetEvmUri := fmt.Sprintf("http://127.0.0.1:9650/ext/bc/%s", "P")
	resp, err := http.Post(subnetEvmUri, "application/json", bodyReader)
	require.NoError(t, err)
	b, _ := io.ReadAll(resp.Body)
	fmt.Printf("%+v", string(b))
}

func avalancheProcessMsg(
	t *testing.T,
	ctx context.Context,
	validator *counterchain.CounterChain,
	destChain *avalanche.AvalancheChain,
	originDomain uint32,
	destDomain uint32,
	dispatchedRecipientAddrHex,
	dispatchedMsgBody,
	dispatchSender,
	dispatchedMsgId string,
	mailboxContract *mailbox.Mailbox,
) error {
	// Find the message ID of the dispatched message and wait for the message to be processed on the destination.
	dispatchedRecipientAddr := hexutil.MustDecode(dispatchedRecipientAddrHex)
	bech32Recipient := sdk.MustBech32ifyAddressBytes(sdk.GetConfig().GetBech32AccountAddrPrefix(), dispatchedRecipientAddr)

	b, err := hexutil.Decode(dispatchedMsgBody)
	require.NoError(t, err)
	message, _ := validator.CreateMessage(dispatchSender, originDomain, destDomain, bech32Recipient, string(b))
	messageId := validator.GetMessageId(message)
	require.Equal(t, dispatchedMsgId, hexutil.Encode(messageId))

	return Await(func() (bool, error) {
		var msg [32]byte
		copy(msg[:], message)
		ctx := context.Background()
		ctx, _ = context.WithTimeout(ctx, 5*time.Second)
		return mailboxContract.Delivered(
			&bind.CallOpts{}, msg)
	}, 10*time.Minute, 5*time.Second)
}

// GetDefaultChainURI returns the default chain URI for a given blockchainID
func GetDefaultChainURI(nodeURI, blockchainID string) string {
	return fmt.Sprintf("%s/ext/bc/%s/rpc", nodeURI, blockchainID)
}

// Spins up a Cosmos node, Avalanche node, and hyperlane validator for Avalanche.
// The validator writes checkpoints to a local file. We read messages from the checkpoints
// and deliver each message to the Cosmos node. No messages are delivered from Cosmos->Ava.
func TestHyperlaneAvalancheCosmosDispatch(t *testing.T) {
	tmpDir1 := t.TempDir()
	buildsEnabled := true
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)
	tarFilePath := filepath.Join(path, "../../")
	goModPath := filepath.Join(path, "../../go.mod")
	hyperlaneConfigPath := filepath.Join(path, "hyperlane-cosmos-avalanche.yaml")

	// TODO: better caching mechanism to prevent rebuilding the same image
	if buildsEnabled {
		// Builds the hyperlane image from the current project (e.g. locally).
		// The directory at 'tarFilePath' will be tarballed and used for the docker context.
		// The args 'buildDir' and 'dockerfilePath' are relative to 'tarFilePath' (the context).
		// Build arguments are derived from 'goModPath' so it must be a full path (not relative).
		docker.BuildHeighlinerHyperlaneImage(docker.HyperlaneImageName, tarFilePath, ".", goModPath, "local.Dockerfile")
	}

	DockerImage := ibc.DockerImage{
		Repository: docker.HyperlaneImageName,
		Version:    "local",
		UidGid:     "1025:1025",
	}

	domain := uint32(23456)
	// Create one chain, simd1, that will be part of the Cosmos <-> Avalanche hyperlane network
	chains := CreateHyperlaneSimds(t, DockerImage, []uint32{domain})

	// Create a new Interchain object which describes the chains
	ic := interchaintest.NewInterchain()

	for _, chain := range chains {
		ic.AddChain(chain)
	}

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)
	ctx := context.Background()

	// Note: make sure that both the 'ic' interchain AND the hyperlane network share this client/network
	client, network := interchaintest.DockerSetup(t)
	opts := interchaintest.InterchainBuildOptions{
		TestName:          t.Name(),
		Client:            client,
		NetworkID:         network,
		SkipPathCreation:  true,
		BlockDatabaseFile: interchaintest.DefaultBlockDatabaseFilepath(),
	}

	err := ic.Build(ctx, eRep, opts)
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = ic.Close()
	})

	hyperlaneCfg, err := hyperlane.ReadHyperlaneConfig(hyperlaneConfigPath, zaptest.NewLogger(t))
	require.NoError(t, err)
	valSimd1, ok := hyperlaneCfg["hyperlane-avalanche-validator"]
	require.True(t, ok)

	mailboxHex := "000000000000000000000000cc2a110c8df654a38749178a04402e88f65091d3"
	prefixedMailboxHex := "0x" + mailboxHex
	require.NoError(t, err)

	_, err = preconfigureHyperlaneValidator(t, valSimd1, tmpDir1, "privKey", "simd1", "simd1", chains[0].GetHostRPCAddress(), chains[0].GetHostGRPCAddress(), prefixedMailboxHex, domain)
	require.NoError(t, err)

	logger := NewLogger(t)

	// Our images are currently local. You must build locally in monorepo, e.g. "cd rust && docker build".
	// Also make sure that the tags in hyperlane.yaml match the local docker image repo and version.
	hyperlaneNetwork := hyperlane.NewHyperlaneNetwork(false, true)
	hyperlaneNetwork.Build(ctx, logger, eRep, opts, *valSimd1)
}
