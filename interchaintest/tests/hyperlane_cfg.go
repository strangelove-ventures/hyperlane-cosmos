package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/strangelove-ventures/interchaintest/v7/chain/hyperlane"
)

// tmpDir is a temporary directory UNIQUE to this hyperlane node. make sure to use a different one per node.
// chainName e.g. simd1 or simd2
// rpcUrl is the node RPC endpoint for e.g. simd1
// hyperlaneDomain is the chain's hyperlane domain, as configured in the chain app state or genesis
func preconfigureHyperlane(t *testing.T, node *hyperlane.HyperlaneChainConfig, tmpDir string, chainName string, chainRpcUrl string, chainGrpcUrl string, originMailbox string, hyperlaneDomain uint32) (valJson string, err error) {
	hyperlaneConfigPath := filepath.Join(tmpDir, chainName+".json")
	fmt.Printf("Chain: %s, RPC Uri: %s, GRPC Uri: %s\n", chainName, chainRpcUrl, chainGrpcUrl)

	// Write the hyperlane CONFIG_FILES to disk where the bind mount will expect it.
	// See also https://docs.hyperlane.xyz/docs/operators/agent-configuration#config-files-with-docker.
	valJson = generateHyperlaneValidatorConfig(chainName, chainRpcUrl, chainGrpcUrl, originMailbox, hyperlaneDomain)
	err = os.WriteFile(hyperlaneConfigPath, []byte(valJson), 777)
	if err != nil {
		return "", err
	}

	// Search and replace for the Docker env vars, cmd-flags, and bind-mounts, see hyperlane.yaml
	node.SetReplacements(map[string]string{
		"${val_dir}":          tmpDir,
		"${val_config_mount}": hyperlaneConfigPath,
		"${chainName}":        chainName,
		"${CHAINNAME}":        strings.ToUpper(chainName),
	})

	return valJson, nil
}

func generateHyperlaneValidatorConfig(chainName, rpcUrl, grpcUrl string, originMailboxHex string, domain uint32) string {
	rawJson := `{
		"chains": {
		  "%s": {
			"connection": { "rpc_url": "%s", "grpc_url": "%s" },
			"name": "%s",
			"domain": %d,
			"addresses": {
			  "mailbox": "%s",
			  "interchainGasPaymaster": "0x6cA0B6D22da47f091B7613223cD4BB03a2d77918",
			  "validatorAnnounce": "0x9bBdef63594D5FFc2f370Fe52115DdFFe97Bc524"
			},
			"protocol": "cosmosModules",
			"finalityBlocks": 1
		  }
		}
	  }`
	return fmt.Sprintf(rawJson, chainName, rpcUrl, grpcUrl, chainName, domain, originMailboxHex)
}

func getMailbox(valJson string, chain string) (mailbox string, err error) {
	data := map[string]interface{}{}
	err = json.Unmarshal([]byte(valJson), &data)
	if err != nil {
		return
	}
	ichains := data["chains"]
	imChains := ichains.(map[string]interface{})
	ichain := imChains[chain]
	imChain := ichain.(map[string]interface{})
	iAddr := imChain["addresses"]
	imAddr := iAddr.(map[string]interface{})
	iMailbox := imAddr["mailbox"]
	mailbox = iMailbox.(string)
	return
}
