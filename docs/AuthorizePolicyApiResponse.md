# AuthorizePolicyApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | 状态码 | [optional] [default to 200]
**Data** | Pointer to [**AuthorizePolicy**](AuthorizePolicy.md) |  | [optional] 
**Error** | Pointer to **NullableString** | 错误信息 | [optional] [default to ""]

## Methods

### NewAuthorizePolicyApiResponse

`func NewAuthorizePolicyApiResponse() *AuthorizePolicyApiResponse`

NewAuthorizePolicyApiResponse instantiates a new AuthorizePolicyApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizePolicyApiResponseWithDefaults

`func NewAuthorizePolicyApiResponseWithDefaults() *AuthorizePolicyApiResponse`

NewAuthorizePolicyApiResponseWithDefaults instantiates a new AuthorizePolicyApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AuthorizePolicyApiResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AuthorizePolicyApiResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AuthorizePolicyApiResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *AuthorizePolicyApiResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *AuthorizePolicyApiResponse) GetData() AuthorizePolicy`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuthorizePolicyApiResponse) GetDataOk() (*AuthorizePolicy, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuthorizePolicyApiResponse) SetData(v AuthorizePolicy)`

SetData sets Data field to given value.

### HasData

`func (o *AuthorizePolicyApiResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *AuthorizePolicyApiResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AuthorizePolicyApiResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AuthorizePolicyApiResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AuthorizePolicyApiResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *AuthorizePolicyApiResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *AuthorizePolicyApiResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


