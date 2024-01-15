# UpsertBillingPeriodRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **string** |  | 
**EndedAt** | Pointer to **NullableTime** |  | [optional] 
**StartedAt** | **time.Time** |  | 

## Methods

### NewUpsertBillingPeriodRequestBody

`func NewUpsertBillingPeriodRequestBody(companyId string, startedAt time.Time, ) *UpsertBillingPeriodRequestBody`

NewUpsertBillingPeriodRequestBody instantiates a new UpsertBillingPeriodRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertBillingPeriodRequestBodyWithDefaults

`func NewUpsertBillingPeriodRequestBodyWithDefaults() *UpsertBillingPeriodRequestBody`

NewUpsertBillingPeriodRequestBodyWithDefaults instantiates a new UpsertBillingPeriodRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *UpsertBillingPeriodRequestBody) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *UpsertBillingPeriodRequestBody) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *UpsertBillingPeriodRequestBody) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetEndedAt

`func (o *UpsertBillingPeriodRequestBody) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *UpsertBillingPeriodRequestBody) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *UpsertBillingPeriodRequestBody) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *UpsertBillingPeriodRequestBody) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *UpsertBillingPeriodRequestBody) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *UpsertBillingPeriodRequestBody) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetStartedAt

`func (o *UpsertBillingPeriodRequestBody) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *UpsertBillingPeriodRequestBody) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *UpsertBillingPeriodRequestBody) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


