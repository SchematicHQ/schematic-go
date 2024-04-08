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

// checks if the GetOrCreateEntityTraitDefinitionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrCreateEntityTraitDefinitionResponse{}

// GetOrCreateEntityTraitDefinitionResponse struct for GetOrCreateEntityTraitDefinitionResponse
type GetOrCreateEntityTraitDefinitionResponse struct {
	Data EntityTraitDefinitionResponseData `json:"data"`
	// Input parameters
	Params map[string]interface{} `json:"params"`
}

type _GetOrCreateEntityTraitDefinitionResponse GetOrCreateEntityTraitDefinitionResponse

// NewGetOrCreateEntityTraitDefinitionResponse instantiates a new GetOrCreateEntityTraitDefinitionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrCreateEntityTraitDefinitionResponse(data EntityTraitDefinitionResponseData, params map[string]interface{}) *GetOrCreateEntityTraitDefinitionResponse {
	this := GetOrCreateEntityTraitDefinitionResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewGetOrCreateEntityTraitDefinitionResponseWithDefaults instantiates a new GetOrCreateEntityTraitDefinitionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrCreateEntityTraitDefinitionResponseWithDefaults() *GetOrCreateEntityTraitDefinitionResponse {
	this := GetOrCreateEntityTraitDefinitionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetOrCreateEntityTraitDefinitionResponse) GetData() EntityTraitDefinitionResponseData {
	if o == nil {
		var ret EntityTraitDefinitionResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetOrCreateEntityTraitDefinitionResponse) GetDataOk() (*EntityTraitDefinitionResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetOrCreateEntityTraitDefinitionResponse) SetData(v EntityTraitDefinitionResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *GetOrCreateEntityTraitDefinitionResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *GetOrCreateEntityTraitDefinitionResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *GetOrCreateEntityTraitDefinitionResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o GetOrCreateEntityTraitDefinitionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrCreateEntityTraitDefinitionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *GetOrCreateEntityTraitDefinitionResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetOrCreateEntityTraitDefinitionResponse := _GetOrCreateEntityTraitDefinitionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetOrCreateEntityTraitDefinitionResponse)

	if err != nil {
		return err
	}

	*o = GetOrCreateEntityTraitDefinitionResponse(varGetOrCreateEntityTraitDefinitionResponse)

	return err
}

type NullableGetOrCreateEntityTraitDefinitionResponse struct {
	value *GetOrCreateEntityTraitDefinitionResponse
	isSet bool
}

func (v NullableGetOrCreateEntityTraitDefinitionResponse) Get() *GetOrCreateEntityTraitDefinitionResponse {
	return v.value
}

func (v *NullableGetOrCreateEntityTraitDefinitionResponse) Set(val *GetOrCreateEntityTraitDefinitionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrCreateEntityTraitDefinitionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrCreateEntityTraitDefinitionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrCreateEntityTraitDefinitionResponse(val *GetOrCreateEntityTraitDefinitionResponse) *NullableGetOrCreateEntityTraitDefinitionResponse {
	return &NullableGetOrCreateEntityTraitDefinitionResponse{value: val, isSet: true}
}

func (v NullableGetOrCreateEntityTraitDefinitionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrCreateEntityTraitDefinitionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

