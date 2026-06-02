package client_test

import (
	"context"
	"testing"

	schematicclient "github.com/schematichq/schematic-go/client"
	"github.com/schematichq/schematic-go/core"
	option "github.com/schematichq/schematic-go/option"
	"github.com/stretchr/testify/assert"
)

// customLevelLogger is a consumer-supplied logger that happens to expose a
// SetLevel method. The SDK must NOT override its level when it is provided
// via option.WithLogger.
type customLevelLogger struct {
	level core.LogLevel
}

func (l *customLevelLogger) Debug(context.Context, string, ...any) {}
func (l *customLevelLogger) Info(context.Context, string, ...any)  {}
func (l *customLevelLogger) Warn(context.Context, string, ...any)  {}
func (l *customLevelLogger) Error(context.Context, string, ...any) {}
func (l *customLevelLogger) SetLevel(level core.LogLevel)          { l.level = level }

// TestWithLogLevelIsExportedFromOptionPackage ensures the public option package
// re-exports WithLogLevel and that it wires through to RequestOptions.LogLevel.
func TestWithLogLevelIsExportedFromOptionPackage(t *testing.T) {
	options := core.NewRequestOptions(option.WithLogLevel(option.LogLevelDebug))
	assert.Equal(t, core.LogLevelDebug, options.LogLevel)
}

// TestCustomLoggerLevelNotOverridden ensures that when a consumer provides their
// own logger, the SDK leaves its level untouched even if WithLogLevel is set.
func TestCustomLoggerLevelNotOverridden(t *testing.T) {
	custom := &customLevelLogger{level: core.LogLevelError}

	client := schematicclient.NewSchematicClient(
		option.WithAPIKey("test-api-key"),
		option.WithLogger(custom),
		option.WithLogLevel(option.LogLevelDebug),
	)
	defer client.Close()

	// The consumer's logger level is the source of truth; WithLogLevel must be
	// ignored when a custom logger is provided.
	assert.Equal(t, core.LogLevelError, custom.level)
}
