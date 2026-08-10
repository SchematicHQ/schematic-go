package rulesengine

import (
	"encoding/json"
	"fmt"
	"time"
)

// Reasons reported on a CheckFlagResult. These match the engine's constants
// verbatim; callers compare against them.
const (
	ReasonNoCompanyOrUser     = "No company or user context; default value for flag"
	ReasonCompanyNotFound     = "Company not found"
	ReasonCompanyNotSpecified = "Must specify a company"
	ReasonFlagNotFound        = "Flag not found"
	ReasonNoRulesMatched      = "No rules matched; default value for flag"
	ReasonServerError         = "Server error; Schematic has been notified"
	ReasonUserNotFound        = "User not found"
)

// CheckFlagResult is the outcome of evaluating a flag.
//
// Its json tags are snake_case, matching the type this SDK previously exposed
// (the external rulesengine.CheckFlagResult), so callers that marshal a result
// returned by DataStreamClient.CheckFlag see an unchanged shape. The engine
// itself emits camelCase (serde's rename_all = "camelCase"); UnmarshalJSON
// bridges that, so the wire casing never leaks into this public type.
type CheckFlagResult struct {
	CompanyID           *string             `json:"company_id,omitempty"`
	Err                 error               `json:"err,omitempty"`
	Entitlement         *FeatureEntitlement `json:"entitlement,omitempty"`
	FeatureAllocation   *int64              `json:"feature_allocation,omitempty"`
	FeatureUsage        *int64              `json:"feature_usage,omitempty"`
	FeatureUsageEvent   *string             `json:"feature_usage_event,omitempty"`
	FeatureUsagePeriod  *MetricPeriod       `json:"feature_usage_period,omitempty"`
	FeatureUsageResetAt *time.Time          `json:"feature_usage_reset_at,omitempty"`
	FlagID              *string             `json:"flag_id,omitempty"`
	FlagKey             string              `json:"flag_key"`
	Reason              string              `json:"reason"`
	RuleID              *string             `json:"rule_id,omitempty"`
	RuleType            *RuleType           `json:"rule_type,omitempty"`
	UserID              *string             `json:"user_id,omitempty"`
	Value               bool                `json:"value"`
}

// UnmarshalJSON decodes the engine's camelCase result into this snake_case type.
// It decodes into a local wire struct with the engine's field names, then maps
// across; Err (the engine sends a string, this carries a Go error) and
// Entitlement (bidirectional casing, see wasmFeatureEntitlement) are converted.
func (r *CheckFlagResult) UnmarshalJSON(data []byte) error {
	var w struct {
		Value               bool                    `json:"value"`
		Reason              string                  `json:"reason"`
		Err                 *string                 `json:"err"`
		RuleID              *string                 `json:"ruleId"`
		RuleType            *RuleType               `json:"ruleType"`
		CompanyID           *string                 `json:"companyId"`
		UserID              *string                 `json:"userId"`
		Entitlement         *wasmFeatureEntitlement `json:"entitlement"`
		FlagID              *string                 `json:"flagId"`
		FlagKey             string                  `json:"flagKey"`
		FeatureAllocation   *int64                  `json:"featureAllocation"`
		FeatureUsage        *int64                  `json:"featureUsage"`
		FeatureUsageEvent   *string                 `json:"featureUsageEvent"`
		FeatureUsagePeriod  *MetricPeriod           `json:"featureUsagePeriod"`
		FeatureUsageResetAt *time.Time              `json:"featureUsageResetAt"`
	}
	if err := json.Unmarshal(data, &w); err != nil {
		return err
	}

	*r = CheckFlagResult{
		Value:               w.Value,
		Reason:              w.Reason,
		CompanyID:           w.CompanyID,
		UserID:              w.UserID,
		Entitlement:         w.Entitlement.toEntitlement(),
		FeatureAllocation:   w.FeatureAllocation,
		FeatureUsage:        w.FeatureUsage,
		FeatureUsageEvent:   w.FeatureUsageEvent,
		FeatureUsagePeriod:  w.FeatureUsagePeriod,
		FeatureUsageResetAt: w.FeatureUsageResetAt,
		FlagID:              w.FlagID,
		FlagKey:             w.FlagKey,
		RuleID:              w.RuleID,
		RuleType:            w.RuleType,
	}
	if w.Err != nil && *w.Err != "" {
		r.Err = newRulesEngineError(*w.Err, 0)
	}
	return nil
}

// checkFlagEnvelope is the input to the engine's checkFlagCombined export.
type checkFlagEnvelope struct {
	Flag    *Flag             `json:"flag"`
	Company *Company          `json:"company,omitempty"`
	User    *User             `json:"user,omitempty"`
	Options *checkFlagOptions `json:"options,omitempty"`
}

// wasmFeatureEntitlement mirrors the engine's FeatureEntitlement wire shape.
// It exists because FeatureEntitlement is bidirectional: snake_case when
// received from the datastream (as Company.Entitlements) and camelCase when
// returned by the engine, and a Go struct field can carry only one json tag.
// This is the only such mirror the result path needs.
type wasmFeatureEntitlement struct {
	FeatureID       string                  `json:"featureId"`
	FeatureKey      string                  `json:"featureKey"`
	ValueType       EntitlementValueType    `json:"valueType"`
	Allocation      *int64                  `json:"allocation"`
	SoftLimit       *int64                  `json:"softLimit"`
	Usage           *int64                  `json:"usage"`
	EventName       *string                 `json:"eventName"`
	EventSubtype    *string                 `json:"eventSubtype"`
	MetricPeriod    *MetricPeriod           `json:"metricPeriod"`
	MonthReset      *MetricPeriodMonthReset `json:"monthReset"`
	MetricResetAt   *time.Time              `json:"metricResetAt"`
	CreditID        *string                 `json:"creditId"`
	CreditTotal     *float64                `json:"creditTotal"`
	CreditUsed      *float64                `json:"creditUsed"`
	CreditRemaining *float64                `json:"creditRemaining"`
	CreditReserved  *float64                `json:"creditReserved"`
	CreditSettled   *float64                `json:"creditSettled"`
	ConsumptionRate *float64                `json:"consumptionRate"`

	// Emitted only by engine builds that carry warning tiers. The Rust type has
	// the field, but it postdates the v0.6.0 release this SDK pins, so against
	// that binary it is simply absent and decodes to nil. Mapping it now means
	// the data flows through the moment WASM_VERSION moves to a build that
	// includes it, rather than being silently dropped here.
	WarningTiers JSONSlice[*WarningTier] `json:"warningTiers"`
}

func (w *wasmFeatureEntitlement) toEntitlement() *FeatureEntitlement {
	if w == nil {
		return nil
	}
	return &FeatureEntitlement{
		FeatureID:       w.FeatureID,
		FeatureKey:      w.FeatureKey,
		ValueType:       w.ValueType,
		Allocation:      w.Allocation,
		SoftLimit:       w.SoftLimit,
		Usage:           w.Usage,
		EventName:       w.EventName,
		EventSubtype:    w.EventSubtype,
		MetricPeriod:    w.MetricPeriod,
		MonthReset:      w.MonthReset,
		MetricResetAt:   w.MetricResetAt,
		CreditID:        w.CreditID,
		CreditTotal:     w.CreditTotal,
		CreditUsed:      w.CreditUsed,
		CreditRemaining: w.CreditRemaining,
		CreditReserved:  w.CreditReserved,
		CreditSettled:   w.CreditSettled,
		ConsumptionRate: w.ConsumptionRate,
		WarningTiers:    w.WarningTiers,
	}
}

// marshalEnvelope serializes the envelope for the engine.
//
// No null-stripping pass is needed, unlike in the Node and Python SDKs. The
// engine's Rust types fill absent fields via serde's #[serde(default)], which
// does not cover a field explicitly set to null: `"rules": null` fails to
// deserialize into Vec<Rule>. The wire types in this package sidestep that by
// construction -- every collection is a JSONSlice (or CompanyMetricCollection),
// which marshals nil as `[]`, and the two map fields are omitempty. The nulls
// that do survive sit on Option-typed scalars like base_plan_id and
// subscription, which serde accepts. rulesengine/models_test.go pins that.
func marshalEnvelope(env *checkFlagEnvelope) ([]byte, error) {
	out, err := json.Marshal(env)
	if err != nil {
		return nil, fmt.Errorf("marshal envelope: %w", err)
	}
	return out, nil
}
