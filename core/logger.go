package core

import "context"

type Logger interface {
	Debug(context.Context, string, ...any)
	Info(context.Context, string, ...any)
	Error(context.Context, string, ...any)
	Warn(context.Context, string, ...any)
	SetLevel(level LogLevel)
	GetLevel() LogLevel
}

type CtxError struct {
	Err error
	Ctx context.Context
}

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)
