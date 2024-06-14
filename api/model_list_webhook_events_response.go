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

// checks if the ListWebhookEventsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWebhookEventsResponse{}

// ListWebhookEventsResponse struct for ListWebhookEventsResponse
type ListWebhookEventsResponse struct {
	// The returned resources
	Data   []WebhookEventDetailResponseData `json:"data"`
	Params ListWebhookEventsParams          `json:"params"`
}

type _ListWebhookEventsResponse ListWebhookEventsResponse

// NewListWebhookEventsResponse instantiates a new ListWebhookEventsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWebhookEventsResponse(data []WebhookEventDetailResponseData, params ListWebhookEventsParams) *ListWebhookEventsResponse {
	this := ListWebhookEventsResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewListWebhookEventsResponseWithDefaults instantiates a new ListWebhookEventsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWebhookEventsResponseWithDefaults() *ListWebhookEventsResponse {
	this := ListWebhookEventsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListWebhookEventsResponse) GetData() []WebhookEventDetailResponseData {
	if o == nil {
		var ret []WebhookEventDetailResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListWebhookEventsResponse) GetDataOk() ([]WebhookEventDetailResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListWebhookEventsResponse) SetData(v []WebhookEventDetailResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *ListWebhookEventsResponse) GetParams() ListWebhookEventsParams {
	if o == nil {
		var ret ListWebhookEventsParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *ListWebhookEventsResponse) GetParamsOk() (*ListWebhookEventsParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *ListWebhookEventsResponse) SetParams(v ListWebhookEventsParams) {
	o.Params = v
}

func (o ListWebhookEventsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWebhookEventsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *ListWebhookEventsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varListWebhookEventsResponse := _ListWebhookEventsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListWebhookEventsResponse)

	if err != nil {
		return err
	}

	*o = ListWebhookEventsResponse(varListWebhookEventsResponse)

	return err
}

type NullableListWebhookEventsResponse struct {
	value *ListWebhookEventsResponse
	isSet bool
}

func (v NullableListWebhookEventsResponse) Get() *ListWebhookEventsResponse {
	return v.value
}

func (v *NullableListWebhookEventsResponse) Set(val *ListWebhookEventsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListWebhookEventsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListWebhookEventsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWebhookEventsResponse(val *ListWebhookEventsResponse) *NullableListWebhookEventsResponse {
	return &NullableListWebhookEventsResponse{value: val, isSet: true}
}

func (v NullableListWebhookEventsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWebhookEventsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
