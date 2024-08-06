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

// checks if the PaymentVoucherBelongMerchantInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentVoucherBelongMerchantInfo{}

// PaymentVoucherBelongMerchantInfo struct for PaymentVoucherBelongMerchantInfo
type PaymentVoucherBelongMerchantInfo struct {
	// 商户支付宝id，默认品牌名和品牌logo将从该商户信息中获取
	MerchantId *string `json:"merchant_id,omitempty"`
	// 商户支付宝id类型。
	MerchantIdType *string `json:"merchant_id_type,omitempty"`
}

// NewPaymentVoucherBelongMerchantInfo instantiates a new PaymentVoucherBelongMerchantInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentVoucherBelongMerchantInfo() *PaymentVoucherBelongMerchantInfo {
	this := PaymentVoucherBelongMerchantInfo{}
	return &this
}

// NewPaymentVoucherBelongMerchantInfoWithDefaults instantiates a new PaymentVoucherBelongMerchantInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentVoucherBelongMerchantInfoWithDefaults() *PaymentVoucherBelongMerchantInfo {
	this := PaymentVoucherBelongMerchantInfo{}
	return &this
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *PaymentVoucherBelongMerchantInfo) GetMerchantId() string {
	if o == nil || IsNil(o.MerchantId) {
		var ret string
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherBelongMerchantInfo) GetMerchantIdOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *PaymentVoucherBelongMerchantInfo) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given string and assigns it to the MerchantId field.
func (o *PaymentVoucherBelongMerchantInfo) SetMerchantId(v string) {
	o.MerchantId = &v
}

// GetMerchantIdType returns the MerchantIdType field value if set, zero value otherwise.
func (o *PaymentVoucherBelongMerchantInfo) GetMerchantIdType() string {
	if o == nil || IsNil(o.MerchantIdType) {
		var ret string
		return ret
	}
	return *o.MerchantIdType
}

// GetMerchantIdTypeOk returns a tuple with the MerchantIdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherBelongMerchantInfo) GetMerchantIdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantIdType) {
		return nil, false
	}
	return o.MerchantIdType, true
}

// HasMerchantIdType returns a boolean if a field has been set.
func (o *PaymentVoucherBelongMerchantInfo) HasMerchantIdType() bool {
	if o != nil && !IsNil(o.MerchantIdType) {
		return true
	}

	return false
}

// SetMerchantIdType gets a reference to the given string and assigns it to the MerchantIdType field.
func (o *PaymentVoucherBelongMerchantInfo) SetMerchantIdType(v string) {
	o.MerchantIdType = &v
}

func (o PaymentVoucherBelongMerchantInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentVoucherBelongMerchantInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantId) {
		toSerialize["merchant_id"] = o.MerchantId
	}
	if !IsNil(o.MerchantIdType) {
		toSerialize["merchant_id_type"] = o.MerchantIdType
	}
	return toSerialize, nil
}

type NullablePaymentVoucherBelongMerchantInfo struct {
	value *PaymentVoucherBelongMerchantInfo
	isSet bool
}

func (v NullablePaymentVoucherBelongMerchantInfo) Get() *PaymentVoucherBelongMerchantInfo {
	return v.value
}

func (v *NullablePaymentVoucherBelongMerchantInfo) Set(val *PaymentVoucherBelongMerchantInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentVoucherBelongMerchantInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentVoucherBelongMerchantInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentVoucherBelongMerchantInfo(val *PaymentVoucherBelongMerchantInfo) *NullablePaymentVoucherBelongMerchantInfo {
	return &NullablePaymentVoucherBelongMerchantInfo{value: val, isSet: true}
}

func (v NullablePaymentVoucherBelongMerchantInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentVoucherBelongMerchantInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


