package leases

import (
	"context"
	"math"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newReservation(id, leaseID string, reserved float64, expiresAtMs float64, clock *virtualClock) ReservationRecord {
	return ReservationRecord{
		ID:               id,
		LeaseID:          leaseID,
		CompanyID:        "co_1",
		CreditTypeID:     "ct_1",
		EventSubtype:     "inference_tokens",
		QuantityReserved: reserved / 10,
		CreditsReserved:  reserved,
		ConsumptionRate:  10,
		ExpiresAt:        clock.at(expiresAtMs),
		Company:          map[string]string{"id": "co_1"},
	}
}

// A NaN actual quantity must clamp to zero rather than debit the lease by NaN,
// which JSON vectors cannot express.
func TestInMemoryReservationStoreClampsNonFiniteConsumption(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	leases := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	installLease(t, leases, clock, "lse_1", "co_1", "ct_1", 1000, 3_600_000)
	reservations := NewInMemoryReservationStore(leases, InMemoryReservationStoreOptions{Clock: clock.Now})

	_, _, err := leases.TryReserve(ctx, "co_1", "ct_1", 100)
	require.NoError(t, err)
	require.NoError(t, reservations.Add(ctx, newReservation("res_1", "lse_1", 100, 60_000, clock)))

	consumed, claimed, err := reservations.Consume(ctx, "res_1", math.NaN())
	require.NoError(t, err)
	assert.True(t, claimed)
	assert.Equal(t, 0.0, consumed)

	entry, err := leases.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 1000.0, entry.LocalRemainingCredits)
}

// The claim is the exactly-once arbiter: of many racing settles only one comes
// away with the record, so the unspent slice refunds exactly once.
func TestInMemoryReservationStoreConcurrentConsumeClaimsOnce(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	leases := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	installLease(t, leases, clock, "lse_1", "co_1", "ct_1", 1000, 3_600_000)
	reservations := NewInMemoryReservationStore(leases, InMemoryReservationStoreOptions{Clock: clock.Now})

	_, _, err := leases.TryReserve(ctx, "co_1", "ct_1", 100)
	require.NoError(t, err)
	require.NoError(t, reservations.Add(ctx, newReservation("res_1", "lse_1", 100, 60_000, clock)))

	var claims atomic.Int64
	var wg sync.WaitGroup
	for i := 0; i < 32; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, claimed, err := reservations.Consume(ctx, "res_1", 40)
			assert.NoError(t, err)
			if claimed {
				claims.Add(1)
			}
		}()
	}
	wg.Wait()

	assert.Equal(t, int64(1), claims.Load())
	entry, err := leases.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 960.0, entry.LocalRemainingCredits)
}

// A sweep and a settle racing on the same reservation must also arbitrate
// through the claim, so the hold refunds once whichever wins.
func TestInMemoryReservationStoreSweepRacesSettleWithoutDoubleRefunding(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	clock := newVirtualClock()
	leases := NewInMemoryLeaseStore(InMemoryLeaseStoreOptions{Clock: clock.Now})
	installLease(t, leases, clock, "lse_1", "co_1", "ct_1", 1000, 3_600_000)
	reservations := NewInMemoryReservationStore(leases, InMemoryReservationStoreOptions{Clock: clock.Now})

	_, _, err := leases.TryReserve(ctx, "co_1", "ct_1", 100)
	require.NoError(t, err)
	// Already past its TTL, so the sweeper is eligible to claim it too.
	require.NoError(t, reservations.Add(ctx, newReservation("res_1", "lse_1", 100, 0, clock)))

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		_, _, err := reservations.Consume(ctx, "res_1", 0)
		assert.NoError(t, err)
	}()
	go func() {
		defer wg.Done()
		_, err := reservations.SweepExpired(ctx)
		assert.NoError(t, err)
	}()
	wg.Wait()

	entry, err := leases.Get(ctx, "co_1", "ct_1")
	require.NoError(t, err)
	assert.Equal(t, 1000.0, entry.LocalRemainingCredits)
	count, err := reservations.Count(ctx)
	require.NoError(t, err)
	assert.Zero(t, count)
}
