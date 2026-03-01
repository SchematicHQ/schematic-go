package client

import (
	"testing"
	"time"

	"github.com/schematichq/rulesengine"
	schematicgo "github.com/schematichq/schematic-go"
	"github.com/stretchr/testify/assert"
)

func TestToRulesEngineEntitlement(t *testing.T) {
	t.Run("nil input", func(t *testing.T) {
		assert.Nil(t, toRulesEngineEntitlement(nil))
	})

	t.Run("all fields populated", func(t *testing.T) {
		allocation := 100
		softLimit := 200
		usage := 50
		eventName := "api-calls"
		metricPeriod := schematicgo.FeatureEntitlementMetricPeriodCurrentMonth
		monthReset := schematicgo.FeatureEntitlementMonthResetBillingCycle
		resetAt := time.Now().UTC()
		creditID := "cred-123"
		creditTotal := 1000.0
		creditUsed := 250.0
		creditRemaining := 750.0

		input := &schematicgo.FeatureEntitlement{
			FeatureID:       "feat-123",
			FeatureKey:      "test-feature",
			ValueType:       schematicgo.EntitlementValueTypeNumeric,
			Allocation:      &allocation,
			SoftLimit:       &softLimit,
			Usage:           &usage,
			EventName:       &eventName,
			MetricPeriod:    &metricPeriod,
			MonthReset:      &monthReset,
			MetricResetAt:   &resetAt,
			CreditID:        &creditID,
			CreditTotal:     &creditTotal,
			CreditUsed:      &creditUsed,
			CreditRemaining: &creditRemaining,
		}

		result := toRulesEngineEntitlement(input)

		assert.Equal(t, "feat-123", result.FeatureID)
		assert.Equal(t, "test-feature", result.FeatureKey)
		assert.Equal(t, rulesengine.EntitlementValueType("numeric"), result.ValueType)
		assert.EqualValues(t, 100, *result.Allocation)
		assert.EqualValues(t, 200, *result.SoftLimit)
		assert.EqualValues(t, 50, *result.Usage)
		assert.Equal(t, "api-calls", *result.EventName)
		assert.Equal(t, rulesengine.MetricPeriod("current_month"), *result.MetricPeriod)
		assert.Equal(t, rulesengine.MetricPeriodMonthReset("billing_cycle"), *result.MonthReset)
		assert.Equal(t, resetAt, *result.MetricResetAt)
		assert.Equal(t, "cred-123", *result.CreditID)
		assert.Equal(t, 1000.0, *result.CreditTotal)
		assert.Equal(t, 250.0, *result.CreditUsed)
		assert.Equal(t, 750.0, *result.CreditRemaining)
	})

	t.Run("optional fields nil", func(t *testing.T) {
		input := &schematicgo.FeatureEntitlement{
			FeatureID:  "feat-456",
			FeatureKey: "boolean-feature",
			ValueType:  schematicgo.EntitlementValueTypeBoolean,
		}

		result := toRulesEngineEntitlement(input)

		assert.Equal(t, "feat-456", result.FeatureID)
		assert.Equal(t, "boolean-feature", result.FeatureKey)
		assert.Equal(t, rulesengine.EntitlementValueType("boolean"), result.ValueType)
		assert.Nil(t, result.Allocation)
		assert.Nil(t, result.SoftLimit)
		assert.Nil(t, result.Usage)
		assert.Nil(t, result.EventName)
		assert.Nil(t, result.MetricPeriod)
		assert.Nil(t, result.MonthReset)
		assert.Nil(t, result.MetricResetAt)
		assert.Nil(t, result.CreditID)
		assert.Nil(t, result.CreditTotal)
		assert.Nil(t, result.CreditUsed)
		assert.Nil(t, result.CreditRemaining)
	})
}

func TestToRulesEngineRuleType(t *testing.T) {
	t.Run("nil input", func(t *testing.T) {
		assert.Nil(t, toRulesEngineRuleType(nil))
	})

	t.Run("valid input", func(t *testing.T) {
		s := "plan_entitlement"
		result := toRulesEngineRuleType(&s)
		assert.Equal(t, rulesengine.RuleType("plan_entitlement"), *result)
	})
}

func TestToCheckFlagResponse(t *testing.T) {
	t.Run("nil input", func(t *testing.T) {
		assert.Nil(t, toCheckFlagResponse(nil))
	})

	t.Run("populated input", func(t *testing.T) {
		companyID := "comp-123"
		flagID := "flag-456"
		ruleID := "rule-789"
		userID := "user-321"
		ruleType := rulesengine.RuleTypePlanEntitlement

		input := &rulesengine.CheckFlagResult{
			CompanyID: &companyID,
			FlagID:    &flagID,
			FlagKey:   "test-flag",
			Reason:    "matched",
			RuleID:    &ruleID,
			RuleType:  &ruleType,
			UserID:    &userID,
			Value:     true,
		}

		result := toCheckFlagResponse(input)

		assert.Equal(t, "comp-123", *result.CompanyID)
		assert.Equal(t, "flag-456", *result.FlagID)
		assert.Equal(t, "test-flag", result.FlagKey)
		assert.Equal(t, "matched", result.Reason)
		assert.Equal(t, "rule-789", *result.RuleID)
		assert.Equal(t, rulesengine.RuleType("plan_entitlement"), *result.RuleType)
		assert.Equal(t, "user-321", *result.UserID)
		assert.True(t, result.Value)
	})
}
