package tests

import (
	"bytes"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

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
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	level := zap.NewAtomicLevelAt(zap.DebugLevel)

	c := zapcore.NewCore(
		consoleEncoder,
		writer,
		cfg.Level,
	)

	// Now create a second logger for logging everything to stdout
	stdout := zapcore.AddSync(os.Stdout)

	// Combine the loggers
	core := zapcore.NewTee(
		c,
		zapcore.NewCore(consoleEncoder, stdout, level),
	)

	return zap.New(
		core,
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
