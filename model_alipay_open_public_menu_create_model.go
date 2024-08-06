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

// checks if the AlipayOpenPublicMenuCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicMenuCreateModel{}

// AlipayOpenPublicMenuCreateModel struct for AlipayOpenPublicMenuCreateModel
type AlipayOpenPublicMenuCreateModel struct {
	// 一级菜单列表。说明： * 如果是文本菜单，最多有4个一级菜单，若开发者在后台打开了\"咨询反馈\"的开关，则只能有3个一级菜单。 * 如果是 ICON 菜单信息，最多有80个一级菜单(忽略二级菜单)。
	Button []ButtonObject `json:"button,omitempty"`
	// 菜单类型，不填时默认为 text（文本型菜单）。枚举值如下： * text：文本型菜单。 * icon：表示 icon 型菜单，当传值为\"icon\"时，菜单节点的icon字段必传。
	Type *string `json:"type,omitempty"`
}

// NewAlipayOpenPublicMenuCreateModel instantiates a new AlipayOpenPublicMenuCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicMenuCreateModel() *AlipayOpenPublicMenuCreateModel {
	this := AlipayOpenPublicMenuCreateModel{}
	return &this
}

// NewAlipayOpenPublicMenuCreateModelWithDefaults instantiates a new AlipayOpenPublicMenuCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicMenuCreateModelWithDefaults() *AlipayOpenPublicMenuCreateModel {
	this := AlipayOpenPublicMenuCreateModel{}
	return &this
}

// GetButton returns the Button field value if set, zero value otherwise.
func (o *AlipayOpenPublicMenuCreateModel) GetButton() []ButtonObject {
	if o == nil || IsNil(o.Button) {
		var ret []ButtonObject
		return ret
	}
	return o.Button
}

// GetButtonOk returns a tuple with the Button field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMenuCreateModel) GetButtonOk() ([]ButtonObject, bool) {
	if o == nil || IsNil(o.Button) {
		return nil, false
	}
	return o.Button, true
}

// HasButton returns a boolean if a field has been set.
func (o *AlipayOpenPublicMenuCreateModel) HasButton() bool {
	if o != nil && !IsNil(o.Button) {
		return true
	}

	return false
}

// SetButton gets a reference to the given []ButtonObject and assigns it to the Button field.
func (o *AlipayOpenPublicMenuCreateModel) SetButton(v []ButtonObject) {
	o.Button = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AlipayOpenPublicMenuCreateModel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicMenuCreateModel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AlipayOpenPublicMenuCreateModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AlipayOpenPublicMenuCreateModel) SetType(v string) {
	o.Type = &v
}

func (o AlipayOpenPublicMenuCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicMenuCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Button) {
		toSerialize["button"] = o.Button
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicMenuCreateModel struct {
	value *AlipayOpenPublicMenuCreateModel
	isSet bool
}

func (v NullableAlipayOpenPublicMenuCreateModel) Get() *AlipayOpenPublicMenuCreateModel {
	return v.value
}

func (v *NullableAlipayOpenPublicMenuCreateModel) Set(val *AlipayOpenPublicMenuCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicMenuCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicMenuCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicMenuCreateModel(val *AlipayOpenPublicMenuCreateModel) *NullableAlipayOpenPublicMenuCreateModel {
	return &NullableAlipayOpenPublicMenuCreateModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicMenuCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicMenuCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

