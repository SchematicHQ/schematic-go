package cache

import (
	"context"
	"encoding/json"
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
		return r.client.Set(ctx, key, data, 0).Err()
	}

	// Store the value in Redis
	return r.client.Set(ctx, key, data, ttl).Err()
}

func (r *redisCache[T]) Delete(ctx context.Context, key string) error {
	// Delete the key from Redis
	return r.client.Del(ctx, key).Err()
}

// DeleteMissing scans Redis under scanPattern and deletes every matching key
// that is not present in keysToKeep. An empty scanPattern is a no-op: the
// caller must provide an explicit scope so unrelated sibling caches aren't
// wiped from a shared Redis keyspace.
func (r *redisCache[T]) DeleteMissing(ctx context.Context, keysToKeep []string, scanPattern string) {
	if scanPattern == "" {
		return
	}
	if len(keysToKeep) == 0 {
		return
	}

	keepSet := make(map[string]struct{}, len(keysToKeep))
	for _, k := range keysToKeep {
		keepSet[k] = struct{}{}
	}

	var cursor uint64
	var keysToDelete []string
	for {
		var batch []string
		var err error
		batch, cursor, err = r.client.Scan(ctx, cursor, scanPattern, 100).Result()
		if err != nil {
			return
		}

		for _, existing := range batch {
			if _, keep := keepSet[existing]; !keep {
				keysToDelete = append(keysToDelete, existing)
			}
		}

		if cursor == 0 {
			break
		}
	}

	if len(keysToDelete) > 0 {
		const batchSize = 100
		for i := 0; i < len(keysToDelete); i += batchSize {
			end := i + batchSize
			if end > len(keysToDelete) {
				end = len(keysToDelete)
			}
			r.client.Del(ctx, keysToDelete[i:end]...)
		}
	}
}
