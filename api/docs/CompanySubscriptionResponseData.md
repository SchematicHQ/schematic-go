# CompanySubscriptionResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerExternalId** | **string** |  | 
**ExpiredAt** | Pointer to **NullableTime** |  | [optional] 
**Interval** | **string** |  | 
**Products** | [**[]BillingProductResponseData**](BillingProductResponseData.md) |  | 
**SubscriptionExternalId** | **string** |  | 

## Methods

### NewCompanySubscriptionResponseData

`func NewCompanySubscriptionResponseData(customerExternalId string, interval string, products []BillingProductResponseData, subscriptionExternalId string, ) *CompanySubscriptionResponseData`

NewCompanySubscriptionResponseData instantiates a new CompanySubscriptionResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanySubscriptionResponseDataWithDefaults

`func NewCompanySubscriptionResponseDataWithDefaults() *CompanySubscriptionResponseData`

NewCompanySubscriptionResponseDataWithDefaults instantiates a new CompanySubscriptionResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerExternalId

`func (o *CompanySubscriptionResponseData) GetCustomerExternalId() string`

GetCustomerExternalId returns the CustomerExternalId field if non-nil, zero value otherwise.

### GetCustomerExternalIdOk

`func (o *CompanySubscriptionResponseData) GetCustomerExternalIdOk() (*string, bool)`

GetCustomerExternalIdOk returns a tuple with the CustomerExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerExternalId

`func (o *CompanySubscriptionResponseData) SetCustomerExternalId(v string)`

SetCustomerExternalId sets CustomerExternalId field to given value.


### GetExpiredAt

`func (o *CompanySubscriptionResponseData) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *CompanySubscriptionResponseData) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *CompanySubscriptionResponseData) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *CompanySubscriptionResponseData) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### SetExpiredAtNil

`func (o *CompanySubscriptionResponseData) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *CompanySubscriptionResponseData) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetInterval

`func (o *CompanySubscriptionResponseData) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CompanySubscriptionResponseData) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CompanySubscriptionResponseData) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetProducts

`func (o *CompanySubscriptionResponseData) GetProducts() []BillingProductResponseData`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *CompanySubscriptionResponseData) GetProductsOk() (*[]BillingProductResponseData, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *CompanySubscriptionResponseData) SetProducts(v []BillingProductResponseData)`

SetProducts sets Products field to given value.


### GetSubscriptionExternalId

`func (o *CompanySubscriptionResponseData) GetSubscriptionExternalId() string`

GetSubscriptionExternalId returns the SubscriptionExternalId field if non-nil, zero value otherwise.

### GetSubscriptionExternalIdOk

`func (o *CompanySubscriptionResponseData) GetSubscriptionExternalIdOk() (*string, bool)`

GetSubscriptionExternalIdOk returns a tuple with the SubscriptionExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionExternalId

`func (o *CompanySubscriptionResponseData) SetSubscriptionExternalId(v string)`

SetSubscriptionExternalId sets SubscriptionExternalId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


