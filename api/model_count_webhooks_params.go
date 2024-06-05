/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CountWebhooksParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountWebhooksParams{}

// CountWebhooksParams Input parameters
type CountWebhooksParams struct {
	Limit  *int32  `json:"limit,omitempty"`
	Offset *int32  `json:"offset,omitempty"`
	Q      *string `json:"q,omitempty"`
}

// NewCountWebhooksParams instantiates a new CountWebhooksParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountWebhooksParams() *CountWebhooksParams {
	this := CountWebhooksParams{}
	return &this
}

// NewCountWebhooksParamsWithDefaults instantiates a new CountWebhooksParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountWebhooksParamsWithDefaults() *CountWebhooksParams {
	this := CountWebhooksParams{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *CountWebhooksParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountWebhooksParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *CountWebhooksParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *CountWebhooksParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *CountWebhooksParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountWebhooksParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *CountWebhooksParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *CountWebhooksParams) SetOffset(v int32) {
	o.Offset = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *CountWebhooksParams) GetQ() string {
	if o == nil || IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountWebhooksParams) GetQOk() (*string, bool) {
	if o == nil || IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *CountWebhooksParams) HasQ() bool {
	if o != nil && !IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *CountWebhooksParams) SetQ(v string) {
	o.Q = &v
}

func (o CountWebhooksParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountWebhooksParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Q) {
		toSerialize["q"] = o.Q
	}
	return toSerialize, nil
}

type NullableCountWebhooksParams struct {
	value *CountWebhooksParams
	isSet bool
}

func (v NullableCountWebhooksParams) Get() *CountWebhooksParams {
	return v.value
}

func (v *NullableCountWebhooksParams) Set(val *CountWebhooksParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCountWebhooksParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCountWebhooksParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountWebhooksParams(val *CountWebhooksParams) *NullableCountWebhooksParams {
	return &NullableCountWebhooksParams{value: val, isSet: true}
}

func (v NullableCountWebhooksParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountWebhooksParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
