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

	err = preconfigureHyperlane(valSimd1, tmpDir1, "simd1", chains[0].GetHostRPCAddress(), domain)
	require.NoError(t, err)

	logger := NewLogger(t)

	// Our images are currently local. You must build locally in monorepo, e.g. "cd rust && docker build".
	// Also make sure that the tags in hyperlane.yaml match the local docker image repo and version.
	hyperlaneNetwork := hyperlane.NewHyperlaneNetwork(false, true)
	hyperlaneNetwork.Build(ctx, logger, eRep, opts, *valSimd1)
}
