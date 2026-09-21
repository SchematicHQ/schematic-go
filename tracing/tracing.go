// Package tracing instruments the SDK with OpenTelemetry.
//
// Nothing here needs configuring. The SDK records on the provider
// option.WithTracerProvider named, or failing that the one behind the caller's
// own span, or failing that the global one. An application that has set
// OpenTelemetry up in any of the usual ways is therefore already covered, and
// one that has not gets the API's default no-op provider, which starts no spans
// and builds no attributes.
//
// See Resolver.For for why the order is what it is.
package tracing

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// instrumentationName identifies this SDK as the source of its spans. It shows
// up on the instrumentation scope of every span the SDK emits, so a consumer
// can tell Schematic's spans from their own.
const instrumentationName = "github.com/schematichq/schematic-go"

// version is the instrumentation version reported alongside instrumentationName.
// It tracks the SDK version in core.RequestOptions.cloneHeader.
const version = "v1.1.1"

// tracer returns the tracer the SDK records spans on. A nil provider falls back
// to the global one, which is what an application that has configured
// OpenTelemetry in the usual way will have set.
func tracer(provider trace.TracerProvider) trace.Tracer {
	if provider == nil {
		provider = otel.GetTracerProvider()
	}
	return provider.Tracer(instrumentationName, trace.WithInstrumentationVersion(version))
}
