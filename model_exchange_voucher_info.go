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

// checks if the ExchangeVoucherInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeVoucherInfo{}

// ExchangeVoucherInfo struct for ExchangeVoucherInfo
type ExchangeVoucherInfo struct {
	// 券的价值
	Amount *string `json:"amount,omitempty"`
	// 兑换券业务类型。 注意：兑换券通过大促活动权益报名能力推广至支付宝会场时，本参数必填。
	BizType *string `json:"biz_type,omitempty"`
	ExchangeGoodsInfo *VoucherExchangeGoodsInfo `json:"exchange_goods_info,omitempty"`
	// 优惠门槛金额，表示只有当订单金额大于等于门槛金额时券才能使用。该字段为空时表示无门槛。  门槛金额的校验由服务商(商户)核销时自行校验，支付宝侧只做展示使用。
	FloorAmount *string `json:"floor_amount,omitempty"`
	VoucherDeductThresholdInfo *VoucherDeductThresholdInfo `json:"voucher_deduct_threshold_info,omitempty"`
}

// NewExchangeVoucherInfo instantiates a new ExchangeVoucherInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeVoucherInfo() *ExchangeVoucherInfo {
	this := ExchangeVoucherInfo{}
	return &this
}

// NewExchangeVoucherInfoWithDefaults instantiates a new ExchangeVoucherInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeVoucherInfoWithDefaults() *ExchangeVoucherInfo {
	this := ExchangeVoucherInfo{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ExchangeVoucherInfo) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeVoucherInfo) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ExchangeVoucherInfo) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *ExchangeVoucherInfo) SetAmount(v string) {
	o.Amount = &v
}

// GetBizType returns the BizType field value if set, zero value otherwise.
func (o *ExchangeVoucherInfo) GetBizType() string {
	if o == nil || IsNil(o.BizType) {
		var ret string
		return ret
	}
	return *o.BizType
}

// GetBizTypeOk returns a tuple with the BizType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeVoucherInfo) GetBizTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BizType) {
		return nil, false
	}
	return o.BizType, true
}

// HasBizType returns a boolean if a field has been set.
func (o *ExchangeVoucherInfo) HasBizType() bool {
	if o != nil && !IsNil(o.BizType) {
		return true
	}

	return false
}

// SetBizType gets a reference to the given string and assigns it to the BizType field.
func (o *ExchangeVoucherInfo) SetBizType(v string) {
	o.BizType = &v
}

// GetExchangeGoodsInfo returns the ExchangeGoodsInfo field value if set, zero value otherwise.
func (o *ExchangeVoucherInfo) GetExchangeGoodsInfo() VoucherExchangeGoodsInfo {
	if o == nil || IsNil(o.ExchangeGoodsInfo) {
		var ret VoucherExchangeGoodsInfo
		return ret
	}
	return *o.ExchangeGoodsInfo
}

// GetExchangeGoodsInfoOk returns a tuple with the ExchangeGoodsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeVoucherInfo) GetExchangeGoodsInfoOk() (*VoucherExchangeGoodsInfo, bool) {
	if o == nil || IsNil(o.ExchangeGoodsInfo) {
		return nil, false
	}
	return o.ExchangeGoodsInfo, true
}

// HasExchangeGoodsInfo returns a boolean if a field has been set.
func (o *ExchangeVoucherInfo) HasExchangeGoodsInfo() bool {
	if o != nil && !IsNil(o.ExchangeGoodsInfo) {
		return true
	}

	return false
}

// SetExchangeGoodsInfo gets a reference to the given VoucherExchangeGoodsInfo and assigns it to the ExchangeGoodsInfo field.
func (o *ExchangeVoucherInfo) SetExchangeGoodsInfo(v VoucherExchangeGoodsInfo) {
	o.ExchangeGoodsInfo = &v
}

// GetFloorAmount returns the FloorAmount field value if set, zero value otherwise.
func (o *ExchangeVoucherInfo) GetFloorAmount() string {
	if o == nil || IsNil(o.FloorAmount) {
		var ret string
		return ret
	}
	return *o.FloorAmount
}

// GetFloorAmountOk returns a tuple with the FloorAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeVoucherInfo) GetFloorAmountOk() (*string, bool) {
	if o == nil || IsNil(o.FloorAmount) {
		return nil, false
	}
	return o.FloorAmount, true
}

// HasFloorAmount returns a boolean if a field has been set.
func (o *ExchangeVoucherInfo) HasFloorAmount() bool {
	if o != nil && !IsNil(o.FloorAmount) {
		return true
	}

	return false
}

// SetFloorAmount gets a reference to the given string and assigns it to the FloorAmount field.
func (o *ExchangeVoucherInfo) SetFloorAmount(v string) {
	o.FloorAmount = &v
}

// GetVoucherDeductThresholdInfo returns the VoucherDeductThresholdInfo field value if set, zero value otherwise.
func (o *ExchangeVoucherInfo) GetVoucherDeductThresholdInfo() VoucherDeductThresholdInfo {
	if o == nil || IsNil(o.VoucherDeductThresholdInfo) {
		var ret VoucherDeductThresholdInfo
		return ret
	}
	return *o.VoucherDeductThresholdInfo
}

// GetVoucherDeductThresholdInfoOk returns a tuple with the VoucherDeductThresholdInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeVoucherInfo) GetVoucherDeductThresholdInfoOk() (*VoucherDeductThresholdInfo, bool) {
	if o == nil || IsNil(o.VoucherDeductThresholdInfo) {
		return nil, false
	}
	return o.VoucherDeductThresholdInfo, true
}

// HasVoucherDeductThresholdInfo returns a boolean if a field has been set.
func (o *ExchangeVoucherInfo) HasVoucherDeductThresholdInfo() bool {
	if o != nil && !IsNil(o.VoucherDeductThresholdInfo) {
		return true
	}

	return false
}

// SetVoucherDeductThresholdInfo gets a reference to the given VoucherDeductThresholdInfo and assigns it to the VoucherDeductThresholdInfo field.
func (o *ExchangeVoucherInfo) SetVoucherDeductThresholdInfo(v VoucherDeductThresholdInfo) {
	o.VoucherDeductThresholdInfo = &v
}

func (o ExchangeVoucherInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeVoucherInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.BizType) {
		toSerialize["biz_type"] = o.BizType
	}
	if !IsNil(o.ExchangeGoodsInfo) {
		toSerialize["exchange_goods_info"] = o.ExchangeGoodsInfo
	}
	if !IsNil(o.FloorAmount) {
		toSerialize["floor_amount"] = o.FloorAmount
	}
	if !IsNil(o.VoucherDeductThresholdInfo) {
		toSerialize["voucher_deduct_threshold_info"] = o.VoucherDeductThresholdInfo
	}
	return toSerialize, nil
}

type NullableExchangeVoucherInfo struct {
	value *ExchangeVoucherInfo
	isSet bool
}

func (v NullableExchangeVoucherInfo) Get() *ExchangeVoucherInfo {
	return v.value
}

func (v *NullableExchangeVoucherInfo) Set(val *ExchangeVoucherInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeVoucherInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeVoucherInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeVoucherInfo(val *ExchangeVoucherInfo) *NullableExchangeVoucherInfo {
	return &NullableExchangeVoucherInfo{value: val, isSet: true}
}

func (v NullableExchangeVoucherInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeVoucherInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


