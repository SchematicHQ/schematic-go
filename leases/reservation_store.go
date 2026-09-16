package leases

import (
	"context"
	"math"
	"sync"
)

// ReservationStore holds the open credit holds carved out of leases.
//
// Add does not debit: the debit already landed in LeaseStore.TryReserve, and
// that ordering is what bounds a crash to a leaked hold rather than a
// double-spend.
type ReservationStore interface {
	// Add registers a reservation. Idempotent on id.
	Add(ctx context.Context, reservation ReservationRecord) error

	// Get looks up a reservation, or returns nil once it has been claimed or
	// swept.
	Get(ctx context.Context, id string) (*ReservationRecord, error)

	// Consume claims a reservation exactly once and refunds its unspent slice.
	//
	// The claim is atomic and comes first: a racing settle or sweep finds
	// nothing to claim, reports claimed=false, and refunds nothing. On a
	// successful claim creditsConsumed is clamped to [0, CreditsReserved], the
	// remainder is refunded to the lease (pinned to the reservation's lease),
	// and the clamped figure is returned. A crash between the claim and the
	// refund loses the refund; it never double-refunds. An error means the
	// caller must treat the settle as not landed even though the claim may
	// already have.
	Consume(ctx context.Context, id string, creditsConsumed float64) (consumed float64, claimed bool, err error)

	// ReservedCredits sums CreditsReserved across the slot's open
	// reservations. A hold counts exactly while it is in the table, so
	// LocalRemainingCredits + ReservedCredits stays exact between operations.
	ReservedCredits(ctx context.Context, companyID, creditTypeID string) (float64, error)

	// SweepExpired removes every reservation past its TTL, refunding each full
	// hold, and returns how many it swept. Refunds are pinned to the
	// originating lease, so a hold carved from a lease that has since expired
	// is dropped rather than credited to its successor.
	SweepExpired(ctx context.Context) (int, error)

	// Count reports the open reservations across every slot.
	Count(ctx context.Context) (int, error)
}

// InMemoryReservationStore keeps the reservation table in this process, and
// refunds into the lease store it is handed.
type InMemoryReservationStore struct {
	leases ReservationRefunder
	clock  Clock

	mu           sync.Mutex
	reservations map[string]ReservationRecord
}

// ReservationRefunder is the slice of LeaseStore a reservation store needs.
// Narrowing it here is what lets the unspent-slice refund stay an ordinary
// single-key step rather than a cross-key script, and lets a test interpose on
// the claim-then-refund window.
type ReservationRefunder interface {
	Refund(ctx context.Context, companyID, creditTypeID string, credits float64, pinLeaseID string) error
}

// InMemoryReservationStoreOptions configures an InMemoryReservationStore. The
// zero value is the production configuration.
type InMemoryReservationStoreOptions struct {
	// Clock defaults to time.Now.
	Clock Clock
}

func NewInMemoryReservationStore(leases ReservationRefunder, opts InMemoryReservationStoreOptions) *InMemoryReservationStore {
	return &InMemoryReservationStore{
		leases:       leases,
		clock:        orNow(opts.Clock),
		reservations: make(map[string]ReservationRecord),
	}
}

func (s *InMemoryReservationStore) Add(_ context.Context, reservation ReservationRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.reservations[reservation.ID] = reservation
	return nil
}

func (s *InMemoryReservationStore) Get(_ context.Context, id string) (*ReservationRecord, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	reservation, ok := s.reservations[id]
	if !ok {
		return nil, nil
	}
	return &reservation, nil
}

func (s *InMemoryReservationStore) Consume(ctx context.Context, id string, creditsConsumed float64) (float64, bool, error) {
	reservation, claimed := s.claim(id)
	if !claimed {
		return 0, false, nil
	}
	consumed := clampConsumption(creditsConsumed, reservation.CreditsReserved)
	refund := reservation.CreditsReserved - consumed
	if refund > 0 {
		// Pinned to the originating lease: if that lease has expired and a
		// successor holds the slot, the refund is dropped, because the expired
		// lease's remainder already went back to the company balance
		// server-side.
		if err := s.leases.Refund(ctx, reservation.CompanyID, reservation.CreditTypeID, refund, reservation.LeaseID); err != nil {
			return 0, false, err
		}
	}
	return consumed, true, nil
}

func (s *InMemoryReservationStore) ReservedCredits(_ context.Context, companyID, creditTypeID string) (float64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	total := 0.0
	for _, reservation := range s.reservations {
		if reservation.CompanyID == companyID && reservation.CreditTypeID == creditTypeID {
			total += reservation.CreditsReserved
		}
	}
	return total, nil
}

func (s *InMemoryReservationStore) SweepExpired(ctx context.Context) (int, error) {
	cutoff := s.clock()
	s.mu.Lock()
	expired := make([]string, 0)
	for id, reservation := range s.reservations {
		if !reservation.ExpiresAt.After(cutoff) {
			expired = append(expired, id)
		}
	}
	s.mu.Unlock()

	swept := 0
	for _, id := range expired {
		// Routed through Consume so the sweep claims exactly once too.
		_, claimed, err := s.Consume(ctx, id, 0)
		if err != nil {
			return swept, err
		}
		if claimed {
			swept++
		}
	}
	return swept, nil
}

func (s *InMemoryReservationStore) Count(_ context.Context) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.reservations), nil
}

// claim removes the reservation from the table under the lock, so of two
// racing callers exactly one comes away with the record.
func (s *InMemoryReservationStore) claim(id string) (ReservationRecord, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	reservation, ok := s.reservations[id]
	if !ok {
		return ReservationRecord{}, false
	}
	delete(s.reservations, id)
	return reservation, true
}

// clampConsumption keeps local bookkeeping from ever debiting a lease past the
// hold it took.
func clampConsumption(creditsConsumed, creditsReserved float64) float64 {
	if math.IsNaN(creditsConsumed) || creditsConsumed < 0 {
		return 0
	}
	return min(creditsConsumed, creditsReserved)
}
