# AppProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**Desc** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAppProperty

`func NewAppProperty() *AppProperty`

NewAppProperty instantiates a new AppProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPropertyWithDefaults

`func NewAppPropertyWithDefaults() *AppProperty`

NewAppPropertyWithDefaults instantiates a new AppProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AppProperty) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AppProperty) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AppProperty) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AppProperty) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *AppProperty) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *AppProperty) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetValue

`func (o *AppProperty) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AppProperty) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AppProperty) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AppProperty) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *AppProperty) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *AppProperty) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetDesc

`func (o *AppProperty) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *AppProperty) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *AppProperty) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *AppProperty) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### SetDescNil

`func (o *AppProperty) SetDescNil(b bool)`

 SetDescNil sets the value for Desc to be an explicit nil

### UnsetDesc
`func (o *AppProperty) UnsetDesc()`

UnsetDesc ensures that no value is present for Desc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


