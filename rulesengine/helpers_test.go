package rulesengine_test

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/schematichq/schematic-go/rulesengine"
)

// Fixture builders ported from the Go rules engine's own test suite, so the
// evaluation tests in flagcheck_test.go can be compared against their originals
// with as little translation as possible.
//
// The originals randomized every field with gofakeit. These use a counter
// instead: a failure should always reproduce, and randomizing IDs bought the
// original suite nothing that a unique counter does not.

var idCounter atomic.Int64

func generateTestID(prefix string) string {
	return fmt.Sprintf("%s_%d", prefix, idCounter.Add(1))
}

func ptr[T any](v T) *T {
	return &v
}

func createTestCompany() *rulesengine.Company {
	return &rulesengine.Company{
		ID:                generateTestID("comp"),
		AccountID:         generateTestID("acct"),
		EnvironmentID:     generateTestID("env"),
		PlanIDs:           []string{generateTestID("plan"), generateTestID("plan")},
		PlanVersionIDs:    []string{generateTestID("plnv"), generateTestID("plnv")},
		BillingProductIDs: []string{generateTestID("bilp"), generateTestID("bilp")},
		BasePlanID:        ptr(generateTestID("plan")),
		Metrics:           make(rulesengine.CompanyMetricCollection, 0),
		Traits:            make(rulesengine.JSONSlice[*rulesengine.Trait], 0),
		Subscription: &rulesengine.Subscription{
			ID:          generateTestID("bilsub"),
			PeriodStart: time.Now().Add(-30 * 24 * time.Hour),
			PeriodEnd:   time.Now().Add(30 * 24 * time.Hour),
		},
	}
}

func createTestUser() *rulesengine.User {
	return &rulesengine.User{
		ID:            generateTestID("user"),
		AccountID:     generateTestID("acct"),
		EnvironmentID: generateTestID("env"),
		Traits:        make(rulesengine.JSONSlice[*rulesengine.Trait], 0),
	}
}

func createTestRule() *rulesengine.Rule {
	return &rulesengine.Rule{
		ID:              generateTestID("rule"),
		AccountID:       generateTestID("acct"),
		EnvironmentID:   generateTestID("env"),
		RuleType:        rulesengine.RuleTypeStandard,
		Name:            generateTestID("name"),
		Priority:        1,
		Conditions:      make(rulesengine.JSONSlice[*rulesengine.Condition], 0),
		ConditionGroups: make(rulesengine.JSONSlice[*rulesengine.ConditionGroup], 0),
		Value:           true,
	}
}

func createTestFlag() *rulesengine.Flag {
	return &rulesengine.Flag{
		ID:            generateTestID("flag"),
		AccountID:     generateTestID("acct"),
		EnvironmentID: generateTestID("env"),
		Key:           generateTestID("key"),
		Rules:         make(rulesengine.JSONSlice[*rulesengine.Rule], 0),
		DefaultValue:  false,
	}
}

func createTestCondition(conditionType rulesengine.ConditionType) *rulesengine.Condition {
	condition := &rulesengine.Condition{
		ID:            generateTestID("cond"),
		AccountID:     generateTestID("acct"),
		EnvironmentID: generateTestID("env"),
		ConditionType: conditionType,
		Operator:      rulesengine.ComparableOperatorEquals,
	}

	switch conditionType {
	case rulesengine.ConditionTypeMetric:
		condition.EventSubtype = ptr(generateTestID("evt"))
		condition.MetricValue = ptr(int64(100))
		condition.MetricPeriod = ptr(rulesengine.MetricPeriodAllTime)
		condition.MetricPeriodMonthReset = ptr(rulesengine.MetricPeriodMonthResetFirst)
	case rulesengine.ConditionTypeTrait:
		condition.TraitDefinition = createTestTraitDefinition(rulesengine.ComparableTypeInt, rulesengine.EntityTypeCompany)
		condition.TraitValue = generateTestID("val")
	}

	return condition
}

func createTestMetric(
	company *rulesengine.Company,
	eventSubtype string,
	period rulesengine.MetricPeriod,
	value int64,
) *rulesengine.CompanyMetric {
	return &rulesengine.CompanyMetric{
		AccountID:     company.AccountID,
		EnvironmentID: company.EnvironmentID,
		CompanyID:     company.ID,
		EventSubtype:  eventSubtype,
		Period:        period,
		MonthReset:    rulesengine.MetricPeriodMonthResetFirst,
		Value:         value,
		CreatedAt:     time.Now(),
	}
}

func createTestTrait(value string, def *rulesengine.TraitDefinition) *rulesengine.Trait {
	if def == nil {
		def = createTestTraitDefinition(rulesengine.ComparableTypeInt, rulesengine.EntityTypeCompany)
	}

	return &rulesengine.Trait{
		TraitDefinition: def,
		Value:           value,
	}
}

func createTestTraitDefinition(
	comparableType rulesengine.ComparableType,
	entityType rulesengine.EntityType,
) *rulesengine.TraitDefinition {
	return &rulesengine.TraitDefinition{
		ID:             generateTestID("trt"),
		ComparableType: comparableType,
		EntityType:     entityType,
	}
}
