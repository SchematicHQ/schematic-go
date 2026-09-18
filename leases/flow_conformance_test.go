package leases

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/schematichq/schematic-go/rulesengine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The flow-level half of the conformance runner: the vectors' `check` and
// `track` ops, driven against the real check and settle flows with a scripted
// rules engine standing in for the WASM. The engine is an oracle here, as
// conformance/SPEC.md says: the vectors pin the orchestration around it, not the
// engine, which is shared across the SDKs and has its own tests.

func init() {
	flowOps["check"] = opCheck
	flowOps["track"] = opTrack
}

// ---------------------------------------------------------------------------
// Vector schema for the flow ops.

type companySpec struct {
	ID             string             `json:"id"`
	CreditBalances map[string]float64 `json:"credit_balances"`
}

type engineEntitlementSpec struct {
	ValueType       string   `json:"value_type"`
	CreditID        *string  `json:"credit_id"`
	ConsumptionRate *float64 `json:"consumption_rate"`
	EventSubtype    *string  `json:"event_subtype"`
	FeatureID       string   `json:"feature_id"`
	FeatureKey      string   `json:"feature_key"`
}

type engineResultSpec struct {
	Value       bool                   `json:"value"`
	Reason      string                 `json:"reason"`
	Entitlement *engineEntitlementSpec `json:"entitlement"`
}

// flowServerScript is the check op's `server` block, which names the wire call
// it scripts rather than scripting one directly the way a manager op does.
type flowServerScript struct {
	Acquire *serverScript `json:"acquire"`
	Extend  *serverScript `json:"extend"`
}

type reservationExpect struct {
	LeaseID          *string  `json:"lease_id"`
	CreditTypeID     *string  `json:"credit_type_id"`
	EventSubtype     *string  `json:"event_subtype"`
	QuantityReserved *float64 `json:"quantity_reserved"`
	CreditsReserved  *float64 `json:"credits_reserved"`
	ConsumptionRate  *float64 `json:"consumption_rate"`
}

type eventUsageExpect struct {
	EventSubtype string `json:"event_subtype"`
	Quantity     int64  `json:"quantity"`
}

type engineCallExpect struct {
	// CreditBalance is a number, or the string "max_safe_integer" for the
	// fail-open substitution.
	CreditBalance json.RawMessage   `json:"credit_balance"`
	CreditCost    *float64          `json:"credit_cost"`
	EventUsage    *eventUsageExpect `json:"event_usage"`
	Usage         *int64            `json:"usage"`
}

type trackExpect struct {
	Event *string `json:"event"`
	// Quantity is the usage the vector declares, which a reference SDK with a
	// numeric event quantity bills as-is. This SDK's track event carries an
	// int64, so a fractional declaration bills as the whole unit it settles.
	Quantity *float64 `json:"quantity"`
	LeaseID  *string  `json:"lease_id"`
}

// ---------------------------------------------------------------------------
// Scripted engine and datastream

// engineCall is one evaluation the flow asked for, kept so a vector can assert
// on what the flow put in front of the engine.
type engineCall struct {
	creditBalances map[string]float64
	options        *EvalOptions
}

// scriptedDataStream serves one flag and one company from "cache" and answers
// each evaluation with the next scripted result, in call order.
type scriptedDataStream struct {
	t       *testing.T
	flagKey string
	flag    *rulesengine.Flag
	company *rulesengine.Company
	results []engineResultSpec
	calls   []engineCall
}

func (d *scriptedDataStream) GetFlag(context.Context, string) (*rulesengine.Flag, bool) {
	return d.flag, true
}

func (d *scriptedDataStream) GetCompany(context.Context, map[string]string) (*rulesengine.Company, error) {
	return d.company, nil
}

func (d *scriptedDataStream) GetUser(context.Context, map[string]string) (*rulesengine.User, error) {
	return nil, nil
}

func (d *scriptedDataStream) EvaluateFlag(
	_ context.Context,
	_ *rulesengine.Flag,
	company *rulesengine.Company,
	_ *rulesengine.User,
	opts *EvalOptions,
) (*rulesengine.CheckFlagResult, error) {
	balances := map[string]float64{}
	if company != nil {
		balances = company.CreditBalances
	}
	d.calls = append(d.calls, engineCall{creditBalances: balances, options: opts})

	if len(d.results) == 0 {
		return nil, fmt.Errorf("unscripted engine call in check op for flag %s", d.flagKey)
	}
	scripted := d.results[0]
	d.results = d.results[1:]

	flagID := "flag_1"
	return &rulesengine.CheckFlagResult{
		Value:       scripted.Value,
		Reason:      scripted.Reason,
		FlagKey:     d.flagKey,
		FlagID:      &flagID,
		Entitlement: entitlementFromSpec(scripted.Entitlement, d.flagKey),
	}, nil
}

func entitlementFromSpec(spec *engineEntitlementSpec, flagKey string) *rulesengine.FeatureEntitlement {
	if spec == nil {
		return nil
	}
	return &rulesengine.FeatureEntitlement{
		FeatureID:       fallbackString(spec.FeatureID, "feat_1"),
		FeatureKey:      fallbackString(spec.FeatureKey, flagKey),
		ValueType:       rulesengine.EntitlementValueType(spec.ValueType),
		CreditID:        spec.CreditID,
		ConsumptionRate: spec.ConsumptionRate,
		EventSubtype:    spec.EventSubtype,
	}
}

// ---------------------------------------------------------------------------
// Ops

func opCheck(h *harness, op operation) {
	h.t.Helper()

	flagKey := fallbackString(op.FlagKey, "flag")
	spec := companySpec{ID: "co_1"}
	if len(op.Company) > 0 {
		require.NoError(h.t, json.Unmarshal(op.Company, &spec))
	}
	var results []engineResultSpec
	if len(op.Engine) > 0 {
		require.NoError(h.t, json.Unmarshal(op.Engine, &results))
	}
	if len(op.Server) > 0 {
		var script flowServerScript
		require.NoError(h.t, json.Unmarshal(op.Server, &script))
		if script.Acquire != nil {
			h.wire.queueAcquire(*script.Acquire)
		}
		if script.Extend != nil {
			h.wire.queueExtend(*script.Extend)
		}
	}

	datastream := &scriptedDataStream{
		t:       h.t,
		flagKey: flagKey,
		flag:    &rulesengine.Flag{ID: "flag_1", Key: flagKey},
		company: &rulesengine.Company{ID: spec.ID, CreditBalances: spec.CreditBalances},
		results: results,
	}

	fellBack := false
	fallback := func(context.Context) *CheckOutcome {
		fellBack = true
		return &CheckOutcome{Allowed: true, Value: true, Reason: "fallback", FlagKey: flagKey}
	}

	usage := 0.0
	if op.Usage != nil {
		usage = *op.Usage
	}
	outcome := CheckWithLease(h.ctx, CheckDeps{
		DataStream:   datastream,
		Leases:       h.leases,
		Reservations: h.reservations,
		Manager:      h.manager,
		Clock:        h.clock.Now,
	}, CheckRequest{
		FlagKey:      flagKey,
		Company:      map[string]string{"id": spec.ID},
		Usage:        usage,
		EventSubtype: op.EventSubtype,
		FailOpen:     op.OnAcquireFailure == "fail-open",
	}, fallback)
	h.manager.drainBackground()

	assertCheckOutcome(h, op, outcome, fellBack)
	assertEngineCalls(h.t, op, datastream.calls, spec)

	_, extends, _ := h.wire.snapshot()
	if op.Expect.has("wire_extends") {
		assert.Len(h.t, extends, int(op.Expect.number(h.t, "wire_extends")))
	}
	if op.Expect.has("last_extend_additional_amount") {
		require.NotEmpty(h.t, extends)
		assert.Equal(h.t, op.Expect.number(h.t, "last_extend_additional_amount"), extends[len(extends)-1].additionalAmount)
	}

	if op.SaveReservationAs != "" && outcome.Reservation != nil {
		h.handles[op.SaveReservationAs] = *outcome.Reservation
	}
}

func assertCheckOutcome(h *harness, op operation, outcome *CheckOutcome, fellBack bool) {
	h.t.Helper()

	if op.Expect.has("allowed") {
		assert.Equal(h.t, op.Expect.boolean(h.t, "allowed"), outcome.Allowed)
	}
	if op.Expect.has("reason") {
		assert.Equal(h.t, op.Expect.str(h.t, "reason"), outcome.Reason)
	}
	if op.Expect.has("err") {
		assert.Equal(h.t, op.Expect.str(h.t, "err"), outcome.Error)
	}
	if op.Expect.has("has_reservation") {
		assert.Equal(h.t, op.Expect.boolean(h.t, "has_reservation"), outcome.Reservation != nil)
	}
	if op.Expect.has("fallback_called") {
		assert.Equal(h.t, op.Expect.boolean(h.t, "fallback_called"), fellBack)
	}
	if !op.Expect.has("reservation") {
		return
	}

	var want reservationExpect
	require.NoError(h.t, json.Unmarshal(op.Expect["reservation"], &want))
	record := outcome.Reservation
	require.NotNil(h.t, record)
	if want.LeaseID != nil {
		assert.Equal(h.t, *want.LeaseID, record.LeaseID)
	}
	if want.CreditTypeID != nil {
		assert.Equal(h.t, *want.CreditTypeID, record.CreditTypeID)
	}
	if want.EventSubtype != nil {
		assert.Equal(h.t, *want.EventSubtype, record.EventSubtype)
	}
	if want.QuantityReserved != nil {
		assert.Equal(h.t, *want.QuantityReserved, record.QuantityReserved)
	}
	if want.CreditsReserved != nil {
		assert.Equal(h.t, *want.CreditsReserved, record.CreditsReserved)
	}
	if want.ConsumptionRate != nil {
		assert.Equal(h.t, *want.ConsumptionRate, record.ConsumptionRate)
	}
}

func assertEngineCalls(t *testing.T, op operation, calls []engineCall, spec companySpec) {
	t.Helper()
	if !op.Expect.has("engine_calls") {
		return
	}

	var want []engineCallExpect
	require.NoError(t, json.Unmarshal(op.Expect["engine_calls"], &want))
	require.Len(t, calls, len(want))

	creditID := creditIDFromVector(t, op, spec)
	for i, expected := range want {
		got := calls[i]
		if len(expected.CreditBalance) > 0 {
			require.NotEmptyf(t, creditID, "engine_calls needs a credit id to assert a balance against")
			assert.Equal(t, expectedCreditBalance(t, expected.CreditBalance), got.creditBalances[creditID])
		}
		if expected.CreditCost != nil {
			require.NotNil(t, got.options)
			assert.Equal(t, *expected.CreditCost, got.options.CreditCost[creditID])
		}
		if expected.EventUsage != nil {
			require.NotNil(t, got.options)
			require.NotNil(t, got.options.EventUsage)
			assert.Equal(t, expected.EventUsage.EventSubtype, got.options.EventUsage.EventSubtype)
			assert.Equal(t, expected.EventUsage.Quantity, got.options.EventUsage.Quantity)
		}
		if expected.Usage != nil {
			require.NotNil(t, got.options)
			require.NotNil(t, got.options.Usage)
			assert.Equal(t, *expected.Usage, *got.options.Usage)
		}
	}
}

// expectedCreditBalance reads a balance expectation, which the vectors write as
// a number or as the name of Node's MAX_SAFE_INTEGER.
func expectedCreditBalance(t *testing.T, raw json.RawMessage) float64 {
	t.Helper()
	if string(raw) == `"max_safe_integer"` {
		return FailOpenBalance
	}
	var balance float64
	require.NoError(t, json.Unmarshal(raw, &balance))
	return balance
}

// creditIDFromVector names the credit a vector's balance and cost expectations
// are about: the one the scripted entitlement meters, or the company's only
// balance when no entitlement names one.
func creditIDFromVector(t *testing.T, op operation, spec companySpec) string {
	t.Helper()
	var results []engineResultSpec
	if len(op.Engine) > 0 {
		require.NoError(t, json.Unmarshal(op.Engine, &results))
	}
	for _, result := range results {
		if result.Entitlement != nil && result.Entitlement.CreditID != nil {
			return *result.Entitlement.CreditID
		}
	}
	for creditID := range spec.CreditBalances {
		return creditID
	}
	return ""
}

func opTrack(h *harness, op operation) {
	h.t.Helper()

	record, ok := h.handles[op.Handle]
	require.Truef(h.t, ok, "unknown reservation handle: %s", op.Handle)
	require.NotNilf(h.t, op.ActualQuantity, "track op needs an actual_quantity")

	outcome := SettleReservation(h.ctx, h.reservations, record, *op.ActualQuantity)
	require.NoError(h.t, outcome.Err)
	if op.Expect.has("settled_locally") {
		assert.Equal(h.t, op.Expect.boolean(h.t, "settled_locally"), outcome.SettledLocally)
	}
	if !op.Expect.has("track") {
		return
	}

	var want trackExpect
	require.NoError(h.t, json.Unmarshal(op.Expect["track"], &want))
	if want.Event != nil {
		assert.Equal(h.t, *want.Event, outcome.Track.Event)
	}
	if want.Quantity != nil {
		require.NotNil(h.t, outcome.Track.Quantity)
		assert.Equal(h.t, SettleQuantity(*want.Quantity), *outcome.Track.Quantity)
	}
	if want.LeaseID != nil {
		require.NotNil(h.t, outcome.Track.LeaseID)
		assert.Equal(h.t, *want.LeaseID, *outcome.Track.LeaseID)
	}
}
