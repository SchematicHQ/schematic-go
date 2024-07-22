# ListCustomersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BillingCustomerWithSubscriptionsResponseData**](BillingCustomerWithSubscriptionsResponseData.md) | The returned resources | 
**Params** | [**ListCustomersParams**](ListCustomersParams.md) |  | 

## Methods

### NewListCustomersResponse

`func NewListCustomersResponse(data []BillingCustomerWithSubscriptionsResponseData, params ListCustomersParams, ) *ListCustomersResponse`

NewListCustomersResponse instantiates a new ListCustomersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomersResponseWithDefaults

`func NewListCustomersResponseWithDefaults() *ListCustomersResponse`

NewListCustomersResponseWithDefaults instantiates a new ListCustomersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListCustomersResponse) GetData() []BillingCustomerWithSubscriptionsResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCustomersResponse) GetDataOk() (*[]BillingCustomerWithSubscriptionsResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCustomersResponse) SetData(v []BillingCustomerWithSubscriptionsResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListCustomersResponse) GetParams() ListCustomersParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListCustomersResponse) GetParamsOk() (*ListCustomersParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListCustomersResponse) SetParams(v ListCustomersParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


