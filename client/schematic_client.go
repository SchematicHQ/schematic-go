package client

import (
	"context"
	"fmt"
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

	dataStream              *datastream.DataStreamClient
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

	var dataStream *datastream.DataStreamClient
	if options.UseDataStream {
		dataStream = datastream.NewDataStream(options.BaseURL, options.APIKey, options.DatastreamOptions)
		dataStream.Start()
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
		stopWorker:              make(chan struct{}),
		workerInterval:          5 * time.Second,
		dataStream:              dataStream,
	}

	// Start background worker which handles async error logging and event buffering
	go client.worker()

	return client
}

func (c *SchematicClient) useDataStream() bool {
	return c.dataStream != nil
}

func (c *SchematicClient) CheckFlag(ctx context.Context, evalCtx *schematicgo.CheckFlagRequestBody, flagKey string) bool {
	if c.isOffline {
		return c.getFlagDefault(flagKey)
	}

	if c.useDataStream() {
		return c.checkFlagDataStream(ctx, evalCtx, flagKey)
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

func (c *SchematicClient) checkFlagDataStream(ctx context.Context, evalCtx *schematicgo.CheckFlagRequestBody, flagKey string) bool {
	var company *rulesengine.Company
	var err error
	if evalCtx.Company != nil {
		company, err = c.dataStream.GetCompany(ctx, evalCtx.Company)
		if err != nil {
			c.logger.Printf("ERROR: Failed to get company from cache: %v", err)
			return false
		}
	}

	var user *rulesengine.User
	if evalCtx.User != nil {
		user, err = c.dataStream.GetUser(ctx, evalCtx.User)
		if err != nil {
			c.logger.Printf("ERROR: Failed to get user from cache: %v", err)
			return false
		}
	}

	// Flags should be loaded already
	// Get flag here
	flag, found := c.dataStream.GetFlag(ctx, flagKey)
	if !found {
		c.logger.Printf("ERROR: Flag %s not found", flagKey)
		return false
	}

	// Evaluate against the rules engine
	resp, err := rulesengine.CheckFlag(ctx, company, user, flag)
	if err != nil {
		c.errors <- err

		return false
	}

	return resp.Value
}

func (c *SchematicClient) Close() {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Error(context.Background(), fmt.Sprintf("Panic occurred while closing client %v", r))
		}
	}()

	close(c.stopWorker)

	if c.dataStream != nil {
		c.dataStream.Close()
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
			c.logger.Error(err.Ctx, "%v", err.Err)
		case <-c.stopWorker:
			buffer.Stop()
			return
		}
	}
}
