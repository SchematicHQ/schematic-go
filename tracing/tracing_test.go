package tracing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// The instrumentation version is read from build info rather than held in a
// constant, so it cannot drift from the released version. Inside this module's
// own tests there is no dependency entry to read, which is the case that must
// degrade to no version rather than to a wrong one.
func TestModuleVersionNeverReportsAStaleConstant(t *testing.T) {
	version := moduleVersion()

	// Whatever it resolves to, it must not be a hardcoded value that release
	// tagging would leave behind.
	assert.NotEqual(t, "(devel)", version, "a local build reports no version at all")

	if version == "" {
		assert.Nil(t, buildTracerOptions(),
			"with no version to report the scope carries no version option")
		return
	}
	assert.NotEmpty(t, buildTracerOptions())
}
