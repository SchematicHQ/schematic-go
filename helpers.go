package schematic

import (
	"context"
	"errors"
	"time"
)

func NewClient(apiKey string) *APIClient {
	configuration := NewConfiguration()
	configuration.AddDefaultHeader("X-Schematic-Api-Key", apiKey)

	return NewAPIClient(configuration)
}

func SendIdentifyEvent(
	ctx context.Context,
	client *APIClient,
	body *EventBodyIdentify,
) error {
	return sendEvent(ctx, client, "identify", EventBodyIdentifyAsEventBody(body))
}

func SendTrackEvent(
	ctx context.Context,
	client *APIClient,
	body *EventBodyTrack,
) error {
	return sendEvent(ctx, client, "track", EventBodyTrackAsEventBody(body))
}

func sendEvent(
	ctx context.Context,
	client *APIClient,
	eventType string,
	body EventBody,
) error {
	if client == nil {
		return errors.New("schematic client is nil")
	}

	eventBody := NewCreateEventRequestBody(eventType)
	if eventBody == nil {
		return errors.New("schematic event body is unexpectedly nil")
	}

	eventBody.SetBody(body)
	eventBody.SetSentAt(time.Now().UTC())

	if _, _, err := client.EventsApi.CreateEvent(ctx).CreateEventRequestBody(*eventBody).Execute(); err != nil {
		return err
	}

	return nil
}
