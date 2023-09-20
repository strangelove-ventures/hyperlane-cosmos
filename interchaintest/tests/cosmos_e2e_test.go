package tests

import (
	"bytes"
	"context"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/strangelove-ventures/hyperlane-cosmos/interchaintest/docker"
	interchaintest "github.com/strangelove-ventures/interchaintest/v7"
	hyperlane "github.com/strangelove-ventures/interchaintest/v7/chain/hyperlane"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"

	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/testreporter"
	"github.com/stretchr/testify/require"
)

// Test that the hyperlane-agents heighliner image initializes with the given args and does not exit
func TestHyperlaneAgentInit(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)
	hyperlaneConfigPath := filepath.Join(path, "hyperlane.yaml")

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)
	ctx := context.Background()

	// docker client/network for hyperlane
	client, network := interchaintest.DockerSetup(t)
	opts := interchaintest.InterchainBuildOptions{
		TestName:          t.Name(),
		Client:            client,
		NetworkID:         network,
		SkipPathCreation:  true,
		BlockDatabaseFile: interchaintest.DefaultBlockDatabaseFilepath(),
	}

	hyperlaneCfg, err := hyperlane.ReadHyperlaneConfig(hyperlaneConfigPath, zaptest.NewLogger(t))
	require.NoError(t, err)
	valSimd1, ok := hyperlaneCfg["hyperlane-validator-simd1"]
	require.True(t, ok)
	valSimd2, ok := hyperlaneCfg["hyperlane-validator-simd2"]
	require.True(t, ok)
	rly, ok := hyperlaneCfg["hyperlane-relayer"]
	require.True(t, ok)

	// Search and replace for the Docker env vars, see hyperlane.yaml env-path
	valSimd1Replacements := map[string]string{
		"${validator_rpc_url}": "http://localhost:26657/validator1", // Fake value, we have no chains in this test case
	}
	valSimd1.SetEnvReplacements(valSimd1Replacements)

	// Search and replace for the Docker env vars, see hyperlane.yaml env-path
	valSimd2Replacements := map[string]string{
		"${validator_rpc_url}": "http://localhost:26657/validator2", // Fake value, we have no chains in this test case
	}
	valSimd2.SetEnvReplacements(valSimd2Replacements)

	logger := NewLogger(t)
	hyperlaneNetwork := hyperlane.HyperlaneNetwork{
		// Our images are currently local. You must build locally in monorepo, e.g. "cd rust && docker build".
		// Also make sure that the tags in hyperlane.yaml match the local docker image repo and version.
		DisableImagePull: true,
	}
	hyperlaneNetwork.Build(ctx, valSimd1, valSimd2, rly, logger, eRep, opts)
}

// LoggerOption configures the test logger built by NewLogger.
type LoggerOption interface {
	applyLoggerOption(*loggerOptions)
}

type loggerOptions struct {
	Level      zapcore.LevelEnabler
	zapOptions []zap.Option
}

type loggerOptionFunc func(*loggerOptions)

func (f loggerOptionFunc) applyLoggerOption(opts *loggerOptions) {
	f(opts)
}

func NewLogger(t zaptest.TestingT) *zap.Logger {
	cfg := loggerOptions{
		Level: zapcore.DebugLevel,
	}

	writer := newTestingWriter(t)
	zapOptions := []zap.Option{
		// Send zap errors to the same writer and mark the test as failed if
		// that happens.
		zap.ErrorOutput(writer.WithMarkFailed(true)),
	}
	zapOptions = append(zapOptions, cfg.zapOptions...)

	return zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			writer,
			cfg.Level,
		),
		zapOptions...,
	)
}

// WithMarkFailed returns a copy of this testingWriter with markFailed set to
// the provided value.
func (w testingWriter) WithMarkFailed(v bool) testingWriter {
	w.markFailed = v
	return w
}

func (w testingWriter) Write(p []byte) (n int, err error) {
	n = len(p)

	// Strip trailing newline because t.Log always adds one.
	p = bytes.TrimRight(p, "\n")

	// Note: t.Log is safe for concurrent use.
	w.t.Logf("%s", p)
	if w.markFailed {
		w.t.Fail()
	}

	return n, nil
}

func (w testingWriter) Sync() error {
	return nil
}

// testingWriter is a WriteSyncer that writes to the given testing.TB.
type testingWriter struct {
	t zaptest.TestingT

	// If true, the test will be marked as failed if this testingWriter is
	// ever used.
	markFailed bool
}

func newTestingWriter(t zaptest.TestingT) testingWriter {
	return testingWriter{t: t, markFailed: true}
}

// e2e style test that spins up two Cosmos nodes (with different origin domains),
// a hyperlane validator and relayer (for Cosmos), and sends messages back and forth.
func TestHyperlaneCosmos(t *testing.T) {
	buildsEnabled := false
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)
	tarFilePath := filepath.Join(path, "../../")
	goModPath := filepath.Join(path, "../../go.mod")
	hyperlaneConfigPath := filepath.Join(path, "hyperlane.yaml")

	// TODO: better caching mechanism to prevent rebuilding the same image
	if buildsEnabled {
		// Builds the hyperlane image from the current project (e.g. locally).
		// The directory at 'tarFilePath' will be tarballed and used for the docker context.
		// The args 'buildDir' and 'dockerfilePath' are relative to 'tarFilePath' (the context).
		// Build arguments are derived from 'goModPath' so it must be a full path (not relative).
		docker.BuildHeighlinerHyperlaneImage(docker.HyperlaneImageName, tarFilePath, ".", goModPath, "local.Dockerfile")
	}

	DockerImage := ibc.DockerImage{
		Repository: docker.HyperlaneImageName,
		Version:    "local",
		UidGid:     "1025:1025",
	}

	// Base setup
	chains := CreateDoubleHyperlaneSimd(t, DockerImage, 23456, 34567)

	// Create a new Interchain object which describes the chains, relayers, and IBC connections we want to use
	ic := interchaintest.NewInterchain()

	for _, chain := range chains {
		ic.AddChain(chain)
	}

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)

	ctx := context.Background()

	// Note: make sure that both the 'ic' interchain AND the hyperlane network share this client/network
	client, network := interchaintest.DockerSetup(t)
	opts := interchaintest.InterchainBuildOptions{
		TestName:          t.Name(),
		Client:            client,
		NetworkID:         network,
		SkipPathCreation:  true,
		BlockDatabaseFile: interchaintest.DefaultBlockDatabaseFilepath(),
	}

	err := ic.Build(ctx, eRep, opts)
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = ic.Close()
	})

	hyperlaneCfg, err := hyperlane.ReadHyperlaneConfig(hyperlaneConfigPath, zaptest.NewLogger(t))
	require.NoError(t, err)
	valSimd1, ok := hyperlaneCfg["hyperlane-validator-simd1"]
	require.True(t, ok)
	valSimd2, ok := hyperlaneCfg["hyperlane-validator-simd2"]
	require.True(t, ok)
	rly, ok := hyperlaneCfg["hyperlane-relayer"]
	require.True(t, ok)

	// Search and replace for the Docker env vars, see hyperlane.yaml env-path
	valSimd1Replacements := map[string]string{
		"${validator_rpc_url}": chains[0].GetHostRPCAddress(),
	}
	valSimd1.SetEnvReplacements(valSimd1Replacements)

	// Search and replace for the Docker env vars, see hyperlane.yaml env-path
	valSimd2Replacements := map[string]string{
		"${validator_rpc_url}": chains[1].GetHostRPCAddress(),
	}
	valSimd2.SetEnvReplacements(valSimd2Replacements)

	hyperlaneNetwork := hyperlane.HyperlaneNetwork{}
	hyperlaneNetwork.Build(ctx, valSimd1, valSimd2, rly, zaptest.NewLogger(t), eRep, opts)
}
