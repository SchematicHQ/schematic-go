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
		`"credit_balances":null`, `"credit_postpaid":null`, `"entitlements":null`,
	} {
		if strings.Contains(string(raw), frag) {
			t.Errorf("collection serialized as null: %s", frag)
		}
	}
}

// TestCompanyCreditPostpaidRoundTrip pins the credit_postpaid wire shape the
// engine reads. A null entry means postpaid off, so decoding must drop it:
// kept as a zero config, it would re-marshal as {} and the engine would read
// that as postpaid on with no limit.
func TestCompanyCreditPostpaidRoundTrip(t *testing.T) {
	const payload = `{"id":"comp-1","account_id":"a","environment_id":"e","credit_postpaid":{` +
		`"limited":{"overdraft_limit":100},"unlimited":{},"off":null}}`

	var c Company
	if err := json.Unmarshal([]byte(payload), &c); err != nil {
		t.Fatal(err)
	}
	if got := len(c.CreditPostpaid); got != 2 {
		t.Fatalf("want 2 postpaid credits, got %d: %v", got, c.CreditPostpaid)
	}
	if _, ok := c.CreditPostpaid["off"]; ok {
		t.Error("null entry should decode as absent (postpaid off)")
	}
	if limited, ok := c.CreditPostpaid["limited"]; !ok || limited.OverdraftLimit == nil || *limited.OverdraftLimit != 100 {
		t.Errorf("limited = %+v, want overdraft_limit 100", limited)
	}
	if unlimited, ok := c.CreditPostpaid["unlimited"]; !ok || unlimited.OverdraftLimit != nil {
		t.Errorf("unlimited = %+v (present %v), want present with no limit", unlimited, ok)
	}

	out, err := json.Marshal(&c)
	if err != nil {
		t.Fatal(err)
	}
	var wire struct {
		CreditPostpaid map[string]json.RawMessage `json:"credit_postpaid"`
	}
	if err := json.Unmarshal(out, &wire); err != nil {
		t.Fatal(err)
	}
	want := map[string]string{"limited": `{"overdraft_limit":100}`, "unlimited": `{}`}
	if len(wire.CreditPostpaid) != len(want) {
		t.Fatalf("re-marshaled credit_postpaid = %s", out)
	}
	for creditID, entry := range want {
		if got := string(wire.CreditPostpaid[creditID]); got != entry {
			t.Errorf("credit_postpaid[%q] = %s, want %s", creditID, got, entry)
		}
	}

	// A company without postpaid omits the field, which the engine reads as off.
	out, err = json.Marshal(&Company{ID: "comp-2"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(out), "credit_postpaid") {
		t.Errorf("empty credit_postpaid should be omitted: %s", out)
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

// TestQuantityRatesRoundTrip pins quantity_rates through each leg it travels:
// a credit condition and a company entitlement decoded from the datastream's
// snake_case payload and re-marshaled for the engine, and the engine's
// camelCase entitlement decoded onto the public snake_case result.
func TestQuantityRatesRoundTrip(t *testing.T) {
	const flagPayload = `{"id":"flag-1","account_id":"a","environment_id":"e","key":"k","rules":[{` +
		`"id":"rule-1","account_id":"a","environment_id":"e","rule_type":"plan_entitlement","conditions":[{` +
		`"id":"cond-1","account_id":"a","environment_id":"e","condition_type":"credit","operator":"lt",` +
		`"credit_id":"credit-abc","consumption_rate":0.5,"quantity_rates":{"input_tokens":0.001,"output_tokens":0.01}}]}]}`

	var flag Flag
	if err := json.Unmarshal([]byte(flagPayload), &flag); err != nil {
		t.Fatal(err)
	}
	rates := flag.Rules[0].Conditions[0].QuantityRates
	if rates["input_tokens"] != 0.001 || rates["output_tokens"] != 0.01 || len(rates) != 2 {
		t.Fatalf("condition quantity_rates = %v", rates)
	}
	out, err := json.Marshal(&flag)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), `"quantity_rates":{"input_tokens":0.001,"output_tokens":0.01}`) {
		t.Errorf("condition did not re-marshal quantity_rates: %s", out)
	}

	const companyPayload = `{"id":"comp-1","account_id":"a","environment_id":"e","entitlements":[{` +
		`"feature_id":"feat-1","feature_key":"chat","value_type":"credit","quantity_rates":{"input_tokens":0.001}}]}`

	var company Company
	if err := json.Unmarshal([]byte(companyPayload), &company); err != nil {
		t.Fatal(err)
	}
	if got := company.Entitlements[0].QuantityRates["input_tokens"]; got != 0.001 {
		t.Fatalf("entitlement quantity_rates = %v", company.Entitlements[0].QuantityRates)
	}
	out, err = json.Marshal(&company)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), `"quantity_rates":{"input_tokens":0.001}`) {
		t.Errorf("entitlement did not re-marshal quantity_rates: %s", out)
	}

	// Empty schedules are omitted, as the API sends them.
	out, err = json.Marshal(&Condition{ID: "cond-2"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(out), "quantity_rates") {
		t.Errorf("empty quantity_rates should be omitted: %s", out)
	}

	const result = `{"value":true,"reason":"r","flagKey":"k","entitlement":{"featureId":"f","featureKey":"k",` +
		`"valueType":"credit","quantityRates":{"output_tokens":0.01}}}`
	var r CheckFlagResult
	if err := json.Unmarshal([]byte(result), &r); err != nil {
		t.Fatal(err)
	}
	if got := r.Entitlement.QuantityRates["output_tokens"]; got != 0.01 {
		t.Fatalf("result entitlement quantity_rates = %v", r.Entitlement.QuantityRates)
	}
	out, err = json.Marshal(&r)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), `"quantity_rates":{"output_tokens":0.01}`) {
		t.Errorf("result did not re-marshal quantity_rates as snake_case: %s", out)
	}
}

// TestEventQuantitiesWireShape pins the options member the engine reads.
func TestEventQuantitiesWireShape(t *testing.T) {
	options := newCheckFlagOptions()
	WithEventQuantities("chat", 0, map[string]float64{"input_tokens": 1000})(options)
	out, err := json.Marshal(options)
	if err != nil {
		t.Fatal(err)
	}
	if want := `{"event_quantities":{"event_subtype":"chat","quantities":{"input_tokens":1000}}}`; string(out) != want {
		t.Errorf("options = %s, want %s", out, want)
	}
	if options.isZero() {
		t.Error("an event_quantities preflight must not read as no options")
	}

	options = newCheckFlagOptions()
	WithEventQuantities("chat", 2, nil)(options)
	out, err = json.Marshal(options)
	if err != nil {
		t.Fatal(err)
	}
	if want := `{"event_quantities":{"event_subtype":"chat","quantity":2}}`; string(out) != want {
		t.Errorf("options = %s, want %s", out, want)
	}
}
