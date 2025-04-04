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

// checks if the AppUserResetPwdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUserResetPwdRequest{}

// AppUserResetPwdRequest struct for AppUserResetPwdRequest
type AppUserResetPwdRequest struct {
	Phone NullableString `json:"phone,omitempty"`
	Email NullableString `json:"email,omitempty"`
	Code string `json:"code" validate:"regexp=\\\\d{4,8}$"`
	Pwd string `json:"pwd" validate:"regexp=^.{6,32}$"`
}

type _AppUserResetPwdRequest AppUserResetPwdRequest

// NewAppUserResetPwdRequest instantiates a new AppUserResetPwdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUserResetPwdRequest(code string, pwd string) *AppUserResetPwdRequest {
	this := AppUserResetPwdRequest{}
	this.Code = code
	this.Pwd = pwd
	return &this
}

// NewAppUserResetPwdRequestWithDefaults instantiates a new AppUserResetPwdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUserResetPwdRequestWithDefaults() *AppUserResetPwdRequest {
	this := AppUserResetPwdRequest{}
	return &this
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppUserResetPwdRequest) GetPhone() string {
	if o == nil || IsNil(o.Phone.Get()) {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppUserResetPwdRequest) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *AppUserResetPwdRequest) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *AppUserResetPwdRequest) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *AppUserResetPwdRequest) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *AppUserResetPwdRequest) UnsetPhone() {
	o.Phone.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppUserResetPwdRequest) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppUserResetPwdRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *AppUserResetPwdRequest) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *AppUserResetPwdRequest) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *AppUserResetPwdRequest) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *AppUserResetPwdRequest) UnsetEmail() {
	o.Email.Unset()
}

// GetCode returns the Code field value
func (o *AppUserResetPwdRequest) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AppUserResetPwdRequest) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AppUserResetPwdRequest) SetCode(v string) {
	o.Code = v
}

// GetPwd returns the Pwd field value
func (o *AppUserResetPwdRequest) GetPwd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pwd
}

// GetPwdOk returns a tuple with the Pwd field value
// and a boolean to check if the value has been set.
func (o *AppUserResetPwdRequest) GetPwdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pwd, true
}

// SetPwd sets field value
func (o *AppUserResetPwdRequest) SetPwd(v string) {
	o.Pwd = v
}

func (o AppUserResetPwdRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUserResetPwdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	toSerialize["code"] = o.Code
	toSerialize["pwd"] = o.Pwd
	return toSerialize, nil
}

func (o *AppUserResetPwdRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"pwd",
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

	varAppUserResetPwdRequest := _AppUserResetPwdRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppUserResetPwdRequest)

	if err != nil {
		return err
	}

	*o = AppUserResetPwdRequest(varAppUserResetPwdRequest)

	return err
}

type NullableAppUserResetPwdRequest struct {
	value *AppUserResetPwdRequest
	isSet bool
}

func (v NullableAppUserResetPwdRequest) Get() *AppUserResetPwdRequest {
	return v.value
}

func (v *NullableAppUserResetPwdRequest) Set(val *AppUserResetPwdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUserResetPwdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUserResetPwdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUserResetPwdRequest(val *AppUserResetPwdRequest) *NullableAppUserResetPwdRequest {
	return &NullableAppUserResetPwdRequest{value: val, isSet: true}
}

func (v NullableAppUserResetPwdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUserResetPwdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


