# ProfileResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**UnionID** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**CreateDate** | Pointer to **time.Time** |  | [optional] 
**UserName** | Pointer to **NullableString** |  | [optional] 
**PhoneIsValid** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**EmailIsValid** | Pointer to **bool** |  | [optional] 
**LastUpdate** | Pointer to **time.Time** |  | [optional] 
**NickName** | Pointer to **NullableString** |  | [optional] 
**Avatar** | Pointer to **NullableString** |  | [optional] 
**Role** | Pointer to **NullableString** |  | [optional] 
**Permission** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProfileResult

`func NewProfileResult() *ProfileResult`

NewProfileResult instantiates a new ProfileResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileResultWithDefaults

`func NewProfileResultWithDefaults() *ProfileResult`

NewProfileResultWithDefaults instantiates a new ProfileResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProfileResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfileResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfileResult) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProfileResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUnionID

`func (o *ProfileResult) GetUnionID() string`

GetUnionID returns the UnionID field if non-nil, zero value otherwise.

### GetUnionIDOk

`func (o *ProfileResult) GetUnionIDOk() (*string, bool)`

GetUnionIDOk returns a tuple with the UnionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnionID

`func (o *ProfileResult) SetUnionID(v string)`

SetUnionID sets UnionID field to given value.

### HasUnionID

`func (o *ProfileResult) HasUnionID() bool`

HasUnionID returns a boolean if a field has been set.

### SetUnionIDNil

`func (o *ProfileResult) SetUnionIDNil(b bool)`

 SetUnionIDNil sets the value for UnionID to be an explicit nil

### UnsetUnionID
`func (o *ProfileResult) UnsetUnionID()`

UnsetUnionID ensures that no value is present for UnionID, not even an explicit nil
### GetPhone

`func (o *ProfileResult) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ProfileResult) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ProfileResult) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ProfileResult) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *ProfileResult) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *ProfileResult) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetCreateDate

`func (o *ProfileResult) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *ProfileResult) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *ProfileResult) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *ProfileResult) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetUserName

`func (o *ProfileResult) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ProfileResult) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ProfileResult) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ProfileResult) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *ProfileResult) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *ProfileResult) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetPhoneIsValid

`func (o *ProfileResult) GetPhoneIsValid() bool`

GetPhoneIsValid returns the PhoneIsValid field if non-nil, zero value otherwise.

### GetPhoneIsValidOk

`func (o *ProfileResult) GetPhoneIsValidOk() (*bool, bool)`

GetPhoneIsValidOk returns a tuple with the PhoneIsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneIsValid

`func (o *ProfileResult) SetPhoneIsValid(v bool)`

SetPhoneIsValid sets PhoneIsValid field to given value.

### HasPhoneIsValid

`func (o *ProfileResult) HasPhoneIsValid() bool`

HasPhoneIsValid returns a boolean if a field has been set.

### GetData

`func (o *ProfileResult) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProfileResult) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProfileResult) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ProfileResult) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ProfileResult) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ProfileResult) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetEmail

`func (o *ProfileResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ProfileResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ProfileResult) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ProfileResult) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ProfileResult) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ProfileResult) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailIsValid

`func (o *ProfileResult) GetEmailIsValid() bool`

GetEmailIsValid returns the EmailIsValid field if non-nil, zero value otherwise.

### GetEmailIsValidOk

`func (o *ProfileResult) GetEmailIsValidOk() (*bool, bool)`

GetEmailIsValidOk returns a tuple with the EmailIsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailIsValid

`func (o *ProfileResult) SetEmailIsValid(v bool)`

SetEmailIsValid sets EmailIsValid field to given value.

### HasEmailIsValid

`func (o *ProfileResult) HasEmailIsValid() bool`

HasEmailIsValid returns a boolean if a field has been set.

### GetLastUpdate

`func (o *ProfileResult) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *ProfileResult) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *ProfileResult) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *ProfileResult) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetNickName

`func (o *ProfileResult) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *ProfileResult) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *ProfileResult) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *ProfileResult) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### SetNickNameNil

`func (o *ProfileResult) SetNickNameNil(b bool)`

 SetNickNameNil sets the value for NickName to be an explicit nil

### UnsetNickName
`func (o *ProfileResult) UnsetNickName()`

UnsetNickName ensures that no value is present for NickName, not even an explicit nil
### GetAvatar

`func (o *ProfileResult) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ProfileResult) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ProfileResult) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *ProfileResult) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *ProfileResult) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *ProfileResult) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetRole

`func (o *ProfileResult) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ProfileResult) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ProfileResult) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ProfileResult) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *ProfileResult) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *ProfileResult) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetPermission

`func (o *ProfileResult) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ProfileResult) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ProfileResult) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *ProfileResult) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### SetPermissionNil

`func (o *ProfileResult) SetPermissionNil(b bool)`

 SetPermissionNil sets the value for Permission to be an explicit nil

### UnsetPermission
`func (o *ProfileResult) UnsetPermission()`

UnsetPermission ensures that no value is present for Permission, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


