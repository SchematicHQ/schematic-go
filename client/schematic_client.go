package client

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/schematichq/rulesengine"
	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/buffer"
	"github.com/schematichq/schematic-go/cache"
	core "github.com/schematichq/schematic-go/core"
	"github.com/schematichq/schematic-go/datastream"
	"github.com/schematichq/schematic-go/flags"
	"github.com/schematichq/schematic-go/logger"
	option "github.com/schematichq/schematic-go/option"
)

type SchematicClient struct {
	*Client

	datastreamClient        *datastream.DataStreamClient
	errors                  chan error
	ctxErrors               chan *core.CtxError
	eventBufferPeriod       *time.Duration
	events                  chan *schematicgo.CreateEventRequestBody
	flagCheckCacheProviders []schematicgo.BoolCacheProvider
	flagDefaults            map[string]bool
	isOffline               bool
	logger                  core.Logger
	options                 *core.RequestOptions
	stopWorker              chan struct{}
	workerInterval          time.Duration
}

// CheckFlagResponse contains the result of a flag check with entitlement information.
// Note: FeatureAllocation and FeatureUsage* fields are deprecated and not included.
// Use Entitlement field for allocation and usage information.
type CheckFlagResponse struct {
	CompanyID   *string                         `json:"company_id,omitempty"`
	Entitlement *rulesengine.FeatureEntitlement `json:"entitlement,omitempty"`
	FlagID      *string                         `json:"flag_id,omitempty"`
	FlagKey     string                          `json:"flag_key"`
	Reason      string                          `json:"reason"`
	RuleID      *string                         `json:"rule_id,omitempty"`
	RuleType    *rulesengine.RuleType           `json:"rule_type,omitempty"`
	UserID      *string                         `json:"user_id,omitempty"`
	Value       bool                            `json:"value"`
}

func NewSchematicClient(opts ...option.RequestOption) *SchematicClient {
	options := core.NewRequestOptions(opts...)
	if options.APIKey == "" {
		// If no API key provided, assume offline mode
		opts = append(opts, core.WithOfflineMode())
	}

	// If no caching behavior is specified, assume a default behavior
	if len(options.FlagCheckCacheProviders) == 0 {
		opts = append(opts, core.WithFlagCheckCacheProvider(cache.NewDefaultCache[bool]()))
	}

	if options.Logger == nil {
		opts = append(opts, core.WithLogger(logger.NewDefaultLogger()))
	}

	// Rebuild options struct in case we added any new options above
	options = core.NewRequestOptions(opts...)

	if setter, ok := options.Logger.(interface{ SetLevel(core.LogLevel) }); ok {
		setter.SetLevel(options.LogLevel)
	}

	client := &SchematicClient{
		Client:                  NewClient(opts...),
		errors:                  make(chan error, 100),
		ctxErrors:               make(chan *core.CtxError, 100),
		eventBufferPeriod:       options.EventBufferPeriod,
		events:                  make(chan *schematicgo.CreateEventRequestBody, 100),
		flagCheckCacheProviders: options.FlagCheckCacheProviders,
		flagDefaults:            options.FlagDefaults,
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

		if err := c.enqueueEvent("flag_check", body); err != nil {
			c.logger.Error(ctx, fmt.Sprintf("Failed to enqueue flag_check event: %v", err))
		}

		return checkFlagResp, nil
	}

	return c.checkFlagAPI(ctx, evalCtx, flagKey), nil
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

	cacheKey := flags.FlagCheckCacheKey(evalCtx, flagKey)
	for _, provider := range c.flagCheckCacheProviders {
		if value, ok := provider.Get(ctx, cacheKey); ok {
			return &CheckFlagResponse{
				FlagKey: flagKey,
				Value:   value,
				Reason:  "cache hit",
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

	go func() {
		for _, provider := range c.flagCheckCacheProviders {
			if err := provider.Set(ctx, cacheKey, resp.Data.Value, nil); err != nil {
				c.ctxErrors <- &core.CtxError{
					Ctx: ctx,
					Err: err,
				}
			}
		}
	}()

	return &CheckFlagResponse{
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

// toRulesEngineRuleType converts a string pointer to a rulesengine.RuleType pointer.
func toRulesEngineRuleType(s *string) *rulesengine.RuleType {
	if s == nil {
		return nil
	}
	v := rulesengine.RuleType(*s)
	return &v
}

func (c *SchematicClient) Close() {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Error(context.Background(), fmt.Sprintf("Panic occurred while closing client %v", r))
		}
	}()

	close(c.stopWorker)

	if c.datastreamClient != nil {
		c.datastreamClient.Close()
	}
}

func (c *SchematicClient) Identify(
	ctx context.Context,
	body *schematicgo.EventBodyIdentify,
) {

	eventBody := schematicgo.EventBody{
		EventBodyIdentify: body,
	}

	if err := c.enqueueEvent("identify", eventBody); err != nil {
		c.ctxErrors <- &core.CtxError{
			Ctx: ctx,
			Err: err,
		}
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
) {

	eventBody := schematicgo.EventBody{
		EventBodyTrack: body,
	}

	if err := c.enqueueEvent("track", eventBody); err != nil {
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
) error {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Error(context.Background(), fmt.Sprintf("Panic occurred while enqueueing event %v", r))
		}
	}()

	if c.isOffline {
		return nil
	}

	now := time.Now().UTC()
	eventBody := &schematicgo.CreateEventRequestBody{
		EventType: schematicgo.EventType(eventType),
		Body:      &body,
		SentAt:    &now,
	}

	c.events <- eventBody

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
		case <-c.stopWorker:
			buffer.Stop()
			return
		}
	}
}
