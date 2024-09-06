package client

import (
	"context"
	"time"

	"github.com/schematichq/schematic-go/cache"
	"github.com/schematichq/schematic-go/core"
)

type ClientOpt interface {
	Apply(context.Context, SchematicClient) error
}

// Override API client host

type ClientOptAPIHost struct {
	host string
}

func (c ClientOptAPIHost) Apply(ctx context.Context, client SchematicClient) error {
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

func (c ClientOptLocalFlagCheckCache) Apply(ctx context.Context, client SchematicClient) error {
	client.AddFlagCheckCacheProvider(ctx, cache.NewLocalCache[bool](c.maxSize, c.ttl))

	return nil
}

func WithLocalFlagCheckCache(maxSize int, ttl time.Duration) ClientOpt {
	return ClientOptLocalFlagCheckCache{maxSize: maxSize, ttl: ttl}
}

// Disable local cache entirely

type ClientOptDisableFlagCheckCache struct {
}

func (c ClientOptDisableFlagCheckCache) Apply(ctx context.Context, client SchematicClient) error {
	client.AddFlagCheckCacheProvider(ctx, cache.NewLocalCache[bool](0, 0))

	return nil
}

func WithDisableFlagCheckCache() ClientOpt {
	return ClientOptDisableFlagCheckCache{}
}

// Specify default flag values

type ClientOptFlagDefaults struct {
	values map[string]bool
}

func (c ClientOptFlagDefaults) Apply(ctx context.Context, client SchematicClient) error {
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

func (c ClientOptEventBufferPeriod) Apply(ctx context.Context, client SchematicClient) error {
	client.SetEventBufferPeriod(c.period)

	return nil
}

func WithEventBufferPeriod(period time.Duration) ClientOpt {
	return ClientOptEventBufferPeriod{period: period}
}

type ClientOptOfflineMode struct {
	isOffline bool
}

func (c ClientOptOfflineMode) Apply(ctx context.Context, client SchematicClient) error {
	if c.isOffline {
		client.SetOfflineMode()
	}

	return nil
}

func WithOfflineMode() ClientOpt {
	return ClientOptOfflineMode{
		isOffline: true,
	}
}

type ClientHTTPClient struct {
	HTTPClient core.HTTPClient
}

func (c ClientHTTPClient) Apply(ctx context.Context, client SchematicClient) error {
	client.SetHTTPClient(c.HTTPClient)

	return nil
}

func WithHTTPClient(httpClient core.HTTPClient) ClientOpt {
	return ClientHTTPClient{
		HTTPClient: httpClient,
	}
}
