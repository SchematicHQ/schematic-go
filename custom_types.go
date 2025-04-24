package schematichq

import (
	"context"

	cache "github.com/schematichq/schematic-go/cache"
)

type BoolCacheProvider = cache.CacheProvider[bool]

type Logger interface {
	Debug(context.Context, string, ...any)
	Info(context.Context, string, ...any)
	Error(context.Context, string, ...any)
	Warn(context.Context, string, ...any)
}
