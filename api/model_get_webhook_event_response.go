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

// checks if the GetWebhookEventResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetWebhookEventResponse{}

// GetWebhookEventResponse struct for GetWebhookEventResponse
type GetWebhookEventResponse struct {
	Data WebhookEventDetailResponseData `json:"data"`
	// Input parameters
	Params               map[string]interface{} `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _GetWebhookEventResponse GetWebhookEventResponse

// NewGetWebhookEventResponse instantiates a new GetWebhookEventResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWebhookEventResponse(data WebhookEventDetailResponseData, params map[string]interface{}) *GetWebhookEventResponse {
	this := GetWebhookEventResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewGetWebhookEventResponseWithDefaults instantiates a new GetWebhookEventResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWebhookEventResponseWithDefaults() *GetWebhookEventResponse {
	this := GetWebhookEventResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetWebhookEventResponse) GetData() WebhookEventDetailResponseData {
	if o == nil {
		var ret WebhookEventDetailResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetWebhookEventResponse) GetDataOk() (*WebhookEventDetailResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetWebhookEventResponse) SetData(v WebhookEventDetailResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *GetWebhookEventResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *GetWebhookEventResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *GetWebhookEventResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o GetWebhookEventResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWebhookEventResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetWebhookEventResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetWebhookEventResponse := _GetWebhookEventResponse{}

	err = json.Unmarshal(data, &varGetWebhookEventResponse)

	if err != nil {
		return err
	}

	*o = GetWebhookEventResponse(varGetWebhookEventResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetWebhookEventResponse struct {
	value *GetWebhookEventResponse
	isSet bool
}

func (v NullableGetWebhookEventResponse) Get() *GetWebhookEventResponse {
	return v.value
}

func (v *NullableGetWebhookEventResponse) Set(val *GetWebhookEventResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWebhookEventResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWebhookEventResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWebhookEventResponse(val *GetWebhookEventResponse) *NullableGetWebhookEventResponse {
	return &NullableGetWebhookEventResponse{value: val, isSet: true}
}

func (v NullableGetWebhookEventResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWebhookEventResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}