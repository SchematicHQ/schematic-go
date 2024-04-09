/*
Schematic API

Schematic API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ListApiKeysResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListApiKeysResponse{}

// ListApiKeysResponse struct for ListApiKeysResponse
type ListApiKeysResponse struct {
	// The returned resources
	Data   []ApiKeyResponseData `json:"data"`
	Params ListApiKeysParams    `json:"params"`
}

type _ListApiKeysResponse ListApiKeysResponse

// NewListApiKeysResponse instantiates a new ListApiKeysResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListApiKeysResponse(data []ApiKeyResponseData, params ListApiKeysParams) *ListApiKeysResponse {
	this := ListApiKeysResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewListApiKeysResponseWithDefaults instantiates a new ListApiKeysResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListApiKeysResponseWithDefaults() *ListApiKeysResponse {
	this := ListApiKeysResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListApiKeysResponse) GetData() []ApiKeyResponseData {
	if o == nil {
		var ret []ApiKeyResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListApiKeysResponse) GetDataOk() ([]ApiKeyResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListApiKeysResponse) SetData(v []ApiKeyResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *ListApiKeysResponse) GetParams() ListApiKeysParams {
	if o == nil {
		var ret ListApiKeysParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *ListApiKeysResponse) GetParamsOk() (*ListApiKeysParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *ListApiKeysResponse) SetParams(v ListApiKeysParams) {
	o.Params = v
}

func (o ListApiKeysResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListApiKeysResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *ListApiKeysResponse) UnmarshalJSON(data []byte) (err error) {
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

	varListApiKeysResponse := _ListApiKeysResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListApiKeysResponse)

	if err != nil {
		return err
	}

	*o = ListApiKeysResponse(varListApiKeysResponse)

	return err
}

type NullableListApiKeysResponse struct {
	value *ListApiKeysResponse
	isSet bool
}

func (v NullableListApiKeysResponse) Get() *ListApiKeysResponse {
	return v.value
}

func (v *NullableListApiKeysResponse) Set(val *ListApiKeysResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListApiKeysResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListApiKeysResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListApiKeysResponse(val *ListApiKeysResponse) *NullableListApiKeysResponse {
	return &NullableListApiKeysResponse{value: val, isSet: true}
}

func (v NullableListApiKeysResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListApiKeysResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
