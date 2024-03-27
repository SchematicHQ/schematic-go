package schematic

import (
	"log"
	"os"
)

type Logger interface {
	Printf(format string, args ...interface{})
}

type defaultLogger struct {
	logger *log.Logger
}

func (l defaultLogger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func newDefaultLogger() Logger {
	return defaultLogger{logger: log.New(os.Stderr, "schematic ", log.LstdFlags)}
}
