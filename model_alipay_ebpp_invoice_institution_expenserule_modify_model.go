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

// checks if the AlipayEbppInvoiceInstitutionExpenseruleModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceInstitutionExpenseruleModifyModel{}

// AlipayEbppInvoiceInstitutionExpenseruleModifyModel struct for AlipayEbppInvoiceInstitutionExpenseruleModifyModel
type AlipayEbppInvoiceInstitutionExpenseruleModifyModel struct {
	// 企业共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	// Deprecated
	AccountId *string `json:"account_id,omitempty"`
	// 修改使用规则
	Action *string `json:"action,omitempty"`
	// 授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	// Deprecated
	AgreementNo          *string               `json:"agreement_no,omitempty"`
	AssetShareSourceInfo *AssetShareSourceInfo `json:"asset_share_source_info,omitempty"`
	// 该使用规则支持的资产消费模式，不填写则为默认模式，默认模式下有余额时使用余额，没有余额则使用规则中的限额，点券模式为只能使用点券，点券+余额模式为可以使用点券和余额。
	ConsumeMode *string `json:"consume_mode,omitempty"`
	// 企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// 使用规则条件列表（已废弃）
	ExpenseCtrlRuleInfoList []ExpenseCtrRuleInfo `json:"expense_ctrl_rule_info_list,omitempty"`
	// 制度id
	InstitutionId *string `json:"institution_id,omitempty"`
	// 开票规则id
	OpenRuleId *string `json:"open_rule_id,omitempty"`
	// 当笔消费金额大于规则可用余额时，用于控制支付策略
	PaymentPolicy *string `json:"payment_policy,omitempty"`
	// 个人收款码转账是否支持因公付，默认为0。可选值：0（不支持）、1（支持）
	PersonalQrcodeMode *int32 `json:"personal_qrcode_mode,omitempty"`
	// 规则条件列表
	StandardConditionInfoList []StandardConditionInfo `json:"standard_condition_info_list,omitempty"`
	// 规则描述
	StandardDesc *string `json:"standard_desc,omitempty"`
	// 使用规则id
	StandardId *string `json:"standard_id,omitempty"`
	// 规则名称
	StandardName *string `json:"standard_name,omitempty"`
}

// NewAlipayEbppInvoiceInstitutionExpenseruleModifyModel instantiates a new AlipayEbppInvoiceInstitutionExpenseruleModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceInstitutionExpenseruleModifyModel() *AlipayEbppInvoiceInstitutionExpenseruleModifyModel {
	this := AlipayEbppInvoiceInstitutionExpenseruleModifyModel{}
	return &this
}

// NewAlipayEbppInvoiceInstitutionExpenseruleModifyModelWithDefaults instantiates a new AlipayEbppInvoiceInstitutionExpenseruleModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceInstitutionExpenseruleModifyModelWithDefaults() *AlipayEbppInvoiceInstitutionExpenseruleModifyModel {
	this := AlipayEbppInvoiceInstitutionExpenseruleModifyModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
// Deprecated
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
// Deprecated
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) SetAction(v string) {
	o.Action = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
// Deprecated
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
// Deprecated
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetAssetShareSourceInfo returns the AssetShareSourceInfo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetAssetShareSourceInfo() AssetShareSourceInfo {
	if o == nil || IsNil(o.AssetShareSourceInfo) {
		var ret AssetShareSourceInfo
		return ret
	}
	return *o.AssetShareSourceInfo
}

// GetAssetShareSourceInfoOk returns a tuple with the AssetShareSourceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetAssetShareSourceInfoOk() (*AssetShareSourceInfo, bool) {
	if o == nil || IsNil(o.AssetShareSourceInfo) {
		return nil, false
	}
	return o.AssetShareSourceInfo, true
}

// HasAssetShareSourceInfo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) HasAssetShareSourceInfo() bool {
	if o != nil && !IsNil(o.AssetShareSourceInfo) {
		return true
	}

	return false
}

// SetAssetShareSourceInfo gets a reference to the given AssetShareSourceInfo and assigns it to the AssetShareSourceInfo field.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) SetAssetShareSourceInfo(v AssetShareSourceInfo) {
	o.AssetShareSourceInfo = &v
}

// GetConsumeMode returns the ConsumeMode field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetConsumeMode() string {
	if o == nil || IsNil(o.ConsumeMode) {
		var ret string
		return ret
	}
	return *o.ConsumeMode
}

// GetConsumeModeOk returns a tuple with the ConsumeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetConsumeModeOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumeMode) {
		return nil, false
	}
	return o.ConsumeMode, true
}

// HasConsumeMode returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) HasConsumeMode() bool {
	if o != nil && !IsNil(o.ConsumeMode) {
		return true
	}

	return false
}

// SetConsumeMode gets a reference to the given string and assigns it to the ConsumeMode field.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) SetConsumeMode(v string) {
	o.ConsumeMode = &v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetEnterpriseId() string {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret string
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetEnterpriseIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given string and assigns it to the EnterpriseId field.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) SetEnterpriseId(v string) {
	o.EnterpriseId = &v
}

// GetExpenseCtrlRuleInfoList returns the ExpenseCtrlRuleInfoList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetExpenseCtrlRuleInfoList() []ExpenseCtrRuleInfo {
	if o == nil || IsNil(o.ExpenseCtrlRuleInfoList) {
		var ret []ExpenseCtrRuleInfo
		return ret
	}
	return o.ExpenseCtrlRuleInfoList
}

// GetExpenseCtrlRuleInfoListOk returns a tuple with the ExpenseCtrlRuleInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetExpenseCtrlRuleInfoListOk() ([]ExpenseCtrRuleInfo, bool) {
	if o == nil || IsNil(o.ExpenseCtrlRuleInfoList) {
		return nil, false
	}
	return o.ExpenseCtrlRuleInfoList, true
}

// HasExpenseCtrlRuleInfoList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) HasExpenseCtrlRuleInfoList() bool {
	if o != nil && !IsNil(o.ExpenseCtrlRuleInfoList) {
		return true
	}

	return false
}

// SetExpenseCtrlRuleInfoList gets a reference to the given []ExpenseCtrRuleInfo and assigns it to the ExpenseCtrlRuleInfoList field.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) SetExpenseCtrlRuleInfoList(v []ExpenseCtrRuleInfo) {
	o.ExpenseCtrlRuleInfoList = v
}

// GetInstitutionId returns the InstitutionId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetInstitutionId() string {
	if o == nil || IsNil(o.InstitutionId) {
		var ret string
		return ret
	}
	return *o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetInstitutionIdOk() (*string, bool) {
	if o == nil || IsNil(o.InstitutionId) {
		return nil, false
	}
	return o.InstitutionId, true
}

// HasInstitutionId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) HasInstitutionId() bool {
	if o != nil && !IsNil(o.InstitutionId) {
		return true
	}

	return false
}

// SetInstitutionId gets a reference to the given string and assigns it to the InstitutionId field.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) SetInstitutionId(v string) {
	o.InstitutionId = &v
}

// GetOpenRuleId returns the OpenRuleId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetOpenRuleId() string {
	if o == nil || IsNil(o.OpenRuleId) {
		var ret string
		return ret
	}
	return *o.OpenRuleId
}

// GetOpenRuleIdOk returns a tuple with the OpenRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetOpenRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenRuleId) {
		return nil, false
	}
	return o.OpenRuleId, true
}

// HasOpenRuleId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) HasOpenRuleId() bool {
	if o != nil && !IsNil(o.OpenRuleId) {
		return true
	}

	return false
}

// SetOpenRuleId gets a reference to the given string and assigns it to the OpenRuleId field.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) SetOpenRuleId(v string) {
	o.OpenRuleId = &v
}

// GetPaymentPolicy returns the PaymentPolicy field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetPaymentPolicy() string {
	if o == nil || IsNil(o.PaymentPolicy) {
		var ret string
		return ret
	}
	return *o.PaymentPolicy
}

// GetPaymentPolicyOk returns a tuple with the PaymentPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetPaymentPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentPolicy) {
		return nil, false
	}
	return o.PaymentPolicy, true
}

// HasPaymentPolicy returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) HasPaymentPolicy() bool {
	if o != nil && !IsNil(o.PaymentPolicy) {
		return true
	}

	return false
}

// SetPaymentPolicy gets a reference to the given string and assigns it to the PaymentPolicy field.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) SetPaymentPolicy(v string) {
	o.PaymentPolicy = &v
}

// GetPersonalQrcodeMode returns the PersonalQrcodeMode field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetPersonalQrcodeMode() int32 {
	if o == nil || IsNil(o.PersonalQrcodeMode) {
		var ret int32
		return ret
	}
	return *o.PersonalQrcodeMode
}

// GetPersonalQrcodeModeOk returns a tuple with the PersonalQrcodeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetPersonalQrcodeModeOk() (*int32, bool) {
	if o == nil || IsNil(o.PersonalQrcodeMode) {
		return nil, false
	}
	return o.PersonalQrcodeMode, true
}

// HasPersonalQrcodeMode returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) HasPersonalQrcodeMode() bool {
	if o != nil && !IsNil(o.PersonalQrcodeMode) {
		return true
	}

	return false
}

// SetPersonalQrcodeMode gets a reference to the given int32 and assigns it to the PersonalQrcodeMode field.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) SetPersonalQrcodeMode(v int32) {
	o.PersonalQrcodeMode = &v
}

// GetStandardConditionInfoList returns the StandardConditionInfoList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetStandardConditionInfoList() []StandardConditionInfo {
	if o == nil || IsNil(o.StandardConditionInfoList) {
		var ret []StandardConditionInfo
		return ret
	}
	return o.StandardConditionInfoList
}

// GetStandardConditionInfoListOk returns a tuple with the StandardConditionInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetStandardConditionInfoListOk() ([]StandardConditionInfo, bool) {
	if o == nil || IsNil(o.StandardConditionInfoList) {
		return nil, false
	}
	return o.StandardConditionInfoList, true
}

// HasStandardConditionInfoList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) HasStandardConditionInfoList() bool {
	if o != nil && !IsNil(o.StandardConditionInfoList) {
		return true
	}

	return false
}

// SetStandardConditionInfoList gets a reference to the given []StandardConditionInfo and assigns it to the StandardConditionInfoList field.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) SetStandardConditionInfoList(v []StandardConditionInfo) {
	o.StandardConditionInfoList = v
}

// GetStandardDesc returns the StandardDesc field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetStandardDesc() string {
	if o == nil || IsNil(o.StandardDesc) {
		var ret string
		return ret
	}
	return *o.StandardDesc
}

// GetStandardDescOk returns a tuple with the StandardDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetStandardDescOk() (*string, bool) {
	if o == nil || IsNil(o.StandardDesc) {
		return nil, false
	}
	return o.StandardDesc, true
}

// HasStandardDesc returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) HasStandardDesc() bool {
	if o != nil && !IsNil(o.StandardDesc) {
		return true
	}

	return false
}

// SetStandardDesc gets a reference to the given string and assigns it to the StandardDesc field.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) SetStandardDesc(v string) {
	o.StandardDesc = &v
}

// GetStandardId returns the StandardId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetStandardId() string {
	if o == nil || IsNil(o.StandardId) {
		var ret string
		return ret
	}
	return *o.StandardId
}

// GetStandardIdOk returns a tuple with the StandardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetStandardIdOk() (*string, bool) {
	if o == nil || IsNil(o.StandardId) {
		return nil, false
	}
	return o.StandardId, true
}

// HasStandardId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) HasStandardId() bool {
	if o != nil && !IsNil(o.StandardId) {
		return true
	}

	return false
}

// SetStandardId gets a reference to the given string and assigns it to the StandardId field.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) SetStandardId(v string) {
	o.StandardId = &v
}

// GetStandardName returns the StandardName field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetStandardName() string {
	if o == nil || IsNil(o.StandardName) {
		var ret string
		return ret
	}
	return *o.StandardName
}

// GetStandardNameOk returns a tuple with the StandardName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) GetStandardNameOk() (*string, bool) {
	if o == nil || IsNil(o.StandardName) {
		return nil, false
	}
	return o.StandardName, true
}

// HasStandardName returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) HasStandardName() bool {
	if o != nil && !IsNil(o.StandardName) {
		return true
	}

	return false
}

// SetStandardName gets a reference to the given string and assigns it to the StandardName field.
func (o *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) SetStandardName(v string) {
	o.StandardName = &v
}

func (o AlipayEbppInvoiceInstitutionExpenseruleModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceInstitutionExpenseruleModifyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.AssetShareSourceInfo) {
		toSerialize["asset_share_source_info"] = o.AssetShareSourceInfo
	}
	if !IsNil(o.ConsumeMode) {
		toSerialize["consume_mode"] = o.ConsumeMode
	}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterprise_id"] = o.EnterpriseId
	}
	if !IsNil(o.ExpenseCtrlRuleInfoList) {
		toSerialize["expense_ctrl_rule_info_list"] = o.ExpenseCtrlRuleInfoList
	}
	if !IsNil(o.InstitutionId) {
		toSerialize["institution_id"] = o.InstitutionId
	}
	if !IsNil(o.OpenRuleId) {
		toSerialize["open_rule_id"] = o.OpenRuleId
	}
	if !IsNil(o.PaymentPolicy) {
		toSerialize["payment_policy"] = o.PaymentPolicy
	}
	if !IsNil(o.PersonalQrcodeMode) {
		toSerialize["personal_qrcode_mode"] = o.PersonalQrcodeMode
	}
	if !IsNil(o.StandardConditionInfoList) {
		toSerialize["standard_condition_info_list"] = o.StandardConditionInfoList
	}
	if !IsNil(o.StandardDesc) {
		toSerialize["standard_desc"] = o.StandardDesc
	}
	if !IsNil(o.StandardId) {
		toSerialize["standard_id"] = o.StandardId
	}
	if !IsNil(o.StandardName) {
		toSerialize["standard_name"] = o.StandardName
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceInstitutionExpenseruleModifyModel struct {
	value *AlipayEbppInvoiceInstitutionExpenseruleModifyModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceInstitutionExpenseruleModifyModel) Get() *AlipayEbppInvoiceInstitutionExpenseruleModifyModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceInstitutionExpenseruleModifyModel) Set(val *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceInstitutionExpenseruleModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceInstitutionExpenseruleModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceInstitutionExpenseruleModifyModel(val *AlipayEbppInvoiceInstitutionExpenseruleModifyModel) *NullableAlipayEbppInvoiceInstitutionExpenseruleModifyModel {
	return &NullableAlipayEbppInvoiceInstitutionExpenseruleModifyModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceInstitutionExpenseruleModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceInstitutionExpenseruleModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
