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

// checks if the AlipayEcoEduKtBillingModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEcoEduKtBillingModifyModel{}

// AlipayEcoEduKtBillingModifyModel struct for AlipayEcoEduKtBillingModifyModel
type AlipayEcoEduKtBillingModifyModel struct {
	// 成功Y，失败N
	BankSuccess *string `json:"bank_success,omitempty"`
	// 退款时，支付宝返回的用户的登录id
	BuyerLogonId *string `json:"buyer_logon_id,omitempty"`
	// 支付宝返回的买家支付宝用户id
	BuyerUserId *string `json:"buyer_user_id,omitempty"`
	// 支付宝返回的买家支付宝用户id加密后信息
	BuyerUserOpenId *string `json:"buyer_user_open_id,omitempty"`
	// 本次退款是否发生了资金变化
	FundChange *string `json:"fund_change,omitempty"`
	// 支付宝返回的退款时间，而不是商户退款申请的时间
	GmtRefund *string `json:"gmt_refund,omitempty"`
	// 标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。（若退款时填写，则同步退款状态时也必须填写）
	OutRequestNo *string `json:"out_request_no,omitempty"`
	// isv系统的订单号
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// 需要退款的金额，该金额不能大于订单金额,单位为元，支持两位小数
	RefundAmount *string `json:"refund_amount,omitempty"`
	// 支付宝返回的退款资金渠道，json格式字符串
	RefundDetailItemList *string `json:"refund_detail_item_list,omitempty"`
	// 退款原因，商家根据客户实际退款原因填写（若退款时填写，则同步退款状态时也必须填写）
	RefundReason *string `json:"refund_reason,omitempty"`
	// 状态：1:缴费成功，2:关闭账单，3、退费  如果为退款状态，需要填写fund_change,   refund_amount, refund_reason,  out_request_no, buyer_logon_id,  gmt_refund,  buyer_user_id, refund_detail_item_list;  4、同步网商返回的状态,如果是网商银行的账单，bank_success这个字段必填
	Status *string `json:"status,omitempty"`
	// 支付宝返回的交易号
	TradeNo *string `json:"trade_no,omitempty"`
}

// NewAlipayEcoEduKtBillingModifyModel instantiates a new AlipayEcoEduKtBillingModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEcoEduKtBillingModifyModel() *AlipayEcoEduKtBillingModifyModel {
	this := AlipayEcoEduKtBillingModifyModel{}
	return &this
}

// NewAlipayEcoEduKtBillingModifyModelWithDefaults instantiates a new AlipayEcoEduKtBillingModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEcoEduKtBillingModifyModelWithDefaults() *AlipayEcoEduKtBillingModifyModel {
	this := AlipayEcoEduKtBillingModifyModel{}
	return &this
}

// GetBankSuccess returns the BankSuccess field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingModifyModel) GetBankSuccess() string {
	if o == nil || IsNil(o.BankSuccess) {
		var ret string
		return ret
	}
	return *o.BankSuccess
}

// GetBankSuccessOk returns a tuple with the BankSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingModifyModel) GetBankSuccessOk() (*string, bool) {
	if o == nil || IsNil(o.BankSuccess) {
		return nil, false
	}
	return o.BankSuccess, true
}

// HasBankSuccess returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingModifyModel) HasBankSuccess() bool {
	if o != nil && !IsNil(o.BankSuccess) {
		return true
	}

	return false
}

// SetBankSuccess gets a reference to the given string and assigns it to the BankSuccess field.
func (o *AlipayEcoEduKtBillingModifyModel) SetBankSuccess(v string) {
	o.BankSuccess = &v
}

// GetBuyerLogonId returns the BuyerLogonId field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingModifyModel) GetBuyerLogonId() string {
	if o == nil || IsNil(o.BuyerLogonId) {
		var ret string
		return ret
	}
	return *o.BuyerLogonId
}

// GetBuyerLogonIdOk returns a tuple with the BuyerLogonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingModifyModel) GetBuyerLogonIdOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerLogonId) {
		return nil, false
	}
	return o.BuyerLogonId, true
}

// HasBuyerLogonId returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingModifyModel) HasBuyerLogonId() bool {
	if o != nil && !IsNil(o.BuyerLogonId) {
		return true
	}

	return false
}

// SetBuyerLogonId gets a reference to the given string and assigns it to the BuyerLogonId field.
func (o *AlipayEcoEduKtBillingModifyModel) SetBuyerLogonId(v string) {
	o.BuyerLogonId = &v
}

// GetBuyerUserId returns the BuyerUserId field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingModifyModel) GetBuyerUserId() string {
	if o == nil || IsNil(o.BuyerUserId) {
		var ret string
		return ret
	}
	return *o.BuyerUserId
}

// GetBuyerUserIdOk returns a tuple with the BuyerUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingModifyModel) GetBuyerUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerUserId) {
		return nil, false
	}
	return o.BuyerUserId, true
}

// HasBuyerUserId returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingModifyModel) HasBuyerUserId() bool {
	if o != nil && !IsNil(o.BuyerUserId) {
		return true
	}

	return false
}

// SetBuyerUserId gets a reference to the given string and assigns it to the BuyerUserId field.
func (o *AlipayEcoEduKtBillingModifyModel) SetBuyerUserId(v string) {
	o.BuyerUserId = &v
}

// GetBuyerUserOpenId returns the BuyerUserOpenId field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingModifyModel) GetBuyerUserOpenId() string {
	if o == nil || IsNil(o.BuyerUserOpenId) {
		var ret string
		return ret
	}
	return *o.BuyerUserOpenId
}

// GetBuyerUserOpenIdOk returns a tuple with the BuyerUserOpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingModifyModel) GetBuyerUserOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.BuyerUserOpenId) {
		return nil, false
	}
	return o.BuyerUserOpenId, true
}

// HasBuyerUserOpenId returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingModifyModel) HasBuyerUserOpenId() bool {
	if o != nil && !IsNil(o.BuyerUserOpenId) {
		return true
	}

	return false
}

// SetBuyerUserOpenId gets a reference to the given string and assigns it to the BuyerUserOpenId field.
func (o *AlipayEcoEduKtBillingModifyModel) SetBuyerUserOpenId(v string) {
	o.BuyerUserOpenId = &v
}

// GetFundChange returns the FundChange field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingModifyModel) GetFundChange() string {
	if o == nil || IsNil(o.FundChange) {
		var ret string
		return ret
	}
	return *o.FundChange
}

// GetFundChangeOk returns a tuple with the FundChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingModifyModel) GetFundChangeOk() (*string, bool) {
	if o == nil || IsNil(o.FundChange) {
		return nil, false
	}
	return o.FundChange, true
}

// HasFundChange returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingModifyModel) HasFundChange() bool {
	if o != nil && !IsNil(o.FundChange) {
		return true
	}

	return false
}

// SetFundChange gets a reference to the given string and assigns it to the FundChange field.
func (o *AlipayEcoEduKtBillingModifyModel) SetFundChange(v string) {
	o.FundChange = &v
}

// GetGmtRefund returns the GmtRefund field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingModifyModel) GetGmtRefund() string {
	if o == nil || IsNil(o.GmtRefund) {
		var ret string
		return ret
	}
	return *o.GmtRefund
}

// GetGmtRefundOk returns a tuple with the GmtRefund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingModifyModel) GetGmtRefundOk() (*string, bool) {
	if o == nil || IsNil(o.GmtRefund) {
		return nil, false
	}
	return o.GmtRefund, true
}

// HasGmtRefund returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingModifyModel) HasGmtRefund() bool {
	if o != nil && !IsNil(o.GmtRefund) {
		return true
	}

	return false
}

// SetGmtRefund gets a reference to the given string and assigns it to the GmtRefund field.
func (o *AlipayEcoEduKtBillingModifyModel) SetGmtRefund(v string) {
	o.GmtRefund = &v
}

// GetOutRequestNo returns the OutRequestNo field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingModifyModel) GetOutRequestNo() string {
	if o == nil || IsNil(o.OutRequestNo) {
		var ret string
		return ret
	}
	return *o.OutRequestNo
}

// GetOutRequestNoOk returns a tuple with the OutRequestNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingModifyModel) GetOutRequestNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutRequestNo) {
		return nil, false
	}
	return o.OutRequestNo, true
}

// HasOutRequestNo returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingModifyModel) HasOutRequestNo() bool {
	if o != nil && !IsNil(o.OutRequestNo) {
		return true
	}

	return false
}

// SetOutRequestNo gets a reference to the given string and assigns it to the OutRequestNo field.
func (o *AlipayEcoEduKtBillingModifyModel) SetOutRequestNo(v string) {
	o.OutRequestNo = &v
}

// GetOutTradeNo returns the OutTradeNo field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingModifyModel) GetOutTradeNo() string {
	if o == nil || IsNil(o.OutTradeNo) {
		var ret string
		return ret
	}
	return *o.OutTradeNo
}

// GetOutTradeNoOk returns a tuple with the OutTradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingModifyModel) GetOutTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutTradeNo) {
		return nil, false
	}
	return o.OutTradeNo, true
}

// HasOutTradeNo returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingModifyModel) HasOutTradeNo() bool {
	if o != nil && !IsNil(o.OutTradeNo) {
		return true
	}

	return false
}

// SetOutTradeNo gets a reference to the given string and assigns it to the OutTradeNo field.
func (o *AlipayEcoEduKtBillingModifyModel) SetOutTradeNo(v string) {
	o.OutTradeNo = &v
}

// GetRefundAmount returns the RefundAmount field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingModifyModel) GetRefundAmount() string {
	if o == nil || IsNil(o.RefundAmount) {
		var ret string
		return ret
	}
	return *o.RefundAmount
}

// GetRefundAmountOk returns a tuple with the RefundAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingModifyModel) GetRefundAmountOk() (*string, bool) {
	if o == nil || IsNil(o.RefundAmount) {
		return nil, false
	}
	return o.RefundAmount, true
}

// HasRefundAmount returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingModifyModel) HasRefundAmount() bool {
	if o != nil && !IsNil(o.RefundAmount) {
		return true
	}

	return false
}

// SetRefundAmount gets a reference to the given string and assigns it to the RefundAmount field.
func (o *AlipayEcoEduKtBillingModifyModel) SetRefundAmount(v string) {
	o.RefundAmount = &v
}

// GetRefundDetailItemList returns the RefundDetailItemList field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingModifyModel) GetRefundDetailItemList() string {
	if o == nil || IsNil(o.RefundDetailItemList) {
		var ret string
		return ret
	}
	return *o.RefundDetailItemList
}

// GetRefundDetailItemListOk returns a tuple with the RefundDetailItemList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingModifyModel) GetRefundDetailItemListOk() (*string, bool) {
	if o == nil || IsNil(o.RefundDetailItemList) {
		return nil, false
	}
	return o.RefundDetailItemList, true
}

// HasRefundDetailItemList returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingModifyModel) HasRefundDetailItemList() bool {
	if o != nil && !IsNil(o.RefundDetailItemList) {
		return true
	}

	return false
}

// SetRefundDetailItemList gets a reference to the given string and assigns it to the RefundDetailItemList field.
func (o *AlipayEcoEduKtBillingModifyModel) SetRefundDetailItemList(v string) {
	o.RefundDetailItemList = &v
}

// GetRefundReason returns the RefundReason field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingModifyModel) GetRefundReason() string {
	if o == nil || IsNil(o.RefundReason) {
		var ret string
		return ret
	}
	return *o.RefundReason
}

// GetRefundReasonOk returns a tuple with the RefundReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingModifyModel) GetRefundReasonOk() (*string, bool) {
	if o == nil || IsNil(o.RefundReason) {
		return nil, false
	}
	return o.RefundReason, true
}

// HasRefundReason returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingModifyModel) HasRefundReason() bool {
	if o != nil && !IsNil(o.RefundReason) {
		return true
	}

	return false
}

// SetRefundReason gets a reference to the given string and assigns it to the RefundReason field.
func (o *AlipayEcoEduKtBillingModifyModel) SetRefundReason(v string) {
	o.RefundReason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingModifyModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingModifyModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingModifyModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlipayEcoEduKtBillingModifyModel) SetStatus(v string) {
	o.Status = &v
}

// GetTradeNo returns the TradeNo field value if set, zero value otherwise.
func (o *AlipayEcoEduKtBillingModifyModel) GetTradeNo() string {
	if o == nil || IsNil(o.TradeNo) {
		var ret string
		return ret
	}
	return *o.TradeNo
}

// GetTradeNoOk returns a tuple with the TradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoEduKtBillingModifyModel) GetTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.TradeNo) {
		return nil, false
	}
	return o.TradeNo, true
}

// HasTradeNo returns a boolean if a field has been set.
func (o *AlipayEcoEduKtBillingModifyModel) HasTradeNo() bool {
	if o != nil && !IsNil(o.TradeNo) {
		return true
	}

	return false
}

// SetTradeNo gets a reference to the given string and assigns it to the TradeNo field.
func (o *AlipayEcoEduKtBillingModifyModel) SetTradeNo(v string) {
	o.TradeNo = &v
}

func (o AlipayEcoEduKtBillingModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEcoEduKtBillingModifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BankSuccess) {
		toSerialize["bank_success"] = o.BankSuccess
	}
	if !IsNil(o.BuyerLogonId) {
		toSerialize["buyer_logon_id"] = o.BuyerLogonId
	}
	if !IsNil(o.BuyerUserId) {
		toSerialize["buyer_user_id"] = o.BuyerUserId
	}
	if !IsNil(o.BuyerUserOpenId) {
		toSerialize["buyer_user_open_id"] = o.BuyerUserOpenId
	}
	if !IsNil(o.FundChange) {
		toSerialize["fund_change"] = o.FundChange
	}
	if !IsNil(o.GmtRefund) {
		toSerialize["gmt_refund"] = o.GmtRefund
	}
	if !IsNil(o.OutRequestNo) {
		toSerialize["out_request_no"] = o.OutRequestNo
	}
	if !IsNil(o.OutTradeNo) {
		toSerialize["out_trade_no"] = o.OutTradeNo
	}
	if !IsNil(o.RefundAmount) {
		toSerialize["refund_amount"] = o.RefundAmount
	}
	if !IsNil(o.RefundDetailItemList) {
		toSerialize["refund_detail_item_list"] = o.RefundDetailItemList
	}
	if !IsNil(o.RefundReason) {
		toSerialize["refund_reason"] = o.RefundReason
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TradeNo) {
		toSerialize["trade_no"] = o.TradeNo
	}
	return toSerialize, nil
}

type NullableAlipayEcoEduKtBillingModifyModel struct {
	value *AlipayEcoEduKtBillingModifyModel
	isSet bool
}

func (v NullableAlipayEcoEduKtBillingModifyModel) Get() *AlipayEcoEduKtBillingModifyModel {
	return v.value
}

func (v *NullableAlipayEcoEduKtBillingModifyModel) Set(val *AlipayEcoEduKtBillingModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoEduKtBillingModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoEduKtBillingModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoEduKtBillingModifyModel(val *AlipayEcoEduKtBillingModifyModel) *NullableAlipayEcoEduKtBillingModifyModel {
	return &NullableAlipayEcoEduKtBillingModifyModel{value: val, isSet: true}
}

func (v NullableAlipayEcoEduKtBillingModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoEduKtBillingModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


