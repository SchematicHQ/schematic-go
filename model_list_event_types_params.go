/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schematic

import (
	"encoding/json"
)

// checks if the ListEventTypesParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListEventTypesParams{}

// ListEventTypesParams Input parameters
type ListEventTypesParams struct {
	Limit *int32 `json:"limit,omitempty"`
	Offset *int32 `json:"offset,omitempty"`
	Q *string `json:"q,omitempty"`
}

// NewListEventTypesParams instantiates a new ListEventTypesParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEventTypesParams() *ListEventTypesParams {
	this := ListEventTypesParams{}
	return &this
}

// NewListEventTypesParamsWithDefaults instantiates a new ListEventTypesParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEventTypesParamsWithDefaults() *ListEventTypesParams {
	this := ListEventTypesParams{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ListEventTypesParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEventTypesParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListEventTypesParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ListEventTypesParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ListEventTypesParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEventTypesParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ListEventTypesParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *ListEventTypesParams) SetOffset(v int32) {
	o.Offset = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *ListEventTypesParams) GetQ() string {
	if o == nil || IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEventTypesParams) GetQOk() (*string, bool) {
	if o == nil || IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *ListEventTypesParams) HasQ() bool {
	if o != nil && !IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *ListEventTypesParams) SetQ(v string) {
	o.Q = &v
}

func (o ListEventTypesParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListEventTypesParams) ToMap() (map[string]interface{}, error) {
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

type NullableListEventTypesParams struct {
	value *ListEventTypesParams
	isSet bool
}

func (v NullableListEventTypesParams) Get() *ListEventTypesParams {
	return v.value
}

func (v *NullableListEventTypesParams) Set(val *ListEventTypesParams) {
	v.value = val
	v.isSet = true
}

func (v NullableListEventTypesParams) IsSet() bool {
	return v.isSet
}

func (v *NullableListEventTypesParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEventTypesParams(val *ListEventTypesParams) *NullableListEventTypesParams {
	return &NullableListEventTypesParams{value: val, isSet: true}
}

func (v NullableListEventTypesParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEventTypesParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


