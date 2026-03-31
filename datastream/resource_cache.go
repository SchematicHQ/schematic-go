package datastream

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/schematichq/schematic-go/cache"
	"github.com/schematichq/schematic-go/core"
)

// resourceCache encapsulates the two-level cache pattern (primary + lookup) shared
// by company and user caches. The primary cache stores the full entity keyed by ID,
// while the lookup cache maps each resource key to the entity's ID.
type resourceCache[T any] struct {
	primaryCache cache.CacheProvider[T]
	lookupCache  cache.CacheProvider[string]
	keyPrefix    string
	getID        func(T) string
	getKeys      func(T) map[string]string
	cacheVersion func() string
}

func (rc *resourceCache[T]) idCacheKey(id string) string {
	return fmt.Sprintf("%s:%s:%s:%s", cacheKeyPrefix, rc.keyPrefix, rc.cacheVersion(), id)
}

func (rc *resourceCache[T]) lookupCacheKey(key, value string) string {
	return fmt.Sprintf("%s:%s:%s:%s:%s", cacheKeyPrefix, rc.keyPrefix, rc.cacheVersion(), strings.ToLower(key), strings.ToLower(value))
}

// cacheForKeys writes the entity to the ID-based primary key and writes lookup
// entries for each of its keys. It also cleans up stale lookup keys by diffing
// against the previously cached version of the entity.
func (rc *resourceCache[T]) cacheForKeys(ctx context.Context, entity T, ttl time.Duration) map[string]error {
	id := rc.getID(entity)
	keys := rc.getKeys(entity)
	if len(keys) == 0 {
		return nil
	}

	cacheResults := make(map[string]error)

	// Clean up stale lookup keys by diffing against existing cached entity
	idKey := rc.idCacheKey(id)
	if existing, found := rc.primaryCache.Get(ctx, idKey); found {
		for key, value := range rc.getKeys(existing) {
			newValue, exists := keys[key]
			if !exists || newValue != value {
				oldLookupKey := rc.lookupCacheKey(key, value)
				cacheResults[oldLookupKey] = rc.lookupCache.Delete(ctx, oldLookupKey)
			}
		}
	}

	// Write full entity to ID-based primary key
	cacheResults[idKey] = rc.primaryCache.Set(ctx, idKey, entity, &ttl)

	// Write entity ID string to each versioned lookup key
	for key, value := range keys {
		lookupKey := rc.lookupCacheKey(key, value)
		cacheResults[lookupKey] = rc.lookupCache.Set(ctx, lookupKey, id, &ttl)
	}

	return cacheResults
}

// deleteEntity fetches the cached entity by ID, deletes all lookup keys (both
// from the cached version and the incoming entity), then deletes the ID key.
func (rc *resourceCache[T]) deleteEntity(ctx context.Context, entity T, logger core.Logger) {
	id := rc.getID(entity)
	idKey := rc.idCacheKey(id)

	// Delete lookup keys from the cached version
	if existing, found := rc.primaryCache.Get(ctx, idKey); found {
		for key, value := range rc.getKeys(existing) {
			lk := rc.lookupCacheKey(key, value)
			if err := rc.lookupCache.Delete(ctx, lk); err != nil {
				logger.Warn(ctx, fmt.Sprintf("Failed to delete %s lookup key '%s': %v", rc.keyPrefix, lk, err))
			}
		}
	}

	// Also delete lookup keys from the incoming entity (in case they differ from cached)
	for key, value := range rc.getKeys(entity) {
		lk := rc.lookupCacheKey(key, value)
		if err := rc.lookupCache.Delete(ctx, lk); err != nil {
			logger.Warn(ctx, fmt.Sprintf("Failed to delete %s lookup key '%s': %v", rc.keyPrefix, lk, err))
		}
	}

	// Delete the ID-based primary key
	if err := rc.primaryCache.Delete(ctx, idKey); err != nil {
		logger.Warn(ctx, fmt.Sprintf("Failed to delete %s ID key '%s': %v", rc.keyPrefix, idKey, err))
	}
}

// getFromCache performs a two-step lookup: for each provided key it looks up
// the entity ID in the lookup cache, then fetches the full entity from the
// primary cache.
func (rc *resourceCache[T]) getFromCache(ctx context.Context, keys map[string]string) (T, bool) {
	for key, value := range keys {
		lookupKey := rc.lookupCacheKey(key, value)
		if entityID, found := rc.lookupCache.Get(ctx, lookupKey); found && entityID != "" {
			idKey := rc.idCacheKey(entityID)
			if entity, exists := rc.primaryCache.Get(ctx, idKey); exists {
				return entity, true
			}
		}
	}

	var zero T
	return zero, false
}
