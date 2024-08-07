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

// checks if the DeductInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeductInfo{}

// DeductInfo struct for DeductInfo
type DeductInfo struct {
	CustomerDefineDeductRule *CustomerDefineDeductRule `json:"customer_define_deduct_rule,omitempty"`
}

// NewDeductInfo instantiates a new DeductInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeductInfo() *DeductInfo {
	this := DeductInfo{}
	return &this
}

// NewDeductInfoWithDefaults instantiates a new DeductInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeductInfoWithDefaults() *DeductInfo {
	this := DeductInfo{}
	return &this
}

// GetCustomerDefineDeductRule returns the CustomerDefineDeductRule field value if set, zero value otherwise.
func (o *DeductInfo) GetCustomerDefineDeductRule() CustomerDefineDeductRule {
	if o == nil || IsNil(o.CustomerDefineDeductRule) {
		var ret CustomerDefineDeductRule
		return ret
	}
	return *o.CustomerDefineDeductRule
}

// GetCustomerDefineDeductRuleOk returns a tuple with the CustomerDefineDeductRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeductInfo) GetCustomerDefineDeductRuleOk() (*CustomerDefineDeductRule, bool) {
	if o == nil || IsNil(o.CustomerDefineDeductRule) {
		return nil, false
	}
	return o.CustomerDefineDeductRule, true
}

// HasCustomerDefineDeductRule returns a boolean if a field has been set.
func (o *DeductInfo) HasCustomerDefineDeductRule() bool {
	if o != nil && !IsNil(o.CustomerDefineDeductRule) {
		return true
	}

	return false
}

// SetCustomerDefineDeductRule gets a reference to the given CustomerDefineDeductRule and assigns it to the CustomerDefineDeductRule field.
func (o *DeductInfo) SetCustomerDefineDeductRule(v CustomerDefineDeductRule) {
	o.CustomerDefineDeductRule = &v
}

func (o DeductInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeductInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerDefineDeductRule) {
		toSerialize["customer_define_deduct_rule"] = o.CustomerDefineDeductRule
	}
	return toSerialize, nil
}

type NullableDeductInfo struct {
	value *DeductInfo
	isSet bool
}

func (v NullableDeductInfo) Get() *DeductInfo {
	return v.value
}

func (v *NullableDeductInfo) Set(val *DeductInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDeductInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDeductInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeductInfo(val *DeductInfo) *NullableDeductInfo {
	return &NullableDeductInfo{value: val, isSet: true}
}

func (v NullableDeductInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeductInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
