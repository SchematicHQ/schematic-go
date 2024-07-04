# WebhookEventDetailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 
**Payload** | Pointer to **NullableString** |  | [optional] 
**RequestType** | **string** |  | 
**ResponseCode** | Pointer to **NullableInt32** |  | [optional] 
**SentAt** | Pointer to **NullableTime** |  | [optional] 
**Status** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**Webhook** | Pointer to [**WebhookResponseData**](WebhookResponseData.md) |  | [optional] 
**WebhookId** | **string** |  | 

## Methods

### NewWebhookEventDetailResponseData

`func NewWebhookEventDetailResponseData(createdAt time.Time, id string, requestType string, status string, updatedAt time.Time, webhookId string, ) *WebhookEventDetailResponseData`

NewWebhookEventDetailResponseData instantiates a new WebhookEventDetailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventDetailResponseDataWithDefaults

`func NewWebhookEventDetailResponseDataWithDefaults() *WebhookEventDetailResponseData`

NewWebhookEventDetailResponseDataWithDefaults instantiates a new WebhookEventDetailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WebhookEventDetailResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookEventDetailResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookEventDetailResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *WebhookEventDetailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookEventDetailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookEventDetailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetPayload

`func (o *WebhookEventDetailResponseData) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WebhookEventDetailResponseData) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WebhookEventDetailResponseData) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *WebhookEventDetailResponseData) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *WebhookEventDetailResponseData) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *WebhookEventDetailResponseData) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetRequestType

`func (o *WebhookEventDetailResponseData) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *WebhookEventDetailResponseData) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *WebhookEventDetailResponseData) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.


### GetResponseCode

`func (o *WebhookEventDetailResponseData) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *WebhookEventDetailResponseData) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *WebhookEventDetailResponseData) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *WebhookEventDetailResponseData) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### SetResponseCodeNil

`func (o *WebhookEventDetailResponseData) SetResponseCodeNil(b bool)`

 SetResponseCodeNil sets the value for ResponseCode to be an explicit nil

### UnsetResponseCode
`func (o *WebhookEventDetailResponseData) UnsetResponseCode()`

UnsetResponseCode ensures that no value is present for ResponseCode, not even an explicit nil
### GetSentAt

`func (o *WebhookEventDetailResponseData) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *WebhookEventDetailResponseData) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *WebhookEventDetailResponseData) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *WebhookEventDetailResponseData) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### SetSentAtNil

`func (o *WebhookEventDetailResponseData) SetSentAtNil(b bool)`

 SetSentAtNil sets the value for SentAt to be an explicit nil

### UnsetSentAt
`func (o *WebhookEventDetailResponseData) UnsetSentAt()`

UnsetSentAt ensures that no value is present for SentAt, not even an explicit nil
### GetStatus

`func (o *WebhookEventDetailResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookEventDetailResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookEventDetailResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *WebhookEventDetailResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WebhookEventDetailResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WebhookEventDetailResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetWebhook

`func (o *WebhookEventDetailResponseData) GetWebhook() WebhookResponseData`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *WebhookEventDetailResponseData) GetWebhookOk() (*WebhookResponseData, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *WebhookEventDetailResponseData) SetWebhook(v WebhookResponseData)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *WebhookEventDetailResponseData) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### GetWebhookId

`func (o *WebhookEventDetailResponseData) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookEventDetailResponseData) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookEventDetailResponseData) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


