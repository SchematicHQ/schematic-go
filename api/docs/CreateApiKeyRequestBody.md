# CreateApiKeyRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**EnvironmentId** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewCreateApiKeyRequestBody

`func NewCreateApiKeyRequestBody(name string, ) *CreateApiKeyRequestBody`

NewCreateApiKeyRequestBody instantiates a new CreateApiKeyRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiKeyRequestBodyWithDefaults

`func NewCreateApiKeyRequestBodyWithDefaults() *CreateApiKeyRequestBody`

NewCreateApiKeyRequestBodyWithDefaults instantiates a new CreateApiKeyRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateApiKeyRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateApiKeyRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateApiKeyRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateApiKeyRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateApiKeyRequestBody) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateApiKeyRequestBody) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnvironmentId

`func (o *CreateApiKeyRequestBody) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *CreateApiKeyRequestBody) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *CreateApiKeyRequestBody) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *CreateApiKeyRequestBody) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### SetEnvironmentIdNil

`func (o *CreateApiKeyRequestBody) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *CreateApiKeyRequestBody) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
### GetName

`func (o *CreateApiKeyRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateApiKeyRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateApiKeyRequestBody) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


