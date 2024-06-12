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

// checks if the GetActiveDealsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetActiveDealsParams{}

// GetActiveDealsParams Input parameters
type GetActiveDealsParams struct {
	CompanyId *string `json:"company_id,omitempty"`
	DealStage *string `json:"deal_stage,omitempty"`
	Limit     *int32  `json:"limit,omitempty"`
	Offset    *int32  `json:"offset,omitempty"`
}

// NewGetActiveDealsParams instantiates a new GetActiveDealsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetActiveDealsParams() *GetActiveDealsParams {
	this := GetActiveDealsParams{}
	return &this
}

// NewGetActiveDealsParamsWithDefaults instantiates a new GetActiveDealsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetActiveDealsParamsWithDefaults() *GetActiveDealsParams {
	this := GetActiveDealsParams{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *GetActiveDealsParams) GetCompanyId() string {
	if o == nil || IsNil(o.CompanyId) {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActiveDealsParams) GetCompanyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyId) {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *GetActiveDealsParams) HasCompanyId() bool {
	if o != nil && !IsNil(o.CompanyId) {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *GetActiveDealsParams) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetDealStage returns the DealStage field value if set, zero value otherwise.
func (o *GetActiveDealsParams) GetDealStage() string {
	if o == nil || IsNil(o.DealStage) {
		var ret string
		return ret
	}
	return *o.DealStage
}

// GetDealStageOk returns a tuple with the DealStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActiveDealsParams) GetDealStageOk() (*string, bool) {
	if o == nil || IsNil(o.DealStage) {
		return nil, false
	}
	return o.DealStage, true
}

// HasDealStage returns a boolean if a field has been set.
func (o *GetActiveDealsParams) HasDealStage() bool {
	if o != nil && !IsNil(o.DealStage) {
		return true
	}

	return false
}

// SetDealStage gets a reference to the given string and assigns it to the DealStage field.
func (o *GetActiveDealsParams) SetDealStage(v string) {
	o.DealStage = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetActiveDealsParams) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActiveDealsParams) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetActiveDealsParams) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *GetActiveDealsParams) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *GetActiveDealsParams) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActiveDealsParams) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *GetActiveDealsParams) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *GetActiveDealsParams) SetOffset(v int32) {
	o.Offset = &v
}

func (o GetActiveDealsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetActiveDealsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanyId) {
		toSerialize["company_id"] = o.CompanyId
	}
	if !IsNil(o.DealStage) {
		toSerialize["deal_stage"] = o.DealStage
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	return toSerialize, nil
}

type NullableGetActiveDealsParams struct {
	value *GetActiveDealsParams
	isSet bool
}

func (v NullableGetActiveDealsParams) Get() *GetActiveDealsParams {
	return v.value
}

func (v *NullableGetActiveDealsParams) Set(val *GetActiveDealsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetActiveDealsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetActiveDealsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetActiveDealsParams(val *GetActiveDealsParams) *NullableGetActiveDealsParams {
	return &NullableGetActiveDealsParams{value: val, isSet: true}
}

func (v NullableGetActiveDealsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetActiveDealsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
