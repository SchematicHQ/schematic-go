package leases

import (
	"context"
	"math"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The vectors pin the arithmetic of every store operation on both backends.
// These tests cover what JSON vectors cannot express: non-finite amounts and
// concurrency.

func TestInMemoryLeaseStoreRejectsNonFiniteAndNegativeDebits(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	store := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	installLease(t, store, clock, "lse_1", "co_1", "ct_1", 100, 60_000)

	for _, credits := range []float64{math.NaN(), math.Inf(1), math.Inf(-1), -1} {
		balance, _, ok, err := store.TryReserve(ctx, "co_1", "ct_1", credits)
		require.NoError(t, err)
		assert.False(t, ok)
		assert.Zero(t, balance)
	}

	// A NaN that reached the arithmetic would poison the balance into approving
	// everything, because every comparison against NaN is false.
	entry, err := store.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	require.NotNil(t, entry)
	assert.Equal(t, 100.0, entry.LocalRemainingCredits)

	_, _, ok, err := store.TryReserve(ctx, "co_1", "ct_1", 101)
	require.NoError(t, err)
	assert.False(t, ok)
}

func TestInMemoryLeaseStoreRefundIgnoresNonFiniteAmounts(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	store := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	installLease(t, store, clock, "lse_1", "co_1", "ct_1", 100, 60_000)
	_, _, _, err := store.TryReserve(ctx, "co_1", "ct_1", 40)
	require.NoError(t, err)

	require.NoError(t, store.Refund(ctx, "co_1", "ct_1", math.NaN(), ""))
	require.NoError(t, store.Refund(ctx, "co_1", "ct_1", math.Inf(1), ""))

	entry, err := store.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	require.NotNil(t, entry)
	assert.Equal(t, 60.0, entry.LocalRemainingCredits)
}

// Per-slot atomicity: concurrent check-and-debits must never oversubscribe the
// grant, and the successes must sum to exactly what was granted.
func TestInMemoryLeaseStoreConcurrentTryReserveIsAtomic(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	store := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	installLease(t, store, clock, "lse_1", "co_1", "ct_1", 100, 3_600_000)

	const attempts = 200
	var granted atomic.Int64
	var wg sync.WaitGroup
	for i := 0; i < attempts; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, _, ok, err := store.TryReserve(ctx, "co_1", "ct_1", 1)
			assert.NoError(t, err)
			if ok {
				granted.Add(1)
			}
		}()
	}
	wg.Wait()

	assert.Equal(t, int64(100), granted.Load())
	entry, err := store.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	require.NotNil(t, entry)
	assert.Equal(t, 0.0, entry.LocalRemainingCredits)
}

// Different slots must not serialize against each other, and each must stay
// exact under concurrent mutation.
func TestInMemoryLeaseStoreConcurrentSlotsStayIndependent(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	store := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	installLease(t, store, clock, "lse_1", "co_1", "ct_1", 100, 3_600_000)
	installLease(t, store, clock, "lse_2", "co_2", "ct_1", 100, 3_600_000)

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			_, _, _, err := store.TryReserve(ctx, "co_1", "ct_1", 1)
			assert.NoError(t, err)
		}()
		go func() {
			defer wg.Done()
			_, _, _, err := store.TryReserve(ctx, "co_2", "ct_1", 2)
			assert.NoError(t, err)
		}()
	}
	wg.Wait()

	first, err := store.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 50.0, first.LocalRemainingCredits)
	second, err := store.Get(ctx, "co_2", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 0.0, second.LocalRemainingCredits)
}

func TestInMemoryLeaseStoreListIsASortedSnapshot(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	store := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	installLease(t, store, clock, "lse_b", "co_2", "ct_1", 100, 60_000)
	installLease(t, store, clock, "lse_a", "co_1", "ct_1", 100, 60_000)

	entries, err := store.List(ctx)
	require.NoError(t, err)
	require.Len(t, entries, 2)
	assert.Equal(t, "lse_a", entries[0].LeaseID)
	assert.Equal(t, "lse_b", entries[1].LeaseID)

	// A snapshot, not a handle on the live row.
	entries[0].LocalRemainingCredits = 0
	entry, err := store.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 100.0, entry.LocalRemainingCredits)
}

func TestInMemoryLeaseStoreDropRemovesTheSlot(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	store := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	installLease(t, store, clock, "lse_1", "co_1", "ct_1", 100, 60_000)

	require.NoError(t, store.Drop(ctx, "co_1", "ct_1"))
	entry, err := store.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Nil(t, entry)
}

func TestResolveLeaseConfigPrefersOverrideThenBaseThenDefault(t *testing.T) {
	t.Parallel()
	overrideDuration := 30 * time.Second
	overrideSize := 7.0
	overrides := map[string]LeaseOverride{
		"ct_override": {LeaseDuration: &overrideDuration, LeaseSize: &overrideSize},
	}
	base := ResolvedLeaseConfig{LeaseDuration: time.Minute, LeaseSize: 500}

	resolved := ResolveLeaseConfig(base, overrides, "ct_override")
	assert.Equal(t, overrideDuration, resolved.LeaseDuration)
	assert.Equal(t, overrideSize, resolved.LeaseSize)
	// Untouched by either the override or the base, so the default stands.
	assert.Equal(t, DefaultReservationTTL, resolved.ReservationTTL)
	assert.Equal(t, DefaultLowWaterMark, resolved.LowWaterMark)

	plain := ResolveLeaseConfig(base, overrides, "ct_other")
	assert.Equal(t, time.Minute, plain.LeaseDuration)
	assert.Equal(t, 500.0, plain.LeaseSize)

	empty := ResolveLeaseConfig(ResolvedLeaseConfig{}, nil, "ct_1")
	assert.Equal(t, DefaultLeaseDuration, empty.LeaseDuration)
	assert.Equal(t, DefaultReservationTTL, empty.ReservationTTL)
	assert.Equal(t, DefaultLeaseSize, empty.LeaseSize)
	assert.Equal(t, DefaultLowWaterMark, empty.LowWaterMark)
}
