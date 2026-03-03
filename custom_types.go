package schematichq

import (
	cache "github.com/schematichq/schematic-go/cache"
	core "github.com/schematichq/schematic-go/core"
)

// CheckFlagResponseCacheProvider is a cache provider for CheckFlagResponse values.
type CheckFlagResponseCacheProvider = cache.CacheProvider[*core.CheckFlagResponse]
