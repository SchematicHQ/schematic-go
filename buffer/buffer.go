package buffer

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/core"
)

const defaultEventBufferPeriod = 5 * time.Second
const maxEvents = 100

// defaultShutdownTimeout bounds how long Stop will wait for the final flush to
// complete, so a wedged HTTP client can't hang a caller's shutdown path
// indefinitely.
const defaultShutdownTimeout = 10 * time.Second

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

	// closed once the periodic flush loop has performed its final flush and exited
	done chan struct{}

	// guards against a double Stop closing the shutdown channel twice
	stopOnce sync.Once

	// whether to accept new events
	stopped atomic.Bool

	// bounds how long the final flush may take
	shutdownTimeout time.Duration
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
		batcher:         NewBatcher(maxEvents),
		sender:          sender,
		errors:          errors,
		interval:        period,
		logger:          logger,
		mutex:           sync.Mutex{},
		shutdown:        make(chan struct{}),
		done:            make(chan struct{}),
		shutdownTimeout: defaultShutdownTimeout,
	}

	// Start ticker to flush events periodically
	go buffer.periodicFlush()

	return buffer
}

// Flush sends any buffered events and blocks until the send (including retries)
// has completed.
func (b *eventBuffer) Flush(ctx context.Context) {
	b.mutex.Lock()
	events := b.batcher.Flush()
	b.mutex.Unlock()

	b.sendEvents(ctx, events)
}

// sendEvents sends events outside the lock so Push callers aren't blocked during HTTP retries.
func (b *eventBuffer) sendEvents(ctx context.Context, events []*schematicgo.CreateEventRequestBody) {
	if events == nil {
		return
	}

	err := b.sender.SendBatch(ctx, events)
	if err == nil {
		return
	}

	// Never block on the error channel: during shutdown the worker goroutine is
	// waiting on this flush and is no longer draining it, so a blocking send
	// would deadlock.
	select {
	case b.errors <- err:
	default:
		b.logger.Error(ctx, "%v", err)
	}
}

func (b *eventBuffer) periodicFlush() {
	defer close(b.done)

	ticker := time.NewTicker(b.interval)
	defer ticker.Stop()

	for {
		select {
		case <-b.shutdown:
			// Flush any remaining events, bounded so shutdown can't hang forever.
			timeout := b.shutdownTimeout
			if timeout <= 0 {
				timeout = defaultShutdownTimeout
			}
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			b.Flush(ctx)
			cancel()

			return
		case <-ticker.C:
			b.Flush(context.Background())
		}
	}
}

func (b *eventBuffer) Push(event *schematicgo.CreateEventRequestBody) {
	if event == nil {
		return
	}

	if b.stopped.Load() {
		b.logger.Error(context.Background(), "Event buffer is stopped, not accepting new events")
		return
	}

	b.mutex.Lock()
	full := b.batcher.Add(event)
	var events []*schematicgo.CreateEventRequestBody
	if full {
		events = b.batcher.Flush()
	}
	b.mutex.Unlock()

	b.sendEvents(context.Background(), events)
}

// Stop stops accepting new events and blocks until the buffered events have
// been flushed, so callers can rely on Stop returning only once in-flight
// events have been delivered (or definitively failed).
func (b *eventBuffer) Stop() {
	defer func() {
		if r := recover(); r != nil {
			b.logger.Error(context.Background(), "Panic occurred while closing client %v", r)
		}
	}()

	b.stopOnce.Do(func() {
		// Stop accepting new events before signalling shutdown, so nothing can be
		// added to the batcher after the final flush has drained it.
		b.stopped.Store(true)
		close(b.shutdown)
	})

	<-b.done
}
