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

// checks if the CreateApiKeyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateApiKeyResponse{}

// CreateApiKeyResponse struct for CreateApiKeyResponse
type CreateApiKeyResponse struct {
	Data ApiKeyCreateResponseData `json:"data"`
	// Input parameters
	Params               map[string]interface{} `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _CreateApiKeyResponse CreateApiKeyResponse

// NewCreateApiKeyResponse instantiates a new CreateApiKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApiKeyResponse(data ApiKeyCreateResponseData, params map[string]interface{}) *CreateApiKeyResponse {
	this := CreateApiKeyResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewCreateApiKeyResponseWithDefaults instantiates a new CreateApiKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApiKeyResponseWithDefaults() *CreateApiKeyResponse {
	this := CreateApiKeyResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CreateApiKeyResponse) GetData() ApiKeyCreateResponseData {
	if o == nil {
		var ret ApiKeyCreateResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateApiKeyResponse) GetDataOk() (*ApiKeyCreateResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateApiKeyResponse) SetData(v ApiKeyCreateResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *CreateApiKeyResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *CreateApiKeyResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *CreateApiKeyResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o CreateApiKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateApiKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateApiKeyResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCreateApiKeyResponse := _CreateApiKeyResponse{}

	err = json.Unmarshal(data, &varCreateApiKeyResponse)

	if err != nil {
		return err
	}

	*o = CreateApiKeyResponse(varCreateApiKeyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateApiKeyResponse struct {
	value *CreateApiKeyResponse
	isSet bool
}

func (v NullableCreateApiKeyResponse) Get() *CreateApiKeyResponse {
	return v.value
}

func (v *NullableCreateApiKeyResponse) Set(val *CreateApiKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApiKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApiKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApiKeyResponse(val *CreateApiKeyResponse) *NullableCreateApiKeyResponse {
	return &NullableCreateApiKeyResponse{value: val, isSet: true}
}

func (v NullableCreateApiKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApiKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
