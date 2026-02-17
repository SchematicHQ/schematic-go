package buffer

import (
	"context"

	schematicgo "github.com/schematichq/schematic-go"
)

// Batcher handles accumulating and preparing events for sending.
// This type is NOT thread-safe; callers must provide external synchronization.
type Batcher struct {
	maxEvents int
	events    []*schematicgo.CreateEventRequestBody
}

// NewBatcher creates a new batcher with the specified max events
func NewBatcher(maxEvents int) *Batcher {
	return &Batcher{
		maxEvents: maxEvents,
		events:    make([]*schematicgo.CreateEventRequestBody, 0, maxEvents),
	}
}

// Add adds an event to the batch. Returns true if the batch is full after adding.
func (b *Batcher) Add(event *schematicgo.CreateEventRequestBody) bool {
	if event == nil {
		return false
	}

	b.events = append(b.events, event)
	return len(b.events) >= b.maxEvents
}

// Flush returns all events and clears the batch
func (b *Batcher) Flush() []*schematicgo.CreateEventRequestBody {
	if len(b.events) == 0 {
		return nil
	}

	// Copy events and filter out nils
	events := make([]*schematicgo.CreateEventRequestBody, 0, len(b.events))
	for _, event := range b.events {
		if event != nil {
			events = append(events, event)
		}
	}

	// Clear the batch
	b.events = b.events[:0]

	return events
}

// EventSender is an interface for sending batched events
type EventSender interface {
	SendBatch(ctx context.Context, events []*schematicgo.CreateEventRequestBody) error
}
