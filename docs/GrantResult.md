# GrantResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 
**ExpiresIn** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGrantResult

`func NewGrantResult() *GrantResult`

NewGrantResult instantiates a new GrantResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantResultWithDefaults

`func NewGrantResultWithDefaults() *GrantResult`

NewGrantResultWithDefaults instantiates a new GrantResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GrantResult) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GrantResult) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GrantResult) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GrantResult) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetError

`func (o *GrantResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GrantResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GrantResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GrantResult) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *GrantResult) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *GrantResult) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetExpiresIn

`func (o *GrantResult) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *GrantResult) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *GrantResult) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *GrantResult) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetData

`func (o *GrantResult) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GrantResult) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GrantResult) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *GrantResult) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *GrantResult) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *GrantResult) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


