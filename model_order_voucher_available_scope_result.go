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

// checks if the OrderVoucherAvailableScopeResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderVoucherAvailableScopeResult{}

// OrderVoucherAvailableScopeResult struct for OrderVoucherAvailableScopeResult
type OrderVoucherAvailableScopeResult struct {
	OrderVoucherAvailableShopResult *OrderVoucherAvailableShopResult `json:"order_voucher_available_shop_result,omitempty"`
	// 可用范围类型。
	VoucherAvailableType *string `json:"voucher_available_type,omitempty"`
}

// NewOrderVoucherAvailableScopeResult instantiates a new OrderVoucherAvailableScopeResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderVoucherAvailableScopeResult() *OrderVoucherAvailableScopeResult {
	this := OrderVoucherAvailableScopeResult{}
	return &this
}

// NewOrderVoucherAvailableScopeResultWithDefaults instantiates a new OrderVoucherAvailableScopeResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderVoucherAvailableScopeResultWithDefaults() *OrderVoucherAvailableScopeResult {
	this := OrderVoucherAvailableScopeResult{}
	return &this
}

// GetOrderVoucherAvailableShopResult returns the OrderVoucherAvailableShopResult field value if set, zero value otherwise.
func (o *OrderVoucherAvailableScopeResult) GetOrderVoucherAvailableShopResult() OrderVoucherAvailableShopResult {
	if o == nil || IsNil(o.OrderVoucherAvailableShopResult) {
		var ret OrderVoucherAvailableShopResult
		return ret
	}
	return *o.OrderVoucherAvailableShopResult
}

// GetOrderVoucherAvailableShopResultOk returns a tuple with the OrderVoucherAvailableShopResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderVoucherAvailableScopeResult) GetOrderVoucherAvailableShopResultOk() (*OrderVoucherAvailableShopResult, bool) {
	if o == nil || IsNil(o.OrderVoucherAvailableShopResult) {
		return nil, false
	}
	return o.OrderVoucherAvailableShopResult, true
}

// HasOrderVoucherAvailableShopResult returns a boolean if a field has been set.
func (o *OrderVoucherAvailableScopeResult) HasOrderVoucherAvailableShopResult() bool {
	if o != nil && !IsNil(o.OrderVoucherAvailableShopResult) {
		return true
	}

	return false
}

// SetOrderVoucherAvailableShopResult gets a reference to the given OrderVoucherAvailableShopResult and assigns it to the OrderVoucherAvailableShopResult field.
func (o *OrderVoucherAvailableScopeResult) SetOrderVoucherAvailableShopResult(v OrderVoucherAvailableShopResult) {
	o.OrderVoucherAvailableShopResult = &v
}

// GetVoucherAvailableType returns the VoucherAvailableType field value if set, zero value otherwise.
func (o *OrderVoucherAvailableScopeResult) GetVoucherAvailableType() string {
	if o == nil || IsNil(o.VoucherAvailableType) {
		var ret string
		return ret
	}
	return *o.VoucherAvailableType
}

// GetVoucherAvailableTypeOk returns a tuple with the VoucherAvailableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderVoucherAvailableScopeResult) GetVoucherAvailableTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherAvailableType) {
		return nil, false
	}
	return o.VoucherAvailableType, true
}

// HasVoucherAvailableType returns a boolean if a field has been set.
func (o *OrderVoucherAvailableScopeResult) HasVoucherAvailableType() bool {
	if o != nil && !IsNil(o.VoucherAvailableType) {
		return true
	}

	return false
}

// SetVoucherAvailableType gets a reference to the given string and assigns it to the VoucherAvailableType field.
func (o *OrderVoucherAvailableScopeResult) SetVoucherAvailableType(v string) {
	o.VoucherAvailableType = &v
}

func (o OrderVoucherAvailableScopeResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderVoucherAvailableScopeResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderVoucherAvailableShopResult) {
		toSerialize["order_voucher_available_shop_result"] = o.OrderVoucherAvailableShopResult
	}
	if !IsNil(o.VoucherAvailableType) {
		toSerialize["voucher_available_type"] = o.VoucherAvailableType
	}
	return toSerialize, nil
}

type NullableOrderVoucherAvailableScopeResult struct {
	value *OrderVoucherAvailableScopeResult
	isSet bool
}

func (v NullableOrderVoucherAvailableScopeResult) Get() *OrderVoucherAvailableScopeResult {
	return v.value
}

func (v *NullableOrderVoucherAvailableScopeResult) Set(val *OrderVoucherAvailableScopeResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderVoucherAvailableScopeResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderVoucherAvailableScopeResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderVoucherAvailableScopeResult(val *OrderVoucherAvailableScopeResult) *NullableOrderVoucherAvailableScopeResult {
	return &NullableOrderVoucherAvailableScopeResult{value: val, isSet: true}
}

func (v NullableOrderVoucherAvailableScopeResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderVoucherAvailableScopeResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


