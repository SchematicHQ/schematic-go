package datastream

import "time"

const defaultBaseURL = "wss://datastream.schematichq.com"

type Action string

const (
	ActionStart Action = "start"
	ActionStop  Action = "stop"
)

type EntityType string

const (
	EntityTypeCompany EntityType = "rulesengine.Company"
	EntityTypeFlags   EntityType = "rulesengine.Flags"
	EntityTypeUser    EntityType = "rulesengine.User"
)

// Time constants for the WebSocket connection
const (
	maxReconnectAttempts = 5
	minReconnectDelay    = 1 * time.Second
	maxReconnectDelay    = 60 * time.Second
	pongWait             = 30 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

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

	RulesEngineVersionKey = "9ecb00ed"
)
