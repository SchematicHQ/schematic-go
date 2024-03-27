package schematic

import (
	"context"
	"time"
)

const defaultCacheSize = 1000
const defaultCacheTTL = 5 * time.Minute

type ClientOpt interface {
	Apply(context.Context, Client) error
	Type() ClientOptType
}

type ClientOptType string

const (
	ClientOptTypeAPIHost      ClientOptType = "api_host"
	ClientOptTypeDisableCache ClientOptType = "disable_cache"
	ClientOptTypeLocalCache   ClientOptType = "local_cache"
)

// Override API client host

type ClientOptAPIHost struct {
	host string
}

func (c ClientOptAPIHost) Apply(ctx context.Context, client Client) error {
	client.SetAPIHost(ctx, c.host)

	return nil
}

func (c ClientOptAPIHost) Type() ClientOptType {
	return ClientOptTypeDisableCache
}

func WithAPIHost(host string) ClientOpt {
	return ClientOptAPIHost{host}
}

// Specify local cache behavior

type ClientOptLocalCache struct {
	maxSize int
	ttl     time.Duration
}

func (c ClientOptLocalCache) Apply(ctx context.Context, client Client) error {
	client.AddCacheProvider(ctx, newLocalCache(c.maxSize, c.ttl))

	return nil
}

func (c ClientOptLocalCache) Type() ClientOptType {
	return ClientOptTypeLocalCache
}

func WithLocalCache(maxSize int, ttl time.Duration) ClientOpt {
	return ClientOptLocalCache{maxSize: maxSize, ttl: ttl}
}

// Disable local cache entirely

type ClientOptDisableCache struct {
}

func (c ClientOptDisableCache) Apply(ctx context.Context, client Client) error {
	client.AddCacheProvider(ctx, newLocalCache(0, 0))

	return nil
}

func (c ClientOptDisableCache) Type() ClientOptType {
	return ClientOptTypeDisableCache
}

func WithDisableCache() ClientOpt {
	return ClientOptDisableCache{}
}
