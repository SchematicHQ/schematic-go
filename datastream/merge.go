package datastream

import (
	"encoding/json"
	"fmt"
	"maps"

	"github.com/schematichq/rulesengine"
)

// partialCompany merges a partial JSON update into an existing Company.
// It copies the existing company by value, then applies only the fields
// present in partialJSON. The "id" field must be present.
func partialCompany(existing *rulesengine.Company, partialJSON json.RawMessage) (*rulesengine.Company, error) {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(partialJSON, &fields); err != nil {
		return nil, fmt.Errorf("unmarshal partial company fields: %w", err)
	}

	if _, ok := fields["id"]; !ok {
		return nil, fmt.Errorf("partial company message missing required field: id")
	}

	// Deep copy existing
	merged := deepCopyCompany(existing)

	for key, raw := range fields {
		switch key {
		case "id":
			if err := json.Unmarshal(raw, &merged.ID); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
		case "account_id":
			if err := json.Unmarshal(raw, &merged.AccountID); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
		case "environment_id":
			if err := json.Unmarshal(raw, &merged.EnvironmentID); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
		case "base_plan_id":
			if err := json.Unmarshal(raw, &merged.BasePlanID); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
		case "billing_product_ids":
			var bpIDs []string
			if err := json.Unmarshal(raw, &bpIDs); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
			merged.BillingProductIDs = bpIDs
		case "credit_balances":
			var cb map[string]float64
			if err := json.Unmarshal(raw, &cb); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
			merged.CreditBalances = cb
		case "entitlements":
			var ents []*rulesengine.FeatureEntitlement
			if err := json.Unmarshal(raw, &ents); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
			merged.Entitlements = ents
		case "keys":
			var keys map[string]string
			if err := json.Unmarshal(raw, &keys); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
			merged.Keys = keys
		case "metrics":
			var metrics rulesengine.CompanyMetricCollection
			if err := json.Unmarshal(raw, &metrics); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
			merged.Metrics = metrics
		case "plan_ids":
			var planIDs []string
			if err := json.Unmarshal(raw, &planIDs); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
			merged.PlanIDs = planIDs
		case "plan_version_ids":
			var pvIDs []string
			if err := json.Unmarshal(raw, &pvIDs); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
			merged.PlanVersionIDs = pvIDs
		case "rules":
			var rules []*rulesengine.Rule
			if err := json.Unmarshal(raw, &rules); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
			merged.Rules = rules
		case "subscription":
			if err := json.Unmarshal(raw, &merged.Subscription); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
		case "traits":
			var traits []*rulesengine.Trait
			if err := json.Unmarshal(raw, &traits); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
			merged.Traits = traits
		}
	}

	return merged, nil
}

// partialUser merges a partial JSON update into an existing User.
// It copies the existing user by value, then applies only the fields
// present in partialJSON. The "id" field must be present.
func partialUser(existing *rulesengine.User, partialJSON json.RawMessage) (*rulesengine.User, error) {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(partialJSON, &fields); err != nil {
		return nil, fmt.Errorf("unmarshal partial user fields: %w", err)
	}

	if _, ok := fields["id"]; !ok {
		return nil, fmt.Errorf("partial user message missing required field: id")
	}

	// Deep copy existing
	merged := deepCopyUser(existing)

	for key, raw := range fields {
		switch key {
		case "id":
			if err := json.Unmarshal(raw, &merged.ID); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
		case "account_id":
			if err := json.Unmarshal(raw, &merged.AccountID); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
		case "environment_id":
			if err := json.Unmarshal(raw, &merged.EnvironmentID); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
		case "keys":
			var keys map[string]string
			if err := json.Unmarshal(raw, &keys); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
			merged.Keys = keys
		case "traits":
			var traits []*rulesengine.Trait
			if err := json.Unmarshal(raw, &traits); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
			merged.Traits = traits
		case "rules":
			var rules []*rulesengine.Rule
			if err := json.Unmarshal(raw, &rules); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
			merged.Rules = rules
		}
	}

	return merged, nil
}

// extractIDFromJSON extracts the "id" field from a raw JSON message.
func extractIDFromJSON(data json.RawMessage) (string, error) {
	var partial struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(data, &partial); err != nil {
		return "", fmt.Errorf("unmarshal id from partial message: %w", err)
	}
	if partial.ID == "" {
		return "", fmt.Errorf("partial message missing required field: id")
	}
	return partial.ID, nil
}

// deepCopyCompany creates a complete deep copy of a Company struct and all its nested fields.
// This ensures that modifying the returned company won't affect the original object.
func deepCopyCompany(c *rulesengine.Company) *rulesengine.Company {
	if c == nil {
		return nil
	}

	cp := &rulesengine.Company{
		ID:                c.ID,
		AccountID:         c.AccountID,
		EnvironmentID:     c.EnvironmentID,
		BillingProductIDs: append([]string{}, c.BillingProductIDs...),
		PlanIDs:           append([]string{}, c.PlanIDs...),
		PlanVersionIDs:    append([]string{}, c.PlanVersionIDs...),
		Entitlements:      c.Entitlements,
		Rules:             c.Rules,
		Keys:              make(map[string]string),
		Metrics:           make([]*rulesengine.CompanyMetric, 0, len(c.Metrics)),
		Traits:            make([]*rulesengine.Trait, 0, len(c.Traits)),
	}

	if c.BasePlanID != nil {
		v := *c.BasePlanID
		cp.BasePlanID = &v
	}

	if c.CreditBalances != nil {
		cp.CreditBalances = make(map[string]float64, len(c.CreditBalances))
		maps.Copy(cp.CreditBalances, c.CreditBalances)
	}

	maps.Copy(cp.Keys, c.Keys)

	if c.Subscription != nil {
		sub := *c.Subscription
		cp.Subscription = &sub
	}

	for _, metric := range c.Metrics {
		if metric == nil {
			cp.Metrics = append(cp.Metrics, nil)
			continue
		}
		metricCopy := &rulesengine.CompanyMetric{
			AccountID:     metric.AccountID,
			EnvironmentID: metric.EnvironmentID,
			CompanyID:     metric.CompanyID,
			EventSubtype:  metric.EventSubtype,
			Period:        metric.Period,
			MonthReset:    metric.MonthReset,
			Value:         metric.Value,
			CreatedAt:     metric.CreatedAt,
		}
		if metric.ValidUntil != nil {
			validUntil := *metric.ValidUntil
			metricCopy.ValidUntil = &validUntil
		}
		cp.Metrics = append(cp.Metrics, metricCopy)
	}

	for _, trait := range c.Traits {
		if trait != nil {
			traitCopy := *trait
			cp.Traits = append(cp.Traits, &traitCopy)
		} else {
			cp.Traits = append(cp.Traits, nil)
		}
	}

	return cp
}

// deepCopyUser creates a complete deep copy of a User struct and all its nested fields.
func deepCopyUser(u *rulesengine.User) *rulesengine.User {
	cp := &rulesengine.User{
		ID:            u.ID,
		AccountID:     u.AccountID,
		EnvironmentID: u.EnvironmentID,
	}

	if u.Keys != nil {
		cp.Keys = make(map[string]string, len(u.Keys))
		maps.Copy(cp.Keys, u.Keys)
	}

	if u.Traits != nil {
		cp.Traits = make([]*rulesengine.Trait, len(u.Traits))
		copy(cp.Traits, u.Traits)
	}

	if u.Rules != nil {
		cp.Rules = make([]*rulesengine.Rule, len(u.Rules))
		copy(cp.Rules, u.Rules)
	}

	return cp
}
