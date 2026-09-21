package datastream_test

import (
	"context"
	"testing"

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
