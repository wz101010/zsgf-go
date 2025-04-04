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

// checks if the GoodsDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GoodsDetail{}

// GoodsDetail struct for GoodsDetail
type GoodsDetail struct {
	AlipayGoodsId NullableString `json:"alipayGoodsId,omitempty"`
	Body NullableString `json:"body,omitempty"`
	CategoriesTree NullableString `json:"categoriesTree,omitempty"`
	GoodsCategory NullableString `json:"goodsCategory,omitempty"`
	GoodsId NullableString `json:"goodsId,omitempty"`
	GoodsName NullableString `json:"goodsName,omitempty"`
	OutItemId NullableString `json:"outItemId,omitempty"`
	OutSkuId NullableString `json:"outSkuId,omitempty"`
	Price NullableString `json:"price,omitempty"`
	Quantity *int64 `json:"quantity,omitempty"`
	ShowUrl NullableString `json:"showUrl,omitempty"`
}

// NewGoodsDetail instantiates a new GoodsDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoodsDetail() *GoodsDetail {
	this := GoodsDetail{}
	return &this
}

// NewGoodsDetailWithDefaults instantiates a new GoodsDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoodsDetailWithDefaults() *GoodsDetail {
	this := GoodsDetail{}
	return &this
}

// GetAlipayGoodsId returns the AlipayGoodsId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoodsDetail) GetAlipayGoodsId() string {
	if o == nil || IsNil(o.AlipayGoodsId.Get()) {
		var ret string
		return ret
	}
	return *o.AlipayGoodsId.Get()
}

// GetAlipayGoodsIdOk returns a tuple with the AlipayGoodsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoodsDetail) GetAlipayGoodsIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlipayGoodsId.Get(), o.AlipayGoodsId.IsSet()
}

// HasAlipayGoodsId returns a boolean if a field has been set.
func (o *GoodsDetail) HasAlipayGoodsId() bool {
	if o != nil && o.AlipayGoodsId.IsSet() {
		return true
	}

	return false
}

// SetAlipayGoodsId gets a reference to the given NullableString and assigns it to the AlipayGoodsId field.
func (o *GoodsDetail) SetAlipayGoodsId(v string) {
	o.AlipayGoodsId.Set(&v)
}
// SetAlipayGoodsIdNil sets the value for AlipayGoodsId to be an explicit nil
func (o *GoodsDetail) SetAlipayGoodsIdNil() {
	o.AlipayGoodsId.Set(nil)
}

// UnsetAlipayGoodsId ensures that no value is present for AlipayGoodsId, not even an explicit nil
func (o *GoodsDetail) UnsetAlipayGoodsId() {
	o.AlipayGoodsId.Unset()
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoodsDetail) GetBody() string {
	if o == nil || IsNil(o.Body.Get()) {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoodsDetail) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *GoodsDetail) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *GoodsDetail) SetBody(v string) {
	o.Body.Set(&v)
}
// SetBodyNil sets the value for Body to be an explicit nil
func (o *GoodsDetail) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *GoodsDetail) UnsetBody() {
	o.Body.Unset()
}

// GetCategoriesTree returns the CategoriesTree field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoodsDetail) GetCategoriesTree() string {
	if o == nil || IsNil(o.CategoriesTree.Get()) {
		var ret string
		return ret
	}
	return *o.CategoriesTree.Get()
}

// GetCategoriesTreeOk returns a tuple with the CategoriesTree field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoodsDetail) GetCategoriesTreeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CategoriesTree.Get(), o.CategoriesTree.IsSet()
}

// HasCategoriesTree returns a boolean if a field has been set.
func (o *GoodsDetail) HasCategoriesTree() bool {
	if o != nil && o.CategoriesTree.IsSet() {
		return true
	}

	return false
}

// SetCategoriesTree gets a reference to the given NullableString and assigns it to the CategoriesTree field.
func (o *GoodsDetail) SetCategoriesTree(v string) {
	o.CategoriesTree.Set(&v)
}
// SetCategoriesTreeNil sets the value for CategoriesTree to be an explicit nil
func (o *GoodsDetail) SetCategoriesTreeNil() {
	o.CategoriesTree.Set(nil)
}

// UnsetCategoriesTree ensures that no value is present for CategoriesTree, not even an explicit nil
func (o *GoodsDetail) UnsetCategoriesTree() {
	o.CategoriesTree.Unset()
}

// GetGoodsCategory returns the GoodsCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoodsDetail) GetGoodsCategory() string {
	if o == nil || IsNil(o.GoodsCategory.Get()) {
		var ret string
		return ret
	}
	return *o.GoodsCategory.Get()
}

// GetGoodsCategoryOk returns a tuple with the GoodsCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoodsDetail) GetGoodsCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GoodsCategory.Get(), o.GoodsCategory.IsSet()
}

// HasGoodsCategory returns a boolean if a field has been set.
func (o *GoodsDetail) HasGoodsCategory() bool {
	if o != nil && o.GoodsCategory.IsSet() {
		return true
	}

	return false
}

// SetGoodsCategory gets a reference to the given NullableString and assigns it to the GoodsCategory field.
func (o *GoodsDetail) SetGoodsCategory(v string) {
	o.GoodsCategory.Set(&v)
}
// SetGoodsCategoryNil sets the value for GoodsCategory to be an explicit nil
func (o *GoodsDetail) SetGoodsCategoryNil() {
	o.GoodsCategory.Set(nil)
}

// UnsetGoodsCategory ensures that no value is present for GoodsCategory, not even an explicit nil
func (o *GoodsDetail) UnsetGoodsCategory() {
	o.GoodsCategory.Unset()
}

// GetGoodsId returns the GoodsId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoodsDetail) GetGoodsId() string {
	if o == nil || IsNil(o.GoodsId.Get()) {
		var ret string
		return ret
	}
	return *o.GoodsId.Get()
}

// GetGoodsIdOk returns a tuple with the GoodsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoodsDetail) GetGoodsIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GoodsId.Get(), o.GoodsId.IsSet()
}

// HasGoodsId returns a boolean if a field has been set.
func (o *GoodsDetail) HasGoodsId() bool {
	if o != nil && o.GoodsId.IsSet() {
		return true
	}

	return false
}

// SetGoodsId gets a reference to the given NullableString and assigns it to the GoodsId field.
func (o *GoodsDetail) SetGoodsId(v string) {
	o.GoodsId.Set(&v)
}
// SetGoodsIdNil sets the value for GoodsId to be an explicit nil
func (o *GoodsDetail) SetGoodsIdNil() {
	o.GoodsId.Set(nil)
}

// UnsetGoodsId ensures that no value is present for GoodsId, not even an explicit nil
func (o *GoodsDetail) UnsetGoodsId() {
	o.GoodsId.Unset()
}

// GetGoodsName returns the GoodsName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoodsDetail) GetGoodsName() string {
	if o == nil || IsNil(o.GoodsName.Get()) {
		var ret string
		return ret
	}
	return *o.GoodsName.Get()
}

// GetGoodsNameOk returns a tuple with the GoodsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoodsDetail) GetGoodsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GoodsName.Get(), o.GoodsName.IsSet()
}

// HasGoodsName returns a boolean if a field has been set.
func (o *GoodsDetail) HasGoodsName() bool {
	if o != nil && o.GoodsName.IsSet() {
		return true
	}

	return false
}

// SetGoodsName gets a reference to the given NullableString and assigns it to the GoodsName field.
func (o *GoodsDetail) SetGoodsName(v string) {
	o.GoodsName.Set(&v)
}
// SetGoodsNameNil sets the value for GoodsName to be an explicit nil
func (o *GoodsDetail) SetGoodsNameNil() {
	o.GoodsName.Set(nil)
}

// UnsetGoodsName ensures that no value is present for GoodsName, not even an explicit nil
func (o *GoodsDetail) UnsetGoodsName() {
	o.GoodsName.Unset()
}

// GetOutItemId returns the OutItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoodsDetail) GetOutItemId() string {
	if o == nil || IsNil(o.OutItemId.Get()) {
		var ret string
		return ret
	}
	return *o.OutItemId.Get()
}

// GetOutItemIdOk returns a tuple with the OutItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoodsDetail) GetOutItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutItemId.Get(), o.OutItemId.IsSet()
}

// HasOutItemId returns a boolean if a field has been set.
func (o *GoodsDetail) HasOutItemId() bool {
	if o != nil && o.OutItemId.IsSet() {
		return true
	}

	return false
}

// SetOutItemId gets a reference to the given NullableString and assigns it to the OutItemId field.
func (o *GoodsDetail) SetOutItemId(v string) {
	o.OutItemId.Set(&v)
}
// SetOutItemIdNil sets the value for OutItemId to be an explicit nil
func (o *GoodsDetail) SetOutItemIdNil() {
	o.OutItemId.Set(nil)
}

// UnsetOutItemId ensures that no value is present for OutItemId, not even an explicit nil
func (o *GoodsDetail) UnsetOutItemId() {
	o.OutItemId.Unset()
}

// GetOutSkuId returns the OutSkuId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoodsDetail) GetOutSkuId() string {
	if o == nil || IsNil(o.OutSkuId.Get()) {
		var ret string
		return ret
	}
	return *o.OutSkuId.Get()
}

// GetOutSkuIdOk returns a tuple with the OutSkuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoodsDetail) GetOutSkuIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutSkuId.Get(), o.OutSkuId.IsSet()
}

// HasOutSkuId returns a boolean if a field has been set.
func (o *GoodsDetail) HasOutSkuId() bool {
	if o != nil && o.OutSkuId.IsSet() {
		return true
	}

	return false
}

// SetOutSkuId gets a reference to the given NullableString and assigns it to the OutSkuId field.
func (o *GoodsDetail) SetOutSkuId(v string) {
	o.OutSkuId.Set(&v)
}
// SetOutSkuIdNil sets the value for OutSkuId to be an explicit nil
func (o *GoodsDetail) SetOutSkuIdNil() {
	o.OutSkuId.Set(nil)
}

// UnsetOutSkuId ensures that no value is present for OutSkuId, not even an explicit nil
func (o *GoodsDetail) UnsetOutSkuId() {
	o.OutSkuId.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoodsDetail) GetPrice() string {
	if o == nil || IsNil(o.Price.Get()) {
		var ret string
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoodsDetail) GetPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *GoodsDetail) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableString and assigns it to the Price field.
func (o *GoodsDetail) SetPrice(v string) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *GoodsDetail) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *GoodsDetail) UnsetPrice() {
	o.Price.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *GoodsDetail) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoodsDetail) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *GoodsDetail) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *GoodsDetail) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetShowUrl returns the ShowUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GoodsDetail) GetShowUrl() string {
	if o == nil || IsNil(o.ShowUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ShowUrl.Get()
}

// GetShowUrlOk returns a tuple with the ShowUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GoodsDetail) GetShowUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShowUrl.Get(), o.ShowUrl.IsSet()
}

// HasShowUrl returns a boolean if a field has been set.
func (o *GoodsDetail) HasShowUrl() bool {
	if o != nil && o.ShowUrl.IsSet() {
		return true
	}

	return false
}

// SetShowUrl gets a reference to the given NullableString and assigns it to the ShowUrl field.
func (o *GoodsDetail) SetShowUrl(v string) {
	o.ShowUrl.Set(&v)
}
// SetShowUrlNil sets the value for ShowUrl to be an explicit nil
func (o *GoodsDetail) SetShowUrlNil() {
	o.ShowUrl.Set(nil)
}

// UnsetShowUrl ensures that no value is present for ShowUrl, not even an explicit nil
func (o *GoodsDetail) UnsetShowUrl() {
	o.ShowUrl.Unset()
}

func (o GoodsDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GoodsDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AlipayGoodsId.IsSet() {
		toSerialize["alipayGoodsId"] = o.AlipayGoodsId.Get()
	}
	if o.Body.IsSet() {
		toSerialize["body"] = o.Body.Get()
	}
	if o.CategoriesTree.IsSet() {
		toSerialize["categoriesTree"] = o.CategoriesTree.Get()
	}
	if o.GoodsCategory.IsSet() {
		toSerialize["goodsCategory"] = o.GoodsCategory.Get()
	}
	if o.GoodsId.IsSet() {
		toSerialize["goodsId"] = o.GoodsId.Get()
	}
	if o.GoodsName.IsSet() {
		toSerialize["goodsName"] = o.GoodsName.Get()
	}
	if o.OutItemId.IsSet() {
		toSerialize["outItemId"] = o.OutItemId.Get()
	}
	if o.OutSkuId.IsSet() {
		toSerialize["outSkuId"] = o.OutSkuId.Get()
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if o.ShowUrl.IsSet() {
		toSerialize["showUrl"] = o.ShowUrl.Get()
	}
	return toSerialize, nil
}

type NullableGoodsDetail struct {
	value *GoodsDetail
	isSet bool
}

func (v NullableGoodsDetail) Get() *GoodsDetail {
	return v.value
}

func (v *NullableGoodsDetail) Set(val *GoodsDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableGoodsDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableGoodsDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoodsDetail(val *GoodsDetail) *NullableGoodsDetail {
	return &NullableGoodsDetail{value: val, isSet: true}
}

func (v NullableGoodsDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoodsDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


