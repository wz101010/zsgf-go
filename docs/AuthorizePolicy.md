# AuthorizePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 鉴权策略的唯一标识符。 | [optional] 
**PolicyName** | Pointer to **NullableString** | 鉴权策略的名称。根据鉴权类型填写不同的名称：角色类型填写角色名称，用户类型填写用户ID，访问令牌类型填写令牌ID。 | [optional] 
**AuthorizeType** | Pointer to **NullableString** | 鉴权策略的类型，可选值为 &#39;access_token&#39;, &#39;user&#39;, 或 &#39;role&#39;。 | [optional] 
**PolicyValue** | Pointer to **NullableString** | 与鉴权策略关联的权限集合，多个权限可以用逗号分隔。 | [optional] 
**CreateDate** | Pointer to **time.Time** | 鉴权策略的创建日期，默认为当前时间。 | [optional] 
**LastUpdate** | Pointer to **time.Time** | 鉴权策略的最后更新日期，默认为当前时间。 | [optional] 

## Methods

### NewAuthorizePolicy

`func NewAuthorizePolicy() *AuthorizePolicy`

NewAuthorizePolicy instantiates a new AuthorizePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizePolicyWithDefaults

`func NewAuthorizePolicyWithDefaults() *AuthorizePolicy`

NewAuthorizePolicyWithDefaults instantiates a new AuthorizePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthorizePolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizePolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizePolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorizePolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPolicyName

`func (o *AuthorizePolicy) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *AuthorizePolicy) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *AuthorizePolicy) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *AuthorizePolicy) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *AuthorizePolicy) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *AuthorizePolicy) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetAuthorizeType

`func (o *AuthorizePolicy) GetAuthorizeType() string`

GetAuthorizeType returns the AuthorizeType field if non-nil, zero value otherwise.

### GetAuthorizeTypeOk

`func (o *AuthorizePolicy) GetAuthorizeTypeOk() (*string, bool)`

GetAuthorizeTypeOk returns a tuple with the AuthorizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeType

`func (o *AuthorizePolicy) SetAuthorizeType(v string)`

SetAuthorizeType sets AuthorizeType field to given value.

### HasAuthorizeType

`func (o *AuthorizePolicy) HasAuthorizeType() bool`

HasAuthorizeType returns a boolean if a field has been set.

### SetAuthorizeTypeNil

`func (o *AuthorizePolicy) SetAuthorizeTypeNil(b bool)`

 SetAuthorizeTypeNil sets the value for AuthorizeType to be an explicit nil

### UnsetAuthorizeType
`func (o *AuthorizePolicy) UnsetAuthorizeType()`

UnsetAuthorizeType ensures that no value is present for AuthorizeType, not even an explicit nil
### GetPolicyValue

`func (o *AuthorizePolicy) GetPolicyValue() string`

GetPolicyValue returns the PolicyValue field if non-nil, zero value otherwise.

### GetPolicyValueOk

`func (o *AuthorizePolicy) GetPolicyValueOk() (*string, bool)`

GetPolicyValueOk returns a tuple with the PolicyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyValue

`func (o *AuthorizePolicy) SetPolicyValue(v string)`

SetPolicyValue sets PolicyValue field to given value.

### HasPolicyValue

`func (o *AuthorizePolicy) HasPolicyValue() bool`

HasPolicyValue returns a boolean if a field has been set.

### SetPolicyValueNil

`func (o *AuthorizePolicy) SetPolicyValueNil(b bool)`

 SetPolicyValueNil sets the value for PolicyValue to be an explicit nil

### UnsetPolicyValue
`func (o *AuthorizePolicy) UnsetPolicyValue()`

UnsetPolicyValue ensures that no value is present for PolicyValue, not even an explicit nil
### GetCreateDate

`func (o *AuthorizePolicy) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *AuthorizePolicy) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *AuthorizePolicy) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *AuthorizePolicy) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetLastUpdate

`func (o *AuthorizePolicy) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *AuthorizePolicy) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *AuthorizePolicy) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *AuthorizePolicy) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


