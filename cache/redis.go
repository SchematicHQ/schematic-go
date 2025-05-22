package cache

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisCache[T any] struct {
	client redis.UniversalClient // Use UniversalClient to support both single instance and cluster
	ttl    time.Duration
}

func NewRedisCache[T any](client redis.UniversalClient, ttl time.Duration) *redisCache[T] {
	return &redisCache[T]{
		client: client,
		ttl:    ttl,
	}
}

func (r *redisCache[T]) Get(ctx context.Context, key string) (T, bool) {
	var empty T

	// Fetch the value from Redis
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		// Key does not exist
		return empty, false
	} else if err != nil {
		// Handle other Redis errors
		return empty, false
	}

	// Deserialize the value
	var result T
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return empty, false
	}

	return result, true
}

func (r *redisCache[T]) Set(ctx context.Context, key string, val T, ttlOverride *time.Duration) error {
	// Serialize the value
	data, err := json.Marshal(val)
	if err != nil {
		return err
	}

	// Determine the TTL
	ttl := r.ttl
	if ttlOverride != nil {
		ttl = *ttlOverride
	}

	// If TTL is 0 or negative, use no expiration
	if ttl <= 0 {
		return r.client.Set(ctx, key, val, 0).Err()
	}

	// Store the value in Redis
	return r.client.Set(ctx, key, data, ttl).Err()
}

func (r *redisCache[T]) Delete(ctx context.Context, key string) error {
	// Delete the key from Redis
	return r.client.Del(ctx, key).Err()
}

// DeleteMissing deletes all keys with a specific prefix that are not in the provided array of keys.
func (r *redisCache[T]) DeleteMissing(ctx context.Context, keys []string) {
	// If no keys provided, nothing to compare against
	if len(keys) == 0 {
		return
	}

	// Create a map for O(1) lookup
	keysMap := make(map[string]struct{}, len(keys))
	for _, k := range keys {
		keysMap[k] = struct{}{}
	}

	// Extract the prefix from the first key to scan for similar keys
	// Assuming keys follow a pattern like "prefix:*"
	var prefix string
	if len(keys) > 0 {
		parts := strings.Split(keys[0], ":")
		if len(parts) > 1 {
			prefix = parts[0] + ":"
		}
	}

	if prefix == "" {
		return
	}

	// Use SCAN to iterate over all keys with the prefix
	var cursor uint64
	var err error
	var keysToDelete []string

	for {
		var batch []string
		batch, cursor, err = r.client.Scan(ctx, cursor, prefix+"*", 100).Result()
		if err != nil {
			return
		}

		// Add keys not in the provided array to the delete list
		for _, existingKey := range batch {
			if _, exists := keysMap[existingKey]; !exists {
				keysToDelete = append(keysToDelete, existingKey)
			}
		}

		// If cursor is 0, we've completed scanning
		if cursor == 0 {
			break
		}
	}

	// Delete keys in batches to avoid overloading Redis
	if len(keysToDelete) > 0 {
		const batchSize = 100
		for i := 0; i < len(keysToDelete); i += batchSize {
			end := i + batchSize
			if end > len(keysToDelete) {
				end = len(keysToDelete)
			}
			batch := keysToDelete[i:end]
			r.client.Del(ctx, batch...)
		}
	}
}
