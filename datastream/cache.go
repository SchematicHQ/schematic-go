package datastream

import (
	"github.com/redis/go-redis/v9"
	"github.com/schematichq/rulesengine"
	"github.com/schematichq/schematic-go/cache"
	"github.com/schematichq/schematic-go/core"
)

func getCacheProviders(options DataStreamClientOptions, configOpt *core.DatastreamOptions, redisClient redis.UniversalClient) (CompanyCacheProvider, UserCacheProvider) {
	// If both cache providers are specified in options, use them
	if options.CompanyCache != nil && options.UserCache != nil {
		return CompanyCacheProvider(options.CompanyCache), UserCacheProvider(options.UserCache)
	}

	// If only company cache provider is specified, create only user cache provider
	if options.CompanyCache != nil {
		_, userCacheProvider := buildCacheProvidersFromConfig(configOpt, redisClient)
		return CompanyCacheProvider(options.CompanyCache), userCacheProvider
	}

	// If only user cache provider is specified, create only company cache provider
	if options.UserCache != nil {
		companyCacheProvider, _ := buildCacheProvidersFromConfig(configOpt, redisClient)
		return companyCacheProvider, UserCacheProvider(options.UserCache)
	}

	// Otherwise build both cache providers based on configuration options
	return buildCacheProvidersFromConfig(configOpt, redisClient)
}

// buildFlagCacheProvider creates a flag cache provider using a shared Redis client.
// Flags use the same cache provider type as company/user caches but with special TTL logic.
func buildFlagCacheProvider(configOpt *core.DatastreamOptions, redisClient redis.UniversalClient) cache.CacheProvider[*rulesengine.Flag] {
	// Calculate flag TTL - use the greater value between configured TTL and max TTL
	flagTTL := maxCacheTTL
	if configOpt != nil && configOpt.CacheTTL > maxCacheTTL {
		flagTTL = configOpt.CacheTTL
	}

	if redisClient != nil {
		return cache.NewRedisCache[*rulesengine.Flag](redisClient, flagTTL)
	}

	return cache.NewLocalCache[*rulesengine.Flag](defaultCacheSize, flagTTL)
}

// buildCompanyLookupCacheProvider creates a string cache provider for company ID lookups
// using a shared Redis client to avoid creating additional connections.
func buildCompanyLookupCacheProvider(configOpt *core.DatastreamOptions, redisClient redis.UniversalClient) CompanyLookupCacheProvider {
	cacheTTL := defaultTTL
	if configOpt != nil && configOpt.CacheTTL > 0 {
		cacheTTL = configOpt.CacheTTL
	}

	if redisClient != nil {
		return cache.NewRedisCache[string](redisClient, cacheTTL)
	}

	return cache.NewLocalCache[string](defaultCacheSize, cacheTTL)
}

// buildRedisClient creates a Redis client from the cache configuration, or returns nil
// if no Redis config is provided. The returned client should be shared across all cache
// providers to respect the user's connection pool settings.
func buildRedisClient(configOpt *core.DatastreamOptions) redis.UniversalClient {
	if configOpt == nil || configOpt.CacheConfig == nil {
		return nil
	}

	switch configOpt.CacheConfig.(type) {
	case *core.RedisCacheConfig:
		config := configOpt.CacheConfig.(*core.RedisCacheConfig)
		return redis.NewClient(ToRedisOptions(config))
	case core.RedisCacheConfig:
		config := configOpt.CacheConfig.(core.RedisCacheConfig)
		return redis.NewClient(ToRedisOptions(&config))
	case *core.RedisCacheClusterConfig:
		config := configOpt.CacheConfig.(*core.RedisCacheClusterConfig)
		return redis.NewClusterClient(ToRedisClusterOptions(config))
	case core.RedisCacheClusterConfig:
		config := configOpt.CacheConfig.(core.RedisCacheClusterConfig)
		return redis.NewClusterClient(ToRedisClusterOptions(&config))
	}

	return nil
}

// Helper function to build cache providers based on configuration options
func buildCacheProvidersFromConfig(configOpt *core.DatastreamOptions, redisClient redis.UniversalClient) (CompanyCacheProvider, UserCacheProvider) {
	if redisClient != nil {
		return buildRedisCache(redisClient, configOpt)
	}
	return buildLocalCache(configOpt)
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
