package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	schematicgo "github.com/schematichq/schematic-go"
	option "github.com/schematichq/schematic-go/option"
	"github.com/schematichq/schematic-go/rulesengine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// A replicator that has lost Schematic (a closed account, an outage) stays up
// and keeps its Redis cache but reports not ready. Flag checks still evaluate
// from that cache, so Track has to keep moving the cached metric, or every
// metered limit freezes at whatever the replicator last wrote.
func TestTrackUpdatesCachedMetricsWhileReplicatorIsNotReady(t *testing.T) {
	const cacheVersion = "v1"
	health := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprintf(w, `{"ready": false, "cache_version": %q}`, cacheVersion)
	}))
	t.Cleanup(health.Close)

	mr := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })

	// Seed the company the way the replicator writes it: the entity under its
	// ID, and a lookup from each of its keys to that ID.
	company := &rulesengine.Company{
		ID:            "comp_1",
		AccountID:     "acct_1",
		EnvironmentID: "env_1",
		Keys:          map[string]string{"id": "comp-key"},
		Metrics: []*rulesengine.CompanyMetric{{
			CompanyID:    "comp_1",
			EventSubtype: "api-calls",
			Value:        10,
			Period:       rulesengine.MetricPeriodAllTime,
		}},
	}
	companyJSON, err := json.Marshal(company)
	require.NoError(t, err)
	require.NoError(t, mr.Set("schematic:company:"+cacheVersion+":comp_1", string(companyJSON)))
	require.NoError(t, mr.Set("schematic:company:"+cacheVersion+":id:comp-key", `"comp_1"`))

	client := NewSchematicClient(
		option.WithAPIKey("test-api-key"),
		option.WithHTTPClient(&pathRecorder{body: map[string]any{}}),
		option.WithDatastream(
			option.WithReplicatorMode(),
			option.WithReplicatorHealthURL(health.URL),
			option.WithReplicatorHealthInterval(time.Hour),
			option.WithRedisClient(rdb),
		),
	)
	t.Cleanup(client.Close)

	ds := client.datastreamClient
	require.Eventually(t, func() bool { return ds.GetCacheVersion() == cacheVersion },
		2*time.Second, 10*time.Millisecond, "the first health check should land")
	require.False(t, ds.IsConnected(), "the replicator reports not ready")

	keys := map[string]string{"id": "comp-key"}
	require.NotNil(t, ds.GetCachedCompany(keys))

	client.Track(context.Background(), &schematicgo.EventBodyTrack{
		Event:    "api-calls",
		Company:  keys,
		Quantity: schematicgo.Int64(5),
	})

	cached := ds.GetCachedCompany(keys)
	require.NotNil(t, cached)
	require.Len(t, cached.Metrics, 1)
	assert.Equal(t, int64(15), cached.Metrics[0].Value)
}
