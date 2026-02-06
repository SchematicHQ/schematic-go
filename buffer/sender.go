package buffer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	schematicgo "github.com/schematichq/schematic-go"
	"github.com/schematichq/schematic-go/core"
)

// EventPayload represents an event in the format expected by the capture service
type EventPayload struct {
	APIKey    string                     `json:"api_key"`
	Body      *schematicgo.EventBody     `json:"body"`
	EventType schematicgo.EventType      `json:"event_type"`
	SentAt    *time.Time                 `json:"sent_at,omitempty"`
}

// BatchPayload represents the batch request body
type BatchPayload struct {
	Events []EventPayload `json:"events"`
}

// HTTPEventSender sends events to the Schematic capture API with retry logic
type HTTPEventSender struct {
	client            core.HTTPClient
	options           *core.RequestOptions
	logger            core.Logger
	maxRetries        int
	initialRetryDelay time.Duration
}

// NewHTTPEventSender creates a new HTTP event sender with retry logic
func NewHTTPEventSender(client core.HTTPClient, options *core.RequestOptions, logger core.Logger) *HTTPEventSender {
	// Ensure EventCaptureBaseURL is set
	if options.EventCaptureBaseURL == "" {
		options.EventCaptureBaseURL = "https://c.schematichq.com"
	}

	if client == nil {
		client = http.DefaultClient
	}

	return &HTTPEventSender{
		client:            client,
		options:           options,
		logger:            logger,
		maxRetries:        3,
		initialRetryDelay: 1 * time.Second,
	}
}

// SendBatch sends a batch of events to the /batch endpoint with retry logic
func (h *HTTPEventSender) SendBatch(ctx context.Context, events []*schematicgo.CreateEventRequestBody) error {
	if len(events) == 0 {
		return nil
	}

	var lastErr error
	
	for attempt := 0; attempt <= h.maxRetries; attempt++ {
		if attempt > 0 {
			h.logger.Info(ctx, fmt.Sprintf("Retrying event batch submission (attempt %d of %d)", attempt, h.maxRetries))
		}
		
		err := h.sendBatch(ctx, events)
		if err == nil {
			if attempt > 0 {
				h.logger.Info(ctx, fmt.Sprintf("Event batch submission succeeded after %d retries", attempt))
			}
			return nil
		}
		
		lastErr = err
		
		// Don't sleep after the last attempt
		if attempt < h.maxRetries {
			delay := h.calculateBackoff(attempt)
			h.logger.Warn(ctx, fmt.Sprintf("Event batch submission failed: %v. Retrying in %.2f seconds...",
				err, float64(delay)/float64(time.Second)))
			
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(delay):
				// Continue to next retry
			}
		}
	}
	
	err := fmt.Errorf("failed after %d retries: %w", h.maxRetries, lastErr)
	h.logger.Error(ctx, fmt.Sprintf("Event batch submission %s", err.Error()))
	return err
}

// sendBatch performs a single send attempt
func (h *HTTPEventSender) sendBatch(ctx context.Context, events []*schematicgo.CreateEventRequestBody) error {
	endpoint := h.options.EventCaptureBaseURL + "/batch"
	
	// Transform events to EventPayload structs
	eventPayloads := make([]EventPayload, len(events))
	for i, event := range events {
		eventPayloads[i] = EventPayload{
			APIKey:    h.options.APIKey,
			Body:      event.Body,
			EventType: event.EventType,
			SentAt:    event.SentAt,
		}
	}
	
	// Create batch payload
	batchPayload := BatchPayload{
		Events: eventPayloads,
	}
	
	// Marshal to JSON
	body, err := json.Marshal(batchPayload)
	if err != nil {
		return fmt.Errorf("failed to marshal batch payload: %w", err)
	}
	
	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	
	// Set headers
	headers := h.options.ToHeader()
	for key, values := range headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	req.Header.Set("Content-Type", "application/json")
	
	// Send the request
	resp, err := h.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()
	
	// Check response status
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		// Read response body for error details
		respBody, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return fmt.Errorf("HTTP %d: (failed to read response body: %v)", resp.StatusCode, readErr)
		}
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(respBody))
	}
	
	return nil
}

// calculateBackoff calculates exponential backoff with jitter
func (h *HTTPEventSender) calculateBackoff(attempt int) time.Duration {
	// Exponential backoff: initialDelay * 2^attempt
	delay := h.initialRetryDelay * time.Duration(1<<uint(attempt))
	
	// Add jitter (±25%)
	jitter := time.Duration(rand.Int63n(int64(delay) / 2))
	delay = delay - delay/4 + jitter
	
	return delay
}
