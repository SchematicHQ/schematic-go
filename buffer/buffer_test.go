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

		// Push and flush
		buffer.Push(&schematicgo.CreateEventRequestBody{EventType: "test"})
		buffer.Flush()

		// Should receive error
		select {
		case err := <-errors:
			assert.Error(t, err)
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
		calls := sender.snapshot()
		assert.Len(t, calls, 1)
		assert.Len(t, calls[0].events, 3)
	})
}

// Mock implementations for testing

// mockSender is written from whichever goroutine is flushing and read from the
// test's, so every access goes through the mutex.
type mockSender struct {
	mu        sync.Mutex
	calls     []mockCall
	responses []error
	callIndex int
}

// snapshot is the calls recorded so far.
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

	// Stop the buffer, which should flush remaining events
	buffer.Stop()

	// Give a moment for the shutdown goroutine to complete
	time.Sleep(50 * time.Millisecond)

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

	// Stop the buffer to flush any remaining events
	buffer.Stop()

	// Give time for the shutdown flush to complete
	time.Sleep(100 * time.Millisecond)

	// Verify all events were eventually sent
	assert.Equal(t, totalExpected, sender.totalEvents())
}

// slowSender holds each batch on the wire for a beat, so a Stop that does not
// wait finishes before the send does.
type slowSender struct {
	delay time.Duration

	mu    sync.Mutex
	batch int
}

func (s *slowSender) SendBatch(_ context.Context, events []*schematicgo.CreateEventRequestBody) error {
	time.Sleep(s.delay)
	s.mu.Lock()
	defer s.mu.Unlock()
	s.batch += len(events)
	return nil
}

func (s *slowSender) sent() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.batch
}

// A flush the client asked for runs on a goroutine of its own, so a Stop that
// did not wait for it would let the process exit with the batch still in the
// air: an identify flushed ahead of a prewarm, most of all.
func TestEventBuffer_StopWaitsForAnAsyncFlush(t *testing.T) {
	sender := &slowSender{delay: 200 * time.Millisecond}
	buffer := &eventBuffer{
		batcher:  NewBatcher(100),
		sender:   sender,
		errors:   make(chan error, 10),
		interval: time.Minute,
		logger:   &mockLogger{},
		shutdown: make(chan struct{}),
	}
	buffer.startPeriodicFlush()

	buffer.Push(&schematicgo.CreateEventRequestBody{EventType: "async-flush"})
	sent := buffer.FlushAsync()
	buffer.Stop()

	assert.Equal(t, 1, sender.sent(), "Stop returned before the flush it started had finished")
	select {
	case <-sent:
	default:
		t.Fatal("the flush should have completed by the time Stop returned")
	}
}

// The ticker loop sends whatever is still buffered on its way out, which is the
// other half of the same promise.
func TestEventBuffer_StopWaitsForTheShutdownFlush(t *testing.T) {
	sender := &slowSender{delay: 200 * time.Millisecond}
	buffer := &eventBuffer{
		batcher:  NewBatcher(100),
		sender:   sender,
		errors:   make(chan error, 10),
		interval: time.Minute,
		logger:   &mockLogger{},
		shutdown: make(chan struct{}),
	}
	buffer.startPeriodicFlush()

	buffer.Push(&schematicgo.CreateEventRequestBody{EventType: "shutdown-flush"})
	buffer.Stop()

	assert.Equal(t, 1, sender.sent(), "Stop returned before the shutdown flush had finished")
}
