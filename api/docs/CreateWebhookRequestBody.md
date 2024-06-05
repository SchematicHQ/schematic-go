# CreateWebhookRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**RequestTypes** | **[]string** |  | 
**Url** | **string** |  | 

## Methods

### NewCreateWebhookRequestBody

`func NewCreateWebhookRequestBody(name string, requestTypes []string, url string, ) *CreateWebhookRequestBody`

NewCreateWebhookRequestBody instantiates a new CreateWebhookRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhookRequestBodyWithDefaults

`func NewCreateWebhookRequestBodyWithDefaults() *CreateWebhookRequestBody`

NewCreateWebhookRequestBodyWithDefaults instantiates a new CreateWebhookRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateWebhookRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWebhookRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWebhookRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetRequestTypes

`func (o *CreateWebhookRequestBody) GetRequestTypes() []string`

GetRequestTypes returns the RequestTypes field if non-nil, zero value otherwise.

### GetRequestTypesOk

`func (o *CreateWebhookRequestBody) GetRequestTypesOk() (*[]string, bool)`

GetRequestTypesOk returns a tuple with the RequestTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTypes

`func (o *CreateWebhookRequestBody) SetRequestTypes(v []string)`

SetRequestTypes sets RequestTypes field to given value.


### GetUrl

`func (o *CreateWebhookRequestBody) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateWebhookRequestBody) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateWebhookRequestBody) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


