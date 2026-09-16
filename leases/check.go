package leases

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/core"
	"github.com/schematichq/schematic-go/rulesengine"
)

// FailOpenBalance is the balance a fail-open evaluation substitutes for the
// metered credit: large enough that the credit gate always passes, and still
// exact as a JSON number, so the engine reads back what the SDK sent. It is
// Node's Number.MAX_SAFE_INTEGER, which the conformance vectors name
// "max_safe_integer".
const FailOpenBalance = float64(1<<53 - 1)

// EventUsage is a simulated quantity scoped to one event subtype.
type EventUsage struct {
	EventSubtype string
	Quantity     int64
}

// EvalOptions is the preflight a lease-bearing check asks the rules engine to
// answer.
//
// It is plain data rather than a slice of rulesengine.CheckFlagOption because
// those close over an unexported struct, so a test could never read back what
// the flow asked for. The client translates this to engine options at the
// DataStream seam.
type EvalOptions struct {
	// CreditCost prices the action per credit type, bypassing the engine's own
	// quantity times rate arithmetic.
	CreditCost map[string]float64
	// Usage is the generic preflight quantity, used when no subtype narrows it.
	Usage *int64
	// EventUsage is the preflight quantity scoped to one event subtype.
	EventUsage *EventUsage
}

// CheckDataStream is the slice of the DataStream client a lease-bearing check
// touches. Narrow on purpose: it keeps this package off the wider DataStream
// surface, keeps the flow importable by the client that owns that surface, and
// lets the conformance runner script the flow without a socket.
type CheckDataStream interface {
	// GetFlag reads a flag from the local cache, reporting whether it was there.
	GetFlag(ctx context.Context, key string) (*rulesengine.Flag, bool)
	// GetCompany resolves company keys, cache first, then over the wire.
	GetCompany(ctx context.Context, keys map[string]string) (*rulesengine.Company, error)
	// GetUser resolves user keys, cache first, then over the wire.
	GetUser(ctx context.Context, keys map[string]string) (*rulesengine.User, error)
	// EvaluateFlag runs the rules engine. A nil opts means no preflight.
	EvaluateFlag(
		ctx context.Context,
		flag *rulesengine.Flag,
		company *rulesengine.Company,
		user *rulesengine.User,
		opts *EvalOptions,
	) (*rulesengine.CheckFlagResult, error)
}

// CheckDeps is everything a lease-bearing check draws on, gathered by the
// client.
type CheckDeps struct {
	DataStream   CheckDataStream
	Leases       LeaseStore
	Reservations ReservationStore
	Manager      *LeaseManager
	// Logger defaults to a no-op logger.
	Logger core.Logger
	// Clock defaults to time.Now. It sizes reservation expiry.
	Clock Clock
	// EmitFlagCheck reports a flag_check event for a check this flow resolved
	// itself. The plain check paths enqueue one per check, so without it a
	// lease-gated check would be invisible to flag-check analytics and to
	// company last-seen. Fallback exits do not call it: the plain check they
	// defer to reports its own.
	EmitFlagCheck func(context.Context, *schematicgo.EventBodyFlagCheck)
	// NewReservationID mints hold ids. Defaults to a UUID.
	NewReservationID func() string
}

// CheckRequest is one caller's ask.
type CheckRequest struct {
	FlagKey string
	// Company and User are the caller's evaluation keys, threaded onto the hold
	// so the settling track event attributes usage to the same entities.
	Company map[string]string
	User    map[string]string
	// Usage is the units of the metered event the operation is about to record.
	Usage float64
	// EventSubtype names the event the usage applies to. Empty defers to the
	// entitlement's own subtype.
	EventSubtype string
	// FailOpen errs on the side of assuming the credits are there when the flow
	// cannot gate. See handleLeaseFailure.
	FailOpen bool
}

// CheckOutcome is what the flow resolved, before the client maps it onto its
// public result type.
type CheckOutcome struct {
	Allowed     bool
	Value       bool
	Reason      string
	FlagKey     string
	FlagID      *string
	Entitlement *rulesengine.FeatureEntitlement
	// Reservation is the hold this check carved out of the lease, when it took
	// one.
	Reservation *ReservationRecord
	Error       string
}

// CheckWithLease gates one check against a local lease, returning a hold when it
// allows.
//
// fallback is the plain flag check. Every step that cannot resolve a credit to
// meter defers to it, since the plain check has its own degradation story and
// issues no hold. A step that can resolve the credit but cannot gate on it goes
// through the caller's fail-open/fail-closed contract instead.
//
// One check runs the rules engine twice. The first run is a probe against the
// company's real balance that names the credit being metered; the second gates
// the call against the lease's local balance, after the credits have already
// been debited. conformance/SPEC.md explains why each step is ordered the way it
// is, and the vectors pin it.
func CheckWithLease(
	ctx context.Context,
	deps CheckDeps,
	req CheckRequest,
	fallback func(context.Context) *CheckOutcome,
) *CheckOutcome {
	log := deps.logger()

	// A malformed usage must never reach the stores: NaN slips through every
	// numeric comparison, and a NaN balance on a possibly shared lease would
	// approve every later reserve. The caller asked for a contract covering
	// exactly this, so resolve it through that rather than letting it surface as
	// an opaque reserve failure.
	if !IsValidQuantity(req.Usage) {
		log.Error(ctx, fmt.Sprintf("Lease check: invalid usage %v for flag %s; must be a finite, non-negative number", req.Usage, req.FlagKey))
		return deps.emit(ctx, req, staticFailureOutcome(req, "invalid_usage", nil), flagCheckIDs{})
	}

	// Nothing to reserve. The plain check still carries the preflight, so every
	// rule evaluates normally; a 0-credit hold would only be a no-op.
	if req.Usage == 0 {
		log.Debug(ctx, fmt.Sprintf("Lease check: usage is 0 for flag %s, nothing to reserve, using a plain check", req.FlagKey))
		return fallback(ctx)
	}

	if deps.DataStream == nil {
		log.Debug(ctx, "Lease check: no DataStream, using a plain check")
		return fallback(ctx)
	}

	flag, found := deps.DataStream.GetFlag(ctx, req.FlagKey)
	if !found || flag == nil {
		log.Debug(ctx, fmt.Sprintf("Lease check: no cached flag for %s, using a plain check", req.FlagKey))
		return fallback(ctx)
	}

	if len(req.Company) == 0 {
		log.Debug(ctx, "Lease check: no company keys, using a plain check")
		return fallback(ctx)
	}

	// Resolve company and user the way a plain DataStream check does: cache
	// first, then a live fetch. Evaluating without an entity the caller named
	// would silently skip its targeted rules and overrides, so a miss defers to
	// the plain check instead.
	company, err := deps.DataStream.GetCompany(ctx, req.Company)
	if err != nil || company == nil {
		log.Debug(ctx, fmt.Sprintf("Lease check: company fetch failed for keys %v (%v), using a plain check", req.Company, err))
		return fallback(ctx)
	}

	var user *rulesengine.User
	if len(req.User) > 0 {
		user, err = deps.DataStream.GetUser(ctx, req.User)
		if err != nil || user == nil {
			log.Debug(ctx, fmt.Sprintf("Lease check: user fetch failed for keys %v (%v), using a plain check", req.User, err))
			return fallback(ctx)
		}
	}

	// Entitlement-first resolution. The probe runs against the real balance with
	// no preflight: applying a credit cost to a lease-depleted server balance
	// could fail the credit condition, drop the engine to a lower-priority rule,
	// and hide the entitlement being looked for.
	probe, err := deps.DataStream.EvaluateFlag(ctx, flag, company, user, nil)
	if err != nil || probe == nil {
		// A probe failure is a resolution miss, not the gate, and no hold exists
		// yet to cancel.
		log.Warn(ctx, fmt.Sprintf("Lease check: entitlement probe failed for flag %s (%v), using a plain check", req.FlagKey, err))
		return fallback(ctx)
	}

	entitlement := probe.Entitlement
	if entitlement == nil || entitlement.ValueType != rulesengine.EntitlementValueTypeCredit {
		// A boolean or override grant, a numeric allocation, unlimited, or
		// simply not entitled. The feature resolves without drawing a credit, so
		// skip the lease and the reserve round trip entirely.
		log.Debug(ctx, fmt.Sprintf("Lease check: flag %s matched a non-credit entitlement (value_type=%s), using a plain check, no reservation", req.FlagKey, entitlementValueType(entitlement)))
		return fallback(ctx)
	}

	creditID := derefString(entitlement.CreditID)
	consumptionRate := derefFloat(entitlement.ConsumptionRate)
	// The caller's subtype wins; otherwise the entitlement names the metered
	// event. A credit entitlement with neither a resolvable subtype nor a
	// positive rate can never be billed, so it is not gateable.
	eventSubtype := req.EventSubtype
	if eventSubtype == "" {
		eventSubtype = derefString(entitlement.EventSubtype)
	}
	if creditID == "" || consumptionRate <= 0 || eventSubtype == "" {
		log.Debug(ctx, fmt.Sprintf(
			"Lease check: flag %s has an incomplete credit entitlement (credit_id=%q, consumption_rate=%v, event_subtype=%q), using a plain check",
			req.FlagKey, creditID, consumptionRate, eventSubtype,
		))
		return fallback(ctx)
	}

	creditCost := req.Usage * consumptionRate
	ids := flagCheckIDs{companyID: &company.ID, userID: userID(user)}

	// Every can't-gate outcome funnels through here, so the fail-open and
	// fail-closed contract holds even when the backing infrastructure is down.
	failure := func(reason string) *CheckOutcome {
		return deps.emit(ctx, req, handleLeaseFailure(ctx, deps, req, reason, flag, company, user, creditID), ids)
	}

	lease := deps.Manager.AcquireIfNeeded(ctx, company.ID, creditID)
	if lease == nil {
		return failure("lease_acquire_failed")
	}

	// TryReserve is the atomic gate: check and debit in one step, returning the
	// post-debit balance so the pre-debit figure needs no second read.
	balance, reserved, err := deps.Leases.TryReserve(ctx, company.ID, creditID, creditCost)
	if err == nil && !reserved {
		// Pass the cost as required credits so a single large request extends
		// even while the ratio still sits above the water mark.
		deps.Manager.MaybeExtend(ctx, company.ID, creditID, &creditCost)
		balance, reserved, err = deps.Leases.TryReserve(ctx, company.ID, creditID, creditCost)
	}
	if err != nil {
		log.Error(ctx, fmt.Sprintf("Lease check: reserve against %s/%s failed: %v", company.ID, creditID, err))
		return failure("lease_store_error")
	}
	if !reserved {
		return failure("insufficient_lease_balance")
	}

	// Record the hold after the debit and before the gate. A crash between the
	// debit and this add leaks at most this one hold, reclaimed when the lease
	// expires server-side; recording first would instead leave a record with no
	// debit, which a later consume would refund into a double-spend.
	resolved := deps.Manager.ResolveConfig(creditID)
	record := ReservationRecord{
		ID:               deps.newReservationID(),
		LeaseID:          lease.LeaseID,
		CompanyID:        company.ID,
		CreditTypeID:     creditID,
		EventSubtype:     eventSubtype,
		QuantityReserved: req.Usage,
		CreditsReserved:  creditCost,
		ConsumptionRate:  consumptionRate,
		ExpiresAt:        deps.now().Add(resolved.ReservationTTL),
		Company:          req.Company,
		User:             req.User,
	}
	if err := deps.Reservations.Add(ctx, record); err != nil {
		log.Error(ctx, fmt.Sprintf("Lease check: failed to persist reservation %s: %v", record.ID, err))
		undoDebit(ctx, deps, record, lease.LeaseID)
		return failure("lease_store_error")
	}

	// Gate against the lease's local view rather than the server's balance. The
	// substituted figure is the pre-reservation balance (what TryReserve
	// returned plus what it debited, exact as of the debit), and the credit cost
	// tells the engine what this call costs, so it evaluates the same arithmetic
	// TryReserve just enforced, plus every non-credit rule.
	substituted := substituteCreditBalance(company, creditID, balance+creditCost)
	gateOptions := &EvalOptions{CreditCost: map[string]float64{creditID: creditCost}}
	result, err := deps.DataStream.EvaluateFlag(ctx, flag, substituted, user, gateOptions)
	if err != nil || result == nil {
		log.Error(ctx, fmt.Sprintf("Lease check: rules evaluation failed for flag %s: %v", req.FlagKey, err))
		// The engine itself is down, so there is no fail-open re-evaluation to
		// run: resolve the mode statically.
		cancelReservation(ctx, deps, record)
		return deps.emit(ctx, req, staticFailureOutcome(req, fmt.Sprintf("wasm_error: %v", err), flag), ids)
	}

	// Engine-evaluated exits report the engine's resolved ids, mirroring the
	// plain DataStream path's flag_check event.
	ids = flagCheckIDs{
		companyID: firstString(result.CompanyID, &company.ID),
		userID:    firstString(result.UserID, userID(user)),
		ruleID:    result.RuleID,
	}

	if !result.Value {
		cancelReservation(ctx, deps, record)
		return deps.emit(ctx, req, &CheckOutcome{
			Reason:      fallbackString(result.Reason, "denied_by_engine"),
			FlagKey:     fallbackString(result.FlagKey, req.FlagKey),
			FlagID:      result.FlagID,
			Entitlement: result.Entitlement,
		}, ids)
	}

	// Allowed against the substituted balance, so the hold stands. Top the lease
	// up in the background now that it has been drawn down: the check that drew
	// it down should not pay for the top-up.
	deps.Manager.ExtendInBackground(ctx, company.ID, creditID)
	return deps.emit(ctx, req, &CheckOutcome{
		Allowed:     true,
		Value:       true,
		Reason:      fallbackString(result.Reason, "lease_reserved"),
		FlagKey:     fallbackString(result.FlagKey, req.FlagKey),
		FlagID:      result.FlagID,
		Entitlement: result.Entitlement,
		Reservation: &record,
	}, ids)
}

// handleLeaseFailure resolves a check that could not gate: acquire failed, store
// unreachable, or the lease is exhausted.
//
// Fail-closed denies. Fail-open means assume the credits are there, not skip the
// evaluation: the rules still run with the balance substituted to an effectively
// unlimited value, so plan targeting, overrides, and every non-credit condition
// still apply, and a company that is not entitled stays denied with the lease
// backend down. Only an error in that evaluation drops to a blanket allow.
func handleLeaseFailure(
	ctx context.Context,
	deps CheckDeps,
	req CheckRequest,
	reason string,
	flag *rulesengine.Flag,
	company *rulesengine.Company,
	user *rulesengine.User,
	creditID string,
) *CheckOutcome {
	if !req.FailOpen {
		return staticFailureOutcome(req, reason, flag)
	}

	substituted := substituteCreditBalance(company, creditID, FailOpenBalance)
	result, err := deps.DataStream.EvaluateFlag(ctx, flag, substituted, user, preflightOptions(req))
	if err != nil || result == nil {
		deps.logger().Warn(ctx, fmt.Sprintf("Lease check: the fail-open evaluation failed (%v); allowing", err))
		return staticFailureOutcome(req, reason, flag)
	}
	return &CheckOutcome{
		Allowed:     result.Value,
		Value:       result.Value,
		Reason:      fmt.Sprintf("%s (%s_fail_open)", fallbackString(result.Reason, "evaluated"), reason),
		FlagKey:     fallbackString(result.FlagKey, req.FlagKey),
		FlagID:      firstString(result.FlagID, &flag.ID),
		Entitlement: result.Entitlement,
		Error:       reason,
	}
}

// staticFailureOutcome resolves a mode with no evaluation behind it: deny for
// fail-closed, blanket allow for fail-open. Used when the engine is the thing
// that failed, and when the fail-open evaluation itself errors.
func staticFailureOutcome(req CheckRequest, reason string, flag *rulesengine.Flag) *CheckOutcome {
	outcome := &CheckOutcome{
		Allowed: req.FailOpen,
		Value:   req.FailOpen,
		Reason:  reason,
		FlagKey: req.FlagKey,
		Error:   reason,
	}
	if req.FailOpen {
		outcome.Reason = reason + "_fail_open"
	}
	if flag != nil {
		outcome.FlagID = &flag.ID
	}
	return outcome
}

// undoDebit returns a debit whose reservation record never landed, rather than
// stranding it until lease expiry. Consume claims whatever slice of the add made
// it to the store and refunds it; nothing claimed means nothing landed, so the
// debit is refunded directly. Both are pinned to this lease. If the undo itself
// fails, accept the bounded leak: the slice comes back at lease expiry, which
// beats risking a double refund.
func undoDebit(ctx context.Context, deps CheckDeps, record ReservationRecord, leaseID string) {
	_, claimed, err := deps.Reservations.Consume(ctx, record.ID, 0)
	if err == nil && !claimed {
		err = deps.Leases.Refund(ctx, record.CompanyID, record.CreditTypeID, record.CreditsReserved, leaseID)
	}
	if err != nil {
		deps.logger().Warn(ctx, fmt.Sprintf("Lease check: could not undo the local debit for %s (%v); the slice is reclaimed at lease expiry", record.ID, err))
	}
}

// cancelReservation claims the hold and refunds all of it. Best effort: a
// failure leaves the hold for the sweeper or for lease expiry.
func cancelReservation(ctx context.Context, deps CheckDeps, record ReservationRecord) {
	if _, _, err := deps.Reservations.Consume(ctx, record.ID, 0); err != nil {
		deps.logger().Warn(ctx, fmt.Sprintf("Lease check: failed to cancel reservation %s (%v); its hold is reclaimed by the sweeper or at lease expiry", record.ID, err))
	}
}

// substituteCreditBalance copies the company with one credit balance replaced,
// leaving the cached entity untouched.
func substituteCreditBalance(company *rulesengine.Company, creditID string, balance float64) *rulesengine.Company {
	copied := *company
	balances := make(map[string]float64, len(company.CreditBalances)+1)
	for id, value := range company.CreditBalances {
		balances[id] = value
	}
	balances[creditID] = balance
	copied.CreditBalances = balances
	return &copied
}

// preflightOptions builds the preflight the caller's usage implies. With an
// event subtype the usage goes out scoped to it so the engine matches the
// subtype's condition; without one it goes out as the generic knob.
func preflightOptions(req CheckRequest) *EvalOptions {
	if !IsValidQuantity(req.Usage) {
		return nil
	}
	quantity := preflightQuantity(req.Usage)
	if req.EventSubtype != "" {
		return &EvalOptions{EventUsage: &EventUsage{EventSubtype: req.EventSubtype, Quantity: quantity}}
	}
	return &EvalOptions{Usage: &quantity}
}

// preflightQuantity casts a usage onto the integer the engine's preflight
// carries. A hold can be sized from a fractional usage, but a preflight asks an
// upper-bound question, so a fraction rounds up: the check must not pass on less
// usage than the operation is about to record.
func preflightQuantity(usage float64) int64 {
	return int64(math.Ceil(usage))
}

// EngineCheckFlagOptions translates the flow's preflight into the rules engine's
// own options. The DataStream seam takes engine options, so the client applies
// this at the adapter.
func EngineCheckFlagOptions(opts *EvalOptions) []rulesengine.CheckFlagOption {
	if opts == nil {
		return nil
	}

	engineOpts := make([]rulesengine.CheckFlagOption, 0, 1+len(opts.CreditCost))
	if opts.Usage != nil {
		engineOpts = append(engineOpts, rulesengine.WithUsage(*opts.Usage))
	}
	if opts.EventUsage != nil {
		engineOpts = append(engineOpts, rulesengine.WithEventUsage(opts.EventUsage.EventSubtype, opts.EventUsage.Quantity))
	}
	for creditID, cost := range opts.CreditCost {
		engineOpts = append(engineOpts, rulesengine.WithCreditCost(creditID, cost))
	}
	return engineOpts
}

// flagCheckIDs are the resolved entity ids a flag_check event carries.
type flagCheckIDs struct {
	companyID *string
	userID    *string
	ruleID    *string
}

// emit reports a lease-path resolution and passes the outcome straight through.
// Analytics must never change a verdict the caller is already acting on, so this
// only ever adds an event.
func (d CheckDeps) emit(ctx context.Context, req CheckRequest, outcome *CheckOutcome, ids flagCheckIDs) *CheckOutcome {
	if d.EmitFlagCheck == nil {
		return outcome
	}
	body := &schematicgo.EventBodyFlagCheck{
		FlagKey:    outcome.FlagKey,
		Value:      outcome.Value,
		Reason:     outcome.Reason,
		FlagID:     outcome.FlagID,
		CompanyID:  ids.companyID,
		UserID:     ids.userID,
		RuleID:     ids.ruleID,
		ReqCompany: req.Company,
		ReqUser:    req.User,
	}
	if outcome.Error != "" {
		body.Error = &outcome.Error
	}
	d.EmitFlagCheck(ctx, body)
	return outcome
}

func (d CheckDeps) logger() core.Logger {
	if d.Logger == nil {
		return noopLogger{}
	}
	return d.Logger
}

func (d CheckDeps) now() time.Time {
	return orNow(d.Clock)()
}

func (d CheckDeps) newReservationID() string {
	if d.NewReservationID == nil {
		return uuid.NewString()
	}
	return d.NewReservationID()
}

func entitlementValueType(entitlement *rulesengine.FeatureEntitlement) string {
	if entitlement == nil {
		return "<none>"
	}
	return string(entitlement.ValueType)
}

func userID(user *rulesengine.User) *string {
	if user == nil {
		return nil
	}
	return &user.ID
}

func derefString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func derefFloat(value *float64) float64 {
	if value == nil {
		return 0
	}
	return *value
}

func firstString(value, fallback *string) *string {
	if value != nil && *value != "" {
		return value
	}
	return fallback
}
