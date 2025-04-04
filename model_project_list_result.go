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

// checks if the ProjectListResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectListResult{}

// ProjectListResult struct for ProjectListResult
type ProjectListResult struct {
	Total *int32 `json:"total,omitempty"`
	Data []Project `json:"data,omitempty"`
}

// NewProjectListResult instantiates a new ProjectListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectListResult() *ProjectListResult {
	this := ProjectListResult{}
	return &this
}

// NewProjectListResultWithDefaults instantiates a new ProjectListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectListResultWithDefaults() *ProjectListResult {
	this := ProjectListResult{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ProjectListResult) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectListResult) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ProjectListResult) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ProjectListResult) SetTotal(v int32) {
	o.Total = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectListResult) GetData() []Project {
	if o == nil {
		var ret []Project
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectListResult) GetDataOk() ([]Project, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ProjectListResult) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Project and assigns it to the Data field.
func (o *ProjectListResult) SetData(v []Project) {
	o.Data = v
}

func (o ProjectListResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectListResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableProjectListResult struct {
	value *ProjectListResult
	isSet bool
}

func (v NullableProjectListResult) Get() *ProjectListResult {
	return v.value
}

func (v *NullableProjectListResult) Set(val *ProjectListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectListResult(val *ProjectListResult) *NullableProjectListResult {
	return &NullableProjectListResult{value: val, isSet: true}
}

func (v NullableProjectListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


