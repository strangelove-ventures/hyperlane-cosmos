package tests

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ava-labs/avalanche-network-runner/local"
	"github.com/ava-labs/avalanche-network-runner/network"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/go-cmd/cmd"
	"go.uber.org/zap"
)

const (
	healthyTimeout = 2 * time.Minute
)

var goPath = os.ExpandEnv("$GOPATH")

func runAvalanche(log logging.Logger, binaryPath string) error {
	// Create the network
	nw, err := local.NewDefaultNetwork(log, binaryPath, true)
	if err != nil {
		return err
	}
	defer func() { // Stop the network when this function returns
		if err := nw.Stop(context.Background()); err != nil {
			log.Info("error stopping network", zap.Error(err))
		}
	}()

	// When we get a SIGINT or SIGTERM, stop the network and close [closedOnShutdownCh]
	signalsChan := make(chan os.Signal, 1)
	signal.Notify(signalsChan, syscall.SIGINT)
	signal.Notify(signalsChan, syscall.SIGTERM)
	closedOnShutdownCh := make(chan struct{})
	go func() {
		shutdownOnSignal(log, nw, signalsChan, closedOnShutdownCh)
	}()

	// Wait until the nodes in the network are ready
	ctx, cancel := context.WithTimeout(context.Background(), healthyTimeout)
	defer cancel()
	log.Info("waiting for all nodes to report healthy...")
	if err := nw.Healthy(ctx); err != nil {
		return err
	}

	log.Info("All nodes healthy. Network will run until you CTRL + C to exit...")
	// Wait until done shutting down network after SIGINT/SIGTERM
	<-closedOnShutdownCh
	return nil
}

// Blocks until a signal is received on [signalChan], upon which
// [n.Stop()] is called. If [signalChan] is closed, does nothing.
// Closes [closedOnShutdownChan] amd [signalChan] when done shutting down network.
// This function should only be called once.
func shutdownOnSignal(
	log logging.Logger,
	n network.Network,
	signalChan chan os.Signal,
	closedOnShutdownChan chan struct{},
) {
	sig := <-signalChan
	log.Info("got OS signal", zap.Stringer("signal", sig))
	if err := n.Stop(context.Background()); err != nil {
		log.Info("error stopping network", zap.Error(err))
	}
	signal.Reset()
	close(signalChan)
	close(closedOnShutdownChan)
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
