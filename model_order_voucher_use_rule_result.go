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

// checks if the OrderVoucherUseRuleResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderVoucherUseRuleResult{}

// OrderVoucherUseRuleResult struct for OrderVoucherUseRuleResult
type OrderVoucherUseRuleResult struct {
	VoucherAvailableScopeResult *OrderVoucherAvailableScopeResult `json:"voucher_available_scope_result,omitempty"`
}

// NewOrderVoucherUseRuleResult instantiates a new OrderVoucherUseRuleResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderVoucherUseRuleResult() *OrderVoucherUseRuleResult {
	this := OrderVoucherUseRuleResult{}
	return &this
}

// NewOrderVoucherUseRuleResultWithDefaults instantiates a new OrderVoucherUseRuleResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderVoucherUseRuleResultWithDefaults() *OrderVoucherUseRuleResult {
	this := OrderVoucherUseRuleResult{}
	return &this
}

// GetVoucherAvailableScopeResult returns the VoucherAvailableScopeResult field value if set, zero value otherwise.
func (o *OrderVoucherUseRuleResult) GetVoucherAvailableScopeResult() OrderVoucherAvailableScopeResult {
	if o == nil || IsNil(o.VoucherAvailableScopeResult) {
		var ret OrderVoucherAvailableScopeResult
		return ret
	}
	return *o.VoucherAvailableScopeResult
}

// GetVoucherAvailableScopeResultOk returns a tuple with the VoucherAvailableScopeResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderVoucherUseRuleResult) GetVoucherAvailableScopeResultOk() (*OrderVoucherAvailableScopeResult, bool) {
	if o == nil || IsNil(o.VoucherAvailableScopeResult) {
		return nil, false
	}
	return o.VoucherAvailableScopeResult, true
}

// HasVoucherAvailableScopeResult returns a boolean if a field has been set.
func (o *OrderVoucherUseRuleResult) HasVoucherAvailableScopeResult() bool {
	if o != nil && !IsNil(o.VoucherAvailableScopeResult) {
		return true
	}

	return false
}

// SetVoucherAvailableScopeResult gets a reference to the given OrderVoucherAvailableScopeResult and assigns it to the VoucherAvailableScopeResult field.
func (o *OrderVoucherUseRuleResult) SetVoucherAvailableScopeResult(v OrderVoucherAvailableScopeResult) {
	o.VoucherAvailableScopeResult = &v
}

func (o OrderVoucherUseRuleResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderVoucherUseRuleResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VoucherAvailableScopeResult) {
		toSerialize["voucher_available_scope_result"] = o.VoucherAvailableScopeResult
	}
	return toSerialize, nil
}

type NullableOrderVoucherUseRuleResult struct {
	value *OrderVoucherUseRuleResult
	isSet bool
}

func (v NullableOrderVoucherUseRuleResult) Get() *OrderVoucherUseRuleResult {
	return v.value
}

func (v *NullableOrderVoucherUseRuleResult) Set(val *OrderVoucherUseRuleResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderVoucherUseRuleResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderVoucherUseRuleResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderVoucherUseRuleResult(val *OrderVoucherUseRuleResult) *NullableOrderVoucherUseRuleResult {
	return &NullableOrderVoucherUseRuleResult{value: val, isSet: true}
}

func (v NullableOrderVoucherUseRuleResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderVoucherUseRuleResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
