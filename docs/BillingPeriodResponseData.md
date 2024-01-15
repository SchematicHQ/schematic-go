# BillingPeriodResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**EndedAt** | Pointer to **NullableTime** |  | [optional] 
**EnvironmentId** | **string** |  | 
**Id** | **string** |  | 
**StartedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewBillingPeriodResponseData

`func NewBillingPeriodResponseData(companyId string, createdAt time.Time, environmentId string, id string, startedAt time.Time, updatedAt time.Time, ) *BillingPeriodResponseData`

NewBillingPeriodResponseData instantiates a new BillingPeriodResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingPeriodResponseDataWithDefaults

`func NewBillingPeriodResponseDataWithDefaults() *BillingPeriodResponseData`

NewBillingPeriodResponseDataWithDefaults instantiates a new BillingPeriodResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *BillingPeriodResponseData) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BillingPeriodResponseData) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BillingPeriodResponseData) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetCreatedAt

`func (o *BillingPeriodResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingPeriodResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingPeriodResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEndedAt

`func (o *BillingPeriodResponseData) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *BillingPeriodResponseData) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *BillingPeriodResponseData) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *BillingPeriodResponseData) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *BillingPeriodResponseData) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *BillingPeriodResponseData) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetEnvironmentId

`func (o *BillingPeriodResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *BillingPeriodResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *BillingPeriodResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetId

`func (o *BillingPeriodResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingPeriodResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingPeriodResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetStartedAt

`func (o *BillingPeriodResponseData) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BillingPeriodResponseData) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BillingPeriodResponseData) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetUpdatedAt

`func (o *BillingPeriodResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingPeriodResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingPeriodResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


