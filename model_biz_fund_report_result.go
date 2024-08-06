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

// checks if the BizFundReportResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BizFundReportResult{}

// BizFundReportResult struct for BizFundReportResult
type BizFundReportResult struct {
	// 实付金额（元）
	AlipayAmount *string `json:"alipay_amount,omitempty"`
	// 是否批量转账
	BatchType *string `json:"batch_type,omitempty"`
	// 业务类型
	BizType *string `json:"biz_type,omitempty"`
	// 业务类型描述
	BizTypeDesc *string `json:"biz_type_desc,omitempty"`
	// 服务费（元）
	ChargeFee *string `json:"charge_fee,omitempty"`
	// 交易时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// 2020xxx
	InstructionId *string `json:"instruction_id,omitempty"`
	// 出借
	Memo *string `json:"memo,omitempty"`
	// 是否脱敏
	NoMbillEncrypt *string `json:"no_mbill_encrypt,omitempty"`
	// 业务订单号
	OrderId *string `json:"order_id,omitempty"`
	// 实付金额（元）
	PayAmount *string `json:"pay_amount,omitempty"`
	// 收款方银行卡号
	PayeeCardNo *string `json:"payee_card_no,omitempty"`
	// 收款方全名
	PayeeFullName *string `json:"payee_full_name,omitempty"`
	// 收款方资金类型
	PayeeFundType *string `json:"payee_fund_type,omitempty"`
	// 支付宝余额
	PayeeFundTypeDesc *string `json:"payee_fund_type_desc,omitempty"`
	// 收款银行中文简称
	PayeeInstName *string `json:"payee_inst_name,omitempty"`
	// 收款方登录Id
	PayeeLoginEmail *string `json:"payee_login_email,omitempty"`
	// 收款方姓名
	PayeeName *string `json:"payee_name,omitempty"`
	// 付款方资金类型
	PayerFundType *string `json:"payer_fund_type,omitempty"`
	// 付款方资金类型描述
	PayerFundTypeDesc *string `json:"payer_fund_type_desc,omitempty"`
	// 退款金额（元）
	RefundAmount *string `json:"refund_amount,omitempty"`
	// 是否展示细节，默认为true
	ShowDetail *string `json:"show_detail,omitempty"`
	// 是否出示凭证
	ShowVoucher *string `json:"show_voucher,omitempty"`
	// 数据来源
	Source *string `json:"source,omitempty"`
	// 单据状态
	Status *string `json:"status,omitempty"`
	// 单据状态描述
	StatusDesc *string `json:"status_desc,omitempty"`
	// 普通转账
	SubBizType *string `json:"sub_biz_type,omitempty"`
	// 业务子类型描述
	SubBizTypeDesc *string `json:"sub_biz_type_desc,omitempty"`
}

// NewBizFundReportResult instantiates a new BizFundReportResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBizFundReportResult() *BizFundReportResult {
	this := BizFundReportResult{}
	return &this
}

// NewBizFundReportResultWithDefaults instantiates a new BizFundReportResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBizFundReportResultWithDefaults() *BizFundReportResult {
	this := BizFundReportResult{}
	return &this
}

// GetAlipayAmount returns the AlipayAmount field value if set, zero value otherwise.
func (o *BizFundReportResult) GetAlipayAmount() string {
	if o == nil || IsNil(o.AlipayAmount) {
		var ret string
		return ret
	}
	return *o.AlipayAmount
}

// GetAlipayAmountOk returns a tuple with the AlipayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetAlipayAmountOk() (*string, bool) {
	if o == nil || IsNil(o.AlipayAmount) {
		return nil, false
	}
	return o.AlipayAmount, true
}

// HasAlipayAmount returns a boolean if a field has been set.
func (o *BizFundReportResult) HasAlipayAmount() bool {
	if o != nil && !IsNil(o.AlipayAmount) {
		return true
	}

	return false
}

// SetAlipayAmount gets a reference to the given string and assigns it to the AlipayAmount field.
func (o *BizFundReportResult) SetAlipayAmount(v string) {
	o.AlipayAmount = &v
}

// GetBatchType returns the BatchType field value if set, zero value otherwise.
func (o *BizFundReportResult) GetBatchType() string {
	if o == nil || IsNil(o.BatchType) {
		var ret string
		return ret
	}
	return *o.BatchType
}

// GetBatchTypeOk returns a tuple with the BatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetBatchTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BatchType) {
		return nil, false
	}
	return o.BatchType, true
}

// HasBatchType returns a boolean if a field has been set.
func (o *BizFundReportResult) HasBatchType() bool {
	if o != nil && !IsNil(o.BatchType) {
		return true
	}

	return false
}

// SetBatchType gets a reference to the given string and assigns it to the BatchType field.
func (o *BizFundReportResult) SetBatchType(v string) {
	o.BatchType = &v
}

// GetBizType returns the BizType field value if set, zero value otherwise.
func (o *BizFundReportResult) GetBizType() string {
	if o == nil || IsNil(o.BizType) {
		var ret string
		return ret
	}
	return *o.BizType
}

// GetBizTypeOk returns a tuple with the BizType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetBizTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BizType) {
		return nil, false
	}
	return o.BizType, true
}

// HasBizType returns a boolean if a field has been set.
func (o *BizFundReportResult) HasBizType() bool {
	if o != nil && !IsNil(o.BizType) {
		return true
	}

	return false
}

// SetBizType gets a reference to the given string and assigns it to the BizType field.
func (o *BizFundReportResult) SetBizType(v string) {
	o.BizType = &v
}

// GetBizTypeDesc returns the BizTypeDesc field value if set, zero value otherwise.
func (o *BizFundReportResult) GetBizTypeDesc() string {
	if o == nil || IsNil(o.BizTypeDesc) {
		var ret string
		return ret
	}
	return *o.BizTypeDesc
}

// GetBizTypeDescOk returns a tuple with the BizTypeDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetBizTypeDescOk() (*string, bool) {
	if o == nil || IsNil(o.BizTypeDesc) {
		return nil, false
	}
	return o.BizTypeDesc, true
}

// HasBizTypeDesc returns a boolean if a field has been set.
func (o *BizFundReportResult) HasBizTypeDesc() bool {
	if o != nil && !IsNil(o.BizTypeDesc) {
		return true
	}

	return false
}

// SetBizTypeDesc gets a reference to the given string and assigns it to the BizTypeDesc field.
func (o *BizFundReportResult) SetBizTypeDesc(v string) {
	o.BizTypeDesc = &v
}

// GetChargeFee returns the ChargeFee field value if set, zero value otherwise.
func (o *BizFundReportResult) GetChargeFee() string {
	if o == nil || IsNil(o.ChargeFee) {
		var ret string
		return ret
	}
	return *o.ChargeFee
}

// GetChargeFeeOk returns a tuple with the ChargeFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetChargeFeeOk() (*string, bool) {
	if o == nil || IsNil(o.ChargeFee) {
		return nil, false
	}
	return o.ChargeFee, true
}

// HasChargeFee returns a boolean if a field has been set.
func (o *BizFundReportResult) HasChargeFee() bool {
	if o != nil && !IsNil(o.ChargeFee) {
		return true
	}

	return false
}

// SetChargeFee gets a reference to the given string and assigns it to the ChargeFee field.
func (o *BizFundReportResult) SetChargeFee(v string) {
	o.ChargeFee = &v
}

// GetGmtCreate returns the GmtCreate field value if set, zero value otherwise.
func (o *BizFundReportResult) GetGmtCreate() string {
	if o == nil || IsNil(o.GmtCreate) {
		var ret string
		return ret
	}
	return *o.GmtCreate
}

// GetGmtCreateOk returns a tuple with the GmtCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetGmtCreateOk() (*string, bool) {
	if o == nil || IsNil(o.GmtCreate) {
		return nil, false
	}
	return o.GmtCreate, true
}

// HasGmtCreate returns a boolean if a field has been set.
func (o *BizFundReportResult) HasGmtCreate() bool {
	if o != nil && !IsNil(o.GmtCreate) {
		return true
	}

	return false
}

// SetGmtCreate gets a reference to the given string and assigns it to the GmtCreate field.
func (o *BizFundReportResult) SetGmtCreate(v string) {
	o.GmtCreate = &v
}

// GetInstructionId returns the InstructionId field value if set, zero value otherwise.
func (o *BizFundReportResult) GetInstructionId() string {
	if o == nil || IsNil(o.InstructionId) {
		var ret string
		return ret
	}
	return *o.InstructionId
}

// GetInstructionIdOk returns a tuple with the InstructionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetInstructionIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstructionId) {
		return nil, false
	}
	return o.InstructionId, true
}

// HasInstructionId returns a boolean if a field has been set.
func (o *BizFundReportResult) HasInstructionId() bool {
	if o != nil && !IsNil(o.InstructionId) {
		return true
	}

	return false
}

// SetInstructionId gets a reference to the given string and assigns it to the InstructionId field.
func (o *BizFundReportResult) SetInstructionId(v string) {
	o.InstructionId = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *BizFundReportResult) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *BizFundReportResult) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *BizFundReportResult) SetMemo(v string) {
	o.Memo = &v
}

// GetNoMbillEncrypt returns the NoMbillEncrypt field value if set, zero value otherwise.
func (o *BizFundReportResult) GetNoMbillEncrypt() string {
	if o == nil || IsNil(o.NoMbillEncrypt) {
		var ret string
		return ret
	}
	return *o.NoMbillEncrypt
}

// GetNoMbillEncryptOk returns a tuple with the NoMbillEncrypt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetNoMbillEncryptOk() (*string, bool) {
	if o == nil || IsNil(o.NoMbillEncrypt) {
		return nil, false
	}
	return o.NoMbillEncrypt, true
}

// HasNoMbillEncrypt returns a boolean if a field has been set.
func (o *BizFundReportResult) HasNoMbillEncrypt() bool {
	if o != nil && !IsNil(o.NoMbillEncrypt) {
		return true
	}

	return false
}

// SetNoMbillEncrypt gets a reference to the given string and assigns it to the NoMbillEncrypt field.
func (o *BizFundReportResult) SetNoMbillEncrypt(v string) {
	o.NoMbillEncrypt = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *BizFundReportResult) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *BizFundReportResult) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *BizFundReportResult) SetOrderId(v string) {
	o.OrderId = &v
}

// GetPayAmount returns the PayAmount field value if set, zero value otherwise.
func (o *BizFundReportResult) GetPayAmount() string {
	if o == nil || IsNil(o.PayAmount) {
		var ret string
		return ret
	}
	return *o.PayAmount
}

// GetPayAmountOk returns a tuple with the PayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetPayAmountOk() (*string, bool) {
	if o == nil || IsNil(o.PayAmount) {
		return nil, false
	}
	return o.PayAmount, true
}

// HasPayAmount returns a boolean if a field has been set.
func (o *BizFundReportResult) HasPayAmount() bool {
	if o != nil && !IsNil(o.PayAmount) {
		return true
	}

	return false
}

// SetPayAmount gets a reference to the given string and assigns it to the PayAmount field.
func (o *BizFundReportResult) SetPayAmount(v string) {
	o.PayAmount = &v
}

// GetPayeeCardNo returns the PayeeCardNo field value if set, zero value otherwise.
func (o *BizFundReportResult) GetPayeeCardNo() string {
	if o == nil || IsNil(o.PayeeCardNo) {
		var ret string
		return ret
	}
	return *o.PayeeCardNo
}

// GetPayeeCardNoOk returns a tuple with the PayeeCardNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetPayeeCardNoOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeCardNo) {
		return nil, false
	}
	return o.PayeeCardNo, true
}

// HasPayeeCardNo returns a boolean if a field has been set.
func (o *BizFundReportResult) HasPayeeCardNo() bool {
	if o != nil && !IsNil(o.PayeeCardNo) {
		return true
	}

	return false
}

// SetPayeeCardNo gets a reference to the given string and assigns it to the PayeeCardNo field.
func (o *BizFundReportResult) SetPayeeCardNo(v string) {
	o.PayeeCardNo = &v
}

// GetPayeeFullName returns the PayeeFullName field value if set, zero value otherwise.
func (o *BizFundReportResult) GetPayeeFullName() string {
	if o == nil || IsNil(o.PayeeFullName) {
		var ret string
		return ret
	}
	return *o.PayeeFullName
}

// GetPayeeFullNameOk returns a tuple with the PayeeFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetPayeeFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeFullName) {
		return nil, false
	}
	return o.PayeeFullName, true
}

// HasPayeeFullName returns a boolean if a field has been set.
func (o *BizFundReportResult) HasPayeeFullName() bool {
	if o != nil && !IsNil(o.PayeeFullName) {
		return true
	}

	return false
}

// SetPayeeFullName gets a reference to the given string and assigns it to the PayeeFullName field.
func (o *BizFundReportResult) SetPayeeFullName(v string) {
	o.PayeeFullName = &v
}

// GetPayeeFundType returns the PayeeFundType field value if set, zero value otherwise.
func (o *BizFundReportResult) GetPayeeFundType() string {
	if o == nil || IsNil(o.PayeeFundType) {
		var ret string
		return ret
	}
	return *o.PayeeFundType
}

// GetPayeeFundTypeOk returns a tuple with the PayeeFundType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetPayeeFundTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeFundType) {
		return nil, false
	}
	return o.PayeeFundType, true
}

// HasPayeeFundType returns a boolean if a field has been set.
func (o *BizFundReportResult) HasPayeeFundType() bool {
	if o != nil && !IsNil(o.PayeeFundType) {
		return true
	}

	return false
}

// SetPayeeFundType gets a reference to the given string and assigns it to the PayeeFundType field.
func (o *BizFundReportResult) SetPayeeFundType(v string) {
	o.PayeeFundType = &v
}

// GetPayeeFundTypeDesc returns the PayeeFundTypeDesc field value if set, zero value otherwise.
func (o *BizFundReportResult) GetPayeeFundTypeDesc() string {
	if o == nil || IsNil(o.PayeeFundTypeDesc) {
		var ret string
		return ret
	}
	return *o.PayeeFundTypeDesc
}

// GetPayeeFundTypeDescOk returns a tuple with the PayeeFundTypeDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetPayeeFundTypeDescOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeFundTypeDesc) {
		return nil, false
	}
	return o.PayeeFundTypeDesc, true
}

// HasPayeeFundTypeDesc returns a boolean if a field has been set.
func (o *BizFundReportResult) HasPayeeFundTypeDesc() bool {
	if o != nil && !IsNil(o.PayeeFundTypeDesc) {
		return true
	}

	return false
}

// SetPayeeFundTypeDesc gets a reference to the given string and assigns it to the PayeeFundTypeDesc field.
func (o *BizFundReportResult) SetPayeeFundTypeDesc(v string) {
	o.PayeeFundTypeDesc = &v
}

// GetPayeeInstName returns the PayeeInstName field value if set, zero value otherwise.
func (o *BizFundReportResult) GetPayeeInstName() string {
	if o == nil || IsNil(o.PayeeInstName) {
		var ret string
		return ret
	}
	return *o.PayeeInstName
}

// GetPayeeInstNameOk returns a tuple with the PayeeInstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetPayeeInstNameOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeInstName) {
		return nil, false
	}
	return o.PayeeInstName, true
}

// HasPayeeInstName returns a boolean if a field has been set.
func (o *BizFundReportResult) HasPayeeInstName() bool {
	if o != nil && !IsNil(o.PayeeInstName) {
		return true
	}

	return false
}

// SetPayeeInstName gets a reference to the given string and assigns it to the PayeeInstName field.
func (o *BizFundReportResult) SetPayeeInstName(v string) {
	o.PayeeInstName = &v
}

// GetPayeeLoginEmail returns the PayeeLoginEmail field value if set, zero value otherwise.
func (o *BizFundReportResult) GetPayeeLoginEmail() string {
	if o == nil || IsNil(o.PayeeLoginEmail) {
		var ret string
		return ret
	}
	return *o.PayeeLoginEmail
}

// GetPayeeLoginEmailOk returns a tuple with the PayeeLoginEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetPayeeLoginEmailOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeLoginEmail) {
		return nil, false
	}
	return o.PayeeLoginEmail, true
}

// HasPayeeLoginEmail returns a boolean if a field has been set.
func (o *BizFundReportResult) HasPayeeLoginEmail() bool {
	if o != nil && !IsNil(o.PayeeLoginEmail) {
		return true
	}

	return false
}

// SetPayeeLoginEmail gets a reference to the given string and assigns it to the PayeeLoginEmail field.
func (o *BizFundReportResult) SetPayeeLoginEmail(v string) {
	o.PayeeLoginEmail = &v
}

// GetPayeeName returns the PayeeName field value if set, zero value otherwise.
func (o *BizFundReportResult) GetPayeeName() string {
	if o == nil || IsNil(o.PayeeName) {
		var ret string
		return ret
	}
	return *o.PayeeName
}

// GetPayeeNameOk returns a tuple with the PayeeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetPayeeNameOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeName) {
		return nil, false
	}
	return o.PayeeName, true
}

// HasPayeeName returns a boolean if a field has been set.
func (o *BizFundReportResult) HasPayeeName() bool {
	if o != nil && !IsNil(o.PayeeName) {
		return true
	}

	return false
}

// SetPayeeName gets a reference to the given string and assigns it to the PayeeName field.
func (o *BizFundReportResult) SetPayeeName(v string) {
	o.PayeeName = &v
}

// GetPayerFundType returns the PayerFundType field value if set, zero value otherwise.
func (o *BizFundReportResult) GetPayerFundType() string {
	if o == nil || IsNil(o.PayerFundType) {
		var ret string
		return ret
	}
	return *o.PayerFundType
}

// GetPayerFundTypeOk returns a tuple with the PayerFundType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetPayerFundTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PayerFundType) {
		return nil, false
	}
	return o.PayerFundType, true
}

// HasPayerFundType returns a boolean if a field has been set.
func (o *BizFundReportResult) HasPayerFundType() bool {
	if o != nil && !IsNil(o.PayerFundType) {
		return true
	}

	return false
}

// SetPayerFundType gets a reference to the given string and assigns it to the PayerFundType field.
func (o *BizFundReportResult) SetPayerFundType(v string) {
	o.PayerFundType = &v
}

// GetPayerFundTypeDesc returns the PayerFundTypeDesc field value if set, zero value otherwise.
func (o *BizFundReportResult) GetPayerFundTypeDesc() string {
	if o == nil || IsNil(o.PayerFundTypeDesc) {
		var ret string
		return ret
	}
	return *o.PayerFundTypeDesc
}

// GetPayerFundTypeDescOk returns a tuple with the PayerFundTypeDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetPayerFundTypeDescOk() (*string, bool) {
	if o == nil || IsNil(o.PayerFundTypeDesc) {
		return nil, false
	}
	return o.PayerFundTypeDesc, true
}

// HasPayerFundTypeDesc returns a boolean if a field has been set.
func (o *BizFundReportResult) HasPayerFundTypeDesc() bool {
	if o != nil && !IsNil(o.PayerFundTypeDesc) {
		return true
	}

	return false
}

// SetPayerFundTypeDesc gets a reference to the given string and assigns it to the PayerFundTypeDesc field.
func (o *BizFundReportResult) SetPayerFundTypeDesc(v string) {
	o.PayerFundTypeDesc = &v
}

// GetRefundAmount returns the RefundAmount field value if set, zero value otherwise.
func (o *BizFundReportResult) GetRefundAmount() string {
	if o == nil || IsNil(o.RefundAmount) {
		var ret string
		return ret
	}
	return *o.RefundAmount
}

// GetRefundAmountOk returns a tuple with the RefundAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetRefundAmountOk() (*string, bool) {
	if o == nil || IsNil(o.RefundAmount) {
		return nil, false
	}
	return o.RefundAmount, true
}

// HasRefundAmount returns a boolean if a field has been set.
func (o *BizFundReportResult) HasRefundAmount() bool {
	if o != nil && !IsNil(o.RefundAmount) {
		return true
	}

	return false
}

// SetRefundAmount gets a reference to the given string and assigns it to the RefundAmount field.
func (o *BizFundReportResult) SetRefundAmount(v string) {
	o.RefundAmount = &v
}

// GetShowDetail returns the ShowDetail field value if set, zero value otherwise.
func (o *BizFundReportResult) GetShowDetail() string {
	if o == nil || IsNil(o.ShowDetail) {
		var ret string
		return ret
	}
	return *o.ShowDetail
}

// GetShowDetailOk returns a tuple with the ShowDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetShowDetailOk() (*string, bool) {
	if o == nil || IsNil(o.ShowDetail) {
		return nil, false
	}
	return o.ShowDetail, true
}

// HasShowDetail returns a boolean if a field has been set.
func (o *BizFundReportResult) HasShowDetail() bool {
	if o != nil && !IsNil(o.ShowDetail) {
		return true
	}

	return false
}

// SetShowDetail gets a reference to the given string and assigns it to the ShowDetail field.
func (o *BizFundReportResult) SetShowDetail(v string) {
	o.ShowDetail = &v
}

// GetShowVoucher returns the ShowVoucher field value if set, zero value otherwise.
func (o *BizFundReportResult) GetShowVoucher() string {
	if o == nil || IsNil(o.ShowVoucher) {
		var ret string
		return ret
	}
	return *o.ShowVoucher
}

// GetShowVoucherOk returns a tuple with the ShowVoucher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetShowVoucherOk() (*string, bool) {
	if o == nil || IsNil(o.ShowVoucher) {
		return nil, false
	}
	return o.ShowVoucher, true
}

// HasShowVoucher returns a boolean if a field has been set.
func (o *BizFundReportResult) HasShowVoucher() bool {
	if o != nil && !IsNil(o.ShowVoucher) {
		return true
	}

	return false
}

// SetShowVoucher gets a reference to the given string and assigns it to the ShowVoucher field.
func (o *BizFundReportResult) SetShowVoucher(v string) {
	o.ShowVoucher = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *BizFundReportResult) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *BizFundReportResult) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *BizFundReportResult) SetSource(v string) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BizFundReportResult) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BizFundReportResult) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BizFundReportResult) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDesc returns the StatusDesc field value if set, zero value otherwise.
func (o *BizFundReportResult) GetStatusDesc() string {
	if o == nil || IsNil(o.StatusDesc) {
		var ret string
		return ret
	}
	return *o.StatusDesc
}

// GetStatusDescOk returns a tuple with the StatusDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetStatusDescOk() (*string, bool) {
	if o == nil || IsNil(o.StatusDesc) {
		return nil, false
	}
	return o.StatusDesc, true
}

// HasStatusDesc returns a boolean if a field has been set.
func (o *BizFundReportResult) HasStatusDesc() bool {
	if o != nil && !IsNil(o.StatusDesc) {
		return true
	}

	return false
}

// SetStatusDesc gets a reference to the given string and assigns it to the StatusDesc field.
func (o *BizFundReportResult) SetStatusDesc(v string) {
	o.StatusDesc = &v
}

// GetSubBizType returns the SubBizType field value if set, zero value otherwise.
func (o *BizFundReportResult) GetSubBizType() string {
	if o == nil || IsNil(o.SubBizType) {
		var ret string
		return ret
	}
	return *o.SubBizType
}

// GetSubBizTypeOk returns a tuple with the SubBizType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetSubBizTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubBizType) {
		return nil, false
	}
	return o.SubBizType, true
}

// HasSubBizType returns a boolean if a field has been set.
func (o *BizFundReportResult) HasSubBizType() bool {
	if o != nil && !IsNil(o.SubBizType) {
		return true
	}

	return false
}

// SetSubBizType gets a reference to the given string and assigns it to the SubBizType field.
func (o *BizFundReportResult) SetSubBizType(v string) {
	o.SubBizType = &v
}

// GetSubBizTypeDesc returns the SubBizTypeDesc field value if set, zero value otherwise.
func (o *BizFundReportResult) GetSubBizTypeDesc() string {
	if o == nil || IsNil(o.SubBizTypeDesc) {
		var ret string
		return ret
	}
	return *o.SubBizTypeDesc
}

// GetSubBizTypeDescOk returns a tuple with the SubBizTypeDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BizFundReportResult) GetSubBizTypeDescOk() (*string, bool) {
	if o == nil || IsNil(o.SubBizTypeDesc) {
		return nil, false
	}
	return o.SubBizTypeDesc, true
}

// HasSubBizTypeDesc returns a boolean if a field has been set.
func (o *BizFundReportResult) HasSubBizTypeDesc() bool {
	if o != nil && !IsNil(o.SubBizTypeDesc) {
		return true
	}

	return false
}

// SetSubBizTypeDesc gets a reference to the given string and assigns it to the SubBizTypeDesc field.
func (o *BizFundReportResult) SetSubBizTypeDesc(v string) {
	o.SubBizTypeDesc = &v
}

func (o BizFundReportResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BizFundReportResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlipayAmount) {
		toSerialize["alipay_amount"] = o.AlipayAmount
	}
	if !IsNil(o.BatchType) {
		toSerialize["batch_type"] = o.BatchType
	}
	if !IsNil(o.BizType) {
		toSerialize["biz_type"] = o.BizType
	}
	if !IsNil(o.BizTypeDesc) {
		toSerialize["biz_type_desc"] = o.BizTypeDesc
	}
	if !IsNil(o.ChargeFee) {
		toSerialize["charge_fee"] = o.ChargeFee
	}
	if !IsNil(o.GmtCreate) {
		toSerialize["gmt_create"] = o.GmtCreate
	}
	if !IsNil(o.InstructionId) {
		toSerialize["instruction_id"] = o.InstructionId
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.NoMbillEncrypt) {
		toSerialize["no_mbill_encrypt"] = o.NoMbillEncrypt
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.PayAmount) {
		toSerialize["pay_amount"] = o.PayAmount
	}
	if !IsNil(o.PayeeCardNo) {
		toSerialize["payee_card_no"] = o.PayeeCardNo
	}
	if !IsNil(o.PayeeFullName) {
		toSerialize["payee_full_name"] = o.PayeeFullName
	}
	if !IsNil(o.PayeeFundType) {
		toSerialize["payee_fund_type"] = o.PayeeFundType
	}
	if !IsNil(o.PayeeFundTypeDesc) {
		toSerialize["payee_fund_type_desc"] = o.PayeeFundTypeDesc
	}
	if !IsNil(o.PayeeInstName) {
		toSerialize["payee_inst_name"] = o.PayeeInstName
	}
	if !IsNil(o.PayeeLoginEmail) {
		toSerialize["payee_login_email"] = o.PayeeLoginEmail
	}
	if !IsNil(o.PayeeName) {
		toSerialize["payee_name"] = o.PayeeName
	}
	if !IsNil(o.PayerFundType) {
		toSerialize["payer_fund_type"] = o.PayerFundType
	}
	if !IsNil(o.PayerFundTypeDesc) {
		toSerialize["payer_fund_type_desc"] = o.PayerFundTypeDesc
	}
	if !IsNil(o.RefundAmount) {
		toSerialize["refund_amount"] = o.RefundAmount
	}
	if !IsNil(o.ShowDetail) {
		toSerialize["show_detail"] = o.ShowDetail
	}
	if !IsNil(o.ShowVoucher) {
		toSerialize["show_voucher"] = o.ShowVoucher
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusDesc) {
		toSerialize["status_desc"] = o.StatusDesc
	}
	if !IsNil(o.SubBizType) {
		toSerialize["sub_biz_type"] = o.SubBizType
	}
	if !IsNil(o.SubBizTypeDesc) {
		toSerialize["sub_biz_type_desc"] = o.SubBizTypeDesc
	}
	return toSerialize, nil
}

type NullableBizFundReportResult struct {
	value *BizFundReportResult
	isSet bool
}

func (v NullableBizFundReportResult) Get() *BizFundReportResult {
	return v.value
}

func (v *NullableBizFundReportResult) Set(val *BizFundReportResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBizFundReportResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBizFundReportResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBizFundReportResult(val *BizFundReportResult) *NullableBizFundReportResult {
	return &NullableBizFundReportResult{value: val, isSet: true}
}

func (v NullableBizFundReportResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBizFundReportResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


