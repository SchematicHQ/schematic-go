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

// checks if the GetWebhookResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetWebhookResponse{}

// GetWebhookResponse struct for GetWebhookResponse
type GetWebhookResponse struct {
	Data WebhookResponseData `json:"data"`
	// Input parameters
	Params               map[string]interface{} `json:"params"`
	AdditionalProperties map[string]interface{}
}

type _GetWebhookResponse GetWebhookResponse

// NewGetWebhookResponse instantiates a new GetWebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetWebhookResponse(data WebhookResponseData, params map[string]interface{}) *GetWebhookResponse {
	this := GetWebhookResponse{}
	this.Data = data
	this.Params = params
	return &this
}

// NewGetWebhookResponseWithDefaults instantiates a new GetWebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetWebhookResponseWithDefaults() *GetWebhookResponse {
	this := GetWebhookResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetWebhookResponse) GetData() WebhookResponseData {
	if o == nil {
		var ret WebhookResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetWebhookResponse) GetDataOk() (*WebhookResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetWebhookResponse) SetData(v WebhookResponseData) {
	o.Data = v
}

// GetParams returns the Params field value
func (o *GetWebhookResponse) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *GetWebhookResponse) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *GetWebhookResponse) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o GetWebhookResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetWebhookResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetWebhookResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetWebhookResponse := _GetWebhookResponse{}

	err = json.Unmarshal(data, &varGetWebhookResponse)

	if err != nil {
		return err
	}

	*o = GetWebhookResponse(varGetWebhookResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetWebhookResponse struct {
	value *GetWebhookResponse
	isSet bool
}

func (v NullableGetWebhookResponse) Get() *GetWebhookResponse {
	return v.value
}

func (v *NullableGetWebhookResponse) Set(val *GetWebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetWebhookResponse(val *GetWebhookResponse) *NullableGetWebhookResponse {
	return &NullableGetWebhookResponse{value: val, isSet: true}
}

func (v NullableGetWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}