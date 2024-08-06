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

// checks if the CommonVoucherUseRuleLiteInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonVoucherUseRuleLiteInfo{}

// CommonVoucherUseRuleLiteInfo struct for CommonVoucherUseRuleLiteInfo
type CommonVoucherUseRuleLiteInfo struct {
	DiscountVoucher *ActivityDiscountVoucher `json:"discount_voucher,omitempty"`
	ExchangeVoucher *ActivityExchangeVoucher `json:"exchange_voucher,omitempty"`
	FixVoucher *ActivityFixVoucher `json:"fix_voucher,omitempty"`
	SpecialVoucher *ActivitySpecialVoucher `json:"special_voucher,omitempty"`
}

// NewCommonVoucherUseRuleLiteInfo instantiates a new CommonVoucherUseRuleLiteInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonVoucherUseRuleLiteInfo() *CommonVoucherUseRuleLiteInfo {
	this := CommonVoucherUseRuleLiteInfo{}
	return &this
}

// NewCommonVoucherUseRuleLiteInfoWithDefaults instantiates a new CommonVoucherUseRuleLiteInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonVoucherUseRuleLiteInfoWithDefaults() *CommonVoucherUseRuleLiteInfo {
	this := CommonVoucherUseRuleLiteInfo{}
	return &this
}

// GetDiscountVoucher returns the DiscountVoucher field value if set, zero value otherwise.
func (o *CommonVoucherUseRuleLiteInfo) GetDiscountVoucher() ActivityDiscountVoucher {
	if o == nil || IsNil(o.DiscountVoucher) {
		var ret ActivityDiscountVoucher
		return ret
	}
	return *o.DiscountVoucher
}

// GetDiscountVoucherOk returns a tuple with the DiscountVoucher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonVoucherUseRuleLiteInfo) GetDiscountVoucherOk() (*ActivityDiscountVoucher, bool) {
	if o == nil || IsNil(o.DiscountVoucher) {
		return nil, false
	}
	return o.DiscountVoucher, true
}

// HasDiscountVoucher returns a boolean if a field has been set.
func (o *CommonVoucherUseRuleLiteInfo) HasDiscountVoucher() bool {
	if o != nil && !IsNil(o.DiscountVoucher) {
		return true
	}

	return false
}

// SetDiscountVoucher gets a reference to the given ActivityDiscountVoucher and assigns it to the DiscountVoucher field.
func (o *CommonVoucherUseRuleLiteInfo) SetDiscountVoucher(v ActivityDiscountVoucher) {
	o.DiscountVoucher = &v
}

// GetExchangeVoucher returns the ExchangeVoucher field value if set, zero value otherwise.
func (o *CommonVoucherUseRuleLiteInfo) GetExchangeVoucher() ActivityExchangeVoucher {
	if o == nil || IsNil(o.ExchangeVoucher) {
		var ret ActivityExchangeVoucher
		return ret
	}
	return *o.ExchangeVoucher
}

// GetExchangeVoucherOk returns a tuple with the ExchangeVoucher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonVoucherUseRuleLiteInfo) GetExchangeVoucherOk() (*ActivityExchangeVoucher, bool) {
	if o == nil || IsNil(o.ExchangeVoucher) {
		return nil, false
	}
	return o.ExchangeVoucher, true
}

// HasExchangeVoucher returns a boolean if a field has been set.
func (o *CommonVoucherUseRuleLiteInfo) HasExchangeVoucher() bool {
	if o != nil && !IsNil(o.ExchangeVoucher) {
		return true
	}

	return false
}

// SetExchangeVoucher gets a reference to the given ActivityExchangeVoucher and assigns it to the ExchangeVoucher field.
func (o *CommonVoucherUseRuleLiteInfo) SetExchangeVoucher(v ActivityExchangeVoucher) {
	o.ExchangeVoucher = &v
}

// GetFixVoucher returns the FixVoucher field value if set, zero value otherwise.
func (o *CommonVoucherUseRuleLiteInfo) GetFixVoucher() ActivityFixVoucher {
	if o == nil || IsNil(o.FixVoucher) {
		var ret ActivityFixVoucher
		return ret
	}
	return *o.FixVoucher
}

// GetFixVoucherOk returns a tuple with the FixVoucher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonVoucherUseRuleLiteInfo) GetFixVoucherOk() (*ActivityFixVoucher, bool) {
	if o == nil || IsNil(o.FixVoucher) {
		return nil, false
	}
	return o.FixVoucher, true
}

// HasFixVoucher returns a boolean if a field has been set.
func (o *CommonVoucherUseRuleLiteInfo) HasFixVoucher() bool {
	if o != nil && !IsNil(o.FixVoucher) {
		return true
	}

	return false
}

// SetFixVoucher gets a reference to the given ActivityFixVoucher and assigns it to the FixVoucher field.
func (o *CommonVoucherUseRuleLiteInfo) SetFixVoucher(v ActivityFixVoucher) {
	o.FixVoucher = &v
}

// GetSpecialVoucher returns the SpecialVoucher field value if set, zero value otherwise.
func (o *CommonVoucherUseRuleLiteInfo) GetSpecialVoucher() ActivitySpecialVoucher {
	if o == nil || IsNil(o.SpecialVoucher) {
		var ret ActivitySpecialVoucher
		return ret
	}
	return *o.SpecialVoucher
}

// GetSpecialVoucherOk returns a tuple with the SpecialVoucher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonVoucherUseRuleLiteInfo) GetSpecialVoucherOk() (*ActivitySpecialVoucher, bool) {
	if o == nil || IsNil(o.SpecialVoucher) {
		return nil, false
	}
	return o.SpecialVoucher, true
}

// HasSpecialVoucher returns a boolean if a field has been set.
func (o *CommonVoucherUseRuleLiteInfo) HasSpecialVoucher() bool {
	if o != nil && !IsNil(o.SpecialVoucher) {
		return true
	}

	return false
}

// SetSpecialVoucher gets a reference to the given ActivitySpecialVoucher and assigns it to the SpecialVoucher field.
func (o *CommonVoucherUseRuleLiteInfo) SetSpecialVoucher(v ActivitySpecialVoucher) {
	o.SpecialVoucher = &v
}

func (o CommonVoucherUseRuleLiteInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonVoucherUseRuleLiteInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DiscountVoucher) {
		toSerialize["discount_voucher"] = o.DiscountVoucher
	}
	if !IsNil(o.ExchangeVoucher) {
		toSerialize["exchange_voucher"] = o.ExchangeVoucher
	}
	if !IsNil(o.FixVoucher) {
		toSerialize["fix_voucher"] = o.FixVoucher
	}
	if !IsNil(o.SpecialVoucher) {
		toSerialize["special_voucher"] = o.SpecialVoucher
	}
	return toSerialize, nil
}

type NullableCommonVoucherUseRuleLiteInfo struct {
	value *CommonVoucherUseRuleLiteInfo
	isSet bool
}

func (v NullableCommonVoucherUseRuleLiteInfo) Get() *CommonVoucherUseRuleLiteInfo {
	return v.value
}

func (v *NullableCommonVoucherUseRuleLiteInfo) Set(val *CommonVoucherUseRuleLiteInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonVoucherUseRuleLiteInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonVoucherUseRuleLiteInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonVoucherUseRuleLiteInfo(val *CommonVoucherUseRuleLiteInfo) *NullableCommonVoucherUseRuleLiteInfo {
	return &NullableCommonVoucherUseRuleLiteInfo{value: val, isSet: true}
}

func (v NullableCommonVoucherUseRuleLiteInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonVoucherUseRuleLiteInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

