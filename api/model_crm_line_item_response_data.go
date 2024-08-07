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

// checks if the CrmLineItemResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CrmLineItemResponseData{}

// CrmLineItemResponseData The created resource
type CrmLineItemResponseData struct {
	AccountId            string         `json:"account_id"`
	CreatedAt            time.Time      `json:"created_at"`
	DealId               NullableString `json:"deal_id,omitempty"`
	EnvironmentId        string         `json:"environment_id"`
	ProductExternalId    NullableString `json:"product_external_id,omitempty"`
	UpdatedAt            time.Time      `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _CrmLineItemResponseData CrmLineItemResponseData

// NewCrmLineItemResponseData instantiates a new CrmLineItemResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrmLineItemResponseData(accountId string, createdAt time.Time, environmentId string, updatedAt time.Time) *CrmLineItemResponseData {
	this := CrmLineItemResponseData{}
	this.AccountId = accountId
	this.CreatedAt = createdAt
	this.EnvironmentId = environmentId
	this.UpdatedAt = updatedAt
	return &this
}

// NewCrmLineItemResponseDataWithDefaults instantiates a new CrmLineItemResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrmLineItemResponseDataWithDefaults() *CrmLineItemResponseData {
	this := CrmLineItemResponseData{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *CrmLineItemResponseData) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *CrmLineItemResponseData) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *CrmLineItemResponseData) SetAccountId(v string) {
	o.AccountId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CrmLineItemResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CrmLineItemResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CrmLineItemResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDealId returns the DealId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CrmLineItemResponseData) GetDealId() string {
	if o == nil || IsNil(o.DealId.Get()) {
		var ret string
		return ret
	}
	return *o.DealId.Get()
}

// GetDealIdOk returns a tuple with the DealId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CrmLineItemResponseData) GetDealIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DealId.Get(), o.DealId.IsSet()
}

// HasDealId returns a boolean if a field has been set.
func (o *CrmLineItemResponseData) HasDealId() bool {
	if o != nil && o.DealId.IsSet() {
		return true
	}

	return false
}

// SetDealId gets a reference to the given NullableString and assigns it to the DealId field.
func (o *CrmLineItemResponseData) SetDealId(v string) {
	o.DealId.Set(&v)
}

// SetDealIdNil sets the value for DealId to be an explicit nil
func (o *CrmLineItemResponseData) SetDealIdNil() {
	o.DealId.Set(nil)
}

// UnsetDealId ensures that no value is present for DealId, not even an explicit nil
func (o *CrmLineItemResponseData) UnsetDealId() {
	o.DealId.Unset()
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *CrmLineItemResponseData) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *CrmLineItemResponseData) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *CrmLineItemResponseData) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetProductExternalId returns the ProductExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CrmLineItemResponseData) GetProductExternalId() string {
	if o == nil || IsNil(o.ProductExternalId.Get()) {
		var ret string
		return ret
	}
	return *o.ProductExternalId.Get()
}

// GetProductExternalIdOk returns a tuple with the ProductExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CrmLineItemResponseData) GetProductExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductExternalId.Get(), o.ProductExternalId.IsSet()
}

// HasProductExternalId returns a boolean if a field has been set.
func (o *CrmLineItemResponseData) HasProductExternalId() bool {
	if o != nil && o.ProductExternalId.IsSet() {
		return true
	}

	return false
}

// SetProductExternalId gets a reference to the given NullableString and assigns it to the ProductExternalId field.
func (o *CrmLineItemResponseData) SetProductExternalId(v string) {
	o.ProductExternalId.Set(&v)
}

// SetProductExternalIdNil sets the value for ProductExternalId to be an explicit nil
func (o *CrmLineItemResponseData) SetProductExternalIdNil() {
	o.ProductExternalId.Set(nil)
}

// UnsetProductExternalId ensures that no value is present for ProductExternalId, not even an explicit nil
func (o *CrmLineItemResponseData) UnsetProductExternalId() {
	o.ProductExternalId.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CrmLineItemResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CrmLineItemResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CrmLineItemResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o CrmLineItemResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CrmLineItemResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	toSerialize["created_at"] = o.CreatedAt
	if o.DealId.IsSet() {
		toSerialize["deal_id"] = o.DealId.Get()
	}
	toSerialize["environment_id"] = o.EnvironmentId
	if o.ProductExternalId.IsSet() {
		toSerialize["product_external_id"] = o.ProductExternalId.Get()
	}
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CrmLineItemResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_id",
		"created_at",
		"environment_id",
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

	varCrmLineItemResponseData := _CrmLineItemResponseData{}

	err = json.Unmarshal(data, &varCrmLineItemResponseData)

	if err != nil {
		return err
	}

	*o = CrmLineItemResponseData(varCrmLineItemResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deal_id")
		delete(additionalProperties, "environment_id")
		delete(additionalProperties, "product_external_id")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCrmLineItemResponseData struct {
	value *CrmLineItemResponseData
	isSet bool
}

func (v NullableCrmLineItemResponseData) Get() *CrmLineItemResponseData {
	return v.value
}

func (v *NullableCrmLineItemResponseData) Set(val *CrmLineItemResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCrmLineItemResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCrmLineItemResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrmLineItemResponseData(val *CrmLineItemResponseData) *NullableCrmLineItemResponseData {
	return &NullableCrmLineItemResponseData{value: val, isSet: true}
}

func (v NullableCrmLineItemResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrmLineItemResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
