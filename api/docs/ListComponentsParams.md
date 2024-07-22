# ListComponentsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Page limit (default 100) | [optional] 
**Offset** | Pointer to **int32** | Page offset (default 0) | [optional] 
**Q** | Pointer to **string** |  | [optional] 

## Methods

### NewListComponentsParams

`func NewListComponentsParams() *ListComponentsParams`

NewListComponentsParams instantiates a new ListComponentsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListComponentsParamsWithDefaults

`func NewListComponentsParamsWithDefaults() *ListComponentsParams`

NewListComponentsParamsWithDefaults instantiates a new ListComponentsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ListComponentsParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListComponentsParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListComponentsParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListComponentsParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ListComponentsParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListComponentsParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListComponentsParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ListComponentsParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetQ

`func (o *ListComponentsParams) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *ListComponentsParams) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *ListComponentsParams) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *ListComponentsParams) HasQ() bool`

HasQ returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


