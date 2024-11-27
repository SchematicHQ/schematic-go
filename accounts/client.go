// This file was auto-generated by Fern from our API Definition.

package accounts

import (
	context "context"
	schematicgo "github.com/schematichq/schematic-go"
	core "github.com/schematichq/schematic-go/core"
	internal "github.com/schematichq/schematic-go/internal"
	option "github.com/schematichq/schematic-go/option"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

func (c *Client) ListAPIKeys(
	ctx context.Context,
	request *schematicgo.ListAPIKeysRequest,
	opts ...option.RequestOption,
) (*schematicgo.ListAPIKeysResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.schematichq.com",
	)
	endpointURL := baseURL + "/api-keys"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &schematicgo.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &schematicgo.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &schematicgo.ForbiddenError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &schematicgo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *schematicgo.ListAPIKeysResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) CreateAPIKey(
	ctx context.Context,
	request *schematicgo.CreateAPIKeyRequestBody,
	opts ...option.RequestOption,
) (*schematicgo.CreateAPIKeyResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.schematichq.com",
	)
	endpointURL := baseURL + "/api-keys"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &schematicgo.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &schematicgo.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &schematicgo.ForbiddenError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &schematicgo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *schematicgo.CreateAPIKeyResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetAPIKey(
	ctx context.Context,
	// api_key_id
	apiKeyID string,
	opts ...option.RequestOption,
) (*schematicgo.GetAPIKeyResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.schematichq.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api-keys/%v",
		apiKeyID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		401: func(apiError *core.APIError) error {
			return &schematicgo.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &schematicgo.ForbiddenError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &schematicgo.NotFoundError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &schematicgo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *schematicgo.GetAPIKeyResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) UpdateAPIKey(
	ctx context.Context,
	// api_key_id
	apiKeyID string,
	request *schematicgo.UpdateAPIKeyRequestBody,
	opts ...option.RequestOption,
) (*schematicgo.UpdateAPIKeyResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.schematichq.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api-keys/%v",
		apiKeyID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &schematicgo.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &schematicgo.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &schematicgo.ForbiddenError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &schematicgo.NotFoundError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &schematicgo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *schematicgo.UpdateAPIKeyResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPut,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) DeleteAPIKey(
	ctx context.Context,
	// api_key_id
	apiKeyID string,
	opts ...option.RequestOption,
) (*schematicgo.DeleteAPIKeyResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.schematichq.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api-keys/%v",
		apiKeyID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &schematicgo.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &schematicgo.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &schematicgo.ForbiddenError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &schematicgo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *schematicgo.DeleteAPIKeyResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) CountAPIKeys(
	ctx context.Context,
	request *schematicgo.CountAPIKeysRequest,
	opts ...option.RequestOption,
) (*schematicgo.CountAPIKeysResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.schematichq.com",
	)
	endpointURL := baseURL + "/api-keys/count"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &schematicgo.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &schematicgo.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &schematicgo.ForbiddenError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &schematicgo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *schematicgo.CountAPIKeysResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) ListAPIRequests(
	ctx context.Context,
	request *schematicgo.ListAPIRequestsRequest,
	opts ...option.RequestOption,
) (*schematicgo.ListAPIRequestsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.schematichq.com",
	)
	endpointURL := baseURL + "/api-requests"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &schematicgo.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &schematicgo.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &schematicgo.ForbiddenError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &schematicgo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *schematicgo.ListAPIRequestsResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetAPIRequest(
	ctx context.Context,
	// api_request_id
	apiRequestID string,
	opts ...option.RequestOption,
) (*schematicgo.GetAPIRequestResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.schematichq.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api-requests/%v",
		apiRequestID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		401: func(apiError *core.APIError) error {
			return &schematicgo.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &schematicgo.ForbiddenError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &schematicgo.NotFoundError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &schematicgo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *schematicgo.GetAPIRequestResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) CountAPIRequests(
	ctx context.Context,
	request *schematicgo.CountAPIRequestsRequest,
	opts ...option.RequestOption,
) (*schematicgo.CountAPIRequestsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.schematichq.com",
	)
	endpointURL := baseURL + "/api-requests/count"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &schematicgo.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &schematicgo.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &schematicgo.ForbiddenError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &schematicgo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *schematicgo.CountAPIRequestsResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) ListEnvironments(
	ctx context.Context,
	request *schematicgo.ListEnvironmentsRequest,
	opts ...option.RequestOption,
) (*schematicgo.ListEnvironmentsResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.schematichq.com",
	)
	endpointURL := baseURL + "/environments"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &schematicgo.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &schematicgo.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &schematicgo.ForbiddenError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &schematicgo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *schematicgo.ListEnvironmentsResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) CreateEnvironment(
	ctx context.Context,
	request *schematicgo.CreateEnvironmentRequestBody,
	opts ...option.RequestOption,
) (*schematicgo.CreateEnvironmentResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.schematichq.com",
	)
	endpointURL := baseURL + "/environments"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &schematicgo.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &schematicgo.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &schematicgo.ForbiddenError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &schematicgo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *schematicgo.CreateEnvironmentResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetEnvironment(
	ctx context.Context,
	// environment_id
	environmentID string,
	opts ...option.RequestOption,
) (*schematicgo.GetEnvironmentResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.schematichq.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/environments/%v",
		environmentID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		401: func(apiError *core.APIError) error {
			return &schematicgo.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &schematicgo.ForbiddenError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &schematicgo.NotFoundError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &schematicgo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *schematicgo.GetEnvironmentResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) UpdateEnvironment(
	ctx context.Context,
	// environment_id
	environmentID string,
	request *schematicgo.UpdateEnvironmentRequestBody,
	opts ...option.RequestOption,
) (*schematicgo.UpdateEnvironmentResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.schematichq.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/environments/%v",
		environmentID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &schematicgo.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &schematicgo.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &schematicgo.ForbiddenError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &schematicgo.NotFoundError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &schematicgo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *schematicgo.UpdateEnvironmentResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPut,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) DeleteEnvironment(
	ctx context.Context,
	// environment_id
	environmentID string,
	opts ...option.RequestOption,
) (*schematicgo.DeleteEnvironmentResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.schematichq.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/environments/%v",
		environmentID,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &schematicgo.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &schematicgo.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &schematicgo.ForbiddenError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &schematicgo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *schematicgo.DeleteEnvironmentResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
