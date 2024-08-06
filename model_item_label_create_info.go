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

// checks if the ItemLabelCreateInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemLabelCreateInfo{}

// ItemLabelCreateInfo struct for ItemLabelCreateInfo
type ItemLabelCreateInfo struct {
	// 标签键  OIL_NUM：油号
	LabelKey *string `json:"label_key,omitempty"`
	// 标签值：  OIL_NUM对应的值为：95# 92# 98# 等等
	LabelValue *string `json:"label_value,omitempty"`
}

// NewItemLabelCreateInfo instantiates a new ItemLabelCreateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemLabelCreateInfo() *ItemLabelCreateInfo {
	this := ItemLabelCreateInfo{}
	return &this
}

// NewItemLabelCreateInfoWithDefaults instantiates a new ItemLabelCreateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemLabelCreateInfoWithDefaults() *ItemLabelCreateInfo {
	this := ItemLabelCreateInfo{}
	return &this
}

// GetLabelKey returns the LabelKey field value if set, zero value otherwise.
func (o *ItemLabelCreateInfo) GetLabelKey() string {
	if o == nil || IsNil(o.LabelKey) {
		var ret string
		return ret
	}
	return *o.LabelKey
}

// GetLabelKeyOk returns a tuple with the LabelKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemLabelCreateInfo) GetLabelKeyOk() (*string, bool) {
	if o == nil || IsNil(o.LabelKey) {
		return nil, false
	}
	return o.LabelKey, true
}

// HasLabelKey returns a boolean if a field has been set.
func (o *ItemLabelCreateInfo) HasLabelKey() bool {
	if o != nil && !IsNil(o.LabelKey) {
		return true
	}

	return false
}

// SetLabelKey gets a reference to the given string and assigns it to the LabelKey field.
func (o *ItemLabelCreateInfo) SetLabelKey(v string) {
	o.LabelKey = &v
}

// GetLabelValue returns the LabelValue field value if set, zero value otherwise.
func (o *ItemLabelCreateInfo) GetLabelValue() string {
	if o == nil || IsNil(o.LabelValue) {
		var ret string
		return ret
	}
	return *o.LabelValue
}

// GetLabelValueOk returns a tuple with the LabelValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemLabelCreateInfo) GetLabelValueOk() (*string, bool) {
	if o == nil || IsNil(o.LabelValue) {
		return nil, false
	}
	return o.LabelValue, true
}

// HasLabelValue returns a boolean if a field has been set.
func (o *ItemLabelCreateInfo) HasLabelValue() bool {
	if o != nil && !IsNil(o.LabelValue) {
		return true
	}

	return false
}

// SetLabelValue gets a reference to the given string and assigns it to the LabelValue field.
func (o *ItemLabelCreateInfo) SetLabelValue(v string) {
	o.LabelValue = &v
}

func (o ItemLabelCreateInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemLabelCreateInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LabelKey) {
		toSerialize["label_key"] = o.LabelKey
	}
	if !IsNil(o.LabelValue) {
		toSerialize["label_value"] = o.LabelValue
	}
	return toSerialize, nil
}

type NullableItemLabelCreateInfo struct {
	value *ItemLabelCreateInfo
	isSet bool
}

func (v NullableItemLabelCreateInfo) Get() *ItemLabelCreateInfo {
	return v.value
}

func (v *NullableItemLabelCreateInfo) Set(val *ItemLabelCreateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableItemLabelCreateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableItemLabelCreateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemLabelCreateInfo(val *ItemLabelCreateInfo) *NullableItemLabelCreateInfo {
	return &NullableItemLabelCreateInfo{value: val, isSet: true}
}

func (v NullableItemLabelCreateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemLabelCreateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


