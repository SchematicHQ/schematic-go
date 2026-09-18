package leases

import (
	"math"
	"time"
)

// Clock reads the current time. Every store and the lease manager take one so
// tests and the conformance runner can drive a virtual clock instead of wall
// time.
type Clock func() time.Time

const (
	// DefaultLeaseDuration is the lease lifetime requested at acquire and
	// extend: expiresAt = now + duration.
	DefaultLeaseDuration = 5 * time.Minute
	// DefaultReservationTTL is the reservation lifetime, and the sweep
	// deadline. Size it above the longest expected gap between a check and its
	// settle: a settle arriving after the TTL still bills the server but no
	// longer re-debits the lease, so the local balance reads high until the
	// lease rolls over.
	DefaultReservationTTL = 60 * time.Second
	// DefaultLeaseSize is the credit amount requested per acquire, and the
	// minimum extend tranche.
	DefaultLeaseSize = 10000.0
	// DefaultLowWaterMark is the remaining/granted ratio at or below which a
	// background extend is kicked off.
	DefaultLowWaterMark = 0.25
	// DefaultSweepInterval is the cadence of the expired-reservation sweep.
	DefaultSweepInterval = time.Second
	// DefaultStopTimeout bounds how long a Stop waits on fire-and-forget lease
	// work, so a stalled wire call cannot hold a shutdown open.
	DefaultStopTimeout = 10 * time.Second
	// DetachedWorkTimeout caps a fire-and-forget lease call. The work is
	// deliberately detached from the caller's context, so without a cap of its
	// own a wire client with no HTTP timeout would leave it running forever.
	DetachedWorkTimeout = 30 * time.Second
	// DefaultPrewarmResolveTimeout is how long a prewarm waits for a freshly
	// identified company to surface in the datastream cache before giving up.
	DefaultPrewarmResolveTimeout = 5 * time.Second

	// MaxReservationTTL is the furthest out the API will hold credits, so a
	// larger server-mode reservation TTL would fail every check.
	MaxReservationTTL = time.Hour
	// ReservationTTLSkewAllowance is held back from MaxReservationTTL because
	// the API measures that hour against its own clock while the SDK computes
	// expiresAt against the caller's: a client running ahead would otherwise be
	// rejected at exactly the cap.
	ReservationTTLSkewAllowance = 60 * time.Second
)

// LeaseState is the local view of the one lease a (company, credit type) slot
// holds.
type LeaseState struct {
	LeaseID      string
	CompanyID    string
	CreditTypeID string
	// GrantedAmount is the server-authoritative total granted to this lease.
	// It grows on extend.
	GrantedAmount float64
	// LocalRemainingCredits is granted minus outstanding holds and
	// consumption. It starts at the full grant when the lease is installed.
	LocalRemainingCredits float64
	// ExpiresAt is the instant past which the lease is dead: the server has
	// refunded the remainder to the company balance, so the local balance is
	// stale and must never serve another reserve.
	ExpiresAt time.Time
}

// LeaseGrant is what the server says a lease is, after an acquire or an
// extend. It is also what installs a lease into a store: the local balance is
// derived, never supplied.
type LeaseGrant struct {
	LeaseID      string
	CompanyID    string
	CreditTypeID string
	// GrantedAmount is the server-authoritative TOTAL, not the increment an
	// extend asked for.
	GrantedAmount float64
	ExpiresAt     time.Time
}

// ReservationRecord is one credit hold carved out of a lease by a check.
type ReservationRecord struct {
	ID string
	// LeaseID is the lease the hold was carved from. It pins refunds, so a
	// hold from an expired lease can never inflate its successor.
	LeaseID      string
	CompanyID    string
	CreditTypeID string
	// EventSubtype is what the settling track event is billed as.
	EventSubtype string
	// QuantityReserved is the caller-declared usage, in event units.
	QuantityReserved float64
	// CreditsReserved is ceil(QuantityReserved) * ConsumptionRate: whole event
	// units, so the hold matches what the settle bills.
	CreditsReserved float64
	ConsumptionRate float64
	// ExpiresAt is the instant past which the sweeper refunds the full hold.
	ExpiresAt time.Time
	// Company and User are the evaluation context the hold was issued for,
	// threaded onto the track event so the server attributes usage to the same
	// company and user.
	Company map[string]string
	User    map[string]string
}

// ResolvedLeaseConfig is the knobs for a single credit type, after overrides
// and defaults. A zero field means "unset": ResolveLeaseConfig falls through to
// the package default for it.
type ResolvedLeaseConfig struct {
	LeaseDuration  time.Duration
	ReservationTTL time.Duration
	LeaseSize      float64
	LowWaterMark   float64
}

// LeaseOverride overrides the four resolvable knobs for one credit type. A nil
// field leaves the client-wide value in place.
type LeaseOverride struct {
	LeaseDuration  *time.Duration
	ReservationTTL *time.Duration
	LeaseSize      *float64
	LowWaterMark   *float64
}

// ResolveLeaseConfig resolves the knobs for one credit type: the credit type's
// override wins, then the client-wide base, then the package default.
func ResolveLeaseConfig(base ResolvedLeaseConfig, overrides map[string]LeaseOverride, creditTypeID string) ResolvedLeaseConfig {
	resolved := ResolvedLeaseConfig{
		LeaseDuration:  DefaultLeaseDuration,
		ReservationTTL: DefaultReservationTTL,
		LeaseSize:      DefaultLeaseSize,
		LowWaterMark:   DefaultLowWaterMark,
	}
	if base.LeaseDuration > 0 {
		resolved.LeaseDuration = base.LeaseDuration
	}
	if base.ReservationTTL > 0 {
		resolved.ReservationTTL = base.ReservationTTL
	}
	if base.LeaseSize > 0 {
		resolved.LeaseSize = base.LeaseSize
	}
	if base.LowWaterMark > 0 {
		resolved.LowWaterMark = base.LowWaterMark
	}
	override, ok := overrides[creditTypeID]
	if !ok {
		return resolved
	}
	if override.LeaseDuration != nil {
		resolved.LeaseDuration = *override.LeaseDuration
	}
	if override.ReservationTTL != nil {
		resolved.ReservationTTL = *override.ReservationTTL
	}
	if override.LeaseSize != nil {
		resolved.LeaseSize = *override.LeaseSize
	}
	if override.LowWaterMark != nil {
		resolved.LowWaterMark = *override.LowWaterMark
	}
	return resolved
}

// IsValidQuantity reports whether a caller-supplied quantity can size a credit
// hold. NaN is the dangerous case: it slips through every numeric comparison,
// and a NaN balance would approve every later reserve on a possibly shared
// lease.
func IsValidQuantity(value float64) bool {
	return !math.IsNaN(value) && !math.IsInf(value, 0) && value >= 0
}
