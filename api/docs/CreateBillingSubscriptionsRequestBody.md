# CreateBillingSubscriptionsRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerExternalId** | **string** |  | 
**ExpiredAt** | **time.Time** |  | 
**ProductExternalIds** | **[]string** |  | 
**SubscriptionExternalId** | **string** |  | 

## Methods

### NewCreateBillingSubscriptionsRequestBody

`func NewCreateBillingSubscriptionsRequestBody(customerExternalId string, expiredAt time.Time, productExternalIds []string, subscriptionExternalId string, ) *CreateBillingSubscriptionsRequestBody`

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


### GetProductExternalIds

`func (o *CreateBillingSubscriptionsRequestBody) GetProductExternalIds() []string`

GetProductExternalIds returns the ProductExternalIds field if non-nil, zero value otherwise.

### GetProductExternalIdsOk

`func (o *CreateBillingSubscriptionsRequestBody) GetProductExternalIdsOk() (*[]string, bool)`

GetProductExternalIdsOk returns a tuple with the ProductExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductExternalIds

`func (o *CreateBillingSubscriptionsRequestBody) SetProductExternalIds(v []string)`

SetProductExternalIds sets ProductExternalIds field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


