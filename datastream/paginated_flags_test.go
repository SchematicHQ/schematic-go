package datastream

import (
	"context"
	"encoding/json"
	"sync"
	"testing"
	"time"

	schematicdatastreamws "github.com/schematichq/schematic-datastream-ws"
	"github.com/schematichq/schematic-go/rulesengine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The server answers a flags subscription in pages when the request asks for
// them. What these pin is that the client treats the set as complete only on
// the last page. Applying a page as if it were the whole snapshot would delete
// every flag the other pages carry and would release a caller waiting on
// bootstrap with a partial set -- both silently, which is why the timing of
// DeleteMissing is worth asserting rather than describing.

// recordingFlagCache is a CacheProvider that keeps what it was given, so a test
// can see which flags are cached and when the delete of everything absent from
// the snapshot ran.
type recordingFlagCache struct {
	mu                 sync.Mutex
	flags              map[string]*rulesengine.Flag
	deleteMissingCalls [][]string
}

func newRecordingFlagCache() *recordingFlagCache {
	return &recordingFlagCache{flags: map[string]*rulesengine.Flag{}}
}

func (c *recordingFlagCache) Get(_ context.Context, key string) (*rulesengine.Flag, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	flag, ok := c.flags[key]
	return flag, ok
}

func (c *recordingFlagCache) Set(_ context.Context, key string, val *rulesengine.Flag, _ *time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.flags[key] = val
	return nil
}

func (c *recordingFlagCache) Delete(_ context.Context, key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.flags, key)
	return nil
}

func (c *recordingFlagCache) DeleteMissing(_ context.Context, keysToKeep []string, _ string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.deleteMissingCalls = append(c.deleteMissingCalls, keysToKeep)
	for key := range c.flags {
		if !containsKey(keysToKeep, key) {
			delete(c.flags, key)
		}
	}
}

func containsKey(keys []string, key string) bool {
	for _, k := range keys {
		if k == key {
			return true
		}
	}
	return false
}

func (c *recordingFlagCache) cachedKeys() []string {
	c.mu.Lock()
	defer c.mu.Unlock()

	keys := make([]string, 0, len(c.flags))
	for key := range c.flags {
		keys = append(keys, key)
	}
	return keys
}

func (c *recordingFlagCache) deleteMissingCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return len(c.deleteMissingCalls)
}

// flagsPageMessage renders one message of a flags snapshot. A nil pagination
// makes it the single-message snapshot a server without pagination sends.
func flagsPageMessage(t *testing.T, keys []string, pagination *schematicdatastreamws.DataStreamPagination) *schematicdatastreamws.DataStreamResp {
	t.Helper()

	flags := make([]*rulesengine.Flag, 0, len(keys))
	for _, key := range keys {
		flags = append(flags, &rulesengine.Flag{
			ID:            "flag_" + key,
			AccountID:     "account_XXXX",
			EnvironmentID: "env_XXXX",
			Key:           key,
			DefaultValue:  true,
		})
	}

	data, err := json.Marshal(flags)
	require.NoError(t, err)

	return &schematicdatastreamws.DataStreamResp{
		EntityType:  string(schematicdatastreamws.EntityTypeFlags),
		Data:        data,
		Pagination:  pagination,
		MessageType: schematicdatastreamws.MessageTypeFull,
	}
}

func page(hasMore bool, number, total int) *schematicdatastreamws.DataStreamPagination {
	return &schematicdatastreamws.DataStreamPagination{HasMore: hasMore, Page: number, Total: total}
}

// quietLogger keeps a cache warning out of the test output.
type quietLogger struct{}

func (quietLogger) Debug(context.Context, string, ...any) {}
func (quietLogger) Info(context.Context, string, ...any)  {}
func (quietLogger) Error(context.Context, string, ...any) {}
func (quietLogger) Warn(context.Context, string, ...any)  {}

// newFlagCacheClient is a client with nothing wired but the flags cache, which
// is all handleFlagsMessage touches. Replicator mode supplies the cache version
// from a field rather than the rules engine, so the client needs no engine.
func newFlagCacheClient(cache *recordingFlagCache) *DataStreamClient {
	return &DataStreamClient{
		flagsCacheProvider:     cache,
		logger:                 quietLogger{},
		replicatorMode:         true,
		replicatorCacheVersion: "testversion",
	}
}

func TestHandleFlagsMessage_PagesApplyAsOneSnapshot(t *testing.T) {
	cache := newRecordingFlagCache()
	client := newFlagCacheClient(cache)
	ctx := context.Background()

	require.NoError(t, client.handleFlagsMessage(ctx, flagsPageMessage(t, []string{"a", "b"}, page(true, 1, 5))))
	require.NoError(t, client.handleFlagsMessage(ctx, flagsPageMessage(t, []string{"c", "d"}, page(true, 2, 5))))

	assert.Zero(t, cache.deleteMissingCount(), "a mid-snapshot page must not delete the pages still to come")

	require.NoError(t, client.handleFlagsMessage(ctx, flagsPageMessage(t, []string{"e"}, page(false, 3, 5))))

	assert.Equal(t, 1, cache.deleteMissingCount(), "the delete runs once, on the last page")
	assert.ElementsMatch(t,
		[]string{
			client.flagCacheKey("a"), client.flagCacheKey("b"), client.flagCacheKey("c"),
			client.flagCacheKey("d"), client.flagCacheKey("e"),
		},
		cache.cachedKeys(),
		"every page's flags survive the snapshot",
	)
}

// A flag the server stops sending has to go, so the delete still has to happen
// -- against the whole snapshot's keys rather than the last page's.
func TestHandleFlagsMessage_FlagsAbsentFromTheSnapshotAreDropped(t *testing.T) {
	cache := newRecordingFlagCache()
	client := newFlagCacheClient(cache)
	ctx := context.Background()

	require.NoError(t, client.handleFlagsMessage(ctx, flagsPageMessage(t, []string{"keep", "retire"}, page(false, 1, 2))))
	require.Len(t, cache.cachedKeys(), 2)

	// A later snapshot, as a reconnect delivers, no longer carries "retire".
	require.NoError(t, client.handleFlagsMessage(ctx, flagsPageMessage(t, []string{"keep"}, page(false, 1, 1))))

	assert.Equal(t, []string{client.flagCacheKey("keep")}, cache.cachedKeys())
}

// An interrupted snapshot must not leave its keys behind: the next snapshot's
// first page starts the set over, so a flag dropped in between still goes.
func TestHandleFlagsMessage_InterruptedSnapshotDoesNotLeakKeys(t *testing.T) {
	cache := newRecordingFlagCache()
	client := newFlagCacheClient(cache)
	ctx := context.Background()

	// First page of a snapshot that never finishes.
	require.NoError(t, client.handleFlagsMessage(ctx, flagsPageMessage(t, []string{"stale"}, page(true, 1, 2))))

	// A fresh snapshot, without that flag, completes in one page.
	require.NoError(t, client.handleFlagsMessage(ctx, flagsPageMessage(t, []string{"current"}, page(false, 1, 1))))

	assert.Equal(t, []string{client.flagCacheKey("current")}, cache.cachedKeys(),
		"the abandoned page's flag was carried into the next snapshot")
}

// A server without pagination sends one message and no pagination field, which
// must still be applied as the whole set.
func TestHandleFlagsMessage_UnpaginatedSnapshotAppliesWhole(t *testing.T) {
	cache := newRecordingFlagCache()
	client := newFlagCacheClient(cache)
	ctx := context.Background()

	require.NoError(t, client.handleFlagsMessage(ctx, flagsPageMessage(t, []string{"a", "b"}, nil)))

	assert.Equal(t, 1, cache.deleteMissingCount())
	assert.ElementsMatch(t, []string{client.flagCacheKey("a"), client.flagCacheKey("b")}, cache.cachedKeys())
}
