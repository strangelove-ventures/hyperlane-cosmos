package helpers

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/stretchr/testify/require"
)

func SetupContract(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, keyname string, fileLoc string, message string) (codeId, contract string) {
	codeId, err := chain.StoreContract(ctx, keyname, fileLoc)
	require.NoError(t, err)

	contractAddr, err := chain.InstantiateContract(ctx, keyname, codeId, message, true)
	require.NoError(t, err)

	return codeId, contractAddr
}

func SetContractsIsm(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, user ibc.Wallet, contractAddr string, ismId uint32) {
	setIsmMsg := ExecuteMsg{
		SetIsmId: &SetIsmId{
			IsmId: ismId,
		},
	}
	setIsmMsgBz, err := json.Marshal(setIsmMsg)
	require.NoError(t, err)
	_, err = chain.ExecuteContract(ctx, user.KeyName(), contractAddr, string(setIsmMsgBz))
	require.NoError(t, err)
}