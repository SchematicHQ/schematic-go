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

// checks if the ListPlanEntitlementsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPlanEntitlementsParams{}

// ListPlanEntitlementsParams Input parameters
type ListPlanEntitlementsParams struct {
	FeatureId *string `json:"feature_id,omitempty"`
	Limit     *int32  `json:"limit,omitempty"`
	Offset    *int32  `json:"offset,omitempty"`
	PlanId    *string `json:"plan_id,omitempty"`
}

// NewListPlanEntitlementsParams instantiates a new ListPlanEntitlementsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPlanEntitlementsParams() *ListPlanEntitlementsParams {
	this := ListPlanEntitlementsParams{}
	return &this
}

// NewListPlanEntitlementsParamsWithDefaults instantiates a new ListPlanEntitlementsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPlanEntitlementsParamsWithDefaults() *ListPlanEntitlementsParams {
	this := ListPlanEntitlementsParams{}
	return &this
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *ListPlanEntitlementsParams) GetFeatureId() string {
	if o == nil || IsNil(o.FeatureId) {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlanEntitlementsParams) GetFeatureIdOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureId) {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *ListPlanEntitlementsParams) HasFeatureId() bool {
	if o != nil && !IsNil(o.FeatureId) {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *ListPlanEntitlementsParams) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ListPlanEntitlementsParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlanEntitlementsParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListPlanEntitlementsParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ListPlanEntitlementsParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ListPlanEntitlementsParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlanEntitlementsParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ListPlanEntitlementsParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *ListPlanEntitlementsParams) SetOffset(v int32) {
	o.Offset = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *ListPlanEntitlementsParams) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlanEntitlementsParams) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *ListPlanEntitlementsParams) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *ListPlanEntitlementsParams) SetPlanId(v string) {
	o.PlanId = &v
}

func (o ListPlanEntitlementsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPlanEntitlementsParams) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.PlanId) {
		toSerialize["plan_id"] = o.PlanId
	}
	return toSerialize, nil
}

type NullableListPlanEntitlementsParams struct {
	value *ListPlanEntitlementsParams
	isSet bool
}

func (v NullableListPlanEntitlementsParams) Get() *ListPlanEntitlementsParams {
	return v.value
}

func (v *NullableListPlanEntitlementsParams) Set(val *ListPlanEntitlementsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableListPlanEntitlementsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableListPlanEntitlementsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPlanEntitlementsParams(val *ListPlanEntitlementsParams) *NullableListPlanEntitlementsParams {
	return &NullableListPlanEntitlementsParams{value: val, isSet: true}
}

func (v NullableListPlanEntitlementsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPlanEntitlementsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
