package tests

import (
	"context"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/docker"
	interchaintest "github.com/strangelove-ventures/interchaintest/v7"
	hyperlane "github.com/strangelove-ventures/interchaintest/v7/chain/hyperlane"
	"go.uber.org/zap/zaptest"

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

	hyperlaneCfg, err := hyperlane.ReadHyperlaneConfig(hyperlaneConfigPath, zaptest.NewLogger(t))
	require.NoError(t, err)
	valSimd1, ok := hyperlaneCfg["hyperlane-validator-simd1"]
	require.True(t, ok)
	valSimd2, ok := hyperlaneCfg["hyperlane-validator-simd2"]
	require.True(t, ok)
	rly, ok := hyperlaneCfg["hyperlane-relayer"]
	require.True(t, ok)

	err = preconfigureHyperlane(valSimd1, tmpDir1, "simd1", "http://simd1-rpc-url", "http://simd1-grpc-url", 23456)
	require.NoError(t, err)
	err = preconfigureHyperlane(valSimd2, tmpDir2, "simd2", "http://simd2-rpc-url", "http://simd1-grpc-url", 34567)
	require.NoError(t, err)

	logger := NewLogger(t)
	// Our images are currently local. You must build locally in monorepo, e.g. "cd rust && docker build".
	// Also make sure that the tags in hyperlane.yaml match the local docker image repo and version.
	hyperlaneNetwork := hyperlane.NewHyperlaneNetwork(false, true)
	hyperlaneNetwork.Build(ctx, logger, eRep, opts, *valSimd1, *valSimd2, *rly)
}

// e2e style test that spins up two Cosmos nodes (with different origin domains),
// a hyperlane validator and relayer (for Cosmos), and sends messages back and forth.
func TestHyperlaneCosmos(t *testing.T) {
	tmpDir1 := t.TempDir()
	tmpDir2 := t.TempDir()
	buildsEnabled := true
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)
	tarFilePath := filepath.Join(path, "../../")
	goModPath := filepath.Join(path, "../../go.mod")
	hyperlaneConfigPath := filepath.Join(path, "hyperlane.yaml")

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

	hyperlaneCfg, err := hyperlane.ReadHyperlaneConfig(hyperlaneConfigPath, zaptest.NewLogger(t))
	require.NoError(t, err)
	valSimd1, ok := hyperlaneCfg["hyperlane-validator-simd1"]
	require.True(t, ok)
	valSimd2, ok := hyperlaneCfg["hyperlane-validator-simd2"]
	require.True(t, ok)
	rly, ok := hyperlaneCfg["hyperlane-relayer"]
	require.True(t, ok)

	err = preconfigureHyperlane(valSimd1, tmpDir1, "simd1", chains[0].GetHostRPCAddress(), chains[0].GetHostGRPCAddress(), 23456)
	require.NoError(t, err)
	err = preconfigureHyperlane(valSimd2, tmpDir2, "simd2", chains[1].GetHostRPCAddress(), chains[1].GetHostGRPCAddress(), 34567)
	require.NoError(t, err)

	logger := NewLogger(t)
	// Our images are currently local. You must build locally in monorepo, e.g. "cd rust && docker build .".
	// Also make sure that the tags in hyperlane.yaml match the local docker image repo and version.
	hyperlaneNetwork := hyperlane.NewHyperlaneNetwork(false, true)
	hyperlaneNetwork.Build(ctx, logger, eRep, opts, *valSimd1, *valSimd2, *rly)
}
