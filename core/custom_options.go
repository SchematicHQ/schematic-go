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

// Datastream options
type DatastreamOption interface {
	applyRequestOptions(opts *DatastreamOptions)
}

type DatastreamOptions struct {
	CacheTTL      time.Duration
	CacheProvider string
}

type CacheTTL struct {
	ttl time.Duration
}

func (c CacheTTL) applyRequestOptions(opts *DatastreamOptions) {
	opts.CacheTTL = c.ttl
}

func WithCacheTTL(ttl time.Duration) DatastreamOption {
	return CacheTTL{ttl: ttl}
}

type CacheProvider struct {
	provider string
}

func (c CacheProvider) applyRequestOptions(opts *DatastreamOptions) {
	opts.CacheProvider = c.provider
}

func WithCacheProvider(provider string) DatastreamOption {
	return CacheProvider{provider: provider}
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
	dataStreamOptons := &DatastreamOptions{
		CacheTTL:      5 * time.Second,
		CacheProvider: "local",
	}

	for _, opt := range opts {
		opt.applyRequestOptions(dataStreamOptons)
	}

	return &ClientOptUseDatastream{
		enabled: true,
		options: dataStreamOptons,
	}
}
