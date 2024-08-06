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

// checks if the AlipayOpenPublicGroupCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicGroupCreateModel{}

// AlipayOpenPublicGroupCreateModel struct for AlipayOpenPublicGroupCreateModel
type AlipayOpenPublicGroupCreateModel struct {
	// 标签规则，满足该规则的粉丝将被圈定，标签id不能重复
	LabelRule []ComplexLabelRule `json:"label_rule,omitempty"`
	// 分组名称，仅支持中文、字母、数字、下划线的组合。
	Name *string `json:"name,omitempty"`
}

// NewAlipayOpenPublicGroupCreateModel instantiates a new AlipayOpenPublicGroupCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicGroupCreateModel() *AlipayOpenPublicGroupCreateModel {
	this := AlipayOpenPublicGroupCreateModel{}
	return &this
}

// NewAlipayOpenPublicGroupCreateModelWithDefaults instantiates a new AlipayOpenPublicGroupCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicGroupCreateModelWithDefaults() *AlipayOpenPublicGroupCreateModel {
	this := AlipayOpenPublicGroupCreateModel{}
	return &this
}

// GetLabelRule returns the LabelRule field value if set, zero value otherwise.
func (o *AlipayOpenPublicGroupCreateModel) GetLabelRule() []ComplexLabelRule {
	if o == nil || IsNil(o.LabelRule) {
		var ret []ComplexLabelRule
		return ret
	}
	return o.LabelRule
}

// GetLabelRuleOk returns a tuple with the LabelRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicGroupCreateModel) GetLabelRuleOk() ([]ComplexLabelRule, bool) {
	if o == nil || IsNil(o.LabelRule) {
		return nil, false
	}
	return o.LabelRule, true
}

// HasLabelRule returns a boolean if a field has been set.
func (o *AlipayOpenPublicGroupCreateModel) HasLabelRule() bool {
	if o != nil && !IsNil(o.LabelRule) {
		return true
	}

	return false
}

// SetLabelRule gets a reference to the given []ComplexLabelRule and assigns it to the LabelRule field.
func (o *AlipayOpenPublicGroupCreateModel) SetLabelRule(v []ComplexLabelRule) {
	o.LabelRule = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlipayOpenPublicGroupCreateModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicGroupCreateModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlipayOpenPublicGroupCreateModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlipayOpenPublicGroupCreateModel) SetName(v string) {
	o.Name = &v
}

func (o AlipayOpenPublicGroupCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicGroupCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LabelRule) {
		toSerialize["label_rule"] = o.LabelRule
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicGroupCreateModel struct {
	value *AlipayOpenPublicGroupCreateModel
	isSet bool
}

func (v NullableAlipayOpenPublicGroupCreateModel) Get() *AlipayOpenPublicGroupCreateModel {
	return v.value
}

func (v *NullableAlipayOpenPublicGroupCreateModel) Set(val *AlipayOpenPublicGroupCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicGroupCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicGroupCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicGroupCreateModel(val *AlipayOpenPublicGroupCreateModel) *NullableAlipayOpenPublicGroupCreateModel {
	return &NullableAlipayOpenPublicGroupCreateModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicGroupCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicGroupCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

