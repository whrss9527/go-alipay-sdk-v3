/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MaterialCreateInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaterialCreateInfo{}

// MaterialCreateInfo struct for MaterialCreateInfo
type MaterialCreateInfo struct {
	// 素材内容（素材地址或素材KEY）
	Content *string `json:"content,omitempty"`
	// 素材类型
	Type *string `json:"type,omitempty"`
}

// NewMaterialCreateInfo instantiates a new MaterialCreateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaterialCreateInfo() *MaterialCreateInfo {
	this := MaterialCreateInfo{}
	return &this
}

// NewMaterialCreateInfoWithDefaults instantiates a new MaterialCreateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaterialCreateInfoWithDefaults() *MaterialCreateInfo {
	this := MaterialCreateInfo{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *MaterialCreateInfo) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialCreateInfo) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *MaterialCreateInfo) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *MaterialCreateInfo) SetContent(v string) {
	o.Content = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MaterialCreateInfo) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialCreateInfo) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MaterialCreateInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MaterialCreateInfo) SetType(v string) {
	o.Type = &v
}

func (o MaterialCreateInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaterialCreateInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableMaterialCreateInfo struct {
	value *MaterialCreateInfo
	isSet bool
}

func (v NullableMaterialCreateInfo) Get() *MaterialCreateInfo {
	return v.value
}

func (v *NullableMaterialCreateInfo) Set(val *MaterialCreateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMaterialCreateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMaterialCreateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaterialCreateInfo(val *MaterialCreateInfo) *NullableMaterialCreateInfo {
	return &NullableMaterialCreateInfo{value: val, isSet: true}
}

func (v NullableMaterialCreateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaterialCreateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


