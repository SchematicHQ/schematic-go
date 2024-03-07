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

// checks if the ListCompanyOverridesParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCompanyOverridesParams{}

// ListCompanyOverridesParams Input parameters
type ListCompanyOverridesParams struct {
	CompanyId *string `json:"company_id,omitempty"`
	FeatureId *string `json:"feature_id,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Offset *int32 `json:"offset,omitempty"`
}

// NewListCompanyOverridesParams instantiates a new ListCompanyOverridesParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCompanyOverridesParams() *ListCompanyOverridesParams {
	this := ListCompanyOverridesParams{}
	return &this
}

// NewListCompanyOverridesParamsWithDefaults instantiates a new ListCompanyOverridesParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCompanyOverridesParamsWithDefaults() *ListCompanyOverridesParams {
	this := ListCompanyOverridesParams{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *ListCompanyOverridesParams) GetCompanyId() string {
	if o == nil || IsNil(o.CompanyId) {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCompanyOverridesParams) GetCompanyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyId) {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *ListCompanyOverridesParams) HasCompanyId() bool {
	if o != nil && !IsNil(o.CompanyId) {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *ListCompanyOverridesParams) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *ListCompanyOverridesParams) GetFeatureId() string {
	if o == nil || IsNil(o.FeatureId) {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCompanyOverridesParams) GetFeatureIdOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureId) {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *ListCompanyOverridesParams) HasFeatureId() bool {
	if o != nil && !IsNil(o.FeatureId) {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *ListCompanyOverridesParams) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ListCompanyOverridesParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCompanyOverridesParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListCompanyOverridesParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ListCompanyOverridesParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ListCompanyOverridesParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCompanyOverridesParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ListCompanyOverridesParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *ListCompanyOverridesParams) SetOffset(v int32) {
	o.Offset = &v
}

func (o ListCompanyOverridesParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCompanyOverridesParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanyId) {
		toSerialize["company_id"] = o.CompanyId
	}
	if !IsNil(o.FeatureId) {
		toSerialize["feature_id"] = o.FeatureId
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	return toSerialize, nil
}

type NullableListCompanyOverridesParams struct {
	value *ListCompanyOverridesParams
	isSet bool
}

func (v NullableListCompanyOverridesParams) Get() *ListCompanyOverridesParams {
	return v.value
}

func (v *NullableListCompanyOverridesParams) Set(val *ListCompanyOverridesParams) {
	v.value = val
	v.isSet = true
}

func (v NullableListCompanyOverridesParams) IsSet() bool {
	return v.isSet
}

func (v *NullableListCompanyOverridesParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCompanyOverridesParams(val *ListCompanyOverridesParams) *NullableListCompanyOverridesParams {
	return &NullableListCompanyOverridesParams{value: val, isSet: true}
}

func (v NullableListCompanyOverridesParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCompanyOverridesParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


