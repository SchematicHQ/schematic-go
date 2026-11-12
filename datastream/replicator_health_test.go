package datastream_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/schematichq/schematic-go/core"
	"github.com/schematichq/schematic-go/datastream"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The replicator's /ready endpoint answers 503 with a JSON body until its cache
// is complete, and 200 once it is. These pin how the health poll reads that
// and what IsCacheReady reports from it.

// healthMode is how healthServer answers a poll.
type healthMode int

const (
	// answer with a JSON body carrying ready and cache_version
	healthAnswer healthMode = iota
	// answer 502 with a body that is not JSON, as a proxy in front of a
	// restarting replicator might
	healthGarbled
	// close the connection without a response
	healthDrop
)

// healthServer serves a /ready endpoint whose answer a test can change between
// polls, and counts the polls it has finished serving.
type healthServer struct {
	mu      sync.Mutex
	mode    healthMode
	status  int
	ready   bool
	version string
	polls   atomic.Int32
	url     string
}

func (h *healthServer) set(status int, ready bool, version string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.mode, h.status, h.ready, h.version = healthAnswer, status, ready, version
}

func (h *healthServer) fail(mode healthMode) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.mode = mode
}

func (h *healthServer) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	defer h.polls.Add(1)
	h.mu.Lock()
	mode, status, ready, version := h.mode, h.status, h.ready, h.version
	h.mu.Unlock()
	switch mode {
	case healthGarbled:
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusBadGateway)
		_, _ = w.Write([]byte("<html>bad gateway</html>"))
	case healthDrop:
		conn, _, err := w.(http.Hijacker).Hijack()
		if err == nil {
			_ = conn.Close()
		}
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(map[string]any{"ready": ready, "cache_version": version})
	}
}

// waitForPolls waits until the client has finished at least n more polls than
// when it was called. Polls run one after another, so once the server has
// served two more, one that started after this call has been applied.
func (h *healthServer) waitForPolls(t *testing.T, n int32) {
	t.Helper()
	seen := h.polls.Load()
	require.Eventually(t, func() bool { return h.polls.Load() >= seen+n+1 }, 2*time.Second, 5*time.Millisecond)
}

func newReplicatorModeClient(t *testing.T, healthURL string) *datastream.DataStreamClient {
	t.Helper()
	client := datastream.NewDataStreamClient(
		createTestClientOptions("", NewMockLogger(), "test-api-key"),
		&core.DatastreamOptions{
			CacheTTL:              5 * time.Minute,
			ReplicatorMode:        true,
			ReplicatorHealthURL:   healthURL,
			ReplicatorHealthCheck: 10 * time.Millisecond,
		},
	)
	client.Start()
	t.Cleanup(client.Close)
	return client
}

func TestReplicatorHealth503RecordsCacheVersionAndIsNotReady(t *testing.T) {
	health := startHealthServer(t, http.StatusServiceUnavailable, false, "vX")
	client := newReplicatorModeClient(t, health.url)

	require.Eventually(t, func() bool { return client.GetCacheVersion() == "vX" }, 2*time.Second, 5*time.Millisecond,
		"the cache_version in a 503 body is recorded")
	assert.False(t, client.IsCacheReady())
	assert.False(t, client.IsConnected(), "IsConnected keeps reporting replicator readiness")

	health.set(http.StatusOK, true, "vX")
	require.Eventually(t, client.IsCacheReady, 2*time.Second, 5*time.Millisecond)
	assert.True(t, client.IsConnected())
}

// startHealthServer serves a healthServer answering with the given body.
func startHealthServer(t *testing.T, status int, ready bool, version string) *healthServer {
	t.Helper()
	health := &healthServer{}
	health.set(status, ready, version)
	server := httptest.NewServer(health)
	t.Cleanup(server.Close)
	health.url = server.URL
	return health
}

func TestReplicatorHealthFailedPollKeepsReady(t *testing.T) {
	for name, mode := range map[string]healthMode{"garbled": healthGarbled, "dropped": healthDrop} {
		t.Run(name, func(t *testing.T) {
			health := startHealthServer(t, http.StatusOK, true, "v1")
			client := newReplicatorModeClient(t, health.url)
			require.Eventually(t, client.IsCacheReady, 2*time.Second, 5*time.Millisecond)

			health.fail(mode)
			health.waitForPolls(t, 3)

			assert.True(t, client.IsCacheReady(), "a poll with no answer keeps the cache ready")
			assert.Equal(t, "v1", client.GetCacheVersion(), "and keeps the cache version")
		})
	}
}

func TestReplicatorHealthFailedPollKeepsNotReady(t *testing.T) {
	health := startHealthServer(t, http.StatusServiceUnavailable, false, "v1")
	client := newReplicatorModeClient(t, health.url)
	require.Eventually(t, func() bool { return client.GetCacheVersion() == "v1" }, 2*time.Second, 5*time.Millisecond)

	health.fail(healthGarbled)
	health.waitForPolls(t, 3)

	assert.False(t, client.IsCacheReady(), "a poll with no answer does not make a loading cache ready")
	assert.Equal(t, "v1", client.GetCacheVersion())
}

func TestReplicatorHealthNeverAnsweredIsNotReady(t *testing.T) {
	health := startHealthServer(t, 0, false, "")
	health.fail(healthGarbled)
	client := newReplicatorModeClient(t, health.url)
	health.waitForPolls(t, 3)

	assert.False(t, client.IsCacheReady(), "no answer yet means not ready")
	assert.Empty(t, client.GetCacheVersion())
}

func TestReplicatorHealthReadyFalseAfterReadyIsNotReady(t *testing.T) {
	health := startHealthServer(t, http.StatusOK, true, "v1")
	client := newReplicatorModeClient(t, health.url)
	require.Eventually(t, client.IsCacheReady, 2*time.Second, 5*time.Millisecond)

	health.set(http.StatusServiceUnavailable, false, "v2")
	require.Eventually(t, func() bool { return !client.IsCacheReady() }, 2*time.Second, 5*time.Millisecond,
		"an explicit ready false gates the cache")
	assert.Equal(t, "v2", client.GetCacheVersion())
}

func TestIsCacheReadyOutsideReplicatorMode(t *testing.T) {
	client := datastream.NewDataStreamClient(
		createTestClientOptions("", NewMockLogger(), "test-api-key"),
		&core.DatastreamOptions{CacheTTL: 5 * time.Minute},
	)
	t.Cleanup(client.Close)

	assert.True(t, client.IsCacheReady(), "WebSocket mode has no readiness gate")
}
