# PhoneSignInRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | **string** |  | 
**VerifyCode** | **string** |  | 
**TwoFactorCode** | Pointer to **NullableString** | 如果启用双因素认证登录，则必填 | [optional] 

## Methods

### NewPhoneSignInRequest

`func NewPhoneSignInRequest(phone string, verifyCode string, ) *PhoneSignInRequest`

NewPhoneSignInRequest instantiates a new PhoneSignInRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneSignInRequestWithDefaults

`func NewPhoneSignInRequestWithDefaults() *PhoneSignInRequest`

NewPhoneSignInRequestWithDefaults instantiates a new PhoneSignInRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *PhoneSignInRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PhoneSignInRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PhoneSignInRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetVerifyCode

`func (o *PhoneSignInRequest) GetVerifyCode() string`

GetVerifyCode returns the VerifyCode field if non-nil, zero value otherwise.

### GetVerifyCodeOk

`func (o *PhoneSignInRequest) GetVerifyCodeOk() (*string, bool)`

GetVerifyCodeOk returns a tuple with the VerifyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCode

`func (o *PhoneSignInRequest) SetVerifyCode(v string)`

SetVerifyCode sets VerifyCode field to given value.


### GetTwoFactorCode

`func (o *PhoneSignInRequest) GetTwoFactorCode() string`

GetTwoFactorCode returns the TwoFactorCode field if non-nil, zero value otherwise.

### GetTwoFactorCodeOk

`func (o *PhoneSignInRequest) GetTwoFactorCodeOk() (*string, bool)`

GetTwoFactorCodeOk returns a tuple with the TwoFactorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorCode

`func (o *PhoneSignInRequest) SetTwoFactorCode(v string)`

SetTwoFactorCode sets TwoFactorCode field to given value.

### HasTwoFactorCode

`func (o *PhoneSignInRequest) HasTwoFactorCode() bool`

HasTwoFactorCode returns a boolean if a field has been set.

### SetTwoFactorCodeNil

`func (o *PhoneSignInRequest) SetTwoFactorCodeNil(b bool)`

 SetTwoFactorCodeNil sets the value for TwoFactorCode to be an explicit nil

### UnsetTwoFactorCode
`func (o *PhoneSignInRequest) UnsetTwoFactorCode()`

UnsetTwoFactorCode ensures that no value is present for TwoFactorCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


