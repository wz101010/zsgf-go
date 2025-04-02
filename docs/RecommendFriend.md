# RecommendFriend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserID** | Pointer to **int64** |  | [optional] 
**NickName** | Pointer to **NullableString** |  | [optional] 
**Avatar** | Pointer to **NullableString** |  | [optional] 
**Gender** | Pointer to **NullableString** |  | [optional] 
**Age** | Pointer to **int32** |  | [optional] 
**InterestTags** | Pointer to **NullableString** |  | [optional] 
**LocationName** | Pointer to **NullableString** |  | [optional] 
**Latitude** | Pointer to **float64** |  | [optional] 
**Longitude** | Pointer to **float64** |  | [optional] 
**Distance** | Pointer to **int64** |  | [optional] 

## Methods

### NewRecommendFriend

`func NewRecommendFriend() *RecommendFriend`

NewRecommendFriend instantiates a new RecommendFriend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendFriendWithDefaults

`func NewRecommendFriendWithDefaults() *RecommendFriend`

NewRecommendFriendWithDefaults instantiates a new RecommendFriend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserID

`func (o *RecommendFriend) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *RecommendFriend) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *RecommendFriend) SetUserID(v int64)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *RecommendFriend) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetNickName

`func (o *RecommendFriend) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *RecommendFriend) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *RecommendFriend) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *RecommendFriend) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### SetNickNameNil

`func (o *RecommendFriend) SetNickNameNil(b bool)`

 SetNickNameNil sets the value for NickName to be an explicit nil

### UnsetNickName
`func (o *RecommendFriend) UnsetNickName()`

UnsetNickName ensures that no value is present for NickName, not even an explicit nil
### GetAvatar

`func (o *RecommendFriend) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *RecommendFriend) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *RecommendFriend) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *RecommendFriend) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *RecommendFriend) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *RecommendFriend) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetGender

`func (o *RecommendFriend) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *RecommendFriend) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *RecommendFriend) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *RecommendFriend) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *RecommendFriend) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *RecommendFriend) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetAge

`func (o *RecommendFriend) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *RecommendFriend) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *RecommendFriend) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *RecommendFriend) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetInterestTags

`func (o *RecommendFriend) GetInterestTags() string`

GetInterestTags returns the InterestTags field if non-nil, zero value otherwise.

### GetInterestTagsOk

`func (o *RecommendFriend) GetInterestTagsOk() (*string, bool)`

GetInterestTagsOk returns a tuple with the InterestTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestTags

`func (o *RecommendFriend) SetInterestTags(v string)`

SetInterestTags sets InterestTags field to given value.

### HasInterestTags

`func (o *RecommendFriend) HasInterestTags() bool`

HasInterestTags returns a boolean if a field has been set.

### SetInterestTagsNil

`func (o *RecommendFriend) SetInterestTagsNil(b bool)`

 SetInterestTagsNil sets the value for InterestTags to be an explicit nil

### UnsetInterestTags
`func (o *RecommendFriend) UnsetInterestTags()`

UnsetInterestTags ensures that no value is present for InterestTags, not even an explicit nil
### GetLocationName

`func (o *RecommendFriend) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *RecommendFriend) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *RecommendFriend) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *RecommendFriend) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### SetLocationNameNil

`func (o *RecommendFriend) SetLocationNameNil(b bool)`

 SetLocationNameNil sets the value for LocationName to be an explicit nil

### UnsetLocationName
`func (o *RecommendFriend) UnsetLocationName()`

UnsetLocationName ensures that no value is present for LocationName, not even an explicit nil
### GetLatitude

`func (o *RecommendFriend) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *RecommendFriend) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *RecommendFriend) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *RecommendFriend) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *RecommendFriend) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *RecommendFriend) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *RecommendFriend) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *RecommendFriend) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetDistance

`func (o *RecommendFriend) GetDistance() int64`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *RecommendFriend) GetDistanceOk() (*int64, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *RecommendFriend) SetDistance(v int64)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *RecommendFriend) HasDistance() bool`

HasDistance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


