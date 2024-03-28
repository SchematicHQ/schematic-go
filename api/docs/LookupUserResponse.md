# LookupUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**UserDetailResponseData**](UserDetailResponseData.md) |  | 
**Params** | [**LookupUserParams**](LookupUserParams.md) |  | 

## Methods

### NewLookupUserResponse

`func NewLookupUserResponse(data UserDetailResponseData, params LookupUserParams, ) *LookupUserResponse`

NewLookupUserResponse instantiates a new LookupUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupUserResponseWithDefaults

`func NewLookupUserResponseWithDefaults() *LookupUserResponse`

NewLookupUserResponseWithDefaults instantiates a new LookupUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *LookupUserResponse) GetData() UserDetailResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LookupUserResponse) GetDataOk() (*UserDetailResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LookupUserResponse) SetData(v UserDetailResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *LookupUserResponse) GetParams() LookupUserParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *LookupUserResponse) GetParamsOk() (*LookupUserParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *LookupUserResponse) SetParams(v LookupUserParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


