# UserFollowersResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalFollowers** | Pointer to **int64** |  | [optional] 
**Data** | Pointer to [**[]FollowerModel**](FollowerModel.md) |  | [optional] 

## Methods

### NewUserFollowersResult

`func NewUserFollowersResult() *UserFollowersResult`

NewUserFollowersResult instantiates a new UserFollowersResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFollowersResultWithDefaults

`func NewUserFollowersResultWithDefaults() *UserFollowersResult`

NewUserFollowersResultWithDefaults instantiates a new UserFollowersResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalFollowers

`func (o *UserFollowersResult) GetTotalFollowers() int64`

GetTotalFollowers returns the TotalFollowers field if non-nil, zero value otherwise.

### GetTotalFollowersOk

`func (o *UserFollowersResult) GetTotalFollowersOk() (*int64, bool)`

GetTotalFollowersOk returns a tuple with the TotalFollowers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFollowers

`func (o *UserFollowersResult) SetTotalFollowers(v int64)`

SetTotalFollowers sets TotalFollowers field to given value.

### HasTotalFollowers

`func (o *UserFollowersResult) HasTotalFollowers() bool`

HasTotalFollowers returns a boolean if a field has been set.

### GetData

`func (o *UserFollowersResult) GetData() []FollowerModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserFollowersResult) GetDataOk() (*[]FollowerModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserFollowersResult) SetData(v []FollowerModel)`

SetData sets Data field to given value.

### HasData

`func (o *UserFollowersResult) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *UserFollowersResult) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *UserFollowersResult) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


