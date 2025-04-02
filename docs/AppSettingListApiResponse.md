# AppSettingListApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | 状态码 | [optional] [default to 200]
**Data** | Pointer to [**[]AppSetting**](AppSetting.md) | 响应数据 | [optional] 
**Error** | Pointer to **NullableString** | 错误信息 | [optional] [default to ""]

## Methods

### NewAppSettingListApiResponse

`func NewAppSettingListApiResponse() *AppSettingListApiResponse`

NewAppSettingListApiResponse instantiates a new AppSettingListApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSettingListApiResponseWithDefaults

`func NewAppSettingListApiResponseWithDefaults() *AppSettingListApiResponse`

NewAppSettingListApiResponseWithDefaults instantiates a new AppSettingListApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AppSettingListApiResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AppSettingListApiResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AppSettingListApiResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *AppSettingListApiResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *AppSettingListApiResponse) GetData() []AppSetting`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppSettingListApiResponse) GetDataOk() (*[]AppSetting, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppSettingListApiResponse) SetData(v []AppSetting)`

SetData sets Data field to given value.

### HasData

`func (o *AppSettingListApiResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *AppSettingListApiResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *AppSettingListApiResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetError

`func (o *AppSettingListApiResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AppSettingListApiResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AppSettingListApiResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AppSettingListApiResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *AppSettingListApiResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *AppSettingListApiResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


