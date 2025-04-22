package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
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

	// Store the value in Redis
	return r.client.Set(ctx, key, data, ttl).Err()
}
