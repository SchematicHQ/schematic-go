# UpdateAudienceRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionGroups** | [**[]CreateOrUpdateConditionGroupRequestBody**](CreateOrUpdateConditionGroupRequestBody.md) |  | 
**Conditions** | [**[]CreateOrUpdateConditionRequestBody**](CreateOrUpdateConditionRequestBody.md) |  | 
**SkipWebhooks** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpdateAudienceRequestBody

`func NewUpdateAudienceRequestBody(conditionGroups []CreateOrUpdateConditionGroupRequestBody, conditions []CreateOrUpdateConditionRequestBody, ) *UpdateAudienceRequestBody`

NewUpdateAudienceRequestBody instantiates a new UpdateAudienceRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAudienceRequestBodyWithDefaults

`func NewUpdateAudienceRequestBodyWithDefaults() *UpdateAudienceRequestBody`

NewUpdateAudienceRequestBodyWithDefaults instantiates a new UpdateAudienceRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionGroups

`func (o *UpdateAudienceRequestBody) GetConditionGroups() []CreateOrUpdateConditionGroupRequestBody`

GetConditionGroups returns the ConditionGroups field if non-nil, zero value otherwise.

### GetConditionGroupsOk

`func (o *UpdateAudienceRequestBody) GetConditionGroupsOk() (*[]CreateOrUpdateConditionGroupRequestBody, bool)`

GetConditionGroupsOk returns a tuple with the ConditionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionGroups

`func (o *UpdateAudienceRequestBody) SetConditionGroups(v []CreateOrUpdateConditionGroupRequestBody)`

SetConditionGroups sets ConditionGroups field to given value.


### GetConditions

`func (o *UpdateAudienceRequestBody) GetConditions() []CreateOrUpdateConditionRequestBody`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *UpdateAudienceRequestBody) GetConditionsOk() (*[]CreateOrUpdateConditionRequestBody, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *UpdateAudienceRequestBody) SetConditions(v []CreateOrUpdateConditionRequestBody)`

SetConditions sets Conditions field to given value.


### GetSkipWebhooks

`func (o *UpdateAudienceRequestBody) GetSkipWebhooks() bool`

GetSkipWebhooks returns the SkipWebhooks field if non-nil, zero value otherwise.

### GetSkipWebhooksOk

`func (o *UpdateAudienceRequestBody) GetSkipWebhooksOk() (*bool, bool)`

GetSkipWebhooksOk returns a tuple with the SkipWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWebhooks

`func (o *UpdateAudienceRequestBody) SetSkipWebhooks(v bool)`

SetSkipWebhooks sets SkipWebhooks field to given value.

### HasSkipWebhooks

`func (o *UpdateAudienceRequestBody) HasSkipWebhooks() bool`

HasSkipWebhooks returns a boolean if a field has been set.

### SetSkipWebhooksNil

`func (o *UpdateAudienceRequestBody) SetSkipWebhooksNil(b bool)`

 SetSkipWebhooksNil sets the value for SkipWebhooks to be an explicit nil

### UnsetSkipWebhooks
`func (o *UpdateAudienceRequestBody) UnsetSkipWebhooks()`

UnsetSkipWebhooks ensures that no value is present for SkipWebhooks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


