# EventBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to [**EventBodyIdentifyCompany**](EventBodyIdentifyCompany.md) |  | [optional] 
**Event** | **string** | The name of the type of track event | 
**Traits** | Pointer to **map[string]interface{}** | A map of trait names to trait values | [optional] 
**User** | Pointer to **map[string]string** | Key-value pairs to identify user associated with track event | [optional] 
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
**Keys** | **map[string]string** | Key-value pairs to identify the user | 
**Name** | Pointer to **string** | The display name of the user being identified; required only if it is a new user | [optional] 

## Methods

### NewEventBody

`func NewEventBody(event string, flagKey string, reason string, value bool, keys map[string]string, ) *EventBody`

NewEventBody instantiates a new EventBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventBodyWithDefaults

`func NewEventBodyWithDefaults() *EventBody`

NewEventBodyWithDefaults instantiates a new EventBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *EventBody) GetCompany() EventBodyIdentifyCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *EventBody) GetCompanyOk() (*EventBodyIdentifyCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *EventBody) SetCompany(v EventBodyIdentifyCompany)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *EventBody) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetEvent

`func (o *EventBody) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventBody) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventBody) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetTraits

`func (o *EventBody) GetTraits() map[string]interface{}`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *EventBody) GetTraitsOk() (*map[string]interface{}, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *EventBody) SetTraits(v map[string]interface{})`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *EventBody) HasTraits() bool`

HasTraits returns a boolean if a field has been set.

### GetUser

`func (o *EventBody) GetUser() map[string]string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EventBody) GetUserOk() (*map[string]string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EventBody) SetUser(v map[string]string)`

SetUser sets User field to given value.

### HasUser

`func (o *EventBody) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCompanyId

`func (o *EventBody) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *EventBody) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *EventBody) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *EventBody) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *EventBody) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *EventBody) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetError

`func (o *EventBody) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EventBody) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EventBody) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *EventBody) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *EventBody) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *EventBody) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetFlagId

`func (o *EventBody) GetFlagId() string`

GetFlagId returns the FlagId field if non-nil, zero value otherwise.

### GetFlagIdOk

`func (o *EventBody) GetFlagIdOk() (*string, bool)`

GetFlagIdOk returns a tuple with the FlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagId

`func (o *EventBody) SetFlagId(v string)`

SetFlagId sets FlagId field to given value.

### HasFlagId

`func (o *EventBody) HasFlagId() bool`

HasFlagId returns a boolean if a field has been set.

### SetFlagIdNil

`func (o *EventBody) SetFlagIdNil(b bool)`

 SetFlagIdNil sets the value for FlagId to be an explicit nil

### UnsetFlagId
`func (o *EventBody) UnsetFlagId()`

UnsetFlagId ensures that no value is present for FlagId, not even an explicit nil
### GetFlagKey

`func (o *EventBody) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *EventBody) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *EventBody) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.


### GetReason

`func (o *EventBody) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EventBody) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EventBody) SetReason(v string)`

SetReason sets Reason field to given value.


### GetReqCompany

`func (o *EventBody) GetReqCompany() map[string]string`

GetReqCompany returns the ReqCompany field if non-nil, zero value otherwise.

### GetReqCompanyOk

`func (o *EventBody) GetReqCompanyOk() (*map[string]string, bool)`

GetReqCompanyOk returns a tuple with the ReqCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqCompany

`func (o *EventBody) SetReqCompany(v map[string]string)`

SetReqCompany sets ReqCompany field to given value.

### HasReqCompany

`func (o *EventBody) HasReqCompany() bool`

HasReqCompany returns a boolean if a field has been set.

### SetReqCompanyNil

`func (o *EventBody) SetReqCompanyNil(b bool)`

 SetReqCompanyNil sets the value for ReqCompany to be an explicit nil

### UnsetReqCompany
`func (o *EventBody) UnsetReqCompany()`

UnsetReqCompany ensures that no value is present for ReqCompany, not even an explicit nil
### GetReqUser

`func (o *EventBody) GetReqUser() map[string]string`

GetReqUser returns the ReqUser field if non-nil, zero value otherwise.

### GetReqUserOk

`func (o *EventBody) GetReqUserOk() (*map[string]string, bool)`

GetReqUserOk returns a tuple with the ReqUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqUser

`func (o *EventBody) SetReqUser(v map[string]string)`

SetReqUser sets ReqUser field to given value.

### HasReqUser

`func (o *EventBody) HasReqUser() bool`

HasReqUser returns a boolean if a field has been set.

### SetReqUserNil

`func (o *EventBody) SetReqUserNil(b bool)`

 SetReqUserNil sets the value for ReqUser to be an explicit nil

### UnsetReqUser
`func (o *EventBody) UnsetReqUser()`

UnsetReqUser ensures that no value is present for ReqUser, not even an explicit nil
### GetRuleId

`func (o *EventBody) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *EventBody) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *EventBody) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *EventBody) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### SetRuleIdNil

`func (o *EventBody) SetRuleIdNil(b bool)`

 SetRuleIdNil sets the value for RuleId to be an explicit nil

### UnsetRuleId
`func (o *EventBody) UnsetRuleId()`

UnsetRuleId ensures that no value is present for RuleId, not even an explicit nil
### GetUserId

`func (o *EventBody) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EventBody) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EventBody) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EventBody) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *EventBody) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *EventBody) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetValue

`func (o *EventBody) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EventBody) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EventBody) SetValue(v bool)`

SetValue sets Value field to given value.


### GetKeys

`func (o *EventBody) GetKeys() map[string]string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *EventBody) GetKeysOk() (*map[string]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *EventBody) SetKeys(v map[string]string)`

SetKeys sets Keys field to given value.


### GetName

`func (o *EventBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventBody) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


