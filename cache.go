package schematic

import (
	"context"
	"math"
	"sync"
	"time"
)

const defaultCacheSize = 10 * 1024 // 10KB
const defaultCacheTTL = 5 * time.Second

type CacheProvider[T any] interface {
	Get(ctx context.Context, key string) (T, bool)
	Set(ctx context.Context, key string, val T, ttlOverride *time.Duration) error
}

type cachedItem[T any] struct {
	value         T
	accessCounter int64
	size          int
	expiration    time.Time
}

type localCache[T any] struct {
	cache map[string]cachedItem[T]

	// Max bytes to store in cache
	maxSize int

	// Current size of the cache in bytes
	currentSize int

	// Access counter for LRU eviction
	accessCounter int64

	// Default TTL for cached items
	ttl time.Duration

	// Function to calculate the size of a cached item
	sizeFunc func(val T) int

	// Mutex to protect concurrent access to the cache
	mu sync.RWMutex
}

func newDefaultCache[T any](sizeFunc func(val T) int) *localCache[T] {
	return newLocalCache(defaultCacheSize, defaultCacheTTL, sizeFunc)
}

func newLocalCache[T any](maxSize int, ttl time.Duration, sizeFunc func(val T) int) *localCache[T] {
	return &localCache[T]{
		cache:    make(map[string]cachedItem[T]),
		maxSize:  maxSize,
		ttl:      ttl,
		sizeFunc: sizeFunc,
	}
}

func (c *localCache[T]) Get(ctx context.Context, key string) (T, bool) {
	var empty T
	if c == nil {
		return empty, false
	}

	if c.maxSize == 0 {
		return empty, false
	}

	c.mu.RLock()
	item, ok := c.cache[key]
	c.mu.RUnlock()
	if !ok {
		return empty, false
	}

	// Check if the item has expired
	if time.Now().After(item.expiration) {
		c.mu.Lock()
		// Re-check the expiration under write lock
		item, ok := c.cache[key]
		if ok && time.Now().After(item.expiration) {
			c.currentSize -= item.size
			delete(c.cache, key)
		}
		c.mu.Unlock()
		return empty, false
	}

	// Update the access counter for LRU eviction
	c.mu.Lock()
	c.accessCounter++
	item.accessCounter = c.accessCounter
	c.cache[key] = item
	c.mu.Unlock()

	return item.value, true
}

func (c *localCache[T]) Set(ctx context.Context, key string, val T, ttlOverride *time.Duration) error {
	// if maxSize = 0, caching is disabled
	if c == nil || c.maxSize == 0 {
		return nil
	}

	ttl := c.ttl
	if ttlOverride != nil {
		ttl = *ttlOverride
	}

	size := c.sizeFunc(val)

	c.mu.Lock()
	defer c.mu.Unlock()

	// Check if the key already exists in the cache
	if item, ok := c.cache[key]; ok {
		// Update the existing item
		c.currentSize -= item.size
		c.currentSize += size
		c.accessCounter++
		c.cache[key] = cachedItem[T]{
			value:         val,
			accessCounter: c.accessCounter,
			size:          size,
			expiration:    time.Now().Add(ttl),
		}
		return nil
	}

	// Evict expired items
	for key, item := range c.cache {
		if time.Now().After(item.expiration) {
			c.currentSize -= item.size
			delete(c.cache, key)
		}
	}

	// Evict records if the cache size exceeds the max size
	for c.currentSize+size > c.maxSize {
		oldestKey := ""
		oldestAccessCounter := int64(math.MaxInt64)

		for k, v := range c.cache {
			if v.accessCounter < oldestAccessCounter {
				oldestKey = k
				oldestAccessCounter = v.accessCounter
			}
		}

		if oldestKey != "" {
			c.currentSize -= c.cache[oldestKey].size
			delete(c.cache, oldestKey)
		}
	}

	// Add the new item to the cache
	c.accessCounter++
	c.cache[key] = cachedItem[T]{
		value:         val,
		accessCounter: c.accessCounter,
		size:          size,
		expiration:    time.Now().Add(ttl),
	}
	c.currentSize += size

	return nil
}

func boolSizeFunc(val bool) int {
	return 1
}
