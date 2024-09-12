// This file was auto-generated by Fern from our API Definition.

package client

import (
	context "context"
	accesstokens "github.com/schematichq/schematic-go/accesstokens"
	accounts "github.com/schematichq/schematic-go/accounts"
	billing "github.com/schematichq/schematic-go/billing"
	companies "github.com/schematichq/schematic-go/companies"
	components "github.com/schematichq/schematic-go/components"
	core "github.com/schematichq/schematic-go/core"
	crm "github.com/schematichq/schematic-go/crm"
	entitlements "github.com/schematichq/schematic-go/entitlements"
	events "github.com/schematichq/schematic-go/events"
	features "github.com/schematichq/schematic-go/features"
	option "github.com/schematichq/schematic-go/option"
	plangroups "github.com/schematichq/schematic-go/plangroups"
	plans "github.com/schematichq/schematic-go/plans"
	webhooks "github.com/schematichq/schematic-go/webhooks"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	Accounts     *accounts.Client
	Features     *features.Client
	Billing      *billing.Client
	Companies    *companies.Client
	Entitlements *entitlements.Client
	Components   *components.Client
	Crm          *crm.Client
	Events       *events.Client
	Plans        *plans.Client
	Plangroups   *plangroups.Client
	Accesstokens *accesstokens.Client
	Webhooks     *webhooks.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header:       options.ToHeader(),
		Accounts:     accounts.NewClient(opts...),
		Features:     features.NewClient(opts...),
		Billing:      billing.NewClient(opts...),
		Companies:    companies.NewClient(opts...),
		Entitlements: entitlements.NewClient(opts...),
		Components:   components.NewClient(opts...),
		Crm:          crm.NewClient(opts...),
		Events:       events.NewClient(opts...),
		Plans:        plans.NewClient(opts...),
		Plangroups:   plangroups.NewClient(opts...),
		Accesstokens: accesstokens.NewClient(opts...),
		Webhooks:     webhooks.NewClient(opts...),
	}
}

func (c *Client) GetCompanyPlans(
	ctx context.Context,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.schematichq.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/company-plans"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodGet,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
		},
	); err != nil {
		return err
	}
	return nil
}
