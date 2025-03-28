package client

import (
	"context"
	"time"

	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/buffer"
	"github.com/schematichq/schematic-go/cache"
	core "github.com/schematichq/schematic-go/core"
	"github.com/schematichq/schematic-go/flags"
	"github.com/schematichq/schematic-go/logger"
	option "github.com/schematichq/schematic-go/option"
)

type SchematicClient struct {
	*Client

	errors                  chan error
	eventBufferPeriod       *time.Duration
	events                  chan *schematicgo.CreateEventRequestBody
	flagCheckCacheProviders []schematicgo.BoolCacheProvider
	flagDefaults            map[string]bool
	isOffline               bool
	logger                  schematicgo.Logger
	stopWorker              chan struct{}
	workerInterval          time.Duration
}

func NewSchematicClient(opts ...option.RequestOption) *SchematicClient {
	options := core.NewRequestOptions(opts...)
	if options.APIKey == "" {
		// If no API key provideed, assume offline mode if no API key provided
		opts = append(opts, core.WithOfflineMode())
	}

	// If no caching behavior is specified, assume a default behavior
	if len(options.FlagCheckCacheProviders) == 0 {
		opts = append(opts, core.WithFlagCheckCacheProvider(cache.NewDefaultCache[bool]()))
	}

	// Rebuild options struct in case we added any new options above
	options = core.NewRequestOptions(opts...)

	client := &SchematicClient{
		Client:                  NewClient(opts...),
		errors:                  make(chan error, 100),
		eventBufferPeriod:       options.EventBufferPeriod,
		events:                  make(chan *schematicgo.CreateEventRequestBody, 100),
		flagCheckCacheProviders: options.FlagCheckCacheProviders,
		flagDefaults:            options.FlagDefaults,
		logger:                  logger.NewDefaultLogger(),
		stopWorker:              make(chan struct{}),
		workerInterval:          5 * time.Second,
	}

	// Start background worker which handles async error logging and event buffering
	go client.worker()

	return client
}

func (c *SchematicClient) CheckFlag(ctx context.Context, evalCtx *schematicgo.CheckFlagRequestBody, flagKey string) bool {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Printf("ERROR: Panic occurred while checking flag %v", r)
		}
	}()

	if c.isOffline {
		return c.getFlagDefault(flagKey)
	}

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
		c.errors <- err

		return c.getFlagDefault(flagKey)
	}

	if resp == nil {
		// if the client was not initialized with an API key, we'll have a no-op here which returns an empty response
		return c.getFlagDefault(flagKey)
	}

	go func() {
		for _, provider := range c.flagCheckCacheProviders {
			if err := provider.Set(ctx, cacheKey, resp.Data.Value, nil); err != nil {
				c.errors <- err
			}
		}
	}()

	return resp.Data.Value
}

func (c *SchematicClient) Close() {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Printf("ERROR: Panic occurred while closing client %v", r)
		}
	}()

	close(c.stopWorker)
}

func (c *SchematicClient) Identify(
	ctx context.Context,
	body *schematicgo.EventBodyIdentify,
) {

	eventBody := schematicgo.EventBody{
		EventBodyIdentify: body,
	}

	if err := c.enqueueEvent("identify", eventBody); err != nil {
		c.errors <- err
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
		c.errors <- err
	}
}

func (c *SchematicClient) enqueueEvent(
	eventType string,
	body schematicgo.EventBody,
) error {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Printf("ERROR: Panic occurred while enqueueing event %v", r)
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
			c.logger.Printf("ERROR: Panic occurred in worker %v", r)
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
			c.logger.Printf("ERROR: %v", err)
		case <-c.stopWorker:
			buffer.Stop()
			return
		}
	}
}
