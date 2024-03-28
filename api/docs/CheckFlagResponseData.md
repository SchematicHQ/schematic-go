# CheckFlagResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **NullableString** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 
**Reason** | **string** |  | 
**RuleId** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**Value** | **bool** |  | 

## Methods

### NewCheckFlagResponseData

`func NewCheckFlagResponseData(reason string, value bool, ) *CheckFlagResponseData`

NewCheckFlagResponseData instantiates a new CheckFlagResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckFlagResponseDataWithDefaults

`func NewCheckFlagResponseDataWithDefaults() *CheckFlagResponseData`

NewCheckFlagResponseDataWithDefaults instantiates a new CheckFlagResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *CheckFlagResponseData) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CheckFlagResponseData) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CheckFlagResponseData) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *CheckFlagResponseData) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *CheckFlagResponseData) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *CheckFlagResponseData) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetError

`func (o *CheckFlagResponseData) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CheckFlagResponseData) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CheckFlagResponseData) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CheckFlagResponseData) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *CheckFlagResponseData) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *CheckFlagResponseData) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetReason

`func (o *CheckFlagResponseData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CheckFlagResponseData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CheckFlagResponseData) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRuleId

`func (o *CheckFlagResponseData) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *CheckFlagResponseData) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *CheckFlagResponseData) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *CheckFlagResponseData) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### SetRuleIdNil

`func (o *CheckFlagResponseData) SetRuleIdNil(b bool)`

 SetRuleIdNil sets the value for RuleId to be an explicit nil

### UnsetRuleId
`func (o *CheckFlagResponseData) UnsetRuleId()`

UnsetRuleId ensures that no value is present for RuleId, not even an explicit nil
### GetUserId

`func (o *CheckFlagResponseData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CheckFlagResponseData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CheckFlagResponseData) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CheckFlagResponseData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *CheckFlagResponseData) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *CheckFlagResponseData) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetValue

`func (o *CheckFlagResponseData) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CheckFlagResponseData) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CheckFlagResponseData) SetValue(v bool)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


