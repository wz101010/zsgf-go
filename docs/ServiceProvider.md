# ServiceProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 服务商的唯一标识符。 | [optional] 
**BizCode** | Pointer to **NullableString** | 服务商的业务代码，用于标识其所属业务领域。 | [optional] 
**Name** | Pointer to **NullableString** | 服务商的名称。 | [optional] 
**Code** | Pointer to **NullableString** | 服务商的唯一代码，用于系统内部标识。 | [optional] 
**Icon** | Pointer to **NullableString** | 服务商图标的URL或路径。 | [optional] 
**Description** | Pointer to **NullableString** | 服务商的详细描述信息。 | [optional] 
**Tags** | Pointer to **NullableString** | 用于分类或标记服务商的标签。 | [optional] 
**Show** | Pointer to **bool** | 指示服务商是否在界面上显示。 | [optional] 
**ShowIndex** | Pointer to **int32** | 服务商在界面上的显示顺序。 | [optional] 
**CreateDate** | Pointer to **time.Time** | 服务商记录的创建日期，默认为当前时间。 | [optional] 
**LastUpdate** | Pointer to **time.Time** | 服务商记录的最后更新日期，默认为当前时间。 | [optional] 

## Methods

### NewServiceProvider

`func NewServiceProvider() *ServiceProvider`

NewServiceProvider instantiates a new ServiceProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProviderWithDefaults

`func NewServiceProviderWithDefaults() *ServiceProvider`

NewServiceProviderWithDefaults instantiates a new ServiceProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceProvider) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceProvider) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceProvider) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBizCode

`func (o *ServiceProvider) GetBizCode() string`

GetBizCode returns the BizCode field if non-nil, zero value otherwise.

### GetBizCodeOk

`func (o *ServiceProvider) GetBizCodeOk() (*string, bool)`

GetBizCodeOk returns a tuple with the BizCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizCode

`func (o *ServiceProvider) SetBizCode(v string)`

SetBizCode sets BizCode field to given value.

### HasBizCode

`func (o *ServiceProvider) HasBizCode() bool`

HasBizCode returns a boolean if a field has been set.

### SetBizCodeNil

`func (o *ServiceProvider) SetBizCodeNil(b bool)`

 SetBizCodeNil sets the value for BizCode to be an explicit nil

### UnsetBizCode
`func (o *ServiceProvider) UnsetBizCode()`

UnsetBizCode ensures that no value is present for BizCode, not even an explicit nil
### GetName

`func (o *ServiceProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ServiceProvider) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ServiceProvider) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCode

`func (o *ServiceProvider) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ServiceProvider) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ServiceProvider) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ServiceProvider) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ServiceProvider) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ServiceProvider) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetIcon

`func (o *ServiceProvider) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ServiceProvider) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ServiceProvider) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *ServiceProvider) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *ServiceProvider) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *ServiceProvider) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetDescription

`func (o *ServiceProvider) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceProvider) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceProvider) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceProvider) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ServiceProvider) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ServiceProvider) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTags

`func (o *ServiceProvider) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceProvider) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceProvider) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceProvider) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ServiceProvider) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ServiceProvider) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetShow

`func (o *ServiceProvider) GetShow() bool`

GetShow returns the Show field if non-nil, zero value otherwise.

### GetShowOk

`func (o *ServiceProvider) GetShowOk() (*bool, bool)`

GetShowOk returns a tuple with the Show field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShow

`func (o *ServiceProvider) SetShow(v bool)`

SetShow sets Show field to given value.

### HasShow

`func (o *ServiceProvider) HasShow() bool`

HasShow returns a boolean if a field has been set.

### GetShowIndex

`func (o *ServiceProvider) GetShowIndex() int32`

GetShowIndex returns the ShowIndex field if non-nil, zero value otherwise.

### GetShowIndexOk

`func (o *ServiceProvider) GetShowIndexOk() (*int32, bool)`

GetShowIndexOk returns a tuple with the ShowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIndex

`func (o *ServiceProvider) SetShowIndex(v int32)`

SetShowIndex sets ShowIndex field to given value.

### HasShowIndex

`func (o *ServiceProvider) HasShowIndex() bool`

HasShowIndex returns a boolean if a field has been set.

### GetCreateDate

`func (o *ServiceProvider) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *ServiceProvider) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *ServiceProvider) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *ServiceProvider) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *ServiceProvider) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *ServiceProvider) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *ServiceProvider) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *ServiceProvider) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


