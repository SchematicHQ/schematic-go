# FlagCheckLogResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckStatus** | **string** |  | 
**CompanyId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**EnvironmentId** | **string** |  | 
**Error** | Pointer to **NullableString** |  | [optional] 
**FlagId** | Pointer to **NullableString** |  | [optional] 
**FlagKey** | **string** |  | 
**Id** | **string** |  | 
**Reason** | **string** |  | 
**ReqCompany** | Pointer to **map[string]string** |  | [optional] 
**ReqUser** | Pointer to **map[string]string** |  | [optional] 
**RuleId** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 
**UserId** | Pointer to **NullableString** |  | [optional] 
**Value** | **bool** |  | 

## Methods

### NewFlagCheckLogResponseData

`func NewFlagCheckLogResponseData(checkStatus string, createdAt time.Time, environmentId string, flagKey string, id string, reason string, updatedAt time.Time, value bool, ) *FlagCheckLogResponseData`

NewFlagCheckLogResponseData instantiates a new FlagCheckLogResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagCheckLogResponseDataWithDefaults

`func NewFlagCheckLogResponseDataWithDefaults() *FlagCheckLogResponseData`

NewFlagCheckLogResponseDataWithDefaults instantiates a new FlagCheckLogResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckStatus

`func (o *FlagCheckLogResponseData) GetCheckStatus() string`

GetCheckStatus returns the CheckStatus field if non-nil, zero value otherwise.

### GetCheckStatusOk

`func (o *FlagCheckLogResponseData) GetCheckStatusOk() (*string, bool)`

GetCheckStatusOk returns a tuple with the CheckStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckStatus

`func (o *FlagCheckLogResponseData) SetCheckStatus(v string)`

SetCheckStatus sets CheckStatus field to given value.


### GetCompanyId

`func (o *FlagCheckLogResponseData) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *FlagCheckLogResponseData) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *FlagCheckLogResponseData) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *FlagCheckLogResponseData) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *FlagCheckLogResponseData) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *FlagCheckLogResponseData) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetCreatedAt

`func (o *FlagCheckLogResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FlagCheckLogResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FlagCheckLogResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnvironmentId

`func (o *FlagCheckLogResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FlagCheckLogResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FlagCheckLogResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetError

`func (o *FlagCheckLogResponseData) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FlagCheckLogResponseData) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FlagCheckLogResponseData) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *FlagCheckLogResponseData) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *FlagCheckLogResponseData) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *FlagCheckLogResponseData) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetFlagId

`func (o *FlagCheckLogResponseData) GetFlagId() string`

GetFlagId returns the FlagId field if non-nil, zero value otherwise.

### GetFlagIdOk

`func (o *FlagCheckLogResponseData) GetFlagIdOk() (*string, bool)`

GetFlagIdOk returns a tuple with the FlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagId

`func (o *FlagCheckLogResponseData) SetFlagId(v string)`

SetFlagId sets FlagId field to given value.

### HasFlagId

`func (o *FlagCheckLogResponseData) HasFlagId() bool`

HasFlagId returns a boolean if a field has been set.

### SetFlagIdNil

`func (o *FlagCheckLogResponseData) SetFlagIdNil(b bool)`

 SetFlagIdNil sets the value for FlagId to be an explicit nil

### UnsetFlagId
`func (o *FlagCheckLogResponseData) UnsetFlagId()`

UnsetFlagId ensures that no value is present for FlagId, not even an explicit nil
### GetFlagKey

`func (o *FlagCheckLogResponseData) GetFlagKey() string`

GetFlagKey returns the FlagKey field if non-nil, zero value otherwise.

### GetFlagKeyOk

`func (o *FlagCheckLogResponseData) GetFlagKeyOk() (*string, bool)`

GetFlagKeyOk returns a tuple with the FlagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagKey

`func (o *FlagCheckLogResponseData) SetFlagKey(v string)`

SetFlagKey sets FlagKey field to given value.


### GetId

`func (o *FlagCheckLogResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlagCheckLogResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlagCheckLogResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetReason

`func (o *FlagCheckLogResponseData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *FlagCheckLogResponseData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *FlagCheckLogResponseData) SetReason(v string)`

SetReason sets Reason field to given value.


### GetReqCompany

`func (o *FlagCheckLogResponseData) GetReqCompany() map[string]string`

GetReqCompany returns the ReqCompany field if non-nil, zero value otherwise.

### GetReqCompanyOk

`func (o *FlagCheckLogResponseData) GetReqCompanyOk() (*map[string]string, bool)`

GetReqCompanyOk returns a tuple with the ReqCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqCompany

`func (o *FlagCheckLogResponseData) SetReqCompany(v map[string]string)`

SetReqCompany sets ReqCompany field to given value.

### HasReqCompany

`func (o *FlagCheckLogResponseData) HasReqCompany() bool`

HasReqCompany returns a boolean if a field has been set.

### SetReqCompanyNil

`func (o *FlagCheckLogResponseData) SetReqCompanyNil(b bool)`

 SetReqCompanyNil sets the value for ReqCompany to be an explicit nil

### UnsetReqCompany
`func (o *FlagCheckLogResponseData) UnsetReqCompany()`

UnsetReqCompany ensures that no value is present for ReqCompany, not even an explicit nil
### GetReqUser

`func (o *FlagCheckLogResponseData) GetReqUser() map[string]string`

GetReqUser returns the ReqUser field if non-nil, zero value otherwise.

### GetReqUserOk

`func (o *FlagCheckLogResponseData) GetReqUserOk() (*map[string]string, bool)`

GetReqUserOk returns a tuple with the ReqUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqUser

`func (o *FlagCheckLogResponseData) SetReqUser(v map[string]string)`

SetReqUser sets ReqUser field to given value.

### HasReqUser

`func (o *FlagCheckLogResponseData) HasReqUser() bool`

HasReqUser returns a boolean if a field has been set.

### SetReqUserNil

`func (o *FlagCheckLogResponseData) SetReqUserNil(b bool)`

 SetReqUserNil sets the value for ReqUser to be an explicit nil

### UnsetReqUser
`func (o *FlagCheckLogResponseData) UnsetReqUser()`

UnsetReqUser ensures that no value is present for ReqUser, not even an explicit nil
### GetRuleId

`func (o *FlagCheckLogResponseData) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *FlagCheckLogResponseData) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *FlagCheckLogResponseData) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *FlagCheckLogResponseData) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### SetRuleIdNil

`func (o *FlagCheckLogResponseData) SetRuleIdNil(b bool)`

 SetRuleIdNil sets the value for RuleId to be an explicit nil

### UnsetRuleId
`func (o *FlagCheckLogResponseData) UnsetRuleId()`

UnsetRuleId ensures that no value is present for RuleId, not even an explicit nil
### GetUpdatedAt

`func (o *FlagCheckLogResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FlagCheckLogResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FlagCheckLogResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *FlagCheckLogResponseData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FlagCheckLogResponseData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FlagCheckLogResponseData) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FlagCheckLogResponseData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *FlagCheckLogResponseData) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *FlagCheckLogResponseData) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetValue

`func (o *FlagCheckLogResponseData) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FlagCheckLogResponseData) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FlagCheckLogResponseData) SetValue(v bool)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


