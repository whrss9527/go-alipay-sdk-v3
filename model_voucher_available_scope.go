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

// checks if the VoucherAvailableScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherAvailableScope{}

// VoucherAvailableScope struct for VoucherAvailableScope
type VoucherAvailableScope struct {
	OrderVoucherAvailableCityCode *OrderVoucherAvailableCityCode `json:"order_voucher_available_city_code,omitempty"`
	OrderVoucherAvailableShop     *OrderVoucherAvailableShop     `json:"order_voucher_available_shop,omitempty"`
	// 可用范围类型。
	VoucherAvailableType *string `json:"voucher_available_type,omitempty"`
}

// NewVoucherAvailableScope instantiates a new VoucherAvailableScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherAvailableScope() *VoucherAvailableScope {
	this := VoucherAvailableScope{}
	return &this
}

// NewVoucherAvailableScopeWithDefaults instantiates a new VoucherAvailableScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherAvailableScopeWithDefaults() *VoucherAvailableScope {
	this := VoucherAvailableScope{}
	return &this
}

// GetOrderVoucherAvailableCityCode returns the OrderVoucherAvailableCityCode field value if set, zero value otherwise.
func (o *VoucherAvailableScope) GetOrderVoucherAvailableCityCode() OrderVoucherAvailableCityCode {
	if o == nil || IsNil(o.OrderVoucherAvailableCityCode) {
		var ret OrderVoucherAvailableCityCode
		return ret
	}
	return *o.OrderVoucherAvailableCityCode
}

// GetOrderVoucherAvailableCityCodeOk returns a tuple with the OrderVoucherAvailableCityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableScope) GetOrderVoucherAvailableCityCodeOk() (*OrderVoucherAvailableCityCode, bool) {
	if o == nil || IsNil(o.OrderVoucherAvailableCityCode) {
		return nil, false
	}
	return o.OrderVoucherAvailableCityCode, true
}

// HasOrderVoucherAvailableCityCode returns a boolean if a field has been set.
func (o *VoucherAvailableScope) HasOrderVoucherAvailableCityCode() bool {
	if o != nil && !IsNil(o.OrderVoucherAvailableCityCode) {
		return true
	}

	return false
}

// SetOrderVoucherAvailableCityCode gets a reference to the given OrderVoucherAvailableCityCode and assigns it to the OrderVoucherAvailableCityCode field.
func (o *VoucherAvailableScope) SetOrderVoucherAvailableCityCode(v OrderVoucherAvailableCityCode) {
	o.OrderVoucherAvailableCityCode = &v
}

// GetOrderVoucherAvailableShop returns the OrderVoucherAvailableShop field value if set, zero value otherwise.
func (o *VoucherAvailableScope) GetOrderVoucherAvailableShop() OrderVoucherAvailableShop {
	if o == nil || IsNil(o.OrderVoucherAvailableShop) {
		var ret OrderVoucherAvailableShop
		return ret
	}
	return *o.OrderVoucherAvailableShop
}

// GetOrderVoucherAvailableShopOk returns a tuple with the OrderVoucherAvailableShop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableScope) GetOrderVoucherAvailableShopOk() (*OrderVoucherAvailableShop, bool) {
	if o == nil || IsNil(o.OrderVoucherAvailableShop) {
		return nil, false
	}
	return o.OrderVoucherAvailableShop, true
}

// HasOrderVoucherAvailableShop returns a boolean if a field has been set.
func (o *VoucherAvailableScope) HasOrderVoucherAvailableShop() bool {
	if o != nil && !IsNil(o.OrderVoucherAvailableShop) {
		return true
	}

	return false
}

// SetOrderVoucherAvailableShop gets a reference to the given OrderVoucherAvailableShop and assigns it to the OrderVoucherAvailableShop field.
func (o *VoucherAvailableScope) SetOrderVoucherAvailableShop(v OrderVoucherAvailableShop) {
	o.OrderVoucherAvailableShop = &v
}

// GetVoucherAvailableType returns the VoucherAvailableType field value if set, zero value otherwise.
func (o *VoucherAvailableScope) GetVoucherAvailableType() string {
	if o == nil || IsNil(o.VoucherAvailableType) {
		var ret string
		return ret
	}
	return *o.VoucherAvailableType
}

// GetVoucherAvailableTypeOk returns a tuple with the VoucherAvailableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableScope) GetVoucherAvailableTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherAvailableType) {
		return nil, false
	}
	return o.VoucherAvailableType, true
}

// HasVoucherAvailableType returns a boolean if a field has been set.
func (o *VoucherAvailableScope) HasVoucherAvailableType() bool {
	if o != nil && !IsNil(o.VoucherAvailableType) {
		return true
	}

	return false
}

// SetVoucherAvailableType gets a reference to the given string and assigns it to the VoucherAvailableType field.
func (o *VoucherAvailableScope) SetVoucherAvailableType(v string) {
	o.VoucherAvailableType = &v
}

func (o VoucherAvailableScope) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherAvailableScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderVoucherAvailableCityCode) {
		toSerialize["order_voucher_available_city_code"] = o.OrderVoucherAvailableCityCode
	}
	if !IsNil(o.OrderVoucherAvailableShop) {
		toSerialize["order_voucher_available_shop"] = o.OrderVoucherAvailableShop
	}
	if !IsNil(o.VoucherAvailableType) {
		toSerialize["voucher_available_type"] = o.VoucherAvailableType
	}
	return toSerialize, nil
}

type NullableVoucherAvailableScope struct {
	value *VoucherAvailableScope
	isSet bool
}

func (v NullableVoucherAvailableScope) Get() *VoucherAvailableScope {
	return v.value
}

func (v *NullableVoucherAvailableScope) Set(val *VoucherAvailableScope) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherAvailableScope) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherAvailableScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherAvailableScope(val *VoucherAvailableScope) *NullableVoucherAvailableScope {
	return &NullableVoucherAvailableScope{value: val, isSet: true}
}

func (v NullableVoucherAvailableScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherAvailableScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
