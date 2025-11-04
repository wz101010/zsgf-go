# UserFollowersResultApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | 状态码 | [optional] [default to 200]
**Data** | Pointer to [**UserFollowersResult**](UserFollowersResult.md) |  | [optional] 
**Error** | Pointer to **NullableString** | 错误信息 | [optional] [default to ""]

## Methods

### NewUserFollowersResultApiResponse

`func NewUserFollowersResultApiResponse() *UserFollowersResultApiResponse`

NewUserFollowersResultApiResponse instantiates a new UserFollowersResultApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFollowersResultApiResponseWithDefaults

`func NewUserFollowersResultApiResponseWithDefaults() *UserFollowersResultApiResponse`

NewUserFollowersResultApiResponseWithDefaults instantiates a new UserFollowersResultApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UserFollowersResultApiResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UserFollowersResultApiResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UserFollowersResultApiResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *UserFollowersResultApiResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *UserFollowersResultApiResponse) GetData() UserFollowersResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserFollowersResultApiResponse) GetDataOk() (*UserFollowersResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserFollowersResultApiResponse) SetData(v UserFollowersResult)`

SetData sets Data field to given value.

### HasData

`func (o *UserFollowersResultApiResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *UserFollowersResultApiResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UserFollowersResultApiResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UserFollowersResultApiResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *UserFollowersResultApiResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *UserFollowersResultApiResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *UserFollowersResultApiResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


