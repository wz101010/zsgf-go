# ObjectListApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | 状态码 | [optional] [default to 200]
**Data** | Pointer to **[]interface{}** | 响应数据 | [optional] 
**Error** | Pointer to **NullableString** | 错误信息 | [optional] [default to ""]

## Methods

### NewObjectListApiResponse

`func NewObjectListApiResponse() *ObjectListApiResponse`

NewObjectListApiResponse instantiates a new ObjectListApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectListApiResponseWithDefaults

`func NewObjectListApiResponseWithDefaults() *ObjectListApiResponse`

NewObjectListApiResponseWithDefaults instantiates a new ObjectListApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ObjectListApiResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ObjectListApiResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ObjectListApiResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ObjectListApiResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *ObjectListApiResponse) GetData() []interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ObjectListApiResponse) GetDataOk() (*[]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ObjectListApiResponse) SetData(v []interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ObjectListApiResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ObjectListApiResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ObjectListApiResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetError

`func (o *ObjectListApiResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ObjectListApiResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ObjectListApiResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ObjectListApiResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ObjectListApiResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ObjectListApiResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


