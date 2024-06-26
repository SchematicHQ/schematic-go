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
	"time"
)

// checks if the FlagCheckLogDetailResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlagCheckLogDetailResponseData{}

// FlagCheckLogDetailResponseData struct for FlagCheckLogDetailResponseData
type FlagCheckLogDetailResponseData struct {
	CheckStatus   string                   `json:"check_status"`
	Company       *CompanyResponseData     `json:"company,omitempty"`
	CompanyId     NullableString           `json:"company_id,omitempty"`
	CreatedAt     time.Time                `json:"created_at"`
	Environment   *EnvironmentResponseData `json:"environment,omitempty"`
	EnvironmentId string                   `json:"environment_id"`
	Error         NullableString           `json:"error,omitempty"`
	Flag          *FlagResponseData        `json:"flag,omitempty"`
	FlagId        NullableString           `json:"flag_id,omitempty"`
	FlagKey       string                   `json:"flag_key"`
	Id            string                   `json:"id"`
	Reason        string                   `json:"reason"`
	ReqCompany    map[string]string        `json:"req_company,omitempty"`
	ReqUser       map[string]string        `json:"req_user,omitempty"`
	Rule          *RuleResponseData        `json:"rule,omitempty"`
	RuleId        NullableString           `json:"rule_id,omitempty"`
	UpdatedAt     time.Time                `json:"updated_at"`
	User          *UserResponseData        `json:"user,omitempty"`
	UserId        NullableString           `json:"user_id,omitempty"`
	Value         bool                     `json:"value"`
}

type _FlagCheckLogDetailResponseData FlagCheckLogDetailResponseData

// NewFlagCheckLogDetailResponseData instantiates a new FlagCheckLogDetailResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlagCheckLogDetailResponseData(checkStatus string, createdAt time.Time, environmentId string, flagKey string, id string, reason string, updatedAt time.Time, value bool) *FlagCheckLogDetailResponseData {
	this := FlagCheckLogDetailResponseData{}
	this.CheckStatus = checkStatus
	this.CreatedAt = createdAt
	this.EnvironmentId = environmentId
	this.FlagKey = flagKey
	this.Id = id
	this.Reason = reason
	this.UpdatedAt = updatedAt
	this.Value = value
	return &this
}

// NewFlagCheckLogDetailResponseDataWithDefaults instantiates a new FlagCheckLogDetailResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlagCheckLogDetailResponseDataWithDefaults() *FlagCheckLogDetailResponseData {
	this := FlagCheckLogDetailResponseData{}
	return &this
}

// GetCheckStatus returns the CheckStatus field value
func (o *FlagCheckLogDetailResponseData) GetCheckStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckStatus
}

// GetCheckStatusOk returns a tuple with the CheckStatus field value
// and a boolean to check if the value has been set.
func (o *FlagCheckLogDetailResponseData) GetCheckStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckStatus, true
}

// SetCheckStatus sets field value
func (o *FlagCheckLogDetailResponseData) SetCheckStatus(v string) {
	o.CheckStatus = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *FlagCheckLogDetailResponseData) GetCompany() CompanyResponseData {
	if o == nil || IsNil(o.Company) {
		var ret CompanyResponseData
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlagCheckLogDetailResponseData) GetCompanyOk() (*CompanyResponseData, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *FlagCheckLogDetailResponseData) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given CompanyResponseData and assigns it to the Company field.
func (o *FlagCheckLogDetailResponseData) SetCompany(v CompanyResponseData) {
	o.Company = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlagCheckLogDetailResponseData) GetCompanyId() string {
	if o == nil || IsNil(o.CompanyId.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyId.Get()
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlagCheckLogDetailResponseData) GetCompanyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyId.Get(), o.CompanyId.IsSet()
}

// HasCompanyId returns a boolean if a field has been set.
func (o *FlagCheckLogDetailResponseData) HasCompanyId() bool {
	if o != nil && o.CompanyId.IsSet() {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given NullableString and assigns it to the CompanyId field.
func (o *FlagCheckLogDetailResponseData) SetCompanyId(v string) {
	o.CompanyId.Set(&v)
}

// SetCompanyIdNil sets the value for CompanyId to be an explicit nil
func (o *FlagCheckLogDetailResponseData) SetCompanyIdNil() {
	o.CompanyId.Set(nil)
}

// UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
func (o *FlagCheckLogDetailResponseData) UnsetCompanyId() {
	o.CompanyId.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *FlagCheckLogDetailResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FlagCheckLogDetailResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *FlagCheckLogDetailResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *FlagCheckLogDetailResponseData) GetEnvironment() EnvironmentResponseData {
	if o == nil || IsNil(o.Environment) {
		var ret EnvironmentResponseData
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlagCheckLogDetailResponseData) GetEnvironmentOk() (*EnvironmentResponseData, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *FlagCheckLogDetailResponseData) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given EnvironmentResponseData and assigns it to the Environment field.
func (o *FlagCheckLogDetailResponseData) SetEnvironment(v EnvironmentResponseData) {
	o.Environment = &v
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *FlagCheckLogDetailResponseData) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *FlagCheckLogDetailResponseData) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *FlagCheckLogDetailResponseData) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlagCheckLogDetailResponseData) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlagCheckLogDetailResponseData) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *FlagCheckLogDetailResponseData) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *FlagCheckLogDetailResponseData) SetError(v string) {
	o.Error.Set(&v)
}

// SetErrorNil sets the value for Error to be an explicit nil
func (o *FlagCheckLogDetailResponseData) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *FlagCheckLogDetailResponseData) UnsetError() {
	o.Error.Unset()
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *FlagCheckLogDetailResponseData) GetFlag() FlagResponseData {
	if o == nil || IsNil(o.Flag) {
		var ret FlagResponseData
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlagCheckLogDetailResponseData) GetFlagOk() (*FlagResponseData, bool) {
	if o == nil || IsNil(o.Flag) {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *FlagCheckLogDetailResponseData) HasFlag() bool {
	if o != nil && !IsNil(o.Flag) {
		return true
	}

	return false
}

// SetFlag gets a reference to the given FlagResponseData and assigns it to the Flag field.
func (o *FlagCheckLogDetailResponseData) SetFlag(v FlagResponseData) {
	o.Flag = &v
}

// GetFlagId returns the FlagId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlagCheckLogDetailResponseData) GetFlagId() string {
	if o == nil || IsNil(o.FlagId.Get()) {
		var ret string
		return ret
	}
	return *o.FlagId.Get()
}

// GetFlagIdOk returns a tuple with the FlagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlagCheckLogDetailResponseData) GetFlagIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlagId.Get(), o.FlagId.IsSet()
}

// HasFlagId returns a boolean if a field has been set.
func (o *FlagCheckLogDetailResponseData) HasFlagId() bool {
	if o != nil && o.FlagId.IsSet() {
		return true
	}

	return false
}

// SetFlagId gets a reference to the given NullableString and assigns it to the FlagId field.
func (o *FlagCheckLogDetailResponseData) SetFlagId(v string) {
	o.FlagId.Set(&v)
}

// SetFlagIdNil sets the value for FlagId to be an explicit nil
func (o *FlagCheckLogDetailResponseData) SetFlagIdNil() {
	o.FlagId.Set(nil)
}

// UnsetFlagId ensures that no value is present for FlagId, not even an explicit nil
func (o *FlagCheckLogDetailResponseData) UnsetFlagId() {
	o.FlagId.Unset()
}

// GetFlagKey returns the FlagKey field value
func (o *FlagCheckLogDetailResponseData) GetFlagKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlagKey
}

// GetFlagKeyOk returns a tuple with the FlagKey field value
// and a boolean to check if the value has been set.
func (o *FlagCheckLogDetailResponseData) GetFlagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlagKey, true
}

// SetFlagKey sets field value
func (o *FlagCheckLogDetailResponseData) SetFlagKey(v string) {
	o.FlagKey = v
}

// GetId returns the Id field value
func (o *FlagCheckLogDetailResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FlagCheckLogDetailResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FlagCheckLogDetailResponseData) SetId(v string) {
	o.Id = v
}

// GetReason returns the Reason field value
func (o *FlagCheckLogDetailResponseData) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *FlagCheckLogDetailResponseData) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *FlagCheckLogDetailResponseData) SetReason(v string) {
	o.Reason = v
}

// GetReqCompany returns the ReqCompany field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlagCheckLogDetailResponseData) GetReqCompany() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ReqCompany
}

// GetReqCompanyOk returns a tuple with the ReqCompany field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlagCheckLogDetailResponseData) GetReqCompanyOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ReqCompany) {
		return nil, false
	}
	return &o.ReqCompany, true
}

// HasReqCompany returns a boolean if a field has been set.
func (o *FlagCheckLogDetailResponseData) HasReqCompany() bool {
	if o != nil && !IsNil(o.ReqCompany) {
		return true
	}

	return false
}

// SetReqCompany gets a reference to the given map[string]string and assigns it to the ReqCompany field.
func (o *FlagCheckLogDetailResponseData) SetReqCompany(v map[string]string) {
	o.ReqCompany = v
}

// GetReqUser returns the ReqUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlagCheckLogDetailResponseData) GetReqUser() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ReqUser
}

// GetReqUserOk returns a tuple with the ReqUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlagCheckLogDetailResponseData) GetReqUserOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ReqUser) {
		return nil, false
	}
	return &o.ReqUser, true
}

// HasReqUser returns a boolean if a field has been set.
func (o *FlagCheckLogDetailResponseData) HasReqUser() bool {
	if o != nil && !IsNil(o.ReqUser) {
		return true
	}

	return false
}

// SetReqUser gets a reference to the given map[string]string and assigns it to the ReqUser field.
func (o *FlagCheckLogDetailResponseData) SetReqUser(v map[string]string) {
	o.ReqUser = v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *FlagCheckLogDetailResponseData) GetRule() RuleResponseData {
	if o == nil || IsNil(o.Rule) {
		var ret RuleResponseData
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlagCheckLogDetailResponseData) GetRuleOk() (*RuleResponseData, bool) {
	if o == nil || IsNil(o.Rule) {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *FlagCheckLogDetailResponseData) HasRule() bool {
	if o != nil && !IsNil(o.Rule) {
		return true
	}

	return false
}

// SetRule gets a reference to the given RuleResponseData and assigns it to the Rule field.
func (o *FlagCheckLogDetailResponseData) SetRule(v RuleResponseData) {
	o.Rule = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlagCheckLogDetailResponseData) GetRuleId() string {
	if o == nil || IsNil(o.RuleId.Get()) {
		var ret string
		return ret
	}
	return *o.RuleId.Get()
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlagCheckLogDetailResponseData) GetRuleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleId.Get(), o.RuleId.IsSet()
}

// HasRuleId returns a boolean if a field has been set.
func (o *FlagCheckLogDetailResponseData) HasRuleId() bool {
	if o != nil && o.RuleId.IsSet() {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given NullableString and assigns it to the RuleId field.
func (o *FlagCheckLogDetailResponseData) SetRuleId(v string) {
	o.RuleId.Set(&v)
}

// SetRuleIdNil sets the value for RuleId to be an explicit nil
func (o *FlagCheckLogDetailResponseData) SetRuleIdNil() {
	o.RuleId.Set(nil)
}

// UnsetRuleId ensures that no value is present for RuleId, not even an explicit nil
func (o *FlagCheckLogDetailResponseData) UnsetRuleId() {
	o.RuleId.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *FlagCheckLogDetailResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *FlagCheckLogDetailResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *FlagCheckLogDetailResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *FlagCheckLogDetailResponseData) GetUser() UserResponseData {
	if o == nil || IsNil(o.User) {
		var ret UserResponseData
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlagCheckLogDetailResponseData) GetUserOk() (*UserResponseData, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *FlagCheckLogDetailResponseData) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserResponseData and assigns it to the User field.
func (o *FlagCheckLogDetailResponseData) SetUser(v UserResponseData) {
	o.User = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlagCheckLogDetailResponseData) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlagCheckLogDetailResponseData) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *FlagCheckLogDetailResponseData) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *FlagCheckLogDetailResponseData) SetUserId(v string) {
	o.UserId.Set(&v)
}

// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *FlagCheckLogDetailResponseData) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *FlagCheckLogDetailResponseData) UnsetUserId() {
	o.UserId.Unset()
}

// GetValue returns the Value field value
func (o *FlagCheckLogDetailResponseData) GetValue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FlagCheckLogDetailResponseData) GetValueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FlagCheckLogDetailResponseData) SetValue(v bool) {
	o.Value = v
}

func (o FlagCheckLogDetailResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlagCheckLogDetailResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["check_status"] = o.CheckStatus
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if o.CompanyId.IsSet() {
		toSerialize["company_id"] = o.CompanyId.Get()
	}
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	toSerialize["environment_id"] = o.EnvironmentId
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if !IsNil(o.Flag) {
		toSerialize["flag"] = o.Flag
	}
	if o.FlagId.IsSet() {
		toSerialize["flag_id"] = o.FlagId.Get()
	}
	toSerialize["flag_key"] = o.FlagKey
	toSerialize["id"] = o.Id
	toSerialize["reason"] = o.Reason
	if o.ReqCompany != nil {
		toSerialize["req_company"] = o.ReqCompany
	}
	if o.ReqUser != nil {
		toSerialize["req_user"] = o.ReqUser
	}
	if !IsNil(o.Rule) {
		toSerialize["rule"] = o.Rule
	}
	if o.RuleId.IsSet() {
		toSerialize["rule_id"] = o.RuleId.Get()
	}
	toSerialize["updated_at"] = o.UpdatedAt
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *FlagCheckLogDetailResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"check_status",
		"created_at",
		"environment_id",
		"flag_key",
		"id",
		"reason",
		"updated_at",
		"value",
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

	varFlagCheckLogDetailResponseData := _FlagCheckLogDetailResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFlagCheckLogDetailResponseData)

	if err != nil {
		return err
	}

	*o = FlagCheckLogDetailResponseData(varFlagCheckLogDetailResponseData)

	return err
}

type NullableFlagCheckLogDetailResponseData struct {
	value *FlagCheckLogDetailResponseData
	isSet bool
}

func (v NullableFlagCheckLogDetailResponseData) Get() *FlagCheckLogDetailResponseData {
	return v.value
}

func (v *NullableFlagCheckLogDetailResponseData) Set(val *FlagCheckLogDetailResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableFlagCheckLogDetailResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableFlagCheckLogDetailResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlagCheckLogDetailResponseData(val *FlagCheckLogDetailResponseData) *NullableFlagCheckLogDetailResponseData {
	return &NullableFlagCheckLogDetailResponseData{value: val, isSet: true}
}

func (v NullableFlagCheckLogDetailResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlagCheckLogDetailResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
