# UpsertCompanyRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Keys** | **map[string]interface{}** |  | 
**LastSeenAt** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**SkipWebhooks** | Pointer to **NullableBool** |  | [optional] 
**Traits** | Pointer to **map[string]interface{}** | A map of trait names to trait values | [optional] 
**UpdateOnly** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpsertCompanyRequestBody

`func NewUpsertCompanyRequestBody(keys map[string]interface{}, ) *UpsertCompanyRequestBody`

NewUpsertCompanyRequestBody instantiates a new UpsertCompanyRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpsertCompanyRequestBodyWithDefaults

`func NewUpsertCompanyRequestBodyWithDefaults() *UpsertCompanyRequestBody`

NewUpsertCompanyRequestBodyWithDefaults instantiates a new UpsertCompanyRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpsertCompanyRequestBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpsertCompanyRequestBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpsertCompanyRequestBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpsertCompanyRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UpsertCompanyRequestBody) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UpsertCompanyRequestBody) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetKeys

`func (o *UpsertCompanyRequestBody) GetKeys() map[string]interface{}`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *UpsertCompanyRequestBody) GetKeysOk() (*map[string]interface{}, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *UpsertCompanyRequestBody) SetKeys(v map[string]interface{})`

SetKeys sets Keys field to given value.


### GetLastSeenAt

`func (o *UpsertCompanyRequestBody) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *UpsertCompanyRequestBody) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *UpsertCompanyRequestBody) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *UpsertCompanyRequestBody) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### SetLastSeenAtNil

`func (o *UpsertCompanyRequestBody) SetLastSeenAtNil(b bool)`

 SetLastSeenAtNil sets the value for LastSeenAt to be an explicit nil

### UnsetLastSeenAt
`func (o *UpsertCompanyRequestBody) UnsetLastSeenAt()`

UnsetLastSeenAt ensures that no value is present for LastSeenAt, not even an explicit nil
### GetName

`func (o *UpsertCompanyRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpsertCompanyRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpsertCompanyRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpsertCompanyRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpsertCompanyRequestBody) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpsertCompanyRequestBody) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSkipWebhooks

`func (o *UpsertCompanyRequestBody) GetSkipWebhooks() bool`

GetSkipWebhooks returns the SkipWebhooks field if non-nil, zero value otherwise.

### GetSkipWebhooksOk

`func (o *UpsertCompanyRequestBody) GetSkipWebhooksOk() (*bool, bool)`

GetSkipWebhooksOk returns a tuple with the SkipWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWebhooks

`func (o *UpsertCompanyRequestBody) SetSkipWebhooks(v bool)`

SetSkipWebhooks sets SkipWebhooks field to given value.

### HasSkipWebhooks

`func (o *UpsertCompanyRequestBody) HasSkipWebhooks() bool`

HasSkipWebhooks returns a boolean if a field has been set.

### SetSkipWebhooksNil

`func (o *UpsertCompanyRequestBody) SetSkipWebhooksNil(b bool)`

 SetSkipWebhooksNil sets the value for SkipWebhooks to be an explicit nil

### UnsetSkipWebhooks
`func (o *UpsertCompanyRequestBody) UnsetSkipWebhooks()`

UnsetSkipWebhooks ensures that no value is present for SkipWebhooks, not even an explicit nil
### GetTraits

`func (o *UpsertCompanyRequestBody) GetTraits() map[string]interface{}`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *UpsertCompanyRequestBody) GetTraitsOk() (*map[string]interface{}, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *UpsertCompanyRequestBody) SetTraits(v map[string]interface{})`

SetTraits sets Traits field to given value.

### HasTraits

`func (o *UpsertCompanyRequestBody) HasTraits() bool`

HasTraits returns a boolean if a field has been set.

### GetUpdateOnly

`func (o *UpsertCompanyRequestBody) GetUpdateOnly() bool`

GetUpdateOnly returns the UpdateOnly field if non-nil, zero value otherwise.

### GetUpdateOnlyOk

`func (o *UpsertCompanyRequestBody) GetUpdateOnlyOk() (*bool, bool)`

GetUpdateOnlyOk returns a tuple with the UpdateOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateOnly

`func (o *UpsertCompanyRequestBody) SetUpdateOnly(v bool)`

SetUpdateOnly sets UpdateOnly field to given value.

### HasUpdateOnly

`func (o *UpsertCompanyRequestBody) HasUpdateOnly() bool`

HasUpdateOnly returns a boolean if a field has been set.

### SetUpdateOnlyNil

`func (o *UpsertCompanyRequestBody) SetUpdateOnlyNil(b bool)`

 SetUpdateOnlyNil sets the value for UpdateOnly to be an explicit nil

### UnsetUpdateOnly
`func (o *UpsertCompanyRequestBody) UnsetUpdateOnly()`

UnsetUpdateOnly ensures that no value is present for UpdateOnly, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


