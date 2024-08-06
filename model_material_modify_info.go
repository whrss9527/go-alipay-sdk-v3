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

// checks if the MaterialModifyInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaterialModifyInfo{}

// MaterialModifyInfo struct for MaterialModifyInfo
type MaterialModifyInfo struct {
	// 素材内容（素材地址或素材KEY）
	Content *string `json:"content,omitempty"`
	// 素材类型
	Type *string `json:"type,omitempty"`
}

// NewMaterialModifyInfo instantiates a new MaterialModifyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaterialModifyInfo() *MaterialModifyInfo {
	this := MaterialModifyInfo{}
	return &this
}

// NewMaterialModifyInfoWithDefaults instantiates a new MaterialModifyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaterialModifyInfoWithDefaults() *MaterialModifyInfo {
	this := MaterialModifyInfo{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *MaterialModifyInfo) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialModifyInfo) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *MaterialModifyInfo) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *MaterialModifyInfo) SetContent(v string) {
	o.Content = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MaterialModifyInfo) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaterialModifyInfo) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MaterialModifyInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MaterialModifyInfo) SetType(v string) {
	o.Type = &v
}

func (o MaterialModifyInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaterialModifyInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableMaterialModifyInfo struct {
	value *MaterialModifyInfo
	isSet bool
}

func (v NullableMaterialModifyInfo) Get() *MaterialModifyInfo {
	return v.value
}

func (v *NullableMaterialModifyInfo) Set(val *MaterialModifyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMaterialModifyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMaterialModifyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaterialModifyInfo(val *MaterialModifyInfo) *NullableMaterialModifyInfo {
	return &NullableMaterialModifyInfo{value: val, isSet: true}
}

func (v NullableMaterialModifyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaterialModifyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


