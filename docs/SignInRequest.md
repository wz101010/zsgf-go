# SignInRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | **string** |  | 
**Pwd** | **string** |  | 
**TwoFactorCode** | Pointer to **NullableString** | 如果启用双因素认证登录，则必填 | [optional] 

## Methods

### NewSignInRequest

`func NewSignInRequest(userName string, pwd string, ) *SignInRequest`

NewSignInRequest instantiates a new SignInRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignInRequestWithDefaults

`func NewSignInRequestWithDefaults() *SignInRequest`

NewSignInRequestWithDefaults instantiates a new SignInRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *SignInRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SignInRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SignInRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetPwd

`func (o *SignInRequest) GetPwd() string`

GetPwd returns the Pwd field if non-nil, zero value otherwise.

### GetPwdOk

`func (o *SignInRequest) GetPwdOk() (*string, bool)`

GetPwdOk returns a tuple with the Pwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwd

`func (o *SignInRequest) SetPwd(v string)`

SetPwd sets Pwd field to given value.


### GetTwoFactorCode

`func (o *SignInRequest) GetTwoFactorCode() string`

GetTwoFactorCode returns the TwoFactorCode field if non-nil, zero value otherwise.

### GetTwoFactorCodeOk

`func (o *SignInRequest) GetTwoFactorCodeOk() (*string, bool)`

GetTwoFactorCodeOk returns a tuple with the TwoFactorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorCode

`func (o *SignInRequest) SetTwoFactorCode(v string)`

SetTwoFactorCode sets TwoFactorCode field to given value.

### HasTwoFactorCode

`func (o *SignInRequest) HasTwoFactorCode() bool`

HasTwoFactorCode returns a boolean if a field has been set.

### SetTwoFactorCodeNil

`func (o *SignInRequest) SetTwoFactorCodeNil(b bool)`

 SetTwoFactorCodeNil sets the value for TwoFactorCode to be an explicit nil

### UnsetTwoFactorCode
`func (o *SignInRequest) UnsetTwoFactorCode()`

UnsetTwoFactorCode ensures that no value is present for TwoFactorCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


