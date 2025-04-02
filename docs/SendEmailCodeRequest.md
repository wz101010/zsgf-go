# SendEmailCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Type** | **string** | 用途。可选值：signup/forgetpwd | 

## Methods

### NewSendEmailCodeRequest

`func NewSendEmailCodeRequest(email string, type_ string, ) *SendEmailCodeRequest`

NewSendEmailCodeRequest instantiates a new SendEmailCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEmailCodeRequestWithDefaults

`func NewSendEmailCodeRequestWithDefaults() *SendEmailCodeRequest`

NewSendEmailCodeRequestWithDefaults instantiates a new SendEmailCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SendEmailCodeRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SendEmailCodeRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SendEmailCodeRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetType

`func (o *SendEmailCodeRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SendEmailCodeRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SendEmailCodeRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


