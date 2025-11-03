# RecommendFriend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserID** | Pointer to **int64** | 用户ID | [optional] 
**NickName** | Pointer to **NullableString** | 昵称 | [optional] 
**Avatar** | Pointer to **NullableString** | 头像 | [optional] 
**Gender** | Pointer to **NullableString** | 性别 | [optional] 
**Age** | Pointer to **int32** | 年龄 | [optional] 
**InterestTags** | Pointer to **NullableString** | 兴趣标签 | [optional] 
**LocationName** | Pointer to **NullableString** | 位置名称 | [optional] 
**Latitude** | Pointer to **float64** | 纬度 | [optional] 
**Longitude** | Pointer to **float64** | 经度 | [optional] 
**Distance** | Pointer to **int64** | 距离 | [optional] 
**Biography** | Pointer to **NullableString** | 个人简介 | [optional] 
**Country** | Pointer to **NullableString** | 国家 | [optional] 
**State** | Pointer to **NullableString** | 省份 | [optional] 
**City** | Pointer to **NullableString** | 城市 | [optional] 
**District** | Pointer to **NullableString** | 区县 | [optional] 

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

### GetBiography

`func (o *RecommendFriend) GetBiography() string`

GetBiography returns the Biography field if non-nil, zero value otherwise.

### GetBiographyOk

`func (o *RecommendFriend) GetBiographyOk() (*string, bool)`

GetBiographyOk returns a tuple with the Biography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiography

`func (o *RecommendFriend) SetBiography(v string)`

SetBiography sets Biography field to given value.

### HasBiography

`func (o *RecommendFriend) HasBiography() bool`

HasBiography returns a boolean if a field has been set.

### SetBiographyNil

`func (o *RecommendFriend) SetBiographyNil(b bool)`

 SetBiographyNil sets the value for Biography to be an explicit nil

### UnsetBiography
`func (o *RecommendFriend) UnsetBiography()`

UnsetBiography ensures that no value is present for Biography, not even an explicit nil
### GetCountry

`func (o *RecommendFriend) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RecommendFriend) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RecommendFriend) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *RecommendFriend) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *RecommendFriend) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *RecommendFriend) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetState

`func (o *RecommendFriend) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RecommendFriend) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RecommendFriend) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RecommendFriend) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *RecommendFriend) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *RecommendFriend) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCity

`func (o *RecommendFriend) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *RecommendFriend) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *RecommendFriend) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *RecommendFriend) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *RecommendFriend) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *RecommendFriend) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetDistrict

`func (o *RecommendFriend) GetDistrict() string`

GetDistrict returns the District field if non-nil, zero value otherwise.

### GetDistrictOk

`func (o *RecommendFriend) GetDistrictOk() (*string, bool)`

GetDistrictOk returns a tuple with the District field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrict

`func (o *RecommendFriend) SetDistrict(v string)`

SetDistrict sets District field to given value.

### HasDistrict

`func (o *RecommendFriend) HasDistrict() bool`

HasDistrict returns a boolean if a field has been set.

### SetDistrictNil

`func (o *RecommendFriend) SetDistrictNil(b bool)`

 SetDistrictNil sets the value for District to be an explicit nil

### UnsetDistrict
`func (o *RecommendFriend) UnsetDistrict()`

UnsetDistrict ensures that no value is present for District, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


