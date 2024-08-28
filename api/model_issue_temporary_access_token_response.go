/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// checks if the IssueTemporaryAccessTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssueTemporaryAccessTokenResponse{}

// IssueTemporaryAccessTokenResponse struct for IssueTemporaryAccessTokenResponse
type IssueTemporaryAccessTokenResponse struct {
	Data IssueTemporaryAccessTokenResponseData `json:"data"`
	// Input parameters
	Params               map[string]interface{} `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _IssueTemporaryAccessTokenResponse IssueTemporaryAccessTokenResponse

// NewIssueTemporaryAccessTokenResponse instantiates a new IssueTemporaryAccessTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueTemporaryAccessTokenResponse(data IssueTemporaryAccessTokenResponseData, params map[string]interface{}) *IssueTemporaryAccessTokenResponse {
	this := IssueTemporaryAccessTokenResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewIssueTemporaryAccessTokenResponseWithDefaults instantiates a new IssueTemporaryAccessTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueTemporaryAccessTokenResponseWithDefaults() *IssueTemporaryAccessTokenResponse {
	this := IssueTemporaryAccessTokenResponse{}
	return &this
}

// GetData returns the Data field value
func (o *IssueTemporaryAccessTokenResponse) GetData() IssueTemporaryAccessTokenResponseData {
	if o == nil {
		var ret IssueTemporaryAccessTokenResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IssueTemporaryAccessTokenResponse) GetDataOk() (*IssueTemporaryAccessTokenResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *IssueTemporaryAccessTokenResponse) SetData(v IssueTemporaryAccessTokenResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *IssueTemporaryAccessTokenResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *IssueTemporaryAccessTokenResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *IssueTemporaryAccessTokenResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o IssueTemporaryAccessTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueTemporaryAccessTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IssueTemporaryAccessTokenResponse) UnmarshalJSON(data []byte) (err error) {
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
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varIssueTemporaryAccessTokenResponse := _IssueTemporaryAccessTokenResponse{}

	err = json.Unmarshal(data, &varIssueTemporaryAccessTokenResponse)

	if err != nil {
		return err
	}

	*o = IssueTemporaryAccessTokenResponse(varIssueTemporaryAccessTokenResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIssueTemporaryAccessTokenResponse struct {
	value *IssueTemporaryAccessTokenResponse
	isSet bool
}

func (v NullableIssueTemporaryAccessTokenResponse) Get() *IssueTemporaryAccessTokenResponse {
	return v.value
}

func (v *NullableIssueTemporaryAccessTokenResponse) Set(val *IssueTemporaryAccessTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueTemporaryAccessTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueTemporaryAccessTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueTemporaryAccessTokenResponse(val *IssueTemporaryAccessTokenResponse) *NullableIssueTemporaryAccessTokenResponse {
	return &NullableIssueTemporaryAccessTokenResponse{value: val, isSet: true}
}

func (v NullableIssueTemporaryAccessTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueTemporaryAccessTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}