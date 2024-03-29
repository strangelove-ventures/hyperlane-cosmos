package helpers

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"
	mbtypes "github.com/strangelove-ventures/hyperlane-cosmos/x/mailbox/types"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

// simd tx hyperlane-mailbox process <metadata> <message>
func CallProcessMsg(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, keyName string, metadata string, message string) []byte {
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
	return stdout
}

func VerifyDispatchEvents(c *cosmos.CosmosChain, txHash string) (destDomain, recipientAddress, msgBody, dispatchId, sender, hyperlaneMsg string, err error) {
	// Look up the events for the TX by hash
	events, err := GetEvents(c, txHash)
	if err != nil {
		return "", "", "", "", "", "", err
	}
	var found bool

	sender, found = GetEventAttribute(events, types.EventTypeDispatch, types.AttributeKeySender)
	if !found {
		return "", "", "", "", "", "", errors.New("sender not found in dispatch TX event attrs")
	}

	destDomain, found = GetEventAttribute(events, types.EventTypeDispatch, types.AttributeKeyDestinationDomain)
	if !found {
		return "", "", "", "", "", "", errors.New("destdomain not found in dispatch TX event attrs")
	}

	recipientAddress, found = GetEventAttribute(events, types.EventTypeDispatch, types.AttributeKeyRecipientAddress)
	if !found {
		return "", "", "", "", "", "", errors.New("msgId not found in dispatch TX event attrs")
	}

	msgBody, found = GetEventAttribute(events, types.EventTypeDispatch, types.AttributeKeyMessage)
	if !found {
		return "", "", "", "", "", "", errors.New("msgBody not found in dispatch TX event attrs")
	}

	dispatchId, found = GetEventAttribute(events, types.EventTypeDispatchId, types.AttributeKeyID)
	if !found {
		return "", "", "", "", "", "", errors.New("dispatchid not found in dispatch TX event attrs")
	}
	hyperlaneMsg, found = GetEventAttribute(events, types.EventTypeDispatch, types.AttributeKeyHyperlaneMessage)
	if !found {
		return "", "", "", "", "", "", errors.New("hyperlane msg not found in dispatch TX event attrs")
	}

	return
}

func HasEvent(ctx context.Context, c *cosmos.CosmosChain, eventName string, attrKey string, attrVal string, height uint64) (bool, error) {
	txs, err := c.FindTxs(ctx, height)
	if err != nil {
		return false, err
	}
	for _, tx := range txs {
		for _, evt := range tx.Events {
			if evt.Type == eventName {
				for _, attr := range evt.Attributes {
					if attr.Key == attrKey && attr.Value == attrVal {
						return true, nil
					}
				}
			}
		}
	}

	return false, nil
}

func VerifyProcessEvents(c *cosmos.CosmosChain, txHash string) (msgId string, err error) {
	// Look up the events for the TX by hash
	events, err := GetEvents(c, txHash)
	if err != nil {
		return "", err
	}
	var found bool

	msgId, found = GetEventAttribute(events, types.EventTypeProcessId, types.AttributeKeyID)
	if !found {
		return "", errors.New("msgId not found in process TX event attrs")
	}
	return
}

// simd query hyperlane-mailbox domain
func QueryDomain(
	t *testing.T,
	ctx context.Context,
	chain *cosmos.CosmosChain,
) (stdout []byte) {
	cmd := []string{
		"simd", "query", "hyperlane-mailbox", "domain",
		"--node", chain.GetRPCAddress(),
		"--home", chain.HomeDir(),
		"--chain-id", chain.Config().ChainID,
	}
	stdout, _, err := chain.Exec(ctx, cmd, nil)
	require.NoError(t, err)

	fmt.Println("QueryDomain stdout: ", string(stdout))

	err = testutil.WaitForBlocks(ctx, 2, chain)
	require.NoError(t, err)
	return stdout
}

// simd query hyperlane-mailbox tree
func QueryCurrentTreeMetadata(
	t *testing.T,
	ctx context.Context,
	chain *cosmos.CosmosChain,
) (stdout []byte) {
	cmd := []string{
		"simd", "query", "hyperlane-mailbox", "tree-metadata",
		"--node", chain.GetRPCAddress(),
		"--home", chain.HomeDir(),
		"--chain-id", chain.Config().ChainID,
	}
	stdout, _, err := chain.Exec(ctx, cmd, nil)
	require.NoError(t, err)

	fmt.Println("QueryCurrentTreeMetadata stdout: ", string(stdout))

	err = testutil.WaitForBlocks(ctx, 2, chain)
	require.NoError(t, err)
	return stdout
}

// simd query hyperlane-mailbox delivered
func QueryMsgDelivered(
	t *testing.T,
	ctx context.Context,
	chain *cosmos.CosmosChain,
	msgId string,
) bool {
	cmd := []string{
		"simd", "query", "hyperlane-mailbox", "delivered", msgId,
		"--node", chain.GetRPCAddress(),
		"--home", chain.HomeDir(),
		"--chain-id", chain.Config().ChainID,
	}
	stdout, _, err := chain.Exec(ctx, cmd, nil)
	require.NoError(t, err)

	fmt.Println("MsgDelivered stdout: ", string(stdout))

	return strings.Contains(string(stdout), "true")
}

// simd query hyperlane-mailbox tree
func QueryCurrentTree(
	t *testing.T,
	ctx context.Context,
	chain *cosmos.CosmosChain,
) (stdout []byte) {
	cmd := []string{
		"simd", "query", "hyperlane-mailbox", "tree",
		"--node", chain.GetRPCAddress(),
		"--home", chain.HomeDir(),
		"--chain-id", chain.Config().ChainID,
	}
	stdout, _, err := chain.Exec(ctx, cmd, nil)
	require.NoError(t, err)

	fmt.Println("QueryCurrentTree stdout: ", string(stdout))

	return stdout
}

func ParseQueryTreeMetadata(input string) (root string, count string) {
	r, _ := regexp.Compile(`(?m)^root:\s(?P<root>.*)$`)

	matches := r.FindStringSubmatch(input)
	rootindex := r.SubexpIndex("root")
	root = matches[rootindex]
	root = strings.Replace(root, "\"", "", -1)

	r2, _ := regexp.Compile(`(?m)^count:\s(?P<count>.*)$`)

	matches = r2.FindStringSubmatch(input)
	cindex := r2.SubexpIndex("count")
	count = matches[cindex]
	count = strings.Replace(count, "\"", "", -1)

	return
}

func ParseQueryDomain(input string) (domain string) {
	r, _ := regexp.Compile(`(?m)^domain:\s(?P<domain>.*)$`)

	matches := r.FindStringSubmatch(input)
	dindex := r.SubexpIndex("domain")
	domain = matches[dindex]
	domain = strings.Replace(domain, "\"", "", -1)

	return
}

func GetMailboxAddress() (string, []byte) {
	mailboxAcc := authtypes.NewModuleAddress(mbtypes.ModuleName)
	mailboxAddr := hex.EncodeToString(mailboxAcc.Bytes())
	padding := 64 - len(mailboxAddr)

	for i := 0; i < padding; i++ {
		mailboxAddr = "0" + mailboxAddr
	}

	mailboxAddrBytes, _ := hex.DecodeString(mailboxAddr)
	return mailboxAddr, mailboxAddrBytes
}

func QueryRecipientsIsmId(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, recipient string) uint32 {
	grpcAddress := chain.GetHostGRPCAddress()
	conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()

	queryClient := mbtypes.NewQueryClient(conn)
	recipientBz, err := sdk.AccAddressFromBech32(recipient)
	require.NoError(t, err)
	res, err := queryClient.RecipientsIsmId(ctx, &mbtypes.QueryRecipientsIsmIdRequest{Recipient: recipientBz})
	require.NoError(t, err)

	return res.IsmId
}
