package leases

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/core"
	"github.com/schematichq/schematic-go/credits"
)

// WireClient is the three lease calls the manager makes. Narrow on purpose: it
// keeps the manager independent of the generated client's request and response
// models, and lets tests script the server.
type WireClient interface {
	Acquire(ctx context.Context, companyID, creditTypeID string, requestedAmount float64, expiresAt time.Time) (*LeaseGrant, error)
	Extend(ctx context.Context, leaseID string, additionalAmount float64, expiresAt time.Time) (*LeaseGrant, error)
	Release(ctx context.Context, leaseID string) error
}

// apiWireClient adapts the generated credits client to WireClient.
type apiWireClient struct {
	credits *credits.Client
}

// NewAPIWireClient adapts the generated credits client for the lease manager.
func NewAPIWireClient(client *credits.Client) WireClient {
	return &apiWireClient{credits: client}
}

// Acquire takes the default retry policy: the server hands back the slot's
// existing active lease rather than opening a second one, so a retry after a
// lost response returns the lease the first attempt created.
func (c *apiWireClient) Acquire(ctx context.Context, companyID, creditTypeID string, requestedAmount float64, expiresAt time.Time) (*LeaseGrant, error) {
	response, err := c.credits.AcquireCreditLease(
		ctx,
		&schematicgo.AcquireCreditLeaseRequestBody{
			CompanyID:       companyID,
			CreditTypeID:    creditTypeID,
			RequestedAmount: requestedAmount,
			ExpiresAt:       &expiresAt,
		},
	)
	if err != nil {
		return nil, err
	}
	return grantFromResponse(response.GetData())
}

func (c *apiWireClient) Extend(ctx context.Context, leaseID string, additionalAmount float64, expiresAt time.Time) (*LeaseGrant, error) {
	// An extend is an increment, so a retry without a key would grant the
	// tranche twice. The key is minted once per extend and the retry loop
	// resends this body, so every attempt of this extend collapses to one grow
	// while a later extend gets its own key.
	idempotencyKey := uuid.NewString()
	response, err := c.credits.ExtendCreditLease(
		ctx,
		leaseID,
		&schematicgo.ExtendCreditLeaseRequestBody{
			AdditionalAmount: additionalAmount,
			ExpiresAt:        &expiresAt,
			IdempotencyKey:   &idempotencyKey,
		},
	)
	if err != nil {
		return nil, err
	}
	return grantFromResponse(response.GetData())
}

func (c *apiWireClient) Release(ctx context.Context, leaseID string) error {
	_, err := c.credits.ReleaseCreditLease(ctx, leaseID)
	return err
}

func grantFromResponse(data *schematicgo.CreditLeaseResponseData) (*LeaseGrant, error) {
	if data == nil {
		return nil, fmt.Errorf("credit lease response carried no data")
	}
	return &LeaseGrant{
		LeaseID:       data.GetID(),
		CompanyID:     data.GetCompanyID(),
		CreditTypeID:  data.GetCreditTypeID(),
		GrantedAmount: data.GetGrantedAmount(),
		ExpiresAt:     data.GetExpiresAt(),
	}, nil
}

// LeaseManager owns lease rows for one client: acquire on first use or after
// expiry, extend when the local view dips below the water mark, release on
// close.
//
// Acquire and extend each get their own best-effort single-flight map keyed by
// slot. Best-effort because callers racing ahead of the registration can still
// issue duplicate wire calls, which is safe: the server is idempotent for an
// active slot, Replace keeps the first live lease, and Extend reconciles to a
// total.
//
// Every path here resolves rather than fails: callers route a missing lease
// through their fail-open/fail-closed handling, and several calls are made
// fire-and-forget, where an error has nowhere to go.
type LeaseManager struct {
	wire         WireClient
	leases       LeaseStore
	reservations ReservationStore
	config       ResolvedLeaseConfig
	overrides    map[string]LeaseOverride
	sweepEvery   time.Duration
	stopTimeout  time.Duration
	clock        Clock
	logger       core.Logger

	// Kept separate so an in-flight extend can never satisfy an acquire, or the
	// other way round.
	acquireFlights flightGroup
	extendFlights  flightGroup

	background sync.WaitGroup

	mu          sync.Mutex
	stopped     bool
	sweepCancel context.CancelFunc
	sweepDone   chan struct{}
}

// LeaseManagerOptions configures a LeaseManager. The zero value resolves every
// knob to its package default.
type LeaseManagerOptions struct {
	// Reservations is swept on the sweep interval by StartSweep. Optional.
	Reservations ReservationStore
	// Config is the client-wide knobs. A zero field falls through to the
	// package default.
	Config ResolvedLeaseConfig
	// Overrides are the per-credit-type knob overrides.
	Overrides map[string]LeaseOverride
	// SweepInterval defaults to DefaultSweepInterval.
	SweepInterval time.Duration
	// StopTimeout bounds how long Stop waits on fire-and-forget work. Defaults
	// to DefaultStopTimeout.
	StopTimeout time.Duration
	// Clock defaults to time.Now.
	Clock Clock
	// Logger defaults to a no-op logger.
	Logger core.Logger
}

func NewLeaseManager(wire WireClient, leases LeaseStore, opts LeaseManagerOptions) *LeaseManager {
	sweepEvery := opts.SweepInterval
	if sweepEvery <= 0 {
		sweepEvery = DefaultSweepInterval
	}
	stopTimeout := opts.StopTimeout
	if stopTimeout <= 0 {
		stopTimeout = DefaultStopTimeout
	}
	logger := opts.Logger
	if logger == nil {
		logger = noopLogger{}
	}
	manager := &LeaseManager{
		wire:         wire,
		leases:       leases,
		reservations: opts.Reservations,
		config:       opts.Config,
		overrides:    opts.Overrides,
		sweepEvery:   sweepEvery,
		stopTimeout:  stopTimeout,
		clock:        orNow(opts.Clock),
		logger:       logger,
	}
	manager.acquireFlights.spawn = manager.spawn
	manager.extendFlights.spawn = manager.spawn
	return manager
}

// ResolveConfig returns the knobs for one credit type.
func (m *LeaseManager) ResolveConfig(creditTypeID string) ResolvedLeaseConfig {
	return ResolveLeaseConfig(m.config, m.overrides, creditTypeID)
}

// AcquireIfNeeded returns the slot's live lease, acquiring one over the wire if
// none is live. It returns nil rather than an error when the wire or the store
// is down, so the caller routes the outcome through fail-open/fail-closed.
func (m *LeaseManager) AcquireIfNeeded(ctx context.Context, companyID, creditTypeID string) *LeaseState {
	if m.isStopped() {
		// A lease installed after ReleaseAllLocalLeases has listed the slots
		// would be held until it expires server-side, with nobody left to
		// release it.
		m.logger.Debug(ctx, fmt.Sprintf("Not acquiring a credit lease for %s/%s: the manager is stopped", companyID, creditTypeID))
		return nil
	}
	existing, err := m.leases.Get(ctx, companyID, creditTypeID)
	if err != nil {
		m.logger.Error(ctx, fmt.Sprintf("Failed to read lease store for %s/%s: %v", companyID, creditTypeID, err))
		return nil
	}
	// Liveness here is judged on this pod's clock, while the Redis store re-reads
	// expiry against the Redis server's clock inside TryReserve. The two can
	// disagree, so a lease this call hands back can still be refused there, and
	// the check routes that through its fail-open handling. Assumed rather than
	// reconciled: the stores keep an expired row for a grace window precisely so
	// clocks within it agree on what is live, and a TIME round trip per check
	// would buy nothing else.
	if existing != nil && existing.ExpiresAt.After(m.clock()) {
		return existing
	}
	// An expired (or absent) slot is left for Replace to overwrite: it guards
	// on expiry and writes atomically. Dropping the stale row first would be a
	// separate, non-atomic op that can interleave between a sibling pod's read
	// and its replace, clobbering a lease that pod just installed. Reading a
	// stale entry in the gap is harmless, since every path that acts on a lease
	// re-guards on expiry.
	return m.acquireFlights.do(ctx, LeaseKey(companyID, creditTypeID), func() *LeaseState {
		// The flight is shared, so it runs detached: the first caller's
		// cancellation or short deadline would otherwise hand every waiter on
		// the slot a nil lease. Each waiter still honors its own context while
		// it waits.
		detached, cancel := detachedContext(ctx)
		defer cancel()
		return m.acquire(detached, companyID, creditTypeID)
	})
}

func (m *LeaseManager) acquire(ctx context.Context, companyID, creditTypeID string) *LeaseState {
	resolved := m.ResolveConfig(creditTypeID)
	grant, err := m.wire.Acquire(ctx, companyID, creditTypeID, resolved.LeaseSize, m.clock().Add(resolved.LeaseDuration))
	if err != nil {
		m.logger.Error(ctx, fmt.Sprintf("Failed to acquire credit lease for %s/%s: %v", companyID, creditTypeID, err))
		return nil
	}
	wrote, err := m.leases.Replace(ctx, LeaseGrant{
		LeaseID:       grant.LeaseID,
		CompanyID:     fallbackString(grant.CompanyID, companyID),
		CreditTypeID:  fallbackString(grant.CreditTypeID, creditTypeID),
		GrantedAmount: grant.GrantedAmount,
		ExpiresAt:     grant.ExpiresAt,
	})
	if err != nil {
		m.logger.Error(ctx, fmt.Sprintf("Failed to install credit lease %s: %v", grant.LeaseID, err))
		return nil
	}

	current, err := m.leases.Get(ctx, companyID, creditTypeID)
	if err != nil {
		m.logger.Error(ctx, fmt.Sprintf("Failed to read lease store for %s/%s: %v", companyID, creditTypeID, err))
		return nil
	}
	if wrote {
		return current
	}

	// A sibling holds the slot with a live lease, or the slot's expired row was
	// reconciled in place. The server is idempotent for an active slot, so a
	// racing acquire is normally handed back the SAME lease the sibling
	// installed, and releasing it would pull the shared lease out from under
	// every pod. Only a different lease is a redundant hold nobody will draw
	// on, so only that one is released. An empty slot (expired in the gap)
	// releases nothing either: this lease is likely what the next acquire is
	// handed.
	if current != nil && current.LeaseID != grant.LeaseID {
		m.logger.Debug(ctx, fmt.Sprintf("Lost acquire race for %s/%s; releasing redundant lease %s", companyID, creditTypeID, grant.LeaseID))
		detached, cancel := detachedContext(ctx)
		if !m.spawn(func() {
			defer cancel()
			if err := m.wire.Release(detached, grant.LeaseID); err != nil {
				m.logger.Warn(ctx, fmt.Sprintf("Failed to release redundant credit lease %s: %v", grant.LeaseID, err))
			}
		}) {
			cancel()
		}
	}
	return current
}

// MaybeExtend extends the slot's lease when the local view warrants it,
// triggered by either the low-water-mark ratio (steady-state refresh) or a
// requiredCredits hint above the local remaining (a check just failed a reserve
// of that size). Pass nil for requiredCredits to ask for the steady-state check
// only. It returns nil rather than an error on any failure.
//
// A caller arriving while an extend is in flight joins it. If its own shortfall
// is larger than what that extend asked for, it waits the flight out and then
// issues exactly one follow-up extend for the remaining difference; otherwise it
// would inherit a tranche-sized ask and fail its post-extend retry with credits
// still sitting on the server.
func (m *LeaseManager) MaybeExtend(ctx context.Context, companyID, creditTypeID string, requiredCredits *float64) *LeaseState {
	return m.maybeExtend(ctx, companyID, creditTypeID, requiredCredits, true)
}

func (m *LeaseManager) maybeExtend(ctx context.Context, companyID, creditTypeID string, requiredCredits *float64, allowFollowUp bool) *LeaseState {
	entry := m.readLiveLease(ctx, companyID, creditTypeID)
	if entry == nil {
		return nil
	}
	resolved := m.ResolveConfig(creditTypeID)
	if !m.needsExtend(entry, resolved, requiredCredits) {
		return entry
	}
	// Size the extend to cover the request that triggered it: a single check
	// needing more than remaining plus one tranche would otherwise fail its
	// post-extend retry forever, however much balance the server has. The
	// steady-state path keeps asking for the configured tranche. Sized here, one
	// level above the wire call, so the flight registered below and the request
	// body provably carry the same number for a joiner to compare against.
	shortfall := 0.0
	if requiredCredits != nil {
		shortfall = *requiredCredits - entry.LocalRemainingCredits
	}
	additionalAmount := max(resolved.LeaseSize, shortfall)

	result, flightAsk, joined := m.extendFlights.doSized(ctx, LeaseKey(companyID, creditTypeID), additionalAmount, func() *LeaseState {
		// Detached for the same reason the acquire flight is: one caller's
		// deadline must not decide what every waiter on the slot gets.
		detached, cancel := detachedContext(ctx)
		defer cancel()
		return m.recheckAndExtend(detached, companyID, creditTypeID, resolved, requiredCredits, additionalAmount)
	})
	// The flight already asked for at least what we need, which covers every
	// watermark-driven joiner and any check the tranche fits. One wire call
	// serves all of them, which is the point of single-flight.
	if !joined || additionalAmount <= flightAsk {
		return result
	}
	// Our shortfall outran the flight's ask. We waited it out rather than racing
	// a second extend onto the same lease, and now top up the difference with
	// exactly one more, re-read against the slot that flight just moved. The
	// follow-up is not allowed one of its own: a company whose balance simply
	// cannot reach the request would otherwise spin. A caller whose own context
	// has already ended has stopped waiting for the answer, so a follow-up would
	// buy it nothing.
	if !allowFollowUp || ctx.Err() != nil {
		return result
	}
	return m.maybeExtend(ctx, companyID, creditTypeID, requiredCredits, false)
}

// readLiveLease reads the slot, reporting nothing when the read fails or the
// lease is absent or expired. An expired lease is never extended: the server
// treats it as released and has already refunded its remainder, so the only
// correct move is a fresh acquire on the next check.
func (m *LeaseManager) readLiveLease(ctx context.Context, companyID, creditTypeID string) *LeaseState {
	entry, err := m.leases.Get(ctx, companyID, creditTypeID)
	if err != nil {
		m.logger.Warn(ctx, fmt.Sprintf("Failed to read lease store for %s/%s: %v", companyID, creditTypeID, err))
		return nil
	}
	if entry == nil {
		return nil
	}
	if !entry.ExpiresAt.After(m.clock()) {
		return nil
	}
	return entry
}

// needsExtend reports whether the slot sits low enough to warrant an extend.
func (m *LeaseManager) needsExtend(entry *LeaseState, resolved ResolvedLeaseConfig, requiredCredits *float64) bool {
	belowWatermark := entry.LocalRemainingCredits/max(entry.GrantedAmount, 1) <= resolved.LowWaterMark
	belowRequired := requiredCredits != nil && entry.LocalRemainingCredits < *requiredCredits
	return belowWatermark || belowRequired
}

// recheckAndExtend re-reads the slot now that this flight owns it, and extends
// only if the fresh row still warrants one. The row that decided this extend was
// read before the flight was registered, so an extend that landed in that gap,
// clearing its own flight on the way out, would otherwise be followed by a
// second extend, under a new idempotency key, for a lease it already topped up.
// The registered ask stands: a joiner compares its shortfall against that
// figure, so the wire body has to carry it.
func (m *LeaseManager) recheckAndExtend(
	ctx context.Context,
	companyID, creditTypeID string,
	resolved ResolvedLeaseConfig,
	requiredCredits *float64,
	additionalAmount float64,
) *LeaseState {
	entry := m.readLiveLease(ctx, companyID, creditTypeID)
	if entry == nil {
		return nil
	}
	if !m.needsExtend(entry, resolved, requiredCredits) {
		return entry
	}
	return m.extend(ctx, *entry, resolved, additionalAmount)
}

// ExtendInBackground kicks off a water-mark extend without waiting for it: a
// check that just drew the lease down should not pay for the top-up.
func (m *LeaseManager) ExtendInBackground(ctx context.Context, companyID, creditTypeID string) {
	detached, cancel := detachedContext(ctx)
	if !m.spawn(func() {
		defer cancel()
		m.MaybeExtend(detached, companyID, creditTypeID, nil)
	}) {
		cancel()
	}
}

// detachedContext survives the caller's cancellation, since the caller is not
// waiting on this work, but still expires: the wire client sets no HTTP timeout
// of its own, so an untimed call would pin a goroutine for as long as the server
// holds the socket.
func detachedContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.WithoutCancel(ctx), DetachedWorkTimeout)
}

func (m *LeaseManager) extend(ctx context.Context, entry LeaseState, resolved ResolvedLeaseConfig, additionalAmount float64) *LeaseState {
	grant, err := m.wire.Extend(ctx, entry.LeaseID, additionalAmount, m.clock().Add(resolved.LeaseDuration))
	if err != nil {
		m.logger.Warn(ctx, fmt.Sprintf("Failed to extend credit lease %s: %v", entry.LeaseID, err))
		return nil
	}
	// Reconcile to the server's authoritative TOTAL, with the store computing
	// the delta against its own current total: per-process single-flight does
	// not cover sibling pods. Pinned to the lease the server extended, so an
	// expiry mid-call cannot mint the delta onto a successor.
	expiresAt := grant.ExpiresAt
	if err := m.leases.Extend(ctx, entry.CompanyID, entry.CreditTypeID, grant.GrantedAmount, &expiresAt, entry.LeaseID); err != nil {
		m.logger.Warn(ctx, fmt.Sprintf("Failed to reconcile extended credit lease %s: %v", entry.LeaseID, err))
		return nil
	}
	current, err := m.leases.Get(ctx, entry.CompanyID, entry.CreditTypeID)
	if err != nil {
		m.logger.Warn(ctx, fmt.Sprintf("Failed to read lease store for %s/%s: %v", entry.CompanyID, entry.CreditTypeID, err))
		return nil
	}
	return current
}

// ReleaseAllLocalLeases releases every live lease this process exclusively
// holds, returning their unspent remainders to the company balance immediately
// instead of waiting out the lease expiry.
//
// Only a per-process store implements LeaseLister; a shared backend is skipped,
// since sibling pods still draw on those leases. Expired leases are skipped
// too: the server already swept them. Best-effort, with failures falling back
// to server-side expiry.
func (m *LeaseManager) ReleaseAllLocalLeases(ctx context.Context) {
	lister, ok := m.leases.(LeaseLister)
	if !ok {
		return
	}
	entries, err := lister.List(ctx)
	if err != nil {
		m.logger.Warn(ctx, fmt.Sprintf("Failed to enumerate leases on close: %v", err))
		return
	}
	now := m.clock()
	for _, entry := range entries {
		if !entry.ExpiresAt.After(now) {
			continue
		}
		if err := m.wire.Release(ctx, entry.LeaseID); err != nil {
			m.logger.Warn(ctx, fmt.Sprintf("Failed to release credit lease %s on close (it will expire server-side): %v", entry.LeaseID, err))
			continue
		}
		if err := m.leases.Drop(ctx, entry.CompanyID, entry.CreditTypeID); err != nil {
			m.logger.Warn(ctx, fmt.Sprintf("Failed to drop released credit lease %s: %v", entry.LeaseID, err))
			continue
		}
		m.logger.Debug(ctx, fmt.Sprintf("Released credit lease %s on close", entry.LeaseID))
	}
}

// StartSweep runs the expired-reservation sweep on the configured interval.
// Safe to call twice; a no-op without a reservation store or after Stop.
func (m *LeaseManager) StartSweep() {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.reservations == nil || m.stopped || m.sweepCancel != nil {
		return
	}
	// The sweep itself runs on this context, not a background one, so a Stop
	// landing behind a Redis backlog cancels the sweep in flight instead of
	// waiting out a whole pass over the reservation table.
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	m.sweepCancel = cancel
	m.sweepDone = done
	go m.sweepLoop(ctx, done)
}

func (m *LeaseManager) sweepLoop(ctx context.Context, done chan<- struct{}) {
	defer close(done)
	ticker := time.NewTicker(m.sweepEvery)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if _, err := m.reservations.SweepExpired(ctx); err != nil && ctx.Err() == nil {
				// Keep the loop alive: a sweep failure is transient (a Redis
				// blip), and the next tick retries.
				m.logger.Debug(ctx, fmt.Sprintf("Reservation sweep failed: %v", err))
			}
		}
	}
}

// Stop halts the sweep loop and waits for it and any fire-and-forget lease work
// to finish. The wait on that work is bounded by StopTimeout: the work is capped
// at DetachedWorkTimeout and unwinds on its own, so a shutdown does not sit
// behind it.
func (m *LeaseManager) Stop() {
	m.StopWithContext(context.Background())
}

// StopWithContext is Stop under the caller's own shutdown budget: every wait
// ends when ctx does, so a client tearing down several subsystems spends one
// deadline across all of them rather than each timeout in turn.
func (m *LeaseManager) StopWithContext(ctx context.Context) {
	m.mu.Lock()
	cancelSweep, done := m.sweepCancel, m.sweepDone
	m.stopped = true
	m.sweepCancel, m.sweepDone = nil, nil
	m.mu.Unlock()

	if cancelSweep != nil {
		cancelSweep()
		select {
		case <-done:
		case <-ctx.Done():
		}
	}

	waited := make(chan struct{})
	go func() {
		m.background.Wait()
		close(waited)
	}()
	timer := time.NewTimer(m.stopTimeout)
	defer timer.Stop()
	select {
	case <-waited:
	case <-ctx.Done():
	case <-timer.C:
		m.logger.Warn(ctx, fmt.Sprintf(
			"Timed out after %s waiting for background credit lease work; it is capped at %s and unwinds on its own",
			m.stopTimeout, DetachedWorkTimeout,
		))
	}
}

// isStopped reads the flag under the lock spawn takes, so an acquire and a Stop
// racing cannot both conclude they got there first.
func (m *LeaseManager) isStopped() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.stopped
}

// spawn runs a fire-and-forget step, reporting whether it started one. It
// refuses after Stop, where an Add would race the wait it is already past, and
// the work would touch a manager that is being torn down. It never propagates a
// panic into the process either: these paths are unawaited, so there is nobody
// to recover for them.
func (m *LeaseManager) spawn(fn func()) bool {
	m.mu.Lock()
	if m.stopped {
		m.mu.Unlock()
		return false
	}
	m.background.Add(1)
	m.mu.Unlock()

	go func() {
		defer m.background.Done()
		defer func() {
			if recovered := recover(); recovered != nil {
				m.logger.Error(context.Background(), fmt.Sprintf("Background lease work panicked: %v", recovered))
			}
		}()
		fn()
	}()
	return true
}

// drainBackground waits out pending fire-and-forget work, so a test can assert
// on what it did.
func (m *LeaseManager) drainBackground() {
	m.background.Wait()
}

// flightGroup collapses concurrent calls for one slot into a single in-flight
// call, which runs on a goroutine of its own so that no one caller's context
// decides what every other caller on the slot gets. Best-effort: a caller
// racing ahead of the registration still issues its own wire call.
type flightGroup struct {
	// spawn runs the shared call, tracked so a Stop waits it out and a panic in
	// it never reaches the host process.
	spawn func(func()) bool

	mu       sync.Mutex
	inFlight map[string]*flight
}

type flight struct {
	done chan struct{}
	// requestedAdditional is what an extend flight's wire call asks the server
	// for, the figure a joiner compares its own shortfall against. Acquire
	// flights are not sized against anything and leave it zero.
	requestedAdditional float64
	result              *LeaseState
}

func (g *flightGroup) do(ctx context.Context, key string, fn func() *LeaseState) *LeaseState {
	result, _, _ := g.doSized(ctx, key, 0, fn)
	return result
}

// doSized is do for a call whose size matters to a joiner: requestedAdditional
// is what this call will ask the server for. It reports the figure the flight it
// joined asked for, and whether it joined one rather than starting its own, so a
// caller the flight does not cover can follow up.
func (g *flightGroup) doSized(ctx context.Context, key string, requestedAdditional float64, fn func() *LeaseState) (*LeaseState, float64, bool) {
	g.mu.Lock()
	if existing, ok := g.inFlight[key]; ok {
		g.mu.Unlock()
		return existing.wait(ctx), existing.requestedAdditional, true
	}
	call := &flight{done: make(chan struct{}), requestedAdditional: requestedAdditional}
	if g.inFlight == nil {
		g.inFlight = make(map[string]*flight)
	}
	g.inFlight[key] = call
	g.mu.Unlock()

	// The cleanup is identity-guarded rather than an unconditional delete: a
	// joiner whose shortfall outran this flight registers a follow-up for the
	// same key, and this flight must not evict it.
	finish := func() {
		g.mu.Lock()
		if g.inFlight[key] == call {
			delete(g.inFlight, key)
		}
		g.mu.Unlock()
		close(call.done)
	}
	if !g.spawn(func() {
		defer finish()
		call.result = fn()
	}) {
		finish()
		return nil, requestedAdditional, false
	}
	return call.wait(ctx), requestedAdditional, false
}

// wait blocks for the shared call, or for the caller's own context to end,
// whichever lands first.
func (f *flight) wait(ctx context.Context) *LeaseState {
	select {
	case <-f.done:
		return f.result
	case <-ctx.Done():
		return nil
	}
}

// noopLogger keeps the manager's logging calls total without forcing a logger
// on the caller.
type noopLogger struct{}

func (noopLogger) Debug(context.Context, string, ...any) {}
func (noopLogger) Info(context.Context, string, ...any)  {}
func (noopLogger) Warn(context.Context, string, ...any)  {}
func (noopLogger) Error(context.Context, string, ...any) {}
