# ListInvoicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]InvoiceResponseData**](InvoiceResponseData.md) | The returned resources | 
**Params** | [**ListInvoicesParams**](ListInvoicesParams.md) |  | 

## Methods

### NewListInvoicesResponse

`func NewListInvoicesResponse(data []InvoiceResponseData, params ListInvoicesParams, ) *ListInvoicesResponse`

NewListInvoicesResponse instantiates a new ListInvoicesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInvoicesResponseWithDefaults

`func NewListInvoicesResponseWithDefaults() *ListInvoicesResponse`

NewListInvoicesResponseWithDefaults instantiates a new ListInvoicesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListInvoicesResponse) GetData() []InvoiceResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListInvoicesResponse) GetDataOk() (*[]InvoiceResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListInvoicesResponse) SetData(v []InvoiceResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListInvoicesResponse) GetParams() ListInvoicesParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListInvoicesResponse) GetParamsOk() (*ListInvoicesParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListInvoicesResponse) SetParams(v ListInvoicesParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


