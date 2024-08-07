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

// checks if the BillingSubscriptionResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingSubscriptionResponseData{}

// BillingSubscriptionResponseData The created resource
type BillingSubscriptionResponseData struct {
	ExpiredAt            NullableTime `json:"expired_at,omitempty"`
	ExternalId           string       `json:"external_id"`
	Id                   int32        `json:"id"`
	UpdatedAt            time.Time    `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _BillingSubscriptionResponseData BillingSubscriptionResponseData

// NewBillingSubscriptionResponseData instantiates a new BillingSubscriptionResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingSubscriptionResponseData(externalId string, id int32, updatedAt time.Time) *BillingSubscriptionResponseData {
	this := BillingSubscriptionResponseData{}
	this.ExternalId = externalId
	this.Id = id
	this.UpdatedAt = updatedAt
	return &this
}

// NewBillingSubscriptionResponseDataWithDefaults instantiates a new BillingSubscriptionResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingSubscriptionResponseDataWithDefaults() *BillingSubscriptionResponseData {
	this := BillingSubscriptionResponseData{}
	return &this
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingSubscriptionResponseData) GetExpiredAt() time.Time {
	if o == nil || IsNil(o.ExpiredAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpiredAt.Get()
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingSubscriptionResponseData) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiredAt.Get(), o.ExpiredAt.IsSet()
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *BillingSubscriptionResponseData) HasExpiredAt() bool {
	if o != nil && o.ExpiredAt.IsSet() {
		return true
	}

	return false
}

// SetExpiredAt gets a reference to the given NullableTime and assigns it to the ExpiredAt field.
func (o *BillingSubscriptionResponseData) SetExpiredAt(v time.Time) {
	o.ExpiredAt.Set(&v)
}

// SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil
func (o *BillingSubscriptionResponseData) SetExpiredAtNil() {
	o.ExpiredAt.Set(nil)
}

// UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
func (o *BillingSubscriptionResponseData) UnsetExpiredAt() {
	o.ExpiredAt.Unset()
}

// GetExternalId returns the ExternalId field value
func (o *BillingSubscriptionResponseData) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *BillingSubscriptionResponseData) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *BillingSubscriptionResponseData) SetExternalId(v string) {
	o.ExternalId = v
}

// GetId returns the Id field value
func (o *BillingSubscriptionResponseData) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BillingSubscriptionResponseData) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BillingSubscriptionResponseData) SetId(v int32) {
	o.Id = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *BillingSubscriptionResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *BillingSubscriptionResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *BillingSubscriptionResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o BillingSubscriptionResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingSubscriptionResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiredAt.IsSet() {
		toSerialize["expired_at"] = o.ExpiredAt.Get()
	}
	toSerialize["external_id"] = o.ExternalId
	toSerialize["id"] = o.Id
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BillingSubscriptionResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"external_id",
		"id",
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

	varBillingSubscriptionResponseData := _BillingSubscriptionResponseData{}

	err = json.Unmarshal(data, &varBillingSubscriptionResponseData)

	if err != nil {
		return err
	}

	*o = BillingSubscriptionResponseData(varBillingSubscriptionResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expired_at")
		delete(additionalProperties, "external_id")
		delete(additionalProperties, "id")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBillingSubscriptionResponseData struct {
	value *BillingSubscriptionResponseData
	isSet bool
}

func (v NullableBillingSubscriptionResponseData) Get() *BillingSubscriptionResponseData {
	return v.value
}

func (v *NullableBillingSubscriptionResponseData) Set(val *BillingSubscriptionResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingSubscriptionResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingSubscriptionResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingSubscriptionResponseData(val *BillingSubscriptionResponseData) *NullableBillingSubscriptionResponseData {
	return &NullableBillingSubscriptionResponseData{value: val, isSet: true}
}

func (v NullableBillingSubscriptionResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingSubscriptionResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
