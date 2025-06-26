# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This repository contains the official Go SDK for Schematic, a feature flag and product management platform. The library allows Go developers to integrate Schematic's functionality into their applications, including:

- Feature flag checks (with caching)
- User/company identity management
- Event tracking and buffering
- Webhook verification
- API client for direct Schematic API access

Much of this repository is generated using [Fern](https//buildwithfern.com). All files that are not listed in the .fernignore file will be regenerated over. If you need to edit any files in this repo, please ensure they are in .fernignore, or add them; if you're editing a generated file, you likely want to create a new file instead and add it to .fernignore, so as to keep generated and non-generated code separate.

## Building and Testing

### Running Tests

Tests use Go's standard testing package with the testify library for assertions:

```bash
# Run all tests
go test ./...

# Run tests in specific package
go test ./client/...

# Run specific test with verbose output
go test -v ./client -run TestNewClient
```

### Generating Mocks

Mock interfaces for testing are generated using mockgen:

```bash
# Generate mocks
./generate.sh
```

## Architecture Overview

### Key Components

1. **SchematicClient**: The main entry point for interacting with Schematic services
   - Located in `client/schematic_client.go`
   - Provides high-level methods like `CheckFlag`, `Identify`, and `Track`
   - Manages configuration options and client lifecycle

2. **EventBuffer**: Handles batching and asynchronous submission of events
   - Located in `buffer/buffer.go`
   - Includes retry mechanism with exponential backoff and jitter
   - Flushes events periodically or when buffer is full

3. **Cache**: Provides local flag value caching
   - Located in `cache/cache.go`
   - Configurable TTL and size limits

4. **API Clients**: Direct API access for various Schematic resources
   - Located in their respective directories (e.g., `features/client.go`, `companies/client.go`)
   - Automatically generated from Fern API definitions

5. **Option Pattern**: Functional options for client configuration
   - Located in `option/custom_options.go`
   - Allows flexible and extensible client configuration

6. **Core Interfaces**: Shared utilities and interfaces
   - HTTP client abstraction in `core/http.go`
   - Error handling in `core/api_error.go`
   - Logging interface in `core/logger.go`

### Recent Features

- **Event Buffer Retry**: The latest feature (commit 71ef110) adds retry functionality with exponential backoff and jitter to the event buffer for handling transient network issues.

## Development Patterns

When working with this codebase, be aware of the following patterns:

1. **Builder Pattern**: Client construction uses functional options
   ```go
   client := schematicclient.NewSchematicClient(
     option.WithAPIKey(apiKey),
     option.WithOfflineMode(),
   )
   ```

2. **Offline Mode**: For testing without network requests
   ```go
   client := schematicclient.NewSchematicClient(option.WithOfflineMode())
   ```

3. **Error Handling**: API errors are structured and can be type-asserted
   ```go
   if apiError, ok := err.(*core.APIError); ok {
     // Handle based on apiError.StatusCode
   }
   ```

4. **Logging Interface**: Custom logger interface in `core/logger.go`
   - Used throughout the codebase for consistent logging
   - Can be replaced with custom implementation via options

5. **Buffer Management**: Event buffering in `buffer/buffer.go`
   - Thread-safe design with mutexes
   - Periodic flushing with configurable interval
   - Retry mechanism for reliability

## Go Conventions

This codebase follows standard Go conventions:

- Error handling through explicit error returns and checks
- Context propagation for cancellation and timeouts
- Pointer parameters for optional values
- Immutable client interfaces
- Functional options for configuration
