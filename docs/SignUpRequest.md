# SignUpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | **string** | 登录账号，4~18位 | 
**Pwd** | **string** | 登录密码，6~32位 | 
**NickName** | Pointer to **NullableString** | 昵称 | [optional] 
**Avatar** | Pointer to **NullableString** | 头像 | [optional] 
**Data** | Pointer to **NullableString** | 自定义数据 | [optional] 
**Email** | Pointer to **NullableString** | 邮箱 | [optional] 
**EmailCode** | Pointer to **NullableString** | 邮箱验证码（只有启用的邮箱验证码功能时，才需要传入） | [optional] 
**Phone** | Pointer to **NullableString** | 手机号 | [optional] 
**PhoneCode** | Pointer to **NullableString** | 手机验证码（只有启用的手机验证码功能时，才需要传入） | [optional] 

## Methods

### NewSignUpRequest

`func NewSignUpRequest(userName string, pwd string, ) *SignUpRequest`

NewSignUpRequest instantiates a new SignUpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignUpRequestWithDefaults

`func NewSignUpRequestWithDefaults() *SignUpRequest`

NewSignUpRequestWithDefaults instantiates a new SignUpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *SignUpRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SignUpRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SignUpRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetPwd

`func (o *SignUpRequest) GetPwd() string`

GetPwd returns the Pwd field if non-nil, zero value otherwise.

### GetPwdOk

`func (o *SignUpRequest) GetPwdOk() (*string, bool)`

GetPwdOk returns a tuple with the Pwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwd

`func (o *SignUpRequest) SetPwd(v string)`

SetPwd sets Pwd field to given value.


### GetNickName

`func (o *SignUpRequest) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *SignUpRequest) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *SignUpRequest) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *SignUpRequest) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### SetNickNameNil

`func (o *SignUpRequest) SetNickNameNil(b bool)`

 SetNickNameNil sets the value for NickName to be an explicit nil

### UnsetNickName
`func (o *SignUpRequest) UnsetNickName()`

UnsetNickName ensures that no value is present for NickName, not even an explicit nil
### GetAvatar

`func (o *SignUpRequest) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *SignUpRequest) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *SignUpRequest) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *SignUpRequest) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *SignUpRequest) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *SignUpRequest) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetData

`func (o *SignUpRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SignUpRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SignUpRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SignUpRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *SignUpRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SignUpRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetEmail

`func (o *SignUpRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SignUpRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SignUpRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SignUpRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *SignUpRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SignUpRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailCode

`func (o *SignUpRequest) GetEmailCode() string`

GetEmailCode returns the EmailCode field if non-nil, zero value otherwise.

### GetEmailCodeOk

`func (o *SignUpRequest) GetEmailCodeOk() (*string, bool)`

GetEmailCodeOk returns a tuple with the EmailCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCode

`func (o *SignUpRequest) SetEmailCode(v string)`

SetEmailCode sets EmailCode field to given value.

### HasEmailCode

`func (o *SignUpRequest) HasEmailCode() bool`

HasEmailCode returns a boolean if a field has been set.

### SetEmailCodeNil

`func (o *SignUpRequest) SetEmailCodeNil(b bool)`

 SetEmailCodeNil sets the value for EmailCode to be an explicit nil

### UnsetEmailCode
`func (o *SignUpRequest) UnsetEmailCode()`

UnsetEmailCode ensures that no value is present for EmailCode, not even an explicit nil
### GetPhone

`func (o *SignUpRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SignUpRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SignUpRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SignUpRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *SignUpRequest) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *SignUpRequest) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetPhoneCode

`func (o *SignUpRequest) GetPhoneCode() string`

GetPhoneCode returns the PhoneCode field if non-nil, zero value otherwise.

### GetPhoneCodeOk

`func (o *SignUpRequest) GetPhoneCodeOk() (*string, bool)`

GetPhoneCodeOk returns a tuple with the PhoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCode

`func (o *SignUpRequest) SetPhoneCode(v string)`

SetPhoneCode sets PhoneCode field to given value.

### HasPhoneCode

`func (o *SignUpRequest) HasPhoneCode() bool`

HasPhoneCode returns a boolean if a field has been set.

### SetPhoneCodeNil

`func (o *SignUpRequest) SetPhoneCodeNil(b bool)`

 SetPhoneCodeNil sets the value for PhoneCode to be an explicit nil

### UnsetPhoneCode
`func (o *SignUpRequest) UnsetPhoneCode()`

UnsetPhoneCode ensures that no value is present for PhoneCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


