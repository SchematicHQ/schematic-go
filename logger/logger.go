package logger

import (
	"context"
	"log"
	"os"

	"github.com/schematichq/schematic-go/core"
)

type defaultLogger struct {
	logger *log.Logger
}

func (l *defaultLogger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

// Info logs an informational message.
func (l *defaultLogger) Info(ctx context.Context, message string, args ...interface{}) {
	l.logger.Printf("[INFO] %s", append([]interface{}{message}, args...)...)
}

// Error logs an error message.
func (l *defaultLogger) Error(ctx context.Context, message string, args ...interface{}) {
	l.logger.Printf("[ERROR] %s", append([]interface{}{message}, args...)...)
}

// Warn logs a warning message.
func (l *defaultLogger) Warn(ctx context.Context, message string, args ...interface{}) {
	l.logger.Printf("[WARN] %s", append([]interface{}{message}, args...)...)
}

// Debug logs a debug message.
func (l *defaultLogger) Debug(ctx context.Context, message string, args ...interface{}) {
	l.logger.Printf("[DEBUG] %s", append([]interface{}{message}, args...)...)
}

func NewDefaultLogger() core.Logger {
	return &defaultLogger{logger: log.New(os.Stderr, "schematic ", log.LstdFlags)}
}
