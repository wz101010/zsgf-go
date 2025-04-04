/*
全部  API 文档

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zsgf

import (
	"encoding/json"
	"time"
)

// checks if the DirectoryItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectoryItem{}

// DirectoryItem struct for DirectoryItem
type DirectoryItem struct {
	Name NullableString `json:"name,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`
	FileSize NullableInt64 `json:"fileSize,omitempty"`
}

// NewDirectoryItem instantiates a new DirectoryItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectoryItem() *DirectoryItem {
	this := DirectoryItem{}
	return &this
}

// NewDirectoryItemWithDefaults instantiates a new DirectoryItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectoryItemWithDefaults() *DirectoryItem {
	this := DirectoryItem{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectoryItem) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectoryItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DirectoryItem) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DirectoryItem) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DirectoryItem) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DirectoryItem) UnsetName() {
	o.Name.Unset()
}

// GetCreateDate returns the CreateDate field value if set, zero value otherwise.
func (o *DirectoryItem) GetCreateDate() time.Time {
	if o == nil || IsNil(o.CreateDate) {
		var ret time.Time
		return ret
	}
	return *o.CreateDate
}

// GetCreateDateOk returns a tuple with the CreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectoryItem) GetCreateDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateDate) {
		return nil, false
	}
	return o.CreateDate, true
}

// HasCreateDate returns a boolean if a field has been set.
func (o *DirectoryItem) HasCreateDate() bool {
	if o != nil && !IsNil(o.CreateDate) {
		return true
	}

	return false
}

// SetCreateDate gets a reference to the given time.Time and assigns it to the CreateDate field.
func (o *DirectoryItem) SetCreateDate(v time.Time) {
	o.CreateDate = &v
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise.
func (o *DirectoryItem) GetLastUpdate() time.Time {
	if o == nil || IsNil(o.LastUpdate) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectoryItem) GetLastUpdateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdate) {
		return nil, false
	}
	return o.LastUpdate, true
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *DirectoryItem) HasLastUpdate() bool {
	if o != nil && !IsNil(o.LastUpdate) {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given time.Time and assigns it to the LastUpdate field.
func (o *DirectoryItem) SetLastUpdate(v time.Time) {
	o.LastUpdate = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectoryItem) GetFileSize() int64 {
	if o == nil || IsNil(o.FileSize.Get()) {
		var ret int64
		return ret
	}
	return *o.FileSize.Get()
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectoryItem) GetFileSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileSize.Get(), o.FileSize.IsSet()
}

// HasFileSize returns a boolean if a field has been set.
func (o *DirectoryItem) HasFileSize() bool {
	if o != nil && o.FileSize.IsSet() {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given NullableInt64 and assigns it to the FileSize field.
func (o *DirectoryItem) SetFileSize(v int64) {
	o.FileSize.Set(&v)
}
// SetFileSizeNil sets the value for FileSize to be an explicit nil
func (o *DirectoryItem) SetFileSizeNil() {
	o.FileSize.Set(nil)
}

// UnsetFileSize ensures that no value is present for FileSize, not even an explicit nil
func (o *DirectoryItem) UnsetFileSize() {
	o.FileSize.Unset()
}

func (o DirectoryItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectoryItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.CreateDate) {
		toSerialize["createDate"] = o.CreateDate
	}
	if !IsNil(o.LastUpdate) {
		toSerialize["lastUpdate"] = o.LastUpdate
	}
	if o.FileSize.IsSet() {
		toSerialize["fileSize"] = o.FileSize.Get()
	}
	return toSerialize, nil
}

type NullableDirectoryItem struct {
	value *DirectoryItem
	isSet bool
}

func (v NullableDirectoryItem) Get() *DirectoryItem {
	return v.value
}

func (v *NullableDirectoryItem) Set(val *DirectoryItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectoryItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectoryItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectoryItem(val *DirectoryItem) *NullableDirectoryItem {
	return &NullableDirectoryItem{value: val, isSet: true}
}

func (v NullableDirectoryItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectoryItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


