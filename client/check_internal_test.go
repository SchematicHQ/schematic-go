package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"

	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/core"
	"github.com/schematichq/schematic-go/datastream"
	"github.com/schematichq/schematic-go/leases"
	option "github.com/schematichq/schematic-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// pathRecorder answers every call with one canned body and remembers which
// paths it was asked for.
type pathRecorder struct {
	body any

	mu    sync.Mutex
	paths []string
}

func (r *pathRecorder) Do(req *http.Request) (*http.Response, error) {
	r.mu.Lock()
	r.paths = append(r.paths, req.URL.Path)
	r.mu.Unlock()

	data, err := json.Marshal(r.body)
	if err != nil {
		return nil, err
	}
	return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(data))}, nil
}

func (r *pathRecorder) recorded() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]string(nil), r.paths...)
}

// Auto mode resolves to client on DataStream alone, so a client whose local
// lease plumbing is missing would otherwise drop to a plain, ungated check. The
// caller asked for credit gating, and server mode needs none of that plumbing,
// so the check gates over the API instead.
//
// The plumbing fields are unexported, so the state is built here rather than
// reached through a constructor that would never produce it.
func TestCheckResolvesMissingClientPlumbingToServerMode(t *testing.T) {
	recorder := &pathRecorder{body: &schematicgo.CheckAndReserveFlagResponse{
		Data: &schematicgo.CheckAndReserveFlagResponseData{
			Flag:   "test-flag",
			Value:  true,
			Reason: "rule matched",
			Reservation: &schematicgo.FlagCheckReservationResponseData{
				ID:           "res_1",
				CompanyID:    "comp_1",
				CreditTypeID: "credit_1",
				EventSubtype: schematicgo.String("inference_tokens"),
				ExpiresAt:    time.Now().UTC().Add(time.Minute),
			},
		},
	}}

	// Auto mode without DataStream builds no lease plumbing.
	client := NewSchematicClient(
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(recorder),
		option.WithCreditLeases(core.CreditLeaseConfig{}),
	)
	t.Cleanup(func() {
		client.datastreamClient = nil
		client.Close()
	})
	// DataStream, as far as the mode resolution can see, with nothing behind it.
	client.datastreamClient = &datastream.DataStreamClient{}
	require.Equal(t, core.CreditLeaseModeClient, client.effectiveLeaseMode())
	require.Nil(t, client.leaseManager)

	result := client.Check(context.Background(), &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"id": "comp-key"},
	}, "test-flag", WithUsage(10))

	assert.True(t, result.Allowed)
	require.NotNil(t, result.Reservation, "the check still gates, and still holds")
	assert.Equal(t, core.CreditLeaseModeServer, result.Reservation.Mode)
	assert.Equal(t, []string{"/flags/test-flag/check-and-reserve"}, recorder.recorded())
}

// A settle that moved no local state must say so, since that is what decides
// whether the cached company metric moves with the event. The event itself still
// goes out: the server is the source of truth for real consumption, and its
// idempotency key is what keeps the retry from billing twice.
func TestSettleClientReservationReportsARepeatSettleAsNotLocal(t *testing.T) {
	leaseStore := leases.NewInMemoryLeaseStore(leases.InMemoryLeaseStoreOptions{})
	_, err := leaseStore.Replace(context.Background(), leases.LeaseGrant{
		LeaseID:       "lse_1",
		CompanyID:     "comp_1",
		CreditTypeID:  "ct_1",
		GrantedAmount: 1000,
		ExpiresAt:     time.Now().Add(5 * time.Minute),
	})
	require.NoError(t, err)
	store := leases.NewInMemoryReservationStore(leaseStore, leases.InMemoryReservationStoreOptions{})

	client := NewSchematicClient(option.WithAPIKey("test-api-key"), option.WithOfflineMode())
	t.Cleanup(client.Close)
	client.reservations = store

	reservation := &Reservation{
		ID:               "res_1",
		LeaseID:          "lse_1",
		Mode:             core.CreditLeaseModeClient,
		CompanyID:        "comp_1",
		CreditTypeID:     "ct_1",
		EventSubtype:     "inference_tokens",
		QuantityReserved: 50,
		CreditsReserved:  50,
		ConsumptionRate:  1,
		ExpiresAt:        time.Now().Add(time.Minute),
		Company:          map[string]string{"id": "comp_1"},
	}
	require.NoError(t, store.Add(context.Background(), reservationRecord(reservation)))

	body, settledLocally := client.settleClientReservation(context.Background(), reservation, 20)
	require.NotNil(t, body)
	assert.True(t, settledLocally)

	body, settledLocally = client.settleClientReservation(context.Background(), reservation, 20)
	require.NotNil(t, body, "the usage is still billed")
	assert.False(t, settledLocally, "the hold was already consumed, so nothing moved locally")
}
