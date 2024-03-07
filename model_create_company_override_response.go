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

// checks if the CreateCompanyOverrideResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCompanyOverrideResponse{}

// CreateCompanyOverrideResponse struct for CreateCompanyOverrideResponse
type CreateCompanyOverrideResponse struct {
	Data CompanyOverrideResponseData `json:"data"`
	// Input parameters
	Params map[string]interface{} `json:"params"`
}

// NewCreateCompanyOverrideResponse instantiates a new CreateCompanyOverrideResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCompanyOverrideResponse(data CompanyOverrideResponseData, params map[string]interface{}) *CreateCompanyOverrideResponse {
	this := CreateCompanyOverrideResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewCreateCompanyOverrideResponseWithDefaults instantiates a new CreateCompanyOverrideResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCompanyOverrideResponseWithDefaults() *CreateCompanyOverrideResponse {
	this := CreateCompanyOverrideResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CreateCompanyOverrideResponse) GetData() CompanyOverrideResponseData {
	if o == nil {
		var ret CompanyOverrideResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateCompanyOverrideResponse) GetDataOk() (*CompanyOverrideResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateCompanyOverrideResponse) SetData(v CompanyOverrideResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *CreateCompanyOverrideResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *CreateCompanyOverrideResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *CreateCompanyOverrideResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o CreateCompanyOverrideResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCompanyOverrideResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

type NullableCreateCompanyOverrideResponse struct {
	value *CreateCompanyOverrideResponse
	isSet bool
}

func (v NullableCreateCompanyOverrideResponse) Get() *CreateCompanyOverrideResponse {
	return v.value
}

func (v *NullableCreateCompanyOverrideResponse) Set(val *CreateCompanyOverrideResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCompanyOverrideResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCompanyOverrideResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCompanyOverrideResponse(val *CreateCompanyOverrideResponse) *NullableCreateCompanyOverrideResponse {
	return &NullableCreateCompanyOverrideResponse{value: val, isSet: true}
}

func (v NullableCreateCompanyOverrideResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCompanyOverrideResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


