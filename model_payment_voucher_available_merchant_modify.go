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

// checks if the PaymentVoucherAvailableMerchantModify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentVoucherAvailableMerchantModify{}

// PaymentVoucherAvailableMerchantModify struct for PaymentVoucherAvailableMerchantModify
type PaymentVoucherAvailableMerchantModify struct {
	// 优惠券可以核销的直连商户PID。默认不修改。  限制：  1、核销商户范围只能增大不能减小。
	AvailablePids []string `json:"available_pids,omitempty"`
	// 优惠券可以核销的间连商户SMID。默认不修改。  限制：   1、核销商户范围只能增大不能减小。
	AvailableSmids []string `json:"available_smids,omitempty"`
}

// NewPaymentVoucherAvailableMerchantModify instantiates a new PaymentVoucherAvailableMerchantModify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentVoucherAvailableMerchantModify() *PaymentVoucherAvailableMerchantModify {
	this := PaymentVoucherAvailableMerchantModify{}
	return &this
}

// NewPaymentVoucherAvailableMerchantModifyWithDefaults instantiates a new PaymentVoucherAvailableMerchantModify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentVoucherAvailableMerchantModifyWithDefaults() *PaymentVoucherAvailableMerchantModify {
	this := PaymentVoucherAvailableMerchantModify{}
	return &this
}

// GetAvailablePids returns the AvailablePids field value if set, zero value otherwise.
func (o *PaymentVoucherAvailableMerchantModify) GetAvailablePids() []string {
	if o == nil || IsNil(o.AvailablePids) {
		var ret []string
		return ret
	}
	return o.AvailablePids
}

// GetAvailablePidsOk returns a tuple with the AvailablePids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherAvailableMerchantModify) GetAvailablePidsOk() ([]string, bool) {
	if o == nil || IsNil(o.AvailablePids) {
		return nil, false
	}
	return o.AvailablePids, true
}

// HasAvailablePids returns a boolean if a field has been set.
func (o *PaymentVoucherAvailableMerchantModify) HasAvailablePids() bool {
	if o != nil && !IsNil(o.AvailablePids) {
		return true
	}

	return false
}

// SetAvailablePids gets a reference to the given []string and assigns it to the AvailablePids field.
func (o *PaymentVoucherAvailableMerchantModify) SetAvailablePids(v []string) {
	o.AvailablePids = v
}

// GetAvailableSmids returns the AvailableSmids field value if set, zero value otherwise.
func (o *PaymentVoucherAvailableMerchantModify) GetAvailableSmids() []string {
	if o == nil || IsNil(o.AvailableSmids) {
		var ret []string
		return ret
	}
	return o.AvailableSmids
}

// GetAvailableSmidsOk returns a tuple with the AvailableSmids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentVoucherAvailableMerchantModify) GetAvailableSmidsOk() ([]string, bool) {
	if o == nil || IsNil(o.AvailableSmids) {
		return nil, false
	}
	return o.AvailableSmids, true
}

// HasAvailableSmids returns a boolean if a field has been set.
func (o *PaymentVoucherAvailableMerchantModify) HasAvailableSmids() bool {
	if o != nil && !IsNil(o.AvailableSmids) {
		return true
	}

	return false
}

// SetAvailableSmids gets a reference to the given []string and assigns it to the AvailableSmids field.
func (o *PaymentVoucherAvailableMerchantModify) SetAvailableSmids(v []string) {
	o.AvailableSmids = v
}

func (o PaymentVoucherAvailableMerchantModify) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentVoucherAvailableMerchantModify) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvailablePids) {
		toSerialize["available_pids"] = o.AvailablePids
	}
	if !IsNil(o.AvailableSmids) {
		toSerialize["available_smids"] = o.AvailableSmids
	}
	return toSerialize, nil
}

type NullablePaymentVoucherAvailableMerchantModify struct {
	value *PaymentVoucherAvailableMerchantModify
	isSet bool
}

func (v NullablePaymentVoucherAvailableMerchantModify) Get() *PaymentVoucherAvailableMerchantModify {
	return v.value
}

func (v *NullablePaymentVoucherAvailableMerchantModify) Set(val *PaymentVoucherAvailableMerchantModify) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentVoucherAvailableMerchantModify) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentVoucherAvailableMerchantModify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentVoucherAvailableMerchantModify(val *PaymentVoucherAvailableMerchantModify) *NullablePaymentVoucherAvailableMerchantModify {
	return &NullablePaymentVoucherAvailableMerchantModify{value: val, isSet: true}
}

func (v NullablePaymentVoucherAvailableMerchantModify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentVoucherAvailableMerchantModify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


