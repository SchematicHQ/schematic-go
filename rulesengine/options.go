package rulesengine

// CheckFlagOption configures a single flag check.
type CheckFlagOption func(*checkFlagOptions)

// eventUsage is a simulated quantity scoped to a specific event subtype.
//
// Field names must match the Rust engine's EventUsage, which -- unlike its
// result types -- is serialized snake_case in both directions.
type eventUsage struct {
	EventSubtype string `json:"event_subtype"`
	Quantity     int64  `json:"quantity"`
}

// checkFlagOptions is the preflight envelope passed to the engine. It mirrors
// the Rust engine::CheckFlagOptions wire shape.
type checkFlagOptions struct {
	CreditCost map[string]float64 `json:"credit_cost,omitempty"`
	Usage      *int64             `json:"usage,omitempty"`
	EventUsage *eventUsage        `json:"event_usage,omitempty"`
}

func newCheckFlagOptions() *checkFlagOptions {
	return &checkFlagOptions{CreditCost: map[string]float64{}}
}

// isZero reports whether any preflight option was supplied. When none were, the
// envelope omits "options" entirely and the engine uses its defaults.
func (o *checkFlagOptions) isZero() bool {
	return len(o.CreditCost) == 0 && o.Usage == nil && o.EventUsage == nil
}

// validate enforces the negative-quantity invariants.
//
// The engine validates these too, but it can only report failures as a string
// on the result. Checking here lets callers keep receiving the typed sentinel
// errors (ErrorNegativePreflightUsage, ErrorNegativePreflightCreditCost) that
// carry a StatusCode, and avoids a wasm round trip for input we know is bad.
func (o *checkFlagOptions) validate() error {
	if o.Usage != nil && *o.Usage < 0 {
		return ErrorNegativePreflightUsage
	}
	if o.EventUsage != nil && o.EventUsage.Quantity < 0 {
		return ErrorNegativePreflightUsage
	}
	for _, c := range o.CreditCost {
		if c < 0 {
			return ErrorNegativePreflightCreditCost
		}
	}
	return nil
}

// WithCreditCost gates a credit-balance condition on `balance >= cost` when the
// condition's credit_id matches. Lets callers supply an already-computed
// per-call cost in credits, bypassing the engine's default
// quantity x consumption_rate math. Highest precedence on credit-balance
// conditions when supplied. Call multiple times to attach costs for several
// credit types in the same check.
//
// Unlike WithUsage / WithEventUsage, a zero cost is not treated as a no-op:
// cost is gated as-is, so WithCreditCost(id, 0) passes whenever the balance is
// non-negative (the "this call is free" semantic). Callers who want to skip the
// override should omit the option entirely. Negative costs are rejected by
// CheckFlag with ErrorNegativePreflightCreditCost.
func WithCreditCost(creditID string, cost float64) CheckFlagOption {
	return func(o *checkFlagOptions) {
		o.CreditCost[creditID] = cost
	}
}

// WithUsage simulates additional usage of a generic quantity for any numeric
// condition encountered while evaluating rules. For metric conditions, the
// quantity is added to the current metric value. For trait conditions with an
// int-comparable trait, it's added to the trait value. For credit-balance
// conditions, the credit cost compared against the balance is
// quantity x consumption_rate. Zero is a no-op; negative quantities are
// rejected by CheckFlag with ErrorNegativePreflightUsage.
func WithUsage(quantity int64) CheckFlagOption {
	return func(o *checkFlagOptions) {
		o.Usage = &quantity
	}
}

// WithEventUsage simulates additional usage of a specific event subtype.
// Applied to metric conditions whose event_subtype matches (quantity is added
// to the current metric value) and to credit-balance conditions whose
// event_subtype matches (compared as quantity x consumption_rate against the
// balance). Preferred over WithUsage on those conditions when the caller knows
// the specific subtype. Zero is a no-op; negative quantities are rejected by
// CheckFlag with ErrorNegativePreflightUsage. Calling this more than once
// replaces the previous pair (last write wins, matching WithUsage).
func WithEventUsage(eventSubtype string, quantity int64) CheckFlagOption {
	return func(o *checkFlagOptions) {
		o.EventUsage = &eventUsage{EventSubtype: eventSubtype, Quantity: quantity}
	}
}
