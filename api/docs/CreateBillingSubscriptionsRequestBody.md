# CreateBillingSubscriptionsRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerExternalId** | **string** |  | 
**ExpiredAt** | **time.Time** |  | 
**Interval** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**ProductExternalIds** | [**[]BillingProductPricing**](BillingProductPricing.md) |  | 
**Status** | Pointer to **NullableString** |  | [optional] 
**SubscriptionExternalId** | **string** |  | 
**TotalPrice** | **int32** |  | 

## Methods

### NewCreateBillingSubscriptionsRequestBody

`func NewCreateBillingSubscriptionsRequestBody(customerExternalId string, expiredAt time.Time, productExternalIds []BillingProductPricing, status string, subscriptionExternalId string, totalPrice int32, ) *CreateBillingSubscriptionsRequestBody`

NewCreateBillingSubscriptionsRequestBody instantiates a new CreateBillingSubscriptionsRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBillingSubscriptionsRequestBodyWithDefaults

`func NewCreateBillingSubscriptionsRequestBodyWithDefaults() *CreateBillingSubscriptionsRequestBody`

NewCreateBillingSubscriptionsRequestBodyWithDefaults instantiates a new CreateBillingSubscriptionsRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerExternalId

`func (o *CreateBillingSubscriptionsRequestBody) GetCustomerExternalId() string`

GetCustomerExternalId returns the CustomerExternalId field if non-nil, zero value otherwise.

### GetCustomerExternalIdOk

`func (o *CreateBillingSubscriptionsRequestBody) GetCustomerExternalIdOk() (*string, bool)`

GetCustomerExternalIdOk returns a tuple with the CustomerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerExternalId

`func (o *CreateBillingSubscriptionsRequestBody) SetCustomerExternalId(v string)`

SetCustomerExternalId sets CustomerExternalId field to given value.


### GetExpiredAt

`func (o *CreateBillingSubscriptionsRequestBody) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *CreateBillingSubscriptionsRequestBody) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *CreateBillingSubscriptionsRequestBody) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.


### GetInterval

`func (o *CreateBillingSubscriptionsRequestBody) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CreateBillingSubscriptionsRequestBody) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CreateBillingSubscriptionsRequestBody) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CreateBillingSubscriptionsRequestBody) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *CreateBillingSubscriptionsRequestBody) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *CreateBillingSubscriptionsRequestBody) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetMetadata

`func (o *CreateBillingSubscriptionsRequestBody) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateBillingSubscriptionsRequestBody) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateBillingSubscriptionsRequestBody) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateBillingSubscriptionsRequestBody) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreateBillingSubscriptionsRequestBody) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreateBillingSubscriptionsRequestBody) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetProductExternalIds

`func (o *CreateBillingSubscriptionsRequestBody) GetProductExternalIds() []BillingProductPricing`

GetProductExternalIds returns the ProductExternalIds field if non-nil, zero value otherwise.

### GetProductExternalIdsOk

`func (o *CreateBillingSubscriptionsRequestBody) GetProductExternalIdsOk() (*[]BillingProductPricing, bool)`

GetProductExternalIdsOk returns a tuple with the ProductExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductExternalIds

`func (o *CreateBillingSubscriptionsRequestBody) SetProductExternalIds(v []BillingProductPricing)`

SetProductExternalIds sets ProductExternalIds field to given value.


### GetStatus

`func (o *CreateBillingSubscriptionsRequestBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateBillingSubscriptionsRequestBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateBillingSubscriptionsRequestBody) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateBillingSubscriptionsRequestBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CreateBillingSubscriptionsRequestBody) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CreateBillingSubscriptionsRequestBody) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSubscriptionExternalId

`func (o *CreateBillingSubscriptionsRequestBody) GetSubscriptionExternalId() string`

GetSubscriptionExternalId returns the SubscriptionExternalId field if non-nil, zero value otherwise.

### GetSubscriptionExternalIdOk

`func (o *CreateBillingSubscriptionsRequestBody) GetSubscriptionExternalIdOk() (*string, bool)`

GetSubscriptionExternalIdOk returns a tuple with the SubscriptionExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExternalId

`func (o *CreateBillingSubscriptionsRequestBody) SetSubscriptionExternalId(v string)`

SetSubscriptionExternalId sets SubscriptionExternalId field to given value.


### GetTotalPrice

`func (o *CreateBillingSubscriptionsRequestBody) GetTotalPrice() int32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *CreateBillingSubscriptionsRequestBody) GetTotalPriceOk() (*int32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *CreateBillingSubscriptionsRequestBody) SetTotalPrice(v int32)`

SetTotalPrice sets TotalPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


