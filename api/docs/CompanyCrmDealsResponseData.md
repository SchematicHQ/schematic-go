# CompanyCrmDealsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DealArr** | Pointer to **NullableFloat32** |  | [optional] 
**DealExternalId** | **string** |  | 
**DealMrr** | Pointer to **NullableFloat32** |  | [optional] 
**DealName** | Pointer to **NullableString** |  | [optional] 
**LineItems** | [**[]CrmDealLineItem**](CrmDealLineItem.md) |  | 

## Methods

### NewCompanyCrmDealsResponseData

`func NewCompanyCrmDealsResponseData(dealExternalId string, lineItems []CrmDealLineItem, ) *CompanyCrmDealsResponseData`

NewCompanyCrmDealsResponseData instantiates a new CompanyCrmDealsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyCrmDealsResponseDataWithDefaults

`func NewCompanyCrmDealsResponseDataWithDefaults() *CompanyCrmDealsResponseData`

NewCompanyCrmDealsResponseDataWithDefaults instantiates a new CompanyCrmDealsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDealArr

`func (o *CompanyCrmDealsResponseData) GetDealArr() float32`

GetDealArr returns the DealArr field if non-nil, zero value otherwise.

### GetDealArrOk

`func (o *CompanyCrmDealsResponseData) GetDealArrOk() (*float32, bool)`

GetDealArrOk returns a tuple with the DealArr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealArr

`func (o *CompanyCrmDealsResponseData) SetDealArr(v float32)`

SetDealArr sets DealArr field to given value.

### HasDealArr

`func (o *CompanyCrmDealsResponseData) HasDealArr() bool`

HasDealArr returns a boolean if a field has been set.

### SetDealArrNil

`func (o *CompanyCrmDealsResponseData) SetDealArrNil(b bool)`

 SetDealArrNil sets the value for DealArr to be an explicit nil

### UnsetDealArr
`func (o *CompanyCrmDealsResponseData) UnsetDealArr()`

UnsetDealArr ensures that no value is present for DealArr, not even an explicit nil
### GetDealExternalId

`func (o *CompanyCrmDealsResponseData) GetDealExternalId() string`

GetDealExternalId returns the DealExternalId field if non-nil, zero value otherwise.

### GetDealExternalIdOk

`func (o *CompanyCrmDealsResponseData) GetDealExternalIdOk() (*string, bool)`

GetDealExternalIdOk returns a tuple with the DealExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealExternalId

`func (o *CompanyCrmDealsResponseData) SetDealExternalId(v string)`

SetDealExternalId sets DealExternalId field to given value.


### GetDealMrr

`func (o *CompanyCrmDealsResponseData) GetDealMrr() float32`

GetDealMrr returns the DealMrr field if non-nil, zero value otherwise.

### GetDealMrrOk

`func (o *CompanyCrmDealsResponseData) GetDealMrrOk() (*float32, bool)`

GetDealMrrOk returns a tuple with the DealMrr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealMrr

`func (o *CompanyCrmDealsResponseData) SetDealMrr(v float32)`

SetDealMrr sets DealMrr field to given value.

### HasDealMrr

`func (o *CompanyCrmDealsResponseData) HasDealMrr() bool`

HasDealMrr returns a boolean if a field has been set.

### SetDealMrrNil

`func (o *CompanyCrmDealsResponseData) SetDealMrrNil(b bool)`

 SetDealMrrNil sets the value for DealMrr to be an explicit nil

### UnsetDealMrr
`func (o *CompanyCrmDealsResponseData) UnsetDealMrr()`

UnsetDealMrr ensures that no value is present for DealMrr, not even an explicit nil
### GetDealName

`func (o *CompanyCrmDealsResponseData) GetDealName() string`

GetDealName returns the DealName field if non-nil, zero value otherwise.

### GetDealNameOk

`func (o *CompanyCrmDealsResponseData) GetDealNameOk() (*string, bool)`

GetDealNameOk returns a tuple with the DealName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealName

`func (o *CompanyCrmDealsResponseData) SetDealName(v string)`

SetDealName sets DealName field to given value.

### HasDealName

`func (o *CompanyCrmDealsResponseData) HasDealName() bool`

HasDealName returns a boolean if a field has been set.

### SetDealNameNil

`func (o *CompanyCrmDealsResponseData) SetDealNameNil(b bool)`

 SetDealNameNil sets the value for DealName to be an explicit nil

### UnsetDealName
`func (o *CompanyCrmDealsResponseData) UnsetDealName()`

UnsetDealName ensures that no value is present for DealName, not even an explicit nil
### GetLineItems

`func (o *CompanyCrmDealsResponseData) GetLineItems() []CrmDealLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CompanyCrmDealsResponseData) GetLineItemsOk() (*[]CrmDealLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CompanyCrmDealsResponseData) SetLineItems(v []CrmDealLineItem)`

SetLineItems sets LineItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


