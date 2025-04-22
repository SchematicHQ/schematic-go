package datastream

import (
	"github.com/go-redis/redis/v8"
	"github.com/schematichq/rulesengine"
	"github.com/schematichq/schematic-go/cache"
	"github.com/schematichq/schematic-go/core"
)

func getCacheProviders(opt *core.DatastreamOptions) (FlagCacheProvider, CompanyCacheProvider, UserCacheProvider) {
	if opt.CacheConfig == nil {
		return buildLocalCache(opt)
	}

	switch opt.CacheConfig.(type) {
	case *core.RedisCacheConfig:
		config := opt.CacheConfig.(*core.RedisCacheConfig)
		client := redis.NewClient(config.Options)
		return buildRedisCache(client, opt)
	case *core.RedisCacheClusterConfig:
		config := opt.CacheConfig.(*core.RedisCacheClusterConfig)
		client := redis.NewClusterClient(config.ClusterOptions)
		return buildRedisCache(client, opt)
	}

	return nil, nil, nil
}

func buildRedisCache(client redis.UniversalClient, opt *core.DatastreamOptions) (FlagCacheProvider, CompanyCacheProvider, UserCacheProvider) {
	flagCacheProvider := cache.NewRedisCache[*rulesengine.Flag](client, opt.CacheTTL)
	companyCacheProvider := cache.NewRedisCache[*rulesengine.Company](client, opt.CacheTTL)
	userCacheProvider := cache.NewRedisCache[*rulesengine.User](client, opt.CacheTTL)

	return flagCacheProvider, companyCacheProvider, userCacheProvider
}

func buildLocalCache(opt *core.DatastreamOptions) (FlagCacheProvider, CompanyCacheProvider, UserCacheProvider) {
	flagCacheProvider := cache.NewLocalCache[*rulesengine.Flag](defaultCacheSize, opt.CacheTTL)
	companyCacheProvider := cache.NewLocalCache[*rulesengine.Company](defaultCacheSize, opt.CacheTTL)
	userCacheProvider := cache.NewLocalCache[*rulesengine.User](defaultCacheSize, opt.CacheTTL)

	return flagCacheProvider, companyCacheProvider, userCacheProvider
}
