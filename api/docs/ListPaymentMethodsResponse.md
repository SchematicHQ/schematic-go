# ListPaymentMethodsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BillingPaymentMethodResponseData**](BillingPaymentMethodResponseData.md) | The returned resources | 
**Params** | [**ListPaymentMethodsParams**](ListPaymentMethodsParams.md) |  | 

## Methods

### NewListPaymentMethodsResponse

`func NewListPaymentMethodsResponse(data []BillingPaymentMethodResponseData, params ListPaymentMethodsParams, ) *ListPaymentMethodsResponse`

NewListPaymentMethodsResponse instantiates a new ListPaymentMethodsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPaymentMethodsResponseWithDefaults

`func NewListPaymentMethodsResponseWithDefaults() *ListPaymentMethodsResponse`

NewListPaymentMethodsResponseWithDefaults instantiates a new ListPaymentMethodsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListPaymentMethodsResponse) GetData() []BillingPaymentMethodResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPaymentMethodsResponse) GetDataOk() (*[]BillingPaymentMethodResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPaymentMethodsResponse) SetData(v []BillingPaymentMethodResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListPaymentMethodsResponse) GetParams() ListPaymentMethodsParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListPaymentMethodsResponse) GetParamsOk() (*ListPaymentMethodsParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListPaymentMethodsResponse) SetParams(v ListPaymentMethodsParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


