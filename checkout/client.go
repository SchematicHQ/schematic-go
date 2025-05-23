// This file was auto-generated by Fern from our API Definition.

package checkout

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

func (c *Client) Internal(
	ctx context.Context,
	request *schematicgo.ChangeSubscriptionInternalRequestBody,
	opts ...option.RequestOption,
) (*schematicgo.CheckoutInternalResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.schematichq.com",
	)
	endpointURL := baseURL + "/checkout-internal"
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

	var response *schematicgo.CheckoutInternalResponse
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

func (c *Client) GetCheckoutData(
	ctx context.Context,
	request *schematicgo.CheckoutDataRequestBody,
	opts ...option.RequestOption,
) (*schematicgo.GetCheckoutDataResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.schematichq.com",
	)
	endpointURL := baseURL + "/checkout-internal/data"
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

	var response *schematicgo.GetCheckoutDataResponse
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

func (c *Client) PreviewCheckoutInternal(
	ctx context.Context,
	request *schematicgo.ChangeSubscriptionInternalRequestBody,
	opts ...option.RequestOption,
) (*schematicgo.PreviewCheckoutInternalResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.schematichq.com",
	)
	endpointURL := baseURL + "/checkout-internal/preview"
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

	var response *schematicgo.PreviewCheckoutInternalResponse
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

func (c *Client) UpdateCustomerSubscriptionTrialEnd(
	ctx context.Context,
	// subscription_id
	subscriptionID string,
	request *schematicgo.UpdateTrialEndRequestBody,
	opts ...option.RequestOption,
) (*schematicgo.UpdateCustomerSubscriptionTrialEndResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.schematichq.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/subscription/%v/edit-trial-end",
		subscriptionID,
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

	var response *schematicgo.UpdateCustomerSubscriptionTrialEndResponse
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
