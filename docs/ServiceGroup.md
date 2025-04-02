# ServiceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 服务功能分组的唯一标识符。 | [optional] 
**ProviderID** | Pointer to **int64** | 关联的服务商的唯一标识符。 | [optional] 
**Name** | Pointer to **NullableString** | 服务功能分组的名称。 | [optional] 
**Code** | Pointer to **NullableString** | 服务功能分组的唯一代码，用于系统内部标识。 | [optional] 
**Icon** | Pointer to **NullableString** | 服务功能分组的图标URL或路径。 | [optional] 
**Description** | Pointer to **NullableString** | 服务功能分组的详细描述信息。 | [optional] 
**Show** | Pointer to **bool** | 指示服务功能分组是否在界面上显示。 | [optional] 
**ShowIndex** | Pointer to **int32** | 服务功能分组在界面上的显示顺序。 | [optional] 
**CreateDate** | Pointer to **time.Time** | 服务功能分组的创建日期，默认为当前时间。 | [optional] 
**LastUpdate** | Pointer to **time.Time** | 服务功能分组的最后更新日期，默认为当前时间。 | [optional] 

## Methods

### NewServiceGroup

`func NewServiceGroup() *ServiceGroup`

NewServiceGroup instantiates a new ServiceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceGroupWithDefaults

`func NewServiceGroupWithDefaults() *ServiceGroup`

NewServiceGroupWithDefaults instantiates a new ServiceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderID

`func (o *ServiceGroup) GetProviderID() int64`

GetProviderID returns the ProviderID field if non-nil, zero value otherwise.

### GetProviderIDOk

`func (o *ServiceGroup) GetProviderIDOk() (*int64, bool)`

GetProviderIDOk returns a tuple with the ProviderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderID

`func (o *ServiceGroup) SetProviderID(v int64)`

SetProviderID sets ProviderID field to given value.

### HasProviderID

`func (o *ServiceGroup) HasProviderID() bool`

HasProviderID returns a boolean if a field has been set.

### GetName

`func (o *ServiceGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ServiceGroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ServiceGroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCode

`func (o *ServiceGroup) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ServiceGroup) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ServiceGroup) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ServiceGroup) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ServiceGroup) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ServiceGroup) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetIcon

`func (o *ServiceGroup) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ServiceGroup) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ServiceGroup) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ServiceGroup) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *ServiceGroup) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *ServiceGroup) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetDescription

`func (o *ServiceGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ServiceGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ServiceGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetShow

`func (o *ServiceGroup) GetShow() bool`

GetShow returns the Show field if non-nil, zero value otherwise.

### GetShowOk

`func (o *ServiceGroup) GetShowOk() (*bool, bool)`

GetShowOk returns a tuple with the Show field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShow

`func (o *ServiceGroup) SetShow(v bool)`

SetShow sets Show field to given value.

### HasShow

`func (o *ServiceGroup) HasShow() bool`

HasShow returns a boolean if a field has been set.

### GetShowIndex

`func (o *ServiceGroup) GetShowIndex() int32`

GetShowIndex returns the ShowIndex field if non-nil, zero value otherwise.

### GetShowIndexOk

`func (o *ServiceGroup) GetShowIndexOk() (*int32, bool)`

GetShowIndexOk returns a tuple with the ShowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIndex

`func (o *ServiceGroup) SetShowIndex(v int32)`

SetShowIndex sets ShowIndex field to given value.

### HasShowIndex

`func (o *ServiceGroup) HasShowIndex() bool`

HasShowIndex returns a boolean if a field has been set.

### GetCreateDate

`func (o *ServiceGroup) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *ServiceGroup) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *ServiceGroup) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *ServiceGroup) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *ServiceGroup) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *ServiceGroup) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *ServiceGroup) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *ServiceGroup) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


