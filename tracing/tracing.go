// Package tracing instruments the SDK with OpenTelemetry.
//
// Tracing is off until option.WithTracing or option.WithTracerProvider asks for
// it, because spans cost money at most vendors and upgrading a library should
// not raise an application's observability bill on its own.
//
// Once on, the SDK records on the provider option.WithTracerProvider named, or
// failing that the one behind the caller's own span, or failing that the global
// one. See Resolver.For for why the order is what it is.
package tracing

import (
	"runtime/debug"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// instrumentationName identifies this SDK as the source of its spans. It shows
// up on the instrumentation scope of every span the SDK emits, so a consumer
// can tell Schematic's spans from their own.
const instrumentationName = "github.com/schematichq/schematic-go"

// tracerOptions carries the instrumentation version, when one can be
// determined. Built once: a tracer is resolved per call on the path that reads
// the provider off the caller's span.
var tracerOptions = buildTracerOptions()

func buildTracerOptions() []trace.TracerOption {
	if version := moduleVersion(); version != "" {
		return []trace.TracerOption{trace.WithInstrumentationVersion(version)}
	}
	return nil
}

// moduleVersion reports the version of this module that the running binary was
// built against, read from the build info the Go toolchain embeds.
//
// Reading it rather than holding a constant keeps the reported version correct
// with nothing to update at release time. It returns the empty string when the
// version cannot be determined — a binary built from a local checkout, or with
// module information stripped — and the instrumentation scope then carries no
// version, which is better than carrying a stale one.
func moduleVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	// Set when the SDK is a dependency, which is every case but its own tests.
	for _, dep := range info.Deps {
		if dep.Path == instrumentationName {
			return dep.Version
		}
	}
	if info.Main.Path == instrumentationName && info.Main.Version != "(devel)" {
		return info.Main.Version
	}
	return ""
}

// tracer returns the tracer the SDK records spans on. A nil provider falls back
// to the global one, which is what an application that has configured
// OpenTelemetry in the usual way will have set.
func tracer(provider trace.TracerProvider) trace.Tracer {
	if provider == nil {
		provider = otel.GetTracerProvider()
	}
	return provider.Tracer(instrumentationName, tracerOptions...)
}
