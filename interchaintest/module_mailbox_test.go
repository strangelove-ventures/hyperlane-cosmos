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

	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/counterchain"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/helpers"
	ismtypes "github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types/legacy_multisig"
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

	// Create counter chain 1, with origin 1, with val set signing legacy multisig
	counterChain1 := counterchain.CreateCounterChain(t, 1, counterchain.LEGACY_MULTISIG)
	counterChain2 := counterchain.CreateCounterChain(t, 2, counterchain.MESSAGE_ID_MULTISIG)
	counterChain3 := counterchain.CreateCounterChain(t, 3, counterchain.MERKLE_ROOT_MULTISIG)

	// Set default isms for counter chains
	helpers.SetDefaultIsm(t, ctx, simd, user.KeyName(), counterChain1, counterChain2, counterChain3)
	res := helpers.QueryAllDefaultIsms(t, ctx, simd)

	var abstractIsm ismtypes.AbstractIsm
	err := simd.Config().EncodingConfig.InterfaceRegistry.UnpackAny(res.DefaultIsms[0].AbstractIsm, &abstractIsm)
	require.NoError(t, err)
	legacyMultiSig := abstractIsm.(*legacy_multisig.LegacyMultiSig)
	require.Equal(t, counterChain1.ValSet.Threshold, uint8(legacyMultiSig.Threshold))
	for i, val := range counterChain1.ValSet.Vals {
		require.Equal(t, val.Addr, legacyMultiSig.ValidatorPubKeys[i])
	}

	// Create first legacy multisig message from counter chain 1
	sender := "0xbcb815f38D481a5EBA4D7ac4c9E74D9D0FC2A7e7"
	destDomain := uint32(12345)
	message, proof := counterChain1.CreateMessage(sender, destDomain, contract, "Legacy Multisig 1")
	metadata := counterChain1.CreateLegacyMetadata(message, proof)
	helpers.CallProcessMsg(t, ctx, simd, user.KeyName(), hexutil.Encode(metadata), hexutil.Encode(message))

	// Create second legacy multisig message from counter chain 1
	message, proof = counterChain1.CreateMessage(sender, destDomain, contract, "Legacy Multisig 2")
	metadata = counterChain1.CreateLegacyMetadata(message, proof)
	helpers.CallProcessMsg(t, ctx, simd, user.KeyName(), hexutil.Encode(metadata), hexutil.Encode(message))

	// Create third legacy multisig message from counter chain 1
	message, proof = counterChain1.CreateMessage(sender, destDomain, contract, "Legacy Multisig 3")
	metadata = counterChain1.CreateLegacyMetadata(message, proof)
	helpers.CallProcessMsg(t, ctx, simd, user.KeyName(), hexutil.Encode(metadata), hexutil.Encode(message))

	// Create first message id multisig message from counter chain 2
	message, _ = counterChain2.CreateMessage(sender, destDomain, contract, "Message Id Multisig 1")
	metadata = counterChain2.CreateMessageIdMetadata(message)
	helpers.CallProcessMsg(t, ctx, simd, user.KeyName(), hexutil.Encode(metadata), hexutil.Encode(message))

	// Create first merkle root multisig message from counter chain 3
	message, proof = counterChain3.CreateMessage(sender, destDomain, contract, "Merkle Root Multisig 1")
	metadata = counterChain3.CreateMerkleRootMetadata(message, proof)
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
