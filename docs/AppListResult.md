# AppListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]App**](App.md) |  | [optional] 

## Methods

### NewAppListResult

`func NewAppListResult() *AppListResult`

NewAppListResult instantiates a new AppListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppListResultWithDefaults

`func NewAppListResultWithDefaults() *AppListResult`

NewAppListResultWithDefaults instantiates a new AppListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *AppListResult) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AppListResult) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AppListResult) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AppListResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *AppListResult) GetData() []App`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppListResult) GetDataOk() (*[]App, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppListResult) SetData(v []App)`

SetData sets Data field to given value.

### HasData

`func (o *AppListResult) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *AppListResult) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *AppListResult) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


