# StorageListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Take** | Pointer to **int32** |  | [optional] 
**Skip** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 
**Data** | Pointer to **interface{}** |  | [optional] 
**Explain** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewStorageListResult

`func NewStorageListResult() *StorageListResult`

NewStorageListResult instantiates a new StorageListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageListResultWithDefaults

`func NewStorageListResultWithDefaults() *StorageListResult`

NewStorageListResultWithDefaults instantiates a new StorageListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTake

`func (o *StorageListResult) GetTake() int32`

GetTake returns the Take field if non-nil, zero value otherwise.

### GetTakeOk

`func (o *StorageListResult) GetTakeOk() (*int32, bool)`

GetTakeOk returns a tuple with the Take field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTake

`func (o *StorageListResult) SetTake(v int32)`

SetTake sets Take field to given value.

### HasTake

`func (o *StorageListResult) HasTake() bool`

HasTake returns a boolean if a field has been set.

### GetSkip

`func (o *StorageListResult) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *StorageListResult) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *StorageListResult) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *StorageListResult) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetTotal

`func (o *StorageListResult) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *StorageListResult) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *StorageListResult) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *StorageListResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *StorageListResult) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StorageListResult) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StorageListResult) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *StorageListResult) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *StorageListResult) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *StorageListResult) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetExplain

`func (o *StorageListResult) GetExplain() interface{}`

GetExplain returns the Explain field if non-nil, zero value otherwise.

### GetExplainOk

`func (o *StorageListResult) GetExplainOk() (*interface{}, bool)`

GetExplainOk returns a tuple with the Explain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplain

`func (o *StorageListResult) SetExplain(v interface{})`

SetExplain sets Explain field to given value.

### HasExplain

`func (o *StorageListResult) HasExplain() bool`

HasExplain returns a boolean if a field has been set.

### SetExplainNil

`func (o *StorageListResult) SetExplainNil(b bool)`

 SetExplainNil sets the value for Explain to be an explicit nil

### UnsetExplain
`func (o *StorageListResult) UnsetExplain()`

UnsetExplain ensures that no value is present for Explain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


