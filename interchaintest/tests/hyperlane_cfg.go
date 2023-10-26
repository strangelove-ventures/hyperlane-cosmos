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
// hyperlaneDomain is the chain's hyperlane domain, as configured in the chain app state or genesis
func preconfigureHyperlaneValidator(
	t *testing.T,
	node *hyperlane.HyperlaneChainConfig,
	tmpDir,
	privKey,
	chainID,
	chainName, // e.g. simd1 or simd2
	chainRpcUrl, // RPC endpoint for e.g. simd1
	chainGrpcUrl, // gRPC endpoint for e.g. simd1
	originMailbox string,
	hyperlaneDomain uint32,
) (valJson string, err error) {
	hyperlaneConfigPath := filepath.Join(tmpDir, chainName+".json")
	fmt.Printf("Chain: %s, RPC Uri: %s, GRPC Uri: %s\n", chainName, chainRpcUrl, chainGrpcUrl)

	// Write the hyperlane CONFIG_FILES to disk where the bind mount will expect it.
	// See also https://docs.hyperlane.xyz/docs/operators/agent-configuration#config-files-with-docker.
	valJson = generateHyperlaneValidatorConfig(privKey, chainID, chainName, chainRpcUrl, chainGrpcUrl, originMailbox, hyperlaneDomain)
	err = os.WriteFile(hyperlaneConfigPath, []byte(valJson), 777)
	if err != nil {
		return
	}
	// Search and replace for the Docker env vars, cmd-flags, and bind-mounts, see hyperlane.yaml
	node.SetReplacements(map[string]string{
		"${val_dir}":          tmpDir,
		"${val_config_mount}": hyperlaneConfigPath,
		"${chainName}":        chainName,
		"${CHAINNAME}":        strings.ToUpper(chainName),
	})
	return
}

// tmpDir is a temporary directory UNIQUE to this hyperlane node. make sure to use a different one per node.
// hyperlaneDomain is the chain's hyperlane domain, as configured in the chain app state or genesis
func preconfigureHyperlaneRelayer(t *testing.T, node *hyperlane.HyperlaneChainConfig, tmpDir string, chains []chainCfg) (rlyJson string, err error) {
	hyperlaneConfigPath := filepath.Join(tmpDir, "rly.json")

	// Write the hyperlane CONFIG_FILES to disk where the bind mount will expect it.
	// See also https://docs.hyperlane.xyz/docs/operators/agent-configuration#config-files-with-docker.
	rlyJson = generateHyperlaneRelayerConfig(chains)
	err = os.WriteFile(hyperlaneConfigPath, []byte(rlyJson), 777)
	if err != nil {
		return "", err
	}

	// create a comma separated list of chain names for the hyperlane `HYP_BASE_RELAYCHAINS` config param
	chainNamesCsv := ""
	for i, chain := range chains {
		chainNamesCsv = chainNamesCsv + chain.chainName
		if i < len(chains)-1 {
			chainNamesCsv = chainNamesCsv + ","
		}
	}

	// Create a bind mount for the directories where the relayer can look for the validator signatures
	valBindMounts := ""
	for i, chain := range chains {
		valBindMounts = valBindMounts + chain.validatorSignaturePath + ":" + chain.validatorSignaturePath
		if i < len(chains)-1 {
			valBindMounts = valBindMounts + ","
		}
	}

	// Search and replace for the Docker env vars, cmd-flags, and bind-mounts, see hyperlane.yaml
	node.SetReplacements(map[string]string{
		"${val-sig-binds}": valBindMounts,
		"${rly_dir}":       tmpDir,
		"${chainNamesCsv}": chainNamesCsv,
	})

	return rlyJson, nil
}

type hyperlaneChainConnection struct {
	RpcUrl  string `json:"rpc_url,omitempty"`
	GrpcUrl string `json:"grpc_url,omitempty"`
	ChainID string `json:"chain_id,omitempty"`
}
type hyperlaneSigner struct {
	Type      string `json:"type,omitempty"`
	Key       string `json:"key,omitempty"`
	Prefix    string `json:"prefix,omitempty"`
	BaseDenom string `json:"base_denom,omitempty"`
}
type hyperlaneAddresses struct {
	Mailbox  string `json:"mailbox,omitempty"`
	Igp      string `json:"interchainGasPaymaster,omitempty"`
	Announce string `json:"validatorAnnounce,omitempty"`
}
type hyperlaneChainCfg struct {
	Connection     *hyperlaneChainConnection `json:"connection,omitempty"`
	Signer         *hyperlaneSigner          `json:"signer,omitempty"`
	Addresses      *hyperlaneAddresses       `json:"addresses,omitempty"`
	Name           string                    `json:"name,omitempty"`
	Domain         uint32                    `json:"domain,omitempty"`
	Protocol       string                    `json:"protocol,omitempty"`
	FinalityBlocks int                       `json:"finalityBlocks,omitempty"`
}

type hyperlaneRelayerConfig struct {
	Chains map[string]hyperlaneChainCfg `json:"chains"`
}

type chainCfg struct {
	privKey                string
	chainID                string
	chainName              string
	rpcUrl                 string
	grpcUrl                string
	originMailboxHex       string
	domain                 uint32
	validatorSignaturePath string
}

func generateHyperlaneRelayerConfig(chains []chainCfg) string {
	cfg := hyperlaneRelayerConfig{
		Chains: map[string]hyperlaneChainCfg{},
	}
	for _, chain := range chains {
		chainCfg := hyperlaneChainCfg{
			Connection: &hyperlaneChainConnection{
				RpcUrl:  chain.rpcUrl,
				GrpcUrl: chain.grpcUrl,
				ChainID: chain.chainID,
			},
			Signer: &hyperlaneSigner{
				Type:      "cosmosKey",
				Key:       chain.privKey,
				Prefix:    "cosmos",
				BaseDenom: "stake",
			},
			Addresses: &hyperlaneAddresses{
				Mailbox:  chain.originMailboxHex,
				Igp:      "0x6cA0B6D22da47f091B7613223cD4BB03a2d77918",
				Announce: "0x9bBdef63594D5FFc2f370Fe52115DdFFe97Bc524",
			},
			Name:           chain.chainName,
			Domain:         chain.domain,
			Protocol:       "cosmosModules",
			FinalityBlocks: 1,
		}
		cfg.Chains[chain.chainName] = chainCfg
	}
	b, err := json.Marshal(cfg)
	if err != nil {
		panic("Could not Marshal hyperlaneRelayerConfig")
	}
	return string(b)
}

func generateHyperlaneValidatorConfig(privKey, chainID, chainName, rpcUrl, grpcUrl string, originMailboxHex string, domain uint32) string {
	rawJson := `{
		"chains": {
		  "%s": {
			"connection": { "rpc_url": "%s", "grpc_url": "%s", "chain_id": "%s" },
			"signer": { "type":"cosmosKey", "key": "%s", "prefix": "cosmos", "base_denom": "stake"},
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
	return fmt.Sprintf(rawJson, chainName, rpcUrl, grpcUrl, chainID, privKey, chainName, domain, originMailboxHex)
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
