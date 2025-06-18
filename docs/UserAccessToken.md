# UserAccessToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 用户令牌的唯一标识符。 | [optional] 
**UserID** | Pointer to **int64** | 与令牌关联的用户ID。 | [optional] 
**Enable** | Pointer to **bool** | 指示令牌是否处于启用状态。 | [optional] 
**Permissions** | Pointer to **NullableString** | 令牌拥有的权限列表，多个权限以逗号分隔。 | [optional] 
**Title** | Pointer to **NullableString** | 令牌的标题或名称，用于标识令牌。 | [optional] 
**AccessToken** | Pointer to **NullableString** | 访问令牌的具体值，用于身份验证。 | [optional] 
**Tags** | Pointer to **NullableString** | 用于分类或标记令牌的标签。 | [optional] 
**Description** | Pointer to **NullableString** | 令牌的详细描述信息。 | [optional] 
**ExpireTime** | Pointer to **time.Time** | 令牌的过期时间，超过该时间令牌将失效。 | [optional] 
**CreateDate** | Pointer to **time.Time** | 令牌的创建日期，默认为当前时间。 | [optional] 
**LastUpdate** | Pointer to **time.Time** | 令牌的最后更新日期，默认为当前时间。 | [optional] 

## Methods

### NewUserAccessToken

`func NewUserAccessToken() *UserAccessToken`

NewUserAccessToken instantiates a new UserAccessToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccessTokenWithDefaults

`func NewUserAccessTokenWithDefaults() *UserAccessToken`

NewUserAccessTokenWithDefaults instantiates a new UserAccessToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserAccessToken) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserAccessToken) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserAccessToken) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserAccessToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserID

`func (o *UserAccessToken) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *UserAccessToken) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *UserAccessToken) SetUserID(v int64)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *UserAccessToken) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetEnable

`func (o *UserAccessToken) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *UserAccessToken) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *UserAccessToken) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *UserAccessToken) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetPermissions

`func (o *UserAccessToken) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UserAccessToken) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UserAccessToken) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UserAccessToken) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *UserAccessToken) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *UserAccessToken) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetTitle

`func (o *UserAccessToken) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UserAccessToken) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UserAccessToken) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UserAccessToken) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *UserAccessToken) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UserAccessToken) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAccessToken

`func (o *UserAccessToken) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *UserAccessToken) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *UserAccessToken) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *UserAccessToken) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *UserAccessToken) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *UserAccessToken) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetTags

`func (o *UserAccessToken) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UserAccessToken) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UserAccessToken) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UserAccessToken) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *UserAccessToken) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *UserAccessToken) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetDescription

`func (o *UserAccessToken) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserAccessToken) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserAccessToken) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserAccessToken) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UserAccessToken) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UserAccessToken) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExpireTime

`func (o *UserAccessToken) GetExpireTime() time.Time`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *UserAccessToken) GetExpireTimeOk() (*time.Time, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *UserAccessToken) SetExpireTime(v time.Time)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *UserAccessToken) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetCreateDate

`func (o *UserAccessToken) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *UserAccessToken) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *UserAccessToken) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *UserAccessToken) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *UserAccessToken) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *UserAccessToken) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *UserAccessToken) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *UserAccessToken) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


