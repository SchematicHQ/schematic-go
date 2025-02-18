// This file was auto-generated by Fern from our API Definition.

package schematichq

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/schematichq/schematic-go/internal"
	time "time"
)

type CheckoutDataRequestBody struct {
	CompanyID      string  `json:"company_id" url:"-"`
	SelectedPlanID *string `json:"selected_plan_id,omitempty" url:"-"`
}

type ChangeSubscriptionInternalRequestBody struct {
	AddOnIDs         []*UpdateAddOnRequestBody        `json:"add_on_ids,omitempty" url:"add_on_ids,omitempty"`
	CompanyID        string                           `json:"company_id" url:"company_id"`
	CouponExternalID *string                          `json:"coupon_external_id,omitempty" url:"coupon_external_id,omitempty"`
	NewPlanID        string                           `json:"new_plan_id" url:"new_plan_id"`
	NewPriceID       string                           `json:"new_price_id" url:"new_price_id"`
	PayInAdvance     []*UpdatePayInAdvanceRequestBody `json:"pay_in_advance,omitempty" url:"pay_in_advance,omitempty"`
	PaymentMethodID  *string                          `json:"payment_method_id,omitempty" url:"payment_method_id,omitempty"`
	PromoCode        *string                          `json:"promo_code,omitempty" url:"promo_code,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *ChangeSubscriptionInternalRequestBody) GetAddOnIDs() []*UpdateAddOnRequestBody {
	if c == nil {
		return nil
	}
	return c.AddOnIDs
}

func (c *ChangeSubscriptionInternalRequestBody) GetCompanyID() string {
	if c == nil {
		return ""
	}
	return c.CompanyID
}

func (c *ChangeSubscriptionInternalRequestBody) GetCouponExternalID() *string {
	if c == nil {
		return nil
	}
	return c.CouponExternalID
}

func (c *ChangeSubscriptionInternalRequestBody) GetNewPlanID() string {
	if c == nil {
		return ""
	}
	return c.NewPlanID
}

func (c *ChangeSubscriptionInternalRequestBody) GetNewPriceID() string {
	if c == nil {
		return ""
	}
	return c.NewPriceID
}

func (c *ChangeSubscriptionInternalRequestBody) GetPayInAdvance() []*UpdatePayInAdvanceRequestBody {
	if c == nil {
		return nil
	}
	return c.PayInAdvance
}

func (c *ChangeSubscriptionInternalRequestBody) GetPaymentMethodID() *string {
	if c == nil {
		return nil
	}
	return c.PaymentMethodID
}

func (c *ChangeSubscriptionInternalRequestBody) GetPromoCode() *string {
	if c == nil {
		return nil
	}
	return c.PromoCode
}

func (c *ChangeSubscriptionInternalRequestBody) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *ChangeSubscriptionInternalRequestBody) UnmarshalJSON(data []byte) error {
	type unmarshaler ChangeSubscriptionInternalRequestBody
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ChangeSubscriptionInternalRequestBody(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *ChangeSubscriptionInternalRequestBody) String() string {
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

// The requested resource
type CheckoutDataResponseData struct {
	ActiveAddOns                   []*PlanDetailResponseData            `json:"active_add_ons,omitempty" url:"active_add_ons,omitempty"`
	ActivePlan                     *PlanDetailResponseData              `json:"active_plan,omitempty" url:"active_plan,omitempty"`
	ActiveUsageBasedEntitlements   []*UsageBasedEntitlementResponseData `json:"active_usage_based_entitlements,omitempty" url:"active_usage_based_entitlements,omitempty"`
	Company                        *CompanyDetailResponseData           `json:"company,omitempty" url:"company,omitempty"`
	FeatureUsage                   *FeatureUsageDetailResponseData      `json:"feature_usage,omitempty" url:"feature_usage,omitempty"`
	SelectedPlan                   *PlanDetailResponseData              `json:"selected_plan,omitempty" url:"selected_plan,omitempty"`
	SelectedUsageBasedEntitlements []*UsageBasedEntitlementResponseData `json:"selected_usage_based_entitlements,omitempty" url:"selected_usage_based_entitlements,omitempty"`
	Subscription                   *CompanySubscriptionResponseData     `json:"subscription,omitempty" url:"subscription,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CheckoutDataResponseData) GetActiveAddOns() []*PlanDetailResponseData {
	if c == nil {
		return nil
	}
	return c.ActiveAddOns
}

func (c *CheckoutDataResponseData) GetActivePlan() *PlanDetailResponseData {
	if c == nil {
		return nil
	}
	return c.ActivePlan
}

func (c *CheckoutDataResponseData) GetActiveUsageBasedEntitlements() []*UsageBasedEntitlementResponseData {
	if c == nil {
		return nil
	}
	return c.ActiveUsageBasedEntitlements
}

func (c *CheckoutDataResponseData) GetCompany() *CompanyDetailResponseData {
	if c == nil {
		return nil
	}
	return c.Company
}

func (c *CheckoutDataResponseData) GetFeatureUsage() *FeatureUsageDetailResponseData {
	if c == nil {
		return nil
	}
	return c.FeatureUsage
}

func (c *CheckoutDataResponseData) GetSelectedPlan() *PlanDetailResponseData {
	if c == nil {
		return nil
	}
	return c.SelectedPlan
}

func (c *CheckoutDataResponseData) GetSelectedUsageBasedEntitlements() []*UsageBasedEntitlementResponseData {
	if c == nil {
		return nil
	}
	return c.SelectedUsageBasedEntitlements
}

func (c *CheckoutDataResponseData) GetSubscription() *CompanySubscriptionResponseData {
	if c == nil {
		return nil
	}
	return c.Subscription
}

func (c *CheckoutDataResponseData) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CheckoutDataResponseData) UnmarshalJSON(data []byte) error {
	type unmarshaler CheckoutDataResponseData
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CheckoutDataResponseData(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CheckoutDataResponseData) String() string {
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

// The requested resource
type PreviewSubscriptionChangeResponseData struct {
	AmountOff        int        `json:"amount_off" url:"amount_off"`
	DueNow           int        `json:"due_now" url:"due_now"`
	NewCharges       int        `json:"new_charges" url:"new_charges"`
	PercentOff       float64    `json:"percent_off" url:"percent_off"`
	PeriodStart      time.Time  `json:"period_start" url:"period_start"`
	PromoCodeApplied bool       `json:"promo_code_applied" url:"promo_code_applied"`
	Proration        int        `json:"proration" url:"proration"`
	TrialEnd         *time.Time `json:"trial_end,omitempty" url:"trial_end,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PreviewSubscriptionChangeResponseData) GetAmountOff() int {
	if p == nil {
		return 0
	}
	return p.AmountOff
}

func (p *PreviewSubscriptionChangeResponseData) GetDueNow() int {
	if p == nil {
		return 0
	}
	return p.DueNow
}

func (p *PreviewSubscriptionChangeResponseData) GetNewCharges() int {
	if p == nil {
		return 0
	}
	return p.NewCharges
}

func (p *PreviewSubscriptionChangeResponseData) GetPercentOff() float64 {
	if p == nil {
		return 0
	}
	return p.PercentOff
}

func (p *PreviewSubscriptionChangeResponseData) GetPeriodStart() time.Time {
	if p == nil {
		return time.Time{}
	}
	return p.PeriodStart
}

func (p *PreviewSubscriptionChangeResponseData) GetPromoCodeApplied() bool {
	if p == nil {
		return false
	}
	return p.PromoCodeApplied
}

func (p *PreviewSubscriptionChangeResponseData) GetProration() int {
	if p == nil {
		return 0
	}
	return p.Proration
}

func (p *PreviewSubscriptionChangeResponseData) GetTrialEnd() *time.Time {
	if p == nil {
		return nil
	}
	return p.TrialEnd
}

func (p *PreviewSubscriptionChangeResponseData) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreviewSubscriptionChangeResponseData) UnmarshalJSON(data []byte) error {
	type embed PreviewSubscriptionChangeResponseData
	var unmarshaler = struct {
		embed
		PeriodStart *internal.DateTime `json:"period_start"`
		TrialEnd    *internal.DateTime `json:"trial_end,omitempty"`
	}{
		embed: embed(*p),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*p = PreviewSubscriptionChangeResponseData(unmarshaler.embed)
	p.PeriodStart = unmarshaler.PeriodStart.Time()
	p.TrialEnd = unmarshaler.TrialEnd.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreviewSubscriptionChangeResponseData) MarshalJSON() ([]byte, error) {
	type embed PreviewSubscriptionChangeResponseData
	var marshaler = struct {
		embed
		PeriodStart *internal.DateTime `json:"period_start"`
		TrialEnd    *internal.DateTime `json:"trial_end,omitempty"`
	}{
		embed:       embed(*p),
		PeriodStart: internal.NewDateTime(p.PeriodStart),
		TrialEnd:    internal.NewOptionalDateTime(p.TrialEnd),
	}
	return json.Marshal(marshaler)
}

func (p *PreviewSubscriptionChangeResponseData) String() string {
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

type UpdateAddOnRequestBody struct {
	AddOnID string `json:"add_on_id" url:"add_on_id"`
	PriceID string `json:"price_id" url:"price_id"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UpdateAddOnRequestBody) GetAddOnID() string {
	if u == nil {
		return ""
	}
	return u.AddOnID
}

func (u *UpdateAddOnRequestBody) GetPriceID() string {
	if u == nil {
		return ""
	}
	return u.PriceID
}

func (u *UpdateAddOnRequestBody) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdateAddOnRequestBody) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateAddOnRequestBody
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateAddOnRequestBody(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdateAddOnRequestBody) String() string {
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

type UpdatePayInAdvanceRequestBody struct {
	PriceID  string `json:"price_id" url:"price_id"`
	Quantity int    `json:"quantity" url:"quantity"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UpdatePayInAdvanceRequestBody) GetPriceID() string {
	if u == nil {
		return ""
	}
	return u.PriceID
}

func (u *UpdatePayInAdvanceRequestBody) GetQuantity() int {
	if u == nil {
		return 0
	}
	return u.Quantity
}

func (u *UpdatePayInAdvanceRequestBody) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdatePayInAdvanceRequestBody) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdatePayInAdvanceRequestBody
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdatePayInAdvanceRequestBody(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdatePayInAdvanceRequestBody) String() string {
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

type CheckoutInternalResponse struct {
	Data *BillingSubscriptionResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CheckoutInternalResponse) GetData() *BillingSubscriptionResponseData {
	if c == nil {
		return nil
	}
	return c.Data
}

func (c *CheckoutInternalResponse) GetParams() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.Params
}

func (c *CheckoutInternalResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CheckoutInternalResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CheckoutInternalResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CheckoutInternalResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CheckoutInternalResponse) String() string {
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

type GetCheckoutDataResponse struct {
	Data *CheckoutDataResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (g *GetCheckoutDataResponse) GetData() *CheckoutDataResponseData {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GetCheckoutDataResponse) GetParams() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.Params
}

func (g *GetCheckoutDataResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetCheckoutDataResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetCheckoutDataResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetCheckoutDataResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	g.rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetCheckoutDataResponse) String() string {
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

type PreviewCheckoutInternalResponse struct {
	Data *PreviewSubscriptionChangeResponseData `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PreviewCheckoutInternalResponse) GetData() *PreviewSubscriptionChangeResponseData {
	if p == nil {
		return nil
	}
	return p.Data
}

func (p *PreviewCheckoutInternalResponse) GetParams() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.Params
}

func (p *PreviewCheckoutInternalResponse) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreviewCheckoutInternalResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler PreviewCheckoutInternalResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreviewCheckoutInternalResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreviewCheckoutInternalResponse) String() string {
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

type UpdateCustomerSubscriptionTrialEndResponse struct {
	Data *BillingSubscriptionView `json:"data,omitempty" url:"data,omitempty"`
	// Input parameters
	Params map[string]interface{} `json:"params,omitempty" url:"params,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (u *UpdateCustomerSubscriptionTrialEndResponse) GetData() *BillingSubscriptionView {
	if u == nil {
		return nil
	}
	return u.Data
}

func (u *UpdateCustomerSubscriptionTrialEndResponse) GetParams() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.Params
}

func (u *UpdateCustomerSubscriptionTrialEndResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdateCustomerSubscriptionTrialEndResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateCustomerSubscriptionTrialEndResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateCustomerSubscriptionTrialEndResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	u.rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdateCustomerSubscriptionTrialEndResponse) String() string {
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

type UpdateTrialEndRequestBody struct {
	TrialEnd *time.Time `json:"trial_end,omitempty" url:"-"`
}

func (u *UpdateTrialEndRequestBody) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateTrialEndRequestBody
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*u = UpdateTrialEndRequestBody(body)
	return nil
}

func (u *UpdateTrialEndRequestBody) MarshalJSON() ([]byte, error) {
	type embed UpdateTrialEndRequestBody
	var marshaler = struct {
		embed
		TrialEnd *internal.DateTime `json:"trial_end,omitempty"`
	}{
		embed:    embed(*u),
		TrialEnd: internal.NewOptionalDateTime(u.TrialEnd),
	}
	return json.Marshal(marshaler)
}
