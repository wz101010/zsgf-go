# UserLoginsListApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | 状态码 | [optional] [default to 200]
**Data** | Pointer to [**[]UserLogins**](UserLogins.md) | 响应数据 | [optional] 
**Error** | Pointer to **NullableString** | 错误信息 | [optional] [default to ""]

## Methods

### NewUserLoginsListApiResponse

`func NewUserLoginsListApiResponse() *UserLoginsListApiResponse`

NewUserLoginsListApiResponse instantiates a new UserLoginsListApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLoginsListApiResponseWithDefaults

`func NewUserLoginsListApiResponseWithDefaults() *UserLoginsListApiResponse`

NewUserLoginsListApiResponseWithDefaults instantiates a new UserLoginsListApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UserLoginsListApiResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UserLoginsListApiResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UserLoginsListApiResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *UserLoginsListApiResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *UserLoginsListApiResponse) GetData() []UserLogins`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserLoginsListApiResponse) GetDataOk() (*[]UserLogins, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserLoginsListApiResponse) SetData(v []UserLogins)`

SetData sets Data field to given value.

### HasData

`func (o *UserLoginsListApiResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *UserLoginsListApiResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *UserLoginsListApiResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetError

`func (o *UserLoginsListApiResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UserLoginsListApiResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UserLoginsListApiResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *UserLoginsListApiResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *UserLoginsListApiResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *UserLoginsListApiResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


