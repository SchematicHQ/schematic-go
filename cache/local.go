package cache

import (
	"container/list"
	"context"
	"slices"
	"sync"
	"time"
)

type localCache[T any] struct {
	cache   map[string]*list.Element
	lruList *list.List
	maxSize int
	ttl     time.Duration
	mu      sync.RWMutex
}

func NewLocalCache[T any](maxSize int, ttl time.Duration) *localCache[T] {
	return &localCache[T]{
		cache:   make(map[string]*list.Element),
		lruList: list.New(),
		maxSize: maxSize,
		ttl:     ttl,
	}
}

func (c *localCache[T]) Get(ctx context.Context, key string) (T, bool) {
	var empty T
	if c == nil || c.maxSize == 0 {
		return empty, false
	}

	c.mu.RLock()
	element, exists := c.cache[key]
	c.mu.RUnlock()

	if !exists {
		return empty, false
	}

	item := element.Value.(*cachedItem[T])

	// Check if the item has expired
	if !item.expiration.IsZero() && time.Now().After(item.expiration) {
		c.mu.Lock()
		c.lruList.Remove(element)
		delete(c.cache, key)
		c.mu.Unlock()
		return empty, false
	}

	// Move the accessed item to the front of the list
	c.mu.Lock()
	c.lruList.MoveToFront(element)
	c.mu.Unlock()

	return item.value, true
}

func (c *localCache[T]) Set(ctx context.Context, key string, val T, ttlOverride *time.Duration) error {
	if c == nil || c.maxSize == 0 {
		return nil
	}

	ttl := c.ttl
	if ttlOverride != nil {
		ttl = *ttlOverride
	}

	// If TTL is 0 or negative, set no expiration
	var expiration time.Time
	if ttl > 0 {
		expiration = time.Now().Add(ttl)
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	// If the key already exists, update it
	if element, exists := c.cache[key]; exists {
		c.lruList.MoveToFront(element)
		item := element.Value.(*cachedItem[T])
		item.value = val
		item.expiration = expiration
		return nil
	}

	// If we're at capacity, remove the least recently used item
	if c.lruList.Len() >= c.maxSize {
		oldest := c.lruList.Back()
		if oldest != nil {
			c.lruList.Remove(oldest)
			delete(c.cache, oldest.Value.(*cachedItem[T]).key)
		}
	}

	// Add the new item
	item := &cachedItem[T]{
		key:        key,
		value:      val,
		expiration: expiration,
	}
	element := c.lruList.PushFront(item)
	c.cache[key] = element

	return nil
}

// Delete removes an item from the cache by its key.
func (c *localCache[T]) Delete(ctx context.Context, key string) error {
	if c == nil || c.maxSize == 0 {
		return nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if element, exists := c.cache[key]; exists {
		c.lruList.Remove(element)
		delete(c.cache, key)
	}

	return nil
}

func (c *localCache[T]) DeleteMissing(ctx context.Context, keys []string) {
	if c == nil || c.maxSize == 0 {
		return
	}

	for key := range c.cache {
		found := slices.Contains(keys, key) // Check if the key is in the keys slice
		if found {
			continue // Skip keys that are in the keys slice
		}
		c.Delete(context.Background(), key) // Delete the key if it doesn't match any in the keys slice
	}
}
