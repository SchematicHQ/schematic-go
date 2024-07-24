# ListCrmProductsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CrmProductResponseData**](CrmProductResponseData.md) | The returned resources | 
**Params** | [**ListCrmProductsParams**](ListCrmProductsParams.md) |  | 

## Methods

### NewListCrmProductsResponse

`func NewListCrmProductsResponse(data []CrmProductResponseData, params ListCrmProductsParams, ) *ListCrmProductsResponse`

NewListCrmProductsResponse instantiates a new ListCrmProductsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCrmProductsResponseWithDefaults

`func NewListCrmProductsResponseWithDefaults() *ListCrmProductsResponse`

NewListCrmProductsResponseWithDefaults instantiates a new ListCrmProductsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListCrmProductsResponse) GetData() []CrmProductResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCrmProductsResponse) GetDataOk() (*[]CrmProductResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCrmProductsResponse) SetData(v []CrmProductResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListCrmProductsResponse) GetParams() ListCrmProductsParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListCrmProductsResponse) GetParamsOk() (*ListCrmProductsParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListCrmProductsResponse) SetParams(v ListCrmProductsParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


