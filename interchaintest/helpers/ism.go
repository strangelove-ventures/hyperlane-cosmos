package helpers

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/counterchain"
	ismtypes "github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types/legacy_multisig"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types/merkle_root_multisig"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types/message_id_multisig"
)

func GetDefaultIsms(counterChains ...*counterchain.CounterChain) (isms []*ismtypes.DefaultIsm) {
	for _, counterChain := range counterChains {
		var valSet []string
		for _, val := range counterChain.ValSet.Vals {
			valSet = append(valSet, val.Addr)
		}
		var ism *codectypes.Any
		switch counterChain.IsmType {
		case counterchain.LEGACY_MULTISIG:
			ism = ismtypes.MustPackAbstractIsm(
				&legacy_multisig.LegacyMultiSig{
					Threshold:        uint32(counterChain.ValSet.Threshold),
					ValidatorPubKeys: valSet,
				},
			)
		case counterchain.MERKLE_ROOT_MULTISIG:
			ism = ismtypes.MustPackAbstractIsm(
				&merkle_root_multisig.MerkleRootMultiSig{
					Threshold:        uint32(counterChain.ValSet.Threshold),
					ValidatorPubKeys: valSet,
				},
			)
		case counterchain.MESSAGE_ID_MULTISIG:
			ism = ismtypes.MustPackAbstractIsm(
				&message_id_multisig.MessageIdMultiSig{
					Threshold:        uint32(counterChain.ValSet.Threshold),
					ValidatorPubKeys: valSet,
				},
			)
		}
		isms = append(isms, &ismtypes.DefaultIsm{
			Origin:      counterChain.Domain,
			AbstractIsm: ism,
		})
	}
	return isms
}

func SetDefaultIsm(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, keyName string, counterChains ...*counterchain.CounterChain) {
	proposal := cosmos.TxProposalv1{
		Metadata: "none",
		Deposit:  "500000000" + chain.Config().Denom, // greater than min deposit
		Title:    "Set hyperlane default ISM",
		Summary:  "Set hyperlane default ISM",
	}

	message := ismtypes.MsgSetDefaultIsm{
		Signer: sdk.MustBech32ifyAddressBytes(chain.Config().Bech32Prefix, authtypes.NewModuleAddress(govtypes.ModuleName)),
		Isms:   GetDefaultIsms(counterChains...),
	}
	msg, err := chain.Config().EncodingConfig.Codec.MarshalInterfaceJSON(&message)
	fmt.Println("Msg: ", string(msg))
	require.NoError(t, err)
	proposal.Messages = append(proposal.Messages, msg)

	tx, err := chain.SubmitProposal(ctx, keyName, proposal)
	require.NoError(t, err)

	height, err := chain.Height(ctx)
	require.NoError(t, err, "error fetching height before submit upgrade proposal")

	err = chain.VoteOnProposalAllValidators(ctx, tx.ProposalID, cosmos.ProposalVoteYes)
	require.NoError(t, err, "failed to submit votes")

	_, err = cosmos.PollForProposalStatus(ctx, chain, height, height+10, tx.ProposalID, cosmos.ProposalStatusPassed)
	require.NoError(t, err, "proposal status did not change to passed in expected number of blocks")

	err = testutil.WaitForBlocks(ctx, 1, chain)
	require.NoError(t, err)
}

func CreateCustomIsm(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, keyName string, counterChain *counterchain.CounterChain) uint32 {
	var valSet []string
	for _, val := range counterChain.ValSet.Vals {
		valSet = append(valSet, val.Addr)
	}

	var ism ismtypes.AbstractIsm
	switch counterChain.IsmType {
	case counterchain.LEGACY_MULTISIG:
		ism = &legacy_multisig.LegacyMultiSig{
				Threshold:        uint32(counterChain.ValSet.Threshold),
				ValidatorPubKeys: valSet,
			}
	case counterchain.MERKLE_ROOT_MULTISIG:
		ism = &merkle_root_multisig.MerkleRootMultiSig{
				Threshold:        uint32(counterChain.ValSet.Threshold),
				ValidatorPubKeys: valSet,
			}
	case counterchain.MESSAGE_ID_MULTISIG:
		ism = &message_id_multisig.MessageIdMultiSig{
				Threshold:        uint32(counterChain.ValSet.Threshold),
				ValidatorPubKeys: valSet,
			}
	}
	
	msgBz, err := chain.Config().EncodingConfig.Codec.MarshalJSON(ism)
	fmt.Println("Msg: ", string(msgBz))
	require.NoError(t, err)

	cmd := []string{
		"simd", "tx", "hyperlane-ism", "create-multisig-ism", counterChain.IsmType,
		string(msgBz),
		"--node", chain.GetRPCAddress(),
		"--home", chain.HomeDir(),
		"--chain-id", chain.Config().ChainID,
		"--from", keyName,
		"--gas", "2500000",
		"--gas-adjustment", "2.0",
		"--keyring-dir", chain.HomeDir(),
		"--keyring-backend", keyring.BackendTest,
		"-y",
	}
	stdout, stderr, err := chain.Exec(ctx, cmd, nil)
	require.NoError(t, err)

	fmt.Println("CreateCustomIsm stdout: ", string(stdout))
	fmt.Println("CreateCustomIsm stderr: ", string(stderr))

	err = testutil.WaitForBlocks(ctx, 2, chain)
	require.NoError(t, err)

	ismTxHash := ParseTxHash(string(stdout))
	
	events, err := GetEvents(chain, ismTxHash)
	require.NoError(t, err)
	ismId, found := GetEventAttribute(events, ismtypes.EventTypeCreateCustomIsm, ismtypes.AttributeKeyIndex)
	require.True(t, found)
	ismId32, err := strconv.ParseUint(ismId, 10, 32)
	require.NoError(t, err)
	return uint32(ismId32)
}

func QueryAllDefaultIsms(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain) *ismtypes.QueryAllDefaultIsmsResponse {
	grpcAddress := chain.GetHostGRPCAddress()
	conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()

	queryClient := ismtypes.NewQueryClient(conn)
	res, err := queryClient.AllDefaultIsms(ctx, &ismtypes.QueryAllDefaultIsmsRequest{})
	require.NoError(t, err)

	return res
}

func QueryCustomIsm(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, ismId uint32) *ismtypes.QueryCustomIsmResponse {
	grpcAddress := chain.GetHostGRPCAddress()
	conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()

	queryClient := ismtypes.NewQueryClient(conn)
	res, err := queryClient.CustomIsm(ctx, &ismtypes.QueryCustomIsmRequest{IsmId: ismId})
	require.NoError(t, err)

	return res
}

func QueryAllCustomIsms(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain) *ismtypes.QueryAllCustomIsmsResponse {
	grpcAddress := chain.GetHostGRPCAddress()
	conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()

	queryClient := ismtypes.NewQueryClient(conn)
	res, err := queryClient.AllCustomIsms(ctx, &ismtypes.QueryAllCustomIsmsRequest{})
	require.NoError(t, err)

	return res
}
