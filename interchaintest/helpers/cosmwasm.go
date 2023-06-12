package helpers

import (
	"context"
	"testing"

	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/stretchr/testify/require"
)

func SetupContract(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, keyname string, fileLoc string, message string) (codeId, contract string) {
	codeId, err := chain.StoreContract(ctx, keyname, fileLoc)
	require.NoError(t, err)

	contractAddr, err := chain.InstantiateContract(ctx, keyname, codeId, message, true)
	require.NoError(t, err)

	return codeId, contractAddr
}
