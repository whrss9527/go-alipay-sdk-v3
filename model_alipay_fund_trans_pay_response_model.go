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

// checks if the AlipayFundTransPayResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundTransPayResponseModel{}

// AlipayFundTransPayResponseModel struct for AlipayFundTransPayResponseModel
type AlipayFundTransPayResponseModel struct {
	// 该笔转账在支付宝系统内部的单据ID
	OrderId *string `json:"order_id,omitempty"`
	// 商户端的唯一订单号
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 支付宝支付资金流水号
	PayFundOrderId *string `json:"pay_fund_order_id,omitempty"`
	// SUCCESS：支付成功；  FAIL：支付失败；  UNKNOWN：未知；建议通过查询确认最新状态
	Status *string `json:"status,omitempty"`
	// 订单支付时间，格式为yyyy-MM-dd HH:mm:ss
	TransPayTime *string `json:"trans_pay_time,omitempty"`
}

// NewAlipayFundTransPayResponseModel instantiates a new AlipayFundTransPayResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundTransPayResponseModel() *AlipayFundTransPayResponseModel {
	this := AlipayFundTransPayResponseModel{}
	return &this
}

// NewAlipayFundTransPayResponseModelWithDefaults instantiates a new AlipayFundTransPayResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundTransPayResponseModelWithDefaults() *AlipayFundTransPayResponseModel {
	this := AlipayFundTransPayResponseModel{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *AlipayFundTransPayResponseModel) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransPayResponseModel) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *AlipayFundTransPayResponseModel) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *AlipayFundTransPayResponseModel) SetOrderId(v string) {
	o.OrderId = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayFundTransPayResponseModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransPayResponseModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayFundTransPayResponseModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayFundTransPayResponseModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetPayFundOrderId returns the PayFundOrderId field value if set, zero value otherwise.
func (o *AlipayFundTransPayResponseModel) GetPayFundOrderId() string {
	if o == nil || IsNil(o.PayFundOrderId) {
		var ret string
		return ret
	}
	return *o.PayFundOrderId
}

// GetPayFundOrderIdOk returns a tuple with the PayFundOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransPayResponseModel) GetPayFundOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayFundOrderId) {
		return nil, false
	}
	return o.PayFundOrderId, true
}

// HasPayFundOrderId returns a boolean if a field has been set.
func (o *AlipayFundTransPayResponseModel) HasPayFundOrderId() bool {
	if o != nil && !IsNil(o.PayFundOrderId) {
		return true
	}

	return false
}

// SetPayFundOrderId gets a reference to the given string and assigns it to the PayFundOrderId field.
func (o *AlipayFundTransPayResponseModel) SetPayFundOrderId(v string) {
	o.PayFundOrderId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayFundTransPayResponseModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransPayResponseModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayFundTransPayResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlipayFundTransPayResponseModel) SetStatus(v string) {
	o.Status = &v
}

// GetTransPayTime returns the TransPayTime field value if set, zero value otherwise.
func (o *AlipayFundTransPayResponseModel) GetTransPayTime() string {
	if o == nil || IsNil(o.TransPayTime) {
		var ret string
		return ret
	}
	return *o.TransPayTime
}

// GetTransPayTimeOk returns a tuple with the TransPayTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransPayResponseModel) GetTransPayTimeOk() (*string, bool) {
	if o == nil || IsNil(o.TransPayTime) {
		return nil, false
	}
	return o.TransPayTime, true
}

// HasTransPayTime returns a boolean if a field has been set.
func (o *AlipayFundTransPayResponseModel) HasTransPayTime() bool {
	if o != nil && !IsNil(o.TransPayTime) {
		return true
	}

	return false
}

// SetTransPayTime gets a reference to the given string and assigns it to the TransPayTime field.
func (o *AlipayFundTransPayResponseModel) SetTransPayTime(v string) {
	o.TransPayTime = &v
}

func (o AlipayFundTransPayResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundTransPayResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.PayFundOrderId) {
		toSerialize["pay_fund_order_id"] = o.PayFundOrderId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TransPayTime) {
		toSerialize["trans_pay_time"] = o.TransPayTime
	}
	return toSerialize, nil
}

type NullableAlipayFundTransPayResponseModel struct {
	value *AlipayFundTransPayResponseModel
	isSet bool
}

func (v NullableAlipayFundTransPayResponseModel) Get() *AlipayFundTransPayResponseModel {
	return v.value
}

func (v *NullableAlipayFundTransPayResponseModel) Set(val *AlipayFundTransPayResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundTransPayResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundTransPayResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundTransPayResponseModel(val *AlipayFundTransPayResponseModel) *NullableAlipayFundTransPayResponseModel {
	return &NullableAlipayFundTransPayResponseModel{value: val, isSet: true}
}

func (v NullableAlipayFundTransPayResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundTransPayResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


