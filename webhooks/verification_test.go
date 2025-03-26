package webhooks_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/schematichq/schematic-go/webhooks"
)

func TestVerifySignature(t *testing.T) {
	tests := []struct {
		name      string
		payload   []byte
		signature string
		timestamp string
		secret    string
		wantErr   bool
	}{
		{
			name:      "valid signature",
			payload:   []byte(`{"event":"test"}`),
			timestamp: "1609459200",
			secret:    "test-secret",
			signature: "", // Will be computed
			wantErr:   false,
		},
		{
			name:      "invalid signature",
			payload:   []byte(`{"event":"test"}`),
			timestamp: "1609459200",
			secret:    "test-secret",
			signature: "invalid",
			wantErr:   true,
		},
		{
			name:      "missing signature",
			payload:   []byte(`{"event":"test"}`),
			timestamp: "1609459200",
			secret:    "test-secret",
			signature: "",
			wantErr:   true,
		},
		{
			name:      "missing timestamp",
			payload:   []byte(`{"event":"test"}`),
			timestamp: "",
			secret:    "test-secret",
			signature: "abcdef",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// If signature is empty, compute a valid one for this test
			if tt.signature == "" && tt.name == "valid signature" {
				tt.signature = webhooks.ComputeHexSignature(tt.payload, tt.timestamp, tt.secret)
			}

			err := webhooks.VerifySignature(tt.payload, tt.signature, tt.timestamp, tt.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifySignature() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVerifyWebhookSignature(t *testing.T) {
	payload := []byte(`{"event":"test"}`)
	timestamp := "1609459200"
	secret := "test-secret"
	signature := webhooks.ComputeHexSignature(payload, timestamp, secret)

	tests := []struct {
		name     string
		setupReq func() *http.Request
		wantErr  bool
	}{
		{
			name: "valid signature",
			setupReq: func() *http.Request {
				req, _ := http.NewRequest("POST", "https://example.com/webhook", bytes.NewReader(payload))
				req.Header.Set(webhooks.WebhookSignatureHeader, signature)
				req.Header.Set(webhooks.WebhookTimestampHeader, timestamp)
				return req
			},
			wantErr: false,
		},
		{
			name: "missing signature header",
			setupReq: func() *http.Request {
				req, _ := http.NewRequest("POST", "https://example.com/webhook", bytes.NewReader(payload))
				req.Header.Set(webhooks.WebhookTimestampHeader, timestamp)
				return req
			},
			wantErr: true,
		},
		{
			name: "missing timestamp header",
			setupReq: func() *http.Request {
				req, _ := http.NewRequest("POST", "https://example.com/webhook", bytes.NewReader(payload))
				req.Header.Set(webhooks.WebhookSignatureHeader, signature)
				return req
			},
			wantErr: true,
		},
		{
			name: "invalid signature",
			setupReq: func() *http.Request {
				req, _ := http.NewRequest("POST", "https://example.com/webhook", bytes.NewReader(payload))
				req.Header.Set(webhooks.WebhookSignatureHeader, "invalid")
				req.Header.Set(webhooks.WebhookTimestampHeader, timestamp)
				return req
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.setupReq()
			
			// Save the original body for verification
			var originalBody []byte
			if req.Body != nil {
				originalBody, _ = io.ReadAll(req.Body)
				req.Body = io.NopCloser(bytes.NewReader(originalBody))
			}

			err := webhooks.VerifyWebhookSignature(req, secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyWebhookSignature() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Verify that body is still available after signature verification
			if req.Body != nil {
				bodyAfterVerify, _ := io.ReadAll(req.Body)
				if !bytes.Equal(bodyAfterVerify, originalBody) {
					t.Errorf("Body was not preserved correctly after verification")
				}
			}
		})
	}
}