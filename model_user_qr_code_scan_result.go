/*
全部  API 文档

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zsgf

import (
	"encoding/json"
)

// checks if the UserQRCodeScanResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserQRCodeScanResult{}

// UserQRCodeScanResult struct for UserQRCodeScanResult
type UserQRCodeScanResult struct {
	AppID *int64 `json:"appID,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Logo NullableString `json:"logo,omitempty"`
	Website NullableString `json:"website,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Tags NullableString `json:"tags,omitempty"`
	Scopes NullableString `json:"scopes,omitempty"`
	Remark NullableString `json:"remark,omitempty"`
	Scheme NullableString `json:"scheme,omitempty"`
}

// NewUserQRCodeScanResult instantiates a new UserQRCodeScanResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserQRCodeScanResult() *UserQRCodeScanResult {
	this := UserQRCodeScanResult{}
	return &this
}

// NewUserQRCodeScanResultWithDefaults instantiates a new UserQRCodeScanResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserQRCodeScanResultWithDefaults() *UserQRCodeScanResult {
	this := UserQRCodeScanResult{}
	return &this
}

// GetAppID returns the AppID field value if set, zero value otherwise.
func (o *UserQRCodeScanResult) GetAppID() int64 {
	if o == nil || IsNil(o.AppID) {
		var ret int64
		return ret
	}
	return *o.AppID
}

// GetAppIDOk returns a tuple with the AppID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserQRCodeScanResult) GetAppIDOk() (*int64, bool) {
	if o == nil || IsNil(o.AppID) {
		return nil, false
	}
	return o.AppID, true
}

// HasAppID returns a boolean if a field has been set.
func (o *UserQRCodeScanResult) HasAppID() bool {
	if o != nil && !IsNil(o.AppID) {
		return true
	}

	return false
}

// SetAppID gets a reference to the given int64 and assigns it to the AppID field.
func (o *UserQRCodeScanResult) SetAppID(v int64) {
	o.AppID = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserQRCodeScanResult) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserQRCodeScanResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UserQRCodeScanResult) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UserQRCodeScanResult) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UserQRCodeScanResult) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UserQRCodeScanResult) UnsetName() {
	o.Name.Unset()
}

// GetLogo returns the Logo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserQRCodeScanResult) GetLogo() string {
	if o == nil || IsNil(o.Logo.Get()) {
		var ret string
		return ret
	}
	return *o.Logo.Get()
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserQRCodeScanResult) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logo.Get(), o.Logo.IsSet()
}

// HasLogo returns a boolean if a field has been set.
func (o *UserQRCodeScanResult) HasLogo() bool {
	if o != nil && o.Logo.IsSet() {
		return true
	}

	return false
}

// SetLogo gets a reference to the given NullableString and assigns it to the Logo field.
func (o *UserQRCodeScanResult) SetLogo(v string) {
	o.Logo.Set(&v)
}
// SetLogoNil sets the value for Logo to be an explicit nil
func (o *UserQRCodeScanResult) SetLogoNil() {
	o.Logo.Set(nil)
}

// UnsetLogo ensures that no value is present for Logo, not even an explicit nil
func (o *UserQRCodeScanResult) UnsetLogo() {
	o.Logo.Unset()
}

// GetWebsite returns the Website field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserQRCodeScanResult) GetWebsite() string {
	if o == nil || IsNil(o.Website.Get()) {
		var ret string
		return ret
	}
	return *o.Website.Get()
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserQRCodeScanResult) GetWebsiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Website.Get(), o.Website.IsSet()
}

// HasWebsite returns a boolean if a field has been set.
func (o *UserQRCodeScanResult) HasWebsite() bool {
	if o != nil && o.Website.IsSet() {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given NullableString and assigns it to the Website field.
func (o *UserQRCodeScanResult) SetWebsite(v string) {
	o.Website.Set(&v)
}
// SetWebsiteNil sets the value for Website to be an explicit nil
func (o *UserQRCodeScanResult) SetWebsiteNil() {
	o.Website.Set(nil)
}

// UnsetWebsite ensures that no value is present for Website, not even an explicit nil
func (o *UserQRCodeScanResult) UnsetWebsite() {
	o.Website.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserQRCodeScanResult) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserQRCodeScanResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UserQRCodeScanResult) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UserQRCodeScanResult) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UserQRCodeScanResult) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UserQRCodeScanResult) UnsetDescription() {
	o.Description.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserQRCodeScanResult) GetTags() string {
	if o == nil || IsNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserQRCodeScanResult) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *UserQRCodeScanResult) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *UserQRCodeScanResult) SetTags(v string) {
	o.Tags.Set(&v)
}
// SetTagsNil sets the value for Tags to be an explicit nil
func (o *UserQRCodeScanResult) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *UserQRCodeScanResult) UnsetTags() {
	o.Tags.Unset()
}

// GetScopes returns the Scopes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserQRCodeScanResult) GetScopes() string {
	if o == nil || IsNil(o.Scopes.Get()) {
		var ret string
		return ret
	}
	return *o.Scopes.Get()
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserQRCodeScanResult) GetScopesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes.Get(), o.Scopes.IsSet()
}

// HasScopes returns a boolean if a field has been set.
func (o *UserQRCodeScanResult) HasScopes() bool {
	if o != nil && o.Scopes.IsSet() {
		return true
	}

	return false
}

// SetScopes gets a reference to the given NullableString and assigns it to the Scopes field.
func (o *UserQRCodeScanResult) SetScopes(v string) {
	o.Scopes.Set(&v)
}
// SetScopesNil sets the value for Scopes to be an explicit nil
func (o *UserQRCodeScanResult) SetScopesNil() {
	o.Scopes.Set(nil)
}

// UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
func (o *UserQRCodeScanResult) UnsetScopes() {
	o.Scopes.Unset()
}

// GetRemark returns the Remark field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserQRCodeScanResult) GetRemark() string {
	if o == nil || IsNil(o.Remark.Get()) {
		var ret string
		return ret
	}
	return *o.Remark.Get()
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserQRCodeScanResult) GetRemarkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Remark.Get(), o.Remark.IsSet()
}

// HasRemark returns a boolean if a field has been set.
func (o *UserQRCodeScanResult) HasRemark() bool {
	if o != nil && o.Remark.IsSet() {
		return true
	}

	return false
}

// SetRemark gets a reference to the given NullableString and assigns it to the Remark field.
func (o *UserQRCodeScanResult) SetRemark(v string) {
	o.Remark.Set(&v)
}
// SetRemarkNil sets the value for Remark to be an explicit nil
func (o *UserQRCodeScanResult) SetRemarkNil() {
	o.Remark.Set(nil)
}

// UnsetRemark ensures that no value is present for Remark, not even an explicit nil
func (o *UserQRCodeScanResult) UnsetRemark() {
	o.Remark.Unset()
}

// GetScheme returns the Scheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserQRCodeScanResult) GetScheme() string {
	if o == nil || IsNil(o.Scheme.Get()) {
		var ret string
		return ret
	}
	return *o.Scheme.Get()
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserQRCodeScanResult) GetSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scheme.Get(), o.Scheme.IsSet()
}

// HasScheme returns a boolean if a field has been set.
func (o *UserQRCodeScanResult) HasScheme() bool {
	if o != nil && o.Scheme.IsSet() {
		return true
	}

	return false
}

// SetScheme gets a reference to the given NullableString and assigns it to the Scheme field.
func (o *UserQRCodeScanResult) SetScheme(v string) {
	o.Scheme.Set(&v)
}
// SetSchemeNil sets the value for Scheme to be an explicit nil
func (o *UserQRCodeScanResult) SetSchemeNil() {
	o.Scheme.Set(nil)
}

// UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
func (o *UserQRCodeScanResult) UnsetScheme() {
	o.Scheme.Unset()
}

func (o UserQRCodeScanResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserQRCodeScanResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppID) {
		toSerialize["appID"] = o.AppID
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Logo.IsSet() {
		toSerialize["logo"] = o.Logo.Get()
	}
	if o.Website.IsSet() {
		toSerialize["website"] = o.Website.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	if o.Scopes.IsSet() {
		toSerialize["scopes"] = o.Scopes.Get()
	}
	if o.Remark.IsSet() {
		toSerialize["remark"] = o.Remark.Get()
	}
	if o.Scheme.IsSet() {
		toSerialize["scheme"] = o.Scheme.Get()
	}
	return toSerialize, nil
}

type NullableUserQRCodeScanResult struct {
	value *UserQRCodeScanResult
	isSet bool
}

func (v NullableUserQRCodeScanResult) Get() *UserQRCodeScanResult {
	return v.value
}

func (v *NullableUserQRCodeScanResult) Set(val *UserQRCodeScanResult) {
	v.value = val
	v.isSet = true
}

func (v NullableUserQRCodeScanResult) IsSet() bool {
	return v.isSet
}

func (v *NullableUserQRCodeScanResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserQRCodeScanResult(val *UserQRCodeScanResult) *NullableUserQRCodeScanResult {
	return &NullableUserQRCodeScanResult{value: val, isSet: true}
}

func (v NullableUserQRCodeScanResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserQRCodeScanResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


