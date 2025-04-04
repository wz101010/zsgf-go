/*
全部  API 文档

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zsgf

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GrantRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantRequest{}

// GrantRequest struct for GrantRequest
type GrantRequest struct {
	// 返回地址。默认无限制，也可在【安全-开放认证网址白名单】配置
	RedirectUri NullableString `json:"redirect_uri,omitempty"`
	// 授权类型。可选：email、phone、unionid、account
	GrantType string `json:"grant_type" validate:"regexp=^(email|phone|unionid|account)$"`
	// 自定义授权范围，用英文空格分隔
	Scopes string `json:"scopes"`
	// 用户名。授权类型为：email/phone/account必填
	UserName NullableString `json:"userName,omitempty"`
	// 登录密码。授权类型为：email/phone/account必填
	Password NullableString `json:"password,omitempty"`
	// unionId。授权类型为：unionid必填
	UnionId NullableString `json:"unionId,omitempty"`
	// 平台。授权类型为：unionid必填
	Platform NullableString `json:"platform,omitempty"`
	// 授权有效期。1~99天
	ExpireInDays *int32 `json:"expireInDays,omitempty"`
}

type _GrantRequest GrantRequest

// NewGrantRequest instantiates a new GrantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantRequest(grantType string, scopes string) *GrantRequest {
	this := GrantRequest{}
	this.GrantType = grantType
	this.Scopes = scopes
	return &this
}

// NewGrantRequestWithDefaults instantiates a new GrantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantRequestWithDefaults() *GrantRequest {
	this := GrantRequest{}
	return &this
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GrantRequest) GetRedirectUri() string {
	if o == nil || IsNil(o.RedirectUri.Get()) {
		var ret string
		return ret
	}
	return *o.RedirectUri.Get()
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GrantRequest) GetRedirectUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUri.Get(), o.RedirectUri.IsSet()
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *GrantRequest) HasRedirectUri() bool {
	if o != nil && o.RedirectUri.IsSet() {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given NullableString and assigns it to the RedirectUri field.
func (o *GrantRequest) SetRedirectUri(v string) {
	o.RedirectUri.Set(&v)
}
// SetRedirectUriNil sets the value for RedirectUri to be an explicit nil
func (o *GrantRequest) SetRedirectUriNil() {
	o.RedirectUri.Set(nil)
}

// UnsetRedirectUri ensures that no value is present for RedirectUri, not even an explicit nil
func (o *GrantRequest) UnsetRedirectUri() {
	o.RedirectUri.Unset()
}

// GetGrantType returns the GrantType field value
func (o *GrantRequest) GetGrantType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value
// and a boolean to check if the value has been set.
func (o *GrantRequest) GetGrantTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantType, true
}

// SetGrantType sets field value
func (o *GrantRequest) SetGrantType(v string) {
	o.GrantType = v
}

// GetScopes returns the Scopes field value
func (o *GrantRequest) GetScopes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *GrantRequest) GetScopesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scopes, true
}

// SetScopes sets field value
func (o *GrantRequest) SetScopes(v string) {
	o.Scopes = v
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GrantRequest) GetUserName() string {
	if o == nil || IsNil(o.UserName.Get()) {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GrantRequest) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *GrantRequest) HasUserName() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *GrantRequest) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *GrantRequest) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *GrantRequest) UnsetUserName() {
	o.UserName.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GrantRequest) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GrantRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *GrantRequest) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *GrantRequest) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *GrantRequest) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *GrantRequest) UnsetPassword() {
	o.Password.Unset()
}

// GetUnionId returns the UnionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GrantRequest) GetUnionId() string {
	if o == nil || IsNil(o.UnionId.Get()) {
		var ret string
		return ret
	}
	return *o.UnionId.Get()
}

// GetUnionIdOk returns a tuple with the UnionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GrantRequest) GetUnionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnionId.Get(), o.UnionId.IsSet()
}

// HasUnionId returns a boolean if a field has been set.
func (o *GrantRequest) HasUnionId() bool {
	if o != nil && o.UnionId.IsSet() {
		return true
	}

	return false
}

// SetUnionId gets a reference to the given NullableString and assigns it to the UnionId field.
func (o *GrantRequest) SetUnionId(v string) {
	o.UnionId.Set(&v)
}
// SetUnionIdNil sets the value for UnionId to be an explicit nil
func (o *GrantRequest) SetUnionIdNil() {
	o.UnionId.Set(nil)
}

// UnsetUnionId ensures that no value is present for UnionId, not even an explicit nil
func (o *GrantRequest) UnsetUnionId() {
	o.UnionId.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GrantRequest) GetPlatform() string {
	if o == nil || IsNil(o.Platform.Get()) {
		var ret string
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GrantRequest) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *GrantRequest) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableString and assigns it to the Platform field.
func (o *GrantRequest) SetPlatform(v string) {
	o.Platform.Set(&v)
}
// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *GrantRequest) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *GrantRequest) UnsetPlatform() {
	o.Platform.Unset()
}

// GetExpireInDays returns the ExpireInDays field value if set, zero value otherwise.
func (o *GrantRequest) GetExpireInDays() int32 {
	if o == nil || IsNil(o.ExpireInDays) {
		var ret int32
		return ret
	}
	return *o.ExpireInDays
}

// GetExpireInDaysOk returns a tuple with the ExpireInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantRequest) GetExpireInDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpireInDays) {
		return nil, false
	}
	return o.ExpireInDays, true
}

// HasExpireInDays returns a boolean if a field has been set.
func (o *GrantRequest) HasExpireInDays() bool {
	if o != nil && !IsNil(o.ExpireInDays) {
		return true
	}

	return false
}

// SetExpireInDays gets a reference to the given int32 and assigns it to the ExpireInDays field.
func (o *GrantRequest) SetExpireInDays(v int32) {
	o.ExpireInDays = &v
}

func (o GrantRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RedirectUri.IsSet() {
		toSerialize["redirect_uri"] = o.RedirectUri.Get()
	}
	toSerialize["grant_type"] = o.GrantType
	toSerialize["scopes"] = o.Scopes
	if o.UserName.IsSet() {
		toSerialize["userName"] = o.UserName.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.UnionId.IsSet() {
		toSerialize["unionId"] = o.UnionId.Get()
	}
	if o.Platform.IsSet() {
		toSerialize["platform"] = o.Platform.Get()
	}
	if !IsNil(o.ExpireInDays) {
		toSerialize["expireInDays"] = o.ExpireInDays
	}
	return toSerialize, nil
}

func (o *GrantRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"grant_type",
		"scopes",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGrantRequest := _GrantRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGrantRequest)

	if err != nil {
		return err
	}

	*o = GrantRequest(varGrantRequest)

	return err
}

type NullableGrantRequest struct {
	value *GrantRequest
	isSet bool
}

func (v NullableGrantRequest) Get() *GrantRequest {
	return v.value
}

func (v *NullableGrantRequest) Set(val *GrantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantRequest(val *GrantRequest) *NullableGrantRequest {
	return &NullableGrantRequest{value: val, isSet: true}
}

func (v NullableGrantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


