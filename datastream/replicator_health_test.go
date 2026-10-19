package datastream_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
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

// healthServer serves a /ready body whose status and contents a test can
// change between polls.
type healthServer struct {
	mu      sync.Mutex
	status  int
	ready   bool
	version string
}

func (h *healthServer) set(status int, ready bool, version string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.status, h.ready, h.version = status, ready, version
}

func (h *healthServer) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	h.mu.Lock()
	status, ready, version := h.status, h.ready, h.version
	h.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{"ready": ready, "cache_version": version})
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
	health := &healthServer{}
	health.set(http.StatusServiceUnavailable, false, "vX")
	server := httptest.NewServer(health)
	t.Cleanup(server.Close)

	client := newReplicatorModeClient(t, server.URL)

	require.Eventually(t, func() bool { return client.GetCacheVersion() == "vX" }, 2*time.Second, 5*time.Millisecond,
		"the cache_version in a 503 body is recorded")
	assert.False(t, client.IsCacheReady())
	assert.False(t, client.IsConnected(), "IsConnected keeps reporting replicator readiness")

	health.set(http.StatusOK, true, "vX")
	require.Eventually(t, client.IsCacheReady, 2*time.Second, 5*time.Millisecond)
	assert.True(t, client.IsConnected())
}

func TestReplicatorHealthUnreachableIsNotReadyAndKeepsCacheVersion(t *testing.T) {
	health := &healthServer{}
	health.set(http.StatusOK, true, "v1")
	server := httptest.NewServer(health)

	client := newReplicatorModeClient(t, server.URL)
	require.Eventually(t, client.IsCacheReady, 2*time.Second, 5*time.Millisecond)
	require.Equal(t, "v1", client.GetCacheVersion())

	server.Close()

	require.Eventually(t, func() bool { return !client.IsCacheReady() }, 2*time.Second, 5*time.Millisecond,
		"a failed poll sets not ready")
	assert.Equal(t, "v1", client.GetCacheVersion(), "a failed poll keeps the last cache version")
}

func TestIsCacheReadyOutsideReplicatorMode(t *testing.T) {
	client := datastream.NewDataStreamClient(
		createTestClientOptions("", NewMockLogger(), "test-api-key"),
		&core.DatastreamOptions{CacheTTL: 5 * time.Minute},
	)
	t.Cleanup(client.Close)

	assert.True(t, client.IsCacheReady(), "WebSocket mode has no readiness gate")
}
