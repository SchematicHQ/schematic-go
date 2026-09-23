package client_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	schematicgo "github.com/schematichq/schematic-go"
	schematicclient "github.com/schematichq/schematic-go/client"
	"github.com/schematichq/schematic-go/mocks"
	option "github.com/schematichq/schematic-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"go.opentelemetry.io/otel/trace"
)

func newTestTracerProvider(t *testing.T) (trace.TracerProvider, func() []sdktrace.ReadOnlySpan) {
	t.Helper()
	exporter := tracetest.NewInMemoryExporter()
	provider := sdktrace.NewTracerProvider(sdktrace.WithSyncer(exporter))
	t.Cleanup(func() { require.NoError(t, provider.Shutdown(context.Background())) })
	// Read lazily: GetSpans snapshots the exporter at call time, so this must
	// run after the spans under test have ended.
	return provider, func() []sdktrace.ReadOnlySpan { return exporter.GetSpans().Snapshots() }
}

func spanNamed(t *testing.T, spans []sdktrace.ReadOnlySpan, name string) sdktrace.ReadOnlySpan {
	t.Helper()
	for _, span := range spans {
		if span.Name() == name {
			return span
		}
	}
	var names []string
	for _, span := range spans {
		names = append(names, span.Name())
	}
	t.Fatalf("no span named %q; got %v", name, names)
	return nil
}

func spanAttrs(span sdktrace.ReadOnlySpan) map[attribute.Key]attribute.Value {
	out := make(map[attribute.Key]attribute.Value)
	for _, kv := range span.Attributes() {
		out[kv.Key] = kv.Value
	}
	return out
}

func tracedFlagResponse(t *testing.T, value bool) *http.Response {
	t.Helper()
	data, err := json.Marshal(&schematicgo.CheckFlagResponse{
		Data: &schematicgo.CheckFlagResponseData{
			Value:     value,
			Reason:    "matched rule",
			CompanyID: schematicgo.String("comp_traced"),
		},
	})
	require.NoError(t, err)
	return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(string(data)))}
}

func tracedClient(t *testing.T, provider trace.TracerProvider, responses ...*http.Response) *schematicclient.SchematicClient {
	t.Helper()
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	for _, resp := range responses {
		mockHTTPClient.EXPECT().Do(gomock.Any()).Return(resp, nil)
	}
	mockHTTPClient.EXPECT().Do(gomock.Any()).Return(tracedFlagResponse(t, false), nil).AnyTimes()

	client := schematicclient.NewSchematicClient(
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(mockHTTPClient),
		option.WithTracerProvider(provider),
		option.WithDisableFlagCheckCache(),
	)
	t.Cleanup(client.Close)
	return client
}

func TestCheckFlagIsTraced(t *testing.T) {
	provider, spans := newTestTracerProvider(t)
	client := tracedClient(t, provider, tracedFlagResponse(t, true))

	ctx, caller := provider.Tracer("test").Start(context.Background(), "handle-request")
	value := client.CheckFlag(ctx, &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"org_id": "acme-internal-id"},
		User:    map[string]string{"email": "someone@example.com"},
	}, "my-flag")
	caller.End()

	assert.True(t, value)

	recorded := spans()
	span := spanNamed(t, recorded, "Schematic.CheckFlag")
	parent := spanNamed(t, recorded, "handle-request")

	assert.Equal(t, parent.SpanContext().SpanID(), span.Parent().SpanID(),
		"the check hangs off the span the caller passed in the context")

	a := spanAttrs(span)
	assert.Equal(t, "my-flag", a["schematic.flag.key"].AsString())
	assert.True(t, a["schematic.flag.value"].AsBool())
	assert.Equal(t, "matched rule", a["schematic.flag.reason"].AsString())
	assert.Equal(t, "comp_traced", a["schematic.company.id"].AsString())
	assert.Equal(t, "api", a["schematic.check.source"].AsString())
}

// A caller's lookup keys are their own identifiers and may be personal data, so
// the span records which dimensions were used and never their values.
func TestCheckFlagRecordsKeyNamesNotKeyValues(t *testing.T) {
	provider, spans := newTestTracerProvider(t)
	client := tracedClient(t, provider, tracedFlagResponse(t, true))

	client.CheckFlag(context.Background(), &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"org_id": "acme-internal-id"},
		User:    map[string]string{"email": "someone@example.com"},
	}, "my-flag")

	span := spanNamed(t, spans(), "Schematic.CheckFlag")
	a := spanAttrs(span)

	assert.Equal(t, "org_id", a["schematic.company.key_names"].AsString())
	assert.Equal(t, "email", a["schematic.user.key_names"].AsString())

	for key, value := range a {
		assert.NotContains(t, value.String(), "someone@example.com",
			"attribute %q leaked a user key value", key)
		assert.NotContains(t, value.String(), "acme-internal-id",
			"attribute %q leaked a company key value", key)
	}
}

func TestCheckFlagOfflineRecordsSource(t *testing.T) {
	provider, spans := newTestTracerProvider(t)
	client := schematicclient.NewSchematicClient(
		option.WithOfflineMode(),
		option.WithTracerProvider(provider),
	)
	defer client.Close()

	client.CheckFlag(context.Background(), &schematicgo.CheckFlagRequestBody{}, "my-flag")

	span := spanNamed(t, spans(), "Schematic.CheckFlag")
	assert.Equal(t, "offline", spanAttrs(span)["schematic.check.source"].AsString())
}

// CheckFlag delegates to the same internal path as CheckFlagWithEntitlement, so
// it must not record a nested duplicate span for the same check. The API call
// underneath it is a separate operation and does get its own span.
func TestCheckFlagRecordsNoDuplicateSpan(t *testing.T) {
	provider, spans := newTestTracerProvider(t)
	client := tracedClient(t, provider, tracedFlagResponse(t, true))

	client.CheckFlag(context.Background(), &schematicgo.CheckFlagRequestBody{}, "my-flag")

	counts := make(map[string]int)
	for _, span := range spans() {
		counts[span.Name()]++
	}

	assert.Equal(t, 1, counts["Schematic.CheckFlag"])
	assert.Zero(t, counts["Schematic.CheckFlagWithEntitlement"],
		"CheckFlag must not nest a duplicate span for the same check")
}

// The API call a check makes gets its own client span, so time on the wire is
// separable from the check wrapped around it.
func TestCheckFlagAPICallIsItsOwnClientSpan(t *testing.T) {
	provider, spans := newTestTracerProvider(t)
	client := tracedClient(t, provider, tracedFlagResponse(t, true))

	client.CheckFlag(context.Background(), &schematicgo.CheckFlagRequestBody{}, "my-flag")

	recorded := spans()
	apiSpan := spanNamed(t, recorded, "Schematic.API.CheckFlag")
	outer := spanNamed(t, recorded, "Schematic.CheckFlag")

	assert.Equal(t, trace.SpanKindClient, apiSpan.SpanKind(),
		"work crossing the network is a client span")
	assert.Equal(t, trace.SpanKindInternal, outer.SpanKind(),
		"the SDK's own wrapper is internal")
	assert.Equal(t, outer.SpanContext().SpanID(), apiSpan.Parent().SpanID(),
		"the API call nests under the check that made it")
	assert.Equal(t, "my-flag", spanAttrs(apiSpan)["schematic.flag.key"].AsString())
}

func TestCheckFlagWithEntitlementIsTraced(t *testing.T) {
	provider, spans := newTestTracerProvider(t)
	client := tracedClient(t, provider, tracedFlagResponse(t, true))

	_, err := client.CheckFlagWithEntitlement(context.Background(), &schematicgo.CheckFlagRequestBody{}, "my-flag")
	require.NoError(t, err)

	span := spanNamed(t, spans(), "Schematic.CheckFlagWithEntitlement")
	assert.Equal(t, "my-flag", spanAttrs(span)["schematic.flag.key"].AsString())
}

func TestCheckFlagsIsTraced(t *testing.T) {
	provider, spans := newTestTracerProvider(t)
	client := tracedClient(t, provider)

	client.CheckFlags(context.Background(), &schematicgo.CheckFlagRequestBody{}, []string{"flag-a", "flag-b"})

	recorded := spans()
	span := spanNamed(t, recorded, "Schematic.CheckFlags")
	a := spanAttrs(span)
	assert.Equal(t, []string{"flag-a", "flag-b"}, a["schematic.flag.keys"].AsStringSlice())
	assert.Equal(t, "api", a["schematic.check.source"].AsString(),
		"a bulk check answered by the API says so")

	// The whole set goes over the API in one call, so it records one client
	// span rather than one per key.
	var apiSpans int
	for _, s := range recorded {
		if s.Name() == "Schematic.API.CheckFlags" {
			apiSpans++
		}
	}
	assert.Equal(t, 1, apiSpans, "one API call for the set, not one per key")
}

func TestCheckIsTraced(t *testing.T) {
	provider, spans := newTestTracerProvider(t)
	client := tracedClient(t, provider, tracedFlagResponse(t, true))

	result := client.Check(context.Background(), &schematicgo.CheckFlagRequestBody{}, "my-flag")
	require.NotNil(t, result)

	span := spanNamed(t, spans(), "Schematic.Check")
	a := spanAttrs(span)
	assert.Equal(t, "my-flag", a["schematic.flag.key"].AsString())
	assert.True(t, a["schematic.check.allowed"].AsBool())
	assert.True(t, a["schematic.flag.value"].AsBool())
	assert.False(t, a["schematic.check.fail_open"].AsBool())
}

func TestIdentifyIsTraced(t *testing.T) {
	provider, spans := newTestTracerProvider(t)
	client := tracedClient(t, provider)

	client.Identify(context.Background(), &schematicgo.EventBodyIdentify{
		Keys: map[string]string{"email": "someone@example.com"},
		Company: &schematicgo.EventBodyIdentifyCompany{
			Keys: map[string]string{"org_id": "acme"},
		},
	})

	span := spanNamed(t, spans(), "Schematic.Identify")
	a := spanAttrs(span)
	assert.Equal(t, "identify", a["schematic.event.type"].AsString())
	assert.Equal(t, "email", a["schematic.user.key_names"].AsString())
	assert.Equal(t, "org_id", a["schematic.company.key_names"].AsString())
}

func TestTrackIsTraced(t *testing.T) {
	provider, spans := newTestTracerProvider(t)
	client := tracedClient(t, provider)

	client.Track(context.Background(), &schematicgo.EventBodyTrack{
		Event:    "api-request",
		Company:  map[string]string{"org_id": "acme"},
		Quantity: schematicgo.Int64(3),
	})

	span := spanNamed(t, spans(), "Schematic.Track")
	a := spanAttrs(span)
	assert.Equal(t, "track", a["schematic.event.type"].AsString())
	assert.Equal(t, "api-request", a["schematic.event.subtype"].AsString())
	assert.Equal(t, int64(3), a["schematic.usage.actual_quantity"].AsInt64())
}

// An application may build a TracerProvider without ever installing it with
// otel.SetTracerProvider. The global is then a no-op, so there is nothing for
// the SDK to record on — except the span the caller passed, which knows which
// provider made it. Borrowing that is what lets this work with no SDK
// configuration at all.
func TestTracesOnProviderNeverInstalledGlobally(t *testing.T) {
	provider, spans := newTestTracerProvider(t)
	// Deliberately no otel.SetTracerProvider, and no option.WithTracerProvider:
	// WithTracing alone, with the provider coming off the caller's span.
	client := schematicclient.NewSchematicClient(
		option.WithOfflineMode(),
		option.WithTracing(),
	)
	defer client.Close()

	ctx, caller := provider.Tracer("app").Start(context.Background(), "handle-request")
	client.CheckFlag(ctx, &schematicgo.CheckFlagRequestBody{}, "my-flag")
	caller.End()

	recorded := spans()
	span := spanNamed(t, recorded, "Schematic.CheckFlag")
	parent := spanNamed(t, recorded, "handle-request")
	assert.Equal(t, parent.SpanContext().SpanID(), span.Parent().SpanID())
	assert.Equal(t, parent.SpanContext().TraceID(), span.SpanContext().TraceID())
}

// An explicitly configured provider must win over the one the caller's span
// carries, so that a caller routing the SDK's spans somewhere specific is not
// silently overridden by their own ambient tracing.
func TestConfiguredProviderWinsOverCallerSpan(t *testing.T) {
	schematicProvider, schematicSpans := newTestTracerProvider(t)
	appProvider, appSpans := newTestTracerProvider(t)

	client := schematicclient.NewSchematicClient(
		option.WithOfflineMode(),
		option.WithTracerProvider(schematicProvider),
	)
	defer client.Close()

	ctx, caller := appProvider.Tracer("app").Start(context.Background(), "handle-request")
	client.CheckFlag(ctx, &schematicgo.CheckFlagRequestBody{}, "my-flag")
	caller.End()

	spanNamed(t, schematicSpans(), "Schematic.CheckFlag")
	for _, span := range appSpans() {
		assert.NotEqual(t, "Schematic.CheckFlag", span.Name(),
			"the SDK's span went to the caller's provider instead of the configured one")
	}
}

// A call with no span in the context has no provider to borrow, so it falls
// through to the global. This is the path a cron job or worker takes.
func TestFallsBackToGlobalWithoutCallerSpan(t *testing.T) {
	provider, spans := newTestTracerProvider(t)
	// Not restored afterwards, and deliberately so: OpenTelemetry refuses a
	// SetTracerProvider call that passes the global provider back to itself, so
	// the obvious save-and-restore is a silent no-op. This is the only test in
	// the package that installs a global, and the rest do not read it, so
	// leaving it set cannot affect them.
	otel.SetTracerProvider(provider)

	client := schematicclient.NewSchematicClient(option.WithOfflineMode(), option.WithTracing())
	defer client.Close()

	client.CheckFlag(context.Background(), &schematicgo.CheckFlagRequestBody{}, "my-flag")

	spanNamed(t, spans(), "Schematic.CheckFlag")
}

// The tracer lookup must stay off the path of a client that is not tracing:
// with no global provider configured and no span in the context, resolution
// falls straight through to the tracer resolved at construction.
func BenchmarkCheckFlagTracingDisabled(b *testing.B) {
	client := schematicclient.NewSchematicClient(option.WithOfflineMode())
	defer client.Close()

	ctx := context.Background()
	evalCtx := &schematicgo.CheckFlagRequestBody{}

	b.ReportAllocs()
	for b.Loop() {
		client.CheckFlag(ctx, evalCtx, "my-flag")
	}
}

// Track and TrackWithReservation share emitTrack, so the span has to live on
// the public entry points rather than on the shared helper. Otherwise settling
// a reservation records a Schematic.Track span nested inside the
// Schematic.TrackWithReservation span for the same event.
func TestTrackWithReservationRecordsNoNestedTrackSpan(t *testing.T) {
	provider, spans := newTestTracerProvider(t)
	client := tracedClient(t, provider)

	client.TrackWithReservation(context.Background(), &schematicclient.Reservation{
		ID:           "res_1",
		CompanyID:    "comp_1",
		CreditTypeID: "credit_1",
		EventSubtype: "api-request",
	}, 3)

	counts := make(map[string]int)
	for _, span := range spans() {
		counts[span.Name()]++
	}

	assert.Equal(t, 1, counts["Schematic.TrackWithReservation"])
	assert.Zero(t, counts["Schematic.Track"],
		"settling a reservation must not nest a Track span for the same event")
}

// Tracing is off unless asked for. Spans cost money at most vendors, so
// upgrading the SDK must not start emitting them in an application that already
// uses OpenTelemetry.
func TestTracingIsOffByDefault(t *testing.T) {
	provider, spans := newTestTracerProvider(t)
	client := schematicclient.NewSchematicClient(option.WithOfflineMode())
	defer client.Close()

	// A caller span in the context is exactly the case that would otherwise
	// hand the SDK a provider to record on.
	ctx, caller := provider.Tracer("app").Start(context.Background(), "handle-request")
	client.CheckFlag(ctx, &schematicgo.CheckFlagRequestBody{}, "my-flag")
	client.Track(ctx, &schematicgo.EventBodyTrack{Event: "api-request"})
	caller.End()

	for _, span := range spans() {
		assert.NotContains(t, span.Name(), "Schematic.",
			"tracing must stay off until WithTracing or WithTracerProvider asks for it")
	}
}

// The global provider is the other way tracing could start on its own.
func TestTracingIsOffByDefaultWithGlobalProvider(t *testing.T) {
	provider, spans := newTestTracerProvider(t)
	otel.SetTracerProvider(provider)

	client := schematicclient.NewSchematicClient(option.WithOfflineMode())
	defer client.Close()

	client.CheckFlag(context.Background(), &schematicgo.CheckFlagRequestBody{}, "my-flag")

	assert.Empty(t, spans(), "a configured global provider does not by itself turn tracing on")
}

// Passing a provider is unambiguously a request to trace, so it must work
// without also passing WithTracing.
func TestTracerProviderAloneEnablesTracing(t *testing.T) {
	provider, spans := newTestTracerProvider(t)
	client := schematicclient.NewSchematicClient(
		option.WithOfflineMode(),
		option.WithTracerProvider(provider),
	)
	defer client.Close()

	client.CheckFlag(context.Background(), &schematicgo.CheckFlagRequestBody{}, "my-flag")

	spanNamed(t, spans(), "Schematic.CheckFlag")
}
