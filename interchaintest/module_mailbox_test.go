package ictest

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"
	"github.com/stretchr/testify/require"

	helpers "github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/helpers"
	ismtypes "github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types/merkle_root_multisig"
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

	users := interchaintest.GetAndFundTestUsers(t, ctx, "default", int64(10_000_000_000), simd)
	user := users[0]

	msg := fmt.Sprintf(`{}`)
	_, contract := helpers.SetupContract(t, ctx, simd, user.KeyName(), "contracts/hyperlane.wasm", msg)
	t.Log("coreContract", contract)

	verifyContractEntryPoints(t, ctx, simd, user, contract)

	counterChain := helpers.CreateCounterChain(t, 1)
	helpers.SetDefaultIsm(t, ctx, simd, user.KeyName(), counterChain)
	res := helpers.QueryAllDefaultIsms(t, ctx, simd)

	var abstractIsm ismtypes.AbstractIsm
	err := simd.Config().EncodingConfig.InterfaceRegistry.UnpackAny(res.DefaultIsms[0].AbstractIsm, &abstractIsm)
	require.NoError(t, err)
	merkleRootMultiSig := abstractIsm.(*merkle_root_multisig.MerkleRootMultiSig)
	require.Equal(t, counterChain.ValSet.Threshold, uint8(merkleRootMultiSig.Threshold))
	for i, val := range counterChain.ValSet.Vals {
		require.Equal(t, val.Addr, merkleRootMultiSig.ValidatorPubKeys[i])
	}

	// Create message
	sender := "0xbcb815f38D481a5EBA4D7ac4c9E74D9D0FC2A7e7"
	destDomain := uint32(12345)
	message, proof := counterChain.CreateMessage(sender, destDomain, contract, "Hello!")
	// Create metadata
	metadata := counterChain.CreateMetadata(message, proof)
	// Process message
	helpers.CallProcessMsg(t, ctx, simd, user.KeyName(), hexutil.Encode(metadata), hexutil.Encode(message))

	message, proof = counterChain.CreateMessage(sender, destDomain, contract, "Hello!2")
	// Create metadata
	metadata = counterChain.CreateMetadata(message, proof)
	// Process message
	helpers.CallProcessMsg(t, ctx, simd, user.KeyName(), hexutil.Encode(metadata), hexutil.Encode(message))

	message, proof = counterChain.CreateMessage(sender, destDomain, contract, "Hello!3")
	// Create metadata
	metadata = counterChain.CreateMetadata(message, proof)
	// Process message
	helpers.CallProcessMsg(t, ctx, simd, user.KeyName(), hexutil.Encode(metadata), hexutil.Encode(message))

	dispatchMsgStruct := helpers.ExecuteMsg{
		DispatchMsg: &helpers.DispatchMsg{
			DestinationAddr: 1,
			RecipientAddr:   "0xbcb815f38D481a5EBA4D7ac4c9E74D9D0FC2A7e7",
			MessageBody:     hexutil.Encode([]byte("MsgDispatchedByContract")),
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
