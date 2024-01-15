# UpdatePlanRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**SkipWebhooks** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpdatePlanRequestBody

`func NewUpdatePlanRequestBody(name string, ) *UpdatePlanRequestBody`

NewUpdatePlanRequestBody instantiates a new UpdatePlanRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePlanRequestBodyWithDefaults

`func NewUpdatePlanRequestBodyWithDefaults() *UpdatePlanRequestBody`

NewUpdatePlanRequestBodyWithDefaults instantiates a new UpdatePlanRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdatePlanRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePlanRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePlanRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetSkipWebhooks

`func (o *UpdatePlanRequestBody) GetSkipWebhooks() bool`

GetSkipWebhooks returns the SkipWebhooks field if non-nil, zero value otherwise.

### GetSkipWebhooksOk

`func (o *UpdatePlanRequestBody) GetSkipWebhooksOk() (*bool, bool)`

GetSkipWebhooksOk returns a tuple with the SkipWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWebhooks

`func (o *UpdatePlanRequestBody) SetSkipWebhooks(v bool)`

SetSkipWebhooks sets SkipWebhooks field to given value.

### HasSkipWebhooks

`func (o *UpdatePlanRequestBody) HasSkipWebhooks() bool`

HasSkipWebhooks returns a boolean if a field has been set.

### SetSkipWebhooksNil

`func (o *UpdatePlanRequestBody) SetSkipWebhooksNil(b bool)`

 SetSkipWebhooksNil sets the value for SkipWebhooks to be an explicit nil

### UnsetSkipWebhooks
`func (o *UpdatePlanRequestBody) UnsetSkipWebhooks()`

UnsetSkipWebhooks ensures that no value is present for SkipWebhooks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


