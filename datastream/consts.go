package datastream

import "time"

// Time constants for the WebSocket connection
const (
	resourceTimeout = 2 * time.Second
)

// Cache constants
const (
	defaultTTL       = 24 * time.Hour
	maxCacheTTL      = 30 * 24 * time.Hour // 30 days maximum TTL for cache items
	defaultCacheSize = 1000

	cacheKeyPrefix        = "schematic"
	cacheKeyPrefixCompany = "company"
	cacheKeyPrefixFlags   = "flags"
	cacheKeyPrefixUser    = "user"
)
