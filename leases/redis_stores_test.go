package leases

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Redis Cluster rejects a multi-key script whose keys span slots, so every
// script here must address exactly one key.
func TestLuaScriptsAreSingleKey(t *testing.T) {
	t.Parallel()
	for name, source := range map[string]string{
		"replace":     replaceScriptSource,
		"try_reserve": tryReserveScriptSource,
		"refund":      refundScriptSource,
		"extend":      extendScriptSource,
		"claim":       claimScriptSource,
	} {
		assert.Containsf(t, source, "KEYS[1]", "%s script should address a key", name)
		for index := 2; index <= 4; index++ {
			assert.NotContainsf(t, source, fmt.Sprintf("KEYS[%d]", index), "%s script must stay single-key", name)
		}
	}
}

func TestRedisLeaseStoreRejectsNonFiniteAndNegativeDebits(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	b := newRedisBackend(t)
	installLease(t, b.leases, b.clock, "lse_1", "co_1", "ct_1", 100, 60_000)

	// The string form of NaN parses back to a Lua nan, slips through the `<`
	// comparison, and would poison the shared balance for every pod.
	for _, credits := range []float64{math.NaN(), math.Inf(1), -1} {
		_, ok, err := b.leases.TryReserve(ctx, "co_1", "ct_1", credits)
		require.NoError(t, err)
		assert.False(t, ok)
	}
	entry, err := b.leases.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 100.0, entry.LocalRemainingCredits)
}

func TestRedisLeaseStoreGetReturnsNilForAnEmptySlot(t *testing.T) {
	t.Parallel()
	b := newRedisBackend(t)
	entry, err := b.leases.Get(context.Background(), "co_missing", "ct_1")
	require.NoError(t, err)
	assert.Nil(t, entry)
}

func TestRedisLeaseStoreExtendFallsBackToTheDefaultDuration(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	b := newRedisBackend(t)
	store := NewRedisLeaseStore(b.client, RedisLeaseStoreOptions{
		DefaultLeaseDuration: 10 * time.Minute,
		Clock:                b.clock.Now,
	})
	installLease(t, store, b.clock, "lse_1", "co_1", "ct_1", 100, 60_000)

	// No expiry from the caller: only a direct caller reaches this, since the
	// manager always passes one.
	require.NoError(t, store.Extend(ctx, "co_1", "ct_1", 200, nil, ""))

	entry, err := store.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 200.0, entry.GrantedAmount)
	assert.Equal(t, b.clock.Now().Add(10*time.Minute), entry.ExpiresAt)
}

// Two pods extending the same lease concurrently must converge on the server
// total instead of each applying a delta from the same stale read.
func TestRedisLeaseStoreConcurrentExtendsConvergeOnTheTotal(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	b := newRedisBackend(t)
	installLease(t, b.leases, b.clock, "lse_1", "co_1", "ct_1", 1000, 3_600_000)
	expiry := b.clock.at(600_000)

	var wg sync.WaitGroup
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			assert.NoError(t, b.leases.Extend(ctx, "co_1", "ct_1", 2000, &expiry, "lse_1"))
		}()
	}
	wg.Wait()

	entry, err := b.leases.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 2000.0, entry.GrantedAmount)
	assert.Equal(t, 2000.0, entry.LocalRemainingCredits)
}

func TestRedisLeaseStoreConcurrentTryReserveIsAtomic(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	b := newRedisBackend(t)
	installLease(t, b.leases, b.clock, "lse_1", "co_1", "ct_1", 50, 3_600_000)

	var mu sync.Mutex
	granted := 0
	var wg sync.WaitGroup
	for i := 0; i < 120; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, ok, err := b.leases.TryReserve(ctx, "co_1", "ct_1", 1)
			assert.NoError(t, err)
			if ok {
				mu.Lock()
				granted++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()

	assert.Equal(t, 50, granted)
	entry, err := b.leases.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 0.0, entry.LocalRemainingCredits)
}

func TestRedisReservationStoreRoundTripsTheEvaluationContext(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	b := newRedisBackend(t)
	installLease(t, b.leases, b.clock, "lse_1", "co_1", "ct_1", 1000, 3_600_000)

	reservation := newReservation("res_1", "lse_1", 100, 60_000, b.clock)
	reservation.User = map[string]string{"id": "user_1"}
	require.NoError(t, b.reservations.Add(ctx, reservation))

	stored, err := b.reservations.Get(ctx, "res_1")
	require.NoError(t, err)
	require.NotNil(t, stored)
	assert.Equal(t, map[string]string{"id": "co_1"}, stored.Company)
	assert.Equal(t, map[string]string{"id": "user_1"}, stored.User)
	assert.Equal(t, reservation.ExpiresAt, stored.ExpiresAt)

	count, err := b.reservations.Count(ctx)
	require.NoError(t, err)
	assert.Equal(t, 1, count)
}

// A hash that TTL-evicted before the sweeper reached it leaves an orphaned
// per-tenant index field. The sweeper reconciles it so ReservedCredits stops
// summing an evicted hold, and deliberately refunds nothing: without the hash,
// exactly-once cannot be arbitrated across racing sweepers.
func TestRedisReservationStoreSweepReconcilesAnEvictedHashWithoutRefunding(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	b := newRedisBackend(t)
	installLease(t, b.leases, b.clock, "lse_1", "co_1", "ct_1", 1000, 3_600_000)

	_, _, err := b.leases.TryReserve(ctx, "co_1", "ct_1", 100)
	require.NoError(t, err)
	require.NoError(t, b.reservations.Add(ctx, newReservation("res_1", "lse_1", 100, 10_000, b.clock)))
	b.server.Del(DefaultKeyPrefix + reservationKeyNamespace + "res_1")

	b.clock.advance(11 * time.Second)
	swept, err := b.reservations.SweepExpired(ctx)
	require.NoError(t, err)
	assert.Zero(t, swept)

	total, err := b.reservations.ReservedCredits(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Zero(t, total)
	entry, err := b.leases.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 900.0, entry.LocalRemainingCredits)
}

func TestRedisReservationStoreSweepDropsAnUnparseableIndexMember(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	b := newRedisBackend(t)
	indexKey := DefaultKeyPrefix + reservationIndexKey
	require.NoError(t, b.client.ZAdd(ctx, indexKey, redis.Z{Score: 0, Member: "not-a-member"}).Err())

	b.clock.advance(time.Second)
	swept, err := b.reservations.SweepExpired(ctx)
	require.NoError(t, err)
	assert.Zero(t, swept)

	count, err := b.client.ZCard(ctx, indexKey).Result()
	require.NoError(t, err)
	assert.Zero(t, count)
}

// Paging keeps one tick's reply bounded; the per-member ZREM is what advances
// the sweeper through a backlog larger than a single batch.
func TestRedisReservationStoreSweepPagesThroughABacklog(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	b := newRedisBackend(t)
	installLease(t, b.leases, b.clock, "lse_1", "co_1", "ct_1", 1000, 3_600_000)

	const backlog = sweepBatchSize + 17
	for i := 0; i < backlog; i++ {
		require.NoError(t, b.reservations.Add(ctx, newReservation("res_"+strconv.Itoa(i), "lse_1", 1, 10_000, b.clock)))
	}

	b.clock.advance(11 * time.Second)
	swept, err := b.reservations.SweepExpired(ctx)
	require.NoError(t, err)
	assert.Equal(t, backlog, swept)

	count, err := b.reservations.Count(ctx)
	require.NoError(t, err)
	assert.Zero(t, count)
}

func TestRedisReservationStoreDoubleConsumeClaimsOnce(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	b := newRedisBackend(t)
	installLease(t, b.leases, b.clock, "lse_1", "co_1", "ct_1", 1000, 3_600_000)

	_, _, err := b.leases.TryReserve(ctx, "co_1", "ct_1", 100)
	require.NoError(t, err)
	require.NoError(t, b.reservations.Add(ctx, newReservation("res_1", "lse_1", 100, 60_000, b.clock)))

	consumed, claimed, err := b.reservations.Consume(ctx, "res_1", 30)
	require.NoError(t, err)
	assert.True(t, claimed)
	assert.Equal(t, 30.0, consumed)

	_, claimed, err = b.reservations.Consume(ctx, "res_1", 30)
	require.NoError(t, err)
	assert.False(t, claimed)

	entry, err := b.leases.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 970.0, entry.LocalRemainingCredits)
}

// Fractional amounts must survive the round trip: a Lua number reply truncates
// to an integer, which is why balances are stored and returned as strings.
func TestRedisLeaseStoreKeepsFractionalAmountsExact(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	b := newRedisBackend(t)
	installLease(t, b.leases, b.clock, "lse_1", "co_1", "ct_1", 10, 3_600_000)

	balance, ok, err := b.leases.TryReserve(ctx, "co_1", "ct_1", 2.5)
	require.NoError(t, err)
	require.True(t, ok)
	assert.Equal(t, 7.5, balance)

	raw, err := b.client.HGet(ctx, DefaultKeyPrefix+leaseKeyNamespace+LeaseKey("co_1", "ct_1"), "localRemainingCredits").Result()
	require.NoError(t, err)
	assert.False(t, strings.ContainsAny(raw, "eE"), "amounts are written without exponent notation")
	assert.Equal(t, "7.5", raw)
}
