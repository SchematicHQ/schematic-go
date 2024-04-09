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

// checks if the GetEventSummariesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEventSummariesResponse{}

// GetEventSummariesResponse struct for GetEventSummariesResponse
type GetEventSummariesResponse struct {
	// The returned resources
	Data   []EventSummaryResponseData `json:"data"`
	Params GetEventSummariesParams    `json:"params"`
}

type _GetEventSummariesResponse GetEventSummariesResponse

// NewGetEventSummariesResponse instantiates a new GetEventSummariesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEventSummariesResponse(data []EventSummaryResponseData, params GetEventSummariesParams) *GetEventSummariesResponse {
	this := GetEventSummariesResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewGetEventSummariesResponseWithDefaults instantiates a new GetEventSummariesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEventSummariesResponseWithDefaults() *GetEventSummariesResponse {
	this := GetEventSummariesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetEventSummariesResponse) GetData() []EventSummaryResponseData {
	if o == nil {
		var ret []EventSummaryResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetEventSummariesResponse) GetDataOk() ([]EventSummaryResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetEventSummariesResponse) SetData(v []EventSummaryResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *GetEventSummariesResponse) GetParams() GetEventSummariesParams {
	if o == nil {
		var ret GetEventSummariesParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *GetEventSummariesResponse) GetParamsOk() (*GetEventSummariesParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *GetEventSummariesResponse) SetParams(v GetEventSummariesParams) {
	o.Params = v
}

func (o GetEventSummariesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEventSummariesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *GetEventSummariesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetEventSummariesResponse := _GetEventSummariesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetEventSummariesResponse)

	if err != nil {
		return err
	}

	*o = GetEventSummariesResponse(varGetEventSummariesResponse)

	return err
}

type NullableGetEventSummariesResponse struct {
	value *GetEventSummariesResponse
	isSet bool
}

func (v NullableGetEventSummariesResponse) Get() *GetEventSummariesResponse {
	return v.value
}

func (v *NullableGetEventSummariesResponse) Set(val *GetEventSummariesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEventSummariesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEventSummariesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEventSummariesResponse(val *GetEventSummariesResponse) *NullableGetEventSummariesResponse {
	return &NullableGetEventSummariesResponse{value: val, isSet: true}
}

func (v NullableGetEventSummariesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEventSummariesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
