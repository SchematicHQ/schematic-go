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

// checks if the CreateBillingSubscriptionsRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBillingSubscriptionsRequestBody{}

// CreateBillingSubscriptionsRequestBody struct for CreateBillingSubscriptionsRequestBody
type CreateBillingSubscriptionsRequestBody struct {
	CustomerExternalId     string                  `json:"customer_external_id"`
	ExpiredAt              time.Time               `json:"expired_at"`
	Interval               NullableString          `json:"interval,omitempty"`
	Metadata               map[string]interface{}  `json:"metadata,omitempty"`
	ProductExternalIds     []BillingProductPricing `json:"product_external_ids"`
	Status                 NullableString          `json:"status,omitempty"`
	SubscriptionExternalId string                  `json:"subscription_external_id"`
	TotalPrice             int32                   `json:"total_price"`
	AdditionalProperties   map[string]interface{}
}

type _CreateBillingSubscriptionsRequestBody CreateBillingSubscriptionsRequestBody

// NewCreateBillingSubscriptionsRequestBody instantiates a new CreateBillingSubscriptionsRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBillingSubscriptionsRequestBody(customerExternalId string, expiredAt time.Time, productExternalIds []BillingProductPricing, subscriptionExternalId string, totalPrice int32) *CreateBillingSubscriptionsRequestBody {
	this := CreateBillingSubscriptionsRequestBody{}
	this.CustomerExternalId = customerExternalId
	this.ExpiredAt = expiredAt
	this.ProductExternalIds = productExternalIds
	this.SubscriptionExternalId = subscriptionExternalId
	this.TotalPrice = totalPrice
	return &this
}

// NewCreateBillingSubscriptionsRequestBodyWithDefaults instantiates a new CreateBillingSubscriptionsRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBillingSubscriptionsRequestBodyWithDefaults() *CreateBillingSubscriptionsRequestBody {
	this := CreateBillingSubscriptionsRequestBody{}
	return &this
}

// GetCustomerExternalId returns the CustomerExternalId field value
func (o *CreateBillingSubscriptionsRequestBody) GetCustomerExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerExternalId
}

// GetCustomerExternalIdOk returns a tuple with the CustomerExternalId field value
// and a boolean to check if the value has been set.
func (o *CreateBillingSubscriptionsRequestBody) GetCustomerExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerExternalId, true
}

// SetCustomerExternalId sets field value
func (o *CreateBillingSubscriptionsRequestBody) SetCustomerExternalId(v string) {
	o.CustomerExternalId = v
}

// GetExpiredAt returns the ExpiredAt field value
func (o *CreateBillingSubscriptionsRequestBody) GetExpiredAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value
// and a boolean to check if the value has been set.
func (o *CreateBillingSubscriptionsRequestBody) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiredAt, true
}

// SetExpiredAt sets field value
func (o *CreateBillingSubscriptionsRequestBody) SetExpiredAt(v time.Time) {
	o.ExpiredAt = v
}

// GetInterval returns the Interval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateBillingSubscriptionsRequestBody) GetInterval() string {
	if o == nil || IsNil(o.Interval.Get()) {
		var ret string
		return ret
	}
	return *o.Interval.Get()
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateBillingSubscriptionsRequestBody) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interval.Get(), o.Interval.IsSet()
}

// HasInterval returns a boolean if a field has been set.
func (o *CreateBillingSubscriptionsRequestBody) HasInterval() bool {
	if o != nil && o.Interval.IsSet() {
		return true
	}

	return false
}

// SetInterval gets a reference to the given NullableString and assigns it to the Interval field.
func (o *CreateBillingSubscriptionsRequestBody) SetInterval(v string) {
	o.Interval.Set(&v)
}

// SetIntervalNil sets the value for Interval to be an explicit nil
func (o *CreateBillingSubscriptionsRequestBody) SetIntervalNil() {
	o.Interval.Set(nil)
}

// UnsetInterval ensures that no value is present for Interval, not even an explicit nil
func (o *CreateBillingSubscriptionsRequestBody) UnsetInterval() {
	o.Interval.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateBillingSubscriptionsRequestBody) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateBillingSubscriptionsRequestBody) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateBillingSubscriptionsRequestBody) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CreateBillingSubscriptionsRequestBody) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetProductExternalIds returns the ProductExternalIds field value
func (o *CreateBillingSubscriptionsRequestBody) GetProductExternalIds() []BillingProductPricing {
	if o == nil {
		var ret []BillingProductPricing
		return ret
	}

	return o.ProductExternalIds
}

// GetProductExternalIdsOk returns a tuple with the ProductExternalIds field value
// and a boolean to check if the value has been set.
func (o *CreateBillingSubscriptionsRequestBody) GetProductExternalIdsOk() ([]BillingProductPricing, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductExternalIds, true
}

// SetProductExternalIds sets field value
func (o *CreateBillingSubscriptionsRequestBody) SetProductExternalIds(v []BillingProductPricing) {
	o.ProductExternalIds = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateBillingSubscriptionsRequestBody) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateBillingSubscriptionsRequestBody) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateBillingSubscriptionsRequestBody) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *CreateBillingSubscriptionsRequestBody) SetStatus(v string) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *CreateBillingSubscriptionsRequestBody) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *CreateBillingSubscriptionsRequestBody) UnsetStatus() {
	o.Status.Unset()
}

// GetSubscriptionExternalId returns the SubscriptionExternalId field value
func (o *CreateBillingSubscriptionsRequestBody) GetSubscriptionExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionExternalId
}

// GetSubscriptionExternalIdOk returns a tuple with the SubscriptionExternalId field value
// and a boolean to check if the value has been set.
func (o *CreateBillingSubscriptionsRequestBody) GetSubscriptionExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionExternalId, true
}

// SetSubscriptionExternalId sets field value
func (o *CreateBillingSubscriptionsRequestBody) SetSubscriptionExternalId(v string) {
	o.SubscriptionExternalId = v
}

// GetTotalPrice returns the TotalPrice field value
func (o *CreateBillingSubscriptionsRequestBody) GetTotalPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalPrice
}

// GetTotalPriceOk returns a tuple with the TotalPrice field value
// and a boolean to check if the value has been set.
func (o *CreateBillingSubscriptionsRequestBody) GetTotalPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPrice, true
}

// SetTotalPrice sets field value
func (o *CreateBillingSubscriptionsRequestBody) SetTotalPrice(v int32) {
	o.TotalPrice = v
}

func (o CreateBillingSubscriptionsRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBillingSubscriptionsRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customer_external_id"] = o.CustomerExternalId
	toSerialize["expired_at"] = o.ExpiredAt
	if o.Interval.IsSet() {
		toSerialize["interval"] = o.Interval.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["product_external_ids"] = o.ProductExternalIds
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	toSerialize["subscription_external_id"] = o.SubscriptionExternalId
	toSerialize["total_price"] = o.TotalPrice

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateBillingSubscriptionsRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customer_external_id",
		"expired_at",
		"product_external_ids",
		"subscription_external_id",
		"total_price",
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

	varCreateBillingSubscriptionsRequestBody := _CreateBillingSubscriptionsRequestBody{}

	err = json.Unmarshal(data, &varCreateBillingSubscriptionsRequestBody)

	if err != nil {
		return err
	}

	*o = CreateBillingSubscriptionsRequestBody(varCreateBillingSubscriptionsRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customer_external_id")
		delete(additionalProperties, "expired_at")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "product_external_ids")
		delete(additionalProperties, "status")
		delete(additionalProperties, "subscription_external_id")
		delete(additionalProperties, "total_price")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateBillingSubscriptionsRequestBody struct {
	value *CreateBillingSubscriptionsRequestBody
	isSet bool
}

func (v NullableCreateBillingSubscriptionsRequestBody) Get() *CreateBillingSubscriptionsRequestBody {
	return v.value
}

func (v *NullableCreateBillingSubscriptionsRequestBody) Set(val *CreateBillingSubscriptionsRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBillingSubscriptionsRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBillingSubscriptionsRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBillingSubscriptionsRequestBody(val *CreateBillingSubscriptionsRequestBody) *NullableCreateBillingSubscriptionsRequestBody {
	return &NullableCreateBillingSubscriptionsRequestBody{value: val, isSet: true}
}

func (v NullableCreateBillingSubscriptionsRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBillingSubscriptionsRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}