package helpers

import (
	"context"
	"fmt"
	"testing"

	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"
	"github.com/stretchr/testify/require"
)

// simd query hyperlane-announce getAnnouncedValidators
func QueryAnnouncedValidators(
	t *testing.T,
	ctx context.Context,
	chain *cosmos.CosmosChain,
) (stdout []byte) {
	cmd := []string{
		"simd", "query", "hyperlane-announce", "getAnnouncedValidators",
		"--node", chain.GetRPCAddress(),
		"--home", chain.HomeDir(),
		"--chain-id", chain.Config().ChainID,
	}
	stdout, _, err := chain.Exec(ctx, cmd, nil)
	require.NoError(t, err)

	fmt.Println("getAnnouncedValidators stdout: ", string(stdout))

	err = testutil.WaitForBlocks(ctx, 2, chain)
	require.NoError(t, err)
	return stdout
}

// simd query hyperlane-announce getAnnouncedStorageLocations
func QueryAnnouncedStorageLocations(
	t *testing.T,
	ctx context.Context,
	chain *cosmos.CosmosChain,
	validator string,
) (stdout []byte) {
	cmd := []string{
		"simd", "query", "hyperlane-announce", "getAnnouncedStorageLocations",
		validator,
		"--node", chain.GetRPCAddress(),
		"--home", chain.HomeDir(),
		"--chain-id", chain.Config().ChainID,
	}
	stdout, _, err := chain.Exec(ctx, cmd, nil)
	require.NoError(t, err)

	fmt.Println("getAnnouncedValidators stdout: ", string(stdout))

	err = testutil.WaitForBlocks(ctx, 2, chain)
	require.NoError(t, err)
	return stdout
}
