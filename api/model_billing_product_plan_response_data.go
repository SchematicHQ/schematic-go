/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the BillingProductPlanResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingProductPlanResponseData{}

// BillingProductPlanResponseData The updated resource
type BillingProductPlanResponseData struct {
	AccountId        string `json:"account_id"`
	BillingProductId string `json:"billing_product_id"`
	EnvironmentId    string `json:"environment_id"`
	PlanId           string `json:"plan_id"`
}

type _BillingProductPlanResponseData BillingProductPlanResponseData

// NewBillingProductPlanResponseData instantiates a new BillingProductPlanResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingProductPlanResponseData(accountId string, billingProductId string, environmentId string, planId string) *BillingProductPlanResponseData {
	this := BillingProductPlanResponseData{}
	this.AccountId = accountId
	this.BillingProductId = billingProductId
	this.EnvironmentId = environmentId
	this.PlanId = planId
	return &this
}

// NewBillingProductPlanResponseDataWithDefaults instantiates a new BillingProductPlanResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingProductPlanResponseDataWithDefaults() *BillingProductPlanResponseData {
	this := BillingProductPlanResponseData{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *BillingProductPlanResponseData) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *BillingProductPlanResponseData) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *BillingProductPlanResponseData) SetAccountId(v string) {
	o.AccountId = v
}

// GetBillingProductId returns the BillingProductId field value
func (o *BillingProductPlanResponseData) GetBillingProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingProductId
}

// GetBillingProductIdOk returns a tuple with the BillingProductId field value
// and a boolean to check if the value has been set.
func (o *BillingProductPlanResponseData) GetBillingProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingProductId, true
}

// SetBillingProductId sets field value
func (o *BillingProductPlanResponseData) SetBillingProductId(v string) {
	o.BillingProductId = v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *BillingProductPlanResponseData) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *BillingProductPlanResponseData) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *BillingProductPlanResponseData) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetPlanId returns the PlanId field value
func (o *BillingProductPlanResponseData) GetPlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *BillingProductPlanResponseData) GetPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *BillingProductPlanResponseData) SetPlanId(v string) {
	o.PlanId = v
}

func (o BillingProductPlanResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingProductPlanResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	toSerialize["billing_product_id"] = o.BillingProductId
	toSerialize["environment_id"] = o.EnvironmentId
	toSerialize["plan_id"] = o.PlanId
	return toSerialize, nil
}

func (o *BillingProductPlanResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_id",
		"billing_product_id",
		"environment_id",
		"plan_id",
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

	varBillingProductPlanResponseData := _BillingProductPlanResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBillingProductPlanResponseData)

	if err != nil {
		return err
	}

	*o = BillingProductPlanResponseData(varBillingProductPlanResponseData)

	return err
}

type NullableBillingProductPlanResponseData struct {
	value *BillingProductPlanResponseData
	isSet bool
}

func (v NullableBillingProductPlanResponseData) Get() *BillingProductPlanResponseData {
	return v.value
}

func (v *NullableBillingProductPlanResponseData) Set(val *BillingProductPlanResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingProductPlanResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingProductPlanResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingProductPlanResponseData(val *BillingProductPlanResponseData) *NullableBillingProductPlanResponseData {
	return &NullableBillingProductPlanResponseData{value: val, isSet: true}
}

func (v NullableBillingProductPlanResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingProductPlanResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
