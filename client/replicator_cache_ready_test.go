package client_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	schematicgo "github.com/schematichq/schematic-go"
	schematicclient "github.com/schematichq/schematic-go/client"
	"github.com/schematichq/schematic-go/core"
	"github.com/schematichq/schematic-go/mocks"
	option "github.com/schematichq/schematic-go/option"
	"github.com/schematichq/schematic-go/rulesengine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// In replicator mode, single and bulk flag checks read the replicator's cache
// only once its health endpoint reports ready. Before that both take the API
// path. The fixtures are arranged so the cache and the API disagree on every
// flag: the cache says true for both, the API and the defaults say false, so a
// result shows which one answered.

const (
	replicatorCacheVersion = "v1"
	replicatorCompanyID    = "comp_1"
	apiReason              = "answered by the api"
)

var replicatorFlagKeys = []string{"flag-a", "flag-b"}

// seedReplicatorCache writes flags and a company into Redis under the keys and
// encoding the replicator uses: JSON values, versioned keys, and a company
// stored by ID with a lookup key per company key pointing at that ID.
func seedReplicatorCache(t *testing.T, mr *miniredis.Miniredis) {
	t.Helper()
	set := func(key string, value any) {
		data, err := json.Marshal(value)
		require.NoError(t, err)
		require.NoError(t, mr.Set(key, string(data)))
	}

	// flag-a is true only for the seeded company, so a true verdict also shows
	// the company came from the cache. flag-b is true for everyone.
	set("schematic:flags:"+replicatorCacheVersion+":flag-a", &rulesengine.Flag{
		ID:            "flag_a",
		AccountID:     "acct",
		EnvironmentID: "env",
		Key:           "flag-a",
		DefaultValue:  false,
		Rules: rulesengine.JSONSlice[*rulesengine.Rule]{{
			ID:            "rule_a",
			AccountID:     "acct",
			EnvironmentID: "env",
			RuleType:      rulesengine.RuleTypeStandard,
			Name:          "Seeded company",
			Priority:      1,
			Value:         true,
			Conditions: rulesengine.JSONSlice[*rulesengine.Condition]{{
				ID:            "cond_a",
				AccountID:     "acct",
				EnvironmentID: "env",
				ConditionType: rulesengine.ConditionTypeCompany,
				Operator:      rulesengine.ComparableOperatorEquals,
				ResourceIDs:   rulesengine.JSONSlice[string]{replicatorCompanyID},
			}},
		}},
	})
	set("schematic:flags:"+replicatorCacheVersion+":flag-b", &rulesengine.Flag{
		ID:            "flag_b",
		AccountID:     "acct",
		EnvironmentID: "env",
		Key:           "flag-b",
		DefaultValue:  true,
	})

	set("schematic:company:"+replicatorCacheVersion+":"+replicatorCompanyID, &rulesengine.Company{
		ID:            replicatorCompanyID,
		AccountID:     "acct",
		EnvironmentID: "env",
		Keys:          map[string]string{"id": replicatorCompanyID},
		Metrics:       make(rulesengine.CompanyMetricCollection, 0),
		Traits:        make(rulesengine.JSONSlice[*rulesengine.Trait], 0),
	})
	set("schematic:company:"+replicatorCacheVersion+":id:"+replicatorCompanyID, replicatorCompanyID)
}

// startReplicatorHealth serves a replicator /ready endpoint with the given
// status and readiness, counting the polls it answers.
func startReplicatorHealth(t *testing.T, ready bool, polls *atomic.Int32) *httptest.Server {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		defer polls.Add(1)
		w.Header().Set("Content-Type", "application/json")
		if !ready {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		_ = json.NewEncoder(w).Encode(map[string]any{
			"ready":         ready,
			"cache_version": replicatorCacheVersion,
		})
	}))
	t.Cleanup(server.Close)
	return server
}

// replicatorAPI answers flag checks false with apiReason, or with a 500 when
// failing is set, and accepts event batches. It counts the flag-check calls.
func replicatorAPI(failing bool, flagChecks *atomic.Int32) func(*http.Request) (*http.Response, error) {
	respond := func(status int, body any) (*http.Response, error) {
		data, _ := json.Marshal(body)
		return &http.Response{
			Status:     http.StatusText(status),
			StatusCode: status,
			Header:     http.Header{"Content-Type": []string{"application/json"}},
			Body:       io.NopCloser(bytes.NewReader(data)),
		}, nil
	}
	return func(req *http.Request) (*http.Response, error) {
		path := req.URL.Path
		if !strings.HasSuffix(path, "/check") {
			return respond(http.StatusOK, map[string]any{})
		}
		flagChecks.Add(1)
		if failing {
			return respond(http.StatusInternalServerError, map[string]any{"error": "unavailable"})
		}
		if strings.HasSuffix(path, "/flags/check") {
			flags := make([]*schematicgo.CheckFlagResponseData, 0, len(replicatorFlagKeys))
			for _, key := range replicatorFlagKeys {
				flags = append(flags, &schematicgo.CheckFlagResponseData{Flag: key, Value: false, Reason: apiReason})
			}
			return respond(http.StatusOK, &schematicgo.CheckFlagsResponse{
				Data: &schematicgo.CheckFlagsResponseData{Flags: flags},
			})
		}
		key := strings.TrimSuffix(strings.TrimPrefix(path[strings.LastIndex(path, "/flags/"):], "/flags/"), "/check")
		return respond(http.StatusOK, &schematicgo.CheckFlagResponse{
			Data: &schematicgo.CheckFlagResponseData{Flag: key, Value: false, Reason: apiReason},
		})
	}
}

type replicatorFixture struct {
	ready      bool
	apiFailing bool
}

// newReplicatorClient builds a replicator-mode client over a seeded Redis and
// waits until it has processed at least one health poll.
func newReplicatorClient(t *testing.T, fixture replicatorFixture) (*schematicclient.SchematicClient, *atomic.Int32) {
	t.Helper()
	mr := miniredis.RunT(t)
	seedReplicatorCache(t, mr)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })

	var polls atomic.Int32
	health := startReplicatorHealth(t, fixture.ready, &polls)

	var flagChecks atomic.Int32
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(replicatorAPI(fixture.apiFailing, &flagChecks)).AnyTimes()

	client := schematicclient.NewSchematicClient(
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(mockHTTPClient),
		option.WithoutRetries(),
		// The API path's own cache would answer a repeat check without a call.
		option.WithDisableFlagCheckCache(),
		option.WithDefaultFlagValues(map[string]bool{"flag-a": false, "flag-b": false}),
		option.WithDatastream(
			option.WithReplicatorMode(),
			option.WithReplicatorHealthURL(health.URL),
			option.WithReplicatorHealthInterval(10*time.Millisecond),
			option.WithRedisClient(rdb),
			option.WithCacheTTL(0),
		),
	)
	t.Cleanup(client.Close)

	// Polls run one after another, so a second poll means the first one's
	// answer has been applied.
	require.Eventually(t, func() bool { return polls.Load() >= 2 }, 2*time.Second, 5*time.Millisecond)
	return client, &flagChecks
}

var replicatorEvalCtx = &schematicgo.CheckFlagRequestBody{
	Company: map[string]string{"id": replicatorCompanyID},
}

func TestReplicatorNotReadySingleAndBulkUseTheAPI(t *testing.T) {
	client, flagChecks := newReplicatorClient(t, replicatorFixture{ready: false})
	ctx := context.Background()

	for _, key := range replicatorFlagKeys {
		assert.False(t, client.CheckFlag(ctx, replicatorEvalCtx, key), "CheckFlag %s", key)

		resp, err := client.CheckFlagWithEntitlement(ctx, replicatorEvalCtx, key)
		require.NoError(t, err)
		assert.False(t, resp.Value, "CheckFlagWithEntitlement %s", key)
		assert.Equal(t, apiReason, resp.Reason, "CheckFlagWithEntitlement %s", key)
	}

	results := client.CheckFlags(ctx, replicatorEvalCtx, replicatorFlagKeys)
	require.Len(t, results, len(replicatorFlagKeys))
	for i, key := range replicatorFlagKeys {
		assert.Equal(t, key, results[i].FlagKey)
		assert.False(t, results[i].Value, "CheckFlags %s", key)
		assert.Equal(t, apiReason, results[i].Reason, "CheckFlags %s", key)
	}

	// Two single checks per key, then one bulk call.
	assert.Equal(t, int32(2*len(replicatorFlagKeys)+1), flagChecks.Load())
}

func TestReplicatorNotReadyWithTheAPIDownReturnsDefaults(t *testing.T) {
	client, flagChecks := newReplicatorClient(t, replicatorFixture{ready: false, apiFailing: true})
	ctx := context.Background()

	for _, key := range replicatorFlagKeys {
		assert.False(t, client.CheckFlag(ctx, replicatorEvalCtx, key), "CheckFlag %s", key)
	}

	results := client.CheckFlags(ctx, replicatorEvalCtx, replicatorFlagKeys)
	require.Len(t, results, len(replicatorFlagKeys))
	for i, key := range replicatorFlagKeys {
		assert.Equal(t, key, results[i].FlagKey)
		assert.False(t, results[i].Value, "CheckFlags %s", key)
		assert.NotEqual(t, apiReason, results[i].Reason, "CheckFlags %s", key)
	}

	assert.Equal(t, int32(len(replicatorFlagKeys)+1), flagChecks.Load(), "every check tried the API")
}

func TestReplicatorReadySingleAndBulkEvaluateFromTheCache(t *testing.T) {
	client, flagChecks := newReplicatorClient(t, replicatorFixture{ready: true})
	ctx := context.Background()

	results := client.CheckFlags(ctx, replicatorEvalCtx, replicatorFlagKeys)
	require.Len(t, results, len(replicatorFlagKeys))

	for i, key := range replicatorFlagKeys {
		assert.True(t, client.CheckFlag(ctx, replicatorEvalCtx, key), "CheckFlag %s", key)

		single, err := client.CheckFlagWithEntitlement(ctx, replicatorEvalCtx, key)
		require.NoError(t, err)
		assert.True(t, single.Value, "CheckFlagWithEntitlement %s", key)

		assert.Equal(t, key, results[i].FlagKey)
		assert.True(t, results[i].Value, "CheckFlags %s", key)
		assert.Equal(t, single.Reason, results[i].Reason, "single and bulk agree on %s", key)
		assert.Equal(t, single.RuleID, results[i].RuleID, "single and bulk agree on %s", key)
	}

	assert.Zero(t, flagChecks.Load(), "a ready cache answers without the API")
}

// A credit check in client lease mode reads the flag and company from the same
// cache, so it waits on the same readiness: before the replicator is ready it
// runs as a plain API check and takes no lease, and once ready it gates on a
// local lease.
func TestReplicatorClientModeCheckWaitsForTheCacheToBeReady(t *testing.T) {
	mr := miniredis.RunT(t)
	cacheKey := func(parts ...string) string {
		return "schematic:" + parts[0] + ":" + replicatorCacheVersion + ":" + strings.Join(parts[1:], ":")
	}
	set := func(key string, value any) {
		data, err := json.Marshal(value)
		require.NoError(t, err)
		require.NoError(t, mr.Set(key, string(data)))
	}
	companyKeys := map[string]string{"id": testCompanyID}
	set(cacheKey("flags", testFlagKey), creditFlag())
	set(cacheKey("company", testCompanyID), creditCompany(companyKeys))
	set(cacheKey("company", "id", testCompanyID), testCompanyID)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })

	var ready atomic.Bool
	var polls atomic.Int32
	health := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		defer polls.Add(1)
		w.Header().Set("Content-Type", "application/json")
		if !ready.Load() {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		_ = json.NewEncoder(w).Encode(map[string]any{
			"ready":         ready.Load(),
			"cache_version": replicatorCacheVersion,
		})
	}))
	t.Cleanup(health.Close)
	// Polls run one after another, so two more polls than now means one that
	// started after this call has been applied.
	waitForPoll := func() {
		seen := polls.Load()
		require.Eventually(t, func() bool { return polls.Load() >= seen+2 }, 2*time.Second, 5*time.Millisecond)
	}

	rec := &requestRecorder{}
	ctrl := gomock.NewController(t)
	mockHTTPClient := mocks.NewMockHTTPClient(ctrl)
	mockHTTPClient.EXPECT().Do(gomock.Any()).DoAndReturn(serveStubs(t, rec, map[string]stub{
		leaseAcquirePath: {body: leaseResponse("lse_1", 10000)},
		"/flags/" + testFlagKey + "/check": {body: &schematicgo.CheckFlagResponse{
			Data: &schematicgo.CheckFlagResponseData{Flag: testFlagKey, Value: true, Reason: apiReason},
		}},
		eventBatchPath: {body: map[string]any{"data": map[string]any{"events": []any{}}}},
	})).AnyTimes()

	client := schematicclient.NewSchematicClient(
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(mockHTTPClient),
		option.WithoutRetries(),
		option.WithDisableFlagCheckCache(),
		option.WithEventBufferPeriod(10*time.Millisecond),
		option.WithDatastream(
			option.WithReplicatorMode(),
			option.WithReplicatorHealthURL(health.URL),
			option.WithReplicatorHealthInterval(10*time.Millisecond),
			option.WithRedisClient(rdb),
			option.WithCacheTTL(0),
		),
		option.WithCreditLeases(core.CreditLeaseConfig{Mode: core.CreditLeaseModeClient}),
	)
	t.Cleanup(client.Close)
	waitForPoll()

	evalCtx := &schematicgo.CheckFlagRequestBody{Company: companyKeys}
	check := func() *schematicclient.CheckResult {
		return client.Check(t.Context(), evalCtx, testFlagKey,
			schematicclient.WithUsage(50),
			schematicclient.WithEventSubtype(testEventSubtype),
		)
	}

	notReady := check()
	assert.True(t, notReady.Allowed, "reason: %s, error: %s", notReady.Reason, notReady.Error)
	assert.Equal(t, apiReason, notReady.Reason, "the API answered")
	assert.Nil(t, notReady.Reservation, "a plain check holds nothing")
	assert.Equal(t, 1, countPaths(rec, "/flags/"+testFlagKey+"/check"))
	assert.Zero(t, countPaths(rec, leaseAcquirePath), "no lease before the cache is ready")

	ready.Store(true)
	waitForPoll()

	result := check()
	require.True(t, result.Allowed, "reason: %s, error: %s", result.Reason, result.Error)
	require.NotNil(t, result.Reservation)
	assert.Equal(t, core.CreditLeaseModeClient, result.Reservation.Mode)
	assert.Equal(t, "lse_1", result.Reservation.LeaseID)
	assert.Equal(t, 50.0, result.Reservation.CreditsReserved)
	assert.Equal(t, 1, countPaths(rec, leaseAcquirePath), "a ready cache gates on a lease")
	assert.Equal(t, 1, countPaths(rec, "/flags/"+testFlagKey+"/check"), "and evaluates without the API")
}
