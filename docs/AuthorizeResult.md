# AuthorizeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 
**AccessToken** | Pointer to **NullableString** |  | [optional] 
**TokenType** | Pointer to **NullableString** |  | [optional] 
**ExpiresIn** | Pointer to **int64** |  | [optional] 

## Methods

### NewAuthorizeResult

`func NewAuthorizeResult() *AuthorizeResult`

NewAuthorizeResult instantiates a new AuthorizeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeResultWithDefaults

`func NewAuthorizeResultWithDefaults() *AuthorizeResult`

NewAuthorizeResultWithDefaults instantiates a new AuthorizeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AuthorizeResult) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AuthorizeResult) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AuthorizeResult) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *AuthorizeResult) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetError

`func (o *AuthorizeResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AuthorizeResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AuthorizeResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AuthorizeResult) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *AuthorizeResult) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *AuthorizeResult) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetAccessToken

`func (o *AuthorizeResult) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AuthorizeResult) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AuthorizeResult) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AuthorizeResult) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *AuthorizeResult) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *AuthorizeResult) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetTokenType

`func (o *AuthorizeResult) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AuthorizeResult) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AuthorizeResult) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *AuthorizeResult) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### SetTokenTypeNil

`func (o *AuthorizeResult) SetTokenTypeNil(b bool)`

 SetTokenTypeNil sets the value for TokenType to be an explicit nil

### UnsetTokenType
`func (o *AuthorizeResult) UnsetTokenType()`

UnsetTokenType ensures that no value is present for TokenType, not even an explicit nil
### GetExpiresIn

`func (o *AuthorizeResult) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *AuthorizeResult) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *AuthorizeResult) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *AuthorizeResult) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


