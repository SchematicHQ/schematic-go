# CheckFlagRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to **map[string]string** |  | [optional] 
**User** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCheckFlagRequestBody

`func NewCheckFlagRequestBody() *CheckFlagRequestBody`

NewCheckFlagRequestBody instantiates a new CheckFlagRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckFlagRequestBodyWithDefaults

`func NewCheckFlagRequestBodyWithDefaults() *CheckFlagRequestBody`

NewCheckFlagRequestBodyWithDefaults instantiates a new CheckFlagRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *CheckFlagRequestBody) GetCompany() map[string]string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CheckFlagRequestBody) GetCompanyOk() (*map[string]string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CheckFlagRequestBody) SetCompany(v map[string]string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CheckFlagRequestBody) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *CheckFlagRequestBody) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *CheckFlagRequestBody) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetUser

`func (o *CheckFlagRequestBody) GetUser() map[string]string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CheckFlagRequestBody) GetUserOk() (*map[string]string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CheckFlagRequestBody) SetUser(v map[string]string)`

SetUser sets User field to given value.

### HasUser

`func (o *CheckFlagRequestBody) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *CheckFlagRequestBody) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *CheckFlagRequestBody) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


