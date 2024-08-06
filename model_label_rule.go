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

// checks if the LabelRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelRule{}

// LabelRule struct for LabelRule
type LabelRule struct {
	// 标签id
	LabelId *string `json:"label_id,omitempty"`
	// 标签值，当有多个取值时用英文\",\"分隔，不允许传入下划线\"_\"、竖线\"|\"或者空格\" \"和方括号\"[\"、\"]\"
	LabelValue *string `json:"label_value,omitempty"`
	// 目前支持EQ（等于）、BETWEEN（范围）、IN（包含）三种操作符；每个标签支持的运算符可以通过<a href=\"https://docs.open.alipay.com/api_6/alipay.open.public.life.label.batchquery#sintq\">标签列表查询接口</a>获得。该字段允许为空，默认运算符为IN
	Operator *string `json:"operator,omitempty"`
}

// NewLabelRule instantiates a new LabelRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelRule() *LabelRule {
	this := LabelRule{}
	return &this
}

// NewLabelRuleWithDefaults instantiates a new LabelRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelRuleWithDefaults() *LabelRule {
	this := LabelRule{}
	return &this
}

// GetLabelId returns the LabelId field value if set, zero value otherwise.
func (o *LabelRule) GetLabelId() string {
	if o == nil || IsNil(o.LabelId) {
		var ret string
		return ret
	}
	return *o.LabelId
}

// GetLabelIdOk returns a tuple with the LabelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelRule) GetLabelIdOk() (*string, bool) {
	if o == nil || IsNil(o.LabelId) {
		return nil, false
	}
	return o.LabelId, true
}

// HasLabelId returns a boolean if a field has been set.
func (o *LabelRule) HasLabelId() bool {
	if o != nil && !IsNil(o.LabelId) {
		return true
	}

	return false
}

// SetLabelId gets a reference to the given string and assigns it to the LabelId field.
func (o *LabelRule) SetLabelId(v string) {
	o.LabelId = &v
}

// GetLabelValue returns the LabelValue field value if set, zero value otherwise.
func (o *LabelRule) GetLabelValue() string {
	if o == nil || IsNil(o.LabelValue) {
		var ret string
		return ret
	}
	return *o.LabelValue
}

// GetLabelValueOk returns a tuple with the LabelValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelRule) GetLabelValueOk() (*string, bool) {
	if o == nil || IsNil(o.LabelValue) {
		return nil, false
	}
	return o.LabelValue, true
}

// HasLabelValue returns a boolean if a field has been set.
func (o *LabelRule) HasLabelValue() bool {
	if o != nil && !IsNil(o.LabelValue) {
		return true
	}

	return false
}

// SetLabelValue gets a reference to the given string and assigns it to the LabelValue field.
func (o *LabelRule) SetLabelValue(v string) {
	o.LabelValue = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *LabelRule) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelRule) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *LabelRule) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *LabelRule) SetOperator(v string) {
	o.Operator = &v
}

func (o LabelRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LabelId) {
		toSerialize["label_id"] = o.LabelId
	}
	if !IsNil(o.LabelValue) {
		toSerialize["label_value"] = o.LabelValue
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	return toSerialize, nil
}

type NullableLabelRule struct {
	value *LabelRule
	isSet bool
}

func (v NullableLabelRule) Get() *LabelRule {
	return v.value
}

func (v *NullableLabelRule) Set(val *LabelRule) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelRule) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelRule(val *LabelRule) *NullableLabelRule {
	return &NullableLabelRule{value: val, isSet: true}
}

func (v NullableLabelRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


