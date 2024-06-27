package schematic

import (
	"context"
	"sync"
	"time"

	"github.com/SchematicHQ/schematic-go/api"
)

const defaultEventBufferPeriod = 5 * time.Second
const maxEvents = 100

type eventBuffer struct {
	// error logging channel
	errors chan error

	// buffer of events
	events []*api.CreateEventRequestBody

	// API client
	eventsAPI api.EventsAPI

	// frequency to flush the buffer
	interval time.Duration

	// logger
	logger Logger

	// max number of events to store in buffer
	maxEvents int

	// mutexes for flushing and pushing to the buffer
	mutexFlush sync.Mutex
	mutexPush  sync.Mutex

	// channel to signal shutdown
	shutdown chan struct{}

	// whether to accept new events
	stopped bool
}

func newEventBuffer(
	eventsAPI api.EventsAPI,
	errors chan error,
	logger Logger,
	_period *time.Duration,
) *eventBuffer {
	period := defaultEventBufferPeriod
	if _period != nil {
		period = *_period
	}

	buffer := &eventBuffer{
		events:     []*api.CreateEventRequestBody{},
		eventsAPI:  eventsAPI,
		errors:     errors,
		interval:   period,
		logger:     logger,
		maxEvents:  maxEvents,
		mutexFlush: sync.Mutex{},
		mutexPush:  sync.Mutex{},
		shutdown:   make(chan struct{}),
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

	events := make([]api.CreateEventRequestBody, len(b.events))
	for i, event := range b.events {
		if event != nil {
			events[i] = *event
		}
	}

	req := api.CreateEventBatchRequestBody{
		Events: events,
	}

	if _, _, err := b.eventsAPI.CreateEventBatch(context.Background()).CreateEventBatchRequestBody(req).Execute(); err != nil {
		b.errors <- err
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

func (b *eventBuffer) push(event *api.CreateEventRequestBody) {
	if event == nil {
		return
	}

	if b.stopped {
		b.logger.Printf("ERROR: Event buffer is stopped, not accepting new events")
		return
	}

	b.mutexPush.Lock()
	defer b.mutexPush.Unlock()

	if len(b.events) >= b.maxEvents {
		b.flush()
	}

	b.events = append(b.events, event)
}

func (b *eventBuffer) stop() {
	defer func() {
		if r := recover(); r != nil {
			b.logger.Printf("ERROR: Panic occurred while closing client %v", r)
		}
	}()

	close(b.shutdown)
}
