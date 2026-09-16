package datastream

import (
	"context"
	"testing"

	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/rulesengine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// creditTestFlag builds a flag whose single rule is a credit-balance condition,
// so a preflight that prices the action above the balance flips the verdict from
// the rule's true to the flag's false default.
func creditTestFlag(creditID string, consumptionRate float64, eventSubtype *string) *rulesengine.Flag {
	condition := &rulesengine.Condition{
		ID:              "cond_preflight",
		AccountID:       "acct_preflight",
		EnvironmentID:   "env_preflight",
		ConditionType:   rulesengine.ConditionTypeCredit,
		Operator:        rulesengine.ComparableOperatorLt,
		CreditID:        &creditID,
		ConsumptionRate: &consumptionRate,
		EventSubtype:    eventSubtype,
	}

	return &rulesengine.Flag{
		ID:            "flag_preflight",
		AccountID:     "acct_preflight",
		EnvironmentID: "env_preflight",
		Key:           "preflight-flag",
		DefaultValue:  false,
		Rules: rulesengine.JSONSlice[*rulesengine.Rule]{
			{
				ID:            "rule_preflight",
				AccountID:     "acct_preflight",
				EnvironmentID: "env_preflight",
				RuleType:      rulesengine.RuleTypeStandard,
				Name:          "credit balance",
				Priority:      1,
				Value:         true,
				Conditions:    rulesengine.JSONSlice[*rulesengine.Condition]{condition},
			},
		},
	}
}

func creditTestCompany(creditID string, balance float64) *rulesengine.Company {
	return &rulesengine.Company{
		ID:             "comp_preflight",
		AccountID:      "acct_preflight",
		EnvironmentID:  "env_preflight",
		CreditBalances: map[string]float64{creditID: balance},
	}
}

func TestPreflightCheckFlagOptions(t *testing.T) {
	ctx := context.Background()
	engine, err := rulesengine.NewEngine(ctx)
	require.NoError(t, err)
	defer func() { _ = engine.Close(ctx) }()

	creditID := "credit-abc"

	// checkWith evaluates the flag through the translated options, the way
	// DataStreamClient.CheckFlag does.
	checkWith := func(t *testing.T, flag *rulesengine.Flag, company *rulesengine.Company, preflight *schematicgo.PreflightRequestBody) bool {
		t.Helper()
		result, err := engine.CheckFlag(ctx, company, nil, flag, preflightCheckFlagOptions(preflight)...)
		require.NoError(t, err)
		return result.Value
	}

	t.Run("No preflight yields no options", func(t *testing.T) {
		assert.Nil(t, preflightCheckFlagOptions(nil))
		assert.Empty(t, preflightCheckFlagOptions(&schematicgo.PreflightRequestBody{}))
	})

	t.Run("Usage is priced at the consumption rate", func(t *testing.T) {
		flag := creditTestFlag(creditID, 0.0001, nil)
		company := creditTestCompany(creditID, 1.0)

		// 50 x 0.0001 = 0.005 credits against a balance of 1.
		assert.True(t, checkWith(t, flag, company, &schematicgo.PreflightRequestBody{Usage: usageOf(50)}))
		// 20000 x 0.0001 = 2 credits against a balance of 1.
		assert.False(t, checkWith(t, flag, company, &schematicgo.PreflightRequestBody{Usage: usageOf(20_000)}))
	})

	t.Run("Event usage only moves the matching subtype", func(t *testing.T) {
		eventSubtype := "api-calls"
		flag := creditTestFlag(creditID, 0.05, &eventSubtype)
		company := creditTestCompany(creditID, 10.0)

		eventUsage := func(subtype string, quantity int64) *schematicgo.PreflightRequestBody {
			return &schematicgo.PreflightRequestBody{
				EventUsage: &schematicgo.PreflightEventUsageRequestBody{EventSubtype: subtype, Quantity: quantity},
			}
		}

		// 100 x 0.05 = 5 credits against a balance of 10.
		assert.True(t, checkWith(t, flag, company, eventUsage(eventSubtype, 100)))
		// 1,000,000 x 0.05 is far past the balance.
		assert.False(t, checkWith(t, flag, company, eventUsage(eventSubtype, 1_000_000)))
		// A different subtype leaves this condition alone.
		assert.True(t, checkWith(t, flag, company, eventUsage("other-events", 1_000_000)))
	})

	t.Run("Credit cost is gated as given", func(t *testing.T) {
		flag := creditTestFlag(creditID, 1.0, nil)
		company := creditTestCompany(creditID, 100.0)

		creditCost := func(cost float64) *schematicgo.PreflightRequestBody {
			return &schematicgo.PreflightRequestBody{CreditCost: map[string]float64{creditID: cost}}
		}

		assert.True(t, checkWith(t, flag, company, creditCost(50)))
		assert.False(t, checkWith(t, flag, company, creditCost(500)))
		// A cost keyed to another credit type does not touch this condition.
		assert.True(t, checkWith(t, flag, company, &schematicgo.PreflightRequestBody{
			CreditCost: map[string]float64{"credit-other": 9999},
		}))
	})

	t.Run("Negative usage reaches the engine's validation", func(t *testing.T) {
		flag := creditTestFlag(creditID, 1.0, nil)
		company := creditTestCompany(creditID, 100.0)

		_, err := engine.CheckFlag(ctx, company, nil, flag, preflightCheckFlagOptions(
			&schematicgo.PreflightRequestBody{Usage: usageOf(-1)},
		)...)

		assert.Equal(t, rulesengine.ErrorNegativePreflightUsage, err)
	})
}

func usageOf(quantity int64) *int64 { return &quantity }
