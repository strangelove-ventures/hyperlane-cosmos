package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/counterchain"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/docker"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/helpers"
	ismtypes "github.com/strangelove-ventures/hyperlane-cosmos/x/ism/types"

	icv7 "github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"
	"github.com/stretchr/testify/require"

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

	// TODO: this is a bug. The ordering of DefaultIsms isn't guaranteed so you CANNOT assume [0] is a 'LegacyMultiSig'
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
	message, proof := counterChain1.CreateMessage(sender, destDomain, destDomain, contract, "Legacy Multisig 1")
	metadata := counterChain1.CreateLegacyMetadata(message, proof)
	helpers.CallProcessMsg(t, ctx, simd, user.KeyName(), hexutil.Encode(metadata), hexutil.Encode(message))

	// Create second legacy multisig message from counter chain 1
	message, proof = counterChain1.CreateMessage(sender, destDomain, destDomain, contract, "Legacy Multisig 2")
	metadata = counterChain1.CreateLegacyMetadata(message, proof)
	helpers.CallProcessMsg(t, ctx, simd, user.KeyName(), hexutil.Encode(metadata), hexutil.Encode(message))

	// Create third legacy multisig message from counter chain 1
	message, proof = counterChain1.CreateMessage(sender, destDomain, destDomain, contract, "Legacy Multisig 3")
	metadata = counterChain1.CreateLegacyMetadata(message, proof)
	helpers.CallProcessMsg(t, ctx, simd, user.KeyName(), hexutil.Encode(metadata), hexutil.Encode(message))

	// Create first message id multisig message from counter chain 2
	message, _ = counterChain2.CreateMessage(sender, destDomain, destDomain, contract, "Message Id Multisig 1")
	metadata = counterChain2.CreateMessageIdMetadata(message)
	helpers.CallProcessMsg(t, ctx, simd, user.KeyName(), hexutil.Encode(metadata), hexutil.Encode(message))

	// Create first merkle root multisig message from counter chain 3
	message, proof = counterChain3.CreateMessage(sender, destDomain, destDomain, contract, "Merkle Root Multisig 1")
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

	buildsEnabled := false
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)
	tarFilePath := filepath.Join(path, "../../")
	goModPath := filepath.Join(path, "../../go.mod")

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

	// Base setup
	chains := CreateDoubleHyperlaneSimd(t, DockerImage, 23456, 34567)
	ctx := BuildInitialChain(t, chains)

	// Chains
	simd := chains[0].(*cosmos.CosmosChain)
	simd2 := chains[1].(*cosmos.CosmosChain)

	simdDomainOutput := helpers.QueryDomain(t, ctx, simd)
	simd2DomainOutput := helpers.QueryDomain(t, ctx, simd2)
	simdDomainStr := helpers.ParseQueryDomain(string(simdDomainOutput))
	simd2DomainStr := helpers.ParseQueryDomain(string(simd2DomainOutput))
	simdDomain, err := strconv.ParseUint(simdDomainStr, 10, 64)
	require.NoError(t, err)
	simd2Domain, err := strconv.ParseUint(simd2DomainStr, 10, 64)
	require.NoError(t, err)
	fmt.Printf("simd mailbox domain: %d, simd2 mailbox domain: %d\n", simdDomain, simd2Domain)

	t.Log("simd.GetHostRPCAddress()", simd.GetHostRPCAddress())
	t.Log("simd2.GetHostRPCAddress()", simd2.GetHostRPCAddress())

	usersSimd := icv7.GetAndFundTestUsers(t, ctx, "default", int64(10_000_000_000), simd)
	userSimd := usersSimd[0]
	usersSimdChain1Users2 := icv7.GetAndFundTestUsers(t, ctx, "default", int64(10_000_000_000), simd)
	chain1Oracle := usersSimdChain1Users2[0]
	usersSimd2 := icv7.GetAndFundTestUsers(t, ctx, "default", int64(10_000_000_000), simd2)
	userSimd2 := usersSimd2[0]

	msg := `{}`
	_, contract := helpers.SetupContract(t, ctx, simd, userSimd.KeyName(), "../contracts/hyperlane.wasm", msg)
	t.Log("coreContract", contract)
	_, contract2 := helpers.SetupContract(t, ctx, simd2, userSimd2.KeyName(), "../contracts/hyperlane.wasm", msg)
	t.Log("coreContract", contract2)

	verifyContractEntryPoints(t, ctx, simd, userSimd, contract)
	verifyContractEntryPoints(t, ctx, simd2, userSimd2, contract2)

	// Create counter chain 1 with val set signing legacy multisig
	simdIsmValidator := counterchain.CreateCounterChain(t, uint32(simdDomain), counterchain.LEGACY_MULTISIG)
	// Create counter chain 2 with val set signing legacy multisig
	simd2IsmValidator := counterchain.CreateCounterChain(t, uint32(simd2Domain), counterchain.LEGACY_MULTISIG)

	// Set default isms for counter chains for SIMD
	helpers.SetDefaultIsm(t, ctx, simd, userSimd.KeyName(), simd2IsmValidator)
	res := helpers.QueryAllDefaultIsms(t, ctx, simd)

	var abstractIsm ismtypes.AbstractIsm
	err = simd.Config().EncodingConfig.InterfaceRegistry.UnpackAny(res.DefaultIsms[0].AbstractIsm, &abstractIsm)
	require.NoError(t, err)
	legacyMultiSig := abstractIsm.(*legacy_multisig.LegacyMultiSig)
	require.Equal(t, simd2IsmValidator.ValSet.Threshold, uint8(legacyMultiSig.Threshold))
	for i, val := range simd2IsmValidator.ValSet.Vals {
		require.Equal(t, val.Addr, legacyMultiSig.ValidatorPubKeys[i])
	}

	// Set default isms for counter chains for SIMD2
	helpers.SetDefaultIsm(t, ctx, simd2, userSimd2.KeyName(), simdIsmValidator)
	res2 := helpers.QueryAllDefaultIsms(t, ctx, simd2)

	var abstractIsm2 ismtypes.AbstractIsm
	err2 := simd2.Config().EncodingConfig.InterfaceRegistry.UnpackAny(res2.DefaultIsms[0].AbstractIsm, &abstractIsm2)
	require.NoError(t, err2)
	legacyMultiSig2 := abstractIsm2.(*legacy_multisig.LegacyMultiSig)
	require.Equal(t, simdIsmValidator.ValSet.Threshold, uint8(legacyMultiSig2.Threshold))
	for i, val := range simdIsmValidator.ValSet.Vals {
		require.Equal(t, val.Addr, legacyMultiSig2.ValidatorPubKeys[i])
	}

	// recipientCosmosBech32 := "cosmos12aqqagjkk3y7mtgkgy5fuun3j84zr3c6e0zr6n"
	recipientAccAddr := sdk.MustAccAddressFromBech32(contract2).Bytes()
	recipientDispatch := hexutil.Encode([]byte(recipientAccAddr))
	fmt.Printf("Recipient dispatch addr hex: %s", recipientDispatch)

	dMsg := []byte("HelloHyperlaneWorld")
	dispatchedMsg := hexutil.Encode(dMsg)

	// Now setup and verification is finished for both chains. Dispatch a message
	dispatchMsgStruct := helpers.ExecuteMsg{
		DispatchMsg: &helpers.DispatchMsg{
			DestinationAddr: uint32(simd2Domain),
			RecipientAddr:   recipientDispatch,
			MessageBody:     dispatchedMsg,
		},
	}
	dipatchMsg, err := json.Marshal(dispatchMsgStruct)
	require.NoError(t, err)
	dispatchedTxHash, err := simd.ExecuteContract(ctx, userSimd.KeyName(), contract, string(dipatchMsg))
	require.NoError(t, err)
	dispatchedDestDomain, dispatchedRecipientAddrHex, dispatchedMsgBody, dispatchedMsgId, dispatchSender, hyperlaneMsg, err := helpers.VerifyDispatchEvents(simd, dispatchedTxHash)
	require.NoError(t, err)
	require.NotEmpty(t, dispatchSender)
	require.NotEmpty(t, dispatchedRecipientAddrHex)

	// Look up the dispatched TX by hash. Note that this means the TX exists in a block on chain,
	// and thus can also be searched by any other available RPC method (websocket event subscription, etc).
	dispatchedTx, err := helpers.GetTransaction(simd, dispatchedTxHash)
	require.NoError(t, err)
	require.NotNil(t, dispatchedTx)

	// The formula for native tokens owed, where 'exchRateScale' is 10^10, is:
	// exchRate := oracle.TokenExchangeRate
	// gasPrice := oracle.GasPrice
	// destGasCost := destGasAmount.Mul(gasPrice)
	// nativePrice := destGasCost.Mul(exchRate).Quo(exchRateScale)

	// Below, exchange rate equals the scale factor, making the formula:
	// native tokens owed = destGasAmount.Mul(gasPrice) = 300,000
	exchangeRate := math.NewInt(1e10)
	gasPrice := math.NewInt(1)
	testGasAmount := math.NewInt(300000)
	quoteExpected := math.NewInt(300000)

	require.NotEqual(t, dispatchedMsgId, "")
	beneficiary := "cosmos12aqqagjkk3y7mtgkgy5fuun3j84zr3c6e0zr6n"

	// This should be IGP 0, which we will ignore and not use for anything
	out := helpers.CallCreateIgp(t, ctx, simd, userSimd.KeyName(), beneficiary)
	igpTxHash := helpers.ParseTxHash(string(out))
	igpIdUint, err := helpers.VerifyIgpEvents(simd, igpTxHash)
	require.NoError(t, err)
	require.Equal(t, uint32(0), igpIdUint)

	// This should be IGP 1
	createigpout := helpers.CallCreateIgp(t, ctx, simd, userSimd.KeyName(), beneficiary)
	igpTxHash = helpers.ParseTxHash(string(createigpout))
	igpIdUint, err = helpers.VerifyIgpEvents(simd, igpTxHash)
	require.NoError(t, err)
	require.Equal(t, uint32(1), igpIdUint)
	igpId := strconv.FormatUint(uint64(igpIdUint), 10)

	destDomainStr := strconv.FormatUint(uint64(simd2Domain), 10)
	// Create the oracle and verify we get the expected address in the events
	oracle := chain1Oracle.FormattedAddress()
	createOracleOutput := helpers.CallCreateOracle(t, ctx, simd, userSimd.KeyName(), oracle, igpId, destDomainStr)
	createOracleTxHash := helpers.ParseTxHash(string(createOracleOutput))
	oracleEvtAddr, err := helpers.VerifyCreateOracleEvents(simd, createOracleTxHash)
	require.NoError(t, err)
	require.Equal(t, oracle, oracleEvtAddr)

	// Make sure this fails, since only the oracle can set the gas prices (e.g. use wrong address intentionally)
	setGasOutput := helpers.CallSetGasPriceMsg(t, ctx, simd, userSimd.KeyName(), igpId, destDomainStr, gasPrice.String(), exchangeRate.String())
	setGasTxHash1 := helpers.ParseTxHash(string(setGasOutput))
	_, _, _, err = helpers.VerifySetGasPriceEvents(simd, setGasTxHash1)
	require.Error(t, err)

	// This should succeed, and we verify the events contain the expected domain/exchange rate/gas price.
	setGasOutput = helpers.CallSetGasPriceMsg(t, ctx, simd, chain1Oracle.KeyName(), igpId, destDomainStr, gasPrice.String(), exchangeRate.String())
	setGasTxHash2 := helpers.ParseTxHash(string(setGasOutput))
	setGasDomain, setGasRate, setGasPrice, err := helpers.VerifySetGasPriceEvents(simd, setGasTxHash2)
	require.NoError(t, err)
	require.Equal(t, setGasDomain, destDomainStr)
	require.Equal(t, setGasRate, exchangeRate.String())
	require.Equal(t, setGasPrice, gasPrice.String())

	// Look up the expected payment and verify it matches what we expected (according to the gas price, exchange rate, and scale).
	quoteGasPaymentOutput := helpers.QueryQuoteGasPayment(t, ctx, simd, userSimd.KeyName(), igpId, destDomainStr, testGasAmount.String())
	nativeTokensOwed, denom := helpers.ParseQuoteGasPayment(string(quoteGasPaymentOutput))
	amountActual, _ := big.NewInt(0).SetString(nativeTokensOwed, 10)
	require.Equal(t, amountActual.String(), quoteExpected.String())
	require.Equal(t, denom, "stake")

	maxPayment := sdk.NewCoin(denom, math.NewIntFromBigInt(amountActual))
	payForGasOutput := helpers.CallPayForGasMsg(t, ctx, simd, userSimd.KeyName(), dispatchedMsgId, destDomainStr, testGasAmount.String(), igpId, maxPayment.String())
	payForGasTxHash := helpers.ParseTxHash(string(payForGasOutput))
	paidMsgId, nativePayment, gasAmount, err := helpers.VerifyPayForGasEvents(simd, payForGasTxHash)
	require.NoError(t, err)
	require.Equal(t, dispatchedMsgId, paidMsgId)
	require.Equal(t, nativePayment, nativeTokensOwed+denom)
	require.Equal(t, gasAmount, testGasAmount.String())

	// sign the message, then send the message to the destination chain
	// Create first legacy multisig message from counter chain 1
	// sender := "0xbcb815f38D481a5EBA4D7ac4c9E74D9D0FC2A7e7"
	dispatchedDestDomainUint, err := strconv.ParseUint(dispatchedDestDomain, 10, 64)
	require.NoError(t, err)
	require.Greater(t, dispatchedDestDomainUint, uint64(0))

	fmt.Printf("Emitted event recipient address hex: %s", dispatchedRecipientAddrHex)
	dispatchedRecipientAddr := hexutil.MustDecode(dispatchedRecipientAddrHex)
	bech32Recipient := sdk.MustBech32ifyAddressBytes(sdk.GetConfig().GetBech32AccountAddrPrefix(), dispatchedRecipientAddr)
	fmt.Printf("Emitted event recipient as bech32: %s. Contract2 which should match: %s", bech32Recipient, contract2)

	b, err := hexutil.Decode(dispatchedMsgBody)
	require.NoError(t, err)
	message, proof := simdIsmValidator.CreateMessage(dispatchSender, uint32(simdDomain), uint32(simd2Domain), bech32Recipient, string(b))
	metadata := simdIsmValidator.CreateLegacyMetadata(message, proof)

	// treeMdOut := helpers.QueryCurrentTreeMetadata(t, ctx, simd)
	// root, count := helpers.ParseQueryTreeMetadata(string(treeMdOut))
	// fmt.Printf("root: %s, count: %s\n", root, count)

	hyperlaneMsgDispatched, err := hexutil.Decode(hyperlaneMsg)
	require.NoError(t, err)

	match := compareBytes(hyperlaneMsgDispatched, message)
	require.True(t, match)

	// CallProcessMsg sends the message and verifies the message and metadata
	processStdout := helpers.CallProcessMsg(t, ctx, simd2, userSimd2.KeyName(), hexutil.Encode(metadata), hexutil.Encode(message))
	processTxHash := helpers.ParseTxHash(string(processStdout))
	processMsgId, err := helpers.VerifyProcessEvents(simd2, processTxHash)
	require.NoError(t, err)
	require.Equal(t, dispatchedMsgId, processMsgId)

	err = testutil.WaitForBlocks(ctx, 2, simd2)
	require.NoError(t, err)
}

func compareBytes(b1 []byte, b2 []byte) bool {
	lenMatch := len(b1) == len(b2)
	if !lenMatch {
		fmt.Printf("byte arrays different length")
	}

	for i, b1Byte := range b1 {
		if len(b2) > i {
			b2Byte := b2[i]
			if b2Byte != b1Byte {
				fmt.Printf("Byte mismatch at index %d. b1: %d, b2: %d\n", i, b1Byte, b2Byte)
				return false
			}
		}
	}

	return lenMatch
}

// GetQueryContext returns a context that includes the height and uses the timeout from the config
func GetQueryContext() (context.Context, context.CancelFunc) {
	timeout, _ := time.ParseDuration("15s")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	// ctx = metadata.AppendToOutgoingContext(ctx, grpctypes.GRPCBlockHeightHeader, height)
	return ctx, cancel
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
