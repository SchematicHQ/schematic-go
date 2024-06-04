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

// checks if the UpdateWebhookResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateWebhookResponse{}

// UpdateWebhookResponse struct for UpdateWebhookResponse
type UpdateWebhookResponse struct {
	Data WebhookResponseData `json:"data"`
	// Input parameters
	Params map[string]interface{} `json:"params"`
}

type _UpdateWebhookResponse UpdateWebhookResponse

// NewUpdateWebhookResponse instantiates a new UpdateWebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWebhookResponse(data WebhookResponseData, params map[string]interface{}) *UpdateWebhookResponse {
	this := UpdateWebhookResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewUpdateWebhookResponseWithDefaults instantiates a new UpdateWebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWebhookResponseWithDefaults() *UpdateWebhookResponse {
	this := UpdateWebhookResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UpdateWebhookResponse) GetData() WebhookResponseData {
	if o == nil {
		var ret WebhookResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UpdateWebhookResponse) GetDataOk() (*WebhookResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UpdateWebhookResponse) SetData(v WebhookResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *UpdateWebhookResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *UpdateWebhookResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *UpdateWebhookResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o UpdateWebhookResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateWebhookResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params
	return toSerialize, nil
}

func (o *UpdateWebhookResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateWebhookResponse := _UpdateWebhookResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateWebhookResponse)

	if err != nil {
		return err
	}

	*o = UpdateWebhookResponse(varUpdateWebhookResponse)

	return err
}

type NullableUpdateWebhookResponse struct {
	value *UpdateWebhookResponse
	isSet bool
}

func (v NullableUpdateWebhookResponse) Get() *UpdateWebhookResponse {
	return v.value
}

func (v *NullableUpdateWebhookResponse) Set(val *UpdateWebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWebhookResponse(val *UpdateWebhookResponse) *NullableUpdateWebhookResponse {
	return &NullableUpdateWebhookResponse{value: val, isSet: true}
}

func (v NullableUpdateWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
