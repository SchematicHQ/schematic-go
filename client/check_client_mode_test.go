package client_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/websocket"
	schematicdatastreamws "github.com/schematichq/schematic-datastream-ws"
	schematicgo "github.com/schematichq/schematic-go"
	schematicclient "github.com/schematichq/schematic-go/client"
	"github.com/schematichq/schematic-go/core"
	"github.com/schematichq/schematic-go/mocks"
	option "github.com/schematichq/schematic-go/option"
	"github.com/schematichq/schematic-go/rulesengine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Client-mode credit leases end to end: a DataStream serving a credit-metered
// flag and its company, the real rules engine, and the lease API stubbed. The
// semantics live in the leases package and its conformance vectors; what is
// pinned here is the wiring, the settle event's shape, and the lifecycle a
// SchematicClient owns.

const (
	leaseAcquirePath = "/billing/credits/lease"
	leaseReleasePath = "/release"
	leaseExtendPath  = "/extend"
	eventBatchPath   = "/batch"
	testFlagKey      = "inference"
	testCreditID     = "ct_1"
	testEventSubtype = "inference_tokens"
	// Prefixed the way a Schematic company id is, so a prewarm whose keys the
	// cache cannot resolve still falls back to it.
	testCompanyID = "comp_1"
)

// creditFlag meters the feature by credit burndown, so a check with usage takes
// the lease path.
func creditFlag() *rulesengine.Flag {
	creditID := testCreditID
	rate := 1.0
	subtype := testEventSubtype
	return &rulesengine.Flag{
		ID:            "flag_1",
		AccountID:     "acct",
		EnvironmentID: "env",
		Key:           testFlagKey,
		DefaultValue:  false,
		Rules: rulesengine.JSONSlice[*rulesengine.Rule]{
			{
				ID:            "rule_1",
				AccountID:     "acct",
				EnvironmentID: "env",
				RuleType:      rulesengine.RuleTypePlanEntitlement,
				Name:          "Credit",
				Priority:      100,
				Value:         true,
				Conditions: rulesengine.JSONSlice[*rulesengine.Condition]{{
					ID:              "cond_1",
					AccountID:       "acct",
					EnvironmentID:   "env",
					ConditionType:   rulesengine.ConditionTypeCredit,
					Operator:        rulesengine.ComparableOperatorLt,
					CreditID:        &creditID,
					ConsumptionRate: &rate,
					EventSubtype:    &subtype,
				}},
			},
		},
	}
}

func creditCompany(keys map[string]string) *rulesengine.Company {
	creditID := testCreditID
	rate := 1.0
	subtype := testEventSubtype
	balance := 100000.0
	return &rulesengine.Company{
		ID:             testCompanyID,
		AccountID:      "acct",
		EnvironmentID:  "env",
		Keys:           keys,
		CreditBalances: map[string]float64{testCreditID: balance},
		Entitlements: rulesengine.JSONSlice[*rulesengine.FeatureEntitlement]{{
			FeatureID:       "feat_1",
			FeatureKey:      testFlagKey,
			ValueType:       rulesengine.EntitlementValueTypeCredit,
			CreditID:        &creditID,
			ConsumptionRate: &rate,
			EventSubtype:    &subtype,
			CreditTotal:     &balance,
			CreditRemaining: &balance,
		}},
		Metrics: make(rulesengine.CompanyMetricCollection, 0),
		Traits:  make(rulesengine.JSONSlice[*rulesengine.Trait], 0),
	}
}

func datastreamMessage(entityType schematicdatastreamws.EntityType, payload any) string {
	data, _ := json.Marshal(payload)
	message, _ := json.Marshal(schematicdatastreamws.DataStreamResp{
		EntityType:  string(entityType),
		Data:        data,
		MessageType: schematicdatastreamws.MessageTypeFull,
	})
	return string(message)
}

// dataStreamFixture says what the fixture DataStream server pushes. A test that
// needs neither flags nor companies keeps the socket silent, which is also the
// only way to test a prewarm that must not wait on one.
type dataStreamFixture struct {
	pushFlags    bool
	serveCompany bool
}

// startDataStreamServer stands in for the DataStream: it pushes the credit flag
// on connect and answers company requests with the fixture company.
func startDataStreamServer(t *testing.T, fixture dataStreamFixture) *httptest.Server {
	t.Helper()
	upgrader := websocket.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer func() { _ = conn.Close() }()

		if fixture.pushFlags {
			if err := conn.WriteMessage(websocket.TextMessage, []byte(datastreamMessage(
				schematicdatastreamws.EntityTypeFlags, []*rulesengine.Flag{creditFlag()},
			))); err != nil {
				return
			}
		}
		for {
			_, raw, err := conn.ReadMessage()
			if err != nil {
				return
			}
			var request schematicdatastreamws.DataStreamBaseReq
			if err := json.Unmarshal(raw, &request); err != nil {
				continue
			}
			if request.Data.EntityType != schematicdatastreamws.EntityTypeCompany || !fixture.serveCompany {
				continue
			}
			_ = conn.WriteMessage(websocket.TextMessage, []byte(datastreamMessage(
				schematicdatastreamws.EntityTypeCompany, creditCompany(request.Data.Keys),
			)))
		}
	}))
	t.Cleanup(server.Close)
	return server
}

func leaseResponse(leaseID string, granted float64) *schematicgo.AcquireCreditLeaseResponse {
	return &schematicgo.AcquireCreditLeaseResponse{
		Data: &schematicgo.CreditLeaseResponseData{
			ID:            leaseID,
			CompanyID:     testCompanyID,
			CreditTypeID:  testCreditID,
			GrantedAmount: granted,
			ExpiresAt:     time.Now().UTC().Add(5 * time.Minute),
		},
	}
}

// clientModeClient is a client whose credit holds are carved out of local
// leases, with DataStream pointed at a fixture server and every API call stubbed.
func clientModeClient(
	t *testing.T,
	rec *requestRecorder,
	stubs map[string]stub,
	serverURL string,
	config core.CreditLeaseConfig,
	extra ...option.RequestOption,
) *schematicclient.SchematicClient {
	t.Helper()
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(serveStubs(t, rec, stubs)).AnyTimes()

	config.Mode = core.CreditLeaseModeClient
	opts := append([]option.RequestOption{
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(mockHTTPClient),
		option.WithBaseURL(serverURL),
		option.WithEventBufferPeriod(10 * time.Millisecond),
		option.WithDatastream(),
		option.WithCreditLeases(config),
	}, extra...)

	client := schematicclient.NewSchematicClient(opts...)
	// Give the socket time to connect and the flags to land.
	time.Sleep(300 * time.Millisecond)
	return client
}

// trackEvents pulls the track events out of whatever batches were posted.
func trackEvents(t *testing.T, rec *requestRecorder) []*schematicgo.EventBodyTrack {
	t.Helper()
	events := make([]*schematicgo.EventBodyTrack, 0)
	for _, request := range rec.all() {
		if !strings.HasSuffix(request.path, eventBatchPath) {
			continue
		}
		var batch struct {
			Events []struct {
				Type string          `json:"type"`
				Body json.RawMessage `json:"body"`
			} `json:"events"`
		}
		require.NoError(t, json.Unmarshal(request.body, &batch))
		for _, event := range batch.Events {
			if event.Type != "track" {
				continue
			}
			body := &schematicgo.EventBodyTrack{}
			require.NoError(t, json.Unmarshal(event.Body, body))
			events = append(events, body)
		}
	}
	return events
}

func TestClientModeCheckIssuesALeaseBackedHoldAndSettlesItByLeaseID(t *testing.T) {
	server := startDataStreamServer(t, dataStreamFixture{pushFlags: true, serveCompany: true})
	rec := &requestRecorder{}
	client := clientModeClient(t, rec, map[string]stub{
		leaseAcquirePath: {body: leaseResponse("lse_1", 10000)},
		eventBatchPath:   {body: map[string]any{"data": map[string]any{"events": []any{}}}},
	}, server.URL, core.CreditLeaseConfig{})
	defer client.Close()

	evalCtx := &schematicgo.CheckFlagRequestBody{Company: map[string]string{"id": testCompanyID}}
	result := client.Check(t.Context(), evalCtx, testFlagKey,
		schematicclient.WithUsage(50),
		schematicclient.WithEventSubtype(testEventSubtype),
	)

	require.True(t, result.Allowed, "reason: %s, error: %s", result.Reason, result.Error)
	require.NotNil(t, result.Reservation)
	assert.Equal(t, core.CreditLeaseModeClient, result.Reservation.Mode)
	assert.Equal(t, "lse_1", result.Reservation.LeaseID)
	assert.Equal(t, testCreditID, result.Reservation.CreditTypeID)
	assert.Equal(t, 50.0, result.Reservation.CreditsReserved)
	_, acquired := rec.find(leaseAcquirePath)
	assert.True(t, acquired, "the check acquires a lease over the wire")

	client.TrackWithReservation(t.Context(), result.Reservation, 20,
		schematicclient.WithTrackTraits(map[string]any{"model": "sonnet"}))
	time.Sleep(100 * time.Millisecond)

	events := trackEvents(t, rec)
	require.NotEmpty(t, events)
	settle := events[len(events)-1]
	assert.Equal(t, testEventSubtype, settle.Event)
	require.NotNil(t, settle.Quantity)
	assert.Equal(t, int64(20), *settle.Quantity)
	require.NotNil(t, settle.LeaseID)
	assert.Equal(t, "lse_1", *settle.LeaseID)
	assert.Nil(t, settle.ReservationID, "a client-mode settle routes by lease, never by hold id")
	assert.Equal(t, map[string]any{"model": "sonnet"}, settle.Traits)
}

func TestClientModeCheckFallsBackToAPlainCheckWithoutUsage(t *testing.T) {
	server := startDataStreamServer(t, dataStreamFixture{pushFlags: true, serveCompany: true})
	rec := &requestRecorder{}
	client := clientModeClient(t, rec, map[string]stub{
		eventBatchPath: {body: map[string]any{"data": map[string]any{"events": []any{}}}},
	}, server.URL, core.CreditLeaseConfig{})
	defer client.Close()

	result := client.Check(t.Context(), &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"id": testCompanyID},
	}, testFlagKey)

	assert.Nil(t, result.Reservation)
	_, acquired := rec.find(leaseAcquirePath)
	assert.False(t, acquired, "a check without usage holds nothing")
}

func TestClientModeCloseReleasesThisProcessLeases(t *testing.T) {
	server := startDataStreamServer(t, dataStreamFixture{pushFlags: true, serveCompany: true})
	rec := &requestRecorder{}
	client := clientModeClient(t, rec, map[string]stub{
		leaseAcquirePath: {body: leaseResponse("lse_1", 10000)},
		leaseReleasePath: {body: leaseResponse("lse_1", 10000)},
		eventBatchPath:   {body: map[string]any{"data": map[string]any{"events": []any{}}}},
	}, server.URL, core.CreditLeaseConfig{})

	result := client.Check(t.Context(), &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"id": testCompanyID},
	}, testFlagKey, schematicclient.WithUsage(50), schematicclient.WithEventSubtype(testEventSubtype))
	require.NotNil(t, result.Reservation)

	client.Close()

	// The remainder goes back to the company balance now rather than at expiry.
	released, ok := rec.find(leaseReleasePath)
	require.True(t, ok)
	assert.Contains(t, released.path, "lse_1")
}

func TestClientModePrewarmAcquiresPerCreditType(t *testing.T) {
	server := startDataStreamServer(t, dataStreamFixture{pushFlags: true, serveCompany: true})
	rec := &requestRecorder{}
	client := clientModeClient(t, rec, map[string]stub{
		leaseAcquirePath: {body: leaseResponse("lse_1", 10000)},
		eventBatchPath:   {body: map[string]any{"data": map[string]any{"events": []any{}}}},
	}, server.URL, core.CreditLeaseConfig{})
	defer client.Close()

	err := client.Prewarm(t.Context(), &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"id": testCompanyID},
	}, []string{testCreditID})

	require.NoError(t, err)
	_, acquired := rec.find(leaseAcquirePath)
	assert.True(t, acquired)
}

func TestClientModePrewarmWithTheWaitOffAnswersFromTheCacheAlone(t *testing.T) {
	// The company is never served over DataStream, so only the cache could
	// answer, and it is empty. A zero resolve timeout says not to wait.
	server := startDataStreamServer(t, dataStreamFixture{})
	rec := &requestRecorder{}
	noWait := time.Duration(0)
	client := clientModeClient(t, rec, map[string]stub{
		leaseAcquirePath: {body: leaseResponse("lse_1", 10000)},
		eventBatchPath:   {body: map[string]any{"data": map[string]any{"events": []any{}}}},
	}, server.URL, core.CreditLeaseConfig{PrewarmResolveTimeout: &noWait})
	defer client.Close()

	start := time.Now()
	err := client.Prewarm(t.Context(), &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"company_key": "acme"},
	}, []string{testCreditID})

	require.Error(t, err)
	assert.Less(t, time.Since(start), time.Second, "the wait is off, so nothing is waited on")
	_, acquired := rec.find(leaseAcquirePath)
	assert.False(t, acquired)
}

func TestClientModeIdentifyPrewarmsInTheBackground(t *testing.T) {
	server := startDataStreamServer(t, dataStreamFixture{pushFlags: true, serveCompany: true})
	rec := &requestRecorder{}
	client := clientModeClient(t, rec, map[string]stub{
		leaseAcquirePath: {body: leaseResponse("lse_1", 10000)},
		eventBatchPath:   {body: map[string]any{"data": map[string]any{"events": []any{}}}},
	}, server.URL, core.CreditLeaseConfig{})
	defer client.Close()

	client.Identify(t.Context(), &schematicgo.EventBodyIdentify{
		Keys:    map[string]string{"id": "user_1"},
		Company: &schematicgo.EventBodyIdentifyCompany{Keys: map[string]string{"id": testCompanyID}},
	}, schematicclient.WithIdentifyPrewarm([]string{testCreditID}))

	require.Eventually(t, func() bool {
		_, acquired := rec.find(leaseAcquirePath)
		return acquired
	}, 2*time.Second, 20*time.Millisecond)
}

func TestTrackTraitsNeverWriteThroughTheCallersBody(t *testing.T) {
	rec := &requestRecorder{}
	client := serverModeClient(t, rec, map[string]stub{
		eventBatchPath: {body: map[string]any{"data": map[string]any{"events": []any{}}}},
	}, option.WithEventBufferPeriod(10*time.Millisecond))
	defer client.Close()

	body := &schematicgo.EventBodyTrack{
		Event:   testEventSubtype,
		Company: map[string]string{"id": testCompanyID},
		Traits:  map[string]any{"model": "haiku", "region": "us"},
	}
	client.Track(t.Context(), body, schematicclient.WithTrackTraits(map[string]any{"model": "sonnet"}))
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, map[string]any{"model": "haiku", "region": "us"}, body.Traits, "the caller's body is untouched")
	events := trackEvents(t, rec)
	require.Len(t, events, 1)
	assert.Equal(t, map[string]any{"model": "sonnet", "region": "us"}, events[0].Traits, "the option wins a key collision")
}

// An identify the server has not seen yet cannot be prewarmed: the buffer would
// hold it past the prewarm's own wait for the company to surface.
func TestClientModeIdentifyFlushesTheBufferBeforePrewarming(t *testing.T) {
	// The company id the keys carry resolves without the socket once the key
	// lookup misses, and a zero resolve wait skips the fetch in between, so the
	// fixture serves nothing and the ordering under test is the only thing
	// moving.
	server := startDataStreamServer(t, dataStreamFixture{})
	rec := &requestRecorder{}
	noWait := time.Duration(0)
	// A buffer period no test would outlast, so only a forced flush can put the
	// identify on the wire before the acquire.
	client := clientModeClient(t, rec, map[string]stub{
		leaseAcquirePath: {body: leaseResponse("lse_1", 10000)},
		eventBatchPath:   {body: map[string]any{"data": map[string]any{"events": []any{}}}},
	}, server.URL, core.CreditLeaseConfig{PrewarmResolveTimeout: &noWait}, option.WithEventBufferPeriod(time.Minute))
	defer client.Close()

	client.Identify(t.Context(), &schematicgo.EventBodyIdentify{
		Keys:    map[string]string{"id": "user_1"},
		Company: &schematicgo.EventBodyIdentifyCompany{Keys: map[string]string{"id": testCompanyID}},
	}, schematicclient.WithIdentifyPrewarm([]string{testCreditID}))

	require.Eventually(t, func() bool {
		_, acquired := rec.find(leaseAcquirePath)
		return acquired
	}, 2*time.Second, 10*time.Millisecond)

	paths := rec.paths()
	batch, acquire := -1, -1
	for i, path := range paths {
		if batch < 0 && strings.HasSuffix(path, eventBatchPath) {
			batch = i
		}
		if acquire < 0 && strings.HasSuffix(path, leaseAcquirePath) {
			acquire = i
		}
	}
	require.GreaterOrEqual(t, batch, 0, "the identify should be flushed, not buffered: %v", paths)
	assert.Less(t, batch, acquire, "the server learns about the company before the prewarm asks for its lease: %v", paths)
}

func TestClientModeIdentifyPrewarmIsSkippedAfterClose(t *testing.T) {
	server := startDataStreamServer(t, dataStreamFixture{})
	rec := &requestRecorder{}
	client := clientModeClient(t, rec, map[string]stub{
		leaseAcquirePath: {body: leaseResponse("lse_1", 10000)},
		eventBatchPath:   {body: map[string]any{"data": map[string]any{"events": []any{}}}},
	}, server.URL, core.CreditLeaseConfig{})

	client.Close()
	client.Identify(t.Context(), &schematicgo.EventBodyIdentify{
		Keys:    map[string]string{"id": "user_1"},
		Company: &schematicgo.EventBodyIdentifyCompany{Keys: map[string]string{"id": testCompanyID}},
	}, schematicclient.WithIdentifyPrewarm([]string{testCreditID}))
	time.Sleep(100 * time.Millisecond)

	_, acquired := rec.find(leaseAcquirePath)
	assert.False(t, acquired, "a closed client has torn down the lease plumbing a prewarm would touch")
}

// Lease state that lives in one process gates that process only, which is the
// misconfiguration that silently loses cross-pod gating.
func TestClientModeWithoutSharedRedisWarns(t *testing.T) {
	server := startDataStreamServer(t, dataStreamFixture{})
	logger := &capturingLogger{}
	client := clientModeClient(t, &requestRecorder{}, map[string]stub{
		eventBatchPath: {body: map[string]any{"data": map[string]any{"events": []any{}}}},
	}, server.URL, core.CreditLeaseConfig{}, option.WithLogger(logger))
	defer client.Close()

	assert.True(t, containsMatch(logger.warnings(), "without a shared Redis backend"), logger.warnings())
}

func TestClientModeCheckHoldsAFractionalUsage(t *testing.T) {
	server := startDataStreamServer(t, dataStreamFixture{pushFlags: true, serveCompany: true})
	rec := &requestRecorder{}
	client := clientModeClient(t, rec, map[string]stub{
		leaseAcquirePath: {body: leaseResponse("lse_1", 10000)},
		eventBatchPath:   {body: map[string]any{"data": map[string]any{"events": []any{}}}},
	}, server.URL, core.CreditLeaseConfig{})
	defer client.Close()

	result := client.Check(t.Context(), &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"id": testCompanyID},
	}, testFlagKey,
		schematicclient.WithUsage(2.5),
		schematicclient.WithEventSubtype(testEventSubtype),
	)

	require.True(t, result.Allowed, "reason: %s, error: %s", result.Reason, result.Error)
	require.NotNil(t, result.Reservation)
	assert.Equal(t, 2.5, result.Reservation.QuantityReserved, "the hold records the quantity the caller declared")
	// The server bills whole events, so the credits the hold costs are sized
	// from the rounded-up quantity the settle will charge.
	assert.Equal(t, 3.0, result.Reservation.CreditsReserved, "the fixture rate is one credit per unit")
}

// overlappingAcquires answers a lease acquire only once the expected number of
// them are in flight together, so a sequential prewarm cannot pass.
type overlappingAcquires struct {
	t        *testing.T
	want     int
	mu       sync.Mutex
	inFlight int
	overlap  chan struct{}
	closing  sync.Once
}

func (o *overlappingAcquires) Do(req *http.Request) (*http.Response, error) {
	body := func(payload any) (*http.Response, error) {
		data, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader(data))}, nil
	}

	if !strings.HasSuffix(req.URL.Path, leaseAcquirePath) {
		return body(map[string]any{"data": map[string]any{"events": []any{}}})
	}

	var request struct {
		CreditTypeID string `json:"credit_type_id"`
	}
	raw, _ := io.ReadAll(req.Body)
	require.NoError(o.t, json.Unmarshal(raw, &request))

	o.mu.Lock()
	o.inFlight++
	if o.inFlight >= o.want {
		o.closing.Do(func() { close(o.overlap) })
	}
	o.mu.Unlock()

	select {
	case <-o.overlap:
	case <-time.After(2 * time.Second):
		return nil, errors.New("the acquires never overlapped")
	}

	return body(&schematicgo.AcquireCreditLeaseResponse{
		Data: &schematicgo.CreditLeaseResponseData{
			ID:            "lse_" + request.CreditTypeID,
			CompanyID:     testCompanyID,
			CreditTypeID:  request.CreditTypeID,
			GrantedAmount: 10000,
			ExpiresAt:     time.Now().UTC().Add(5 * time.Minute),
		},
	})
}

func TestClientModePrewarmAcquiresCreditTypesConcurrently(t *testing.T) {
	server := startDataStreamServer(t, dataStreamFixture{})
	httpClient := &overlappingAcquires{t: t, want: 2, overlap: make(chan struct{})}

	client := schematicclient.NewSchematicClient(
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(httpClient),
		option.WithBaseURL(server.URL),
		option.WithEventBufferPeriod(10*time.Millisecond),
		option.WithDatastream(),
		option.WithCreditLeases(core.CreditLeaseConfig{Mode: core.CreditLeaseModeClient}),
	)
	defer client.Close()
	time.Sleep(300 * time.Millisecond)

	err := client.Prewarm(t.Context(), &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"id": testCompanyID},
	}, []string{testCreditID, "ct_2"})

	require.NoError(t, err)
}

// The event worker is the only reader of the client's event channel, so it must
// never block on the network. An identify's prewarm asks it for a flush; if the
// worker sent that batch itself, its retries would back up every later event,
// a client-mode check's flag_check included.
func TestClientModeIdentifyFlushDoesNotBlockEventEnqueue(t *testing.T) {
	server := startDataStreamServer(t, dataStreamFixture{})
	rec := &requestRecorder{}
	blocked := make(chan struct{})
	// A buffer period no test would outlast, so the only send in play is the
	// one the identify's prewarm asks for.
	client := clientModeClient(t, rec, map[string]stub{
		leaseAcquirePath: {body: leaseResponse("lse_1", 10000)},
		eventBatchPath:   {body: map[string]any{"data": map[string]any{"events": []any{}}}, block: blocked},
	}, server.URL, core.CreditLeaseConfig{}, option.WithEventBufferPeriod(time.Minute))
	defer client.Close()
	defer close(blocked)

	client.Identify(t.Context(), &schematicgo.EventBodyIdentify{
		Keys:    map[string]string{"id": "user_1"},
		Company: &schematicgo.EventBodyIdentifyCompany{Keys: map[string]string{"id": testCompanyID}},
	}, schematicclient.WithIdentifyPrewarm([]string{testCreditID}))
	require.Eventually(t, func() bool {
		_, sending := rec.find(eventBatchPath)
		return sending
	}, 2*time.Second, 10*time.Millisecond, "the flush should reach the API, where it stalls")

	// More events than the channel holds, so anything blocking the worker
	// blocks the caller too.
	enqueued := make(chan struct{})
	go func() {
		defer close(enqueued)
		for i := 0; i < 150; i++ {
			client.Track(t.Context(), &schematicgo.EventBodyTrack{Event: testEventSubtype})
		}
	}()
	select {
	case <-enqueued:
	case <-time.After(5 * time.Second):
		t.Fatal("enqueueing events blocked behind the stalled flush")
	}
}

// Server mode keeps no lease to warm, so an identify that asks for a prewarm
// has nothing to flush for and its events ride the buffer as usual.
func TestServerModeIdentifyPrewarmDoesNotFlush(t *testing.T) {
	rec := &requestRecorder{}
	client := serverModeClient(t, rec, map[string]stub{
		eventBatchPath: {body: map[string]any{"data": map[string]any{"events": []any{}}}},
	}, option.WithEventBufferPeriod(time.Minute))

	client.Identify(t.Context(), &schematicgo.EventBodyIdentify{
		Keys:    map[string]string{"id": "user_1"},
		Company: &schematicgo.EventBodyIdentifyCompany{Keys: map[string]string{"id": testCompanyID}},
	}, schematicclient.WithIdentifyPrewarm([]string{testCreditID}))
	time.Sleep(200 * time.Millisecond)

	_, flushed := rec.find(eventBatchPath)
	assert.False(t, flushed, "no flush is owed when there is no lease to warm: %v", rec.paths())
}

// Client mode with nothing to gate against: the flag never reaches the cache,
// so the check falls back to the plain API check, and that fails too. The
// caller's fail-open choice is what decides the verdict, not the flag default.
func TestClientModeCheckFailsOpenWhenDataStreamAndTheAPIAreBothDown(t *testing.T) {
	server := startDataStreamServer(t, dataStreamFixture{})
	rec := &requestRecorder{}
	client := clientModeClient(t, rec, map[string]stub{
		checkPath:      {status: http.StatusInternalServerError, body: map[string]any{"error": "boom"}},
		eventBatchPath: {body: map[string]any{"data": map[string]any{"events": []any{}}}},
	}, server.URL, core.CreditLeaseConfig{})
	defer client.Close()

	evalCtx := &schematicgo.CheckFlagRequestBody{Company: map[string]string{"id": testCompanyID}}
	denied := client.Check(t.Context(), evalCtx, testFlagKey,
		schematicclient.WithUsage(50),
		schematicclient.WithEventSubtype(testEventSubtype),
	)
	assert.False(t, denied.Allowed)
	assert.Equal(t, "check_failed", denied.Reason)
	assert.NotEmpty(t, denied.Error)

	allowed := client.Check(t.Context(), evalCtx, testFlagKey,
		schematicclient.WithUsage(50),
		schematicclient.WithEventSubtype(testEventSubtype),
		schematicclient.WithFailOpen(),
		schematicclient.WithCheckDefault(true),
	)
	assert.True(t, allowed.Allowed, "reason: %s, error: %s", allowed.Reason, allowed.Error)
	assert.Equal(t, "check_failed_fail_open", allowed.Reason)
	assert.NotEmpty(t, allowed.Error)
}

// A prewarm stuck on the wire used to eat the whole close budget, leaving the
// releases to run on a context that had already expired, so every one of them
// failed and the credits sat held until the leases expired server-side.
func TestCloseReleasesLeasesWhileAPrewarmIsStuckOnTheWire(t *testing.T) {
	server := startDataStreamServer(t, dataStreamFixture{pushFlags: true, serveCompany: true})
	rec := &requestRecorder{}
	blocked := make(chan struct{})
	defer close(blocked)
	// A buffer period no test would outlast, so the only send in play is the
	// one the prewarm asks for, and it never comes back.
	client := clientModeClient(t, rec, map[string]stub{
		leaseAcquirePath: {body: leaseResponse("lse_1", 10000)},
		leaseReleasePath: {body: leaseResponse("lse_1", 10000)},
		eventBatchPath:   {body: map[string]any{"data": map[string]any{"events": []any{}}}, block: blocked},
	}, server.URL, core.CreditLeaseConfig{}, option.WithEventBufferPeriod(time.Minute))

	// A lease for the close to hand back.
	result := client.Check(t.Context(), &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"id": testCompanyID},
	}, testFlagKey, schematicclient.WithUsage(50), schematicclient.WithEventSubtype(testEventSubtype))
	require.NotNil(t, result.Reservation)

	// The prewarm flushes the identify before warming anything, and this API
	// never answers that flush, so the prewarm is stuck on the wire.
	client.Identify(t.Context(), &schematicgo.EventBodyIdentify{
		Keys:    map[string]string{"id": "user_1"},
		Company: &schematicgo.EventBodyIdentifyCompany{Keys: map[string]string{"id": testCompanyID}},
	}, schematicclient.WithIdentifyPrewarm([]string{testCreditID}))
	require.Eventually(t, func() bool {
		_, flushing := rec.find(eventBatchPath)
		return flushing
	}, 2*time.Second, 10*time.Millisecond, "the flush should reach the API, where it stalls")

	start := time.Now()
	client.Close()

	assert.Less(t, time.Since(start), 2*time.Second, "a stalled prewarm is cancelled, not waited out")
	released, ok := rec.find(leaseReleasePath)
	require.True(t, ok, "the lease is released on close: %v", rec.paths())
	assert.False(t, released.onDeadCtx, "the release runs on a context that still has budget")
}

// An acquire that started before the close can still install a lease. The close
// waits it out before listing, so the lease it installs is released rather than
// left held until it expires server-side.
func TestCloseReleasesALeaseInstalledByAnAcquireRacingIt(t *testing.T) {
	server := startDataStreamServer(t, dataStreamFixture{pushFlags: true, serveCompany: true})
	rec := &requestRecorder{}
	acquiring := make(chan struct{})
	client := clientModeClient(t, rec, map[string]stub{
		leaseAcquirePath: {body: leaseResponse("lse_1", 10000), block: acquiring},
		leaseReleasePath: {body: leaseResponse("lse_1", 10000)},
		eventBatchPath:   {body: map[string]any{"data": map[string]any{"events": []any{}}}},
	}, server.URL, core.CreditLeaseConfig{})

	checked := make(chan struct{})
	go func() {
		defer close(checked)
		client.Check(t.Context(), &schematicgo.CheckFlagRequestBody{
			Company: map[string]string{"id": testCompanyID},
		}, testFlagKey, schematicclient.WithUsage(50), schematicclient.WithEventSubtype(testEventSubtype))
	}()
	require.Eventually(t, func() bool {
		_, acquired := rec.find(leaseAcquirePath)
		return acquired
	}, 2*time.Second, 10*time.Millisecond, "the acquire should be in flight")

	closed := make(chan struct{})
	go func() {
		defer close(closed)
		client.Close()
	}()
	// Let the close reach its wait before the acquire lands, so the lease is
	// installed while the client is already shutting down.
	time.Sleep(200 * time.Millisecond)
	close(acquiring)

	select {
	case <-closed:
	case <-time.After(15 * time.Second):
		t.Fatal("the close never finished")
	}
	<-checked

	released, ok := rec.find(leaseReleasePath)
	require.True(t, ok, "no lease survives the close: %v", rec.paths())
	assert.Contains(t, released.path, "lse_1")
}

// The ticket's case, end to end with the first extend held mid-flight: a
// sub-watermark check fires a fire-and-forget extend for one tranche, and a
// check needing more than that tranche arrives while it is in flight.
// Inheriting the tranche leaves the second check failing
// insufficient_lease_balance with the credits sitting on the server.
func TestClientModeCheckThatNeedsMoreThanTheExtendInFlightAskedFor(t *testing.T) {
	server := startDataStreamServer(t, dataStreamFixture{pushFlags: true, serveCompany: true})
	rec := &requestRecorder{}
	releaseExtend := make(chan struct{})
	client := clientModeClient(t, rec, map[string]stub{
		leaseAcquirePath: {body: leaseResponse("lse_1", 1000)},
		leaseExtendPath: {
			// The server's authoritative total after each extend.
			bodies: []any{leaseResponse("lse_1", 2000), leaseResponse("lse_1", 3000)},
			block:  releaseExtend,
		},
		leaseReleasePath: {body: leaseResponse("lse_1", 3000)},
		eventBatchPath:   {body: map[string]any{"data": map[string]any{"events": []any{}}}},
	}, server.URL, core.CreditLeaseConfig{DefaultLeaseSize: 1000, LowWaterMark: 0.25})
	defer client.Close()

	evalCtx := &schematicgo.CheckFlagRequestBody{Company: map[string]string{"id": testCompanyID}}
	// 800 of the 1000-credit lease, leaving 200: below the water mark, so this
	// check's fire-and-forget extend goes out and hangs.
	first := client.Check(t.Context(), evalCtx, testFlagKey,
		schematicclient.WithUsage(800),
		schematicclient.WithEventSubtype(testEventSubtype),
	)
	require.True(t, first.Allowed, "reason: %s, error: %s", first.Reason, first.Error)
	require.Eventually(t, func() bool {
		return countPaths(rec, leaseExtendPath) == 1
	}, 2*time.Second, 10*time.Millisecond, "the watermark extend should be in flight")

	// 1500 against 200 remaining: the reserve fails and the check asks for an
	// extend, joining the tranche-sized flight.
	second := make(chan *schematicclient.CheckResult, 1)
	go func() {
		second <- client.Check(t.Context(), evalCtx, testFlagKey,
			schematicclient.WithUsage(1500),
			schematicclient.WithEventSubtype(testEventSubtype),
		)
	}()
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, 1, countPaths(rec, leaseExtendPath), "the joiner waits the flight out rather than racing it")

	close(releaseExtend)
	result := <-second

	require.True(t, result.Allowed, "reason: %s, error: %s", result.Reason, result.Error)
	require.NotNil(t, result.Reservation)
	assert.Equal(t, 1500.0, result.Reservation.CreditsReserved)
	assert.GreaterOrEqual(t, countPaths(rec, leaseExtendPath), 2, "the joiner topped the lease up itself: %v", rec.paths())
}

// countPaths is how many recorded requests ended in suffix.
func countPaths(rec *requestRecorder, suffix string) int {
	count := 0
	for _, path := range rec.paths() {
		if strings.HasSuffix(path, suffix) {
			count++
		}
	}
	return count
}

// An account is free to define an entity key called `id` holding its own
// identifier, so every key is looked up before any value is read as a Schematic
// company id.
func TestClientModePrewarmResolvesAnAccountDefinedIDKeyThroughTheLookup(t *testing.T) {
	server := startDataStreamServer(t, dataStreamFixture{serveCompany: true})
	rec := &requestRecorder{}
	client := clientModeClient(t, rec, map[string]stub{
		leaseAcquirePath: {body: leaseResponse("lse_1", 10000)},
		eventBatchPath:   {body: map[string]any{"data": map[string]any{"events": []any{}}}},
	}, server.URL, core.CreditLeaseConfig{})
	defer client.Close()

	err := client.Prewarm(t.Context(), &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"id": "acme"},
	}, []string{testCreditID})

	require.NoError(t, err)
	req, ok := rec.find(leaseAcquirePath)
	require.True(t, ok)
	assert.Equal(t, testCompanyID, decodeBody(t, req.body)["company_id"], "the lease is keyed by the id the lookup returned")
}

// Only once the keys resolve nothing is a value read as the company's own id, by
// its `comp_` prefix rather than by the name of the key it sits under.
func TestClientModePrewarmFallsBackToASchematicIDTheKeysCarry(t *testing.T) {
	server := startDataStreamServer(t, dataStreamFixture{})
	rec := &requestRecorder{}
	noWait := time.Duration(0)
	// The server answers with the company it was asked for, which here is the
	// id the keys carried.
	granted := leaseResponse("lse_1", 10000)
	granted.Data.CompanyID = "comp_9"
	client := clientModeClient(t, rec, map[string]stub{
		leaseAcquirePath: {body: granted},
		eventBatchPath:   {body: map[string]any{"data": map[string]any{"events": []any{}}}},
	}, server.URL, core.CreditLeaseConfig{PrewarmResolveTimeout: &noWait})
	defer client.Close()

	err := client.Prewarm(t.Context(), &schematicgo.CheckFlagRequestBody{
		Company: map[string]string{"account_id": "comp_9"},
	}, []string{testCreditID})

	require.NoError(t, err)
	req, ok := rec.find(leaseAcquirePath)
	require.True(t, ok)
	assert.Equal(t, "comp_9", decodeBody(t, req.body)["company_id"])
}
