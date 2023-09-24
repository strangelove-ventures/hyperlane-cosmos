package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-cmd/cmd"
)

// RunCommand starts the command [bin] with the given [args] and returns the command to the caller
// TODO cmd package mentions we can do this more efficiently with cmd.NewCmdOptions rather than looping
// and calling Status().
func RunCommand(bin string, args ...string) (*cmd.Cmd, error) {
	curCmd := cmd.NewCmd(bin, args...)
	_ = curCmd.Start()

	// to stream outputs
	ticker := time.NewTicker(10 * time.Millisecond)
	go func() {
		prevLine := ""
		for range ticker.C {
			status := curCmd.Status()
			n := len(status.Stdout)
			if n == 0 {
				continue
			}

			line := status.Stdout[n-1]
			if prevLine != line && line != "" {
				fmt.Println("[streaming output]", line)
			}

			prevLine = line
		}
	}()

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

func AwaitHealthy(avaNodeHealthcheckUri string, maxWait time.Duration, retryInterval time.Duration) error {
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
			return errors.New("chain unhealthy after maxWait exceeded")
		case _ = <-ticker.C:
			isHealthy, _ := healthCheck(avaNodeHealthcheckUri)
			if isHealthy {
				return nil
			}
		}
	}
}

// Launch Avalanche local node network.
// subnetEvmPath - The path to the subnet-evm repo cloned from github.com/ava-labs/subnet-evm.git.
// localNodeUri - Will usually be "http://127.0.0.1:9650"
func launchAvalanche(subnetEvmPath, localNodeUri string) error {
	// TODO: wait for build to finish somehow
	_, err := RunCommand(subnetEvmPath + "/scripts/build.sh")
	if err != nil {
		return err
	}
	time.Sleep(2 * time.Second)

	cmd, err := RunCommand(subnetEvmPath + "/scripts/run.sh")
	if err != nil {
		return err
	}
	defer cmd.Stop()

	return AwaitHealthy(localNodeUri+"/ext/health", 5*time.Minute, 5*time.Second)
}
