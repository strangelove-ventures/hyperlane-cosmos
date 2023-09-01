package helpers

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"
	"github.com/stretchr/testify/require"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/igp/types"
)

func VerifyIgpEvents(c *cosmos.CosmosChain, txHash string) (igpId uint32, err error) {
	// Look up the events for the TX by hash
	events, err := GetEvents(c, txHash)
	if err != nil {
		return 0, err
	}
	igp, found := GetEventAttribute(events, types.EventTypeCreateIgp, types.AttributeIgpId)
	if !found {
		return 0, errors.New("IGP ID not found in TX event attrs")
	}
	igpId64, err := strconv.ParseUint(igp, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint32(igpId64), err
}

func VerifyCreateOracleEvents(c *cosmos.CosmosChain, txHash string) (oracleAddr string, err error) {
	// Look up the events for the TX by hash
	events, err := GetEvents(c, txHash)
	if err != nil {
		return "", err
	}

	oracleAddr, found := GetEventAttribute(events, types.EventTypeCreateOracle, types.AttributeOracleAddress)
	if !found {
		return "", errors.New("oracle address not found in TX event attrs")
	}

	return oracleAddr, err
}

func VerifySetGasPriceEvents(c *cosmos.CosmosChain, txHash string) (remoteDomain string, exchRate string, gasPrice string, err error) {
	// Look up the events for the TX by hash
	events, err := GetEvents(c, txHash)
	if err != nil {
		return "", "", "", err
	}

	remoteDomain, found := GetEventAttribute(events, types.EventTypeGasDataSet, types.AttributeRemoteDomain)
	if !found {
		return "", "", "", errors.New("remote domain not found in setGasPrice TX event attrs")
	}

	exchRate, found = GetEventAttribute(events, types.EventTypeGasDataSet, types.AttributeTokenExchangeRate)
	if !found {
		return "", "", "", errors.New("exchRate not found in setGasPrice TX event attrs")
	}

	gasPrice, found = GetEventAttribute(events, types.EventTypeGasDataSet, types.AttributeGasPrice)
	if !found {
		return "", "", "", errors.New("gasPrice not found in setGasPrice TX event attrs")
	}

	return
}

func VerifyPayForGasEvents(c *cosmos.CosmosChain, txHash string) (msgId string, payment string, gasAmount string, err error) {
	// Look up the events for the TX by hash
	events, err := GetEvents(c, txHash)
	if err != nil {
		return "", "", "", err
	}
	var found bool
	msgId, found = GetEventAttribute(events, types.EventTypePayForGas, types.AttributeMessageId)
	if !found {
		return "", "", "", errors.New("message ID not found in PayForGas TX event attrs")
	}

	payment, found = GetEventAttribute(events, types.EventTypePayForGas, types.AttributePayment)
	if !found {
		return "", "", "", errors.New("payment not found in PayForGas TX event attrs")
	}

	gasAmount, found = GetEventAttribute(events, types.EventTypePayForGas, types.AttributeGasAmount)
	if !found {
		return "", "", "", errors.New("gas price not found in PayForGas TX event attrs")
	}

	return
}

func ParseQuoteGasPayment(input string) (amount, denom string) {
	// Input will look like this:
	// amount: "2250000000000000000000"
	// denom: stake

	r, _ := regexp.Compile(`(?m)^amount:\s(?P<amount>.*)$`)
	r2, _ := regexp.Compile(`(?m)^denom:\s(?P<denom>.*)$`)

	matches := r.FindStringSubmatch(input)
	amtIndex := r.SubexpIndex("amount")
	amount = matches[amtIndex]
	amount = strings.Replace(amount, "\"", "", -1)

	matches = r2.FindStringSubmatch(input)
	denomIndex := r2.SubexpIndex("denom")
	denom = matches[denomIndex]
	denom = strings.Replace(denom, "\"", "", -1)

	return
}

// simd tx hyperlane-igp createigp <beneficiary>
func CallCreateIgp(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, keyName, beneficiary string) (stdout []byte) {
	cmd := []string{
		"simd", "tx", "hyperlane-igp", "createigp",
		beneficiary,
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

	fmt.Println("CallCreateIgp stdout: ", string(stdout))

	err = testutil.WaitForBlocks(ctx, 2, chain)
	require.NoError(t, err)
	return stdout
}

// simd tx hyperlane-igp createoracle <oracle-addr> <igp-id> <remote-domain>
func CallCreateOracle(
	t *testing.T,
	ctx context.Context,
	chain *cosmos.CosmosChain,
	keyName,
	oracleAddr string,
	igpId string,
	remoteDomain string,
) (stdout []byte) {
	cmd := []string{
		"simd", "tx", "hyperlane-igp", "createoracle",
		oracleAddr,
		igpId,
		remoteDomain,
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

	fmt.Println("CallCreateOracle stdout: ", string(stdout))

	err = testutil.WaitForBlocks(ctx, 2, chain)
	require.NoError(t, err)
	return stdout
}

// simd tx hyperlane-igp msgpay <message-id> <destination-domain> <destination-gas-amount> <igp-id> <max-payment>
func CallPayForGasMsg(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, keyName, msgId, domain, destGas, igpId, maxPayment string) (stdout []byte) {
	cmd := []string{
		"simd", "tx", "hyperlane-igp", "msgpay",
		msgId,
		domain,
		destGas,
		igpId,
		maxPayment,
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

	fmt.Println("CallPayForGasMsg stdout: ", string(stdout))

	err = testutil.WaitForBlocks(ctx, 2, chain)
	require.NoError(t, err)
	return stdout
}

// simd tx hyperlane-igp setgasprice <igp-id> <remote-domain> <gas-price> <exch-rate>
func CallSetGasPriceMsg(
	t *testing.T,
	ctx context.Context,
	chain *cosmos.CosmosChain,
	keyName,
	igpId,
	domain,
	gasPrice,
	exchRate string,
) (stdout []byte) {
	cmd := []string{
		"simd", "tx", "hyperlane-igp", "setgasprice",
		igpId,
		domain,
		gasPrice,
		exchRate,
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

	fmt.Println("CallSetGasPriceMsg stdout: ", string(stdout))

	err = testutil.WaitForBlocks(ctx, 2, chain)
	require.NoError(t, err)
	return stdout
}

// simd query hyperlane-igp setgasprice <igp-id> <remote-domain> <gas-price> <exch-rate>
func QueryQuoteGasPayment(
	t *testing.T,
	ctx context.Context,
	chain *cosmos.CosmosChain,
	keyName,
	igpId,
	domain,
	gasAmount string,
) (stdout []byte) {
	cmd := []string{
		"simd", "query", "hyperlane-igp", "quoteGasPayment",
		igpId,
		domain,
		gasAmount,
		"--node", chain.GetRPCAddress(),
		"--home", chain.HomeDir(),
		"--chain-id", chain.Config().ChainID,
	}
	stdout, _, err := chain.Exec(ctx, cmd, nil)
	require.NoError(t, err)

	fmt.Println("CallSetGasPriceMsg stdout: ", string(stdout))

	err = testutil.WaitForBlocks(ctx, 2, chain)
	require.NoError(t, err)
	return stdout
}
