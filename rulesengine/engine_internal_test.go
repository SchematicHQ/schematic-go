package rulesengine

import (
	"context"
	"errors"
	"testing"
	"time"
)

// TestCheckFlagContextCancellation pins that CheckFlag returns ctx.Err() instead
// of blocking when every pool instance is busy and the context is done. It drains
// the pool directly so the wait-for-instance path is exercised deterministically.
func TestCheckFlagContextCancellation(t *testing.T) {
	e, err := NewEngine(context.Background(), WithPoolSize(1))
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = e.Close(context.Background()) }()

	// Take the only instance out of the pool and keep it, so any CheckFlag must
	// wait on the pool channel.
	held := <-e.pool
	defer func() { e.pool <- held }()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	done := make(chan error, 1)
	go func() {
		_, err := e.CheckFlag(ctx, nil, nil, &Flag{ID: "f", Key: "f", DefaultValue: true})
		done <- err
	}()

	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("expected context.Canceled, got %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("CheckFlag blocked on a cancelled context instead of returning")
	}
}
