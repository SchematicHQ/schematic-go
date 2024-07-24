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

// checks if the ListEntityTraitDefinitionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListEntityTraitDefinitionsResponse{}

// ListEntityTraitDefinitionsResponse struct for ListEntityTraitDefinitionsResponse
type ListEntityTraitDefinitionsResponse struct {
	// The returned resources
	Data                 []EntityTraitDefinitionResponseData `json:"data"`
	Params               ListEntityTraitDefinitionsParams    `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _ListEntityTraitDefinitionsResponse ListEntityTraitDefinitionsResponse

// NewListEntityTraitDefinitionsResponse instantiates a new ListEntityTraitDefinitionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEntityTraitDefinitionsResponse(data []EntityTraitDefinitionResponseData, params ListEntityTraitDefinitionsParams) *ListEntityTraitDefinitionsResponse {
	this := ListEntityTraitDefinitionsResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewListEntityTraitDefinitionsResponseWithDefaults instantiates a new ListEntityTraitDefinitionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEntityTraitDefinitionsResponseWithDefaults() *ListEntityTraitDefinitionsResponse {
	this := ListEntityTraitDefinitionsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListEntityTraitDefinitionsResponse) GetData() []EntityTraitDefinitionResponseData {
	if o == nil {
		var ret []EntityTraitDefinitionResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListEntityTraitDefinitionsResponse) GetDataOk() ([]EntityTraitDefinitionResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListEntityTraitDefinitionsResponse) SetData(v []EntityTraitDefinitionResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *ListEntityTraitDefinitionsResponse) GetParams() ListEntityTraitDefinitionsParams {
	if o == nil {
		var ret ListEntityTraitDefinitionsParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *ListEntityTraitDefinitionsResponse) GetParamsOk() (*ListEntityTraitDefinitionsParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *ListEntityTraitDefinitionsResponse) SetParams(v ListEntityTraitDefinitionsParams) {
	o.Params = v
}

func (o ListEntityTraitDefinitionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListEntityTraitDefinitionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListEntityTraitDefinitionsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varListEntityTraitDefinitionsResponse := _ListEntityTraitDefinitionsResponse{}

	err = json.Unmarshal(data, &varListEntityTraitDefinitionsResponse)

	if err != nil {
		return err
	}

	*o = ListEntityTraitDefinitionsResponse(varListEntityTraitDefinitionsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListEntityTraitDefinitionsResponse struct {
	value *ListEntityTraitDefinitionsResponse
	isSet bool
}

func (v NullableListEntityTraitDefinitionsResponse) Get() *ListEntityTraitDefinitionsResponse {
	return v.value
}

func (v *NullableListEntityTraitDefinitionsResponse) Set(val *ListEntityTraitDefinitionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListEntityTraitDefinitionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListEntityTraitDefinitionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEntityTraitDefinitionsResponse(val *ListEntityTraitDefinitionsResponse) *NullableListEntityTraitDefinitionsResponse {
	return &NullableListEntityTraitDefinitionsResponse{value: val, isSet: true}
}

func (v NullableListEntityTraitDefinitionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEntityTraitDefinitionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
