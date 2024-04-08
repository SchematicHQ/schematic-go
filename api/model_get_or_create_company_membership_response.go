/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetOrCreateCompanyMembershipResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrCreateCompanyMembershipResponse{}

// GetOrCreateCompanyMembershipResponse struct for GetOrCreateCompanyMembershipResponse
type GetOrCreateCompanyMembershipResponse struct {
	Data CompanyMembershipDetailResponseData `json:"data"`
	// Input parameters
	Params map[string]interface{} `json:"params"`
}

type _GetOrCreateCompanyMembershipResponse GetOrCreateCompanyMembershipResponse

// NewGetOrCreateCompanyMembershipResponse instantiates a new GetOrCreateCompanyMembershipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrCreateCompanyMembershipResponse(data CompanyMembershipDetailResponseData, params map[string]interface{}) *GetOrCreateCompanyMembershipResponse {
	this := GetOrCreateCompanyMembershipResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewGetOrCreateCompanyMembershipResponseWithDefaults instantiates a new GetOrCreateCompanyMembershipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrCreateCompanyMembershipResponseWithDefaults() *GetOrCreateCompanyMembershipResponse {
	this := GetOrCreateCompanyMembershipResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetOrCreateCompanyMembershipResponse) GetData() CompanyMembershipDetailResponseData {
	if o == nil {
		var ret CompanyMembershipDetailResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetOrCreateCompanyMembershipResponse) GetDataOk() (*CompanyMembershipDetailResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetOrCreateCompanyMembershipResponse) SetData(v CompanyMembershipDetailResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *GetOrCreateCompanyMembershipResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *GetOrCreateCompanyMembershipResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *GetOrCreateCompanyMembershipResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o GetOrCreateCompanyMembershipResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrCreateCompanyMembershipResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *GetOrCreateCompanyMembershipResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"params",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetOrCreateCompanyMembershipResponse := _GetOrCreateCompanyMembershipResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetOrCreateCompanyMembershipResponse)

	if err != nil {
		return err
	}

	*o = GetOrCreateCompanyMembershipResponse(varGetOrCreateCompanyMembershipResponse)

	return err
}

type NullableGetOrCreateCompanyMembershipResponse struct {
	value *GetOrCreateCompanyMembershipResponse
	isSet bool
}

func (v NullableGetOrCreateCompanyMembershipResponse) Get() *GetOrCreateCompanyMembershipResponse {
	return v.value
}

func (v *NullableGetOrCreateCompanyMembershipResponse) Set(val *GetOrCreateCompanyMembershipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrCreateCompanyMembershipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrCreateCompanyMembershipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrCreateCompanyMembershipResponse(val *GetOrCreateCompanyMembershipResponse) *NullableGetOrCreateCompanyMembershipResponse {
	return &NullableGetOrCreateCompanyMembershipResponse{value: val, isSet: true}
}

func (v NullableGetOrCreateCompanyMembershipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrCreateCompanyMembershipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


