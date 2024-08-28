/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the CreateBillingInvoiceRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBillingInvoiceRequestBody{}

// CreateBillingInvoiceRequestBody struct for CreateBillingInvoiceRequestBody
type CreateBillingInvoiceRequestBody struct {
	AmountDue              int32          `json:"amount_due"`
	AmountPaid             int32          `json:"amount_paid"`
	AmountRemaining        int32          `json:"amount_remaining"`
	CollectionMethod       string         `json:"collection_method"`
	CompanyId              NullableString `json:"company_id,omitempty"`
	Currency               string         `json:"currency"`
	CustomerExternalId     string         `json:"customer_external_id"`
	DueDate                NullableTime   `json:"due_date,omitempty"`
	ExternalId             string         `json:"external_id"`
	SubscriptionExternalId NullableString `json:"subscription_external_id,omitempty"`
	Subtotal               int32          `json:"subtotal"`
	AdditionalProperties   map[string]interface{}
}

type _CreateBillingInvoiceRequestBody CreateBillingInvoiceRequestBody

// NewCreateBillingInvoiceRequestBody instantiates a new CreateBillingInvoiceRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBillingInvoiceRequestBody(amountDue int32, amountPaid int32, amountRemaining int32, collectionMethod string, currency string, customerExternalId string, externalId string, subtotal int32) *CreateBillingInvoiceRequestBody {
	this := CreateBillingInvoiceRequestBody{}
	this.AmountDue = amountDue
	this.AmountPaid = amountPaid
	this.AmountRemaining = amountRemaining
	this.CollectionMethod = collectionMethod
	this.Currency = currency
	this.CustomerExternalId = customerExternalId
	this.ExternalId = externalId
	this.Subtotal = subtotal
	return &this
}

// NewCreateBillingInvoiceRequestBodyWithDefaults instantiates a new CreateBillingInvoiceRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBillingInvoiceRequestBodyWithDefaults() *CreateBillingInvoiceRequestBody {
	this := CreateBillingInvoiceRequestBody{}
	return &this
}

// GetAmountDue returns the AmountDue field value
func (o *CreateBillingInvoiceRequestBody) GetAmountDue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountDue
}

// GetAmountDueOk returns a tuple with the AmountDue field value
// and a boolean to check if the value has been set.
func (o *CreateBillingInvoiceRequestBody) GetAmountDueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountDue, true
}

// SetAmountDue sets field value
func (o *CreateBillingInvoiceRequestBody) SetAmountDue(v int32) {
	o.AmountDue = v
}

// GetAmountPaid returns the AmountPaid field value
func (o *CreateBillingInvoiceRequestBody) GetAmountPaid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountPaid
}

// GetAmountPaidOk returns a tuple with the AmountPaid field value
// and a boolean to check if the value has been set.
func (o *CreateBillingInvoiceRequestBody) GetAmountPaidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountPaid, true
}

// SetAmountPaid sets field value
func (o *CreateBillingInvoiceRequestBody) SetAmountPaid(v int32) {
	o.AmountPaid = v
}

// GetAmountRemaining returns the AmountRemaining field value
func (o *CreateBillingInvoiceRequestBody) GetAmountRemaining() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountRemaining
}

// GetAmountRemainingOk returns a tuple with the AmountRemaining field value
// and a boolean to check if the value has been set.
func (o *CreateBillingInvoiceRequestBody) GetAmountRemainingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountRemaining, true
}

// SetAmountRemaining sets field value
func (o *CreateBillingInvoiceRequestBody) SetAmountRemaining(v int32) {
	o.AmountRemaining = v
}

// GetCollectionMethod returns the CollectionMethod field value
func (o *CreateBillingInvoiceRequestBody) GetCollectionMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionMethod
}

// GetCollectionMethodOk returns a tuple with the CollectionMethod field value
// and a boolean to check if the value has been set.
func (o *CreateBillingInvoiceRequestBody) GetCollectionMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionMethod, true
}

// SetCollectionMethod sets field value
func (o *CreateBillingInvoiceRequestBody) SetCollectionMethod(v string) {
	o.CollectionMethod = v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateBillingInvoiceRequestBody) GetCompanyId() string {
	if o == nil || IsNil(o.CompanyId.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyId.Get()
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateBillingInvoiceRequestBody) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyId.Get(), o.CompanyId.IsSet()
}

// HasCompanyId returns a boolean if a field has been set.
func (o *CreateBillingInvoiceRequestBody) HasCompanyId() bool {
	if o != nil && o.CompanyId.IsSet() {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given NullableString and assigns it to the CompanyId field.
func (o *CreateBillingInvoiceRequestBody) SetCompanyId(v string) {
	o.CompanyId.Set(&v)
}

// SetCompanyIdNil sets the value for CompanyId to be an explicit nil
func (o *CreateBillingInvoiceRequestBody) SetCompanyIdNil() {
	o.CompanyId.Set(nil)
}

// UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
func (o *CreateBillingInvoiceRequestBody) UnsetCompanyId() {
	o.CompanyId.Unset()
}

// GetCurrency returns the Currency field value
func (o *CreateBillingInvoiceRequestBody) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CreateBillingInvoiceRequestBody) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CreateBillingInvoiceRequestBody) SetCurrency(v string) {
	o.Currency = v
}

// GetCustomerExternalId returns the CustomerExternalId field value
func (o *CreateBillingInvoiceRequestBody) GetCustomerExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerExternalId
}

// GetCustomerExternalIdOk returns a tuple with the CustomerExternalId field value
// and a boolean to check if the value has been set.
func (o *CreateBillingInvoiceRequestBody) GetCustomerExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerExternalId, true
}

// SetCustomerExternalId sets field value
func (o *CreateBillingInvoiceRequestBody) SetCustomerExternalId(v string) {
	o.CustomerExternalId = v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateBillingInvoiceRequestBody) GetDueDate() time.Time {
	if o == nil || IsNil(o.DueDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DueDate.Get()
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateBillingInvoiceRequestBody) GetDueDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DueDate.Get(), o.DueDate.IsSet()
}

// HasDueDate returns a boolean if a field has been set.
func (o *CreateBillingInvoiceRequestBody) HasDueDate() bool {
	if o != nil && o.DueDate.IsSet() {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given NullableTime and assigns it to the DueDate field.
func (o *CreateBillingInvoiceRequestBody) SetDueDate(v time.Time) {
	o.DueDate.Set(&v)
}

// SetDueDateNil sets the value for DueDate to be an explicit nil
func (o *CreateBillingInvoiceRequestBody) SetDueDateNil() {
	o.DueDate.Set(nil)
}

// UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
func (o *CreateBillingInvoiceRequestBody) UnsetDueDate() {
	o.DueDate.Unset()
}

// GetExternalId returns the ExternalId field value
func (o *CreateBillingInvoiceRequestBody) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *CreateBillingInvoiceRequestBody) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *CreateBillingInvoiceRequestBody) SetExternalId(v string) {
	o.ExternalId = v
}

// GetSubscriptionExternalId returns the SubscriptionExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateBillingInvoiceRequestBody) GetSubscriptionExternalId() string {
	if o == nil || IsNil(o.SubscriptionExternalId.Get()) {
		var ret string
		return ret
	}
	return *o.SubscriptionExternalId.Get()
}

// GetSubscriptionExternalIdOk returns a tuple with the SubscriptionExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateBillingInvoiceRequestBody) GetSubscriptionExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionExternalId.Get(), o.SubscriptionExternalId.IsSet()
}

// HasSubscriptionExternalId returns a boolean if a field has been set.
func (o *CreateBillingInvoiceRequestBody) HasSubscriptionExternalId() bool {
	if o != nil && o.SubscriptionExternalId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionExternalId gets a reference to the given NullableString and assigns it to the SubscriptionExternalId field.
func (o *CreateBillingInvoiceRequestBody) SetSubscriptionExternalId(v string) {
	o.SubscriptionExternalId.Set(&v)
}

// SetSubscriptionExternalIdNil sets the value for SubscriptionExternalId to be an explicit nil
func (o *CreateBillingInvoiceRequestBody) SetSubscriptionExternalIdNil() {
	o.SubscriptionExternalId.Set(nil)
}

// UnsetSubscriptionExternalId ensures that no value is present for SubscriptionExternalId, not even an explicit nil
func (o *CreateBillingInvoiceRequestBody) UnsetSubscriptionExternalId() {
	o.SubscriptionExternalId.Unset()
}

// GetSubtotal returns the Subtotal field value
func (o *CreateBillingInvoiceRequestBody) GetSubtotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Subtotal
}

// GetSubtotalOk returns a tuple with the Subtotal field value
// and a boolean to check if the value has been set.
func (o *CreateBillingInvoiceRequestBody) GetSubtotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtotal, true
}

// SetSubtotal sets field value
func (o *CreateBillingInvoiceRequestBody) SetSubtotal(v int32) {
	o.Subtotal = v
}

func (o CreateBillingInvoiceRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBillingInvoiceRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount_due"] = o.AmountDue
	toSerialize["amount_paid"] = o.AmountPaid
	toSerialize["amount_remaining"] = o.AmountRemaining
	toSerialize["collection_method"] = o.CollectionMethod
	if o.CompanyId.IsSet() {
		toSerialize["company_id"] = o.CompanyId.Get()
	}
	toSerialize["currency"] = o.Currency
	toSerialize["customer_external_id"] = o.CustomerExternalId
	if o.DueDate.IsSet() {
		toSerialize["due_date"] = o.DueDate.Get()
	}
	toSerialize["external_id"] = o.ExternalId
	if o.SubscriptionExternalId.IsSet() {
		toSerialize["subscription_external_id"] = o.SubscriptionExternalId.Get()
	}
	toSerialize["subtotal"] = o.Subtotal

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateBillingInvoiceRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount_due",
		"amount_paid",
		"amount_remaining",
		"collection_method",
		"currency",
		"customer_external_id",
		"external_id",
		"subtotal",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateBillingInvoiceRequestBody := _CreateBillingInvoiceRequestBody{}

	err = json.Unmarshal(data, &varCreateBillingInvoiceRequestBody)

	if err != nil {
		return err
	}

	*o = CreateBillingInvoiceRequestBody(varCreateBillingInvoiceRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount_due")
		delete(additionalProperties, "amount_paid")
		delete(additionalProperties, "amount_remaining")
		delete(additionalProperties, "collection_method")
		delete(additionalProperties, "company_id")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "customer_external_id")
		delete(additionalProperties, "due_date")
		delete(additionalProperties, "external_id")
		delete(additionalProperties, "subscription_external_id")
		delete(additionalProperties, "subtotal")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateBillingInvoiceRequestBody struct {
	value *CreateBillingInvoiceRequestBody
	isSet bool
}

func (v NullableCreateBillingInvoiceRequestBody) Get() *CreateBillingInvoiceRequestBody {
	return v.value
}

func (v *NullableCreateBillingInvoiceRequestBody) Set(val *CreateBillingInvoiceRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBillingInvoiceRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBillingInvoiceRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBillingInvoiceRequestBody(val *CreateBillingInvoiceRequestBody) *NullableCreateBillingInvoiceRequestBody {
	return &NullableCreateBillingInvoiceRequestBody{value: val, isSet: true}
}

func (v NullableCreateBillingInvoiceRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBillingInvoiceRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}