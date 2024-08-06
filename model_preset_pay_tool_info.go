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

// checks if the PresetPayToolInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PresetPayToolInfo{}

// PresetPayToolInfo struct for PresetPayToolInfo
type PresetPayToolInfo struct {
	// 前置资产金额。单位：元。
	Amount []string `json:"amount,omitempty"`
	// 前置资产类型编码，和收单支付传入的preset_pay_tool里面的类型编码保持一致。
	AssertTypeCode *string `json:"assert_type_code,omitempty"`
}

// NewPresetPayToolInfo instantiates a new PresetPayToolInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPresetPayToolInfo() *PresetPayToolInfo {
	this := PresetPayToolInfo{}
	return &this
}

// NewPresetPayToolInfoWithDefaults instantiates a new PresetPayToolInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPresetPayToolInfoWithDefaults() *PresetPayToolInfo {
	this := PresetPayToolInfo{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PresetPayToolInfo) GetAmount() []string {
	if o == nil || IsNil(o.Amount) {
		var ret []string
		return ret
	}
	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresetPayToolInfo) GetAmountOk() ([]string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PresetPayToolInfo) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given []string and assigns it to the Amount field.
func (o *PresetPayToolInfo) SetAmount(v []string) {
	o.Amount = v
}

// GetAssertTypeCode returns the AssertTypeCode field value if set, zero value otherwise.
func (o *PresetPayToolInfo) GetAssertTypeCode() string {
	if o == nil || IsNil(o.AssertTypeCode) {
		var ret string
		return ret
	}
	return *o.AssertTypeCode
}

// GetAssertTypeCodeOk returns a tuple with the AssertTypeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresetPayToolInfo) GetAssertTypeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AssertTypeCode) {
		return nil, false
	}
	return o.AssertTypeCode, true
}

// HasAssertTypeCode returns a boolean if a field has been set.
func (o *PresetPayToolInfo) HasAssertTypeCode() bool {
	if o != nil && !IsNil(o.AssertTypeCode) {
		return true
	}

	return false
}

// SetAssertTypeCode gets a reference to the given string and assigns it to the AssertTypeCode field.
func (o *PresetPayToolInfo) SetAssertTypeCode(v string) {
	o.AssertTypeCode = &v
}

func (o PresetPayToolInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PresetPayToolInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.AssertTypeCode) {
		toSerialize["assert_type_code"] = o.AssertTypeCode
	}
	return toSerialize, nil
}

type NullablePresetPayToolInfo struct {
	value *PresetPayToolInfo
	isSet bool
}

func (v NullablePresetPayToolInfo) Get() *PresetPayToolInfo {
	return v.value
}

func (v *NullablePresetPayToolInfo) Set(val *PresetPayToolInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePresetPayToolInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePresetPayToolInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePresetPayToolInfo(val *PresetPayToolInfo) *NullablePresetPayToolInfo {
	return &NullablePresetPayToolInfo{value: val, isSet: true}
}

func (v NullablePresetPayToolInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePresetPayToolInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


