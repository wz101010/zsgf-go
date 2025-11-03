# AppUserResetPhoneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **NullableString** | 手机号码 | [optional] 
**Code** | **string** | 手机验证码 | 

## Methods

### NewAppUserResetPhoneRequest

`func NewAppUserResetPhoneRequest(code string, ) *AppUserResetPhoneRequest`

NewAppUserResetPhoneRequest instantiates a new AppUserResetPhoneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUserResetPhoneRequestWithDefaults

`func NewAppUserResetPhoneRequestWithDefaults() *AppUserResetPhoneRequest`

NewAppUserResetPhoneRequestWithDefaults instantiates a new AppUserResetPhoneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *AppUserResetPhoneRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AppUserResetPhoneRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AppUserResetPhoneRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AppUserResetPhoneRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *AppUserResetPhoneRequest) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *AppUserResetPhoneRequest) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetCode

`func (o *AppUserResetPhoneRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AppUserResetPhoneRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AppUserResetPhoneRequest) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


