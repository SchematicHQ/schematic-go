package leases

import (
	"context"
	"testing"

	"github.com/schematichq/schematic-go/rulesengine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The rest of this package's flow tests script the rules engine, so the contract
// that matters most is never exercised against the real thing: resolving the
// matched credit entitlement from the probe, substituting the lease balance into
// the company's credit balances, and letting the engine's credit_cost gate
// decide. These run the bundled WebAssembly end to end, so a drift in the option
// envelope or the entity shape fails here rather than mis-gating in production.

const wasmFlagKey = "infer"

func wasmCreditCondition() *rulesengine.Condition {
	creditID := testCreditID
	rate := 1.0
	subtype := testSubtype
	return &rulesengine.Condition{
		ID:              "cond_credit",
		AccountID:       "acct",
		EnvironmentID:   "env",
		ConditionType:   rulesengine.ConditionTypeCredit,
		Operator:        rulesengine.ComparableOperatorLt,
		CreditID:        &creditID,
		ConsumptionRate: &rate,
		EventSubtype:    &subtype,
	}
}

// wasmCompanyCondition is a membership condition a test can point at some other
// company, to make the engine deny for a reason that has nothing to do with the
// credit balance.
func wasmCompanyCondition(companyIDs ...string) *rulesengine.Condition {
	return &rulesengine.Condition{
		ID:            "cond_company",
		AccountID:     "acct",
		EnvironmentID: "env",
		ConditionType: rulesengine.ConditionTypeCompany,
		Operator:      rulesengine.ComparableOperatorEquals,
		ResourceIDs:   companyIDs,
	}
}

func wasmCreditFlag(extra ...*rulesengine.Condition) *rulesengine.Flag {
	conditions := append(rulesengine.JSONSlice[*rulesengine.Condition]{wasmCreditCondition()}, extra...)
	return &rulesengine.Flag{
		ID:            "flag_infer",
		AccountID:     "acct",
		EnvironmentID: "env",
		Key:           wasmFlagKey,
		DefaultValue:  false,
		Rules: rulesengine.JSONSlice[*rulesengine.Rule]{
			{
				ID:            "rule_credit",
				AccountID:     "acct",
				EnvironmentID: "env",
				RuleType:      rulesengine.RuleTypePlanEntitlement,
				Name:          "Credit",
				Priority:      100,
				Value:         true,
				Conditions:    conditions,
			},
		},
	}
}

// wasmCompany carries the company's resolved entitlement for the feature, which
// is what the probe reads the credit id, rate and subtype off.
func wasmCompany(balance float64, entitlements ...*rulesengine.FeatureEntitlement) *rulesengine.Company {
	if len(entitlements) == 0 {
		creditID := testCreditID
		rate := 1.0
		subtype := testSubtype
		entitlements = []*rulesengine.FeatureEntitlement{{
			FeatureID:       "feat_infer",
			FeatureKey:      wasmFlagKey,
			ValueType:       rulesengine.EntitlementValueTypeCredit,
			CreditID:        &creditID,
			ConsumptionRate: &rate,
			EventSubtype:    &subtype,
			CreditTotal:     &balance,
			CreditRemaining: &balance,
		}}
	}
	return &rulesengine.Company{
		ID:             "co_1",
		AccountID:      "acct",
		EnvironmentID:  "env",
		CreditBalances: map[string]float64{testCreditID: balance},
		Entitlements:   entitlements,
		Metrics:        make(rulesengine.CompanyMetricCollection, 0),
		Traits:         make(rulesengine.JSONSlice[*rulesengine.Trait], 0),
	}
}

// engineDataStream serves fixed fixtures and the real rules engine.
type engineDataStream struct {
	engine  *rulesengine.Engine
	flag    *rulesengine.Flag
	company *rulesengine.Company
}

func (d *engineDataStream) GetFlag(context.Context, string) (*rulesengine.Flag, bool) {
	return d.flag, true
}

func (d *engineDataStream) GetCompany(context.Context, map[string]string) (*rulesengine.Company, error) {
	return d.company, nil
}

func (d *engineDataStream) GetUser(context.Context, map[string]string) (*rulesengine.User, error) {
	return nil, nil
}

func (d *engineDataStream) EvaluateFlag(
	ctx context.Context,
	flag *rulesengine.Flag,
	company *rulesengine.Company,
	user *rulesengine.User,
	opts *EvalOptions,
) (*rulesengine.CheckFlagResult, error) {
	return d.engine.CheckFlag(ctx, company, user, flag, EngineCheckFlagOptions(opts)...)
}

func newWasmEngine(t *testing.T) *rulesengine.Engine {
	t.Helper()
	engine, err := rulesengine.NewEngine(context.Background())
	require.NoError(t, err)
	t.Cleanup(func() { _ = engine.Close(context.Background()) })
	return engine
}

// newWasmFixture wires the real engine into a flow fixture. It reuses
// checkFixture's stores and manager, swapping in the engine-backed datastream.
func newWasmFixture(t *testing.T, flag *rulesengine.Flag, company *rulesengine.Company) *checkFixture {
	t.Helper()
	return newCheckFixture(t, &engineDataStream{engine: newWasmEngine(t), flag: flag, company: company})
}

func TestWasmGateSubstitutesTheLeaseBalanceAndIssuesAHold(t *testing.T) {
	fixture := newWasmFixture(t, wasmCreditFlag(), wasmCompany(100))
	fixture.installTestLease(t, "co_1", 10000)

	request := baseRequest()
	request.Usage = 50
	outcome := fixture.run(t, request)

	require.False(t, fixture.fellBack)
	assert.True(t, outcome.Allowed)
	require.NotNil(t, outcome.Reservation)
	assert.Equal(t, testCreditID, outcome.Reservation.CreditTypeID)
	assert.Equal(t, 50.0, outcome.Reservation.CreditsReserved)

	entry, err := fixture.leases.Get(context.Background(), "co_1", testCreditID)
	require.NoError(t, err)
	assert.Equal(t, 9950.0, entry.LocalRemainingCredits, "the hold stays debited from the lease's local view")
	count, err := fixture.store.Count(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestWasmGateCancelsTheHoldWhenTheRuleDeniesForANonCreditReason(t *testing.T) {
	// The balance is plentiful, but the membership condition excludes this
	// company, so the engine denies. The hold taken before the gate must go back.
	flag := wasmCreditFlag(wasmCompanyCondition("co_other"))
	fixture := newWasmFixture(t, flag, wasmCompany(10000))
	fixture.installTestLease(t, "co_1", 10000)

	request := baseRequest()
	request.Usage = 50
	outcome := fixture.run(t, request)

	require.False(t, fixture.fellBack)
	assert.False(t, outcome.Allowed)
	assert.Nil(t, outcome.Reservation)

	entry, err := fixture.leases.Get(context.Background(), "co_1", testCreditID)
	require.NoError(t, err)
	assert.Equal(t, 10000.0, entry.LocalRemainingCredits)
	count, err := fixture.store.Count(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 0, count)
}

func TestWasmGateSkipsTheLeaseWhenTheMatchedEntitlementIsNotCreditMetered(t *testing.T) {
	// A company override grants the feature outright, so the company's effective
	// entitlement is boolean. The probe surfaces that, and the flow defers to the
	// plain check rather than paying a reserve it would only have to cancel.
	flag := wasmCreditFlag()
	flag.Rules = append(rulesengine.JSONSlice[*rulesengine.Rule]{{
		ID:            "rule_override",
		AccountID:     "acct",
		EnvironmentID: "env",
		RuleType:      rulesengine.RuleTypeCompanyOverride,
		Name:          "Override",
		Priority:      1,
		Value:         true,
		Conditions:    rulesengine.JSONSlice[*rulesengine.Condition]{wasmCompanyCondition("co_1")},
	}}, flag.Rules...)
	company := wasmCompany(100, &rulesengine.FeatureEntitlement{
		FeatureID:  "feat_infer",
		FeatureKey: wasmFlagKey,
		ValueType:  rulesengine.EntitlementValueTypeBoolean,
	})
	fixture := newWasmFixture(t, flag, company)

	request := baseRequest()
	request.Usage = 50
	outcome := fixture.run(t, request)

	assert.True(t, fixture.fellBack)
	assert.True(t, outcome.Allowed)
	assert.Nil(t, outcome.Reservation)

	entry, err := fixture.leases.Get(context.Background(), "co_1", testCreditID)
	require.NoError(t, err)
	assert.Nil(t, entry, "no lease is acquired for a feature that draws no credit")
	acquires, _, _ := fixture.wire.snapshot()
	assert.Empty(t, acquires)
}

func TestWasmGatesExactlyAtTheBalanceBoundary(t *testing.T) {
	// The contract the flow leans on: the preflight the SDK builds reaches the
	// engine as the event-scoped usage it gates the credit condition with.
	engine := newWasmEngine(t)
	flag := wasmCreditFlag()
	company := wasmCompany(100)
	ctx := context.Background()

	under, err := engine.CheckFlag(ctx, company, nil, flag,
		EngineCheckFlagOptions(&EvalOptions{EventUsage: &EventUsage{EventSubtype: testSubtype, Quantity: 50}})...)
	require.NoError(t, err)
	over, err := engine.CheckFlag(ctx, company, nil, flag,
		EngineCheckFlagOptions(&EvalOptions{EventUsage: &EventUsage{EventSubtype: testSubtype, Quantity: 150}})...)
	require.NoError(t, err)

	assert.True(t, under.Value)
	assert.False(t, over.Value)
}

// The engine reads usage and event_usage.quantity as i64, so a fraction cannot
// reach it: Go's option types are integers, and the flow rounds up on the way
// in. Rounding up rather than down is what keeps a check from passing on less
// usage than the operation is about to record.
func TestWasmPreflightRoundsAFractionalQuantityUp(t *testing.T) {
	engine := newWasmEngine(t)
	flag := wasmCreditFlag()
	company := wasmCompany(1.5)
	ctx := context.Background()

	opts := preflightOptions(CheckRequest{Usage: 1.5, EventSubtype: testSubtype})
	require.NotNil(t, opts)
	require.NotNil(t, opts.EventUsage)
	assert.Equal(t, int64(2), opts.EventUsage.Quantity)

	// Two whole events at a rate of 1 cost more than the balance of 1.5.
	rounded, err := engine.CheckFlag(ctx, company, nil, flag, EngineCheckFlagOptions(opts)...)
	require.NoError(t, err)
	assert.False(t, rounded.Value)

	// The same usage truncated instead of rounded would have passed.
	truncated, err := engine.CheckFlag(ctx, company, nil, flag,
		EngineCheckFlagOptions(&EvalOptions{EventUsage: &EventUsage{EventSubtype: testSubtype, Quantity: 1}})...)
	require.NoError(t, err)
	assert.True(t, truncated.Value)
}
