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

// checks if the CountCompanyOverridesParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountCompanyOverridesParams{}

// CountCompanyOverridesParams Input parameters
type CountCompanyOverridesParams struct {
	CompanyId  *string  `json:"company_id,omitempty"`
	CompanyIds []string `json:"company_ids,omitempty"`
	FeatureId  *string  `json:"feature_id,omitempty"`
	FeatureIds []string `json:"feature_ids,omitempty"`
	Ids        []string `json:"ids,omitempty"`
	// Page limit (default 100)
	Limit *int32 `json:"limit,omitempty"`
	// Page offset (default 0)
	Offset *int32 `json:"offset,omitempty"`
}

// NewCountCompanyOverridesParams instantiates a new CountCompanyOverridesParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountCompanyOverridesParams() *CountCompanyOverridesParams {
	this := CountCompanyOverridesParams{}
	return &this
}

// NewCountCompanyOverridesParamsWithDefaults instantiates a new CountCompanyOverridesParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountCompanyOverridesParamsWithDefaults() *CountCompanyOverridesParams {
	this := CountCompanyOverridesParams{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *CountCompanyOverridesParams) GetCompanyId() string {
	if o == nil || IsNil(o.CompanyId) {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountCompanyOverridesParams) GetCompanyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyId) {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *CountCompanyOverridesParams) HasCompanyId() bool {
	if o != nil && !IsNil(o.CompanyId) {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *CountCompanyOverridesParams) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetCompanyIds returns the CompanyIds field value if set, zero value otherwise.
func (o *CountCompanyOverridesParams) GetCompanyIds() []string {
	if o == nil || IsNil(o.CompanyIds) {
		var ret []string
		return ret
	}
	return o.CompanyIds
}

// GetCompanyIdsOk returns a tuple with the CompanyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountCompanyOverridesParams) GetCompanyIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CompanyIds) {
		return nil, false
	}
	return o.CompanyIds, true
}

// HasCompanyIds returns a boolean if a field has been set.
func (o *CountCompanyOverridesParams) HasCompanyIds() bool {
	if o != nil && !IsNil(o.CompanyIds) {
		return true
	}

	return false
}

// SetCompanyIds gets a reference to the given []string and assigns it to the CompanyIds field.
func (o *CountCompanyOverridesParams) SetCompanyIds(v []string) {
	o.CompanyIds = v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *CountCompanyOverridesParams) GetFeatureId() string {
	if o == nil || IsNil(o.FeatureId) {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountCompanyOverridesParams) GetFeatureIdOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureId) {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *CountCompanyOverridesParams) HasFeatureId() bool {
	if o != nil && !IsNil(o.FeatureId) {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *CountCompanyOverridesParams) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetFeatureIds returns the FeatureIds field value if set, zero value otherwise.
func (o *CountCompanyOverridesParams) GetFeatureIds() []string {
	if o == nil || IsNil(o.FeatureIds) {
		var ret []string
		return ret
	}
	return o.FeatureIds
}

// GetFeatureIdsOk returns a tuple with the FeatureIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountCompanyOverridesParams) GetFeatureIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.FeatureIds) {
		return nil, false
	}
	return o.FeatureIds, true
}

// HasFeatureIds returns a boolean if a field has been set.
func (o *CountCompanyOverridesParams) HasFeatureIds() bool {
	if o != nil && !IsNil(o.FeatureIds) {
		return true
	}

	return false
}

// SetFeatureIds gets a reference to the given []string and assigns it to the FeatureIds field.
func (o *CountCompanyOverridesParams) SetFeatureIds(v []string) {
	o.FeatureIds = v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *CountCompanyOverridesParams) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountCompanyOverridesParams) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *CountCompanyOverridesParams) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *CountCompanyOverridesParams) SetIds(v []string) {
	o.Ids = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *CountCompanyOverridesParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountCompanyOverridesParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *CountCompanyOverridesParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *CountCompanyOverridesParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *CountCompanyOverridesParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountCompanyOverridesParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *CountCompanyOverridesParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *CountCompanyOverridesParams) SetOffset(v int32) {
	o.Offset = &v
}

func (o CountCompanyOverridesParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountCompanyOverridesParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanyId) {
		toSerialize["company_id"] = o.CompanyId
	}
	if !IsNil(o.CompanyIds) {
		toSerialize["company_ids"] = o.CompanyIds
	}
	if !IsNil(o.FeatureId) {
		toSerialize["feature_id"] = o.FeatureId
	}
	if !IsNil(o.FeatureIds) {
		toSerialize["feature_ids"] = o.FeatureIds
	}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	return toSerialize, nil
}

type NullableCountCompanyOverridesParams struct {
	value *CountCompanyOverridesParams
	isSet bool
}

func (v NullableCountCompanyOverridesParams) Get() *CountCompanyOverridesParams {
	return v.value
}

func (v *NullableCountCompanyOverridesParams) Set(val *CountCompanyOverridesParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCountCompanyOverridesParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCountCompanyOverridesParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountCompanyOverridesParams(val *CountCompanyOverridesParams) *NullableCountCompanyOverridesParams {
	return &NullableCountCompanyOverridesParams{value: val, isSet: true}
}

func (v NullableCountCompanyOverridesParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountCompanyOverridesParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
