package option

import core "github.com/schematichq/schematic-go/core"

// These custom options must be defined in the core package in order to be able
// to fulfill the options interface, but we alias them here so that clients can
// import all of their options from a single package.

var WithDefaultFlagValues = core.WithDefaultFlagValues
var WithDisableFlagCheckCache = core.WithDisableFlagCheckCache
var WithEventBufferPeriod = core.WithEventBufferPeriod
var WithEventCaptureBaseURL = core.WithEventCaptureBaseURL
var WithFlagCheckCacheProvider = core.WithFlagCheckCacheProvider
var WithLocalFlagCheckCache = core.WithLocalFlagCheckCache
var WithOfflineMode = core.WithOfflineMode
var WithLogger = core.WithLogger
var WithDatastream = core.WithDatastream
var WithRedisCache = core.WithRedisCache
var WithCacheTTL = core.WithCacheTTL
var WithReplicatorMode = core.WithReplicatorMode
var WithReplicatorHealthURL = core.WithReplicatorHealthURL
var WithReplicatorHealthInterval = core.WithReplicatorHealthInterval

type RedisCacheConfig = core.RedisCacheConfig
type RedisCacheClusterConfig = core.RedisCacheClusterConfig
type LogLevel = core.LogLevel

const (
	LogLevelDebug LogLevel = core.LogLevelDebug
	LogLevelInfo  LogLevel = core.LogLevelInfo
	LogLevelWarn  LogLevel = core.LogLevelWarn
	LogLevelError LogLevel = core.LogLevelError
)
