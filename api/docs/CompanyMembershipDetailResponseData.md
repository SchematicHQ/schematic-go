# CompanyMembershipDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to [**CompanyResponseData**](CompanyResponseData.md) |  | [optional] 
**CompanyId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 

## Methods

### NewCompanyMembershipDetailResponseData

`func NewCompanyMembershipDetailResponseData(companyId string, createdAt time.Time, id string, updatedAt time.Time, userId string, ) *CompanyMembershipDetailResponseData`

NewCompanyMembershipDetailResponseData instantiates a new CompanyMembershipDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyMembershipDetailResponseDataWithDefaults

`func NewCompanyMembershipDetailResponseDataWithDefaults() *CompanyMembershipDetailResponseData`

NewCompanyMembershipDetailResponseDataWithDefaults instantiates a new CompanyMembershipDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *CompanyMembershipDetailResponseData) GetCompany() CompanyResponseData`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CompanyMembershipDetailResponseData) GetCompanyOk() (*CompanyResponseData, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CompanyMembershipDetailResponseData) SetCompany(v CompanyResponseData)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CompanyMembershipDetailResponseData) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyId

`func (o *CompanyMembershipDetailResponseData) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *CompanyMembershipDetailResponseData) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *CompanyMembershipDetailResponseData) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetCreatedAt

`func (o *CompanyMembershipDetailResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CompanyMembershipDetailResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CompanyMembershipDetailResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *CompanyMembershipDetailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyMembershipDetailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyMembershipDetailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetUpdatedAt

`func (o *CompanyMembershipDetailResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CompanyMembershipDetailResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CompanyMembershipDetailResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *CompanyMembershipDetailResponseData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CompanyMembershipDetailResponseData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CompanyMembershipDetailResponseData) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


