# CreateEnvironmentRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentType** | **string** |  | 
**Name** | **string** |  | 
**SkipWebhooks** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewCreateEnvironmentRequestBody

`func NewCreateEnvironmentRequestBody(environmentType string, name string, ) *CreateEnvironmentRequestBody`

NewCreateEnvironmentRequestBody instantiates a new CreateEnvironmentRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEnvironmentRequestBodyWithDefaults

`func NewCreateEnvironmentRequestBodyWithDefaults() *CreateEnvironmentRequestBody`

NewCreateEnvironmentRequestBodyWithDefaults instantiates a new CreateEnvironmentRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentType

`func (o *CreateEnvironmentRequestBody) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *CreateEnvironmentRequestBody) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *CreateEnvironmentRequestBody) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.


### GetName

`func (o *CreateEnvironmentRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEnvironmentRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEnvironmentRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetSkipWebhooks

`func (o *CreateEnvironmentRequestBody) GetSkipWebhooks() bool`

GetSkipWebhooks returns the SkipWebhooks field if non-nil, zero value otherwise.

### GetSkipWebhooksOk

`func (o *CreateEnvironmentRequestBody) GetSkipWebhooksOk() (*bool, bool)`

GetSkipWebhooksOk returns a tuple with the SkipWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWebhooks

`func (o *CreateEnvironmentRequestBody) SetSkipWebhooks(v bool)`

SetSkipWebhooks sets SkipWebhooks field to given value.

### HasSkipWebhooks

`func (o *CreateEnvironmentRequestBody) HasSkipWebhooks() bool`

HasSkipWebhooks returns a boolean if a field has been set.

### SetSkipWebhooksNil

`func (o *CreateEnvironmentRequestBody) SetSkipWebhooksNil(b bool)`

 SetSkipWebhooksNil sets the value for SkipWebhooks to be an explicit nil

### UnsetSkipWebhooks
`func (o *CreateEnvironmentRequestBody) UnsetSkipWebhooks()`

UnsetSkipWebhooks ensures that no value is present for SkipWebhooks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


