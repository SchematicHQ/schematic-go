# UpsertBillingProductRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingProductId** | **string** |  | 
**MonthlyPriceId** | Pointer to **NullableString** |  | [optional] 
**YearlyPriceId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpsertBillingProductRequestBody

`func NewUpsertBillingProductRequestBody(billingProductId string, ) *UpsertBillingProductRequestBody`

NewUpsertBillingProductRequestBody instantiates a new UpsertBillingProductRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertBillingProductRequestBodyWithDefaults

`func NewUpsertBillingProductRequestBodyWithDefaults() *UpsertBillingProductRequestBody`

NewUpsertBillingProductRequestBodyWithDefaults instantiates a new UpsertBillingProductRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingProductId

`func (o *UpsertBillingProductRequestBody) GetBillingProductId() string`

GetBillingProductId returns the BillingProductId field if non-nil, zero value otherwise.

### GetBillingProductIdOk

`func (o *UpsertBillingProductRequestBody) GetBillingProductIdOk() (*string, bool)`

GetBillingProductIdOk returns a tuple with the BillingProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProductId

`func (o *UpsertBillingProductRequestBody) SetBillingProductId(v string)`

SetBillingProductId sets BillingProductId field to given value.


### GetMonthlyPriceId

`func (o *UpsertBillingProductRequestBody) GetMonthlyPriceId() string`

GetMonthlyPriceId returns the MonthlyPriceId field if non-nil, zero value otherwise.

### GetMonthlyPriceIdOk

`func (o *UpsertBillingProductRequestBody) GetMonthlyPriceIdOk() (*string, bool)`

GetMonthlyPriceIdOk returns a tuple with the MonthlyPriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPriceId

`func (o *UpsertBillingProductRequestBody) SetMonthlyPriceId(v string)`

SetMonthlyPriceId sets MonthlyPriceId field to given value.

### HasMonthlyPriceId

`func (o *UpsertBillingProductRequestBody) HasMonthlyPriceId() bool`

HasMonthlyPriceId returns a boolean if a field has been set.

### SetMonthlyPriceIdNil

`func (o *UpsertBillingProductRequestBody) SetMonthlyPriceIdNil(b bool)`

 SetMonthlyPriceIdNil sets the value for MonthlyPriceId to be an explicit nil

### UnsetMonthlyPriceId
`func (o *UpsertBillingProductRequestBody) UnsetMonthlyPriceId()`

UnsetMonthlyPriceId ensures that no value is present for MonthlyPriceId, not even an explicit nil
### GetYearlyPriceId

`func (o *UpsertBillingProductRequestBody) GetYearlyPriceId() string`

GetYearlyPriceId returns the YearlyPriceId field if non-nil, zero value otherwise.

### GetYearlyPriceIdOk

`func (o *UpsertBillingProductRequestBody) GetYearlyPriceIdOk() (*string, bool)`

GetYearlyPriceIdOk returns a tuple with the YearlyPriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyPriceId

`func (o *UpsertBillingProductRequestBody) SetYearlyPriceId(v string)`

SetYearlyPriceId sets YearlyPriceId field to given value.

### HasYearlyPriceId

`func (o *UpsertBillingProductRequestBody) HasYearlyPriceId() bool`

HasYearlyPriceId returns a boolean if a field has been set.

### SetYearlyPriceIdNil

`func (o *UpsertBillingProductRequestBody) SetYearlyPriceIdNil(b bool)`

 SetYearlyPriceIdNil sets the value for YearlyPriceId to be an explicit nil

### UnsetYearlyPriceId
`func (o *UpsertBillingProductRequestBody) UnsetYearlyPriceId()`

UnsetYearlyPriceId ensures that no value is present for YearlyPriceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


