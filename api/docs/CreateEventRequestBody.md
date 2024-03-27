# CreateEventRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to [**EventBody**](EventBody.md) |  | [optional] 
**EventType** | **string** | Either &#39;identify&#39; or &#39;track&#39; | 
**SentAt** | Pointer to **NullableTime** | Optionally provide a timestamp at which the event was sent to Schematic | [optional] 

## Methods

### NewCreateEventRequestBody

`func NewCreateEventRequestBody(eventType string, ) *CreateEventRequestBody`

NewCreateEventRequestBody instantiates a new CreateEventRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEventRequestBodyWithDefaults

`func NewCreateEventRequestBodyWithDefaults() *CreateEventRequestBody`

NewCreateEventRequestBodyWithDefaults instantiates a new CreateEventRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *CreateEventRequestBody) GetBody() EventBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateEventRequestBody) GetBodyOk() (*EventBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateEventRequestBody) SetBody(v EventBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *CreateEventRequestBody) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetEventType

`func (o *CreateEventRequestBody) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CreateEventRequestBody) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CreateEventRequestBody) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetSentAt

`func (o *CreateEventRequestBody) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *CreateEventRequestBody) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *CreateEventRequestBody) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *CreateEventRequestBody) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### SetSentAtNil

`func (o *CreateEventRequestBody) SetSentAtNil(b bool)`

 SetSentAtNil sets the value for SentAt to be an explicit nil

### UnsetSentAt
`func (o *CreateEventRequestBody) UnsetSentAt()`

UnsetSentAt ensures that no value is present for SentAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


