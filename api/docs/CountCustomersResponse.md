# CountCustomersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**CountResponse**](CountResponse.md) |  | 
**Params** | [**CountCustomersParams**](CountCustomersParams.md) |  | 

## Methods

### NewCountCustomersResponse

`func NewCountCustomersResponse(data CountResponse, params CountCustomersParams, ) *CountCustomersResponse`

NewCountCustomersResponse instantiates a new CountCustomersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountCustomersResponseWithDefaults

`func NewCountCustomersResponseWithDefaults() *CountCustomersResponse`

NewCountCustomersResponseWithDefaults instantiates a new CountCustomersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CountCustomersResponse) GetData() CountResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CountCustomersResponse) GetDataOk() (*CountResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CountCustomersResponse) SetData(v CountResponse)`

SetData sets Data field to given value.


### GetParams

`func (o *CountCustomersResponse) GetParams() CountCustomersParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CountCustomersResponse) GetParamsOk() (*CountCustomersParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CountCustomersResponse) SetParams(v CountCustomersParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


