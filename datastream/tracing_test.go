package datastream_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	schematicdatastreamws "github.com/schematichq/schematic-datastream-ws"
	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/core"

	"github.com/schematichq/schematic-go/datastream"
	"github.com/schematichq/schematic-go/rulesengine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"go.opentelemetry.io/otel/trace"
)

// The rules-engine evaluation is local, but it is the one place in a
// DataStream-served check where time can go unexplained, so it records its own
// span under whatever the caller passed.
func TestEvaluateFlagRecordsSpan(t *testing.T) {
	exporter := tracetest.NewInMemoryExporter()
	provider := sdktrace.NewTracerProvider(sdktrace.WithSyncer(exporter))
	t.Cleanup(func() { require.NoError(t, provider.Shutdown(context.Background())) })

	client := datastream.NewDataStreamClient(datastream.DataStreamClientOptions{
		ApiKey:         "test",
		Logger:         NewMockLogger(),
		BaseURL:        "http://invalid.example.com",
		TracerProvider: provider,
	}, nil)
	t.Cleanup(client.Close)

	ctx, caller := provider.Tracer("app").Start(context.Background(), "handle-request")
	_, _ = client.EvaluateFlag(ctx, &rulesengine.Flag{Key: "my-flag"}, nil, nil)
	caller.End()

	var engineSpan, parent sdktrace.ReadOnlySpan
	for _, span := range exporter.GetSpans().Snapshots() {
		switch span.Name() {
		case "Schematic.RulesEngine.CheckFlag":
			engineSpan = span
		case "handle-request":
			parent = span
		}
	}
	require.NotNil(t, engineSpan, "rules engine evaluation recorded no span")
	require.NotNil(t, parent)

	assert.Equal(t, trace.SpanKindInternal, engineSpan.SpanKind(), "evaluation is local work")
	assert.Equal(t, parent.SpanContext().SpanID(), engineSpan.Parent().SpanID())

	var flagKey string
	for _, kv := range engineSpan.Attributes() {
		if kv.Key == "schematic.flag.key" {
			flagKey = kv.Value.AsString()
		}
	}
	assert.Equal(t, "my-flag", flagKey)
}

// newTracedCheckFlagsClient mirrors newCheckFlagsTestClient, with a provider so
// the spans can be read back.
func newTracedCheckFlagsClient(t *testing.T, provider trace.TracerProvider) *datastream.DataStreamClient {
	t.Helper()

	server, incomingMessages, outgoingMessages := setupMockWebSocketServer()
	t.Cleanup(func() {
		server.Close()
		close(outgoingMessages)
	})

	options := createTestClientOptions(server.URL, NewMockLogger(), "test-api-key")
	options.TracerProvider = provider

	client := datastream.NewDataStreamClient(options, &core.DatastreamOptions{CacheTTL: 5 * time.Minute})
	client.Start()
	t.Cleanup(client.Close)

	time.Sleep(300 * time.Millisecond)

	go func() {
		for msg := range incomingMessages {
			var req schematicdatastreamws.DataStreamBaseReq
			_ = json.Unmarshal([]byte(msg), &req)
			if req.Data.EntityType == schematicdatastreamws.EntityTypeCompany {
				outgoingMessages <- createMockCompanyData(req.Data.Keys["company_id"], schematicdatastreamws.MessageTypeFull)
			}
		}
	}()

	return client
}

// A bulk check crosses into the rules engine once for the whole set, and the
// span has to say the same thing. One span per flag would report the same
// interval several times over and would multiply the cost of tracing a check by
// the number of flags in it.
func TestCheckFlagsRecordsOneBulkEngineSpan(t *testing.T) {
	exporter := tracetest.NewInMemoryExporter()
	provider := sdktrace.NewTracerProvider(sdktrace.WithSyncer(exporter))
	t.Cleanup(func() { require.NoError(t, provider.Shutdown(context.Background())) })

	client := newTracedCheckFlagsClient(t, provider)

	ctx, caller := provider.Tracer("app").Start(context.Background(), "handle-request")
	keys := []string{"test-flag-1", "test-flag-2"}
	results, err := client.CheckFlags(ctx, &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"company_id": "123"},
	}, keys)
	require.NoError(t, err)
	require.Len(t, results, len(keys))
	caller.End()

	counts := make(map[string]int)
	var bulk sdktrace.ReadOnlySpan
	for _, span := range exporter.GetSpans().Snapshots() {
		counts[span.Name()]++
		if span.Name() == "Schematic.RulesEngine.CheckFlags" {
			bulk = span
		}
	}

	assert.Equal(t, 1, counts["Schematic.RulesEngine.CheckFlags"], "one span for the set")
	assert.Zero(t, counts["Schematic.RulesEngine.CheckFlag"],
		"a bulk check must not fan out into per-flag engine spans")

	require.NotNil(t, bulk)
	var flagKeys []string
	for _, kv := range bulk.Attributes() {
		if kv.Key == "schematic.flag.keys" {
			flagKeys = kv.Value.AsStringSlice()
		}
	}
	assert.Equal(t, keys, flagKeys, "the span names the flags that were evaluated")
}
