# ListProductPricesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BillingPriceResponseData**](BillingPriceResponseData.md) | The returned resources | 
**Params** | [**ListProductPricesParams**](ListProductPricesParams.md) |  | 

## Methods

### NewListProductPricesResponse

`func NewListProductPricesResponse(data []BillingPriceResponseData, params ListProductPricesParams, ) *ListProductPricesResponse`

NewListProductPricesResponse instantiates a new ListProductPricesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProductPricesResponseWithDefaults

`func NewListProductPricesResponseWithDefaults() *ListProductPricesResponse`

NewListProductPricesResponseWithDefaults instantiates a new ListProductPricesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListProductPricesResponse) GetData() []BillingPriceResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListProductPricesResponse) GetDataOk() (*[]BillingPriceResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListProductPricesResponse) SetData(v []BillingPriceResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListProductPricesResponse) GetParams() ListProductPricesParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListProductPricesResponse) GetParamsOk() (*ListProductPricesParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListProductPricesResponse) SetParams(v ListProductPricesParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


