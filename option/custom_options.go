package option

import core "github.com/schematichq/schematic-go/core"

// These custom options must be defined in the core package in order to be able
// to fulfill the options interface, but we alias them here so that clients can
// import all of their options from a single package.

var WithDefaultFlagValues = core.WithDefaultFlagValues
var WithDisableFlagCheckCache = core.WithDisableFlagCheckCache
var WithEventBufferPeriod = core.WithEventBufferPeriod
var WithFlagCheckCacheProvider = core.WithFlagCheckCacheProvider
var WithLocalFlagCheckCache = core.WithLocalFlagCheckCache
var WithOfflineMode = core.WithOfflineMode
var WithUseDatastream = core.WithUseDatastream
