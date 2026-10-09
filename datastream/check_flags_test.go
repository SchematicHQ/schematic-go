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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// CheckFlags evaluates a set of flags in one crossing into the rules engine.
// What these pin is that batching changes nothing but the crossing --
// the verdicts must be the ones CheckFlag gives key by key -- and that the
// cases where a batch cannot answer fail the whole call rather than returning
// a short list a caller would misread as negative verdicts.

// newCheckFlagsTestClient starts a client against a mock datastream that
// answers company requests, mirroring the setup the CheckFlag tests use.
func newCheckFlagsTestClient(t *testing.T) *datastream.DataStreamClient {
	t.Helper()

	server, incomingMessages, outgoingMessages := setupMockWebSocketServer()
	t.Cleanup(func() {
		server.Close()
		close(outgoingMessages)
	})

	logger := NewMockLogger()
	configOptions := &core.DatastreamOptions{CacheTTL: 5 * time.Minute}
	clientOptions := createTestClientOptions(server.URL, logger, "test-api-key")

	client := datastream.NewDataStreamClient(clientOptions, configOptions)
	client.Start()
	t.Cleanup(client.Close)

	// Wait for connection and the initial flag set.
	time.Sleep(300 * time.Millisecond)

	go func() {
		for msg := range incomingMessages {
			var req schematicdatastreamws.DataStreamBaseReq
			_ = json.Unmarshal([]byte(msg), &req)

			if req.Data.EntityType == schematicdatastreamws.EntityTypeCompany {
				companyID := req.Data.Keys["company_id"]
				outgoingMessages <- createMockCompanyData(companyID, schematicdatastreamws.MessageTypeFull)
			}
		}
	}()

	return client
}

// TestCheckFlagsMatchesCheckFlag is the property that matters: a batched
// verdict must be the verdict the per-key path gives, in the order the keys
// were passed.
func TestCheckFlagsMatchesCheckFlag(t *testing.T) {
	client := newCheckFlagsTestClient(t)

	ctx := context.Background()
	evalCtx := &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"company_id": "123"},
	}
	keys := []string{"test-flag-1", "test-flag-2", "test-flag-1"}

	results, err := client.CheckFlags(ctx, evalCtx, keys)
	require.NoError(t, err)
	require.Len(t, results, len(keys), "one result per key, including repeats")

	for i, key := range keys {
		single, err := client.CheckFlag(ctx, evalCtx, key)
		require.NoError(t, err)

		assert.Equal(t, single.Value, results[i].Value, "value for %s", key)
		assert.Equal(t, single.FlagKey, results[i].FlagKey)
		assert.Equal(t, single.Reason, results[i].Reason)
	}

	// The fixture's flags disagree, so a batch that returned one verdict for
	// everything would pass the loop above only if both were the same.
	assert.True(t, results[0].Value, "test-flag-1 is true")
	assert.False(t, results[1].Value, "test-flag-2 is false")
}

// TestCheckFlagsRefusesPreflight pins the refusal. A preflight asks whether one
// action would be allowed; spread over a set it lands on every numeric
// condition and flips flags unrelated to the action, which is why the API
// answers one with a 400 on its multi-flag routes.
func TestCheckFlagsRefusesPreflight(t *testing.T) {
	client := newCheckFlagsTestClient(t)

	ctx := context.Background()
	evalCtx := &schematicgo.CheckFlagRequestBody{
		Company:   map[string]string{"company_id": "123"},
		Preflight: &schematicgo.PreflightRequestBody{Usage: schematicgo.Int64(1)},
	}

	results, err := client.CheckFlags(ctx, evalCtx, []string{"test-flag-1", "test-flag-2"})
	assert.ErrorIs(t, err, datastream.ErrPreflightMultipleFlags)
	assert.Nil(t, results, "a refused batch answers nothing")

	// The single-flag path still takes it: one action, one flag.
	single, err := client.CheckFlag(ctx, evalCtx, "test-flag-1")
	assert.NoError(t, err, "preflight remains supported for a single flag")
	assert.NotNil(t, single)
}

// TestCheckFlagsUnknownFlagFailsTheBatch pins that a missing flag is not
// quietly dropped. A short result list would line up against the wrong keys,
// so the whole call fails and the caller falls back to the API for the set.
func TestCheckFlagsUnknownFlagFailsTheBatch(t *testing.T) {
	client := newCheckFlagsTestClient(t)

	ctx := context.Background()
	evalCtx := &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"company_id": "123"},
	}

	results, err := client.CheckFlags(ctx, evalCtx, []string{"test-flag-1", "non-existent-flag"})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "flag not found")
	assert.Nil(t, results)
}

// TestCheckFlagsEmptyKeys covers the degenerate case: no keys, no engine call,
// an empty result rather than an error.
func TestCheckFlagsEmptyKeys(t *testing.T) {
	client := newCheckFlagsTestClient(t)

	results, err := client.CheckFlags(context.Background(), &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"company_id": "123"},
	}, nil)

	assert.NoError(t, err)
	assert.Empty(t, results)
}

// TestCheckFlagsWithoutCompanyOrUser pins that an eval context naming neither
// entity still evaluates, reporting each flag's default rather than failing.
func TestCheckFlagsWithoutCompanyOrUser(t *testing.T) {
	client := newCheckFlagsTestClient(t)

	results, err := client.CheckFlags(context.Background(), &schematicgo.CheckFlagRequestBody{},
		[]string{"test-flag-1", "test-flag-2"})

	require.NoError(t, err)
	require.Len(t, results, 2)
}
