package leases

import (
	"context"
	"errors"
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
		_, _, ok, err := b.leases.TryReserve(ctx, "co_1", "ct_1", credits)
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
			_, _, ok, err := b.leases.TryReserve(ctx, "co_1", "ct_1", 1)
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

	_, _, _, err := b.leases.TryReserve(ctx, "co_1", "ct_1", 100)
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

	_, _, _, err := b.leases.TryReserve(ctx, "co_1", "ct_1", 100)
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

	balance, _, ok, err := b.leases.TryReserve(ctx, "co_1", "ct_1", 2.5)
	require.NoError(t, err)
	require.True(t, ok)
	assert.Equal(t, 7.5, balance)

	raw, err := b.client.HGet(ctx, DefaultKeyPrefix+leaseKeyNamespace+LeaseKey("co_1", "ct_1"), "localRemainingCredits").Result()
	require.NoError(t, err)
	assert.False(t, strings.ContainsAny(raw, "eE"), "amounts are written without exponent notation")
	assert.Equal(t, "7.5", raw)
}

// The hash, its expiry and both indexes go out as one transaction. Issued
// separately, an expiry that failed on its own left a hash with no TTL and no
// index entry, which nothing would ever reap.
func TestRedisReservationStoreAddLandsAsOneTransaction(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	b := newRedisBackend(t)
	recorder := &commandRecorder{}
	b.client.AddHook(recorder)

	require.NoError(t, b.reservations.Add(ctx, newReservation("res_1", "lse_1", 100, 60_000, b.clock)))

	assert.Equal(t,
		[]string{"multi", "hset", "pexpireat", "zadd", "hset", "exec"},
		recorder.lastBatch(),
	)
	assert.Positive(t, b.server.TTL(DefaultKeyPrefix+reservationKeyNamespace+"res_1"),
		"the hash carries the expiry that reaps it")
}

func TestRedisReservationStoreAddLeavesNothingBehindWhenTheWriteFails(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	b := newRedisBackend(t)
	recorder := &commandRecorder{}
	b.client.AddHook(recorder)
	recorder.failTransactions(errors.New("redis went away mid-transaction"))

	err := b.reservations.Add(ctx, newReservation("res_1", "lse_1", 100, 60_000, b.clock))

	require.Error(t, err)
	assert.False(t, b.server.Exists(DefaultKeyPrefix+reservationKeyNamespace+"res_1"),
		"a hash with no expiry would never be reaped")
	count, err := b.reservations.Count(ctx)
	require.NoError(t, err)
	assert.Zero(t, count)
	total, err := b.reservations.ReservedCredits(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Zero(t, total)
}

// The expiry index is how the sweeper reaches a per-tenant field whose hash is
// gone, so the field goes first: dropping the index entry and then failing here
// would inflate ReservedCredits forever.
func TestRedisReservationStoreConsumeDropsThePerTenantFieldBeforeTheIndex(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	b := newRedisBackend(t)
	installLease(t, b.leases, b.clock, "lse_1", "co_1", "ct_1", 1000, 3_600_000)
	_, _, _, err := b.leases.TryReserve(ctx, "co_1", "ct_1", 100)
	require.NoError(t, err)
	require.NoError(t, b.reservations.Add(ctx, newReservation("res_1", "lse_1", 100, 60_000, b.clock)))

	recorder := &commandRecorder{}
	b.client.AddHook(recorder)
	_, claimed, err := b.reservations.Consume(ctx, "res_1", 30)
	require.NoError(t, err)
	require.True(t, claimed)

	commands := recorder.order()
	hdel, zrem := indexOf(commands, "hdel"), indexOf(commands, "zrem")
	require.NotEqual(t, -1, hdel, commands)
	require.NotEqual(t, -1, zrem, commands)
	assert.Less(t, hdel, zrem, "the per-tenant field is dropped first: %v", commands)
}

// A verdict that flips to denied after the credits are debited has to hand the
// whole hold back, on the shared store as well as in memory.
func TestCancelReservationRefundsTheWholeHoldOnRedis(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	b := newRedisBackend(t)
	installLease(t, b.leases, b.clock, "lse_1", "co_1", "ct_1", 1000, 3_600_000)
	_, _, _, err := b.leases.TryReserve(ctx, "co_1", "ct_1", 100)
	require.NoError(t, err)
	record := newReservation("res_1", "lse_1", 100, 60_000, b.clock)
	require.NoError(t, b.reservations.Add(ctx, record))

	cancelReservation(ctx, CheckDeps{Leases: b.leases, Reservations: b.reservations}, record)

	entry, err := b.leases.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 1000.0, entry.LocalRemainingCredits)
	total, err := b.reservations.ReservedCredits(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Zero(t, total)
	count, err := b.reservations.Count(ctx)
	require.NoError(t, err)
	assert.Zero(t, count)
}

// A debit whose record never landed has nothing to claim, so the refund goes
// straight to the lease rather than stranding the slice until expiry.
func TestUndoDebitRefundsADebitWithNoRecordOnRedis(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	b := newRedisBackend(t)
	installLease(t, b.leases, b.clock, "lse_1", "co_1", "ct_1", 1000, 3_600_000)
	_, _, _, err := b.leases.TryReserve(ctx, "co_1", "ct_1", 100)
	require.NoError(t, err)

	undoDebit(ctx, CheckDeps{Leases: b.leases, Reservations: b.reservations}, newReservation("res_1", "lse_1", 100, 60_000, b.clock))

	entry, err := b.leases.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 1000.0, entry.LocalRemainingCredits)
}

// An expired lease row outlives its expiry by the grace window, so a sweeper
// still has somewhere to refund expired holds; reserving against it is refused
// from the moment it expires.
func TestRedisLeaseStoreKeepsAnExpiredLeaseForTheGraceWindow(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	b := newRedisBackend(t)
	installLease(t, b.leases, b.clock, "lse_1", "co_1", "ct_1", 1000, 60_000)

	elapse(b, 61*time.Second)

	_, _, reserved, err := b.leases.TryReserve(ctx, "co_1", "ct_1", 1)
	require.NoError(t, err)
	assert.False(t, reserved, "an expired lease gates nothing, grace window or not")
	entry, err := b.leases.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	require.NotNil(t, entry, "the row survives the grace window")

	elapse(b, time.Duration(leaseTTLGraceMs)*time.Millisecond)

	entry, err = b.leases.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Nil(t, entry, "past the grace window Redis has evicted the row")
}

// Past its own grace window the reservation hash is evicted, and the sweeper
// meets the orphaned index entry it was built for: it reconciles the per-tenant
// field so ReservedCredits stops summing an evicted hold, and deliberately
// refunds nothing, since exactly-once cannot be arbitrated without the hash.
func TestRedisReservationStoreSweepAfterTheGraceWindowReconcilesWithoutRefunding(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	b := newRedisBackend(t)
	installLease(t, b.leases, b.clock, "lse_1", "co_1", "ct_1", 1000, 3_600_000)
	_, _, _, err := b.leases.TryReserve(ctx, "co_1", "ct_1", 100)
	require.NoError(t, err)
	require.NoError(t, b.reservations.Add(ctx, newReservation("res_1", "lse_1", 100, 10_000, b.clock)))

	// Expired, but inside the grace window, so the sweeper can still refund.
	elapse(b, 11*time.Second)
	require.True(t, b.server.Exists(DefaultKeyPrefix+reservationKeyNamespace+"res_1"))

	elapse(b, time.Duration(reservationTTLGraceMs)*time.Millisecond)
	require.False(t, b.server.Exists(DefaultKeyPrefix+reservationKeyNamespace+"res_1"))

	swept, err := b.reservations.SweepExpired(ctx)
	require.NoError(t, err)
	assert.Zero(t, swept)

	total, err := b.reservations.ReservedCredits(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Zero(t, total, "the orphaned per-tenant field is reconciled away")
	entry, err := b.leases.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 900.0, entry.LocalRemainingCredits, "the slice waits for the lease to expire server-side")
	count, err := b.reservations.Count(ctx)
	require.NoError(t, err)
	assert.Zero(t, count)
}
