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

// stopFlushTimeout bounds how long Stop waits for the sends still on the wire.
// Bounded because the sender retries for seconds of its own, and a shutdown
// must not sit behind an API that has stopped answering.
const stopFlushTimeout = 5 * time.Second

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

	// sending tracks the goroutines that put events on the wire: the ticker
	// loop and every send FlushAsync started. Stop waits on them, so a batch
	// flushed just before a Close is not lost when the process exits behind it.
	sending sync.WaitGroup

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
	buffer.startPeriodicFlush()

	return buffer
}

// Flush sends whatever is buffered right now instead of waiting for the next
// tick, on the calling goroutine.
func (b *eventBuffer) Flush() {
	b.sendEvents(b.drain())
}

// FlushAsync sends whatever is buffered right now on a goroutine of its own, and
// closes the returned channel once that send has finished.
//
// It is what a caller that must not block on network I/O uses: the client's
// event worker services flush requests, and a synchronous send there would hold
// every enqueued event behind a batch the API is slow to accept. Draining first
// means the events are out of the buffer before this returns, so a later flush
// cannot send them twice.
func (b *eventBuffer) FlushAsync() <-chan struct{} {
	events := b.drain()
	sent := make(chan struct{})
	b.sending.Add(1)
	go func() {
		defer b.sending.Done()
		defer close(sent)
		b.sendEvents(events)
	}()
	return sent
}

// drain takes everything buffered right now, under the lock.
func (b *eventBuffer) drain() []*schematicgo.CreateEventRequestBody {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.batcher.Flush()
}

// sendEvents sends events outside the lock so Push callers aren't blocked during HTTP retries.
func (b *eventBuffer) sendEvents(events []*schematicgo.CreateEventRequestBody) {
	if events == nil {
		return
	}

	err := b.sender.SendBatch(context.Background(), events)
	if err != nil {
		b.errors <- err
	}
}

// startPeriodicFlush runs the ticker loop, tracked so Stop waits for the flush
// it does on its way out.
func (b *eventBuffer) startPeriodicFlush() {
	b.sending.Add(1)
	go func() {
		defer b.sending.Done()
		b.periodicFlush()
	}()
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
			b.Flush()

			return
		case <-ticker.C:
			b.Flush()
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
	full := b.batcher.Add(event)
	var events []*schematicgo.CreateEventRequestBody
	if full {
		events = b.batcher.Flush()
	}
	b.mutex.Unlock()

	b.sendEvents(events)
}

// Stop shuts the buffer down and waits for the events already on the wire,
// including the final flush the ticker loop does on its way out. Without the
// wait, a process exiting right behind a Close would take an in-flight batch
// with it.
func (b *eventBuffer) Stop() {
	defer func() {
		if r := recover(); r != nil {
			b.logger.Error(context.Background(), "Panic occurred while closing client %v", r)
		}
	}()

	close(b.shutdown)

	sent := make(chan struct{})
	go func() {
		b.sending.Wait()
		close(sent)
	}()
	timer := time.NewTimer(stopFlushTimeout)
	defer timer.Stop()
	select {
	case <-sent:
	case <-timer.C:
		b.logger.Error(context.Background(), "Gave up waiting for buffered events to be sent; some may be lost")
	}
}
