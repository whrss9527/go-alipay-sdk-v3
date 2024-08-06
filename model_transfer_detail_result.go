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

// checks if the TransferDetailResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferDetailResult{}

// TransferDetailResult struct for TransferDetailResult
type TransferDetailResult struct {
	// 付款/收款账户。充值记录中是付款账户。提现、转账记录中是收款账户。支付宝名称及账号脱敏；银行账户的户名脱敏，银行账户显示银行名称+银行卡号后四位
	Account *string `json:"account,omitempty"`
	// 金额
	Amount *string `json:"amount,omitempty"`
	// 资金来源/去向类型。在充值记录中，表示资金来源类型，在转账和提现类型中，表示去向类型
	FundDesc *string `json:"fund_desc,omitempty"`
	// 银行单据号。对账使用，无需脱敏
	InstructionId *string `json:"instruction_id,omitempty"`
	// 备注信息
	Memo *string `json:"memo,omitempty"`
	// 业务订单号。该笔业务单据的唯一识别编号
	OrderNo *string `json:"order_no,omitempty"`
	// 服务费金额
	ServiceFee *string `json:"service_fee,omitempty"`
	// 资金状态
	Status *string `json:"status,omitempty"`
	// 子类型。“充值类型”，普通充值、大额充值。“转账类型”，暂无实现。转账至支付宝账户、转账至银行卡、批量转账支付宝账户、批量转账至银行卡、批量付款。“提现类型”，暂无实现。普通提现、批量委托提现。对账使用，无需脱敏
	SubTypeDesc *string `json:"sub_type_desc,omitempty"`
	// 业务发生时间
	TransDt *string `json:"trans_dt,omitempty"`
	// 查询类型描述：充值、转账、提现
	TypeDesc *string `json:"type_desc,omitempty"`
}

// NewTransferDetailResult instantiates a new TransferDetailResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferDetailResult() *TransferDetailResult {
	this := TransferDetailResult{}
	return &this
}

// NewTransferDetailResultWithDefaults instantiates a new TransferDetailResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferDetailResultWithDefaults() *TransferDetailResult {
	this := TransferDetailResult{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *TransferDetailResult) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferDetailResult) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *TransferDetailResult) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *TransferDetailResult) SetAccount(v string) {
	o.Account = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *TransferDetailResult) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferDetailResult) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *TransferDetailResult) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *TransferDetailResult) SetAmount(v string) {
	o.Amount = &v
}

// GetFundDesc returns the FundDesc field value if set, zero value otherwise.
func (o *TransferDetailResult) GetFundDesc() string {
	if o == nil || IsNil(o.FundDesc) {
		var ret string
		return ret
	}
	return *o.FundDesc
}

// GetFundDescOk returns a tuple with the FundDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferDetailResult) GetFundDescOk() (*string, bool) {
	if o == nil || IsNil(o.FundDesc) {
		return nil, false
	}
	return o.FundDesc, true
}

// HasFundDesc returns a boolean if a field has been set.
func (o *TransferDetailResult) HasFundDesc() bool {
	if o != nil && !IsNil(o.FundDesc) {
		return true
	}

	return false
}

// SetFundDesc gets a reference to the given string and assigns it to the FundDesc field.
func (o *TransferDetailResult) SetFundDesc(v string) {
	o.FundDesc = &v
}

// GetInstructionId returns the InstructionId field value if set, zero value otherwise.
func (o *TransferDetailResult) GetInstructionId() string {
	if o == nil || IsNil(o.InstructionId) {
		var ret string
		return ret
	}
	return *o.InstructionId
}

// GetInstructionIdOk returns a tuple with the InstructionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferDetailResult) GetInstructionIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstructionId) {
		return nil, false
	}
	return o.InstructionId, true
}

// HasInstructionId returns a boolean if a field has been set.
func (o *TransferDetailResult) HasInstructionId() bool {
	if o != nil && !IsNil(o.InstructionId) {
		return true
	}

	return false
}

// SetInstructionId gets a reference to the given string and assigns it to the InstructionId field.
func (o *TransferDetailResult) SetInstructionId(v string) {
	o.InstructionId = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *TransferDetailResult) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferDetailResult) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *TransferDetailResult) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *TransferDetailResult) SetMemo(v string) {
	o.Memo = &v
}

// GetOrderNo returns the OrderNo field value if set, zero value otherwise.
func (o *TransferDetailResult) GetOrderNo() string {
	if o == nil || IsNil(o.OrderNo) {
		var ret string
		return ret
	}
	return *o.OrderNo
}

// GetOrderNoOk returns a tuple with the OrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferDetailResult) GetOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OrderNo) {
		return nil, false
	}
	return o.OrderNo, true
}

// HasOrderNo returns a boolean if a field has been set.
func (o *TransferDetailResult) HasOrderNo() bool {
	if o != nil && !IsNil(o.OrderNo) {
		return true
	}

	return false
}

// SetOrderNo gets a reference to the given string and assigns it to the OrderNo field.
func (o *TransferDetailResult) SetOrderNo(v string) {
	o.OrderNo = &v
}

// GetServiceFee returns the ServiceFee field value if set, zero value otherwise.
func (o *TransferDetailResult) GetServiceFee() string {
	if o == nil || IsNil(o.ServiceFee) {
		var ret string
		return ret
	}
	return *o.ServiceFee
}

// GetServiceFeeOk returns a tuple with the ServiceFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferDetailResult) GetServiceFeeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceFee) {
		return nil, false
	}
	return o.ServiceFee, true
}

// HasServiceFee returns a boolean if a field has been set.
func (o *TransferDetailResult) HasServiceFee() bool {
	if o != nil && !IsNil(o.ServiceFee) {
		return true
	}

	return false
}

// SetServiceFee gets a reference to the given string and assigns it to the ServiceFee field.
func (o *TransferDetailResult) SetServiceFee(v string) {
	o.ServiceFee = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TransferDetailResult) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferDetailResult) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TransferDetailResult) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TransferDetailResult) SetStatus(v string) {
	o.Status = &v
}

// GetSubTypeDesc returns the SubTypeDesc field value if set, zero value otherwise.
func (o *TransferDetailResult) GetSubTypeDesc() string {
	if o == nil || IsNil(o.SubTypeDesc) {
		var ret string
		return ret
	}
	return *o.SubTypeDesc
}

// GetSubTypeDescOk returns a tuple with the SubTypeDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferDetailResult) GetSubTypeDescOk() (*string, bool) {
	if o == nil || IsNil(o.SubTypeDesc) {
		return nil, false
	}
	return o.SubTypeDesc, true
}

// HasSubTypeDesc returns a boolean if a field has been set.
func (o *TransferDetailResult) HasSubTypeDesc() bool {
	if o != nil && !IsNil(o.SubTypeDesc) {
		return true
	}

	return false
}

// SetSubTypeDesc gets a reference to the given string and assigns it to the SubTypeDesc field.
func (o *TransferDetailResult) SetSubTypeDesc(v string) {
	o.SubTypeDesc = &v
}

// GetTransDt returns the TransDt field value if set, zero value otherwise.
func (o *TransferDetailResult) GetTransDt() string {
	if o == nil || IsNil(o.TransDt) {
		var ret string
		return ret
	}
	return *o.TransDt
}

// GetTransDtOk returns a tuple with the TransDt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferDetailResult) GetTransDtOk() (*string, bool) {
	if o == nil || IsNil(o.TransDt) {
		return nil, false
	}
	return o.TransDt, true
}

// HasTransDt returns a boolean if a field has been set.
func (o *TransferDetailResult) HasTransDt() bool {
	if o != nil && !IsNil(o.TransDt) {
		return true
	}

	return false
}

// SetTransDt gets a reference to the given string and assigns it to the TransDt field.
func (o *TransferDetailResult) SetTransDt(v string) {
	o.TransDt = &v
}

// GetTypeDesc returns the TypeDesc field value if set, zero value otherwise.
func (o *TransferDetailResult) GetTypeDesc() string {
	if o == nil || IsNil(o.TypeDesc) {
		var ret string
		return ret
	}
	return *o.TypeDesc
}

// GetTypeDescOk returns a tuple with the TypeDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferDetailResult) GetTypeDescOk() (*string, bool) {
	if o == nil || IsNil(o.TypeDesc) {
		return nil, false
	}
	return o.TypeDesc, true
}

// HasTypeDesc returns a boolean if a field has been set.
func (o *TransferDetailResult) HasTypeDesc() bool {
	if o != nil && !IsNil(o.TypeDesc) {
		return true
	}

	return false
}

// SetTypeDesc gets a reference to the given string and assigns it to the TypeDesc field.
func (o *TransferDetailResult) SetTypeDesc(v string) {
	o.TypeDesc = &v
}

func (o TransferDetailResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferDetailResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.FundDesc) {
		toSerialize["fund_desc"] = o.FundDesc
	}
	if !IsNil(o.InstructionId) {
		toSerialize["instruction_id"] = o.InstructionId
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.OrderNo) {
		toSerialize["order_no"] = o.OrderNo
	}
	if !IsNil(o.ServiceFee) {
		toSerialize["service_fee"] = o.ServiceFee
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubTypeDesc) {
		toSerialize["sub_type_desc"] = o.SubTypeDesc
	}
	if !IsNil(o.TransDt) {
		toSerialize["trans_dt"] = o.TransDt
	}
	if !IsNil(o.TypeDesc) {
		toSerialize["type_desc"] = o.TypeDesc
	}
	return toSerialize, nil
}

type NullableTransferDetailResult struct {
	value *TransferDetailResult
	isSet bool
}

func (v NullableTransferDetailResult) Get() *TransferDetailResult {
	return v.value
}

func (v *NullableTransferDetailResult) Set(val *TransferDetailResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferDetailResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferDetailResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferDetailResult(val *TransferDetailResult) *NullableTransferDetailResult {
	return &NullableTransferDetailResult{value: val, isSet: true}
}

func (v NullableTransferDetailResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferDetailResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
