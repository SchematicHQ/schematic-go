package rulesengine_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/schematichq/schematic-go/rulesengine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const inferenceSubtype = "chat"

// inferenceFlag wraps a single credit-balance rule priced like an inference
// entitlement: requests at consumptionRate, tokens at their own rates.
func inferenceFlag(creditID string, consumptionRate float64, quantityRates map[string]float64) (*rulesengine.Flag, *rulesengine.Rule) {
	rule := createTestRule()
	condition := createTestCondition(rulesengine.ConditionTypeCredit)
	condition.Operator = rulesengine.ComparableOperatorLt
	condition.CreditID = &creditID
	condition.ConsumptionRate = ptr(consumptionRate)
	condition.EventSubtype = ptr(inferenceSubtype)
	condition.QuantityRates = quantityRates
	rule.Conditions = []*rulesengine.Condition{condition}

	flag := createTestFlag()
	flag.DefaultValue = false
	flag.Rules = []*rulesengine.Rule{rule}
	return flag, rule
}

func companyWithBalance(creditID string, balance float64) *rulesengine.Company {
	company := createTestCompany()
	company.CreditBalances = map[string]float64{creditID: balance}
	return company
}

func TestWithEventQuantities(t *testing.T) {
	ctx := context.Background()
	engine := newTestEngine(t)

	const creditID = "credit-abc"
	rates := map[string]float64{"input_tokens": 0.001, "output_tokens": 0.01}
	// 1 request × 0.5 + (1000 − 400 cached) × 0.001 + 100 × 0.01 = 2.1. The
	// cached tokens are unrated, so they cost nothing but still come out of
	// input.
	call := rulesengine.WithEventQuantities(inferenceSubtype, 0, map[string]float64{
		"input_tokens":        1000,
		"cached_input_tokens": 400,
		"output_tokens":       100,
	})

	t.Run("passes when the balance covers the call", func(t *testing.T) {
		flag, rule := inferenceFlag(creditID, 0.5, rates)

		result, err := engine.CheckFlag(ctx, companyWithBalance(creditID, 2.1), nil, flag, call)

		require.NoError(t, err)
		assert.Equal(t, &rule.ID, result.RuleID)
		assert.True(t, result.Value)
	})

	t.Run("refuses when the balance falls short", func(t *testing.T) {
		// Covers the request and the legacy single unit, not the tokens.
		flag, _ := inferenceFlag(creditID, 0.5, rates)

		result, err := engine.CheckFlag(ctx, companyWithBalance(creditID, 2.0), nil, flag, call)

		require.NoError(t, err)
		assert.Nil(t, result.RuleID)
		assert.False(t, result.Value)
	})

	t.Run("ignored for another subtype", func(t *testing.T) {
		flag, rule := inferenceFlag(creditID, 0.5, rates)
		other := rulesengine.WithEventQuantities("other", 0, map[string]float64{"input_tokens": 1e6})

		result, err := engine.CheckFlag(ctx, companyWithBalance(creditID, 1.0), nil, flag, other)

		require.NoError(t, err)
		assert.Equal(t, &rule.ID, result.RuleID)
	})

	t.Run("quantity multiplies the base, not the quantities", func(t *testing.T) {
		// 3 × 0.5 + 600 × 0.001 + 100 × 0.01 = 3.1.
		flag, rule := inferenceFlag(creditID, 0.5, rates)
		opt := rulesengine.WithEventQuantities(inferenceSubtype, 3, map[string]float64{
			"input_tokens":        1000,
			"cached_input_tokens": 400,
			"output_tokens":       100,
		})

		result, err := engine.CheckFlag(ctx, companyWithBalance(creditID, 3.1), nil, flag, opt)
		require.NoError(t, err)
		assert.Equal(t, &rule.ID, result.RuleID)

		result, err = engine.CheckFlag(ctx, companyWithBalance(creditID, 3.0), nil, flag, opt)
		require.NoError(t, err)
		assert.Nil(t, result.RuleID)
	})

	t.Run("credit_cost beats event_quantities", func(t *testing.T) {
		flag, rule := inferenceFlag(creditID, 0.5, rates)

		result, err := engine.CheckFlag(ctx, companyWithBalance(creditID, 1.0), nil, flag,
			call, rulesengine.WithCreditCost(creditID, 1.0))

		require.NoError(t, err)
		assert.Equal(t, &rule.ID, result.RuleID)
	})

	t.Run("event_quantities beats event_usage", func(t *testing.T) {
		// event_usage alone would ask 1 × 0.5, which 2.0 covers.
		flag, _ := inferenceFlag(creditID, 0.5, rates)

		result, err := engine.CheckFlag(ctx, companyWithBalance(creditID, 2.0), nil, flag,
			rulesengine.WithEventUsage(inferenceSubtype, 1), call)

		require.NoError(t, err)
		assert.Nil(t, result.RuleID)
	})

	t.Run("later call replaces the earlier one", func(t *testing.T) {
		flag, rule := inferenceFlag(creditID, 0.5, rates)
		huge := rulesengine.WithEventQuantities(inferenceSubtype, 0, map[string]float64{"output_tokens": 1e6})

		result, err := engine.CheckFlag(ctx, companyWithBalance(creditID, 2.1), nil, flag, huge, call)

		require.NoError(t, err)
		assert.Equal(t, &rule.ID, result.RuleID)
	})

	t.Run("rejects negative values", func(t *testing.T) {
		flag, _ := inferenceFlag(creditID, 0.5, rates)
		for name, opt := range map[string]rulesengine.CheckFlagOption{
			"quantity":   rulesengine.WithEventQuantities(inferenceSubtype, -1, nil),
			"quantities": rulesengine.WithEventQuantities(inferenceSubtype, 1, map[string]float64{"input_tokens": -1}),
		} {
			t.Run(name, func(t *testing.T) {
				result, err := engine.CheckFlag(ctx, companyWithBalance(creditID, 100), nil, flag, opt)

				assert.Equal(t, rulesengine.ErrorNegativePreflightUsage, err)
				assert.Equal(t, rulesengine.ErrorNegativePreflightUsage, result.Err)
			})
		}
	})

	t.Run("the option keeps its own copy of quantities", func(t *testing.T) {
		flag, rule := inferenceFlag(creditID, 0.5, rates)
		quantities := map[string]float64{"output_tokens": 100}
		opt := rulesengine.WithEventQuantities(inferenceSubtype, 0, quantities)
		quantities["output_tokens"] = 1e6

		// 0.5 + 100 × 0.01 = 1.5.
		result, err := engine.CheckFlag(ctx, companyWithBalance(creditID, 1.5), nil, flag, opt)

		require.NoError(t, err)
		assert.Equal(t, &rule.ID, result.RuleID)
	})
}

// TestEntitlementQuantityRatesThroughEngine pins that a company entitlement's
// quantity_rates survives the trip into the engine and back onto the result.
func TestEntitlementQuantityRatesThroughEngine(t *testing.T) {
	ctx := context.Background()
	engine := newTestEngine(t)

	flag := createTestFlag()
	company := createTestCompany()
	company.Entitlements = rulesengine.JSONSlice[*rulesengine.FeatureEntitlement]{{
		FeatureID:     "feat-1",
		FeatureKey:    flag.Key,
		ValueType:     rulesengine.EntitlementValueTypeCredit,
		QuantityRates: map[string]float64{"input_tokens": 0.001, "output_tokens": 0.01},
	}}

	result, err := engine.CheckFlag(ctx, company, nil, flag)

	require.NoError(t, err)
	require.NotNil(t, result.Entitlement)
	assert.Equal(t, map[string]float64{"input_tokens": 0.001, "output_tokens": 0.01}, result.Entitlement.QuantityRates)
}

// quantityCostCase is one case of testdata/quantity_cost.json, copied verbatim
// from schematic-api's api/lib/rulesengine/testdata/quantity_cost.json. The API's
// burn and the engine's quantity_cost both price every case there; running them
// through the WebAssembly engine here pins that this SDK's wire shape for
// quantity_rates and event_quantities reaches that pricing intact.
type quantityCostCase struct {
	Name            string             `json:"name"`
	ConsumptionRate float64            `json:"consumption_rate"`
	QuantityRates   map[string]float64 `json:"quantity_rates"`
	Quantity        float64            `json:"quantity"`
	Quantities      map[string]float64 `json:"quantities"`
	ExpectedCost    float64            `json:"expected_cost"`
}

func TestWithEventQuantitiesSharedFixture(t *testing.T) {
	ctx := context.Background()
	engine := newTestEngine(t)

	raw, err := os.ReadFile(filepath.Join("testdata", "quantity_cost.json"))
	require.NoError(t, err)
	var fixture struct {
		Cases []quantityCostCase `json:"cases"`
	}
	require.NoError(t, json.Unmarshal(raw, &fixture))
	require.NotEmpty(t, fixture.Cases)

	const creditID = "credit-abc"
	for _, tc := range fixture.Cases {
		t.Run(tc.Name, func(t *testing.T) {
			if hasNegative(tc) {
				t.Skip("the preflight rejects negative quantities before pricing")
			}

			flag, rule := inferenceFlag(creditID, tc.ConsumptionRate, tc.QuantityRates)
			opt := rulesengine.WithEventQuantities(inferenceSubtype, tc.Quantity, tc.Quantities)

			// A cost priced to zero gates on balance > 0, so the smallest
			// positive balance passes and zero does not.
			covers, short := tc.ExpectedCost*(1+1e-9)+1e-9, tc.ExpectedCost*(1-1e-6)
			if tc.ExpectedCost == 0 {
				short = 0
			}

			result, err := engine.CheckFlag(ctx, companyWithBalance(creditID, covers), nil, flag, opt)
			require.NoError(t, err)
			assert.Equal(t, &rule.ID, result.RuleID, "balance %v should cover cost %v", covers, tc.ExpectedCost)

			result, err = engine.CheckFlag(ctx, companyWithBalance(creditID, short), nil, flag, opt)
			require.NoError(t, err)
			assert.Nil(t, result.RuleID, "balance %v should not cover cost %v", short, tc.ExpectedCost)
		})
	}
}

func hasNegative(tc quantityCostCase) bool {
	if tc.Quantity < 0 {
		return true
	}
	for _, q := range tc.Quantities {
		if q < 0 {
			return true
		}
	}
	return false
}
