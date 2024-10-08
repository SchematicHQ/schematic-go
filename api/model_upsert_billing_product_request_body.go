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
)

// checks if the UpsertBillingProductRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpsertBillingProductRequestBody{}

// UpsertBillingProductRequestBody struct for UpsertBillingProductRequestBody
type UpsertBillingProductRequestBody struct {
	BillingProductId     string         `json:"billing_product_id"`
	MonthlyPriceId       NullableString `json:"monthly_price_id,omitempty"`
	YearlyPriceId        NullableString `json:"yearly_price_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpsertBillingProductRequestBody UpsertBillingProductRequestBody

// NewUpsertBillingProductRequestBody instantiates a new UpsertBillingProductRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpsertBillingProductRequestBody(billingProductId string) *UpsertBillingProductRequestBody {
	this := UpsertBillingProductRequestBody{}
	this.BillingProductId = billingProductId
	return &this
}

// NewUpsertBillingProductRequestBodyWithDefaults instantiates a new UpsertBillingProductRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpsertBillingProductRequestBodyWithDefaults() *UpsertBillingProductRequestBody {
	this := UpsertBillingProductRequestBody{}
	return &this
}

// GetBillingProductId returns the BillingProductId field value
func (o *UpsertBillingProductRequestBody) GetBillingProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingProductId
}

// GetBillingProductIdOk returns a tuple with the BillingProductId field value
// and a boolean to check if the value has been set.
func (o *UpsertBillingProductRequestBody) GetBillingProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingProductId, true
}

// SetBillingProductId sets field value
func (o *UpsertBillingProductRequestBody) SetBillingProductId(v string) {
	o.BillingProductId = v
}

// GetMonthlyPriceId returns the MonthlyPriceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpsertBillingProductRequestBody) GetMonthlyPriceId() string {
	if o == nil || IsNil(o.MonthlyPriceId.Get()) {
		var ret string
		return ret
	}
	return *o.MonthlyPriceId.Get()
}

// GetMonthlyPriceIdOk returns a tuple with the MonthlyPriceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpsertBillingProductRequestBody) GetMonthlyPriceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonthlyPriceId.Get(), o.MonthlyPriceId.IsSet()
}

// HasMonthlyPriceId returns a boolean if a field has been set.
func (o *UpsertBillingProductRequestBody) HasMonthlyPriceId() bool {
	if o != nil && o.MonthlyPriceId.IsSet() {
		return true
	}

	return false
}

// SetMonthlyPriceId gets a reference to the given NullableString and assigns it to the MonthlyPriceId field.
func (o *UpsertBillingProductRequestBody) SetMonthlyPriceId(v string) {
	o.MonthlyPriceId.Set(&v)
}

// SetMonthlyPriceIdNil sets the value for MonthlyPriceId to be an explicit nil
func (o *UpsertBillingProductRequestBody) SetMonthlyPriceIdNil() {
	o.MonthlyPriceId.Set(nil)
}

// UnsetMonthlyPriceId ensures that no value is present for MonthlyPriceId, not even an explicit nil
func (o *UpsertBillingProductRequestBody) UnsetMonthlyPriceId() {
	o.MonthlyPriceId.Unset()
}

// GetYearlyPriceId returns the YearlyPriceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpsertBillingProductRequestBody) GetYearlyPriceId() string {
	if o == nil || IsNil(o.YearlyPriceId.Get()) {
		var ret string
		return ret
	}
	return *o.YearlyPriceId.Get()
}

// GetYearlyPriceIdOk returns a tuple with the YearlyPriceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpsertBillingProductRequestBody) GetYearlyPriceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.YearlyPriceId.Get(), o.YearlyPriceId.IsSet()
}

// HasYearlyPriceId returns a boolean if a field has been set.
func (o *UpsertBillingProductRequestBody) HasYearlyPriceId() bool {
	if o != nil && o.YearlyPriceId.IsSet() {
		return true
	}

	return false
}

// SetYearlyPriceId gets a reference to the given NullableString and assigns it to the YearlyPriceId field.
func (o *UpsertBillingProductRequestBody) SetYearlyPriceId(v string) {
	o.YearlyPriceId.Set(&v)
}

// SetYearlyPriceIdNil sets the value for YearlyPriceId to be an explicit nil
func (o *UpsertBillingProductRequestBody) SetYearlyPriceIdNil() {
	o.YearlyPriceId.Set(nil)
}

// UnsetYearlyPriceId ensures that no value is present for YearlyPriceId, not even an explicit nil
func (o *UpsertBillingProductRequestBody) UnsetYearlyPriceId() {
	o.YearlyPriceId.Unset()
}

func (o UpsertBillingProductRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpsertBillingProductRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["billing_product_id"] = o.BillingProductId
	if o.MonthlyPriceId.IsSet() {
		toSerialize["monthly_price_id"] = o.MonthlyPriceId.Get()
	}
	if o.YearlyPriceId.IsSet() {
		toSerialize["yearly_price_id"] = o.YearlyPriceId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpsertBillingProductRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"billing_product_id",
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

	varUpsertBillingProductRequestBody := _UpsertBillingProductRequestBody{}

	err = json.Unmarshal(data, &varUpsertBillingProductRequestBody)

	if err != nil {
		return err
	}

	*o = UpsertBillingProductRequestBody(varUpsertBillingProductRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "billing_product_id")
		delete(additionalProperties, "monthly_price_id")
		delete(additionalProperties, "yearly_price_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpsertBillingProductRequestBody struct {
	value *UpsertBillingProductRequestBody
	isSet bool
}

func (v NullableUpsertBillingProductRequestBody) Get() *UpsertBillingProductRequestBody {
	return v.value
}

func (v *NullableUpsertBillingProductRequestBody) Set(val *UpsertBillingProductRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpsertBillingProductRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpsertBillingProductRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpsertBillingProductRequestBody(val *UpsertBillingProductRequestBody) *NullableUpsertBillingProductRequestBody {
	return &NullableUpsertBillingProductRequestBody{value: val, isSet: true}
}

func (v NullableUpsertBillingProductRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpsertBillingProductRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
