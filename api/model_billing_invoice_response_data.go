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

// checks if the BillingInvoiceResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingInvoiceResponseData{}

// BillingInvoiceResponseData struct for BillingInvoiceResponseData
type BillingInvoiceResponseData struct {
	AmountDue              int32          `json:"amount_due"`
	AmountPaid             int32          `json:"amount_paid"`
	AmountRemaining        int32          `json:"amount_remaining"`
	CollectionMethod       string         `json:"collection_method"`
	CompanyId              NullableString `json:"company_id,omitempty"`
	CreatedAt              time.Time      `json:"created_at"`
	Currency               string         `json:"currency"`
	CustomerExternalId     string         `json:"customer_external_id"`
	DueDate                NullableTime   `json:"due_date,omitempty"`
	EnvironmentId          string         `json:"environment_id"`
	ExternalId             string         `json:"external_id"`
	Id                     string         `json:"id"`
	SubscriptionExternalId NullableString `json:"subscription_external_id,omitempty"`
	Subtotal               int32          `json:"subtotal"`
	UpdatedAt              time.Time      `json:"updated_at"`
	AdditionalProperties   map[string]interface{}
}

type _BillingInvoiceResponseData BillingInvoiceResponseData

// NewBillingInvoiceResponseData instantiates a new BillingInvoiceResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInvoiceResponseData(amountDue int32, amountPaid int32, amountRemaining int32, collectionMethod string, createdAt time.Time, currency string, customerExternalId string, environmentId string, externalId string, id string, subtotal int32, updatedAt time.Time) *BillingInvoiceResponseData {
	this := BillingInvoiceResponseData{}
	this.AmountDue = amountDue
	this.AmountPaid = amountPaid
	this.AmountRemaining = amountRemaining
	this.CollectionMethod = collectionMethod
	this.CreatedAt = createdAt
	this.Currency = currency
	this.CustomerExternalId = customerExternalId
	this.EnvironmentId = environmentId
	this.ExternalId = externalId
	this.Id = id
	this.Subtotal = subtotal
	this.UpdatedAt = updatedAt
	return &this
}

// NewBillingInvoiceResponseDataWithDefaults instantiates a new BillingInvoiceResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInvoiceResponseDataWithDefaults() *BillingInvoiceResponseData {
	this := BillingInvoiceResponseData{}
	return &this
}

// GetAmountDue returns the AmountDue field value
func (o *BillingInvoiceResponseData) GetAmountDue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountDue
}

// GetAmountDueOk returns a tuple with the AmountDue field value
// and a boolean to check if the value has been set.
func (o *BillingInvoiceResponseData) GetAmountDueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountDue, true
}

// SetAmountDue sets field value
func (o *BillingInvoiceResponseData) SetAmountDue(v int32) {
	o.AmountDue = v
}

// GetAmountPaid returns the AmountPaid field value
func (o *BillingInvoiceResponseData) GetAmountPaid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountPaid
}

// GetAmountPaidOk returns a tuple with the AmountPaid field value
// and a boolean to check if the value has been set.
func (o *BillingInvoiceResponseData) GetAmountPaidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountPaid, true
}

// SetAmountPaid sets field value
func (o *BillingInvoiceResponseData) SetAmountPaid(v int32) {
	o.AmountPaid = v
}

// GetAmountRemaining returns the AmountRemaining field value
func (o *BillingInvoiceResponseData) GetAmountRemaining() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountRemaining
}

// GetAmountRemainingOk returns a tuple with the AmountRemaining field value
// and a boolean to check if the value has been set.
func (o *BillingInvoiceResponseData) GetAmountRemainingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountRemaining, true
}

// SetAmountRemaining sets field value
func (o *BillingInvoiceResponseData) SetAmountRemaining(v int32) {
	o.AmountRemaining = v
}

// GetCollectionMethod returns the CollectionMethod field value
func (o *BillingInvoiceResponseData) GetCollectionMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionMethod
}

// GetCollectionMethodOk returns a tuple with the CollectionMethod field value
// and a boolean to check if the value has been set.
func (o *BillingInvoiceResponseData) GetCollectionMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionMethod, true
}

// SetCollectionMethod sets field value
func (o *BillingInvoiceResponseData) SetCollectionMethod(v string) {
	o.CollectionMethod = v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInvoiceResponseData) GetCompanyId() string {
	if o == nil || IsNil(o.CompanyId.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyId.Get()
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInvoiceResponseData) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyId.Get(), o.CompanyId.IsSet()
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BillingInvoiceResponseData) HasCompanyId() bool {
	if o != nil && o.CompanyId.IsSet() {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given NullableString and assigns it to the CompanyId field.
func (o *BillingInvoiceResponseData) SetCompanyId(v string) {
	o.CompanyId.Set(&v)
}

// SetCompanyIdNil sets the value for CompanyId to be an explicit nil
func (o *BillingInvoiceResponseData) SetCompanyIdNil() {
	o.CompanyId.Set(nil)
}

// UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
func (o *BillingInvoiceResponseData) UnsetCompanyId() {
	o.CompanyId.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *BillingInvoiceResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BillingInvoiceResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *BillingInvoiceResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCurrency returns the Currency field value
func (o *BillingInvoiceResponseData) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *BillingInvoiceResponseData) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *BillingInvoiceResponseData) SetCurrency(v string) {
	o.Currency = v
}

// GetCustomerExternalId returns the CustomerExternalId field value
func (o *BillingInvoiceResponseData) GetCustomerExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerExternalId
}

// GetCustomerExternalIdOk returns a tuple with the CustomerExternalId field value
// and a boolean to check if the value has been set.
func (o *BillingInvoiceResponseData) GetCustomerExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerExternalId, true
}

// SetCustomerExternalId sets field value
func (o *BillingInvoiceResponseData) SetCustomerExternalId(v string) {
	o.CustomerExternalId = v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInvoiceResponseData) GetDueDate() time.Time {
	if o == nil || IsNil(o.DueDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DueDate.Get()
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInvoiceResponseData) GetDueDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DueDate.Get(), o.DueDate.IsSet()
}

// HasDueDate returns a boolean if a field has been set.
func (o *BillingInvoiceResponseData) HasDueDate() bool {
	if o != nil && o.DueDate.IsSet() {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given NullableTime and assigns it to the DueDate field.
func (o *BillingInvoiceResponseData) SetDueDate(v time.Time) {
	o.DueDate.Set(&v)
}

// SetDueDateNil sets the value for DueDate to be an explicit nil
func (o *BillingInvoiceResponseData) SetDueDateNil() {
	o.DueDate.Set(nil)
}

// UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
func (o *BillingInvoiceResponseData) UnsetDueDate() {
	o.DueDate.Unset()
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *BillingInvoiceResponseData) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *BillingInvoiceResponseData) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *BillingInvoiceResponseData) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetExternalId returns the ExternalId field value
func (o *BillingInvoiceResponseData) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *BillingInvoiceResponseData) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *BillingInvoiceResponseData) SetExternalId(v string) {
	o.ExternalId = v
}

// GetId returns the Id field value
func (o *BillingInvoiceResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BillingInvoiceResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BillingInvoiceResponseData) SetId(v string) {
	o.Id = v
}

// GetSubscriptionExternalId returns the SubscriptionExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInvoiceResponseData) GetSubscriptionExternalId() string {
	if o == nil || IsNil(o.SubscriptionExternalId.Get()) {
		var ret string
		return ret
	}
	return *o.SubscriptionExternalId.Get()
}

// GetSubscriptionExternalIdOk returns a tuple with the SubscriptionExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInvoiceResponseData) GetSubscriptionExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionExternalId.Get(), o.SubscriptionExternalId.IsSet()
}

// HasSubscriptionExternalId returns a boolean if a field has been set.
func (o *BillingInvoiceResponseData) HasSubscriptionExternalId() bool {
	if o != nil && o.SubscriptionExternalId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionExternalId gets a reference to the given NullableString and assigns it to the SubscriptionExternalId field.
func (o *BillingInvoiceResponseData) SetSubscriptionExternalId(v string) {
	o.SubscriptionExternalId.Set(&v)
}

// SetSubscriptionExternalIdNil sets the value for SubscriptionExternalId to be an explicit nil
func (o *BillingInvoiceResponseData) SetSubscriptionExternalIdNil() {
	o.SubscriptionExternalId.Set(nil)
}

// UnsetSubscriptionExternalId ensures that no value is present for SubscriptionExternalId, not even an explicit nil
func (o *BillingInvoiceResponseData) UnsetSubscriptionExternalId() {
	o.SubscriptionExternalId.Unset()
}

// GetSubtotal returns the Subtotal field value
func (o *BillingInvoiceResponseData) GetSubtotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Subtotal
}

// GetSubtotalOk returns a tuple with the Subtotal field value
// and a boolean to check if the value has been set.
func (o *BillingInvoiceResponseData) GetSubtotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtotal, true
}

// SetSubtotal sets field value
func (o *BillingInvoiceResponseData) SetSubtotal(v int32) {
	o.Subtotal = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *BillingInvoiceResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *BillingInvoiceResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *BillingInvoiceResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o BillingInvoiceResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingInvoiceResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount_due"] = o.AmountDue
	toSerialize["amount_paid"] = o.AmountPaid
	toSerialize["amount_remaining"] = o.AmountRemaining
	toSerialize["collection_method"] = o.CollectionMethod
	if o.CompanyId.IsSet() {
		toSerialize["company_id"] = o.CompanyId.Get()
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["currency"] = o.Currency
	toSerialize["customer_external_id"] = o.CustomerExternalId
	if o.DueDate.IsSet() {
		toSerialize["due_date"] = o.DueDate.Get()
	}
	toSerialize["environment_id"] = o.EnvironmentId
	toSerialize["external_id"] = o.ExternalId
	toSerialize["id"] = o.Id
	if o.SubscriptionExternalId.IsSet() {
		toSerialize["subscription_external_id"] = o.SubscriptionExternalId.Get()
	}
	toSerialize["subtotal"] = o.Subtotal
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BillingInvoiceResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount_due",
		"amount_paid",
		"amount_remaining",
		"collection_method",
		"created_at",
		"currency",
		"customer_external_id",
		"environment_id",
		"external_id",
		"id",
		"subtotal",
		"updated_at",
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

	varBillingInvoiceResponseData := _BillingInvoiceResponseData{}

	err = json.Unmarshal(data, &varBillingInvoiceResponseData)

	if err != nil {
		return err
	}

	*o = BillingInvoiceResponseData(varBillingInvoiceResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount_due")
		delete(additionalProperties, "amount_paid")
		delete(additionalProperties, "amount_remaining")
		delete(additionalProperties, "collection_method")
		delete(additionalProperties, "company_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "customer_external_id")
		delete(additionalProperties, "due_date")
		delete(additionalProperties, "environment_id")
		delete(additionalProperties, "external_id")
		delete(additionalProperties, "id")
		delete(additionalProperties, "subscription_external_id")
		delete(additionalProperties, "subtotal")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBillingInvoiceResponseData struct {
	value *BillingInvoiceResponseData
	isSet bool
}

func (v NullableBillingInvoiceResponseData) Get() *BillingInvoiceResponseData {
	return v.value
}

func (v *NullableBillingInvoiceResponseData) Set(val *BillingInvoiceResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInvoiceResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInvoiceResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInvoiceResponseData(val *BillingInvoiceResponseData) *NullableBillingInvoiceResponseData {
	return &NullableBillingInvoiceResponseData{value: val, isSet: true}
}

func (v NullableBillingInvoiceResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInvoiceResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}