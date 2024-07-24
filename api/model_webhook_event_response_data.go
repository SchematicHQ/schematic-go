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
	"time"
)

// checks if the WebhookEventResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookEventResponseData{}

// WebhookEventResponseData struct for WebhookEventResponseData
type WebhookEventResponseData struct {
	CreatedAt            time.Time      `json:"created_at"`
	Id                   string         `json:"id"`
	Payload              NullableString `json:"payload,omitempty"`
	RequestType          string         `json:"request_type"`
	ResponseCode         NullableInt32  `json:"response_code,omitempty"`
	SentAt               NullableTime   `json:"sent_at,omitempty"`
	Status               string         `json:"status"`
	UpdatedAt            time.Time      `json:"updated_at"`
	WebhookId            string         `json:"webhook_id"`
	AdditionalProperties map[string]interface{}
}

type _WebhookEventResponseData WebhookEventResponseData

// NewWebhookEventResponseData instantiates a new WebhookEventResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEventResponseData(createdAt time.Time, id string, requestType string, status string, updatedAt time.Time, webhookId string) *WebhookEventResponseData {
	this := WebhookEventResponseData{}
	this.CreatedAt = createdAt
	this.Id = id
	this.RequestType = requestType
	this.Status = status
	this.UpdatedAt = updatedAt
	this.WebhookId = webhookId
	return &this
}

// NewWebhookEventResponseDataWithDefaults instantiates a new WebhookEventResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEventResponseDataWithDefaults() *WebhookEventResponseData {
	this := WebhookEventResponseData{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *WebhookEventResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WebhookEventResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WebhookEventResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *WebhookEventResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhookEventResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhookEventResponseData) SetId(v string) {
	o.Id = v
}

// GetPayload returns the Payload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookEventResponseData) GetPayload() string {
	if o == nil || IsNil(o.Payload.Get()) {
		var ret string
		return ret
	}
	return *o.Payload.Get()
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookEventResponseData) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payload.Get(), o.Payload.IsSet()
}

// HasPayload returns a boolean if a field has been set.
func (o *WebhookEventResponseData) HasPayload() bool {
	if o != nil && o.Payload.IsSet() {
		return true
	}

	return false
}

// SetPayload gets a reference to the given NullableString and assigns it to the Payload field.
func (o *WebhookEventResponseData) SetPayload(v string) {
	o.Payload.Set(&v)
}

// SetPayloadNil sets the value for Payload to be an explicit nil
func (o *WebhookEventResponseData) SetPayloadNil() {
	o.Payload.Set(nil)
}

// UnsetPayload ensures that no value is present for Payload, not even an explicit nil
func (o *WebhookEventResponseData) UnsetPayload() {
	o.Payload.Unset()
}

// GetRequestType returns the RequestType field value
func (o *WebhookEventResponseData) GetRequestType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *WebhookEventResponseData) GetRequestTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value
func (o *WebhookEventResponseData) SetRequestType(v string) {
	o.RequestType = v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookEventResponseData) GetResponseCode() int32 {
	if o == nil || IsNil(o.ResponseCode.Get()) {
		var ret int32
		return ret
	}
	return *o.ResponseCode.Get()
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookEventResponseData) GetResponseCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseCode.Get(), o.ResponseCode.IsSet()
}

// HasResponseCode returns a boolean if a field has been set.
func (o *WebhookEventResponseData) HasResponseCode() bool {
	if o != nil && o.ResponseCode.IsSet() {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given NullableInt32 and assigns it to the ResponseCode field.
func (o *WebhookEventResponseData) SetResponseCode(v int32) {
	o.ResponseCode.Set(&v)
}

// SetResponseCodeNil sets the value for ResponseCode to be an explicit nil
func (o *WebhookEventResponseData) SetResponseCodeNil() {
	o.ResponseCode.Set(nil)
}

// UnsetResponseCode ensures that no value is present for ResponseCode, not even an explicit nil
func (o *WebhookEventResponseData) UnsetResponseCode() {
	o.ResponseCode.Unset()
}

// GetSentAt returns the SentAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookEventResponseData) GetSentAt() time.Time {
	if o == nil || IsNil(o.SentAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.SentAt.Get()
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookEventResponseData) GetSentAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SentAt.Get(), o.SentAt.IsSet()
}

// HasSentAt returns a boolean if a field has been set.
func (o *WebhookEventResponseData) HasSentAt() bool {
	if o != nil && o.SentAt.IsSet() {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given NullableTime and assigns it to the SentAt field.
func (o *WebhookEventResponseData) SetSentAt(v time.Time) {
	o.SentAt.Set(&v)
}

// SetSentAtNil sets the value for SentAt to be an explicit nil
func (o *WebhookEventResponseData) SetSentAtNil() {
	o.SentAt.Set(nil)
}

// UnsetSentAt ensures that no value is present for SentAt, not even an explicit nil
func (o *WebhookEventResponseData) UnsetSentAt() {
	o.SentAt.Unset()
}

// GetStatus returns the Status field value
func (o *WebhookEventResponseData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WebhookEventResponseData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WebhookEventResponseData) SetStatus(v string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *WebhookEventResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *WebhookEventResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *WebhookEventResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetWebhookId returns the WebhookId field value
func (o *WebhookEventResponseData) GetWebhookId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value
// and a boolean to check if the value has been set.
func (o *WebhookEventResponseData) GetWebhookIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebhookId, true
}

// SetWebhookId sets field value
func (o *WebhookEventResponseData) SetWebhookId(v string) {
	o.WebhookId = v
}

func (o WebhookEventResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookEventResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["id"] = o.Id
	if o.Payload.IsSet() {
		toSerialize["payload"] = o.Payload.Get()
	}
	toSerialize["request_type"] = o.RequestType
	if o.ResponseCode.IsSet() {
		toSerialize["response_code"] = o.ResponseCode.Get()
	}
	if o.SentAt.IsSet() {
		toSerialize["sent_at"] = o.SentAt.Get()
	}
	toSerialize["status"] = o.Status
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["webhook_id"] = o.WebhookId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebhookEventResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"id",
		"request_type",
		"status",
		"updated_at",
		"webhook_id",
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

	varWebhookEventResponseData := _WebhookEventResponseData{}

	err = json.Unmarshal(data, &varWebhookEventResponseData)

	if err != nil {
		return err
	}

	*o = WebhookEventResponseData(varWebhookEventResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "payload")
		delete(additionalProperties, "request_type")
		delete(additionalProperties, "response_code")
		delete(additionalProperties, "sent_at")
		delete(additionalProperties, "status")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "webhook_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebhookEventResponseData struct {
	value *WebhookEventResponseData
	isSet bool
}

func (v NullableWebhookEventResponseData) Get() *WebhookEventResponseData {
	return v.value
}

func (v *NullableWebhookEventResponseData) Set(val *WebhookEventResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEventResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEventResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEventResponseData(val *WebhookEventResponseData) *NullableWebhookEventResponseData {
	return &NullableWebhookEventResponseData{value: val, isSet: true}
}

func (v NullableWebhookEventResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEventResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
