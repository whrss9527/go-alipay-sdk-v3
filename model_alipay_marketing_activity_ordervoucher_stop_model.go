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

// checks if the AlipayMarketingActivityOrdervoucherStopModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityOrdervoucherStopModel{}

// AlipayMarketingActivityOrdervoucherStopModel struct for AlipayMarketingActivityOrdervoucherStopModel
type AlipayMarketingActivityOrdervoucherStopModel struct {
	// 商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// 外部业务单号，用作幂等控制。  幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。  外部接入方需保证业务单号唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
}

// NewAlipayMarketingActivityOrdervoucherStopModel instantiates a new AlipayMarketingActivityOrdervoucherStopModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityOrdervoucherStopModel() *AlipayMarketingActivityOrdervoucherStopModel {
	this := AlipayMarketingActivityOrdervoucherStopModel{}
	return &this
}

// NewAlipayMarketingActivityOrdervoucherStopModelWithDefaults instantiates a new AlipayMarketingActivityOrdervoucherStopModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityOrdervoucherStopModelWithDefaults() *AlipayMarketingActivityOrdervoucherStopModel {
	this := AlipayMarketingActivityOrdervoucherStopModel{}
	return &this
}

// GetMerchantAccessMode returns the MerchantAccessMode field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherStopModel) GetMerchantAccessMode() string {
	if o == nil || IsNil(o.MerchantAccessMode) {
		var ret string
		return ret
	}
	return *o.MerchantAccessMode
}

// GetMerchantAccessModeOk returns a tuple with the MerchantAccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherStopModel) GetMerchantAccessModeOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantAccessMode) {
		return nil, false
	}
	return o.MerchantAccessMode, true
}

// HasMerchantAccessMode returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherStopModel) HasMerchantAccessMode() bool {
	if o != nil && !IsNil(o.MerchantAccessMode) {
		return true
	}

	return false
}

// SetMerchantAccessMode gets a reference to the given string and assigns it to the MerchantAccessMode field.
func (o *AlipayMarketingActivityOrdervoucherStopModel) SetMerchantAccessMode(v string) {
	o.MerchantAccessMode = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherStopModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherStopModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherStopModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayMarketingActivityOrdervoucherStopModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

func (o AlipayMarketingActivityOrdervoucherStopModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityOrdervoucherStopModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantAccessMode) {
		toSerialize["merchant_access_mode"] = o.MerchantAccessMode
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	return toSerialize, nil
}

type NullableAlipayMarketingActivityOrdervoucherStopModel struct {
	value *AlipayMarketingActivityOrdervoucherStopModel
	isSet bool
}

func (v NullableAlipayMarketingActivityOrdervoucherStopModel) Get() *AlipayMarketingActivityOrdervoucherStopModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityOrdervoucherStopModel) Set(val *AlipayMarketingActivityOrdervoucherStopModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityOrdervoucherStopModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityOrdervoucherStopModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityOrdervoucherStopModel(val *AlipayMarketingActivityOrdervoucherStopModel) *NullableAlipayMarketingActivityOrdervoucherStopModel {
	return &NullableAlipayMarketingActivityOrdervoucherStopModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityOrdervoucherStopModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityOrdervoucherStopModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
