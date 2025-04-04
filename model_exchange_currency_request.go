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

// checks if the ExchangeCurrencyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeCurrencyRequest{}

// ExchangeCurrencyRequest struct for ExchangeCurrencyRequest
type ExchangeCurrencyRequest struct {
	// 源虚拟币代码
	FromCurrency string `json:"fromCurrency"`
	// 目标虚拟币代码
	Currency string `json:"currency"`
	// 兑换额
	Balance int32 `json:"balance"`
	// 备注
	Remark NullableString `json:"remark,omitempty"`
	// 描述
	Description NullableString `json:"description,omitempty"`
	// 标签
	Tags NullableString `json:"tags,omitempty"`
}

type _ExchangeCurrencyRequest ExchangeCurrencyRequest

// NewExchangeCurrencyRequest instantiates a new ExchangeCurrencyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeCurrencyRequest(fromCurrency string, currency string, balance int32) *ExchangeCurrencyRequest {
	this := ExchangeCurrencyRequest{}
	this.FromCurrency = fromCurrency
	this.Currency = currency
	this.Balance = balance
	return &this
}

// NewExchangeCurrencyRequestWithDefaults instantiates a new ExchangeCurrencyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeCurrencyRequestWithDefaults() *ExchangeCurrencyRequest {
	this := ExchangeCurrencyRequest{}
	return &this
}

// GetFromCurrency returns the FromCurrency field value
func (o *ExchangeCurrencyRequest) GetFromCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromCurrency
}

// GetFromCurrencyOk returns a tuple with the FromCurrency field value
// and a boolean to check if the value has been set.
func (o *ExchangeCurrencyRequest) GetFromCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromCurrency, true
}

// SetFromCurrency sets field value
func (o *ExchangeCurrencyRequest) SetFromCurrency(v string) {
	o.FromCurrency = v
}

// GetCurrency returns the Currency field value
func (o *ExchangeCurrencyRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *ExchangeCurrencyRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *ExchangeCurrencyRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetBalance returns the Balance field value
func (o *ExchangeCurrencyRequest) GetBalance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *ExchangeCurrencyRequest) GetBalanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *ExchangeCurrencyRequest) SetBalance(v int32) {
	o.Balance = v
}

// GetRemark returns the Remark field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExchangeCurrencyRequest) GetRemark() string {
	if o == nil || IsNil(o.Remark.Get()) {
		var ret string
		return ret
	}
	return *o.Remark.Get()
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExchangeCurrencyRequest) GetRemarkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Remark.Get(), o.Remark.IsSet()
}

// HasRemark returns a boolean if a field has been set.
func (o *ExchangeCurrencyRequest) HasRemark() bool {
	if o != nil && o.Remark.IsSet() {
		return true
	}

	return false
}

// SetRemark gets a reference to the given NullableString and assigns it to the Remark field.
func (o *ExchangeCurrencyRequest) SetRemark(v string) {
	o.Remark.Set(&v)
}
// SetRemarkNil sets the value for Remark to be an explicit nil
func (o *ExchangeCurrencyRequest) SetRemarkNil() {
	o.Remark.Set(nil)
}

// UnsetRemark ensures that no value is present for Remark, not even an explicit nil
func (o *ExchangeCurrencyRequest) UnsetRemark() {
	o.Remark.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExchangeCurrencyRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExchangeCurrencyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ExchangeCurrencyRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ExchangeCurrencyRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ExchangeCurrencyRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ExchangeCurrencyRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExchangeCurrencyRequest) GetTags() string {
	if o == nil || IsNil(o.Tags.Get()) {
		var ret string
		return ret
	}
	return *o.Tags.Get()
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExchangeCurrencyRequest) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags.Get(), o.Tags.IsSet()
}

// HasTags returns a boolean if a field has been set.
func (o *ExchangeCurrencyRequest) HasTags() bool {
	if o != nil && o.Tags.IsSet() {
		return true
	}

	return false
}

// SetTags gets a reference to the given NullableString and assigns it to the Tags field.
func (o *ExchangeCurrencyRequest) SetTags(v string) {
	o.Tags.Set(&v)
}
// SetTagsNil sets the value for Tags to be an explicit nil
func (o *ExchangeCurrencyRequest) SetTagsNil() {
	o.Tags.Set(nil)
}

// UnsetTags ensures that no value is present for Tags, not even an explicit nil
func (o *ExchangeCurrencyRequest) UnsetTags() {
	o.Tags.Unset()
}

func (o ExchangeCurrencyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeCurrencyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fromCurrency"] = o.FromCurrency
	toSerialize["currency"] = o.Currency
	toSerialize["balance"] = o.Balance
	if o.Remark.IsSet() {
		toSerialize["remark"] = o.Remark.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Tags.IsSet() {
		toSerialize["tags"] = o.Tags.Get()
	}
	return toSerialize, nil
}

func (o *ExchangeCurrencyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fromCurrency",
		"currency",
		"balance",
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

	varExchangeCurrencyRequest := _ExchangeCurrencyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExchangeCurrencyRequest)

	if err != nil {
		return err
	}

	*o = ExchangeCurrencyRequest(varExchangeCurrencyRequest)

	return err
}

type NullableExchangeCurrencyRequest struct {
	value *ExchangeCurrencyRequest
	isSet bool
}

func (v NullableExchangeCurrencyRequest) Get() *ExchangeCurrencyRequest {
	return v.value
}

func (v *NullableExchangeCurrencyRequest) Set(val *ExchangeCurrencyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeCurrencyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeCurrencyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeCurrencyRequest(val *ExchangeCurrencyRequest) *NullableExchangeCurrencyRequest {
	return &NullableExchangeCurrencyRequest{value: val, isSet: true}
}

func (v NullableExchangeCurrencyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeCurrencyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


