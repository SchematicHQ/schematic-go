package schematic

import (
	"context"
	"errors"
	"log"
	"net/url"
	"time"

	"github.com/schematichq/schematic-go/api"
)

// Alias some types from API package so that callers don't need to import the API submodule for common operations
type CheckFlagRequestBody = api.CheckFlagRequestBody
type EventBodyIdentify = api.EventBodyIdentify
type EventBodyTrack = api.EventBodyTrack

type Client interface {
	AddCacheProvider(context.Context, CacheProvider)
	API() *api.APIClient
	SetAPIHost(context.Context, string)
	Identify(context.Context, *EventBodyIdentify) error
	Track(context.Context, *EventBodyTrack) error
}

type client struct {
	*api.APIClient

	apiClientConfig *api.Configuration
	cacheProviders  []CacheProvider
	events          chan *api.CreateEventRequestBody
	stopWorker      chan struct{}
	errors          chan error
	workerInterval  time.Duration
}

func NewClient(apiKey string, opts ...ClientOpt) Client {
	ctx := context.Background()
	apiConfig := api.NewConfiguration()
	apiConfig.AddDefaultHeader("X-Schematic-Api-Key", apiKey)
	apiConfig.UserAgent = "schematic-go" // TODO: Add version

	client := &client{
		apiClientConfig: apiConfig,
		cacheProviders:  make([]CacheProvider, 0),
		errors:          make(chan error, 100),
		events:          make(chan *api.CreateEventRequestBody, 100),
		stopWorker:      make(chan struct{}),
		workerInterval:  5 * time.Second,
	}

	for _, opt := range opts {
		if err := opt.Apply(ctx, client); err != nil {
			log.Println(err) // TODO error log
		}
	}

	client.APIClient = api.NewAPIClient(client.apiClientConfig)

	if len(client.cacheProviders) == 0 {
		client.cacheProviders = append(client.cacheProviders, newDefaultCache())
	}

	go client.worker()

	return client
}

func (c *client) SetAPIHost(ctx context.Context, host string) {
	// Host can be provided with or without scheme; if it's provided without scheme, assume HTTPS
	if u, err := url.Parse(host); err == nil && u.Scheme != "" && u.Host != "" {
		c.apiClientConfig.Scheme = u.Scheme
		c.apiClientConfig.Host = u.Host
	} else if err != nil {
		c.apiClientConfig.Host = host
	} else {
		log.Println("Invalid host URL") // TODO error log
		return
	}

	// In case this is called after initialization, recreate the API client
	c.APIClient = api.NewAPIClient(c.apiClientConfig)
}

func (c *client) AddCacheProvider(ctx context.Context, provider CacheProvider) {
	c.cacheProviders = append(c.cacheProviders, provider)
}

func (c *client) API() *api.APIClient {
	return c.APIClient
}

// TODO: Fallbacks?
// TODO: Should I return error?
func (c *client) CheckFlag(ctx context.Context, evalCtx *CheckFlagRequestBody, flagKey string) (bool, error) {
	// If nil, check flag with empty context
	if evalCtx == nil {
		evalCtx = &CheckFlagRequestBody{}
	}

	// TODO: Check and update cache providers first

	resp, _, err := c.API().FeaturesAPI.CheckFlag(ctx, flagKey).CheckFlagRequestBody(*evalCtx).Execute()
	if err != nil {
		return false, err
	}

	// TODO: Update cache providers

	return resp.Data.Value, nil
}

func (c *client) Close() (err error) {
	defer func() {
		if recover() != nil {
			err = errors.New("panic occurred while closing client")
		}
	}()

	close(c.stopWorker)
	return
}

func (c *client) Identify(
	ctx context.Context,
	body *EventBodyIdentify,
) error {
	return c.enqueueEvent("identify", api.EventBodyIdentifyAsEventBody(body))
}

func (c *client) Track(
	ctx context.Context,
	body *EventBodyTrack,
) error {
	return c.enqueueEvent("track", api.EventBodyTrackAsEventBody(body))
}

func (c *client) enqueueEvent(
	eventType string,
	body api.EventBody,
) error {
	eventBody := api.NewCreateEventRequestBody(eventType)
	if eventBody == nil {
		return errors.New("schematic event body is unexpectedly nil")
	}

	eventBody.SetBody(body)
	eventBody.SetSentAt(time.Now().UTC())

	c.events <- eventBody

	return nil
}

func (c *client) worker() {
	tick := time.NewTicker(c.workerInterval)
	defer tick.Stop()

	for {
		select {
		case event := <-c.events:
			// TODO: Batch events
			if _, _, err := c.API().EventsAPI.CreateEvent(context.Background()).CreateEventRequestBody(*event).Execute(); err != nil {
				c.errors <- err
			}
		case err := <-c.errors:
			log.Println(err) // TODO error log
		case <-c.stopWorker:
			return
		}
	}
}
