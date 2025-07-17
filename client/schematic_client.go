package client

import (
	"context"
	"fmt"
	"time"

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
	datastreamConnected     bool
	datastreamConnection    chan bool
	errors                  chan error
	ctxErrors               chan *core.CtxError
	eventBufferPeriod       *time.Duration
	events                  chan *schematicgo.CreateEventRequestBody
	flagCheckCacheProviders []schematicgo.BoolCacheProvider
	flagDefaults            map[string]bool
	isOffline               bool
	logger                  core.Logger
	stopWorker              chan struct{}
	workerInterval          time.Duration
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

	datastreamConnection := make(chan bool, 1)

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
		stopWorker:              make(chan struct{}),
		workerInterval:          5 * time.Second,
		datastreamConnection:    datastreamConnection,
	}

	// Start background worker which handles async error logging and event buffering
	go client.worker()

	if options.UseDataStream {
		go client.monitorDatastreamConnection()

		datastreamOptions := datastream.DataStreamClientOptions{
			ApiKey:         options.APIKey,
			BaseURL:        options.BaseURL,
			MonitorChannel: datastreamConnection,
			Logger:         options.Logger,
		}

		client.datastreamClient = datastream.NewDataStreamClient(datastreamOptions, options.DatastreamOptions)
		client.datastreamClient.Start()
	}

	return client
}

func (c *SchematicClient) useDataStream() bool {
	return c.datastreamConnected
}

func (c *SchematicClient) monitorDatastreamConnection() {
	for {
		select {
		case connected := <-c.datastreamConnection:
			c.datastreamConnected = connected
		}
	}
}

func (c *SchematicClient) CheckFlag(ctx context.Context, evalCtx *schematicgo.CheckFlagRequestBody, flagKey string) bool {
	if c.isOffline {
		return c.getFlagDefault(flagKey)
	}

	if c.useDataStream() {
		resp := c.datastreamClient.CheckFlag(ctx, evalCtx, flagKey)
		if resp == nil {
			return false
		}

		body := schematicgo.EventBody{
			EventBodyFlagCheck: &schematicgo.EventBodyFlagCheck{
				FlagKey:    flagKey,
				Value:      resp.Value,
				CompanyID:  resp.CompanyID,
				UserID:     resp.UserID,
				FlagID:     resp.FlagID,
				ReqCompany: evalCtx.Company,
				ReqUser:    evalCtx.User,
				RuleID:     resp.RuleID,
				Reason:     resp.Reason,
			},
		}

		c.enqueueEvent("flag_check", body)

		return resp.Value
	}
	return c.checkFlag(ctx, evalCtx, flagKey)
}

func (c *SchematicClient) checkFlag(ctx context.Context, evalCtx *schematicgo.CheckFlagRequestBody, flagKey string) bool {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Error(ctx, "Panic occurred while checking flag %v", r)
		}
	}()

	// If nil, check flag with empty context
	if evalCtx == nil {
		evalCtx = &schematicgo.CheckFlagRequestBody{}
	}

	cacheKey := flags.FlagCheckCacheKey(evalCtx, flagKey)
	for _, provider := range c.flagCheckCacheProviders {
		if value, ok := provider.Get(ctx, cacheKey); ok {
			return value
		}
	}

	resp, err := c.Features.CheckFlag(ctx, flagKey, evalCtx)
	if err != nil {
		c.ctxErrors <- &core.CtxError{
			Ctx: ctx,
			Err: err,
		}

		return c.getFlagDefault(flagKey)
	}

	if resp == nil {
		// if the client was not initialized with an API key, we'll have a no-op here which returns an empty response
		return c.getFlagDefault(flagKey)
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

	return resp.Data.Value
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

	if body.Company != nil && c.datastreamConnected {
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
		EventType: schematicgo.CreateEventRequestBodyEventType(eventType),
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

	buffer := buffer.NewEventBuffer(
		c.Events,
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
