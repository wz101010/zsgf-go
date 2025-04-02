# AppUserResetPwdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Code** | **string** |  | 
**Pwd** | **string** |  | 

## Methods

### NewAppUserResetPwdRequest

`func NewAppUserResetPwdRequest(code string, pwd string, ) *AppUserResetPwdRequest`

NewAppUserResetPwdRequest instantiates a new AppUserResetPwdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUserResetPwdRequestWithDefaults

`func NewAppUserResetPwdRequestWithDefaults() *AppUserResetPwdRequest`

NewAppUserResetPwdRequestWithDefaults instantiates a new AppUserResetPwdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *AppUserResetPwdRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AppUserResetPwdRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AppUserResetPwdRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AppUserResetPwdRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *AppUserResetPwdRequest) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *AppUserResetPwdRequest) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetEmail

`func (o *AppUserResetPwdRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AppUserResetPwdRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AppUserResetPwdRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AppUserResetPwdRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *AppUserResetPwdRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AppUserResetPwdRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCode

`func (o *AppUserResetPwdRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AppUserResetPwdRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AppUserResetPwdRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetPwd

`func (o *AppUserResetPwdRequest) GetPwd() string`

GetPwd returns the Pwd field if non-nil, zero value otherwise.

### GetPwdOk

`func (o *AppUserResetPwdRequest) GetPwdOk() (*string, bool)`

GetPwdOk returns a tuple with the Pwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwd

`func (o *AppUserResetPwdRequest) SetPwd(v string)`

SetPwd sets Pwd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


