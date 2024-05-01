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

// checks if the CountFeatureCompaniesParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountFeatureCompaniesParams{}

// CountFeatureCompaniesParams Input parameters
type CountFeatureCompaniesParams struct {
	FeatureId *string `json:"feature_id,omitempty"`
	Limit     *int32  `json:"limit,omitempty"`
	Offset    *int32  `json:"offset,omitempty"`
	Q         *string `json:"q,omitempty"`
}

// NewCountFeatureCompaniesParams instantiates a new CountFeatureCompaniesParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountFeatureCompaniesParams() *CountFeatureCompaniesParams {
	this := CountFeatureCompaniesParams{}
	return &this
}

// NewCountFeatureCompaniesParamsWithDefaults instantiates a new CountFeatureCompaniesParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountFeatureCompaniesParamsWithDefaults() *CountFeatureCompaniesParams {
	this := CountFeatureCompaniesParams{}
	return &this
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *CountFeatureCompaniesParams) GetFeatureId() string {
	if o == nil || IsNil(o.FeatureId) {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountFeatureCompaniesParams) GetFeatureIdOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureId) {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *CountFeatureCompaniesParams) HasFeatureId() bool {
	if o != nil && !IsNil(o.FeatureId) {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *CountFeatureCompaniesParams) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *CountFeatureCompaniesParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountFeatureCompaniesParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *CountFeatureCompaniesParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *CountFeatureCompaniesParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *CountFeatureCompaniesParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountFeatureCompaniesParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *CountFeatureCompaniesParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *CountFeatureCompaniesParams) SetOffset(v int32) {
	o.Offset = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *CountFeatureCompaniesParams) GetQ() string {
	if o == nil || IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountFeatureCompaniesParams) GetQOk() (*string, bool) {
	if o == nil || IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *CountFeatureCompaniesParams) HasQ() bool {
	if o != nil && !IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *CountFeatureCompaniesParams) SetQ(v string) {
	o.Q = &v
}

func (o CountFeatureCompaniesParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountFeatureCompaniesParams) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableCountFeatureCompaniesParams struct {
	value *CountFeatureCompaniesParams
	isSet bool
}

func (v NullableCountFeatureCompaniesParams) Get() *CountFeatureCompaniesParams {
	return v.value
}

func (v *NullableCountFeatureCompaniesParams) Set(val *CountFeatureCompaniesParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCountFeatureCompaniesParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCountFeatureCompaniesParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountFeatureCompaniesParams(val *CountFeatureCompaniesParams) *NullableCountFeatureCompaniesParams {
	return &NullableCountFeatureCompaniesParams{value: val, isSet: true}
}

func (v NullableCountFeatureCompaniesParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountFeatureCompaniesParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
