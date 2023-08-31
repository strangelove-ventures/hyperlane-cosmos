package helpers

import (
	"errors"
	"regexp"
	"time"

	retry "github.com/avast/retry-go/v4"
	comettypes "github.com/cometbft/cometbft/abci/types"
	authTx "github.com/cosmos/cosmos-sdk/x/auth/tx"

	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
)

func getFullNode(c *cosmos.CosmosChain) *cosmos.ChainNode {
	if len(c.FullNodes) > 0 {
		// use first full node
		return c.FullNodes[0]
	}
	// use first validator
	return c.Validators[0]
}

func GetTransaction(c *cosmos.CosmosChain, txHash string) (*sdk.TxResponse, error) {
	// Retry because sometimes the tx is not committed to state yet.
	var txResp *types.TxResponse
	err := retry.Do(func() error {
		var err error
		txResp, err = authTx.QueryTx(getFullNode(c).CliContext(), txHash)
		return err
	},
		// retry for total of 3 seconds
		retry.Attempts(15),
		retry.Delay(200*time.Millisecond),
		retry.DelayType(retry.FixedDelay),
		retry.LastErrorOnly(true),
	)
	return txResp, err
}

func GetEvents(c *cosmos.CosmosChain, txHash string) ([]comettypes.Event, error) {
	// Look up the dispatched TX by hash
	txResp, err := GetTransaction(c, txHash)
	if err != nil {
		return nil, err
	} else if txResp.Code > 0 {
		return nil, errors.New(txResp.RawLog)
	}

	return txResp.Events, nil
}

func GetEventAttribute(events []comettypes.Event, eventType string, attributeKey string) (attrVal string, found bool) {
	for _, evt := range events {
		if evt.Type == eventType {
			for _, attr := range evt.Attributes {
				if attr.Key == attributeKey {
					attrVal = attr.Value
					found = true
					return
				}
			}
		}
	}

	return
}

func ParseTxHash(input string) string {
	// STDOUT will include the hash in its own line in the format "txhash: 6A9F49069B8D3E8B4CA92F76C695263C2F0E8C59B1BDD9A65E5A7C699A9F32DA"
	r, _ := regexp.Compile(`(?m)^txhash:\s(?P<hash>.*)$`)
	matches := r.FindStringSubmatch(input)
	hashIndex := r.SubexpIndex("hash")
	return matches[hashIndex]
}
