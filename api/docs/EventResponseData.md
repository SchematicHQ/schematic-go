# EventResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **NullableString** |  | [optional] 
**Body** | **map[string]interface{}** |  | 
**BodyPreview** | **string** |  | 
**CapturedAt** | **time.Time** |  | 
**CompanyId** | Pointer to **NullableString** |  | [optional] 
**EnrichedAt** | Pointer to **NullableTime** |  | [optional] 
**EnvironmentId** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**FeatureIds** | **[]string** |  | 
**Id** | **string** |  | 
**LoadedAt** | Pointer to **NullableTime** |  | [optional] 
**ProcessedAt** | Pointer to **NullableTime** |  | [optional] 
**SentAt** | Pointer to **NullableTime** |  | [optional] 
**Status** | **string** |  | 
**Subtype** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**UserId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEventResponseData

`func NewEventResponseData(body map[string]interface{}, bodyPreview string, capturedAt time.Time, featureIds []string, id string, status string, type_ string, updatedAt time.Time, ) *EventResponseData`

NewEventResponseData instantiates a new EventResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventResponseDataWithDefaults

`func NewEventResponseDataWithDefaults() *EventResponseData`

NewEventResponseDataWithDefaults instantiates a new EventResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *EventResponseData) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *EventResponseData) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *EventResponseData) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *EventResponseData) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *EventResponseData) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *EventResponseData) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetBody

`func (o *EventResponseData) GetBody() map[string]interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EventResponseData) GetBodyOk() (*map[string]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EventResponseData) SetBody(v map[string]interface{})`

SetBody sets Body field to given value.


### GetBodyPreview

`func (o *EventResponseData) GetBodyPreview() string`

GetBodyPreview returns the BodyPreview field if non-nil, zero value otherwise.

### GetBodyPreviewOk

`func (o *EventResponseData) GetBodyPreviewOk() (*string, bool)`

GetBodyPreviewOk returns a tuple with the BodyPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyPreview

`func (o *EventResponseData) SetBodyPreview(v string)`

SetBodyPreview sets BodyPreview field to given value.


### GetCapturedAt

`func (o *EventResponseData) GetCapturedAt() time.Time`

GetCapturedAt returns the CapturedAt field if non-nil, zero value otherwise.

### GetCapturedAtOk

`func (o *EventResponseData) GetCapturedAtOk() (*time.Time, bool)`

GetCapturedAtOk returns a tuple with the CapturedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedAt

`func (o *EventResponseData) SetCapturedAt(v time.Time)`

SetCapturedAt sets CapturedAt field to given value.


### GetCompanyId

`func (o *EventResponseData) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *EventResponseData) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *EventResponseData) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *EventResponseData) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *EventResponseData) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *EventResponseData) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetEnrichedAt

`func (o *EventResponseData) GetEnrichedAt() time.Time`

GetEnrichedAt returns the EnrichedAt field if non-nil, zero value otherwise.

### GetEnrichedAtOk

`func (o *EventResponseData) GetEnrichedAtOk() (*time.Time, bool)`

GetEnrichedAtOk returns a tuple with the EnrichedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrichedAt

`func (o *EventResponseData) SetEnrichedAt(v time.Time)`

SetEnrichedAt sets EnrichedAt field to given value.

### HasEnrichedAt

`func (o *EventResponseData) HasEnrichedAt() bool`

HasEnrichedAt returns a boolean if a field has been set.

### SetEnrichedAtNil

`func (o *EventResponseData) SetEnrichedAtNil(b bool)`

 SetEnrichedAtNil sets the value for EnrichedAt to be an explicit nil

### UnsetEnrichedAt
`func (o *EventResponseData) UnsetEnrichedAt()`

UnsetEnrichedAt ensures that no value is present for EnrichedAt, not even an explicit nil
### GetEnvironmentId

`func (o *EventResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *EventResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *EventResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *EventResponseData) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### SetEnvironmentIdNil

`func (o *EventResponseData) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *EventResponseData) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
### GetErrorMessage

`func (o *EventResponseData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *EventResponseData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *EventResponseData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *EventResponseData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *EventResponseData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *EventResponseData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetFeatureIds

`func (o *EventResponseData) GetFeatureIds() []string`

GetFeatureIds returns the FeatureIds field if non-nil, zero value otherwise.

### GetFeatureIdsOk

`func (o *EventResponseData) GetFeatureIdsOk() (*[]string, bool)`

GetFeatureIdsOk returns a tuple with the FeatureIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureIds

`func (o *EventResponseData) SetFeatureIds(v []string)`

SetFeatureIds sets FeatureIds field to given value.


### GetId

`func (o *EventResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetLoadedAt

`func (o *EventResponseData) GetLoadedAt() time.Time`

GetLoadedAt returns the LoadedAt field if non-nil, zero value otherwise.

### GetLoadedAtOk

`func (o *EventResponseData) GetLoadedAtOk() (*time.Time, bool)`

GetLoadedAtOk returns a tuple with the LoadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadedAt

`func (o *EventResponseData) SetLoadedAt(v time.Time)`

SetLoadedAt sets LoadedAt field to given value.

### HasLoadedAt

`func (o *EventResponseData) HasLoadedAt() bool`

HasLoadedAt returns a boolean if a field has been set.

### SetLoadedAtNil

`func (o *EventResponseData) SetLoadedAtNil(b bool)`

 SetLoadedAtNil sets the value for LoadedAt to be an explicit nil

### UnsetLoadedAt
`func (o *EventResponseData) UnsetLoadedAt()`

UnsetLoadedAt ensures that no value is present for LoadedAt, not even an explicit nil
### GetProcessedAt

`func (o *EventResponseData) GetProcessedAt() time.Time`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *EventResponseData) GetProcessedAtOk() (*time.Time, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *EventResponseData) SetProcessedAt(v time.Time)`

SetProcessedAt sets ProcessedAt field to given value.

### HasProcessedAt

`func (o *EventResponseData) HasProcessedAt() bool`

HasProcessedAt returns a boolean if a field has been set.

### SetProcessedAtNil

`func (o *EventResponseData) SetProcessedAtNil(b bool)`

 SetProcessedAtNil sets the value for ProcessedAt to be an explicit nil

### UnsetProcessedAt
`func (o *EventResponseData) UnsetProcessedAt()`

UnsetProcessedAt ensures that no value is present for ProcessedAt, not even an explicit nil
### GetSentAt

`func (o *EventResponseData) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *EventResponseData) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *EventResponseData) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *EventResponseData) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### SetSentAtNil

`func (o *EventResponseData) SetSentAtNil(b bool)`

 SetSentAtNil sets the value for SentAt to be an explicit nil

### UnsetSentAt
`func (o *EventResponseData) UnsetSentAt()`

UnsetSentAt ensures that no value is present for SentAt, not even an explicit nil
### GetStatus

`func (o *EventResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubtype

`func (o *EventResponseData) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *EventResponseData) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *EventResponseData) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *EventResponseData) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### SetSubtypeNil

`func (o *EventResponseData) SetSubtypeNil(b bool)`

 SetSubtypeNil sets the value for Subtype to be an explicit nil

### UnsetSubtype
`func (o *EventResponseData) UnsetSubtype()`

UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
### GetType

`func (o *EventResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventResponseData) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *EventResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EventResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EventResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *EventResponseData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EventResponseData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EventResponseData) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EventResponseData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *EventResponseData) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *EventResponseData) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


