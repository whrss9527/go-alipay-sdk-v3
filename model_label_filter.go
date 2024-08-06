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

// checks if the LabelFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelFilter{}

// LabelFilter struct for LabelFilter
type LabelFilter struct {
	// 标签组名，商户自定义的标签固定为label_id_list，支付宝开放的标签详见<a href=\"http://aopsdkdownload.cn-hangzhou.alipay-pub.aliyun-inc.com/doc/FirstPartOfTagsFromAlipay.xlsx\">支付宝开放标签</a>
	ColumnName *string `json:"column_name,omitempty"`
	// 操作符，支持=、!=、in三个操作符；其中in表示是某几个标签中的一个
	Op *string `json:"op,omitempty"`
	// 标签数组，用于组装最后的表达式
	Values []string `json:"values,omitempty"`
}

// NewLabelFilter instantiates a new LabelFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelFilter() *LabelFilter {
	this := LabelFilter{}
	return &this
}

// NewLabelFilterWithDefaults instantiates a new LabelFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelFilterWithDefaults() *LabelFilter {
	this := LabelFilter{}
	return &this
}

// GetColumnName returns the ColumnName field value if set, zero value otherwise.
func (o *LabelFilter) GetColumnName() string {
	if o == nil || IsNil(o.ColumnName) {
		var ret string
		return ret
	}
	return *o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelFilter) GetColumnNameOk() (*string, bool) {
	if o == nil || IsNil(o.ColumnName) {
		return nil, false
	}
	return o.ColumnName, true
}

// HasColumnName returns a boolean if a field has been set.
func (o *LabelFilter) HasColumnName() bool {
	if o != nil && !IsNil(o.ColumnName) {
		return true
	}

	return false
}

// SetColumnName gets a reference to the given string and assigns it to the ColumnName field.
func (o *LabelFilter) SetColumnName(v string) {
	o.ColumnName = &v
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *LabelFilter) GetOp() string {
	if o == nil || IsNil(o.Op) {
		var ret string
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelFilter) GetOpOk() (*string, bool) {
	if o == nil || IsNil(o.Op) {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *LabelFilter) HasOp() bool {
	if o != nil && !IsNil(o.Op) {
		return true
	}

	return false
}

// SetOp gets a reference to the given string and assigns it to the Op field.
func (o *LabelFilter) SetOp(v string) {
	o.Op = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *LabelFilter) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelFilter) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *LabelFilter) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *LabelFilter) SetValues(v []string) {
	o.Values = v
}

func (o LabelFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ColumnName) {
		toSerialize["column_name"] = o.ColumnName
	}
	if !IsNil(o.Op) {
		toSerialize["op"] = o.Op
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableLabelFilter struct {
	value *LabelFilter
	isSet bool
}

func (v NullableLabelFilter) Get() *LabelFilter {
	return v.value
}

func (v *NullableLabelFilter) Set(val *LabelFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelFilter(val *LabelFilter) *NullableLabelFilter {
	return &NullableLabelFilter{value: val, isSet: true}
}

func (v NullableLabelFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

