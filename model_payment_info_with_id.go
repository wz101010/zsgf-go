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

// checks if the PaymentInfoWithId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentInfoWithId{}

// PaymentInfoWithId struct for PaymentInfoWithId
type PaymentInfoWithId struct {
	PaymentIds []string `json:"paymentIds,omitempty"`
	Type NullableString `json:"type,omitempty"`
}

// NewPaymentInfoWithId instantiates a new PaymentInfoWithId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInfoWithId() *PaymentInfoWithId {
	this := PaymentInfoWithId{}
	return &this
}

// NewPaymentInfoWithIdWithDefaults instantiates a new PaymentInfoWithId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInfoWithIdWithDefaults() *PaymentInfoWithId {
	this := PaymentInfoWithId{}
	return &this
}

// GetPaymentIds returns the PaymentIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInfoWithId) GetPaymentIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PaymentIds
}

// GetPaymentIdsOk returns a tuple with the PaymentIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInfoWithId) GetPaymentIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PaymentIds) {
		return nil, false
	}
	return o.PaymentIds, true
}

// HasPaymentIds returns a boolean if a field has been set.
func (o *PaymentInfoWithId) HasPaymentIds() bool {
	if o != nil && !IsNil(o.PaymentIds) {
		return true
	}

	return false
}

// SetPaymentIds gets a reference to the given []string and assigns it to the PaymentIds field.
func (o *PaymentInfoWithId) SetPaymentIds(v []string) {
	o.PaymentIds = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInfoWithId) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInfoWithId) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *PaymentInfoWithId) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *PaymentInfoWithId) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *PaymentInfoWithId) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *PaymentInfoWithId) UnsetType() {
	o.Type.Unset()
}

func (o PaymentInfoWithId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentInfoWithId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentIds != nil {
		toSerialize["paymentIds"] = o.PaymentIds
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	return toSerialize, nil
}

type NullablePaymentInfoWithId struct {
	value *PaymentInfoWithId
	isSet bool
}

func (v NullablePaymentInfoWithId) Get() *PaymentInfoWithId {
	return v.value
}

func (v *NullablePaymentInfoWithId) Set(val *PaymentInfoWithId) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInfoWithId) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInfoWithId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInfoWithId(val *PaymentInfoWithId) *NullablePaymentInfoWithId {
	return &NullablePaymentInfoWithId{value: val, isSet: true}
}

func (v NullablePaymentInfoWithId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInfoWithId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


