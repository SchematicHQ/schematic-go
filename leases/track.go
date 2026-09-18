package leases

import (
	"context"
	"math"

	schematicgo "github.com/schematichq/schematic-go"
)

// SettleOutcome is what a settle did locally, and what it owes the server.
type SettleOutcome struct {
	// Track is the billing event to emit.
	Track *schematicgo.EventBodyTrack
	// SettledLocally is true when the hold was still open and this call debited
	// the consumed slice and refunded the rest. False when it had already been
	// swept at its TTL, already settled, or the store was unreachable: the lease
	// balance was not touched here, so it reads high until the lease rolls over,
	// and the track event is a recovery emit.
	SettledLocally bool
	// Err is the store failure behind a settle that did not land. The event is
	// built either way, since the server is the source of truth for real
	// consumption and the usage still has to be billed.
	Err error
}

// SettleReservation consumes a client-mode hold against its lease and builds the
// event that bills it.
//
// The event comes from the caller-held record rather than the store, so the
// usage is still billed once the hold has been swept. Only the local bookkeeping
// clamps to the reserved amount; the event carries the unclamped actual.
func SettleReservation(
	ctx context.Context,
	reservations ReservationStore,
	record ReservationRecord,
	actualQuantity float64,
) SettleOutcome {
	// Rounded up for the same reason the hold is (see CheckWithLease): the debit
	// has to move the local ledger by exactly what the track event bills.
	_, claimed, err := reservations.Consume(ctx, record.ID, math.Ceil(actualQuantity)*record.ConsumptionRate)
	return SettleOutcome{
		Track:          BuildTrackEvent(record, SettleQuantity(actualQuantity)),
		SettledLocally: err == nil && claimed,
		Err:            err,
	}
}

// BuildTrackEvent builds the event that settles a client-mode hold, from the
// record alone.
//
// Kept free of store access so the client can still bill the usage when the
// local settle fails against an unreachable store. The lease id routes the
// server-side consumption through the lease's sub-ledger instead of
// decrementing a grant the acquire already pre-debited. Caller traits ride on
// the event through the client's own track options, so they are not this
// function's business.
func BuildTrackEvent(record ReservationRecord, quantity int64) *schematicgo.EventBodyTrack {
	return &schematicgo.EventBodyTrack{
		Company:  record.Company,
		Event:    record.EventSubtype,
		LeaseID:  &record.LeaseID,
		Quantity: &quantity,
		User:     record.User,
	}
}

// SettleQuantity casts a settled usage onto the integer a track event records. A
// hold can be sized from a fractional usage, but the event's quantity is an
// integer, so a partial unit settles as a whole one rather than as none.
func SettleQuantity(actualQuantity float64) int64 {
	return int64(math.Ceil(actualQuantity))
}
