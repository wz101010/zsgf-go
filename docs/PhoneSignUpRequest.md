# PhoneSignUpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | **string** |  | 
**PhoneCode** | **string** | 手机验证码 | 
**Pwd** | **string** |  | 
**Email** | Pointer to **NullableString** | 邮箱 | [optional] 
**EmailCode** | Pointer to **NullableString** | 邮箱验证码（只有启用的邮箱验证码功能时，才需要传入） | [optional] 
**NickName** | Pointer to **NullableString** | 昵称 | [optional] 
**Avatar** | Pointer to **NullableString** | 头像 | [optional] 
**Data** | Pointer to **NullableString** | 自定义数据 | [optional] 

## Methods

### NewPhoneSignUpRequest

`func NewPhoneSignUpRequest(phone string, phoneCode string, pwd string, ) *PhoneSignUpRequest`

NewPhoneSignUpRequest instantiates a new PhoneSignUpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneSignUpRequestWithDefaults

`func NewPhoneSignUpRequestWithDefaults() *PhoneSignUpRequest`

NewPhoneSignUpRequestWithDefaults instantiates a new PhoneSignUpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *PhoneSignUpRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PhoneSignUpRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PhoneSignUpRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetPhoneCode

`func (o *PhoneSignUpRequest) GetPhoneCode() string`

GetPhoneCode returns the PhoneCode field if non-nil, zero value otherwise.

### GetPhoneCodeOk

`func (o *PhoneSignUpRequest) GetPhoneCodeOk() (*string, bool)`

GetPhoneCodeOk returns a tuple with the PhoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCode

`func (o *PhoneSignUpRequest) SetPhoneCode(v string)`

SetPhoneCode sets PhoneCode field to given value.


### GetPwd

`func (o *PhoneSignUpRequest) GetPwd() string`

GetPwd returns the Pwd field if non-nil, zero value otherwise.

### GetPwdOk

`func (o *PhoneSignUpRequest) GetPwdOk() (*string, bool)`

GetPwdOk returns a tuple with the Pwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwd

`func (o *PhoneSignUpRequest) SetPwd(v string)`

SetPwd sets Pwd field to given value.


### GetEmail

`func (o *PhoneSignUpRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PhoneSignUpRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PhoneSignUpRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PhoneSignUpRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PhoneSignUpRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PhoneSignUpRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailCode

`func (o *PhoneSignUpRequest) GetEmailCode() string`

GetEmailCode returns the EmailCode field if non-nil, zero value otherwise.

### GetEmailCodeOk

`func (o *PhoneSignUpRequest) GetEmailCodeOk() (*string, bool)`

GetEmailCodeOk returns a tuple with the EmailCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCode

`func (o *PhoneSignUpRequest) SetEmailCode(v string)`

SetEmailCode sets EmailCode field to given value.

### HasEmailCode

`func (o *PhoneSignUpRequest) HasEmailCode() bool`

HasEmailCode returns a boolean if a field has been set.

### SetEmailCodeNil

`func (o *PhoneSignUpRequest) SetEmailCodeNil(b bool)`

 SetEmailCodeNil sets the value for EmailCode to be an explicit nil

### UnsetEmailCode
`func (o *PhoneSignUpRequest) UnsetEmailCode()`

UnsetEmailCode ensures that no value is present for EmailCode, not even an explicit nil
### GetNickName

`func (o *PhoneSignUpRequest) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *PhoneSignUpRequest) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *PhoneSignUpRequest) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *PhoneSignUpRequest) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### SetNickNameNil

`func (o *PhoneSignUpRequest) SetNickNameNil(b bool)`

 SetNickNameNil sets the value for NickName to be an explicit nil

### UnsetNickName
`func (o *PhoneSignUpRequest) UnsetNickName()`

UnsetNickName ensures that no value is present for NickName, not even an explicit nil
### GetAvatar

`func (o *PhoneSignUpRequest) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *PhoneSignUpRequest) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *PhoneSignUpRequest) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *PhoneSignUpRequest) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *PhoneSignUpRequest) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *PhoneSignUpRequest) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetData

`func (o *PhoneSignUpRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PhoneSignUpRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PhoneSignUpRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *PhoneSignUpRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *PhoneSignUpRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *PhoneSignUpRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


