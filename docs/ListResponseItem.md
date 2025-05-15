# ListResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**UserID** | Pointer to **int64** |  | [optional] 
**UserName** | Pointer to **NullableString** |  | [optional] 
**NickName** | Pointer to **NullableString** |  | [optional] 
**Avatar** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**Role** | Pointer to **NullableString** |  | [optional] 
**Permission** | Pointer to **NullableString** |  | [optional] 
**LastUpdate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListResponseItem

`func NewListResponseItem() *ListResponseItem`

NewListResponseItem instantiates a new ListResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResponseItemWithDefaults

`func NewListResponseItemWithDefaults() *ListResponseItem`

NewListResponseItemWithDefaults instantiates a new ListResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListResponseItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListResponseItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListResponseItem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListResponseItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserID

`func (o *ListResponseItem) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *ListResponseItem) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *ListResponseItem) SetUserID(v int64)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *ListResponseItem) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetUserName

`func (o *ListResponseItem) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ListResponseItem) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ListResponseItem) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ListResponseItem) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *ListResponseItem) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *ListResponseItem) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetNickName

`func (o *ListResponseItem) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *ListResponseItem) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *ListResponseItem) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *ListResponseItem) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### SetNickNameNil

`func (o *ListResponseItem) SetNickNameNil(b bool)`

 SetNickNameNil sets the value for NickName to be an explicit nil

### UnsetNickName
`func (o *ListResponseItem) UnsetNickName()`

UnsetNickName ensures that no value is present for NickName, not even an explicit nil
### GetAvatar

`func (o *ListResponseItem) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ListResponseItem) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ListResponseItem) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *ListResponseItem) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *ListResponseItem) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *ListResponseItem) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetEmail

`func (o *ListResponseItem) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ListResponseItem) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ListResponseItem) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ListResponseItem) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ListResponseItem) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ListResponseItem) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *ListResponseItem) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ListResponseItem) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ListResponseItem) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ListResponseItem) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *ListResponseItem) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *ListResponseItem) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetRole

`func (o *ListResponseItem) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ListResponseItem) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ListResponseItem) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ListResponseItem) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *ListResponseItem) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *ListResponseItem) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetPermission

`func (o *ListResponseItem) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ListResponseItem) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ListResponseItem) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *ListResponseItem) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### SetPermissionNil

`func (o *ListResponseItem) SetPermissionNil(b bool)`

 SetPermissionNil sets the value for Permission to be an explicit nil

### UnsetPermission
`func (o *ListResponseItem) UnsetPermission()`

UnsetPermission ensures that no value is present for Permission, not even an explicit nil
### GetLastUpdate

`func (o *ListResponseItem) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *ListResponseItem) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *ListResponseItem) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *ListResponseItem) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


