# ApiKeyRequestResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | **string** |  | 
**EndedAt** | Pointer to **NullableTime** |  | [optional] 
**EnvironmentId** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**Method** | **string** |  | 
**ReqBody** | Pointer to **NullableString** |  | [optional] 
**RequestType** | Pointer to **NullableString** |  | [optional] 
**ResourceId** | Pointer to **NullableInt32** |  | [optional] 
**ResourceIdString** | Pointer to **NullableString** |  | [optional] 
**ResourceName** | Pointer to **NullableString** |  | [optional] 
**ResourceType** | Pointer to **NullableString** |  | [optional] 
**RespBody** | Pointer to **NullableString** |  | [optional] 
**RespCode** | Pointer to **NullableInt32** |  | [optional] 
**SecondaryResource** | Pointer to **NullableString** |  | [optional] 
**StartedAt** | **time.Time** |  | 
**Url** | **string** |  | 
**UserAgent** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**UserName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewApiKeyRequestResponseData

`func NewApiKeyRequestResponseData(apiKeyId string, id string, method string, startedAt time.Time, url string, ) *ApiKeyRequestResponseData`

NewApiKeyRequestResponseData instantiates a new ApiKeyRequestResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyRequestResponseDataWithDefaults

`func NewApiKeyRequestResponseDataWithDefaults() *ApiKeyRequestResponseData`

NewApiKeyRequestResponseDataWithDefaults instantiates a new ApiKeyRequestResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *ApiKeyRequestResponseData) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *ApiKeyRequestResponseData) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *ApiKeyRequestResponseData) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.


### GetEndedAt

`func (o *ApiKeyRequestResponseData) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *ApiKeyRequestResponseData) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *ApiKeyRequestResponseData) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *ApiKeyRequestResponseData) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *ApiKeyRequestResponseData) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *ApiKeyRequestResponseData) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetEnvironmentId

`func (o *ApiKeyRequestResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ApiKeyRequestResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ApiKeyRequestResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ApiKeyRequestResponseData) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### SetEnvironmentIdNil

`func (o *ApiKeyRequestResponseData) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *ApiKeyRequestResponseData) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
### GetId

`func (o *ApiKeyRequestResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKeyRequestResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiKeyRequestResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetMethod

`func (o *ApiKeyRequestResponseData) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ApiKeyRequestResponseData) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ApiKeyRequestResponseData) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetReqBody

`func (o *ApiKeyRequestResponseData) GetReqBody() string`

GetReqBody returns the ReqBody field if non-nil, zero value otherwise.

### GetReqBodyOk

`func (o *ApiKeyRequestResponseData) GetReqBodyOk() (*string, bool)`

GetReqBodyOk returns a tuple with the ReqBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqBody

`func (o *ApiKeyRequestResponseData) SetReqBody(v string)`

SetReqBody sets ReqBody field to given value.

### HasReqBody

`func (o *ApiKeyRequestResponseData) HasReqBody() bool`

HasReqBody returns a boolean if a field has been set.

### SetReqBodyNil

`func (o *ApiKeyRequestResponseData) SetReqBodyNil(b bool)`

 SetReqBodyNil sets the value for ReqBody to be an explicit nil

### UnsetReqBody
`func (o *ApiKeyRequestResponseData) UnsetReqBody()`

UnsetReqBody ensures that no value is present for ReqBody, not even an explicit nil
### GetRequestType

`func (o *ApiKeyRequestResponseData) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *ApiKeyRequestResponseData) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *ApiKeyRequestResponseData) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *ApiKeyRequestResponseData) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### SetRequestTypeNil

`func (o *ApiKeyRequestResponseData) SetRequestTypeNil(b bool)`

 SetRequestTypeNil sets the value for RequestType to be an explicit nil

### UnsetRequestType
`func (o *ApiKeyRequestResponseData) UnsetRequestType()`

UnsetRequestType ensures that no value is present for RequestType, not even an explicit nil
### GetResourceId

`func (o *ApiKeyRequestResponseData) GetResourceId() int32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ApiKeyRequestResponseData) GetResourceIdOk() (*int32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ApiKeyRequestResponseData) SetResourceId(v int32)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *ApiKeyRequestResponseData) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *ApiKeyRequestResponseData) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *ApiKeyRequestResponseData) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetResourceIdString

`func (o *ApiKeyRequestResponseData) GetResourceIdString() string`

GetResourceIdString returns the ResourceIdString field if non-nil, zero value otherwise.

### GetResourceIdStringOk

`func (o *ApiKeyRequestResponseData) GetResourceIdStringOk() (*string, bool)`

GetResourceIdStringOk returns a tuple with the ResourceIdString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIdString

`func (o *ApiKeyRequestResponseData) SetResourceIdString(v string)`

SetResourceIdString sets ResourceIdString field to given value.

### HasResourceIdString

`func (o *ApiKeyRequestResponseData) HasResourceIdString() bool`

HasResourceIdString returns a boolean if a field has been set.

### SetResourceIdStringNil

`func (o *ApiKeyRequestResponseData) SetResourceIdStringNil(b bool)`

 SetResourceIdStringNil sets the value for ResourceIdString to be an explicit nil

### UnsetResourceIdString
`func (o *ApiKeyRequestResponseData) UnsetResourceIdString()`

UnsetResourceIdString ensures that no value is present for ResourceIdString, not even an explicit nil
### GetResourceName

`func (o *ApiKeyRequestResponseData) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ApiKeyRequestResponseData) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ApiKeyRequestResponseData) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *ApiKeyRequestResponseData) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *ApiKeyRequestResponseData) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *ApiKeyRequestResponseData) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetResourceType

`func (o *ApiKeyRequestResponseData) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ApiKeyRequestResponseData) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ApiKeyRequestResponseData) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ApiKeyRequestResponseData) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *ApiKeyRequestResponseData) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *ApiKeyRequestResponseData) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetRespBody

`func (o *ApiKeyRequestResponseData) GetRespBody() string`

GetRespBody returns the RespBody field if non-nil, zero value otherwise.

### GetRespBodyOk

`func (o *ApiKeyRequestResponseData) GetRespBodyOk() (*string, bool)`

GetRespBodyOk returns a tuple with the RespBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespBody

`func (o *ApiKeyRequestResponseData) SetRespBody(v string)`

SetRespBody sets RespBody field to given value.

### HasRespBody

`func (o *ApiKeyRequestResponseData) HasRespBody() bool`

HasRespBody returns a boolean if a field has been set.

### SetRespBodyNil

`func (o *ApiKeyRequestResponseData) SetRespBodyNil(b bool)`

 SetRespBodyNil sets the value for RespBody to be an explicit nil

### UnsetRespBody
`func (o *ApiKeyRequestResponseData) UnsetRespBody()`

UnsetRespBody ensures that no value is present for RespBody, not even an explicit nil
### GetRespCode

`func (o *ApiKeyRequestResponseData) GetRespCode() int32`

GetRespCode returns the RespCode field if non-nil, zero value otherwise.

### GetRespCodeOk

`func (o *ApiKeyRequestResponseData) GetRespCodeOk() (*int32, bool)`

GetRespCodeOk returns a tuple with the RespCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespCode

`func (o *ApiKeyRequestResponseData) SetRespCode(v int32)`

SetRespCode sets RespCode field to given value.

### HasRespCode

`func (o *ApiKeyRequestResponseData) HasRespCode() bool`

HasRespCode returns a boolean if a field has been set.

### SetRespCodeNil

`func (o *ApiKeyRequestResponseData) SetRespCodeNil(b bool)`

 SetRespCodeNil sets the value for RespCode to be an explicit nil

### UnsetRespCode
`func (o *ApiKeyRequestResponseData) UnsetRespCode()`

UnsetRespCode ensures that no value is present for RespCode, not even an explicit nil
### GetSecondaryResource

`func (o *ApiKeyRequestResponseData) GetSecondaryResource() string`

GetSecondaryResource returns the SecondaryResource field if non-nil, zero value otherwise.

### GetSecondaryResourceOk

`func (o *ApiKeyRequestResponseData) GetSecondaryResourceOk() (*string, bool)`

GetSecondaryResourceOk returns a tuple with the SecondaryResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryResource

`func (o *ApiKeyRequestResponseData) SetSecondaryResource(v string)`

SetSecondaryResource sets SecondaryResource field to given value.

### HasSecondaryResource

`func (o *ApiKeyRequestResponseData) HasSecondaryResource() bool`

HasSecondaryResource returns a boolean if a field has been set.

### SetSecondaryResourceNil

`func (o *ApiKeyRequestResponseData) SetSecondaryResourceNil(b bool)`

 SetSecondaryResourceNil sets the value for SecondaryResource to be an explicit nil

### UnsetSecondaryResource
`func (o *ApiKeyRequestResponseData) UnsetSecondaryResource()`

UnsetSecondaryResource ensures that no value is present for SecondaryResource, not even an explicit nil
### GetStartedAt

`func (o *ApiKeyRequestResponseData) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ApiKeyRequestResponseData) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ApiKeyRequestResponseData) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetUrl

`func (o *ApiKeyRequestResponseData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApiKeyRequestResponseData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApiKeyRequestResponseData) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUserAgent

`func (o *ApiKeyRequestResponseData) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *ApiKeyRequestResponseData) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *ApiKeyRequestResponseData) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *ApiKeyRequestResponseData) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### SetUserAgentNil

`func (o *ApiKeyRequestResponseData) SetUserAgentNil(b bool)`

 SetUserAgentNil sets the value for UserAgent to be an explicit nil

### UnsetUserAgent
`func (o *ApiKeyRequestResponseData) UnsetUserAgent()`

UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
### GetUserId

`func (o *ApiKeyRequestResponseData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApiKeyRequestResponseData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApiKeyRequestResponseData) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApiKeyRequestResponseData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ApiKeyRequestResponseData) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ApiKeyRequestResponseData) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetUserName

`func (o *ApiKeyRequestResponseData) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ApiKeyRequestResponseData) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ApiKeyRequestResponseData) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ApiKeyRequestResponseData) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *ApiKeyRequestResponseData) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *ApiKeyRequestResponseData) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


