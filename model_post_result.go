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

// checks if the PostResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostResult{}

// PostResult struct for PostResult
type PostResult struct {
	Id *int64 `json:"id,omitempty"`
}

// NewPostResult instantiates a new PostResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostResult() *PostResult {
	this := PostResult{}
	return &this
}

// NewPostResultWithDefaults instantiates a new PostResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostResultWithDefaults() *PostResult {
	this := PostResult{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PostResult) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostResult) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PostResult) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *PostResult) SetId(v int64) {
	o.Id = &v
}

func (o PostResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePostResult struct {
	value *PostResult
	isSet bool
}

func (v NullablePostResult) Get() *PostResult {
	return v.value
}

func (v *NullablePostResult) Set(val *PostResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePostResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePostResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostResult(val *PostResult) *NullablePostResult {
	return &NullablePostResult{value: val, isSet: true}
}

func (v NullablePostResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


