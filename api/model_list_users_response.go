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

// checks if the ListUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListUsersResponse{}

// ListUsersResponse struct for ListUsersResponse
type ListUsersResponse struct {
	// The returned resources
	Data                 []UserDetailResponseData `json:"data"`
	Params               ListUsersParams          `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _ListUsersResponse ListUsersResponse

// NewListUsersResponse instantiates a new ListUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUsersResponse(data []UserDetailResponseData, params ListUsersParams) *ListUsersResponse {
	this := ListUsersResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewListUsersResponseWithDefaults instantiates a new ListUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUsersResponseWithDefaults() *ListUsersResponse {
	this := ListUsersResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListUsersResponse) GetData() []UserDetailResponseData {
	if o == nil {
		var ret []UserDetailResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListUsersResponse) GetDataOk() ([]UserDetailResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListUsersResponse) SetData(v []UserDetailResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *ListUsersResponse) GetParams() ListUsersParams {
	if o == nil {
		var ret ListUsersParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *ListUsersResponse) GetParamsOk() (*ListUsersParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *ListUsersResponse) SetParams(v ListUsersParams) {
	o.Params = v
}

func (o ListUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListUsersResponse) UnmarshalJSON(data []byte) (err error) {
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

	varListUsersResponse := _ListUsersResponse{}

	err = json.Unmarshal(data, &varListUsersResponse)

	if err != nil {
		return err
	}

	*o = ListUsersResponse(varListUsersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListUsersResponse struct {
	value *ListUsersResponse
	isSet bool
}

func (v NullableListUsersResponse) Get() *ListUsersResponse {
	return v.value
}

func (v *NullableListUsersResponse) Set(val *ListUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUsersResponse(val *ListUsersResponse) *NullableListUsersResponse {
	return &NullableListUsersResponse{value: val, isSet: true}
}

func (v NullableListUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
