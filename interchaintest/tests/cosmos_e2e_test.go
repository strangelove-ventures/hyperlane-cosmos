package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/counterchain"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/helpers"
	icv7 "github.com/strangelove-ventures/interchaintest/v7"
	interchaintest "github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	hyperlane "github.com/strangelove-ventures/interchaintest/v7/chain/hyperlane"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"

	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/docker"

	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/testreporter"
	"github.com/stretchr/testify/require"
)

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

	err = preconfigureHyperlane(valSimd1, tmpDir1, "simd1", "http://simd1-rpc-url", "http://simd1-grpc-url", 23456)
	require.NoError(t, err)
	err = preconfigureHyperlane(valSimd2, tmpDir2, "simd2", "http://simd2-rpc-url", "http://simd1-grpc-url", 34567)
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

	DockerImage := ibc.DockerImage{
		Repository: docker.HyperlaneImageName,
		Version:    "local",
		UidGid:     "1025:1025",
	}

	// Base setup
	chains := CreateHyperlaneSimds(t, DockerImage, []uint32{23456, 34567})
	simd1 := chains[0].(*cosmos.CosmosChain)
	simd2 := chains[1].(*cosmos.CosmosChain)

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

	userSimd := icv7.GetAndFundTestUsers(t, ctx, "default", int64(10_000_000_000), simd1)[0]
	userSimd2 := icv7.GetAndFundTestUsers(t, ctx, "default", int64(10_000_000_000), simd2)[0]

	msg := `{}`
	_, contract := helpers.SetupContract(t, ctx, simd1, userSimd.KeyName(), "../contracts/hyperlane.wasm", msg)
	t.Log("coreContract", contract)
	_, contract2 := helpers.SetupContract(t, ctx, simd2, userSimd2.KeyName(), "../contracts/hyperlane.wasm", msg)
	t.Log("coreContract", contract2)

	verifyContractEntryPoints(t, ctx, simd1, userSimd, contract)
	verifyContractEntryPoints(t, ctx, simd2, userSimd2, contract2)

	// TODO: Right now the test case is not working because we need the validator private key in order to properly
	// set up the counterchain (and set the chain's ISM).
	// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	// Create counter chain 1 with val set signing legacy multisig
	simd1IsmValidator := counterchain.CreateCounterChain(t, uint32(simdDomain), counterchain.LEGACY_MULTISIG)
	// Create counter chain 2 with val set signing legacy multisig
	simd2IsmValidator := counterchain.CreateCounterChain(t, uint32(simd2Domain), counterchain.LEGACY_MULTISIG)

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

	logger.Info("Preconfiguring Hyperlane (getting configs)")
	err = preconfigureHyperlane(valSimd1, tmpDir1, chains[0].Config().Name, chains[0].GetRPCAddress(), "http://"+chains[0].GetGRPCAddress(), 23456)
	require.NoError(t, err)
	err = preconfigureHyperlane(valSimd2, tmpDir2, chains[1].Config().Name, chains[1].GetRPCAddress(), "http://"+chains[1].GetGRPCAddress(), 34567)
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

	// Wait for the hyperlane validator to sign it. The first message will show up as 0_with_id.json
	// TODO: ask the hyperlane team to explain what 1.json is.
	simd1FirstSignedCheckpoint := filepath.Join(simd1ValidatorSignaturesDir, "0_with_id.json")

	// Wait for the 0_with_id.json file to show up in the validator's bind mount on the host
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

	// Find the message that the hyperlane validator signed, and Process() it on SIMD2.
	dispatchedRecipientAddr := hexutil.MustDecode(dispatchedRecipientAddrHex)
	bech32Recipient := sdk.MustBech32ifyAddressBytes(sdk.GetConfig().GetBech32AccountAddrPrefix(), dispatchedRecipientAddr)
	b, err := hexutil.Decode(dispatchedMsgBody)
	require.NoError(t, err)

	// First we must 'fake' the relayer's portion of the data. We DO NOT SIGN, since we get the real signature from the validator.
	message, proof := simd1IsmValidator.CreateMessage(dispatchSender, uint32(simdDomain), uint32(simd2Domain), bech32Recipient, string(b))
	metadata := simd1IsmValidator.CreateRelayerLegacyMetadata(message, proof)

	//Append the signature from the validator to the metadata.
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

	err = testutil.WaitForBlocks(ctx, 2, simd2)
	require.NoError(t, err)
}

type ValidatorCheckpoint struct {
	SerializedSignature string `json:"serialized_signature"`
}

func dispatchMessage() {

}
