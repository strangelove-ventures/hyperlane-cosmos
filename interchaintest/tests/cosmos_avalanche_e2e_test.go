package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/oriser/regroup"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/docker"
	interchaintest "github.com/strangelove-ventures/interchaintest/v7"
	hyperlane "github.com/strangelove-ventures/interchaintest/v7/chain/hyperlane"
	"go.uber.org/zap/zaptest"

	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/testreporter"
	"github.com/stretchr/testify/require"
)

func TestLaunchAvalanche(t *testing.T) {
	// The path to the subnet-evm repo cloned from github.com/ava-labs/subnet-evm.git.
	subnetEvmPath, ok := os.LookupEnv("AVALANCHE_SUBNETEVM_PATH")
	if !ok {
		fmt.Print("set AVALANCHE_SUBNETEVM_PATH to the directory where github.com/ava-labs/subnet-evm resides")
		t.FailNow()
	}

	localNodeUri := "http://127.0.0.1:9650"
	cmd, err := launchAvalanche(subnetEvmPath, localNodeUri)
	require.NoError(t, err)
	defer cmd.Stop()
}

// Gets the subnet-evm RPC Uri
func TestAvalancheGetRpcUri(t *testing.T) {
	// The path to the subnet-evm repo cloned from github.com/ava-labs/subnet-evm.git.
	subnetEvmPath, ok := os.LookupEnv("AVALANCHE_SUBNETEVM_PATH")
	if !ok {
		fmt.Print("set AVALANCHE_SUBNETEVM_PATH to the directory where github.com/ava-labs/subnet-evm resides")
		t.FailNow()
	}

	localNodeUri := "http://127.0.0.1:9650"
	cmd, err := launchAvalanche(subnetEvmPath, localNodeUri)
	require.NoError(t, err)
	defer cmd.Stop()

	jsonBody := []byte(`{"jsonrpc": "2.0","id": 1,"method": "info.getNodeID"}`)
	bodyReader := bytes.NewReader(jsonBody)

	requestURL := "http://127.0.0.1:9650/ext/info"
	resp, err := http.Post(requestURL, "application/json", bodyReader)
	require.NoError(t, err)
	b, _ := io.ReadAll(resp.Body)
	fmt.Printf("%+v", string(b))

	var res map[string]interface{}
	err = json.Unmarshal(b, &res)
	require.NoError(t, err)
	iNodeData := res["result"]
	nodeData := iNodeData.(map[string]interface{})
	iID := nodeData["nodeID"]
	fullID := iID.(string)
	//fmt.Printf("%s", fullID)
	r := regroup.MustCompile("NodeID-(?P<ID>.*)")
	match, err := r.Groups(fullID)
	nodeID := match["ID"]
	fmt.Printf("%s", nodeID)

	getBlockchains(t)

	jsonBody = []byte(`{"jsonrpc": "2.0","id": 1,"method": "eth_getChainConfig", "params" :[]}`)
	bodyReader = bytes.NewReader(jsonBody)
	subnetEvmUri := "http://127.0.0.1:9650/ext/bc/C/rpc"
	resp, err = http.Post(subnetEvmUri, "application/json", bodyReader)
	require.NoError(t, err)
	b, _ = io.ReadAll(resp.Body)
	fmt.Printf("%+v", string(b))
}

// This API call is deprecated according to https://docs.avax.network/reference/avalanchego/p-chain/api.
// There is no explanation of what the new API is, though, even in the release notes. So for now we use this.
// I believe we may need this to look up the blockchain ID for subnet-evm so we can query it later.
// But for now we are only using C-chain queries which are simpler, "http://127.0.0.1:9650/ext/bc/C/rpc"
func getBlockchains(t *testing.T) {
	jsonBody := []byte(`{"jsonrpc": "2.0","id": 1,"method": "platform.getBlockchains", "params" :[]}`)
	bodyReader := bytes.NewReader(jsonBody)
	subnetEvmUri := fmt.Sprintf("http://127.0.0.1:9650/ext/bc/%s", "P")
	resp, err := http.Post(subnetEvmUri, "application/json", bodyReader)
	require.NoError(t, err)
	b, _ := io.ReadAll(resp.Body)
	fmt.Printf("%+v", string(b))
}

// Testing Go commands to Avalanche local network.
// To spin up the local network, first:
// - Build the avalanchego binary by running ./script/build.sh from avalanchego root directory.
// - Copy the ./build/avalanchego to e.g. ${go env GOPATH}/src/github.com/ava-labs/avalanchego/build/avalanchego
// - Build subnet-evm by running ./scripts/build.sh in github.com/ConsiderItDone/subnet-evm AILC-141 branch
// - ls ${go env GOPATH}/github.com/ava-labs/avalanchego/build/plugins, there should be a plugin (binary) from above in there.
// - Finally, run the network with ./scripts/run.sh in github.com/ConsiderItDone/subnet-evm AILC-141
//
// TODO: if we're using subnet-evm as below, when we configure the hyperlane validator it may require a non-standard RPC endpoint.
func TestAvalancheLocalNetwork(t *testing.T) {
	// nodeURI := "http://127.0.0.1:9650"
	// ctx, cancel := context.WithTimeout(context.Background(), 100*time.Minute)
	// defer cancel()

	// kc := secp256k1fx.NewKeychain(genesis.EWOQKey)

	// // NewWalletFromURI fetches the available UTXOs owned by [kc] on the network
	// // that [LocalAPIURI] is hosting.
	// _, err := wallet.NewWalletFromURI(ctx, nodeURI, kc)
	// require.NoError(t, err)

	//pWallet := wallet.P()

	// owner := &secp256k1fx.OutputOwners{
	// 	Threshold: 1,
	// 	Addrs: []ids.ShortID{
	// 		genesis.EWOQKey.PublicKey().Address(),
	// 	},
	// }

	// genesisBytes, err := os.ReadFile("./avagenesis/ibc.json")
	// require.NoError(t, err)

	// createSubnetTxID, err := pWallet.IssueCreateSubnetTx(owner)
	// require.NoError(t, err)
	// t.Logf("new subnet id: %s", createSubnetTxID)

	// genesis := new(core.Genesis)
	// require.NoError(t, genesis.UnmarshalJSON(genesisBytes))

	// createChainTxID, err := pWallet.IssueCreateChainTx(
	// 	createSubnetTxID,
	// 	genesisBytes,
	// 	evm.ID,
	// 	nil,
	// 	"testChain",
	// )
	// require.NoError(t, err)
	// t.Logf("new chain id: %s", createSubnetTxID)

	// // Confirm the new blockchain is ready by waiting for the readiness endpoint
	// infoClient := info.NewClient(nodeURI)
	// bootstrapped, err := info.AwaitBootstrapped(ctx, infoClient, createChainTxID.String(), 2*time.Second)
	// require.NoError(t, err)
	// require.True(t, bootstrapped, "network isn't bootstaped")

	// chainURI := GetDefaultChainURI(nodeURI, createChainTxID.String())
	// t.Logf("subnet successfully created: %s", chainURI)

	// rpcClient, err := rpc.DialContext(ctx, chainURI)
	// require.NoError(t, err)

	// ethClient := ethclient.NewClient(rpcClient)
	// //subnetClient := subnetevmclient.New(rpcClient)

	// // TODO: check that this maps to the avagenesis/ibc.json
	// testKey, _ := crypto.HexToECDSA("56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027")
	// chainId := big.NewInt(99999)

	// auth, err = bind.NewKeyedTransactorWithChainID(testKey, chainId)
	// require.NoError(t, err)
	// t.Log("transactor created")
	// t.Log("eth client created")

	// utx := ethtypes.NewTransaction(senderNonce, toAddress, big.NewInt(amount.Amount), 21000, gasPrice, nil)
	// signedTx, err := ethtypes.SignTx(utx, ethtypes.NewEIP155Signer(chainID), privateKey)

}

// GetDefaultChainURI returns the default chain URI for a given blockchainID
func GetDefaultChainURI(nodeURI, blockchainID string) string {
	return fmt.Sprintf("%s/ext/bc/%s/rpc", nodeURI, blockchainID)
}

// Spins up a Cosmos node, Avalanche node, and hyperlane validator for Avalanche.
// The validator writes checkpoints to a local file. We read messages from the checkpoints
// and deliver each message to the Cosmos node. No messages are delivered from Cosmos->Ava.
func TestHyperlaneAvalancheCosmosDispatch(t *testing.T) {
	tmpDir1 := t.TempDir()
	buildsEnabled := true
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)
	tarFilePath := filepath.Join(path, "../../")
	goModPath := filepath.Join(path, "../../go.mod")
	hyperlaneConfigPath := filepath.Join(path, "hyperlane-cosmos-avalanche.yaml")

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

	domain := uint32(23456)
	// Create one chain, simd1, that will be part of the Cosmos <-> Avalanche hyperlane network
	chains := CreateHyperlaneSimds(t, DockerImage, []uint32{domain})

	// Create a new Interchain object which describes the chains
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

	hyperlaneCfg, err := hyperlane.ReadHyperlaneConfig(hyperlaneConfigPath, zaptest.NewLogger(t))
	require.NoError(t, err)
	valSimd1, ok := hyperlaneCfg["hyperlane-avalanche-validator"]
	require.True(t, ok)

	err = preconfigureHyperlane(valSimd1, tmpDir1, "simd1", chains[0].GetHostRPCAddress(), chains[0].GetHostGRPCAddress(), domain)
	require.NoError(t, err)

	logger := NewLogger(t)

	// Our images are currently local. You must build locally in monorepo, e.g. "cd rust && docker build".
	// Also make sure that the tags in hyperlane.yaml match the local docker image repo and version.
	hyperlaneNetwork := hyperlane.NewHyperlaneNetwork(false, true)
	hyperlaneNetwork.Build(ctx, logger, eRep, opts, *valSimd1)
}
