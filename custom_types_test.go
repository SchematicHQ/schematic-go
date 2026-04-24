package schematichq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSDKVersion_NonEmpty(t *testing.T) {
	// Under `go test`, schematic-go is the main module and has no tagged
	// version, so this returns "(devel)" or "unknown" depending on the build.
	// The only contract we enforce here is that it never returns empty string —
	// consumers that build against a tagged release will see that tag instead.
	v := SDKVersion()
	assert.NotEmpty(t, v)
}

func TestSDKVersion_Stable(t *testing.T) {
	// The result is cached via sync.Once; repeated calls must agree.
	assert.Equal(t, SDKVersion(), SDKVersion())
}

func TestSDKName(t *testing.T) {
	assert.Equal(t, "schematic-go", SDKName)
}
