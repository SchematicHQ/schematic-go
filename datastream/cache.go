package datastream

import (
	"github.com/redis/go-redis/v9"
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
		Network:               config.Network,
		Addr:                  config.Addr,
		ClientName:            config.ClientName,
		Protocol:              config.Protocol,
		Username:              config.Username,
		Password:              config.Password,
		DB:                    config.DB,
		MaxRetries:            config.MaxRetries,
		MinRetryBackoff:       config.MinRetryBackoff,
		MaxRetryBackoff:       config.MaxRetryBackoff,
		DialTimeout:           config.DialTimeout,
		ReadTimeout:           config.ReadTimeout,
		WriteTimeout:          config.WriteTimeout,
		ContextTimeoutEnabled: config.ContextTimeoutEnabled,
		PoolFIFO:              config.PoolFIFO,
		PoolSize:              config.PoolSize,
		PoolTimeout:           config.PoolTimeout,
		MinIdleConns:          config.MinIdleConns,
		MaxIdleConns:          config.MaxIdleConns,
		MaxActiveConns:        config.MaxActiveConns,
		ConnMaxIdleTime:       config.ConnMaxIdleTime,
		ConnMaxLifetime:       config.ConnMaxLifetime,
		DisableIndentity:      config.DisableIndentity,
		DisableIdentity:       config.DisableIdentity,
		IdentitySuffix:        config.IdentitySuffix,
		UnstableResp3:         config.UnstableResp3,
	}
}

func ToRedisClusterOptions(config *core.RedisCacheClusterConfig) *redis.ClusterOptions {
	return &redis.ClusterOptions{
		Addrs:                 config.Addrs,
		MaxRedirects:          config.MaxRedirects,
		RouteByLatency:        config.RouteByLatency,
		RouteRandomly:         config.RouteRandomly,
		Protocol:              config.Protocol,
		Username:              config.Username,
		Password:              config.Password,
		MaxRetries:            config.MaxRetries,
		MinRetryBackoff:       config.MinRetryBackoff,
		MaxRetryBackoff:       config.MaxRetryBackoff,
		DialTimeout:           config.DialTimeout,
		ReadTimeout:           config.ReadTimeout,
		WriteTimeout:          config.WriteTimeout,
		ContextTimeoutEnabled: config.ContextTimeoutEnabled,
		PoolFIFO:              config.PoolFIFO,
		PoolSize:              config.PoolSize,
		PoolTimeout:           config.PoolTimeout,
		MinIdleConns:          config.MinIdleConns,
		MaxIdleConns:          config.MaxIdleConns,
		MaxActiveConns:        config.MaxActiveConns,
		ConnMaxIdleTime:       config.ConnMaxIdleTime,
		ConnMaxLifetime:       config.ConnMaxLifetime,
		DisableIndentity:      config.DisableIndentity,
		DisableIdentity:       config.DisableIdentity,
		IdentitySuffix:        config.IdentitySuffix,
		UnstableResp3:         config.UnstableResp3,
	}
}
