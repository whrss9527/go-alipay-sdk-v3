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

// checks if the RecruitEnrollRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecruitEnrollRule{}

// RecruitEnrollRule struct for RecruitEnrollRule
type RecruitEnrollRule struct {
	// 最大数量
	MaxSize *int32 `json:"max_size,omitempty"`
	// 最小数量
	MinSize *int32 `json:"min_size,omitempty"`
	// 是否必填
	Required *bool                  `json:"required,omitempty"`
	RuleData *RecruitEnrollRuleData `json:"rule_data,omitempty"`
	// 招商报名时提交的内容类型: 券:VOUCHER; 小程序:MINI_APP; 报名商户:ENROLL_MERCHANT; 素材:MATERIAL; 活动城市:CITY;
	Type *string `json:"type,omitempty"`
}

// NewRecruitEnrollRule instantiates a new RecruitEnrollRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecruitEnrollRule() *RecruitEnrollRule {
	this := RecruitEnrollRule{}
	return &this
}

// NewRecruitEnrollRuleWithDefaults instantiates a new RecruitEnrollRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecruitEnrollRuleWithDefaults() *RecruitEnrollRule {
	this := RecruitEnrollRule{}
	return &this
}

// GetMaxSize returns the MaxSize field value if set, zero value otherwise.
func (o *RecruitEnrollRule) GetMaxSize() int32 {
	if o == nil || IsNil(o.MaxSize) {
		var ret int32
		return ret
	}
	return *o.MaxSize
}

// GetMaxSizeOk returns a tuple with the MaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecruitEnrollRule) GetMaxSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxSize) {
		return nil, false
	}
	return o.MaxSize, true
}

// HasMaxSize returns a boolean if a field has been set.
func (o *RecruitEnrollRule) HasMaxSize() bool {
	if o != nil && !IsNil(o.MaxSize) {
		return true
	}

	return false
}

// SetMaxSize gets a reference to the given int32 and assigns it to the MaxSize field.
func (o *RecruitEnrollRule) SetMaxSize(v int32) {
	o.MaxSize = &v
}

// GetMinSize returns the MinSize field value if set, zero value otherwise.
func (o *RecruitEnrollRule) GetMinSize() int32 {
	if o == nil || IsNil(o.MinSize) {
		var ret int32
		return ret
	}
	return *o.MinSize
}

// GetMinSizeOk returns a tuple with the MinSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecruitEnrollRule) GetMinSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.MinSize) {
		return nil, false
	}
	return o.MinSize, true
}

// HasMinSize returns a boolean if a field has been set.
func (o *RecruitEnrollRule) HasMinSize() bool {
	if o != nil && !IsNil(o.MinSize) {
		return true
	}

	return false
}

// SetMinSize gets a reference to the given int32 and assigns it to the MinSize field.
func (o *RecruitEnrollRule) SetMinSize(v int32) {
	o.MinSize = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *RecruitEnrollRule) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecruitEnrollRule) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *RecruitEnrollRule) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *RecruitEnrollRule) SetRequired(v bool) {
	o.Required = &v
}

// GetRuleData returns the RuleData field value if set, zero value otherwise.
func (o *RecruitEnrollRule) GetRuleData() RecruitEnrollRuleData {
	if o == nil || IsNil(o.RuleData) {
		var ret RecruitEnrollRuleData
		return ret
	}
	return *o.RuleData
}

// GetRuleDataOk returns a tuple with the RuleData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecruitEnrollRule) GetRuleDataOk() (*RecruitEnrollRuleData, bool) {
	if o == nil || IsNil(o.RuleData) {
		return nil, false
	}
	return o.RuleData, true
}

// HasRuleData returns a boolean if a field has been set.
func (o *RecruitEnrollRule) HasRuleData() bool {
	if o != nil && !IsNil(o.RuleData) {
		return true
	}

	return false
}

// SetRuleData gets a reference to the given RecruitEnrollRuleData and assigns it to the RuleData field.
func (o *RecruitEnrollRule) SetRuleData(v RecruitEnrollRuleData) {
	o.RuleData = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RecruitEnrollRule) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecruitEnrollRule) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RecruitEnrollRule) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RecruitEnrollRule) SetType(v string) {
	o.Type = &v
}

func (o RecruitEnrollRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecruitEnrollRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxSize) {
		toSerialize["max_size"] = o.MaxSize
	}
	if !IsNil(o.MinSize) {
		toSerialize["min_size"] = o.MinSize
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.RuleData) {
		toSerialize["rule_data"] = o.RuleData
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableRecruitEnrollRule struct {
	value *RecruitEnrollRule
	isSet bool
}

func (v NullableRecruitEnrollRule) Get() *RecruitEnrollRule {
	return v.value
}

func (v *NullableRecruitEnrollRule) Set(val *RecruitEnrollRule) {
	v.value = val
	v.isSet = true
}

func (v NullableRecruitEnrollRule) IsSet() bool {
	return v.isSet
}

func (v *NullableRecruitEnrollRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecruitEnrollRule(val *RecruitEnrollRule) *NullableRecruitEnrollRule {
	return &NullableRecruitEnrollRule{value: val, isSet: true}
}

func (v NullableRecruitEnrollRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecruitEnrollRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
