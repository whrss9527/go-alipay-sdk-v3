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

// checks if the AlipayEbppPdeductAsyncPayModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppPdeductAsyncPayModel{}

// AlipayEbppPdeductAsyncPayModel struct for AlipayEbppPdeductAsyncPayModel
type AlipayEbppPdeductAsyncPayModel struct {
	// 分配给外部机构发起扣款时的渠道码。朗新为LANGXIN
	AgentChannel *string `json:"agent_channel,omitempty"`
	// 二级渠道码，预留字段
	AgentCode *string `json:"agent_code,omitempty"`
	// 支付宝代扣协议Id
	AgreementId *string `json:"agreement_id,omitempty"`
	// 账期
	BillDate *string `json:"bill_date,omitempty"`
	// 户号
	BillKey *string `json:"bill_key,omitempty"`
	// 扩展字段
	ExtendField *string `json:"extend_field,omitempty"`
	// 滞纳金
	FineAmount *string `json:"fine_amount,omitempty"`
	// 备注信息
	Memo *string `json:"memo,omitempty"`
	// 用户UserId在应用AppId下的唯一用户标识
	OpenId *string `json:"open_id,omitempty"`
	// 商户外部业务流水号
	OutOrderNo *string `json:"out_order_no,omitempty"`
	// 扣款金额，支付总金额，包含滞纳金
	PayAmount *string `json:"pay_amount,omitempty"`
	// 商户partnerId
	Pid *string `json:"pid,omitempty"`
	// 用户ID
	UserId *string `json:"user_id,omitempty"`
}

// NewAlipayEbppPdeductAsyncPayModel instantiates a new AlipayEbppPdeductAsyncPayModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppPdeductAsyncPayModel() *AlipayEbppPdeductAsyncPayModel {
	this := AlipayEbppPdeductAsyncPayModel{}
	return &this
}

// NewAlipayEbppPdeductAsyncPayModelWithDefaults instantiates a new AlipayEbppPdeductAsyncPayModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppPdeductAsyncPayModelWithDefaults() *AlipayEbppPdeductAsyncPayModel {
	this := AlipayEbppPdeductAsyncPayModel{}
	return &this
}

// GetAgentChannel returns the AgentChannel field value if set, zero value otherwise.
func (o *AlipayEbppPdeductAsyncPayModel) GetAgentChannel() string {
	if o == nil || IsNil(o.AgentChannel) {
		var ret string
		return ret
	}
	return *o.AgentChannel
}

// GetAgentChannelOk returns a tuple with the AgentChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductAsyncPayModel) GetAgentChannelOk() (*string, bool) {
	if o == nil || IsNil(o.AgentChannel) {
		return nil, false
	}
	return o.AgentChannel, true
}

// HasAgentChannel returns a boolean if a field has been set.
func (o *AlipayEbppPdeductAsyncPayModel) HasAgentChannel() bool {
	if o != nil && !IsNil(o.AgentChannel) {
		return true
	}

	return false
}

// SetAgentChannel gets a reference to the given string and assigns it to the AgentChannel field.
func (o *AlipayEbppPdeductAsyncPayModel) SetAgentChannel(v string) {
	o.AgentChannel = &v
}

// GetAgentCode returns the AgentCode field value if set, zero value otherwise.
func (o *AlipayEbppPdeductAsyncPayModel) GetAgentCode() string {
	if o == nil || IsNil(o.AgentCode) {
		var ret string
		return ret
	}
	return *o.AgentCode
}

// GetAgentCodeOk returns a tuple with the AgentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductAsyncPayModel) GetAgentCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AgentCode) {
		return nil, false
	}
	return o.AgentCode, true
}

// HasAgentCode returns a boolean if a field has been set.
func (o *AlipayEbppPdeductAsyncPayModel) HasAgentCode() bool {
	if o != nil && !IsNil(o.AgentCode) {
		return true
	}

	return false
}

// SetAgentCode gets a reference to the given string and assigns it to the AgentCode field.
func (o *AlipayEbppPdeductAsyncPayModel) SetAgentCode(v string) {
	o.AgentCode = &v
}

// GetAgreementId returns the AgreementId field value if set, zero value otherwise.
func (o *AlipayEbppPdeductAsyncPayModel) GetAgreementId() string {
	if o == nil || IsNil(o.AgreementId) {
		var ret string
		return ret
	}
	return *o.AgreementId
}

// GetAgreementIdOk returns a tuple with the AgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductAsyncPayModel) GetAgreementIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementId) {
		return nil, false
	}
	return o.AgreementId, true
}

// HasAgreementId returns a boolean if a field has been set.
func (o *AlipayEbppPdeductAsyncPayModel) HasAgreementId() bool {
	if o != nil && !IsNil(o.AgreementId) {
		return true
	}

	return false
}

// SetAgreementId gets a reference to the given string and assigns it to the AgreementId field.
func (o *AlipayEbppPdeductAsyncPayModel) SetAgreementId(v string) {
	o.AgreementId = &v
}

// GetBillDate returns the BillDate field value if set, zero value otherwise.
func (o *AlipayEbppPdeductAsyncPayModel) GetBillDate() string {
	if o == nil || IsNil(o.BillDate) {
		var ret string
		return ret
	}
	return *o.BillDate
}

// GetBillDateOk returns a tuple with the BillDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductAsyncPayModel) GetBillDateOk() (*string, bool) {
	if o == nil || IsNil(o.BillDate) {
		return nil, false
	}
	return o.BillDate, true
}

// HasBillDate returns a boolean if a field has been set.
func (o *AlipayEbppPdeductAsyncPayModel) HasBillDate() bool {
	if o != nil && !IsNil(o.BillDate) {
		return true
	}

	return false
}

// SetBillDate gets a reference to the given string and assigns it to the BillDate field.
func (o *AlipayEbppPdeductAsyncPayModel) SetBillDate(v string) {
	o.BillDate = &v
}

// GetBillKey returns the BillKey field value if set, zero value otherwise.
func (o *AlipayEbppPdeductAsyncPayModel) GetBillKey() string {
	if o == nil || IsNil(o.BillKey) {
		var ret string
		return ret
	}
	return *o.BillKey
}

// GetBillKeyOk returns a tuple with the BillKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductAsyncPayModel) GetBillKeyOk() (*string, bool) {
	if o == nil || IsNil(o.BillKey) {
		return nil, false
	}
	return o.BillKey, true
}

// HasBillKey returns a boolean if a field has been set.
func (o *AlipayEbppPdeductAsyncPayModel) HasBillKey() bool {
	if o != nil && !IsNil(o.BillKey) {
		return true
	}

	return false
}

// SetBillKey gets a reference to the given string and assigns it to the BillKey field.
func (o *AlipayEbppPdeductAsyncPayModel) SetBillKey(v string) {
	o.BillKey = &v
}

// GetExtendField returns the ExtendField field value if set, zero value otherwise.
func (o *AlipayEbppPdeductAsyncPayModel) GetExtendField() string {
	if o == nil || IsNil(o.ExtendField) {
		var ret string
		return ret
	}
	return *o.ExtendField
}

// GetExtendFieldOk returns a tuple with the ExtendField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductAsyncPayModel) GetExtendFieldOk() (*string, bool) {
	if o == nil || IsNil(o.ExtendField) {
		return nil, false
	}
	return o.ExtendField, true
}

// HasExtendField returns a boolean if a field has been set.
func (o *AlipayEbppPdeductAsyncPayModel) HasExtendField() bool {
	if o != nil && !IsNil(o.ExtendField) {
		return true
	}

	return false
}

// SetExtendField gets a reference to the given string and assigns it to the ExtendField field.
func (o *AlipayEbppPdeductAsyncPayModel) SetExtendField(v string) {
	o.ExtendField = &v
}

// GetFineAmount returns the FineAmount field value if set, zero value otherwise.
func (o *AlipayEbppPdeductAsyncPayModel) GetFineAmount() string {
	if o == nil || IsNil(o.FineAmount) {
		var ret string
		return ret
	}
	return *o.FineAmount
}

// GetFineAmountOk returns a tuple with the FineAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductAsyncPayModel) GetFineAmountOk() (*string, bool) {
	if o == nil || IsNil(o.FineAmount) {
		return nil, false
	}
	return o.FineAmount, true
}

// HasFineAmount returns a boolean if a field has been set.
func (o *AlipayEbppPdeductAsyncPayModel) HasFineAmount() bool {
	if o != nil && !IsNil(o.FineAmount) {
		return true
	}

	return false
}

// SetFineAmount gets a reference to the given string and assigns it to the FineAmount field.
func (o *AlipayEbppPdeductAsyncPayModel) SetFineAmount(v string) {
	o.FineAmount = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *AlipayEbppPdeductAsyncPayModel) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductAsyncPayModel) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *AlipayEbppPdeductAsyncPayModel) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *AlipayEbppPdeductAsyncPayModel) SetMemo(v string) {
	o.Memo = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *AlipayEbppPdeductAsyncPayModel) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductAsyncPayModel) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *AlipayEbppPdeductAsyncPayModel) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *AlipayEbppPdeductAsyncPayModel) SetOpenId(v string) {
	o.OpenId = &v
}

// GetOutOrderNo returns the OutOrderNo field value if set, zero value otherwise.
func (o *AlipayEbppPdeductAsyncPayModel) GetOutOrderNo() string {
	if o == nil || IsNil(o.OutOrderNo) {
		var ret string
		return ret
	}
	return *o.OutOrderNo
}

// GetOutOrderNoOk returns a tuple with the OutOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductAsyncPayModel) GetOutOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutOrderNo) {
		return nil, false
	}
	return o.OutOrderNo, true
}

// HasOutOrderNo returns a boolean if a field has been set.
func (o *AlipayEbppPdeductAsyncPayModel) HasOutOrderNo() bool {
	if o != nil && !IsNil(o.OutOrderNo) {
		return true
	}

	return false
}

// SetOutOrderNo gets a reference to the given string and assigns it to the OutOrderNo field.
func (o *AlipayEbppPdeductAsyncPayModel) SetOutOrderNo(v string) {
	o.OutOrderNo = &v
}

// GetPayAmount returns the PayAmount field value if set, zero value otherwise.
func (o *AlipayEbppPdeductAsyncPayModel) GetPayAmount() string {
	if o == nil || IsNil(o.PayAmount) {
		var ret string
		return ret
	}
	return *o.PayAmount
}

// GetPayAmountOk returns a tuple with the PayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductAsyncPayModel) GetPayAmountOk() (*string, bool) {
	if o == nil || IsNil(o.PayAmount) {
		return nil, false
	}
	return o.PayAmount, true
}

// HasPayAmount returns a boolean if a field has been set.
func (o *AlipayEbppPdeductAsyncPayModel) HasPayAmount() bool {
	if o != nil && !IsNil(o.PayAmount) {
		return true
	}

	return false
}

// SetPayAmount gets a reference to the given string and assigns it to the PayAmount field.
func (o *AlipayEbppPdeductAsyncPayModel) SetPayAmount(v string) {
	o.PayAmount = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *AlipayEbppPdeductAsyncPayModel) GetPid() string {
	if o == nil || IsNil(o.Pid) {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductAsyncPayModel) GetPidOk() (*string, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *AlipayEbppPdeductAsyncPayModel) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *AlipayEbppPdeductAsyncPayModel) SetPid(v string) {
	o.Pid = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AlipayEbppPdeductAsyncPayModel) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductAsyncPayModel) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AlipayEbppPdeductAsyncPayModel) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AlipayEbppPdeductAsyncPayModel) SetUserId(v string) {
	o.UserId = &v
}

func (o AlipayEbppPdeductAsyncPayModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppPdeductAsyncPayModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgentChannel) {
		toSerialize["agent_channel"] = o.AgentChannel
	}
	if !IsNil(o.AgentCode) {
		toSerialize["agent_code"] = o.AgentCode
	}
	if !IsNil(o.AgreementId) {
		toSerialize["agreement_id"] = o.AgreementId
	}
	if !IsNil(o.BillDate) {
		toSerialize["bill_date"] = o.BillDate
	}
	if !IsNil(o.BillKey) {
		toSerialize["bill_key"] = o.BillKey
	}
	if !IsNil(o.ExtendField) {
		toSerialize["extend_field"] = o.ExtendField
	}
	if !IsNil(o.FineAmount) {
		toSerialize["fine_amount"] = o.FineAmount
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.OutOrderNo) {
		toSerialize["out_order_no"] = o.OutOrderNo
	}
	if !IsNil(o.PayAmount) {
		toSerialize["pay_amount"] = o.PayAmount
	}
	if !IsNil(o.Pid) {
		toSerialize["pid"] = o.Pid
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableAlipayEbppPdeductAsyncPayModel struct {
	value *AlipayEbppPdeductAsyncPayModel
	isSet bool
}

func (v NullableAlipayEbppPdeductAsyncPayModel) Get() *AlipayEbppPdeductAsyncPayModel {
	return v.value
}

func (v *NullableAlipayEbppPdeductAsyncPayModel) Set(val *AlipayEbppPdeductAsyncPayModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppPdeductAsyncPayModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppPdeductAsyncPayModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppPdeductAsyncPayModel(val *AlipayEbppPdeductAsyncPayModel) *NullableAlipayEbppPdeductAsyncPayModel {
	return &NullableAlipayEbppPdeductAsyncPayModel{value: val, isSet: true}
}

func (v NullableAlipayEbppPdeductAsyncPayModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppPdeductAsyncPayModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
