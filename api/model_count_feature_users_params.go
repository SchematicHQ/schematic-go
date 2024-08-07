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

// checks if the CountFeatureUsersParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountFeatureUsersParams{}

// CountFeatureUsersParams Input parameters
type CountFeatureUsersParams struct {
	FeatureId *string `json:"feature_id,omitempty"`
	// Page limit (default 100)
	Limit *int32 `json:"limit,omitempty"`
	// Page offset (default 0)
	Offset               *int32  `json:"offset,omitempty"`
	Q                    *string `json:"q,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CountFeatureUsersParams CountFeatureUsersParams

// NewCountFeatureUsersParams instantiates a new CountFeatureUsersParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountFeatureUsersParams() *CountFeatureUsersParams {
	this := CountFeatureUsersParams{}
	return &this
}

// NewCountFeatureUsersParamsWithDefaults instantiates a new CountFeatureUsersParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountFeatureUsersParamsWithDefaults() *CountFeatureUsersParams {
	this := CountFeatureUsersParams{}
	return &this
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *CountFeatureUsersParams) GetFeatureId() string {
	if o == nil || IsNil(o.FeatureId) {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountFeatureUsersParams) GetFeatureIdOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureId) {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *CountFeatureUsersParams) HasFeatureId() bool {
	if o != nil && !IsNil(o.FeatureId) {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *CountFeatureUsersParams) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *CountFeatureUsersParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountFeatureUsersParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *CountFeatureUsersParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *CountFeatureUsersParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *CountFeatureUsersParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountFeatureUsersParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *CountFeatureUsersParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *CountFeatureUsersParams) SetOffset(v int32) {
	o.Offset = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *CountFeatureUsersParams) GetQ() string {
	if o == nil || IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountFeatureUsersParams) GetQOk() (*string, bool) {
	if o == nil || IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *CountFeatureUsersParams) HasQ() bool {
	if o != nil && !IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *CountFeatureUsersParams) SetQ(v string) {
	o.Q = &v
}

func (o CountFeatureUsersParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountFeatureUsersParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeatureId) {
		toSerialize["feature_id"] = o.FeatureId
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Q) {
		toSerialize["q"] = o.Q
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CountFeatureUsersParams) UnmarshalJSON(data []byte) (err error) {
	varCountFeatureUsersParams := _CountFeatureUsersParams{}

	err = json.Unmarshal(data, &varCountFeatureUsersParams)

	if err != nil {
		return err
	}

	*o = CountFeatureUsersParams(varCountFeatureUsersParams)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "feature_id")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "q")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCountFeatureUsersParams struct {
	value *CountFeatureUsersParams
	isSet bool
}

func (v NullableCountFeatureUsersParams) Get() *CountFeatureUsersParams {
	return v.value
}

func (v *NullableCountFeatureUsersParams) Set(val *CountFeatureUsersParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCountFeatureUsersParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCountFeatureUsersParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountFeatureUsersParams(val *CountFeatureUsersParams) *NullableCountFeatureUsersParams {
	return &NullableCountFeatureUsersParams{value: val, isSet: true}
}

func (v NullableCountFeatureUsersParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountFeatureUsersParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
