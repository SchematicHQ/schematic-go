package leases

import (
	"context"
	"slices"
	"sync"
	"time"
)

// LeaseKey is the (company, credit type) slot key. Exported so the Redis
// backends and any external store implementation namespace slots alike.
func LeaseKey(companyID, creditTypeID string) string {
	return companyID + ":" + creditTypeID
}

// LeaseStore holds at most one lease per (company, credit type) slot. Every
// mutation is atomic per slot: InMemoryLeaseStore gets that from a per-slot
// mutex, RedisLeaseStore from single-key Lua.
type LeaseStore interface {
	// Get returns a snapshot of the slot, expired or not, or nil when the slot
	// is empty. Callers re-guard on expiry.
	Get(ctx context.Context, companyID, creditTypeID string) (*LeaseState, error)

	// Replace installs a fresh lease at its full grant, if the slot is free to
	// take, and reports whether it wrote.
	//
	// A live lease holds the slot even when it carries a different id (a
	// sibling pod won the acquire race): its already-debited balance wins and
	// this reports false. An expired row carrying the SAME id is not rewritten
	// either, since that would reset the balance and erase debits whose
	// reservations are still open; it is reconciled like an extend (granted to
	// the incoming total, expiry forward only, balance untouched) and also
	// reports false. Only a fresh write reports true, which is what tells the
	// manager whether the lease it just acquired is redundant.
	Replace(ctx context.Context, grant LeaseGrant) (bool, error)

	// TryReserve atomically checks and debits, returning the post-debit
	// balance and the lease the credits came out of. It reports ok=false,
	// touching nothing, when there is no lease, the lease has expired, the
	// balance is short, or credits is not a finite non-negative number.
	// Returning the balance rather than a bool lets the caller derive the
	// pre-debit figure as balance+credits without a racy follow-up read.
	//
	// leaseID is read in the same atomic step as the debit. The slot's lease
	// can be replaced between a caller's acquire and its reserve, so a hold
	// pinned to the lease the caller last saw would send its refunds to a lease
	// that never held the credits, and bill that lease for the usage.
	TryReserve(ctx context.Context, companyID, creditTypeID string, credits float64) (balance float64, leaseID string, ok bool, err error)

	// Refund returns credits to the slot's balance, clamped at the granted
	// amount. With a non-empty pinLeaseID the refund applies only while the
	// slot still holds that lease: a hold carved out of an expired lease must
	// never inflate its successor, whose grant the server already issued whole.
	Refund(ctx context.Context, companyID, creditTypeID string, credits float64, pinLeaseID string) error

	// Extend reconciles the slot to the server-authoritative total. The delta
	// is computed inside the store against the currently stored total, never
	// from a caller-held pre-wire-call read: two pods extending concurrently
	// from the same stale read would each apply a delta and mint phantom
	// credits. A total a sibling already applied is a no-op, so applies
	// converge in any order. Expiry only ever moves forward. A non-empty
	// pinLeaseID drops the whole extend when the slot holds a different lease.
	Extend(ctx context.Context, companyID, creditTypeID string, grantedTotal float64, newExpiresAt *time.Time, pinLeaseID string) error

	// Drop removes the slot entry, after a remote release.
	Drop(ctx context.Context, companyID, creditTypeID string) error
}

// LeaseLister is implemented only by a per-process store, whose leases are
// exclusively this process's, so releasing them on close is safe. A shared
// backend must never enumerate and release: sibling pods still draw on those
// leases.
type LeaseLister interface {
	List(ctx context.Context) ([]LeaseState, error)
}

// InMemoryLeaseStore keeps lease slots in this process only, so it gates a
// single pod. Swap in RedisLeaseStore to gate across pods; both implement
// LeaseStore.
type InMemoryLeaseStore struct {
	clock  Clock
	locks  slotLocks
	mu     sync.Mutex
	leases map[string]LeaseState
}

// InMemoryLeaseStoreOptions configures an InMemoryLeaseStore. The zero value
// is the production configuration.
type InMemoryLeaseStoreOptions struct {
	// Clock defaults to time.Now.
	Clock Clock
}

func NewInMemoryLeaseStore(opts InMemoryLeaseStoreOptions) *InMemoryLeaseStore {
	return &InMemoryLeaseStore{
		clock:  orNow(opts.Clock),
		locks:  newSlotLocks(),
		leases: make(map[string]LeaseState),
	}
}

func (s *InMemoryLeaseStore) Get(_ context.Context, companyID, creditTypeID string) (*LeaseState, error) {
	key := LeaseKey(companyID, creditTypeID)
	release := s.locks.lock(key)
	defer release()
	return s.load(key), nil
}

func (s *InMemoryLeaseStore) Replace(_ context.Context, grant LeaseGrant) (bool, error) {
	key := LeaseKey(grant.CompanyID, grant.CreditTypeID)
	release := s.locks.lock(key)
	defer release()

	existing := s.load(key)
	if existing != nil && existing.ExpiresAt.After(s.clock()) {
		return false, nil
	}
	if existing != nil && existing.LeaseID == grant.LeaseID {
		// The same lease coming back over its own expired row: a stale acquire
		// response for a lease the idempotent server also handed a racing
		// sibling, which may since have extended it. Rewriting would reset the
		// balance and erase debits whose reservations are still open.
		if add := grant.GrantedAmount - existing.GrantedAmount; add > 0 {
			existing.GrantedAmount = grant.GrantedAmount
			existing.LocalRemainingCredits += add
		}
		if grant.ExpiresAt.After(existing.ExpiresAt) {
			existing.ExpiresAt = grant.ExpiresAt
		}
		s.store(key, *existing)
		return false, nil
	}
	s.store(key, LeaseState{
		LeaseID:               grant.LeaseID,
		CompanyID:             grant.CompanyID,
		CreditTypeID:          grant.CreditTypeID,
		GrantedAmount:         grant.GrantedAmount,
		LocalRemainingCredits: grant.GrantedAmount,
		ExpiresAt:             grant.ExpiresAt,
	})
	return true, nil
}

func (s *InMemoryLeaseStore) TryReserve(_ context.Context, companyID, creditTypeID string, credits float64) (float64, string, bool, error) {
	// NaN passes every comparison below, and a NaN balance would approve every
	// later reserve, so it never reaches the arithmetic.
	if !IsValidQuantity(credits) {
		return 0, "", false, nil
	}
	key := LeaseKey(companyID, creditTypeID)
	release := s.locks.lock(key)
	defer release()

	entry := s.load(key)
	if entry == nil {
		return 0, "", false, nil
	}
	if !entry.ExpiresAt.After(s.clock()) {
		return 0, "", false, nil
	}
	if entry.LocalRemainingCredits < credits {
		return 0, "", false, nil
	}
	entry.LocalRemainingCredits -= credits
	s.store(key, *entry)
	// Read under the slot lock, so it names the lease this debit actually came
	// out of rather than whichever one the slot holds by the time the caller
	// looks again.
	return entry.LocalRemainingCredits, entry.LeaseID, true, nil
}

func (s *InMemoryLeaseStore) Refund(_ context.Context, companyID, creditTypeID string, credits float64, pinLeaseID string) error {
	if !IsValidQuantity(credits) || credits <= 0 {
		return nil
	}
	key := LeaseKey(companyID, creditTypeID)
	release := s.locks.lock(key)
	defer release()

	entry := s.load(key)
	if entry == nil {
		return nil
	}
	if pinLeaseID != "" && entry.LeaseID != pinLeaseID {
		return nil
	}
	entry.LocalRemainingCredits = min(entry.LocalRemainingCredits+credits, entry.GrantedAmount)
	s.store(key, *entry)
	return nil
}

func (s *InMemoryLeaseStore) Extend(_ context.Context, companyID, creditTypeID string, grantedTotal float64, newExpiresAt *time.Time, pinLeaseID string) error {
	key := LeaseKey(companyID, creditTypeID)
	release := s.locks.lock(key)
	defer release()

	entry := s.load(key)
	if entry == nil {
		return nil
	}
	if pinLeaseID != "" && entry.LeaseID != pinLeaseID {
		return nil
	}
	if add := grantedTotal - entry.GrantedAmount; add > 0 {
		entry.GrantedAmount = grantedTotal
		entry.LocalRemainingCredits += add
	}
	if newExpiresAt != nil && newExpiresAt.After(entry.ExpiresAt) {
		entry.ExpiresAt = *newExpiresAt
	}
	s.store(key, *entry)
	return nil
}

func (s *InMemoryLeaseStore) Drop(_ context.Context, companyID, creditTypeID string) error {
	key := LeaseKey(companyID, creditTypeID)
	release := s.locks.lock(key)
	defer release()

	s.mu.Lock()
	delete(s.leases, key)
	s.mu.Unlock()
	return nil
}

// List returns every slot this process holds, ordered by slot key so a caller
// releasing them on close does so deterministically.
func (s *InMemoryLeaseStore) List(_ context.Context) ([]LeaseState, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	keys := make([]string, 0, len(s.leases))
	for key := range s.leases {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	entries := make([]LeaseState, 0, len(keys))
	for _, key := range keys {
		entries = append(entries, s.leases[key])
	}
	return entries, nil
}

func (s *InMemoryLeaseStore) load(key string) *LeaseState {
	s.mu.Lock()
	defer s.mu.Unlock()
	entry, ok := s.leases[key]
	if !ok {
		return nil
	}
	return &entry
}

func (s *InMemoryLeaseStore) store(key string, entry LeaseState) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.leases[key] = entry
}

// slotLocks hands out one mutex per slot, refcounted so idle slots do not
// accumulate.
type slotLocks struct {
	mu    sync.Mutex
	locks map[string]*slotLock
}

type slotLock struct {
	mu      sync.Mutex
	holders int
}

func newSlotLocks() slotLocks {
	return slotLocks{locks: make(map[string]*slotLock)}
}

// lock blocks until the slot is free and returns the release func.
func (s *slotLocks) lock(key string) func() {
	s.mu.Lock()
	if s.locks == nil {
		s.locks = make(map[string]*slotLock)
	}
	lock, ok := s.locks[key]
	if !ok {
		lock = &slotLock{}
		s.locks[key] = lock
	}
	lock.holders++
	s.mu.Unlock()

	lock.mu.Lock()
	return func() {
		lock.mu.Unlock()
		s.mu.Lock()
		lock.holders--
		if lock.holders == 0 {
			delete(s.locks, key)
		}
		s.mu.Unlock()
	}
}

func orNow(clock Clock) Clock {
	if clock != nil {
		return clock
	}
	return time.Now
}
