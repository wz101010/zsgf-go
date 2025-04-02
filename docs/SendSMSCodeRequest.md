# SendSMSCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | **string** | 手机号 | 
**Type** | **string** | 用途。可选值：signup/forgetpwd | 

## Methods

### NewSendSMSCodeRequest

`func NewSendSMSCodeRequest(phone string, type_ string, ) *SendSMSCodeRequest`

NewSendSMSCodeRequest instantiates a new SendSMSCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendSMSCodeRequestWithDefaults

`func NewSendSMSCodeRequestWithDefaults() *SendSMSCodeRequest`

NewSendSMSCodeRequestWithDefaults instantiates a new SendSMSCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *SendSMSCodeRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendSMSCodeRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendSMSCodeRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetType

`func (o *SendSMSCodeRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SendSMSCodeRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SendSMSCodeRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


