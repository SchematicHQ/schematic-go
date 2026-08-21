package rulesengine

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/tetratelabs/wazero/api"
)

// noPeriodStart is what the module returns when a period has no start: an
// all-time period, or a company with no billing subscription. It is a sentinel
// rather than an error because "no boundary" is an ordinary answer, not a
// failure.
const noPeriodStart int64 = -1

// metricPeriodToInt mirrors MetricPeriod::from_int on the Rust side, which is
// the integer the module actually reads. Keep the two in step: an unrecognized
// value is silently treated as all-time there, so a mismatch here degrades to
// "no boundary" rather than surfacing.
func metricPeriodToInt(p MetricPeriod) int32 {
	switch p {
	case MetricPeriodAllTime:
		return 0
	case MetricPeriodCurrentDay:
		return 1
	case MetricPeriodCurrentMonth:
		return 2
	case MetricPeriodCurrentWeek:
		return 3
	default:
		return 0
	}
}

// CurrentMetricPeriodStart returns when the current calendar metric period
// began, or nil for all-time.
//
// The boundaries are the engine's, not the caller's: weeks start Sunday and
// every boundary is UTC midnight. Deriving them independently is what makes a
// usage counter and a flag evaluation disagree about which period an event
// falls in, so read them from here rather than recomputing.
func (e *Engine) CurrentMetricPeriodStart(ctx context.Context, period MetricPeriod) (*time.Time, error) {
	return e.callPeriodCalendar(ctx, period, func(i *instance) api.Function { return i.curPeriodCalendar })
}

// NextMetricPeriodStart returns when the current calendar metric period ends
// (equivalently, when the next one begins), or nil for all-time.
func (e *Engine) NextMetricPeriodStart(ctx context.Context, period MetricPeriod) (*time.Time, error) {
	return e.callPeriodCalendar(ctx, period, func(i *instance) api.Function { return i.nextPeriodCalendar })
}

// CurrentMetricPeriodStartForSubscription returns when the company's current
// billing period began, or nil when it has no subscription.
func (e *Engine) CurrentMetricPeriodStartForSubscription(ctx context.Context, company *Company) (*time.Time, error) {
	return e.callPeriodSubscription(ctx, company, func(i *instance) api.Function { return i.curPeriodSubscript })
}

// NextMetricPeriodStartForSubscription returns when the company's current
// billing period ends, or nil when it has no subscription.
func (e *Engine) NextMetricPeriodStartForSubscription(ctx context.Context, company *Company) (*time.Time, error) {
	return e.callPeriodSubscription(ctx, company, func(i *instance) api.Function { return i.nextPeriodSubscript })
}

// acquire takes an instance from the pool, honoring cancellation the same way
// CheckFlag does: a caller with an expired context returns rather than blocking
// on a fully busy pool.
func (e *Engine) acquire(ctx context.Context) (*instance, func(), error) {
	select {
	case inst := <-e.pool:
		return inst, func() { e.pool <- inst }, nil
	case <-ctx.Done():
		return nil, nil, ctx.Err()
	}
}

// syncClock hands the module the current time. The raw wasm32-unknown-unknown
// build has no system clock, and every metric-period boundary is computed
// relative to now, so skipping this yields boundaries measured from the epoch.
func (i *instance) syncClock(ctx context.Context) error {
	if i.setTime == nil {
		return nil
	}
	if _, err := i.setTime.Call(ctx, api.EncodeI64(time.Now().UnixMilli())); err != nil {
		return fmt.Errorf("rules engine setCurrentTimeMillis: %w", err)
	}
	return nil
}

// marshalCompany serializes a company for the subscription-based exports, which
// take the company alone rather than a check envelope.
func marshalCompany(company *Company) ([]byte, error) {
	out, err := json.Marshal(company)
	if err != nil {
		return nil, fmt.Errorf("marshal company: %w", err)
	}
	return out, nil
}

func (e *Engine) callPeriodCalendar(
	ctx context.Context,
	period MetricPeriod,
	pick func(*instance) api.Function,
) (*time.Time, error) {
	inst, release, err := e.acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer release()

	fn := pick(inst)
	if fn == nil {
		return nil, ErrorUnexpected
	}
	if err := inst.syncClock(ctx); err != nil {
		return nil, err
	}

	res, err := fn.Call(ctx, api.EncodeI32(metricPeriodToInt(period)))
	if err != nil {
		return nil, fmt.Errorf("rules engine metric period: %w", err)
	}

	return decodePeriodStart(res), nil
}

func (e *Engine) callPeriodSubscription(
	ctx context.Context,
	company *Company,
	pick func(*instance) api.Function,
) (*time.Time, error) {
	// A nil company is passed as a zero-length buffer, which the module reads
	// as "no company" and answers with the no-boundary sentinel.
	var input []byte
	if company != nil {
		var err error
		if input, err = marshalCompany(company); err != nil {
			return nil, err
		}
	}

	inst, release, err := e.acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer release()

	fn := pick(inst)
	if fn == nil {
		return nil, ErrorUnexpected
	}
	if err := inst.syncClock(ctx); err != nil {
		return nil, err
	}

	var ptr uint32
	if len(input) > 0 {
		alloc, err := inst.alloc.Call(ctx, uint64(len(input)))
		if err != nil {
			return nil, fmt.Errorf("rules engine alloc: %w", err)
		}
		ptr = api.DecodeU32(alloc[0])
		defer func() { _, _ = inst.dealloc.Call(ctx, uint64(ptr), uint64(len(input))) }()

		if !inst.mem.Write(ptr, input) {
			return nil, fmt.Errorf("rules engine: writing %d bytes at %d exceeds memory", len(input), ptr)
		}
	}

	res, err := fn.Call(ctx, uint64(ptr), uint64(len(input)))
	if err != nil {
		return nil, fmt.Errorf("rules engine metric period: %w", err)
	}

	return decodePeriodStart(res), nil
}

// decodePeriodStart turns the module's unix-seconds return into a time,
// mapping the sentinel to nil.
func decodePeriodStart(res []uint64) *time.Time {
	if len(res) == 0 {
		return nil
	}
	// wazero returns i64 results as a raw uint64; there is no DecodeI64.
	secs := int64(res[0])
	if secs == noPeriodStart {
		return nil
	}
	t := time.Unix(secs, 0).UTC()

	return &t
}
