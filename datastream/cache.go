package datastream

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"github.com/redis/go-redis/v9"
	"github.com/schematichq/schematic-go/cache"
	"github.com/schematichq/schematic-go/core"
	"github.com/schematichq/schematic-go/rulesengine"
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

// buildLookupCacheProvider creates a string cache provider for resource ID lookups
// using a shared Redis client to avoid creating additional connections.
func buildLookupCacheProvider(configOpt *core.DatastreamOptions, redisClient redis.UniversalClient) cache.CacheProvider[string] {
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
// providers to respect the user's connection pool settings. The logger, which may be
// nil, receives a warning if a configured connection URL cannot be parsed.
func buildRedisClient(configOpt *core.DatastreamOptions, logger core.Logger) redis.UniversalClient {
	if configOpt == nil || configOpt.CacheConfig == nil {
		return nil
	}

	switch configOpt.CacheConfig.(type) {
	case *core.RedisCacheConfig:
		config := configOpt.CacheConfig.(*core.RedisCacheConfig)
		return newRedisClient(config, logger)
	case core.RedisCacheConfig:
		config := configOpt.CacheConfig.(core.RedisCacheConfig)
		return newRedisClient(&config, logger)
	case *core.RedisCacheClusterConfig:
		config := configOpt.CacheConfig.(*core.RedisCacheClusterConfig)
		return newRedisClusterClient(config, logger)
	case core.RedisCacheClusterConfig:
		config := configOpt.CacheConfig.(core.RedisCacheClusterConfig)
		return newRedisClusterClient(&config, logger)
	case *core.RedisClientConfig:
		return callerSuppliedClient(configOpt.CacheConfig.(*core.RedisClientConfig).Client)
	case core.RedisClientConfig:
		return callerSuppliedClient(configOpt.CacheConfig.(core.RedisClientConfig).Client)
	}

	return nil
}

// callerSuppliedClient guards against a typed nil. A nil client means "no Redis" and
// sends buildCacheProvidersFromConfig to the local cache, but a nil *redis.Client stored
// in the interface is not nil and would panic on first use instead.
func callerSuppliedClient(client redis.UniversalClient) redis.UniversalClient {
	if client == nil || isNilValue(client) {
		return nil
	}

	return client
}

// isNilValue reports whether the interface holds a nil pointer or other nilable zero.
// The kind is checked first because reflect.Value.IsNil panics on kinds that cannot be
// nil, and a caller implementing the interface by value must not blow up the constructor.
func isNilValue(client redis.UniversalClient) bool {
	value := reflect.ValueOf(client)

	switch value.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Map, reflect.Slice, reflect.Func, reflect.Chan, reflect.UnsafePointer:
		return value.IsNil()
	default:
		return false
	}
}

// newRedisClient builds a single-instance client, or nothing at all if the connection
// URL is unusable. Dialing on regardless would mean go-redis' localhost:6379 default or
// an address that never reached a Redis; returning nil degrades to the local cache
// instead, the same path a nil config takes.
func newRedisClient(config *core.RedisCacheConfig, logger core.Logger) redis.UniversalClient {
	opts, err := toRedisOptions(config)
	if err != nil {
		if logger != nil {
			connURL := redisConnectionURL(config)
			logger.Warn(context.Background(), fmt.Sprintf(
				"Could not parse Redis connection URL '%s', falling back to the local cache: %s",
				redactRedisURL(connURL), redactRedisError(err, connURL)))
		}

		return nil
	}

	return redis.NewClient(opts)
}

// newRedisClusterClient builds a cluster client, declining an unusable URL for the same
// reason newRedisClient does.
func newRedisClusterClient(config *core.RedisCacheClusterConfig, logger core.Logger) redis.UniversalClient {
	opts, err := toRedisClusterOptions(config)
	if err != nil {
		if logger != nil {
			connURL, _, _ := clusterConnectionURL(config)
			logger.Warn(context.Background(), fmt.Sprintf(
				"Could not use Redis cluster connection URL '%s', falling back to the local cache: %s",
				redactRedisURL(connURL), redactRedisError(err, connURL)))
		}

		return nil
	}

	return redis.NewClusterClient(opts)
}

// redactRedisURL blanks any credentials in a connection URL so it can be logged. The
// URL is malformed by the time this runs, so it cannot lean on url.Parse; it drops
// whatever sits before the last '@' of the authority instead.
//
// An unescaped '/' or '?' in a password hides that '@' from the scan, so a '@' beyond the
// guessed authority is treated as proof the guess is wrong and the whole URL is elided.
// Losing the host on a URL whose only '@' is in a query value is the cheaper mistake.
func redactRedisURL(connURL string) string {
	scheme, authority := "", connURL
	if i := strings.Index(connURL, "://"); i >= 0 {
		scheme, authority = connURL[:i+len("://")], connURL[i+len("://"):]
	}

	end := strings.IndexAny(authority, "/?")
	if end < 0 {
		end = len(authority)
	}

	at := strings.LastIndex(authority[:end], "@")
	if at < 0 {
		if strings.Contains(authority[end:], "@") {
			return scheme + "redacted"
		}

		return connURL
	}

	return scheme + "redacted@" + authority[at+1:]
}

// redactRedisError reports how much of a parse failure is safe to log. The URL cannot be
// scrubbed out of the error text: net/url quotes it in %q-escaped, defeating a literal
// replacement, and quotes fragments besides — a password holding a '/' comes back named
// as the invalid port. So the text is logged only when there is no userinfo to leak.
func redactRedisError(err error, connURL string) string {
	if strings.Contains(connURL, "@") {
		return "error detail withheld because the URL carries credentials"
	}

	return err.Error()
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

// ToRedisOptions maps a cache config onto go-redis options, resolving a connection URL
// when the config carries one. An unusable URL is left in Addr verbatim, so a client
// built from these options fails loudly rather than against go-redis' localhost default.
func ToRedisOptions(config *core.RedisCacheConfig) *redis.Options {
	opts, _ := toRedisOptions(config)
	return opts
}

// toRedisOptions is ToRedisOptions with the parse failure surfaced; the exported name is
// public API and cannot grow a second return value. It always returns usable options.
func toRedisOptions(config *core.RedisCacheConfig) (*redis.Options, error) {
	connURL := redisConnectionURL(config)
	if connURL == "" {
		return applyRedisConfig(&redis.Options{Addr: config.Addr}, config), nil
	}

	parsed, err := redis.ParseURL(connURL)
	if err == nil {
		err = checkRedisURLHost(connURL)
	}
	if err != nil {
		// connURL, never Addr: an empty Addr would become localhost:6379.
		return applyRedisConfig(&redis.Options{Addr: connURL}, config), err
	}

	return applyRedisConfig(parsed, config), nil
}

// errRedisURLNoHost reports a connection URL that names no server to connect to.
var errRedisURLNoHost = errors.New("redis: connection URL has no host")

// checkRedisURLHost rejects a URL that names no host, which redis.ParseURL accepts and
// turns into localhost:6379 — the shape an unset environment variable takes. A unix://
// URL carries a socket path instead and is left alone.
func checkRedisURLHost(connURL string) error {
	parsed, err := url.Parse(connURL)
	if err != nil || parsed.Scheme == "unix" {
		return nil
	}

	if parsed.Hostname() == "" {
		return errRedisURLNoHost
	}

	return nil
}

// redisConnectionURL reports the connection URL the config asks for, if any. Callers
// may set URL, or leave a scheme-bearing string in Addr: go-redis wants a bare
// host:port there, so such a value only ever reached net.Dial and failed, and treating
// it as a URL breaks nothing that worked.
func redisConnectionURL(config *core.RedisCacheConfig) string {
	if config.URL != "" {
		return config.URL
	}

	if strings.Contains(config.Addr, "://") {
		return config.Addr
	}

	return ""
}

// applyRedisConfig overlays the config's non-zero fields onto opts, which may already
// carry values a connection URL derived. Leaving zero values alone is what lets
// URL-derived settings survive; on a fresh redis.Options it reproduces the previous
// field-for-field mapping.
func applyRedisConfig(opts *redis.Options, config *core.RedisCacheConfig) *redis.Options {
	if config.Network != "" {
		opts.Network = config.Network
	}
	if config.ClientName != "" {
		opts.ClientName = config.ClientName
	}
	if config.Protocol != 0 {
		opts.Protocol = config.Protocol
	}
	if config.Username != "" {
		opts.Username = config.Username
	}
	if config.Password != "" {
		opts.Password = config.Password
	}
	if config.DB != 0 {
		opts.DB = config.DB
	}
	if config.MaxRetries != 0 {
		opts.MaxRetries = config.MaxRetries
	}
	if config.MinRetryBackoff != 0 {
		opts.MinRetryBackoff = config.MinRetryBackoff
	}
	if config.MaxRetryBackoff != 0 {
		opts.MaxRetryBackoff = config.MaxRetryBackoff
	}
	if config.DialTimeout != 0 {
		opts.DialTimeout = config.DialTimeout
	}
	if config.ReadTimeout != 0 {
		opts.ReadTimeout = config.ReadTimeout
	}
	if config.WriteTimeout != 0 {
		opts.WriteTimeout = config.WriteTimeout
	}
	if config.ContextTimeoutEnabled {
		opts.ContextTimeoutEnabled = true
	}
	if config.PoolFIFO {
		opts.PoolFIFO = true
	}
	if config.PoolSize != 0 {
		opts.PoolSize = config.PoolSize
	}
	if config.PoolTimeout != 0 {
		opts.PoolTimeout = config.PoolTimeout
	}
	if config.MinIdleConns != 0 {
		opts.MinIdleConns = config.MinIdleConns
	}
	if config.MaxIdleConns != 0 {
		opts.MaxIdleConns = config.MaxIdleConns
	}
	if config.MaxActiveConns != 0 {
		opts.MaxActiveConns = config.MaxActiveConns
	}
	if config.ConnMaxIdleTime != 0 {
		opts.ConnMaxIdleTime = config.ConnMaxIdleTime
	}
	if config.ConnMaxLifetime != 0 {
		opts.ConnMaxLifetime = config.ConnMaxLifetime
	}
	if config.DisableIndentity {
		//nolint:staticcheck // the config still exposes go-redis' misspelled field; keep mapping it
		opts.DisableIndentity = true
	}
	if config.DisableIdentity {
		opts.DisableIdentity = true
	}
	if config.IdentitySuffix != "" {
		opts.IdentitySuffix = config.IdentitySuffix
	}
	if config.UnstableResp3 {
		//nolint:staticcheck // deprecated no-op in go-redis, but the config still exposes it
		opts.UnstableResp3 = true
	}
	if config.TLSConfig != nil {
		opts.TLSConfig = config.TLSConfig
	}

	return opts
}

// ToRedisClusterOptions maps a cluster cache config onto go-redis options, resolving a
// connection URL when the config carries one. An unusable URL is left in Addrs verbatim,
// so a client built from these options fails loudly rather than against go-redis'
// localhost default.
func ToRedisClusterOptions(config *core.RedisCacheClusterConfig) *redis.ClusterOptions {
	opts, _ := toRedisClusterOptions(config)
	return opts
}

// toRedisClusterOptions is ToRedisClusterOptions with the parse failure surfaced, split
// out for the same reason toRedisOptions is.
func toRedisClusterOptions(config *core.RedisCacheClusterConfig) (*redis.ClusterOptions, error) {
	connURL, addrs, err := clusterConnectionURL(config)
	if err != nil {
		return applyRedisClusterConfig(&redis.ClusterOptions{Addrs: config.Addrs}, config), err
	}

	if connURL == "" {
		return applyRedisClusterConfig(&redis.ClusterOptions{Addrs: config.Addrs}, config), nil
	}

	parsed, err := redis.ParseClusterURL(connURL)
	if err == nil {
		err = checkRedisURLHost(connURL)
	}
	if err != nil {
		// connURL, never Addrs: an empty Addrs would become localhost:6379.
		return applyRedisClusterConfig(&redis.ClusterOptions{Addrs: []string{connURL}}, config), err
	}

	parsed.Addrs = append(parsed.Addrs, addrs...)

	return applyRedisClusterConfig(parsed, config), nil
}

// errRedisMultipleClusterURLs reports a config naming more than one connection URL.
var errRedisMultipleClusterURLs = errors.New("redis: cluster config names more than one connection URL")

// clusterConnectionURL reports the cluster's connection URL, if any, with the bare
// host:port entries that belong beside it. URL wins outright; otherwise a lone
// scheme-bearing entry in Addrs is the URL. Two of them carry two sets of credentials
// and TLS, with nothing to choose between them.
func clusterConnectionURL(config *core.RedisCacheClusterConfig) (connURL string, addrs []string, err error) {
	if config.URL != "" {
		return config.URL, nil, nil
	}

	for _, addr := range config.Addrs {
		if !strings.Contains(addr, "://") {
			addrs = append(addrs, addr)
			continue
		}

		if connURL != "" {
			return connURL, addrs, errRedisMultipleClusterURLs
		}

		connURL = addr
	}

	if connURL == "" {
		return "", nil, nil
	}

	return connURL, addrs, nil
}

// applyRedisClusterConfig is applyRedisConfig for a cluster. Addrs is not among the
// fields it overlays: the URL and the bare entries beside it have already settled that.
func applyRedisClusterConfig(opts *redis.ClusterOptions, config *core.RedisCacheClusterConfig) *redis.ClusterOptions {
	if config.MaxRedirects != 0 {
		opts.MaxRedirects = config.MaxRedirects
	}
	if config.RouteByLatency {
		opts.RouteByLatency = true
	}
	if config.RouteRandomly {
		opts.RouteRandomly = true
	}
	if config.Protocol != 0 {
		opts.Protocol = config.Protocol
	}
	if config.Username != "" {
		opts.Username = config.Username
	}
	if config.Password != "" {
		opts.Password = config.Password
	}
	if config.MaxRetries != 0 {
		opts.MaxRetries = config.MaxRetries
	}
	if config.MinRetryBackoff != 0 {
		opts.MinRetryBackoff = config.MinRetryBackoff
	}
	if config.MaxRetryBackoff != 0 {
		opts.MaxRetryBackoff = config.MaxRetryBackoff
	}
	if config.DialTimeout != 0 {
		opts.DialTimeout = config.DialTimeout
	}
	if config.ReadTimeout != 0 {
		opts.ReadTimeout = config.ReadTimeout
	}
	if config.WriteTimeout != 0 {
		opts.WriteTimeout = config.WriteTimeout
	}
	if config.ContextTimeoutEnabled {
		opts.ContextTimeoutEnabled = true
	}
	if config.PoolFIFO {
		opts.PoolFIFO = true
	}
	if config.PoolSize != 0 {
		opts.PoolSize = config.PoolSize
	}
	if config.PoolTimeout != 0 {
		opts.PoolTimeout = config.PoolTimeout
	}
	if config.MinIdleConns != 0 {
		opts.MinIdleConns = config.MinIdleConns
	}
	if config.MaxIdleConns != 0 {
		opts.MaxIdleConns = config.MaxIdleConns
	}
	if config.MaxActiveConns != 0 {
		opts.MaxActiveConns = config.MaxActiveConns
	}
	if config.ConnMaxIdleTime != 0 {
		opts.ConnMaxIdleTime = config.ConnMaxIdleTime
	}
	if config.ConnMaxLifetime != 0 {
		opts.ConnMaxLifetime = config.ConnMaxLifetime
	}
	if config.DisableIndentity {
		//nolint:staticcheck // the config still exposes go-redis' misspelled field; keep mapping it
		opts.DisableIndentity = true
	}
	if config.DisableIdentity {
		opts.DisableIdentity = true
	}
	if config.IdentitySuffix != "" {
		opts.IdentitySuffix = config.IdentitySuffix
	}
	if config.UnstableResp3 {
		//nolint:staticcheck // deprecated no-op in go-redis, but the config still exposes it
		opts.UnstableResp3 = true
	}
	if config.TLSConfig != nil {
		opts.TLSConfig = config.TLSConfig
	}

	return opts
}
