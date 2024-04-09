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

// checks if the GetEventSummaryBySubtypeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEventSummaryBySubtypeResponse{}

// GetEventSummaryBySubtypeResponse struct for GetEventSummaryBySubtypeResponse
type GetEventSummaryBySubtypeResponse struct {
	Data EventSummaryResponseData `json:"data"`
	// Input parameters
	Params map[string]interface{} `json:"params"`
}

type _GetEventSummaryBySubtypeResponse GetEventSummaryBySubtypeResponse

// NewGetEventSummaryBySubtypeResponse instantiates a new GetEventSummaryBySubtypeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEventSummaryBySubtypeResponse(data EventSummaryResponseData, params map[string]interface{}) *GetEventSummaryBySubtypeResponse {
	this := GetEventSummaryBySubtypeResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewGetEventSummaryBySubtypeResponseWithDefaults instantiates a new GetEventSummaryBySubtypeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEventSummaryBySubtypeResponseWithDefaults() *GetEventSummaryBySubtypeResponse {
	this := GetEventSummaryBySubtypeResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetEventSummaryBySubtypeResponse) GetData() EventSummaryResponseData {
	if o == nil {
		var ret EventSummaryResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetEventSummaryBySubtypeResponse) GetDataOk() (*EventSummaryResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetEventSummaryBySubtypeResponse) SetData(v EventSummaryResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *GetEventSummaryBySubtypeResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *GetEventSummaryBySubtypeResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *GetEventSummaryBySubtypeResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o GetEventSummaryBySubtypeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEventSummaryBySubtypeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *GetEventSummaryBySubtypeResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetEventSummaryBySubtypeResponse := _GetEventSummaryBySubtypeResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetEventSummaryBySubtypeResponse)

	if err != nil {
		return err
	}

	*o = GetEventSummaryBySubtypeResponse(varGetEventSummaryBySubtypeResponse)

	return err
}

type NullableGetEventSummaryBySubtypeResponse struct {
	value *GetEventSummaryBySubtypeResponse
	isSet bool
}

func (v NullableGetEventSummaryBySubtypeResponse) Get() *GetEventSummaryBySubtypeResponse {
	return v.value
}

func (v *NullableGetEventSummaryBySubtypeResponse) Set(val *GetEventSummaryBySubtypeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEventSummaryBySubtypeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEventSummaryBySubtypeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEventSummaryBySubtypeResponse(val *GetEventSummaryBySubtypeResponse) *NullableGetEventSummaryBySubtypeResponse {
	return &NullableGetEventSummaryBySubtypeResponse{value: val, isSet: true}
}

func (v NullableGetEventSummaryBySubtypeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEventSummaryBySubtypeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
