package tracing

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

// Resolver picks the tracer to record an operation on.
//
// It exists so that every part of the SDK resolves a tracer the same way, and
// so that the ordering below is documented in one place rather than repeated
// wherever a span is started.
type Resolver struct {
	// configured is the tracer for the provider option.WithTracerProvider
	// named, and is nil when the caller configured none.
	configured trace.Tracer
	// global is the last-resort fallback, resolved once. The global provider
	// hands back a tracer that delegates to whatever is installed later, so
	// resolving it up front is safe and keeps the lookup off the per-call path
	// of a client that is not tracing.
	global trace.Tracer
}

// NewResolver builds a Resolver for an optionally configured provider. A nil
// provider means the caller configured none, which is the ordinary case.
func NewResolver(provider trace.TracerProvider) *Resolver {
	r := &Resolver{global: tracer(nil)}
	if provider != nil {
		r.configured = tracer(provider)
	}
	return r
}

// For returns the tracer to record one operation on, in three steps:
//
//  1. The provider option.WithTracerProvider configured, if there was one. An
//     explicit instruction always wins: a caller who deliberately routes the
//     SDK's spans to a particular provider must not have that quietly
//     overridden by whatever their context happens to carry.
//  2. Otherwise the provider behind the caller's own span. This is what lets an
//     application that built a TracerProvider but never installed it with
//     otel.SetTracerProvider still see the SDK's work: the span they passed in
//     the context knows where it came from, so the SDK can record alongside it
//     with nothing configured at all.
//  3. Otherwise the global provider, whose default is a no-op. This is the
//     ordinary case, and the only step that can serve a call whose context
//     carries no span — a cron job or a worker calling in on
//     context.Background() has nothing to borrow a provider from.
//
// This mirrors what otelhttp does. Only step 2 is resolved per call, since only
// it depends on the context.
func (r *Resolver) For(ctx context.Context) trace.Tracer {
	if r.configured != nil {
		return r.configured
	}
	if span := trace.SpanFromContext(ctx); span.SpanContext().IsValid() {
		return tracer(span.TracerProvider())
	}
	return r.global
}

// InternalSpanOptions starts a span for work happening inside the caller's
// process, such as a flag evaluation or the SDK's own bookkeeping.
var InternalSpanOptions = []trace.SpanStartOption{trace.WithSpanKind(trace.SpanKindInternal)}

// ClientSpanOptions starts a span for work that leaves the process — an API
// call, and anything else crossing the network. Keeping these distinct from
// InternalSpanOptions is what lets a reader tell the SDK's own wrappers from
// the time actually spent on the wire.
var ClientSpanOptions = []trace.SpanStartOption{trace.WithSpanKind(trace.SpanKindClient)}
