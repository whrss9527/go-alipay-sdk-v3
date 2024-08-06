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

// checks if the VoucherPackageSalesLiteInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherPackageSalesLiteInfo{}

// VoucherPackageSalesLiteInfo struct for VoucherPackageSalesLiteInfo
type VoucherPackageSalesLiteInfo struct {
	// 券包预算，单位是份数
	Budget *int32 `json:"budget,omitempty"`
	// 券包售卖价格，单位是元
	SalePrice *string `json:"sale_price,omitempty"`
}

// NewVoucherPackageSalesLiteInfo instantiates a new VoucherPackageSalesLiteInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherPackageSalesLiteInfo() *VoucherPackageSalesLiteInfo {
	this := VoucherPackageSalesLiteInfo{}
	return &this
}

// NewVoucherPackageSalesLiteInfoWithDefaults instantiates a new VoucherPackageSalesLiteInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherPackageSalesLiteInfoWithDefaults() *VoucherPackageSalesLiteInfo {
	this := VoucherPackageSalesLiteInfo{}
	return &this
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *VoucherPackageSalesLiteInfo) GetBudget() int32 {
	if o == nil || IsNil(o.Budget) {
		var ret int32
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherPackageSalesLiteInfo) GetBudgetOk() (*int32, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *VoucherPackageSalesLiteInfo) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given int32 and assigns it to the Budget field.
func (o *VoucherPackageSalesLiteInfo) SetBudget(v int32) {
	o.Budget = &v
}

// GetSalePrice returns the SalePrice field value if set, zero value otherwise.
func (o *VoucherPackageSalesLiteInfo) GetSalePrice() string {
	if o == nil || IsNil(o.SalePrice) {
		var ret string
		return ret
	}
	return *o.SalePrice
}

// GetSalePriceOk returns a tuple with the SalePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherPackageSalesLiteInfo) GetSalePriceOk() (*string, bool) {
	if o == nil || IsNil(o.SalePrice) {
		return nil, false
	}
	return o.SalePrice, true
}

// HasSalePrice returns a boolean if a field has been set.
func (o *VoucherPackageSalesLiteInfo) HasSalePrice() bool {
	if o != nil && !IsNil(o.SalePrice) {
		return true
	}

	return false
}

// SetSalePrice gets a reference to the given string and assigns it to the SalePrice field.
func (o *VoucherPackageSalesLiteInfo) SetSalePrice(v string) {
	o.SalePrice = &v
}

func (o VoucherPackageSalesLiteInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherPackageSalesLiteInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Budget) {
		toSerialize["budget"] = o.Budget
	}
	if !IsNil(o.SalePrice) {
		toSerialize["sale_price"] = o.SalePrice
	}
	return toSerialize, nil
}

type NullableVoucherPackageSalesLiteInfo struct {
	value *VoucherPackageSalesLiteInfo
	isSet bool
}

func (v NullableVoucherPackageSalesLiteInfo) Get() *VoucherPackageSalesLiteInfo {
	return v.value
}

func (v *NullableVoucherPackageSalesLiteInfo) Set(val *VoucherPackageSalesLiteInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherPackageSalesLiteInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherPackageSalesLiteInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherPackageSalesLiteInfo(val *VoucherPackageSalesLiteInfo) *NullableVoucherPackageSalesLiteInfo {
	return &NullableVoucherPackageSalesLiteInfo{value: val, isSet: true}
}

func (v NullableVoucherPackageSalesLiteInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherPackageSalesLiteInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
