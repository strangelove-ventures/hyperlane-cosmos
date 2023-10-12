package helpers

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"
	"github.com/stretchr/testify/require"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/announce/types"
)

// simd tx hyperlane-announce announce <hex-validator-address> <storageLocation> <hex-validator-signature>
func CallAnnounceMsg(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, keyName, hexValAddr, storageLoc, hexValSig string) []byte {
	cmd := []string{
		"simd", "tx", "hyperlane-announce", "announce",
		hexValAddr,
		storageLoc,
		hexValSig,
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

	fmt.Println("CallAnnounceMsg stdout: ", string(stdout))

	err = testutil.WaitForBlocks(ctx, 2, chain)
	require.NoError(t, err)
	return stdout
}

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

func VerifyAnnounceEvents(c *cosmos.CosmosChain, txHash string) (storageLocation, validatorAddr string, err error) {
	// Look up the events for the TX by hash
	events, err := GetEvents(c, txHash)
	if err != nil {
		return "", "", err
	}
	var found bool

	storageLocation, found = GetEventAttribute(events, types.EventTypeAnnounce, types.AttributeStorageLocation)
	if !found {
		return "", "", errors.New("storage location not found in TX event attrs")
	}

	validatorAddr, found = GetEventAttribute(events, types.EventTypeAnnounce, types.AttributeValidatorAddress)
	if !found {
		return "", "", errors.New("validator addr not found in TX event attrs")
	}

	return
}

func ParseEnvVar(envVar, key string) (value string) {
	//HYP_BASE_CHECKPOINTSYNCER_PATH=${val_dir}/signatures-simd1
	pattern := fmt.Sprintf(`(?m)^%s=(?P<val>.*)$`, key)
	r, _ := regexp.Compile(pattern)

	matches := r.FindStringSubmatch(envVar)
	if matches == nil {
		return
	}
	index := r.SubexpIndex("val")
	if index != -1 {
		value = matches[index]
		value = strings.Replace(value, "\"", "", -1)
	}
	return
}
