package client

import (
	"time"
)

// eventOptions accumulates the optional metadata fields applied by
// TrackOption / IdentifyOption.
type eventOptions struct {
	idempotencyKey     *string
	sentAt             *time.Time
	trustedClientClock *bool
	backfill           *bool
	traits             map[string]any
	prewarm            []string
}

// TrackOption configures optional metadata on a Track call. Fields map
// directly to the equivalent properties on the capture-service event
// payload; unset fields are omitted from the wire payload.
type TrackOption func(*eventOptions)

// IdentifyOption configures optional metadata on an Identify call. Only
// IdempotencyKey is exposed for identify events.
type IdentifyOption func(*eventOptions)

// WithTrackIdempotencyKey supplies a client-side dedupe key for a Track
// event. Duplicate events with the same key (scoped to the environment)
// are dropped server-side for 24 hours.
func WithTrackIdempotencyKey(key string) TrackOption {
	return func(o *eventOptions) { o.idempotencyKey = &key }
}

// WithTrackSentAt overrides the SDK's default SentAt (UtcNow at enqueue
// time) for a Track event. Required when WithTrustedClientClock is set.
func WithTrackSentAt(t time.Time) TrackOption {
	return func(o *eventOptions) { o.sentAt = &t }
}

// WithTrustedClientClock instructs the capture service to use the
// client-supplied SentAt as the effective event timestamp instead of
// server receipt time. Requires a secret API key and a SentAt value.
func WithTrustedClientClock(b bool) TrackOption {
	return func(o *eventOptions) { o.trustedClientClock = &b }
}

// WithBackfill imports historical data without affecting billing.
// Requires a secret API key and implies trusted_client_clock.
func WithBackfill(b bool) TrackOption {
	return func(o *eventOptions) { o.backfill = &b }
}

// WithTrackTraits attaches traits to a Track event. Go carries traits on the
// event body, so this is for the paths that build the body for you, notably
// TrackWithReservation. On a Track whose body already names traits, these are
// merged in and win on a key collision; the caller's body is never written
// through.
func WithTrackTraits(traits map[string]any) TrackOption {
	return func(o *eventOptions) { o.traits = traits }
}

// WithIdentifyPrewarm acquires a credit lease per credit type once the identify
// is enqueued, so the session's first Check does not pay the acquire round trip.
// Client mode only, and best effort: it runs in the background and logs a
// failure rather than holding up the identify.
func WithIdentifyPrewarm(creditTypeIDs []string) IdentifyOption {
	return func(o *eventOptions) { o.prewarm = creditTypeIDs }
}

// WithIdentifyIdempotencyKey supplies a client-side dedupe key for an
// Identify event. Duplicate events with the same key (scoped to the
// environment) are dropped server-side for 24 hours.
func WithIdentifyIdempotencyKey(key string) IdentifyOption {
	return func(o *eventOptions) { o.idempotencyKey = &key }
}
