package datastream

import (
	"encoding/json"
	"fmt"
	"maps"

	"github.com/schematichq/rulesengine"
	schematicdatastreamws "github.com/schematichq/schematic-datastream-ws"
)

// ApplyPartialCompany merges a partial update into an existing Company and
// returns the result. The original is not mutated.
func ApplyPartialCompany(existing *rulesengine.Company, partialType schematicdatastreamws.PartialType, data json.RawMessage) (*rulesengine.Company, error) {
	merged := DeepCopyCompany(existing)

	switch partialType {
	case schematicdatastreamws.PartialTypeCreditBalances:
		var cb map[string]float64
		if err := json.Unmarshal(data, &cb); err != nil {
			return nil, fmt.Errorf("unmarshal credit balances: %w", err)
		}
		if merged.CreditBalances == nil {
			merged.CreditBalances = make(map[string]float64, len(cb))
		}
		maps.Copy(merged.CreditBalances, cb)

	case schematicdatastreamws.PartialTypeCompanyMetric:
		var metric *rulesengine.CompanyMetric
		if err := json.Unmarshal(data, &metric); err != nil {
			return nil, fmt.Errorf("unmarshal company metric: %w", err)
		}
		if metric == nil {
			return merged, nil
		}
		merged.Metrics = upsertMetrics(merged.Metrics, rulesengine.CompanyMetricCollection{metric})

	default:
		return nil, fmt.Errorf("unknown partial_type: %s", partialType)
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

	merged := DeepCopyUser(existing)

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
			if merged.Keys == nil {
				merged.Keys = make(map[string]string)
			}
			maps.Copy(merged.Keys, keys)
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

// DeepCopyCompany creates a complete deep copy of a Company struct and all its nested fields.
func DeepCopyCompany(c *rulesengine.Company) *rulesengine.Company {
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

// DeepCopyUser creates a complete deep copy of a User struct and all its nested fields.
func DeepCopyUser(u *rulesengine.User) *rulesengine.User {
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

type metricKey struct {
	EventSubtype string
	Period       rulesengine.MetricPeriod
	MonthReset   rulesengine.MetricPeriodMonthReset
}

// upsertMetrics merges incoming metrics into existing ones. Metrics are
// matched by (EventSubtype, Period, MonthReset); matches are replaced,
// new metrics are appended.
func upsertMetrics(existing, incoming rulesengine.CompanyMetricCollection) rulesengine.CompanyMetricCollection {
	index := make(map[metricKey]int, len(existing))
	result := make(rulesengine.CompanyMetricCollection, len(existing))
	copy(result, existing)

	for i, m := range result {
		if m != nil {
			index[metricKey{m.EventSubtype, m.Period, m.MonthReset}] = i
		}
	}

	for _, m := range incoming {
		if m == nil {
			continue
		}
		k := metricKey{m.EventSubtype, m.Period, m.MonthReset}
		if idx, ok := index[k]; ok {
			result[idx] = m
		} else {
			result = append(result, m)
		}
	}

	return result
}
