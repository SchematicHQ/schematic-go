package leases

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/core"
	"github.com/schematichq/schematic-go/credits"
	"github.com/schematichq/schematic-go/mocks"
	"github.com/schematichq/schematic-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// What NewAPIWireClient puts on the wire, and what it does with what comes back.
// Both lease calls are safe to retry, acquire because the server hands back the
// slot's existing lease and extend because it carries an idempotency key; these
// pin that a retried attempt grows the lease once.

type wireRequest struct {
	method string
	path   string
	body   []byte
}

type wireRecorder struct {
	mu       sync.Mutex
	requests []wireRequest
}

func (r *wireRecorder) all() []wireRequest {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]wireRequest(nil), r.requests...)
}

// scriptedResponse is one answer the test server hands back.
type scriptedResponse struct {
	status int
	body   any
}

// newTestWireClient answers every request with status and body, recording what
// was asked.
func newTestWireClient(t *testing.T, status int, body any) (WireClient, *wireRecorder) {
	t.Helper()
	return newScriptedWireClient(t, scriptedResponse{status: status, body: body})
}

// newScriptedWireClient answers each request with the next scripted response,
// staying on the last one once the script runs out, and records what was asked.
func newScriptedWireClient(t *testing.T, script ...scriptedResponse) (WireClient, *wireRecorder) {
	t.Helper()
	require.NotEmpty(t, script)
	recorder := &wireRecorder{}
	ctrl := gomock.NewController(t)
	httpClient := mocks.NewMockHTTPClient(ctrl)
	httpClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		var raw []byte
		if req.Body != nil {
			raw, _ = io.ReadAll(req.Body)
		}
		recorder.mu.Lock()
		attempt := len(recorder.requests)
		recorder.requests = append(recorder.requests, wireRequest{method: req.Method, path: req.URL.Path, body: raw})
		recorder.mu.Unlock()

		if attempt >= len(script) {
			attempt = len(script) - 1
		}
		encoded, err := json.Marshal(script[attempt].body)
		if err != nil {
			return nil, err
		}
		return &http.Response{
			Status:     fmt.Sprint(script[attempt].status),
			StatusCode: script[attempt].status,
			Body:       io.NopCloser(bytes.NewReader(encoded)),
		}, nil
	}).AnyTimes()

	options := core.NewRequestOptions(option.WithAPIKey("test-api-key"), option.WithHTTPClient(httpClient))
	return NewAPIWireClient(credits.NewClient(options)), recorder
}

// idempotencyKeys is the key each recorded request carried, empty string for a
// request that carried none.
func (r *wireRecorder) idempotencyKeys(t *testing.T) []string {
	t.Helper()
	keys := make([]string, 0, len(r.all()))
	for _, request := range r.all() {
		var sent struct {
			IdempotencyKey string `json:"idempotency_key"`
		}
		require.NoError(t, json.Unmarshal(request.body, &sent))
		keys = append(keys, sent.IdempotencyKey)
	}
	return keys
}

func leaseResponseData() *schematicgo.CreditLeaseResponseData {
	return &schematicgo.CreditLeaseResponseData{
		ID:            "lse_1",
		CompanyID:     "co_1",
		CreditTypeID:  "ct_1",
		GrantedAmount: 1000,
		ExpiresAt:     time.Date(2026, 1, 1, 0, 5, 0, 0, time.UTC),
	}
}

func TestAPIWireClientAcquireSendsTheRequestedTrancheAndExpiry(t *testing.T) {
	wire, recorder := newTestWireClient(t, http.StatusOK, &schematicgo.AcquireCreditLeaseResponse{Data: leaseResponseData()})
	expiresAt := time.Date(2026, 1, 1, 0, 5, 0, 0, time.UTC)

	grant, err := wire.Acquire(context.Background(), "co_1", "ct_1", 1000, expiresAt)

	require.NoError(t, err)
	require.NotNil(t, grant)
	assert.Equal(t, "lse_1", grant.LeaseID)
	assert.Equal(t, "co_1", grant.CompanyID)
	assert.Equal(t, "ct_1", grant.CreditTypeID)
	assert.Equal(t, 1000.0, grant.GrantedAmount)
	assert.Equal(t, expiresAt, grant.ExpiresAt)

	requests := recorder.all()
	require.Len(t, requests, 1)
	var sent map[string]any
	require.NoError(t, json.Unmarshal(requests[0].body, &sent))
	assert.Equal(t, "co_1", sent["company_id"])
	assert.Equal(t, "ct_1", sent["credit_type_id"])
	assert.InDelta(t, 1000.0, sent["requested_amount"], 0)
	assert.NotEmpty(t, sent["expires_at"])
}

func TestAPIWireClientExtendSendsTheIncrementNotTheTotal(t *testing.T) {
	data := leaseResponseData()
	data.GrantedAmount = 2000
	wire, recorder := newTestWireClient(t, http.StatusOK, &schematicgo.ExtendCreditLeaseResponse{Data: data})

	grant, err := wire.Extend(context.Background(), "lse_1", 1000, time.Date(2026, 1, 1, 0, 10, 0, 0, time.UTC))

	require.NoError(t, err)
	// The server answers with the authoritative total, which is what the store
	// reconciles against; the request carries only the increment.
	assert.Equal(t, 2000.0, grant.GrantedAmount)

	requests := recorder.all()
	require.Len(t, requests, 1)
	assert.Contains(t, requests[0].path, "lse_1")
	var sent map[string]any
	require.NoError(t, json.Unmarshal(requests[0].body, &sent))
	assert.InDelta(t, 1000.0, sent["additional_amount"], 0)
}

func TestAPIWireClientReleaseNamesTheLease(t *testing.T) {
	wire, recorder := newTestWireClient(t, http.StatusOK, &schematicgo.ReleaseCreditLeaseResponse{Data: leaseResponseData()})

	require.NoError(t, wire.Release(context.Background(), "lse_1"))

	requests := recorder.all()
	require.Len(t, requests, 1)
	assert.Contains(t, requests[0].path, "lse_1")
}

func TestAPIWireClientExtendKeysEveryCallSeparately(t *testing.T) {
	data := leaseResponseData()
	data.GrantedAmount = 2000
	wire, recorder := newTestWireClient(t, http.StatusOK, &schematicgo.ExtendCreditLeaseResponse{Data: data})

	_, err := wire.Extend(context.Background(), "lse_1", 1000, time.Date(2026, 1, 1, 0, 10, 0, 0, time.UTC))
	require.NoError(t, err)
	_, err = wire.Extend(context.Background(), "lse_1", 1000, time.Date(2026, 1, 1, 0, 15, 0, 0, time.UTC))
	require.NoError(t, err)

	keys := recorder.idempotencyKeys(t)
	require.Len(t, keys, 2)
	assert.NotEmpty(t, keys[0])
	// Two deliberate grows have to be told apart, or the second would be
	// answered with the first one's lease.
	assert.NotEqual(t, keys[0], keys[1])
}

func TestAPIWireClientExtendRetriesUnderOneKey(t *testing.T) {
	t.Parallel()
	data := leaseResponseData()
	data.GrantedAmount = 2000
	wire, recorder := newScriptedWireClient(t,
		scriptedResponse{status: http.StatusBadGateway, body: &schematicgo.ExtendCreditLeaseResponse{}},
		scriptedResponse{status: http.StatusOK, body: &schematicgo.ExtendCreditLeaseResponse{Data: data}},
	)

	grant, err := wire.Extend(context.Background(), "lse_1", 1000, time.Date(2026, 1, 1, 0, 10, 0, 0, time.UTC))

	require.NoError(t, err)
	assert.Equal(t, 2000.0, grant.GrantedAmount)
	keys := recorder.idempotencyKeys(t)
	require.Len(t, keys, 2)
	assert.NotEmpty(t, keys[0])
	// One key across the attempts is what keeps the retry from growing the
	// lease a second time.
	assert.Equal(t, keys[0], keys[1])
}

func TestAPIWireClientAcquireRetriesAfterALostResponse(t *testing.T) {
	t.Parallel()
	wire, recorder := newScriptedWireClient(t,
		scriptedResponse{status: http.StatusBadGateway, body: &schematicgo.AcquireCreditLeaseResponse{}},
		scriptedResponse{status: http.StatusOK, body: &schematicgo.AcquireCreditLeaseResponse{Data: leaseResponseData()}},
	)

	grant, err := wire.Acquire(context.Background(), "co_1", "ct_1", 1000, time.Date(2026, 1, 1, 0, 5, 0, 0, time.UTC))

	require.NoError(t, err)
	require.NotNil(t, grant)
	assert.Equal(t, "lse_1", grant.LeaseID)
	assert.Equal(t, 1000.0, grant.GrantedAmount)
	assert.Len(t, recorder.all(), 2)
}

// The manager sits on top of the retrying wire client, so a retried extend has
// to leave the slot holding the server's total rather than twice the tranche.
func TestManagerAppliesARetriedExtendOnce(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	store := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	installLease(t, store, clock, "lse_1", "co_1", "ct_1", 1000, 300_000)

	data := leaseResponseData()
	data.GrantedAmount = 2000
	wire, recorder := newScriptedWireClient(t,
		scriptedResponse{status: http.StatusBadGateway, body: &schematicgo.ExtendCreditLeaseResponse{}},
		scriptedResponse{status: http.StatusOK, body: &schematicgo.ExtendCreditLeaseResponse{Data: data}},
	)
	manager := NewLeaseManager(wire, store, LeaseManagerOptions{Clock: clock.Now})
	t.Cleanup(manager.Stop)

	required := 1500.0
	entry := manager.MaybeExtend(ctx, "co_1", "ct_1", &required)

	require.NotNil(t, entry)
	assert.Equal(t, 2000.0, entry.GrantedAmount)
	assert.Equal(t, 2000.0, entry.LocalRemainingCredits)

	stored, err := store.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 2000.0, stored.GrantedAmount)

	keys := recorder.idempotencyKeys(t)
	require.Len(t, keys, 2)
	assert.Equal(t, keys[0], keys[1])
}

func TestAPIWireClientRejectsAResponseWithoutALease(t *testing.T) {
	wire, _ := newTestWireClient(t, http.StatusOK, &schematicgo.AcquireCreditLeaseResponse{})

	grant, err := wire.Acquire(context.Background(), "co_1", "ct_1", 1000, time.Now())

	assert.Nil(t, grant)
	require.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "no data"))
}
