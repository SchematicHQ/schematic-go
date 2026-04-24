package schematichq

import (
	"runtime/debug"
	"sync"

	cache "github.com/schematichq/schematic-go/cache"
	core "github.com/schematichq/schematic-go/core"
)

// CheckFlagResponseCacheProvider is a cache provider for CheckFlagResponse values.
type CheckFlagResponseCacheProvider = cache.CacheProvider[*core.CheckFlagResponse]

// SDKName is the client identifier sent in datastream handshake headers so the
// backend can tell the direct-SDK path apart from schematic-datastream-replicator.
const SDKName = "schematic-go"

const sdkModulePath = "github.com/schematichq/schematic-go"

var (
	sdkVersionOnce sync.Once
	sdkVersion     string
)

// SDKVersion returns the module version of schematic-go recorded in the build.
// For consumers importing a tagged release (e.g. v1.4.6) this returns that tag
// — which is what the Fern generate action stamps when publishing the SDK.
// For builds against an untagged checkout or a replace directive it returns
// "(devel)"; if build info is unavailable it returns "unknown".
func SDKVersion() string {
	sdkVersionOnce.Do(func() {
		sdkVersion = resolveSDKVersion()
	})
	return sdkVersion
}

func resolveSDKVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}
	if info.Main.Path == sdkModulePath && info.Main.Version != "" {
		return info.Main.Version
	}
	for _, dep := range info.Deps {
		if dep == nil || dep.Path != sdkModulePath {
			continue
		}
		if dep.Replace != nil && dep.Replace.Version != "" {
			return dep.Replace.Version
		}
		if dep.Version != "" {
			return dep.Version
		}
	}
	return "unknown"
}
