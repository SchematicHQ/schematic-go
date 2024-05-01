# ListEntityKeyDefinitionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]EntityKeyDefinitionResponseData**](EntityKeyDefinitionResponseData.md) | The returned resources | 
**Params** | [**ListEntityKeyDefinitionsParams**](ListEntityKeyDefinitionsParams.md) |  | 

## Methods

### NewListEntityKeyDefinitionsResponse

`func NewListEntityKeyDefinitionsResponse(data []EntityKeyDefinitionResponseData, params ListEntityKeyDefinitionsParams, ) *ListEntityKeyDefinitionsResponse`

NewListEntityKeyDefinitionsResponse instantiates a new ListEntityKeyDefinitionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEntityKeyDefinitionsResponseWithDefaults

`func NewListEntityKeyDefinitionsResponseWithDefaults() *ListEntityKeyDefinitionsResponse`

NewListEntityKeyDefinitionsResponseWithDefaults instantiates a new ListEntityKeyDefinitionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListEntityKeyDefinitionsResponse) GetData() []EntityKeyDefinitionResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListEntityKeyDefinitionsResponse) GetDataOk() (*[]EntityKeyDefinitionResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListEntityKeyDefinitionsResponse) SetData(v []EntityKeyDefinitionResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *ListEntityKeyDefinitionsResponse) GetParams() ListEntityKeyDefinitionsParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ListEntityKeyDefinitionsResponse) GetParamsOk() (*ListEntityKeyDefinitionsParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ListEntityKeyDefinitionsResponse) SetParams(v ListEntityKeyDefinitionsParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


