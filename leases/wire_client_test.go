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
// Both lease calls move credits and carry no idempotency key, so a retry would
// take a second hold or grant a tranche twice; these pin that neither retries.

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

// newTestWireClient answers every request with status and body, recording what
// was asked.
func newTestWireClient(t *testing.T, status int, body any) (WireClient, *wireRecorder) {
	t.Helper()
	recorder := &wireRecorder{}
	ctrl := gomock.NewController(t)
	httpClient := mocks.NewMockHTTPClient(ctrl)
	httpClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(req *http.Request) (*http.Response, error) {
		var raw []byte
		if req.Body != nil {
			raw, _ = io.ReadAll(req.Body)
		}
		recorder.mu.Lock()
		recorder.requests = append(recorder.requests, wireRequest{method: req.Method, path: req.URL.Path, body: raw})
		recorder.mu.Unlock()

		encoded, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		return &http.Response{
			Status:     fmt.Sprint(status),
			StatusCode: status,
			Body:       io.NopCloser(bytes.NewReader(encoded)),
		}, nil
	}).AnyTimes()

	options := core.NewRequestOptions(option.WithAPIKey("test-api-key"), option.WithHTTPClient(httpClient))
	return NewAPIWireClient(credits.NewClient(options)), recorder
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

func TestAPIWireClientNeverRetriesACallThatMovesCredits(t *testing.T) {
	// A retried acquire is a second hold against the company balance, and a
	// retried extend grants the tranche twice. Neither request carries an
	// idempotency key, so the default retry policy must not apply.
	for _, testCase := range []struct {
		name string
		call func(WireClient) error
		body any
	}{
		{
			name: "acquire",
			body: &schematicgo.AcquireCreditLeaseResponse{},
			call: func(wire WireClient) error {
				_, err := wire.Acquire(context.Background(), "co_1", "ct_1", 1000, time.Now())
				return err
			},
		},
		{
			name: "extend",
			body: &schematicgo.ExtendCreditLeaseResponse{},
			call: func(wire WireClient) error {
				_, err := wire.Extend(context.Background(), "lse_1", 1000, time.Now())
				return err
			},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			wire, recorder := newTestWireClient(t, http.StatusInternalServerError, testCase.body)

			require.Error(t, testCase.call(wire))
			assert.Len(t, recorder.all(), 1)
		})
	}
}

func TestAPIWireClientRejectsAResponseWithoutALease(t *testing.T) {
	wire, _ := newTestWireClient(t, http.StatusOK, &schematicgo.AcquireCreditLeaseResponse{})

	grant, err := wire.Acquire(context.Background(), "co_1", "ct_1", 1000, time.Now())

	assert.Nil(t, grant)
	require.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "no data"))
}
