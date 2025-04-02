# EmailSignUpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Pwd** | **string** |  | 
**EmailCode** | Pointer to **NullableString** | 邮箱验证码 | [optional] 
**Phone** | Pointer to **NullableString** | 手机号 | [optional] 
**PhoneCode** | Pointer to **NullableString** | 手机验证码（只有启用的手机验证码功能时，才需要传入） | [optional] 
**NickName** | Pointer to **NullableString** | 昵称 | [optional] 
**Avatar** | Pointer to **NullableString** | 头像 | [optional] 
**Data** | Pointer to **NullableString** | 自定义数据 | [optional] 

## Methods

### NewEmailSignUpRequest

`func NewEmailSignUpRequest(email string, pwd string, ) *EmailSignUpRequest`

NewEmailSignUpRequest instantiates a new EmailSignUpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSignUpRequestWithDefaults

`func NewEmailSignUpRequestWithDefaults() *EmailSignUpRequest`

NewEmailSignUpRequestWithDefaults instantiates a new EmailSignUpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EmailSignUpRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailSignUpRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailSignUpRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPwd

`func (o *EmailSignUpRequest) GetPwd() string`

GetPwd returns the Pwd field if non-nil, zero value otherwise.

### GetPwdOk

`func (o *EmailSignUpRequest) GetPwdOk() (*string, bool)`

GetPwdOk returns a tuple with the Pwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwd

`func (o *EmailSignUpRequest) SetPwd(v string)`

SetPwd sets Pwd field to given value.


### GetEmailCode

`func (o *EmailSignUpRequest) GetEmailCode() string`

GetEmailCode returns the EmailCode field if non-nil, zero value otherwise.

### GetEmailCodeOk

`func (o *EmailSignUpRequest) GetEmailCodeOk() (*string, bool)`

GetEmailCodeOk returns a tuple with the EmailCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCode

`func (o *EmailSignUpRequest) SetEmailCode(v string)`

SetEmailCode sets EmailCode field to given value.

### HasEmailCode

`func (o *EmailSignUpRequest) HasEmailCode() bool`

HasEmailCode returns a boolean if a field has been set.

### SetEmailCodeNil

`func (o *EmailSignUpRequest) SetEmailCodeNil(b bool)`

 SetEmailCodeNil sets the value for EmailCode to be an explicit nil

### UnsetEmailCode
`func (o *EmailSignUpRequest) UnsetEmailCode()`

UnsetEmailCode ensures that no value is present for EmailCode, not even an explicit nil
### GetPhone

`func (o *EmailSignUpRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *EmailSignUpRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *EmailSignUpRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *EmailSignUpRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *EmailSignUpRequest) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *EmailSignUpRequest) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetPhoneCode

`func (o *EmailSignUpRequest) GetPhoneCode() string`

GetPhoneCode returns the PhoneCode field if non-nil, zero value otherwise.

### GetPhoneCodeOk

`func (o *EmailSignUpRequest) GetPhoneCodeOk() (*string, bool)`

GetPhoneCodeOk returns a tuple with the PhoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCode

`func (o *EmailSignUpRequest) SetPhoneCode(v string)`

SetPhoneCode sets PhoneCode field to given value.

### HasPhoneCode

`func (o *EmailSignUpRequest) HasPhoneCode() bool`

HasPhoneCode returns a boolean if a field has been set.

### SetPhoneCodeNil

`func (o *EmailSignUpRequest) SetPhoneCodeNil(b bool)`

 SetPhoneCodeNil sets the value for PhoneCode to be an explicit nil

### UnsetPhoneCode
`func (o *EmailSignUpRequest) UnsetPhoneCode()`

UnsetPhoneCode ensures that no value is present for PhoneCode, not even an explicit nil
### GetNickName

`func (o *EmailSignUpRequest) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *EmailSignUpRequest) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *EmailSignUpRequest) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *EmailSignUpRequest) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### SetNickNameNil

`func (o *EmailSignUpRequest) SetNickNameNil(b bool)`

 SetNickNameNil sets the value for NickName to be an explicit nil

### UnsetNickName
`func (o *EmailSignUpRequest) UnsetNickName()`

UnsetNickName ensures that no value is present for NickName, not even an explicit nil
### GetAvatar

`func (o *EmailSignUpRequest) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *EmailSignUpRequest) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *EmailSignUpRequest) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *EmailSignUpRequest) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *EmailSignUpRequest) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *EmailSignUpRequest) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetData

`func (o *EmailSignUpRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EmailSignUpRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EmailSignUpRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *EmailSignUpRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *EmailSignUpRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *EmailSignUpRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


