# GetEventSummariesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventSubtypes** | Pointer to **[]string** |  | [optional] 
**Limit** | Pointer to **int32** | Page limit (default 100) | [optional] 
**Offset** | Pointer to **int32** | Page offset (default 0) | [optional] 
**Q** | Pointer to **string** |  | [optional] 

## Methods

### NewGetEventSummariesParams

`func NewGetEventSummariesParams() *GetEventSummariesParams`

NewGetEventSummariesParams instantiates a new GetEventSummariesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEventSummariesParamsWithDefaults

`func NewGetEventSummariesParamsWithDefaults() *GetEventSummariesParams`

NewGetEventSummariesParamsWithDefaults instantiates a new GetEventSummariesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventSubtypes

`func (o *GetEventSummariesParams) GetEventSubtypes() []string`

GetEventSubtypes returns the EventSubtypes field if non-nil, zero value otherwise.

### GetEventSubtypesOk

`func (o *GetEventSummariesParams) GetEventSubtypesOk() (*[]string, bool)`

GetEventSubtypesOk returns a tuple with the EventSubtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubtypes

`func (o *GetEventSummariesParams) SetEventSubtypes(v []string)`

SetEventSubtypes sets EventSubtypes field to given value.

### HasEventSubtypes

`func (o *GetEventSummariesParams) HasEventSubtypes() bool`

HasEventSubtypes returns a boolean if a field has been set.

### GetLimit

`func (o *GetEventSummariesParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetEventSummariesParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetEventSummariesParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetEventSummariesParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *GetEventSummariesParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetEventSummariesParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetEventSummariesParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetEventSummariesParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetQ

`func (o *GetEventSummariesParams) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *GetEventSummariesParams) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *GetEventSummariesParams) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *GetEventSummariesParams) HasQ() bool`

HasQ returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


