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

// checks if the EventResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventResponseData{}

// EventResponseData struct for EventResponseData
type EventResponseData struct {
	ApiKey               NullableString         `json:"api_key,omitempty"`
	Body                 map[string]interface{} `json:"body"`
	BodyPreview          string                 `json:"body_preview"`
	CapturedAt           time.Time              `json:"captured_at"`
	CompanyId            NullableString         `json:"company_id,omitempty"`
	EnrichedAt           NullableTime           `json:"enriched_at,omitempty"`
	EnvironmentId        NullableString         `json:"environment_id,omitempty"`
	ErrorMessage         NullableString         `json:"error_message,omitempty"`
	FeatureIds           []string               `json:"feature_ids"`
	Id                   string                 `json:"id"`
	LoadedAt             NullableTime           `json:"loaded_at,omitempty"`
	ProcessedAt          NullableTime           `json:"processed_at,omitempty"`
	SentAt               NullableTime           `json:"sent_at,omitempty"`
	Status               string                 `json:"status"`
	Subtype              NullableString         `json:"subtype,omitempty"`
	Type                 string                 `json:"type"`
	UpdatedAt            time.Time              `json:"updated_at"`
	UserId               NullableString         `json:"user_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EventResponseData EventResponseData

// NewEventResponseData instantiates a new EventResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventResponseData(body map[string]interface{}, bodyPreview string, capturedAt time.Time, featureIds []string, id string, status string, type_ string, updatedAt time.Time) *EventResponseData {
	this := EventResponseData{}
	this.Body = body
	this.BodyPreview = bodyPreview
	this.CapturedAt = capturedAt
	this.FeatureIds = featureIds
	this.Id = id
	this.Status = status
	this.Type = type_
	this.UpdatedAt = updatedAt
	return &this
}

// NewEventResponseDataWithDefaults instantiates a new EventResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventResponseDataWithDefaults() *EventResponseData {
	this := EventResponseData{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventResponseData) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey.Get()) {
		var ret string
		return ret
	}
	return *o.ApiKey.Get()
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventResponseData) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiKey.Get(), o.ApiKey.IsSet()
}

// HasApiKey returns a boolean if a field has been set.
func (o *EventResponseData) HasApiKey() bool {
	if o != nil && o.ApiKey.IsSet() {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given NullableString and assigns it to the ApiKey field.
func (o *EventResponseData) SetApiKey(v string) {
	o.ApiKey.Set(&v)
}

// SetApiKeyNil sets the value for ApiKey to be an explicit nil
func (o *EventResponseData) SetApiKeyNil() {
	o.ApiKey.Set(nil)
}

// UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
func (o *EventResponseData) UnsetApiKey() {
	o.ApiKey.Unset()
}

// GetBody returns the Body field value
func (o *EventResponseData) GetBody() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *EventResponseData) GetBodyOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Body, true
}

// SetBody sets field value
func (o *EventResponseData) SetBody(v map[string]interface{}) {
	o.Body = v
}

// GetBodyPreview returns the BodyPreview field value
func (o *EventResponseData) GetBodyPreview() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BodyPreview
}

// GetBodyPreviewOk returns a tuple with the BodyPreview field value
// and a boolean to check if the value has been set.
func (o *EventResponseData) GetBodyPreviewOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BodyPreview, true
}

// SetBodyPreview sets field value
func (o *EventResponseData) SetBodyPreview(v string) {
	o.BodyPreview = v
}

// GetCapturedAt returns the CapturedAt field value
func (o *EventResponseData) GetCapturedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CapturedAt
}

// GetCapturedAtOk returns a tuple with the CapturedAt field value
// and a boolean to check if the value has been set.
func (o *EventResponseData) GetCapturedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapturedAt, true
}

// SetCapturedAt sets field value
func (o *EventResponseData) SetCapturedAt(v time.Time) {
	o.CapturedAt = v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventResponseData) GetCompanyId() string {
	if o == nil || IsNil(o.CompanyId.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyId.Get()
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventResponseData) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyId.Get(), o.CompanyId.IsSet()
}

// HasCompanyId returns a boolean if a field has been set.
func (o *EventResponseData) HasCompanyId() bool {
	if o != nil && o.CompanyId.IsSet() {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given NullableString and assigns it to the CompanyId field.
func (o *EventResponseData) SetCompanyId(v string) {
	o.CompanyId.Set(&v)
}

// SetCompanyIdNil sets the value for CompanyId to be an explicit nil
func (o *EventResponseData) SetCompanyIdNil() {
	o.CompanyId.Set(nil)
}

// UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
func (o *EventResponseData) UnsetCompanyId() {
	o.CompanyId.Unset()
}

// GetEnrichedAt returns the EnrichedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventResponseData) GetEnrichedAt() time.Time {
	if o == nil || IsNil(o.EnrichedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EnrichedAt.Get()
}

// GetEnrichedAtOk returns a tuple with the EnrichedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventResponseData) GetEnrichedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrichedAt.Get(), o.EnrichedAt.IsSet()
}

// HasEnrichedAt returns a boolean if a field has been set.
func (o *EventResponseData) HasEnrichedAt() bool {
	if o != nil && o.EnrichedAt.IsSet() {
		return true
	}

	return false
}

// SetEnrichedAt gets a reference to the given NullableTime and assigns it to the EnrichedAt field.
func (o *EventResponseData) SetEnrichedAt(v time.Time) {
	o.EnrichedAt.Set(&v)
}

// SetEnrichedAtNil sets the value for EnrichedAt to be an explicit nil
func (o *EventResponseData) SetEnrichedAtNil() {
	o.EnrichedAt.Set(nil)
}

// UnsetEnrichedAt ensures that no value is present for EnrichedAt, not even an explicit nil
func (o *EventResponseData) UnsetEnrichedAt() {
	o.EnrichedAt.Unset()
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventResponseData) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId.Get()) {
		var ret string
		return ret
	}
	return *o.EnvironmentId.Get()
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventResponseData) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvironmentId.Get(), o.EnvironmentId.IsSet()
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *EventResponseData) HasEnvironmentId() bool {
	if o != nil && o.EnvironmentId.IsSet() {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given NullableString and assigns it to the EnvironmentId field.
func (o *EventResponseData) SetEnvironmentId(v string) {
	o.EnvironmentId.Set(&v)
}

// SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil
func (o *EventResponseData) SetEnvironmentIdNil() {
	o.EnvironmentId.Set(nil)
}

// UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
func (o *EventResponseData) UnsetEnvironmentId() {
	o.EnvironmentId.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventResponseData) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventResponseData) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *EventResponseData) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *EventResponseData) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}

// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *EventResponseData) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *EventResponseData) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetFeatureIds returns the FeatureIds field value
func (o *EventResponseData) GetFeatureIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FeatureIds
}

// GetFeatureIdsOk returns a tuple with the FeatureIds field value
// and a boolean to check if the value has been set.
func (o *EventResponseData) GetFeatureIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeatureIds, true
}

// SetFeatureIds sets field value
func (o *EventResponseData) SetFeatureIds(v []string) {
	o.FeatureIds = v
}

// GetId returns the Id field value
func (o *EventResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EventResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EventResponseData) SetId(v string) {
	o.Id = v
}

// GetLoadedAt returns the LoadedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventResponseData) GetLoadedAt() time.Time {
	if o == nil || IsNil(o.LoadedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LoadedAt.Get()
}

// GetLoadedAtOk returns a tuple with the LoadedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventResponseData) GetLoadedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoadedAt.Get(), o.LoadedAt.IsSet()
}

// HasLoadedAt returns a boolean if a field has been set.
func (o *EventResponseData) HasLoadedAt() bool {
	if o != nil && o.LoadedAt.IsSet() {
		return true
	}

	return false
}

// SetLoadedAt gets a reference to the given NullableTime and assigns it to the LoadedAt field.
func (o *EventResponseData) SetLoadedAt(v time.Time) {
	o.LoadedAt.Set(&v)
}

// SetLoadedAtNil sets the value for LoadedAt to be an explicit nil
func (o *EventResponseData) SetLoadedAtNil() {
	o.LoadedAt.Set(nil)
}

// UnsetLoadedAt ensures that no value is present for LoadedAt, not even an explicit nil
func (o *EventResponseData) UnsetLoadedAt() {
	o.LoadedAt.Unset()
}

// GetProcessedAt returns the ProcessedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventResponseData) GetProcessedAt() time.Time {
	if o == nil || IsNil(o.ProcessedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ProcessedAt.Get()
}

// GetProcessedAtOk returns a tuple with the ProcessedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventResponseData) GetProcessedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessedAt.Get(), o.ProcessedAt.IsSet()
}

// HasProcessedAt returns a boolean if a field has been set.
func (o *EventResponseData) HasProcessedAt() bool {
	if o != nil && o.ProcessedAt.IsSet() {
		return true
	}

	return false
}

// SetProcessedAt gets a reference to the given NullableTime and assigns it to the ProcessedAt field.
func (o *EventResponseData) SetProcessedAt(v time.Time) {
	o.ProcessedAt.Set(&v)
}

// SetProcessedAtNil sets the value for ProcessedAt to be an explicit nil
func (o *EventResponseData) SetProcessedAtNil() {
	o.ProcessedAt.Set(nil)
}

// UnsetProcessedAt ensures that no value is present for ProcessedAt, not even an explicit nil
func (o *EventResponseData) UnsetProcessedAt() {
	o.ProcessedAt.Unset()
}

// GetSentAt returns the SentAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventResponseData) GetSentAt() time.Time {
	if o == nil || IsNil(o.SentAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.SentAt.Get()
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventResponseData) GetSentAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SentAt.Get(), o.SentAt.IsSet()
}

// HasSentAt returns a boolean if a field has been set.
func (o *EventResponseData) HasSentAt() bool {
	if o != nil && o.SentAt.IsSet() {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given NullableTime and assigns it to the SentAt field.
func (o *EventResponseData) SetSentAt(v time.Time) {
	o.SentAt.Set(&v)
}

// SetSentAtNil sets the value for SentAt to be an explicit nil
func (o *EventResponseData) SetSentAtNil() {
	o.SentAt.Set(nil)
}

// UnsetSentAt ensures that no value is present for SentAt, not even an explicit nil
func (o *EventResponseData) UnsetSentAt() {
	o.SentAt.Unset()
}

// GetStatus returns the Status field value
func (o *EventResponseData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EventResponseData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EventResponseData) SetStatus(v string) {
	o.Status = v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventResponseData) GetSubtype() string {
	if o == nil || IsNil(o.Subtype.Get()) {
		var ret string
		return ret
	}
	return *o.Subtype.Get()
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventResponseData) GetSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subtype.Get(), o.Subtype.IsSet()
}

// HasSubtype returns a boolean if a field has been set.
func (o *EventResponseData) HasSubtype() bool {
	if o != nil && o.Subtype.IsSet() {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given NullableString and assigns it to the Subtype field.
func (o *EventResponseData) SetSubtype(v string) {
	o.Subtype.Set(&v)
}

// SetSubtypeNil sets the value for Subtype to be an explicit nil
func (o *EventResponseData) SetSubtypeNil() {
	o.Subtype.Set(nil)
}

// UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
func (o *EventResponseData) UnsetSubtype() {
	o.Subtype.Unset()
}

// GetType returns the Type field value
func (o *EventResponseData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EventResponseData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EventResponseData) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *EventResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *EventResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *EventResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventResponseData) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventResponseData) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *EventResponseData) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *EventResponseData) SetUserId(v string) {
	o.UserId.Set(&v)
}

// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *EventResponseData) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *EventResponseData) UnsetUserId() {
	o.UserId.Unset()
}

func (o EventResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiKey.IsSet() {
		toSerialize["api_key"] = o.ApiKey.Get()
	}
	toSerialize["body"] = o.Body
	toSerialize["body_preview"] = o.BodyPreview
	toSerialize["captured_at"] = o.CapturedAt
	if o.CompanyId.IsSet() {
		toSerialize["company_id"] = o.CompanyId.Get()
	}
	if o.EnrichedAt.IsSet() {
		toSerialize["enriched_at"] = o.EnrichedAt.Get()
	}
	if o.EnvironmentId.IsSet() {
		toSerialize["environment_id"] = o.EnvironmentId.Get()
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["error_message"] = o.ErrorMessage.Get()
	}
	toSerialize["feature_ids"] = o.FeatureIds
	toSerialize["id"] = o.Id
	if o.LoadedAt.IsSet() {
		toSerialize["loaded_at"] = o.LoadedAt.Get()
	}
	if o.ProcessedAt.IsSet() {
		toSerialize["processed_at"] = o.ProcessedAt.Get()
	}
	if o.SentAt.IsSet() {
		toSerialize["sent_at"] = o.SentAt.Get()
	}
	toSerialize["status"] = o.Status
	if o.Subtype.IsSet() {
		toSerialize["subtype"] = o.Subtype.Get()
	}
	toSerialize["type"] = o.Type
	toSerialize["updated_at"] = o.UpdatedAt
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"body",
		"body_preview",
		"captured_at",
		"feature_ids",
		"id",
		"status",
		"type",
		"updated_at",
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

	varEventResponseData := _EventResponseData{}

	err = json.Unmarshal(data, &varEventResponseData)

	if err != nil {
		return err
	}

	*o = EventResponseData(varEventResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "api_key")
		delete(additionalProperties, "body")
		delete(additionalProperties, "body_preview")
		delete(additionalProperties, "captured_at")
		delete(additionalProperties, "company_id")
		delete(additionalProperties, "enriched_at")
		delete(additionalProperties, "environment_id")
		delete(additionalProperties, "error_message")
		delete(additionalProperties, "feature_ids")
		delete(additionalProperties, "id")
		delete(additionalProperties, "loaded_at")
		delete(additionalProperties, "processed_at")
		delete(additionalProperties, "sent_at")
		delete(additionalProperties, "status")
		delete(additionalProperties, "subtype")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "user_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventResponseData struct {
	value *EventResponseData
	isSet bool
}

func (v NullableEventResponseData) Get() *EventResponseData {
	return v.value
}

func (v *NullableEventResponseData) Set(val *EventResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableEventResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableEventResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventResponseData(val *EventResponseData) *NullableEventResponseData {
	return &NullableEventResponseData{value: val, isSet: true}
}

func (v NullableEventResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
