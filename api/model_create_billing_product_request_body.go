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

// checks if the CreateBillingProductRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBillingProductRequestBody{}

// CreateBillingProductRequestBody struct for CreateBillingProductRequestBody
type CreateBillingProductRequestBody struct {
	Currency             string  `json:"currency"`
	ExternalId           string  `json:"external_id"`
	Name                 string  `json:"name"`
	Price                float32 `json:"price"`
	Quantity             int32   `json:"quantity"`
	AdditionalProperties map[string]interface{}
}

type _CreateBillingProductRequestBody CreateBillingProductRequestBody

// NewCreateBillingProductRequestBody instantiates a new CreateBillingProductRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBillingProductRequestBody(currency string, externalId string, name string, price float32, quantity int32) *CreateBillingProductRequestBody {
	this := CreateBillingProductRequestBody{}
	this.Currency = currency
	this.ExternalId = externalId
	this.Name = name
	this.Price = price
	this.Quantity = quantity
	return &this
}

// NewCreateBillingProductRequestBodyWithDefaults instantiates a new CreateBillingProductRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBillingProductRequestBodyWithDefaults() *CreateBillingProductRequestBody {
	this := CreateBillingProductRequestBody{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *CreateBillingProductRequestBody) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CreateBillingProductRequestBody) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CreateBillingProductRequestBody) SetCurrency(v string) {
	o.Currency = v
}

// GetExternalId returns the ExternalId field value
func (o *CreateBillingProductRequestBody) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *CreateBillingProductRequestBody) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *CreateBillingProductRequestBody) SetExternalId(v string) {
	o.ExternalId = v
}

// GetName returns the Name field value
func (o *CreateBillingProductRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateBillingProductRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateBillingProductRequestBody) SetName(v string) {
	o.Name = v
}

// GetPrice returns the Price field value
func (o *CreateBillingProductRequestBody) GetPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *CreateBillingProductRequestBody) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *CreateBillingProductRequestBody) SetPrice(v float32) {
	o.Price = v
}

// GetQuantity returns the Quantity field value
func (o *CreateBillingProductRequestBody) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *CreateBillingProductRequestBody) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *CreateBillingProductRequestBody) SetQuantity(v int32) {
	o.Quantity = v
}

func (o CreateBillingProductRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBillingProductRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	toSerialize["external_id"] = o.ExternalId
	toSerialize["name"] = o.Name
	toSerialize["price"] = o.Price
	toSerialize["quantity"] = o.Quantity

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateBillingProductRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency",
		"external_id",
		"name",
		"price",
		"quantity",
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

	varCreateBillingProductRequestBody := _CreateBillingProductRequestBody{}

	err = json.Unmarshal(data, &varCreateBillingProductRequestBody)

	if err != nil {
		return err
	}

	*o = CreateBillingProductRequestBody(varCreateBillingProductRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "currency")
		delete(additionalProperties, "external_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "price")
		delete(additionalProperties, "quantity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateBillingProductRequestBody struct {
	value *CreateBillingProductRequestBody
	isSet bool
}

func (v NullableCreateBillingProductRequestBody) Get() *CreateBillingProductRequestBody {
	return v.value
}

func (v *NullableCreateBillingProductRequestBody) Set(val *CreateBillingProductRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBillingProductRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBillingProductRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBillingProductRequestBody(val *CreateBillingProductRequestBody) *NullableCreateBillingProductRequestBody {
	return &NullableCreateBillingProductRequestBody{value: val, isSet: true}
}

func (v NullableCreateBillingProductRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBillingProductRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
