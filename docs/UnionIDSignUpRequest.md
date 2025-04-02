# UnionIDSignUpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | Pointer to **NullableString** | 登录账号，可空 | [optional] 
**UnionID** | **string** | UnionID | 
**Platform** | **string** | 平台标识 | 
**Pwd** | **string** | 密码，6~32位 | 
**NickName** | Pointer to **NullableString** | 昵称 | [optional] 
**Avatar** | Pointer to **NullableString** | 头像 | [optional] 
**Data** | Pointer to **NullableString** | 自定义数据 | [optional] 
**Email** | Pointer to **NullableString** | 邮箱 | [optional] 
**EmailCode** | Pointer to **NullableString** | 邮箱验证码（只有启用的邮箱验证码功能时，才需要传入） | [optional] 
**Phone** | Pointer to **NullableString** | 手机号 | [optional] 
**PhoneCode** | Pointer to **NullableString** | 手机验证码（只有启用的手机验证码功能时，才需要传入） | [optional] 

## Methods

### NewUnionIDSignUpRequest

`func NewUnionIDSignUpRequest(unionID string, platform string, pwd string, ) *UnionIDSignUpRequest`

NewUnionIDSignUpRequest instantiates a new UnionIDSignUpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnionIDSignUpRequestWithDefaults

`func NewUnionIDSignUpRequestWithDefaults() *UnionIDSignUpRequest`

NewUnionIDSignUpRequestWithDefaults instantiates a new UnionIDSignUpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *UnionIDSignUpRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UnionIDSignUpRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UnionIDSignUpRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UnionIDSignUpRequest) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *UnionIDSignUpRequest) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *UnionIDSignUpRequest) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetUnionID

`func (o *UnionIDSignUpRequest) GetUnionID() string`

GetUnionID returns the UnionID field if non-nil, zero value otherwise.

### GetUnionIDOk

`func (o *UnionIDSignUpRequest) GetUnionIDOk() (*string, bool)`

GetUnionIDOk returns a tuple with the UnionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnionID

`func (o *UnionIDSignUpRequest) SetUnionID(v string)`

SetUnionID sets UnionID field to given value.


### GetPlatform

`func (o *UnionIDSignUpRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UnionIDSignUpRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UnionIDSignUpRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetPwd

`func (o *UnionIDSignUpRequest) GetPwd() string`

GetPwd returns the Pwd field if non-nil, zero value otherwise.

### GetPwdOk

`func (o *UnionIDSignUpRequest) GetPwdOk() (*string, bool)`

GetPwdOk returns a tuple with the Pwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwd

`func (o *UnionIDSignUpRequest) SetPwd(v string)`

SetPwd sets Pwd field to given value.


### GetNickName

`func (o *UnionIDSignUpRequest) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *UnionIDSignUpRequest) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *UnionIDSignUpRequest) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *UnionIDSignUpRequest) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### SetNickNameNil

`func (o *UnionIDSignUpRequest) SetNickNameNil(b bool)`

 SetNickNameNil sets the value for NickName to be an explicit nil

### UnsetNickName
`func (o *UnionIDSignUpRequest) UnsetNickName()`

UnsetNickName ensures that no value is present for NickName, not even an explicit nil
### GetAvatar

`func (o *UnionIDSignUpRequest) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UnionIDSignUpRequest) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UnionIDSignUpRequest) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UnionIDSignUpRequest) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *UnionIDSignUpRequest) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *UnionIDSignUpRequest) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetData

`func (o *UnionIDSignUpRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UnionIDSignUpRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UnionIDSignUpRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *UnionIDSignUpRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *UnionIDSignUpRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *UnionIDSignUpRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetEmail

`func (o *UnionIDSignUpRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnionIDSignUpRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnionIDSignUpRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnionIDSignUpRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UnionIDSignUpRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UnionIDSignUpRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailCode

`func (o *UnionIDSignUpRequest) GetEmailCode() string`

GetEmailCode returns the EmailCode field if non-nil, zero value otherwise.

### GetEmailCodeOk

`func (o *UnionIDSignUpRequest) GetEmailCodeOk() (*string, bool)`

GetEmailCodeOk returns a tuple with the EmailCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCode

`func (o *UnionIDSignUpRequest) SetEmailCode(v string)`

SetEmailCode sets EmailCode field to given value.

### HasEmailCode

`func (o *UnionIDSignUpRequest) HasEmailCode() bool`

HasEmailCode returns a boolean if a field has been set.

### SetEmailCodeNil

`func (o *UnionIDSignUpRequest) SetEmailCodeNil(b bool)`

 SetEmailCodeNil sets the value for EmailCode to be an explicit nil

### UnsetEmailCode
`func (o *UnionIDSignUpRequest) UnsetEmailCode()`

UnsetEmailCode ensures that no value is present for EmailCode, not even an explicit nil
### GetPhone

`func (o *UnionIDSignUpRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UnionIDSignUpRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UnionIDSignUpRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UnionIDSignUpRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *UnionIDSignUpRequest) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *UnionIDSignUpRequest) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetPhoneCode

`func (o *UnionIDSignUpRequest) GetPhoneCode() string`

GetPhoneCode returns the PhoneCode field if non-nil, zero value otherwise.

### GetPhoneCodeOk

`func (o *UnionIDSignUpRequest) GetPhoneCodeOk() (*string, bool)`

GetPhoneCodeOk returns a tuple with the PhoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCode

`func (o *UnionIDSignUpRequest) SetPhoneCode(v string)`

SetPhoneCode sets PhoneCode field to given value.

### HasPhoneCode

`func (o *UnionIDSignUpRequest) HasPhoneCode() bool`

HasPhoneCode returns a boolean if a field has been set.

### SetPhoneCodeNil

`func (o *UnionIDSignUpRequest) SetPhoneCodeNil(b bool)`

 SetPhoneCodeNil sets the value for PhoneCode to be an explicit nil

### UnsetPhoneCode
`func (o *UnionIDSignUpRequest) UnsetPhoneCode()`

UnsetPhoneCode ensures that no value is present for PhoneCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


