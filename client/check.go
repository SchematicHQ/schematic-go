package client

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	schematicgo "github.com/schematichq/schematic-go"
	core "github.com/schematichq/schematic-go/core"
	"github.com/schematichq/schematic-go/datastream"
	"github.com/schematichq/schematic-go/leases"
	option "github.com/schematichq/schematic-go/option"
	"github.com/schematichq/schematic-go/rulesengine"
)

const (
	// insufficientCreditsReason mirrors what the API reports on a 200 with
	// value false for the same denial, so a caller matching on Reason has one
	// string to match either way.
	insufficientCreditsReason = "Insufficient credits"
	// reservationTrackIdempotencyPrefix namespaces the settling track event's
	// dedupe key, which is derived from the reservation ID.
	reservationTrackIdempotencyPrefix = "lease-reservation:"
	// releaseTimeout bounds a best-effort release of a hold or a lease, which
	// nobody is waiting on and which the server would eventually do itself at
	// expiry. Close gives the releases it runs this budget of their own.
	releaseTimeout = 5 * time.Second
	// closeTimeout is the budget a Close spends waiting out the work already
	// running, shared across every one of those waits.
	closeTimeout = 10 * time.Second
	// identifyFlushTimeout bounds the flush an identify does before prewarming,
	// so an API that is slow to take the event cannot hold the prewarm, or a
	// Close waiting on it, open.
	identifyFlushTimeout = 5 * time.Second
	// prewarmPollInterval is how often a prewarm retries the company fetch while
	// waiting for the entity to surface over DataStream.
	prewarmPollInterval = 100 * time.Millisecond
)

// clientOnlyLeaseFields are the config knobs that steer the local lease
// plumbing, which server mode never builds.
var clientOnlyLeaseFields = []string{
	"DefaultLeaseSize",
	"DefaultLeaseDuration",
	"LowWaterMark",
	"SweepInterval",
	"PrewarmResolveTimeout",
	"RedisClient",
	"RedisKeyPrefix",
	"Overrides",
}

// checkOptions accumulates the optional fields applied by CheckOption.
type checkOptions struct {
	usage        *float64
	eventSubtype *string
	failOpen     bool
	defaultValue *bool
	timeout      *time.Duration
}

// CheckOption configures a single SchematicClient.Check call.
type CheckOption func(*checkOptions)

// WithUsage declares how many units of the feature the operation is about to
// consume. The check holds quantity times the entitlement's consumption rate
// from the company's credit balance, and returns a Reservation to settle with
// TrackWithReservation. Fractional quantities are allowed, since a credit cost
// need not fall on a whole unit. Without it, Check is a plain flag check that
// holds nothing.
func WithUsage(quantity float64) CheckOption {
	return func(o *checkOptions) { o.usage = &quantity }
}

// WithEventSubtype names the event the usage applies to, for example
// "inference_tokens". Needed only when the flag meters more than one event.
func WithEventSubtype(s string) CheckOption {
	return func(o *checkOptions) { o.eventSubtype = &s }
}

// WithFailOpen returns the caller's default value when the check cannot gate,
// rather than denying. Checks fail closed by default.
func WithFailOpen() CheckOption {
	return func(o *checkOptions) { o.failOpen = true }
}

// WithCheckDefault overrides the client's configured flag default for this
// check, which is what a fail-open check returns when it cannot gate.
func WithCheckDefault(value bool) CheckOption {
	return func(o *checkOptions) { o.defaultValue = &value }
}

// WithCheckTimeout caps how long the API calls this check makes may take.
func WithCheckTimeout(d time.Duration) CheckOption {
	return func(o *checkOptions) { o.timeout = &d }
}

// Reservation is the handle Check returns when a credit hold was taken. Pass it
// to TrackWithReservation once the work completes.
type Reservation struct {
	// ID identifies the hold.
	ID string
	// LeaseID is the lease the hold draws from. Server mode has no lease, so
	// this mirrors ID and the field stays populated for code that reads it.
	LeaseID string
	// Mode says where the hold lives.
	Mode core.CreditLeaseMode
	// CompanyID owns the held credits.
	CompanyID string
	// CreditTypeID is the credit type the hold draws down.
	CreditTypeID string
	// EventSubtype is what the settling track event is recorded under.
	EventSubtype string
	// QuantityReserved is the units of usage the hold covers, as requested.
	QuantityReserved float64
	// CreditsReserved is QuantityReserved times ConsumptionRate.
	CreditsReserved float64
	// ConsumptionRate is the credits per unit of usage the hold was priced at.
	ConsumptionRate float64
	// ExpiresAt is when the unspent hold is refunded if nothing settles it.
	ExpiresAt time.Time
	// Company and User are the evaluation context the hold was issued for, so
	// the settling track event attributes the usage to the same entities.
	Company map[string]string
	User    map[string]string
}

// CheckResult is what Check returns. It is never nil.
type CheckResult struct {
	// Allowed says whether the caller may proceed. It mirrors Value outside the
	// credit paths.
	Allowed bool
	// Value is the flag's boolean value.
	Value bool
	// Reason explains the verdict.
	Reason string
	// FlagKey is the key that was checked.
	FlagKey string
	// FlagID is the flag's ID, when one was found.
	FlagID *string
	// Entitlement carries the matched feature entitlement, when there was one.
	Entitlement *rulesengine.FeatureEntitlement
	// Reservation is the credit hold taken for this check, when one was taken.
	Reservation *Reservation
	// Error carries the server's or the SDK's explanation when something went
	// wrong, even where the check still produced a verdict.
	Error string
}

// configureCreditLeases stores the credit lease config and says once, at
// construction, where the caller's settings will not gate the way they look
// like they will.
func (c *SchematicClient) configureCreditLeases(cfg *core.CreditLeaseConfig) {
	if cfg == nil {
		return
	}

	ctx := context.Background()
	c.creditLeases = cfg

	if c.isOffline {
		c.logger.Warn(ctx, "Credit leases are configured but the client is offline; Check returns flag defaults and holds no credits")
		return
	}

	// The mode a check will actually resolve to. "auto" without DataStream is
	// server mode, and so is every mode on a client that asked for DataStream
	// and did not get it.
	serverMode := cfg.Mode == core.CreditLeaseModeServer || (cfg.Mode != core.CreditLeaseModeClient && !c.useDataStream())

	ttl := cfg.DefaultReservationTTL
	if ttl <= 0 {
		ttl = leases.DefaultReservationTTL
	}
	// Only server mode sends this to the API. In client mode it sizes the local
	// sweep, so clamping it there would shorten holds for no reason and the
	// warning would be untrue.
	if maxTTL := leases.MaxReservationTTL - leases.ReservationTTLSkewAllowance; ttl > maxTTL {
		if serverMode {
			c.logger.Warn(ctx, fmt.Sprintf(
				"Credit lease DefaultReservationTTL of %s is longer than the API will hold credits for; server-mode holds are clamped to %s",
				ttl, maxTTL,
			))
		}
		ttl = maxTTL
	}
	c.serverReservationTTL = ttl

	if serverMode {
		if ignored := setClientOnlyLeaseFields(cfg); len(ignored) > 0 {
			c.logger.Warn(ctx, fmt.Sprintf(
				"Credit leases resolve to server mode, so %s will be ignored; those options only apply to client mode, where leases are carved up locally over DataStream",
				strings.Join(ignored, ", "),
			))
		}
	}

	if cfg.Mode == core.CreditLeaseModeClient && !c.useDataStream() {
		c.logger.Warn(ctx, "Credit lease mode is client, which requires DataStream, and DataStream is not enabled; checks are ungated and fall back to plain flag checks. Use server mode, or the auto default, to gate on credits without DataStream")
		return
	}

	if cfg.Mode != core.CreditLeaseModeClient && cfg.Mode != core.CreditLeaseModeServer && !c.useDataStream() {
		// Not a misconfiguration: auto without DataStream is the server-mode
		// default, which gates over the API instead.
		c.logger.Info(ctx, "Credit leases are configured and DataStream is not enabled, so holds are taken in server mode, one check-and-reserve call per check. Enable DataStream for client-side leases")
	}

	// Client mode carves holds out of leases this process tracks, so it needs
	// the stores and the manager. Server mode builds none of it.
	if c.effectiveLeaseMode() == core.CreditLeaseModeClient {
		c.buildClientLeasePlumbing(ctx, cfg)
	}
}

// buildClientLeasePlumbing builds the lease store, the reservation table, their
// manager, and the sweeper that returns expired holds.
//
// Lease state belongs in a shared cache so gating holds across processes. An
// explicit RedisClient wins; otherwise the DataStream cache's Redis is reused,
// so an existing setup backs leases with no second connection pool to wire up.
func (c *SchematicClient) buildClientLeasePlumbing(ctx context.Context, cfg *core.CreditLeaseConfig) {
	redisClient := cfg.RedisClient
	if redisClient == nil && c.datastreamClient != nil {
		redisClient = c.datastreamClient.RedisClient()
	}
	prefix := cfg.RedisKeyPrefix
	if prefix == "" {
		prefix = leases.DefaultKeyPrefix
	}

	if redisClient != nil {
		store := leases.NewRedisLeaseStore(redisClient, leases.RedisLeaseStoreOptions{
			KeyPrefix:            prefix,
			DefaultLeaseDuration: cfg.DefaultLeaseDuration,
		})
		c.leaseStore = store
		c.reservations = leases.NewRedisReservationStore(redisClient, store, leases.RedisReservationStoreOptions{KeyPrefix: prefix})
	} else {
		c.logger.Warn(ctx, "Credit leases are enabled without a shared Redis backend, so lease and reservation state stays in this process and gates within it only. Set CreditLeaseConfig.RedisClient, or give DataStream a Redis cache, so leases gate across every SDK instance")
		store := leases.NewInMemoryLeaseStore(leases.InMemoryLeaseStoreOptions{})
		c.leaseStore = store
		c.reservations = leases.NewInMemoryReservationStore(store, leases.InMemoryReservationStoreOptions{})
	}

	c.leaseManager = leases.NewLeaseManager(leases.NewAPIWireClient(c.Credits), c.leaseStore, leases.LeaseManagerOptions{
		Reservations: c.reservations,
		Config: leases.ResolvedLeaseConfig{
			LeaseDuration:  cfg.DefaultLeaseDuration,
			ReservationTTL: cfg.DefaultReservationTTL,
			LeaseSize:      cfg.DefaultLeaseSize,
			LowWaterMark:   cfg.LowWaterMark,
		},
		Overrides:     leaseOverrides(cfg.Overrides),
		SweepInterval: cfg.SweepInterval,
		Logger:        c.logger,
	})
	c.leaseManager.StartSweep()

	c.prewarmResolveTimeout = leases.DefaultPrewarmResolveTimeout
	if cfg.PrewarmResolveTimeout != nil {
		c.prewarmResolveTimeout = *cfg.PrewarmResolveTimeout
	}
}

// leaseOverrides maps the public per-credit-type overrides onto the lease
// package's own.
func leaseOverrides(overrides map[string]core.CreditLeaseOverride) map[string]leases.LeaseOverride {
	if len(overrides) == 0 {
		return nil
	}

	mapped := make(map[string]leases.LeaseOverride, len(overrides))
	for creditTypeID, override := range overrides {
		mapped[creditTypeID] = leases.LeaseOverride{
			LeaseDuration:  override.LeaseDuration,
			ReservationTTL: override.ReservationTTL,
			LeaseSize:      override.LeaseSize,
			LowWaterMark:   override.LowWaterMark,
		}
	}
	return mapped
}

// Prewarm acquires a lease per credit type up front, so a session's first Check
// does not pay the acquire round trip. A no-op outside client mode.
//
// When the company keys carry no ID, this fetches the company over DataStream,
// waiting up to CreditLeaseConfig.PrewarmResolveTimeout for it to surface, which
// covers a company the server has only just ingested. The fetch also primes the
// cache, so the first real Check takes the lease path instead of falling back.
func (c *SchematicClient) Prewarm(
	ctx context.Context,
	evalCtx *schematicgo.CheckFlagRequestBody,
	creditTypeIDs []string,
) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic while prewarming credit leases: %v", r)
		}
	}()

	if c.leaseManager == nil {
		if c.effectiveLeaseMode() == core.CreditLeaseModeServer {
			c.logger.Debug(ctx, "Prewarm is a no-op in server mode; there is no local lease to warm")
		} else {
			c.logger.Debug(ctx, "Prewarm was called but client-mode credit leases are not configured")
		}
		return nil
	}
	if evalCtx == nil || len(evalCtx.Company) == 0 {
		c.logger.Debug(ctx, "Prewarm needs company keys")
		return nil
	}
	if len(creditTypeIDs) == 0 {
		return nil
	}

	companyID, err := c.resolveCompanyIDForPrewarm(ctx, evalCtx.Company)
	if err != nil {
		return err
	}

	// One acquire per credit type, concurrently: they are independent round
	// trips, and a caller warming several should wait for the slowest, not for
	// the sum.
	failures := make([]error, len(creditTypeIDs))
	var warming sync.WaitGroup
	for i, creditTypeID := range creditTypeIDs {
		warming.Add(1)
		go func() {
			defer warming.Done()
			if c.leaseManager.AcquireIfNeeded(ctx, companyID, creditTypeID) == nil {
				failures[i] = fmt.Errorf("could not warm a credit lease for %s; the first check acquires instead", creditTypeID)
			}
		}()
	}
	warming.Wait()

	// Indexed rather than collected as they land, so the error a caller sees
	// does not depend on which goroutine lost.
	for _, err := range failures {
		if err != nil {
			return err
		}
	}
	return nil
}

// resolveCompanyIDForPrewarm turns company keys into the ID a lease is keyed by,
// waiting for the company to surface over DataStream when only secondary keys
// were given.
//
// An identify does not push a company into the DataStream cache, since companies
// are only streamed on request, so this fetches rather than watching an empty
// cache. A PrewarmResolveTimeout of zero keeps the cache lookup and skips the
// wait, so an already-seen company still warms.
func (c *SchematicClient) resolveCompanyIDForPrewarm(ctx context.Context, keys map[string]string) (string, error) {
	if id := keys["id"]; id != "" {
		return id, nil
	}
	if c.datastreamClient == nil {
		return "", errors.New("prewarm needs DataStream to resolve company keys to an ID")
	}
	// An earlier check or prewarm may already have cached this company, and that
	// answer costs nothing.
	if cached := c.datastreamClient.GetCachedCompany(keys); cached != nil && cached.ID != "" {
		return cached.ID, nil
	}
	if c.prewarmResolveTimeout <= 0 {
		return "", fmt.Errorf("prewarm: company %v is not cached and the resolve wait is off", keys)
	}

	deadline := time.Now().Add(c.prewarmResolveTimeout)
	for {
		company, err := c.datastreamClient.GetCompany(ctx, keys)
		if err == nil && company != nil && company.ID != "" {
			return company.ID, nil
		}
		// Expected while the socket is still connecting, and while the server
		// has yet to ingest a preceding identify.
		c.logger.Debug(ctx, fmt.Sprintf("Prewarm: DataStream company fetch failed (%v)", err))
		if !time.Now().Before(deadline) {
			return "", fmt.Errorf("prewarm: company %v did not resolve within %s; the first check acquires instead", keys, c.prewarmResolveTimeout)
		}
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-time.After(prewarmPollInterval):
		}
	}
}

// bodyWithTraits applies an option's traits to a track event, copying the
// caller's body rather than writing through it. Traits already on the body are
// kept, with the option winning a key collision.
func bodyWithTraits(body *schematicgo.EventBodyTrack, traits map[string]any) *schematicgo.EventBodyTrack {
	if body == nil || len(traits) == 0 {
		return body
	}

	copied := *body
	if len(body.Traits) == 0 {
		copied.Traits = traits
		return &copied
	}

	merged := make(map[string]any, len(body.Traits)+len(traits))
	for name, value := range body.Traits {
		merged[name] = value
	}
	for name, value := range traits {
		merged[name] = value
	}
	copied.Traits = merged
	return &copied
}

// setClientOnlyLeaseFields names the client-mode-only fields the caller set.
func setClientOnlyLeaseFields(cfg *core.CreditLeaseConfig) []string {
	set := map[string]bool{
		"DefaultLeaseSize":      cfg.DefaultLeaseSize != 0,
		"DefaultLeaseDuration":  cfg.DefaultLeaseDuration != 0,
		"LowWaterMark":          cfg.LowWaterMark != 0,
		"SweepInterval":         cfg.SweepInterval != 0,
		"PrewarmResolveTimeout": cfg.PrewarmResolveTimeout != nil,
		"RedisClient":           cfg.RedisClient != nil,
		"RedisKeyPrefix":        cfg.RedisKeyPrefix != "",
		"Overrides":             len(cfg.Overrides) > 0,
	}

	ignored := make([]string, 0, len(clientOnlyLeaseFields))
	for _, name := range clientOnlyLeaseFields {
		if set[name] {
			ignored = append(ignored, name)
		}
	}
	return ignored
}

// effectiveLeaseMode is which mode a check with usage resolves to on this
// client. The empty string means no credit gating at all.
//
// DataStream counts as available when the client built a DataStream client at
// construction, not when the socket is up, so the answer is the same for the
// client's life. A socket that drops mid-run therefore stays in client mode,
// where the flow degrades per check instead: the entity lookups it needs fail,
// and each check falls back to a plain API check that holds nothing.
func (c *SchematicClient) effectiveLeaseMode() core.CreditLeaseMode {
	if c.creditLeases == nil || c.isOffline {
		return ""
	}

	switch c.creditLeases.Mode {
	case core.CreditLeaseModeServer:
		return core.CreditLeaseModeServer
	case core.CreditLeaseModeClient:
		// Client mode rides on DataStream. Without it there is nothing to gate
		// against, so the check stays plain.
		if c.useDataStream() {
			return core.CreditLeaseModeClient
		}
		return ""
	default:
		if c.useDataStream() {
			return core.CreditLeaseModeClient
		}
		return core.CreditLeaseModeServer
	}
}

// Check evaluates a flag, holding credits for the work the caller is about to
// do when WithUsage is set and credit leases are configured.
//
// On success with a hold, the result carries a Reservation; pass it to
// TrackWithReservation once the work completes, and the unspent slice is
// refunded. Without WithUsage, or without credit leases configured, this is a
// plain flag check that holds nothing, with the caller's usage still threaded
// through as a preflight so the verdict accounts for the usage about to be
// recorded.
//
// The result is never nil.
func (c *SchematicClient) Check(
	ctx context.Context,
	evalCtx *schematicgo.CheckFlagRequestBody,
	flagKey string,
	opts ...CheckOption,
) (result *CheckResult) {
	o := &checkOptions{}
	for _, apply := range opts {
		apply(o)
	}

	defer func() {
		if r := recover(); r != nil {
			c.logger.Error(ctx, fmt.Sprintf("Panic occurred while checking flag %v", r))
			result = c.checkFailureResult(flagKey, o, "error")
		}
	}()

	if evalCtx == nil {
		evalCtx = &schematicgo.CheckFlagRequestBody{}
	}

	mode := c.effectiveLeaseMode()
	if o.usage == nil || mode == "" {
		return c.checkFallback(ctx, evalCtx, flagKey, o)
	}
	if mode == core.CreditLeaseModeServer {
		return c.checkWithServerReservation(ctx, evalCtx, flagKey, o)
	}
	// Client mode resolved from DataStream alone, with the local plumbing
	// missing. Gate over the API rather than dropping to an ungated check: the
	// caller asked for credit gating and server mode needs none of that
	// plumbing.
	if c.leaseManager == nil || c.leaseStore == nil || c.reservations == nil || c.datastreamClient == nil {
		return c.checkWithServerReservation(ctx, evalCtx, flagKey, o)
	}

	return c.checkWithClientLease(ctx, evalCtx, flagKey, o)
}

// checkWithClientLease gates one check against a lease this process holds: the
// SDK draws a tranche of credits per company and credit type, carves the hold
// out of it locally, and evaluates the flag against the lease's own balance. No
// API call per check, and a shared Redis makes the gate hold across processes.
//
// leases.CheckWithLease owns the flow, so the conformance vectors that pin it
// drive the same code this does.
func (c *SchematicClient) checkWithClientLease(
	ctx context.Context,
	evalCtx *schematicgo.CheckFlagRequestBody,
	flagKey string,
	o *checkOptions,
) *CheckResult {
	// The caller's per-check timeout governs the lease wire calls the same way
	// it governs the plain check's.
	callCtx, cancel := o.withTimeout(ctx)
	defer cancel()

	deps := leases.CheckDeps{
		DataStream:   &leaseDataStream{client: c.datastreamClient},
		Leases:       c.leaseStore,
		Reservations: c.reservations,
		Manager:      c.leaseManager,
		Logger:       c.logger,
		EmitFlagCheck: func(ctx context.Context, body *schematicgo.EventBodyFlagCheck) {
			if err := c.enqueueEvent("flag_check", schematicgo.EventBody{EventBodyFlagCheck: body}, &eventOptions{}); err != nil {
				c.logger.Error(ctx, fmt.Sprintf("Failed to enqueue flag_check event: %v", err))
			}
		},
	}
	request := leases.CheckRequest{
		FlagKey:  flagKey,
		Company:  evalCtx.Company,
		User:     evalCtx.User,
		Usage:    *o.usage,
		FailOpen: o.failOpen,
	}
	if o.eventSubtype != nil {
		request.EventSubtype = *o.eventSubtype
	}

	outcome := leases.CheckWithLease(callCtx, deps, request, func(ctx context.Context) *leases.CheckOutcome {
		return outcomeFromCheckResult(c.checkFallback(ctx, evalCtx, flagKey, o))
	})
	return checkResultFromOutcome(outcome)
}

// leaseDataStream adapts the DataStream client to the narrow surface the
// credit-lease flow drives. The flow lives in its own package, which the
// DataStream client cannot import, so the translation lands here.
type leaseDataStream struct {
	client *datastream.DataStreamClient
}

func (d *leaseDataStream) GetFlag(ctx context.Context, key string) (*rulesengine.Flag, bool) {
	return d.client.GetFlag(ctx, key)
}

func (d *leaseDataStream) GetCompany(ctx context.Context, keys map[string]string) (*rulesengine.Company, error) {
	return d.client.GetCompany(ctx, keys)
}

func (d *leaseDataStream) GetUser(ctx context.Context, keys map[string]string) (*rulesengine.User, error) {
	return d.client.GetUser(ctx, keys)
}

func (d *leaseDataStream) EvaluateFlag(
	ctx context.Context,
	flag *rulesengine.Flag,
	company *rulesengine.Company,
	user *rulesengine.User,
	opts *leases.EvalOptions,
) (*rulesengine.CheckFlagResult, error) {
	return d.client.EvaluateFlag(ctx, flag, company, user, leases.EngineCheckFlagOptions(opts)...)
}

// outcomeFromCheckResult carries a plain check's verdict back into the flow,
// which owns whether the caller ever sees it.
func outcomeFromCheckResult(result *CheckResult) *leases.CheckOutcome {
	return &leases.CheckOutcome{
		Allowed:     result.Allowed,
		Value:       result.Value,
		Reason:      result.Reason,
		FlagKey:     result.FlagKey,
		FlagID:      result.FlagID,
		Entitlement: result.Entitlement,
		Error:       result.Error,
	}
}

// checkResultFromOutcome maps what the flow resolved onto the public result,
// turning the hold it recorded into the handle the caller settles with.
func checkResultFromOutcome(outcome *leases.CheckOutcome) *CheckResult {
	result := &CheckResult{
		Allowed:     outcome.Allowed,
		Value:       outcome.Value,
		Reason:      outcome.Reason,
		FlagKey:     outcome.FlagKey,
		FlagID:      outcome.FlagID,
		Entitlement: outcome.Entitlement,
		Error:       outcome.Error,
	}
	if record := outcome.Reservation; record != nil {
		result.Reservation = &Reservation{
			ID:               record.ID,
			LeaseID:          record.LeaseID,
			Mode:             core.CreditLeaseModeClient,
			CompanyID:        record.CompanyID,
			CreditTypeID:     record.CreditTypeID,
			EventSubtype:     record.EventSubtype,
			QuantityReserved: record.QuantityReserved,
			CreditsReserved:  record.CreditsReserved,
			ConsumptionRate:  record.ConsumptionRate,
			ExpiresAt:        record.ExpiresAt,
			Company:          record.Company,
			User:             record.User,
		}
	}
	return result
}

// reservationRecord is the handle as the lease stores hold it.
func reservationRecord(reservation *Reservation) leases.ReservationRecord {
	return leases.ReservationRecord{
		ID:               reservation.ID,
		LeaseID:          reservation.LeaseID,
		CompanyID:        reservation.CompanyID,
		CreditTypeID:     reservation.CreditTypeID,
		EventSubtype:     reservation.EventSubtype,
		QuantityReserved: reservation.QuantityReserved,
		CreditsReserved:  reservation.CreditsReserved,
		ConsumptionRate:  reservation.ConsumptionRate,
		ExpiresAt:        reservation.ExpiresAt,
		Company:          reservation.Company,
		User:             reservation.User,
	}
}

// checkFallback answers a check without holding anything. The caller's usage
// still goes out as a preflight, so the verdict accounts for the usage about to
// be recorded.
func (c *SchematicClient) checkFallback(
	ctx context.Context,
	evalCtx *schematicgo.CheckFlagRequestBody,
	flagKey string,
	o *checkOptions,
) *CheckResult {
	ctx, cancel := o.withTimeout(ctx)
	defer cancel()

	// The variant that keeps the failure: the public method answers an
	// unreachable API with the flag default and no error, which would decide
	// the verdict here without the caller's fail-open choice ever being read.
	resp, err := c.checkFlagWithEntitlement(ctx, evalCtxWithPreflight(evalCtx, o), flagKey)
	if err != nil || resp == nil {
		if err == nil {
			err = fmt.Errorf("the check of flag %s returned no response", flagKey)
		}
		c.logger.Error(ctx, fmt.Sprintf("Check: flag %s could not be evaluated: %v", flagKey, err))
		return c.checkErrorResult(flagKey, o, err)
	}

	return &CheckResult{
		Allowed:     resp.Value,
		Value:       resp.Value,
		Reason:      resp.Reason,
		FlagKey:     resp.FlagKey,
		FlagID:      resp.FlagID,
		Entitlement: resp.Entitlement,
	}
}

// checkWithServerReservation gates one check on the server: a single
// check-and-reserve call evaluates the flag against the company's real balance
// and takes the hold in the same round trip. There is no lease, no local store
// and no rules engine involved.
//
// Fail-open here returns the caller's default rather than re-running the rules
// with an assumed-sufficient balance, since the call that would have answered is
// the one that failed and there is no local engine to fall back on. No
// flag_check event is enqueued: the server logs the check, the same way the
// plain REST path does.
func (c *SchematicClient) checkWithServerReservation(
	ctx context.Context,
	evalCtx *schematicgo.CheckFlagRequestBody,
	flagKey string,
	o *checkOptions,
) *CheckResult {
	usage := *o.usage
	if !leases.IsValidQuantity(usage) {
		c.logger.Error(ctx, fmt.Sprintf("Server reservation: invalid usage %v for flag %s; must be a finite, non-negative number", usage, flagKey))
		return c.checkFailureResult(flagKey, o, "invalid_usage")
	}
	if usage == 0 {
		c.logger.Debug(ctx, fmt.Sprintf("Server reservation: usage is 0 for flag %s, nothing to hold, using a plain check", flagKey))
		return c.checkFallback(ctx, evalCtx, flagKey, o)
	}

	callCtx, cancel := o.withTimeout(ctx)
	defer cancel()

	expiresAt := time.Now().UTC().Add(c.serverReservationTTL)
	// A 502 arriving after the API committed the hold leaves a retry no way to
	// tell a lost response from a rejected request. The key is minted once per
	// check and the retry loop resends this body, so the server answers every
	// attempt of this check with the one hold it already took, while the next
	// check gets its own key and its own hold.
	idempotencyKey := uuid.NewString()
	body := &schematicgo.CheckAndReserveFlagRequestBody{
		Company:        evalCtx.Company,
		User:           evalCtx.User,
		Quantity:       &usage,
		ExpiresAt:      &expiresAt,
		IdempotencyKey: &idempotencyKey,
		Preflight:      mergedPreflight(evalCtx.Preflight, o.preflight()),
	}

	resp, err := c.Features.CheckAndReserveFlag(callCtx, flagKey, body)
	if err != nil {
		// A 402 is the server's answer, not a failure to answer: it knows the
		// credits are not there. Deny regardless of fail-open, which would
		// otherwise hand out credit the balance cannot cover.
		var paymentRequired *schematicgo.PaymentRequiredError
		if errors.As(err, &paymentRequired) {
			return &CheckResult{
				FlagKey: flagKey,
				Reason:  insufficientCreditsReason,
				Error:   paymentRequiredMessage(paymentRequired),
			}
		}
		c.logger.Error(ctx, fmt.Sprintf("Server reservation: check-and-reserve for flag %s failed: %v", flagKey, err))
		return c.checkFailureResult(flagKey, o, "server_reservation_failed")
	}
	if resp == nil || resp.Data == nil {
		c.logger.Error(ctx, fmt.Sprintf("Server reservation: check-and-reserve for flag %s returned no data", flagKey))
		return c.checkFailureResult(flagKey, o, "server_reservation_failed")
	}

	data := resp.Data
	result := &CheckResult{
		Allowed:     data.Value,
		Value:       data.Value,
		Reason:      data.Reason,
		FlagKey:     flagKey,
		FlagID:      data.FlagID,
		Entitlement: toRulesEngineEntitlement(data.Entitlement),
	}
	if data.Flag != "" {
		result.FlagKey = data.Flag
	}
	if data.Error != nil {
		result.Error = *data.Error
	}

	// No hold comes back when the flag denied, the credits were short, or the
	// feature is not credit-metered. Nothing was held, so nothing to release.
	held := data.Reservation
	if !data.Value || held == nil {
		return result
	}

	// The settling track event is named by the event subtype; the caller's wins,
	// otherwise the server names it on the hold. With neither, the hold could
	// never be settled, so release it now instead of parking the credits until
	// the TTL.
	eventSubtype := ""
	if o.eventSubtype != nil {
		eventSubtype = *o.eventSubtype
	} else if held.EventSubtype != nil {
		eventSubtype = *held.EventSubtype
	}
	if eventSubtype == "" {
		c.logger.Error(ctx, fmt.Sprintf("Server reservation: reservation %s for flag %s names no event subtype; releasing it, since it could never be settled", held.ID, flagKey))
		// Detached and unretried: the caller is not waiting on this, a release
		// the caller's cancellation cut short would park the credits until the
		// hold expires, and the call is idempotent, so a retry buys nothing the
		// expiry does not.
		releaseCtx, cancelRelease := context.WithTimeout(context.WithoutCancel(ctx), releaseTimeout)
		if _, err := c.Credits.ReleaseCreditReservation(releaseCtx, held.ID, option.WithoutRetries()); err != nil {
			c.logger.Warn(ctx, fmt.Sprintf("Server reservation: failed to release %s (%v); its hold is refunded when it expires", held.ID, err))
		}
		cancelRelease()
		if !o.failOpen {
			return c.checkFailureResult(flagKey, o, "missing_event_subtype")
		}
		// Fail-open means assume the credits are there, and the server has
		// already evaluated the flag and said yes. Only the settle is
		// impossible, so its verdict stands.
		result.Error = "missing_event_subtype"
		return result
	}

	result.Reservation = &Reservation{
		ID: held.ID,
		// No lease exists in server mode; mirror the ID so the field stays
		// populated and a handle round-trips through code that reads it.
		LeaseID:          held.ID,
		Mode:             core.CreditLeaseModeServer,
		CompanyID:        held.CompanyID,
		CreditTypeID:     held.CreditTypeID,
		EventSubtype:     eventSubtype,
		QuantityReserved: held.QuantityReserved,
		CreditsReserved:  held.CreditsReserved,
		ConsumptionRate:  held.ConsumptionRate,
		ExpiresAt:        held.ExpiresAt,
		Company:          evalCtx.Company,
		User:             evalCtx.User,
	}
	return result
}

// TrackWithReservation settles a reservation issued by Check with the actual
// usage. The track event carries the reservation ID, and the server settles the
// hold, refunding the unspent slice, when it processes the event. The event's
// idempotency key is derived from the reservation ID, so a duplicate or retried
// settle is dropped server-side rather than billed twice.
func (c *SchematicClient) TrackWithReservation(
	ctx context.Context,
	reservation *Reservation,
	actualQuantity int64,
	opts ...TrackOption,
) {
	if c.isOffline {
		return
	}
	// A check can allow without taking a hold, for instance when the feature is
	// not credit-metered or the check failed open, and callers pass
	// result.Reservation straight through. The usage still has to be recorded,
	// but only a plain Track can do it.
	if reservation == nil {
		c.logger.Error(ctx, "TrackWithReservation: no reservation to settle; the check allowed without taking a hold. Report the usage with Track instead")
		return
	}
	// A quantity the server cannot bill must reach neither the event nor the
	// hold: skip the settle and let the hold refund itself at its TTL.
	if actualQuantity < 0 {
		c.logger.Error(ctx, fmt.Sprintf("TrackWithReservation: invalid quantity %d for reservation %s; must be non-negative. Skipping the settle, the hold is refunded at its TTL", actualQuantity, reservation.ID))
		return
	}

	// A server-mode hold lives on the server and settles by ID. A client-mode
	// one is consumed against its local lease first and routes through that
	// lease's sub-ledger instead; never send both, since the server prefers the
	// lease ID.
	var body *schematicgo.EventBodyTrack
	if reservation.Mode == core.CreditLeaseModeServer {
		body = &schematicgo.EventBodyTrack{
			Company:       reservation.Company,
			Event:         reservation.EventSubtype,
			Quantity:      &actualQuantity,
			ReservationID: &reservation.ID,
			User:          reservation.User,
		}
	} else {
		body = c.settleClientReservation(ctx, reservation, actualQuantity)
	}

	// The caller's options come last so an explicit idempotency key still wins.
	trackOpts := append(
		[]TrackOption{WithTrackIdempotencyKey(reservationTrackIdempotencyPrefix + reservation.ID)},
		opts...,
	)
	c.Track(ctx, body, trackOpts...)
}

// settleClientReservation consumes a client-mode hold against its lease and
// hands back the event that bills it.
//
// The server is the source of truth for real consumption, so a settle that
// cannot run locally still emits: the event's idempotency key is what keeps the
// retry from billing twice.
func (c *SchematicClient) settleClientReservation(
	ctx context.Context,
	reservation *Reservation,
	actualQuantity int64,
) *schematicgo.EventBodyTrack {
	record := reservationRecord(reservation)
	if c.reservations == nil {
		// The handle came from a lease-configured client, so the event still
		// needs its lease ID and its dedupe key even though this client holds
		// nothing to settle.
		c.logger.Warn(ctx, "TrackWithReservation: client-mode credit leases are not configured here; emitting an unsettled track")
		return leases.BuildTrackEvent(record, actualQuantity)
	}

	outcome := leases.SettleReservation(ctx, c.reservations, record, float64(actualQuantity))
	switch {
	case outcome.Err != nil:
		c.logger.Warn(ctx, fmt.Sprintf("TrackWithReservation: failed to settle reservation %s locally (%v); emitting the track anyway", reservation.ID, outcome.Err))
	case !outcome.SettledLocally:
		c.logger.Debug(ctx, fmt.Sprintf("TrackWithReservation: reservation %s was not settled locally (swept at its TTL, already settled, or the store is unreachable); the track is keyed for server-side dedupe", reservation.ID))
	}
	return outcome.Track
}

// checkFailureResult resolves a check that could not gate. Fail-closed denies;
// fail-open returns the caller's resolved default, since server mode has no
// local engine to re-run with the balance assumed sufficient.
func (c *SchematicClient) checkFailureResult(flagKey string, o *checkOptions, reason string) *CheckResult {
	if !o.failOpen {
		return &CheckResult{FlagKey: flagKey, Reason: reason, Error: reason}
	}

	value := c.resolveCheckDefault(flagKey, o)
	return &CheckResult{
		Allowed: value,
		Value:   value,
		FlagKey: flagKey,
		Reason:  reason + "_fail_open",
		Error:   reason,
	}
}

// checkErrorResult resolves a plain check that never reached a verdict, through
// the same fail-open contract the credit paths use, and carries the underlying
// error: a defaulted value on its own cannot tell a caller that nothing
// evaluated the flag.
func (c *SchematicClient) checkErrorResult(flagKey string, o *checkOptions, err error) *CheckResult {
	result := c.checkFailureResult(flagKey, o, "check_failed")
	result.Error = err.Error()
	return result
}

func (c *SchematicClient) resolveCheckDefault(flagKey string, o *checkOptions) bool {
	if o.defaultValue != nil {
		return *o.defaultValue
	}
	return c.getFlagDefault(flagKey)
}

// withTimeout applies the caller's per-check timeout, if any. The returned
// cancel is always safe to call.
func (o *checkOptions) withTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	if o.timeout == nil {
		return ctx, func() {}
	}
	return context.WithTimeout(ctx, *o.timeout)
}

// preflight builds the hypothetical the flag is evaluated against. With an
// event subtype the usage goes out as the event_usage pair so the engine matches
// it to that subtype's condition; without one it goes out as the generic usage
// knob. A usage no hold could be sized from is dropped rather than threaded,
// since the server rejects it and it would only turn a check into an error.
func (o *checkOptions) preflight() *schematicgo.PreflightRequestBody {
	if o.usage == nil || !leases.IsValidQuantity(*o.usage) {
		return nil
	}

	// A hold can be sized from a fractional usage, but the preflight's quantity
	// is an integer and the question it asks is an upper bound, so a fraction
	// rounds up: the check must not pass on less usage than the operation is
	// about to record.
	quantity := int64(math.Ceil(*o.usage))
	if o.eventSubtype != nil {
		return &schematicgo.PreflightRequestBody{
			EventUsage: &schematicgo.PreflightEventUsageRequestBody{
				EventSubtype: *o.eventSubtype,
				Quantity:     quantity,
			},
		}
	}
	return &schematicgo.PreflightRequestBody{Usage: &quantity}
}

// evalCtxWithPreflight copies the caller's evaluation context with the check's
// preflight attached, leaving the caller's own value untouched.
func evalCtxWithPreflight(evalCtx *schematicgo.CheckFlagRequestBody, o *checkOptions) *schematicgo.CheckFlagRequestBody {
	preflight := mergedPreflight(evalCtx.Preflight, o.preflight())
	if preflight == evalCtx.Preflight {
		return evalCtx
	}

	copied := *evalCtx
	copied.Preflight = preflight
	return &copied
}

// mergedPreflight combines a preflight the caller put on the evaluation context
// with the one WithUsage implies. The option owns the usage knobs, since it is
// the quantity the hold is sized from, and the caller's own credit cost rides
// along: nothing here can recompute a cost the caller priced itself.
func mergedPreflight(caller, fromUsage *schematicgo.PreflightRequestBody) *schematicgo.PreflightRequestBody {
	if fromUsage == nil {
		return caller
	}
	if caller == nil {
		return fromUsage
	}

	merged := *fromUsage
	merged.CreditCost = caller.CreditCost
	return &merged
}

// paymentRequiredMessage is the server's own explanation for a 402, when the
// body carries one.
func paymentRequiredMessage(err *schematicgo.PaymentRequiredError) string {
	if err == nil {
		return ""
	}
	if err.Body != nil && err.Body.Error != "" {
		return err.Body.Error
	}
	if err.APIError == nil {
		return ""
	}
	return err.Error()
}
