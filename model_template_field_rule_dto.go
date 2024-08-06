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

// checks if the TemplateFieldRuleDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateFieldRuleDTO{}

// TemplateFieldRuleDTO struct for TemplateFieldRuleDTO
type TemplateFieldRuleDTO struct {
	// 模板字段规则 字段名称(用于定义会员卡开卡接口时卡相应的参数) 余额：Balance 积分：Point 等级：Level 开卡日期：OpenDate 过期日期：ValidDate
	FieldName *string `json:"field_name,omitempty"`
	// 规则名  1、ASSIGN_FROM_REQUEST:   以rule_value为key值，表示该栏位的值从会员卡开卡接口中获取，会员卡开卡接口的card_info中获取对应参数值  2、DATE_IN_FUTURE: 生成一个未来的日期（格式YYYY-MM-DD)，当选择DATE_IN_FUTURE的时候，field_name  必须是OpenDate或ValidDate， 值为(10m或10d 分别表示10个月或10天)  3、CONST: 常量，会员卡开卡接口进行开卡的时候使用模板创建时候设置的值，即取rule_value的值
	RuleName *string `json:"rule_name,omitempty"`
	// 规则值，根据rule_name，采取相应取值策略  CONST：直接取rule_value作为卡属性值  DATE_IN_FUTURE：10m或10d 分别表示10个月或10天  ASSIGN_FROM_REQUEST：在开卡Reuqest请求中按rule_value取值，现在和field_name对应的为（OpenDate、ValidDate、Level、Point、Balance）
	RuleValue *string `json:"rule_value,omitempty"`
}

// NewTemplateFieldRuleDTO instantiates a new TemplateFieldRuleDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateFieldRuleDTO() *TemplateFieldRuleDTO {
	this := TemplateFieldRuleDTO{}
	return &this
}

// NewTemplateFieldRuleDTOWithDefaults instantiates a new TemplateFieldRuleDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateFieldRuleDTOWithDefaults() *TemplateFieldRuleDTO {
	this := TemplateFieldRuleDTO{}
	return &this
}

// GetFieldName returns the FieldName field value if set, zero value otherwise.
func (o *TemplateFieldRuleDTO) GetFieldName() string {
	if o == nil || IsNil(o.FieldName) {
		var ret string
		return ret
	}
	return *o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldRuleDTO) GetFieldNameOk() (*string, bool) {
	if o == nil || IsNil(o.FieldName) {
		return nil, false
	}
	return o.FieldName, true
}

// HasFieldName returns a boolean if a field has been set.
func (o *TemplateFieldRuleDTO) HasFieldName() bool {
	if o != nil && !IsNil(o.FieldName) {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given string and assigns it to the FieldName field.
func (o *TemplateFieldRuleDTO) SetFieldName(v string) {
	o.FieldName = &v
}

// GetRuleName returns the RuleName field value if set, zero value otherwise.
func (o *TemplateFieldRuleDTO) GetRuleName() string {
	if o == nil || IsNil(o.RuleName) {
		var ret string
		return ret
	}
	return *o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldRuleDTO) GetRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RuleName) {
		return nil, false
	}
	return o.RuleName, true
}

// HasRuleName returns a boolean if a field has been set.
func (o *TemplateFieldRuleDTO) HasRuleName() bool {
	if o != nil && !IsNil(o.RuleName) {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given string and assigns it to the RuleName field.
func (o *TemplateFieldRuleDTO) SetRuleName(v string) {
	o.RuleName = &v
}

// GetRuleValue returns the RuleValue field value if set, zero value otherwise.
func (o *TemplateFieldRuleDTO) GetRuleValue() string {
	if o == nil || IsNil(o.RuleValue) {
		var ret string
		return ret
	}
	return *o.RuleValue
}

// GetRuleValueOk returns a tuple with the RuleValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFieldRuleDTO) GetRuleValueOk() (*string, bool) {
	if o == nil || IsNil(o.RuleValue) {
		return nil, false
	}
	return o.RuleValue, true
}

// HasRuleValue returns a boolean if a field has been set.
func (o *TemplateFieldRuleDTO) HasRuleValue() bool {
	if o != nil && !IsNil(o.RuleValue) {
		return true
	}

	return false
}

// SetRuleValue gets a reference to the given string and assigns it to the RuleValue field.
func (o *TemplateFieldRuleDTO) SetRuleValue(v string) {
	o.RuleValue = &v
}

func (o TemplateFieldRuleDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateFieldRuleDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FieldName) {
		toSerialize["field_name"] = o.FieldName
	}
	if !IsNil(o.RuleName) {
		toSerialize["rule_name"] = o.RuleName
	}
	if !IsNil(o.RuleValue) {
		toSerialize["rule_value"] = o.RuleValue
	}
	return toSerialize, nil
}

type NullableTemplateFieldRuleDTO struct {
	value *TemplateFieldRuleDTO
	isSet bool
}

func (v NullableTemplateFieldRuleDTO) Get() *TemplateFieldRuleDTO {
	return v.value
}

func (v *NullableTemplateFieldRuleDTO) Set(val *TemplateFieldRuleDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateFieldRuleDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateFieldRuleDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateFieldRuleDTO(val *TemplateFieldRuleDTO) *NullableTemplateFieldRuleDTO {
	return &NullableTemplateFieldRuleDTO{value: val, isSet: true}
}

func (v NullableTemplateFieldRuleDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateFieldRuleDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


