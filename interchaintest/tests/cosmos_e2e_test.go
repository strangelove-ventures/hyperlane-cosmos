package tests

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	icv7 "github.com/strangelove-ventures/interchaintest/v7"
	interchaintest "github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	hyperlane "github.com/strangelove-ventures/interchaintest/v7/chain/hyperlane"
	"go.uber.org/zap"

	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/counterchain"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/helpers"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"

	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/docker"

	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/testreporter"
	"github.com/stretchr/testify/require"
)

const (
	mnemonic        = "spare number knock scan copper method lunch camera trap inject fine suspect edit sure design crowd sorry actual better spatial cover grit entire raccoon" // Testing only, do NOT use this mnemonic
	bech32Addr      = "cosmos13gpsgkxaavz3kcvh8y55xzat9umg944qnwxq4k"                                                                                                            // for the mnemonic above
	mnemonicPrivKey = "fe257759a16d7085ba4df68773c94647966e9fc8c7a7e1eb3311c40bbe1a0ed3"                                                                                         // for the mnemonic above                                                                                                      // Corresponds to the key above
	valPrivKey      = "8166f546bab6da521a8369cab06c5d2b9e46670292d85c875ee9ec20e84ffb61"                                                                                         // Testing only, do NOT use this key. Corresponds to the hyperlane validator signing key, not the mnemonic above
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

// Test that the hyperlane-agents heighliner image initializes with the given args and does not exit
func TestHyperlaneAgentInit(t *testing.T) {
	tmpDir1 := t.TempDir()
	tmpDir2 := t.TempDir()
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)
	hyperlaneConfigPath := filepath.Join(path, "hyperlane.yaml")
	logger := NewLogger(t)

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)
	ctx := context.Background()

	// docker client/network for hyperlane
	client, network := interchaintest.DockerSetup(t)
	opts := interchaintest.InterchainBuildOptions{
		TestName:          t.Name(),
		Client:            client,
		NetworkID:         network,
		SkipPathCreation:  true,
		BlockDatabaseFile: interchaintest.DefaultBlockDatabaseFilepath(),
	}

	hyperlaneCfg, err := hyperlane.ReadHyperlaneConfig(hyperlaneConfigPath, logger)
	require.NoError(t, err)
	valSimd1, ok := hyperlaneCfg["hyperlane-validator-simd1"]
	require.True(t, ok)
	valSimd2, ok := hyperlaneCfg["hyperlane-validator-simd2"]
	require.True(t, ok)
	// rly, ok := hyperlaneCfg["hyperlane-relayer"]
	// require.True(t, ok)

	mailboxHex := "000000000000000000000000cc2a110c8df654a38749178a04402e88f65091d3"
	prefixedMailboxHex := "0x" + mailboxHex
	// originMailbox, err := hex.DecodeString(mailboxHex)
	require.NoError(t, err)

	_, err = preconfigureHyperlane(t, valSimd1, tmpDir1, valPrivKey, "bech32address", "simd1", "simd1", "http://simd1-rpc-url", "http://simd1-grpc-url", prefixedMailboxHex, 23456)
	require.NoError(t, err)
	_, err = preconfigureHyperlane(t, valSimd2, tmpDir2, valPrivKey, "bech32address", "simd2", "simd2", "http://simd2-rpc-url", "http://simd1-grpc-url", prefixedMailboxHex, 34567)
	require.NoError(t, err)

	// Our images are currently local. You must build locally in monorepo, e.g. "cd rust && docker build".
	// Also make sure that the tags in hyperlane.yaml match the local docker image repo and version.
	hyperlaneNetwork := hyperlane.NewHyperlaneNetwork(false, true)
	err = hyperlaneNetwork.Build(ctx, logger, eRep, opts, *valSimd1, *valSimd2)
	require.NoError(t, err)
}

// e2e style test that spins up two Cosmos nodes (with different origin domains),
// a hyperlane validator and relayer (for Cosmos), and sends messages back and forth.
// IMPORTANT:
// Prior to running this test you must build the hyperlane-monorepo locally.
// You MUST tag the image it builds locally as hyperlane-monorepo:latest.
// Command will look like: docker tag 2dc725db78e3 hyperlane-monorepo:latest.
func TestHyperlaneCosmos(t *testing.T) {
	tmpDir1 := t.TempDir()
	tmpDir2 := t.TempDir()
	buildsEnabled := false
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)
	tarFilePath := filepath.Join(path, "../../")
	goModPath := filepath.Join(path, "../../go.mod")
	hyperlaneConfigPath := filepath.Join(path, "hyperlane.yaml")
	logger := NewLogger(t)

	// TODO: better caching mechanism to prevent rebuilding the same image
	if buildsEnabled {
		// Builds the hyperlane image from the current project (e.g. locally).
		// The directory at 'tarFilePath' will be tarballed and used for the docker context.
		// The args 'buildDir' and 'dockerfilePath' are relative to 'tarFilePath' (the context).
		// Build arguments are derived from 'goModPath' so it must be a full path (not relative).
		docker.BuildHeighlinerHyperlaneImage(docker.HyperlaneImageName, tarFilePath, ".", goModPath, "local.Dockerfile")
	}

	if t.Failed() {
		logger.Fatal("Test marked failed")
	}

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

	err := ic.Build(ctx, eRep, opts)
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
	simd1IsmValidator := counterchain.CreateEmperorValidator(t, uint32(simdDomain), counterchain.LEGACY_MULTISIG, valPrivKey)
	// Create counter chain 2 with val set signing legacy multisig
	simd2IsmValidator := counterchain.CreateEmperorValidator(t, uint32(simd2Domain), counterchain.LEGACY_MULTISIG, valPrivKey)

	// Set default isms for counter chains for SIMD
	helpers.SetDefaultIsm(t, ctx, simd1, userSimd.KeyName(), simd2IsmValidator)
	// Set default isms for counter chains for SIMD2
	helpers.SetDefaultIsm(t, ctx, simd2, userSimd2.KeyName(), simd1IsmValidator)

	recipientAccAddr := sdk.MustAccAddressFromBech32(contract2).Bytes()
	recipientDispatch := hexutil.Encode([]byte(recipientAccAddr))
	fmt.Printf("Recipient dispatch addr hex: %s", recipientDispatch)

	hyperlaneCfg, err := hyperlane.ReadHyperlaneConfig(hyperlaneConfigPath, logger)
	require.NoError(t, err)
	valSimd1, ok := hyperlaneCfg["hyperlane-validator-simd1"]
	require.True(t, ok)
	valSimd2, ok := hyperlaneCfg["hyperlane-validator-simd2"]
	require.True(t, ok)
	// rly, ok := hyperlaneCfg["hyperlane-relayer"]
	// require.True(t, ok)

	mailboxHex := "000000000000000000000000cc2a110c8df654a38749178a04402e88f65091d3"
	prefixedMailboxHex := "0x" + mailboxHex
	logger.Info("Preconfiguring Hyperlane (getting configs)")

	valJson, err := preconfigureHyperlane(t, valSimd1, tmpDir1, bech32Addr, mnemonicPrivKey, chains[0].Config().ChainID, chains[0].Config().Name, chains[0].GetRPCAddress(), "http://"+chains[0].GetGRPCAddress(), prefixedMailboxHex, 23456)
	require.NoError(t, err)

	simd1MailboxHex, err := getMailbox(valJson, chains[0].Config().Name)
	require.NoError(t, err)
	expectedMailbox, _ := hex.DecodeString("000000000000000000000000cc2a110c8df654a38749178a04402e88f65091d3")
	_, simd1MailboxUnprefixed, found := strings.Cut(simd1MailboxHex, "0x")
	require.True(t, found)
	simd1Mailbox, err := hex.DecodeString(simd1MailboxUnprefixed)
	require.NoError(t, err)
	originMailboxB := []byte(simd1Mailbox)
	require.Equal(t, expectedMailbox, originMailboxB)

	valJson, err = preconfigureHyperlane(t, valSimd2, tmpDir2, bech32Addr, mnemonicPrivKey, chains[1].Config().ChainID, chains[1].Config().Name, chains[1].GetRPCAddress(), "http://"+chains[1].GetGRPCAddress(), prefixedMailboxHex, 34567)
	require.NoError(t, err)
	simd1ValidatorSignaturesDir := filepath.Join(tmpDir1, "signatures-"+chains[0].Config().Name) //${val_dir}/signatures-${chainName}

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

// Test the announce module 'announce' TX
func TestAnnounce(t *testing.T) {
	tmpDir1 := t.TempDir()
	tmpDir2 := t.TempDir()
	buildsEnabled := false
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)
	tarFilePath := filepath.Join(path, "../../")
	goModPath := filepath.Join(path, "../../go.mod")
	hyperlaneConfigPath := filepath.Join(path, "hyperlane.yaml")
	logger := NewLogger(t)

	// TODO: better caching mechanism to prevent rebuilding the same image
	if buildsEnabled {
		// Builds the hyperlane image from the current project (e.g. locally).
		// The directory at 'tarFilePath' will be tarballed and used for the docker context.
		// The args 'buildDir' and 'dockerfilePath' are relative to 'tarFilePath' (the context).
		// Build arguments are derived from 'goModPath' so it must be a full path (not relative).
		docker.BuildHeighlinerHyperlaneImage(docker.HyperlaneImageName, tarFilePath, ".", goModPath, "local.Dockerfile")
	}

	if t.Failed() {
		logger.Fatal("Test marked failed")
	}

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

	err := ic.Build(ctx, eRep, opts)
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
	announceWallet, err := icv7.GetAndFundTestUserWithMnemonic(ctx, "valannounce", mnemonic, int64(10_000_000_000), simd1)
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
	simd1IsmValidator := counterchain.CreateEmperorValidator(t, uint32(simdDomain), counterchain.LEGACY_MULTISIG, valPrivKey)
	// Create counter chain 2 with val set signing legacy multisig
	simd2IsmValidator := counterchain.CreateEmperorValidator(t, uint32(simd2Domain), counterchain.LEGACY_MULTISIG, valPrivKey)

	// Set default isms for counter chains for SIMD
	helpers.SetDefaultIsm(t, ctx, simd1, userSimd.KeyName(), simd2IsmValidator)
	// Set default isms for counter chains for SIMD2
	helpers.SetDefaultIsm(t, ctx, simd2, userSimd2.KeyName(), simd1IsmValidator)

	recipientAccAddr := sdk.MustAccAddressFromBech32(contract2).Bytes()
	recipientDispatch := hexutil.Encode([]byte(recipientAccAddr))
	fmt.Printf("Recipient dispatch addr hex: %s", recipientDispatch)

	hyperlaneCfg, err := hyperlane.ReadHyperlaneConfig(hyperlaneConfigPath, logger)
	require.NoError(t, err)
	valSimd1, ok := hyperlaneCfg["hyperlane-validator-simd1"]
	require.True(t, ok)
	valSimd2, ok := hyperlaneCfg["hyperlane-validator-simd2"]
	require.True(t, ok)
	// rly, ok := hyperlaneCfg["hyperlane-relayer"]
	// require.True(t, ok)

	mailboxHex := "000000000000000000000000cc2a110c8df654a38749178a04402e88f65091d3"
	prefixedMailboxHex := "0x" + mailboxHex
	logger.Info("Preconfiguring Hyperlane (getting configs)")

	valJson, err := preconfigureHyperlane(t, valSimd1, tmpDir1, bech32Addr, mnemonicPrivKey, chains[0].Config().ChainID, chains[0].Config().Name, chains[0].GetRPCAddress(), "http://"+chains[0].GetGRPCAddress(), prefixedMailboxHex, 23456)
	require.NoError(t, err)

	simd1MailboxHex, err := getMailbox(valJson, chains[0].Config().Name)
	require.NoError(t, err)
	expectedMailbox, _ := hex.DecodeString("000000000000000000000000cc2a110c8df654a38749178a04402e88f65091d3")
	_, simd1MailboxUnprefixed, found := strings.Cut(simd1MailboxHex, "0x")
	require.True(t, found)
	simd1Mailbox, err := hex.DecodeString(simd1MailboxUnprefixed)
	require.NoError(t, err)
	originMailboxB := []byte(simd1Mailbox)
	require.Equal(t, expectedMailbox, originMailboxB)

	valJson, err = preconfigureHyperlane(t, valSimd2, tmpDir2, bech32Addr, mnemonicPrivKey, chains[1].Config().ChainID, chains[1].Config().Name, chains[1].GetRPCAddress(), "http://"+chains[1].GetGRPCAddress(), prefixedMailboxHex, 34567)
	require.NoError(t, err)
	//simd1ValidatorSignaturesDir := filepath.Join(tmpDir1, "signatures-"+chains[0].Config().Name) //${val_dir}/signatures-${chainName}

	// Our images are currently local. You must build locally in monorepo, e.g. "cd rust && docker build .".
	// Also make sure that the tags in hyperlane.yaml match the local docker image repo and version.
	hyperlaneNetwork := hyperlane.NewHyperlaneNetwork(false, true)
	err = hyperlaneNetwork.Build(ctx, logger, eRep, opts, *valSimd1, *valSimd2)
	require.NoError(t, err)

	// Give the hyperlane validators time to start up and start watching the mailbox for the chain
	time.Sleep(10 * time.Second)

	validatorPrivateKey, err := crypto.HexToECDSA(valPrivKey)
	require.NoError(t, err)
	valAddr := crypto.PubkeyToAddress(validatorPrivateKey.PublicKey)
	valAddrHex := hex.EncodeToString(valAddr.Bytes())
	storageLocation := "file:///tmp//signatures-simd1"

	digest, err := types.GetAnnouncementDigest(uint32(simdDomain), simd1Mailbox, storageLocation)
	require.NoError(t, err)
	valSignature := simd1IsmValidator.Sign(digest)
	valSigHex := hex.EncodeToString(valSignature)

	// Announcement sends the announcement to the chain
	processStdout := helpers.CallAnnounceMsg(t, ctx, simd1, announceWallet.KeyName(), valAddrHex, storageLocation, valSigHex)
	announcementTxHash := helpers.ParseTxHash(string(processStdout))
	fmt.Printf(announcementTxHash)
	evtStorageLocation, evtValAddr, err := helpers.VerifyAnnounceEvents(simd1, announcementTxHash)
	require.NoError(t, err)
	require.Equal(t, evtStorageLocation, storageLocation)
	require.Equal(t, evtValAddr, valAddrHex)

	// (2) Now use the validator 'announce' module to check if the validator announcement succeeded.
	announcedValidators := helpers.QueryAnnouncedValidators(t, ctx, simd1)
	announcedVals := string(announcedValidators)
	valExpected := string(valAddr.Bytes())
	require.Contains(t, announcedVals, valExpected)
	fmt.Printf(announcedVals)
}
