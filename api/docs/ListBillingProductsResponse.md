# ListBillingProductsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BillingProductDetailResponseData**](BillingProductDetailResponseData.md) | The returned resources | 
**Params** | [**ListBillingProductsParams**](ListBillingProductsParams.md) |  | 

## Methods

### NewListBillingProductsResponse

`func NewListBillingProductsResponse(data []BillingProductDetailResponseData, params ListBillingProductsParams, ) *ListBillingProductsResponse`

NewListBillingProductsResponse instantiates a new ListBillingProductsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBillingProductsResponseWithDefaults

`func NewListBillingProductsResponseWithDefaults() *ListBillingProductsResponse`

NewListBillingProductsResponseWithDefaults instantiates a new ListBillingProductsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListBillingProductsResponse) GetData() []BillingProductDetailResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListBillingProductsResponse) GetDataOk() (*[]BillingProductDetailResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListBillingProductsResponse) SetData(v []BillingProductDetailResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListBillingProductsResponse) GetParams() ListBillingProductsParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListBillingProductsResponse) GetParamsOk() (*ListBillingProductsParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListBillingProductsResponse) SetParams(v ListBillingProductsParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


