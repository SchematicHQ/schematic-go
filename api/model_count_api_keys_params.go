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

// checks if the CountApiKeysParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountApiKeysParams{}

// CountApiKeysParams Input parameters
type CountApiKeysParams struct {
	EnvironmentId *string `json:"environment_id,omitempty"`
	// Page limit (default 100)
	Limit *int32 `json:"limit,omitempty"`
	// Page offset (default 0)
	Offset               *int32 `json:"offset,omitempty"`
	RequireEnvironment   *bool  `json:"require_environment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CountApiKeysParams CountApiKeysParams

// NewCountApiKeysParams instantiates a new CountApiKeysParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountApiKeysParams() *CountApiKeysParams {
	this := CountApiKeysParams{}
	return &this
}

// NewCountApiKeysParamsWithDefaults instantiates a new CountApiKeysParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountApiKeysParamsWithDefaults() *CountApiKeysParams {
	this := CountApiKeysParams{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *CountApiKeysParams) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountApiKeysParams) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *CountApiKeysParams) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *CountApiKeysParams) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *CountApiKeysParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountApiKeysParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *CountApiKeysParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *CountApiKeysParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *CountApiKeysParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountApiKeysParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *CountApiKeysParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *CountApiKeysParams) SetOffset(v int32) {
	o.Offset = &v
}

// GetRequireEnvironment returns the RequireEnvironment field value if set, zero value otherwise.
func (o *CountApiKeysParams) GetRequireEnvironment() bool {
	if o == nil || IsNil(o.RequireEnvironment) {
		var ret bool
		return ret
	}
	return *o.RequireEnvironment
}

// GetRequireEnvironmentOk returns a tuple with the RequireEnvironment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountApiKeysParams) GetRequireEnvironmentOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireEnvironment) {
		return nil, false
	}
	return o.RequireEnvironment, true
}

// HasRequireEnvironment returns a boolean if a field has been set.
func (o *CountApiKeysParams) HasRequireEnvironment() bool {
	if o != nil && !IsNil(o.RequireEnvironment) {
		return true
	}

	return false
}

// SetRequireEnvironment gets a reference to the given bool and assigns it to the RequireEnvironment field.
func (o *CountApiKeysParams) SetRequireEnvironment(v bool) {
	o.RequireEnvironment = &v
}

func (o CountApiKeysParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountApiKeysParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnvironmentId) {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.RequireEnvironment) {
		toSerialize["require_environment"] = o.RequireEnvironment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CountApiKeysParams) UnmarshalJSON(data []byte) (err error) {
	varCountApiKeysParams := _CountApiKeysParams{}

	err = json.Unmarshal(data, &varCountApiKeysParams)

	if err != nil {
		return err
	}

	*o = CountApiKeysParams(varCountApiKeysParams)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "environment_id")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "require_environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCountApiKeysParams struct {
	value *CountApiKeysParams
	isSet bool
}

func (v NullableCountApiKeysParams) Get() *CountApiKeysParams {
	return v.value
}

func (v *NullableCountApiKeysParams) Set(val *CountApiKeysParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCountApiKeysParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCountApiKeysParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountApiKeysParams(val *CountApiKeysParams) *NullableCountApiKeysParams {
	return &NullableCountApiKeysParams{value: val, isSet: true}
}

func (v NullableCountApiKeysParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountApiKeysParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
