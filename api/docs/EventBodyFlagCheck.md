# EventBodyFlagCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** | Schematic company ID (starting with &#39;comp_&#39;) of the company evaluated, if any | [optional] 
**Error** | Pointer to **NullableString** | Report an error that occurred during the flag check | [optional] 
**FlagId** | Pointer to **NullableString** | Schematic flag ID (starting with &#39;flag_&#39;) for the flag matching the key, if any | [optional] 
**FlagKey** | **string** | The key of the flag being checked | 
**Reason** | **string** | The reason why the value was returned | 
**ReqCompany** | Pointer to **map[string]string** | Key-value pairs used to to identify company for which the flag was checked | [optional] 
**ReqUser** | Pointer to **map[string]string** | Key-value pairs used to to identify user for which the flag was checked | [optional] 
**RuleId** | Pointer to **NullableString** | Schematic rule ID (starting with &#39;rule_&#39;) of the rule that matched for the flag, if any | [optional] 
**UserId** | Pointer to **NullableString** | Schematic user ID (starting with &#39;user_&#39;) of the user evaluated, if any | [optional] 
**Value** | **bool** | The value of the flag for the given company and/or user | 

## Methods

### NewEventBodyFlagCheck

`func NewEventBodyFlagCheck(flagKey string, reason string, value bool, ) *EventBodyFlagCheck`

NewEventBodyFlagCheck instantiates a new EventBodyFlagCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventBodyFlagCheckWithDefaults

`func NewEventBodyFlagCheckWithDefaults() *EventBodyFlagCheck`

NewEventBodyFlagCheckWithDefaults instantiates a new EventBodyFlagCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *EventBodyFlagCheck) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *EventBodyFlagCheck) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *EventBodyFlagCheck) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *EventBodyFlagCheck) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *EventBodyFlagCheck) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *EventBodyFlagCheck) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetError

`func (o *EventBodyFlagCheck) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EventBodyFlagCheck) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EventBodyFlagCheck) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *EventBodyFlagCheck) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *EventBodyFlagCheck) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *EventBodyFlagCheck) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetFlagId

`func (o *EventBodyFlagCheck) GetFlagId() string`

GetFlagId returns the FlagId field if non-nil, zero value otherwise.

### GetFlagIdOk

`func (o *EventBodyFlagCheck) GetFlagIdOk() (*string, bool)`

GetFlagIdOk returns a tuple with the FlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagId

`func (o *EventBodyFlagCheck) SetFlagId(v string)`

SetFlagId sets FlagId field to given value.

### HasFlagId

`func (o *EventBodyFlagCheck) HasFlagId() bool`

HasFlagId returns a boolean if a field has been set.

### SetFlagIdNil

`func (o *EventBodyFlagCheck) SetFlagIdNil(b bool)`

 SetFlagIdNil sets the value for FlagId to be an explicit nil

### UnsetFlagId
`func (o *EventBodyFlagCheck) UnsetFlagId()`

UnsetFlagId ensures that no value is present for FlagId, not even an explicit nil
### GetFlagKey

`func (o *EventBodyFlagCheck) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *EventBodyFlagCheck) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *EventBodyFlagCheck) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.


### GetReason

`func (o *EventBodyFlagCheck) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EventBodyFlagCheck) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EventBodyFlagCheck) SetReason(v string)`

SetReason sets Reason field to given value.


### GetReqCompany

`func (o *EventBodyFlagCheck) GetReqCompany() map[string]string`

GetReqCompany returns the ReqCompany field if non-nil, zero value otherwise.

### GetReqCompanyOk

`func (o *EventBodyFlagCheck) GetReqCompanyOk() (*map[string]string, bool)`

GetReqCompanyOk returns a tuple with the ReqCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqCompany

`func (o *EventBodyFlagCheck) SetReqCompany(v map[string]string)`

SetReqCompany sets ReqCompany field to given value.

### HasReqCompany

`func (o *EventBodyFlagCheck) HasReqCompany() bool`

HasReqCompany returns a boolean if a field has been set.

### SetReqCompanyNil

`func (o *EventBodyFlagCheck) SetReqCompanyNil(b bool)`

 SetReqCompanyNil sets the value for ReqCompany to be an explicit nil

### UnsetReqCompany
`func (o *EventBodyFlagCheck) UnsetReqCompany()`

UnsetReqCompany ensures that no value is present for ReqCompany, not even an explicit nil
### GetReqUser

`func (o *EventBodyFlagCheck) GetReqUser() map[string]string`

GetReqUser returns the ReqUser field if non-nil, zero value otherwise.

### GetReqUserOk

`func (o *EventBodyFlagCheck) GetReqUserOk() (*map[string]string, bool)`

GetReqUserOk returns a tuple with the ReqUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqUser

`func (o *EventBodyFlagCheck) SetReqUser(v map[string]string)`

SetReqUser sets ReqUser field to given value.

### HasReqUser

`func (o *EventBodyFlagCheck) HasReqUser() bool`

HasReqUser returns a boolean if a field has been set.

### SetReqUserNil

`func (o *EventBodyFlagCheck) SetReqUserNil(b bool)`

 SetReqUserNil sets the value for ReqUser to be an explicit nil

### UnsetReqUser
`func (o *EventBodyFlagCheck) UnsetReqUser()`

UnsetReqUser ensures that no value is present for ReqUser, not even an explicit nil
### GetRuleId

`func (o *EventBodyFlagCheck) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *EventBodyFlagCheck) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *EventBodyFlagCheck) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *EventBodyFlagCheck) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### SetRuleIdNil

`func (o *EventBodyFlagCheck) SetRuleIdNil(b bool)`

 SetRuleIdNil sets the value for RuleId to be an explicit nil

### UnsetRuleId
`func (o *EventBodyFlagCheck) UnsetRuleId()`

UnsetRuleId ensures that no value is present for RuleId, not even an explicit nil
### GetUserId

`func (o *EventBodyFlagCheck) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EventBodyFlagCheck) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EventBodyFlagCheck) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EventBodyFlagCheck) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *EventBodyFlagCheck) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *EventBodyFlagCheck) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetValue

`func (o *EventBodyFlagCheck) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EventBodyFlagCheck) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EventBodyFlagCheck) SetValue(v bool)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


