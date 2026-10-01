package datastream

import "time"

// Time constants for the WebSocket connection
const (
	resourceTimeout = 2 * time.Second
)

// defaultFlagSnapshotPageSize is how many flags we ask the server to put in one
// message of the flags snapshot. The snapshot used to arrive whole, and for an
// ordinary set of flags and rules that ran past a megabyte -- over the frame
// limit for some hosts, and compressible only by a constant factor.
//
// Counted in flags rather than bytes, so it bounds a page's flag count and not
// its size; one flag carrying enough rules can still overflow a frame on its
// own. core.WithFlagPageSize overrides it per client.
const defaultFlagSnapshotPageSize = 100

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
