package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/buffer"
	"github.com/schematichq/schematic-go/cache"
	core "github.com/schematichq/schematic-go/core"
	"github.com/schematichq/schematic-go/datastream"
	"github.com/schematichq/schematic-go/flags"
	"github.com/schematichq/schematic-go/leases"
	"github.com/schematichq/schematic-go/logger"
	option "github.com/schematichq/schematic-go/option"
	"github.com/schematichq/schematic-go/rulesengine"
)

type SchematicClient struct {
	*Client

	creditLeases            *core.CreditLeaseConfig
	datastreamClient        *datastream.DataStreamClient
	errors                  chan error
	ctxErrors               chan *core.CtxError
	eventBufferPeriod       *time.Duration
	events                  chan *schematicgo.CreateEventRequestBody
	flagCheckCacheProviders []cache.CacheProvider[*core.CheckFlagResponse]
	flagDefaults            map[string]bool
	flushRequests           chan chan struct{}
	isOffline               bool
	leaseManager            *leases.LeaseManager
	leaseStore              leases.LeaseStore
	logger                  core.Logger
	options                 *core.RequestOptions
	prewarmResolveTimeout   time.Duration
	reservations            leases.ReservationStore
	serverReservationTTL    time.Duration
	stopWorker              chan struct{}
	workerInterval          time.Duration

	// background tracks the goroutines the client spawns on a caller's behalf,
	// so Close can wait them out instead of letting them call into plumbing it
	// is about to tear down.
	backgroundMu sync.Mutex
	background   sync.WaitGroup
	closed       bool
}

// CheckFlagResponse is an alias for core.CheckFlagResponse to preserve the public API.
type CheckFlagResponse = core.CheckFlagResponse

func NewSchematicClient(opts ...option.RequestOption) *SchematicClient {
	options := core.NewRequestOptions(opts...)
	if options.APIKey == "" {
		// If no API key provided, assume offline mode
		opts = append(opts, core.WithOfflineMode())
	}

	// If no caching behavior is specified, assume a default behavior
	if len(options.FlagCheckCacheProviders) == 0 {
		opts = append(opts, core.WithFlagCheckCacheProvider(cache.NewDefaultCache[*core.CheckFlagResponse]()))
	}

	// Track whether we're supplying the SDK-default logger so we only apply the
	// configured log level to our own logger. A consumer-provided logger's level
	// is the source of truth and must not be overridden.
	usingDefaultLogger := options.Logger == nil
	if usingDefaultLogger {
		opts = append(opts, core.WithLogger(logger.NewDefaultLogger()))
	}

	// Rebuild options struct in case we added any new options above
	options = core.NewRequestOptions(opts...)

	if usingDefaultLogger {
		if setter, ok := options.Logger.(interface{ SetLevel(core.LogLevel) }); ok {
			setter.SetLevel(options.LogLevel)
		}
	}

	client := &SchematicClient{
		Client:                  NewClient(opts...),
		errors:                  make(chan error, 100),
		ctxErrors:               make(chan *core.CtxError, 100),
		eventBufferPeriod:       options.EventBufferPeriod,
		events:                  make(chan *schematicgo.CreateEventRequestBody, 100),
		flagCheckCacheProviders: options.FlagCheckCacheProviders,
		flagDefaults:            options.FlagDefaults,
		flushRequests:           make(chan chan struct{}),
		isOffline:               options.OfflineMode,
		logger:                  options.Logger,
		options:                 options,
		stopWorker:              make(chan struct{}),
		workerInterval:          5 * time.Second,
	}

	// Start background worker which handles async error logging and event buffering
	go client.worker()

	if options.UseDataStream {
		datastreamOptions := datastream.DataStreamClientOptions{
			ApiKey:  options.APIKey,
			BaseURL: options.BaseURL,
			Logger:  options.Logger,
		}

		client.datastreamClient = datastream.NewDataStreamClient(datastreamOptions, options.DatastreamOptions)
		client.datastreamClient.Start()
	}

	// After the datastream client, so the resolved mode knows whether DataStream
	// is actually enabled.
	client.configureCreditLeases(options.CreditLeases)

	return client
}

func (c *SchematicClient) useDataStream() bool {
	return c.datastreamClient != nil
}

func (c *SchematicClient) CheckFlag(ctx context.Context, evalCtx *schematicgo.CheckFlagRequestBody, flagKey string) bool {
	resp, _ := c.CheckFlagWithEntitlement(ctx, evalCtx, flagKey)
	if resp == nil {
		return c.getFlagDefault(flagKey)
	}
	return resp.Value
}

// CheckFlagWithEntitlement checks a flag and returns the full response including entitlement information.
// Note: The deprecated FeatureAllocation and FeatureUsage* fields are not included in the response.
func (c *SchematicClient) CheckFlagWithEntitlement(ctx context.Context, evalCtx *schematicgo.CheckFlagRequestBody, flagKey string) (*CheckFlagResponse, error) {
	if c.isOffline {
		return &CheckFlagResponse{
			FlagKey: flagKey,
			Value:   c.getFlagDefault(flagKey),
			Reason:  "offline mode",
		}, nil
	}

	if c.useDataStream() {
		resp, err := c.datastreamClient.CheckFlag(ctx, evalCtx, flagKey)
		if err != nil {
			c.logger.Debug(ctx, fmt.Sprintf("Datastream flag check failed (%v), falling back to API", err))
			return c.checkFlagAPI(ctx, evalCtx, flagKey), nil
		}

		checkFlagResp := toCheckFlagResponse(resp)
		if checkFlagResp == nil {
			return c.checkFlagAPI(ctx, evalCtx, flagKey), nil
		}

		body := schematicgo.EventBody{
			EventBodyFlagCheck: &schematicgo.EventBodyFlagCheck{
				FlagKey:    flagKey,
				Value:      checkFlagResp.Value,
				CompanyID:  checkFlagResp.CompanyID,
				UserID:     checkFlagResp.UserID,
				FlagID:     checkFlagResp.FlagID,
				ReqCompany: evalCtx.Company,
				ReqUser:    evalCtx.User,
				RuleID:     checkFlagResp.RuleID,
				Reason:     checkFlagResp.Reason,
			},
		}

		if err := c.enqueueEvent("flag_check", body, &eventOptions{}); err != nil {
			c.logger.Error(ctx, fmt.Sprintf("Failed to enqueue flag_check event: %v", err))
		}

		return checkFlagResp, nil
	}

	return c.checkFlagAPI(ctx, evalCtx, flagKey), nil
}

// CheckFlags evaluates multiple flags for the given context.
//
// When keys is empty or nil, all flags for the context are returned via the
// bulk /flags/check endpoint (cache is bypassed for this shape).
//
// When keys are provided:
//   - In offline mode, each key resolves to its configured default.
//   - If datastream is enabled and connected, all keys are evaluated locally;
//     any failure causes a fallback to the API for the entire batch.
//   - Otherwise, cached values are returned only when every requested key is
//     present in the cache. Any miss triggers a single API call to refresh all
//     keys, guaranteeing the returned values come from a consistent evaluation.
//
// On error, defaults are returned for each requested key with Reason set.
// Results are returned in the same order as keys; entries for keys missing
// from the API response use the configured default.
func (c *SchematicClient) CheckFlags(ctx context.Context, evalCtx *schematicgo.CheckFlagRequestBody, keys []string) (results []*CheckFlagResponse) {
	results = c.flagDefaultsFor(keys, "error")
	defer func() {
		if r := recover(); r != nil {
			c.logger.Error(ctx, fmt.Sprintf("Panic occurred while checking flags %v", r))
		}
	}()

	if evalCtx == nil {
		evalCtx = &schematicgo.CheckFlagRequestBody{}
	}

	if c.isOffline {
		return c.flagDefaultsFor(keys, "offline mode")
	}

	// Datastream path: evaluate all keys locally if possible.
	if c.useDataStream() && len(keys) > 0 {
		if dsResults, ok := c.checkFlagsViaDataStream(ctx, evalCtx, keys); ok {
			return dsResults
		}
		c.logger.Debug(ctx, "Datastream check_flags failed for one or more keys, falling back to API")
	}

	return c.checkFlagsAPI(ctx, evalCtx, keys)
}

// flagDefaultsFor returns a slice of CheckFlagResponse, one per requested key,
// populated from SetFlagDefault values with the given reason. An empty keys
// slice yields an empty (non-nil) slice.
func (c *SchematicClient) flagDefaultsFor(keys []string, reason string) []*CheckFlagResponse {
	results := make([]*CheckFlagResponse, 0, len(keys))
	for _, key := range keys {
		results = append(results, &CheckFlagResponse{
			FlagKey: key,
			Value:   c.getFlagDefault(key),
			Reason:  reason,
		})
	}
	return results
}

// checkFlagsViaDataStream evaluates all keys via the local datastream cache.
// Returns (results, true) if every key was evaluated successfully; otherwise
// (nil, false) to signal the caller should fall back to the API.
func (c *SchematicClient) checkFlagsViaDataStream(ctx context.Context, evalCtx *schematicgo.CheckFlagRequestBody, keys []string) ([]*CheckFlagResponse, bool) {
	results := make([]*CheckFlagResponse, 0, len(keys))
	for _, key := range keys {
		resp, err := c.datastreamClient.CheckFlag(ctx, evalCtx, key)
		if err != nil {
			c.logger.Debug(ctx, fmt.Sprintf("Datastream flag check failed for '%s' (%v), falling back to API", key, err))
			return nil, false
		}
		checkFlagResp := toCheckFlagResponse(resp)
		if checkFlagResp == nil {
			return nil, false
		}

		results = append(results, checkFlagResp)
	}
	return results, true
}

func (c *SchematicClient) checkFlagsAPI(ctx context.Context, evalCtx *schematicgo.CheckFlagRequestBody, keys []string) []*CheckFlagResponse {
	// No keys: return the full set of flags for the context directly from the API.
	if len(keys) == 0 {
		resp, err := c.Features.CheckFlags(ctx, evalCtx)
		if err != nil {
			c.ctxErrors <- &core.CtxError{Ctx: ctx, Err: err}
			return []*CheckFlagResponse{}
		}
		if resp == nil || resp.Data == nil {
			return []*CheckFlagResponse{}
		}
		results := make([]*CheckFlagResponse, 0, len(resp.Data.Flags))
		for _, f := range resp.Data.Flags {
			results = append(results, flagResponseDataToCheckFlagResponse(f))
		}
		return results
	}

	// A preflighted check asks a hypothetical, and the cache key is only
	// (flag, company, user), so it neither reads nor writes the cache. See
	// checkFlagAPI.
	useCache := evalCtx.Preflight == nil

	// Check cache for all requested keys; return cached results only if every key hits.
	cachedResults := make(map[string]*CheckFlagResponse, len(keys))
	allCached := useCache
	if useCache {
		for _, key := range keys {
			cacheKey := flags.FlagCheckCacheKey(evalCtx, key)
			found := false
			for _, provider := range c.flagCheckCacheProviders {
				if cached, ok := provider.Get(ctx, cacheKey); ok && cached != nil {
					cachedCopy := *cached
					cachedResults[key] = &cachedCopy
					found = true
					break
				}
			}
			if !found {
				allCached = false
			}
		}
	}

	if allCached {
		results := make([]*CheckFlagResponse, 0, len(keys))
		for _, key := range keys {
			results = append(results, cachedResults[key])
		}
		return results
	}

	// Any cache miss — refresh all keys from the API for a consistent snapshot.
	resp, err := c.Features.CheckFlags(ctx, evalCtx)
	if err != nil {
		c.ctxErrors <- &core.CtxError{Ctx: ctx, Err: err}
		return c.flagDefaultsFor(keys, "error")
	}

	apiResults := make(map[string]*CheckFlagResponse)
	if resp != nil && resp.Data != nil {
		for _, f := range resp.Data.Flags {
			apiResults[f.Flag] = flagResponseDataToCheckFlagResponse(f)
		}
	}

	// Cache all fresh values asynchronously; keyed per flag+context.
	if useCache {
		toCache := make(map[string]*CheckFlagResponse, len(apiResults))
		for flagKey, result := range apiResults {
			cacheKey := flags.FlagCheckCacheKey(evalCtx, flagKey)
			resultCopy := *result
			toCache[cacheKey] = &resultCopy
		}
		go func() {
			for cacheKey, value := range toCache {
				for _, provider := range c.flagCheckCacheProviders {
					if err := provider.Set(ctx, cacheKey, value, nil); err != nil {
						c.ctxErrors <- &core.CtxError{Ctx: ctx, Err: err}
					}
				}
			}
		}()
	}

	results := make([]*CheckFlagResponse, 0, len(keys))
	for _, key := range keys {
		if api, ok := apiResults[key]; ok {
			results = append(results, api)
			continue
		}
		results = append(results, &CheckFlagResponse{
			FlagKey: key,
			Value:   c.getFlagDefault(key),
			Reason:  "flag not found",
		})
	}
	return results
}

// flagResponseDataToCheckFlagResponse adapts the bulk endpoint's flag entry to
// the Go SDK's response type. Deprecated FeatureAllocation/FeatureUsage* fields
// on the source are intentionally dropped.
func flagResponseDataToCheckFlagResponse(f *schematicgo.CheckFlagResponseData) *CheckFlagResponse {
	if f == nil {
		return nil
	}
	return &CheckFlagResponse{
		CompanyID:   f.CompanyID,
		Entitlement: toRulesEngineEntitlement(f.Entitlement),
		FlagID:      f.FlagID,
		FlagKey:     f.Flag,
		Reason:      f.Reason,
		RuleID:      f.RuleID,
		RuleType:    toRulesEngineRuleType(f.RuleType),
		UserID:      f.UserID,
		Value:       f.Value,
	}
}

func (c *SchematicClient) checkFlagAPI(ctx context.Context, evalCtx *schematicgo.CheckFlagRequestBody, flagKey string) (result *CheckFlagResponse) {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Error(ctx, "Panic occurred while checking flag %v", r)
			result = &CheckFlagResponse{
				FlagKey: flagKey,
				Value:   c.getFlagDefault(flagKey),
				Reason:  "error",
			}
		}
	}()

	// If nil, check flag with empty context
	if evalCtx == nil {
		evalCtx = &schematicgo.CheckFlagRequestBody{}
	}

	// A preflighted check asks a hypothetical ("would this action be allowed?"),
	// and the cache key is only (flag, company, user). Reading the cache would
	// answer the hypothetical with the plain verdict, and writing it would serve
	// the hypothetical to every later plain check, so a preflight skips both.
	useCache := evalCtx.Preflight == nil

	cacheKey := flags.FlagCheckCacheKey(evalCtx, flagKey)
	if useCache {
		for _, provider := range c.flagCheckCacheProviders {
			if cached, ok := provider.Get(ctx, cacheKey); ok && cached != nil {
				result := *cached
				return &result
			}
		}
	}

	resp, err := c.Features.CheckFlag(ctx, flagKey, evalCtx)
	if err != nil {
		c.ctxErrors <- &core.CtxError{
			Ctx: ctx,
			Err: err,
		}

		return &CheckFlagResponse{
			FlagKey: flagKey,
			Value:   c.getFlagDefault(flagKey),
			Reason:  "error",
		}
	}

	if resp == nil {
		// if the client was not initialized with an API key, we'll have a no-op here which returns an empty response
		return &CheckFlagResponse{
			FlagKey: flagKey,
			Value:   c.getFlagDefault(flagKey),
			Reason:  "no response",
		}
	}

	checkFlagResp := &CheckFlagResponse{
		CompanyID:   resp.Data.CompanyID,
		Entitlement: toRulesEngineEntitlement(resp.Data.Entitlement),
		FlagID:      resp.Data.FlagID,
		FlagKey:     flagKey,
		Reason:      resp.Data.Reason,
		RuleID:      resp.Data.RuleID,
		RuleType:    toRulesEngineRuleType(resp.Data.RuleType),
		UserID:      resp.Data.UserID,
		Value:       resp.Data.Value,
	}

	if useCache {
		cachedCopy := *checkFlagResp
		go func() {
			for _, provider := range c.flagCheckCacheProviders {
				if err := provider.Set(ctx, cacheKey, &cachedCopy, nil); err != nil {
					c.ctxErrors <- &core.CtxError{
						Ctx: ctx,
						Err: err,
					}
				}
			}
		}()
	}

	return checkFlagResp
}

// toCheckFlagResponse converts a rulesengine.CheckFlagResult to CheckFlagResponse,
// excluding deprecated FeatureAllocation and FeatureUsage* fields.
func toCheckFlagResponse(result *rulesengine.CheckFlagResult) *CheckFlagResponse {
	if result == nil {
		return nil
	}

	return &CheckFlagResponse{
		CompanyID:   result.CompanyID,
		Entitlement: result.Entitlement,
		FlagID:      result.FlagID,
		FlagKey:     result.FlagKey,
		Reason:      result.Reason,
		RuleID:      result.RuleID,
		RuleType:    result.RuleType,
		UserID:      result.UserID,
		Value:       result.Value,
	}
}

// toRulesEngineEntitlement converts a schematicgo.FeatureEntitlement to a rulesengine.FeatureEntitlement.
func toRulesEngineEntitlement(e *schematicgo.FeatureEntitlement) *rulesengine.FeatureEntitlement {
	if e == nil {
		return nil
	}

	result := &rulesengine.FeatureEntitlement{
		FeatureID:       e.FeatureID,
		FeatureKey:      e.FeatureKey,
		ValueType:       rulesengine.EntitlementValueType(e.ValueType),
		EventName:       e.EventName,
		MetricResetAt:   e.MetricResetAt,
		CreditID:        e.CreditID,
		CreditTotal:     e.CreditTotal,
		CreditUsed:      e.CreditUsed,
		CreditRemaining: e.CreditRemaining,
	}

	if e.Allocation != nil {
		v := int64(*e.Allocation)
		result.Allocation = &v
	}
	if e.SoftLimit != nil {
		v := int64(*e.SoftLimit)
		result.SoftLimit = &v
	}
	if e.Usage != nil {
		v := int64(*e.Usage)
		result.Usage = &v
	}
	if e.MetricPeriod != nil {
		v := rulesengine.MetricPeriod(string(*e.MetricPeriod))
		result.MetricPeriod = &v
	}
	if e.MonthReset != nil {
		v := rulesengine.MetricPeriodMonthReset(string(*e.MonthReset))
		result.MonthReset = &v
	}

	return result
}

// toRulesEngineRuleType converts a schematicgo.RuleType pointer to a rulesengine.RuleType pointer.
func toRulesEngineRuleType(s *schematicgo.RuleType) *rulesengine.RuleType {
	if s == nil {
		return nil
	}
	v := rulesengine.RuleType(string(*s))
	return &v
}

func (c *SchematicClient) Close() {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Error(context.Background(), fmt.Sprintf("Panic occurred while closing client %v", r))
		}
	}()

	// One budget for the whole teardown, rather than each step's timeout in
	// turn: a caller closing a client wants a bounded wait, not the sum of
	// every wait inside it.
	ctx, cancel := context.WithTimeout(context.Background(), closeTimeout)
	defer cancel()

	// Before the lease plumbing goes away, since these goroutines call into it.
	c.stopBackground(ctx)

	if c.leaseManager != nil {
		c.leaseManager.StopWithContext(ctx)
		// Releasing hands the unspent remainder of this process's leases back to
		// the company balance now instead of at expiry. The manager skips a
		// shared store, whose leases sibling processes still draw on.
		c.leaseManager.ReleaseAllLocalLeases(ctx)
	}

	close(c.stopWorker)

	if c.datastreamClient != nil {
		c.datastreamClient.Close()
	}
}

// spawnBackground runs work on the caller's behalf, tracked so Close can wait
// it out. It refuses once the client is closing, since the plumbing that work
// would touch is being torn down.
func (c *SchematicClient) spawnBackground(fn func()) bool {
	c.backgroundMu.Lock()
	if c.closed {
		c.backgroundMu.Unlock()
		return false
	}
	c.background.Add(1)
	c.backgroundMu.Unlock()

	go func() {
		defer c.background.Done()
		defer func() {
			if r := recover(); r != nil {
				c.logger.Error(context.Background(), fmt.Sprintf("Panic occurred in background client work %v", r))
			}
		}()
		fn()
	}()
	return true
}

// stopBackground refuses further background work and waits out what is already
// running, bounded by the caller's close budget so a stalled call cannot hold up
// a shutdown.
func (c *SchematicClient) stopBackground(ctx context.Context) {
	c.backgroundMu.Lock()
	c.closed = true
	c.backgroundMu.Unlock()

	waited := make(chan struct{})
	go func() {
		c.background.Wait()
		close(waited)
	}()
	select {
	case <-waited:
	case <-ctx.Done():
		c.logger.Warn(ctx, "Gave up waiting for background client work to finish; closing anyway")
	}
}

// flushEvents sends everything buffered now rather than at the next tick. The
// buffer belongs to the worker goroutine, so the request goes through it and
// this waits for the worker's acknowledgement.
func (c *SchematicClient) flushEvents(ctx context.Context) error {
	ack := make(chan struct{})
	select {
	case c.flushRequests <- ack:
	case <-c.stopWorker:
		return errors.New("the event worker has stopped")
	case <-ctx.Done():
		return ctx.Err()
	}

	select {
	case <-ack:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (c *SchematicClient) Identify(
	ctx context.Context,
	body *schematicgo.EventBodyIdentify,
	opts ...IdentifyOption,
) {

	eventBody := schematicgo.EventBody{
		EventBodyIdentify: body,
	}

	o := &eventOptions{}
	for _, apply := range opts {
		apply(o)
	}

	if err := c.enqueueEvent("identify", eventBody, o); err != nil {
		c.ctxErrors <- &core.CtxError{
			Ctx: ctx,
			Err: err,
		}
	}

	if len(o.prewarm) == 0 {
		return
	}
	if body == nil || body.Company == nil || len(body.Company.Keys) == 0 {
		c.logger.Debug(ctx, "Identify: a prewarm needs company keys on the identify event")
		return
	}
	// Detached, so a caller that cancels the request context still gets its
	// leases warmed, and unawaited, so the identify does not wait on the wire.
	evalCtx := &schematicgo.CheckFlagRequestBody{Company: body.Company.Keys, User: body.Keys}
	prewarmCtx := context.WithoutCancel(ctx)
	spawned := c.spawnBackground(func() {
		// The buffer otherwise holds the identify for up to its flush period,
		// which outlasts the prewarm's own wait for the company to surface, so
		// the prewarm would be polling for an entity the server has not been
		// told about yet. Only a client-mode prewarm does that polling, and the
		// wait is capped: an identify the API is slow to take must not hold a
		// lease warm-up, or a Close behind it, open indefinitely.
		if c.leaseManager != nil {
			flushCtx, cancelFlush := context.WithTimeout(prewarmCtx, identifyFlushTimeout)
			if err := c.flushEvents(flushCtx); err != nil {
				c.logger.Debug(prewarmCtx, fmt.Sprintf("Identify: flushing before the prewarm failed: %v", err))
			}
			cancelFlush()
		}
		if err := c.Prewarm(prewarmCtx, evalCtx, o.prewarm); err != nil {
			c.logger.Warn(prewarmCtx, fmt.Sprintf("Identify: prewarm failed: %v", err))
		}
	})
	if !spawned {
		c.logger.Debug(ctx, "Identify: the client is closing, so the prewarm was skipped")
	}
}

func (c *SchematicClient) SetFlagDefault(
	flag string,
	value bool,
) {
	if c.flagDefaults == nil {
		c.flagDefaults = make(map[string]bool)
	}
	c.flagDefaults[flag] = value
}

func (c *SchematicClient) SetFlagDefaults(
	values map[string]bool,
) {
	for flag, value := range values {
		c.SetFlagDefault(flag, value)
	}
}

func (c *SchematicClient) Track(
	ctx context.Context,
	body *schematicgo.EventBodyTrack,
	opts ...TrackOption,
) {

	o := &eventOptions{}
	for _, apply := range opts {
		apply(o)
	}

	body = bodyWithTraits(body, o.traits)
	eventBody := schematicgo.EventBody{
		EventBodyTrack: body,
	}

	if err := c.enqueueEvent("track", eventBody, o); err != nil {
		c.ctxErrors <- &core.CtxError{
			Ctx: ctx,
			Err: err,
		}
	}

	if body.Company != nil && c.useDataStream() && c.datastreamClient.IsConnected() {
		err := c.datastreamClient.UpdateCompanyMetrics(ctx, body)
		if err != nil {
			c.ctxErrors <- &core.CtxError{
				Ctx: ctx,
				Err: fmt.Errorf("failed to update company metrics: %w", err),
			}
		}
	}
}

func (c *SchematicClient) enqueueEvent(
	eventType string,
	body schematicgo.EventBody,
	opts *eventOptions,
) error {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Error(context.Background(), fmt.Sprintf("Panic occurred while enqueueing event %v", r))
		}
	}()

	if c.isOffline {
		return nil
	}

	sentAt := opts.sentAt
	if sentAt == nil {
		now := time.Now().UTC()
		sentAt = &now
	}

	c.events <- &schematicgo.CreateEventRequestBody{
		EventType:          schematicgo.EventType(eventType),
		Body:               &body,
		IdempotencyKey:     opts.idempotencyKey,
		SentAt:             sentAt,
		TrustedClientClock: opts.trustedClientClock,
		Backfill:           opts.backfill,
	}

	return nil
}

func (c *SchematicClient) getFlagDefault(
	flag string,
) bool {
	if c.flagDefaults == nil {
		return false
	}

	if value, ok := c.flagDefaults[flag]; ok {
		return value
	}

	return false
}

func (c *SchematicClient) worker() {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Error(context.Background(), fmt.Sprintf("Panic occurred in worker %v", r))
		}
	}()

	// Start buffered event worker
	httpClient := c.options.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	buffer := buffer.NewEventBuffer(
		c.options,
		httpClient,
		c.errors,
		c.logger,
		c.eventBufferPeriod,
	)

	for {
		select {
		case event := <-c.events:
			buffer.Push(event)
		case err := <-c.errors:
			c.logger.Error(context.Background(), fmt.Sprintf("%v", err))
		case err := <-c.ctxErrors:
			c.logger.Error(err.Ctx, fmt.Sprintf("%v", err.Err))
		case ack := <-c.flushRequests:
			// Enqueueing an event only hands it to this loop, so drain what is
			// already queued first: a caller that enqueued and then asked for a
			// flush means the flush to cover that event.
			for drained := true; drained; {
				select {
				case event := <-c.events:
					buffer.Push(event)
				default:
					drained = false
				}
			}
			// The send itself runs elsewhere. It retries for seconds, and this
			// loop is the only reader of c.events, so a send here would back
			// every enqueueEvent up behind the API, a client-mode check's
			// flag_check included. The acknowledgement still waits on the send
			// having happened, since that is what the caller asked for.
			sent := buffer.FlushAsync()
			go func() {
				<-sent
				close(ack)
			}()
		case <-c.stopWorker:
			buffer.Stop()
			return
		}
	}
}
