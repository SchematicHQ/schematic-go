// This file was auto-generated by Fern from our API Definition.

package schematichq

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/schematichq/schematic-go/core"
)

type CountCompanyOverridesRequest struct {
	CompanyID  *string   `json:"-" url:"company_id,omitempty"`
	CompanyIDs []*string `json:"-" url:"company_ids,omitempty"`
	FeatureID  *string   `json:"-" url:"feature_id,omitempty"`
	FeatureIDs []*string `json:"-" url:"feature_ids,omitempty"`
	IDs        []*string `json:"-" url:"ids,omitempty"`
	Q          *string   `json:"-" url:"q,omitempty"`
	// Page limit (default 100)
	Limit *int `json:"-" url:"limit,omitempty"`
	// Page offset (default 0)
	Offset *int `json:"-" url:"offset,omitempty"`
}

type CountFeatureCompaniesRequest struct {
	FeatureID string  `json:"-" url:"feature_id"`
	Q         *string `json:"-" url:"q,omitempty"`
	// Page limit (default 100)
	Limit *int `json:"-" url:"limit,omitempty"`
	// Page offset (default 0)
	Offset *int `json:"-" url:"offset,omitempty"`
}

type CountFeatureUsageRequest struct {
	CompanyID   *string            `json:"-" url:"company_id,omitempty"`
	CompanyKeys map[string]*string `json:"-" url:"company_keys,omitempty"`
	FeatureIDs  []*string          `json:"-" url:"feature_ids,omitempty"`
	Q           *string            `json:"-" url:"q,omitempty"`
	// Page limit (default 100)
	Limit *int `json:"-" url:"limit,omitempty"`
	// Page offset (default 0)
	Offset *int `json:"-" url:"offset,omitempty"`
}

type CountFeatureUsersRequest struct {
	FeatureID string  `json:"-" url:"feature_id"`
	Q         *string `json:"-" url:"q,omitempty"`
	// Page limit (default 100)
	Limit *int `json:"-" url:"limit,omitempty"`
	// Page offset (default 0)
	Offset *int `json:"-" url:"offset,omitempty"`
}

type CountPlanEntitlementsRequest struct {
	FeatureID  *string   `json:"-" url:"feature_id,omitempty"`
	FeatureIDs []*string `json:"-" url:"feature_ids,omitempty"`
	IDs        []*string `json:"-" url:"ids,omitempty"`
	PlanID     *string   `json:"-" url:"plan_id,omitempty"`
	PlanIDs    []*string `json:"-" url:"plan_ids,omitempty"`
	Q          *string   `json:"-" url:"q,omitempty"`
	// Page limit (default 100)
	Limit *int `json:"-" url:"limit,omitempty"`
	// Page offset (default 0)
	Offset *int `json:"-" url:"offset,omitempty"`
}

type CreateCompanyOverrideRequestBody struct {
	CompanyID    string                                        `json:"company_id" url:"-"`
	FeatureID    string                                        `json:"feature_id" url:"-"`
	MetricPeriod *CreateCompanyOverrideRequestBodyMetricPeriod `json:"metric_period,omitempty" url:"-"`
	ValueBool    *bool                                         `json:"value_bool,omitempty" url:"-"`
	ValueNumeric *int                                          `json:"value_numeric,omitempty" url:"-"`
	ValueTraitID *string                                       `json:"value_trait_id,omitempty" url:"-"`
	ValueType    CreateCompanyOverrideRequestBodyValueType     `json:"value_type" url:"-"`
}

type CreatePlanEntitlementRequestBody struct {
	FeatureID    string                                        `json:"feature_id" url:"-"`
	MetricPeriod *CreatePlanEntitlementRequestBodyMetricPeriod `json:"metric_period,omitempty" url:"-"`
	PlanID       string                                        `json:"plan_id" url:"-"`
	ValueBool    *bool                                         `json:"value_bool,omitempty" url:"-"`
	ValueNumeric *int                                          `json:"value_numeric,omitempty" url:"-"`
	ValueTraitID *string                                       `json:"value_trait_id,omitempty" url:"-"`
	ValueType    CreatePlanEntitlementRequestBodyValueType     `json:"value_type" url:"-"`
}

type GetFeatureUsageByCompanyRequest struct {
	// Key/value pairs
	Keys map[string]interface{} `json:"-" url:"keys,omitempty"`
}

type ListCompanyOverridesRequest struct {
	CompanyID  *string   `json:"-" url:"company_id,omitempty"`
	CompanyIDs []*string `json:"-" url:"company_ids,omitempty"`
	FeatureID  *string   `json:"-" url:"feature_id,omitempty"`
	FeatureIDs []*string `json:"-" url:"feature_ids,omitempty"`
	IDs        []*string `json:"-" url:"ids,omitempty"`
	Q          *string   `json:"-" url:"q,omitempty"`
	// Page limit (default 100)
	Limit *int `json:"-" url:"limit,omitempty"`
	// Page offset (default 0)
	Offset *int `json:"-" url:"offset,omitempty"`
}

type ListFeatureCompaniesRequest struct {
	FeatureID string  `json:"-" url:"feature_id"`
	Q         *string `json:"-" url:"q,omitempty"`
	// Page limit (default 100)
	Limit *int `json:"-" url:"limit,omitempty"`
	// Page offset (default 0)
	Offset *int `json:"-" url:"offset,omitempty"`
}

type ListFeatureUsageRequest struct {
	CompanyID   *string            `json:"-" url:"company_id,omitempty"`
	CompanyKeys map[string]*string `json:"-" url:"company_keys,omitempty"`
	FeatureIDs  []*string          `json:"-" url:"feature_ids,omitempty"`
	Q           *string            `json:"-" url:"q,omitempty"`
	// Page limit (default 100)
	Limit *int `json:"-" url:"limit,omitempty"`
	// Page offset (default 0)
	Offset *int `json:"-" url:"offset,omitempty"`
}

type ListFeatureUsersRequest struct {
	FeatureID string  `json:"-" url:"feature_id"`
	Q         *string `json:"-" url:"q,omitempty"`
	// Page limit (default 100)
	Limit *int `json:"-" url:"limit,omitempty"`
	// Page offset (default 0)
	Offset *int `json:"-" url:"offset,omitempty"`
}

type ListPlanEntitlementsRequest struct {
	FeatureID  *string   `json:"-" url:"feature_id,omitempty"`
	FeatureIDs []*string `json:"-" url:"feature_ids,omitempty"`
	IDs        []*string `json:"-" url:"ids,omitempty"`
	PlanID     *string   `json:"-" url:"plan_id,omitempty"`
	PlanIDs    []*string `json:"-" url:"plan_ids,omitempty"`
	Q          *string   `json:"-" url:"q,omitempty"`
	// Page limit (default 100)
	Limit *int `json:"-" url:"limit,omitempty"`
	// Page offset (default 0)
	Offset *int `json:"-" url:"offset,omitempty"`
}

type CountCompanyOverridesResponse struct {
	Data *CountResponse `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params *CountCompanyOverridesParams `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CountCompanyOverridesResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CountCompanyOverridesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CountCompanyOverridesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CountCompanyOverridesResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CountCompanyOverridesResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CountFeatureCompaniesResponse struct {
	Data *CountResponse `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params *CountFeatureCompaniesParams `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CountFeatureCompaniesResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CountFeatureCompaniesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CountFeatureCompaniesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CountFeatureCompaniesResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CountFeatureCompaniesResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CountFeatureUsageResponse struct {
	Data *CountResponse `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params *CountFeatureUsageParams `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CountFeatureUsageResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CountFeatureUsageResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CountFeatureUsageResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CountFeatureUsageResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CountFeatureUsageResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CountFeatureUsersResponse struct {
	Data *CountResponse `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params *CountFeatureUsersParams `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CountFeatureUsersResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CountFeatureUsersResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CountFeatureUsersResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CountFeatureUsersResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CountFeatureUsersResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CountPlanEntitlementsResponse struct {
	Data *CountResponse `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params *CountPlanEntitlementsParams `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CountPlanEntitlementsResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CountPlanEntitlementsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CountPlanEntitlementsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CountPlanEntitlementsResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CountPlanEntitlementsResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CreateCompanyOverrideRequestBodyMetricPeriod string

const (
	CreateCompanyOverrideRequestBodyMetricPeriodCurrentMonth CreateCompanyOverrideRequestBodyMetricPeriod = "current_month"
	CreateCompanyOverrideRequestBodyMetricPeriodCurrentWeek  CreateCompanyOverrideRequestBodyMetricPeriod = "current_week"
	CreateCompanyOverrideRequestBodyMetricPeriodCurrentDay   CreateCompanyOverrideRequestBodyMetricPeriod = "current_day"
)

func NewCreateCompanyOverrideRequestBodyMetricPeriodFromString(s string) (CreateCompanyOverrideRequestBodyMetricPeriod, error) {
	switch s {
	case "current_month":
		return CreateCompanyOverrideRequestBodyMetricPeriodCurrentMonth, nil
	case "current_week":
		return CreateCompanyOverrideRequestBodyMetricPeriodCurrentWeek, nil
	case "current_day":
		return CreateCompanyOverrideRequestBodyMetricPeriodCurrentDay, nil
	}
	var t CreateCompanyOverrideRequestBodyMetricPeriod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CreateCompanyOverrideRequestBodyMetricPeriod) Ptr() *CreateCompanyOverrideRequestBodyMetricPeriod {
	return &c
}

type CreateCompanyOverrideRequestBodyValueType string

const (
	CreateCompanyOverrideRequestBodyValueTypeBoolean   CreateCompanyOverrideRequestBodyValueType = "boolean"
	CreateCompanyOverrideRequestBodyValueTypeNumeric   CreateCompanyOverrideRequestBodyValueType = "numeric"
	CreateCompanyOverrideRequestBodyValueTypeTrait     CreateCompanyOverrideRequestBodyValueType = "trait"
	CreateCompanyOverrideRequestBodyValueTypeUnlimited CreateCompanyOverrideRequestBodyValueType = "unlimited"
)

func NewCreateCompanyOverrideRequestBodyValueTypeFromString(s string) (CreateCompanyOverrideRequestBodyValueType, error) {
	switch s {
	case "boolean":
		return CreateCompanyOverrideRequestBodyValueTypeBoolean, nil
	case "numeric":
		return CreateCompanyOverrideRequestBodyValueTypeNumeric, nil
	case "trait":
		return CreateCompanyOverrideRequestBodyValueTypeTrait, nil
	case "unlimited":
		return CreateCompanyOverrideRequestBodyValueTypeUnlimited, nil
	}
	var t CreateCompanyOverrideRequestBodyValueType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CreateCompanyOverrideRequestBodyValueType) Ptr() *CreateCompanyOverrideRequestBodyValueType {
	return &c
}

type CreateCompanyOverrideResponse struct {
	Data *CompanyOverrideResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CreateCompanyOverrideResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateCompanyOverrideResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateCompanyOverrideResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateCompanyOverrideResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateCompanyOverrideResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CreatePlanEntitlementRequestBodyMetricPeriod string

const (
	CreatePlanEntitlementRequestBodyMetricPeriodCurrentMonth CreatePlanEntitlementRequestBodyMetricPeriod = "current_month"
	CreatePlanEntitlementRequestBodyMetricPeriodCurrentWeek  CreatePlanEntitlementRequestBodyMetricPeriod = "current_week"
	CreatePlanEntitlementRequestBodyMetricPeriodCurrentDay   CreatePlanEntitlementRequestBodyMetricPeriod = "current_day"
)

func NewCreatePlanEntitlementRequestBodyMetricPeriodFromString(s string) (CreatePlanEntitlementRequestBodyMetricPeriod, error) {
	switch s {
	case "current_month":
		return CreatePlanEntitlementRequestBodyMetricPeriodCurrentMonth, nil
	case "current_week":
		return CreatePlanEntitlementRequestBodyMetricPeriodCurrentWeek, nil
	case "current_day":
		return CreatePlanEntitlementRequestBodyMetricPeriodCurrentDay, nil
	}
	var t CreatePlanEntitlementRequestBodyMetricPeriod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CreatePlanEntitlementRequestBodyMetricPeriod) Ptr() *CreatePlanEntitlementRequestBodyMetricPeriod {
	return &c
}

type CreatePlanEntitlementRequestBodyValueType string

const (
	CreatePlanEntitlementRequestBodyValueTypeBoolean   CreatePlanEntitlementRequestBodyValueType = "boolean"
	CreatePlanEntitlementRequestBodyValueTypeNumeric   CreatePlanEntitlementRequestBodyValueType = "numeric"
	CreatePlanEntitlementRequestBodyValueTypeTrait     CreatePlanEntitlementRequestBodyValueType = "trait"
	CreatePlanEntitlementRequestBodyValueTypeUnlimited CreatePlanEntitlementRequestBodyValueType = "unlimited"
)

func NewCreatePlanEntitlementRequestBodyValueTypeFromString(s string) (CreatePlanEntitlementRequestBodyValueType, error) {
	switch s {
	case "boolean":
		return CreatePlanEntitlementRequestBodyValueTypeBoolean, nil
	case "numeric":
		return CreatePlanEntitlementRequestBodyValueTypeNumeric, nil
	case "trait":
		return CreatePlanEntitlementRequestBodyValueTypeTrait, nil
	case "unlimited":
		return CreatePlanEntitlementRequestBodyValueTypeUnlimited, nil
	}
	var t CreatePlanEntitlementRequestBodyValueType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CreatePlanEntitlementRequestBodyValueType) Ptr() *CreatePlanEntitlementRequestBodyValueType {
	return &c
}

type CreatePlanEntitlementResponse struct {
	Data *PlanEntitlementResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CreatePlanEntitlementResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreatePlanEntitlementResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreatePlanEntitlementResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreatePlanEntitlementResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreatePlanEntitlementResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type DeleteCompanyOverrideResponse struct {
	Data *DeleteResponse `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *DeleteCompanyOverrideResponse) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeleteCompanyOverrideResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteCompanyOverrideResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteCompanyOverrideResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeleteCompanyOverrideResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type DeletePlanEntitlementResponse struct {
	Data *DeleteResponse `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *DeletePlanEntitlementResponse) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeletePlanEntitlementResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DeletePlanEntitlementResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeletePlanEntitlementResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeletePlanEntitlementResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type GetCompanyOverrideResponse struct {
	Data *CompanyOverrideResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GetCompanyOverrideResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetCompanyOverrideResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetCompanyOverrideResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetCompanyOverrideResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetCompanyOverrideResponse) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := core.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

type GetFeatureUsageByCompanyResponse struct {
	Data *FeatureUsageDetailResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params *GetFeatureUsageByCompanyParams `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GetFeatureUsageByCompanyResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetFeatureUsageByCompanyResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetFeatureUsageByCompanyResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetFeatureUsageByCompanyResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetFeatureUsageByCompanyResponse) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := core.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

type GetPlanEntitlementResponse struct {
	Data *PlanEntitlementResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GetPlanEntitlementResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetPlanEntitlementResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetPlanEntitlementResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetPlanEntitlementResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetPlanEntitlementResponse) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := core.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

type ListCompanyOverridesResponse struct {
	// The returned resources
	Data []*CompanyOverrideResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params *ListCompanyOverridesParams `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *ListCompanyOverridesResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListCompanyOverridesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListCompanyOverridesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListCompanyOverridesResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListCompanyOverridesResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type ListFeatureCompaniesResponse struct {
	// The returned resources
	Data []*FeatureCompanyResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params *ListFeatureCompaniesParams `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *ListFeatureCompaniesResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListFeatureCompaniesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListFeatureCompaniesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListFeatureCompaniesResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListFeatureCompaniesResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type ListFeatureUsageResponse struct {
	// The returned resources
	Data []*FeatureUsageResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params *ListFeatureUsageParams `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *ListFeatureUsageResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListFeatureUsageResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListFeatureUsageResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListFeatureUsageResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListFeatureUsageResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type ListFeatureUsersResponse struct {
	// The returned resources
	Data []*FeatureCompanyUserResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params *ListFeatureUsersParams `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *ListFeatureUsersResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListFeatureUsersResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListFeatureUsersResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListFeatureUsersResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListFeatureUsersResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type ListPlanEntitlementsResponse struct {
	// The returned resources
	Data []*PlanEntitlementResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params *ListPlanEntitlementsParams `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *ListPlanEntitlementsResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListPlanEntitlementsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListPlanEntitlementsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListPlanEntitlementsResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListPlanEntitlementsResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type UpdateCompanyOverrideRequestBodyMetricPeriod string

const (
	UpdateCompanyOverrideRequestBodyMetricPeriodCurrentMonth UpdateCompanyOverrideRequestBodyMetricPeriod = "current_month"
	UpdateCompanyOverrideRequestBodyMetricPeriodCurrentWeek  UpdateCompanyOverrideRequestBodyMetricPeriod = "current_week"
	UpdateCompanyOverrideRequestBodyMetricPeriodCurrentDay   UpdateCompanyOverrideRequestBodyMetricPeriod = "current_day"
)

func NewUpdateCompanyOverrideRequestBodyMetricPeriodFromString(s string) (UpdateCompanyOverrideRequestBodyMetricPeriod, error) {
	switch s {
	case "current_month":
		return UpdateCompanyOverrideRequestBodyMetricPeriodCurrentMonth, nil
	case "current_week":
		return UpdateCompanyOverrideRequestBodyMetricPeriodCurrentWeek, nil
	case "current_day":
		return UpdateCompanyOverrideRequestBodyMetricPeriodCurrentDay, nil
	}
	var t UpdateCompanyOverrideRequestBodyMetricPeriod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (u UpdateCompanyOverrideRequestBodyMetricPeriod) Ptr() *UpdateCompanyOverrideRequestBodyMetricPeriod {
	return &u
}

type UpdateCompanyOverrideRequestBodyValueType string

const (
	UpdateCompanyOverrideRequestBodyValueTypeBoolean   UpdateCompanyOverrideRequestBodyValueType = "boolean"
	UpdateCompanyOverrideRequestBodyValueTypeNumeric   UpdateCompanyOverrideRequestBodyValueType = "numeric"
	UpdateCompanyOverrideRequestBodyValueTypeTrait     UpdateCompanyOverrideRequestBodyValueType = "trait"
	UpdateCompanyOverrideRequestBodyValueTypeUnlimited UpdateCompanyOverrideRequestBodyValueType = "unlimited"
)

func NewUpdateCompanyOverrideRequestBodyValueTypeFromString(s string) (UpdateCompanyOverrideRequestBodyValueType, error) {
	switch s {
	case "boolean":
		return UpdateCompanyOverrideRequestBodyValueTypeBoolean, nil
	case "numeric":
		return UpdateCompanyOverrideRequestBodyValueTypeNumeric, nil
	case "trait":
		return UpdateCompanyOverrideRequestBodyValueTypeTrait, nil
	case "unlimited":
		return UpdateCompanyOverrideRequestBodyValueTypeUnlimited, nil
	}
	var t UpdateCompanyOverrideRequestBodyValueType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (u UpdateCompanyOverrideRequestBodyValueType) Ptr() *UpdateCompanyOverrideRequestBodyValueType {
	return &u
}

type UpdateCompanyOverrideResponse struct {
	Data *CompanyOverrideResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *UpdateCompanyOverrideResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdateCompanyOverrideResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateCompanyOverrideResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateCompanyOverrideResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdateCompanyOverrideResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UpdatePlanEntitlementRequestBodyMetricPeriod string

const (
	UpdatePlanEntitlementRequestBodyMetricPeriodCurrentMonth UpdatePlanEntitlementRequestBodyMetricPeriod = "current_month"
	UpdatePlanEntitlementRequestBodyMetricPeriodCurrentWeek  UpdatePlanEntitlementRequestBodyMetricPeriod = "current_week"
	UpdatePlanEntitlementRequestBodyMetricPeriodCurrentDay   UpdatePlanEntitlementRequestBodyMetricPeriod = "current_day"
)

func NewUpdatePlanEntitlementRequestBodyMetricPeriodFromString(s string) (UpdatePlanEntitlementRequestBodyMetricPeriod, error) {
	switch s {
	case "current_month":
		return UpdatePlanEntitlementRequestBodyMetricPeriodCurrentMonth, nil
	case "current_week":
		return UpdatePlanEntitlementRequestBodyMetricPeriodCurrentWeek, nil
	case "current_day":
		return UpdatePlanEntitlementRequestBodyMetricPeriodCurrentDay, nil
	}
	var t UpdatePlanEntitlementRequestBodyMetricPeriod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (u UpdatePlanEntitlementRequestBodyMetricPeriod) Ptr() *UpdatePlanEntitlementRequestBodyMetricPeriod {
	return &u
}

type UpdatePlanEntitlementRequestBodyValueType string

const (
	UpdatePlanEntitlementRequestBodyValueTypeBoolean   UpdatePlanEntitlementRequestBodyValueType = "boolean"
	UpdatePlanEntitlementRequestBodyValueTypeNumeric   UpdatePlanEntitlementRequestBodyValueType = "numeric"
	UpdatePlanEntitlementRequestBodyValueTypeTrait     UpdatePlanEntitlementRequestBodyValueType = "trait"
	UpdatePlanEntitlementRequestBodyValueTypeUnlimited UpdatePlanEntitlementRequestBodyValueType = "unlimited"
)

func NewUpdatePlanEntitlementRequestBodyValueTypeFromString(s string) (UpdatePlanEntitlementRequestBodyValueType, error) {
	switch s {
	case "boolean":
		return UpdatePlanEntitlementRequestBodyValueTypeBoolean, nil
	case "numeric":
		return UpdatePlanEntitlementRequestBodyValueTypeNumeric, nil
	case "trait":
		return UpdatePlanEntitlementRequestBodyValueTypeTrait, nil
	case "unlimited":
		return UpdatePlanEntitlementRequestBodyValueTypeUnlimited, nil
	}
	var t UpdatePlanEntitlementRequestBodyValueType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (u UpdatePlanEntitlementRequestBodyValueType) Ptr() *UpdatePlanEntitlementRequestBodyValueType {
	return &u
}

type UpdatePlanEntitlementResponse struct {
	Data *PlanEntitlementResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *UpdatePlanEntitlementResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdatePlanEntitlementResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdatePlanEntitlementResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdatePlanEntitlementResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdatePlanEntitlementResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UpdateCompanyOverrideRequestBody struct {
	MetricPeriod *UpdateCompanyOverrideRequestBodyMetricPeriod `json:"metric_period,omitempty" url:"-"`
	ValueBool    *bool                                         `json:"value_bool,omitempty" url:"-"`
	ValueNumeric *int                                          `json:"value_numeric,omitempty" url:"-"`
	ValueTraitID *string                                       `json:"value_trait_id,omitempty" url:"-"`
	ValueType    UpdateCompanyOverrideRequestBodyValueType     `json:"value_type" url:"-"`
}

type UpdatePlanEntitlementRequestBody struct {
	MetricPeriod *UpdatePlanEntitlementRequestBodyMetricPeriod `json:"metric_period,omitempty" url:"-"`
	ValueBool    *bool                                         `json:"value_bool,omitempty" url:"-"`
	ValueNumeric *int                                          `json:"value_numeric,omitempty" url:"-"`
	ValueTraitID *string                                       `json:"value_trait_id,omitempty" url:"-"`
	ValueType    UpdatePlanEntitlementRequestBodyValueType     `json:"value_type" url:"-"`
}
