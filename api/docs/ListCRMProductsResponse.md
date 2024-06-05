# ListCRMProductsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CRMProductResponseData**](CRMProductResponseData.md) | The returned resources | 
**Params** | [**ListCRMProductsParams**](ListCRMProductsParams.md) |  | 

## Methods

### NewListCRMProductsResponse

`func NewListCRMProductsResponse(data []CRMProductResponseData, params ListCRMProductsParams, ) *ListCRMProductsResponse`

NewListCRMProductsResponse instantiates a new ListCRMProductsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCRMProductsResponseWithDefaults

`func NewListCRMProductsResponseWithDefaults() *ListCRMProductsResponse`

NewListCRMProductsResponseWithDefaults instantiates a new ListCRMProductsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListCRMProductsResponse) GetData() []CRMProductResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCRMProductsResponse) GetDataOk() (*[]CRMProductResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCRMProductsResponse) SetData(v []CRMProductResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListCRMProductsResponse) GetParams() ListCRMProductsParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListCRMProductsResponse) GetParamsOk() (*ListCRMProductsParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListCRMProductsResponse) SetParams(v ListCRMProductsParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


