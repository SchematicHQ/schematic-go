# GetEntityTraitValuesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]EntityTraitValue**](EntityTraitValue.md) | The returned resources | 
**Params** | [**GetEntityTraitValuesParams**](GetEntityTraitValuesParams.md) |  | 

## Methods

### NewGetEntityTraitValuesResponse

`func NewGetEntityTraitValuesResponse(data []EntityTraitValue, params GetEntityTraitValuesParams, ) *GetEntityTraitValuesResponse`

NewGetEntityTraitValuesResponse instantiates a new GetEntityTraitValuesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntityTraitValuesResponseWithDefaults

`func NewGetEntityTraitValuesResponseWithDefaults() *GetEntityTraitValuesResponse`

NewGetEntityTraitValuesResponseWithDefaults instantiates a new GetEntityTraitValuesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetEntityTraitValuesResponse) GetData() []EntityTraitValue`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetEntityTraitValuesResponse) GetDataOk() (*[]EntityTraitValue, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetEntityTraitValuesResponse) SetData(v []EntityTraitValue)`

SetData sets Data field to given value.


### GetParams

`func (o *GetEntityTraitValuesResponse) GetParams() GetEntityTraitValuesParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *GetEntityTraitValuesResponse) GetParamsOk() (*GetEntityTraitValuesParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *GetEntityTraitValuesResponse) SetParams(v GetEntityTraitValuesParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


