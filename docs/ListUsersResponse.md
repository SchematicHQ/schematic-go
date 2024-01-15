# ListUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]UserResponseData**](UserResponseData.md) | The returned resources | 
**Params** | [**ListUsersParams**](ListUsersParams.md) |  | 

## Methods

### NewListUsersResponse

`func NewListUsersResponse(data []UserResponseData, params ListUsersParams, ) *ListUsersResponse`

NewListUsersResponse instantiates a new ListUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsersResponseWithDefaults

`func NewListUsersResponseWithDefaults() *ListUsersResponse`

NewListUsersResponseWithDefaults instantiates a new ListUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListUsersResponse) GetData() []UserResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListUsersResponse) GetDataOk() (*[]UserResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListUsersResponse) SetData(v []UserResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListUsersResponse) GetParams() ListUsersParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListUsersResponse) GetParamsOk() (*ListUsersParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListUsersResponse) SetParams(v ListUsersParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


