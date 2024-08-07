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

// checks if the PaymentVoucherBudgetInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentVoucherBudgetInfo{}

// PaymentVoucherBudgetInfo struct for PaymentVoucherBudgetInfo
type PaymentVoucherBudgetInfo struct {
	// 营销资金预算,单位元。 总预算=优惠金额*总发券张数
	Amount *string `json:"amount,omitempty"`
	// 营销资金预算类型 1、免充值：制券方无需提前充值资金，用户核销支付券时，直接从订单原价中扣除优惠减价金额，最终只将用户实际支付的金额结算给核销商户，商户实收少于订单原价。 2.预充值:  制券方需将优惠预算提前充值到支付宝指定营销账户中，用户核销支付券时，系统从该账户可用余额中扣除优惠减价部分对应的资金，连同用户实际支付的资金，一并结算给核销商户，不影响实收。
	BudgetType   *string                     `json:"budget_type,omitempty"`
	RechargeInfo *PaymentVoucherRechargeInfo `json:"recharge_info,omitempty"`
}

// NewPaymentVoucherBudgetInfo instantiates a new PaymentVoucherBudgetInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentVoucherBudgetInfo() *PaymentVoucherBudgetInfo {
	this := PaymentVoucherBudgetInfo{}
	return &this
}

// NewPaymentVoucherBudgetInfoWithDefaults instantiates a new PaymentVoucherBudgetInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentVoucherBudgetInfoWithDefaults() *PaymentVoucherBudgetInfo {
	this := PaymentVoucherBudgetInfo{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PaymentVoucherBudgetInfo) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherBudgetInfo) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PaymentVoucherBudgetInfo) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *PaymentVoucherBudgetInfo) SetAmount(v string) {
	o.Amount = &v
}

// GetBudgetType returns the BudgetType field value if set, zero value otherwise.
func (o *PaymentVoucherBudgetInfo) GetBudgetType() string {
	if o == nil || IsNil(o.BudgetType) {
		var ret string
		return ret
	}
	return *o.BudgetType
}

// GetBudgetTypeOk returns a tuple with the BudgetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherBudgetInfo) GetBudgetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BudgetType) {
		return nil, false
	}
	return o.BudgetType, true
}

// HasBudgetType returns a boolean if a field has been set.
func (o *PaymentVoucherBudgetInfo) HasBudgetType() bool {
	if o != nil && !IsNil(o.BudgetType) {
		return true
	}

	return false
}

// SetBudgetType gets a reference to the given string and assigns it to the BudgetType field.
func (o *PaymentVoucherBudgetInfo) SetBudgetType(v string) {
	o.BudgetType = &v
}

// GetRechargeInfo returns the RechargeInfo field value if set, zero value otherwise.
func (o *PaymentVoucherBudgetInfo) GetRechargeInfo() PaymentVoucherRechargeInfo {
	if o == nil || IsNil(o.RechargeInfo) {
		var ret PaymentVoucherRechargeInfo
		return ret
	}
	return *o.RechargeInfo
}

// GetRechargeInfoOk returns a tuple with the RechargeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherBudgetInfo) GetRechargeInfoOk() (*PaymentVoucherRechargeInfo, bool) {
	if o == nil || IsNil(o.RechargeInfo) {
		return nil, false
	}
	return o.RechargeInfo, true
}

// HasRechargeInfo returns a boolean if a field has been set.
func (o *PaymentVoucherBudgetInfo) HasRechargeInfo() bool {
	if o != nil && !IsNil(o.RechargeInfo) {
		return true
	}

	return false
}

// SetRechargeInfo gets a reference to the given PaymentVoucherRechargeInfo and assigns it to the RechargeInfo field.
func (o *PaymentVoucherBudgetInfo) SetRechargeInfo(v PaymentVoucherRechargeInfo) {
	o.RechargeInfo = &v
}

func (o PaymentVoucherBudgetInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentVoucherBudgetInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.BudgetType) {
		toSerialize["budget_type"] = o.BudgetType
	}
	if !IsNil(o.RechargeInfo) {
		toSerialize["recharge_info"] = o.RechargeInfo
	}
	return toSerialize, nil
}

type NullablePaymentVoucherBudgetInfo struct {
	value *PaymentVoucherBudgetInfo
	isSet bool
}

func (v NullablePaymentVoucherBudgetInfo) Get() *PaymentVoucherBudgetInfo {
	return v.value
}

func (v *NullablePaymentVoucherBudgetInfo) Set(val *PaymentVoucherBudgetInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentVoucherBudgetInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentVoucherBudgetInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentVoucherBudgetInfo(val *PaymentVoucherBudgetInfo) *NullablePaymentVoucherBudgetInfo {
	return &NullablePaymentVoucherBudgetInfo{value: val, isSet: true}
}

func (v NullablePaymentVoucherBudgetInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentVoucherBudgetInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
