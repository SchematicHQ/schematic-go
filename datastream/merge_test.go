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

func TestPartialCompany_OnlyTraits(t *testing.T) {
	existing := baseCompany()
	partial := json.RawMessage(`{"id":"co-1","traits":[{"value":"Startup","trait_definition":{"id":"plan"}}]}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	assert.Len(t, merged.Traits, 1)
	assert.Equal(t, "Startup", merged.Traits[0].Value)

	assert.Equal(t, "acc-1", merged.AccountID)
	assert.Equal(t, "env-1", merged.EnvironmentID)
	assert.Equal(t, map[string]string{"domain": "example.com"}, merged.Keys)
	assert.Equal(t, []string{"bp-1"}, merged.BillingProductIDs)
	assert.NotNil(t, merged.BasePlanID)
	assert.Equal(t, "plan-1", *merged.BasePlanID)
}

func TestPartialCompany_ReplacesKeys(t *testing.T) {
	existing := baseCompany()
	partial := json.RawMessage(`{"id":"co-1","keys":{"slug":"new-slug"}}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	assert.Equal(t, map[string]string{"slug": "new-slug"}, merged.Keys)
	assert.Len(t, merged.Traits, 1)
}

func TestPartialCompany_EmptyEntitlements(t *testing.T) {
	existing := baseCompany()
	partial := json.RawMessage(`{"id":"co-1","entitlements":[]}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	assert.Empty(t, merged.Entitlements)
	assert.Equal(t, "acc-1", merged.AccountID)
}

func TestPartialCompany_NullBasePlanID(t *testing.T) {
	existing := baseCompany()
	partial := json.RawMessage(`{"id":"co-1","base_plan_id":null}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	assert.Nil(t, merged.BasePlanID)
	assert.Equal(t, []string{"bp-1"}, merged.BillingProductIDs)
}

func TestPartialCompany_MissingID(t *testing.T) {
	existing := baseCompany()
	partial := json.RawMessage(`{"traits":[]}`)

	_, err := PartialCompany(existing, partial)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "missing required field: id")
}

func TestPartialCompany_DoesNotMutateOriginal(t *testing.T) {
	existing := baseCompany()
	origKeys := make(map[string]string)
	for k, v := range existing.Keys {
		origKeys[k] = v
	}

	partial := json.RawMessage(`{"id":"co-1","keys":{"slug":"new-slug"},"traits":[]}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	assert.Equal(t, origKeys, existing.Keys)
	assert.Len(t, existing.Traits, 1)

	assert.Equal(t, map[string]string{"slug": "new-slug"}, merged.Keys)
	assert.Empty(t, merged.Traits)
}

func TestPartialCompany_Rules(t *testing.T) {
	existing := baseCompany()
	existing.Rules = []*rulesengine.Rule{{ID: "rule-old"}}
	partial := json.RawMessage(`{"id":"co-1","rules":[{"id":"rule-new"}]}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	require.Len(t, merged.Rules, 1)
	assert.Equal(t, "rule-new", merged.Rules[0].ID)
	assert.Equal(t, "rule-old", existing.Rules[0].ID)
}

func TestPartialUser_OnlyTraits(t *testing.T) {
	existing := baseUser()
	partial := json.RawMessage(`{"id":"user-1","traits":[{"value":"Free","trait_definition":{"id":"tier"}}]}`)

	merged, err := PartialUser(existing, partial)
	require.NoError(t, err)

	assert.Len(t, merged.Traits, 1)
	assert.Equal(t, "Free", merged.Traits[0].Value)
	assert.Equal(t, map[string]string{"email": "user@example.com"}, merged.Keys)
}

func TestPartialUser_ReplacesKeys(t *testing.T) {
	existing := baseUser()
	partial := json.RawMessage(`{"id":"user-1","keys":{"email":"new@example.com"}}`)

	merged, err := PartialUser(existing, partial)
	require.NoError(t, err)

	assert.Equal(t, map[string]string{"email": "new@example.com"}, merged.Keys)
	assert.Len(t, merged.Traits, 1)
}

func TestPartialUser_MissingID(t *testing.T) {
	existing := baseUser()
	partial := json.RawMessage(`{"keys":{"email":"new@example.com"}}`)

	_, err := PartialUser(existing, partial)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "missing required field: id")
}

func TestPartialUser_DoesNotMutateOriginal(t *testing.T) {
	existing := baseUser()
	origKeys := make(map[string]string)
	for k, v := range existing.Keys {
		origKeys[k] = v
	}

	partial := json.RawMessage(`{"id":"user-1","keys":{"slug":"new"},"traits":[]}`)

	merged, err := PartialUser(existing, partial)
	require.NoError(t, err)

	assert.Equal(t, origKeys, existing.Keys)
	assert.Len(t, existing.Traits, 1)

	assert.Equal(t, map[string]string{"slug": "new"}, merged.Keys)
	assert.Empty(t, merged.Traits)
}

func TestExtractIDFromJSON(t *testing.T) {
	id, err := ExtractIDFromJSON(json.RawMessage(`{"id":"co-1","traits":[]}`))
	require.NoError(t, err)
	assert.Equal(t, "co-1", id)
}

func TestExtractIDFromJSON_Missing(t *testing.T) {
	_, err := ExtractIDFromJSON(json.RawMessage(`{"traits":[]}`))
	assert.Error(t, err)
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
	assert.Equal(t, int64(42), cp.Metrics[0].Value)
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
