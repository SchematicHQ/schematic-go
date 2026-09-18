package leases

import (
	"context"
	"errors"
	"strings"
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

func (w *gatedWireClient) Acquire(ctx context.Context, _, _ string, _ float64, _ time.Time) (*LeaseGrant, error) {
	w.acquires.Add(1)
	// Honors the context the way an HTTP call would, so a test can cancel a
	// call in flight.
	select {
	case <-w.gate:
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	if w.duringAcquire != nil {
		during := w.duringAcquire
		w.duringAcquire = nil
		during()
	}
	grant := w.grant
	return &grant, nil
}

func (w *gatedWireClient) Extend(ctx context.Context, _ string, _ float64, _ time.Time) (*LeaseGrant, error) {
	w.extends.Add(1)
	select {
	case <-w.gate:
	case <-ctx.Done():
		return nil, ctx.Err()
	}
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
	_, _, _, err := store.TryReserve(ctx, "co_1", "ct_1", 900)
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

	_, _, _, err := store.TryReserve(ctx, "co_1", "ct_1", 100)
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
	group := flightGroup{spawn: goSpawn}
	gate := make(chan struct{})
	started := make(chan struct{})
	var calls atomic.Int64

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			entry := group.do(t.Context(), "slot", func() *LeaseState {
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
	other := group.do(t.Context(), "other", func() *LeaseState { return &LeaseState{LeaseID: "lse_2"} })
	assert.Equal(t, "lse_2", other.LeaseID)
}

// Acquire and extend keep separate registries, so an in-flight extend can never
// satisfy an acquire for the same slot, or the other way round.
func TestSeparateFlightGroupsDoNotSatisfyEachOther(t *testing.T) {
	t.Parallel()
	acquires, extends := flightGroup{spawn: goSpawn}, flightGroup{spawn: goSpawn}
	gate := make(chan struct{})
	started := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		entry := acquires.do(t.Context(), "slot", func() *LeaseState {
			close(started)
			<-gate
			return &LeaseState{LeaseID: "lse_acquired"}
		})
		assert.Equal(t, "lse_acquired", entry.LeaseID)
	}()
	<-started

	// The acquire for this very slot is still blocked; the extend runs its own
	// call rather than riding that result.
	extended := extends.do(t.Context(), "slot", func() *LeaseState { return &LeaseState{LeaseID: "lse_extended"} })
	assert.Equal(t, "lse_extended", extended.LeaseID)

	close(gate)
	wg.Wait()
}

// goSpawn stands in for the manager's own spawn, which is what tracks a flight
// and recovers a panic in it.
func goSpawn(fn func()) bool {
	go fn()
	return true
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
	_, _, _, err := store.TryReserve(ctx, "co_1", "ct_1", 900)
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
	_, _, _, err := store.TryReserve(ctx, "co_1", "ct_1", 900)
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

// A shared acquire runs on a context of its own: the first caller's
// cancellation or short deadline must not decide what every other caller
// waiting on the slot gets.
func TestAcquireFlightOutlivesTheFirstCallersContext(t *testing.T) {
	t.Parallel()
	clock := newVirtualClock()
	store := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	wire := &gatedWireClient{
		gate:  make(chan struct{}),
		grant: LeaseGrant{LeaseID: "lse_1", CompanyID: "co_1", CreditTypeID: "ct_1", GrantedAmount: 1000, ExpiresAt: clock.at(300_000)},
	}
	manager := NewLeaseManager(wire, store, LeaseManagerOptions{Clock: clock.Now})
	t.Cleanup(manager.Stop)

	firstCtx, cancelFirst := context.WithCancel(context.Background())
	first := make(chan *LeaseState, 1)
	go func() { first <- manager.AcquireIfNeeded(firstCtx, "co_1", "ct_1") }()
	require.Eventually(t, func() bool { return wire.acquires.Load() == 1 }, time.Second, time.Millisecond)

	second := make(chan *LeaseState, 1)
	go func() { second <- manager.AcquireIfNeeded(context.Background(), "co_1", "ct_1") }()
	time.Sleep(20 * time.Millisecond)

	cancelFirst()
	assert.Nil(t, <-first, "a caller that gave up waiting gets nothing")

	close(wire.gate)
	entry := <-second
	require.NotNil(t, entry, "the shared call is not the first caller's to cancel")
	assert.Equal(t, "lse_1", entry.LeaseID)
	assert.Equal(t, int64(1), wire.acquires.Load(), "the flight is still shared")
}

// A check racing a Close must not install a lease after ReleaseAllLocalLeases
// has listed the slots: nobody would be left to release it.
func TestAcquireAfterStopIsRefused(t *testing.T) {
	t.Parallel()
	clock := newVirtualClock()
	store := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	wire := &gatedWireClient{
		gate:  closedGate(),
		grant: LeaseGrant{LeaseID: "lse_1", CompanyID: "co_1", CreditTypeID: "ct_1", GrantedAmount: 1000, ExpiresAt: clock.at(300_000)},
	}
	manager := NewLeaseManager(wire, store, LeaseManagerOptions{Clock: clock.Now})
	manager.Stop()

	assert.Nil(t, manager.AcquireIfNeeded(context.Background(), "co_1", "ct_1"))
	assert.Zero(t, wire.acquires.Load(), "a refused acquire makes no wire call")
	entries, err := store.List(context.Background())
	require.NoError(t, err)
	assert.Empty(t, entries, "and installs nothing for the release pass to miss")
}

// extendScriptWire answers each extend from a queue of server totals, repeating
// the last, and holds the first one open until the test releases it, so a caller
// can arrive while an extend is in flight.
type extendScriptWire struct {
	release chan struct{}

	mu     sync.Mutex
	grants []LeaseGrant
	calls  []extendCall
}

func (w *extendScriptWire) Acquire(context.Context, string, string, float64, time.Time) (*LeaseGrant, error) {
	return nil, errors.New("no acquire is expected here")
}

func (w *extendScriptWire) Extend(ctx context.Context, leaseID string, additionalAmount float64, expiresAt time.Time) (*LeaseGrant, error) {
	w.mu.Lock()
	index := len(w.calls)
	w.calls = append(w.calls, extendCall{leaseID: leaseID, additionalAmount: additionalAmount, expiresAt: expiresAt})
	grant := w.grants[min(index, len(w.grants)-1)]
	w.mu.Unlock()

	if index == 0 {
		select {
		case <-w.release:
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
	grant.LeaseID = leaseID
	return &grant, nil
}

func (w *extendScriptWire) Release(context.Context, string) error { return nil }

func (w *extendScriptWire) extends() []extendCall {
	w.mu.Lock()
	defer w.mu.Unlock()
	return append([]extendCall(nil), w.calls...)
}

// drawnDownSlot is a lease of 1000 with 800 already reserved, so the slot sits
// at 200 remaining, below the water mark every joiner test starts from.
func drawnDownSlot(t *testing.T, wire WireClient) (*LeaseManager, LeaseStore, *virtualClock) {
	t.Helper()
	clock := newVirtualClock()
	store := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	installLease(t, store, clock, "lse_1", "co_1", "ct_1", 1000, 300_000)
	_, _, _, err := store.TryReserve(context.Background(), "co_1", "ct_1", 800)
	require.NoError(t, err)

	manager := NewLeaseManager(wire, store, LeaseManagerOptions{
		Clock:  clock.Now,
		Config: ResolvedLeaseConfig{LeaseSize: 1000, LowWaterMark: 0.25},
	})
	t.Cleanup(manager.Stop)
	return manager, store, clock
}

func serverTotal(clock *virtualClock, total float64) LeaseGrant {
	return LeaseGrant{CompanyID: "co_1", CreditTypeID: "ct_1", GrantedAmount: total, ExpiresAt: clock.at(600_000)}
}

// A watermark extend asking for one tranche is in flight when a check needing
// far more arrives. Taking the tranche would leave that check's post-extend
// reserve failing with the credits sitting on the server.
func TestExtendJoinerWhoseShortfallOutranTheFlightGetsItsOwnTopUp(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	wire := &extendScriptWire{release: make(chan struct{})}
	wire.grants = []LeaseGrant{serverTotal(clock, 2000), serverTotal(clock, 5800)}
	manager, store, _ := drawnDownSlot(t, wire)

	watermark := make(chan *LeaseState, 1)
	go func() { watermark <- manager.MaybeExtend(ctx, "co_1", "ct_1", nil) }()
	require.Eventually(t, func() bool { return len(wire.extends()) == 1 }, time.Second, time.Millisecond)

	required := 5000.0
	joiner := make(chan *LeaseState, 1)
	go func() { joiner <- manager.MaybeExtend(ctx, "co_1", "ct_1", &required) }()
	// The joiner waits the flight out rather than racing a second extend onto
	// the same lease, so nothing reaches the wire while it is held.
	time.Sleep(20 * time.Millisecond)
	assert.Len(t, wire.extends(), 1)

	close(wire.release)
	require.NotNil(t, <-watermark)
	joined := <-joiner

	calls := wire.extends()
	require.Len(t, calls, 2, "exactly one follow-up")
	assert.Equal(t, 1000.0, calls[0].additionalAmount, "the steady-state refresh asks for one tranche")
	// 5000 required against the 1200 the first extend left: sized against the
	// slot that flight just moved, not against the joiner's own stale read.
	assert.Equal(t, 3800.0, calls[1].additionalAmount)
	require.NotNil(t, joined)
	assert.Equal(t, 5000.0, joined.LocalRemainingCredits)

	entry, err := store.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 5000.0, entry.LocalRemainingCredits)
}

// The common case, and the fan-out the follow-up must not introduce: both asks
// fit inside the tranche the flight already carries.
func TestExtendJoinerTheFlightCoversSharesTheOneWireCall(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	wire := &extendScriptWire{release: make(chan struct{})}
	wire.grants = []LeaseGrant{serverTotal(clock, 2000)}
	manager, _, _ := drawnDownSlot(t, wire)

	watermark := make(chan *LeaseState, 1)
	go func() { watermark <- manager.MaybeExtend(ctx, "co_1", "ct_1", nil) }()
	require.Eventually(t, func() bool { return len(wire.extends()) == 1 }, time.Second, time.Millisecond)

	// 900 against 200 remaining is a 700 shortfall, inside the tranche.
	required := 900.0
	joiner := make(chan *LeaseState, 1)
	go func() { joiner <- manager.MaybeExtend(ctx, "co_1", "ct_1", &required) }()
	time.Sleep(20 * time.Millisecond)

	close(wire.release)
	first, joined := <-watermark, <-joiner

	assert.Len(t, wire.extends(), 1, "one wire call serves both")
	require.NotNil(t, joined)
	assert.Equal(t, first, joined)
	assert.Equal(t, 1200.0, joined.LocalRemainingCredits)
}

// The follow-up never chains: a company whose balance simply cannot reach the
// request would otherwise spin extending.
func TestExtendFollowUpDoesNotChainWhenTheServerCannotCover(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	wire := &extendScriptWire{release: make(chan struct{})}
	// The server grants what it has, still far short of the ask.
	wire.grants = []LeaseGrant{serverTotal(clock, 2000), serverTotal(clock, 3000)}
	manager, _, _ := drawnDownSlot(t, wire)

	watermark := make(chan *LeaseState, 1)
	go func() { watermark <- manager.MaybeExtend(ctx, "co_1", "ct_1", nil) }()
	require.Eventually(t, func() bool { return len(wire.extends()) == 1 }, time.Second, time.Millisecond)

	required := 50_000.0
	joiner := make(chan *LeaseState, 1)
	go func() { joiner <- manager.MaybeExtend(ctx, "co_1", "ct_1", &required) }()
	time.Sleep(20 * time.Millisecond)

	close(wire.release)
	require.NotNil(t, <-watermark)
	joined := <-joiner

	calls := wire.extends()
	require.Len(t, calls, 2, "the follow-up spawns no follow-up of its own")
	assert.Equal(t, 48_800.0, calls[1].additionalAmount)
	require.NotNil(t, joined)
	// Resolves short rather than chaining: the caller's reserve fails and the
	// check reports insufficient balance, which is the honest answer.
	assert.Equal(t, 2200.0, joined.LocalRemainingCredits)
}

// Two watermark-driven callers ask for the same tranche, so neither follows up
// and the pair costs one wire call.
func TestConcurrentWatermarkExtendsShareOneWireCall(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	wire := &extendScriptWire{release: make(chan struct{})}
	wire.grants = []LeaseGrant{serverTotal(clock, 2000)}
	manager, _, _ := drawnDownSlot(t, wire)

	first := make(chan *LeaseState, 1)
	go func() { first <- manager.MaybeExtend(ctx, "co_1", "ct_1", nil) }()
	require.Eventually(t, func() bool { return len(wire.extends()) == 1 }, time.Second, time.Millisecond)

	second := make(chan *LeaseState, 1)
	go func() { second <- manager.MaybeExtend(ctx, "co_1", "ct_1", nil) }()
	time.Sleep(20 * time.Millisecond)

	close(wire.release)
	leading, joined := <-first, <-second

	assert.Len(t, wire.extends(), 1)
	require.NotNil(t, joined)
	assert.Equal(t, leading, joined)
	assert.Equal(t, 1200.0, joined.LocalRemainingCredits)
}

// A joiner whose shortfall outran the flight registers a follow-up for the same
// key, so the finishing flight must clear only its own registration.
func TestFlightGroupCleanupLeavesASuccessorRegistered(t *testing.T) {
	t.Parallel()
	group := flightGroup{spawn: goSpawn}
	started, release := make(chan struct{}), make(chan struct{})
	done := make(chan *LeaseState, 1)
	go func() {
		done <- group.do(context.Background(), "slot", func() *LeaseState {
			close(started)
			<-release
			return &LeaseState{LeaseID: "lse_1"}
		})
	}()
	<-started

	successor := &flight{done: make(chan struct{})}
	group.mu.Lock()
	group.inFlight["slot"] = successor
	group.mu.Unlock()

	close(release)
	<-done

	group.mu.Lock()
	registered := group.inFlight["slot"]
	group.mu.Unlock()
	assert.Same(t, successor, registered, "the first flight's cleanup does not evict the follow-up")
}

// staleFirstReadStore hands its first reader a row from before an extend landed,
// so a trigger can reach the flight registration on a view the wire has already
// moved past.
type staleFirstReadStore struct {
	LeaseStore

	mu    sync.Mutex
	stale *LeaseState
}

func (s *staleFirstReadStore) Get(ctx context.Context, companyID, creditTypeID string) (*LeaseState, error) {
	s.mu.Lock()
	stale := s.stale
	s.stale = nil
	s.mu.Unlock()
	if stale != nil {
		return stale, nil
	}
	return s.LeaseStore.Get(ctx, companyID, creditTypeID)
}

// The second trigger reads the slot as it was before the first extend landed:
// that flight is already gone, so nothing stops it reaching the wire but the
// re-read the flight registration now makes.
func TestExtendTriggeredByARowThePreviousExtendAlreadyMovedSendsNothing(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	wire := &extendScriptWire{release: make(chan struct{})}
	close(wire.release)
	wire.grants = []LeaseGrant{serverTotal(clock, 2000)}

	inner := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	installLease(t, inner, clock, "lse_1", "co_1", "ct_1", 1000, 300_000)
	_, _, _, err := inner.TryReserve(ctx, "co_1", "ct_1", 900)
	require.NoError(t, err)
	before, err := inner.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)

	store := &staleFirstReadStore{LeaseStore: inner}
	manager := NewLeaseManager(wire, store, LeaseManagerOptions{
		Clock:  clock.Now,
		Config: ResolvedLeaseConfig{LeaseSize: 1000, LowWaterMark: 0.25},
	})
	t.Cleanup(manager.Stop)

	require.NotNil(t, manager.MaybeExtend(ctx, "co_1", "ct_1", nil))
	require.Len(t, wire.extends(), 1)

	store.mu.Lock()
	store.stale = before
	store.mu.Unlock()
	manager.MaybeExtend(ctx, "co_1", "ct_1", nil)

	assert.Len(t, wire.extends(), 1)
	entry, err := inner.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 2000.0, entry.GrantedAmount)
}

// capturingLogger keeps the warnings a test wants to read back.
type capturingLogger struct {
	noopLogger

	mu       sync.Mutex
	warnings []string
}

func (l *capturingLogger) Warn(_ context.Context, message string, _ ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.warnings = append(l.warnings, message)
}

func (l *capturingLogger) warned(substring string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, warning := range l.warnings {
		if strings.Contains(warning, substring) {
			return true
		}
	}
	return false
}

// stallingReleaseWire answers a release by hanging until the caller's budget
// runs out, the way a wire call to a server that has stopped answering would.
type stallingReleaseWire struct {
	releases atomic.Int64
}

func (w *stallingReleaseWire) Acquire(context.Context, string, string, float64, time.Time) (*LeaseGrant, error) {
	return nil, errors.New("no acquire is expected here")
}

func (w *stallingReleaseWire) Extend(context.Context, string, float64, time.Time) (*LeaseGrant, error) {
	return nil, errors.New("no extend is expected here")
}

func (w *stallingReleaseWire) Release(ctx context.Context, _ string) error {
	w.releases.Add(1)
	<-ctx.Done()
	return ctx.Err()
}

func TestReleaseAllLocalLeasesGivesUpOnAReleaseThatNeverLands(t *testing.T) {
	t.Parallel()
	clock := newVirtualClock()
	store := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	installLease(t, store, clock, "lse_1", "co_1", "ct_1", 1000, 300_000)
	installLease(t, store, clock, "lse_2", "co_2", "ct_1", 1000, 300_000)
	logger := &capturingLogger{}
	wire := &stallingReleaseWire{}
	manager := NewLeaseManager(wire, store, LeaseManagerOptions{Clock: clock.Now, Logger: logger})
	t.Cleanup(manager.Stop)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	started := time.Now()
	manager.ReleaseAllLocalLeases(ctx)

	assert.Less(t, time.Since(started), time.Second)
	assert.True(t, logger.warned("releasing credit leases on close"))
	// The budget ran out inside the first release, so the second is never tried.
	assert.Equal(t, int64(1), wire.releases.Load())
}

// callerKey names the caller a context belongs to, so a test store can hold one
// caller's read and pin an interleaving that is otherwise up to the scheduler.
type callerKey struct{}

// heldReadStore holds the named caller's nth read until the gate opens.
type heldReadStore struct {
	LeaseStore

	caller string
	nth    int
	gate   chan struct{}

	mu    sync.Mutex
	reads int
}

func (s *heldReadStore) Get(ctx context.Context, companyID, creditTypeID string) (*LeaseState, error) {
	if name, _ := ctx.Value(callerKey{}).(string); name == s.caller {
		s.mu.Lock()
		s.reads++
		read := s.reads
		s.mu.Unlock()
		if read == s.nth {
			<-s.gate
		}
	}
	return s.LeaseStore.Get(ctx, companyID, creditTypeID)
}

// sequencedExtendWire announces each extend as it reaches the wire and holds it
// until the test lets it answer.
type sequencedExtendWire struct {
	grants  []LeaseGrant
	started []chan struct{}
	release []chan struct{}

	mu    sync.Mutex
	calls []extendCall
}

func newSequencedExtendWire(grants ...LeaseGrant) *sequencedExtendWire {
	wire := &sequencedExtendWire{grants: grants}
	for range grants {
		wire.started = append(wire.started, make(chan struct{}))
		wire.release = append(wire.release, make(chan struct{}))
	}
	return wire
}

func (w *sequencedExtendWire) Acquire(context.Context, string, string, float64, time.Time) (*LeaseGrant, error) {
	return nil, errors.New("no acquire is expected here")
}

func (w *sequencedExtendWire) Extend(ctx context.Context, leaseID string, additionalAmount float64, expiresAt time.Time) (*LeaseGrant, error) {
	w.mu.Lock()
	index := len(w.calls)
	w.calls = append(w.calls, extendCall{leaseID: leaseID, additionalAmount: additionalAmount, expiresAt: expiresAt})
	w.mu.Unlock()
	if index >= len(w.grants) {
		return nil, errors.New("unscripted extend")
	}

	close(w.started[index])
	select {
	case <-w.release[index]:
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	grant := w.grants[index]
	grant.LeaseID = leaseID
	return &grant, nil
}

func (w *sequencedExtendWire) Release(context.Context, string) error { return nil }

func (w *sequencedExtendWire) extends() []extendCall {
	w.mu.Lock()
	defer w.mu.Unlock()
	return append([]extendCall(nil), w.calls...)
}

// Two checks join one refill, both needing more than it asked for. The smaller
// one's follow-up registers first; taking its result would send the larger one's
// retry back to a lease it already knows is short, with the credits sitting on
// the server.
func TestExtendFollowUpDoesNotInheritASmallerFollowUp(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	wire := newSequencedExtendWire(
		serverTotal(clock, 20_000), // the refill both checks join
		serverTotal(clock, 22_000), // the smaller check's follow-up
		serverTotal(clock, 40_000), // the larger check's own extend
	)
	close(wire.release[2])

	inner := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	installLease(t, inner, clock, "lse_1", "co_1", "ct_1", 10_000, 300_000)
	_, _, _, err := inner.TryReserve(ctx, "co_1", "ct_1", 10_000)
	require.NoError(t, err)
	// The larger check's second read waits for the smaller one's follow-up to
	// reach the wire, so it is the flight the larger check finds on its way back.
	store := &heldReadStore{LeaseStore: inner, caller: "larger", nth: 2, gate: wire.started[1]}
	manager := NewLeaseManager(wire, store, LeaseManagerOptions{
		Clock:  clock.Now,
		Config: ResolvedLeaseConfig{LeaseSize: 1_000, LowWaterMark: 0.25},
	})
	t.Cleanup(manager.Stop)

	refillNeed, smallerNeed, largerNeed := 10_000.0, 12_000.0, 28_000.0
	refill := make(chan *LeaseState, 1)
	go func() { refill <- manager.MaybeExtend(ctx, "co_1", "ct_1", &refillNeed) }()
	<-wire.started[0]

	smaller := make(chan *LeaseState, 1)
	go func() { smaller <- manager.MaybeExtend(ctx, "co_1", "ct_1", &smallerNeed) }()
	time.Sleep(20 * time.Millisecond)
	larger := make(chan *LeaseState, 1)
	go func() {
		larger <- manager.MaybeExtend(context.WithValue(ctx, callerKey{}, "larger"), "co_1", "ct_1", &largerNeed)
	}()
	time.Sleep(20 * time.Millisecond)
	require.Len(t, wire.extends(), 1, "both checks join the refill")

	close(wire.release[0])
	// Let the larger check reach the smaller one's flight and join it before
	// that flight answers.
	<-wire.started[1]
	time.Sleep(20 * time.Millisecond)
	close(wire.release[1])

	require.NotNil(t, <-refill)
	smallerResult, largerResult := <-smaller, <-larger

	calls := wire.extends()
	require.Len(t, calls, 3)
	assert.Equal(t, 10_000.0, calls[0].additionalAmount, "the refill")
	assert.Equal(t, 2_000.0, calls[1].additionalAmount, "the smaller check's shortfall")
	// Sized against the 12,000 the smaller check left behind rather than
	// inherited from its 2,000 ask.
	assert.Equal(t, 16_000.0, calls[2].additionalAmount)
	require.NotNil(t, smallerResult)
	assert.GreaterOrEqual(t, smallerResult.LocalRemainingCredits, smallerNeed)
	require.NotNil(t, largerResult)
	assert.GreaterOrEqual(t, largerResult.LocalRemainingCredits, largerNeed)
}
