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

// checks if the AlipayEbppInvoiceExpenserulesProjectrulesModifyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceExpenserulesProjectrulesModifyModel{}

// AlipayEbppInvoiceExpenserulesProjectrulesModifyModel struct for AlipayEbppInvoiceExpenserulesProjectrulesModifyModel
type AlipayEbppInvoiceExpenserulesProjectrulesModifyModel struct {
	// 企业ID
	AccountId *string `json:"account_id,omitempty"`
	// 修改操作 枚举值：MODIFY_RULE（修改费控条件），仅支持MODIFY_RULE
	Action *string `json:"action,omitempty"`
	// 授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 费控规则列表
	ExpenseCtrlRuleInfoGroupList []ExpenseCtrRuleGroupInfo `json:"expense_ctrl_rule_info_group_list,omitempty"`
	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`
}

// NewAlipayEbppInvoiceExpenserulesProjectrulesModifyModel instantiates a new AlipayEbppInvoiceExpenserulesProjectrulesModifyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceExpenserulesProjectrulesModifyModel() *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel {
	this := AlipayEbppInvoiceExpenserulesProjectrulesModifyModel{}
	return &this
}

// NewAlipayEbppInvoiceExpenserulesProjectrulesModifyModelWithDefaults instantiates a new AlipayEbppInvoiceExpenserulesProjectrulesModifyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceExpenserulesProjectrulesModifyModelWithDefaults() *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel {
	this := AlipayEbppInvoiceExpenserulesProjectrulesModifyModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) SetAction(v string) {
	o.Action = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetExpenseCtrlRuleInfoGroupList returns the ExpenseCtrlRuleInfoGroupList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) GetExpenseCtrlRuleInfoGroupList() []ExpenseCtrRuleGroupInfo {
	if o == nil || IsNil(o.ExpenseCtrlRuleInfoGroupList) {
		var ret []ExpenseCtrRuleGroupInfo
		return ret
	}
	return o.ExpenseCtrlRuleInfoGroupList
}

// GetExpenseCtrlRuleInfoGroupListOk returns a tuple with the ExpenseCtrlRuleInfoGroupList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) GetExpenseCtrlRuleInfoGroupListOk() ([]ExpenseCtrRuleGroupInfo, bool) {
	if o == nil || IsNil(o.ExpenseCtrlRuleInfoGroupList) {
		return nil, false
	}
	return o.ExpenseCtrlRuleInfoGroupList, true
}

// HasExpenseCtrlRuleInfoGroupList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) HasExpenseCtrlRuleInfoGroupList() bool {
	if o != nil && !IsNil(o.ExpenseCtrlRuleInfoGroupList) {
		return true
	}

	return false
}

// SetExpenseCtrlRuleInfoGroupList gets a reference to the given []ExpenseCtrRuleGroupInfo and assigns it to the ExpenseCtrlRuleInfoGroupList field.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) SetExpenseCtrlRuleInfoGroupList(v []ExpenseCtrRuleGroupInfo) {
	o.ExpenseCtrlRuleInfoGroupList = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) SetProjectId(v string) {
	o.ProjectId = &v
}

func (o AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ExpenseCtrlRuleInfoGroupList) {
		toSerialize["expense_ctrl_rule_info_group_list"] = o.ExpenseCtrlRuleInfoGroupList
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceExpenserulesProjectrulesModifyModel struct {
	value *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceExpenserulesProjectrulesModifyModel) Get() *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceExpenserulesProjectrulesModifyModel) Set(val *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceExpenserulesProjectrulesModifyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceExpenserulesProjectrulesModifyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceExpenserulesProjectrulesModifyModel(val *AlipayEbppInvoiceExpenserulesProjectrulesModifyModel) *NullableAlipayEbppInvoiceExpenserulesProjectrulesModifyModel {
	return &NullableAlipayEbppInvoiceExpenserulesProjectrulesModifyModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceExpenserulesProjectrulesModifyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceExpenserulesProjectrulesModifyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
