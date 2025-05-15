# UserSettingApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | 状态码 | [optional] [default to 200]
**Data** | Pointer to [**UserSetting**](UserSetting.md) |  | [optional] 
**Error** | Pointer to **NullableString** | 错误信息 | [optional] [default to ""]

## Methods

### NewUserSettingApiResponse

`func NewUserSettingApiResponse() *UserSettingApiResponse`

NewUserSettingApiResponse instantiates a new UserSettingApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSettingApiResponseWithDefaults

`func NewUserSettingApiResponseWithDefaults() *UserSettingApiResponse`

NewUserSettingApiResponseWithDefaults instantiates a new UserSettingApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UserSettingApiResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UserSettingApiResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UserSettingApiResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *UserSettingApiResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *UserSettingApiResponse) GetData() UserSetting`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserSettingApiResponse) GetDataOk() (*UserSetting, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserSettingApiResponse) SetData(v UserSetting)`

SetData sets Data field to given value.

### HasData

`func (o *UserSettingApiResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *UserSettingApiResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UserSettingApiResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UserSettingApiResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *UserSettingApiResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *UserSettingApiResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *UserSettingApiResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


