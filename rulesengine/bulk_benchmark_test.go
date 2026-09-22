package rulesengine_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/schematichq/schematic-go/rulesengine"
	"github.com/stretchr/testify/require"
)

// Engine.CheckFlags against calling Engine.CheckFlag once per flag.
//
// Evaluation is identical either way, so what is being measured is the wasm
// boundary: marshaling the company, copying it into the module's linear memory
// and parsing it there, once per flag versus once for the whole set.
//
// The individual path trims the company per flag, keeping only the rules
// naming that flag and the single entitlement matching its feature key, since
// those are all the engine will consult for it. That is the fair baseline: a
// caller checking many flags for one company should already be trimming, so
// benchmarking bulk against an untrimmed loop would credit it with a saving
// the trim gets on its own. The untrimmed series is kept alongside to show how
// much that is.
//
// What the trim cannot reach is traits and metrics. Conditions look traits up
// by definition, so which ones matter is not knowable from the flag alone, and
// both have to be sent whole. They are what the individual path re-sends per
// flag, and what bulk sends once.
//
// Run with:
//
//	go test ./rulesengine -run XXX -bench 'BenchmarkCheckFlags' -benchmem

// benchFlags builds entitlement-shaped flags, one per feature, each gated on a
// metric the company reports.
func benchFlags(count int) []*rulesengine.Flag {
	flags := make([]*rulesengine.Flag, 0, count)

	for i := 0; i < count; i++ {
		condition := createTestCondition(rulesengine.ConditionTypeMetric)
		condition.EventSubtype = ptr(fmt.Sprintf("event_%d", i))
		condition.MetricValue = ptr(int64(1000))
		condition.Operator = rulesengine.ComparableOperatorLt

		rule := createTestRule()
		rule.RuleType = rulesengine.RuleTypePlanEntitlement
		rule.Conditions = rulesengine.NewJSONSlice([]*rulesengine.Condition{condition})

		flag := createTestFlag()
		flag.Key = fmt.Sprintf("feature_%d", i)
		flag.Rules = rulesengine.NewJSONSlice([]*rulesengine.Rule{rule})
		rule.FlagID = ptr(flag.ID)

		flags = append(flags, flag)
	}

	return flags
}

// benchCompany builds a company shaped like one the datastream hands the API:
// traits it was identified with, a metric and an entitlement per metered
// feature, and a company override rule naming each flag.
//
// The rules name real flags and the entitlements carry real feature keys, so
// the trim below has something to keep as well as something to drop. A fixture
// whose rules had no flag id would let the trim discard all of them and flatter
// the individual path.
func benchCompany(flags []*rulesengine.Flag, traitCount int) *rulesengine.Company {
	company := createTestCompany()

	traits := make([]*rulesengine.Trait, 0, traitCount)
	for i := 0; i < traitCount; i++ {
		traits = append(traits, createTestTrait(fmt.Sprintf("value_%d", i), nil))
	}
	company.Traits = rulesengine.NewJSONSlice(traits)

	metrics := make(rulesengine.CompanyMetricCollection, 0, len(flags))
	entitlements := make([]*rulesengine.FeatureEntitlement, 0, len(flags))
	rules := make([]*rulesengine.Rule, 0, len(flags))

	for i, flag := range flags {
		metrics = append(metrics, createTestMetric(
			company,
			fmt.Sprintf("event_%d", i),
			rulesengine.MetricPeriodAllTime,
			int64(i%900),
		))

		entitlements = append(entitlements, &rulesengine.FeatureEntitlement{
			FeatureID:  fmt.Sprintf("feat_%d", i),
			FeatureKey: flag.Key,
			ValueType:  rulesengine.EntitlementValueTypeNumeric,
			Allocation: ptr(int64(1000)),
			Usage:      ptr(int64(i % 900)),
		})

		rule := createTestRule()
		rule.RuleType = rulesengine.RuleTypeCompanyOverride
		rule.Priority = int64(i)
		rule.FlagID = ptr(flag.ID)
		rule.Conditions = rulesengine.NewJSONSlice([]*rulesengine.Condition{
			createTestCondition(rulesengine.ConditionTypeTrait),
		})
		rules = append(rules, rule)
	}

	company.Metrics = metrics
	company.Entitlements = rulesengine.NewJSONSlice(entitlements)
	company.Rules = rulesengine.NewJSONSlice(rules)

	return company
}

// trimCompanyForFlag keeps only what the engine will consult for this flag:
// the rules naming it, and the first entitlement matching its feature key,
// since the engine takes that one and stops. It copies rather than mutating,
// because one company is checked against many flags.
func trimCompanyForFlag(company *rulesengine.Company, flag *rulesengine.Flag) *rulesengine.Company {
	if company == nil {
		return nil
	}

	trimmed := *company
	trimmed.Rules = rulesForFlags(company.Rules, flag.ID)

	trimmed.Entitlements = nil
	for _, ent := range company.Entitlements {
		if ent != nil && ent.FeatureKey == flag.Key {
			trimmed.Entitlements = rulesengine.JSONSlice[*rulesengine.FeatureEntitlement]{ent}
			break
		}
	}

	return &trimmed
}

// trimCompanyForFlags is the trim a bulk check can do: the union across every
// flag in the set. Checking the whole environment makes it close to a no-op,
// which is the point -- bulk's saving is sending one company, not a smaller one.
func trimCompanyForFlags(company *rulesengine.Company, flags []*rulesengine.Flag) *rulesengine.Company {
	if company == nil {
		return nil
	}

	ids := make([]string, 0, len(flags))
	keys := make(map[string]bool, len(flags))
	for _, flag := range flags {
		ids = append(ids, flag.ID)
		keys[flag.Key] = true
	}

	trimmed := *company
	trimmed.Rules = rulesForFlags(company.Rules, ids...)

	kept := make([]*rulesengine.FeatureEntitlement, 0, len(flags))
	for _, ent := range company.Entitlements {
		if ent != nil && keys[ent.FeatureKey] {
			kept = append(kept, ent)
		}
	}
	trimmed.Entitlements = rulesengine.NewJSONSlice(kept)

	return &trimmed
}

// rulesForFlags keeps the rules naming any of the given flags. A rule without a
// flag id is dropped, matching the engine, which skips it.
func rulesForFlags(rules rulesengine.JSONSlice[*rulesengine.Rule], flagIDs ...string) rulesengine.JSONSlice[*rulesengine.Rule] {
	wanted := make(map[string]bool, len(flagIDs))
	for _, id := range flagIDs {
		wanted[id] = true
	}

	var kept rulesengine.JSONSlice[*rulesengine.Rule]
	for _, rule := range rules {
		if rule != nil && rule.FlagID != nil && wanted[*rule.FlagID] {
			kept = append(kept, rule)
		}
	}

	return kept
}

var benchFlagCounts = []int{1, 10, 25, 63, 150}

func newBenchEngine(b *testing.B) (context.Context, *rulesengine.Engine) {
	b.Helper()

	ctx := context.Background()
	e, err := rulesengine.NewEngine(ctx, rulesengine.WithPoolSize(1))
	require.NoError(b, err)
	b.Cleanup(func() { _ = e.Close(ctx) })

	return ctx, e
}

// BenchmarkCheckFlagsIndividually is how the API evaluates a multi-flag check
// today: one engine call per flag, each carrying the company trimmed to that
// flag. The trim is inside the timed loop because the API pays it per check.
func BenchmarkCheckFlagsIndividually(b *testing.B) {
	for _, count := range benchFlagCounts {
		b.Run(fmt.Sprintf("flags=%d", count), func(b *testing.B) {
			ctx, e := newBenchEngine(b)

			flags := benchFlags(count)
			company := benchCompany(flags, 20)

			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				for _, flag := range flags {
					if _, err := e.CheckFlag(ctx, trimCompanyForFlag(company, flag), nil, flag); err != nil {
						b.Fatal(err)
					}
				}
			}
		})
	}
}

// BenchmarkCheckFlagsIndividuallyUntrimmed is the same loop sending the whole
// company every time. Kept for contrast: the distance between this and the
// trimmed series is what the API's trim already recovered, and measuring bulk
// against this one would claim that saving twice.
func BenchmarkCheckFlagsIndividuallyUntrimmed(b *testing.B) {
	for _, count := range benchFlagCounts {
		b.Run(fmt.Sprintf("flags=%d", count), func(b *testing.B) {
			ctx, e := newBenchEngine(b)

			flags := benchFlags(count)
			company := benchCompany(flags, 20)

			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				for _, flag := range flags {
					if _, err := e.CheckFlag(ctx, company, nil, flag); err != nil {
						b.Fatal(err)
					}
				}
			}
		})
	}
}

// BenchmarkCheckFlagsBulk is the same work through the bulk export, trimmed to
// the union across the set as a bulk caller would.
func BenchmarkCheckFlagsBulk(b *testing.B) {
	for _, count := range benchFlagCounts {
		b.Run(fmt.Sprintf("flags=%d", count), func(b *testing.B) {
			ctx, e := newBenchEngine(b)

			flags := benchFlags(count)
			company := benchCompany(flags, 20)

			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				if _, err := e.CheckFlags(ctx, trimCompanyForFlags(company, flags), nil, flags); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkCheckFlagsByCompanySize holds the flag count at the 63 the API's
// bulk route evaluates and grows the company instead. With the trim in place
// this isolates what the trim cannot reach -- traits and metrics -- which is
// the part the individual path still re-sends per flag.
func BenchmarkCheckFlagsByCompanySize(b *testing.B) {
	const flagCount = 63

	for _, size := range []struct {
		features int
		traits   int
	}{{5, 2}, {25, 10}, {63, 20}, {200, 50}} {
		flags := benchFlags(flagCount)
		company := benchCompany(benchFlags(size.features), size.traits)

		label := fmt.Sprintf("features=%d", size.features)

		b.Run(label+"/individual", func(b *testing.B) {
			ctx, e := newBenchEngine(b)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, flag := range flags {
					if _, err := e.CheckFlag(ctx, trimCompanyForFlag(company, flag), nil, flag); err != nil {
						b.Fatal(err)
					}
				}
			}
		})

		b.Run(label+"/bulk", func(b *testing.B) {
			ctx, e := newBenchEngine(b)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if _, err := e.CheckFlags(ctx, trimCompanyForFlags(company, flags), nil, flags); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
