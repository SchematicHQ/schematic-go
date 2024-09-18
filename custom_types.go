package schematichq

import (
	"context"
	time "time"
)

type CacheProvider[T any] interface {
	Get(ctx context.Context, key string) (T, bool)
	Set(ctx context.Context, key string, val T, ttlOverride *time.Duration) error
}

// gomock generator doesn't like generic args so adding this for the SchematicClient interface
type CacheProviderBool = CacheProvider[bool]

type Logger interface {
	Printf(format string, args ...interface{})
}
