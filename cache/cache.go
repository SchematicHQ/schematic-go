package cache

import (
	"context"
	"time"
)

type CacheProvider[T any] interface {
	Get(ctx context.Context, key string) (T, bool)
	Set(ctx context.Context, key string, val T, ttlOverride *time.Duration) error
	Delete(ctx context.Context, key string) error
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
