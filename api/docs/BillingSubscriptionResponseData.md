# BillingSubscriptionResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiredAt** | Pointer to **NullableTime** |  | [optional] 
**ExternalId** | **string** |  | 
**Id** | **int32** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewBillingSubscriptionResponseData

`func NewBillingSubscriptionResponseData(externalId string, id int32, updatedAt time.Time, ) *BillingSubscriptionResponseData`

NewBillingSubscriptionResponseData instantiates a new BillingSubscriptionResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingSubscriptionResponseDataWithDefaults

`func NewBillingSubscriptionResponseDataWithDefaults() *BillingSubscriptionResponseData`

NewBillingSubscriptionResponseDataWithDefaults instantiates a new BillingSubscriptionResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiredAt

`func (o *BillingSubscriptionResponseData) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *BillingSubscriptionResponseData) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *BillingSubscriptionResponseData) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *BillingSubscriptionResponseData) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### SetExpiredAtNil

`func (o *BillingSubscriptionResponseData) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *BillingSubscriptionResponseData) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetExternalId

`func (o *BillingSubscriptionResponseData) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BillingSubscriptionResponseData) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BillingSubscriptionResponseData) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetId

`func (o *BillingSubscriptionResponseData) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingSubscriptionResponseData) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingSubscriptionResponseData) SetId(v int32)`

SetId sets Id field to given value.


### GetUpdatedAt

`func (o *BillingSubscriptionResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BillingSubscriptionResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BillingSubscriptionResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


