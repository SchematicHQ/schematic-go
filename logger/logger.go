package logger

import (
	"log"
	"os"

	schematicgo "github.com/schematichq/schematic-go"
)

type defaultLogger struct {
	logger *log.Logger
}

func (l defaultLogger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func NewDefaultLogger() schematicgo.Logger {
	return defaultLogger{logger: log.New(os.Stderr, "schematic ", log.LstdFlags)}
}
