# ProjectListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]Project**](Project.md) |  | [optional] 

## Methods

### NewProjectListResult

`func NewProjectListResult() *ProjectListResult`

NewProjectListResult instantiates a new ProjectListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectListResultWithDefaults

`func NewProjectListResultWithDefaults() *ProjectListResult`

NewProjectListResultWithDefaults instantiates a new ProjectListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ProjectListResult) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProjectListResult) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProjectListResult) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ProjectListResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *ProjectListResult) GetData() []Project`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProjectListResult) GetDataOk() (*[]Project, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProjectListResult) SetData(v []Project)`

SetData sets Data field to given value.

### HasData

`func (o *ProjectListResult) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ProjectListResult) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ProjectListResult) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


