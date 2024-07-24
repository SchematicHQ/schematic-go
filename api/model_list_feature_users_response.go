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

// checks if the ListFeatureUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFeatureUsersResponse{}

// ListFeatureUsersResponse struct for ListFeatureUsersResponse
type ListFeatureUsersResponse struct {
	// The returned resources
	Data                 []FeatureCompanyUserResponseData `json:"data"`
	Params               ListFeatureUsersParams           `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _ListFeatureUsersResponse ListFeatureUsersResponse

// NewListFeatureUsersResponse instantiates a new ListFeatureUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFeatureUsersResponse(data []FeatureCompanyUserResponseData, params ListFeatureUsersParams) *ListFeatureUsersResponse {
	this := ListFeatureUsersResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewListFeatureUsersResponseWithDefaults instantiates a new ListFeatureUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFeatureUsersResponseWithDefaults() *ListFeatureUsersResponse {
	this := ListFeatureUsersResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListFeatureUsersResponse) GetData() []FeatureCompanyUserResponseData {
	if o == nil {
		var ret []FeatureCompanyUserResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListFeatureUsersResponse) GetDataOk() ([]FeatureCompanyUserResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListFeatureUsersResponse) SetData(v []FeatureCompanyUserResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *ListFeatureUsersResponse) GetParams() ListFeatureUsersParams {
	if o == nil {
		var ret ListFeatureUsersParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *ListFeatureUsersResponse) GetParamsOk() (*ListFeatureUsersParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *ListFeatureUsersResponse) SetParams(v ListFeatureUsersParams) {
	o.Params = v
}

func (o ListFeatureUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFeatureUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListFeatureUsersResponse) UnmarshalJSON(data []byte) (err error) {
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

	varListFeatureUsersResponse := _ListFeatureUsersResponse{}

	err = json.Unmarshal(data, &varListFeatureUsersResponse)

	if err != nil {
		return err
	}

	*o = ListFeatureUsersResponse(varListFeatureUsersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListFeatureUsersResponse struct {
	value *ListFeatureUsersResponse
	isSet bool
}

func (v NullableListFeatureUsersResponse) Get() *ListFeatureUsersResponse {
	return v.value
}

func (v *NullableListFeatureUsersResponse) Set(val *ListFeatureUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListFeatureUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListFeatureUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFeatureUsersResponse(val *ListFeatureUsersResponse) *NullableListFeatureUsersResponse {
	return &NullableListFeatureUsersResponse{value: val, isSet: true}
}

func (v NullableListFeatureUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFeatureUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
