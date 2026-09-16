package client_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	http "net/http"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	schematicgo "github.com/schematichq/schematic-go"
	schematicclient "github.com/schematichq/schematic-go/client"
	"github.com/schematichq/schematic-go/core"
	"github.com/schematichq/schematic-go/mocks"
	option "github.com/schematichq/schematic-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	checkAndReservePath = "/check-and-reserve"
	checkPath           = "/check"
	releasePath         = "/release"
)

// recordedRequest is one outgoing HTTP request, kept so a test can assert on
// what the SDK actually put on the wire.
type recordedRequest struct {
	path        string
	body        []byte
	hasDeadline bool
}

type requestRecorder struct {
	mu       sync.Mutex
	requests []recordedRequest
}

func (r *requestRecorder) record(req recordedRequest) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.requests = append(r.requests, req)
}

func (r *requestRecorder) all() []recordedRequest {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]recordedRequest(nil), r.requests...)
}

// paths returns the path of each recorded request, in order.
func (r *requestRecorder) paths() []string {
	paths := make([]string, 0)
	for _, req := range r.all() {
		paths = append(paths, req.path)
	}
	return paths
}

// find returns the first recorded request whose path ends in suffix.
func (r *requestRecorder) find(suffix string) (recordedRequest, bool) {
	for _, req := range r.all() {
		if strings.HasSuffix(req.path, suffix) {
			return req, true
		}
	}
	return recordedRequest{}, false
}

// stub is the canned answer for one endpoint.
type stub struct {
	status int
	body   any
	// block, when set, holds the request open until the channel is closed, so a
	// test can drive an API that is slow to answer.
	block chan struct{}
}

// serveStubs records every request and answers it with the stub whose path
// suffix matches, so a test can drive a check and the release that follows it
// from one mock.
func serveStubs(t *testing.T, rec *requestRecorder, stubs map[string]stub) func(*http.Request) (*http.Response, error) {
	t.Helper()
	return func(req *http.Request) (*http.Response, error) {
		var body []byte
		if req.Body != nil {
			body, _ = io.ReadAll(req.Body)
		}
		_, hasDeadline := req.Context().Deadline()
		rec.record(recordedRequest{path: req.URL.Path, body: body, hasDeadline: hasDeadline})

		for suffix, s := range stubs {
			if !strings.HasSuffix(req.URL.Path, suffix) {
				continue
			}
			if s.block != nil {
				<-s.block
			}
			data, err := json.Marshal(s.body)
			if err != nil {
				return nil, fmt.Errorf("could not marshal the stub for %s: %w", suffix, err)
			}
			status := s.status
			if status == 0 {
				status = http.StatusOK
			}
			return &http.Response{
				Status:     fmt.Sprint(status),
				StatusCode: status,
				Body:       io.NopCloser(bytes.NewReader(data)),
			}, nil
		}
		return nil, fmt.Errorf("unexpected request to %s", req.URL.Path)
	}
}

// capturingLogger keeps every message so a test can assert on a warning.
type capturingLogger struct {
	mu       sync.Mutex
	debugs   []string
	infos    []string
	warns    []string
	errs     []string
	logLevel core.LogLevel
}

func (l *capturingLogger) Debug(_ context.Context, message string, _ ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.debugs = append(l.debugs, message)
}

func (l *capturingLogger) Info(_ context.Context, message string, _ ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.infos = append(l.infos, message)
}

func (l *capturingLogger) Warn(_ context.Context, message string, _ ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.warns = append(l.warns, message)
}

func (l *capturingLogger) Error(_ context.Context, message string, _ ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.errs = append(l.errs, message)
}

func (l *capturingLogger) SetLevel(level core.LogLevel) { l.logLevel = level }
func (l *capturingLogger) GetLevel() core.LogLevel      { return l.logLevel }
func (l *capturingLogger) warnings() []string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return append([]string(nil), l.warns...)
}
func (l *capturingLogger) infoMessages() []string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return append([]string(nil), l.infos...)
}
func containsMatch(messages []string, substr string) bool {
	for _, m := range messages {
		if strings.Contains(m, substr) {
			return true
		}
	}
	return false
}

// serverModeClient is a client configured for server-mode credit holds, with
// every HTTP call answered from stubs.
func serverModeClient(t *testing.T, rec *requestRecorder, stubs map[string]stub, extra ...option.RequestOption) *schematicclient.SchematicClient {
	t.Helper()
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(serveStubs(t, rec, stubs)).AnyTimes()

	opts := append([]option.RequestOption{
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(mockHTTPClient),
		option.WithCreditLeases(core.CreditLeaseConfig{Mode: core.CreditLeaseModeServer}),
	}, extra...)

	client := schematicclient.NewSchematicClient(opts...)
	t.Cleanup(client.Close)
	return client
}

func reserveResponse(value bool, reservation *schematicgo.FlagCheckReservationResponseData) *schematicgo.CheckAndReserveFlagResponse {
	return &schematicgo.CheckAndReserveFlagResponse{
		Data: &schematicgo.CheckAndReserveFlagResponseData{
			Flag:        "test-flag",
			Value:       value,
			Reason:      "rule matched",
			FlagID:      schematicgo.String("flag_123"),
			Reservation: reservation,
		},
	}
}

func heldReservation(eventSubtype *string) *schematicgo.FlagCheckReservationResponseData {
	return &schematicgo.FlagCheckReservationResponseData{
		ID:               "res_123",
		CompanyID:        "comp_123",
		CreditTypeID:     "credit_123",
		EventSubtype:     eventSubtype,
		QuantityReserved: 1000,
		CreditsReserved:  50,
		ConsumptionRate:  0.05,
		ExpiresAt:        time.Now().UTC().Add(time.Minute),
	}
}

func checkFlagResponse(value bool) *schematicgo.CheckFlagResponse {
	return &schematicgo.CheckFlagResponse{
		Data: &schematicgo.CheckFlagResponseData{
			Flag:   "test-flag",
			Value:  value,
			Reason: "rule matched",
		},
	}
}

func testEvalCtx() *schematicgo.CheckFlagRequestBody {
	return &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"id": "comp-key"},
		User:    map[string]string{"id": "user-key"},
	}
}

// decodeBody unmarshals a recorded request body into a generic map.
func decodeBody(t *testing.T, body []byte) map[string]any {
	t.Helper()
	var decoded map[string]any
	require.NoError(t, json.Unmarshal(body, &decoded))
	return decoded
}

func TestCheckWithoutCreditLeases(t *testing.T) {
	rec := &requestRecorder{}
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(serveStubs(t, rec, map[string]stub{
		checkPath: {body: checkFlagResponse(true)},
	})).AnyTimes()

	client := schematicclient.NewSchematicClient(
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(mockHTTPClient),
	)
	defer client.Close()

	result := client.Check(context.Background(), testEvalCtx(), "test-flag", schematicclient.WithUsage(10))

	assert.True(t, result.Allowed)
	assert.True(t, result.Value)
	assert.Nil(t, result.Reservation)
	assert.Equal(t, []string{"/flags/test-flag/check"}, rec.paths())

	// The usage still goes out as a preflight, so the plain verdict accounts
	// for the usage about to be recorded.
	req, ok := rec.find(checkPath)
	require.True(t, ok)
	body := decodeBody(t, req.body)
	preflight, ok := body["preflight"].(map[string]any)
	require.True(t, ok, "preflight should be present: %s", req.body)
	assert.Equal(t, float64(10), preflight["usage"])
}

func TestCheckPreflightBypassesFlagCheckCache(t *testing.T) {
	rec := &requestRecorder{}
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(serveStubs(t, rec, map[string]stub{
		checkPath: {body: checkFlagResponse(true)},
	})).AnyTimes()

	client := schematicclient.NewSchematicClient(
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(mockHTTPClient),
	)
	defer client.Close()

	ctx := context.Background()

	// A plain check populates the cache, and a second one is served from it.
	assert.True(t, client.CheckFlag(ctx, testEvalCtx(), "test-flag"))
	time.Sleep(10 * time.Millisecond)
	assert.True(t, client.CheckFlag(ctx, testEvalCtx(), "test-flag"))
	assert.Len(t, rec.all(), 1)

	// A preflighted check is a different question, so it neither reads that
	// entry nor writes over it.
	preflighted := testEvalCtx()
	preflighted.Preflight = &schematicgo.PreflightRequestBody{Usage: schematicgo.Int64(5)}
	assert.True(t, client.CheckFlag(ctx, preflighted, "test-flag"))
	time.Sleep(10 * time.Millisecond)
	assert.True(t, client.CheckFlag(ctx, preflighted, "test-flag"))
	assert.Len(t, rec.all(), 3)

	// The plain answer is still cached.
	assert.True(t, client.CheckFlag(ctx, testEvalCtx(), "test-flag"))
	assert.Len(t, rec.all(), 3)
}

func TestCheckFlagsPreflightBypassesCache(t *testing.T) {
	rec := &requestRecorder{}
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(serveStubs(t, rec, map[string]stub{
		checkPath: {body: &schematicgo.CheckFlagsResponse{
			Data: &schematicgo.CheckFlagsResponseData{
				Flags: []*schematicgo.CheckFlagResponseData{
					{Flag: "test-flag", Value: true, Reason: "rule matched"},
				},
			},
		}},
	})).AnyTimes()

	client := schematicclient.NewSchematicClient(
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(mockHTTPClient),
	)
	defer client.Close()

	ctx := context.Background()
	keys := []string{"test-flag"}

	client.CheckFlags(ctx, testEvalCtx(), keys)
	time.Sleep(10 * time.Millisecond)
	client.CheckFlags(ctx, testEvalCtx(), keys)
	assert.Len(t, rec.all(), 1)

	preflighted := testEvalCtx()
	preflighted.Preflight = &schematicgo.PreflightRequestBody{Usage: schematicgo.Int64(5)}
	client.CheckFlags(ctx, preflighted, keys)
	time.Sleep(10 * time.Millisecond)
	client.CheckFlags(ctx, preflighted, keys)
	assert.Len(t, rec.all(), 3)
}

func TestCheckOfflineFallsBack(t *testing.T) {
	client := schematicclient.NewSchematicClient(
		option.WithOfflineMode(),
		option.WithDefaultFlagValues(map[string]bool{"test-flag": true}),
		option.WithCreditLeases(core.CreditLeaseConfig{Mode: core.CreditLeaseModeServer}),
	)
	defer client.Close()

	result := client.Check(context.Background(), testEvalCtx(), "test-flag", schematicclient.WithUsage(10))

	assert.True(t, result.Allowed)
	assert.Equal(t, "offline mode", result.Reason)
	assert.Nil(t, result.Reservation)
}

func TestCheckZeroUsageUsesPlainCheck(t *testing.T) {
	rec := &requestRecorder{}
	client := serverModeClient(t, rec, map[string]stub{checkPath: {body: checkFlagResponse(true)}})

	result := client.Check(context.Background(), testEvalCtx(), "test-flag", schematicclient.WithUsage(0))

	assert.True(t, result.Allowed)
	assert.Nil(t, result.Reservation)
	assert.Equal(t, []string{"/flags/test-flag/check"}, rec.paths())
}

func TestCheckNegativeUsage(t *testing.T) {
	t.Run("Fails closed", func(t *testing.T) {
		rec := &requestRecorder{}
		client := serverModeClient(t, rec, map[string]stub{})

		result := client.Check(context.Background(), testEvalCtx(), "test-flag", schematicclient.WithUsage(-1))

		assert.False(t, result.Allowed)
		assert.False(t, result.Value)
		assert.Equal(t, "invalid_usage", result.Reason)
		assert.Equal(t, "invalid_usage", result.Error)
		assert.Empty(t, rec.all(), "a malformed usage must never reach the wire")
	})

	t.Run("Fails open to the caller's default", func(t *testing.T) {
		rec := &requestRecorder{}
		client := serverModeClient(t, rec, map[string]stub{})

		result := client.Check(
			context.Background(), testEvalCtx(), "test-flag",
			schematicclient.WithUsage(-1),
			schematicclient.WithFailOpen(),
			schematicclient.WithCheckDefault(true),
		)

		assert.True(t, result.Allowed)
		assert.True(t, result.Value)
		assert.Equal(t, "invalid_usage_fail_open", result.Reason)
		assert.Equal(t, "invalid_usage", result.Error)
		assert.Empty(t, rec.all())
	})
}

func TestCheckPaymentRequiredDeniesRegardlessOfFailOpen(t *testing.T) {
	rec := &requestRecorder{}
	client := serverModeClient(t, rec, map[string]stub{
		checkAndReservePath: {status: http.StatusPaymentRequired, body: map[string]any{"error": "no credits left"}},
	}, option.WithDefaultFlagValues(map[string]bool{"test-flag": true}))

	result := client.Check(
		context.Background(), testEvalCtx(), "test-flag",
		schematicclient.WithUsage(10),
		schematicclient.WithFailOpen(),
		schematicclient.WithCheckDefault(true),
	)

	assert.False(t, result.Allowed)
	assert.False(t, result.Value)
	assert.Equal(t, "Insufficient credits", result.Reason)
	assert.Equal(t, "no credits left", result.Error)
	assert.Nil(t, result.Reservation)
}

func TestCheckServerError(t *testing.T) {
	t.Run("Fails closed", func(t *testing.T) {
		rec := &requestRecorder{}
		client := serverModeClient(t, rec, map[string]stub{
			checkAndReservePath: {status: http.StatusInternalServerError, body: map[string]any{"error": "boom"}},
		})

		result := client.Check(context.Background(), testEvalCtx(), "test-flag", schematicclient.WithUsage(10))

		assert.False(t, result.Allowed)
		assert.Equal(t, "server_reservation_failed", result.Reason)
		assert.Equal(t, "server_reservation_failed", result.Error)
		// A check-and-reserve carries no idempotency key, so a retry would take
		// a second hold. The call must go out exactly once.
		assert.Len(t, rec.all(), 1)
	})

	t.Run("Fails open to the explicit default", func(t *testing.T) {
		rec := &requestRecorder{}
		client := serverModeClient(t, rec, map[string]stub{
			checkAndReservePath: {status: http.StatusInternalServerError, body: map[string]any{"error": "boom"}},
		})

		result := client.Check(
			context.Background(), testEvalCtx(), "test-flag",
			schematicclient.WithUsage(10),
			schematicclient.WithFailOpen(),
			schematicclient.WithCheckDefault(true),
		)

		assert.True(t, result.Allowed)
		assert.True(t, result.Value)
		assert.Equal(t, "server_reservation_failed_fail_open", result.Reason)
		assert.Equal(t, "server_reservation_failed", result.Error)
	})

	t.Run("Fails open to the client's flag default", func(t *testing.T) {
		rec := &requestRecorder{}
		client := serverModeClient(t, rec, map[string]stub{
			checkAndReservePath: {status: http.StatusInternalServerError, body: map[string]any{"error": "boom"}},
		}, option.WithDefaultFlagValues(map[string]bool{"test-flag": true}))

		result := client.Check(
			context.Background(), testEvalCtx(), "test-flag",
			schematicclient.WithUsage(10),
			schematicclient.WithFailOpen(),
		)

		assert.True(t, result.Allowed)
		assert.Equal(t, "server_reservation_failed_fail_open", result.Reason)
	})
}

func TestCheckDeniedFlagHoldsNothing(t *testing.T) {
	rec := &requestRecorder{}
	client := serverModeClient(t, rec, map[string]stub{
		checkAndReservePath: {body: reserveResponse(false, nil)},
	})

	result := client.Check(context.Background(), testEvalCtx(), "test-flag", schematicclient.WithUsage(10))

	assert.False(t, result.Allowed)
	assert.False(t, result.Value)
	assert.Equal(t, "rule matched", result.Reason)
	assert.Nil(t, result.Reservation)
}

func TestCheckWithoutReservationInResponse(t *testing.T) {
	rec := &requestRecorder{}
	client := serverModeClient(t, rec, map[string]stub{
		checkAndReservePath: {body: reserveResponse(true, nil)},
	})

	result := client.Check(context.Background(), testEvalCtx(), "test-flag", schematicclient.WithUsage(10))

	assert.True(t, result.Allowed)
	assert.True(t, result.Value)
	assert.Nil(t, result.Reservation, "an unmetered feature allows without holding anything")
}

func TestCheckMissingEventSubtypeReleasesTheHold(t *testing.T) {
	t.Run("Releases and fails closed", func(t *testing.T) {
		rec := &requestRecorder{}
		client := serverModeClient(t, rec, map[string]stub{
			checkAndReservePath: {body: reserveResponse(true, heldReservation(nil))},
			releasePath:         {body: map[string]any{"data": map[string]any{}}},
		})

		result := client.Check(context.Background(), testEvalCtx(), "test-flag", schematicclient.WithUsage(10))

		assert.False(t, result.Allowed)
		assert.Equal(t, "missing_event_subtype", result.Reason)
		assert.Equal(t, "missing_event_subtype", result.Error)
		assert.Nil(t, result.Reservation)
		_, released := rec.find("/billing/credits/reservations/res_123/release")
		assert.True(t, released, "an unsettleable hold must be released, not parked until its TTL: %v", rec.paths())
	})

	t.Run("Swallows a failed release", func(t *testing.T) {
		rec := &requestRecorder{}
		client := serverModeClient(t, rec, map[string]stub{
			checkAndReservePath: {body: reserveResponse(true, heldReservation(nil))},
			releasePath:         {status: http.StatusInternalServerError, body: map[string]any{"error": "boom"}},
		})

		result := client.Check(context.Background(), testEvalCtx(), "test-flag", schematicclient.WithUsage(10))

		assert.False(t, result.Allowed)
		assert.Equal(t, "missing_event_subtype", result.Reason)
		releases := 0
		for _, path := range rec.paths() {
			if strings.HasSuffix(path, releasePath) {
				releases++
			}
		}
		assert.Equal(t, 1, releases, "the release is best effort, and the hold expires anyway, so it is not retried")
	})

	t.Run("Releases under a context of its own", func(t *testing.T) {
		rec := &requestRecorder{}
		client := serverModeClient(t, rec, map[string]stub{
			checkAndReservePath: {body: reserveResponse(true, heldReservation(nil))},
			releasePath:         {body: map[string]any{"data": map[string]any{}}},
		})

		// This check sets no timeout, so the caller's context carries no
		// deadline: a release that has one is a release detached from the
		// caller and bounded on its own.
		result := client.Check(context.Background(), testEvalCtx(), "test-flag", schematicclient.WithUsage(10))

		assert.False(t, result.Allowed)
		released, ok := rec.find("/billing/credits/reservations/res_123/release")
		require.True(t, ok)
		assert.True(t, released.hasDeadline, "a best-effort release does not ride the caller's context")
	})

	t.Run("Keeps the server's verdict when failing open", func(t *testing.T) {
		rec := &requestRecorder{}
		client := serverModeClient(t, rec, map[string]stub{
			checkAndReservePath: {body: reserveResponse(true, heldReservation(nil))},
			releasePath:         {body: map[string]any{"data": map[string]any{}}},
		})

		result := client.Check(
			context.Background(), testEvalCtx(), "test-flag",
			schematicclient.WithUsage(10),
			schematicclient.WithFailOpen(),
		)

		// The server already evaluated the flag and allowed it; only the settle
		// is impossible.
		assert.True(t, result.Allowed)
		assert.Equal(t, "rule matched", result.Reason)
		assert.Equal(t, "missing_event_subtype", result.Error)
		assert.Nil(t, result.Reservation)
	})

	t.Run("The caller's subtype settles a hold the server did not name", func(t *testing.T) {
		rec := &requestRecorder{}
		client := serverModeClient(t, rec, map[string]stub{
			checkAndReservePath: {body: reserveResponse(true, heldReservation(nil))},
		})

		result := client.Check(
			context.Background(), testEvalCtx(), "test-flag",
			schematicclient.WithUsage(10),
			schematicclient.WithEventSubtype("inference_tokens"),
		)

		require.NotNil(t, result.Reservation)
		assert.Equal(t, "inference_tokens", result.Reservation.EventSubtype)
		_, released := rec.find(releasePath)
		assert.False(t, released)
	})
}

func TestCheckHappyPath(t *testing.T) {
	rec := &requestRecorder{}
	client := serverModeClient(t, rec, map[string]stub{
		checkAndReservePath: {body: reserveResponse(true, heldReservation(schematicgo.String("inference_tokens")))},
	})

	before := time.Now().UTC()
	result := client.Check(
		context.Background(), testEvalCtx(), "test-flag",
		schematicclient.WithUsage(1000),
		schematicclient.WithEventSubtype("inference_tokens"),
		schematicclient.WithCheckTimeout(5*time.Second),
	)

	assert.True(t, result.Allowed)
	assert.True(t, result.Value)
	assert.Equal(t, "test-flag", result.FlagKey)
	require.NotNil(t, result.FlagID)
	assert.Equal(t, "flag_123", *result.FlagID)

	require.NotNil(t, result.Reservation)
	reservation := result.Reservation
	assert.Equal(t, "res_123", reservation.ID)
	assert.Equal(t, "res_123", reservation.LeaseID, "server mode has no lease, so the ID mirrors")
	assert.Equal(t, core.CreditLeaseModeServer, reservation.Mode)
	assert.Equal(t, "comp_123", reservation.CompanyID)
	assert.Equal(t, "credit_123", reservation.CreditTypeID)
	assert.Equal(t, "inference_tokens", reservation.EventSubtype)
	assert.Equal(t, float64(1000), reservation.QuantityReserved)
	assert.Equal(t, float64(50), reservation.CreditsReserved)
	assert.Equal(t, 0.05, reservation.ConsumptionRate)
	assert.Equal(t, map[string]string{"id": "comp-key"}, reservation.Company)
	assert.Equal(t, map[string]string{"id": "user-key"}, reservation.User)

	req, ok := rec.find(checkAndReservePath)
	require.True(t, ok)
	assert.Equal(t, "/flags/test-flag/check-and-reserve", req.path)
	assert.True(t, req.hasDeadline, "WithCheckTimeout should bound the call")

	body := decodeBody(t, req.body)
	assert.Equal(t, float64(1000), body["quantity"])
	assert.Equal(t, map[string]any{"id": "comp-key"}, body["company"])
	assert.Equal(t, map[string]any{"id": "user-key"}, body["user"])

	expiresAt, err := time.Parse(time.RFC3339Nano, body["expires_at"].(string))
	require.NoError(t, err)
	assert.WithinDuration(t, before.Add(time.Minute), expiresAt, 5*time.Second)

	preflight, ok := body["preflight"].(map[string]any)
	require.True(t, ok, "preflight should be present: %s", req.body)
	eventUsage, ok := preflight["event_usage"].(map[string]any)
	require.True(t, ok, "a named subtype goes out as event_usage: %s", req.body)
	assert.Equal(t, "inference_tokens", eventUsage["event_subtype"])
	assert.Equal(t, float64(1000), eventUsage["quantity"])
	assert.Nil(t, preflight["usage"])
}

func TestCheckPreflightWithoutEventSubtype(t *testing.T) {
	rec := &requestRecorder{}
	client := serverModeClient(t, rec, map[string]stub{
		checkAndReservePath: {body: reserveResponse(true, heldReservation(schematicgo.String("inference_tokens")))},
	})

	client.Check(context.Background(), testEvalCtx(), "test-flag", schematicclient.WithUsage(1000))

	req, ok := rec.find(checkAndReservePath)
	require.True(t, ok)
	preflight := decodeBody(t, req.body)["preflight"].(map[string]any)
	assert.Equal(t, float64(1000), preflight["usage"])
	assert.Nil(t, preflight["event_usage"])
}

func TestCheckCustomReservationTTL(t *testing.T) {
	rec := &requestRecorder{}
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(serveStubs(t, rec, map[string]stub{
		checkAndReservePath: {body: reserveResponse(true, heldReservation(schematicgo.String("inference_tokens")))},
	})).AnyTimes()

	client := schematicclient.NewSchematicClient(
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(mockHTTPClient),
		option.WithCreditLeases(core.CreditLeaseConfig{
			Mode:                  core.CreditLeaseModeServer,
			DefaultReservationTTL: 10 * time.Minute,
		}),
	)
	defer client.Close()

	before := time.Now().UTC()
	client.Check(context.Background(), testEvalCtx(), "test-flag", schematicclient.WithUsage(10))

	req, ok := rec.find(checkAndReservePath)
	require.True(t, ok)
	expiresAt, err := time.Parse(time.RFC3339Nano, decodeBody(t, req.body)["expires_at"].(string))
	require.NoError(t, err)
	assert.WithinDuration(t, before.Add(10*time.Minute), expiresAt, 5*time.Second)
}

func TestTrackWithReservation(t *testing.T) {
	serverReservation := func() *schematicclient.Reservation {
		return &schematicclient.Reservation{
			ID:           "res_123",
			LeaseID:      "res_123",
			Mode:         core.CreditLeaseModeServer,
			CompanyID:    "comp_123",
			CreditTypeID: "credit_123",
			EventSubtype: "inference_tokens",
			Company:      map[string]string{"id": "comp-key"},
			User:         map[string]string{"id": "user-key"},
		}
	}

	t.Run("Settles the hold by ID", func(t *testing.T) {
		rec := &requestRecorder{}
		ctrl := gomock.NewController(t)
		mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
		mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(serveStubs(t, rec, map[string]stub{
			"/batch": {body: &schematicgo.CreateEventBatchResponse{
				Data: &schematicgo.RawEventBatchResponseData{Events: []*schematicgo.RawEventResponseData{}},
			}},
		})).AnyTimes()

		client := schematicclient.NewSchematicClient(
			option.WithAPIKey("test-api-key"),
			option.WithHTTPClient(mockHTTPClient),
			option.WithEventBufferPeriod(10*time.Millisecond),
			option.WithCreditLeases(core.CreditLeaseConfig{Mode: core.CreditLeaseModeServer}),
		)
		defer client.Close()

		client.TrackWithReservation(context.Background(), serverReservation(), 750)
		time.Sleep(50 * time.Millisecond)

		req, ok := rec.find("/batch")
		require.True(t, ok, "the settle should be flushed: %v", rec.paths())

		var batch struct {
			Events []struct {
				IdempotencyKey string `json:"idempotency_key"`
				Body           struct {
					Company       map[string]string `json:"company"`
					Event         string            `json:"event"`
					LeaseID       *string           `json:"lease_id"`
					Quantity      *int64            `json:"quantity"`
					ReservationID *string           `json:"reservation_id"`
					User          map[string]string `json:"user"`
				} `json:"body"`
			} `json:"events"`
		}
		require.NoError(t, json.Unmarshal(req.body, &batch))
		require.Len(t, batch.Events, 1)

		event := batch.Events[0]
		assert.Equal(t, "lease-reservation:res_123", event.IdempotencyKey)
		assert.Equal(t, "inference_tokens", event.Body.Event)
		require.NotNil(t, event.Body.ReservationID)
		assert.Equal(t, "res_123", *event.Body.ReservationID)
		assert.Nil(t, event.Body.LeaseID, "a server-mode hold has no lease behind it")
		require.NotNil(t, event.Body.Quantity)
		assert.Equal(t, int64(750), *event.Body.Quantity)
		assert.Equal(t, map[string]string{"id": "comp-key"}, event.Body.Company)
		assert.Equal(t, map[string]string{"id": "user-key"}, event.Body.User)
	})

	t.Run("Skips a nil reservation", func(t *testing.T) {
		rec := &requestRecorder{}
		client := serverModeClient(t, rec, map[string]stub{}, option.WithEventBufferPeriod(10*time.Millisecond))

		client.TrackWithReservation(context.Background(), nil, 750)
		time.Sleep(50 * time.Millisecond)

		assert.Empty(t, rec.all())
	})

	t.Run("Skips a negative quantity", func(t *testing.T) {
		rec := &requestRecorder{}
		client := serverModeClient(t, rec, map[string]stub{}, option.WithEventBufferPeriod(10*time.Millisecond))

		client.TrackWithReservation(context.Background(), serverReservation(), -1)
		time.Sleep(50 * time.Millisecond)

		assert.Empty(t, rec.all(), "the untouched hold is refunded at its TTL")
	})
}

func TestCreditLeaseConstructorWarnings(t *testing.T) {
	t.Run("Clamps and warns about a too-long TTL in server mode", func(t *testing.T) {
		logger := &capturingLogger{}
		client := schematicclient.NewSchematicClient(
			option.WithAPIKey("test-api-key"),
			option.WithLogger(logger),
			option.WithCreditLeases(core.CreditLeaseConfig{
				Mode:                  core.CreditLeaseModeServer,
				DefaultReservationTTL: 3 * time.Hour,
			}),
		)
		defer client.Close()

		assert.True(t, containsMatch(logger.warnings(), "longer than the API will hold credits for"), logger.warnings())
	})

	t.Run("Stays quiet about a too-long TTL in client mode", func(t *testing.T) {
		logger := &capturingLogger{}
		client := schematicclient.NewSchematicClient(
			option.WithAPIKey("test-api-key"),
			option.WithLogger(logger),
			option.WithCreditLeases(core.CreditLeaseConfig{
				Mode:                  core.CreditLeaseModeClient,
				DefaultReservationTTL: 3 * time.Hour,
			}),
		)
		defer client.Close()

		// The TTL only drives the local sweeper in client mode, so clamping it
		// would be untrue.
		assert.False(t, containsMatch(logger.warnings(), "longer than the API will hold credits for"), logger.warnings())
	})

	t.Run("Names the client-only options server mode will ignore", func(t *testing.T) {
		logger := &capturingLogger{}
		client := schematicclient.NewSchematicClient(
			option.WithAPIKey("test-api-key"),
			option.WithLogger(logger),
			option.WithCreditLeases(core.CreditLeaseConfig{
				Mode:                 core.CreditLeaseModeServer,
				DefaultLeaseSize:     10_000,
				DefaultLeaseDuration: 5 * time.Minute,
				RedisKeyPrefix:       "schematic:",
			}),
		)
		defer client.Close()

		warnings := logger.warnings()
		require.True(t, containsMatch(warnings, "resolve to server mode"), warnings)
		assert.True(t, containsMatch(warnings, "DefaultLeaseSize"), warnings)
		assert.True(t, containsMatch(warnings, "DefaultLeaseDuration"), warnings)
		assert.True(t, containsMatch(warnings, "RedisKeyPrefix"), warnings)
		assert.False(t, containsMatch(warnings, "LowWaterMark"), warnings)
	})

	t.Run("Warns that client mode without DataStream gates nothing", func(t *testing.T) {
		logger := &capturingLogger{}
		client := schematicclient.NewSchematicClient(
			option.WithAPIKey("test-api-key"),
			option.WithLogger(logger),
			option.WithCreditLeases(core.CreditLeaseConfig{Mode: core.CreditLeaseModeClient}),
		)
		defer client.Close()

		assert.True(t, containsMatch(logger.warnings(), "requires DataStream"), logger.warnings())
	})

	t.Run("Says auto without DataStream lands on server mode", func(t *testing.T) {
		logger := &capturingLogger{}
		client := schematicclient.NewSchematicClient(
			option.WithAPIKey("test-api-key"),
			option.WithLogger(logger),
			option.WithCreditLeases(core.CreditLeaseConfig{}),
		)
		defer client.Close()

		assert.True(t, containsMatch(logger.infoMessages(), "server mode"), logger.infoMessages())
		assert.Empty(t, logger.warnings())
	})
}

func TestCheckInClientModeWithoutDataStreamIsUngated(t *testing.T) {
	rec := &requestRecorder{}
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(serveStubs(t, rec, map[string]stub{
		checkPath: {body: checkFlagResponse(true)},
	})).AnyTimes()

	client := schematicclient.NewSchematicClient(
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(mockHTTPClient),
		option.WithCreditLeases(core.CreditLeaseConfig{Mode: core.CreditLeaseModeClient}),
	)
	defer client.Close()

	result := client.Check(context.Background(), testEvalCtx(), "test-flag", schematicclient.WithUsage(10))

	assert.True(t, result.Allowed)
	assert.Nil(t, result.Reservation)
	assert.Equal(t, []string{"/flags/test-flag/check"}, rec.paths())
}

func TestCheckKeepsTheCallersPreflight(t *testing.T) {
	callerPreflight := func() *schematicgo.PreflightRequestBody {
		return &schematicgo.PreflightRequestBody{
			CreditCost: map[string]float64{"credit_123": 7},
			Usage:      schematicgo.Int64(999),
		}
	}

	t.Run("Carries the caller's credit cost under the option's usage", func(t *testing.T) {
		rec := &requestRecorder{}
		client := serverModeClient(t, rec, map[string]stub{
			checkAndReservePath: {body: reserveResponse(true, heldReservation(schematicgo.String("inference_tokens")))},
		})

		evalCtx := testEvalCtx()
		evalCtx.Preflight = callerPreflight()
		client.Check(context.Background(), evalCtx, "test-flag",
			schematicclient.WithUsage(10),
			schematicclient.WithEventSubtype("inference_tokens"),
		)

		req, ok := rec.find(checkAndReservePath)
		require.True(t, ok)
		preflight := decodeBody(t, req.body)["preflight"].(map[string]any)
		assert.Equal(t, map[string]any{"credit_123": float64(7)}, preflight["credit_cost"], "a cost only the caller could price survives")
		eventUsage, ok := preflight["event_usage"].(map[string]any)
		require.True(t, ok, "the option owns the usage knobs: %s", req.body)
		assert.Equal(t, float64(10), eventUsage["quantity"])
		assert.Nil(t, preflight["usage"], "the option's shape replaces the caller's")
	})

	t.Run("Uses the caller's preflight as-is without WithUsage", func(t *testing.T) {
		rec := &requestRecorder{}
		client := serverModeClient(t, rec, map[string]stub{checkPath: {body: checkFlagResponse(true)}})

		evalCtx := testEvalCtx()
		evalCtx.Preflight = callerPreflight()
		client.Check(context.Background(), evalCtx, "test-flag")

		req, ok := rec.find(checkPath)
		require.True(t, ok)
		preflight := decodeBody(t, req.body)["preflight"].(map[string]any)
		assert.Equal(t, map[string]any{"credit_123": float64(7)}, preflight["credit_cost"])
		assert.Equal(t, float64(999), preflight["usage"])
	})
}

func TestCheckFractionalUsage(t *testing.T) {
	rec := &requestRecorder{}
	client := serverModeClient(t, rec, map[string]stub{
		checkAndReservePath: {body: reserveResponse(true, heldReservation(schematicgo.String("inference_tokens")))},
	})

	client.Check(context.Background(), testEvalCtx(), "test-flag",
		schematicclient.WithUsage(2.5),
		schematicclient.WithEventSubtype("inference_tokens"),
	)

	req, ok := rec.find(checkAndReservePath)
	require.True(t, ok)
	body := decodeBody(t, req.body)
	assert.Equal(t, 2.5, body["quantity"], "the hold is sized from the fraction the caller declared")

	// The preflight's quantity is an integer and its question is an upper
	// bound, so the fraction rounds up rather than down.
	eventUsage := body["preflight"].(map[string]any)["event_usage"].(map[string]any)
	assert.Equal(t, float64(3), eventUsage["quantity"])
}
