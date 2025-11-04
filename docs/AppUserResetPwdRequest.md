# AppUserResetPwdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pwd** | **string** | 新的密码 | 
**Oldpwd** | **string** | 旧的密码 | 

## Methods

### NewAppUserResetPwdRequest

`func NewAppUserResetPwdRequest(pwd string, oldpwd string, ) *AppUserResetPwdRequest`

NewAppUserResetPwdRequest instantiates a new AppUserResetPwdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUserResetPwdRequestWithDefaults

`func NewAppUserResetPwdRequestWithDefaults() *AppUserResetPwdRequest`

NewAppUserResetPwdRequestWithDefaults instantiates a new AppUserResetPwdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetOldpwd

`func (o *AppUserResetPwdRequest) GetOldpwd() string`

GetOldpwd returns the Oldpwd field if non-nil, zero value otherwise.

### GetOldpwdOk

`func (o *AppUserResetPwdRequest) GetOldpwdOk() (*string, bool)`

GetOldpwdOk returns a tuple with the Oldpwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldpwd

`func (o *AppUserResetPwdRequest) SetOldpwd(v string)`

SetOldpwd sets Oldpwd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


