package cache

import (
	"container/list"
	"context"
	"runtime"
	"slices"
	"sync"
	"time"
)

type localCache[T any] struct {
	cache        map[string]*list.Element
	lruList      *list.List
	maxSize      int
	ttl          time.Duration
	mu           sync.RWMutex
	cleanupTimer *time.Ticker
	stopCleanup  chan struct{}
	cleanupMu    sync.Mutex // Protects cleanupTimer and stopCleanup
}

func NewLocalCache[T any](maxSize int, ttl time.Duration) *localCache[T] {
	cache := &localCache[T]{
		cache:   make(map[string]*list.Element),
		lruList: list.New(),
		maxSize: maxSize,
		ttl:     ttl,
	}

	// Only start periodic cleanup if TTL is set AND maxSize > 0
	if ttl > 0 && maxSize > 0 {
		// Run cleanup at 1/4 of TTL duration, but not less than 1 second
		// and not more than 1 minute
		cleanupInterval := ttl / 4
		if cleanupInterval < time.Second {
			cleanupInterval = time.Second
		} else if cleanupInterval > time.Minute {
			cleanupInterval = time.Minute
		}

		// Initialize cleanup resources under lock
		cache.cleanupMu.Lock()
		cache.stopCleanup = make(chan struct{})
		cache.cleanupTimer = time.NewTicker(cleanupInterval)
		cache.cleanupMu.Unlock()

		// Start cleanup routine
		go cache.startCleanupRoutine()

		// Set up finalizer to ensure cleanup goroutine is stopped if cache is garbage collected
		runtime.SetFinalizer(cache, func(c *localCache[T]) {
			if c != nil {
				c.Stop()
			}
		})
	}

	return cache
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

	// We no longer need to check expiration here as the background cleanup will handle it
	// However, for extra safety, we still check to avoid returning expired data
	// This is just a quick double-check in case the cleanup routine hasn't run yet
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
	if c == nil || c.maxSize == 0 || c.cache == nil {
		return
	}

	// Normalize nil slice to empty slice to avoid issues
	if keys == nil {
		keys = []string{}
	}

	// Create a list of keys to delete to avoid modifying while iterating
	var keysToDelete []string
	c.mu.RLock()
	for key := range c.cache {
		found := slices.Contains(keys, key) // Check if the key is in the keys slice
		if !found {
			keysToDelete = append(keysToDelete, key)
		}
	}
	c.mu.RUnlock()

	// Delete the keys that weren't found in the provided list
	for _, key := range keysToDelete {
		c.Delete(ctx, key) // Use the provided context
	}
}

// startCleanupRoutine runs a background goroutine that periodically removes expired entries
func (c *localCache[T]) startCleanupRoutine() {
	// Safety check for nil receiver
	if c == nil {
		return
	}

	// Use local references to avoid race conditions
	c.cleanupMu.Lock()
	timer := c.cleanupTimer
	stopChan := c.stopCleanup
	c.cleanupMu.Unlock()

	// Safety check for nil channels/timers
	if timer == nil || stopChan == nil {
		return
	}

	for {
		select {
		case <-timer.C:
			c.removeExpiredItems()
		case <-stopChan:
			return
		}
	}
}

// removeExpiredItems scans the cache and removes expired entries
func (c *localCache[T]) removeExpiredItems() {
	// Safety check for nil receiver
	if c == nil {
		return
	}

	now := time.Now()

	// Use more granular locking strategy to minimize contention:
	// 1. First identify expired items with a read lock
	// 2. Then acquire write lock only for actual deletion
	var expiredElements []*list.Element
	var expiredKeys []string

	// Step 1: Identify expired items with read lock
	c.mu.RLock()
	for e := c.lruList.Front(); e != nil; e = e.Next() {
		item := e.Value.(*cachedItem[T])
		if !item.expiration.IsZero() && now.After(item.expiration) {
			expiredElements = append(expiredElements, e)
			expiredKeys = append(expiredKeys, item.key)
		}
	}
	c.mu.RUnlock()

	// If nothing to delete, avoid acquiring write lock
	if len(expiredElements) == 0 {
		return
	}

	// Step 2: Delete with write lock
	c.mu.Lock()
	defer c.mu.Unlock()

	// Delete the expired items
	for i, e := range expiredElements {
		// Verify the element is still in the list and still expired
		// (it might have been updated or removed between our read and write locks)
		if elem, ok := c.cache[expiredKeys[i]]; ok && elem == e {
			item := e.Value.(*cachedItem[T])
			if !item.expiration.IsZero() && now.After(item.expiration) {
				c.lruList.Remove(e)
				delete(c.cache, item.key)
			}
		}
	}
}

// Stop shuts down the cleanup goroutine. It is safe to call Stop multiple times.
func (c *localCache[T]) Stop() {
	if c == nil {
		return
	}

	// Use a separate mutex for cleanup resources to avoid deadlocks
	// and reduce contention on the main cache mutex
	c.cleanupMu.Lock()
	defer c.cleanupMu.Unlock()

	if c.cleanupTimer != nil {
		c.cleanupTimer.Stop()

		if c.stopCleanup != nil {
			// Use non-blocking send/close to avoid panic if channel is already closed
			select {
			case <-c.stopCleanup:
				// Channel already closed, do nothing
			default:
				close(c.stopCleanup)
			}
		}

		c.cleanupTimer = nil

		// Remove the finalizer since we've manually cleaned up
		runtime.SetFinalizer(c, nil)
	}
}

// Ensure localCache implements CacheProvider
var _ CacheProvider[any] = (*localCache[any])(nil)
