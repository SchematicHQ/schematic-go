package core

import (
	"time"

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
}

func (c RedisCacheConfig) applyDatastreamOptions(opts *DatastreamOptions) {
	opts.CacheConfig = c
}

type RedisCacheClusterConfig struct {
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
}

func (c RedisCacheClusterConfig) applyDatastreamOptions(opts *DatastreamOptions) {
	opts.CacheConfig = c
}

func WithRedisCache(opts CacheConfig) CacheConfig {
	return opts
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
