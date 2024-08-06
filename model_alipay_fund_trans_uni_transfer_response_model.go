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

// checks if the AlipayFundTransUniTransferResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundTransUniTransferResponseModel{}

// AlipayFundTransUniTransferResponseModel struct for AlipayFundTransUniTransferResponseModel
type AlipayFundTransUniTransferResponseModel struct {
	// 转账金额，单位为元，默认为空，特殊场景提供。
	Amount *string `json:"amount,omitempty"`
	// 支付宝转账订单号
	OrderId *string `json:"order_id,omitempty"`
	// 商户订单号
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 支付宝支付资金流水号
	PayFundOrderId *string `json:"pay_fund_order_id,omitempty"`
	// 金融机构发起签约类、支付类、差错类业务时，应为每笔业务分配唯一的交易流水号。31位交易流水号组成规则为：“8位日期”+“16位序列号”+“1位预留位”+“6位控制位”，其中： a）“8位日期”为系统当前日期，ISODate格式：“YYYYMMDD” b）“16位序列号”由金融机构生成，金融机构应确保该值在网联当日唯一 c）“1位预留位”由平台分配 d）“6位控制位”由金融机构通过平台获取 例如：2023052993044491260542090100400
	SettleSerialNo *string `json:"settle_serial_no,omitempty"`
	// 转账单据状态。 SUCCESS（该笔转账交易成功）：成功； FAIL：失败（具体失败原因请参见error_code以及fail_reason返回值）； DEALING：处理中（转账到支付宝账户不涉及）； REFUND：退票（转账到支付宝账户不涉及）；
	Status *string `json:"status,omitempty"`
	// 订单支付时间，格式为yyyy-MM-dd HH:mm:ss
	TransDate *string `json:"trans_date,omitempty"`
}

// NewAlipayFundTransUniTransferResponseModel instantiates a new AlipayFundTransUniTransferResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundTransUniTransferResponseModel() *AlipayFundTransUniTransferResponseModel {
	this := AlipayFundTransUniTransferResponseModel{}
	return &this
}

// NewAlipayFundTransUniTransferResponseModelWithDefaults instantiates a new AlipayFundTransUniTransferResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundTransUniTransferResponseModelWithDefaults() *AlipayFundTransUniTransferResponseModel {
	this := AlipayFundTransUniTransferResponseModel{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferResponseModel) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferResponseModel) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferResponseModel) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AlipayFundTransUniTransferResponseModel) SetAmount(v string) {
	o.Amount = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferResponseModel) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferResponseModel) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferResponseModel) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *AlipayFundTransUniTransferResponseModel) SetOrderId(v string) {
	o.OrderId = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferResponseModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferResponseModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferResponseModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayFundTransUniTransferResponseModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetPayFundOrderId returns the PayFundOrderId field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferResponseModel) GetPayFundOrderId() string {
	if o == nil || IsNil(o.PayFundOrderId) {
		var ret string
		return ret
	}
	return *o.PayFundOrderId
}

// GetPayFundOrderIdOk returns a tuple with the PayFundOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferResponseModel) GetPayFundOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayFundOrderId) {
		return nil, false
	}
	return o.PayFundOrderId, true
}

// HasPayFundOrderId returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferResponseModel) HasPayFundOrderId() bool {
	if o != nil && !IsNil(o.PayFundOrderId) {
		return true
	}

	return false
}

// SetPayFundOrderId gets a reference to the given string and assigns it to the PayFundOrderId field.
func (o *AlipayFundTransUniTransferResponseModel) SetPayFundOrderId(v string) {
	o.PayFundOrderId = &v
}

// GetSettleSerialNo returns the SettleSerialNo field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferResponseModel) GetSettleSerialNo() string {
	if o == nil || IsNil(o.SettleSerialNo) {
		var ret string
		return ret
	}
	return *o.SettleSerialNo
}

// GetSettleSerialNoOk returns a tuple with the SettleSerialNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferResponseModel) GetSettleSerialNoOk() (*string, bool) {
	if o == nil || IsNil(o.SettleSerialNo) {
		return nil, false
	}
	return o.SettleSerialNo, true
}

// HasSettleSerialNo returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferResponseModel) HasSettleSerialNo() bool {
	if o != nil && !IsNil(o.SettleSerialNo) {
		return true
	}

	return false
}

// SetSettleSerialNo gets a reference to the given string and assigns it to the SettleSerialNo field.
func (o *AlipayFundTransUniTransferResponseModel) SetSettleSerialNo(v string) {
	o.SettleSerialNo = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferResponseModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferResponseModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlipayFundTransUniTransferResponseModel) SetStatus(v string) {
	o.Status = &v
}

// GetTransDate returns the TransDate field value if set, zero value otherwise.
func (o *AlipayFundTransUniTransferResponseModel) GetTransDate() string {
	if o == nil || IsNil(o.TransDate) {
		var ret string
		return ret
	}
	return *o.TransDate
}

// GetTransDateOk returns a tuple with the TransDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundTransUniTransferResponseModel) GetTransDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransDate) {
		return nil, false
	}
	return o.TransDate, true
}

// HasTransDate returns a boolean if a field has been set.
func (o *AlipayFundTransUniTransferResponseModel) HasTransDate() bool {
	if o != nil && !IsNil(o.TransDate) {
		return true
	}

	return false
}

// SetTransDate gets a reference to the given string and assigns it to the TransDate field.
func (o *AlipayFundTransUniTransferResponseModel) SetTransDate(v string) {
	o.TransDate = &v
}

func (o AlipayFundTransUniTransferResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundTransUniTransferResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.PayFundOrderId) {
		toSerialize["pay_fund_order_id"] = o.PayFundOrderId
	}
	if !IsNil(o.SettleSerialNo) {
		toSerialize["settle_serial_no"] = o.SettleSerialNo
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TransDate) {
		toSerialize["trans_date"] = o.TransDate
	}
	return toSerialize, nil
}

type NullableAlipayFundTransUniTransferResponseModel struct {
	value *AlipayFundTransUniTransferResponseModel
	isSet bool
}

func (v NullableAlipayFundTransUniTransferResponseModel) Get() *AlipayFundTransUniTransferResponseModel {
	return v.value
}

func (v *NullableAlipayFundTransUniTransferResponseModel) Set(val *AlipayFundTransUniTransferResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundTransUniTransferResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundTransUniTransferResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundTransUniTransferResponseModel(val *AlipayFundTransUniTransferResponseModel) *NullableAlipayFundTransUniTransferResponseModel {
	return &NullableAlipayFundTransUniTransferResponseModel{value: val, isSet: true}
}

func (v NullableAlipayFundTransUniTransferResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundTransUniTransferResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
