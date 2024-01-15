# PaginationFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **NullableInt32** | Page limit (default 100) | [optional] 
**Offset** | Pointer to **NullableInt32** | Page offset (default 0) | [optional] 

## Methods

### NewPaginationFilter

`func NewPaginationFilter() *PaginationFilter`

NewPaginationFilter instantiates a new PaginationFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationFilterWithDefaults

`func NewPaginationFilterWithDefaults() *PaginationFilter`

NewPaginationFilterWithDefaults instantiates a new PaginationFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *PaginationFilter) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PaginationFilter) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PaginationFilter) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PaginationFilter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *PaginationFilter) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *PaginationFilter) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetOffset

`func (o *PaginationFilter) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PaginationFilter) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PaginationFilter) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PaginationFilter) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *PaginationFilter) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *PaginationFilter) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


