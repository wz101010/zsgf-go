# UserCommonInterestsResultApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | 状态码 | [optional] [default to 200]
**Data** | Pointer to [**UserCommonInterestsResult**](UserCommonInterestsResult.md) |  | [optional] 
**Error** | Pointer to **NullableString** | 错误信息 | [optional] [default to ""]

## Methods

### NewUserCommonInterestsResultApiResponse

`func NewUserCommonInterestsResultApiResponse() *UserCommonInterestsResultApiResponse`

NewUserCommonInterestsResultApiResponse instantiates a new UserCommonInterestsResultApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCommonInterestsResultApiResponseWithDefaults

`func NewUserCommonInterestsResultApiResponseWithDefaults() *UserCommonInterestsResultApiResponse`

NewUserCommonInterestsResultApiResponseWithDefaults instantiates a new UserCommonInterestsResultApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UserCommonInterestsResultApiResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UserCommonInterestsResultApiResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UserCommonInterestsResultApiResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *UserCommonInterestsResultApiResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *UserCommonInterestsResultApiResponse) GetData() UserCommonInterestsResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserCommonInterestsResultApiResponse) GetDataOk() (*UserCommonInterestsResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserCommonInterestsResultApiResponse) SetData(v UserCommonInterestsResult)`

SetData sets Data field to given value.

### HasData

`func (o *UserCommonInterestsResultApiResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *UserCommonInterestsResultApiResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UserCommonInterestsResultApiResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UserCommonInterestsResultApiResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *UserCommonInterestsResultApiResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *UserCommonInterestsResultApiResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *UserCommonInterestsResultApiResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


