package core

import (
	fmt "fmt"
	http "net/http"
	url "net/url"
	"time"

	cache "github.com/schematichq/schematic-go/cache"
)

// RequestOption adapts the behavior of the client or an individual request.
type RequestOption interface {
	applyRequestOptions(*RequestOptions)
}

// RequestOptions defines all of the possible request options.
//
// This type is primarily used by the generated code and is not meant
// to be used directly; use the option package instead.
type RequestOptions struct {
	BaseURL         string
	HTTPClient      HTTPClient
	HTTPHeader      http.Header
	BodyProperties  map[string]interface{}
	QueryParameters url.Values
	MaxAttempts     uint
	APIKey          string

	// Schematic custom request option fields
	EventBufferPeriod       *time.Duration
	FlagCheckCacheProviders []cache.CacheProvider[bool]
	FlagDefaults            map[string]bool
	OfflineMode             bool
}

// NewRequestOptions returns a new *RequestOptions value.
//
// This function is primarily used by the generated code and is not meant
// to be used directly; use RequestOption instead.
func NewRequestOptions(opts ...RequestOption) *RequestOptions {
	options := &RequestOptions{
		HTTPHeader:      make(http.Header),
		BodyProperties:  make(map[string]interface{}),
		QueryParameters: make(url.Values),
	}
	for _, opt := range opts {
		opt.applyRequestOptions(options)
	}
	return options
}

// ToHeader maps the configured request options into a http.Header used
// for the request(s).
func (r *RequestOptions) ToHeader() http.Header {
	header := r.cloneHeader()
	if r.APIKey != "" {
		header.Set("X-Schematic-Api-Key", fmt.Sprintf("%v", r.APIKey))
	}
	return header
}

func (r *RequestOptions) cloneHeader() http.Header {
	headers := r.HTTPHeader.Clone()
	headers.Set("X-Fern-Language", "Go")
	headers.Set("X-Fern-SDK-Name", "github.com/schematichq/schematic-go")
	headers.Set("X-Fern-SDK-Version", "v1.1.1")
	return headers
}

// BaseURLOption implements the RequestOption interface.
type BaseURLOption struct {
	BaseURL string
}

func (b *BaseURLOption) applyRequestOptions(opts *RequestOptions) {
	opts.BaseURL = b.BaseURL
}

// HTTPClientOption implements the RequestOption interface.
type HTTPClientOption struct {
	HTTPClient HTTPClient
}

func (h *HTTPClientOption) applyRequestOptions(opts *RequestOptions) {
	opts.HTTPClient = h.HTTPClient
}

// HTTPHeaderOption implements the RequestOption interface.
type HTTPHeaderOption struct {
	HTTPHeader http.Header
}

func (h *HTTPHeaderOption) applyRequestOptions(opts *RequestOptions) {
	opts.HTTPHeader = h.HTTPHeader
}

// BodyPropertiesOption implements the RequestOption interface.
type BodyPropertiesOption struct {
	BodyProperties map[string]interface{}
}

func (b *BodyPropertiesOption) applyRequestOptions(opts *RequestOptions) {
	opts.BodyProperties = b.BodyProperties
}

// QueryParametersOption implements the RequestOption interface.
type QueryParametersOption struct {
	QueryParameters url.Values
}

func (q *QueryParametersOption) applyRequestOptions(opts *RequestOptions) {
	opts.QueryParameters = q.QueryParameters
}

// MaxAttemptsOption implements the RequestOption interface.
type MaxAttemptsOption struct {
	MaxAttempts uint
}

func (m *MaxAttemptsOption) applyRequestOptions(opts *RequestOptions) {
	opts.MaxAttempts = m.MaxAttempts
}

// APIKeyOption implements the RequestOption interface.
type APIKeyOption struct {
	APIKey string
}

func (a *APIKeyOption) applyRequestOptions(opts *RequestOptions) {
	opts.APIKey = a.APIKey
}
