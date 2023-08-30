package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	retry "github.com/avast/retry-go/v4"
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authTx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/counterchain"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/docker"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/helpers"
	ismtypes "github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"
	mbtypes "github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types/legacy_multisig"
	icv7 "github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"
	"github.com/stretchr/testify/require"
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

	users := icv7.GetAndFundTestUsers(t, ctx, "default", int64(10_000_000_000), simd)
	user := users[0]

	msg := fmt.Sprintf(`{}`)
	_, contract := helpers.SetupContract(t, ctx, simd, user.KeyName(), "../contracts/hyperlane.wasm", msg)
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

func TestHyperlaneIgp(t *testing.T) {
	t.Parallel()
	// _, filename, _, _ := runtime.Caller(0)
	// path := filepath.Dir(filename)
	// tarFilePath := filepath.Join(path, "../../")
	//goModPath := filepath.Join(path, "../../go.mod")

	// Builds the hyperlane image from the current project (e.g. locally).
	// The directory at 'tarFilePath' will be tarballed and used for the docker context.
	// The args 'buildDir' and 'dockerfilePath' are relative to 'tarFilePath' (the context).
	// Build arguments are derived from 'goModPath' so it must be a full path (not relative).
	//docker.BuildHeighlinerHyperlaneImage(docker.HyperlaneImageName, tarFilePath, ".", goModPath, "local.Dockerfile")

	DockerImage := ibc.DockerImage{
		Repository: docker.HyperlaneImageName,
		Version:    "local",
		UidGid:     "1025:1025",
	}

	// Base setup
	chains := CreateDoubleHyperlaneSimd(t, DockerImage)
	ctx := BuildInitialChain(t, chains)

	// Chains
	simd := chains[0].(*cosmos.CosmosChain)
	simd2 := chains[1].(*cosmos.CosmosChain)

	t.Log("simd.GetHostRPCAddress()", simd.GetHostRPCAddress())
	t.Log("simd2.GetHostRPCAddress()", simd2.GetHostRPCAddress())

	usersSimd := icv7.GetAndFundTestUsers(t, ctx, "default", int64(10_000_000_000), simd)
	userSimd := usersSimd[0]
	usersSimd2 := icv7.GetAndFundTestUsers(t, ctx, "default", int64(10_000_000_000), simd2)
	userSimd2 := usersSimd2[0]

	msg := fmt.Sprintf(`{}`)
	_, contract := helpers.SetupContract(t, ctx, simd, userSimd.KeyName(), "../contracts/hyperlane.wasm", msg)
	t.Log("coreContract", contract)
	_, contract2 := helpers.SetupContract(t, ctx, simd2, userSimd2.KeyName(), "../contracts/hyperlane.wasm", msg)
	t.Log("coreContract", contract2)

	verifyContractEntryPoints(t, ctx, simd, userSimd, contract)
	verifyContractEntryPoints(t, ctx, simd2, userSimd2, contract2)

	// Create counter chain 1, with origin 1, with val set signing legacy multisig
	counterChainSimd1 := counterchain.CreateCounterChain(t, 1, counterchain.LEGACY_MULTISIG)
	// Create counter chain 2, with origin 2, with val set signing legacy multisig
	counterChainSimd2 := counterchain.CreateCounterChain(t, 2, counterchain.LEGACY_MULTISIG)

	// Set default isms for counter chains for SIMD
	helpers.SetDefaultIsm(t, ctx, simd, userSimd.KeyName(), counterChainSimd1)
	res := helpers.QueryAllDefaultIsms(t, ctx, simd)

	var abstractIsm ismtypes.AbstractIsm
	err := simd.Config().EncodingConfig.InterfaceRegistry.UnpackAny(res.DefaultIsms[0].AbstractIsm, &abstractIsm)
	require.NoError(t, err)
	legacyMultiSig := abstractIsm.(*legacy_multisig.LegacyMultiSig)
	require.Equal(t, counterChainSimd1.ValSet.Threshold, uint8(legacyMultiSig.Threshold))
	for i, val := range counterChainSimd1.ValSet.Vals {
		require.Equal(t, val.Addr, legacyMultiSig.ValidatorPubKeys[i])
	}

	// Set default isms for counter chains for SIMD2
	helpers.SetDefaultIsm(t, ctx, simd2, userSimd2.KeyName(), counterChainSimd2)
	res2 := helpers.QueryAllDefaultIsms(t, ctx, simd2)

	var abstractIsm2 ismtypes.AbstractIsm
	err2 := simd2.Config().EncodingConfig.InterfaceRegistry.UnpackAny(res2.DefaultIsms[0].AbstractIsm, &abstractIsm2)
	require.NoError(t, err2)
	legacyMultiSig2 := abstractIsm2.(*legacy_multisig.LegacyMultiSig)
	require.Equal(t, counterChainSimd2.ValSet.Threshold, uint8(legacyMultiSig.Threshold))
	for i, val := range counterChainSimd2.ValSet.Vals {
		require.Equal(t, val.Addr, legacyMultiSig2.ValidatorPubKeys[i])
	}

	dMsg := []byte("HelloHyperlaneWorld")
	dispatchedMsg := hexutil.Encode(dMsg)
	destDomain := 1
	//Now setup and verification is finished for both chains. Dispatch a message
	dispatchMsgStruct := helpers.ExecuteMsg{
		DispatchMsg: &helpers.DispatchMsg{
			DestinationAddr: uint32(destDomain),
			RecipientAddr:   "0xbcb815f38D481a5EBA4D7ac4c9E74D9D0FC2A7e7",
			MessageBody:     dispatchedMsg,
		},
	}
	dipatchMsg, err := json.Marshal(dispatchMsgStruct)
	require.NoError(t, err)
	dispatchedTxHash, err := simd.ExecuteContract(ctx, userSimd.KeyName(), contract, string(dipatchMsg))
	require.NoError(t, err)

	// Look up the dispatched TX by hash. Note that this means the TX exists in a block on chain,
	// and thus can also be searched by any other available RPC method (websocket event subscription, etc).
	dispatchedTx, err := getTransaction(simd, dispatchedTxHash)
	require.NoError(t, err)
	require.NotNil(t, dispatchedTx)

	// TODO: why does this not work?? Note that getTransaction() as seen above does the same thing, but this would be cleaner.

	// simdGrpcConn, err := grpc.Dial(
	// 	simd.GetGRPCAddress(), // your gRPC server address.
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// 	grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	// )
	// require.NoError(t, err)
	// defer simdGrpcConn.Close()
	// querySimdTxsClient := txTypes.NewServiceClient(simdGrpcConn)
	// ctx, cancel := GetQueryContext()
	// defer cancel()
	// dispatchedTx, err := querySimdTxsClient.GetTx(ctx, &txTypes.GetTxRequest{Hash: dispatchedTxHash})
	// require.NoError(t, err)
	// require.NotNil(t, dispatchedTx)

	// Check that the queried TX has the events we expect.
	msgId := ""   //hyperlane ID of the message
	msgBody := "" //message to be processed
	allEvents := dispatchedTx.Events
	for _, evt := range allEvents {
		// this tx only has a single message broadcast so it will be the first 'dispatch_id' event found.
		if evt.Type == mbtypes.EventTypeDispatchId {
			for _, attr := range evt.Attributes {
				if attr.Key == mbtypes.AttributeKeyID {
					msgId = attr.Value
				}
			}
		}
		if evt.Type == mbtypes.EventTypeDispatch {
			for _, attr := range evt.Attributes {
				if attr.Key == mbtypes.AttributeKeyMessage {
					msgBody = attr.Value
				}
			}
		}
	}

	require.NotEqual(t, msgId, "")
	destDomainStr := "1"
	igpId := "1"
	destGasStr := "1000000000"
	maxPayment := "1000000000ustake"
	beneficiary := "cosmos12aqqagjkk3y7mtgkgy5fuun3j84zr3c6e0zr6n"

	// Pay for gas
	helpers.CallCreateIgp(t, ctx, simd, userSimd.KeyName(), beneficiary)
	helpers.CallPayForGasMsg(t, ctx, simd, userSimd.KeyName(), msgId, destDomainStr, destGasStr, igpId, maxPayment)

	// sign the message, then send the message to the destination chain
	// Create first legacy multisig message from counter chain 1
	sender := "0xbcb815f38D481a5EBA4D7ac4c9E74D9D0FC2A7e7"
	message, proof := counterChainSimd2.CreateMessage(sender, uint32(destDomain), contract, msgBody)
	metadata := counterChainSimd2.CreateLegacyMetadata(message, proof)

	//CallProcessMsg sends the message and verifies the message and metadata
	helpers.CallProcessMsg(t, ctx, simd2, userSimd2.KeyName(), hexutil.Encode(metadata), hexutil.Encode(message))

	err = testutil.WaitForBlocks(ctx, 2, simd2)
	require.NoError(t, err)
}

// GetQueryContext returns a context that includes the height and uses the timeout from the config
func GetQueryContext() (context.Context, context.CancelFunc) {
	timeout, _ := time.ParseDuration("15s")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	//ctx = metadata.AppendToOutgoingContext(ctx, grpctypes.GRPCBlockHeightHeader, height)
	return ctx, cancel
}

func getTransaction(c *cosmos.CosmosChain, txHash string) (*sdk.TxResponse, error) {
	// Retry because sometimes the tx is not committed to state yet.
	var txResp *types.TxResponse
	err := retry.Do(func() error {
		var err error
		txResp, err = authTx.QueryTx(getFullNode(c).CliContext(), txHash)
		return err
	},
		// retry for total of 3 seconds
		retry.Attempts(15),
		retry.Delay(200*time.Millisecond),
		retry.DelayType(retry.FixedDelay),
		retry.LastErrorOnly(true),
	)
	return txResp, err
}

func getFullNode(c *cosmos.CosmosChain) *cosmos.ChainNode {
	if len(c.FullNodes) > 0 {
		// use first full node
		return c.FullNodes[0]
	}
	// use first validator
	return c.Validators[0]
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
