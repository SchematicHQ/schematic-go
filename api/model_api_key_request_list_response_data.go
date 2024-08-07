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

// checks if the ApiKeyRequestListResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyRequestListResponseData{}

// ApiKeyRequestListResponseData struct for ApiKeyRequestListResponseData
type ApiKeyRequestListResponseData struct {
	ApiKeyId             string         `json:"api_key_id"`
	EndedAt              NullableTime   `json:"ended_at,omitempty"`
	EnvironmentId        NullableString `json:"environment_id,omitempty"`
	Id                   string         `json:"id"`
	Method               string         `json:"method"`
	ReqBody              NullableString `json:"req_body,omitempty"`
	RequestType          NullableString `json:"request_type,omitempty"`
	ResourceId           NullableInt32  `json:"resource_id,omitempty"`
	ResourceIdString     NullableString `json:"resource_id_string,omitempty"`
	ResourceName         NullableString `json:"resource_name,omitempty"`
	ResourceType         NullableString `json:"resource_type,omitempty"`
	RespCode             NullableInt32  `json:"resp_code,omitempty"`
	SecondaryResource    NullableString `json:"secondary_resource,omitempty"`
	StartedAt            time.Time      `json:"started_at"`
	Url                  string         `json:"url"`
	UserName             NullableString `json:"user_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApiKeyRequestListResponseData ApiKeyRequestListResponseData

// NewApiKeyRequestListResponseData instantiates a new ApiKeyRequestListResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyRequestListResponseData(apiKeyId string, id string, method string, startedAt time.Time, url string) *ApiKeyRequestListResponseData {
	this := ApiKeyRequestListResponseData{}
	this.ApiKeyId = apiKeyId
	this.Id = id
	this.Method = method
	this.StartedAt = startedAt
	this.Url = url
	return &this
}

// NewApiKeyRequestListResponseDataWithDefaults instantiates a new ApiKeyRequestListResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyRequestListResponseDataWithDefaults() *ApiKeyRequestListResponseData {
	this := ApiKeyRequestListResponseData{}
	return &this
}

// GetApiKeyId returns the ApiKeyId field value
func (o *ApiKeyRequestListResponseData) GetApiKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKeyId
}

// GetApiKeyIdOk returns a tuple with the ApiKeyId field value
// and a boolean to check if the value has been set.
func (o *ApiKeyRequestListResponseData) GetApiKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKeyId, true
}

// SetApiKeyId sets field value
func (o *ApiKeyRequestListResponseData) SetApiKeyId(v string) {
	o.ApiKeyId = v
}

// GetEndedAt returns the EndedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKeyRequestListResponseData) GetEndedAt() time.Time {
	if o == nil || IsNil(o.EndedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EndedAt.Get()
}

// GetEndedAtOk returns a tuple with the EndedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyRequestListResponseData) GetEndedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndedAt.Get(), o.EndedAt.IsSet()
}

// HasEndedAt returns a boolean if a field has been set.
func (o *ApiKeyRequestListResponseData) HasEndedAt() bool {
	if o != nil && o.EndedAt.IsSet() {
		return true
	}

	return false
}

// SetEndedAt gets a reference to the given NullableTime and assigns it to the EndedAt field.
func (o *ApiKeyRequestListResponseData) SetEndedAt(v time.Time) {
	o.EndedAt.Set(&v)
}

// SetEndedAtNil sets the value for EndedAt to be an explicit nil
func (o *ApiKeyRequestListResponseData) SetEndedAtNil() {
	o.EndedAt.Set(nil)
}

// UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
func (o *ApiKeyRequestListResponseData) UnsetEndedAt() {
	o.EndedAt.Unset()
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKeyRequestListResponseData) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId.Get()) {
		var ret string
		return ret
	}
	return *o.EnvironmentId.Get()
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyRequestListResponseData) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvironmentId.Get(), o.EnvironmentId.IsSet()
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *ApiKeyRequestListResponseData) HasEnvironmentId() bool {
	if o != nil && o.EnvironmentId.IsSet() {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given NullableString and assigns it to the EnvironmentId field.
func (o *ApiKeyRequestListResponseData) SetEnvironmentId(v string) {
	o.EnvironmentId.Set(&v)
}

// SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil
func (o *ApiKeyRequestListResponseData) SetEnvironmentIdNil() {
	o.EnvironmentId.Set(nil)
}

// UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
func (o *ApiKeyRequestListResponseData) UnsetEnvironmentId() {
	o.EnvironmentId.Unset()
}

// GetId returns the Id field value
func (o *ApiKeyRequestListResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiKeyRequestListResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiKeyRequestListResponseData) SetId(v string) {
	o.Id = v
}

// GetMethod returns the Method field value
func (o *ApiKeyRequestListResponseData) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *ApiKeyRequestListResponseData) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *ApiKeyRequestListResponseData) SetMethod(v string) {
	o.Method = v
}

// GetReqBody returns the ReqBody field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKeyRequestListResponseData) GetReqBody() string {
	if o == nil || IsNil(o.ReqBody.Get()) {
		var ret string
		return ret
	}
	return *o.ReqBody.Get()
}

// GetReqBodyOk returns a tuple with the ReqBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyRequestListResponseData) GetReqBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReqBody.Get(), o.ReqBody.IsSet()
}

// HasReqBody returns a boolean if a field has been set.
func (o *ApiKeyRequestListResponseData) HasReqBody() bool {
	if o != nil && o.ReqBody.IsSet() {
		return true
	}

	return false
}

// SetReqBody gets a reference to the given NullableString and assigns it to the ReqBody field.
func (o *ApiKeyRequestListResponseData) SetReqBody(v string) {
	o.ReqBody.Set(&v)
}

// SetReqBodyNil sets the value for ReqBody to be an explicit nil
func (o *ApiKeyRequestListResponseData) SetReqBodyNil() {
	o.ReqBody.Set(nil)
}

// UnsetReqBody ensures that no value is present for ReqBody, not even an explicit nil
func (o *ApiKeyRequestListResponseData) UnsetReqBody() {
	o.ReqBody.Unset()
}

// GetRequestType returns the RequestType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKeyRequestListResponseData) GetRequestType() string {
	if o == nil || IsNil(o.RequestType.Get()) {
		var ret string
		return ret
	}
	return *o.RequestType.Get()
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyRequestListResponseData) GetRequestTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestType.Get(), o.RequestType.IsSet()
}

// HasRequestType returns a boolean if a field has been set.
func (o *ApiKeyRequestListResponseData) HasRequestType() bool {
	if o != nil && o.RequestType.IsSet() {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given NullableString and assigns it to the RequestType field.
func (o *ApiKeyRequestListResponseData) SetRequestType(v string) {
	o.RequestType.Set(&v)
}

// SetRequestTypeNil sets the value for RequestType to be an explicit nil
func (o *ApiKeyRequestListResponseData) SetRequestTypeNil() {
	o.RequestType.Set(nil)
}

// UnsetRequestType ensures that no value is present for RequestType, not even an explicit nil
func (o *ApiKeyRequestListResponseData) UnsetRequestType() {
	o.RequestType.Unset()
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKeyRequestListResponseData) GetResourceId() int32 {
	if o == nil || IsNil(o.ResourceId.Get()) {
		var ret int32
		return ret
	}
	return *o.ResourceId.Get()
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyRequestListResponseData) GetResourceIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceId.Get(), o.ResourceId.IsSet()
}

// HasResourceId returns a boolean if a field has been set.
func (o *ApiKeyRequestListResponseData) HasResourceId() bool {
	if o != nil && o.ResourceId.IsSet() {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given NullableInt32 and assigns it to the ResourceId field.
func (o *ApiKeyRequestListResponseData) SetResourceId(v int32) {
	o.ResourceId.Set(&v)
}

// SetResourceIdNil sets the value for ResourceId to be an explicit nil
func (o *ApiKeyRequestListResponseData) SetResourceIdNil() {
	o.ResourceId.Set(nil)
}

// UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
func (o *ApiKeyRequestListResponseData) UnsetResourceId() {
	o.ResourceId.Unset()
}

// GetResourceIdString returns the ResourceIdString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKeyRequestListResponseData) GetResourceIdString() string {
	if o == nil || IsNil(o.ResourceIdString.Get()) {
		var ret string
		return ret
	}
	return *o.ResourceIdString.Get()
}

// GetResourceIdStringOk returns a tuple with the ResourceIdString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyRequestListResponseData) GetResourceIdStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceIdString.Get(), o.ResourceIdString.IsSet()
}

// HasResourceIdString returns a boolean if a field has been set.
func (o *ApiKeyRequestListResponseData) HasResourceIdString() bool {
	if o != nil && o.ResourceIdString.IsSet() {
		return true
	}

	return false
}

// SetResourceIdString gets a reference to the given NullableString and assigns it to the ResourceIdString field.
func (o *ApiKeyRequestListResponseData) SetResourceIdString(v string) {
	o.ResourceIdString.Set(&v)
}

// SetResourceIdStringNil sets the value for ResourceIdString to be an explicit nil
func (o *ApiKeyRequestListResponseData) SetResourceIdStringNil() {
	o.ResourceIdString.Set(nil)
}

// UnsetResourceIdString ensures that no value is present for ResourceIdString, not even an explicit nil
func (o *ApiKeyRequestListResponseData) UnsetResourceIdString() {
	o.ResourceIdString.Unset()
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKeyRequestListResponseData) GetResourceName() string {
	if o == nil || IsNil(o.ResourceName.Get()) {
		var ret string
		return ret
	}
	return *o.ResourceName.Get()
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyRequestListResponseData) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceName.Get(), o.ResourceName.IsSet()
}

// HasResourceName returns a boolean if a field has been set.
func (o *ApiKeyRequestListResponseData) HasResourceName() bool {
	if o != nil && o.ResourceName.IsSet() {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given NullableString and assigns it to the ResourceName field.
func (o *ApiKeyRequestListResponseData) SetResourceName(v string) {
	o.ResourceName.Set(&v)
}

// SetResourceNameNil sets the value for ResourceName to be an explicit nil
func (o *ApiKeyRequestListResponseData) SetResourceNameNil() {
	o.ResourceName.Set(nil)
}

// UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
func (o *ApiKeyRequestListResponseData) UnsetResourceName() {
	o.ResourceName.Unset()
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKeyRequestListResponseData) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType.Get()) {
		var ret string
		return ret
	}
	return *o.ResourceType.Get()
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyRequestListResponseData) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceType.Get(), o.ResourceType.IsSet()
}

// HasResourceType returns a boolean if a field has been set.
func (o *ApiKeyRequestListResponseData) HasResourceType() bool {
	if o != nil && o.ResourceType.IsSet() {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given NullableString and assigns it to the ResourceType field.
func (o *ApiKeyRequestListResponseData) SetResourceType(v string) {
	o.ResourceType.Set(&v)
}

// SetResourceTypeNil sets the value for ResourceType to be an explicit nil
func (o *ApiKeyRequestListResponseData) SetResourceTypeNil() {
	o.ResourceType.Set(nil)
}

// UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
func (o *ApiKeyRequestListResponseData) UnsetResourceType() {
	o.ResourceType.Unset()
}

// GetRespCode returns the RespCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKeyRequestListResponseData) GetRespCode() int32 {
	if o == nil || IsNil(o.RespCode.Get()) {
		var ret int32
		return ret
	}
	return *o.RespCode.Get()
}

// GetRespCodeOk returns a tuple with the RespCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyRequestListResponseData) GetRespCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RespCode.Get(), o.RespCode.IsSet()
}

// HasRespCode returns a boolean if a field has been set.
func (o *ApiKeyRequestListResponseData) HasRespCode() bool {
	if o != nil && o.RespCode.IsSet() {
		return true
	}

	return false
}

// SetRespCode gets a reference to the given NullableInt32 and assigns it to the RespCode field.
func (o *ApiKeyRequestListResponseData) SetRespCode(v int32) {
	o.RespCode.Set(&v)
}

// SetRespCodeNil sets the value for RespCode to be an explicit nil
func (o *ApiKeyRequestListResponseData) SetRespCodeNil() {
	o.RespCode.Set(nil)
}

// UnsetRespCode ensures that no value is present for RespCode, not even an explicit nil
func (o *ApiKeyRequestListResponseData) UnsetRespCode() {
	o.RespCode.Unset()
}

// GetSecondaryResource returns the SecondaryResource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKeyRequestListResponseData) GetSecondaryResource() string {
	if o == nil || IsNil(o.SecondaryResource.Get()) {
		var ret string
		return ret
	}
	return *o.SecondaryResource.Get()
}

// GetSecondaryResourceOk returns a tuple with the SecondaryResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyRequestListResponseData) GetSecondaryResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecondaryResource.Get(), o.SecondaryResource.IsSet()
}

// HasSecondaryResource returns a boolean if a field has been set.
func (o *ApiKeyRequestListResponseData) HasSecondaryResource() bool {
	if o != nil && o.SecondaryResource.IsSet() {
		return true
	}

	return false
}

// SetSecondaryResource gets a reference to the given NullableString and assigns it to the SecondaryResource field.
func (o *ApiKeyRequestListResponseData) SetSecondaryResource(v string) {
	o.SecondaryResource.Set(&v)
}

// SetSecondaryResourceNil sets the value for SecondaryResource to be an explicit nil
func (o *ApiKeyRequestListResponseData) SetSecondaryResourceNil() {
	o.SecondaryResource.Set(nil)
}

// UnsetSecondaryResource ensures that no value is present for SecondaryResource, not even an explicit nil
func (o *ApiKeyRequestListResponseData) UnsetSecondaryResource() {
	o.SecondaryResource.Unset()
}

// GetStartedAt returns the StartedAt field value
func (o *ApiKeyRequestListResponseData) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *ApiKeyRequestListResponseData) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *ApiKeyRequestListResponseData) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetUrl returns the Url field value
func (o *ApiKeyRequestListResponseData) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ApiKeyRequestListResponseData) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ApiKeyRequestListResponseData) SetUrl(v string) {
	o.Url = v
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiKeyRequestListResponseData) GetUserName() string {
	if o == nil || IsNil(o.UserName.Get()) {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiKeyRequestListResponseData) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *ApiKeyRequestListResponseData) HasUserName() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *ApiKeyRequestListResponseData) SetUserName(v string) {
	o.UserName.Set(&v)
}

// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *ApiKeyRequestListResponseData) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *ApiKeyRequestListResponseData) UnsetUserName() {
	o.UserName.Unset()
}

func (o ApiKeyRequestListResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyRequestListResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["api_key_id"] = o.ApiKeyId
	if o.EndedAt.IsSet() {
		toSerialize["ended_at"] = o.EndedAt.Get()
	}
	if o.EnvironmentId.IsSet() {
		toSerialize["environment_id"] = o.EnvironmentId.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["method"] = o.Method
	if o.ReqBody.IsSet() {
		toSerialize["req_body"] = o.ReqBody.Get()
	}
	if o.RequestType.IsSet() {
		toSerialize["request_type"] = o.RequestType.Get()
	}
	if o.ResourceId.IsSet() {
		toSerialize["resource_id"] = o.ResourceId.Get()
	}
	if o.ResourceIdString.IsSet() {
		toSerialize["resource_id_string"] = o.ResourceIdString.Get()
	}
	if o.ResourceName.IsSet() {
		toSerialize["resource_name"] = o.ResourceName.Get()
	}
	if o.ResourceType.IsSet() {
		toSerialize["resource_type"] = o.ResourceType.Get()
	}
	if o.RespCode.IsSet() {
		toSerialize["resp_code"] = o.RespCode.Get()
	}
	if o.SecondaryResource.IsSet() {
		toSerialize["secondary_resource"] = o.SecondaryResource.Get()
	}
	toSerialize["started_at"] = o.StartedAt
	toSerialize["url"] = o.Url
	if o.UserName.IsSet() {
		toSerialize["user_name"] = o.UserName.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApiKeyRequestListResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"api_key_id",
		"id",
		"method",
		"started_at",
		"url",
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

	varApiKeyRequestListResponseData := _ApiKeyRequestListResponseData{}

	err = json.Unmarshal(data, &varApiKeyRequestListResponseData)

	if err != nil {
		return err
	}

	*o = ApiKeyRequestListResponseData(varApiKeyRequestListResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "api_key_id")
		delete(additionalProperties, "ended_at")
		delete(additionalProperties, "environment_id")
		delete(additionalProperties, "id")
		delete(additionalProperties, "method")
		delete(additionalProperties, "req_body")
		delete(additionalProperties, "request_type")
		delete(additionalProperties, "resource_id")
		delete(additionalProperties, "resource_id_string")
		delete(additionalProperties, "resource_name")
		delete(additionalProperties, "resource_type")
		delete(additionalProperties, "resp_code")
		delete(additionalProperties, "secondary_resource")
		delete(additionalProperties, "started_at")
		delete(additionalProperties, "url")
		delete(additionalProperties, "user_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApiKeyRequestListResponseData struct {
	value *ApiKeyRequestListResponseData
	isSet bool
}

func (v NullableApiKeyRequestListResponseData) Get() *ApiKeyRequestListResponseData {
	return v.value
}

func (v *NullableApiKeyRequestListResponseData) Set(val *ApiKeyRequestListResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyRequestListResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyRequestListResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyRequestListResponseData(val *ApiKeyRequestListResponseData) *NullableApiKeyRequestListResponseData {
	return &NullableApiKeyRequestListResponseData{value: val, isSet: true}
}

func (v NullableApiKeyRequestListResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyRequestListResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
