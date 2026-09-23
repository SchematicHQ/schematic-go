package client_test

import (
	"context"
	"net/http"
	"testing"

	schematicgo "github.com/schematichq/schematic-go"
	schematicclient "github.com/schematichq/schematic-go/client"
	option "github.com/schematichq/schematic-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/trace"
)

// instrumentedHTTPClient stands in for a caller's own otelhttp-wrapped client.
type instrumentedHTTPClient struct {
	next   *mockRespClient
	tracer trace.Tracer
}

type mockRespClient struct{ t *testing.T }

func (m *mockRespClient) Do(*http.Request) (*http.Response, error) {
	return tracedFlagResponse(m.t, true), nil
}

func (c *instrumentedHTTPClient) Do(req *http.Request) (*http.Response, error) {
	// What otelhttp does: start a span from the request's context.
	_, span := c.tracer.Start(req.Context(), "HTTP POST", trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()
	return c.next.Do(req)
}

func TestUserInstrumentedHTTPClientNestsUnderSchematicSpans(t *testing.T) {
	provider, spans := newTestTracerProvider(t)

	client := schematicclient.NewSchematicClient(
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(&instrumentedHTTPClient{
			next:   &mockRespClient{t: t},
			tracer: provider.Tracer("user-app"),
		}),
		option.WithTracerProvider(provider),
		option.WithDisableFlagCheckCache(),
	)
	defer client.Close()

	client.CheckFlag(context.Background(), &schematicgo.CheckFlagRequestBody{}, "my-flag")

	recorded := spans()
	httpSpan := spanNamed(t, recorded, "HTTP POST")
	apiSpan := spanNamed(t, recorded, "Schematic.API.CheckFlag")
	outer := spanNamed(t, recorded, "Schematic.CheckFlag")

	require.Equal(t, apiSpan.SpanContext().SpanID(), httpSpan.Parent().SpanID(),
		"the caller's HTTP span nests under the Schematic API span that issued it")
	assert.Equal(t, outer.SpanContext().SpanID(), apiSpan.Parent().SpanID())
	assert.Equal(t, outer.SpanContext().TraceID(), httpSpan.SpanContext().TraceID())
}
