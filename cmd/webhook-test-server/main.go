// Webhook Test Server
//
// This is a testing utility to verify Schematic webhook signatures. It exposes
// a simple HTTP server that logs and validates incoming webhook requests.
//
// Usage:
//
//  1. Set up:
//     - Export your webhook secret: export SCHEMATIC_WEBHOOK_SECRET="your-webhook-secret"
//     - Optionally set a custom port: export PORT=9000 (default: 8080)
//     - Build and run: go run cmd/webhook-test-server/main.go
//
//  2. Expose to the internet:
//     - Since Schematic can't access localhost directly, use a tunneling service like ngrok:
//     ngrok http 8080
//
//  3. Configure in Schematic:
//     - Use the ngrok URL (e.g., https://a1b2c3d4.ngrok.io/webhook) as your webhook endpoint
//     - Set the same webhook secret in Schematic that you used in SCHEMATIC_WEBHOOK_SECRET
//
//  4. Test:
//     - Trigger a webhook from Schematic
//     - This server will verify the signature and output the payload
//
// This tool helps validate the webhook signature verification implementation before
// integrating it into production applications.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/schematichq/schematic-go/webhooks"
)

func main() {
	// Get webhook secret from environment variable
	webhookSecret := os.Getenv("SCHEMATIC_WEBHOOK_SECRET")
	if webhookSecret == "" {
		log.Fatal("SCHEMATIC_WEBHOOK_SECRET environment variable is required")
	}

	// Define webhook handler
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received webhook request")

		// Verify signature
		err := webhooks.VerifyWebhookSignature(r, webhookSecret)
		if err != nil {
			log.Printf("Signature verification failed: %v", err)
			http.Error(w, fmt.Sprintf("Signature verification failed: %v", err), http.StatusUnauthorized)
			return
		}

		log.Println("Signature verification successful")

		// Read and print request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read request body: %v", err)
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		// Log headers
		log.Println("Headers:")
		for name, values := range r.Header {
			for _, value := range values {
				log.Printf("  %s: %s", name, value)
			}
		}

		// Try to pretty-print JSON
		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, body, "", "  ")
		if err == nil {
			log.Printf("Webhook payload:\n%s", prettyJSON.String())
		} else {
			// If not valid JSON, print raw body
			log.Printf("Webhook payload (raw):\n%s", string(body))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Webhook received and verified successfully"))
	})

	// Start server
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	log.Printf("Starting webhook test server on port %s", port)
	log.Printf("Local URL: http://localhost:%s/webhook", port)
	log.Printf("Remember: Use a tool like ngrok to expose this server to the internet")
	log.Printf("Example: ngrok http %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
