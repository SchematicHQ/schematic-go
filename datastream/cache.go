package datastream

import (
	"github.com/go-redis/redis/v8"
	"github.com/schematichq/rulesengine"
	"github.com/schematichq/schematic-go/cache"
	"github.com/schematichq/schematic-go/core"
)

func getCacheProviders(opt *core.DatastreamOptions) (CompanyCacheProvider, UserCacheProvider) {
	if opt.CacheConfig == nil {
		return buildLocalCache(opt)
	}

	switch opt.CacheConfig.(type) {
	case *core.RedisCacheConfig:
		config := opt.CacheConfig.(*core.RedisCacheConfig)
		client := redis.NewClient(ToRedisOptions(config))
		return buildRedisCache(client, opt)
	case core.RedisCacheConfig:
		config := opt.CacheConfig.(core.RedisCacheConfig)
		client := redis.NewClient(ToRedisOptions(&config))
		return buildRedisCache(client, opt)
	case *core.RedisCacheClusterConfig:
		config := opt.CacheConfig.(*core.RedisCacheClusterConfig)
		client := redis.NewClusterClient(ToRedisClusterOptions(config))
		return buildRedisCache(client, opt)
	case core.RedisCacheClusterConfig:
		config := opt.CacheConfig.(core.RedisCacheClusterConfig)
		client := redis.NewClusterClient(ToRedisClusterOptions(&config))
		return buildRedisCache(client, opt)
	}

	return nil, nil
}

func buildRedisCache(client redis.UniversalClient, opt *core.DatastreamOptions) (CompanyCacheProvider, UserCacheProvider) {
	companyCacheProvider := cache.NewRedisCache[*rulesengine.Company](client, opt.CacheTTL)
	userCacheProvider := cache.NewRedisCache[*rulesengine.User](client, opt.CacheTTL)

	return companyCacheProvider, userCacheProvider
}

func buildLocalCache(opt *core.DatastreamOptions) (CompanyCacheProvider, UserCacheProvider) {
	companyCacheProvider := cache.NewLocalCache[*rulesengine.Company](defaultCacheSize, opt.CacheTTL)
	userCacheProvider := cache.NewLocalCache[*rulesengine.User](defaultCacheSize, opt.CacheTTL)

	return companyCacheProvider, userCacheProvider
}

func ToRedisOptions(config *core.RedisCacheConfig) *redis.Options {
	return &redis.Options{
		Addr:               config.Addr,
		DB:                 config.DB,
		Password:           config.Password,
		MaxRetries:         config.MaxRetries,
		MinRetryBackoff:    config.MinRetryBackoff,
		MaxRetryBackoff:    config.MaxRetryBackoff,
		DialTimeout:        config.DialTimeout,
		ReadTimeout:        config.ReadTimeout,
		WriteTimeout:       config.WriteTimeout,
		PoolSize:           config.PoolSize,
		MinIdleConns:       config.MinIdleConns,
		MaxConnAge:         config.MaxConnAge,
		PoolTimeout:        config.PoolTimeout,
		IdleTimeout:        config.IdleTimeout,
		IdleCheckFrequency: config.IdleCheckFrequency,
	}
}

func ToRedisClusterOptions(config *core.RedisCacheClusterConfig) *redis.ClusterOptions {
	return &redis.ClusterOptions{
		Addrs:              config.Addrs,
		MaxRedirects:       config.MaxRedirects,
		ReadOnly:           config.ReadOnly,
		RouteByLatency:     config.RouteByLatency,
		RouteRandomly:      config.RouteRandomly,
		Password:           config.Password,
		MaxRetries:         config.MaxRetries,
		MinRetryBackoff:    config.MinRetryBackoff,
		MaxRetryBackoff:    config.MaxRetryBackoff,
		DialTimeout:        config.DialTimeout,
		ReadTimeout:        config.ReadTimeout,
		WriteTimeout:       config.WriteTimeout,
		PoolSize:           config.PoolSize,
		MinIdleConns:       config.MinIdleConns,
		MaxConnAge:         config.MaxConnAge,
		PoolTimeout:        config.PoolTimeout,
		IdleTimeout:        config.IdleTimeout,
		IdleCheckFrequency: config.IdleCheckFrequency,
	}
}
