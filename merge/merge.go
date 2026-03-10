package merge

import (
	"encoding/json"
	"fmt"

	"github.com/schematichq/rulesengine"
)

// PartialCompany merges a partial JSON update into an existing Company.
// It copies the existing company by value, then applies only the fields
// present in partialJSON. The "id" field must be present.
func PartialCompany(existing *rulesengine.Company, partialJSON json.RawMessage) (*rulesengine.Company, error) {
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

// PartialUser merges a partial JSON update into an existing User.
// It copies the existing user by value, then applies only the fields
// present in partialJSON. The "id" field must be present.
func PartialUser(existing *rulesengine.User, partialJSON json.RawMessage) (*rulesengine.User, error) {
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

// ExtractIDFromJSON extracts the "id" field from a raw JSON message.
func ExtractIDFromJSON(data json.RawMessage) (string, error) {
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

func deepCopyCompany(c *rulesengine.Company) *rulesengine.Company {
	cp := &rulesengine.Company{
		ID:            c.ID,
		AccountID:     c.AccountID,
		EnvironmentID: c.EnvironmentID,
	}

	if c.BasePlanID != nil {
		v := *c.BasePlanID
		cp.BasePlanID = &v
	}

	if c.BillingProductIDs != nil {
		cp.BillingProductIDs = make([]string, len(c.BillingProductIDs))
		copy(cp.BillingProductIDs, c.BillingProductIDs)
	}

	if c.CreditBalances != nil {
		cp.CreditBalances = make(map[string]float64, len(c.CreditBalances))
		for k, v := range c.CreditBalances {
			cp.CreditBalances[k] = v
		}
	}

	if c.Entitlements != nil {
		cp.Entitlements = make([]*rulesengine.FeatureEntitlement, len(c.Entitlements))
		copy(cp.Entitlements, c.Entitlements)
	}

	if c.Keys != nil {
		cp.Keys = make(map[string]string, len(c.Keys))
		for k, v := range c.Keys {
			cp.Keys[k] = v
		}
	}

	if c.Metrics != nil {
		cp.Metrics = make(rulesengine.CompanyMetricCollection, len(c.Metrics))
		copy(cp.Metrics, c.Metrics)
	}

	if c.PlanIDs != nil {
		cp.PlanIDs = make([]string, len(c.PlanIDs))
		copy(cp.PlanIDs, c.PlanIDs)
	}

	if c.PlanVersionIDs != nil {
		cp.PlanVersionIDs = make([]string, len(c.PlanVersionIDs))
		copy(cp.PlanVersionIDs, c.PlanVersionIDs)
	}

	if c.Rules != nil {
		cp.Rules = make([]*rulesengine.Rule, len(c.Rules))
		copy(cp.Rules, c.Rules)
	}

	if c.Subscription != nil {
		sub := *c.Subscription
		cp.Subscription = &sub
	}

	if c.Traits != nil {
		cp.Traits = make([]*rulesengine.Trait, len(c.Traits))
		copy(cp.Traits, c.Traits)
	}

	return cp
}

func deepCopyUser(u *rulesengine.User) *rulesengine.User {
	cp := &rulesengine.User{
		ID:            u.ID,
		AccountID:     u.AccountID,
		EnvironmentID: u.EnvironmentID,
	}

	if u.Keys != nil {
		cp.Keys = make(map[string]string, len(u.Keys))
		for k, v := range u.Keys {
			cp.Keys[k] = v
		}
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
