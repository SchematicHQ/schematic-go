package leases

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

// t0 is the fixed virtual instant every vector and lease test starts from.
var t0 = time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)

// virtualClock only moves when a test advances it, so nothing here depends on
// wall time.
type virtualClock struct {
	mu        sync.Mutex
	now       time.Time
	onAdvance func(time.Time)
}

func newVirtualClock() *virtualClock {
	return &virtualClock{now: t0}
}

func (c *virtualClock) Now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.now
}

func (c *virtualClock) advance(d time.Duration) {
	c.mu.Lock()
	c.now = c.now.Add(d)
	now, hook := c.now, c.onAdvance
	c.mu.Unlock()
	if hook != nil {
		hook(now)
	}
}

// at is an absolute position on the virtual timeline, as the vectors express
// every *_at_ms field.
func (c *virtualClock) at(offsetMs float64) time.Time {
	return t0.Add(time.Duration(offsetMs) * time.Millisecond)
}

// crashingRefund is a lease-store refunder that fails once while armed,
// reproducing a process death between a reservation's claim and its refund.
// The reservation store refunds through whatever refunder it is handed, so
// wrapping that one leaves the test body reading the real store.
type crashingRefund struct {
	target ReservationRefunder

	mu    sync.Mutex
	armed bool
}

var errSimulatedCrash = errors.New("simulated crash before refund")

func (c *crashingRefund) arm() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.armed = true
}

func (c *crashingRefund) Refund(ctx context.Context, companyID, creditTypeID string, credits float64, pinLeaseID string) error {
	c.mu.Lock()
	armed := c.armed
	c.armed = false
	c.mu.Unlock()
	if armed {
		return errSimulatedCrash
	}
	return c.target.Refund(ctx, companyID, creditTypeID, credits, pinLeaseID)
}

// backend is one store pair plus the seams a vector needs to drive it. server
// and client are set only for the Redis backend.
type backend struct {
	name         string
	clock        *virtualClock
	leases       LeaseStore
	reservations ReservationStore
	crash        *crashingRefund
	server       *miniredis.Miniredis
	client       redis.UniversalClient
}

func newInMemoryBackend(_ *testing.T) *backend {
	clock := newVirtualClock()
	leases := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	crash := &crashingRefund{target: leases}
	return &backend{
		name:         "in_memory",
		clock:        clock,
		leases:       leases,
		reservations: NewInMemoryReservationStore(crash, InMemoryReservationStoreOptions{Clock: clock.Now}),
		crash:        crash,
	}
}

func newRedisBackend(t *testing.T) *backend {
	t.Helper()
	server := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{Addr: server.Addr()})
	t.Cleanup(func() { _ = client.Close() })

	clock := newVirtualClock()
	// The Redis store decides expiry against the store's own clock (TIME), so
	// the virtual clock has to drive miniredis too.
	server.SetTime(clock.Now())
	clock.onAdvance = server.SetTime

	leases := NewRedisLeaseStore(client, RedisLeaseStoreOptions{Clock: clock.Now})
	crash := &crashingRefund{target: leases}
	return &backend{
		name:         "redis",
		clock:        clock,
		leases:       leases,
		reservations: NewRedisReservationStore(client, crash, RedisReservationStoreOptions{Clock: clock.Now}),
		crash:        crash,
		server:       server,
		client:       client,
	}
}

var backendFactories = []struct {
	name string
	make func(*testing.T) *backend
}{
	{name: "in_memory", make: newInMemoryBackend},
	{name: "redis", make: newRedisBackend},
}

// serverLeaseSpec is the lease a vector scripts the server to answer with.
type serverLeaseSpec struct {
	LeaseID       string  `json:"lease_id"`
	GrantedAmount float64 `json:"granted_amount"`
	GrantedTotal  float64 `json:"granted_total"`
	ExpiresAtMs   float64 `json:"expires_at_ms"`
}

type serverScript struct {
	Lease *serverLeaseSpec `json:"lease"`
	Error string           `json:"error"`
}

type acquireCall struct {
	companyID       string
	creditTypeID    string
	requestedAmount float64
	expiresAt       time.Time
}

type extendCall struct {
	leaseID          string
	additionalAmount float64
	expiresAt        time.Time
}

// scriptedWireClient stands in for the lease API: queued responses in, recorded
// calls out.
type scriptedWireClient struct {
	clock *virtualClock

	mu               sync.Mutex
	acquireResponses []serverScript
	extendResponses  []serverScript
	acquireCalls     []acquireCall
	extendCalls      []extendCall
	releaseCalls     []string
	// duringAcquire runs while an acquire is in flight, for emulating a sibling
	// pod winning the race.
	duringAcquire func()
}

func (w *scriptedWireClient) queueAcquire(script serverScript) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.acquireResponses = append(w.acquireResponses, script)
}

func (w *scriptedWireClient) queueExtend(script serverScript) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.extendResponses = append(w.extendResponses, script)
}

func (w *scriptedWireClient) Acquire(_ context.Context, companyID, creditTypeID string, requestedAmount float64, expiresAt time.Time) (*LeaseGrant, error) {
	w.mu.Lock()
	w.acquireCalls = append(w.acquireCalls, acquireCall{
		companyID:       companyID,
		creditTypeID:    creditTypeID,
		requestedAmount: requestedAmount,
		expiresAt:       expiresAt,
	})
	during := w.duringAcquire
	w.duringAcquire = nil
	var script *serverScript
	if len(w.acquireResponses) > 0 {
		script = &w.acquireResponses[0]
		w.acquireResponses = w.acquireResponses[1:]
	}
	w.mu.Unlock()

	if during != nil {
		during()
	}
	lease, err := scriptedLease(script, "unscripted acquire wire call")
	if err != nil {
		return nil, err
	}
	return &LeaseGrant{
		LeaseID:       fallbackString(lease.LeaseID, "lse_unnamed"),
		CompanyID:     companyID,
		CreditTypeID:  creditTypeID,
		GrantedAmount: lease.GrantedAmount,
		ExpiresAt:     w.clock.at(lease.ExpiresAtMs),
	}, nil
}

func (w *scriptedWireClient) Extend(_ context.Context, leaseID string, additionalAmount float64, expiresAt time.Time) (*LeaseGrant, error) {
	w.mu.Lock()
	w.extendCalls = append(w.extendCalls, extendCall{
		leaseID:          leaseID,
		additionalAmount: additionalAmount,
		expiresAt:        expiresAt,
	})
	var script *serverScript
	if len(w.extendResponses) > 0 {
		script = &w.extendResponses[0]
		w.extendResponses = w.extendResponses[1:]
	}
	w.mu.Unlock()

	lease, err := scriptedLease(script, "unscripted extend wire call")
	if err != nil {
		return nil, err
	}
	granted := lease.GrantedTotal
	if granted == 0 {
		granted = lease.GrantedAmount
	}
	return &LeaseGrant{
		LeaseID:       leaseID,
		CompanyID:     "co_wire",
		CreditTypeID:  "ct_wire",
		GrantedAmount: granted,
		ExpiresAt:     w.clock.at(lease.ExpiresAtMs),
	}, nil
}

func (w *scriptedWireClient) Release(_ context.Context, leaseID string) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.releaseCalls = append(w.releaseCalls, leaseID)
	return nil
}

func (w *scriptedWireClient) snapshot() (acquires []acquireCall, extends []extendCall, releases []string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return append([]acquireCall(nil), w.acquireCalls...),
		append([]extendCall(nil), w.extendCalls...),
		append([]string(nil), w.releaseCalls...)
}

func scriptedLease(script *serverScript, missing string) (*serverLeaseSpec, error) {
	if script == nil {
		return nil, errors.New(missing)
	}
	if script.Error != "" || script.Lease == nil {
		return nil, errors.New(fallbackString(script.Error, missing))
	}
	return script.Lease, nil
}

// installLease writes a lease into a store the way a vector's `given` block
// does, asserting the install took.
func installLease(t *testing.T, store LeaseStore, clock *virtualClock, leaseID, companyID, creditTypeID string, granted, expiresAtMs float64) {
	t.Helper()
	wrote, err := store.Replace(context.Background(), LeaseGrant{
		LeaseID:       leaseID,
		CompanyID:     companyID,
		CreditTypeID:  creditTypeID,
		GrantedAmount: granted,
		ExpiresAt:     clock.at(expiresAtMs),
	})
	require.NoError(t, err)
	require.True(t, wrote)
}
