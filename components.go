// This file was auto-generated by Fern from our API Definition.

package schematichq

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/schematichq/schematic-go/internal"
	time "time"
)

type CountComponentsRequest struct {
	Q *string `json:"-" url:"q,omitempty"`
	// Page limit (default 100)
	Limit *int `json:"-" url:"limit,omitempty"`
	// Page offset (default 0)
	Offset *int `json:"-" url:"offset,omitempty"`
}

type CreateComponentRequestBody struct {
	Ast        map[string]float64                   `json:"ast,omitempty" url:"-"`
	EntityType CreateComponentRequestBodyEntityType `json:"entity_type" url:"-"`
	Name       string                               `json:"name" url:"-"`
}

type ListComponentsRequest struct {
	Q *string `json:"-" url:"q,omitempty"`
	// Page limit (default 100)
	Limit *int `json:"-" url:"limit,omitempty"`
	// Page offset (default 0)
	Offset *int `json:"-" url:"offset,omitempty"`
}

type PreviewComponentDataRequest struct {
	CompanyID   *string `json:"-" url:"company_id,omitempty"`
	ComponentID *string `json:"-" url:"component_id,omitempty"`
}

type CompanyPlanDetailResponseData struct {
	AudienceType     *string                           `json:"audience_type,omitempty" url:"audience_type,omitempty"`
	BillingProduct   *BillingProductDetailResponseData `json:"billing_product,omitempty" url:"billing_product,omitempty"`
	CompanyCanTrial  bool                              `json:"company_can_trial" url:"company_can_trial"`
	CompanyCount     int                               `json:"company_count" url:"company_count"`
	CreatedAt        time.Time                         `json:"created_at" url:"created_at"`
	Current          bool                              `json:"current" url:"current"`
	Custom           bool                              `json:"custom" url:"custom"`
	CustomPlanConfig *CustomPlanConfig                 `json:"custom_plan_config,omitempty" url:"custom_plan_config,omitempty"`
	Description      string                            `json:"description" url:"description"`
	Entitlements     []*PlanEntitlementResponseData    `json:"entitlements,omitempty" url:"entitlements,omitempty"`
	Features         []*FeatureDetailResponseData      `json:"features,omitempty" url:"features,omitempty"`
	Icon             string                            `json:"icon" url:"icon"`
	ID               string                            `json:"id" url:"id"`
	IsCustom         bool                              `json:"is_custom" url:"is_custom"`
	IsDefault        bool                              `json:"is_default" url:"is_default"`
	IsFree           bool                              `json:"is_free" url:"is_free"`
	IsTrialable      bool                              `json:"is_trialable" url:"is_trialable"`
	MonthlyPrice     *BillingPriceResponseData         `json:"monthly_price,omitempty" url:"monthly_price,omitempty"`
	Name             string                            `json:"name" url:"name"`
	PlanType         string                            `json:"plan_type" url:"plan_type"`
	TrialDays        *int                              `json:"trial_days,omitempty" url:"trial_days,omitempty"`
	UpdatedAt        time.Time                         `json:"updated_at" url:"updated_at"`
	Valid            bool                              `json:"valid" url:"valid"`
	YearlyPrice      *BillingPriceResponseData         `json:"yearly_price,omitempty" url:"yearly_price,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CompanyPlanDetailResponseData) GetAudienceType() *string {
	if c == nil {
		return nil
	}
	return c.AudienceType
}

func (c *CompanyPlanDetailResponseData) GetBillingProduct() *BillingProductDetailResponseData {
	if c == nil {
		return nil
	}
	return c.BillingProduct
}

func (c *CompanyPlanDetailResponseData) GetCompanyCanTrial() bool {
	if c == nil {
		return false
	}
	return c.CompanyCanTrial
}

func (c *CompanyPlanDetailResponseData) GetCompanyCount() int {
	if c == nil {
		return 0
	}
	return c.CompanyCount
}

func (c *CompanyPlanDetailResponseData) GetCreatedAt() time.Time {
	if c == nil {
		return time.Time{}
	}
	return c.CreatedAt
}

func (c *CompanyPlanDetailResponseData) GetCurrent() bool {
	if c == nil {
		return false
	}
	return c.Current
}

func (c *CompanyPlanDetailResponseData) GetCustom() bool {
	if c == nil {
		return false
	}
	return c.Custom
}

func (c *CompanyPlanDetailResponseData) GetCustomPlanConfig() *CustomPlanConfig {
	if c == nil {
		return nil
	}
	return c.CustomPlanConfig
}

func (c *CompanyPlanDetailResponseData) GetDescription() string {
	if c == nil {
		return ""
	}
	return c.Description
}

func (c *CompanyPlanDetailResponseData) GetEntitlements() []*PlanEntitlementResponseData {
	if c == nil {
		return nil
	}
	return c.Entitlements
}

func (c *CompanyPlanDetailResponseData) GetFeatures() []*FeatureDetailResponseData {
	if c == nil {
		return nil
	}
	return c.Features
}

func (c *CompanyPlanDetailResponseData) GetIcon() string {
	if c == nil {
		return ""
	}
	return c.Icon
}

func (c *CompanyPlanDetailResponseData) GetID() string {
	if c == nil {
		return ""
	}
	return c.ID
}

func (c *CompanyPlanDetailResponseData) GetIsCustom() bool {
	if c == nil {
		return false
	}
	return c.IsCustom
}

func (c *CompanyPlanDetailResponseData) GetIsDefault() bool {
	if c == nil {
		return false
	}
	return c.IsDefault
}

func (c *CompanyPlanDetailResponseData) GetIsFree() bool {
	if c == nil {
		return false
	}
	return c.IsFree
}

func (c *CompanyPlanDetailResponseData) GetIsTrialable() bool {
	if c == nil {
		return false
	}
	return c.IsTrialable
}

func (c *CompanyPlanDetailResponseData) GetMonthlyPrice() *BillingPriceResponseData {
	if c == nil {
		return nil
	}
	return c.MonthlyPrice
}

func (c *CompanyPlanDetailResponseData) GetName() string {
	if c == nil {
		return ""
	}
	return c.Name
}

func (c *CompanyPlanDetailResponseData) GetPlanType() string {
	if c == nil {
		return ""
	}
	return c.PlanType
}

func (c *CompanyPlanDetailResponseData) GetTrialDays() *int {
	if c == nil {
		return nil
	}
	return c.TrialDays
}

func (c *CompanyPlanDetailResponseData) GetUpdatedAt() time.Time {
	if c == nil {
		return time.Time{}
	}
	return c.UpdatedAt
}

func (c *CompanyPlanDetailResponseData) GetValid() bool {
	if c == nil {
		return false
	}
	return c.Valid
}

func (c *CompanyPlanDetailResponseData) GetYearlyPrice() *BillingPriceResponseData {
	if c == nil {
		return nil
	}
	return c.YearlyPrice
}

func (c *CompanyPlanDetailResponseData) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CompanyPlanDetailResponseData) UnmarshalJSON(data []byte) error {
	type embed CompanyPlanDetailResponseData
	var unmarshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at"`
		UpdatedAt *internal.DateTime `json:"updated_at"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = CompanyPlanDetailResponseData(unmarshaler.embed)
	c.CreatedAt = unmarshaler.CreatedAt.Time()
	c.UpdatedAt = unmarshaler.UpdatedAt.Time()
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CompanyPlanDetailResponseData) MarshalJSON() ([]byte, error) {
	type embed CompanyPlanDetailResponseData
	var marshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at"`
		UpdatedAt *internal.DateTime `json:"updated_at"`
	}{
		embed:     embed(*c),
		CreatedAt: internal.NewDateTime(c.CreatedAt),
		UpdatedAt: internal.NewDateTime(c.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (c *CompanyPlanDetailResponseData) String() string {
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type ComponentCapabilities struct {
	BadgeVisibility bool `json:"badge_visibility" url:"badge_visibility"`
	Checkout        bool `json:"checkout" url:"checkout"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *ComponentCapabilities) GetBadgeVisibility() bool {
	if c == nil {
		return false
	}
	return c.BadgeVisibility
}

func (c *ComponentCapabilities) GetCheckout() bool {
	if c == nil {
		return false
	}
	return c.Checkout
}

func (c *ComponentCapabilities) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *ComponentCapabilities) UnmarshalJSON(data []byte) error {
	type unmarshaler ComponentCapabilities
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ComponentCapabilities(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *ComponentCapabilities) String() string {
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

// The returned resource
type ComponentPreviewResponseData struct {
	ActiveAddOns                 []*CompanyPlanDetailResponseData     `json:"active_add_ons,omitempty" url:"active_add_ons,omitempty"`
	ActivePlans                  []*CompanyPlanDetailResponseData     `json:"active_plans,omitempty" url:"active_plans,omitempty"`
	ActiveUsageBasedEntitlements []*UsageBasedEntitlementResponseData `json:"active_usage_based_entitlements,omitempty" url:"active_usage_based_entitlements,omitempty"`
	Capabilities                 *ComponentCapabilities               `json:"capabilities,omitempty" url:"capabilities,omitempty"`
	Company                      *CompanyDetailResponseData           `json:"company,omitempty" url:"company,omitempty"`
	Component                    *ComponentResponseData               `json:"component,omitempty" url:"component,omitempty"`
	DefaultPlan                  *PlanDetailResponseData              `json:"default_plan,omitempty" url:"default_plan,omitempty"`
	FeatureUsage                 *FeatureUsageDetailResponseData      `json:"feature_usage,omitempty" url:"feature_usage,omitempty"`
	Invoices                     []*InvoiceResponseData               `json:"invoices,omitempty" url:"invoices,omitempty"`
	StripeEmbed                  *StripeEmbedInfo                     `json:"stripe_embed,omitempty" url:"stripe_embed,omitempty"`
	Subscription                 *CompanySubscriptionResponseData     `json:"subscription,omitempty" url:"subscription,omitempty"`
	TrialPaymentMethodRequired   *bool                                `json:"trial_payment_method_required,omitempty" url:"trial_payment_method_required,omitempty"`
	UpcomingInvoice              *InvoiceResponseData                 `json:"upcoming_invoice,omitempty" url:"upcoming_invoice,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *ComponentPreviewResponseData) GetActiveAddOns() []*CompanyPlanDetailResponseData {
	if c == nil {
		return nil
	}
	return c.ActiveAddOns
}

func (c *ComponentPreviewResponseData) GetActivePlans() []*CompanyPlanDetailResponseData {
	if c == nil {
		return nil
	}
	return c.ActivePlans
}

func (c *ComponentPreviewResponseData) GetActiveUsageBasedEntitlements() []*UsageBasedEntitlementResponseData {
	if c == nil {
		return nil
	}
	return c.ActiveUsageBasedEntitlements
}

func (c *ComponentPreviewResponseData) GetCapabilities() *ComponentCapabilities {
	if c == nil {
		return nil
	}
	return c.Capabilities
}

func (c *ComponentPreviewResponseData) GetCompany() *CompanyDetailResponseData {
	if c == nil {
		return nil
	}
	return c.Company
}

func (c *ComponentPreviewResponseData) GetComponent() *ComponentResponseData {
	if c == nil {
		return nil
	}
	return c.Component
}

func (c *ComponentPreviewResponseData) GetDefaultPlan() *PlanDetailResponseData {
	if c == nil {
		return nil
	}
	return c.DefaultPlan
}

func (c *ComponentPreviewResponseData) GetFeatureUsage() *FeatureUsageDetailResponseData {
	if c == nil {
		return nil
	}
	return c.FeatureUsage
}

func (c *ComponentPreviewResponseData) GetInvoices() []*InvoiceResponseData {
	if c == nil {
		return nil
	}
	return c.Invoices
}

func (c *ComponentPreviewResponseData) GetStripeEmbed() *StripeEmbedInfo {
	if c == nil {
		return nil
	}
	return c.StripeEmbed
}

func (c *ComponentPreviewResponseData) GetSubscription() *CompanySubscriptionResponseData {
	if c == nil {
		return nil
	}
	return c.Subscription
}

func (c *ComponentPreviewResponseData) GetTrialPaymentMethodRequired() *bool {
	if c == nil {
		return nil
	}
	return c.TrialPaymentMethodRequired
}

func (c *ComponentPreviewResponseData) GetUpcomingInvoice() *InvoiceResponseData {
	if c == nil {
		return nil
	}
	return c.UpcomingInvoice
}

func (c *ComponentPreviewResponseData) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *ComponentPreviewResponseData) UnmarshalJSON(data []byte) error {
	type unmarshaler ComponentPreviewResponseData
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ComponentPreviewResponseData(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *ComponentPreviewResponseData) String() string {
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

// The updated resource
type ComponentResponseData struct {
	Ast       map[string]float64 `json:"ast,omitempty" url:"ast,omitempty"`
	CreatedAt time.Time          `json:"created_at" url:"created_at"`
	ID        string             `json:"id" url:"id"`
	Name      string             `json:"name" url:"name"`
	State     string             `json:"state" url:"state"`
	Type      string             `json:"type" url:"type"`
	UpdatedAt time.Time          `json:"updated_at" url:"updated_at"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *ComponentResponseData) GetAst() map[string]float64 {
	if c == nil {
		return nil
	}
	return c.Ast
}

func (c *ComponentResponseData) GetCreatedAt() time.Time {
	if c == nil {
		return time.Time{}
	}
	return c.CreatedAt
}

func (c *ComponentResponseData) GetID() string {
	if c == nil {
		return ""
	}
	return c.ID
}

func (c *ComponentResponseData) GetName() string {
	if c == nil {
		return ""
	}
	return c.Name
}

func (c *ComponentResponseData) GetState() string {
	if c == nil {
		return ""
	}
	return c.State
}

func (c *ComponentResponseData) GetType() string {
	if c == nil {
		return ""
	}
	return c.Type
}

func (c *ComponentResponseData) GetUpdatedAt() time.Time {
	if c == nil {
		return time.Time{}
	}
	return c.UpdatedAt
}

func (c *ComponentResponseData) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *ComponentResponseData) UnmarshalJSON(data []byte) error {
	type embed ComponentResponseData
	var unmarshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at"`
		UpdatedAt *internal.DateTime `json:"updated_at"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = ComponentResponseData(unmarshaler.embed)
	c.CreatedAt = unmarshaler.CreatedAt.Time()
	c.UpdatedAt = unmarshaler.UpdatedAt.Time()
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *ComponentResponseData) MarshalJSON() ([]byte, error) {
	type embed ComponentResponseData
	var marshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at"`
		UpdatedAt *internal.DateTime `json:"updated_at"`
	}{
		embed:     embed(*c),
		CreatedAt: internal.NewDateTime(c.CreatedAt),
		UpdatedAt: internal.NewDateTime(c.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (c *ComponentResponseData) String() string {
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type StripeEmbedInfo struct {
	PublishableKey          string  `json:"publishable_key" url:"publishable_key"`
	SetupIntentClientSecret *string `json:"setup_intent_client_secret,omitempty" url:"setup_intent_client_secret,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *StripeEmbedInfo) GetPublishableKey() string {
	if s == nil {
		return ""
	}
	return s.PublishableKey
}

func (s *StripeEmbedInfo) GetSetupIntentClientSecret() *string {
	if s == nil {
		return nil
	}
	return s.SetupIntentClientSecret
}

func (s *StripeEmbedInfo) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *StripeEmbedInfo) UnmarshalJSON(data []byte) error {
	type unmarshaler StripeEmbedInfo
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = StripeEmbedInfo(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *StripeEmbedInfo) String() string {
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// Input parameters
type CountComponentsParams struct {
	// Page limit (default 100)
	Limit *int `json:"limit,omitempty" url:"limit,omitempty"`
	// Page offset (default 0)
	Offset *int    `json:"offset,omitempty" url:"offset,omitempty"`
	Q      *string `json:"q,omitempty" url:"q,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CountComponentsParams) GetLimit() *int {
	if c == nil {
		return nil
	}
	return c.Limit
}

func (c *CountComponentsParams) GetOffset() *int {
	if c == nil {
		return nil
	}
	return c.Offset
}

func (c *CountComponentsParams) GetQ() *string {
	if c == nil {
		return nil
	}
	return c.Q
}

func (c *CountComponentsParams) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CountComponentsParams) UnmarshalJSON(data []byte) error {
	type unmarshaler CountComponentsParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CountComponentsParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CountComponentsParams) String() string {
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CountComponentsResponse struct {
	Data *CountResponse `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params *CountComponentsParams `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CountComponentsResponse) GetData() *CountResponse {
	if c == nil {
		return nil
	}
	return c.Data
}

func (c *CountComponentsResponse) GetParams() *CountComponentsParams {
	if c == nil {
		return nil
	}
	return c.Params
}

func (c *CountComponentsResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CountComponentsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CountComponentsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CountComponentsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CountComponentsResponse) String() string {
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CreateComponentRequestBodyEntityType string

const (
	CreateComponentRequestBodyEntityTypeEntitlement CreateComponentRequestBodyEntityType = "entitlement"
	CreateComponentRequestBodyEntityTypeBilling     CreateComponentRequestBodyEntityType = "billing"
)

func NewCreateComponentRequestBodyEntityTypeFromString(s string) (CreateComponentRequestBodyEntityType, error) {
	switch s {
	case "entitlement":
		return CreateComponentRequestBodyEntityTypeEntitlement, nil
	case "billing":
		return CreateComponentRequestBodyEntityTypeBilling, nil
	}
	var t CreateComponentRequestBodyEntityType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CreateComponentRequestBodyEntityType) Ptr() *CreateComponentRequestBodyEntityType {
	return &c
}

type CreateComponentResponse struct {
	Data *ComponentResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CreateComponentResponse) GetData() *ComponentResponseData {
	if c == nil {
		return nil
	}
	return c.Data
}

func (c *CreateComponentResponse) GetParams() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.Params
}

func (c *CreateComponentResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateComponentResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateComponentResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateComponentResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateComponentResponse) String() string {
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type DeleteComponentResponse struct {
	Data *DeleteResponse `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *DeleteComponentResponse) GetData() *DeleteResponse {
	if d == nil {
		return nil
	}
	return d.Data
}

func (d *DeleteComponentResponse) GetParams() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.Params
}

func (d *DeleteComponentResponse) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeleteComponentResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteComponentResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteComponentResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeleteComponentResponse) String() string {
	if len(d.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(d.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type GetComponentResponse struct {
	Data *ComponentResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetComponentResponse) GetData() *ComponentResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetComponentResponse) GetParams() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.Params
}

func (g *GetComponentResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetComponentResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetComponentResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetComponentResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetComponentResponse) String() string {
	if len(g.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(g.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

// Input parameters
type ListComponentsParams struct {
	// Page limit (default 100)
	Limit *int `json:"limit,omitempty" url:"limit,omitempty"`
	// Page offset (default 0)
	Offset *int    `json:"offset,omitempty" url:"offset,omitempty"`
	Q      *string `json:"q,omitempty" url:"q,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListComponentsParams) GetLimit() *int {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListComponentsParams) GetOffset() *int {
	if l == nil {
		return nil
	}
	return l.Offset
}

func (l *ListComponentsParams) GetQ() *string {
	if l == nil {
		return nil
	}
	return l.Q
}

func (l *ListComponentsParams) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListComponentsParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ListComponentsParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListComponentsParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListComponentsParams) String() string {
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type ListComponentsResponse struct {
	// The returned resources
	Data []*ComponentResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params *ListComponentsParams `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *ListComponentsResponse) GetData() []*ComponentResponseData {
	if l == nil {
		return nil
	}
	return l.Data
}

func (l *ListComponentsResponse) GetParams() *ListComponentsParams {
	if l == nil {
		return nil
	}
	return l.Params
}

func (l *ListComponentsResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListComponentsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListComponentsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListComponentsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListComponentsResponse) String() string {
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// Input parameters
type PreviewComponentDataParams struct {
	CompanyID   *string `json:"company_id,omitempty" url:"company_id,omitempty"`
	ComponentID *string `json:"component_id,omitempty" url:"component_id,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PreviewComponentDataParams) GetCompanyID() *string {
	if p == nil {
		return nil
	}
	return p.CompanyID
}

func (p *PreviewComponentDataParams) GetComponentID() *string {
	if p == nil {
		return nil
	}
	return p.ComponentID
}

func (p *PreviewComponentDataParams) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreviewComponentDataParams) UnmarshalJSON(data []byte) error {
	type unmarshaler PreviewComponentDataParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreviewComponentDataParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreviewComponentDataParams) String() string {
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PreviewComponentDataResponse struct {
	Data *ComponentPreviewResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params *PreviewComponentDataParams `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PreviewComponentDataResponse) GetData() *ComponentPreviewResponseData {
	if p == nil {
		return nil
	}
	return p.Data
}

func (p *PreviewComponentDataResponse) GetParams() *PreviewComponentDataParams {
	if p == nil {
		return nil
	}
	return p.Params
}

func (p *PreviewComponentDataResponse) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreviewComponentDataResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler PreviewComponentDataResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreviewComponentDataResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreviewComponentDataResponse) String() string {
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type UpdateComponentRequestBodyEntityType string

const (
	UpdateComponentRequestBodyEntityTypeEntitlement UpdateComponentRequestBodyEntityType = "entitlement"
	UpdateComponentRequestBodyEntityTypeBilling     UpdateComponentRequestBodyEntityType = "billing"
)

func NewUpdateComponentRequestBodyEntityTypeFromString(s string) (UpdateComponentRequestBodyEntityType, error) {
	switch s {
	case "entitlement":
		return UpdateComponentRequestBodyEntityTypeEntitlement, nil
	case "billing":
		return UpdateComponentRequestBodyEntityTypeBilling, nil
	}
	var t UpdateComponentRequestBodyEntityType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (u UpdateComponentRequestBodyEntityType) Ptr() *UpdateComponentRequestBodyEntityType {
	return &u
}

type UpdateComponentRequestBodyState string

const (
	UpdateComponentRequestBodyStateDraft UpdateComponentRequestBodyState = "draft"
	UpdateComponentRequestBodyStateLive  UpdateComponentRequestBodyState = "live"
)

func NewUpdateComponentRequestBodyStateFromString(s string) (UpdateComponentRequestBodyState, error) {
	switch s {
	case "draft":
		return UpdateComponentRequestBodyStateDraft, nil
	case "live":
		return UpdateComponentRequestBodyStateLive, nil
	}
	var t UpdateComponentRequestBodyState
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (u UpdateComponentRequestBodyState) Ptr() *UpdateComponentRequestBodyState {
	return &u
}

type UpdateComponentResponse struct {
	Data *ComponentResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UpdateComponentResponse) GetData() *ComponentResponseData {
	if u == nil {
		return nil
	}
	return u.Data
}

func (u *UpdateComponentResponse) GetParams() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.Params
}

func (u *UpdateComponentResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdateComponentResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateComponentResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateComponentResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdateComponentResponse) String() string {
	if len(u.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(u.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UpdateComponentRequestBody struct {
	Ast        map[string]float64                    `json:"ast,omitempty" url:"-"`
	EntityType *UpdateComponentRequestBodyEntityType `json:"entity_type,omitempty" url:"-"`
	Name       *string                               `json:"name,omitempty" url:"-"`
	State      *UpdateComponentRequestBodyState      `json:"state,omitempty" url:"-"`
}
