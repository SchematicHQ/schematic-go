package schematic

import (
	"context"
	"math"
	"time"
)

type CacheProvider interface {
	Get(ctx *context.Context, key string) ([]byte, bool)
	Set(ctx *context.Context, key string, val []byte)
}

type cachedItem struct {
	value         []byte
	accessCounter int64
	size          int
	expiration    time.Time
}

type localCache struct {
	cache map[string]cachedItem

	// Max bytes to store in cache
	maxSize int

	// Current size of the cache in bytes
	currentSize int

	// Access counter for LRU eviction
	accessCounter int64

	// Default TTL for cached items
	ttl time.Duration
}

func newDefaultCache() *localCache {
	return newLocalCache(defaultCacheSize, defaultCacheTTL)
}

func newLocalCache(maxSize int, ttl time.Duration) *localCache {
	return &localCache{
		cache:   make(map[string]cachedItem),
		maxSize: maxSize,
		ttl:     ttl,
	}
}

func (c *localCache) Get(ctx *context.Context, key string) ([]byte, bool) {
	if c.maxSize == 0 {
		return nil, false
	}

	item, ok := c.cache[key]
	if !ok {
		return nil, false
	}

	// Check if the item has expired
	if time.Now().After(item.expiration) {
		c.currentSize -= item.size
		delete(c.cache, key)
		return nil, false
	}

	// Update the access counter for LRU eviction
	c.accessCounter++
	item.accessCounter = c.accessCounter
	c.cache[key] = item

	return item.value, true
}

func (c *localCache) Set(ctx *context.Context, key string, val []byte) {
	size := len(val)

	// Check if the key already exists in the cache
	if item, ok := c.cache[key]; ok {
		// Update the existing item
		c.currentSize -= item.size
		c.currentSize += size
		c.accessCounter++
		c.cache[key] = cachedItem{
			value:         val,
			accessCounter: c.accessCounter,
			size:          size,
			expiration:    time.Now().Add(c.ttl),
		}
		return
	}

	// Evict expired items
	for key, item := range c.cache {
		if time.Now().After(item.expiration) {
			c.currentSize -= item.size
			delete(c.cache, key)
		}
	}

	// Evict older records if the cache size exceeds the max size
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
	c.cache[key] = cachedItem{
		value:         val,
		accessCounter: c.accessCounter,
		size:          size,
		expiration:    time.Now().Add(c.ttl),
	}
	c.currentSize += size
}
