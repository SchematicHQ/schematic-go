package rulesengine

import (
	"encoding/json"
	"strings"
	"testing"
)

// TestEnvelopeHasNoNulls pins the invariant that lets marshalEnvelope skip the
// null-stripping pass the Node and Python SDKs need: no collection field ever
// serializes as null. If a future wire field is declared []T instead of
// JSONSlice[T], this fails and the engine would start rejecting checks with -1.
func TestEnvelopeHasNoNulls(t *testing.T) {
	// Worst case: every optional field left nil.
	env := &checkFlagEnvelope{
		Flag:    &Flag{ID: "flag-1", AccountID: "a", EnvironmentID: "e", Key: "k"},
		Company: &Company{ID: "comp-1", AccountID: "a", EnvironmentID: "e"},
		User:    &User{ID: "user-1", AccountID: "a", EnvironmentID: "e"},
	}
	raw, err := json.Marshal(env)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("plain marshal:\n  %s", raw)
	// Nulls on Option-typed scalar fields (base_plan_id, subscription) are fine --
	// serde accepts null for Option<T>. Only collection fields matter: those
	// deserialize into Vec/HashMap, where null is an error rather than a default.
	for _, frag := range []string{
		`"rules":null`, `"traits":null`, `"metrics":null`, `"keys":null`,
		`"billing_product_ids":null`, `"plan_ids":null`, `"plan_version_ids":null`,
		`"credit_balances":null`, `"entitlements":null`,
	} {
		if strings.Contains(string(raw), frag) {
			t.Errorf("collection serialized as null: %s", frag)
		}
	}
}

// TestResultDecodesWarningTiers pins the warning-tier leg of the camelCase ->
// snake_case bridge in CheckFlagResult.UnmarshalJSON. The pinned v0.6.0 engine
// does not emit warningTiers, so the populated case is exercised against a
// payload shaped like the build that will, and the absent case covers today.
func TestResultDecodesWarningTiers(t *testing.T) {
	const withTiers = `{
		"value": true,
		"reason": "Plan entitlement",
		"flagKey": "seats-flag",
		"entitlement": {
			"featureId": "feat-1",
			"featureKey": "seats",
			"valueType": "numeric",
			"warningTiers": [{"key": "soft", "value": 80}, {"key": "hard", "value": 95}]
		}
	}`

	var r CheckFlagResult
	if err := json.Unmarshal([]byte(withTiers), &r); err != nil {
		t.Fatal(err)
	}
	if r.Entitlement == nil {
		t.Fatal("entitlement failed to decode")
	}
	if got := len(r.Entitlement.WarningTiers); got != 2 {
		t.Fatalf("want 2 warning tiers, got %d", got)
	}
	if k, v := r.Entitlement.WarningTiers[0].Key, r.Entitlement.WarningTiers[0].Value; k != "soft" || v != 80 {
		t.Errorf("first tier = %q/%d, want soft/80", k, v)
	}

	// The public type is snake_case, so a decoded result must re-marshal as
	// warning_tiers -- the engine's casing must not leak to SDK consumers.
	out, err := json.Marshal(&r)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), `"warning_tiers":[{"key":"soft","value":80}`) {
		t.Errorf("result did not re-marshal warning tiers as snake_case: %s", out)
	}

	// v0.6.0 omits the field entirely; that has to stay benign.
	const withoutTiers = `{"value":true,"reason":"r","flagKey":"k",` +
		`"entitlement":{"featureId":"f","featureKey":"k","valueType":"numeric"}}`
	var absent CheckFlagResult
	if err := json.Unmarshal([]byte(withoutTiers), &absent); err != nil {
		t.Fatal(err)
	}
	if absent.Entitlement.WarningTiers != nil {
		t.Errorf("absent warningTiers should decode to nil, got %v", absent.Entitlement.WarningTiers)
	}
}

// TestNestedRuleHasNoNulls covers the same invariant at depth, where a nil
// collection hides inside a rule or metric rather than on the top-level entity.
func TestNestedRuleHasNoNulls(t *testing.T) {
	// A rule with nil Conditions/ConditionGroups, and a metric with nil ValidUntil.
	env := &checkFlagEnvelope{
		Flag: &Flag{
			ID: "flag-1", AccountID: "a", EnvironmentID: "e", Key: "k",
			Rules: JSONSlice[*Rule]{{ID: "rule-1", AccountID: "a", EnvironmentID: "e", RuleType: RuleTypeStandard}},
		},
		Company: &Company{
			ID: "comp-1", AccountID: "a", EnvironmentID: "e",
			Metrics: CompanyMetricCollection{{AccountID: "a", EnvironmentID: "e", CompanyID: "comp-1"}},
		},
	}
	raw, err := json.Marshal(env)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("nested marshal:\n  %s", raw)
	for _, frag := range []string{`"conditions":null`, `"condition_groups":null`, `"resource_ids":null`, `"rules":null`, `"traits":null`, `"metrics":null`} {
		if strings.Contains(string(raw), frag) {
			t.Errorf("found collection null: %s", frag)
		}
	}
}
