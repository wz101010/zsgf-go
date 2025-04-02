# OrderListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]Order**](Order.md) |  | [optional] 

## Methods

### NewOrderListResult

`func NewOrderListResult() *OrderListResult`

NewOrderListResult instantiates a new OrderListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderListResultWithDefaults

`func NewOrderListResultWithDefaults() *OrderListResult`

NewOrderListResultWithDefaults instantiates a new OrderListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *OrderListResult) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OrderListResult) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OrderListResult) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *OrderListResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *OrderListResult) GetData() []Order`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrderListResult) GetDataOk() (*[]Order, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrderListResult) SetData(v []Order)`

SetData sets Data field to given value.

### HasData

`func (o *OrderListResult) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *OrderListResult) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *OrderListResult) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


