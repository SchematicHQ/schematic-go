package webhooks

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"net/http"
)

// WebhookSignatureHeader is the HTTP header containing the webhook signature
const WebhookSignatureHeader = "X-Schematic-Webhook-Signature"

// WebhookTimestampHeader is the HTTP header containing the webhook timestamp
const WebhookTimestampHeader = "X-Schematic-Webhook-Timestamp"

// ErrInvalidSignature is returned when the webhook signature is invalid
var ErrInvalidSignature = errors.New("invalid webhook signature")

// ErrMissingSignature is returned when the webhook signature header is missing
var ErrMissingSignature = errors.New("missing webhook signature header")

// ErrMissingTimestamp is returned when the webhook timestamp header is missing
var ErrMissingTimestamp = errors.New("missing webhook timestamp header")

// ComputeSignature computes the HMAC-SHA256 signature for a webhook payload.
// It takes the request body, timestamp, and webhook secret.
func ComputeSignature(body []byte, timestamp string, secret string) []byte {
	// Compute signature: HMAC(body + "+" + timestamp, secret)
	signatureContent := append(body, []byte("+"+timestamp)...)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(signatureContent)
	return mac.Sum(nil)
}

// ComputeHexSignature computes and hex-encodes the HMAC-SHA256 signature for a webhook payload.
// It takes the request body, timestamp, and webhook secret, and returns a hex-encoded string.
func ComputeHexSignature(body []byte, timestamp string, secret string) string {
	return hex.EncodeToString(ComputeSignature(body, timestamp, secret))
}

// VerifySignature verifies a webhook signature against the payload and timestamp.
// The signature should be the hex-encoded HMAC-SHA256 of the payload and timestamp.
func VerifySignature(payload []byte, signature string, timestamp string, secret string) error {
	if signature == "" {
		return ErrMissingSignature
	}

	if timestamp == "" {
		return ErrMissingTimestamp
	}

	// Decode signature
	decodedSignature, err := hex.DecodeString(signature)
	if err != nil {
		return err
	}

	// Compute expected signature
	expectedSignature := ComputeSignature(payload, timestamp, secret)

	// Compare signatures
	if !hmac.Equal(decodedSignature, expectedSignature) {
		return ErrInvalidSignature
	}

	return nil
}

// VerifyWebhookSignature verifies the signature of a webhook request.
// It takes the request and the webhook secret.
// The request body will be preserved for later reading.
func VerifyWebhookSignature(r *http.Request, secret string) error {
	// Get headers
	signature := r.Header.Get(WebhookSignatureHeader)
	timestamp := r.Header.Get(WebhookTimestampHeader)

	// Read the body without consuming it
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	// Restore the body for later use
	r.Body = io.NopCloser(bytes.NewReader(body))

	// Verify the signature
	return VerifySignature(body, signature, timestamp, secret)
}

