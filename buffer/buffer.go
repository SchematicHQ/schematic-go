package buffer

import (
	"context"
	"sync"
	"time"

	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/core"
)

const defaultEventBufferPeriod = 5 * time.Second
const maxEvents = 100

type eventBuffer struct {
	// error logging channel
	errors chan error

	// batcher handles accumulating events
	batcher *Batcher

	// sender handles sending events with retries
	sender EventSender

	// frequency to flush the buffer
	interval time.Duration

	// logger
	logger core.Logger

	// mutex for buffer operations
	mutex sync.Mutex

	// channel to signal shutdown
	shutdown chan struct{}

	// whether to accept new events
	stopped bool
}

func NewEventBuffer(
	options *core.RequestOptions,
	client core.HTTPClient,
	errors chan error,
	logger core.Logger,
	_period *time.Duration,
) *eventBuffer {
	period := defaultEventBufferPeriod
	if _period != nil {
		period = *_period
	}

	// Create HTTP sender with built-in retry logic
	sender := NewHTTPEventSender(client, options, logger)

	buffer := &eventBuffer{
		batcher:  NewBatcher(maxEvents),
		sender:   sender,
		errors:   errors,
		interval: period,
		logger:   logger,
		mutex:    sync.Mutex{},
		shutdown: make(chan struct{}),
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
	events := b.batcher.Flush()
	if events == nil {
		return
	}

	// Send with retry logic
	err := b.sender.SendBatch(context.Background(), events)
	if err != nil {
		b.errors <- err
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

	// Add returns true if batch is full
	if b.batcher.Add(event) {
		b.flushLocked()
	}
}

func (b *eventBuffer) Stop() {
	defer func() {
		if r := recover(); r != nil {
			b.logger.Error(context.Background(), "Panic occurred while closing client %v", r)
		}
	}()

	close(b.shutdown)
}