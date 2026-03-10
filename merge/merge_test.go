package merge

import (
	"encoding/json"
	"testing"

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

	// Traits updated
	assert.Len(t, merged.Traits, 1)
	assert.Equal(t, "Startup", merged.Traits[0].Value)

	// Other fields unchanged
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
	// Original traits still present
	assert.Len(t, merged.Traits, 1)
}

func TestPartialCompany_EmptyEntitlements(t *testing.T) {
	existing := baseCompany()
	partial := json.RawMessage(`{"id":"co-1","entitlements":[]}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	assert.Empty(t, merged.Entitlements)
	// Other fields unchanged
	assert.Equal(t, "acc-1", merged.AccountID)
}

func TestPartialCompany_NullBasePlanID(t *testing.T) {
	existing := baseCompany()
	partial := json.RawMessage(`{"id":"co-1","base_plan_id":null}`)

	merged, err := PartialCompany(existing, partial)
	require.NoError(t, err)

	assert.Nil(t, merged.BasePlanID)
	// Other fields unchanged
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

	// Original not mutated
	assert.Equal(t, origKeys, existing.Keys)
	assert.Len(t, existing.Traits, 1)

	// Merged is different
	assert.Equal(t, map[string]string{"slug": "new-slug"}, merged.Keys)
	assert.Empty(t, merged.Traits)
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
