package buffer

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Integration test that doesn't require actual HTTP calls
func TestEventBuffer_Integration(t *testing.T) {
	t.Run("flushes when batch is full", func(t *testing.T) {
		// Create a mock sender to capture calls
		mockSender := &mockSender{}
		logger := &mockLogger{}

		// Create buffer with small batch size for testing
		errors := make(chan error, 10)
		buffer := &eventBuffer{
			batcher:  NewBatcher(2), // Small batch size
			sender:   mockSender,
			errors:   errors,
			interval: 100 * time.Millisecond,
			logger:   logger,
			shutdown: make(chan struct{}),
		}

		// Push events - should auto-flush when full
		buffer.Push(&schematicgo.CreateEventRequestBody{EventType: "test1"})
		buffer.Push(&schematicgo.CreateEventRequestBody{EventType: "test2"})

		// Give a moment for async operations
		time.Sleep(10 * time.Millisecond)

		// Should have flushed
		calls := mockSender.snapshot()
		require.Len(t, calls, 1)
		assert.Len(t, calls[0].events, 2)
	})

	t.Run("periodic flush works", func(t *testing.T) {
		mockSender := &mockSender{}
		logger := &mockLogger{}

		options := core.NewRequestOptions()
		options.APIKey = "test-key"

		// Create a real buffer with very short interval
		errors := make(chan error, 10)
		client := http.DefaultClient

		buffer := NewEventBuffer(options, client, errors, logger, func() *time.Duration {
			d := 50 * time.Millisecond
			return &d
		}())

		// Replace the sender with our mock after creation
		buffer.sender = mockSender

		// Push an event
		buffer.Push(&schematicgo.CreateEventRequestBody{EventType: "test"})

		// Wait for periodic flush
		time.Sleep(100 * time.Millisecond)

		// Should have been flushed
		assert.GreaterOrEqual(t, len(mockSender.snapshot()), 1)

		// Cleanup
		buffer.Stop()
	})

	t.Run("handles errors gracefully", func(t *testing.T) {
		// Mock sender that always fails
		mockSender := &mockSender{
			responses: []error{assert.AnError, assert.AnError},
		}
		logger := &mockLogger{}

		errors := make(chan error, 10)
		buffer := &eventBuffer{
			batcher:  NewBatcher(10),
			sender:   mockSender,
			errors:   errors,
			logger:   logger,
			shutdown: make(chan struct{}),
		}

		// An explicit Flush returns the send error to its caller rather than
		// routing it to the error channel, so a caller that needs an event to be
		// durable can tell whether it landed.
		buffer.Push(&schematicgo.CreateEventRequestBody{EventType: "test"})
		assert.ErrorIs(t, buffer.Flush(context.Background()), assert.AnError)

		// Flushes with no caller to return to -- the periodic loop and shutdown --
		// report on the error channel instead.
		buffer.Push(&schematicgo.CreateEventRequestBody{EventType: "test"})
		buffer.flushAndReport(context.Background())

		select {
		case err := <-errors:
			assert.ErrorIs(t, err, assert.AnError)
		case <-time.After(100 * time.Millisecond):
			t.Fatal("Expected error but got timeout")
		}
	})
}

// Example of how easy it is to test individual components now
func TestComponentIntegration(t *testing.T) {
	t.Run("batcher and sender work together", func(t *testing.T) {
		// Setup
		batcher := NewBatcher(5)
		sender := &mockSender{}

		// Add events
		for i := 0; i < 3; i++ {
			batcher.Add(&schematicgo.CreateEventRequestBody{
				EventType: "test",
			})
		}

		// Flush and send
		events := batcher.Flush()
		err := sender.SendBatch(context.Background(), events)

		// Verify
		assert.NoError(t, err)
		assert.Len(t, sender.calls, 1)
		assert.Len(t, sender.calls[0].events, 3)
	})
}

// Mock implementations for testing

// mockSender is written by the buffer's flush goroutine and read by the test
// goroutine, so it must be safe for concurrent use.
type mockSender struct {
	mu        sync.Mutex
	calls     []mockCall
	responses []error
	callIndex int
}

// snapshot returns a copy of the calls recorded so far.
func (m *mockSender) snapshot() []mockCall {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]mockCall(nil), m.calls...)
}

type mockCall struct {
	ctx    context.Context
	events []*schematicgo.CreateEventRequestBody
}

func (m *mockSender) SendBatch(ctx context.Context, events []*schematicgo.CreateEventRequestBody) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.calls = append(m.calls, mockCall{ctx: ctx, events: events})
	if m.callIndex < len(m.responses) {
		err := m.responses[m.callIndex]
		m.callIndex++
		return err
	}
	return nil
}

type mockLogger struct {
	infos  []string
	warns  []string
	errors []string
	debugs []string
}

func (m *mockLogger) Debug(ctx context.Context, message string, args ...interface{}) {
	m.debugs = append(m.debugs, message)
}

func (m *mockLogger) Info(ctx context.Context, message string, args ...interface{}) {
	m.infos = append(m.infos, message)
}

func (m *mockLogger) Warn(ctx context.Context, message string, args ...interface{}) {
	m.warns = append(m.warns, message)
}

func (m *mockLogger) Error(ctx context.Context, message string, args ...interface{}) {
	m.errors = append(m.errors, message)
}

func TestEventBuffer_ShutdownFlushesRemaining(t *testing.T) {
	mockSender := &mockSender{}
	logger := &mockLogger{}

	errors := make(chan error, 10)
	buffer := &eventBuffer{
		batcher:  NewBatcher(100), // Large batch so it won't auto-flush
		sender:   mockSender,
		errors:   errors,
		interval: 10 * time.Second, // Long interval so periodic flush won't trigger
		logger:   logger,
		shutdown: make(chan struct{}),
		done:     make(chan struct{}),
	}

	// Start the periodic flush goroutine (mirrors NewEventBuffer behavior)
	go buffer.periodicFlush()

	// Push several events (fewer than batch size so no auto-flush)
	for i := 0; i < 5; i++ {
		buffer.Push(&schematicgo.CreateEventRequestBody{
			EventType: schematicgo.EventType(fmt.Sprintf("shutdown-test-%d", i)),
		})
	}

	// No flush should have happened yet
	require.Len(t, mockSender.snapshot(), 0)

	// Stop the buffer, which should flush remaining events. Stop blocks until the
	// flush has completed, so no sleep is needed here.
	buffer.Stop()

	// Verify all 5 events were flushed
	totalEvents := 0
	for _, call := range mockSender.snapshot() {
		totalEvents += len(call.events)
	}
	assert.Equal(t, 5, totalEvents)
}

// syncMockSender is a thread-safe mock sender for concurrent tests
type syncMockSender struct {
	mu        sync.Mutex
	calls     []mockCall
	responses []error
	callIndex int
}

func (m *syncMockSender) SendBatch(ctx context.Context, events []*schematicgo.CreateEventRequestBody) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.calls = append(m.calls, mockCall{ctx: ctx, events: events})
	if m.callIndex < len(m.responses) {
		err := m.responses[m.callIndex]
		m.callIndex++
		return err
	}
	return nil
}

func (m *syncMockSender) totalEvents() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	total := 0
	for _, call := range m.calls {
		total += len(call.events)
	}
	return total
}

func TestEventBuffer_ConcurrentPush(t *testing.T) {
	sender := &syncMockSender{}
	logger := &mockLogger{}

	options := core.NewRequestOptions()
	options.APIKey = "test-key"

	errors := make(chan error, 100)
	client := http.DefaultClient

	buffer := NewEventBuffer(options, client, errors, logger, func() *time.Duration {
		d := 50 * time.Millisecond
		return &d
	}())

	// Replace the sender with our thread-safe mock
	buffer.sender = sender

	numGoroutines := 10
	eventsPerGoroutine := 20
	totalExpected := numGoroutines * eventsPerGoroutine

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Launch multiple goroutines pushing events concurrently
	for g := 0; g < numGoroutines; g++ {
		go func(goroutineID int) {
			defer wg.Done()
			for i := 0; i < eventsPerGoroutine; i++ {
				buffer.Push(&schematicgo.CreateEventRequestBody{
					EventType: schematicgo.EventType(fmt.Sprintf("concurrent-%d-%d", goroutineID, i)),
				})
			}
		}(g)
	}

	// Wait for all goroutines to finish pushing
	wg.Wait()

	// Stop the buffer to flush any remaining events; Stop blocks until done
	buffer.Stop()

	// Verify all events were eventually sent
	assert.Equal(t, totalExpected, sender.totalEvents())
}

// slowSender blocks for a fixed duration inside SendBatch, so a Stop that
// returned before the flush completed would be observable.
type slowSender struct {
	mu    sync.Mutex
	delay time.Duration
	sent  int
}

func (s *slowSender) SendBatch(ctx context.Context, events []*schematicgo.CreateEventRequestBody) error {
	time.Sleep(s.delay)
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sent += len(events)
	return nil
}

func (s *slowSender) total() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.sent
}

// TestEventBuffer_StopBlocksUntilFlushCompletes is the regression test for the
// lost-events-on-shutdown bug: Stop used to close the shutdown channel and
// return immediately, leaving the final flush running in an unwaited goroutine
// that the exiting process would never wait for.
func TestEventBuffer_StopBlocksUntilFlushCompletes(t *testing.T) {
	sender := &slowSender{delay: 150 * time.Millisecond}

	buffer := &eventBuffer{
		batcher:  NewBatcher(100),
		sender:   sender,
		errors:   make(chan error, 10),
		interval: time.Hour, // periodic flush must not fire
		logger:   &mockLogger{},
		shutdown: make(chan struct{}),
		done:     make(chan struct{}),
	}
	go buffer.periodicFlush()

	buffer.Push(&schematicgo.CreateEventRequestBody{EventType: "track"})

	buffer.Stop()

	// No sleep: if Stop returned before the slow send finished, this is 0.
	assert.Equal(t, 1, sender.total())
}

// TestEventBuffer_StopIsIdempotent guards the sync.Once around the shutdown
// channel; a second Stop must neither panic nor block.
func TestEventBuffer_StopIsIdempotent(t *testing.T) {
	sender := &slowSender{}

	buffer := &eventBuffer{
		batcher:  NewBatcher(100),
		sender:   sender,
		errors:   make(chan error, 10),
		interval: time.Hour,
		logger:   &mockLogger{},
		shutdown: make(chan struct{}),
		done:     make(chan struct{}),
	}
	go buffer.periodicFlush()

	buffer.Push(&schematicgo.CreateEventRequestBody{EventType: "track"})
	buffer.Stop()
	buffer.Stop()

	assert.Equal(t, 1, sender.total())
}

// TestEventBuffer_StopDoesNotBlockOnFullErrorChannel ensures a failing final
// flush can't deadlock shutdown when nothing is draining the error channel.
func TestEventBuffer_StopDoesNotBlockOnFullErrorChannel(t *testing.T) {
	errors := make(chan error) // unbuffered, nobody reading

	buffer := &eventBuffer{
		batcher:  NewBatcher(100),
		sender:   &mockSender{responses: []error{assert.AnError}},
		errors:   errors,
		interval: time.Hour,
		logger:   &mockLogger{},
		shutdown: make(chan struct{}),
		done:     make(chan struct{}),
	}
	go buffer.periodicFlush()

	buffer.Push(&schematicgo.CreateEventRequestBody{EventType: "track"})

	stopped := make(chan struct{})
	go func() {
		buffer.Stop()
		close(stopped)
	}()

	select {
	case <-stopped:
	case <-time.After(5 * time.Second):
		t.Fatal("Stop deadlocked on the error channel")
	}
}

// hangingSender ignores context cancellation entirely, standing in for a sender
// wedged somewhere Go's context plumbing can't reach.
type hangingSender struct {
	released chan struct{}
}

func (h *hangingSender) SendBatch(ctx context.Context, events []*schematicgo.CreateEventRequestBody) error {
	<-h.released
	return nil
}

// TestEventBuffer_StopIsBounded ensures Stop honours the shutdown budget even
// when the sender never returns, so a wedged endpoint delays shutdown but
// cannot hang it.
func TestEventBuffer_StopIsBounded(t *testing.T) {
	sender := &hangingSender{released: make(chan struct{})}
	defer close(sender.released)

	buffer := &eventBuffer{
		batcher:         NewBatcher(100),
		sender:          sender,
		errors:          make(chan error, 10),
		interval:        time.Hour,
		logger:          &mockLogger{},
		shutdown:        make(chan struct{}),
		done:            make(chan struct{}),
		shutdownTimeout: 100 * time.Millisecond,
	}
	go buffer.periodicFlush()

	buffer.Push(&schematicgo.CreateEventRequestBody{EventType: "track"})

	start := time.Now()
	buffer.Stop()
	elapsed := time.Since(start)

	// 2*shutdownTimeout + shutdownWaitGrace, plus slack for slow CI.
	assert.Less(t, elapsed, 2*time.Second, "Stop should give up rather than block on a wedged sender")
}

// TestEventBuffer_ShutdownFlushIsBoundedByTimeout ensures the final flush stops
// retrying once the shutdown budget is spent, rather than running until the
// sender's own retry schedule is exhausted.
func TestEventBuffer_ShutdownFlushIsBoundedByTimeout(t *testing.T) {
	deadlines := make(chan time.Time, 1)
	sender := &deadlineSender{deadlines: deadlines}

	buffer := &eventBuffer{
		batcher:         NewBatcher(100),
		sender:          sender,
		errors:          make(chan error, 10),
		interval:        time.Hour,
		logger:          &mockLogger{},
		shutdown:        make(chan struct{}),
		done:            make(chan struct{}),
		shutdownTimeout: 250 * time.Millisecond,
	}
	go buffer.periodicFlush()

	buffer.Push(&schematicgo.CreateEventRequestBody{EventType: "track"})
	buffer.Stop()

	select {
	case deadline := <-deadlines:
		assert.WithinDuration(t, time.Now().Add(250*time.Millisecond), deadline, time.Second,
			"final flush should carry the shutdown budget as its deadline")
	default:
		t.Fatal("final flush never reached the sender")
	}
}

// deadlineSender records the deadline of the context it was handed.
type deadlineSender struct {
	deadlines chan time.Time
}

func (d *deadlineSender) SendBatch(ctx context.Context, events []*schematicgo.CreateEventRequestBody) error {
	if deadline, ok := ctx.Deadline(); ok {
		select {
		case d.deadlines <- deadline:
		default:
		}
	}
	return nil
}
