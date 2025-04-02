# CreateOrderResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**OrderNo** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateOrderResult

`func NewCreateOrderResult() *CreateOrderResult`

NewCreateOrderResult instantiates a new CreateOrderResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderResultWithDefaults

`func NewCreateOrderResultWithDefaults() *CreateOrderResult`

NewCreateOrderResultWithDefaults instantiates a new CreateOrderResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateOrderResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateOrderResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateOrderResult) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateOrderResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderNo

`func (o *CreateOrderResult) GetOrderNo() string`

GetOrderNo returns the OrderNo field if non-nil, zero value otherwise.

### GetOrderNoOk

`func (o *CreateOrderResult) GetOrderNoOk() (*string, bool)`

GetOrderNoOk returns a tuple with the OrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNo

`func (o *CreateOrderResult) SetOrderNo(v string)`

SetOrderNo sets OrderNo field to given value.

### HasOrderNo

`func (o *CreateOrderResult) HasOrderNo() bool`

HasOrderNo returns a boolean if a field has been set.

### SetOrderNoNil

`func (o *CreateOrderResult) SetOrderNoNil(b bool)`

 SetOrderNoNil sets the value for OrderNo to be an explicit nil

### UnsetOrderNo
`func (o *CreateOrderResult) UnsetOrderNo()`

UnsetOrderNo ensures that no value is present for OrderNo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


