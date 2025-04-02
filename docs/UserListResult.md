# UserListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Data** | Pointer to [**[]AppUserListResponse**](AppUserListResponse.md) |  | [optional] 

## Methods

### NewUserListResult

`func NewUserListResult() *UserListResult`

NewUserListResult instantiates a new UserListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserListResultWithDefaults

`func NewUserListResultWithDefaults() *UserListResult`

NewUserListResultWithDefaults instantiates a new UserListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *UserListResult) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UserListResult) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UserListResult) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UserListResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *UserListResult) GetData() []AppUserListResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserListResult) GetDataOk() (*[]AppUserListResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserListResult) SetData(v []AppUserListResponse)`

SetData sets Data field to given value.

### HasData

`func (o *UserListResult) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *UserListResult) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *UserListResult) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


