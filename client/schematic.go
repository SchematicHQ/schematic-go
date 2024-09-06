package client

import (
	"context"
	"time"

	schematicgo "github.com/schematichq/schematic-go"
	core "github.com/schematichq/schematic-go/core"
)

type SchematicClient interface {
	// Check a flag to get a boolean value
	CheckFlag(context.Context, *schematicgo.CheckFlagRequestBody, string) bool

	// Identify a user
	Identify(context.Context, *schematicgo.EventBodyIdentify)

	// Track an action
	Track(context.Context, *schematicgo.EventBodyTrack)

	// Access bindings for Schematic API
	API() *Client

	// Close the client
	Close()

	// Provide a custom cache provide for flag checks
	AddFlagCheckCacheProvider(context.Context, schematicgo.CacheProviderBool)

	// Set a custom API client
	SetAPIClient(*Client)

	// Set a custom API host
	SetAPIHost(context.Context, string)

	// Event buffer period determines how often batches of events are sent to Schematic
	SetEventBufferPeriod(period time.Duration)

	// Set a default flag value
	SetFlagDefault(string, bool)

	// Set multiple default flag values
	SetFlagDefaults(map[string]bool)

	// Allow to specify which http client to use
	SetHTTPClient(httpClient core.HTTPClient)

	// Sets offline mode to true
	SetOfflineMode()
}

// Alias some types so that callers don't need to directly import the API submodule for common operations
