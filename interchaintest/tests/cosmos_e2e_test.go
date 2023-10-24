package tests

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
	
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"

	icv7 "github.com/strangelove-ventures/interchaintest/v7"
	interchaintest "github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	hyperlane "github.com/strangelove-ventures/interchaintest/v7/chain/hyperlane"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/counterchain"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/helpers"

	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/docker"

	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/testreporter"
	"github.com/stretchr/testify/require"
)

const (
	COSMOS_E2E_CONFIG = "hyperlane.yaml"
	mnemonic          = "spare number knock scan copper method lunch camera trap inject fine suspect edit sure design crowd sorry actual better spatial cover grit entire raccoon" // Testing only, do NOT use this mnemonic
	bech32Addr        = "cosmos13gpsgkxaavz3kcvh8y55xzat9umg944qnwxq4k"                                                                                                            // for the mnemonic above
	mnemonicPrivKey   = "fe257759a16d7085ba4df68773c94647966e9fc8c7a7e1eb3311c40bbe1a0ed3"                                                                                         // for the mnemonic above                                                                                                      // Corresponds to the key above
	// valPrivKey        = "8166f546bab6da521a8369cab06c5d2b9e46670292d85c875ee9ec20e84ffb61"                                                                                         // Testing only, do NOT use this key. Corresponds to the hyperlane validator signing key, not the mnemonic above
)

var (
	SupportedAlgorithms       = keyring.SigningAlgoList{hd.Secp256k1}
	SupportedAlgorithmsLedger = keyring.SigningAlgoList{hd.Secp256k1}
)

func KeyringAlgoOptions() keyring.Option {
	return func(options *keyring.Options) {
		options.SupportedAlgos = SupportedAlgorithms
		options.SupportedAlgosLedger = SupportedAlgorithmsLedger
	}
}

func TestHyperlaneCosmosE2E(t *testing.T) {
	logger := NewLogger(t)

	// Mailbox address - will be used in the hyperlane validator config
	mailboxHex, expectedMailbox := helpers.GetMailboxAddress()
	prefixedMailboxHex := "0x" + mailboxHex

	// Directories where files related to this test will be stored
	val1TmpDir := t.TempDir()
	val2TmpDir := t.TempDir()
	rlyTmpDir := t.TempDir()

	// Get the hyperlane agent raw configs (before variable replacements)
	valSimd1, valSimd2, rly := readHyperlaneConfig(t, COSMOS_E2E_CONFIG, logger)

	// Get the validator key for the agents. We also need this key to configure the chain ISM.
	valSimd1PrivKey, err := getHyperlaneBaseValidatorKey(valSimd1)
	require.NoError(t, err)
	valSimd2PrivKey, err := getHyperlaneBaseValidatorKey(valSimd2)
	require.NoError(t, err)

	// Build the chain docker image from the local repo
	optionalBuildChainImage()

	DockerImage := ibc.DockerImage{
		Repository: docker.HyperlaneImageName,
		Version:    "local",
		UidGid:     "1025:1025",
	}

	simd1Domain := uint32(23456)
	simd2Domain := uint32(34567)

	// Set up two Cosmos chains (with our hyperlane modules) for the test.
	// The images must be in our local registry so we skip image pull.
	chains := CreateHyperlaneSimds(t, DockerImage, []uint32{simd1Domain, simd2Domain})
	simd1 := chains[0].(*cosmos.CosmosChain)
	simd1.SkipImagePull = true
	simd2 := chains[1].(*cosmos.CosmosChain)
	simd2.SkipImagePull = true

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
		TestName:          t.Name(),
		Client:            client,
		NetworkID:         network,
		SkipPathCreation:  true,
		BlockDatabaseFile: interchaintest.DefaultBlockDatabaseFilepath(),
	}

	err = ic.Build(ctx, eRep, opts)
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = ic.Close()
	})

	// Query the cosmos chains and ensure the hyperlane domain configuration is as expected
	verifyDomain(t, ctx, simd1, uint64(simd1Domain))
	verifyDomain(t, ctx, simd2, uint64(simd2Domain))

	// Now we need to configure the Hyperlane modules and setup some test users...
	userSimd1 := icv7.GetAndFundTestUsers(t, ctx, "default", int64(10_000_000_000), simd1)[0]
	userSimd2 := icv7.GetAndFundTestUsers(t, ctx, "default", int64(10_000_000_000), simd2)[0]
	simd1Oracle := icv7.GetAndFundTestUsers(t, ctx, "default", int64(10_000_000_000), simd1)[0]

	// Fund this wallet since the validator will announce its storage location with this mnemonic's private key.
	_, err = icv7.GetAndFundTestUserWithMnemonic(ctx, "valannounce", mnemonic, int64(10_000_000_000), simd1)
	require.NoError(t, err)
	_, err = icv7.GetAndFundTestUserWithMnemonic(ctx, "valannounce", mnemonic, int64(10_000_000_000), simd2)
	require.NoError(t, err)

	msg := `{}`
	_, contract := helpers.SetupContract(t, ctx, simd1, userSimd1.KeyName(), "../contracts/hyperlane.wasm", msg)
	logger.Info("simd1 contract", zap.String("address", contract))
	_, contract2 := helpers.SetupContract(t, ctx, simd2, userSimd2.KeyName(), "../contracts/hyperlane.wasm", msg)
	logger.Info("simd2 contract", zap.String("address", contract2))

	verifyContractEntryPoints(t, ctx, simd1, userSimd1, contract)
	verifyContractEntryPoints(t, ctx, simd2, userSimd2, contract2)

	// Create counter chain 1 with val set signing legacy multisig.
	// The private key used here MUST be the one from the validator config file.
	simd1IsmValidator := counterchain.CreateEmperorValidator(t, simd1Domain, counterchain.LEGACY_MULTISIG, valSimd1PrivKey)
	// Create counter chain 2 with val set signing legacy multisig
	simd2IsmValidator := counterchain.CreateEmperorValidator(t, simd2Domain, counterchain.LEGACY_MULTISIG, valSimd2PrivKey)

	// Set default isms for counter chains for SIMD
	helpers.SetDefaultIsm(t, ctx, simd1, userSimd1.KeyName(), simd2IsmValidator)
	// Set default isms for counter chains for SIMD2
	helpers.SetDefaultIsm(t, ctx, simd2, userSimd2.KeyName(), simd1IsmValidator)

	recipientAccAddr := sdk.MustAccAddressFromBech32(contract2).Bytes()
	recipientDispatch := hexutil.Encode([]byte(recipientAccAddr))
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

	valJson, err = preconfigureHyperlaneValidator(t, valSimd2, val2TmpDir, mnemonicPrivKey, chains[1].Config().ChainID, chains[1].Config().Name, chains[1].GetRPCAddress(), "http://"+chains[1].GetGRPCAddress(), prefixedMailboxHex, 34567)
	require.NoError(t, err)
	simd1ValidatorSignaturesDir := filepath.Join(val1TmpDir, "signatures-"+chains[0].Config().Name) //${val_dir}/signatures-${chainName}
	simd2ValidatorSignaturesDir := filepath.Join(val2TmpDir, "signatures-"+chains[1].Config().Name) //${val_dir}/signatures-${chainName}

	rlyCfgs := []chainCfg{
		{
			privKey:                mnemonicPrivKey,
			chainID:                chains[0].Config().ChainID,
			chainName:              chains[0].Config().Name,
			rpcUrl:                 chains[0].GetRPCAddress(),
			grpcUrl:                "http://" + chains[0].GetGRPCAddress(),
			originMailboxHex:       prefixedMailboxHex,
			domain:                 23456,
			validatorSignaturePath: simd1ValidatorSignaturesDir,
		},
		{
			privKey:                mnemonicPrivKey,
			chainID:                chains[1].Config().ChainID,
			chainName:              chains[1].Config().Name,
			rpcUrl:                 chains[1].GetRPCAddress(),
			grpcUrl:                "http://" + chains[1].GetGRPCAddress(),
			originMailboxHex:       prefixedMailboxHex,
			domain:                 34567,
			validatorSignaturePath: simd2ValidatorSignaturesDir,
		},
	}
	_, err = preconfigureHyperlaneRelayer(t, rly, rlyTmpDir, rlyCfgs)
	require.NoError(t, err)

	// Our images are currently local. You must build locally in monorepo, e.g. "cd rust && docker build .".
	// Also make sure that the tags in hyperlane.yaml match the local docker image repo and version.
	hyperlaneNetwork := hyperlane.NewHyperlaneNetwork(true, true)
	err = hyperlaneNetwork.Build(ctx, logger, eRep, opts, *valSimd1, *valSimd2, *rly)
	require.NoError(t, err)

	// Give the hyperlane validators and relayer time to start up and start watching the mailbox for the chain
	time.Sleep(10 * time.Second)

	// Dispatch a message to SIMD1
	dMsg := []byte("CosmosSimd1ToCosmosSimd2")
	dispatchedMsg := hexutil.Encode(dMsg)

	dispatchMsgStruct := helpers.ExecuteMsg{
		DispatchMsg: &helpers.DispatchMsg{
			DestinationAddr: uint32(simd2Domain),
			RecipientAddr:   recipientDispatch,
			MessageBody:     dispatchedMsg,
		},
	}

	// Dispatch the hyperlane message to simd1
	dipatchMsg, err := json.Marshal(dispatchMsgStruct)
	require.NoError(t, err)
	dispatchedTxHash, err := simd1.ExecuteContract(ctx, userSimd1.KeyName(), contract, string(dipatchMsg))
	require.NoError(t, err)
	logger.Info("Message dispatched to simd1")
	dispatchedDestDomain, dispatchedRecipientAddrHex, dispatchedMsgBody, dispatchedMsgId, dispatchSender, _, err := helpers.VerifyDispatchEvents(simd1, dispatchedTxHash)
	require.NoError(t, err)
	require.NotEmpty(t, dispatchSender)
	require.NotEmpty(t, dispatchedRecipientAddrHex)
	require.Equal(t, fmt.Sprintf("%d", simd2Domain), dispatchedDestDomain)
	// Finished sending message to simd1!

	// ***************** IGP setup, copied from  ***************************************
	exchangeRate := math.NewInt(1e10)
	gasPrice := math.NewInt(1)
	testGasAmount := math.NewInt(300000)
	quoteExpected := math.NewInt(300000)
	beneficiary := "cosmos12aqqagjkk3y7mtgkgy5fuun3j84zr3c6e0zr6n"

	// This should be IGP 0, which we will ignore and not use for anything
	out := helpers.CallCreateIgp(t, ctx, simd1, userSimd1.KeyName(), beneficiary)
	igpTxHash := helpers.ParseTxHash(string(out))
	igpIdUint, err := helpers.VerifyIgpEvents(simd1, igpTxHash)
	require.NoError(t, err)
	require.Equal(t, uint32(0), igpIdUint)

	// This should be IGP 1
	createigpout := helpers.CallCreateIgp(t, ctx, simd1, userSimd1.KeyName(), beneficiary)
	igpTxHash = helpers.ParseTxHash(string(createigpout))
	igpIdUint, err = helpers.VerifyIgpEvents(simd1, igpTxHash)
	require.NoError(t, err)
	require.Equal(t, uint32(1), igpIdUint)
	igpId := strconv.FormatUint(uint64(igpIdUint), 10)

	destDomainStr := strconv.FormatUint(uint64(simd2Domain), 10)
	createOracleOutput := helpers.CallCreateOracle(t, ctx, simd1, userSimd1.KeyName(), simd1Oracle.FormattedAddress(), igpId, destDomainStr)
	createOracleTxHash := helpers.ParseTxHash(string(createOracleOutput))
	oracleEvtAddr, err := helpers.VerifyCreateOracleEvents(simd1, createOracleTxHash)
	require.NoError(t, err)
	require.Equal(t, simd1Oracle.FormattedAddress(), oracleEvtAddr)


	// This should succeed, and we verify the events contain the expected domain/exchange rate/gas price.
	setGasOutput := helpers.CallSetGasPriceMsg(t, ctx, simd1, simd1Oracle.KeyName(), igpId, destDomainStr, gasPrice.String(), exchangeRate.String())
	setGasTxHash2 := helpers.ParseTxHash(string(setGasOutput))
	setGasDomain, setGasRate, setGasPrice, err := helpers.VerifySetGasPriceEvents(simd1, setGasTxHash2)
	require.NoError(t, err)
	require.Equal(t, setGasDomain, destDomainStr)
	require.Equal(t, setGasRate, exchangeRate.String())
	require.Equal(t, setGasPrice, gasPrice.String())

	// Look up the expected payment and verify it matches what we expected (according to the gas price, exchange rate, and scale).
	quoteGasPaymentOutput := helpers.QueryQuoteGasPayment(t, ctx, simd1, userSimd1.KeyName(), igpId, destDomainStr, testGasAmount.String())
	nativeTokensOwed, denom := helpers.ParseQuoteGasPayment(string(quoteGasPaymentOutput))
	amountActual, _ := big.NewInt(0).SetString(nativeTokensOwed, 10)
	require.Equal(t, amountActual.String(), quoteExpected.String())
	require.Equal(t, denom, "stake")

	maxPayment := sdk.NewCoin(denom, math.NewIntFromBigInt(amountActual))
	payForGasOutput := helpers.CallPayForGasMsg(t, ctx, simd1, simd1Oracle.KeyName(), dispatchedMsgId, destDomainStr, testGasAmount.String(), igpId, maxPayment.String())
	payForGasTxHash := helpers.ParseTxHash(string(payForGasOutput))
	paidMsgId, nativePayment, gasAmount, err := helpers.VerifyPayForGasEvents(simd1, payForGasTxHash)
	require.NoError(t, err)
	require.Equal(t, dispatchedMsgId, paidMsgId)
	require.Equal(t, nativePayment, nativeTokensOwed)
	require.Equal(t, gasAmount, testGasAmount.String())

	// *************** Continue with non-IGP testing ***********************

	// Wait for the hyperlane validator to sign it. The first message will show up as 0.json
	expectedSigFile := "0.json"
	simd1FirstSignedCheckpoint := filepath.Join(simd1ValidatorSignaturesDir, expectedSigFile)

	// Find the message that the hyperlane validator signed.
	// Here, we know what location the signed files should show up in based on our configuration.
	// (1) Wait for the 0_with_id.json file to show up in the validator's bind mount on the host.
	err = Await(func() (bool, error) {
		return fileExists(simd1FirstSignedCheckpoint), nil
	}, 1*time.Minute, 1*time.Second)
	require.NoError(t, err)
	valSig, err := os.ReadFile(simd1FirstSignedCheckpoint)
	require.NoError(t, err)
	signature := &ValidatorCheckpoint{}
	err = json.Unmarshal(valSig, signature)
	require.NoError(t, err)
	require.NotEmpty(t, signature.SerializedSignature)
	_, err = hexutil.Decode(signature.SerializedSignature)
	require.NoError(t, err)

	// Based on the message that was dispatched, find the message ID.
	// We will use the message ID to wait for the message to be processed on the destination.
	dispatchedRecipientAddr := hexutil.MustDecode(dispatchedRecipientAddrHex)
	bech32Recipient := sdk.MustBech32ifyAddressBytes(sdk.GetConfig().GetBech32AccountAddrPrefix(), dispatchedRecipientAddr)
	b, err := hexutil.Decode(dispatchedMsgBody)
	require.NoError(t, err)
	message, _ := simd1IsmValidator.CreateMessage(dispatchSender, simd1Domain, simd2Domain, bech32Recipient, string(b))
	messageId := simd1IsmValidator.GetMessageId(message)
	require.Equal(t, dispatchedMsgId, hexutil.Encode(messageId))

	err = Await(func() (bool, error) {
		return helpers.QueryMsgDelivered(t, ctx, simd2, dispatchedMsgId), nil
	}, 10*time.Minute, 5*time.Second)
	require.NoError(t, err)
}

func TestHyperlaneCosmosMultiMessageE2E(t *testing.T) {
	logger := NewLogger(t)

	// Mailbox address - will be used in the hyperlane validator config
	mailboxHex, expectedMailbox := helpers.GetMailboxAddress()
	prefixedMailboxHex := "0x" + mailboxHex

	// Directories where files related to this test will be stored
	val1TmpDir := t.TempDir()
	val2TmpDir := t.TempDir()
	rlyTmpDir := t.TempDir()

	// Get the hyperlane agent raw configs (before variable replacements)
	valSimd1, valSimd2, rly := readHyperlaneConfig(t, COSMOS_E2E_CONFIG, logger)

	// Get the validator key for the agents. We also need this key to configure the chain ISM.
	valSimd1PrivKey, err := getHyperlaneBaseValidatorKey(valSimd1)
	require.NoError(t, err)
	valSimd2PrivKey, err := getHyperlaneBaseValidatorKey(valSimd2)
	require.NoError(t, err)

	// Build the chain docker image from the local repo
	optionalBuildChainImage()

	DockerImage := ibc.DockerImage{
		Repository: docker.HyperlaneImageName,
		Version:    "local",
		UidGid:     "1025:1025",
	}

	simd1Domain := uint32(23456)
	simd2Domain := uint32(34567)

	// Set up two Cosmos chains (with our hyperlane modules) for the test.
	// The images must be in our local registry so we skip image pull.
	chains := CreateHyperlaneSimds(t, DockerImage, []uint32{simd1Domain, simd2Domain})
	simd1 := chains[0].(*cosmos.CosmosChain)
	simd1.SkipImagePull = true
	simd2 := chains[1].(*cosmos.CosmosChain)
	simd2.SkipImagePull = true

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
		TestName:          t.Name(),
		Client:            client,
		NetworkID:         network,
		SkipPathCreation:  true,
		BlockDatabaseFile: interchaintest.DefaultBlockDatabaseFilepath(),
	}

	err = ic.Build(ctx, eRep, opts)
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = ic.Close()
	})

	// Query the cosmos chains and ensure the hyperlane domain configuration is as expected
	verifyDomain(t, ctx, simd1, uint64(simd1Domain))
	verifyDomain(t, ctx, simd2, uint64(simd2Domain))

	// Now we need to configure the Hyperlane modules and setup some test users...
	userSimd := icv7.GetAndFundTestUsers(t, ctx, "default", int64(10_000_000_000), simd1)[0]
	userSimd2 := icv7.GetAndFundTestUsers(t, ctx, "default", int64(10_000_000_000), simd2)[0]

	// Fund this wallet since the validator will announce its storage location with this mnemonic's private key.
	_, err = icv7.GetAndFundTestUserWithMnemonic(ctx, "valannounce", mnemonic, int64(10_000_000_000), simd1)
	require.NoError(t, err)
	_, err = icv7.GetAndFundTestUserWithMnemonic(ctx, "valannounce", mnemonic, int64(10_000_000_000), simd2)
	require.NoError(t, err)

	msg := `{}`
	_, contract := helpers.SetupContract(t, ctx, simd1, userSimd.KeyName(), "../contracts/hyperlane.wasm", msg)
	logger.Info("simd1 contract", zap.String("address", contract))
	_, contract2 := helpers.SetupContract(t, ctx, simd2, userSimd2.KeyName(), "../contracts/hyperlane.wasm", msg)
	logger.Info("simd2 contract", zap.String("address", contract2))

	verifyContractEntryPoints(t, ctx, simd1, userSimd, contract)
	verifyContractEntryPoints(t, ctx, simd2, userSimd2, contract2)

	// Create counter chain 1 with val set signing legacy multisig.
	// The private key used here MUST be the one from the validator config file.
	simd1IsmValidator := counterchain.CreateEmperorValidator(t, simd1Domain, counterchain.LEGACY_MULTISIG, valSimd1PrivKey)
	// Create counter chain 2 with val set signing legacy multisig
	simd2IsmValidator := counterchain.CreateEmperorValidator(t, simd2Domain, counterchain.LEGACY_MULTISIG, valSimd2PrivKey)

	// Set default isms for counter chains for SIMD
	helpers.SetDefaultIsm(t, ctx, simd1, userSimd.KeyName(), simd2IsmValidator)
	// Set default isms for counter chains for SIMD2
	helpers.SetDefaultIsm(t, ctx, simd2, userSimd2.KeyName(), simd1IsmValidator)

	recipientAccAddr := sdk.MustAccAddressFromBech32(contract2).Bytes()
	recipientDispatch := hexutil.Encode([]byte(recipientAccAddr))
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

	valJson, err = preconfigureHyperlaneValidator(t, valSimd2, val2TmpDir, mnemonicPrivKey, chains[1].Config().ChainID, chains[1].Config().Name, chains[1].GetRPCAddress(), "http://"+chains[1].GetGRPCAddress(), prefixedMailboxHex, 34567)
	require.NoError(t, err)
	simd1ValidatorSignaturesDir := filepath.Join(val1TmpDir, "signatures-"+chains[0].Config().Name) //${val_dir}/signatures-${chainName}
	simd2ValidatorSignaturesDir := filepath.Join(val2TmpDir, "signatures-"+chains[1].Config().Name) //${val_dir}/signatures-${chainName}

	rlyCfgs := []chainCfg{
		{
			privKey:                mnemonicPrivKey,
			chainID:                chains[0].Config().ChainID,
			chainName:              chains[0].Config().Name,
			rpcUrl:                 chains[0].GetRPCAddress(),
			grpcUrl:                "http://" + chains[0].GetGRPCAddress(),
			originMailboxHex:       prefixedMailboxHex,
			domain:                 23456,
			validatorSignaturePath: simd1ValidatorSignaturesDir,
		},
		{
			privKey:                mnemonicPrivKey,
			chainID:                chains[1].Config().ChainID,
			chainName:              chains[1].Config().Name,
			rpcUrl:                 chains[1].GetRPCAddress(),
			grpcUrl:                "http://" + chains[1].GetGRPCAddress(),
			originMailboxHex:       prefixedMailboxHex,
			domain:                 34567,
			validatorSignaturePath: simd2ValidatorSignaturesDir,
		},
	}
	_, err = preconfigureHyperlaneRelayer(t, rly, rlyTmpDir, rlyCfgs)
	require.NoError(t, err)

	// Our images are currently local. You must build locally in monorepo, e.g. "cd rust && docker build .".
	// Also make sure that the tags in hyperlane.yaml match the local docker image repo and version.
	hyperlaneNetwork := hyperlane.NewHyperlaneNetwork(true, true)
	err = hyperlaneNetwork.Build(ctx, logger, eRep, opts, *valSimd1, *valSimd2, *rly)
	require.NoError(t, err)

	// Give the hyperlane validators and relayer time to start up and start watching the mailbox for the chain
	time.Sleep(10 * time.Second)

	var eg errgroup.Group
	numMsgs := 100

	for i := 0; i < numMsgs; i++ {
		i := i
		eg.Go(func() (err error) {
			dispatchedRecipientAddrHex, dispatchedMsgBody, dispatchSender, dispatchedMsgId := dispatchMsg(t, i, ctx, simd2Domain, recipientDispatch, userSimd.KeyName(), contract, simd1, logger)
			return processMsg(t, ctx, simd1IsmValidator, simd2, simd1Domain, simd2Domain, dispatchedRecipientAddrHex, dispatchedMsgBody, dispatchSender, dispatchedMsgId)
		})
	}

	err = eg.Wait()
	require.NoError(t, err)
}

func processMsg(
	t *testing.T,
	ctx context.Context,
	validator *counterchain.CounterChain,
	destChain *cosmos.CosmosChain,
	originDomain uint32,
	destDomain uint32,
	dispatchedRecipientAddrHex,
	dispatchedMsgBody,
	dispatchSender,
	dispatchedMsgId string) error {
	// Find the message ID of the dispatched message and wait for the message to be processed on the destination.
	dispatchedRecipientAddr := hexutil.MustDecode(dispatchedRecipientAddrHex)
	bech32Recipient := sdk.MustBech32ifyAddressBytes(sdk.GetConfig().GetBech32AccountAddrPrefix(), dispatchedRecipientAddr)

	b, err := hexutil.Decode(dispatchedMsgBody)
	require.NoError(t, err)
	message, _ := validator.CreateMessage(dispatchSender, originDomain, destDomain, bech32Recipient, string(b))
	messageId := validator.GetMessageId(message)
	require.Equal(t, dispatchedMsgId, hexutil.Encode(messageId))

	return Await(func() (bool, error) {
		return helpers.QueryMsgDelivered(t, ctx, destChain, dispatchedMsgId), nil
	}, 10*time.Minute, 5*time.Second)
}

// Dispatch the message to the chain and verify it was sent successfully
func dispatchMsg(
	t *testing.T,
	msgIndex int,
	ctx context.Context,
	destDomain uint32,
	recipientAddr string,
	keyName string,
	contract string,
	chain *cosmos.CosmosChain,
	logger *zap.Logger,
) (dispatchedRecipientAddrHex, dispatchedMsgBody, dispatchSender, dispatchedMsgId string) {
	// Dispatch a message to SIMD1
	dMsg := []byte("CosmosSimd1ToCosmosSimd2" + fmt.Sprintf("%d", msgIndex))
	dispatchedMsg := hexutil.Encode(dMsg)

	dispatchMsgStruct := helpers.ExecuteMsg{
		DispatchMsg: &helpers.DispatchMsg{
			DestinationAddr: destDomain,
			RecipientAddr:   recipientAddr,
			MessageBody:     dispatchedMsg,
		},
	}

	// Dispatch the hyperlane message to simd1
	dipatchMsg, err := json.Marshal(dispatchMsgStruct)
	require.NoError(t, err)
	dispatchedTxHash, err := chain.ExecuteContract(ctx, keyName, contract, string(dipatchMsg))
	require.NoError(t, err)
	logger.Info("Message dispatched to simd1")
	var dispatchedDestDomain string
	dispatchedDestDomain, dispatchedRecipientAddrHex, dispatchedMsgBody, dispatchedMsgId, dispatchSender, _, err = helpers.VerifyDispatchEvents(chain, dispatchedTxHash)
	require.NoError(t, err)
	require.NotEmpty(t, dispatchSender)
	require.NotEmpty(t, dispatchedRecipientAddrHex)
	require.Equal(t, fmt.Sprintf("%d", destDomain), dispatchedDestDomain)
	return
}

// e2e style test that spins up two Cosmos nodes (with different origin domains),
// a hyperlane validator (for Cosmos), and sends messages back and forth.
// Does not use a hyperlane relayer.
//
// IMPORTANT:
// Prior to running this test you must build the hyperlane-monorepo locally.
// You MUST tag the image it builds locally as hyperlane-monorepo:latest.
// Command will look like: docker tag 2dc725db78e3 hyperlane-monorepo:latest.
func TestHyperlaneCosmosValidator(t *testing.T) {
	logger := NewLogger(t)

	// Mailbox address - will be used in the hyperlane validator config
	mailboxHex, expectedMailbox := helpers.GetMailboxAddress()
	prefixedMailboxHex := "0x" + mailboxHex

	// Directories where files related to this test will be stored
	val1TmpDir := t.TempDir()
	val2TmpDir := t.TempDir()

	// Get the hyperlane agent raw configs (before variable replacements)
	valSimd1, valSimd2, _ := readHyperlaneConfig(t, COSMOS_E2E_CONFIG, logger)

	// Get the validator key for the agents. We also need this key to configure the chain ISM.
	valSimd1PrivKey, err := getHyperlaneBaseValidatorKey(valSimd1)
	require.NoError(t, err)
	valSimd2PrivKey, err := getHyperlaneBaseValidatorKey(valSimd2)
	require.NoError(t, err)

	// Build the chain docker image from the local repo
	optionalBuildChainImage()

	DockerImage := ibc.DockerImage{
		Repository: docker.HyperlaneImageName,
		Version:    "local",
		UidGid:     "1025:1025",
	}

	// Base setup
	chains := CreateHyperlaneSimds(t, DockerImage, []uint32{23456, 34567})
	simd1 := chains[0].(*cosmos.CosmosChain)
	simd1.SkipImagePull = true
	simd2 := chains[1].(*cosmos.CosmosChain)
	simd2.SkipImagePull = true

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

	// The initialization stage just finished and now the docker network is running for simd1 and simd2.
	// Now we need to configure the Hyperlane modules and setup some test users...
	simdDomainOutput := helpers.QueryDomain(t, ctx, simd1)
	simd2DomainOutput := helpers.QueryDomain(t, ctx, simd2)
	simdDomainStr := helpers.ParseQueryDomain(string(simdDomainOutput))
	simd2DomainStr := helpers.ParseQueryDomain(string(simd2DomainOutput))
	simdDomain, err := strconv.ParseUint(simdDomainStr, 10, 64)
	require.NoError(t, err)
	simd2Domain, err := strconv.ParseUint(simd2DomainStr, 10, 64)
	require.NoError(t, err)
	fmt.Printf("simd mailbox domain: %d, simd2 mailbox domain: %d\n", simdDomain, simd2Domain)
	_, err = icv7.GetAndFundTestUserWithMnemonic(ctx, "valannounce", mnemonic, int64(10_000_000_000), simd1)
	require.NoError(t, err)

	userSimd := icv7.GetAndFundTestUsers(t, ctx, "default", int64(10_000_000_000), simd1)[0]
	userSimd2 := icv7.GetAndFundTestUsers(t, ctx, "default", int64(10_000_000_000), simd2)[0]

	msg := `{}`
	_, contract := helpers.SetupContract(t, ctx, simd1, userSimd.KeyName(), "../contracts/hyperlane.wasm", msg)
	logger.Info("simd1 contract", zap.String("address", contract))
	_, contract2 := helpers.SetupContract(t, ctx, simd2, userSimd2.KeyName(), "../contracts/hyperlane.wasm", msg)
	logger.Info("simd2 contract", zap.String("address", contract2))

	verifyContractEntryPoints(t, ctx, simd1, userSimd, contract)
	verifyContractEntryPoints(t, ctx, simd2, userSimd2, contract2)

	// Create counter chain 1 with val set signing legacy multisig
	// The private key used here MUST be the one from the validator config file. TODO: cleanup this test to read it from the file.
	simd1IsmValidator := counterchain.CreateEmperorValidator(t, uint32(simdDomain), counterchain.LEGACY_MULTISIG, valSimd1PrivKey)
	// Create counter chain 2 with val set signing legacy multisig
	simd2IsmValidator := counterchain.CreateEmperorValidator(t, uint32(simd2Domain), counterchain.LEGACY_MULTISIG, valSimd2PrivKey)

	// Set default isms for counter chains for SIMD
	helpers.SetDefaultIsm(t, ctx, simd1, userSimd.KeyName(), simd2IsmValidator)
	// Set default isms for counter chains for SIMD2
	helpers.SetDefaultIsm(t, ctx, simd2, userSimd2.KeyName(), simd1IsmValidator)

	recipientAccAddr := sdk.MustAccAddressFromBech32(contract2).Bytes()
	recipientDispatch := hexutil.Encode([]byte(recipientAccAddr))
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

	valJson, err = preconfigureHyperlaneValidator(t, valSimd2, val2TmpDir, mnemonicPrivKey, chains[1].Config().ChainID, chains[1].Config().Name, chains[1].GetRPCAddress(), "http://"+chains[1].GetGRPCAddress(), prefixedMailboxHex, 34567)
	require.NoError(t, err)

	simd1ValidatorSignaturesDir := filepath.Join(val1TmpDir, "signatures-"+chains[0].Config().Name) //${val_dir}/signatures-${chainName}
	// simd2ValidatorSignaturesDir := filepath.Join(tmpDir2, "signatures-"+chains[1].Config().Name) //${val_dir}/signatures-${chainName}

	// Our images are currently local. You must build locally in monorepo, e.g. "cd rust && docker build .".
	// Also make sure that the tags in hyperlane.yaml match the local docker image repo and version.
	hyperlaneNetwork := hyperlane.NewHyperlaneNetwork(false, true)
	err = hyperlaneNetwork.Build(ctx, logger, eRep, opts, *valSimd1, *valSimd2)
	require.NoError(t, err)

	// Give the hyperlane validators time to start up and start watching the mailbox for the chain
	time.Sleep(10 * time.Second)

	// Dispatch a message to SIMD1
	dMsg := []byte("CosmosSimd1ToCosmosSimd2")
	dispatchedMsg := hexutil.Encode(dMsg)

	dispatchMsgStruct := helpers.ExecuteMsg{
		DispatchMsg: &helpers.DispatchMsg{
			DestinationAddr: uint32(simd2Domain),
			RecipientAddr:   recipientDispatch,
			MessageBody:     dispatchedMsg,
		},
	}

	dipatchMsg, err := json.Marshal(dispatchMsgStruct)
	require.NoError(t, err)
	dispatchedTxHash, err := simd1.ExecuteContract(ctx, userSimd.KeyName(), contract, string(dipatchMsg))
	require.NoError(t, err)
	logger.Info("Message dispatched to simd1")
	dispatchedDestDomain, dispatchedRecipientAddrHex, dispatchedMsgBody, dispatchedMsgId, dispatchSender, hyperlaneMsg, err := helpers.VerifyDispatchEvents(simd1, dispatchedTxHash)
	require.NoError(t, err)
	require.NotEmpty(t, dispatchSender)
	require.NotEmpty(t, dispatchedRecipientAddrHex)
	require.Equal(t, fmt.Sprintf("%d", simd2Domain), dispatchedDestDomain)
	// Finished sending message to simd1!

	// Wait for the hyperlane validator to sign it. The first message will show up as 0.json
	expectedSigFile := "0.json"
	simd1FirstSignedCheckpoint := filepath.Join(simd1ValidatorSignaturesDir, expectedSigFile)

	// Find the message that the hyperlane validator signed, and Process() it on SIMD2.
	// Here, we know what location the signed files should show up in based on our configuration.
	// (1) Wait for the 0_with_id.json file to show up in the validator's bind mount on the host.
	err = Await(func() (bool, error) {
		return fileExists(simd1FirstSignedCheckpoint), nil
	}, 1*time.Minute, 1*time.Second)
	require.NoError(t, err)
	valSig, err := os.ReadFile(simd1FirstSignedCheckpoint)
	require.NoError(t, err)
	signature := &ValidatorCheckpoint{}
	err = json.Unmarshal(valSig, signature)
	require.NoError(t, err)
	require.NotEmpty(t, signature.SerializedSignature)
	decodedValidatorSignature, err := hexutil.Decode(signature.SerializedSignature)
	require.NoError(t, err)

	// (2) Now use the validator 'announce' module to check if the validator announcement succeeded.
	announcedValidators := helpers.QueryAnnouncedValidators(t, ctx, simd1)
	fmt.Printf(string(announcedValidators))

	// (3) Process the signed message on SIMD2.
	dispatchedRecipientAddr := hexutil.MustDecode(dispatchedRecipientAddrHex)
	bech32Recipient := sdk.MustBech32ifyAddressBytes(sdk.GetConfig().GetBech32AccountAddrPrefix(), dispatchedRecipientAddr)
	b, err := hexutil.Decode(dispatchedMsgBody)
	require.NoError(t, err)

	// First we must 'fake' the relayer's portion of the data. We DO NOT SIGN, since we get the real signature from the validator.
	message, proof := simd1IsmValidator.CreateMessage(dispatchSender, uint32(simdDomain), uint32(simd2Domain), bech32Recipient, string(b))
	metadata := simd1IsmValidator.CreateRelayerLegacyMetadata(message, proof, originMailboxB)

	// Append the signature from the validator to the metadata.
	metadata = append(metadata, decodedValidatorSignature...)

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
}

type ValidatorCheckpoint struct {
	SerializedSignature string `json:"serialized_signature"`
}

func dispatchMessage() {
}

// TODO: per Steve, we should be waiting for the first mailbox message :)
// Check if the agent keeps trying to announce itself OR if it just waits after making a single announcement

func envVarValByKey(envVars []string, key string) string {
	for _, v := range envVars {
		res := helpers.ParseEnvVar(v, key)
		if res != "" {
			return res
		}
	}
	return ""
}

func TestEnvRegex(t *testing.T) {
	str := "HYP_BASE_CHECKPOINTSYNCER_PATH=/tmp/signatures-simd1"
	v := helpers.ParseEnvVar(str, "HYP_BASE_CHECKPOINTSYNCER_PATH")
	require.NotEmpty(t, v)
}
