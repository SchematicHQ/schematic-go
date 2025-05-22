package logger

import (
	"context"
	"log"
	"os"

	"github.com/schematichq/schematic-go/core"
)

type defaultLogger struct {
	logger *log.Logger
	level  core.LogLevel
}

func (l *defaultLogger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

// Error logs an error message.
func (l *defaultLogger) Error(ctx context.Context, message string, args ...interface{}) {
	if l.level <= core.LogLevelError {
		l.logger.Printf("[ERROR] %s", append([]interface{}{message}, args...)...)
	}
}

// Warn logs a warning message.
func (l *defaultLogger) Warn(ctx context.Context, message string, args ...interface{}) {
	if l.level <= core.LogLevelWarn {
		l.logger.Printf("[WARN] %s", append([]interface{}{message}, args...)...)
	}
}

func (l *defaultLogger) Info(ctx context.Context, message string, args ...interface{}) {
	if l.level <= core.LogLevelInfo {
		l.logger.Printf("[INFO] %s", append([]interface{}{message}, args...)...)
	}
}

func (l *defaultLogger) Debug(ctx context.Context, message string, args ...interface{}) {
	if l.level <= core.LogLevelDebug {
		l.logger.Printf("[DEBUG] %s", append([]interface{}{message}, args...)...)
	}
}

func (l *defaultLogger) SetLevel(level core.LogLevel) {
	l.level = level
}

func (l *defaultLogger) GetLevel() core.LogLevel {
	return l.level
}

func NewDefaultLogger() core.Logger {
	return &defaultLogger{
		logger: log.New(os.Stderr, "schematic ", log.LstdFlags),
		level:  core.LogLevelInfo, // Default level
	}
}
