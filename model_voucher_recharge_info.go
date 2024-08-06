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

// checks if the VoucherRechargeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherRechargeInfo{}

// VoucherRechargeInfo struct for VoucherRechargeInfo
type VoucherRechargeInfo struct {
	// 预充值方式。
	RechargeType               *string                     `json:"recharge_type,omitempty"`
	VoucherBalanceRechargeInfo *VoucherBalanceRechargeInfo `json:"voucher_balance_recharge_info,omitempty"`
}

// NewVoucherRechargeInfo instantiates a new VoucherRechargeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherRechargeInfo() *VoucherRechargeInfo {
	this := VoucherRechargeInfo{}
	return &this
}

// NewVoucherRechargeInfoWithDefaults instantiates a new VoucherRechargeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherRechargeInfoWithDefaults() *VoucherRechargeInfo {
	this := VoucherRechargeInfo{}
	return &this
}

// GetRechargeType returns the RechargeType field value if set, zero value otherwise.
func (o *VoucherRechargeInfo) GetRechargeType() string {
	if o == nil || IsNil(o.RechargeType) {
		var ret string
		return ret
	}
	return *o.RechargeType
}

// GetRechargeTypeOk returns a tuple with the RechargeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherRechargeInfo) GetRechargeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RechargeType) {
		return nil, false
	}
	return o.RechargeType, true
}

// HasRechargeType returns a boolean if a field has been set.
func (o *VoucherRechargeInfo) HasRechargeType() bool {
	if o != nil && !IsNil(o.RechargeType) {
		return true
	}

	return false
}

// SetRechargeType gets a reference to the given string and assigns it to the RechargeType field.
func (o *VoucherRechargeInfo) SetRechargeType(v string) {
	o.RechargeType = &v
}

// GetVoucherBalanceRechargeInfo returns the VoucherBalanceRechargeInfo field value if set, zero value otherwise.
func (o *VoucherRechargeInfo) GetVoucherBalanceRechargeInfo() VoucherBalanceRechargeInfo {
	if o == nil || IsNil(o.VoucherBalanceRechargeInfo) {
		var ret VoucherBalanceRechargeInfo
		return ret
	}
	return *o.VoucherBalanceRechargeInfo
}

// GetVoucherBalanceRechargeInfoOk returns a tuple with the VoucherBalanceRechargeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherRechargeInfo) GetVoucherBalanceRechargeInfoOk() (*VoucherBalanceRechargeInfo, bool) {
	if o == nil || IsNil(o.VoucherBalanceRechargeInfo) {
		return nil, false
	}
	return o.VoucherBalanceRechargeInfo, true
}

// HasVoucherBalanceRechargeInfo returns a boolean if a field has been set.
func (o *VoucherRechargeInfo) HasVoucherBalanceRechargeInfo() bool {
	if o != nil && !IsNil(o.VoucherBalanceRechargeInfo) {
		return true
	}

	return false
}

// SetVoucherBalanceRechargeInfo gets a reference to the given VoucherBalanceRechargeInfo and assigns it to the VoucherBalanceRechargeInfo field.
func (o *VoucherRechargeInfo) SetVoucherBalanceRechargeInfo(v VoucherBalanceRechargeInfo) {
	o.VoucherBalanceRechargeInfo = &v
}

func (o VoucherRechargeInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherRechargeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RechargeType) {
		toSerialize["recharge_type"] = o.RechargeType
	}
	if !IsNil(o.VoucherBalanceRechargeInfo) {
		toSerialize["voucher_balance_recharge_info"] = o.VoucherBalanceRechargeInfo
	}
	return toSerialize, nil
}

type NullableVoucherRechargeInfo struct {
	value *VoucherRechargeInfo
	isSet bool
}

func (v NullableVoucherRechargeInfo) Get() *VoucherRechargeInfo {
	return v.value
}

func (v *NullableVoucherRechargeInfo) Set(val *VoucherRechargeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherRechargeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherRechargeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherRechargeInfo(val *VoucherRechargeInfo) *NullableVoucherRechargeInfo {
	return &NullableVoucherRechargeInfo{value: val, isSet: true}
}

func (v NullableVoucherRechargeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherRechargeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
