package buffer

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"sync"
	"time"

	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/core"
	"github.com/schematichq/schematic-go/events"
	"github.com/schematichq/schematic-go/internal"
)

const defaultEventBufferPeriod = 5 * time.Second
const maxEvents = 100
const defaultMaxRetries = 3
const defaultInitialRetryDelay = 1 * time.Second

type eventBuffer struct {
	// error logging channel
	errors chan error

	// buffer of events
	events []*schematicgo.CreateEventRequestBody

	// API client
	eventClient *events.Client

	// request options (includes API key, base URL, etc.)
	options *core.RequestOptions

	// caller for making HTTP requests with retry logic
	caller *internal.Caller

	// frequency to flush the buffer
	interval time.Duration

	// logger
	logger core.Logger

	// max number of events to store in buffer
	maxEvents int

	// retry configuration
	maxRetries        int
	initialRetryDelay time.Duration

	// mutex for buffer operations
	mutex sync.Mutex

	// channel to signal shutdown
	shutdown chan struct{}

	// whether to accept new events
	stopped bool
}

func NewEventBuffer(
	client *events.Client,
	options *core.RequestOptions,
	caller *internal.Caller,
	errors chan error,
	logger core.Logger,
	_period *time.Duration,
) *eventBuffer {
	period := defaultEventBufferPeriod
	if _period != nil {
		period = *_period
	}

	// Set default CaptureURL if not explicitly provided
	options.CaptureURL = internal.ResolveBaseURL(
		options.CaptureURL,
		"https://capture.schematichq.com",
	)

	buffer := &eventBuffer{
		events:            []*schematicgo.CreateEventRequestBody{},
		eventClient:       client,
		options:           options,
		caller:            caller,
		errors:            errors,
		interval:          period,
		logger:            logger,
		maxEvents:         maxEvents,
		maxRetries:        defaultMaxRetries,
		initialRetryDelay: defaultInitialRetryDelay,
		mutex:             sync.Mutex{},
		shutdown:          make(chan struct{}),
	}

	// Start ticker to flush events periodically
	go buffer.periodicFlush()

	return buffer
}

func (b *eventBuffer) flush() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.flushLocked()
}

func (b *eventBuffer) flushLocked() {
	if len(b.events) == 0 {
		return
	}

	eventsToProcess := make([]*schematicgo.CreateEventRequestBody, len(b.events))
	copy(eventsToProcess, b.events)

	b.events = b.events[:0]

	events := make([]*schematicgo.CreateEventRequestBody, 0, len(eventsToProcess))
	for _, event := range eventsToProcess {
		if event != nil {
			events = append(events, event)
		}
	}

	if len(events) == 0 {
		return
	}

	// Initialize retry counter and success flag
	retryCount := 0
	success := false
	var lastErr error

	// Try with retries and exponential backoff
	for retryCount <= b.maxRetries && !success {
		if retryCount > 0 {
			// Log retry attempt
			b.logger.Info(context.Background(), fmt.Sprintf("Retrying event batch submission (attempt %d of %d)", retryCount, b.maxRetries))
		}

		// Attempt to send events to custom batch endpoint
		err := b.sendBatchEvents(events)
		if err == nil {
			success = true
		} else {
			lastErr = err
			retryCount++

			if retryCount <= b.maxRetries {
				// Calculate backoff with jitter
				delay := float64(b.initialRetryDelay.Nanoseconds()) * math.Pow(2, float64(retryCount-1))
				jitter := rand.Float64() * 0.1 * delay // 10% jitter
				waitTime := time.Duration(delay + jitter)

				b.logger.Warn(
					context.Background(),
					fmt.Sprintf("Event batch submission failed: %v. Retrying in %.2f seconds...",
						err, float64(waitTime)/float64(time.Second)),
				)

				// Wait before retry
				time.Sleep(waitTime)
			}
		}
	}

	// After all retries, if still not successful, log the error
	if !success {
		b.logger.Error(
			context.Background(),
			fmt.Sprintf("Event batch submission failed after %d retries: %v", b.maxRetries, lastErr),
		)
		b.errors <- lastErr
	} else if retryCount > 0 {
		b.logger.Info(context.Background(), fmt.Sprintf("Event batch submission succeeded after %d retries", retryCount))
	}
}

func (b *eventBuffer) periodicFlush() {
	ticker := time.NewTicker(b.interval)
	defer ticker.Stop()

	for {
		select {
		case <-b.shutdown:
			// stop accepting new events
			b.stopped = true

			// flush any remaining events
			b.flush()

			return
		case <-ticker.C:
			b.flush()
		}
	}
}

func (b *eventBuffer) Push(event *schematicgo.CreateEventRequestBody) {
	if event == nil {
		return
	}

	if b.stopped {
		b.logger.Error(context.Background(), "Event buffer is stopped, not accepting new events")
		return
	}

	b.mutex.Lock()
	defer b.mutex.Unlock()

	if len(b.events) >= b.maxEvents {
		b.flushLocked()
	}

	b.events = append(b.events, event)
}

func (b *eventBuffer) Stop() {
	defer func() {
		if r := recover(); r != nil {
			b.logger.Error(context.Background(), "Panic occurred while closing client %v", r)
		}
	}()

	close(b.shutdown)
}

// sendBatchEvents sends events to the /batch endpoint on the capture subdomain
func (b *eventBuffer) sendBatchEvents(events []*schematicgo.CreateEventRequestBody) error {
	// Use CaptureURL from options (set during initialization)
	endpoint := b.options.CaptureURL + "/batch"

	// Transform CreateEventRequestBody to EventPayload format expected by capture service
	// Each event needs: api_key, body, type, sent_at
	eventPayloads := make([]map[string]interface{}, 0, len(events))
	for _, event := range events {
		payload := map[string]interface{}{
			"api_key": b.options.APIKey,
			"body":    event.Body,
			"type":    event.EventType,
			"sent_at": event.SentAt,
		}
		eventPayloads = append(eventPayloads, payload)
	}

	// Create batch payload with transformed events
	batchPayload := map[string]interface{}{
		"events": eventPayloads,
	}

	// Use SDK utilities to build headers (includes API key and SDK tracking headers)
	headers := b.options.ToHeader()
	headers.Add("Content-Type", "application/json")

	// Use internal.Caller to make the request with built-in retry logic
	_, err := b.caller.Call(context.Background(), &internal.CallParams{
		URL:            endpoint,
		Method:         http.MethodPost,
		Headers:        headers,
		BodyProperties: batchPayload,
		MaxAttempts:    b.options.MaxAttempts,
		Response:       nil, // We don't need to parse the response
		ErrorDecoder:   nil, // Could use schematichq.ErrorCodes if capture API returns same format
	})

	if err != nil {
		return fmt.Errorf("failed to send batch events: %w", err)
	}

	return nil
}
