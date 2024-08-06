/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

import (
	"encoding/json"
)

// checks if the AlipayOpenPublicDefaultExtensionCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicDefaultExtensionCreateModel{}

// AlipayOpenPublicDefaultExtensionCreateModel struct for AlipayOpenPublicDefaultExtensionCreateModel
type AlipayOpenPublicDefaultExtensionCreateModel struct {
	// 默认扩展区列表，最多包含3个扩展区
	Areas []ExtensionArea `json:"areas,omitempty"`
}

// NewAlipayOpenPublicDefaultExtensionCreateModel instantiates a new AlipayOpenPublicDefaultExtensionCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicDefaultExtensionCreateModel() *AlipayOpenPublicDefaultExtensionCreateModel {
	this := AlipayOpenPublicDefaultExtensionCreateModel{}
	return &this
}

// NewAlipayOpenPublicDefaultExtensionCreateModelWithDefaults instantiates a new AlipayOpenPublicDefaultExtensionCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicDefaultExtensionCreateModelWithDefaults() *AlipayOpenPublicDefaultExtensionCreateModel {
	this := AlipayOpenPublicDefaultExtensionCreateModel{}
	return &this
}

// GetAreas returns the Areas field value if set, zero value otherwise.
func (o *AlipayOpenPublicDefaultExtensionCreateModel) GetAreas() []ExtensionArea {
	if o == nil || IsNil(o.Areas) {
		var ret []ExtensionArea
		return ret
	}
	return o.Areas
}

// GetAreasOk returns a tuple with the Areas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicDefaultExtensionCreateModel) GetAreasOk() ([]ExtensionArea, bool) {
	if o == nil || IsNil(o.Areas) {
		return nil, false
	}
	return o.Areas, true
}

// HasAreas returns a boolean if a field has been set.
func (o *AlipayOpenPublicDefaultExtensionCreateModel) HasAreas() bool {
	if o != nil && !IsNil(o.Areas) {
		return true
	}

	return false
}

// SetAreas gets a reference to the given []ExtensionArea and assigns it to the Areas field.
func (o *AlipayOpenPublicDefaultExtensionCreateModel) SetAreas(v []ExtensionArea) {
	o.Areas = v
}

func (o AlipayOpenPublicDefaultExtensionCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicDefaultExtensionCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Areas) {
		toSerialize["areas"] = o.Areas
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicDefaultExtensionCreateModel struct {
	value *AlipayOpenPublicDefaultExtensionCreateModel
	isSet bool
}

func (v NullableAlipayOpenPublicDefaultExtensionCreateModel) Get() *AlipayOpenPublicDefaultExtensionCreateModel {
	return v.value
}

func (v *NullableAlipayOpenPublicDefaultExtensionCreateModel) Set(val *AlipayOpenPublicDefaultExtensionCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicDefaultExtensionCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicDefaultExtensionCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicDefaultExtensionCreateModel(val *AlipayOpenPublicDefaultExtensionCreateModel) *NullableAlipayOpenPublicDefaultExtensionCreateModel {
	return &NullableAlipayOpenPublicDefaultExtensionCreateModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicDefaultExtensionCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicDefaultExtensionCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
