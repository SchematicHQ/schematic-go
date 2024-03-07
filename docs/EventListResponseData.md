# EventListResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **NullableString** |  | [optional] 
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

### NewEventListResponseData

`func NewEventListResponseData(bodyPreview string, capturedAt time.Time, featureIds []string, id string, status string, type_ string, updatedAt time.Time, ) *EventListResponseData`

NewEventListResponseData instantiates a new EventListResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventListResponseDataWithDefaults

`func NewEventListResponseDataWithDefaults() *EventListResponseData`

NewEventListResponseDataWithDefaults instantiates a new EventListResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *EventListResponseData) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *EventListResponseData) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *EventListResponseData) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *EventListResponseData) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *EventListResponseData) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *EventListResponseData) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetBodyPreview

`func (o *EventListResponseData) GetBodyPreview() string`

GetBodyPreview returns the BodyPreview field if non-nil, zero value otherwise.

### GetBodyPreviewOk

`func (o *EventListResponseData) GetBodyPreviewOk() (*string, bool)`

GetBodyPreviewOk returns a tuple with the BodyPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyPreview

`func (o *EventListResponseData) SetBodyPreview(v string)`

SetBodyPreview sets BodyPreview field to given value.


### GetCapturedAt

`func (o *EventListResponseData) GetCapturedAt() time.Time`

GetCapturedAt returns the CapturedAt field if non-nil, zero value otherwise.

### GetCapturedAtOk

`func (o *EventListResponseData) GetCapturedAtOk() (*time.Time, bool)`

GetCapturedAtOk returns a tuple with the CapturedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedAt

`func (o *EventListResponseData) SetCapturedAt(v time.Time)`

SetCapturedAt sets CapturedAt field to given value.


### GetCompanyId

`func (o *EventListResponseData) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *EventListResponseData) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *EventListResponseData) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *EventListResponseData) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *EventListResponseData) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *EventListResponseData) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetEnrichedAt

`func (o *EventListResponseData) GetEnrichedAt() time.Time`

GetEnrichedAt returns the EnrichedAt field if non-nil, zero value otherwise.

### GetEnrichedAtOk

`func (o *EventListResponseData) GetEnrichedAtOk() (*time.Time, bool)`

GetEnrichedAtOk returns a tuple with the EnrichedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrichedAt

`func (o *EventListResponseData) SetEnrichedAt(v time.Time)`

SetEnrichedAt sets EnrichedAt field to given value.

### HasEnrichedAt

`func (o *EventListResponseData) HasEnrichedAt() bool`

HasEnrichedAt returns a boolean if a field has been set.

### SetEnrichedAtNil

`func (o *EventListResponseData) SetEnrichedAtNil(b bool)`

 SetEnrichedAtNil sets the value for EnrichedAt to be an explicit nil

### UnsetEnrichedAt
`func (o *EventListResponseData) UnsetEnrichedAt()`

UnsetEnrichedAt ensures that no value is present for EnrichedAt, not even an explicit nil
### GetEnvironmentId

`func (o *EventListResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *EventListResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *EventListResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *EventListResponseData) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### SetEnvironmentIdNil

`func (o *EventListResponseData) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *EventListResponseData) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
### GetErrorMessage

`func (o *EventListResponseData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *EventListResponseData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *EventListResponseData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *EventListResponseData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *EventListResponseData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *EventListResponseData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetFeatureIds

`func (o *EventListResponseData) GetFeatureIds() []string`

GetFeatureIds returns the FeatureIds field if non-nil, zero value otherwise.

### GetFeatureIdsOk

`func (o *EventListResponseData) GetFeatureIdsOk() (*[]string, bool)`

GetFeatureIdsOk returns a tuple with the FeatureIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureIds

`func (o *EventListResponseData) SetFeatureIds(v []string)`

SetFeatureIds sets FeatureIds field to given value.


### GetId

`func (o *EventListResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventListResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventListResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetLoadedAt

`func (o *EventListResponseData) GetLoadedAt() time.Time`

GetLoadedAt returns the LoadedAt field if non-nil, zero value otherwise.

### GetLoadedAtOk

`func (o *EventListResponseData) GetLoadedAtOk() (*time.Time, bool)`

GetLoadedAtOk returns a tuple with the LoadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadedAt

`func (o *EventListResponseData) SetLoadedAt(v time.Time)`

SetLoadedAt sets LoadedAt field to given value.

### HasLoadedAt

`func (o *EventListResponseData) HasLoadedAt() bool`

HasLoadedAt returns a boolean if a field has been set.

### SetLoadedAtNil

`func (o *EventListResponseData) SetLoadedAtNil(b bool)`

 SetLoadedAtNil sets the value for LoadedAt to be an explicit nil

### UnsetLoadedAt
`func (o *EventListResponseData) UnsetLoadedAt()`

UnsetLoadedAt ensures that no value is present for LoadedAt, not even an explicit nil
### GetProcessedAt

`func (o *EventListResponseData) GetProcessedAt() time.Time`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *EventListResponseData) GetProcessedAtOk() (*time.Time, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *EventListResponseData) SetProcessedAt(v time.Time)`

SetProcessedAt sets ProcessedAt field to given value.

### HasProcessedAt

`func (o *EventListResponseData) HasProcessedAt() bool`

HasProcessedAt returns a boolean if a field has been set.

### SetProcessedAtNil

`func (o *EventListResponseData) SetProcessedAtNil(b bool)`

 SetProcessedAtNil sets the value for ProcessedAt to be an explicit nil

### UnsetProcessedAt
`func (o *EventListResponseData) UnsetProcessedAt()`

UnsetProcessedAt ensures that no value is present for ProcessedAt, not even an explicit nil
### GetSentAt

`func (o *EventListResponseData) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *EventListResponseData) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *EventListResponseData) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *EventListResponseData) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### SetSentAtNil

`func (o *EventListResponseData) SetSentAtNil(b bool)`

 SetSentAtNil sets the value for SentAt to be an explicit nil

### UnsetSentAt
`func (o *EventListResponseData) UnsetSentAt()`

UnsetSentAt ensures that no value is present for SentAt, not even an explicit nil
### GetStatus

`func (o *EventListResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventListResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventListResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubtype

`func (o *EventListResponseData) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *EventListResponseData) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *EventListResponseData) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *EventListResponseData) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### SetSubtypeNil

`func (o *EventListResponseData) SetSubtypeNil(b bool)`

 SetSubtypeNil sets the value for Subtype to be an explicit nil

### UnsetSubtype
`func (o *EventListResponseData) UnsetSubtype()`

UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
### GetType

`func (o *EventListResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventListResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventListResponseData) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *EventListResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EventListResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EventListResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *EventListResponseData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EventListResponseData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EventListResponseData) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EventListResponseData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *EventListResponseData) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *EventListResponseData) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


