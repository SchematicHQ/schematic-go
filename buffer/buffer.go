package buffer

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/core"
)

const defaultEventBufferPeriod = 5 * time.Second
const maxEvents = 100

// DefaultShutdownTimeout is the budget for the buffer's final flush: how long
// Stop will spend delivering whatever is still buffered before giving up, so a
// wedged HTTP client can't hang a caller's shutdown path indefinitely.
const DefaultShutdownTimeout = 10 * time.Second

// sendTimeout bounds a single batch send outside of shutdown. It sits
// comfortably above the sender's retry schedule, so it never truncates a
// legitimate retry -- it exists only so a send cannot hang the worker
// goroutine forever.
const sendTimeout = 30 * time.Second

// shutdownWaitGrace is the extra time Stop allows the flush goroutine to unwind
// after its flush budgets have expired. Without it Stop's timer and the flush's
// own deadline would expire together, and Stop could report a timeout for a
// flush that was about to finish.
const shutdownWaitGrace = 500 * time.Millisecond

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

	shutdownTimeout := DefaultShutdownTimeout
	if options != nil && options.ShutdownTimeout != nil && *options.ShutdownTimeout > 0 {
		shutdownTimeout = *options.ShutdownTimeout
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
		shutdownTimeout: shutdownTimeout,
	}

	// Start ticker to flush events periodically
	go buffer.periodicFlush()

	return buffer
}

// timeout is the buffer's shutdown budget, tolerating a zero value so a
// directly-constructed buffer still has a bound.
func (b *eventBuffer) timeout() time.Duration {
	if b.shutdownTimeout <= 0 {
		return DefaultShutdownTimeout
	}
	return b.shutdownTimeout
}

// Flush sends any buffered events and blocks until the send (including retries)
// has completed. The send error, if any, is returned rather than routed to the
// error channel, so a caller that needs an event to be durable can tell whether
// it actually landed.
func (b *eventBuffer) Flush(ctx context.Context) error {
	b.mutex.Lock()
	events := b.batcher.Flush()
	b.mutex.Unlock()

	if len(events) == 0 {
		return nil
	}

	return b.sender.SendBatch(ctx, events)
}

// flushAndReport flushes on behalf of the periodic loop and shutdown, which
// have no caller to return an error to.
func (b *eventBuffer) flushAndReport(ctx context.Context) {
	if err := b.Flush(ctx); err != nil {
		b.reportError(ctx, err)
	}
}

// sendAndReport sends an already-drained batch, reporting rather than returning
// any error.
func (b *eventBuffer) sendAndReport(ctx context.Context, events []*schematicgo.CreateEventRequestBody) {
	if len(events) == 0 {
		return
	}

	if err := b.sender.SendBatch(ctx, events); err != nil {
		b.reportError(ctx, err)
	}
}

// reportError never blocks on the error channel: during shutdown the worker
// goroutine is waiting on the flush and is no longer draining it, so a blocking
// send would deadlock.
func (b *eventBuffer) reportError(ctx context.Context, err error) {
	select {
	case b.errors <- err:
	default:
		b.logger.Error(ctx, fmt.Sprintf("%v", err))
	}
}

func (b *eventBuffer) periodicFlush() {
	defer close(b.done)

	ticker := time.NewTicker(b.interval)
	defer ticker.Stop()

	for {
		select {
		case <-b.shutdown:
			// Flush what remains, bounded so shutdown can't hang forever.
			ctx, cancel := context.WithTimeout(context.Background(), b.timeout())
			b.flushAndReport(ctx)
			cancel()

			return
		case <-ticker.C:
			// Bounded like the final flush: the loop cannot observe b.shutdown
			// while a flush is in flight, so an unbounded periodic flush would
			// delay shutdown indefinitely.
			ctx, cancel := context.WithTimeout(context.Background(), b.timeout())
			b.flushAndReport(ctx)
			cancel()
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

	if len(events) == 0 {
		return
	}

	// Bounded: Push runs on the client's worker goroutine, including while it is
	// draining events on the shutdown path, so an unbounded send here would keep
	// the worker alive indefinitely.
	ctx, cancel := context.WithTimeout(context.Background(), sendTimeout)
	defer cancel()

	b.sendAndReport(ctx, events)
}

// Stop stops accepting new events and blocks until the buffered events have
// been flushed, so callers can rely on Stop returning only once in-flight
// events have been delivered (or definitively failed). The wait is bounded, so
// a wedged sender delays shutdown but cannot hang it.
//
// Stop is safe to call more than once.
func (b *eventBuffer) Stop() {
	defer func() {
		if r := recover(); r != nil {
			b.logger.Error(context.Background(), fmt.Sprintf("Panic occurred while closing client %v", r))
		}
	}()

	b.stopOnce.Do(func() {
		// Stop accepting new events before signalling shutdown, so nothing can be
		// added to the batcher after the final flush has drained it.
		b.stopped.Store(true)
		close(b.shutdown)
	})

	// Budget for the worst case: a periodic flush already in flight when shutdown
	// was signalled, followed by the final flush, each bounded by b.timeout().
	// The flushes bound their own HTTP work, but waiting on b.done
	// unconditionally would still let a sender that ignores context cancellation
	// hang the caller.
	timer := time.NewTimer(2*b.timeout() + shutdownWaitGrace)
	defer timer.Stop()

	select {
	case <-b.done:
	case <-timer.C:
		b.logger.Error(context.Background(), "Timed out waiting for buffered events to flush on shutdown")
	}
}
