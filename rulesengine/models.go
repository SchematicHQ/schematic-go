package rulesengine

import (
	"encoding/json"
	"time"
)

// The types below are the rules engine's wire format: what the datastream sends,
// what gets cached, and what is handed to the WebAssembly engine for evaluation.
//
// They live here, rather than being taken from the generated Rulesengine* types
// in the root package, because the root package imports core (Fern generates
// errors.go that way) and core exposes CheckFlagResponse, whose fields are typed
// from this package. Importing the root package here would close that loop into
// an import cycle. Keeping this package a leaf is what lets core keep its
// current public API.
//
// Only the wire shape lives here. Evaluation itself is the WebAssembly engine's
// job, so the rule-matching helpers the Go engine carried alongside these types
// are deliberately absent.

// ComparableType is the type a trait's value is compared as.
type ComparableType string

const (
	ComparableTypeBool   ComparableType = "bool"
	ComparableTypeDate   ComparableType = "date"
	ComparableTypeInt    ComparableType = "int"
	ComparableTypeString ComparableType = "string"
)

// ComparableOperator is the comparison a condition applies.
type ComparableOperator string

const (
	ComparableOperatorEquals    ComparableOperator = "eq"
	ComparableOperatorNotEquals ComparableOperator = "ne"
	ComparableOperatorGt        ComparableOperator = "gt"
	ComparableOperatorLt        ComparableOperator = "lt"
	ComparableOperatorGte       ComparableOperator = "gte"
	ComparableOperatorLte       ComparableOperator = "lte"
	ComparableOperatorIsEmpty   ComparableOperator = "is_empty"
	ComparableOperatorNotEmpty  ComparableOperator = "not_empty"
)

// EntityType is the kind of entity a trait definition applies to.
type EntityType string

const (
	EntityTypeCompany EntityType = "company"
	EntityTypeUser    EntityType = "user"
)

// ConditionType is what a condition matches on.
type ConditionType string

const (
	ConditionTypeBasePlan       ConditionType = "base_plan"
	ConditionTypeBillingProduct ConditionType = "billing_product"
	ConditionTypeCompany        ConditionType = "company"
	ConditionTypeCredit         ConditionType = "credit"
	ConditionTypeMetric         ConditionType = "metric"
	ConditionTypePlan           ConditionType = "plan"
	ConditionTypePlanVersion    ConditionType = "plan_version"
	ConditionTypeTrait          ConditionType = "trait"
	ConditionTypeUser           ConditionType = "user"
)

// RuleType is the kind of rule that produced a flag value.
type RuleType string

const (
	RuleTypeCompanyOverride              RuleType = "company_override"
	RuleTypeCompanyOverrideUsageExceeded RuleType = "company_override_usage_exceeded"
	RuleTypeDefault                      RuleType = "default"
	RuleTypeGlobalOverride               RuleType = "global_override"
	RuleTypePlanEntitlement              RuleType = "plan_entitlement"
	RuleTypePlanEntitlementUsageExceeded RuleType = "plan_entitlement_usage_exceeded"
	RuleTypeStandard                     RuleType = "standard"
)

// EntitlementValueType is the type of a feature entitlement's value.
type EntitlementValueType string

const (
	EntitlementValueTypeBoolean   EntitlementValueType = "boolean"
	EntitlementValueTypeCredit    EntitlementValueType = "credit"
	EntitlementValueTypeNumeric   EntitlementValueType = "numeric"
	EntitlementValueTypeTrait     EntitlementValueType = "trait"
	EntitlementValueTypeUnknown   EntitlementValueType = "unknown"
	EntitlementValueTypeUnlimited EntitlementValueType = "unlimited"
)

// MetricPeriod is the window over which usage is counted.
type MetricPeriod string

const (
	MetricPeriodAllTime      MetricPeriod = "all_time"
	MetricPeriodCurrentDay   MetricPeriod = "current_day"
	MetricPeriodCurrentMonth MetricPeriod = "current_month"
	MetricPeriodCurrentWeek  MetricPeriod = "current_week"
)

// MetricPeriodMonthReset is how a monthly period resets.
type MetricPeriodMonthReset string

const (
	MetricPeriodMonthResetFirst   MetricPeriodMonthReset = "first_of_month"
	MetricPeriodMonthResetBilling MetricPeriodMonthReset = "billing_cycle"
)

type TraitDefinition struct {
	ID             string         `json:"id"`
	ComparableType ComparableType `json:"comparable_type"`
	EntityType     EntityType     `json:"entity_type"`
}

type Flag struct {
	ID            string           `json:"id"`
	AccountID     string           `json:"account_id"`
	EnvironmentID string           `json:"environment_id"`
	Key           string           `json:"key"`
	Rules         JSONSlice[*Rule] `json:"rules"`
	DefaultValue  bool             `json:"default_value"`
}

type Rule struct {
	ID              string                     `json:"id"`
	FlagID          *string                    `json:"flag_id"`
	AccountID       string                     `json:"account_id"`
	EnvironmentID   string                     `json:"environment_id"`
	RuleType        RuleType                   `json:"rule_type"`
	Name            string                     `json:"name"`
	Priority        int64                      `json:"priority"`
	Conditions      JSONSlice[*Condition]      `json:"conditions"`
	ConditionGroups JSONSlice[*ConditionGroup] `json:"condition_groups"`
	Value           bool                       `json:"value"`
}

type Condition struct {
	ID            string             `json:"id"`
	AccountID     string             `json:"account_id"`
	EnvironmentID string             `json:"environment_id"`
	ConditionType ConditionType      `json:"condition_type"`
	Operator      ComparableOperator `json:"operator"`

	// Relevant when ConditionType is one of Company, User, Plan, Plan Version,
	// Base Plan, Billing Product, or Billing Credit
	ResourceIDs JSONSlice[string] `json:"resource_ids"`

	// Relevant when ConditionType = Metric
	EventSubtype           *string                 `json:"event_subtype"`
	MetricValue            *int64                  `json:"metric_value"`
	MetricPeriod           *MetricPeriod           `json:"metric_period"`
	MetricPeriodMonthReset *MetricPeriodMonthReset `json:"metric_period_month_reset"`

	// Relevant when ConditionType = Credit
	CreditID        *string  `json:"credit_id"`
	ConsumptionRate *float64 `json:"consumption_rate"`

	// Relevant when ConditionType = Trait
	TraitDefinition *TraitDefinition `json:"trait_definition"`
	TraitValue      string           `json:"trait_value"`

	// Relevant when ConditionType is either Metric or Trait
	ComparisonTraitDefinition *TraitDefinition `json:"comparison_trait_definition"`
}

type ConditionGroup struct {
	Conditions JSONSlice[*Condition] `json:"conditions"`
}

type CompanyMetric struct {
	AccountID     string                 `json:"account_id"`
	EnvironmentID string                 `json:"environment_id"`
	CompanyID     string                 `json:"company_id"`
	EventSubtype  string                 `json:"event_subtype"`
	Period        MetricPeriod           `json:"period"`
	MonthReset    MetricPeriodMonthReset `json:"month_reset"`
	Value         int64                  `json:"value"`
	CreatedAt     time.Time              `json:"created_at"`
	ValidUntil    *time.Time             `json:"valid_until"`
}

type CompanyMetricCollection []*CompanyMetric

// MarshalJSON ensures a nil collection serializes as `[]` rather than `null`,
// matching the JSONSlice contract the rest of these wire types follow.
func (c CompanyMetricCollection) MarshalJSON() ([]byte, error) {
	if c == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]*CompanyMetric(c))
}

type Trait struct {
	TraitDefinition *TraitDefinition `json:"trait_definition"`
	Value           string           `json:"value"`
}

type Subscription struct {
	ID          string    `json:"id"`
	PeriodStart time.Time `json:"period_start"`
	PeriodEnd   time.Time `json:"period_end"`
}

type FeatureEntitlement struct {
	Allocation      *int64                  `json:"allocation"`
	ConsumptionRate *float64                `json:"consumption_rate,omitempty"`
	CreditID        *string                 `json:"credit_id"`
	CreditRemaining *float64                `json:"credit_remaining"`
	CreditReserved  *float64                `json:"credit_reserved,omitempty"`
	CreditSettled   *float64                `json:"credit_settled,omitempty"`
	CreditTotal     *float64                `json:"credit_total"`
	CreditUsed      *float64                `json:"credit_used"`
	EventName       *string                 `json:"event_name"`
	EventSubtype    *string                 `json:"event_subtype,omitempty"`
	FeatureID       string                  `json:"feature_id"`
	FeatureKey      string                  `json:"feature_key"`
	MetricPeriod    *MetricPeriod           `json:"metric_period"`
	MetricResetAt   *time.Time              `json:"metric_reset_at"`
	MonthReset      *MetricPeriodMonthReset `json:"month_reset"`
	SoftLimit       *int64                  `json:"soft_limit"`
	Usage           *int64                  `json:"usage"`
	ValueType       EntitlementValueType    `json:"value_type"`
	WarningTiers    JSONSlice[*WarningTier] `json:"warning_tiers,omitempty"`
}

// WarningTier is a customer-defined usage warning threshold on a feature
// entitlement. Key is a customer label identifying the tier; Value is the
// threshold, expressed in the entitlement's usage units.
type WarningTier struct {
	Key   string `json:"key"`
	Value int64  `json:"value"`
}

type Company struct {
	ID            string `json:"id"`
	AccountID     string `json:"account_id"`
	EnvironmentID string `json:"environment_id"`

	BasePlanID        *string                        `json:"base_plan_id"`
	BillingProductIDs JSONSlice[string]              `json:"billing_product_ids"`
	CreditBalances    map[string]float64             `json:"credit_balances,omitempty"`
	Entitlements      JSONSlice[*FeatureEntitlement] `json:"entitlements,omitempty"`
	Keys              map[string]string              `json:"keys,omitempty"`
	Metrics           CompanyMetricCollection        `json:"metrics"`
	PlanIDs           JSONSlice[string]              `json:"plan_ids"`
	PlanVersionIDs    JSONSlice[string]              `json:"plan_version_ids"`
	Rules             JSONSlice[*Rule]               `json:"rules"`
	Subscription      *Subscription                  `json:"subscription"`
	Traits            JSONSlice[*Trait]              `json:"traits"`
}

type User struct {
	ID            string `json:"id"`
	AccountID     string `json:"account_id"`
	EnvironmentID string `json:"environment_id"`

	Keys   map[string]string `json:"keys,omitempty"`
	Traits JSONSlice[*Trait] `json:"traits"`
	Rules  JSONSlice[*Rule]  `json:"rules"`
}
