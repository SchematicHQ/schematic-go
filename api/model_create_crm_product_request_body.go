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

// checks if the CreateCrmProductRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCrmProductRequestBody{}

// CreateCrmProductRequestBody struct for CreateCrmProductRequestBody
type CreateCrmProductRequestBody struct {
	Currency             string `json:"currency"`
	Description          string `json:"description"`
	ExternalId           string `json:"external_id"`
	Interval             string `json:"interval"`
	Name                 string `json:"name"`
	Price                string `json:"price"`
	Quantity             int32  `json:"quantity"`
	Sku                  string `json:"sku"`
	AdditionalProperties map[string]interface{}
}

type _CreateCrmProductRequestBody CreateCrmProductRequestBody

// NewCreateCrmProductRequestBody instantiates a new CreateCrmProductRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCrmProductRequestBody(currency string, description string, externalId string, interval string, name string, price string, quantity int32, sku string) *CreateCrmProductRequestBody {
	this := CreateCrmProductRequestBody{}
	this.Currency = currency
	this.Description = description
	this.ExternalId = externalId
	this.Interval = interval
	this.Name = name
	this.Price = price
	this.Quantity = quantity
	this.Sku = sku
	return &this
}

// NewCreateCrmProductRequestBodyWithDefaults instantiates a new CreateCrmProductRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCrmProductRequestBodyWithDefaults() *CreateCrmProductRequestBody {
	this := CreateCrmProductRequestBody{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *CreateCrmProductRequestBody) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CreateCrmProductRequestBody) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CreateCrmProductRequestBody) SetCurrency(v string) {
	o.Currency = v
}

// GetDescription returns the Description field value
func (o *CreateCrmProductRequestBody) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateCrmProductRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateCrmProductRequestBody) SetDescription(v string) {
	o.Description = v
}

// GetExternalId returns the ExternalId field value
func (o *CreateCrmProductRequestBody) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *CreateCrmProductRequestBody) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *CreateCrmProductRequestBody) SetExternalId(v string) {
	o.ExternalId = v
}

// GetInterval returns the Interval field value
func (o *CreateCrmProductRequestBody) GetInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *CreateCrmProductRequestBody) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *CreateCrmProductRequestBody) SetInterval(v string) {
	o.Interval = v
}

// GetName returns the Name field value
func (o *CreateCrmProductRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateCrmProductRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateCrmProductRequestBody) SetName(v string) {
	o.Name = v
}

// GetPrice returns the Price field value
func (o *CreateCrmProductRequestBody) GetPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *CreateCrmProductRequestBody) GetPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *CreateCrmProductRequestBody) SetPrice(v string) {
	o.Price = v
}

// GetQuantity returns the Quantity field value
func (o *CreateCrmProductRequestBody) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *CreateCrmProductRequestBody) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *CreateCrmProductRequestBody) SetQuantity(v int32) {
	o.Quantity = v
}

// GetSku returns the Sku field value
func (o *CreateCrmProductRequestBody) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *CreateCrmProductRequestBody) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *CreateCrmProductRequestBody) SetSku(v string) {
	o.Sku = v
}

func (o CreateCrmProductRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCrmProductRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	toSerialize["description"] = o.Description
	toSerialize["external_id"] = o.ExternalId
	toSerialize["interval"] = o.Interval
	toSerialize["name"] = o.Name
	toSerialize["price"] = o.Price
	toSerialize["quantity"] = o.Quantity
	toSerialize["sku"] = o.Sku

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateCrmProductRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency",
		"description",
		"external_id",
		"interval",
		"name",
		"price",
		"quantity",
		"sku",
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

	varCreateCrmProductRequestBody := _CreateCrmProductRequestBody{}

	err = json.Unmarshal(data, &varCreateCrmProductRequestBody)

	if err != nil {
		return err
	}

	*o = CreateCrmProductRequestBody(varCreateCrmProductRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "currency")
		delete(additionalProperties, "description")
		delete(additionalProperties, "external_id")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "name")
		delete(additionalProperties, "price")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "sku")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateCrmProductRequestBody struct {
	value *CreateCrmProductRequestBody
	isSet bool
}

func (v NullableCreateCrmProductRequestBody) Get() *CreateCrmProductRequestBody {
	return v.value
}

func (v *NullableCreateCrmProductRequestBody) Set(val *CreateCrmProductRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCrmProductRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCrmProductRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCrmProductRequestBody(val *CreateCrmProductRequestBody) *NullableCreateCrmProductRequestBody {
	return &NullableCreateCrmProductRequestBody{value: val, isSet: true}
}

func (v NullableCreateCrmProductRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCrmProductRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}