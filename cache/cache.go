package cache

import (
	"context"
	"time"
)

type CacheProvider[T any] interface {
	Get(ctx context.Context, key string) (T, bool)
	Set(ctx context.Context, key string, val T, ttlOverride *time.Duration) error
	Delete(ctx context.Context, key string) error
	// DeleteMissing removes every entry under scanPattern whose key is not in
	// keysToKeep. scanPattern scopes the operation to the caller's own slice of
	// the keyspace — this is required for backends like Redis where the cache
	// is shared with other data. Backends that own their keyspace (e.g. the
	// in-process local cache) may ignore scanPattern. An empty scanPattern is
	// treated as a no-op on scan-based backends to avoid wiping unrelated keys.
	DeleteMissing(ctx context.Context, keysToKeep []string, scanPattern string)
}

// KeepTTLSetter is implemented by providers that can overwrite a value while
// keeping the entry's current expiry. Use it when another process owns the
// key's lifetime and a write from this SDK must not change it.
type KeepTTLSetter[T any] interface {
	SetKeepTTL(ctx context.Context, key string, val T) error
}

const defaultCacheSize = 1000 // 1000 records
const defaultCacheTTL = 5 * time.Second

type cachedItem[T any] struct {
	key        string
	value      T
	expiration time.Time
}

func NewDefaultCache[T any]() *localCache[T] {
	return NewLocalCache[T](defaultCacheSize, defaultCacheTTL)
}
