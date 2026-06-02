package datastream

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/schematichq/rulesengine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func baseCompany() *rulesengine.Company {
	plan := "plan-1"
	return &rulesengine.Company{
		ID:                "co-1",
		AccountID:         "acc-1",
		EnvironmentID:     "env-1",
		BasePlanID:        &plan,
		BillingProductIDs: []string{"bp-1"},
		CreditBalances:    map[string]float64{"credit-1": 100.0},
		Keys:              map[string]string{"domain": "example.com"},
		PlanIDs:           []string{"plan-1"},
		PlanVersionIDs:    []string{"pv-1"},
		Traits: []*rulesengine.Trait{
			{Value: "Enterprise", TraitDefinition: &rulesengine.TraitDefinition{ID: "plan"}},
		},
		Entitlements: []*rulesengine.FeatureEntitlement{
			{FeatureID: "feat-1", FeatureKey: "feature-one"},
		},
	}
}

func baseUser() *rulesengine.User {
	return &rulesengine.User{
		ID:            "user-1",
		AccountID:     "acc-1",
		EnvironmentID: "env-1",
		Keys:          map[string]string{"email": "user@example.com"},
		Traits: []*rulesengine.Trait{
			{Value: "Premium", TraitDefinition: &rulesengine.TraitDefinition{ID: "tier"}},
		},
	}
}

func TestPartialCompany_AddsCreditBalance(t *testing.T) {
	existing := baseCompany()
	partial := json.RawMessage(`{"credit_balances":{"credit-2":200.0}}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	assert.Equal(t, map[string]float64{"credit-1": 100.0, "credit-2": 200.0}, merged.CreditBalances)
	assert.Equal(t, "acc-1", merged.AccountID)
}

func TestPartialCompany_OverwritesCreditBalance(t *testing.T) {
	existing := baseCompany()
	partial := json.RawMessage(`{"credit_balances":{"credit-1":50.0}}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	assert.Equal(t, map[string]float64{"credit-1": 50.0}, merged.CreditBalances)
}

func TestPartialCompany_CreditBalances_NilExistingMap(t *testing.T) {
	existing := baseCompany()
	existing.CreditBalances = nil
	partial := json.RawMessage(`{"credit_balances":{"credit-1":42.5}}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	assert.Equal(t, map[string]float64{"credit-1": 42.5}, merged.CreditBalances)
}

func TestPartialCompany_UpsertsMetric(t *testing.T) {
	existing := baseCompany()
	existing.Metrics = rulesengine.CompanyMetricCollection{
		{EventSubtype: "event-a", Period: "all_time", MonthReset: "first_of_month", Value: 10},
		{EventSubtype: "event-b", Period: "current_month", MonthReset: "first_of_month", Value: 5},
	}
	// Update event-a in place + append event-c.
	partial := json.RawMessage(`{"metrics":[
		{"event_subtype":"event-a","period":"all_time","month_reset":"first_of_month","value":42},
		{"event_subtype":"event-c","period":"current_day","month_reset":"billing_cycle","value":1}
	]}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	require.Len(t, merged.Metrics, 3)
	assert.Equal(t, "event-a", merged.Metrics[0].EventSubtype)
	assert.EqualValues(t, 42, merged.Metrics[0].Value)
	assert.Equal(t, "event-b", merged.Metrics[1].EventSubtype)
	assert.EqualValues(t, 5, merged.Metrics[1].Value)
	assert.Equal(t, "event-c", merged.Metrics[2].EventSubtype)
	assert.EqualValues(t, 1, merged.Metrics[2].Value)

	// Original not mutated.
	assert.EqualValues(t, 10, existing.Metrics[0].Value)
}

func TestPartialCompany_MergesKeys(t *testing.T) {
	existing := baseCompany()
	partial := json.RawMessage(`{"keys":{"slug":"new-slug"}}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	assert.Equal(t, map[string]string{"domain": "example.com", "slug": "new-slug"}, merged.Keys)
}

func TestPartialCompany_ReplacesTraits(t *testing.T) {
	existing := baseCompany()
	partial := json.RawMessage(`{"traits":[{"value":"Startup","trait_definition":{"id":"plan"}}]}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	require.Len(t, merged.Traits, 1)
	assert.Equal(t, "Startup", merged.Traits[0].Value)
	// Original not mutated.
	require.Len(t, existing.Traits, 1)
	assert.Equal(t, "Enterprise", existing.Traits[0].Value)
}

func TestPartialCompany_ReplacesEntitlements(t *testing.T) {
	existing := baseCompany()
	partial := json.RawMessage(`{"entitlements":[
		{"feature_id":"feat-new","feature_key":"feature-new"},
		{"feature_id":"feat-2","feature_key":"feature-two"}
	]}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	require.Len(t, merged.Entitlements, 2)
	assert.Equal(t, "feat-new", merged.Entitlements[0].FeatureID)
	assert.Equal(t, "feat-2", merged.Entitlements[1].FeatureID)
}

func TestPartialCompany_ReplacesRules(t *testing.T) {
	existing := baseCompany()
	existing.Rules = []*rulesengine.Rule{{ID: "rule-old"}}
	partial := json.RawMessage(`{"rules":[{"id":"rule-new"}]}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	require.Len(t, merged.Rules, 1)
	assert.Equal(t, "rule-new", merged.Rules[0].ID)
	// Original not mutated.
	assert.Equal(t, "rule-old", existing.Rules[0].ID)
}

func TestPartialCompany_ReplacesSubscription(t *testing.T) {
	existing := baseCompany()
	existing.Subscription = &rulesengine.Subscription{ID: "sub-old"}
	partial := json.RawMessage(`{"subscription":{"id":"sub-new"}}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	require.NotNil(t, merged.Subscription)
	assert.Equal(t, "sub-new", merged.Subscription.ID)
	// Original not mutated.
	assert.Equal(t, "sub-old", existing.Subscription.ID)
}

func TestPartialCompany_ClearsSubscription(t *testing.T) {
	existing := baseCompany()
	existing.Subscription = &rulesengine.Subscription{ID: "sub-old"}
	partial := json.RawMessage(`{"subscription":null}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	assert.Nil(t, merged.Subscription)
	// Original not mutated.
	require.NotNil(t, existing.Subscription)
}

func TestPartialCompany_ReplacesBillingProductIDs(t *testing.T) {
	existing := baseCompany()
	partial := json.RawMessage(`{"billing_product_ids":["bp-10","bp-20"]}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	assert.Equal(t, rulesengine.JSONSlice[string]{"bp-10", "bp-20"}, merged.BillingProductIDs)
	assert.Equal(t, rulesengine.JSONSlice[string]{"bp-1"}, existing.BillingProductIDs)
}

func TestPartialCompany_ReplacesPlanIDs(t *testing.T) {
	existing := baseCompany()
	partial := json.RawMessage(`{"plan_ids":["plan-99","plan-100"]}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	assert.Equal(t, rulesengine.JSONSlice[string]{"plan-99", "plan-100"}, merged.PlanIDs)
	assert.Equal(t, rulesengine.JSONSlice[string]{"plan-1"}, existing.PlanIDs)
}

func TestPartialCompany_ReplacesPlanVersionIDs(t *testing.T) {
	existing := baseCompany()
	partial := json.RawMessage(`{"plan_version_ids":["pv-99"]}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	assert.Equal(t, rulesengine.JSONSlice[string]{"pv-99"}, merged.PlanVersionIDs)
	assert.Equal(t, rulesengine.JSONSlice[string]{"pv-1"}, existing.PlanVersionIDs)
}

func TestPartialCompany_NullBasePlanID(t *testing.T) {
	existing := baseCompany()
	partial := json.RawMessage(`{"base_plan_id":null}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	assert.Nil(t, merged.BasePlanID)
	require.NotNil(t, existing.BasePlanID)
	assert.Equal(t, "plan-1", *existing.BasePlanID)
}

func TestPartialCompany_UnknownFieldsAreIgnored(t *testing.T) {
	existing := baseCompany()
	partial := json.RawMessage(`{"something_new":{"x":1}}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	// Unknown field is silently dropped — the SDK doesn't error on partial
	// types it doesn't understand, so producers can ship new fields ahead of
	// SDK support without breaking older clients.
	assert.Equal(t, "acc-1", merged.AccountID)
	assert.Equal(t, map[string]float64{"credit-1": 100.0}, merged.CreditBalances)
}

func TestPartialCompany_MultipleFieldsInOnePayload(t *testing.T) {
	existing := baseCompany()
	partial := json.RawMessage(`{
		"credit_balances":{"credit-2":50.0},
		"keys":{"slug":"new"}
	}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	assert.Equal(t, map[string]float64{"credit-1": 100.0, "credit-2": 50.0}, merged.CreditBalances)
	assert.Equal(t, map[string]string{"domain": "example.com", "slug": "new"}, merged.Keys)
}

func TestPartialCompany_DoesNotMutateOriginal(t *testing.T) {
	existing := baseCompany()
	origBalances := map[string]float64{}
	for k, v := range existing.CreditBalances {
		origBalances[k] = v
	}

	partial := json.RawMessage(`{"credit_balances":{"credit-1":999.0,"credit-2":50.0}}`)
	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	assert.Equal(t, origBalances, existing.CreditBalances)
	assert.Equal(t, map[string]float64{"credit-1": 999.0, "credit-2": 50.0}, merged.CreditBalances)
}

// Credit-balance partials don't include refreshed entitlements, so the SDK
// syncs credit_remaining locally to keep the two in step for consumers who
// read the cached entitlement object between full company refreshes.

func makeEntitlement(featureID, featureKey string) *rulesengine.FeatureEntitlement {
	return &rulesengine.FeatureEntitlement{FeatureID: featureID, FeatureKey: featureKey}
}

func withCredit(ent *rulesengine.FeatureEntitlement, creditID string, remaining float64) *rulesengine.FeatureEntitlement {
	id := creditID
	r := remaining
	ent.CreditID = &id
	ent.CreditRemaining = &r
	return ent
}

func withEvent(ent *rulesengine.FeatureEntitlement, eventName string, period *rulesengine.MetricPeriod, monthReset *rulesengine.MetricPeriodMonthReset, usage int64) *rulesengine.FeatureEntitlement {
	name := eventName
	u := usage
	ent.EventName = &name
	ent.MetricPeriod = period
	ent.MonthReset = monthReset
	ent.Usage = &u
	return ent
}

func TestPartialCompany_SyncsCreditRemaining_MatchingCreditID(t *testing.T) {
	existing := baseCompany()
	existing.CreditBalances = map[string]float64{"credit-1": 100.0}
	existing.Entitlements = []*rulesengine.FeatureEntitlement{
		withCredit(makeEntitlement("feat-1", "f1"), "credit-1", 100.0),
		makeEntitlement("feat-2", "f2"), // no credit_id — must stay untouched
	}

	partial := json.RawMessage(`{"credit_balances":{"credit-1":25.0}}`)
	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	require.Len(t, merged.Entitlements, 2)
	require.NotNil(t, merged.Entitlements[0].CreditRemaining)
	assert.Equal(t, 25.0, *merged.Entitlements[0].CreditRemaining)
	assert.Nil(t, merged.Entitlements[1].CreditRemaining)

	// Original not mutated.
	require.NotNil(t, existing.Entitlements[0].CreditRemaining)
	assert.Equal(t, 100.0, *existing.Entitlements[0].CreditRemaining)
}

func TestPartialCompany_SyncsCreditRemaining_AcrossMultipleCreditIDs(t *testing.T) {
	existing := baseCompany()
	existing.CreditBalances = map[string]float64{"credit-1": 100.0, "credit-2": 50.0}
	existing.Entitlements = []*rulesengine.FeatureEntitlement{
		withCredit(makeEntitlement("feat-1", "f1"), "credit-1", 100.0),
		withCredit(makeEntitlement("feat-2", "f2"), "credit-2", 50.0),
	}

	partial := json.RawMessage(`{"credit_balances":{"credit-1":75.0,"credit-2":10.0}}`)
	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	require.Len(t, merged.Entitlements, 2)
	require.NotNil(t, merged.Entitlements[0].CreditRemaining)
	assert.Equal(t, 75.0, *merged.Entitlements[0].CreditRemaining)
	require.NotNil(t, merged.Entitlements[1].CreditRemaining)
	assert.Equal(t, 10.0, *merged.Entitlements[1].CreditRemaining)
}

func TestPartialCompany_UnmatchedCreditID_LeftAlone(t *testing.T) {
	existing := baseCompany()
	existing.CreditBalances = map[string]float64{"credit-1": 100.0}
	existing.Entitlements = []*rulesengine.FeatureEntitlement{
		withCredit(makeEntitlement("feat-1", "f1"), "credit-other", 999.0),
	}

	// Partial only updates credit-1; entitlement points at credit-other and
	// must keep its existing credit_remaining.
	partial := json.RawMessage(`{"credit_balances":{"credit-1":25.0}}`)
	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	require.Len(t, merged.Entitlements, 1)
	require.NotNil(t, merged.Entitlements[0].CreditRemaining)
	assert.Equal(t, 999.0, *merged.Entitlements[0].CreditRemaining)
}

func TestPartialCompany_SingleCreditFansOutToMultipleEntitlements(t *testing.T) {
	// Common when one credit type funds multiple features — each gets its
	// own entitlement sharing the same credit_id.
	existing := baseCompany()
	existing.CreditBalances = map[string]float64{"credit-shared": 500.0}
	existing.Entitlements = []*rulesengine.FeatureEntitlement{
		withCredit(makeEntitlement("feat-a", "feature-a"), "credit-shared", 500.0),
		withCredit(makeEntitlement("feat-b", "feature-b"), "credit-shared", 500.0),
		withCredit(makeEntitlement("feat-c", "feature-c"), "credit-shared", 500.0),
		makeEntitlement("feat-d", "feature-d"), // unrelated, no credit
	}

	partial := json.RawMessage(`{"credit_balances":{"credit-shared":120.0}}`)
	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	require.Len(t, merged.Entitlements, 4)
	for i := range 3 {
		require.NotNil(t, merged.Entitlements[i].CreditRemaining, "entitlement %d", i)
		assert.Equal(t, 120.0, *merged.Entitlements[i].CreditRemaining, "entitlement %d", i)
	}
	assert.Nil(t, merged.Entitlements[3].CreditRemaining)
}

func TestPartialCompany_SyncSkippedWhenPartialAlsoSendsEntitlements(t *testing.T) {
	// If the partial carries entitlements, we trust those wholesale and don't
	// re-derive credit_remaining from credit_balances.
	existing := baseCompany()
	existing.CreditBalances = map[string]float64{"credit-1": 100.0}
	existing.Entitlements = []*rulesengine.FeatureEntitlement{
		withCredit(makeEntitlement("feat-1", "f1"), "credit-1", 100.0),
	}

	partial := json.RawMessage(`{
		"credit_balances":{"credit-1":25.0},
		"entitlements":[{"feature_id":"feat-1","feature_key":"f1","value_type":"boolean","credit_id":"credit-1","credit_remaining":17.0}]
	}`)
	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	require.Len(t, merged.Entitlements, 1)
	require.NotNil(t, merged.Entitlements[0].CreditRemaining)
	assert.Equal(t, 17.0, *merged.Entitlements[0].CreditRemaining)
}

func TestPartialCompany_SyncNoOpWhenNoEntitlementsExist(t *testing.T) {
	existing := baseCompany()
	existing.Entitlements = nil

	partial := json.RawMessage(`{"credit_balances":{"credit-1":25.0}}`)
	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	assert.Equal(t, 25.0, merged.CreditBalances["credit-1"])
	assert.Empty(t, merged.Entitlements)
}

func TestPartialCompany_SyncsUsage_EventBasedEntitlement(t *testing.T) {
	existing := baseCompany()
	currentMonth := rulesengine.MetricPeriodCurrentMonth
	firstOfMonth := rulesengine.MetricPeriodMonthResetFirst
	existing.Metrics = rulesengine.CompanyMetricCollection{
		{EventSubtype: "credits_used", Period: currentMonth, MonthReset: firstOfMonth, Value: 10},
	}
	existing.Entitlements = []*rulesengine.FeatureEntitlement{
		withEvent(makeEntitlement("feat-1", "f1"), "credits_used", &currentMonth, &firstOfMonth, 10),
	}

	partial := json.RawMessage(`{"metrics":[{"event_subtype":"credits_used","period":"current_month","month_reset":"first_of_month","value":42}]}`)
	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	require.Len(t, merged.Entitlements, 1)
	require.NotNil(t, merged.Entitlements[0].Usage)
	assert.EqualValues(t, 42, *merged.Entitlements[0].Usage)
}

func TestPartialCompany_UsageMatchRequiresFullTriple(t *testing.T) {
	// The server matches metrics to entitlements on the full triple
	// (event_subtype, period, month_reset). A metric with a different period
	// must not satisfy an entitlement's lookup.
	existing := baseCompany()
	allTime := rulesengine.MetricPeriodAllTime
	currentMonth := rulesengine.MetricPeriodCurrentMonth
	firstOfMonth := rulesengine.MetricPeriodMonthResetFirst
	existing.Metrics = rulesengine.CompanyMetricCollection{
		{EventSubtype: "api_calls", Period: allTime, MonthReset: firstOfMonth, Value: 100},
	}
	existing.Entitlements = []*rulesengine.FeatureEntitlement{
		withEvent(makeEntitlement("feat-1", "f1"), "api_calls", &currentMonth, &firstOfMonth, 5),
	}

	// Partial updates the all_time metric. Entitlement points at current_month
	// so its usage must NOT change.
	partial := json.RawMessage(`{"metrics":[{"event_subtype":"api_calls","period":"all_time","month_reset":"first_of_month","value":999}]}`)
	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	require.Len(t, merged.Entitlements, 1)
	require.NotNil(t, merged.Entitlements[0].Usage)
	assert.EqualValues(t, 5, *merged.Entitlements[0].Usage)
}

func TestPartialCompany_UsageDefaultsToAllTimeFirstOfMonth(t *testing.T) {
	// When entitlement period/month_reset are nil, the server's metric lookup
	// defaults to all_time / first_of_month — the sync must do the same.
	existing := baseCompany()
	allTime := rulesengine.MetricPeriodAllTime
	firstOfMonth := rulesengine.MetricPeriodMonthResetFirst
	existing.Metrics = rulesengine.CompanyMetricCollection{
		{EventSubtype: "api_calls", Period: allTime, MonthReset: firstOfMonth, Value: 0},
	}
	existing.Entitlements = []*rulesengine.FeatureEntitlement{
		withEvent(makeEntitlement("feat-1", "f1"), "api_calls", nil, nil, 0),
	}

	partial := json.RawMessage(`{"metrics":[{"event_subtype":"api_calls","period":"all_time","month_reset":"first_of_month","value":7}]}`)
	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	require.Len(t, merged.Entitlements, 1)
	require.NotNil(t, merged.Entitlements[0].Usage)
	assert.EqualValues(t, 7, *merged.Entitlements[0].Usage)
}

func TestPartialCompany_UsageUnchangedWhenNoMatchingMetricInPartial(t *testing.T) {
	// If the partial's metric upsert leaves the entitlement's metric untouched
	// (different event_subtype), keep the existing usage. The sync re-derives
	// from the merged metrics list, which still holds event-a at value 50.
	existing := baseCompany()
	allTime := rulesengine.MetricPeriodAllTime
	firstOfMonth := rulesengine.MetricPeriodMonthResetFirst
	existing.Metrics = rulesengine.CompanyMetricCollection{
		{EventSubtype: "event-a", Period: allTime, MonthReset: firstOfMonth, Value: 50},
	}
	existing.Entitlements = []*rulesengine.FeatureEntitlement{
		withEvent(makeEntitlement("feat-1", "f1"), "event-a", &allTime, &firstOfMonth, 50),
	}

	partial := json.RawMessage(`{"metrics":[{"event_subtype":"event-b","period":"all_time","month_reset":"first_of_month","value":999}]}`)
	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	require.Len(t, merged.Entitlements, 1)
	require.NotNil(t, merged.Entitlements[0].Usage)
	assert.EqualValues(t, 50, *merged.Entitlements[0].Usage)
}

func TestPartialCompany_SyncsCreditRemainingAndUsageTogether(t *testing.T) {
	// A partial can carry both credit_balances and metrics changes; both
	// derived fields must be applied in the same entitlements rebuild.
	existing := baseCompany()
	allTime := rulesengine.MetricPeriodAllTime
	firstOfMonth := rulesengine.MetricPeriodMonthResetFirst
	existing.CreditBalances = map[string]float64{"credit-1": 100.0}
	existing.Metrics = rulesengine.CompanyMetricCollection{
		{EventSubtype: "event-a", Period: allTime, MonthReset: firstOfMonth, Value: 5},
	}
	ent := withCredit(makeEntitlement("feat-1", "f1"), "credit-1", 100.0)
	ent = withEvent(ent, "event-a", &allTime, &firstOfMonth, 5)
	existing.Entitlements = []*rulesengine.FeatureEntitlement{ent}

	partial := json.RawMessage(`{
		"credit_balances":{"credit-1":25.0},
		"metrics":[{"event_subtype":"event-a","period":"all_time","month_reset":"first_of_month","value":80}]
	}`)
	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	require.Len(t, merged.Entitlements, 1)
	require.NotNil(t, merged.Entitlements[0].CreditRemaining)
	assert.Equal(t, 25.0, *merged.Entitlements[0].CreditRemaining)
	require.NotNil(t, merged.Entitlements[0].Usage)
	assert.EqualValues(t, 80, *merged.Entitlements[0].Usage)
}

func TestPartialCompany_CreditTotalAndUsedNotDerivedFromPartial(t *testing.T) {
	// credit_total and credit_used aggregate across a grant ledger the SDK
	// doesn't see. They must keep their last full-message value across
	// credit_balances partials; grant-lifecycle events trigger a full
	// company message that refreshes them.
	existing := baseCompany()
	existing.CreditBalances = map[string]float64{"credit-1": 100.0}
	total := 200.0
	used := 100.0
	ent := withCredit(makeEntitlement("feat-1", "f1"), "credit-1", 100.0)
	ent.CreditTotal = &total
	ent.CreditUsed = &used
	existing.Entitlements = []*rulesengine.FeatureEntitlement{ent}

	partial := json.RawMessage(`{"credit_balances":{"credit-1":25.0}}`)
	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	require.Len(t, merged.Entitlements, 1)
	require.NotNil(t, merged.Entitlements[0].CreditTotal)
	assert.Equal(t, 200.0, *merged.Entitlements[0].CreditTotal)
	require.NotNil(t, merged.Entitlements[0].CreditUsed)
	assert.Equal(t, 100.0, *merged.Entitlements[0].CreditUsed)
}

func TestPartialUser_ReplacesTraits(t *testing.T) {
	existing := baseUser()
	partial := json.RawMessage(`{"traits":[{"value":"Free","trait_definition":{"id":"tier"}}]}`)

	merged, err := PartialUser(existing, partial)
	require.NoError(t, err)

	assert.Len(t, merged.Traits, 1)
	assert.Equal(t, "Free", merged.Traits[0].Value)
	assert.Equal(t, map[string]string{"email": "user@example.com"}, merged.Keys)
}

func TestPartialUser_MergesKeys(t *testing.T) {
	existing := baseUser()
	partial := json.RawMessage(`{"keys":{"slack_id":"U123"}}`)

	merged, err := PartialUser(existing, partial)
	require.NoError(t, err)

	// New key added, existing key preserved.
	assert.Equal(t, map[string]string{"email": "user@example.com", "slack_id": "U123"}, merged.Keys)
	assert.Len(t, merged.Traits, 1)
}

func TestPartialUser_DoesNotMutateOriginal(t *testing.T) {
	existing := baseUser()
	origKeys := make(map[string]string)
	for k, v := range existing.Keys {
		origKeys[k] = v
	}

	partial := json.RawMessage(`{"keys":{"slug":"new"},"traits":[]}`)

	merged, err := PartialUser(existing, partial)
	require.NoError(t, err)

	assert.Equal(t, origKeys, existing.Keys)
	assert.Len(t, existing.Traits, 1)

	assert.Equal(t, map[string]string{"email": "user@example.com", "slug": "new"}, merged.Keys)
	assert.Empty(t, merged.Traits)
}

func TestPartialUser_MultipleFieldsInOnePayload(t *testing.T) {
	existing := baseUser()
	existing.Rules = []*rulesengine.Rule{{ID: "rule-1"}}

	partial := json.RawMessage(`{
		"keys": {"email": "new@example.com", "slack_id": "U999"},
		"traits": [{"value": "Free", "trait_definition": {"id": "tier"}}],
		"rules": [{"id": "rule-new"}]
	}`)

	merged, err := PartialUser(existing, partial)
	require.NoError(t, err)

	// Keys merge.
	assert.Equal(t, map[string]string{"email": "new@example.com", "slack_id": "U999"}, merged.Keys)
	require.Len(t, merged.Traits, 1)
	assert.Equal(t, "Free", merged.Traits[0].Value)
	require.Len(t, merged.Rules, 1)
	assert.Equal(t, "rule-new", merged.Rules[0].ID)

	// Original not mutated.
	assert.Equal(t, map[string]string{"email": "user@example.com"}, existing.Keys)
	assert.Equal(t, "rule-1", existing.Rules[0].ID)
}

func TestDeepCopyCompany_Nil(t *testing.T) {
	assert.Nil(t, DeepCopyCompany(nil))
}

func TestDeepCopyCompany_FullCopy(t *testing.T) {
	validUntil := time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC)
	plan := "plan-1"
	orig := &rulesengine.Company{
		ID:                "co-1",
		AccountID:         "acc-1",
		EnvironmentID:     "env-1",
		BasePlanID:        &plan,
		BillingProductIDs: []string{"bp-1", "bp-2"},
		CreditBalances:    map[string]float64{"credit-1": 50.0},
		Keys:              map[string]string{"domain": "example.com"},
		PlanIDs:           []string{"plan-1"},
		PlanVersionIDs:    []string{"pv-1"},
		Subscription:      &rulesengine.Subscription{ID: "sub-1"},
		Entitlements:      []*rulesengine.FeatureEntitlement{{FeatureID: "f1"}},
		Rules:             []*rulesengine.Rule{{ID: "r1"}},
		Metrics: []*rulesengine.CompanyMetric{
			{
				CompanyID:    "co-1",
				EventSubtype: "event-1",
				Value:        42,
				ValidUntil:   &validUntil,
			},
			nil,
		},
		Traits: []*rulesengine.Trait{
			{Value: "Enterprise", TraitDefinition: &rulesengine.TraitDefinition{ID: "plan"}},
			nil,
		},
	}

	cp := DeepCopyCompany(orig)

	assert.Equal(t, orig.ID, cp.ID)
	assert.Equal(t, orig.AccountID, cp.AccountID)
	assert.Equal(t, orig.EnvironmentID, cp.EnvironmentID)

	// BasePlanID deep copied
	require.NotNil(t, cp.BasePlanID)
	assert.Equal(t, *orig.BasePlanID, *cp.BasePlanID)
	newPlan := "plan-2"
	cp.BasePlanID = &newPlan
	assert.Equal(t, "plan-1", *orig.BasePlanID)

	// Slices are independent
	cp.BillingProductIDs[0] = "changed"
	assert.Equal(t, "bp-1", orig.BillingProductIDs[0])

	cp.PlanIDs[0] = "changed"
	assert.Equal(t, "plan-1", orig.PlanIDs[0])

	cp.PlanVersionIDs[0] = "changed"
	assert.Equal(t, "pv-1", orig.PlanVersionIDs[0])

	// Maps are independent
	cp.Keys["domain"] = "changed.com"
	assert.Equal(t, "example.com", orig.Keys["domain"])

	cp.CreditBalances["credit-1"] = 999
	assert.Equal(t, 50.0, orig.CreditBalances["credit-1"])

	// Subscription deep copied
	require.NotNil(t, cp.Subscription)
	assert.Equal(t, "sub-1", cp.Subscription.ID)
	cp.Subscription.ID = "changed"
	assert.Equal(t, "sub-1", orig.Subscription.ID)

	// Metrics deep copied (including ValidUntil pointer)
	require.Len(t, cp.Metrics, 2)
	require.NotNil(t, cp.Metrics[0])
	assert.EqualValues(t, 42, cp.Metrics[0].Value)
	require.NotNil(t, cp.Metrics[0].ValidUntil)
	assert.Equal(t, validUntil, *cp.Metrics[0].ValidUntil)
	newTime := time.Date(2099, 1, 1, 0, 0, 0, 0, time.UTC)
	cp.Metrics[0].ValidUntil = &newTime
	assert.Equal(t, validUntil, *orig.Metrics[0].ValidUntil)
	assert.Nil(t, cp.Metrics[1])

	// Traits deep copied
	require.Len(t, cp.Traits, 2)
	require.NotNil(t, cp.Traits[0])
	assert.Equal(t, "Enterprise", cp.Traits[0].Value)
	cp.Traits[0].Value = "changed"
	assert.Equal(t, "Enterprise", orig.Traits[0].Value)
	assert.Nil(t, cp.Traits[1])
}

func TestDeepCopyUser_EmptyFields(t *testing.T) {
	cp := DeepCopyUser(&rulesengine.User{ID: "u1"})
	assert.Equal(t, "u1", cp.ID)
	assert.Nil(t, cp.Keys)
	assert.Nil(t, cp.Traits)
	assert.Nil(t, cp.Rules)
}

func TestDeepCopyUser_FullCopy(t *testing.T) {
	orig := &rulesengine.User{
		ID:            "user-1",
		AccountID:     "acc-1",
		EnvironmentID: "env-1",
		Keys:          map[string]string{"email": "a@b.com"},
		Traits:        []*rulesengine.Trait{{Value: "Premium"}},
		Rules:         []*rulesengine.Rule{{ID: "r1"}},
	}

	cp := DeepCopyUser(orig)

	assert.Equal(t, orig.ID, cp.ID)
	assert.Equal(t, orig.AccountID, cp.AccountID)

	cp.Keys["email"] = "changed"
	assert.Equal(t, "a@b.com", orig.Keys["email"])

	cp.Traits[0] = &rulesengine.Trait{Value: "Free"}
	assert.Equal(t, "Premium", orig.Traits[0].Value)

	cp.Rules[0] = &rulesengine.Rule{ID: "r2"}
	assert.Equal(t, "r1", orig.Rules[0].ID)
}
