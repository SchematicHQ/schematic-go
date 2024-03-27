# ApiKeyRequestListResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | **string** |  | 
**EndedAt** | Pointer to **NullableTime** |  | [optional] 
**EnvironmentId** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**Method** | **string** |  | 
**ReqBody** | Pointer to **NullableString** |  | [optional] 
**RespCode** | Pointer to **NullableInt32** |  | [optional] 
**StartedAt** | **time.Time** |  | 
**Url** | **string** |  | 

## Methods

### NewApiKeyRequestListResponseData

`func NewApiKeyRequestListResponseData(apiKeyId string, id string, method string, startedAt time.Time, url string, ) *ApiKeyRequestListResponseData`

NewApiKeyRequestListResponseData instantiates a new ApiKeyRequestListResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyRequestListResponseDataWithDefaults

`func NewApiKeyRequestListResponseDataWithDefaults() *ApiKeyRequestListResponseData`

NewApiKeyRequestListResponseDataWithDefaults instantiates a new ApiKeyRequestListResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *ApiKeyRequestListResponseData) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *ApiKeyRequestListResponseData) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *ApiKeyRequestListResponseData) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.


### GetEndedAt

`func (o *ApiKeyRequestListResponseData) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *ApiKeyRequestListResponseData) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *ApiKeyRequestListResponseData) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *ApiKeyRequestListResponseData) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *ApiKeyRequestListResponseData) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *ApiKeyRequestListResponseData) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetEnvironmentId

`func (o *ApiKeyRequestListResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ApiKeyRequestListResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ApiKeyRequestListResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ApiKeyRequestListResponseData) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### SetEnvironmentIdNil

`func (o *ApiKeyRequestListResponseData) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *ApiKeyRequestListResponseData) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
### GetId

`func (o *ApiKeyRequestListResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKeyRequestListResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiKeyRequestListResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetMethod

`func (o *ApiKeyRequestListResponseData) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ApiKeyRequestListResponseData) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ApiKeyRequestListResponseData) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetReqBody

`func (o *ApiKeyRequestListResponseData) GetReqBody() string`

GetReqBody returns the ReqBody field if non-nil, zero value otherwise.

### GetReqBodyOk

`func (o *ApiKeyRequestListResponseData) GetReqBodyOk() (*string, bool)`

GetReqBodyOk returns a tuple with the ReqBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqBody

`func (o *ApiKeyRequestListResponseData) SetReqBody(v string)`

SetReqBody sets ReqBody field to given value.

### HasReqBody

`func (o *ApiKeyRequestListResponseData) HasReqBody() bool`

HasReqBody returns a boolean if a field has been set.

### SetReqBodyNil

`func (o *ApiKeyRequestListResponseData) SetReqBodyNil(b bool)`

 SetReqBodyNil sets the value for ReqBody to be an explicit nil

### UnsetReqBody
`func (o *ApiKeyRequestListResponseData) UnsetReqBody()`

UnsetReqBody ensures that no value is present for ReqBody, not even an explicit nil
### GetRespCode

`func (o *ApiKeyRequestListResponseData) GetRespCode() int32`

GetRespCode returns the RespCode field if non-nil, zero value otherwise.

### GetRespCodeOk

`func (o *ApiKeyRequestListResponseData) GetRespCodeOk() (*int32, bool)`

GetRespCodeOk returns a tuple with the RespCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespCode

`func (o *ApiKeyRequestListResponseData) SetRespCode(v int32)`

SetRespCode sets RespCode field to given value.

### HasRespCode

`func (o *ApiKeyRequestListResponseData) HasRespCode() bool`

HasRespCode returns a boolean if a field has been set.

### SetRespCodeNil

`func (o *ApiKeyRequestListResponseData) SetRespCodeNil(b bool)`

 SetRespCodeNil sets the value for RespCode to be an explicit nil

### UnsetRespCode
`func (o *ApiKeyRequestListResponseData) UnsetRespCode()`

UnsetRespCode ensures that no value is present for RespCode, not even an explicit nil
### GetStartedAt

`func (o *ApiKeyRequestListResponseData) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ApiKeyRequestListResponseData) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ApiKeyRequestListResponseData) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetUrl

`func (o *ApiKeyRequestListResponseData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApiKeyRequestListResponseData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApiKeyRequestListResponseData) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


