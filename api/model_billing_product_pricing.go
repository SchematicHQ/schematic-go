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

// checks if the BillingProductPricing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingProductPricing{}

// BillingProductPricing struct for BillingProductPricing
type BillingProductPricing struct {
	Price                int32  `json:"price"`
	ProductExternalId    string `json:"product_external_id"`
	AdditionalProperties map[string]interface{}
}

type _BillingProductPricing BillingProductPricing

// NewBillingProductPricing instantiates a new BillingProductPricing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingProductPricing(price int32, productExternalId string) *BillingProductPricing {
	this := BillingProductPricing{}
	this.Price = price
	this.ProductExternalId = productExternalId
	return &this
}

// NewBillingProductPricingWithDefaults instantiates a new BillingProductPricing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingProductPricingWithDefaults() *BillingProductPricing {
	this := BillingProductPricing{}
	return &this
}

// GetPrice returns the Price field value
func (o *BillingProductPricing) GetPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *BillingProductPricing) GetPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *BillingProductPricing) SetPrice(v int32) {
	o.Price = v
}

// GetProductExternalId returns the ProductExternalId field value
func (o *BillingProductPricing) GetProductExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductExternalId
}

// GetProductExternalIdOk returns a tuple with the ProductExternalId field value
// and a boolean to check if the value has been set.
func (o *BillingProductPricing) GetProductExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductExternalId, true
}

// SetProductExternalId sets field value
func (o *BillingProductPricing) SetProductExternalId(v string) {
	o.ProductExternalId = v
}

func (o BillingProductPricing) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingProductPricing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["price"] = o.Price
	toSerialize["product_external_id"] = o.ProductExternalId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BillingProductPricing) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"price",
		"product_external_id",
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

	varBillingProductPricing := _BillingProductPricing{}

	err = json.Unmarshal(data, &varBillingProductPricing)

	if err != nil {
		return err
	}

	*o = BillingProductPricing(varBillingProductPricing)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "price")
		delete(additionalProperties, "product_external_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBillingProductPricing struct {
	value *BillingProductPricing
	isSet bool
}

func (v NullableBillingProductPricing) Get() *BillingProductPricing {
	return v.value
}

func (v *NullableBillingProductPricing) Set(val *BillingProductPricing) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingProductPricing) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingProductPricing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingProductPricing(val *BillingProductPricing) *NullableBillingProductPricing {
	return &NullableBillingProductPricing{value: val, isSet: true}
}

func (v NullableBillingProductPricing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingProductPricing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}