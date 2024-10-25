package core

import (
	"time"

	cache "github.com/schematichq/schematic-go/cache"
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
}

func WithOfflineMode() RequestOption {
	return ClientOptOfflineMode{
		isOffline: true,
	}
}
