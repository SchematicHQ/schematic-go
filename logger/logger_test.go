package logger

import (
	"bytes"
	"context"
	"log"
	"strings"
	"testing"

	"github.com/schematichq/schematic-go/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// newTestLogger returns a defaultLogger writing to a buffer so emissions
// can be asserted without touching stderr.
func newTestLogger() (*defaultLogger, *bytes.Buffer) {
	var buf bytes.Buffer
	return &defaultLogger{
		logger: log.New(&buf, "", 0),
		level:  core.LogLevelWarn,
	}, &buf
}

func TestNewDefaultLogger_DefaultsToWarn(t *testing.T) {
	l := NewDefaultLogger()
	getter, ok := l.(interface{ GetLevel() core.LogLevel })
	require.True(t, ok, "default logger must expose GetLevel")
	assert.Equal(t, core.LogLevelWarn, getter.GetLevel())
}

func TestDefaultLogger_FiltersBelowWarn(t *testing.T) {
	l, buf := newTestLogger()
	ctx := context.Background()

	l.Debug(ctx, "debug-message")
	l.Info(ctx, "info-message")
	l.Warn(ctx, "warn-message")
	l.Error(ctx, "error-message")

	out := buf.String()
	assert.NotContains(t, out, "debug-message")
	assert.NotContains(t, out, "info-message")
	assert.Contains(t, out, "warn-message")
	assert.Contains(t, out, "error-message")
}

func TestDefaultLogger_SetLevelEnablesLowerLevels(t *testing.T) {
	l, buf := newTestLogger()
	l.SetLevel(core.LogLevelDebug)
	ctx := context.Background()

	l.Debug(ctx, "debug-message")
	l.Info(ctx, "info-message")

	out := buf.String()
	assert.True(t, strings.Contains(out, "debug-message"))
	assert.True(t, strings.Contains(out, "info-message"))
}
