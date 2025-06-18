# ExternalAccountPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **NullableString** | 头像 | [optional] 
**Data** | Pointer to **NullableString** | 自定义数据 | [optional] 
**Enable** | Pointer to **bool** | 启用 | [optional] 

## Methods

### NewExternalAccountPutRequest

`func NewExternalAccountPutRequest() *ExternalAccountPutRequest`

NewExternalAccountPutRequest instantiates a new ExternalAccountPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAccountPutRequestWithDefaults

`func NewExternalAccountPutRequestWithDefaults() *ExternalAccountPutRequest`

NewExternalAccountPutRequestWithDefaults instantiates a new ExternalAccountPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *ExternalAccountPutRequest) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ExternalAccountPutRequest) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ExternalAccountPutRequest) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *ExternalAccountPutRequest) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *ExternalAccountPutRequest) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *ExternalAccountPutRequest) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetData

`func (o *ExternalAccountPutRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExternalAccountPutRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExternalAccountPutRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ExternalAccountPutRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ExternalAccountPutRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ExternalAccountPutRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetEnable

`func (o *ExternalAccountPutRequest) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ExternalAccountPutRequest) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ExternalAccountPutRequest) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ExternalAccountPutRequest) HasEnable() bool`

HasEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


