# FlagCheckLogDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckStatus** | **string** |  | 
**Company** | Pointer to [**CompanyResponseData**](CompanyResponseData.md) |  | [optional] 
**CompanyId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Environment** | Pointer to [**EnvironmentResponseData**](EnvironmentResponseData.md) |  | [optional] 
**EnvironmentId** | **string** |  | 
**Error** | Pointer to **NullableString** |  | [optional] 
**Flag** | Pointer to [**FlagResponseData**](FlagResponseData.md) |  | [optional] 
**FlagId** | Pointer to **NullableString** |  | [optional] 
**FlagKey** | **string** |  | 
**Id** | **string** |  | 
**Reason** | **string** |  | 
**ReqCompany** | Pointer to **map[string]string** |  | [optional] 
**ReqUser** | Pointer to **map[string]string** |  | [optional] 
**Rule** | Pointer to [**RuleResponseData**](RuleResponseData.md) |  | [optional] 
**RuleId** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 
**User** | Pointer to [**UserResponseData**](UserResponseData.md) |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**Value** | **bool** |  | 

## Methods

### NewFlagCheckLogDetailResponseData

`func NewFlagCheckLogDetailResponseData(checkStatus string, createdAt time.Time, environmentId string, flagKey string, id string, reason string, updatedAt time.Time, value bool, ) *FlagCheckLogDetailResponseData`

NewFlagCheckLogDetailResponseData instantiates a new FlagCheckLogDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagCheckLogDetailResponseDataWithDefaults

`func NewFlagCheckLogDetailResponseDataWithDefaults() *FlagCheckLogDetailResponseData`

NewFlagCheckLogDetailResponseDataWithDefaults instantiates a new FlagCheckLogDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckStatus

`func (o *FlagCheckLogDetailResponseData) GetCheckStatus() string`

GetCheckStatus returns the CheckStatus field if non-nil, zero value otherwise.

### GetCheckStatusOk

`func (o *FlagCheckLogDetailResponseData) GetCheckStatusOk() (*string, bool)`

GetCheckStatusOk returns a tuple with the CheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckStatus

`func (o *FlagCheckLogDetailResponseData) SetCheckStatus(v string)`

SetCheckStatus sets CheckStatus field to given value.


### GetCompany

`func (o *FlagCheckLogDetailResponseData) GetCompany() CompanyResponseData`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *FlagCheckLogDetailResponseData) GetCompanyOk() (*CompanyResponseData, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *FlagCheckLogDetailResponseData) SetCompany(v CompanyResponseData)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *FlagCheckLogDetailResponseData) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyId

`func (o *FlagCheckLogDetailResponseData) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *FlagCheckLogDetailResponseData) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *FlagCheckLogDetailResponseData) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *FlagCheckLogDetailResponseData) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *FlagCheckLogDetailResponseData) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *FlagCheckLogDetailResponseData) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCreatedAt

`func (o *FlagCheckLogDetailResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FlagCheckLogDetailResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FlagCheckLogDetailResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnvironment

`func (o *FlagCheckLogDetailResponseData) GetEnvironment() EnvironmentResponseData`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FlagCheckLogDetailResponseData) GetEnvironmentOk() (*EnvironmentResponseData, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FlagCheckLogDetailResponseData) SetEnvironment(v EnvironmentResponseData)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FlagCheckLogDetailResponseData) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *FlagCheckLogDetailResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FlagCheckLogDetailResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FlagCheckLogDetailResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetError

`func (o *FlagCheckLogDetailResponseData) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FlagCheckLogDetailResponseData) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FlagCheckLogDetailResponseData) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *FlagCheckLogDetailResponseData) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *FlagCheckLogDetailResponseData) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *FlagCheckLogDetailResponseData) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetFlag

`func (o *FlagCheckLogDetailResponseData) GetFlag() FlagResponseData`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *FlagCheckLogDetailResponseData) GetFlagOk() (*FlagResponseData, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *FlagCheckLogDetailResponseData) SetFlag(v FlagResponseData)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *FlagCheckLogDetailResponseData) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetFlagId

`func (o *FlagCheckLogDetailResponseData) GetFlagId() string`

GetFlagId returns the FlagId field if non-nil, zero value otherwise.

### GetFlagIdOk

`func (o *FlagCheckLogDetailResponseData) GetFlagIdOk() (*string, bool)`

GetFlagIdOk returns a tuple with the FlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagId

`func (o *FlagCheckLogDetailResponseData) SetFlagId(v string)`

SetFlagId sets FlagId field to given value.

### HasFlagId

`func (o *FlagCheckLogDetailResponseData) HasFlagId() bool`

HasFlagId returns a boolean if a field has been set.

### SetFlagIdNil

`func (o *FlagCheckLogDetailResponseData) SetFlagIdNil(b bool)`

 SetFlagIdNil sets the value for FlagId to be an explicit nil

### UnsetFlagId
`func (o *FlagCheckLogDetailResponseData) UnsetFlagId()`

UnsetFlagId ensures that no value is present for FlagId, not even an explicit nil
### GetFlagKey

`func (o *FlagCheckLogDetailResponseData) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *FlagCheckLogDetailResponseData) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *FlagCheckLogDetailResponseData) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.


### GetId

`func (o *FlagCheckLogDetailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlagCheckLogDetailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlagCheckLogDetailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetReason

`func (o *FlagCheckLogDetailResponseData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *FlagCheckLogDetailResponseData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *FlagCheckLogDetailResponseData) SetReason(v string)`

SetReason sets Reason field to given value.


### GetReqCompany

`func (o *FlagCheckLogDetailResponseData) GetReqCompany() map[string]string`

GetReqCompany returns the ReqCompany field if non-nil, zero value otherwise.

### GetReqCompanyOk

`func (o *FlagCheckLogDetailResponseData) GetReqCompanyOk() (*map[string]string, bool)`

GetReqCompanyOk returns a tuple with the ReqCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqCompany

`func (o *FlagCheckLogDetailResponseData) SetReqCompany(v map[string]string)`

SetReqCompany sets ReqCompany field to given value.

### HasReqCompany

`func (o *FlagCheckLogDetailResponseData) HasReqCompany() bool`

HasReqCompany returns a boolean if a field has been set.

### SetReqCompanyNil

`func (o *FlagCheckLogDetailResponseData) SetReqCompanyNil(b bool)`

 SetReqCompanyNil sets the value for ReqCompany to be an explicit nil

### UnsetReqCompany
`func (o *FlagCheckLogDetailResponseData) UnsetReqCompany()`

UnsetReqCompany ensures that no value is present for ReqCompany, not even an explicit nil
### GetReqUser

`func (o *FlagCheckLogDetailResponseData) GetReqUser() map[string]string`

GetReqUser returns the ReqUser field if non-nil, zero value otherwise.

### GetReqUserOk

`func (o *FlagCheckLogDetailResponseData) GetReqUserOk() (*map[string]string, bool)`

GetReqUserOk returns a tuple with the ReqUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqUser

`func (o *FlagCheckLogDetailResponseData) SetReqUser(v map[string]string)`

SetReqUser sets ReqUser field to given value.

### HasReqUser

`func (o *FlagCheckLogDetailResponseData) HasReqUser() bool`

HasReqUser returns a boolean if a field has been set.

### SetReqUserNil

`func (o *FlagCheckLogDetailResponseData) SetReqUserNil(b bool)`

 SetReqUserNil sets the value for ReqUser to be an explicit nil

### UnsetReqUser
`func (o *FlagCheckLogDetailResponseData) UnsetReqUser()`

UnsetReqUser ensures that no value is present for ReqUser, not even an explicit nil
### GetRule

`func (o *FlagCheckLogDetailResponseData) GetRule() RuleResponseData`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *FlagCheckLogDetailResponseData) GetRuleOk() (*RuleResponseData, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *FlagCheckLogDetailResponseData) SetRule(v RuleResponseData)`

SetRule sets Rule field to given value.

### HasRule

`func (o *FlagCheckLogDetailResponseData) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetRuleId

`func (o *FlagCheckLogDetailResponseData) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *FlagCheckLogDetailResponseData) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *FlagCheckLogDetailResponseData) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *FlagCheckLogDetailResponseData) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### SetRuleIdNil

`func (o *FlagCheckLogDetailResponseData) SetRuleIdNil(b bool)`

 SetRuleIdNil sets the value for RuleId to be an explicit nil

### UnsetRuleId
`func (o *FlagCheckLogDetailResponseData) UnsetRuleId()`

UnsetRuleId ensures that no value is present for RuleId, not even an explicit nil
### GetUpdatedAt

`func (o *FlagCheckLogDetailResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FlagCheckLogDetailResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FlagCheckLogDetailResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUser

`func (o *FlagCheckLogDetailResponseData) GetUser() UserResponseData`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FlagCheckLogDetailResponseData) GetUserOk() (*UserResponseData, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FlagCheckLogDetailResponseData) SetUser(v UserResponseData)`

SetUser sets User field to given value.

### HasUser

`func (o *FlagCheckLogDetailResponseData) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *FlagCheckLogDetailResponseData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FlagCheckLogDetailResponseData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FlagCheckLogDetailResponseData) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FlagCheckLogDetailResponseData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *FlagCheckLogDetailResponseData) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *FlagCheckLogDetailResponseData) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetValue

`func (o *FlagCheckLogDetailResponseData) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FlagCheckLogDetailResponseData) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FlagCheckLogDetailResponseData) SetValue(v bool)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


