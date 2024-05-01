# EventDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **NullableString** |  | [optional] 
**Body** | **map[string]interface{}** |  | 
**BodyPreview** | **string** |  | 
**CapturedAt** | **time.Time** |  | 
**Company** | Pointer to [**PreviewObject**](PreviewObject.md) |  | [optional] 
**CompanyId** | Pointer to **NullableString** |  | [optional] 
**EnrichedAt** | Pointer to **NullableTime** |  | [optional] 
**EnvironmentId** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**FeatureIds** | **[]string** |  | 
**Features** | [**[]PreviewObject**](PreviewObject.md) |  | 
**Id** | **string** |  | 
**LoadedAt** | Pointer to **NullableTime** |  | [optional] 
**ProcessedAt** | Pointer to **NullableTime** |  | [optional] 
**SentAt** | Pointer to **NullableTime** |  | [optional] 
**Status** | **string** |  | 
**Subtype** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**User** | Pointer to [**PreviewObject**](PreviewObject.md) |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEventDetailResponseData

`func NewEventDetailResponseData(body map[string]interface{}, bodyPreview string, capturedAt time.Time, featureIds []string, features []PreviewObject, id string, status string, type_ string, updatedAt time.Time, ) *EventDetailResponseData`

NewEventDetailResponseData instantiates a new EventDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDetailResponseDataWithDefaults

`func NewEventDetailResponseDataWithDefaults() *EventDetailResponseData`

NewEventDetailResponseDataWithDefaults instantiates a new EventDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *EventDetailResponseData) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *EventDetailResponseData) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *EventDetailResponseData) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *EventDetailResponseData) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *EventDetailResponseData) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *EventDetailResponseData) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetBody

`func (o *EventDetailResponseData) GetBody() map[string]interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EventDetailResponseData) GetBodyOk() (*map[string]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EventDetailResponseData) SetBody(v map[string]interface{})`

SetBody sets Body field to given value.


### GetBodyPreview

`func (o *EventDetailResponseData) GetBodyPreview() string`

GetBodyPreview returns the BodyPreview field if non-nil, zero value otherwise.

### GetBodyPreviewOk

`func (o *EventDetailResponseData) GetBodyPreviewOk() (*string, bool)`

GetBodyPreviewOk returns a tuple with the BodyPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyPreview

`func (o *EventDetailResponseData) SetBodyPreview(v string)`

SetBodyPreview sets BodyPreview field to given value.


### GetCapturedAt

`func (o *EventDetailResponseData) GetCapturedAt() time.Time`

GetCapturedAt returns the CapturedAt field if non-nil, zero value otherwise.

### GetCapturedAtOk

`func (o *EventDetailResponseData) GetCapturedAtOk() (*time.Time, bool)`

GetCapturedAtOk returns a tuple with the CapturedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedAt

`func (o *EventDetailResponseData) SetCapturedAt(v time.Time)`

SetCapturedAt sets CapturedAt field to given value.


### GetCompany

`func (o *EventDetailResponseData) GetCompany() PreviewObject`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *EventDetailResponseData) GetCompanyOk() (*PreviewObject, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *EventDetailResponseData) SetCompany(v PreviewObject)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *EventDetailResponseData) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyId

`func (o *EventDetailResponseData) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *EventDetailResponseData) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *EventDetailResponseData) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *EventDetailResponseData) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### SetCompanyIdNil

`func (o *EventDetailResponseData) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *EventDetailResponseData) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetEnrichedAt

`func (o *EventDetailResponseData) GetEnrichedAt() time.Time`

GetEnrichedAt returns the EnrichedAt field if non-nil, zero value otherwise.

### GetEnrichedAtOk

`func (o *EventDetailResponseData) GetEnrichedAtOk() (*time.Time, bool)`

GetEnrichedAtOk returns a tuple with the EnrichedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrichedAt

`func (o *EventDetailResponseData) SetEnrichedAt(v time.Time)`

SetEnrichedAt sets EnrichedAt field to given value.

### HasEnrichedAt

`func (o *EventDetailResponseData) HasEnrichedAt() bool`

HasEnrichedAt returns a boolean if a field has been set.

### SetEnrichedAtNil

`func (o *EventDetailResponseData) SetEnrichedAtNil(b bool)`

 SetEnrichedAtNil sets the value for EnrichedAt to be an explicit nil

### UnsetEnrichedAt
`func (o *EventDetailResponseData) UnsetEnrichedAt()`

UnsetEnrichedAt ensures that no value is present for EnrichedAt, not even an explicit nil
### GetEnvironmentId

`func (o *EventDetailResponseData) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *EventDetailResponseData) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *EventDetailResponseData) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *EventDetailResponseData) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### SetEnvironmentIdNil

`func (o *EventDetailResponseData) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *EventDetailResponseData) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
### GetErrorMessage

`func (o *EventDetailResponseData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *EventDetailResponseData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *EventDetailResponseData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *EventDetailResponseData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *EventDetailResponseData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *EventDetailResponseData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetFeatureIds

`func (o *EventDetailResponseData) GetFeatureIds() []string`

GetFeatureIds returns the FeatureIds field if non-nil, zero value otherwise.

### GetFeatureIdsOk

`func (o *EventDetailResponseData) GetFeatureIdsOk() (*[]string, bool)`

GetFeatureIdsOk returns a tuple with the FeatureIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureIds

`func (o *EventDetailResponseData) SetFeatureIds(v []string)`

SetFeatureIds sets FeatureIds field to given value.


### GetFeatures

`func (o *EventDetailResponseData) GetFeatures() []PreviewObject`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *EventDetailResponseData) GetFeaturesOk() (*[]PreviewObject, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *EventDetailResponseData) SetFeatures(v []PreviewObject)`

SetFeatures sets Features field to given value.


### GetId

`func (o *EventDetailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventDetailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventDetailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetLoadedAt

`func (o *EventDetailResponseData) GetLoadedAt() time.Time`

GetLoadedAt returns the LoadedAt field if non-nil, zero value otherwise.

### GetLoadedAtOk

`func (o *EventDetailResponseData) GetLoadedAtOk() (*time.Time, bool)`

GetLoadedAtOk returns a tuple with the LoadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadedAt

`func (o *EventDetailResponseData) SetLoadedAt(v time.Time)`

SetLoadedAt sets LoadedAt field to given value.

### HasLoadedAt

`func (o *EventDetailResponseData) HasLoadedAt() bool`

HasLoadedAt returns a boolean if a field has been set.

### SetLoadedAtNil

`func (o *EventDetailResponseData) SetLoadedAtNil(b bool)`

 SetLoadedAtNil sets the value for LoadedAt to be an explicit nil

### UnsetLoadedAt
`func (o *EventDetailResponseData) UnsetLoadedAt()`

UnsetLoadedAt ensures that no value is present for LoadedAt, not even an explicit nil
### GetProcessedAt

`func (o *EventDetailResponseData) GetProcessedAt() time.Time`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *EventDetailResponseData) GetProcessedAtOk() (*time.Time, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *EventDetailResponseData) SetProcessedAt(v time.Time)`

SetProcessedAt sets ProcessedAt field to given value.

### HasProcessedAt

`func (o *EventDetailResponseData) HasProcessedAt() bool`

HasProcessedAt returns a boolean if a field has been set.

### SetProcessedAtNil

`func (o *EventDetailResponseData) SetProcessedAtNil(b bool)`

 SetProcessedAtNil sets the value for ProcessedAt to be an explicit nil

### UnsetProcessedAt
`func (o *EventDetailResponseData) UnsetProcessedAt()`

UnsetProcessedAt ensures that no value is present for ProcessedAt, not even an explicit nil
### GetSentAt

`func (o *EventDetailResponseData) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *EventDetailResponseData) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *EventDetailResponseData) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *EventDetailResponseData) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### SetSentAtNil

`func (o *EventDetailResponseData) SetSentAtNil(b bool)`

 SetSentAtNil sets the value for SentAt to be an explicit nil

### UnsetSentAt
`func (o *EventDetailResponseData) UnsetSentAt()`

UnsetSentAt ensures that no value is present for SentAt, not even an explicit nil
### GetStatus

`func (o *EventDetailResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventDetailResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventDetailResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubtype

`func (o *EventDetailResponseData) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *EventDetailResponseData) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *EventDetailResponseData) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *EventDetailResponseData) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### SetSubtypeNil

`func (o *EventDetailResponseData) SetSubtypeNil(b bool)`

 SetSubtypeNil sets the value for Subtype to be an explicit nil

### UnsetSubtype
`func (o *EventDetailResponseData) UnsetSubtype()`

UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
### GetType

`func (o *EventDetailResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventDetailResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventDetailResponseData) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *EventDetailResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EventDetailResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EventDetailResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUser

`func (o *EventDetailResponseData) GetUser() PreviewObject`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EventDetailResponseData) GetUserOk() (*PreviewObject, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EventDetailResponseData) SetUser(v PreviewObject)`

SetUser sets User field to given value.

### HasUser

`func (o *EventDetailResponseData) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *EventDetailResponseData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EventDetailResponseData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EventDetailResponseData) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EventDetailResponseData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *EventDetailResponseData) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *EventDetailResponseData) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


