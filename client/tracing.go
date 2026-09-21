package client

import (
	"context"
	"time"

	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/leases"
	"github.com/schematichq/schematic-go/tracing"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// startSpan starts a span for one of the SDK's own operations, as a child of
// whatever span the caller's context already carries.
//
// Attributes are deliberately not taken here. Building them costs allocations
// on every call, and a check served from the local cache is hot enough for that
// to matter, so callers set them behind span.IsRecording() instead. With the
// default no-op provider that check is false and nothing is built at all.
func (c *SchematicClient) startSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	return c.tracers.For(ctx).Start(ctx, name, tracing.InternalSpanOptions...)
}

// startAPISpan starts a span for a call that leaves the process. It is marked
// SpanKindClient so that time spent on the wire is separable from the SDK's own
// wrapper around it, which is what tells a slow check caused by the API apart
// from one caused by local evaluation.
func (c *SchematicClient) startAPISpan(ctx context.Context, name string) (context.Context, trace.Span) {
	return c.tracers.For(ctx).Start(ctx, name, tracing.ClientSpanOptions...)
}

// recordSource notes what answered a check — the local cache, DataStream, the
// API, or a default — on whichever span is current.
//
// The decision is made several layers below the span that should carry it, and
// threading a span through those layers would mean changing their signatures
// for tracing's benefit. Reading the span back out of the context keeps the
// instrumentation to one line at the point the decision is actually made.
func recordSource(ctx context.Context, source string) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}
	span.SetAttributes(tracing.AttrCheckSource.String(source))
}

// recordEvalContext records which dimensions the caller identified a company
// and user by. Only the key names: see tracing.KeyNames.
func recordEvalContext(span trace.Span, evalCtx *schematicgo.CheckFlagRequestBody) {
	if evalCtx == nil {
		return
	}
	if names := tracing.KeyNames(evalCtx.Company); names != "" {
		span.SetAttributes(tracing.AttrCompanyKeyNames.String(names))
	}
	if names := tracing.KeyNames(evalCtx.User); names != "" {
		span.SetAttributes(tracing.AttrUserKeyNames.String(names))
	}
}

// recordFlagResponse puts a resolved check on the span.
func recordFlagResponse(span trace.Span, resp *CheckFlagResponse) {
	if resp == nil {
		return
	}
	span.SetAttributes(
		tracing.AttrFlagValue.Bool(resp.Value),
		tracing.AttrFlagReason.String(resp.Reason),
	)
	if resp.CompanyID != nil {
		span.SetAttributes(tracing.AttrCompanyID.String(*resp.CompanyID))
	}
	if resp.UserID != nil {
		span.SetAttributes(tracing.AttrUserID.String(*resp.UserID))
	}
	if resp.FlagID != nil {
		span.SetAttributes(tracing.AttrFlagID.String(*resp.FlagID))
	}
}

// recordCheckResult puts a Check verdict on the span, including the credit hold
// it took and any failure behind a defaulted answer.
//
// A check that failed open still returns a usable verdict, so the span is
// marked an error to keep that visible: the caller was allowed through by their
// own fail-open choice, not by an entitlement.
func recordCheckResult(span trace.Span, result *CheckResult) {
	if result == nil {
		return
	}
	span.SetAttributes(
		tracing.AttrCheckAllowed.Bool(result.Allowed),
		tracing.AttrFlagValue.Bool(result.Value),
		tracing.AttrFlagReason.String(result.Reason),
	)
	if result.FlagID != nil {
		span.SetAttributes(tracing.AttrFlagID.String(*result.FlagID))
	}
	if r := result.Reservation; r != nil {
		span.SetAttributes(
			tracing.AttrReservationID.String(r.ID),
			tracing.AttrCreditTypeID.String(r.CreditTypeID),
			tracing.AttrCreditsReserved.Float64(r.CreditsReserved),
			tracing.AttrLeaseMode.String(string(r.Mode)),
		)
		if r.CompanyID != "" {
			span.SetAttributes(tracing.AttrCompanyID.String(r.CompanyID))
		}
	}
	if result.Error != "" {
		span.SetStatus(codes.Error, result.Error)
	}
}

// callCheckFlagAPI issues the single-flag check over the API inside its own
// client span, so that time on the wire is separable from the check around it.
func (c *SchematicClient) callCheckFlagAPI(
	ctx context.Context,
	flagKey string,
	evalCtx *schematicgo.CheckFlagRequestBody,
) (*schematicgo.CheckFlagResponse, error) {
	ctx, span := c.startAPISpan(ctx, "Schematic.API.CheckFlag")
	defer span.End()
	if span.IsRecording() {
		span.SetAttributes(tracing.AttrFlagKey.String(flagKey))
	}

	resp, err := c.Features.CheckFlag(ctx, flagKey, evalCtx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	return resp, err
}

// callCheckFlagsAPI issues the bulk check over the API inside its own client
// span.
func (c *SchematicClient) callCheckFlagsAPI(
	ctx context.Context,
	evalCtx *schematicgo.CheckFlagRequestBody,
) (*schematicgo.CheckFlagsResponse, error) {
	ctx, span := c.startAPISpan(ctx, "Schematic.API.CheckFlags")
	defer span.End()

	resp, err := c.Features.CheckFlags(ctx, evalCtx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	return resp, err
}

// tracedWireClient records the lease API calls the credit paths make.
//
// It decorates leases.WireClient rather than instrumenting the lease manager,
// so the lease logic stays free of tracing and the spans land exactly on the
// three calls that cross the network.
//
// A lease acquired on the check path nests under the check that needed it. The
// lease manager also extends and releases from its own background goroutines,
// off a context that carries no span, so those record as standalone traces.
type tracedWireClient struct {
	next    leases.WireClient
	tracers *tracing.Resolver
}

func (w *tracedWireClient) start(ctx context.Context, name string) (context.Context, trace.Span) {
	return w.tracers.For(ctx).Start(ctx, name, tracing.ClientSpanOptions...)
}

func (w *tracedWireClient) Acquire(
	ctx context.Context,
	companyID, creditTypeID string,
	requestedAmount float64,
	expiresAt time.Time,
) (*leases.LeaseGrant, error) {
	ctx, span := w.start(ctx, "Schematic.Lease.Acquire")
	defer span.End()
	if span.IsRecording() {
		span.SetAttributes(
			tracing.AttrCompanyID.String(companyID),
			tracing.AttrCreditTypeID.String(creditTypeID),
			tracing.AttrLeaseRequested.Float64(requestedAmount),
		)
	}

	grant, err := w.next.Acquire(ctx, companyID, creditTypeID, requestedAmount, expiresAt)
	recordLeaseGrant(span, grant, err)
	return grant, err
}

func (w *tracedWireClient) Extend(
	ctx context.Context,
	leaseID string,
	additionalAmount float64,
	expiresAt time.Time,
) (*leases.LeaseGrant, error) {
	ctx, span := w.start(ctx, "Schematic.Lease.Extend")
	defer span.End()
	if span.IsRecording() {
		span.SetAttributes(
			tracing.AttrLeaseID.String(leaseID),
			tracing.AttrLeaseRequested.Float64(additionalAmount),
		)
	}

	grant, err := w.next.Extend(ctx, leaseID, additionalAmount, expiresAt)
	recordLeaseGrant(span, grant, err)
	return grant, err
}

func (w *tracedWireClient) Release(ctx context.Context, leaseID string) error {
	ctx, span := w.start(ctx, "Schematic.Lease.Release")
	defer span.End()
	if span.IsRecording() {
		span.SetAttributes(tracing.AttrLeaseID.String(leaseID))
	}

	err := w.next.Release(ctx, leaseID)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	return err
}

func recordLeaseGrant(span trace.Span, grant *leases.LeaseGrant, err error) {
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return
	}
	if !span.IsRecording() || grant == nil {
		return
	}
	// GrantedAmount is the server's total for the lease, not the increment an
	// extend asked for, so it reads as the lease's size either way.
	span.SetAttributes(
		tracing.AttrLeaseID.String(grant.LeaseID),
		tracing.AttrLeaseGranted.Float64(grant.GrantedAmount),
	)
}

// tracedWire wraps a WireClient so its API calls are recorded.
func (c *SchematicClient) tracedWire(next leases.WireClient) leases.WireClient {
	return &tracedWireClient{next: next, tracers: c.tracers}
}

// recordTrackEvent puts a track event's shape on the span. Only the key names
// of the company and user maps, per tracing.KeyNames.
func recordTrackEvent(span trace.Span, body *schematicgo.EventBodyTrack) {
	span.SetAttributes(tracing.AttrEventType.String("track"))
	if body == nil {
		return
	}
	span.SetAttributes(tracing.AttrEventSubtype.String(body.Event))
	if names := tracing.KeyNames(body.Company); names != "" {
		span.SetAttributes(tracing.AttrCompanyKeyNames.String(names))
	}
	if names := tracing.KeyNames(body.User); names != "" {
		span.SetAttributes(tracing.AttrUserKeyNames.String(names))
	}
	if body.Quantity != nil {
		span.SetAttributes(tracing.AttrActualQuantity.Int64(*body.Quantity))
	}
	if body.ReservationID != nil {
		span.SetAttributes(tracing.AttrReservationID.String(*body.ReservationID))
	}
}

// recordError marks whichever span is current as failed. Like recordSource, it
// reads the span back out of the context so that a helper shared by several
// entry points records against the one the caller actually started.
func recordError(ctx context.Context, err error) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())
}
