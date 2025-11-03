# AppUserResetEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | 邮箱 | [optional] 
**Code** | **string** | 邮箱验证码 | 

## Methods

### NewAppUserResetEmailRequest

`func NewAppUserResetEmailRequest(code string, ) *AppUserResetEmailRequest`

NewAppUserResetEmailRequest instantiates a new AppUserResetEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUserResetEmailRequestWithDefaults

`func NewAppUserResetEmailRequestWithDefaults() *AppUserResetEmailRequest`

NewAppUserResetEmailRequestWithDefaults instantiates a new AppUserResetEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AppUserResetEmailRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AppUserResetEmailRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AppUserResetEmailRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AppUserResetEmailRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *AppUserResetEmailRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AppUserResetEmailRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCode

`func (o *AppUserResetEmailRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AppUserResetEmailRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AppUserResetEmailRequest) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


