# GetRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**RuleDetailResponseData**](RuleDetailResponseData.md) |  | 
**Params** | **map[string]interface{}** | Input parameters | 

## Methods

### NewGetRuleResponse

`func NewGetRuleResponse(data RuleDetailResponseData, params map[string]interface{}, ) *GetRuleResponse`

NewGetRuleResponse instantiates a new GetRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRuleResponseWithDefaults

`func NewGetRuleResponseWithDefaults() *GetRuleResponse`

NewGetRuleResponseWithDefaults instantiates a new GetRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetRuleResponse) GetData() RuleDetailResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRuleResponse) GetDataOk() (*RuleDetailResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRuleResponse) SetData(v RuleDetailResponseData)`

SetData sets Data field to given value.


### GetParams

`func (o *GetRuleResponse) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *GetRuleResponse) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *GetRuleResponse) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


