package rulesengine

import (
	"encoding/json"
	"strings"
	"testing"
)

// The bulk envelope carries the same wire types as the single one, so it
// inherits the same hazard: the engine's Rust models fill an absent key via
// serde's #[serde(default)], which does not cover a key explicitly set to null.
// `"rules": null` fails to deserialize into Vec<Rule>, and the whole check
// comes back -1. These pin that no collection in a bulk envelope can reach the
// engine as null, including the flag list itself, which is new here.

// collectionNulls are the fragments that would sink an envelope. Nulls on
// Option-typed scalars (base_plan_id, subscription) are fine -- serde accepts
// null for Option<T>.
var collectionNulls = []string{
	`"flags":null`, `"rules":null`, `"traits":null`, `"metrics":null`, `"keys":null`,
	`"billing_product_ids":null`, `"plan_ids":null`, `"plan_version_ids":null`,
	`"credit_balances":null`, `"credit_postpaid":null`, `"entitlements":null`,
	`"conditions":null`, `"condition_groups":null`, `"resource_ids":null`,
}

func assertNoCollectionNulls(t *testing.T, raw []byte) {
	t.Helper()

	for _, frag := range collectionNulls {
		if strings.Contains(string(raw), frag) {
			t.Errorf("collection serialized as null: %s", frag)
		}
	}
}

// TestFlagsEnvelopeHasNoNulls is TestEnvelopeHasNoNulls for the bulk envelope:
// worst case, every optional field left nil.
func TestFlagsEnvelopeHasNoNulls(t *testing.T) {
	env := &checkFlagsEnvelope{
		Flags: JSONSlice[*Flag]{
			{ID: "flag-1", AccountID: "a", EnvironmentID: "e", Key: "k"},
			{ID: "flag-2", AccountID: "a", EnvironmentID: "e", Key: "k2"},
		},
		Company: &Company{ID: "comp-1", AccountID: "a", EnvironmentID: "e"},
		User:    &User{ID: "user-1", AccountID: "a", EnvironmentID: "e"},
	}

	raw, err := marshalFlagsEnvelope(env)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("bulk marshal:\n  %s", raw)
	assertNoCollectionNulls(t, raw)
}

// TestFlagsEnvelopeNilFlagsIsEmptyArray covers the field the bulk envelope adds.
// A nil Flags must cross as [], not null: the engine reads "flags" into a
// Vec<Flag>, where null is an error rather than an empty list.
func TestFlagsEnvelopeNilFlagsIsEmptyArray(t *testing.T) {
	raw, err := marshalFlagsEnvelope(&checkFlagsEnvelope{
		Company: &Company{ID: "comp-1", AccountID: "a", EnvironmentID: "e"},
	})
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(raw), `"flags":[]`) {
		t.Errorf("nil flags should marshal as [], got: %s", raw)
	}
	assertNoCollectionNulls(t, raw)
}

// TestFlagsEnvelopeNestedHasNoNulls covers the invariant at depth, where a nil
// collection hides inside a flag's rule or the company's metrics rather than on
// the top-level entity. NewJSONSlice is deliberately not used on the inner
// slices, so this fails if a wire field is ever declared []T.
func TestFlagsEnvelopeNestedHasNoNulls(t *testing.T) {
	env := &checkFlagsEnvelope{
		Flags: JSONSlice[*Flag]{
			{
				ID: "flag-1", AccountID: "a", EnvironmentID: "e", Key: "k",
				Rules: JSONSlice[*Rule]{
					{ID: "rule-1", AccountID: "a", EnvironmentID: "e", RuleType: RuleTypeStandard},
				},
			},
		},
		Company: &Company{
			ID: "comp-1", AccountID: "a", EnvironmentID: "e",
			Metrics: CompanyMetricCollection{{AccountID: "a", EnvironmentID: "e", CompanyID: "comp-1"}},
		},
	}

	raw, err := marshalFlagsEnvelope(env)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("nested bulk marshal:\n  %s", raw)
	assertNoCollectionNulls(t, raw)
}

// TestFlagsEnvelopeOmitsAbsentContext pins that a nil company or user is left
// out rather than sent as null. The engine treats an absent and a null member
// the same, so this is about not relying on that: every other member of the
// envelope is a required key.
func TestFlagsEnvelopeOmitsAbsentContext(t *testing.T) {
	raw, err := marshalFlagsEnvelope(&checkFlagsEnvelope{
		Flags: JSONSlice[*Flag]{{ID: "flag-1", AccountID: "a", EnvironmentID: "e", Key: "k"}},
	})
	if err != nil {
		t.Fatal(err)
	}

	var decoded map[string]json.RawMessage
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatal(err)
	}

	if _, ok := decoded["company"]; ok {
		t.Errorf("a nil company should be omitted, got: %s", raw)
	}
	if _, ok := decoded["user"]; ok {
		t.Errorf("a nil user should be omitted, got: %s", raw)
	}
	if _, ok := decoded["options"]; ok {
		t.Errorf("the bulk envelope carries no options, got: %s", raw)
	}
}

// TestFlagsEnvelopeMatchesSingleShape pins that the members shared with the
// single envelope serialize identically. They are the same types, so a
// divergence would mean one of the two grew its own copy.
func TestFlagsEnvelopeMatchesSingleShape(t *testing.T) {
	flag := &Flag{ID: "flag-1", AccountID: "a", EnvironmentID: "e", Key: "k"}
	company := &Company{ID: "comp-1", AccountID: "a", EnvironmentID: "e"}
	user := &User{ID: "user-1", AccountID: "a", EnvironmentID: "e"}

	single, err := marshalEnvelope(&checkFlagEnvelope{Flag: flag, Company: company, User: user})
	if err != nil {
		t.Fatal(err)
	}
	bulk, err := marshalFlagsEnvelope(&checkFlagsEnvelope{
		Flags: JSONSlice[*Flag]{flag}, Company: company, User: user,
	})
	if err != nil {
		t.Fatal(err)
	}

	var singleEnv, bulkEnv map[string]json.RawMessage
	if err := json.Unmarshal(single, &singleEnv); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(bulk, &bulkEnv); err != nil {
		t.Fatal(err)
	}

	for _, key := range []string{"company", "user"} {
		if string(singleEnv[key]) != string(bulkEnv[key]) {
			t.Errorf("%s differs between envelopes:\n  single: %s\n  bulk:   %s",
				key, singleEnv[key], bulkEnv[key])
		}
	}

	if want := "[" + string(singleEnv["flag"]) + "]"; string(bulkEnv["flags"]) != want {
		t.Errorf("flags should be the single envelope's flag in an array:\n  got:  %s\n  want: %s",
			bulkEnv["flags"], want)
	}
}
