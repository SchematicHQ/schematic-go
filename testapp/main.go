// SDK E2E Test App for schematic-go.
//
// A lightweight HTTP server implementing the standard test app contract
// defined in the SDK spec. The E2E harness (SchematicHQ/actions/sdk-e2e)
// calls POST /configure after startup to pass an env-scoped API key,
// then runs assertions against the other endpoints.
//
// Usage:
//
//	go run ./testapp
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	schematic "github.com/schematichq/schematic-go"
	"github.com/redis/go-redis/v9"
	"github.com/schematichq/schematic-go/cache"
	schematicclient "github.com/schematichq/schematic-go/client"
	"github.com/schematichq/schematic-go/core"
	"github.com/schematichq/schematic-go/option"
)

const cacheTTL = 2 * time.Second // Short TTL for E2E — tests sleep past this to verify cache expiration

// --- State ---

var (
	mu            sync.RWMutex
	client        *schematicclient.SchematicClient
	currentConfig map[string]any
)

// --- Helpers ---

func readJSON(r *http.Request, v any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if len(body) == 0 {
		return nil
	}
	return json.Unmarshal(body, v)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func getClient() *schematicclient.SchematicClient {
	mu.RLock()
	defer mu.RUnlock()
	return client
}

func getConfigString(key string) string {
	if currentConfig == nil {
		return ""
	}
	v, _ := currentConfig[key].(string)
	return v
}

func getConfigBool(key string) bool {
	if currentConfig == nil {
		return false
	}
	v, _ := currentConfig[key].(bool)
	return v
}

// --- Route handlers ---

func handleConfigure(w http.ResponseWriter, r *http.Request) {
	var config map[string]any
	if err := readJSON(r, &config); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"success": false, "error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// Close existing client
	if client != nil {
		client.Close()
	}

	currentConfig = config

	apiKey := getConfigString("apiKey")
	opts := []option.RequestOption{
		option.WithAPIKey(apiKey),
	}

	if baseURL := getConfigString("baseUrl"); baseURL != "" {
		opts = append(opts, option.WithBaseURL(baseURL))
	}

	// Flag defaults
	if rawDefaults, ok := config["flagDefaults"].(map[string]any); ok {
		defaults := make(map[string]bool, len(rawDefaults))
		for k, v := range rawDefaults {
			if b, ok := v.(bool); ok {
				defaults[k] = b
			}
		}
		opts = append(opts, core.WithDefaultFlagValues(defaults))
	}

	// Offline mode
	if getConfigBool("offline") {
		opts = append(opts, core.WithOfflineMode())
	}

	// Cache configuration
	redisAddr := strings.TrimPrefix(getConfigString("redisUrl"), "redis://")
	if getConfigBool("noCache") {
		opts = append(opts, core.WithDisableFlagCheckCache())
	} else if redisAddr != "" && !getConfigBool("useDataStream") {
		// Redis cache without DataStream — create a Redis cache provider directly
		redisClient := redis.NewClient(&redis.Options{Addr: redisAddr})
		redisCache := cache.NewRedisCache[*core.CheckFlagResponse](redisClient, cacheTTL)
		opts = append(opts, core.WithFlagCheckCacheProvider(redisCache))
	}

	// DataStream / Replicator
	if getConfigBool("useDataStream") {
		var dsOpts []core.DatastreamOption

		if redisAddr != "" {
			dsOpts = append(dsOpts, core.WithRedisCache(core.RedisCacheConfig{
				Addr: redisAddr,
			}))
		}

		if replicatorURL := getConfigString("replicatorUrl"); replicatorURL != "" {
			dsOpts = append(dsOpts, core.WithReplicatorMode())
			dsOpts = append(dsOpts, core.WithReplicatorHealthURL(replicatorURL+"/ready"))
		}

		opts = append(opts, core.WithDatastream(dsOpts...))
	}

	client = schematicclient.NewSchematicClient(opts...)

	writeJSON(w, http.StatusOK, map[string]any{"success": true, "cacheTtlMs": cacheTTL.Milliseconds()})
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	c := getClient()
	status := "waiting"
	if c != nil {
		status = "configured"
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"status": status,
		"config": map[string]any{
			"offline":        getConfigBool("offline"),
			"useDataStream":  getConfigBool("useDataStream"),
			"cacheTtlMs":     cacheTTL.Milliseconds(),
		},
	})
}

func handleCheckFlag(w http.ResponseWriter, r *http.Request) {
	c := getClient()
	if c == nil {
		writeJSON(w, http.StatusServiceUnavailable, map[string]any{"error": "not configured"})
		return
	}

	var body map[string]any
	if err := readJSON(r, &body); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	flagKey, _ := body["flagKey"].(string)
	evalCtx := &schematic.CheckFlagRequestBody{}
	if company := getMapFromBody(body, "company"); company != nil {
		evalCtx.Company = company
	}
	if user := getMapFromBody(body, "user"); user != nil {
		evalCtx.User = user
	}

	value := c.CheckFlag(r.Context(), evalCtx, flagKey)
	writeJSON(w, http.StatusOK, map[string]any{"value": value})
}

func handleIdentify(w http.ResponseWriter, r *http.Request) {
	c := getClient()
	if c == nil {
		writeJSON(w, http.StatusServiceUnavailable, map[string]any{"error": "not configured"})
		return
	}

	var body map[string]any
	if err := readJSON(r, &body); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	// Translate E2E contract to Go SDK's EventBodyIdentify:
	//   E2E:  { company: {k:v}, user: {k:v}, keys: {k:v} }
	//   SDK:  { Keys: {k:v}, Company: { Keys: {k:v} } }
	identifyBody := &schematic.EventBodyIdentify{}
	if keys := getMapFromBody(body, "keys"); keys != nil {
		identifyBody.Keys = keys
	} else if user := getMapFromBody(body, "user"); user != nil {
		identifyBody.Keys = user
	}
	if company := getMapFromBody(body, "company"); company != nil {
		identifyBody.Company = &schematic.EventBodyIdentifyCompany{
			Keys: company,
		}
	}

	c.Identify(r.Context(), identifyBody)
	writeJSON(w, http.StatusOK, map[string]any{"success": true})
}

func handleTrack(w http.ResponseWriter, r *http.Request) {
	c := getClient()
	if c == nil {
		writeJSON(w, http.StatusServiceUnavailable, map[string]any{"error": "not configured"})
		return
	}

	var body map[string]any
	if err := readJSON(r, &body); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	trackBody := &schematic.EventBodyTrack{}
	if event, ok := body["event"].(string); ok {
		trackBody.Event = event
	}
	if company := getMapFromBody(body, "company"); company != nil {
		trackBody.Company = company
	}
	if user := getMapFromBody(body, "user"); user != nil {
		trackBody.User = user
	}
	if quantity, ok := body["quantity"].(float64); ok {
		q := int(quantity)
		trackBody.Quantity = &q
	}

	c.Track(r.Context(), trackBody)
	writeJSON(w, http.StatusOK, map[string]any{"success": true})
}

func handleSetFlagDefault(w http.ResponseWriter, r *http.Request) {
	c := getClient()
	if c == nil {
		writeJSON(w, http.StatusServiceUnavailable, map[string]any{"error": "not configured"})
		return
	}

	var body struct {
		FlagKey string `json:"flagKey"`
		Value   bool   `json:"value"`
	}
	if err := readJSON(r, &body); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	c.SetFlagDefault(body.FlagKey, body.Value)
	writeJSON(w, http.StatusOK, map[string]any{"success": true})
}

// getMapFromBody extracts a map[string]string from a JSON body field.
func getMapFromBody(body map[string]any, key string) map[string]string {
	raw, ok := body[key].(map[string]any)
	if !ok {
		return nil
	}
	result := make(map[string]string, len(raw))
	for k, v := range raw {
		if s, ok := v.(string); ok {
			result[k] = s
		}
	}
	return result
}

// --- Server ---

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handleHealth)
	mux.HandleFunc("POST /configure", handleConfigure)
	mux.HandleFunc("POST /check-flag", handleCheckFlag)
	mux.HandleFunc("POST /identify", handleIdentify)
	mux.HandleFunc("POST /track", handleTrack)
	mux.HandleFunc("POST /set-flag-default", handleSetFlagDefault)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// Graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt)
		<-sigCh
		log.Println("Shutting down...")
		mu.RLock()
		if client != nil {
			client.Close()
		}
		mu.RUnlock()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()

	fmt.Printf("SDK E2E test app listening on http://localhost:%s\n", port)
	fmt.Println("Waiting for POST /configure to initialize SchematicClient...")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
