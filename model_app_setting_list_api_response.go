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

// checks if the AppSettingListApiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppSettingListApiResponse{}

// AppSettingListApiResponse struct for AppSettingListApiResponse
type AppSettingListApiResponse struct {
	// 状态码
	Code *int32 `json:"code,omitempty"`
	// 响应数据
	Data []AppSetting `json:"data,omitempty"`
	// 错误信息
	Error NullableString `json:"error,omitempty"`
}

// NewAppSettingListApiResponse instantiates a new AppSettingListApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppSettingListApiResponse() *AppSettingListApiResponse {
	this := AppSettingListApiResponse{}
	var code int32 = 200
	this.Code = &code
	var error_ string = ""
	this.Error = *NewNullableString(&error_)
	return &this
}

// NewAppSettingListApiResponseWithDefaults instantiates a new AppSettingListApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppSettingListApiResponseWithDefaults() *AppSettingListApiResponse {
	this := AppSettingListApiResponse{}
	var code int32 = 200
	this.Code = &code
	var error_ string = ""
	this.Error = *NewNullableString(&error_)
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AppSettingListApiResponse) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSettingListApiResponse) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AppSettingListApiResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *AppSettingListApiResponse) SetCode(v int32) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppSettingListApiResponse) GetData() []AppSetting {
	if o == nil {
		var ret []AppSetting
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppSettingListApiResponse) GetDataOk() ([]AppSetting, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppSettingListApiResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AppSetting and assigns it to the Data field.
func (o *AppSettingListApiResponse) SetData(v []AppSetting) {
	o.Data = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppSettingListApiResponse) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppSettingListApiResponse) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *AppSettingListApiResponse) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *AppSettingListApiResponse) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *AppSettingListApiResponse) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *AppSettingListApiResponse) UnsetError() {
	o.Error.Unset()
}

func (o AppSettingListApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppSettingListApiResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	return toSerialize, nil
}

type NullableAppSettingListApiResponse struct {
	value *AppSettingListApiResponse
	isSet bool
}

func (v NullableAppSettingListApiResponse) Get() *AppSettingListApiResponse {
	return v.value
}

func (v *NullableAppSettingListApiResponse) Set(val *AppSettingListApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppSettingListApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppSettingListApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppSettingListApiResponse(val *AppSettingListApiResponse) *NullableAppSettingListApiResponse {
	return &NullableAppSettingListApiResponse{value: val, isSet: true}
}

func (v NullableAppSettingListApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppSettingListApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


