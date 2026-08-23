package datastream

import (
	"crypto/tls"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/schematichq/schematic-go/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToRedisOptionsTLSConfig(t *testing.T) {
	t.Run("nil TLSConfig stays nil", func(t *testing.T) {
		opts := ToRedisOptions(&core.RedisCacheConfig{Addr: "localhost:6379"})

		assert.Nil(t, opts.TLSConfig)
		assert.Equal(t, "localhost:6379", opts.Addr)
	})

	t.Run("TLSConfig is propagated", func(t *testing.T) {
		tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}
		opts := ToRedisOptions(&core.RedisCacheConfig{Addr: "localhost:6379", TLSConfig: tlsConfig})

		require.NotNil(t, opts.TLSConfig)
		assert.Same(t, tlsConfig, opts.TLSConfig)
		assert.Equal(t, uint16(tls.VersionTLS12), opts.TLSConfig.MinVersion)
	})
}

func TestToRedisClusterOptionsTLSConfig(t *testing.T) {
	t.Run("nil TLSConfig stays nil", func(t *testing.T) {
		opts := ToRedisClusterOptions(&core.RedisCacheClusterConfig{Addrs: []string{"localhost:6379"}})

		assert.Nil(t, opts.TLSConfig)
	})

	t.Run("TLSConfig is propagated", func(t *testing.T) {
		tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}
		opts := ToRedisClusterOptions(&core.RedisCacheClusterConfig{
			Addrs:     []string{"localhost:6379"},
			TLSConfig: tlsConfig,
		})

		require.NotNil(t, opts.TLSConfig)
		assert.Same(t, tlsConfig, opts.TLSConfig)
	})
}

func TestToRedisOptionsCredentialsStillMapped(t *testing.T) {
	opts := ToRedisOptions(&core.RedisCacheConfig{
		Addr:     "localhost:6379",
		Username: "user",
		Password: "pass",
		DB:       3,
	})

	assert.Equal(t, "user", opts.Username)
	assert.Equal(t, "pass", opts.Password)
	assert.Equal(t, 3, opts.DB)
}

func TestToRedisOptionsAddrAsURL(t *testing.T) {
	t.Run("redis scheme is reduced to host:port", func(t *testing.T) {
		opts := ToRedisOptions(&core.RedisCacheConfig{Addr: "redis://localhost:6379"})

		assert.Equal(t, "localhost:6379", opts.Addr)
		assert.Nil(t, opts.TLSConfig)
	})

	t.Run("rediss scheme implies TLS", func(t *testing.T) {
		opts := ToRedisOptions(&core.RedisCacheConfig{Addr: "rediss://cache.example.com:6380"})

		assert.Equal(t, "cache.example.com:6380", opts.Addr)
		require.NotNil(t, opts.TLSConfig)
		assert.Equal(t, "cache.example.com", opts.TLSConfig.ServerName)
	})

	t.Run("URL credentials and db are used when config omits them", func(t *testing.T) {
		opts := ToRedisOptions(&core.RedisCacheConfig{Addr: "redis://user:pass@localhost:6379/4"})

		assert.Equal(t, "localhost:6379", opts.Addr)
		assert.Equal(t, "user", opts.Username)
		assert.Equal(t, "pass", opts.Password)
		assert.Equal(t, 4, opts.DB)
	})

	t.Run("explicit config wins over URL", func(t *testing.T) {
		explicitTLS := &tls.Config{MinVersion: tls.VersionTLS13}
		opts := ToRedisOptions(&core.RedisCacheConfig{
			Addr:      "rediss://urluser:urlpass@cache.example.com:6380/4",
			Username:  "configuser",
			Password:  "configpass",
			DB:        1,
			TLSConfig: explicitTLS,
		})

		assert.Equal(t, "cache.example.com:6380", opts.Addr)
		assert.Equal(t, "configuser", opts.Username)
		assert.Equal(t, "configpass", opts.Password)
		assert.Equal(t, 1, opts.DB)
		assert.Same(t, explicitTLS, opts.TLSConfig)
	})

	t.Run("unparseable URL is left alone", func(t *testing.T) {
		opts := ToRedisOptions(&core.RedisCacheConfig{Addr: "memcached://localhost:11211"})

		assert.Equal(t, "memcached://localhost:11211", opts.Addr)
	})
}

func TestBuildRedisClientUsesSuppliedClient(t *testing.T) {
	supplied := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	t.Cleanup(func() { _ = supplied.Close() })

	t.Run("by value", func(t *testing.T) {
		client := buildRedisClient(&core.DatastreamOptions{
			CacheConfig: core.RedisClientConfig{Client: supplied},
		})

		assert.Same(t, supplied, client)
	})

	t.Run("by pointer", func(t *testing.T) {
		client := buildRedisClient(&core.DatastreamOptions{
			CacheConfig: &core.RedisClientConfig{Client: supplied},
		})

		assert.Same(t, supplied, client)
	})

	t.Run("via WithRedisClient", func(t *testing.T) {
		config := core.WithRedisClient(supplied)
		client := buildRedisClient(&core.DatastreamOptions{CacheConfig: config})

		assert.Same(t, supplied, client)
	})
}

func TestBuildRedisClientWithoutConfig(t *testing.T) {
	assert.Nil(t, buildRedisClient(nil))
	assert.Nil(t, buildRedisClient(&core.DatastreamOptions{}))
}
