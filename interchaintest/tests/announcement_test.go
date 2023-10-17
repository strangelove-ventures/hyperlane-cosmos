package tests

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/counterchain"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/docker"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/helpers"
	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
	icv7 "github.com/strangelove-ventures/interchaintest/v7"
	interchaintest "github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	hyperlane "github.com/strangelove-ventures/interchaintest/v7/chain/hyperlane"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/testreporter"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// Test the announce module 'announce' TX
func TestAnnounce(t *testing.T) {
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

	validatorPrivateKey, err := crypto.HexToECDSA(valSimd1PrivKey)
	require.NoError(t, err)
	valAddr := crypto.PubkeyToAddress(validatorPrivateKey.PublicKey)
	valAddrHex := hex.EncodeToString(valAddr.Bytes())

	valExpected := hexutil.Encode(valAddr.Bytes())
	val1Env, err := valSimd1.ReadEnvFile()
	require.NoError(t, err)
	expectedStorageLocation := envVarValByKey(val1Env, "HYP_BASE_CHECKPOINTSYNCER_PATH")
	require.NotEmpty(t, expectedStorageLocation)
	expectedStorageLocation = "file://" + expectedStorageLocation

	digest, err := types.GetAnnouncementDigest(uint32(simdDomain), simd1Mailbox, expectedStorageLocation)
	require.NoError(t, err)
	valSignature := simd1IsmValidator.Sign(digest)
	valSigHex := hex.EncodeToString(valSignature)

	//Announcement sends the announcement to the chain
	processStdout := helpers.CallAnnounceMsg(t, ctx, simd1, announceWallet.KeyName(), valAddrHex, expectedStorageLocation, valSigHex)
	announcementTxHash := helpers.ParseTxHash(string(processStdout))

	// Give the chain time to index the TX
	time.Sleep(10 * time.Second)

	fmt.Printf(announcementTxHash)
	evtStorageLocation, evtValAddr, err := helpers.VerifyAnnounceEvents(simd1, announcementTxHash)
	require.NoError(t, err)
	require.Equal(t, evtStorageLocation, expectedStorageLocation)
	require.Equal(t, evtValAddr, valAddrHex)

	// (2) Now use the validator 'announce' module to check if the validator announcement succeeded.
	err = Await(func() (bool, error) {
		announcedValidators := helpers.QueryAnnouncedValidators(t, ctx, simd1)
		announcedVals := string(announcedValidators)
		hasVal := strings.Contains(announcedVals, valExpected)
		return hasVal, nil
	}, 1*time.Minute, 5*time.Second)
	require.NoError(t, err)

	// (3) Now get the announced storage location and ensure it matches the expected one
	err = Await(func() (bool, error) {
		storageLocations := helpers.QueryAnnouncedStorageLocations(t, ctx, simd1, valExpected)
		locations := string(storageLocations)
		hasLoc := strings.Contains(locations, expectedStorageLocation)
		return hasLoc, nil
	}, 1*time.Minute, 5*time.Second)
	require.NoError(t, err)
}

// Test the hyperlane announce module with the monorepo (rust) validator agent making the announcement
func TestHyperlaneAnnounceWithValidator(t *testing.T) {
	logger := NewLogger(t)

	// Mailbox address - will be used in the hyperlane validator config
	mailboxHex, expectedMailbox := helpers.GetMailboxAddress()
	prefixedMailboxHex := "0x" + mailboxHex

	// Directories where files related to this test will be stored
	tmpDir := t.TempDir()

	// Get the hyperlane agent raw configs (before variable replacements)
	valSimd1, _, _ := readHyperlaneConfig(t, COSMOS_E2E_CONFIG, logger)

	// Get the validator key for the agents. We also need this key to configure the chain ISM.
	valSimd1PrivKey, err := getHyperlaneBaseValidatorKey(valSimd1)
	require.NoError(t, err)

	// Build the chain docker image from the local repo
	optionalBuildChainImage()

	DockerImage := ibc.DockerImage{
		Repository: docker.HyperlaneImageName,
		Version:    "local",
		UidGid:     "1025:1025",
	}

	// Base setup
	chains := CreateHyperlaneSimds(t, DockerImage, []uint32{23456})
	simd1 := chains[0].(*cosmos.CosmosChain)
	simd1.SkipImagePull = true

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
	simdDomainStr := helpers.ParseQueryDomain(string(simdDomainOutput))
	simdDomain, err := strconv.ParseUint(simdDomainStr, 10, 64)
	require.NoError(t, err)
	fmt.Printf("simd mailbox domain: %d\n", simdDomain)
	_, err = icv7.GetAndFundTestUserWithMnemonic(ctx, "valannounce", mnemonic, int64(10_000_000_000), simd1)
	//announceWallet, err := icv7.GetAndFundTestUserWithMnemonic(ctx, "valannounce", mnemonic, int64(10_000_000_000), simd1)
	require.NoError(t, err)

	userSimd := icv7.GetAndFundTestUsers(t, ctx, "default", int64(10_000_000_000), simd1)[0]

	msg := `{}`
	_, contract := helpers.SetupContract(t, ctx, simd1, userSimd.KeyName(), "../contracts/hyperlane.wasm", msg)
	logger.Info("simd1 contract", zap.String("address", contract))
	verifyContractEntryPoints(t, ctx, simd1, userSimd, contract)

	// Create counter chain with val set signing legacy multisig
	// The private key used here MUST be the one from the validator config file. TODO: cleanup this test to read it from the file.
	simd2IsmValidator := counterchain.CreateEmperorValidator(t, uint32(34567), counterchain.LEGACY_MULTISIG, valSimd1PrivKey)

	// Set default isms for counter chains for SIMD
	helpers.SetDefaultIsm(t, ctx, simd1, userSimd.KeyName(), simd2IsmValidator)

	valJson, err := preconfigureHyperlaneValidator(t, valSimd1, tmpDir, mnemonicPrivKey, chains[0].Config().ChainID, chains[0].Config().Name, chains[0].GetRPCAddress(), "http://"+chains[0].GetGRPCAddress(), prefixedMailboxHex, 23456)
	require.NoError(t, err)

	simd1MailboxHex, err := getMailbox(valJson, chains[0].Config().Name)
	require.NoError(t, err)
	_, simd1MailboxUnprefixed, found := strings.Cut(simd1MailboxHex, "0x")
	require.True(t, found)
	simd1Mailbox, err := hex.DecodeString(simd1MailboxUnprefixed)
	require.NoError(t, err)
	originMailboxB := []byte(simd1Mailbox)
	require.Equal(t, expectedMailbox, originMailboxB)

	// Our images are currently local. You must build locally in monorepo, e.g. "cd rust && docker build .".
	// Also make sure that the tags in hyperlane.yaml match the local docker image repo and version.
	hyperlaneNetwork := hyperlane.NewHyperlaneNetwork(false, true)
	err = hyperlaneNetwork.Build(ctx, logger, eRep, opts, *valSimd1)
	require.NoError(t, err)

	// Give the hyperlane validator time to start up and start watching the announcements for the chain
	time.Sleep(10 * time.Second)

	validatorPrivateKey, err := crypto.HexToECDSA(valSimd1PrivKey)
	require.NoError(t, err)
	valAddr := crypto.PubkeyToAddress(validatorPrivateKey.PublicKey)

	valExpected := hexutil.Encode(valAddr.Bytes())
	val1Env, err := valSimd1.ReadEnvFile()
	require.NoError(t, err)
	expectedStorageLocation := envVarValByKey(val1Env, "HYP_BASE_CHECKPOINTSYNCER_PATH")
	require.NotEmpty(t, expectedStorageLocation)

	// (2) Now use the validator 'announce' module to check if the validator announcement succeeded.
	err = Await(func() (bool, error) {
		announcedValidators := helpers.QueryAnnouncedValidators(t, ctx, simd1)
		announcedVals := string(announcedValidators)
		hasVal := strings.Contains(announcedVals, valExpected)
		return hasVal, nil
	}, 1*time.Minute, 5*time.Second)
	require.NoError(t, err)

	// (3) Now get the announced storage location and ensure it matches the expected one
	err = Await(func() (bool, error) {
		storageLocations := helpers.QueryAnnouncedStorageLocations(t, ctx, simd1, valExpected)
		locations := string(storageLocations)
		hasLoc := strings.Contains(locations, expectedStorageLocation)
		return hasLoc, nil
	}, 1*time.Minute, 5*time.Second)
	require.NoError(t, err)
}
