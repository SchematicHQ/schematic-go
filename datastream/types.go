package datastream

import (
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	schematicdatastreamws "github.com/schematichq/schematic-datastream-ws"
	"github.com/schematichq/schematic-go/cache"
	"github.com/schematichq/schematic-go/core"
	"github.com/schematichq/schematic-go/rulesengine"
	"github.com/schematichq/schematic-go/tracing"
	"go.opentelemetry.io/otel/trace"
)

type CompanyCacheProvider cache.CacheProvider[*rulesengine.Company]
type FlagCacheProvider cache.CacheProvider[*rulesengine.Flag]
type UserCacheProvider cache.CacheProvider[*rulesengine.User]

type DataStreamClientOptions struct {
	ApiKey       string
	BaseURL      string
	Logger       core.Logger
	CompanyCache cache.CacheProvider[*rulesengine.Company]
	UserCache    cache.CacheProvider[*rulesengine.User]
	FlagCache    cache.CacheProvider[*rulesengine.Flag]
	// TracerProvider is the provider option.WithTracerProvider named, or nil,
	// and TracingEnabled is whether option.WithTracing opted in. Together they
	// build the tracing.Resolver the rules-engine spans are recorded on.
	TracerProvider trace.TracerProvider
	TracingEnabled bool
}

type DataStreamClient struct {
	// tracers picks which tracer the rules-engine spans are recorded on.
	tracers            *tracing.Resolver
	cacheTTL           time.Duration
	wsClient           *schematicdatastreamws.Client
	logger             core.Logger
	companyCache       *resourceCache[*rulesengine.Company]
	userCache          *resourceCache[*rulesengine.User]
	flagsCacheProvider FlagCacheProvider
	apiKey             string
	// redisClient is the client every cache provider shares, or nil when the
	// caches are local. Kept so credit leases can reuse it rather than open a
	// second connection pool to the same Redis.
	redisClient redis.UniversalClient

	// engine evaluates flags locally. It owns a compiled WebAssembly module and a
	// pool of instances, so it is built once per client and closed alongside it.
	engine *rulesengine.Engine

	pendingCompanyRequests map[string][]chan *rulesengine.Company
	pendingUserRequests    map[string][]chan *rulesengine.User
	pendingFlagRequest     chan bool

	// snapshotFlagKeys collects the cache keys a paginated flags snapshot has
	// written so far, so the delete of everything absent from the snapshot runs
	// against the whole set rather than the last page. Guarded by flagsMu, and
	// reset when a snapshot's first page arrives -- an interrupted snapshot
	// therefore leaves nothing for the next one to inherit.
	snapshotFlagKeys []string
	// flagPageSize is how many flags the server is asked to put in one message
	// of the flags snapshot, from core.WithFlagPageSize or the SDK default.
	flagPageSize int
	// snapshotNextPage is the page number the snapshot in progress expects
	// next, so a gap is caught rather than applied. The transport drops a
	// message when its queue is full, and a snapshot is now many messages: a
	// lost middle page would otherwise let the last page complete a snapshot
	// missing flags, and the delete of everything absent from it would remove
	// them. Guarded by flagsMu.
	//
	// Zero means no snapshot is in progress, which is the value a client holds
	// before its first snapshot and the one it returns to after a snapshot
	// completes or is abandoned. Only page 1 is accepted in that state -- it
	// sets the field to 2 and starts the set over -- so a page arriving without
	// its snapshot's beginning is out of sequence and abandoned like any other
	// gap. That matters because the server sends the whole snapshot before any
	// update: a stray page means the pages that preceded it were dropped, not
	// that they are still coming.
	snapshotNextPage int

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
	replicatorMu     sync.RWMutex // For replicator state operations
}
