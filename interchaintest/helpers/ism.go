package helpers

import (
	"context"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	ismtypes "github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types/merkle_root_multisig"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/counterchain"
)

func SetDefaultIsm(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, keyName string, counterChain *counterchain.CounterChain) {
	proposal := cosmos.TxProposalv1{
		Metadata: "none",
		Deposit:  "500000000" + chain.Config().Denom, // greater than min deposit
		Title:    "Set hyperlane default ISM",
		Summary:  "Set hyperlane default ISM",
	}

	var valSet []string
	for _, val := range counterChain.ValSet.Vals {
		valSet = append(valSet, val.Addr)
	}

	message := ismtypes.MsgSetDefaultIsm{
		Signer: sdk.MustBech32ifyAddressBytes(chain.Config().Bech32Prefix, authtypes.NewModuleAddress(govtypes.ModuleName)),
		Isms: []*ismtypes.Ism{
			{
				Origin: 1,
				AbstractIsm: ismtypes.MustPackAbstractIsm(
					&merkle_root_multisig.MerkleRootMultiSig{
						Threshold:        2,
						ValidatorPubKeys: valSet,
					},
				),
			},
		},
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
