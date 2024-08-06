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

// checks if the SearchPartAgreeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchPartAgreeInfo{}

// SearchPartAgreeInfo struct for SearchPartAgreeInfo
type SearchPartAgreeInfo struct {
	// 剔除数据名称
	AuditInfo *string `json:"audit_info,omitempty"`
	// 审核operator信息
	AuditOperator *string `json:"audit_operator,omitempty"`
	// 审核原因
	AuditReason *string `json:"audit_reason,omitempty"`
	// 剔除数据类型
	AuditType *string `json:"audit_type,omitempty"`
}

// NewSearchPartAgreeInfo instantiates a new SearchPartAgreeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchPartAgreeInfo() *SearchPartAgreeInfo {
	this := SearchPartAgreeInfo{}
	return &this
}

// NewSearchPartAgreeInfoWithDefaults instantiates a new SearchPartAgreeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchPartAgreeInfoWithDefaults() *SearchPartAgreeInfo {
	this := SearchPartAgreeInfo{}
	return &this
}

// GetAuditInfo returns the AuditInfo field value if set, zero value otherwise.
func (o *SearchPartAgreeInfo) GetAuditInfo() string {
	if o == nil || IsNil(o.AuditInfo) {
		var ret string
		return ret
	}
	return *o.AuditInfo
}

// GetAuditInfoOk returns a tuple with the AuditInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPartAgreeInfo) GetAuditInfoOk() (*string, bool) {
	if o == nil || IsNil(o.AuditInfo) {
		return nil, false
	}
	return o.AuditInfo, true
}

// HasAuditInfo returns a boolean if a field has been set.
func (o *SearchPartAgreeInfo) HasAuditInfo() bool {
	if o != nil && !IsNil(o.AuditInfo) {
		return true
	}

	return false
}

// SetAuditInfo gets a reference to the given string and assigns it to the AuditInfo field.
func (o *SearchPartAgreeInfo) SetAuditInfo(v string) {
	o.AuditInfo = &v
}

// GetAuditOperator returns the AuditOperator field value if set, zero value otherwise.
func (o *SearchPartAgreeInfo) GetAuditOperator() string {
	if o == nil || IsNil(o.AuditOperator) {
		var ret string
		return ret
	}
	return *o.AuditOperator
}

// GetAuditOperatorOk returns a tuple with the AuditOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPartAgreeInfo) GetAuditOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.AuditOperator) {
		return nil, false
	}
	return o.AuditOperator, true
}

// HasAuditOperator returns a boolean if a field has been set.
func (o *SearchPartAgreeInfo) HasAuditOperator() bool {
	if o != nil && !IsNil(o.AuditOperator) {
		return true
	}

	return false
}

// SetAuditOperator gets a reference to the given string and assigns it to the AuditOperator field.
func (o *SearchPartAgreeInfo) SetAuditOperator(v string) {
	o.AuditOperator = &v
}

// GetAuditReason returns the AuditReason field value if set, zero value otherwise.
func (o *SearchPartAgreeInfo) GetAuditReason() string {
	if o == nil || IsNil(o.AuditReason) {
		var ret string
		return ret
	}
	return *o.AuditReason
}

// GetAuditReasonOk returns a tuple with the AuditReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPartAgreeInfo) GetAuditReasonOk() (*string, bool) {
	if o == nil || IsNil(o.AuditReason) {
		return nil, false
	}
	return o.AuditReason, true
}

// HasAuditReason returns a boolean if a field has been set.
func (o *SearchPartAgreeInfo) HasAuditReason() bool {
	if o != nil && !IsNil(o.AuditReason) {
		return true
	}

	return false
}

// SetAuditReason gets a reference to the given string and assigns it to the AuditReason field.
func (o *SearchPartAgreeInfo) SetAuditReason(v string) {
	o.AuditReason = &v
}

// GetAuditType returns the AuditType field value if set, zero value otherwise.
func (o *SearchPartAgreeInfo) GetAuditType() string {
	if o == nil || IsNil(o.AuditType) {
		var ret string
		return ret
	}
	return *o.AuditType
}

// GetAuditTypeOk returns a tuple with the AuditType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPartAgreeInfo) GetAuditTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuditType) {
		return nil, false
	}
	return o.AuditType, true
}

// HasAuditType returns a boolean if a field has been set.
func (o *SearchPartAgreeInfo) HasAuditType() bool {
	if o != nil && !IsNil(o.AuditType) {
		return true
	}

	return false
}

// SetAuditType gets a reference to the given string and assigns it to the AuditType field.
func (o *SearchPartAgreeInfo) SetAuditType(v string) {
	o.AuditType = &v
}

func (o SearchPartAgreeInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchPartAgreeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuditInfo) {
		toSerialize["audit_info"] = o.AuditInfo
	}
	if !IsNil(o.AuditOperator) {
		toSerialize["audit_operator"] = o.AuditOperator
	}
	if !IsNil(o.AuditReason) {
		toSerialize["audit_reason"] = o.AuditReason
	}
	if !IsNil(o.AuditType) {
		toSerialize["audit_type"] = o.AuditType
	}
	return toSerialize, nil
}

type NullableSearchPartAgreeInfo struct {
	value *SearchPartAgreeInfo
	isSet bool
}

func (v NullableSearchPartAgreeInfo) Get() *SearchPartAgreeInfo {
	return v.value
}

func (v *NullableSearchPartAgreeInfo) Set(val *SearchPartAgreeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchPartAgreeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchPartAgreeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchPartAgreeInfo(val *SearchPartAgreeInfo) *NullableSearchPartAgreeInfo {
	return &NullableSearchPartAgreeInfo{value: val, isSet: true}
}

func (v NullableSearchPartAgreeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchPartAgreeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
