package schematic

import (
	"context"
	"time"
)

type ClientOpt interface {
	Apply(context.Context, Client) error
}

// Override API client host

type ClientOptAPIHost struct {
	host string
}

func (c ClientOptAPIHost) Apply(ctx context.Context, client Client) error {
	client.SetAPIHost(ctx, c.host)

	return nil
}

func WithAPIHost(host string) ClientOpt {
	return ClientOptAPIHost{host}
}

// Specify local cache behavior

type ClientOptLocalFlagCheckCache struct {
	maxSize int
	ttl     time.Duration
}

func (c ClientOptLocalFlagCheckCache) Apply(ctx context.Context, client Client) error {
	client.AddFlagCheckCacheProvider(ctx, newLocalCache[bool](c.maxSize, c.ttl))

	return nil
}

func WithLocalFlagCheckCache(maxSize int, ttl time.Duration) ClientOpt {
	return ClientOptLocalFlagCheckCache{maxSize: maxSize, ttl: ttl}
}

// Disable local cache entirely

type ClientOptDisableFlagCheckCache struct {
}

func (c ClientOptDisableFlagCheckCache) Apply(ctx context.Context, client Client) error {
	client.AddFlagCheckCacheProvider(ctx, newLocalCache[bool](0, 0))

	return nil
}

func WithDisableFlagCheckCache() ClientOpt {
	return ClientOptDisableFlagCheckCache{}
}

// Specify default flag values

type ClientOptFlagDefaults struct {
	values map[string]bool
}

func (c ClientOptFlagDefaults) Apply(ctx context.Context, client Client) error {
	client.SetFlagDefaults(c.values)

	return nil
}

func WithDefaultFlagValues(values map[string]bool) ClientOpt {
	return ClientOptFlagDefaults{values: values}
}

// Event buffer period

type ClientOptEventBufferPeriod struct {
	period time.Duration
}

func (c ClientOptEventBufferPeriod) Apply(ctx context.Context, client Client) error {
	client.SetEventBufferPeriod(c.period)

	return nil
}

func WithEventBufferPeriod(period time.Duration) ClientOpt {
	return ClientOptEventBufferPeriod{period: period}
}
