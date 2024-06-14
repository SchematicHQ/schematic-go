# CreateCrmLineItemRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TermMonth** | Pointer to **NullableInt32** |  | [optional] 
**Amount** | **string** |  | 
**DiscountPercentage** | Pointer to **NullableString** |  | [optional] 
**Interval** | **string** |  | 
**LineItemExternalId** | **string** |  | 
**ProductExternalId** | **string** |  | 
**Quantity** | **int32** |  | 
**TotalDiscount** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateCrmLineItemRequestBody

`func NewCreateCrmLineItemRequestBody(amount string, interval string, lineItemExternalId string, productExternalId string, quantity int32, ) *CreateCrmLineItemRequestBody`

NewCreateCrmLineItemRequestBody instantiates a new CreateCrmLineItemRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCrmLineItemRequestBodyWithDefaults

`func NewCreateCrmLineItemRequestBodyWithDefaults() *CreateCrmLineItemRequestBody`

NewCreateCrmLineItemRequestBodyWithDefaults instantiates a new CreateCrmLineItemRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTermMonth

`func (o *CreateCrmLineItemRequestBody) GetTermMonth() int32`

GetTermMonth returns the TermMonth field if non-nil, zero value otherwise.

### GetTermMonthOk

`func (o *CreateCrmLineItemRequestBody) GetTermMonthOk() (*int32, bool)`

GetTermMonthOk returns a tuple with the TermMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermMonth

`func (o *CreateCrmLineItemRequestBody) SetTermMonth(v int32)`

SetTermMonth sets TermMonth field to given value.

### HasTermMonth

`func (o *CreateCrmLineItemRequestBody) HasTermMonth() bool`

HasTermMonth returns a boolean if a field has been set.

### SetTermMonthNil

`func (o *CreateCrmLineItemRequestBody) SetTermMonthNil(b bool)`

 SetTermMonthNil sets the value for TermMonth to be an explicit nil

### UnsetTermMonth
`func (o *CreateCrmLineItemRequestBody) UnsetTermMonth()`

UnsetTermMonth ensures that no value is present for TermMonth, not even an explicit nil
### GetAmount

`func (o *CreateCrmLineItemRequestBody) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateCrmLineItemRequestBody) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateCrmLineItemRequestBody) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDiscountPercentage

`func (o *CreateCrmLineItemRequestBody) GetDiscountPercentage() string`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *CreateCrmLineItemRequestBody) GetDiscountPercentageOk() (*string, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *CreateCrmLineItemRequestBody) SetDiscountPercentage(v string)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *CreateCrmLineItemRequestBody) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### SetDiscountPercentageNil

`func (o *CreateCrmLineItemRequestBody) SetDiscountPercentageNil(b bool)`

 SetDiscountPercentageNil sets the value for DiscountPercentage to be an explicit nil

### UnsetDiscountPercentage
`func (o *CreateCrmLineItemRequestBody) UnsetDiscountPercentage()`

UnsetDiscountPercentage ensures that no value is present for DiscountPercentage, not even an explicit nil
### GetInterval

`func (o *CreateCrmLineItemRequestBody) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CreateCrmLineItemRequestBody) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CreateCrmLineItemRequestBody) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetLineItemExternalId

`func (o *CreateCrmLineItemRequestBody) GetLineItemExternalId() string`

GetLineItemExternalId returns the LineItemExternalId field if non-nil, zero value otherwise.

### GetLineItemExternalIdOk

`func (o *CreateCrmLineItemRequestBody) GetLineItemExternalIdOk() (*string, bool)`

GetLineItemExternalIdOk returns a tuple with the LineItemExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemExternalId

`func (o *CreateCrmLineItemRequestBody) SetLineItemExternalId(v string)`

SetLineItemExternalId sets LineItemExternalId field to given value.


### GetProductExternalId

`func (o *CreateCrmLineItemRequestBody) GetProductExternalId() string`

GetProductExternalId returns the ProductExternalId field if non-nil, zero value otherwise.

### GetProductExternalIdOk

`func (o *CreateCrmLineItemRequestBody) GetProductExternalIdOk() (*string, bool)`

GetProductExternalIdOk returns a tuple with the ProductExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductExternalId

`func (o *CreateCrmLineItemRequestBody) SetProductExternalId(v string)`

SetProductExternalId sets ProductExternalId field to given value.


### GetQuantity

`func (o *CreateCrmLineItemRequestBody) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateCrmLineItemRequestBody) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateCrmLineItemRequestBody) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetTotalDiscount

`func (o *CreateCrmLineItemRequestBody) GetTotalDiscount() string`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *CreateCrmLineItemRequestBody) GetTotalDiscountOk() (*string, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *CreateCrmLineItemRequestBody) SetTotalDiscount(v string)`

SetTotalDiscount sets TotalDiscount field to given value.

### HasTotalDiscount

`func (o *CreateCrmLineItemRequestBody) HasTotalDiscount() bool`

HasTotalDiscount returns a boolean if a field has been set.

### SetTotalDiscountNil

`func (o *CreateCrmLineItemRequestBody) SetTotalDiscountNil(b bool)`

 SetTotalDiscountNil sets the value for TotalDiscount to be an explicit nil

### UnsetTotalDiscount
`func (o *CreateCrmLineItemRequestBody) UnsetTotalDiscount()`

UnsetTotalDiscount ensures that no value is present for TotalDiscount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


