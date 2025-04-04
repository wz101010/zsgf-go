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

// checks if the TapPayInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TapPayInfo{}

// TapPayInfo struct for TapPayInfo
type TapPayInfo struct {
	PaymentMediumType NullableString `json:"paymentMediumType,omitempty"`
}

// NewTapPayInfo instantiates a new TapPayInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTapPayInfo() *TapPayInfo {
	this := TapPayInfo{}
	return &this
}

// NewTapPayInfoWithDefaults instantiates a new TapPayInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTapPayInfoWithDefaults() *TapPayInfo {
	this := TapPayInfo{}
	return &this
}

// GetPaymentMediumType returns the PaymentMediumType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TapPayInfo) GetPaymentMediumType() string {
	if o == nil || IsNil(o.PaymentMediumType.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentMediumType.Get()
}

// GetPaymentMediumTypeOk returns a tuple with the PaymentMediumType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TapPayInfo) GetPaymentMediumTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMediumType.Get(), o.PaymentMediumType.IsSet()
}

// HasPaymentMediumType returns a boolean if a field has been set.
func (o *TapPayInfo) HasPaymentMediumType() bool {
	if o != nil && o.PaymentMediumType.IsSet() {
		return true
	}

	return false
}

// SetPaymentMediumType gets a reference to the given NullableString and assigns it to the PaymentMediumType field.
func (o *TapPayInfo) SetPaymentMediumType(v string) {
	o.PaymentMediumType.Set(&v)
}
// SetPaymentMediumTypeNil sets the value for PaymentMediumType to be an explicit nil
func (o *TapPayInfo) SetPaymentMediumTypeNil() {
	o.PaymentMediumType.Set(nil)
}

// UnsetPaymentMediumType ensures that no value is present for PaymentMediumType, not even an explicit nil
func (o *TapPayInfo) UnsetPaymentMediumType() {
	o.PaymentMediumType.Unset()
}

func (o TapPayInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TapPayInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentMediumType.IsSet() {
		toSerialize["paymentMediumType"] = o.PaymentMediumType.Get()
	}
	return toSerialize, nil
}

type NullableTapPayInfo struct {
	value *TapPayInfo
	isSet bool
}

func (v NullableTapPayInfo) Get() *TapPayInfo {
	return v.value
}

func (v *NullableTapPayInfo) Set(val *TapPayInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTapPayInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTapPayInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTapPayInfo(val *TapPayInfo) *NullableTapPayInfo {
	return &NullableTapPayInfo{value: val, isSet: true}
}

func (v NullableTapPayInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTapPayInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


