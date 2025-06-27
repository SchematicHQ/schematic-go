package datastream

import (
	"github.com/redis/go-redis/v9"
	"github.com/schematichq/rulesengine"
	"github.com/schematichq/schematic-go/cache"
	"github.com/schematichq/schematic-go/core"
)

func getCacheProviders(options DataStreamClientOptions, configOpt *core.DatastreamOptions) (CompanyCacheProvider, UserCacheProvider) {
	// If both cache providers are specified in options, use them
	if options.CompanyCache != nil && options.UserCache != nil {
		return CompanyCacheProvider(options.CompanyCache), UserCacheProvider(options.UserCache)
	}

	// If only company cache provider is specified, create only user cache provider
	if options.CompanyCache != nil {
		_, userCacheProvider := buildCacheProvidersFromConfig(configOpt)
		return CompanyCacheProvider(options.CompanyCache), userCacheProvider
	}

	// If only user cache provider is specified, create only company cache provider
	if options.UserCache != nil {
		companyCacheProvider, _ := buildCacheProvidersFromConfig(configOpt)
		return companyCacheProvider, UserCacheProvider(options.UserCache)
	}

	// Otherwise build both cache providers based on configuration options
	return buildCacheProvidersFromConfig(configOpt)
}

// Helper function to build cache providers based on configuration options
func buildCacheProvidersFromConfig(configOpt *core.DatastreamOptions) (CompanyCacheProvider, UserCacheProvider) {
	if configOpt == nil || configOpt.CacheConfig == nil {
		return buildLocalCache(configOpt)
	}

	switch configOpt.CacheConfig.(type) {
	case *core.RedisCacheConfig:
		config := configOpt.CacheConfig.(*core.RedisCacheConfig)
		client := redis.NewClient(ToRedisOptions(config))
		return buildRedisCache(client, configOpt)
	case core.RedisCacheConfig:
		config := configOpt.CacheConfig.(core.RedisCacheConfig)
		client := redis.NewClient(ToRedisOptions(&config))
		return buildRedisCache(client, configOpt)
	case *core.RedisCacheClusterConfig:
		config := configOpt.CacheConfig.(*core.RedisCacheClusterConfig)
		client := redis.NewClusterClient(ToRedisClusterOptions(config))
		return buildRedisCache(client, configOpt)
	case core.RedisCacheClusterConfig:
		config := configOpt.CacheConfig.(core.RedisCacheClusterConfig)
		client := redis.NewClusterClient(ToRedisClusterOptions(&config))
		return buildRedisCache(client, configOpt)
	}

	return nil, nil
}

func buildRedisCache(client redis.UniversalClient, opt *core.DatastreamOptions) (CompanyCacheProvider, UserCacheProvider) {
	// Default cache TTL if opt is nil or CacheTTL is zero
	cacheTTL := defaultTTL
	if opt != nil && opt.CacheTTL > 0 {
		cacheTTL = opt.CacheTTL
	}

	companyCacheProvider := cache.NewRedisCache[*rulesengine.Company](client, cacheTTL)
	userCacheProvider := cache.NewRedisCache[*rulesengine.User](client, cacheTTL)

	return companyCacheProvider, userCacheProvider
}

func buildLocalCache(opt *core.DatastreamOptions) (CompanyCacheProvider, UserCacheProvider) {
	// Default cache TTL if opt is nil or CacheTTL is zero
	cacheTTL := defaultTTL
	if opt != nil && opt.CacheTTL > 0 {
		cacheTTL = opt.CacheTTL
	}

	companyCacheProvider := cache.NewLocalCache[*rulesengine.Company](defaultCacheSize, cacheTTL)
	userCacheProvider := cache.NewLocalCache[*rulesengine.User](defaultCacheSize, cacheTTL)

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
