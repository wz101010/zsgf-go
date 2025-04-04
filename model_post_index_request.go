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

// checks if the PostIndexRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostIndexRequest{}

// PostIndexRequest struct for PostIndexRequest
type PostIndexRequest struct {
	Model map[string][]interface{} `json:"model"`
	Options map[string][]interface{} `json:"options,omitempty"`
}

type _PostIndexRequest PostIndexRequest

// NewPostIndexRequest instantiates a new PostIndexRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostIndexRequest(model map[string][]interface{}) *PostIndexRequest {
	this := PostIndexRequest{}
	this.Model = model
	return &this
}

// NewPostIndexRequestWithDefaults instantiates a new PostIndexRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostIndexRequestWithDefaults() *PostIndexRequest {
	this := PostIndexRequest{}
	return &this
}

// GetModel returns the Model field value
func (o *PostIndexRequest) GetModel() map[string][]interface{} {
	if o == nil {
		var ret map[string][]interface{}
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *PostIndexRequest) GetModelOk() (*map[string][]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *PostIndexRequest) SetModel(v map[string][]interface{}) {
	o.Model = v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostIndexRequest) GetOptions() map[string][]interface{} {
	if o == nil {
		var ret map[string][]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostIndexRequest) GetOptionsOk() (*map[string][]interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *PostIndexRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string][]interface{} and assigns it to the Options field.
func (o *PostIndexRequest) SetOptions(v map[string][]interface{}) {
	o.Options = v
}

func (o PostIndexRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostIndexRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["model"] = o.Model
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

func (o *PostIndexRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"model",
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

	varPostIndexRequest := _PostIndexRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostIndexRequest)

	if err != nil {
		return err
	}

	*o = PostIndexRequest(varPostIndexRequest)

	return err
}

type NullablePostIndexRequest struct {
	value *PostIndexRequest
	isSet bool
}

func (v NullablePostIndexRequest) Get() *PostIndexRequest {
	return v.value
}

func (v *NullablePostIndexRequest) Set(val *PostIndexRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostIndexRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostIndexRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostIndexRequest(val *PostIndexRequest) *NullablePostIndexRequest {
	return &NullablePostIndexRequest{value: val, isSet: true}
}

func (v NullablePostIndexRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostIndexRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


