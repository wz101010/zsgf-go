# ExecuteFunctionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionKey** | Pointer to **NullableString** | 函数标识 | [optional] 

## Methods

### NewExecuteFunctionRequest

`func NewExecuteFunctionRequest() *ExecuteFunctionRequest`

NewExecuteFunctionRequest instantiates a new ExecuteFunctionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteFunctionRequestWithDefaults

`func NewExecuteFunctionRequestWithDefaults() *ExecuteFunctionRequest`

NewExecuteFunctionRequestWithDefaults instantiates a new ExecuteFunctionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionKey

`func (o *ExecuteFunctionRequest) GetFunctionKey() string`

GetFunctionKey returns the FunctionKey field if non-nil, zero value otherwise.

### GetFunctionKeyOk

`func (o *ExecuteFunctionRequest) GetFunctionKeyOk() (*string, bool)`

GetFunctionKeyOk returns a tuple with the FunctionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionKey

`func (o *ExecuteFunctionRequest) SetFunctionKey(v string)`

SetFunctionKey sets FunctionKey field to given value.

### HasFunctionKey

`func (o *ExecuteFunctionRequest) HasFunctionKey() bool`

HasFunctionKey returns a boolean if a field has been set.

### SetFunctionKeyNil

`func (o *ExecuteFunctionRequest) SetFunctionKeyNil(b bool)`

 SetFunctionKeyNil sets the value for FunctionKey to be an explicit nil

### UnsetFunctionKey
`func (o *ExecuteFunctionRequest) UnsetFunctionKey()`

UnsetFunctionKey ensures that no value is present for FunctionKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


