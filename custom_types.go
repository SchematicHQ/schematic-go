package schematichq

import (
	cache "github.com/schematichq/schematic-go/cache"
)

type BoolCacheProvider = cache.CacheProvider[bool]

type Logger interface {
	Printf(format string, args ...interface{})
}
