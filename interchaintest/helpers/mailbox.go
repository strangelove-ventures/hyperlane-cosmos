package helpers

import (
	"context"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"
	"github.com/stretchr/testify/require"
)

// simd tx hyperlane-mailbox process <metadata> <message>
func CallProcessMsg(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, keyName string, metadata string, message string) {
	cmd := []string{
		"simd", "tx", "hyperlane-mailbox", "process",
		metadata,
		message,
		"--node", chain.GetRPCAddress(),
		"--home", chain.HomeDir(),
		"--chain-id", chain.Config().ChainID,
		"--from", keyName,
		"--gas", "2500000",
		"--gas-adjustment", "2.0",
		"--keyring-dir", chain.HomeDir(),
		"--keyring-backend", keyring.BackendTest,
		"-y",
	}
	stdout, _, err := chain.Exec(ctx, cmd, nil)
	require.NoError(t, err)

	fmt.Println("CallProcessMsg stdout: ", string(stdout))

	err = testutil.WaitForBlocks(ctx, 2, chain)
	require.NoError(t, err)
}
