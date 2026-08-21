package rulesengine

import (
	"context"
	"testing"
	"time"
)

// The boundaries here are the engine's, and the point of the bindings is that
// callers read them rather than deriving their own. These assertions therefore
// pin the engine's actual conventions -- UTC midnight, Sunday-start weeks --
// because those are the things an independent implementation gets wrong.
func TestMetricPeriodStart(t *testing.T) {
	ctx := context.Background()
	e, err := NewEngine(ctx)
	if err != nil {
		t.Fatalf("NewEngine: %v", err)
	}
	defer func() { _ = e.Close(ctx) }()

	now := time.Now().UTC()

	t.Run("current day is UTC midnight today", func(t *testing.T) {
		got, err := e.CurrentMetricPeriodStart(ctx, MetricPeriodCurrentDay)
		if err != nil {
			t.Fatalf("CurrentMetricPeriodStart: %v", err)
		}
		if got == nil {
			t.Fatal("expected a boundary for current_day")
		}
		want := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		if !got.Equal(want) {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("current month is the first of the month", func(t *testing.T) {
		got, err := e.CurrentMetricPeriodStart(ctx, MetricPeriodCurrentMonth)
		if err != nil {
			t.Fatalf("CurrentMetricPeriodStart: %v", err)
		}
		if got == nil {
			t.Fatal("expected a boundary for current_month")
		}
		want := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
		if !got.Equal(want) {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("weeks start Sunday, not Monday", func(t *testing.T) {
		// The convention an independent implementation is most likely to get
		// wrong: ISO weeks start Monday, this engine starts Sunday.
		got, err := e.CurrentMetricPeriodStart(ctx, MetricPeriodCurrentWeek)
		if err != nil {
			t.Fatalf("CurrentMetricPeriodStart: %v", err)
		}
		if got == nil {
			t.Fatal("expected a boundary for current_week")
		}
		if got.Weekday() != time.Sunday {
			t.Errorf("week starts on %s, expected Sunday", got.Weekday())
		}
		if got.After(now) {
			t.Errorf("week start %s is in the future", got)
		}
	})

	t.Run("all time has no boundary", func(t *testing.T) {
		got, err := e.CurrentMetricPeriodStart(ctx, MetricPeriodAllTime)
		if err != nil {
			t.Fatalf("CurrentMetricPeriodStart: %v", err)
		}
		if got != nil {
			t.Errorf("expected nil for all_time, got %s", got)
		}
	})

	t.Run("next period is after current", func(t *testing.T) {
		cur, err := e.CurrentMetricPeriodStart(ctx, MetricPeriodCurrentMonth)
		if err != nil {
			t.Fatalf("current: %v", err)
		}
		next, err := e.NextMetricPeriodStart(ctx, MetricPeriodCurrentMonth)
		if err != nil {
			t.Fatalf("next: %v", err)
		}
		if cur == nil || next == nil {
			t.Fatal("expected both boundaries")
		}
		if !next.After(*cur) {
			t.Errorf("next %s is not after current %s", next, cur)
		}
	})

	// A company without a subscription is not an error and not nil: the engine
	// falls back to the calendar month. Worth pinning, because it is the kind of
	// rule a reimplementation would omit -- it is stated only in the Rust.
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)

	t.Run("no subscription falls back to the calendar month", func(t *testing.T) {
		got, err := e.CurrentMetricPeriodStartForSubscription(ctx, &Company{ID: "co_1"})
		if err != nil {
			t.Fatalf("CurrentMetricPeriodStartForSubscription: %v", err)
		}
		if got == nil || !got.Equal(monthStart) {
			t.Errorf("got %v, want %s", got, monthStart)
		}
	})

	t.Run("a nil company takes the same fallback, with no wasm write", func(t *testing.T) {
		got, err := e.CurrentMetricPeriodStartForSubscription(ctx, nil)
		if err != nil {
			t.Fatalf("nil company: %v", err)
		}
		if got == nil || !got.Equal(monthStart) {
			t.Errorf("got %v, want %s", got, monthStart)
		}
	})

	t.Run("a subscription period is used when present", func(t *testing.T) {
		// period_start two months back on the 9th: the current period should
		// begin on the 9th of a month, not the 1st.
		start := now.AddDate(0, -2, 0)
		start = time.Date(start.Year(), start.Month(), 9, 3, 0, 0, 0, time.UTC)
		company := &Company{
			ID: "co_2",
			Subscription: &Subscription{
				ID:          "sub_1",
				PeriodStart: start,
				PeriodEnd:   start.AddDate(0, 3, 0),
			},
		}
		got, err := e.CurrentMetricPeriodStartForSubscription(ctx, company)
		if err != nil {
			t.Fatalf("CurrentMetricPeriodStartForSubscription: %v", err)
		}
		if got == nil {
			t.Fatal("expected a boundary from the subscription")
		}
		if got.Day() != 9 {
			t.Errorf("period starts on day %d, expected the subscription day 9 (got %s)", got.Day(), got)
		}
	})
}
