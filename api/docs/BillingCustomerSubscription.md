# BillingCustomerSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiredAt** | Pointer to **NullableTime** |  | [optional] 
**TotalPrice** | **int32** |  | 

## Methods

### NewBillingCustomerSubscription

`func NewBillingCustomerSubscription(totalPrice int32, ) *BillingCustomerSubscription`

NewBillingCustomerSubscription instantiates a new BillingCustomerSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCustomerSubscriptionWithDefaults

`func NewBillingCustomerSubscriptionWithDefaults() *BillingCustomerSubscription`

NewBillingCustomerSubscriptionWithDefaults instantiates a new BillingCustomerSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiredAt

`func (o *BillingCustomerSubscription) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *BillingCustomerSubscription) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *BillingCustomerSubscription) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *BillingCustomerSubscription) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### SetExpiredAtNil

`func (o *BillingCustomerSubscription) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *BillingCustomerSubscription) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetTotalPrice

`func (o *BillingCustomerSubscription) GetTotalPrice() int32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *BillingCustomerSubscription) GetTotalPriceOk() (*int32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *BillingCustomerSubscription) SetTotalPrice(v int32)`

SetTotalPrice sets TotalPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


