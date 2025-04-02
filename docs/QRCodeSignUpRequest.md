# QRCodeSignUpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignInKey** | Pointer to **int64** |  | [optional] 
**UnionID** | **string** |  | 
**NickName** | Pointer to **NullableString** | 昵称 | [optional] 
**Avatar** | Pointer to **NullableString** | 头像 | [optional] 
**Data** | Pointer to **NullableString** | 自定义数据 | [optional] 
**Email** | Pointer to **NullableString** | 邮箱 | [optional] 
**EmailCode** | Pointer to **NullableString** | 邮箱验证码（只有启用的邮箱验证码功能时，才需要传入） | [optional] 
**Phone** | Pointer to **NullableString** | 手机号 | [optional] 
**PhoneCode** | Pointer to **NullableString** | 手机验证码（只有启用的手机验证码功能时，才需要传入） | [optional] 

## Methods

### NewQRCodeSignUpRequest

`func NewQRCodeSignUpRequest(unionID string, ) *QRCodeSignUpRequest`

NewQRCodeSignUpRequest instantiates a new QRCodeSignUpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQRCodeSignUpRequestWithDefaults

`func NewQRCodeSignUpRequestWithDefaults() *QRCodeSignUpRequest`

NewQRCodeSignUpRequestWithDefaults instantiates a new QRCodeSignUpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignInKey

`func (o *QRCodeSignUpRequest) GetSignInKey() int64`

GetSignInKey returns the SignInKey field if non-nil, zero value otherwise.

### GetSignInKeyOk

`func (o *QRCodeSignUpRequest) GetSignInKeyOk() (*int64, bool)`

GetSignInKeyOk returns a tuple with the SignInKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInKey

`func (o *QRCodeSignUpRequest) SetSignInKey(v int64)`

SetSignInKey sets SignInKey field to given value.

### HasSignInKey

`func (o *QRCodeSignUpRequest) HasSignInKey() bool`

HasSignInKey returns a boolean if a field has been set.

### GetUnionID

`func (o *QRCodeSignUpRequest) GetUnionID() string`

GetUnionID returns the UnionID field if non-nil, zero value otherwise.

### GetUnionIDOk

`func (o *QRCodeSignUpRequest) GetUnionIDOk() (*string, bool)`

GetUnionIDOk returns a tuple with the UnionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnionID

`func (o *QRCodeSignUpRequest) SetUnionID(v string)`

SetUnionID sets UnionID field to given value.


### GetNickName

`func (o *QRCodeSignUpRequest) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *QRCodeSignUpRequest) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *QRCodeSignUpRequest) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *QRCodeSignUpRequest) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### SetNickNameNil

`func (o *QRCodeSignUpRequest) SetNickNameNil(b bool)`

 SetNickNameNil sets the value for NickName to be an explicit nil

### UnsetNickName
`func (o *QRCodeSignUpRequest) UnsetNickName()`

UnsetNickName ensures that no value is present for NickName, not even an explicit nil
### GetAvatar

`func (o *QRCodeSignUpRequest) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *QRCodeSignUpRequest) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *QRCodeSignUpRequest) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *QRCodeSignUpRequest) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *QRCodeSignUpRequest) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *QRCodeSignUpRequest) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetData

`func (o *QRCodeSignUpRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QRCodeSignUpRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QRCodeSignUpRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *QRCodeSignUpRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *QRCodeSignUpRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *QRCodeSignUpRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetEmail

`func (o *QRCodeSignUpRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *QRCodeSignUpRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *QRCodeSignUpRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *QRCodeSignUpRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *QRCodeSignUpRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *QRCodeSignUpRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailCode

`func (o *QRCodeSignUpRequest) GetEmailCode() string`

GetEmailCode returns the EmailCode field if non-nil, zero value otherwise.

### GetEmailCodeOk

`func (o *QRCodeSignUpRequest) GetEmailCodeOk() (*string, bool)`

GetEmailCodeOk returns a tuple with the EmailCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCode

`func (o *QRCodeSignUpRequest) SetEmailCode(v string)`

SetEmailCode sets EmailCode field to given value.

### HasEmailCode

`func (o *QRCodeSignUpRequest) HasEmailCode() bool`

HasEmailCode returns a boolean if a field has been set.

### SetEmailCodeNil

`func (o *QRCodeSignUpRequest) SetEmailCodeNil(b bool)`

 SetEmailCodeNil sets the value for EmailCode to be an explicit nil

### UnsetEmailCode
`func (o *QRCodeSignUpRequest) UnsetEmailCode()`

UnsetEmailCode ensures that no value is present for EmailCode, not even an explicit nil
### GetPhone

`func (o *QRCodeSignUpRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *QRCodeSignUpRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *QRCodeSignUpRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *QRCodeSignUpRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *QRCodeSignUpRequest) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *QRCodeSignUpRequest) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetPhoneCode

`func (o *QRCodeSignUpRequest) GetPhoneCode() string`

GetPhoneCode returns the PhoneCode field if non-nil, zero value otherwise.

### GetPhoneCodeOk

`func (o *QRCodeSignUpRequest) GetPhoneCodeOk() (*string, bool)`

GetPhoneCodeOk returns a tuple with the PhoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCode

`func (o *QRCodeSignUpRequest) SetPhoneCode(v string)`

SetPhoneCode sets PhoneCode field to given value.

### HasPhoneCode

`func (o *QRCodeSignUpRequest) HasPhoneCode() bool`

HasPhoneCode returns a boolean if a field has been set.

### SetPhoneCodeNil

`func (o *QRCodeSignUpRequest) SetPhoneCodeNil(b bool)`

 SetPhoneCodeNil sets the value for PhoneCode to be an explicit nil

### UnsetPhoneCode
`func (o *QRCodeSignUpRequest) UnsetPhoneCode()`

UnsetPhoneCode ensures that no value is present for PhoneCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


