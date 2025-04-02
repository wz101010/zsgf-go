# UnionIDSignInRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnionID** | **string** |  | 
**Platform** | **string** |  | 
**TwoFactorCode** | Pointer to **NullableString** | 如果启用双因素认证登录，则必填 | [optional] 

## Methods

### NewUnionIDSignInRequest

`func NewUnionIDSignInRequest(unionID string, platform string, ) *UnionIDSignInRequest`

NewUnionIDSignInRequest instantiates a new UnionIDSignInRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnionIDSignInRequestWithDefaults

`func NewUnionIDSignInRequestWithDefaults() *UnionIDSignInRequest`

NewUnionIDSignInRequestWithDefaults instantiates a new UnionIDSignInRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnionID

`func (o *UnionIDSignInRequest) GetUnionID() string`

GetUnionID returns the UnionID field if non-nil, zero value otherwise.

### GetUnionIDOk

`func (o *UnionIDSignInRequest) GetUnionIDOk() (*string, bool)`

GetUnionIDOk returns a tuple with the UnionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnionID

`func (o *UnionIDSignInRequest) SetUnionID(v string)`

SetUnionID sets UnionID field to given value.


### GetPlatform

`func (o *UnionIDSignInRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UnionIDSignInRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UnionIDSignInRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetTwoFactorCode

`func (o *UnionIDSignInRequest) GetTwoFactorCode() string`

GetTwoFactorCode returns the TwoFactorCode field if non-nil, zero value otherwise.

### GetTwoFactorCodeOk

`func (o *UnionIDSignInRequest) GetTwoFactorCodeOk() (*string, bool)`

GetTwoFactorCodeOk returns a tuple with the TwoFactorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorCode

`func (o *UnionIDSignInRequest) SetTwoFactorCode(v string)`

SetTwoFactorCode sets TwoFactorCode field to given value.

### HasTwoFactorCode

`func (o *UnionIDSignInRequest) HasTwoFactorCode() bool`

HasTwoFactorCode returns a boolean if a field has been set.

### SetTwoFactorCodeNil

`func (o *UnionIDSignInRequest) SetTwoFactorCodeNil(b bool)`

 SetTwoFactorCodeNil sets the value for TwoFactorCode to be an explicit nil

### UnsetTwoFactorCode
`func (o *UnionIDSignInRequest) UnsetTwoFactorCode()`

UnsetTwoFactorCode ensures that no value is present for TwoFactorCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


