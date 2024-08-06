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

// checks if the WaitRepaymentOrderInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WaitRepaymentOrderInfo{}

// WaitRepaymentOrderInfo struct for WaitRepaymentOrderInfo
type WaitRepaymentOrderInfo struct {
	// 垫资单id
	AdvanceOrderId *string `json:"advance_order_id,omitempty"`
	// 买家的支付宝用户id
	AlipayUserId *string `json:"alipay_user_id,omitempty"`
	// 通常为商户签约的收单产品码
	BizProduct *string `json:"biz_product,omitempty"`
	// 买家的支付宝用户id
	OpenId *string `json:"open_id,omitempty"`
	// 原始的业务单号，通常为支付宝交易号
	OrigBizOrderId *string `json:"orig_biz_order_id,omitempty"`
	// 垫资金额
	WaitRepaymentAmount *int32 `json:"wait_repayment_amount,omitempty"`
}

// NewWaitRepaymentOrderInfo instantiates a new WaitRepaymentOrderInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWaitRepaymentOrderInfo() *WaitRepaymentOrderInfo {
	this := WaitRepaymentOrderInfo{}
	return &this
}

// NewWaitRepaymentOrderInfoWithDefaults instantiates a new WaitRepaymentOrderInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWaitRepaymentOrderInfoWithDefaults() *WaitRepaymentOrderInfo {
	this := WaitRepaymentOrderInfo{}
	return &this
}

// GetAdvanceOrderId returns the AdvanceOrderId field value if set, zero value otherwise.
func (o *WaitRepaymentOrderInfo) GetAdvanceOrderId() string {
	if o == nil || IsNil(o.AdvanceOrderId) {
		var ret string
		return ret
	}
	return *o.AdvanceOrderId
}

// GetAdvanceOrderIdOk returns a tuple with the AdvanceOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaitRepaymentOrderInfo) GetAdvanceOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdvanceOrderId) {
		return nil, false
	}
	return o.AdvanceOrderId, true
}

// HasAdvanceOrderId returns a boolean if a field has been set.
func (o *WaitRepaymentOrderInfo) HasAdvanceOrderId() bool {
	if o != nil && !IsNil(o.AdvanceOrderId) {
		return true
	}

	return false
}

// SetAdvanceOrderId gets a reference to the given string and assigns it to the AdvanceOrderId field.
func (o *WaitRepaymentOrderInfo) SetAdvanceOrderId(v string) {
	o.AdvanceOrderId = &v
}

// GetAlipayUserId returns the AlipayUserId field value if set, zero value otherwise.
func (o *WaitRepaymentOrderInfo) GetAlipayUserId() string {
	if o == nil || IsNil(o.AlipayUserId) {
		var ret string
		return ret
	}
	return *o.AlipayUserId
}

// GetAlipayUserIdOk returns a tuple with the AlipayUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaitRepaymentOrderInfo) GetAlipayUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlipayUserId) {
		return nil, false
	}
	return o.AlipayUserId, true
}

// HasAlipayUserId returns a boolean if a field has been set.
func (o *WaitRepaymentOrderInfo) HasAlipayUserId() bool {
	if o != nil && !IsNil(o.AlipayUserId) {
		return true
	}

	return false
}

// SetAlipayUserId gets a reference to the given string and assigns it to the AlipayUserId field.
func (o *WaitRepaymentOrderInfo) SetAlipayUserId(v string) {
	o.AlipayUserId = &v
}

// GetBizProduct returns the BizProduct field value if set, zero value otherwise.
func (o *WaitRepaymentOrderInfo) GetBizProduct() string {
	if o == nil || IsNil(o.BizProduct) {
		var ret string
		return ret
	}
	return *o.BizProduct
}

// GetBizProductOk returns a tuple with the BizProduct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaitRepaymentOrderInfo) GetBizProductOk() (*string, bool) {
	if o == nil || IsNil(o.BizProduct) {
		return nil, false
	}
	return o.BizProduct, true
}

// HasBizProduct returns a boolean if a field has been set.
func (o *WaitRepaymentOrderInfo) HasBizProduct() bool {
	if o != nil && !IsNil(o.BizProduct) {
		return true
	}

	return false
}

// SetBizProduct gets a reference to the given string and assigns it to the BizProduct field.
func (o *WaitRepaymentOrderInfo) SetBizProduct(v string) {
	o.BizProduct = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *WaitRepaymentOrderInfo) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaitRepaymentOrderInfo) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *WaitRepaymentOrderInfo) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *WaitRepaymentOrderInfo) SetOpenId(v string) {
	o.OpenId = &v
}

// GetOrigBizOrderId returns the OrigBizOrderId field value if set, zero value otherwise.
func (o *WaitRepaymentOrderInfo) GetOrigBizOrderId() string {
	if o == nil || IsNil(o.OrigBizOrderId) {
		var ret string
		return ret
	}
	return *o.OrigBizOrderId
}

// GetOrigBizOrderIdOk returns a tuple with the OrigBizOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaitRepaymentOrderInfo) GetOrigBizOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrigBizOrderId) {
		return nil, false
	}
	return o.OrigBizOrderId, true
}

// HasOrigBizOrderId returns a boolean if a field has been set.
func (o *WaitRepaymentOrderInfo) HasOrigBizOrderId() bool {
	if o != nil && !IsNil(o.OrigBizOrderId) {
		return true
	}

	return false
}

// SetOrigBizOrderId gets a reference to the given string and assigns it to the OrigBizOrderId field.
func (o *WaitRepaymentOrderInfo) SetOrigBizOrderId(v string) {
	o.OrigBizOrderId = &v
}

// GetWaitRepaymentAmount returns the WaitRepaymentAmount field value if set, zero value otherwise.
func (o *WaitRepaymentOrderInfo) GetWaitRepaymentAmount() int32 {
	if o == nil || IsNil(o.WaitRepaymentAmount) {
		var ret int32
		return ret
	}
	return *o.WaitRepaymentAmount
}

// GetWaitRepaymentAmountOk returns a tuple with the WaitRepaymentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WaitRepaymentOrderInfo) GetWaitRepaymentAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.WaitRepaymentAmount) {
		return nil, false
	}
	return o.WaitRepaymentAmount, true
}

// HasWaitRepaymentAmount returns a boolean if a field has been set.
func (o *WaitRepaymentOrderInfo) HasWaitRepaymentAmount() bool {
	if o != nil && !IsNil(o.WaitRepaymentAmount) {
		return true
	}

	return false
}

// SetWaitRepaymentAmount gets a reference to the given int32 and assigns it to the WaitRepaymentAmount field.
func (o *WaitRepaymentOrderInfo) SetWaitRepaymentAmount(v int32) {
	o.WaitRepaymentAmount = &v
}

func (o WaitRepaymentOrderInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WaitRepaymentOrderInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdvanceOrderId) {
		toSerialize["advance_order_id"] = o.AdvanceOrderId
	}
	if !IsNil(o.AlipayUserId) {
		toSerialize["alipay_user_id"] = o.AlipayUserId
	}
	if !IsNil(o.BizProduct) {
		toSerialize["biz_product"] = o.BizProduct
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.OrigBizOrderId) {
		toSerialize["orig_biz_order_id"] = o.OrigBizOrderId
	}
	if !IsNil(o.WaitRepaymentAmount) {
		toSerialize["wait_repayment_amount"] = o.WaitRepaymentAmount
	}
	return toSerialize, nil
}

type NullableWaitRepaymentOrderInfo struct {
	value *WaitRepaymentOrderInfo
	isSet bool
}

func (v NullableWaitRepaymentOrderInfo) Get() *WaitRepaymentOrderInfo {
	return v.value
}

func (v *NullableWaitRepaymentOrderInfo) Set(val *WaitRepaymentOrderInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWaitRepaymentOrderInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWaitRepaymentOrderInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWaitRepaymentOrderInfo(val *WaitRepaymentOrderInfo) *NullableWaitRepaymentOrderInfo {
	return &NullableWaitRepaymentOrderInfo{value: val, isSet: true}
}

func (v NullableWaitRepaymentOrderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWaitRepaymentOrderInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
