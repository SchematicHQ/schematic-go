# CountCustomersParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedToImport** | Pointer to **bool** |  | [optional] 
**Limit** | Pointer to **int32** | Page limit (default 100) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int32** | Page offset (default 0) | [optional] 
**Q** | Pointer to **string** |  | [optional] 

## Methods

### NewCountCustomersParams

`func NewCountCustomersParams() *CountCustomersParams`

NewCountCustomersParams instantiates a new CountCustomersParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountCustomersParamsWithDefaults

`func NewCountCustomersParamsWithDefaults() *CountCustomersParams`

NewCountCustomersParamsWithDefaults instantiates a new CountCustomersParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedToImport

`func (o *CountCustomersParams) GetFailedToImport() bool`

GetFailedToImport returns the FailedToImport field if non-nil, zero value otherwise.

### GetFailedToImportOk

`func (o *CountCustomersParams) GetFailedToImportOk() (*bool, bool)`

GetFailedToImportOk returns a tuple with the FailedToImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedToImport

`func (o *CountCustomersParams) SetFailedToImport(v bool)`

SetFailedToImport sets FailedToImport field to given value.

### HasFailedToImport

`func (o *CountCustomersParams) HasFailedToImport() bool`

HasFailedToImport returns a boolean if a field has been set.

### GetLimit

`func (o *CountCustomersParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CountCustomersParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CountCustomersParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CountCustomersParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetName

`func (o *CountCustomersParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CountCustomersParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CountCustomersParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CountCustomersParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOffset

`func (o *CountCustomersParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *CountCustomersParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *CountCustomersParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *CountCustomersParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetQ

`func (o *CountCustomersParams) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *CountCustomersParams) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *CountCustomersParams) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *CountCustomersParams) HasQ() bool`

HasQ returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


