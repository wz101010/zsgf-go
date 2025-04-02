# Team

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | ID | [optional] 
**UserID** | Pointer to **int64** | 用户ID | [optional] 
**ChannelCode** | Pointer to **NullableString** | 渠道代码 | [optional] 
**ChannelAppID** | Pointer to **NullableString** | 渠道应用ID | [optional] 
**Role** | Pointer to **NullableString** | 角色 | [optional] 
**Permission** | Pointer to **NullableString** | 权限 | [optional] 
**Show** | Pointer to **bool** | 是否显示 | [optional] 
**ShowIndex** | Pointer to **int64** | 显示顺序 | [optional] 
**CreateDate** | Pointer to **time.Time** | 创建日期 | [optional] 
**LastUpdate** | Pointer to **time.Time** | 最后更新日期 | [optional] 

## Methods

### NewTeam

`func NewTeam() *Team`

NewTeam instantiates a new Team object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamWithDefaults

`func NewTeamWithDefaults() *Team`

NewTeamWithDefaults instantiates a new Team object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Team) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Team) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Team) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Team) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserID

`func (o *Team) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *Team) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *Team) SetUserID(v int64)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *Team) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetChannelCode

`func (o *Team) GetChannelCode() string`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *Team) GetChannelCodeOk() (*string, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *Team) SetChannelCode(v string)`

SetChannelCode sets ChannelCode field to given value.

### HasChannelCode

`func (o *Team) HasChannelCode() bool`

HasChannelCode returns a boolean if a field has been set.

### SetChannelCodeNil

`func (o *Team) SetChannelCodeNil(b bool)`

 SetChannelCodeNil sets the value for ChannelCode to be an explicit nil

### UnsetChannelCode
`func (o *Team) UnsetChannelCode()`

UnsetChannelCode ensures that no value is present for ChannelCode, not even an explicit nil
### GetChannelAppID

`func (o *Team) GetChannelAppID() string`

GetChannelAppID returns the ChannelAppID field if non-nil, zero value otherwise.

### GetChannelAppIDOk

`func (o *Team) GetChannelAppIDOk() (*string, bool)`

GetChannelAppIDOk returns a tuple with the ChannelAppID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelAppID

`func (o *Team) SetChannelAppID(v string)`

SetChannelAppID sets ChannelAppID field to given value.

### HasChannelAppID

`func (o *Team) HasChannelAppID() bool`

HasChannelAppID returns a boolean if a field has been set.

### SetChannelAppIDNil

`func (o *Team) SetChannelAppIDNil(b bool)`

 SetChannelAppIDNil sets the value for ChannelAppID to be an explicit nil

### UnsetChannelAppID
`func (o *Team) UnsetChannelAppID()`

UnsetChannelAppID ensures that no value is present for ChannelAppID, not even an explicit nil
### GetRole

`func (o *Team) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Team) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Team) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Team) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *Team) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *Team) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetPermission

`func (o *Team) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *Team) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *Team) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *Team) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### SetPermissionNil

`func (o *Team) SetPermissionNil(b bool)`

 SetPermissionNil sets the value for Permission to be an explicit nil

### UnsetPermission
`func (o *Team) UnsetPermission()`

UnsetPermission ensures that no value is present for Permission, not even an explicit nil
### GetShow

`func (o *Team) GetShow() bool`

GetShow returns the Show field if non-nil, zero value otherwise.

### GetShowOk

`func (o *Team) GetShowOk() (*bool, bool)`

GetShowOk returns a tuple with the Show field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShow

`func (o *Team) SetShow(v bool)`

SetShow sets Show field to given value.

### HasShow

`func (o *Team) HasShow() bool`

HasShow returns a boolean if a field has been set.

### GetShowIndex

`func (o *Team) GetShowIndex() int64`

GetShowIndex returns the ShowIndex field if non-nil, zero value otherwise.

### GetShowIndexOk

`func (o *Team) GetShowIndexOk() (*int64, bool)`

GetShowIndexOk returns a tuple with the ShowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIndex

`func (o *Team) SetShowIndex(v int64)`

SetShowIndex sets ShowIndex field to given value.

### HasShowIndex

`func (o *Team) HasShowIndex() bool`

HasShowIndex returns a boolean if a field has been set.

### GetCreateDate

`func (o *Team) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *Team) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *Team) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *Team) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *Team) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *Team) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *Team) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *Team) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


