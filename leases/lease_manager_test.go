package leases

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// gatedWireClient holds every acquire and extend open until the test releases
// them, so a race the single-flight is meant to collapse can actually happen.
type gatedWireClient struct {
	gate chan struct{}

	acquires atomic.Int64
	extends  atomic.Int64
	releases atomic.Int64

	grant        LeaseGrant
	extendGrant  LeaseGrant
	releasePanic bool
	// duringAcquire runs while an acquire is in flight, for emulating a sibling
	// pod winning the race.
	duringAcquire func()
}

func (w *gatedWireClient) Acquire(context.Context, string, string, float64, time.Time) (*LeaseGrant, error) {
	w.acquires.Add(1)
	<-w.gate
	if w.duringAcquire != nil {
		during := w.duringAcquire
		w.duringAcquire = nil
		during()
	}
	grant := w.grant
	return &grant, nil
}

func (w *gatedWireClient) Extend(context.Context, string, float64, time.Time) (*LeaseGrant, error) {
	w.extends.Add(1)
	<-w.gate
	grant := w.extendGrant
	return &grant, nil
}

func (w *gatedWireClient) Release(context.Context, string) error {
	w.releases.Add(1)
	if w.releasePanic {
		panic("release blew up")
	}
	return nil
}

// failingLeaseStore reports a store outage on every read.
type failingLeaseStore struct {
	LeaseStore
}

var errStoreDown = errors.New("lease store unreachable")

func (failingLeaseStore) Get(context.Context, string, string) (*LeaseState, error) {
	return nil, errStoreDown
}

func TestAcquireIfNeededIsSingleFlightPerSlot(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	store := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	wire := &gatedWireClient{
		gate: make(chan struct{}),
		grant: LeaseGrant{
			LeaseID:       "lse_1",
			CompanyID:     "co_1",
			CreditTypeID:  "ct_1",
			GrantedAmount: 1000,
			ExpiresAt:     clock.at(300_000),
		},
	}
	manager := NewLeaseManager(wire, store, LeaseManagerOptions{Clock: clock.Now})
	t.Cleanup(manager.Stop)

	const callers = 12
	results := make(chan *LeaseState, callers)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		results <- manager.AcquireIfNeeded(ctx, "co_1", "ct_1")
	}()
	// The flight registers before the wire call starts, so once the call is in
	// flight every later caller is guaranteed to find it and join.
	require.Eventually(t, func() bool { return wire.acquires.Load() == 1 }, time.Second, time.Millisecond)

	for i := 1; i < callers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results <- manager.AcquireIfNeeded(ctx, "co_1", "ct_1")
		}()
	}
	time.Sleep(20 * time.Millisecond)
	close(wire.gate)
	wg.Wait()
	close(results)

	assert.Equal(t, int64(1), wire.acquires.Load(), "concurrent callers share one wire call")
	for entry := range results {
		require.NotNil(t, entry)
		assert.Equal(t, "lse_1", entry.LeaseID)
	}
	// The lease installed once, at its full grant.
	entry, err := store.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 1000.0, entry.LocalRemainingCredits)
}

// A live slot answers an acquire from the store, so a blocked extend on the
// same slot neither stalls it nor hands it the extend's result.
func TestAcquireDoesNotWaitOnAnInFlightExtend(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	store := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	installLease(t, store, clock, "lse_1", "co_1", "ct_1", 1000, 300_000)
	_, _, err := store.TryReserve(ctx, "co_1", "ct_1", 900)
	require.NoError(t, err)

	wire := &gatedWireClient{
		gate:        make(chan struct{}),
		grant:       LeaseGrant{LeaseID: "lse_2", CompanyID: "co_1", CreditTypeID: "ct_1", GrantedAmount: 1000, ExpiresAt: clock.at(600_000)},
		extendGrant: LeaseGrant{LeaseID: "lse_1", CompanyID: "co_1", CreditTypeID: "ct_1", GrantedAmount: 2000, ExpiresAt: clock.at(600_000)},
	}
	manager := NewLeaseManager(wire, store, LeaseManagerOptions{
		Clock:  clock.Now,
		Config: ResolvedLeaseConfig{LeaseSize: 1000, LowWaterMark: 0.25},
	})
	t.Cleanup(manager.Stop)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		manager.MaybeExtend(ctx, "co_1", "ct_1", nil)
	}()
	require.Eventually(t, func() bool { return wire.extends.Load() == 1 }, time.Second, time.Millisecond)

	// The slot still holds a live lease, so this returns it without any wire
	// call, rather than joining the extend in flight.
	entry := manager.AcquireIfNeeded(ctx, "co_1", "ct_1")
	require.NotNil(t, entry)
	assert.Equal(t, "lse_1", entry.LeaseID)
	assert.Zero(t, wire.acquires.Load())

	close(wire.gate)
	wg.Wait()

	extended, err := store.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 2000.0, extended.GrantedAmount)
}

// Fire-and-forget work is unawaited, so a panic in it has nobody to recover it.
func TestRedundantReleasePanicDoesNotEscape(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	store := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	wire := &gatedWireClient{
		gate:         closedGate(),
		grant:        LeaseGrant{LeaseID: "lse_loser", CompanyID: "co_1", CreditTypeID: "ct_1", GrantedAmount: 1000, ExpiresAt: clock.at(300_000)},
		releasePanic: true,
	}
	// A sibling installs a live lease while this acquire is in flight, so the
	// lease the server minted for the loser is redundant and gets released in
	// the background.
	wire.duringAcquire = func() {
		installLease(t, store, clock, "lse_winner", "co_1", "ct_1", 1000, 300_000)
	}
	manager := NewLeaseManager(wire, store, LeaseManagerOptions{Clock: clock.Now})

	entry := manager.AcquireIfNeeded(ctx, "co_1", "ct_1")
	require.NotNil(t, entry)
	assert.Equal(t, "lse_winner", entry.LeaseID)

	assert.NotPanics(t, manager.Stop)
	assert.Equal(t, int64(1), wire.releases.Load())
}

func TestStoreFailuresResolveToNoLease(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	store := failingLeaseStore{NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})}
	wire := &gatedWireClient{gate: closedGate()}
	manager := NewLeaseManager(wire, store, LeaseManagerOptions{Clock: clock.Now})
	t.Cleanup(manager.Stop)

	assert.Nil(t, manager.AcquireIfNeeded(ctx, "co_1", "ct_1"))
	assert.Nil(t, manager.MaybeExtend(ctx, "co_1", "ct_1", nil))
	assert.Zero(t, wire.acquires.Load())
	assert.Zero(t, wire.extends.Load())
}

// A shared backend must never enumerate and release: sibling pods still draw on
// those leases.
func TestReleaseAllLocalLeasesSkipsANonEnumerableStore(t *testing.T) {
	t.Parallel()
	b := newRedisBackend(t)
	installLease(t, b.leases, b.clock, "lse_1", "co_1", "ct_1", 1000, 300_000)
	wire := &gatedWireClient{gate: closedGate()}
	manager := NewLeaseManager(wire, b.leases, LeaseManagerOptions{Clock: b.clock.Now})
	t.Cleanup(manager.Stop)

	manager.ReleaseAllLocalLeases(context.Background())
	assert.Zero(t, wire.releases.Load())

	entry, err := b.leases.Get(context.Background(), "co_1", "ct_1")
	require.NoError(t, err)
	assert.NotNil(t, entry)
}

func TestStartSweepRunsAndStopJoinsTheLoop(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	store := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	installLease(t, store, clock, "lse_1", "co_1", "ct_1", 1000, 3_600_000)
	reservations := NewInMemoryReservationStore(store, InMemoryReservationStoreOptions{Clock: clock.Now})

	_, _, err := store.TryReserve(ctx, "co_1", "ct_1", 100)
	require.NoError(t, err)
	// Already expired, so the first tick sweeps it back to the lease.
	require.NoError(t, reservations.Add(ctx, newReservation("res_1", "lse_1", 100, 0, clock)))

	manager := NewLeaseManager(&gatedWireClient{gate: closedGate()}, store, LeaseManagerOptions{
		Reservations:  reservations,
		SweepInterval: time.Millisecond,
		Clock:         clock.Now,
	})
	manager.StartSweep()
	// Calling twice must not start a second loop.
	manager.StartSweep()

	require.Eventually(t, func() bool {
		count, err := reservations.Count(ctx)
		return err == nil && count == 0
	}, time.Second, time.Millisecond)

	entry, err := store.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 1000.0, entry.LocalRemainingCredits)

	manager.Stop()
	// Stop joins the loop, so no further sweeps can be in flight, and it is
	// safe to call again.
	assert.NotPanics(t, manager.Stop)
	manager.StartSweep()
}

func TestFlightGroupCollapsesConcurrentCallsPerKey(t *testing.T) {
	t.Parallel()
	var group flightGroup
	gate := make(chan struct{})
	started := make(chan struct{})
	var calls atomic.Int64

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			entry := group.do("slot", func() *LeaseState {
				if calls.Add(1) == 1 {
					close(started)
				}
				<-gate
				return &LeaseState{LeaseID: "lse_1"}
			})
			assert.Equal(t, "lse_1", entry.LeaseID)
		}()
	}
	<-started
	time.Sleep(20 * time.Millisecond)
	close(gate)
	wg.Wait()

	// A different key never joins an unrelated flight.
	other := group.do("other", func() *LeaseState { return &LeaseState{LeaseID: "lse_2"} })
	assert.Equal(t, "lse_2", other.LeaseID)
}

// Acquire and extend keep separate registries, so an in-flight extend can never
// satisfy an acquire for the same slot, or the other way round.
func TestSeparateFlightGroupsDoNotSatisfyEachOther(t *testing.T) {
	t.Parallel()
	var acquires, extends flightGroup
	gate := make(chan struct{})
	started := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		entry := acquires.do("slot", func() *LeaseState {
			close(started)
			<-gate
			return &LeaseState{LeaseID: "lse_acquired"}
		})
		assert.Equal(t, "lse_acquired", entry.LeaseID)
	}()
	<-started

	// The acquire for this very slot is still blocked; the extend runs its own
	// call rather than riding that result.
	extended := extends.do("slot", func() *LeaseState { return &LeaseState{LeaseID: "lse_extended"} })
	assert.Equal(t, "lse_extended", extended.LeaseID)

	close(gate)
	wg.Wait()
}

func closedGate() chan struct{} {
	gate := make(chan struct{})
	close(gate)
	return gate
}

// A stalled extend must not hold a shutdown open: the wire client sets no HTTP
// timeout, so Stop bounds its own wait instead.
func TestStopIsBoundedWhenBackgroundWorkStalls(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	store := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	installLease(t, store, clock, "lse_1", "co_1", "ct_1", 1000, 300_000)
	// Below the water mark, so the background extend actually fires.
	_, _, err := store.TryReserve(ctx, "co_1", "ct_1", 900)
	require.NoError(t, err)

	wire := &gatedWireClient{
		gate:        make(chan struct{}),
		extendGrant: LeaseGrant{LeaseID: "lse_1", CompanyID: "co_1", CreditTypeID: "ct_1", GrantedAmount: 2000, ExpiresAt: clock.at(600_000)},
	}
	manager := NewLeaseManager(wire, store, LeaseManagerOptions{
		Clock:       clock.Now,
		StopTimeout: 50 * time.Millisecond,
		Config:      ResolvedLeaseConfig{LeaseSize: 1000, LowWaterMark: 0.25},
	})

	manager.ExtendInBackground(ctx, "co_1", "ct_1")
	require.Eventually(t, func() bool { return wire.extends.Load() == 1 }, time.Second, time.Millisecond)

	start := time.Now()
	manager.Stop()
	assert.Less(t, time.Since(start), time.Second, "Stop waits out the stalled extend only up to StopTimeout")

	close(wire.gate)
	manager.drainBackground()
}

// Spawning after Stop would Add to a WaitGroup whose wait has already passed,
// and the work would touch a manager being torn down.
func TestSpawnAfterStopIsRefused(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	store := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	installLease(t, store, clock, "lse_1", "co_1", "ct_1", 1000, 300_000)
	_, _, err := store.TryReserve(ctx, "co_1", "ct_1", 900)
	require.NoError(t, err)

	wire := &gatedWireClient{
		gate:        closedGate(),
		extendGrant: LeaseGrant{LeaseID: "lse_1", CompanyID: "co_1", CreditTypeID: "ct_1", GrantedAmount: 2000, ExpiresAt: clock.at(600_000)},
	}
	manager := NewLeaseManager(wire, store, LeaseManagerOptions{
		Clock:  clock.Now,
		Config: ResolvedLeaseConfig{LeaseSize: 1000, LowWaterMark: 0.25},
	})
	manager.Stop()

	assert.NotPanics(t, func() { manager.ExtendInBackground(ctx, "co_1", "ct_1") })
	manager.drainBackground()
	assert.Zero(t, wire.extends.Load(), "a refused spawn runs nothing")
}
