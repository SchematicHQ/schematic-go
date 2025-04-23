package datastream

import "time"

const defaultBaseURL = "api.schematic.com"

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
	reconnectDelay       = 2 * time.Second
	pongWait             = 30 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
)

// Cache constants
const (
	defaultCacheProvider = "local"
	defaultTTL           = 24 * time.Hour

	cacheKeyPrefixCompany = "company"
	cacheKeyPrefixFlags   = "flags"
	cacheKeyPrefixUser    = "user"
)
