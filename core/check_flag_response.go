package core

import "github.com/schematichq/rulesengine"

// CheckFlagResponse contains the result of a flag check with entitlement information.
// Note: FeatureAllocation and FeatureUsage* fields are deprecated and not included.
// Use Entitlement field for allocation and usage information.
type CheckFlagResponse struct {
	CompanyID   *string                         `json:"company_id,omitempty"`
	Entitlement *rulesengine.FeatureEntitlement `json:"entitlement,omitempty"`
	FlagID      *string                         `json:"flag_id,omitempty"`
	FlagKey     string                          `json:"flag_key"`
	Reason      string                          `json:"reason"`
	RuleID      *string                         `json:"rule_id,omitempty"`
	RuleType    *rulesengine.RuleType           `json:"rule_type,omitempty"`
	UserID      *string                         `json:"user_id,omitempty"`
	Value       bool                            `json:"value"`
}
