package leases

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Runs the language-agnostic conformance vectors against this SDK.
//
// The vectors and the semantics they pin live in conformance/ at the repo root,
// copied verbatim from schematic-node (the reference implementation). This
// runner is the only language-specific piece; every port reimplements it and
// must pass the same vectors, on every store backend it ships.

// ---------------------------------------------------------------------------
// Vector schema (see conformance/SPEC.md; keys are snake_case JSON).

type vectorDoc struct {
	Category string   `json:"category"`
	Vectors  []vector `json:"vectors"`
}

type vector struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Backends    []string    `json:"backends"`
	Given       *givenBlock `json:"given"`
	Operations  []operation `json:"operations"`
}

type givenBlock struct {
	Config *configSpec `json:"config"`
	Leases []leaseSpec `json:"leases"`
}

type configSpec struct {
	LeaseDurationMs  *float64 `json:"lease_duration_ms"`
	ReservationTTLMs *float64 `json:"reservation_ttl_ms"`
	LeaseSize        *float64 `json:"lease_size"`
	LowWaterMark     *float64 `json:"low_water_mark"`
}

type leaseSpec struct {
	LeaseID       string  `json:"lease_id"`
	CompanyID     string  `json:"company_id"`
	CreditTypeID  string  `json:"credit_type_id"`
	GrantedAmount float64 `json:"granted_amount"`
	ExpiresAtMs   float64 `json:"expires_at_ms"`
}

type operation struct {
	Op     string       `json:"op"`
	Expect expectations `json:"expect"`

	// advance_clock
	MS float64 `json:"ms"`

	// lease and reservation store ops
	LeaseID           string   `json:"lease_id"`
	CompanyID         string   `json:"company_id"`
	CreditTypeID      string   `json:"credit_type_id"`
	GrantedAmount     float64  `json:"granted_amount"`
	ExpiresAtMs       *float64 `json:"expires_at_ms"`
	Credits           float64  `json:"credits"`
	PinLeaseID        string   `json:"pin_lease_id"`
	GrantedTotal      float64  `json:"granted_total"`
	ID                string   `json:"id"`
	Handle            string   `json:"handle"`
	EventSubtype      string   `json:"event_subtype"`
	QuantityReserved  float64  `json:"quantity_reserved"`
	CreditsReserved   float64  `json:"credits_reserved"`
	ConsumptionRate   float64  `json:"consumption_rate"`
	CrashBeforeRefund bool     `json:"crash_before_refund"`

	// manager ops
	RequiredCredits   *float64        `json:"required_credits"`
	Server            json.RawMessage `json:"server"`
	InstallDuringWire *leaseSpec      `json:"install_during_wire"`

	// flow ops, read by the check/track handlers
	FlagKey           string          `json:"flag_key"`
	Company           json.RawMessage `json:"company"`
	Usage             *float64        `json:"usage"`
	OnAcquireFailure  string          `json:"on_acquire_failure"`
	Engine            json.RawMessage `json:"engine"`
	SaveReservationAs string          `json:"save_reservation_as"`
	ActualQuantity    *float64        `json:"actual_quantity"`
}

// expectations keeps assertions as raw JSON so a key that is present and null
// (the refused result) stays distinguishable from a key that is absent.
type expectations map[string]json.RawMessage

func (e expectations) has(key string) bool {
	_, ok := e[key]
	return ok
}

func (e expectations) isNull(key string) bool {
	return string(e[key]) == "null"
}

func (e expectations) number(t *testing.T, key string) float64 {
	t.Helper()
	var value float64
	require.NoError(t, json.Unmarshal(e[key], &value))
	return value
}

func (e expectations) boolean(t *testing.T, key string) bool {
	t.Helper()
	var value bool
	require.NoError(t, json.Unmarshal(e[key], &value))
	return value
}

func (e expectations) str(t *testing.T, key string) string {
	t.Helper()
	var value string
	require.NoError(t, json.Unmarshal(e[key], &value))
	return value
}

func (e expectations) strSlice(t *testing.T, key string) []string {
	t.Helper()
	var value []string
	require.NoError(t, json.Unmarshal(e[key], &value))
	return value
}

// ---------------------------------------------------------------------------
// Harness

// harness carries one vector's stores, manager, clock, and reservation handles.
type harness struct {
	t            *testing.T
	ctx          context.Context
	clock        *virtualClock
	leases       LeaseStore
	reservations ReservationStore
	crash        *crashingRefund
	wire         *scriptedWireClient
	manager      *LeaseManager
	handles      map[string]ReservationRecord
}

// flowOpHandler runs a flow-level conformance op against the harness. The
// check and track ops register here from flow_conformance_test.go.
type flowOpHandler func(h *harness, op operation)

var flowOps = map[string]flowOpHandler{}

// resolveReservationID maps a vector's `handle` back to the id the check that
// issued it returned, or takes the literal `id`.
func (h *harness) resolveReservationID(op operation) string {
	h.t.Helper()
	if op.Handle != "" {
		reservation, ok := h.handles[op.Handle]
		require.Truef(h.t, ok, "unknown reservation handle: %s", op.Handle)
		return reservation.ID
	}
	require.NotEmptyf(h.t, op.ID, "op %s needs an id or handle", op.Op)
	return op.ID
}

func TestConformanceVectors(t *testing.T) {
	documents := loadVectorDocs(t)
	for _, factory := range backendFactories {
		t.Run(factory.name, func(t *testing.T) {
			for _, doc := range documents {
				t.Run(doc.Category, func(t *testing.T) {
					for _, vec := range doc.Vectors {
						if len(vec.Backends) > 0 && !slices.Contains(vec.Backends, factory.name) {
							continue
						}
						t.Run(vec.Name, func(t *testing.T) {
							runVector(t, factory.make(t), vec)
						})
					}
				})
			}
		})
	}
}

func loadVectorDocs(t *testing.T) []vectorDoc {
	t.Helper()
	paths, err := filepath.Glob(filepath.Join("..", "conformance", "vectors", "*.json"))
	require.NoError(t, err)
	require.NotEmpty(t, paths, "no conformance vectors found")
	slices.Sort(paths)

	documents := make([]vectorDoc, 0, len(paths))
	for _, path := range paths {
		raw, err := os.ReadFile(path) //nolint:gosec // fixed in-repo path
		require.NoError(t, err)
		var doc vectorDoc
		require.NoErrorf(t, json.Unmarshal(raw, &doc), "parsing %s", path)
		documents = append(documents, doc)
	}
	return documents
}

func runVector(t *testing.T, b *backend, vec vector) {
	t.Helper()
	config := configSpec{}
	if vec.Given != nil && vec.Given.Config != nil {
		config = *vec.Given.Config
	}
	wire := &scriptedWireClient{clock: b.clock}
	manager := NewLeaseManager(wire, b.leases, LeaseManagerOptions{
		Reservations: b.reservations,
		Config: ResolvedLeaseConfig{
			LeaseDuration:  durationFromMs(config.LeaseDurationMs),
			ReservationTTL: durationFromMs(config.ReservationTTLMs),
			LeaseSize:      floatOrZero(config.LeaseSize),
			LowWaterMark:   floatOrZero(config.LowWaterMark),
		},
		Clock: b.clock.Now,
	})
	t.Cleanup(manager.Stop)

	h := &harness{
		t:            t,
		ctx:          context.Background(),
		clock:        b.clock,
		leases:       b.leases,
		reservations: b.reservations,
		crash:        b.crash,
		wire:         wire,
		manager:      manager,
		handles:      make(map[string]ReservationRecord),
	}

	if vec.Given != nil {
		for _, lease := range vec.Given.Leases {
			installLease(t, b.leases, b.clock, lease.LeaseID, lease.CompanyID, lease.CreditTypeID, lease.GrantedAmount, lease.ExpiresAtMs)
		}
	}
	for _, op := range vec.Operations {
		runOperation(h, op)
	}
}

func runOperation(h *harness, op operation) {
	h.t.Helper()
	if handler, ok := storeOps[op.Op]; ok {
		handler(h, op)
		return
	}
	if handler, ok := flowOps[op.Op]; ok {
		handler(h, op)
		return
	}
	h.t.Fatalf("unknown conformance op: %s", op.Op)
}

// storeOps covers every store-level and manager-level op in the SPEC's tables.
var storeOps = map[string]flowOpHandler{
	"advance_clock":            opAdvanceClock,
	"replace_lease":            opReplaceLease,
	"drop_lease":               opDropLease,
	"try_reserve":              opTryReserve,
	"refund_lease":             opRefundLease,
	"extend_lease":             opExtendLease,
	"get_lease":                opGetLease,
	"add_reservation":          opAddReservation,
	"consume_reservation":      opConsumeReservation,
	"get_reservation":          opGetReservation,
	"reserved_credits":         opReservedCredits,
	"reservation_count":        opReservationCount,
	"sweep_expired":            opSweepExpired,
	"acquire_if_needed":        opAcquireIfNeeded,
	"maybe_extend":             opMaybeExtend,
	"release_all_local_leases": opReleaseAllLocalLeases,
}

func opAdvanceClock(h *harness, op operation) {
	h.clock.advance(time.Duration(op.MS) * time.Millisecond)
}

func opReplaceLease(h *harness, op operation) {
	h.t.Helper()
	wrote, err := h.leases.Replace(h.ctx, LeaseGrant{
		LeaseID:       op.LeaseID,
		CompanyID:     op.CompanyID,
		CreditTypeID:  op.CreditTypeID,
		GrantedAmount: op.GrantedAmount,
		ExpiresAt:     h.clock.at(*op.ExpiresAtMs),
	})
	require.NoError(h.t, err)
	if op.Expect.has("written") {
		assert.Equal(h.t, op.Expect.boolean(h.t, "written"), wrote)
	}
}

func opDropLease(h *harness, op operation) {
	h.t.Helper()
	require.NoError(h.t, h.leases.Drop(h.ctx, op.CompanyID, op.CreditTypeID))
}

func opTryReserve(h *harness, op operation) {
	h.t.Helper()
	balance, ok, err := h.leases.TryReserve(h.ctx, op.CompanyID, op.CreditTypeID, op.Credits)
	require.NoError(h.t, err)
	assertNullableNumber(h.t, op.Expect, "balance", balance, ok)
}

func opRefundLease(h *harness, op operation) {
	h.t.Helper()
	require.NoError(h.t, h.leases.Refund(h.ctx, op.CompanyID, op.CreditTypeID, op.Credits, op.PinLeaseID))
}

func opExtendLease(h *harness, op operation) {
	h.t.Helper()
	var expiresAt *time.Time
	if op.ExpiresAtMs != nil {
		at := h.clock.at(*op.ExpiresAtMs)
		expiresAt = &at
	}
	require.NoError(h.t, h.leases.Extend(h.ctx, op.CompanyID, op.CreditTypeID, op.GrantedTotal, expiresAt, op.PinLeaseID))
}

func opGetLease(h *harness, op operation) {
	h.t.Helper()
	entry, err := h.leases.Get(h.ctx, op.CompanyID, op.CreditTypeID)
	require.NoError(h.t, err)
	if op.Expect.has("exists") {
		assert.Equal(h.t, op.Expect.boolean(h.t, "exists"), entry != nil)
	}
	if op.Expect.has("lease_id") {
		var got string
		if entry != nil {
			got = entry.LeaseID
		}
		if op.Expect.isNull("lease_id") {
			assert.Nil(h.t, entry)
		} else {
			assert.Equal(h.t, op.Expect.str(h.t, "lease_id"), got)
		}
	}
	if op.Expect.has("granted_amount") {
		require.NotNil(h.t, entry)
		assert.Equal(h.t, op.Expect.number(h.t, "granted_amount"), entry.GrantedAmount)
	}
	if op.Expect.has("local_remaining_credits") {
		require.NotNil(h.t, entry)
		assert.Equal(h.t, op.Expect.number(h.t, "local_remaining_credits"), entry.LocalRemainingCredits)
	}
}

func opAddReservation(h *harness, op operation) {
	h.t.Helper()
	require.NoError(h.t, h.reservations.Add(h.ctx, ReservationRecord{
		ID:               op.ID,
		LeaseID:          op.LeaseID,
		CompanyID:        op.CompanyID,
		CreditTypeID:     op.CreditTypeID,
		EventSubtype:     op.EventSubtype,
		QuantityReserved: op.QuantityReserved,
		CreditsReserved:  op.CreditsReserved,
		ConsumptionRate:  op.ConsumptionRate,
		ExpiresAt:        h.clock.at(*op.ExpiresAtMs),
		Company:          map[string]string{"id": op.CompanyID},
	}))
}

func opConsumeReservation(h *harness, op operation) {
	h.t.Helper()
	id := h.resolveReservationID(op)
	if op.CrashBeforeRefund {
		h.crash.arm()
		_, _, err := h.reservations.Consume(h.ctx, id, op.Credits)
		require.ErrorIs(h.t, err, errSimulatedCrash)
		assert.True(h.t, op.Expect.boolean(h.t, "throws"))
		return
	}
	consumed, claimed, err := h.reservations.Consume(h.ctx, id, op.Credits)
	require.NoError(h.t, err)
	assertNullableNumber(h.t, op.Expect, "consumed", consumed, claimed)
}

func opGetReservation(h *harness, op operation) {
	h.t.Helper()
	reservation, err := h.reservations.Get(h.ctx, h.resolveReservationID(op))
	require.NoError(h.t, err)
	if op.Expect.has("exists") {
		assert.Equal(h.t, op.Expect.boolean(h.t, "exists"), reservation != nil)
	}
}

func opReservedCredits(h *harness, op operation) {
	h.t.Helper()
	total, err := h.reservations.ReservedCredits(h.ctx, op.CompanyID, op.CreditTypeID)
	require.NoError(h.t, err)
	assert.Equal(h.t, op.Expect.number(h.t, "total"), total)
}

func opReservationCount(h *harness, op operation) {
	h.t.Helper()
	count, err := h.reservations.Count(h.ctx)
	require.NoError(h.t, err)
	assert.Equal(h.t, int(op.Expect.number(h.t, "count")), count)
}

func opSweepExpired(h *harness, op operation) {
	h.t.Helper()
	swept, err := h.reservations.SweepExpired(h.ctx)
	require.NoError(h.t, err)
	if op.Expect.has("swept") {
		assert.Equal(h.t, int(op.Expect.number(h.t, "swept")), swept)
	}
}

func opAcquireIfNeeded(h *harness, op operation) {
	h.t.Helper()
	if len(op.Server) > 0 {
		h.wire.queueAcquire(decodeServerScript(h.t, op.Server))
	}
	if install := op.InstallDuringWire; install != nil {
		h.wire.duringAcquire = func() {
			_, err := h.leases.Replace(h.ctx, LeaseGrant{
				LeaseID:       install.LeaseID,
				CompanyID:     install.CompanyID,
				CreditTypeID:  install.CreditTypeID,
				GrantedAmount: install.GrantedAmount,
				ExpiresAt:     h.clock.at(install.ExpiresAtMs),
			})
			assert.NoError(h.t, err)
		}
	}
	entry := h.manager.AcquireIfNeeded(h.ctx, op.CompanyID, op.CreditTypeID)
	h.manager.drainBackground()

	acquires, _, releases := h.wire.snapshot()
	if op.Expect.has("lease_id") {
		if op.Expect.isNull("lease_id") {
			assert.Nil(h.t, entry)
		} else {
			require.NotNil(h.t, entry)
			assert.Equal(h.t, op.Expect.str(h.t, "lease_id"), entry.LeaseID)
		}
	}
	if op.Expect.has("wire_acquires") {
		assert.Len(h.t, acquires, int(op.Expect.number(h.t, "wire_acquires")))
	}
	if op.Expect.has("last_acquire_requested_amount") {
		require.NotEmpty(h.t, acquires)
		assert.Equal(h.t, op.Expect.number(h.t, "last_acquire_requested_amount"), acquires[len(acquires)-1].requestedAmount)
	}
	if op.Expect.has("released_lease_ids") {
		assert.Equal(h.t, op.Expect.strSlice(h.t, "released_lease_ids"), orEmpty(releases))
	}
}

func opMaybeExtend(h *harness, op operation) {
	h.t.Helper()
	if len(op.Server) > 0 {
		h.wire.queueExtend(decodeServerScript(h.t, op.Server))
	}
	h.manager.MaybeExtend(h.ctx, op.CompanyID, op.CreditTypeID, op.RequiredCredits)
	h.manager.drainBackground()

	_, extends, _ := h.wire.snapshot()
	if op.Expect.has("wire_extends") {
		assert.Len(h.t, extends, int(op.Expect.number(h.t, "wire_extends")))
	}
	if op.Expect.has("last_extend_additional_amount") {
		require.NotEmpty(h.t, extends)
		assert.Equal(h.t, op.Expect.number(h.t, "last_extend_additional_amount"), extends[len(extends)-1].additionalAmount)
	}
	if op.Expect.has("last_extend_lease_id") {
		require.NotEmpty(h.t, extends)
		assert.Equal(h.t, op.Expect.str(h.t, "last_extend_lease_id"), extends[len(extends)-1].leaseID)
	}
}

func opReleaseAllLocalLeases(h *harness, op operation) {
	h.t.Helper()
	h.manager.ReleaseAllLocalLeases(h.ctx)
	_, _, releases := h.wire.snapshot()
	if op.Expect.has("released_lease_ids") {
		assert.Equal(h.t, op.Expect.strSlice(h.t, "released_lease_ids"), orEmpty(releases))
	}
	if op.Expect.has("remaining_slots") {
		lister, ok := h.leases.(LeaseLister)
		require.True(h.t, ok, "remaining_slots needs an enumerable store")
		remaining, err := lister.List(h.ctx)
		require.NoError(h.t, err)
		assert.Len(h.t, remaining, int(op.Expect.number(h.t, "remaining_slots")))
	}
}

// assertNullableNumber compares a result that may be the refused outcome, which
// the vectors write as JSON null and never as a figure.
func assertNullableNumber(t *testing.T, expect expectations, key string, value float64, ok bool) {
	t.Helper()
	if !expect.has(key) {
		return
	}
	if expect.isNull(key) {
		assert.Falsef(t, ok, "expected %s to be refused, got %v", key, value)
		return
	}
	require.Truef(t, ok, "expected %s to be %s, got the refused result", key, expect[key])
	assert.Equal(t, expect.number(t, key), value)
}

func decodeServerScript(t *testing.T, raw json.RawMessage) serverScript {
	t.Helper()
	var script serverScript
	require.NoError(t, json.Unmarshal(raw, &script))
	return script
}

func durationFromMs(ms *float64) time.Duration {
	if ms == nil {
		return 0
	}
	return time.Duration(*ms) * time.Millisecond
}

func floatOrZero(value *float64) float64 {
	if value == nil {
		return 0
	}
	return *value
}

func orEmpty(values []string) []string {
	if values == nil {
		return []string{}
	}
	return values
}
