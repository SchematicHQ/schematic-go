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

// checks if the ListCompanyMembershipsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCompanyMembershipsResponse{}

// ListCompanyMembershipsResponse struct for ListCompanyMembershipsResponse
type ListCompanyMembershipsResponse struct {
	// The returned resources
	Data []CompanyMembershipDetailResponseData `json:"data"`
	Params ListCompanyMembershipsParams `json:"params"`
}

// NewListCompanyMembershipsResponse instantiates a new ListCompanyMembershipsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCompanyMembershipsResponse(data []CompanyMembershipDetailResponseData, params ListCompanyMembershipsParams) *ListCompanyMembershipsResponse {
	this := ListCompanyMembershipsResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewListCompanyMembershipsResponseWithDefaults instantiates a new ListCompanyMembershipsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCompanyMembershipsResponseWithDefaults() *ListCompanyMembershipsResponse {
	this := ListCompanyMembershipsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListCompanyMembershipsResponse) GetData() []CompanyMembershipDetailResponseData {
	if o == nil {
		var ret []CompanyMembershipDetailResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListCompanyMembershipsResponse) GetDataOk() ([]CompanyMembershipDetailResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListCompanyMembershipsResponse) SetData(v []CompanyMembershipDetailResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *ListCompanyMembershipsResponse) GetParams() ListCompanyMembershipsParams {
	if o == nil {
		var ret ListCompanyMembershipsParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *ListCompanyMembershipsResponse) GetParamsOk() (*ListCompanyMembershipsParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *ListCompanyMembershipsResponse) SetParams(v ListCompanyMembershipsParams) {
	o.Params = v
}

func (o ListCompanyMembershipsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCompanyMembershipsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

type NullableListCompanyMembershipsResponse struct {
	value *ListCompanyMembershipsResponse
	isSet bool
}

func (v NullableListCompanyMembershipsResponse) Get() *ListCompanyMembershipsResponse {
	return v.value
}

func (v *NullableListCompanyMembershipsResponse) Set(val *ListCompanyMembershipsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCompanyMembershipsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCompanyMembershipsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCompanyMembershipsResponse(val *ListCompanyMembershipsResponse) *NullableListCompanyMembershipsResponse {
	return &NullableListCompanyMembershipsResponse{value: val, isSet: true}
}

func (v NullableListCompanyMembershipsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCompanyMembershipsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


