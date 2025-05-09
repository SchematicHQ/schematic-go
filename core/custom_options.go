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
	opts.FlagCheckCacheProviders = append(opts.FlagCheckCacheProviders, cache.NewLocalCache[bool](0, 0))
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

// Specify custom cache behavior

type ClientOptFlagCheckCacheProvider struct {
	provider cache.CacheProvider[bool]
}

func (c ClientOptFlagCheckCacheProvider) applyRequestOptions(opts *RequestOptions) {
	opts.FlagCheckCacheProviders = append(opts.FlagCheckCacheProviders, c.provider)
}

func WithFlagCheckCacheProvider(provider cache.CacheProvider[bool]) RequestOption {
	return ClientOptFlagCheckCacheProvider{provider: provider}
}

// Specify local cache behavior

type ClientOptLocalFlagCheckCache struct {
	maxSize int
	ttl     time.Duration
}

func (c ClientOptLocalFlagCheckCache) applyRequestOptions(opts *RequestOptions) {
	opts.FlagCheckCacheProviders = append(opts.FlagCheckCacheProviders, cache.NewLocalCache[bool](c.maxSize, c.ttl))
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

// Datastream options
type DatastreamOption interface {
	applyDatastreamOptions(opts *DatastreamOptions)
}

type DatastreamOptions struct {
	CacheTTL    time.Duration
	CacheConfig CacheConfig
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
	Addr               string
	DB                 int
	Password           string
	MaxRetries         int
	MinRetryBackoff    time.Duration
	MaxRetryBackoff    time.Duration
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	PoolSize           int
	MinIdleConns       int
	MaxConnAge         time.Duration
	PoolTimeout        time.Duration
	IdleTimeout        time.Duration
	IdleCheckFrequency time.Duration
}

func (c RedisCacheConfig) applyDatastreamOptions(opts *DatastreamOptions) {
	opts.CacheConfig = c
}

type RedisCacheClusterConfig struct {
	Addrs              []string
	MaxRedirects       int
	ReadOnly           bool
	RouteByLatency     bool
	RouteRandomly      bool
	Password           string
	MaxRetries         int
	MinRetryBackoff    time.Duration
	MaxRetryBackoff    time.Duration
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	PoolSize           int
	MinIdleConns       int
	MaxConnAge         time.Duration
	PoolTimeout        time.Duration
	IdleTimeout        time.Duration
	IdleCheckFrequency time.Duration
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

	return &ClientOptUseDatastream{
		enabled: true,
		options: dataStreamOptions,
	}
}
