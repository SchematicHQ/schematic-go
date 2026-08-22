package client

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// captureServer stands in for the event capture API, recording the event types
// it received and optionally stalling to widen the shutdown race window.
type captureServer struct {
	*httptest.Server

	mu     sync.Mutex
	events []string
	delay  time.Duration
}

func newCaptureServer(delay time.Duration) *captureServer {
	c := &captureServer{delay: delay}
	c.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var payload struct {
			Events []struct {
				EventType string `json:"type"`
			} `json:"events"`
		}
		if err := json.Unmarshal(body, &payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if c.delay > 0 {
			time.Sleep(c.delay)
		}

		c.mu.Lock()
		for _, e := range payload.Events {
			c.events = append(c.events, e.EventType)
		}
		c.mu.Unlock()

		w.WriteHeader(http.StatusOK)
	}))
	return c
}

func (c *captureServer) received() []string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return append([]string(nil), c.events...)
}

func newTestClient(t *testing.T, captureURL string) *SchematicClient {
	t.Helper()

	return NewSchematicClient(
		option.WithAPIKey("test-key"),
		option.WithEventCaptureBaseURL(captureURL),
		// Long buffer period: nothing should be delivered by the periodic flush
		// during these tests, so any delivery is attributable to Close/Flush.
		option.WithEventBufferPeriod(time.Hour),
	)
}

// TestClose_FlushesBufferedEvents is the regression test for the reported bug:
// `defer client.Close()` used to return immediately while the flush ran in an
// unwaited goroutine, so an event tracked just before shutdown (e.g. a credit
// lease redemption) could be lost when the process exited.
func TestClose_FlushesBufferedEvents(t *testing.T) {
	server := newCaptureServer(50 * time.Millisecond)
	defer server.Close()

	client := newTestClient(t, server.URL)

	client.Track(context.Background(), &schematicgo.EventBodyTrack{
		Event: "lease-redemption",
		Company: map[string]string{
			"company_id": "co-1",
		},
	})

	client.Close()

	// No sleep: Close must not return until the flush has completed.
	assert.Equal(t, []string{"track"}, server.received())
}

// TestClose_IsIdempotent guards the sync.Once around the stop channel; Close
// used to rely on recovering from a "close of closed channel" panic.
func TestClose_IsIdempotent(t *testing.T) {
	server := newCaptureServer(0)
	defer server.Close()

	client := newTestClient(t, server.URL)

	client.Track(context.Background(), &schematicgo.EventBodyTrack{
		Event:   "test",
		Company: map[string]string{"company_id": "co-1"},
	})

	client.Close()
	client.Close()

	assert.Equal(t, []string{"track"}, server.received())
}

// TestFlush_DeliversWithoutClosing covers the non-terminal durability path: a
// caller that needs one specific event to be durable (a lease redemption)
// without shutting the client down.
func TestFlush_DeliversWithoutClosing(t *testing.T) {
	server := newCaptureServer(0)
	defer server.Close()

	client := newTestClient(t, server.URL)
	defer client.Close()

	client.Track(context.Background(), &schematicgo.EventBodyTrack{
		Event:   "lease-redemption",
		Company: map[string]string{"company_id": "co-1"},
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	require.NoError(t, client.Flush(ctx))

	assert.Equal(t, []string{"track"}, server.received())

	// The client is still usable after a flush.
	client.Track(context.Background(), &schematicgo.EventBodyTrack{
		Event:   "second",
		Company: map[string]string{"company_id": "co-1"},
	})
	require.NoError(t, client.Flush(ctx))
	assert.Len(t, server.received(), 2)
}

// TestFlush_AfterCloseReturnsError ensures Flush fails fast rather than
// blocking forever once the worker has exited.
func TestFlush_AfterCloseReturnsError(t *testing.T) {
	server := newCaptureServer(0)
	defer server.Close()

	client := newTestClient(t, server.URL)
	client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	assert.ErrorIs(t, client.Flush(ctx), ErrClientClosed)
}

// TestFlush_HonorsContext ensures a cancelled context unblocks Flush.
func TestFlush_HonorsContext(t *testing.T) {
	server := newCaptureServer(0)
	defer server.Close()

	client := newTestClient(t, server.URL)
	defer client.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	assert.ErrorIs(t, client.Flush(ctx), context.Canceled)
}

// TestClose_IsBoundedWhenCaptureHangs ensures a wedged capture endpoint delays
// shutdown but cannot hang it: Close gives up once the shutdown budget is
// spent rather than blocking the caller's termination path forever.
func TestClose_IsBoundedWhenCaptureHangs(t *testing.T) {
	blocked := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-blocked
	}))
	defer server.Close()
	defer close(blocked)

	client := NewSchematicClient(
		option.WithAPIKey("test-key"),
		option.WithEventCaptureBaseURL(server.URL),
		option.WithEventBufferPeriod(time.Hour),
		option.WithShutdownTimeout(200*time.Millisecond),
	)

	client.Track(context.Background(), &schematicgo.EventBodyTrack{
		Event:   "lease-redemption",
		Company: map[string]string{"company_id": "co-1"},
	})

	done := make(chan time.Duration, 1)
	go func() {
		start := time.Now()
		client.Close()
		done <- time.Since(start)
	}()

	select {
	case elapsed := <-done:
		// 2*shutdownTimeout + closeGrace, with slack for slow CI.
		assert.Less(t, elapsed, 10*time.Second, "Close should give up rather than block on a wedged endpoint")
	case <-time.After(15 * time.Second):
		t.Fatal("Close never returned")
	}
}

// TestFlush_ReturnsSendError covers the durability contract: a caller that
// needs an event to have landed must be able to tell that it did not.
func TestFlush_ReturnsSendError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := newTestClient(t, server.URL)
	defer client.Close()

	client.Track(context.Background(), &schematicgo.EventBodyTrack{
		Event:   "lease-redemption",
		Company: map[string]string{"company_id": "co-1"},
	})

	// Short context: the sender retries with backoff, and the context must bound
	// the send itself rather than just the wait for it.
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	assert.Error(t, client.Flush(ctx), "Flush must not report success when delivery failed")
}

// TestTrack_AfterCloseDoesNotBlock ensures tracking past shutdown fails fast.
// The events channel is buffered, so a plain send would hang the caller once
// the worker has exited and the buffer has filled.
func TestTrack_AfterCloseDoesNotBlock(t *testing.T) {
	server := newCaptureServer(0)
	defer server.Close()

	client := newTestClient(t, server.URL)
	client.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := 0; i < 500; i++ {
			client.Track(context.Background(), &schematicgo.EventBodyTrack{
				Event:   "after-close",
				Company: map[string]string{"company_id": "co-1"},
			})
		}
	}()

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("Track blocked after Close")
	}
}

// TestConcurrentTrackAndClose exercises the shutdown path against in-flight
// producers: Close must return, and no Track may block or panic.
func TestConcurrentTrackAndClose(t *testing.T) {
	server := newCaptureServer(0)
	defer server.Close()

	client := newTestClient(t, server.URL)

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				client.Track(context.Background(), &schematicgo.EventBodyTrack{
					Event:   "concurrent",
					Company: map[string]string{"company_id": "co-1"},
				})
			}
		}()
	}

	// Close concurrently with the producers, then again to exercise idempotency.
	client.Close()
	client.Close()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("Track blocked while the client was closing")
	}
}
