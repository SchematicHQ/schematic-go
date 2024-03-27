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
**RequestType** | Pointer to **NullableString** |  | [optional] 
**ResourceId** | Pointer to **NullableInt32** |  | [optional] 
**ResourceIdString** | Pointer to **NullableString** |  | [optional] 
**ResourceName** | Pointer to **NullableString** |  | [optional] 
**ResourceType** | Pointer to **NullableString** |  | [optional] 
**RespCode** | Pointer to **NullableInt32** |  | [optional] 
**SecondaryResource** | Pointer to **NullableString** |  | [optional] 
**StartedAt** | **time.Time** |  | 
**Url** | **string** |  | 
**UserName** | Pointer to **NullableString** |  | [optional] 

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
### GetRequestType

`func (o *ApiKeyRequestListResponseData) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *ApiKeyRequestListResponseData) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *ApiKeyRequestListResponseData) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *ApiKeyRequestListResponseData) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### SetRequestTypeNil

`func (o *ApiKeyRequestListResponseData) SetRequestTypeNil(b bool)`

 SetRequestTypeNil sets the value for RequestType to be an explicit nil

### UnsetRequestType
`func (o *ApiKeyRequestListResponseData) UnsetRequestType()`

UnsetRequestType ensures that no value is present for RequestType, not even an explicit nil
### GetResourceId

`func (o *ApiKeyRequestListResponseData) GetResourceId() int32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ApiKeyRequestListResponseData) GetResourceIdOk() (*int32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ApiKeyRequestListResponseData) SetResourceId(v int32)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *ApiKeyRequestListResponseData) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *ApiKeyRequestListResponseData) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *ApiKeyRequestListResponseData) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetResourceIdString

`func (o *ApiKeyRequestListResponseData) GetResourceIdString() string`

GetResourceIdString returns the ResourceIdString field if non-nil, zero value otherwise.

### GetResourceIdStringOk

`func (o *ApiKeyRequestListResponseData) GetResourceIdStringOk() (*string, bool)`

GetResourceIdStringOk returns a tuple with the ResourceIdString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIdString

`func (o *ApiKeyRequestListResponseData) SetResourceIdString(v string)`

SetResourceIdString sets ResourceIdString field to given value.

### HasResourceIdString

`func (o *ApiKeyRequestListResponseData) HasResourceIdString() bool`

HasResourceIdString returns a boolean if a field has been set.

### SetResourceIdStringNil

`func (o *ApiKeyRequestListResponseData) SetResourceIdStringNil(b bool)`

 SetResourceIdStringNil sets the value for ResourceIdString to be an explicit nil

### UnsetResourceIdString
`func (o *ApiKeyRequestListResponseData) UnsetResourceIdString()`

UnsetResourceIdString ensures that no value is present for ResourceIdString, not even an explicit nil
### GetResourceName

`func (o *ApiKeyRequestListResponseData) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ApiKeyRequestListResponseData) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ApiKeyRequestListResponseData) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *ApiKeyRequestListResponseData) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *ApiKeyRequestListResponseData) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *ApiKeyRequestListResponseData) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetResourceType

`func (o *ApiKeyRequestListResponseData) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ApiKeyRequestListResponseData) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ApiKeyRequestListResponseData) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ApiKeyRequestListResponseData) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *ApiKeyRequestListResponseData) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *ApiKeyRequestListResponseData) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
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
### GetSecondaryResource

`func (o *ApiKeyRequestListResponseData) GetSecondaryResource() string`

GetSecondaryResource returns the SecondaryResource field if non-nil, zero value otherwise.

### GetSecondaryResourceOk

`func (o *ApiKeyRequestListResponseData) GetSecondaryResourceOk() (*string, bool)`

GetSecondaryResourceOk returns a tuple with the SecondaryResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryResource

`func (o *ApiKeyRequestListResponseData) SetSecondaryResource(v string)`

SetSecondaryResource sets SecondaryResource field to given value.

### HasSecondaryResource

`func (o *ApiKeyRequestListResponseData) HasSecondaryResource() bool`

HasSecondaryResource returns a boolean if a field has been set.

### SetSecondaryResourceNil

`func (o *ApiKeyRequestListResponseData) SetSecondaryResourceNil(b bool)`

 SetSecondaryResourceNil sets the value for SecondaryResource to be an explicit nil

### UnsetSecondaryResource
`func (o *ApiKeyRequestListResponseData) UnsetSecondaryResource()`

UnsetSecondaryResource ensures that no value is present for SecondaryResource, not even an explicit nil
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


### GetUserName

`func (o *ApiKeyRequestListResponseData) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ApiKeyRequestListResponseData) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ApiKeyRequestListResponseData) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ApiKeyRequestListResponseData) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *ApiKeyRequestListResponseData) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *ApiKeyRequestListResponseData) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


