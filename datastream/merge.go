package datastream

import (
	"encoding/json"
	"fmt"
	"maps"

	"github.com/schematichq/rulesengine"
)

// PartialCompany merges a partial JSON update into an existing Company and
// returns the result. The original is not mutated. Each key in partialJSON
// names a Company JSON field; the SDK dispatches per-field merge semantics:
// maps shallow-merge, keyed collections upsert by their natural key, and
// scalar fields replace.
//
// Partials don't carry refreshed entitlements, so when credit_balances or
// metrics change without an accompanying entitlements field, we re-derive
// entitlement.credit_remaining and entitlement.usage locally to match what
// the server computes when assembling a full company message.
func PartialCompany(existing *rulesengine.Company, partialJSON json.RawMessage) (*rulesengine.Company, error) {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(partialJSON, &fields); err != nil {
		return nil, fmt.Errorf("unmarshal partial company fields: %w", err)
	}

	merged := DeepCopyCompany(existing)

	var (
		creditBalanceUpdates  map[string]float64
		metricsUpdated        bool
		entitlementsInPartial bool
	)

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
			var ids []string
			if err := json.Unmarshal(raw, &ids); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
			merged.BillingProductIDs = ids
		case "credit_balances":
			var cb map[string]float64
			if err := json.Unmarshal(raw, &cb); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
			if merged.CreditBalances == nil {
				merged.CreditBalances = make(map[string]float64, len(cb))
			}
			maps.Copy(merged.CreditBalances, cb)
			creditBalanceUpdates = cb
		case "entitlements":
			var ents []*rulesengine.FeatureEntitlement
			if err := json.Unmarshal(raw, &ents); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
			merged.Entitlements = ents
			entitlementsInPartial = true
		case "keys":
			var keys map[string]string
			if err := json.Unmarshal(raw, &keys); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
			if merged.Keys == nil {
				merged.Keys = make(map[string]string, len(keys))
			}
			maps.Copy(merged.Keys, keys)
		case "metrics":
			var incoming rulesengine.CompanyMetricCollection
			if err := json.Unmarshal(raw, &incoming); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
			merged.Metrics = upsertMetrics(merged.Metrics, incoming)
			metricsUpdated = true
		case "plan_ids":
			var ids []string
			if err := json.Unmarshal(raw, &ids); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
			merged.PlanIDs = ids
		case "plan_version_ids":
			var ids []string
			if err := json.Unmarshal(raw, &ids); err != nil {
				return nil, fmt.Errorf("unmarshal field %q: %w", key, err)
			}
			merged.PlanVersionIDs = ids
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

	if !entitlementsInPartial && (creditBalanceUpdates != nil || metricsUpdated) && len(merged.Entitlements) > 0 {
		merged.Entitlements = syncEntitlementDerivedFields(merged.Entitlements, creditBalanceUpdates, metricsUpdated, merged.Metrics)
	}

	return merged, nil
}

// syncEntitlementDerivedFields re-derives credit_remaining and usage on
// entitlements when a partial updates credit_balances or metrics without
// sending refreshed entitlements. Returns a new slice of new pointers — the
// input entitlements are not mutated.
//
// credit_remaining is matched per-entitlement against the incoming balance
// map (entitlements pointing at credits not in the partial are untouched).
// usage is matched against the merged metrics list using the full triple
// (event_name, metric_period, month_reset), defaulting period to all_time
// and month_reset to first_of_month when the entitlement leaves them unset.
//
// credit_total and credit_used are deliberately left alone — they aggregate
// across a grant ledger the SDK doesn't see.
func syncEntitlementDerivedFields(
	entitlements []*rulesengine.FeatureEntitlement,
	creditBalanceUpdates map[string]float64,
	metricsUpdated bool,
	mergedMetrics rulesengine.CompanyMetricCollection,
) []*rulesengine.FeatureEntitlement {
	var metricLookup map[metricKey]int64
	if metricsUpdated {
		metricLookup = make(map[metricKey]int64, len(mergedMetrics))
		for _, m := range mergedMetrics {
			if m == nil {
				continue
			}
			metricLookup[metricKey{m.EventSubtype, m.Period, m.MonthReset}] = m.Value
		}
	}

	result := make([]*rulesengine.FeatureEntitlement, len(entitlements))
	for i, ent := range entitlements {
		if ent == nil {
			continue
		}
		updated := *ent

		if creditBalanceUpdates != nil && ent.CreditID != nil {
			if balance, ok := creditBalanceUpdates[*ent.CreditID]; ok {
				v := balance
				updated.CreditRemaining = &v
			}
		}

		if metricLookup != nil && ent.EventName != nil {
			period := rulesengine.MetricPeriodAllTime
			if ent.MetricPeriod != nil {
				period = *ent.MetricPeriod
			}
			monthReset := rulesengine.MetricPeriodMonthResetFirst
			if ent.MonthReset != nil {
				monthReset = *ent.MonthReset
			}
			if value, ok := metricLookup[metricKey{*ent.EventName, period, monthReset}]; ok {
				v := value
				updated.Usage = &v
			}
		}

		result[i] = &updated
	}
	return result
}

// PartialUser merges a partial JSON update into an existing User and returns
// the result. The original is not mutated.
func PartialUser(existing *rulesengine.User, partialJSON json.RawMessage) (*rulesengine.User, error) {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(partialJSON, &fields); err != nil {
		return nil, fmt.Errorf("unmarshal partial user fields: %w", err)
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
		Entitlements:      append([]*rulesengine.FeatureEntitlement{}, c.Entitlements...),
		Rules:             append([]*rulesengine.Rule{}, c.Rules...),
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

