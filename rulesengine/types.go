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
// Its json tags are camelCase because the engine emits results that way
// (serde's rename_all = "camelCase") and this type decodes that output
// directly -- see UnmarshalJSON. The engine's input types accept snake_case via
// aliases, so only the result direction is camelCase; the wire types in
// models.go stay snake_case to match the datastream.
type CheckFlagResult struct {
	CompanyID           *string       `json:"companyId,omitempty"`
	FeatureAllocation   *int64        `json:"featureAllocation,omitempty"`
	FeatureUsage        *int64        `json:"featureUsage,omitempty"`
	FeatureUsageEvent   *string       `json:"featureUsageEvent,omitempty"`
	FeatureUsagePeriod  *MetricPeriod `json:"featureUsagePeriod,omitempty"`
	FeatureUsageResetAt *time.Time    `json:"featureUsageResetAt,omitempty"`
	FlagID              *string       `json:"flagId,omitempty"`
	FlagKey             string        `json:"flagKey"`
	Reason              string        `json:"reason"`
	RuleID              *string       `json:"ruleId,omitempty"`
	RuleType            *RuleType     `json:"ruleType,omitempty"`
	UserID              *string       `json:"userId,omitempty"`
	Value               bool          `json:"value"`

	// Handled in UnmarshalJSON rather than decoded directly. Err is a Go error
	// but the engine sends a string; Entitlement is a FeatureEntitlement, whose
	// tags are snake_case for datastream input while the engine emits it
	// camelCase, so it needs the wasmFeatureEntitlement mirror.
	Err         error               `json:"-"`
	Entitlement *FeatureEntitlement `json:"-"`
}

// UnmarshalJSON decodes the engine's camelCase result. Every field except Err
// and Entitlement decodes straight into CheckFlagResult via the alias (which
// drops this method to avoid recursion); those two are decoded from their raw
// wire forms and converted.
func (r *CheckFlagResult) UnmarshalJSON(data []byte) error {
	type alias CheckFlagResult
	aux := struct {
		*alias
		Err         *string                 `json:"err"`
		Entitlement *wasmFeatureEntitlement `json:"entitlement"`
	}{alias: (*alias)(r)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	r.Entitlement = aux.Entitlement.toEntitlement()
	if aux.Err != nil && *aux.Err != "" {
		r.Err = newRulesEngineError(*aux.Err, 0)
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
