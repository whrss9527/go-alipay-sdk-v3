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

// checks if the AlipayTradeCustomsDeclareResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayTradeCustomsDeclareResponseModel{}

// AlipayTradeCustomsDeclareResponseModel struct for AlipayTradeCustomsDeclareResponseModel
type AlipayTradeCustomsDeclareResponseModel struct {
	// 支付宝报关流水号。
	AlipayDeclareNo *string `json:"alipay_declare_no,omitempty"`
	// 币种
	Currency *string `json:"currency,omitempty"`
	// 订购人身份信息和支付人的身份信息验证结果。T表示商户传入的订购人姓名和身份证号和支付人的姓名和身份证号一致。F代表商户传入的订购人姓名和身份证号和支付人的姓名和身份证号不一致。对于同一笔报关单支付宝只会校验一次，如果多次重推不会返回此参数。
	IdentityCheck *string `json:"identity_check,omitempty"`
	// 国际站外部订单号
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// 支付机构注册号
	PayCode *string `json:"pay_code,omitempty"`
	// 清算流水号
	PayTransactionId *string `json:"pay_transaction_id,omitempty"`
	// 交易总金额(单位:分)
	TotalAmount *string `json:"total_amount,omitempty"`
	// 支付宝推送到海关的支付单据号。
	TradeNo *string `json:"trade_no,omitempty"`
	// 清算机构标志（1-cup,2-null,3-other）
	VerDept *string `json:"ver_dept,omitempty"`
}

// NewAlipayTradeCustomsDeclareResponseModel instantiates a new AlipayTradeCustomsDeclareResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayTradeCustomsDeclareResponseModel() *AlipayTradeCustomsDeclareResponseModel {
	this := AlipayTradeCustomsDeclareResponseModel{}
	return &this
}

// NewAlipayTradeCustomsDeclareResponseModelWithDefaults instantiates a new AlipayTradeCustomsDeclareResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayTradeCustomsDeclareResponseModelWithDefaults() *AlipayTradeCustomsDeclareResponseModel {
	this := AlipayTradeCustomsDeclareResponseModel{}
	return &this
}

// GetAlipayDeclareNo returns the AlipayDeclareNo field value if set, zero value otherwise.
func (o *AlipayTradeCustomsDeclareResponseModel) GetAlipayDeclareNo() string {
	if o == nil || IsNil(o.AlipayDeclareNo) {
		var ret string
		return ret
	}
	return *o.AlipayDeclareNo
}

// GetAlipayDeclareNoOk returns a tuple with the AlipayDeclareNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeCustomsDeclareResponseModel) GetAlipayDeclareNoOk() (*string, bool) {
	if o == nil || IsNil(o.AlipayDeclareNo) {
		return nil, false
	}
	return o.AlipayDeclareNo, true
}

// HasAlipayDeclareNo returns a boolean if a field has been set.
func (o *AlipayTradeCustomsDeclareResponseModel) HasAlipayDeclareNo() bool {
	if o != nil && !IsNil(o.AlipayDeclareNo) {
		return true
	}

	return false
}

// SetAlipayDeclareNo gets a reference to the given string and assigns it to the AlipayDeclareNo field.
func (o *AlipayTradeCustomsDeclareResponseModel) SetAlipayDeclareNo(v string) {
	o.AlipayDeclareNo = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AlipayTradeCustomsDeclareResponseModel) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeCustomsDeclareResponseModel) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AlipayTradeCustomsDeclareResponseModel) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AlipayTradeCustomsDeclareResponseModel) SetCurrency(v string) {
	o.Currency = &v
}

// GetIdentityCheck returns the IdentityCheck field value if set, zero value otherwise.
func (o *AlipayTradeCustomsDeclareResponseModel) GetIdentityCheck() string {
	if o == nil || IsNil(o.IdentityCheck) {
		var ret string
		return ret
	}
	return *o.IdentityCheck
}

// GetIdentityCheckOk returns a tuple with the IdentityCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeCustomsDeclareResponseModel) GetIdentityCheckOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityCheck) {
		return nil, false
	}
	return o.IdentityCheck, true
}

// HasIdentityCheck returns a boolean if a field has been set.
func (o *AlipayTradeCustomsDeclareResponseModel) HasIdentityCheck() bool {
	if o != nil && !IsNil(o.IdentityCheck) {
		return true
	}

	return false
}

// SetIdentityCheck gets a reference to the given string and assigns it to the IdentityCheck field.
func (o *AlipayTradeCustomsDeclareResponseModel) SetIdentityCheck(v string) {
	o.IdentityCheck = &v
}

// GetOutTradeNo returns the OutTradeNo field value if set, zero value otherwise.
func (o *AlipayTradeCustomsDeclareResponseModel) GetOutTradeNo() string {
	if o == nil || IsNil(o.OutTradeNo) {
		var ret string
		return ret
	}
	return *o.OutTradeNo
}

// GetOutTradeNoOk returns a tuple with the OutTradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeCustomsDeclareResponseModel) GetOutTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutTradeNo) {
		return nil, false
	}
	return o.OutTradeNo, true
}

// HasOutTradeNo returns a boolean if a field has been set.
func (o *AlipayTradeCustomsDeclareResponseModel) HasOutTradeNo() bool {
	if o != nil && !IsNil(o.OutTradeNo) {
		return true
	}

	return false
}

// SetOutTradeNo gets a reference to the given string and assigns it to the OutTradeNo field.
func (o *AlipayTradeCustomsDeclareResponseModel) SetOutTradeNo(v string) {
	o.OutTradeNo = &v
}

// GetPayCode returns the PayCode field value if set, zero value otherwise.
func (o *AlipayTradeCustomsDeclareResponseModel) GetPayCode() string {
	if o == nil || IsNil(o.PayCode) {
		var ret string
		return ret
	}
	return *o.PayCode
}

// GetPayCodeOk returns a tuple with the PayCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeCustomsDeclareResponseModel) GetPayCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PayCode) {
		return nil, false
	}
	return o.PayCode, true
}

// HasPayCode returns a boolean if a field has been set.
func (o *AlipayTradeCustomsDeclareResponseModel) HasPayCode() bool {
	if o != nil && !IsNil(o.PayCode) {
		return true
	}

	return false
}

// SetPayCode gets a reference to the given string and assigns it to the PayCode field.
func (o *AlipayTradeCustomsDeclareResponseModel) SetPayCode(v string) {
	o.PayCode = &v
}

// GetPayTransactionId returns the PayTransactionId field value if set, zero value otherwise.
func (o *AlipayTradeCustomsDeclareResponseModel) GetPayTransactionId() string {
	if o == nil || IsNil(o.PayTransactionId) {
		var ret string
		return ret
	}
	return *o.PayTransactionId
}

// GetPayTransactionIdOk returns a tuple with the PayTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeCustomsDeclareResponseModel) GetPayTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayTransactionId) {
		return nil, false
	}
	return o.PayTransactionId, true
}

// HasPayTransactionId returns a boolean if a field has been set.
func (o *AlipayTradeCustomsDeclareResponseModel) HasPayTransactionId() bool {
	if o != nil && !IsNil(o.PayTransactionId) {
		return true
	}

	return false
}

// SetPayTransactionId gets a reference to the given string and assigns it to the PayTransactionId field.
func (o *AlipayTradeCustomsDeclareResponseModel) SetPayTransactionId(v string) {
	o.PayTransactionId = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *AlipayTradeCustomsDeclareResponseModel) GetTotalAmount() string {
	if o == nil || IsNil(o.TotalAmount) {
		var ret string
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeCustomsDeclareResponseModel) GetTotalAmountOk() (*string, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *AlipayTradeCustomsDeclareResponseModel) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given string and assigns it to the TotalAmount field.
func (o *AlipayTradeCustomsDeclareResponseModel) SetTotalAmount(v string) {
	o.TotalAmount = &v
}

// GetTradeNo returns the TradeNo field value if set, zero value otherwise.
func (o *AlipayTradeCustomsDeclareResponseModel) GetTradeNo() string {
	if o == nil || IsNil(o.TradeNo) {
		var ret string
		return ret
	}
	return *o.TradeNo
}

// GetTradeNoOk returns a tuple with the TradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeCustomsDeclareResponseModel) GetTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.TradeNo) {
		return nil, false
	}
	return o.TradeNo, true
}

// HasTradeNo returns a boolean if a field has been set.
func (o *AlipayTradeCustomsDeclareResponseModel) HasTradeNo() bool {
	if o != nil && !IsNil(o.TradeNo) {
		return true
	}

	return false
}

// SetTradeNo gets a reference to the given string and assigns it to the TradeNo field.
func (o *AlipayTradeCustomsDeclareResponseModel) SetTradeNo(v string) {
	o.TradeNo = &v
}

// GetVerDept returns the VerDept field value if set, zero value otherwise.
func (o *AlipayTradeCustomsDeclareResponseModel) GetVerDept() string {
	if o == nil || IsNil(o.VerDept) {
		var ret string
		return ret
	}
	return *o.VerDept
}

// GetVerDeptOk returns a tuple with the VerDept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayTradeCustomsDeclareResponseModel) GetVerDeptOk() (*string, bool) {
	if o == nil || IsNil(o.VerDept) {
		return nil, false
	}
	return o.VerDept, true
}

// HasVerDept returns a boolean if a field has been set.
func (o *AlipayTradeCustomsDeclareResponseModel) HasVerDept() bool {
	if o != nil && !IsNil(o.VerDept) {
		return true
	}

	return false
}

// SetVerDept gets a reference to the given string and assigns it to the VerDept field.
func (o *AlipayTradeCustomsDeclareResponseModel) SetVerDept(v string) {
	o.VerDept = &v
}

func (o AlipayTradeCustomsDeclareResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayTradeCustomsDeclareResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlipayDeclareNo) {
		toSerialize["alipay_declare_no"] = o.AlipayDeclareNo
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.IdentityCheck) {
		toSerialize["identity_check"] = o.IdentityCheck
	}
	if !IsNil(o.OutTradeNo) {
		toSerialize["out_trade_no"] = o.OutTradeNo
	}
	if !IsNil(o.PayCode) {
		toSerialize["pay_code"] = o.PayCode
	}
	if !IsNil(o.PayTransactionId) {
		toSerialize["pay_transaction_id"] = o.PayTransactionId
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["total_amount"] = o.TotalAmount
	}
	if !IsNil(o.TradeNo) {
		toSerialize["trade_no"] = o.TradeNo
	}
	if !IsNil(o.VerDept) {
		toSerialize["ver_dept"] = o.VerDept
	}
	return toSerialize, nil
}

type NullableAlipayTradeCustomsDeclareResponseModel struct {
	value *AlipayTradeCustomsDeclareResponseModel
	isSet bool
}

func (v NullableAlipayTradeCustomsDeclareResponseModel) Get() *AlipayTradeCustomsDeclareResponseModel {
	return v.value
}

func (v *NullableAlipayTradeCustomsDeclareResponseModel) Set(val *AlipayTradeCustomsDeclareResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayTradeCustomsDeclareResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayTradeCustomsDeclareResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayTradeCustomsDeclareResponseModel(val *AlipayTradeCustomsDeclareResponseModel) *NullableAlipayTradeCustomsDeclareResponseModel {
	return &NullableAlipayTradeCustomsDeclareResponseModel{value: val, isSet: true}
}

func (v NullableAlipayTradeCustomsDeclareResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayTradeCustomsDeclareResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
