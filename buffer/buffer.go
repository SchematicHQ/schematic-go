package buffer

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/core"
	"github.com/schematichq/schematic-go/events"
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

	// frequency to flush the buffer
	interval time.Duration

	// logger
	logger core.Logger

	// max number of events to store in buffer
	maxEvents int

	// retry configuration
	maxRetries        int
	initialRetryDelay time.Duration

	// mutexes for flushing and pushing to the buffer
	mutexFlush sync.Mutex
	mutexPush  sync.Mutex

	// channel to signal shutdown
	shutdown chan struct{}

	// whether to accept new events
	stopped bool
}

func NewEventBuffer(
	client *events.Client,
	errors chan error,
	logger core.Logger,
	_period *time.Duration,
) *eventBuffer {
	period := defaultEventBufferPeriod
	if _period != nil {
		period = *_period
	}

	buffer := &eventBuffer{
		events:            []*schematicgo.CreateEventRequestBody{},
		eventClient:       client,
		errors:            errors,
		interval:          period,
		logger:            logger,
		maxEvents:         maxEvents,
		maxRetries:        defaultMaxRetries,
		initialRetryDelay: defaultInitialRetryDelay,
		mutexFlush:        sync.Mutex{},
		mutexPush:         sync.Mutex{},
		shutdown:          make(chan struct{}),
	}

	// Start ticker to flush events periodically
	go buffer.periodicFlush()

	return buffer
}

func (b *eventBuffer) flush() {
	b.mutexFlush.Lock()
	defer b.mutexFlush.Unlock()

	if len(b.events) == 0 {
		return
	}

	events := make([]*schematicgo.CreateEventRequestBody, len(b.events))
	for i, event := range b.events {
		if event != nil {
			events[i] = event
		}
	}

	req := &schematicgo.CreateEventBatchRequestBody{
		Events: events,
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

		// Attempt to send events
		_, err := b.eventClient.CreateEventBatch(context.Background(), req)
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

	b.events = b.events[:0]
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

	b.mutexPush.Lock()
	defer b.mutexPush.Unlock()

	if len(b.events) >= b.maxEvents {
		b.flush()
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
