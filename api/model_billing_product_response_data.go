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
	"time"
)

// checks if the BillingProductResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingProductResponseData{}

// BillingProductResponseData struct for BillingProductResponseData
type BillingProductResponseData struct {
	AccountId     string       `json:"account_id"`
	CreatedAt     time.Time    `json:"created_at"`
	DeletedAt     NullableTime `json:"deleted_at,omitempty"`
	EnvironmentId string       `json:"environment_id"`
	ExternalId    string       `json:"external_id"`
	Name          string       `json:"name"`
	Currency      *string      `json:"currency,omitempty"`
	Price         float32      `json:"price"`
	ProductId     string       `json:"product_id"`
	Quantity      float32      `json:"quantity"`
	UpdatedAt     time.Time    `json:"updated_at"`
}

type _BillingProductResponseData BillingProductResponseData

// NewBillingProductResponseData instantiates a new BillingProductResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingProductResponseData(accountId string, createdAt time.Time, environmentId string, externalId string, name string, price float32, productId string, quantity float32, updatedAt time.Time) *BillingProductResponseData {
	this := BillingProductResponseData{}
	this.AccountId = accountId
	this.CreatedAt = createdAt
	this.EnvironmentId = environmentId
	this.ExternalId = externalId
	this.Name = name
	this.Price = price
	this.ProductId = productId
	this.Quantity = quantity
	this.UpdatedAt = updatedAt
	return &this
}

// NewBillingProductResponseDataWithDefaults instantiates a new BillingProductResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingProductResponseDataWithDefaults() *BillingProductResponseData {
	this := BillingProductResponseData{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *BillingProductResponseData) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *BillingProductResponseData) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *BillingProductResponseData) SetAccountId(v string) {
	o.AccountId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *BillingProductResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BillingProductResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *BillingProductResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingProductResponseData) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingProductResponseData) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *BillingProductResponseData) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *BillingProductResponseData) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *BillingProductResponseData) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *BillingProductResponseData) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *BillingProductResponseData) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *BillingProductResponseData) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *BillingProductResponseData) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetExternalId returns the ExternalId field value
func (o *BillingProductResponseData) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *BillingProductResponseData) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *BillingProductResponseData) SetExternalId(v string) {
	o.ExternalId = v
}

// GetName returns the Name field value
func (o *BillingProductResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BillingProductResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BillingProductResponseData) SetName(v string) {
	o.Name = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *BillingProductResponseData) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingProductResponseData) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *BillingProductResponseData) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *BillingProductResponseData) SetCurrency(v string) {
	o.Currency = &v
}

// GetPrice returns the Price field value
func (o *BillingProductResponseData) GetPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *BillingProductResponseData) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *BillingProductResponseData) SetPrice(v float32) {
	o.Price = v
}

// GetProductId returns the ProductId field value
func (o *BillingProductResponseData) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *BillingProductResponseData) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *BillingProductResponseData) SetProductId(v string) {
	o.ProductId = v
}

// GetQuantity returns the Quantity field value
func (o *BillingProductResponseData) GetQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *BillingProductResponseData) GetQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *BillingProductResponseData) SetQuantity(v float32) {
	o.Quantity = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *BillingProductResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *BillingProductResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *BillingProductResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o BillingProductResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingProductResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	toSerialize["created_at"] = o.CreatedAt
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	toSerialize["environment_id"] = o.EnvironmentId
	toSerialize["external_id"] = o.ExternalId
	toSerialize["name"] = o.Name
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	toSerialize["price"] = o.Price
	toSerialize["product_id"] = o.ProductId
	toSerialize["quantity"] = o.Quantity
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *BillingProductResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_id",
		"created_at",
		"environment_id",
		"external_id",
		"name",
		"price",
		"product_id",
		"quantity",
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

	varBillingProductResponseData := _BillingProductResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBillingProductResponseData)

	if err != nil {
		return err
	}

	*o = BillingProductResponseData(varBillingProductResponseData)

	return err
}

type NullableBillingProductResponseData struct {
	value *BillingProductResponseData
	isSet bool
}

func (v NullableBillingProductResponseData) Get() *BillingProductResponseData {
	return v.value
}

func (v *NullableBillingProductResponseData) Set(val *BillingProductResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingProductResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingProductResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingProductResponseData(val *BillingProductResponseData) *NullableBillingProductResponseData {
	return &NullableBillingProductResponseData{value: val, isSet: true}
}

func (v NullableBillingProductResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingProductResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}