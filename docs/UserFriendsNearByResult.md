# UserFriendsNearByResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalFollowers** | Pointer to **int64** |  | [optional] 
**Data** | Pointer to [**[]RecommendFriend**](RecommendFriend.md) |  | [optional] 

## Methods

### NewUserFriendsNearByResult

`func NewUserFriendsNearByResult() *UserFriendsNearByResult`

NewUserFriendsNearByResult instantiates a new UserFriendsNearByResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFriendsNearByResultWithDefaults

`func NewUserFriendsNearByResultWithDefaults() *UserFriendsNearByResult`

NewUserFriendsNearByResultWithDefaults instantiates a new UserFriendsNearByResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalFollowers

`func (o *UserFriendsNearByResult) GetTotalFollowers() int64`

GetTotalFollowers returns the TotalFollowers field if non-nil, zero value otherwise.

### GetTotalFollowersOk

`func (o *UserFriendsNearByResult) GetTotalFollowersOk() (*int64, bool)`

GetTotalFollowersOk returns a tuple with the TotalFollowers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFollowers

`func (o *UserFriendsNearByResult) SetTotalFollowers(v int64)`

SetTotalFollowers sets TotalFollowers field to given value.

### HasTotalFollowers

`func (o *UserFriendsNearByResult) HasTotalFollowers() bool`

HasTotalFollowers returns a boolean if a field has been set.

### GetData

`func (o *UserFriendsNearByResult) GetData() []RecommendFriend`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserFriendsNearByResult) GetDataOk() (*[]RecommendFriend, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserFriendsNearByResult) SetData(v []RecommendFriend)`

SetData sets Data field to given value.

### HasData

`func (o *UserFriendsNearByResult) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *UserFriendsNearByResult) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *UserFriendsNearByResult) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


