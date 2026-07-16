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
