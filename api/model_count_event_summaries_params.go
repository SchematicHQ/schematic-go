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

// checks if the CountEventSummariesParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountEventSummariesParams{}

// CountEventSummariesParams Input parameters
type CountEventSummariesParams struct {
	Limit *int32 `json:"limit,omitempty"`
	Offset *int32 `json:"offset,omitempty"`
	Q *string `json:"q,omitempty"`
}

// NewCountEventSummariesParams instantiates a new CountEventSummariesParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountEventSummariesParams() *CountEventSummariesParams {
	this := CountEventSummariesParams{}
	return &this
}

// NewCountEventSummariesParamsWithDefaults instantiates a new CountEventSummariesParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountEventSummariesParamsWithDefaults() *CountEventSummariesParams {
	this := CountEventSummariesParams{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *CountEventSummariesParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountEventSummariesParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *CountEventSummariesParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *CountEventSummariesParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *CountEventSummariesParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountEventSummariesParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *CountEventSummariesParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *CountEventSummariesParams) SetOffset(v int32) {
	o.Offset = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *CountEventSummariesParams) GetQ() string {
	if o == nil || IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountEventSummariesParams) GetQOk() (*string, bool) {
	if o == nil || IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *CountEventSummariesParams) HasQ() bool {
	if o != nil && !IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *CountEventSummariesParams) SetQ(v string) {
	o.Q = &v
}

func (o CountEventSummariesParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountEventSummariesParams) ToMap() (map[string]interface{}, error) {
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

type NullableCountEventSummariesParams struct {
	value *CountEventSummariesParams
	isSet bool
}

func (v NullableCountEventSummariesParams) Get() *CountEventSummariesParams {
	return v.value
}

func (v *NullableCountEventSummariesParams) Set(val *CountEventSummariesParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCountEventSummariesParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCountEventSummariesParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountEventSummariesParams(val *CountEventSummariesParams) *NullableCountEventSummariesParams {
	return &NullableCountEventSummariesParams{value: val, isSet: true}
}

func (v NullableCountEventSummariesParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountEventSummariesParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

