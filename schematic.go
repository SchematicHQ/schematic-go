package schematic

import (
	"context"
	"time"

	"github.com/SchematicHQ/schematic-go/api"
)

type Client interface {
	// Check a flag to get a boolean value
	CheckFlag(context.Context, *CheckFlagRequestBody, string) bool

	// Identify a user
	Identify(context.Context, *EventBodyIdentify)

	// Track an action
	Track(context.Context, *EventBodyTrack)

	// Access bindings for Schematic API
	API() *api.APIClient

	// Close the client
	Close()

	// Add default headers to the API client
	AddDefaultHeaders(ctx context.Context, headers map[string]string)

	// Provide a custom cache provide for flag checks
	AddFlagCheckCacheProvider(context.Context, CacheProvider[bool])

	// Set a custom API client
	SetAPIClient(*api.APIClient)

	// Set a custom API host
	SetAPIHost(context.Context, string)

	// Event buffer period determines how often batches of events are sent to Schematic
	SetEventBufferPeriod(period time.Duration)

	// Set a default flag value
	SetFlagDefault(string, bool)

	// Set multiple default flag values
	SetFlagDefaults(map[string]bool)
}

// Alias some types so that callers don't need to directly import the API submodule for common operations
type CheckFlagRequestBody = api.CheckFlagRequestBody
type EventBodyIdentify = api.EventBodyIdentify
type EventBodyTrack = api.EventBodyTrack
