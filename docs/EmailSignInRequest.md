# EmailSignInRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**VerifyCode** | **string** |  | 
**TwoFactorCode** | Pointer to **NullableString** | 如果启用双因素认证登录，则必填 | [optional] 

## Methods

### NewEmailSignInRequest

`func NewEmailSignInRequest(email string, verifyCode string, ) *EmailSignInRequest`

NewEmailSignInRequest instantiates a new EmailSignInRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSignInRequestWithDefaults

`func NewEmailSignInRequestWithDefaults() *EmailSignInRequest`

NewEmailSignInRequestWithDefaults instantiates a new EmailSignInRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EmailSignInRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailSignInRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailSignInRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetVerifyCode

`func (o *EmailSignInRequest) GetVerifyCode() string`

GetVerifyCode returns the VerifyCode field if non-nil, zero value otherwise.

### GetVerifyCodeOk

`func (o *EmailSignInRequest) GetVerifyCodeOk() (*string, bool)`

GetVerifyCodeOk returns a tuple with the VerifyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCode

`func (o *EmailSignInRequest) SetVerifyCode(v string)`

SetVerifyCode sets VerifyCode field to given value.


### GetTwoFactorCode

`func (o *EmailSignInRequest) GetTwoFactorCode() string`

GetTwoFactorCode returns the TwoFactorCode field if non-nil, zero value otherwise.

### GetTwoFactorCodeOk

`func (o *EmailSignInRequest) GetTwoFactorCodeOk() (*string, bool)`

GetTwoFactorCodeOk returns a tuple with the TwoFactorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorCode

`func (o *EmailSignInRequest) SetTwoFactorCode(v string)`

SetTwoFactorCode sets TwoFactorCode field to given value.

### HasTwoFactorCode

`func (o *EmailSignInRequest) HasTwoFactorCode() bool`

HasTwoFactorCode returns a boolean if a field has been set.

### SetTwoFactorCodeNil

`func (o *EmailSignInRequest) SetTwoFactorCodeNil(b bool)`

 SetTwoFactorCodeNil sets the value for TwoFactorCode to be an explicit nil

### UnsetTwoFactorCode
`func (o *EmailSignInRequest) UnsetTwoFactorCode()`

UnsetTwoFactorCode ensures that no value is present for TwoFactorCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


