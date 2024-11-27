// This file was auto-generated by Fern from our API Definition.

package accesstokens

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

func (c *Client) IssueTemporaryAccessToken(
	ctx context.Context,
	request *schematicgo.IssueTemporaryAccessTokenRequestBody,
	opts ...option.RequestOption,
) (*schematicgo.IssueTemporaryAccessTokenResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.schematichq.com",
	)
	endpointURL := baseURL + "/temporary-access-tokens"
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

	var response *schematicgo.IssueTemporaryAccessTokenResponse
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
