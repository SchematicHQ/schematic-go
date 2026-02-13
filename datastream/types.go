package datastream

import (
	"sync"
	"time"

	"github.com/schematichq/rulesengine"
	schematicdatastreamws "github.com/schematichq/schematic-datastream-ws"
	"github.com/schematichq/schematic-go/cache"
	"github.com/schematichq/schematic-go/core"
)

type CompanyCacheProvider cache.CacheProvider[*rulesengine.Company]
type CompanyLookupCacheProvider cache.CacheProvider[string]
type FlagCacheProvider cache.CacheProvider[*rulesengine.Flag]
type UserCacheProvider cache.CacheProvider[*rulesengine.User]

type DataStreamClientOptions struct {
	ApiKey       string
	BaseURL      string
	Logger       core.Logger
	CompanyCache cache.CacheProvider[*rulesengine.Company]
	UserCache    cache.CacheProvider[*rulesengine.User]
	FlagCache    cache.CacheProvider[*rulesengine.Flag]
}

type DataStreamClient struct {
	cacheTTL                   time.Duration
	wsClient                   *schematicdatastreamws.Client
	logger                     core.Logger
	companyCacheProvider       CompanyCacheProvider
	companyLookupCacheProvider CompanyLookupCacheProvider
	flagsCacheProvider         FlagCacheProvider
	userCacheProvider          UserCacheProvider
	companyCache               map[string]*rulesengine.Company
	apiKey                     string

	pendingCompanyRequests map[string][]chan *rulesengine.Company
	pendingUserRequests    map[string][]chan *rulesengine.User
	pendingFlagRequest     chan bool

	// Connection state tracking
	connected bool

	// Replicator mode configuration
	replicatorMode         bool
	replicatorHealthURL    string
	replicatorHealthCheck  time.Duration
	replicatorReady        bool
	replicatorCacheVersion string
	replicatorHealthDone   chan bool

	// Locks
	flagsMu          sync.RWMutex // For flags cache operations
	companyMu        sync.RWMutex // For company cache operations
	userMu           sync.RWMutex // For user cache operations
	pendingCompReqMu sync.Mutex   // For pending company request operations
	pendingUserReqMu sync.Mutex   // For pending user request operations
	pendingFlagReqMu sync.Mutex   // For pending flag request operations
	writeMu          sync.Mutex   // Existing mutex for WebSocket writes
	connectedMu      sync.RWMutex // For connection state operations
	replicatorMu     sync.RWMutex // For replicator state operations
}
