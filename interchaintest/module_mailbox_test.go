package interchaintest

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"
	"github.com/stretchr/testify/require"

	helpers "github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/helpers"
)

// TestHyperlaneMailbox ensures the mailbox module & bindings work properly.
func TestHyperlaneMailbox(t *testing.T) {
	t.Parallel()

	// Base setup
	chains := CreateSingleHyperlaneSimd(t)
	ctx := BuildInitialChain(t, chains)

	// Chains
	simd := chains[0].(*cosmos.CosmosChain)
	t.Log("simd.GetHostRPCAddress()", simd.GetHostRPCAddress())

	users := interchaintest.GetAndFundTestUsers(t, ctx, "default", int64(10_000_000), simd)
	user := users[0]

	msg := fmt.Sprintf(`{}`)
	_, contract := helpers.SetupContract(t, ctx, simd, user.KeyName(), "contracts/hyperlane.wasm", msg)
	t.Log("coreContract", contract)

	verifyContractEntryPoints(t, ctx, simd, user, contract)

	processMsgStruct := helpers.ExecuteMsg{
		ProcessMsg: &helpers.ProcessMsg{
			Msg: "MsgProcessedByContract",
		},
	}
	processMsg, err := json.Marshal(processMsgStruct)
	require.NoError(t, err)
	simd.ExecuteContract(ctx, user.KeyName(), contract, string(processMsg))

	dispatchMsgStruct := helpers.ExecuteMsg{
		DispatchMsg: &helpers.DispatchMsg{
			DestinationAddr: 1,
			RecipientAddr:   "cosmos10qa7yajp3fp869mdegtpap5zg056exja3chkw5",
			MessageBody:     "MsgDispatchedByContract",
		},
	}
	dipatchMsg, err := json.Marshal(dispatchMsgStruct)
	require.NoError(t, err)
	simd.ExecuteContract(ctx, user.KeyName(), contract, string(dipatchMsg))

	err = testutil.WaitForBlocks(ctx, 2, simd)
	require.NoError(t, err)
}

func verifyContractEntryPoints(t *testing.T, ctx context.Context, simd *cosmos.CosmosChain, user ibc.Wallet, contract string) {
	queryMsg := helpers.QueryMsg{Owner: &struct{}{}}
	var queryRsp helpers.QueryRsp
	err := simd.QueryContract(ctx, contract, queryMsg, &queryRsp)
	require.NoError(t, err)
	require.Equal(t, user.FormattedAddress(), queryRsp.Data.Address)

	randomAddr := "cosmos10qa7yajp3fp869mdegtpap5zg056exja3chkw5"
	newContractOwnerStruct := helpers.ExecuteMsg{
		ChangeContractOwner: &helpers.ChangeContractOwner{
			NewOwner: randomAddr,
		},
	}
	newContractOwner, err := json.Marshal(newContractOwnerStruct)
	require.NoError(t, err)
	simd.ExecuteContract(ctx, user.KeyName(), contract, string(newContractOwner))

	err = simd.QueryContract(ctx, contract, queryMsg, &queryRsp)
	require.NoError(t, err)
	require.Equal(t, randomAddr, queryRsp.Data.Address)
}
