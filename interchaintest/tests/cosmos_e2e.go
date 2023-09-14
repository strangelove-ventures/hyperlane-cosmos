package tests

import (
	"context"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/docker"
	interchaintest "github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/testreporter"
	"github.com/stretchr/testify/require"
)

// e2e style test that spins up two Cosmos nodes (with different origin domains),
// a hyperlane validator and relayer (for Cosmos), and sends messages back and forth.
func TestHyperlaneCosmos(t *testing.T) {
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

	err := ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:          t.Name(),
		Client:            client,
		NetworkID:         network,
		SkipPathCreation:  true,
		BlockDatabaseFile: interchaintest.DefaultBlockDatabaseFilepath(),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = ic.Close()
	})
}
