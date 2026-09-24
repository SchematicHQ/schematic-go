package rulesengine_test

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/schematichq/schematic-go/rulesengine"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// testClock is a controllable time source, safe for concurrent use as WithNow
// requires.
type testClock struct{ nanos atomic.Int64 }

func newTestClock(t time.Time) *testClock {
	c := &testClock{}
	c.Set(t)
	return c
}

func (c *testClock) Now() time.Time  { return time.Unix(0, c.nanos.Load()).UTC() }
func (c *testClock) Set(t time.Time) { c.nanos.Store(t.UnixNano()) }

// The injected instants sit years from today, so a pass cannot come from the
// wall clock landing on the expected boundary by chance.
var leapDay = time.Date(2020, time.February, 29, 15, 30, 0, 0, time.UTC)

func newClockedEngine(t *testing.T, opts ...rulesengine.EngineOption) *rulesengine.Engine {
	t.Helper()
	ctx := context.Background()
	e, err := rulesengine.NewEngine(ctx, opts...)
	require.NoError(t, err)
	t.Cleanup(func() { _ = e.Close(ctx) })
	return e
}

func TestWithNowDrivesMetricPeriods(t *testing.T) {
	ctx := context.Background()
	clock := newTestClock(leapDay)
	e := newClockedEngine(t, rulesengine.WithNow(clock.Now))

	cases := []struct {
		period    rulesengine.MetricPeriod
		cur, next time.Time
	}{
		{
			rulesengine.MetricPeriodCurrentDay,
			time.Date(2020, time.February, 29, 0, 0, 0, 0, time.UTC),
			time.Date(2020, time.March, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			// 2020-02-29 is a Saturday; weeks start Sunday.
			rulesengine.MetricPeriodCurrentWeek,
			time.Date(2020, time.February, 23, 0, 0, 0, 0, time.UTC),
			time.Date(2020, time.March, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			rulesengine.MetricPeriodCurrentMonth,
			time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2020, time.March, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tc := range cases {
		t.Run(string(tc.period), func(t *testing.T) {
			cur, err := e.CurrentMetricPeriodStart(ctx, tc.period)
			require.NoError(t, err)
			require.NotNil(t, cur)
			assert.Equal(t, tc.cur, *cur)

			next, err := e.NextMetricPeriodStart(ctx, tc.period)
			require.NoError(t, err)
			require.NotNil(t, next)
			assert.Equal(t, tc.next, *next)
		})
	}
}

// The source must be read per call, not captured at construction: pooled
// instances serve every caller, and a test that advances its clock expects the
// next evaluation to see it. A pool of one forces both calls onto the same
// instance.
func TestWithNowReadEachCall(t *testing.T) {
	ctx := context.Background()
	clock := newTestClock(leapDay)
	e := newClockedEngine(t, rulesengine.WithNow(clock.Now), rulesengine.WithPoolSize(1))

	before, err := e.CurrentMetricPeriodStart(ctx, rulesengine.MetricPeriodCurrentMonth)
	require.NoError(t, err)
	require.NotNil(t, before)
	assert.Equal(t, time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC), *before)

	// One millisecond past midnight crosses into March.
	clock.Set(time.Date(2020, time.March, 1, 0, 0, 0, int(time.Millisecond), time.UTC))

	after, err := e.CurrentMetricPeriodStart(ctx, rulesengine.MetricPeriodCurrentMonth)
	require.NoError(t, err)
	require.NotNil(t, after)
	assert.Equal(t, time.Date(2020, time.March, 1, 0, 0, 0, 0, time.UTC), *after)
}

// Flag checks read the same source: a monthly metric's reset time follows the
// injected clock, including after it advances.
func TestWithNowDrivesCheckFlagResetTime(t *testing.T) {
	ctx := context.Background()
	clock := newTestClock(leapDay)
	e := newClockedEngine(t, rulesengine.WithNow(clock.Now), rulesengine.WithPoolSize(1))

	company := createTestCompany()
	rule := createTestRule()
	// The engine fills usage fields, reset time included, only for
	// entitlement rules.
	rule.RuleType = rulesengine.RuleTypePlanEntitlement
	condition := createTestCondition(rulesengine.ConditionTypeMetric)
	condition.Operator = rulesengine.ComparableOperatorLte
	condition.MetricPeriod = ptr(rulesengine.MetricPeriodCurrentMonth)
	rule.Conditions = []*rulesengine.Condition{condition}
	company.Metrics = append(company.Metrics,
		createTestMetric(company, *condition.EventSubtype, rulesengine.MetricPeriodCurrentMonth, 5))
	flag := createTestFlag()
	flag.Rules = []*rulesengine.Rule{rule}

	result, err := e.CheckFlag(ctx, company, nil, flag)
	require.NoError(t, err)
	require.NotNil(t, result.FeatureUsageResetAt)
	assert.True(t, result.FeatureUsageResetAt.Equal(time.Date(2020, time.March, 1, 0, 0, 0, 0, time.UTC)),
		"got %s", result.FeatureUsageResetAt)

	clock.Set(time.Date(2020, time.March, 15, 0, 0, 0, 0, time.UTC))

	result, err = e.CheckFlag(ctx, company, nil, flag)
	require.NoError(t, err)
	require.NotNil(t, result.FeatureUsageResetAt)
	assert.True(t, result.FeatureUsageResetAt.Equal(time.Date(2020, time.April, 1, 0, 0, 0, 0, time.UTC)),
		"got %s", result.FeatureUsageResetAt)
}

// Omitting the option, or passing nil, keeps the wall clock.
func TestDefaultClockIsWallClock(t *testing.T) {
	ctx := context.Background()
	for name, opts := range map[string][]rulesengine.EngineOption{
		"omitted": nil,
		"nil":     {rulesengine.WithNow(nil)},
	} {
		t.Run(name, func(t *testing.T) {
			e := newClockedEngine(t, opts...)

			before := time.Now().UTC()
			got, err := e.CurrentMetricPeriodStart(ctx, rulesengine.MetricPeriodCurrentDay)
			require.NoError(t, err)
			after := time.Now().UTC()
			require.NotNil(t, got)

			// Bracket the call so a run straddling UTC midnight still passes.
			want := []time.Time{
				time.Date(before.Year(), before.Month(), before.Day(), 0, 0, 0, 0, time.UTC),
				time.Date(after.Year(), after.Month(), after.Day(), 0, 0, 0, 0, time.UTC),
			}
			assert.Contains(t, want, *got)
		})
	}
}
