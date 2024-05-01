# ListProductsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BillingProductResponseData**](BillingProductResponseData.md) | The returned resources | 
**Params** | [**ListProductsParams**](ListProductsParams.md) |  | 

## Methods

### NewListProductsResponse

`func NewListProductsResponse(data []BillingProductResponseData, params ListProductsParams, ) *ListProductsResponse`

NewListProductsResponse instantiates a new ListProductsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProductsResponseWithDefaults

`func NewListProductsResponseWithDefaults() *ListProductsResponse`

NewListProductsResponseWithDefaults instantiates a new ListProductsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListProductsResponse) GetData() []BillingProductResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListProductsResponse) GetDataOk() (*[]BillingProductResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListProductsResponse) SetData(v []BillingProductResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListProductsResponse) GetParams() ListProductsParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListProductsResponse) GetParamsOk() (*ListProductsParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListProductsResponse) SetParams(v ListProductsParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


