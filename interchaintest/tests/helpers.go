package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"testing"
	"time"

	"github.com/go-cmd/cmd"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/docker"
	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/helpers"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/chain/hyperlane"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

const (
	CHAIN_BUILD_ENABLED = false
)

// RunCommand starts the command [bin] with the given [args] and returns the command to the caller
// TODO cmd package mentions we can do this more efficiently with cmd.NewCmdOptions rather than looping
// and calling Status().
func RunCommand(bin string, args ...string) (*cmd.Cmd, error) {
	curCmd := cmd.NewCmd(bin, args...)
	_ = curCmd.Start()

	// to stream outputs
	// ticker := time.NewTicker(10 * time.Millisecond)
	// go func() {
	// 	prevLine := ""
	// 	for range ticker.C {
	// 		status := curCmd.Status()
	// 		n := len(status.Stdout)
	// 		if n == 0 {
	// 			continue
	// 		}

	// 		line := status.Stdout[n-1]
	// 		if prevLine != line && line != "" {
	// 			fmt.Println("[streaming output]", line)
	// 		}

	// 		prevLine = line
	// 	}
	// }()

	return curCmd, nil
}

func healthCheck(avaNodeHealthcheckUri string) (bool, error) {
	jsonBody := []byte(`{"jsonrpc": "2.0","id": 1,"method": "health.health"}`)
	bodyReader := bytes.NewReader(jsonBody)
	resp, err := http.Post(avaNodeHealthcheckUri, "application/json", bodyReader)
	if err != nil {
		return false, err
	}

	b, _ := io.ReadAll(resp.Body)
	fmt.Printf("%+v", string(b))

	var res map[string]interface{}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return false, err
	}
	iNodeData := res["result"]
	nodeData := iNodeData.(map[string]interface{})
	iHealthy := nodeData["healthy"]
	healthy := iHealthy.(bool)
	return healthy, nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func Await(f func() (bool, error), maxWait time.Duration, retryInterval time.Duration) error {
	if retryInterval > maxWait {
		return errors.New("retryInterval must be less than maxWait")
	}

	ticker := time.NewTicker(retryInterval)
	done := make(chan bool)
	go func() {
		time.Sleep(maxWait)
		ticker.Stop()
		done <- true
	}()

	for {
		select {
		case <-done:
			return errors.New("maxWait exceeded")
		case _ = <-ticker.C:
			success, _ := f()
			if success {
				return nil
			}
		}
	}
}

// Launch Avalanche local node network.
// subnetEvmPath - The path to the subnet-evm repo cloned from github.com/ava-labs/subnet-evm.git.
// localNodeUri - Will usually be "http://127.0.0.1:9650"
func launchAvalanche(subnetEvmPath, localNodeUri string) (*cmd.Cmd, error) {
	// TODO: wait for build to finish somehow
	_, err := RunCommand(subnetEvmPath + "/scripts/build.sh")
	if err != nil {
		return nil, err
	}
	time.Sleep(2 * time.Second)

	cmd, err := RunCommand(subnetEvmPath + "/scripts/run.sh")
	if err != nil {
		return nil, err
	}

	f := func() (bool, error) {
		return healthCheck(localNodeUri + "/ext/health")
	}
	return cmd, Await(f, 5*time.Minute, 5*time.Second)
}

func readHyperlaneConfig(t *testing.T, configName string, logger *zap.Logger) (
	valSimd1 *hyperlane.HyperlaneChainConfig,
	valSimd2 *hyperlane.HyperlaneChainConfig,
	rly *hyperlane.HyperlaneChainConfig,
) {
	var ok bool

	// Build the paths to the hyperlane config
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)
	hyperlaneConfigPath := filepath.Join(path, configName)

	// Read the hyperlane agent config file
	hyperlaneCfg, err := hyperlane.ReadHyperlaneConfig(hyperlaneConfigPath, logger)
	require.NoError(t, err)
	valSimd1, ok = hyperlaneCfg["hyperlane-validator-simd1"]
	require.True(t, ok)
	valSimd2, ok = hyperlaneCfg["hyperlane-validator-simd2"]
	require.True(t, ok)
	rly, ok = hyperlaneCfg["hyperlane-relayer"]
	require.True(t, ok)

	return
}

func verifyDomain(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, expected uint64) {
	domainOutput := helpers.QueryDomain(t, ctx, chain)
	domainStr := helpers.ParseQueryDomain(string(domainOutput))
	domain, err := strconv.ParseUint(domainStr, 10, 64)
	require.NoError(t, err)
	require.Equal(t, expected, domain)
}

func optionalBuildChainImage() {
	if !CHAIN_BUILD_ENABLED {
		return
	}

	// Build the paths to the root repository
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)
	tarFilePath := filepath.Join(path, "../../")
	goModPath := filepath.Join(path, "../../go.mod")

	// TODO: better caching mechanism to prevent rebuilding the same image
	// Builds the hyperlane image from the current project (e.g. locally).
	// The directory at 'tarFilePath' will be tarballed and used for the docker context.
	// The args 'buildDir' and 'dockerfilePath' are relative to 'tarFilePath' (the context).
	// Build arguments are derived from 'goModPath' so it must be a full path (not relative).
	docker.BuildHeighlinerHyperlaneImage(docker.HyperlaneImageName, tarFilePath, ".", goModPath, "local.Dockerfile")
}

func getHyperlaneBaseValidatorKey(cfg *hyperlane.HyperlaneChainConfig) (string, error) {
	if cfg.Type != "validator" {
		return "", errors.New(cfg.Name + " is not a hyperlane validator")
	}
	txt, _ := os.ReadFile(cfg.EnvPath)
	r, _ := regexp.Compile("(?m)^HYP_BASE_VALIDATOR_KEY=(?P<key>.*)$")
	matches := r.FindStringSubmatch(string(txt))
	if matches == nil {
		return "", errors.New("Key not found for validator " + cfg.Name)
	}
	idx := r.SubexpIndex("key")
	return matches[idx], nil
}
