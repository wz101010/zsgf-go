# UserLogins

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 唯一标识符 | [optional] 
**UserID** | Pointer to **int64** | 关联的用户ID | [optional] 
**PlatformName** | Pointer to **NullableString** | 第三方平台名称 | [optional] 
**Platform** | Pointer to **NullableString** | 第三方平台 | [optional] 
**UnionID** | Pointer to **NullableString** | 第三方平台唯一标识 | [optional] 
**Avatar** | Pointer to **NullableString** | 用户头像 | [optional] 
**Data** | Pointer to **NullableString** | 扩展数据 | [optional] 
**Enable** | Pointer to **bool** | 启用 | [optional] 
**CreateDate** | Pointer to **time.Time** | 创建时间 | [optional] 
**LastUpdate** | Pointer to **time.Time** | 最后更新时间 | [optional] 

## Methods

### NewUserLogins

`func NewUserLogins() *UserLogins`

NewUserLogins instantiates a new UserLogins object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLoginsWithDefaults

`func NewUserLoginsWithDefaults() *UserLogins`

NewUserLoginsWithDefaults instantiates a new UserLogins object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserLogins) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserLogins) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserLogins) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserLogins) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserID

`func (o *UserLogins) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *UserLogins) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *UserLogins) SetUserID(v int64)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *UserLogins) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetPlatformName

`func (o *UserLogins) GetPlatformName() string`

GetPlatformName returns the PlatformName field if non-nil, zero value otherwise.

### GetPlatformNameOk

`func (o *UserLogins) GetPlatformNameOk() (*string, bool)`

GetPlatformNameOk returns a tuple with the PlatformName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformName

`func (o *UserLogins) SetPlatformName(v string)`

SetPlatformName sets PlatformName field to given value.

### HasPlatformName

`func (o *UserLogins) HasPlatformName() bool`

HasPlatformName returns a boolean if a field has been set.

### SetPlatformNameNil

`func (o *UserLogins) SetPlatformNameNil(b bool)`

 SetPlatformNameNil sets the value for PlatformName to be an explicit nil

### UnsetPlatformName
`func (o *UserLogins) UnsetPlatformName()`

UnsetPlatformName ensures that no value is present for PlatformName, not even an explicit nil
### GetPlatform

`func (o *UserLogins) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UserLogins) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UserLogins) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *UserLogins) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *UserLogins) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *UserLogins) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetUnionID

`func (o *UserLogins) GetUnionID() string`

GetUnionID returns the UnionID field if non-nil, zero value otherwise.

### GetUnionIDOk

`func (o *UserLogins) GetUnionIDOk() (*string, bool)`

GetUnionIDOk returns a tuple with the UnionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnionID

`func (o *UserLogins) SetUnionID(v string)`

SetUnionID sets UnionID field to given value.

### HasUnionID

`func (o *UserLogins) HasUnionID() bool`

HasUnionID returns a boolean if a field has been set.

### SetUnionIDNil

`func (o *UserLogins) SetUnionIDNil(b bool)`

 SetUnionIDNil sets the value for UnionID to be an explicit nil

### UnsetUnionID
`func (o *UserLogins) UnsetUnionID()`

UnsetUnionID ensures that no value is present for UnionID, not even an explicit nil
### GetAvatar

`func (o *UserLogins) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserLogins) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserLogins) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UserLogins) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *UserLogins) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *UserLogins) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetData

`func (o *UserLogins) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserLogins) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserLogins) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *UserLogins) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *UserLogins) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *UserLogins) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetEnable

`func (o *UserLogins) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *UserLogins) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *UserLogins) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *UserLogins) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetCreateDate

`func (o *UserLogins) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *UserLogins) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *UserLogins) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *UserLogins) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *UserLogins) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *UserLogins) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *UserLogins) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *UserLogins) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


