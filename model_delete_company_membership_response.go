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

// checks if the DeleteCompanyMembershipResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteCompanyMembershipResponse{}

// DeleteCompanyMembershipResponse struct for DeleteCompanyMembershipResponse
type DeleteCompanyMembershipResponse struct {
	Data DeleteResponse `json:"data"`
	// Input parameters
	Params map[string]interface{} `json:"params"`
}

// NewDeleteCompanyMembershipResponse instantiates a new DeleteCompanyMembershipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteCompanyMembershipResponse(data DeleteResponse, params map[string]interface{}) *DeleteCompanyMembershipResponse {
	this := DeleteCompanyMembershipResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewDeleteCompanyMembershipResponseWithDefaults instantiates a new DeleteCompanyMembershipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteCompanyMembershipResponseWithDefaults() *DeleteCompanyMembershipResponse {
	this := DeleteCompanyMembershipResponse{}
	return &this
}

// GetData returns the Data field value
func (o *DeleteCompanyMembershipResponse) GetData() DeleteResponse {
	if o == nil {
		var ret DeleteResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeleteCompanyMembershipResponse) GetDataOk() (*DeleteResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeleteCompanyMembershipResponse) SetData(v DeleteResponse) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *DeleteCompanyMembershipResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *DeleteCompanyMembershipResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *DeleteCompanyMembershipResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o DeleteCompanyMembershipResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteCompanyMembershipResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

type NullableDeleteCompanyMembershipResponse struct {
	value *DeleteCompanyMembershipResponse
	isSet bool
}

func (v NullableDeleteCompanyMembershipResponse) Get() *DeleteCompanyMembershipResponse {
	return v.value
}

func (v *NullableDeleteCompanyMembershipResponse) Set(val *DeleteCompanyMembershipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteCompanyMembershipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteCompanyMembershipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteCompanyMembershipResponse(val *DeleteCompanyMembershipResponse) *NullableDeleteCompanyMembershipResponse {
	return &NullableDeleteCompanyMembershipResponse{value: val, isSet: true}
}

func (v NullableDeleteCompanyMembershipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteCompanyMembershipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


