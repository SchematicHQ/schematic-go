package core

import (
	"crypto/tls"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/schematichq/schematic-go/cache"
	"github.com/schematichq/schematic-go/http"
)

// Specify default flag values

type ClientOptFlagDefaults struct {
	values map[string]bool
}

func (c ClientOptFlagDefaults) applyRequestOptions(opts *RequestOptions) {
	opts.FlagDefaults = c.values
}

func WithDefaultFlagValues(values map[string]bool) RequestOption {
	return ClientOptFlagDefaults{values: values}
}

// Disable local cache entirely

type ClientOptDisableFlagCheckCache struct {
}

func (c ClientOptDisableFlagCheckCache) applyRequestOptions(opts *RequestOptions) {
	opts.FlagCheckCacheProviders = append(opts.FlagCheckCacheProviders, cache.NewLocalCache[*CheckFlagResponse](0, 0))
}

func WithDisableFlagCheckCache() RequestOption {
	return ClientOptDisableFlagCheckCache{}
}

// Event buffer period

type ClientOptEventBufferPeriod struct {
	period time.Duration
}

func (c ClientOptEventBufferPeriod) applyRequestOptions(opts *RequestOptions) {
	opts.EventBufferPeriod = &c.period
}

func WithEventBufferPeriod(period time.Duration) RequestOption {
	return ClientOptEventBufferPeriod{
		period: period,
	}
}

// Shutdown timeout

type ClientOptShutdownTimeout struct {
	timeout time.Duration
}

func (c ClientOptShutdownTimeout) applyRequestOptions(opts *RequestOptions) {
	opts.ShutdownTimeout = &c.timeout
}

// WithShutdownTimeout bounds how long Close spends flushing buffered events
// before giving up. Lower it when the process has a short termination grace
// period; the default is 10 seconds.
func WithShutdownTimeout(timeout time.Duration) RequestOption {
	return ClientOptShutdownTimeout{
		timeout: timeout,
	}
}

// Event Capture Base URL

type ClientOptEventCaptureBaseURL struct {
	url string
}

func (c ClientOptEventCaptureBaseURL) applyRequestOptions(opts *RequestOptions) {
	opts.EventCaptureBaseURL = c.url
}

func WithEventCaptureBaseURL(url string) RequestOption {
	return ClientOptEventCaptureBaseURL{
		url: url,
	}
}

// Specify custom cache behavior

type ClientOptFlagCheckCacheProvider struct {
	provider cache.CacheProvider[*CheckFlagResponse]
}

func (c ClientOptFlagCheckCacheProvider) applyRequestOptions(opts *RequestOptions) {
	opts.FlagCheckCacheProviders = append(opts.FlagCheckCacheProviders, c.provider)
}

func WithFlagCheckCacheProvider(provider cache.CacheProvider[*CheckFlagResponse]) RequestOption {
	return ClientOptFlagCheckCacheProvider{provider: provider}
}

// Specify local cache behavior

type ClientOptLocalFlagCheckCache struct {
	maxSize int
	ttl     time.Duration
}

func (c ClientOptLocalFlagCheckCache) applyRequestOptions(opts *RequestOptions) {
	opts.FlagCheckCacheProviders = append(opts.FlagCheckCacheProviders, cache.NewLocalCache[*CheckFlagResponse](c.maxSize, c.ttl))
}

func WithLocalFlagCheckCache(maxSize int, ttl time.Duration) RequestOption {
	return ClientOptLocalFlagCheckCache{maxSize: maxSize, ttl: ttl}
}

// Offline mode

type ClientOptOfflineMode struct {
	isOffline bool
}

func (c ClientOptOfflineMode) applyRequestOptions(opts *RequestOptions) {
	opts.OfflineMode = c.isOffline
	opts.HTTPClient = http.NewNoopClient()
}

func WithOfflineMode() RequestOption {
	return ClientOptOfflineMode{
		isOffline: true,
	}
}

type CustomLogger struct {
	logger Logger
}

func (c CustomLogger) applyRequestOptions(opts *RequestOptions) {
	opts.Logger = c.logger
}
func WithLogger(logger Logger) RequestOption {
	return CustomLogger{logger: logger}
}

type LogLevelOption struct {
	level LogLevel
}

func (o LogLevelOption) applyRequestOptions(opts *RequestOptions) {
	opts.LogLevel = o.level
}

func WithLogLevel(level LogLevel) RequestOption {
	return LogLevelOption{level: level}
}

// Datastream options
type DatastreamOption interface {
	applyDatastreamOptions(opts *DatastreamOptions)
}

type DatastreamOptions struct {
	CacheTTL              time.Duration
	CacheConfig           CacheConfig
	ReplicatorMode        bool
	ReplicatorHealthURL   string
	ReplicatorHealthCheck time.Duration
}

type CacheTTL struct {
	ttl time.Duration
}

func (c CacheTTL) applyDatastreamOptions(opts *DatastreamOptions) {
	opts.CacheTTL = c.ttl
}

func WithCacheTTL(ttl time.Duration) DatastreamOption {
	return CacheTTL{ttl: ttl}
}

// Define an interface for Redis options
type CacheConfig interface {
	applyDatastreamOptions(opts *DatastreamOptions)
}

type RedisCacheConfig struct {
	// URL is a Redis connection URL, e.g. redis://host:6379/0 or
	// rediss://user:pass@host:6379 for TLS, which yields a TLSConfig with
	// ServerName already set. It takes precedence over Addr, and an Addr carrying
	// a scheme is treated as a URL too.
	//
	// Any field below set to a non-zero value overrides what the URL derived. Only
	// non-zero ones, so a URL-derived setting cannot be turned back off here: drop
	// it from the URL instead.
	//
	// A URL that does not parse, or names no host, leaves the datastream on its
	// local cache rather than being dialed as-is.
	URL                   string
	Network               string
	Addr                  string
	ClientName            string
	Protocol              int
	Username              string
	Password              string
	DB                    int
	MaxRetries            int
	MinRetryBackoff       time.Duration
	MaxRetryBackoff       time.Duration
	DialTimeout           time.Duration
	ReadTimeout           time.Duration
	WriteTimeout          time.Duration
	ContextTimeoutEnabled bool
	PoolFIFO              bool
	PoolSize              int
	PoolTimeout           time.Duration
	MinIdleConns          int
	MaxIdleConns          int
	MaxActiveConns        int
	ConnMaxIdleTime       time.Duration
	ConnMaxLifetime       time.Duration
	DisableIndentity      bool
	DisableIdentity       bool
	IdentitySuffix        string
	UnstableResp3         bool
	// TLSConfig enables TLS for the connection. Nil, the zero value, means
	// plaintext, which is go-redis' default.
	TLSConfig *tls.Config
}

func (c RedisCacheConfig) applyDatastreamOptions(opts *DatastreamOptions) {
	opts.CacheConfig = c
}

type RedisCacheClusterConfig struct {
	// URL is a Redis connection URL for the cluster — an ElastiCache or Valkey
	// configuration endpoint, typically. It follows RedisCacheConfig.URL's rules,
	// and takes precedence over Addrs, which is then ignored; name further nodes
	// with go-redis' ?addr=host:port parameter.
	//
	// A single entry in Addrs carrying a scheme is treated as this URL too, with
	// the remaining bare entries kept as additional nodes. Two of them is an error:
	// each carries its own credentials and TLS, with nothing to choose between them.
	URL                   string
	Addrs                 []string
	MaxRedirects          int
	RouteByLatency        bool
	RouteRandomly         bool
	Protocol              int
	Username              string
	Password              string
	MaxRetries            int
	MinRetryBackoff       time.Duration
	MaxRetryBackoff       time.Duration
	DialTimeout           time.Duration
	ReadTimeout           time.Duration
	WriteTimeout          time.Duration
	ContextTimeoutEnabled bool
	PoolFIFO              bool
	PoolSize              int
	PoolTimeout           time.Duration
	MinIdleConns          int
	MaxIdleConns          int
	MaxActiveConns        int
	ConnMaxIdleTime       time.Duration
	ConnMaxLifetime       time.Duration
	DisableIndentity      bool
	DisableIdentity       bool
	IdentitySuffix        string
	UnstableResp3         bool
	// TLSConfig enables TLS for connections to every node in the cluster. Nil,
	// the zero value, means plaintext, which is go-redis' default.
	TLSConfig *tls.Config
}

func (c RedisCacheClusterConfig) applyDatastreamOptions(opts *DatastreamOptions) {
	opts.CacheConfig = c
}

// RedisClientConfig hands the datastream a Redis client the caller has already
// built, instead of describing one field by field. Prefer it when the
// application constructs its own client: the field-by-field configs map a fixed
// set of go-redis options, so anything outside that set — a custom dialer, a
// hook, a tracing wrapper — cannot be expressed through them at all.
//
// The SDK shares the client across every datastream cache and never closes it;
// its lifecycle stays with the caller.
type RedisClientConfig struct {
	Client redis.UniversalClient
}

func (c RedisClientConfig) applyDatastreamOptions(opts *DatastreamOptions) {
	opts.CacheConfig = c
}

func WithRedisCache(opts CacheConfig) CacheConfig {
	return opts
}

// WithRedisClient wraps an existing Redis client as a datastream cache config.
// The client may be any redis.UniversalClient — single-node, cluster or
// failover.
func WithRedisClient(client redis.UniversalClient) CacheConfig {
	return RedisClientConfig{Client: client}
}

type ClientOptUseDatastream struct {
	enabled bool
	options *DatastreamOptions
}

func (c ClientOptUseDatastream) applyRequestOptions(opts *RequestOptions) {
	opts.UseDataStream = c.enabled
	opts.DatastreamOptions = c.options
}

func WithDatastream(opts ...DatastreamOption) RequestOption {
	dataStreamOptions := &DatastreamOptions{
		CacheTTL: 24 * time.Hour,
	}

	for _, opt := range opts {
		opt.applyDatastreamOptions(dataStreamOptions)
	}

	// Apply replicator defaults if replicator mode is enabled but specific options weren't provided
	if dataStreamOptions.ReplicatorMode {
		if dataStreamOptions.ReplicatorHealthURL == "" {
			dataStreamOptions.ReplicatorHealthURL = "http://localhost:8090/ready"
		}
		if dataStreamOptions.ReplicatorHealthCheck == 0 {
			dataStreamOptions.ReplicatorHealthCheck = 30 * time.Second
		}
	}

	return &ClientOptUseDatastream{
		enabled: true,
		options: dataStreamOptions,
	}
}

// ReplicatorMode enables replicator mode for datastream client
type ReplicatorMode struct{}

func (s ReplicatorMode) applyDatastreamOptions(opts *DatastreamOptions) {
	opts.ReplicatorMode = true
}

func WithReplicatorMode() DatastreamOption {
	return ReplicatorMode{}
}

// ReplicatorHealthURL configures the health check URL for replicator mode
type ReplicatorHealthURL struct {
	url string
}

func (s ReplicatorHealthURL) applyDatastreamOptions(opts *DatastreamOptions) {
	opts.ReplicatorHealthURL = s.url
}

func WithReplicatorHealthURL(url string) DatastreamOption {
	return ReplicatorHealthURL{url: url}
}

// ReplicatorHealthInterval configures the health check interval for replicator mode
type ReplicatorHealthInterval struct {
	interval time.Duration
}

func (s ReplicatorHealthInterval) applyDatastreamOptions(opts *DatastreamOptions) {
	opts.ReplicatorHealthCheck = s.interval
}

func WithReplicatorHealthInterval(interval time.Duration) DatastreamOption {
	return ReplicatorHealthInterval{interval: interval}
}
