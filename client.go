package schematic

import (
	"context"
	"errors"
	"net/url"
	"time"

	"github.com/SchematicHQ/schematic-go/api"
)

type client struct {
	*api.APIClient

	apiClientConfig         *api.Configuration
	errors                  chan error
	eventBufferPeriod       *time.Duration
	events                  chan *api.CreateEventRequestBody
	flagCheckCacheProviders []CacheProvider[bool]
	flagDefaults            map[string]bool
	logger                  Logger
	stopWorker              chan struct{}
	workerInterval          time.Duration
}

func NewClient(apiKey string, opts ...ClientOpt) Client {
	apiConfig := api.NewConfiguration()
	apiConfig.AddDefaultHeader("X-Schematic-Api-Key", apiKey)

	client := &client{
		apiClientConfig:         apiConfig,
		flagCheckCacheProviders: make([]CacheProvider[bool], 0),
		errors:                  make(chan error, 100),
		events:                  make(chan *api.CreateEventRequestBody, 100),
		flagDefaults:            make(map[string]bool),
		logger:                  newDefaultLogger(),
		stopWorker:              make(chan struct{}),
		workerInterval:          5 * time.Second,
	}

	// Apply initialization options
	ctx := context.Background()
	for _, opt := range opts {
		if err := opt.Apply(ctx, client); err != nil {
			client.errors <- err
		}
	}

	// Once initialization options are applied, we can create the API client
	client.APIClient = api.NewAPIClient(client.apiClientConfig)

	// If no caching behavior is specified, assume a default behavior
	if len(client.flagCheckCacheProviders) == 0 {
		client.flagCheckCacheProviders = append(client.flagCheckCacheProviders, newDefaultCache[bool]())
	}

	// Start background worker which handles async error logging and event buffering
	go client.worker()

	return client
}

func (c *client) AddDefaultHeaders(ctx context.Context, headers map[string]string) {
	for key, value := range headers {
		c.apiClientConfig.AddDefaultHeader(key, value)
	}

	// In case this is called after initialization, recreate the API client
	c.APIClient = api.NewAPIClient(c.apiClientConfig)
}

func (c *client) AddFlagCheckCacheProvider(ctx context.Context, provider CacheProvider[bool]) {
	c.flagCheckCacheProviders = append(c.flagCheckCacheProviders, provider)
}

func (c *client) API() *api.APIClient {
	return c.APIClient
}

func (c *client) CheckFlag(ctx context.Context, evalCtx *CheckFlagRequestBody, flagKey string) bool {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Printf("ERROR: Panic occurred while checking flag %v", r)
		}
	}()

	// If nil, check flag with empty context
	if evalCtx == nil {
		evalCtx = &CheckFlagRequestBody{}
	}

	cacheKey := flagCheckCacheKey(evalCtx, flagKey)
	for _, provider := range c.flagCheckCacheProviders {
		if value, ok := provider.Get(ctx, cacheKey); ok {
			return value
		}
	}

	resp, _, err := c.API().FeaturesAPI.CheckFlag(ctx, flagKey).CheckFlagRequestBody(*evalCtx).Execute()
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

func (c *client) Close() {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Printf("ERROR: Panic occurred while closing client %v", r)
		}
	}()

	close(c.stopWorker)
}

func (c *client) Identify(
	ctx context.Context,
	body *EventBodyIdentify,
) {
	if err := c.enqueueEvent("identify", api.EventBodyIdentifyAsEventBody(body)); err != nil {
		c.errors <- err
	}
}

func (c *client) SetAPIClient(apiClient *api.APIClient) {
	c.APIClient = apiClient
}

func (c *client) SetAPIHost(ctx context.Context, host string) {
	// Host can be provided with or without scheme; if it's provided without scheme, assume HTTPS
	if u, err := url.Parse(host); err == nil && u.Scheme != "" && u.Host != "" {
		c.apiClientConfig.Scheme = u.Scheme
		c.apiClientConfig.Host = u.Host
	} else if err != nil {
		c.apiClientConfig.Host = host
	} else {
		c.logger.Printf("ERROR: Invalid host URL: %s", host)
		return
	}

	// In case this is called after initialization, recreate the API client
	c.APIClient = api.NewAPIClient(c.apiClientConfig)
}

func (c *client) SetEventBufferPeriod(period time.Duration) {
	c.eventBufferPeriod = &period
}

func (c *client) SetFlagDefault(
	flag string,
	value bool,
) {
	c.flagDefaults[flag] = value
}

func (c *client) SetFlagDefaults(
	values map[string]bool,
) {
	for flag, value := range values {
		c.SetFlagDefault(flag, value)
	}
}

func (c *client) Track(
	ctx context.Context,
	body *EventBodyTrack,
) {
	if err := c.enqueueEvent("track", api.EventBodyTrackAsEventBody(body)); err != nil {
		c.errors <- err
	}
}

func (c *client) enqueueEvent(
	eventType string,
	body api.EventBody,
) error {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Printf("ERROR: Panic occurred while enqueueing event %v", r)
		}
	}()

	eventBody := api.NewCreateEventRequestBody(eventType)
	if eventBody == nil {
		return errors.New("schematic event body is unexpectedly nil")
	}

	eventBody.SetBody(body)
	eventBody.SetSentAt(time.Now().UTC())

	c.events <- eventBody

	return nil
}

func (c *client) getFlagDefault(
	flag string,
) bool {
	if value, ok := c.flagDefaults[flag]; ok {
		return value
	}

	return false
}

func (c *client) worker() {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Printf("ERROR: Panic occurred in worker %v", r)
		}
	}()

	buffer := newEventBuffer(
		c.API().EventsAPI,
		c.errors,
		c.logger,
		c.eventBufferPeriod,
	)

	for {
		select {
		case event := <-c.events:
			buffer.push(event)
		case err := <-c.errors:
			c.logger.Printf("ERROR: %v", err)
		case <-c.stopWorker:
			buffer.stop()
			return
		}
	}
}
