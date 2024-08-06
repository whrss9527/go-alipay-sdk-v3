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

// checks if the AlipayEbppInvoiceExpenserulesSceneruleCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceExpenserulesSceneruleCreateModel{}

// AlipayEbppInvoiceExpenserulesSceneruleCreateModel struct for AlipayEbppInvoiceExpenserulesSceneruleCreateModel
type AlipayEbppInvoiceExpenserulesSceneruleCreateModel struct {
	// 共同账号ID
	AccountId *string `json:"account_id,omitempty"`
	// 授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// 有效期截止
	EffectiveEndDate *string `json:"effective_end_date,omitempty"`
	// 有效期起始
	EffectiveStartDate *string `json:"effective_start_date,omitempty"`
	// 切换open_id前请使用：员工支付宝uid列表 特殊说明：单次传入的最大员工数为50，后续通过修改费控规则员工接口进行员工调整
	EmployeeList []string `json:"employee_list,omitempty"`
	// 切换open_id后请使用：员工open_id/企业码员工ID列表 特殊说明：单次传入的最大员工数为50，后续通过修改费控规则员工接口进行员工调整
	EmployeeOpenIdList []string `json:"employee_open_id_list,omitempty"`
	// 企业码企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// 费控条件列表 特殊说明： 1）至少存在日额度（QUOTA_DAY）、月额度（QUOTA_MONTH）、有效期总额度（QUOTA_TOTAL）三者中的一个额度条件； 2）如果费用类型为MEAL，费控维度商户（MEAL_MERCHANT）和商户类型（MCC）对应的费控条件必须存在其一且不能同时存在； 3）如果费用类型为METRO，费控维度地铁卡类型（CARD_TYPE）对应的费控条件必须存在； 4）如果因公场景为OVERTIME，费控维度时间段（ALARM_CLOCK_TIME）对应的费控条件必须存在； 5）不能存在重复的费控维度对应的费控条件； 6）非MEAL费用类型，商户仅支持MERCHANT，不支持MEAL_MERCHANT
	ExpenseCtrlRuleInfoList []ExpenseCtrRuleInfo `json:"expense_ctrl_rule_info_list,omitempty"`
	// 费用类型
	ExpenseType *string `json:"expense_type,omitempty"`
	// 当笔消费金额大于规则可用余额时，用于控制支付策略，该字段缺省时采取因公账户和个人账户组合支付策略， 枚举值：PERSONAL（全部个人账户支付）, COMBINATION（因公账户和个人账户组合支付）
	PaymentPolicy *string `json:"payment_policy,omitempty"`
	// 因公场景
	SceneType *string `json:"scene_type,omitempty"`
	// 费控规则说明 特殊说明：敏感词校验
	StandardDesc *string `json:"standard_desc,omitempty"`
	// 费控规则名称 特殊说明： 1）敏感词校验 2）不能重复，若需要重复联系支持人员定向放开
	StandardName *string `json:"standard_name,omitempty"`
}

// NewAlipayEbppInvoiceExpenserulesSceneruleCreateModel instantiates a new AlipayEbppInvoiceExpenserulesSceneruleCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceExpenserulesSceneruleCreateModel() *AlipayEbppInvoiceExpenserulesSceneruleCreateModel {
	this := AlipayEbppInvoiceExpenserulesSceneruleCreateModel{}
	return &this
}

// NewAlipayEbppInvoiceExpenserulesSceneruleCreateModelWithDefaults instantiates a new AlipayEbppInvoiceExpenserulesSceneruleCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceExpenserulesSceneruleCreateModelWithDefaults() *AlipayEbppInvoiceExpenserulesSceneruleCreateModel {
	this := AlipayEbppInvoiceExpenserulesSceneruleCreateModel{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAgreementNo returns the AgreementNo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetAgreementNo() string {
	if o == nil || IsNil(o.AgreementNo) {
		var ret string
		return ret
	}
	return *o.AgreementNo
}

// GetAgreementNoOk returns a tuple with the AgreementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetAgreementNoOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementNo) {
		return nil, false
	}
	return o.AgreementNo, true
}

// HasAgreementNo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) HasAgreementNo() bool {
	if o != nil && !IsNil(o.AgreementNo) {
		return true
	}

	return false
}

// SetAgreementNo gets a reference to the given string and assigns it to the AgreementNo field.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) SetAgreementNo(v string) {
	o.AgreementNo = &v
}

// GetEffectiveEndDate returns the EffectiveEndDate field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetEffectiveEndDate() string {
	if o == nil || IsNil(o.EffectiveEndDate) {
		var ret string
		return ret
	}
	return *o.EffectiveEndDate
}

// GetEffectiveEndDateOk returns a tuple with the EffectiveEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetEffectiveEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveEndDate) {
		return nil, false
	}
	return o.EffectiveEndDate, true
}

// HasEffectiveEndDate returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) HasEffectiveEndDate() bool {
	if o != nil && !IsNil(o.EffectiveEndDate) {
		return true
	}

	return false
}

// SetEffectiveEndDate gets a reference to the given string and assigns it to the EffectiveEndDate field.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) SetEffectiveEndDate(v string) {
	o.EffectiveEndDate = &v
}

// GetEffectiveStartDate returns the EffectiveStartDate field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetEffectiveStartDate() string {
	if o == nil || IsNil(o.EffectiveStartDate) {
		var ret string
		return ret
	}
	return *o.EffectiveStartDate
}

// GetEffectiveStartDateOk returns a tuple with the EffectiveStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetEffectiveStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveStartDate) {
		return nil, false
	}
	return o.EffectiveStartDate, true
}

// HasEffectiveStartDate returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) HasEffectiveStartDate() bool {
	if o != nil && !IsNil(o.EffectiveStartDate) {
		return true
	}

	return false
}

// SetEffectiveStartDate gets a reference to the given string and assigns it to the EffectiveStartDate field.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) SetEffectiveStartDate(v string) {
	o.EffectiveStartDate = &v
}

// GetEmployeeList returns the EmployeeList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetEmployeeList() []string {
	if o == nil || IsNil(o.EmployeeList) {
		var ret []string
		return ret
	}
	return o.EmployeeList
}

// GetEmployeeListOk returns a tuple with the EmployeeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetEmployeeListOk() ([]string, bool) {
	if o == nil || IsNil(o.EmployeeList) {
		return nil, false
	}
	return o.EmployeeList, true
}

// HasEmployeeList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) HasEmployeeList() bool {
	if o != nil && !IsNil(o.EmployeeList) {
		return true
	}

	return false
}

// SetEmployeeList gets a reference to the given []string and assigns it to the EmployeeList field.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) SetEmployeeList(v []string) {
	o.EmployeeList = v
}

// GetEmployeeOpenIdList returns the EmployeeOpenIdList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetEmployeeOpenIdList() []string {
	if o == nil || IsNil(o.EmployeeOpenIdList) {
		var ret []string
		return ret
	}
	return o.EmployeeOpenIdList
}

// GetEmployeeOpenIdListOk returns a tuple with the EmployeeOpenIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetEmployeeOpenIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.EmployeeOpenIdList) {
		return nil, false
	}
	return o.EmployeeOpenIdList, true
}

// HasEmployeeOpenIdList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) HasEmployeeOpenIdList() bool {
	if o != nil && !IsNil(o.EmployeeOpenIdList) {
		return true
	}

	return false
}

// SetEmployeeOpenIdList gets a reference to the given []string and assigns it to the EmployeeOpenIdList field.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) SetEmployeeOpenIdList(v []string) {
	o.EmployeeOpenIdList = v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetEnterpriseId() string {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret string
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetEnterpriseIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given string and assigns it to the EnterpriseId field.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) SetEnterpriseId(v string) {
	o.EnterpriseId = &v
}

// GetExpenseCtrlRuleInfoList returns the ExpenseCtrlRuleInfoList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetExpenseCtrlRuleInfoList() []ExpenseCtrRuleInfo {
	if o == nil || IsNil(o.ExpenseCtrlRuleInfoList) {
		var ret []ExpenseCtrRuleInfo
		return ret
	}
	return o.ExpenseCtrlRuleInfoList
}

// GetExpenseCtrlRuleInfoListOk returns a tuple with the ExpenseCtrlRuleInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetExpenseCtrlRuleInfoListOk() ([]ExpenseCtrRuleInfo, bool) {
	if o == nil || IsNil(o.ExpenseCtrlRuleInfoList) {
		return nil, false
	}
	return o.ExpenseCtrlRuleInfoList, true
}

// HasExpenseCtrlRuleInfoList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) HasExpenseCtrlRuleInfoList() bool {
	if o != nil && !IsNil(o.ExpenseCtrlRuleInfoList) {
		return true
	}

	return false
}

// SetExpenseCtrlRuleInfoList gets a reference to the given []ExpenseCtrRuleInfo and assigns it to the ExpenseCtrlRuleInfoList field.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) SetExpenseCtrlRuleInfoList(v []ExpenseCtrRuleInfo) {
	o.ExpenseCtrlRuleInfoList = v
}

// GetExpenseType returns the ExpenseType field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetExpenseType() string {
	if o == nil || IsNil(o.ExpenseType) {
		var ret string
		return ret
	}
	return *o.ExpenseType
}

// GetExpenseTypeOk returns a tuple with the ExpenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetExpenseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpenseType) {
		return nil, false
	}
	return o.ExpenseType, true
}

// HasExpenseType returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) HasExpenseType() bool {
	if o != nil && !IsNil(o.ExpenseType) {
		return true
	}

	return false
}

// SetExpenseType gets a reference to the given string and assigns it to the ExpenseType field.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) SetExpenseType(v string) {
	o.ExpenseType = &v
}

// GetPaymentPolicy returns the PaymentPolicy field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetPaymentPolicy() string {
	if o == nil || IsNil(o.PaymentPolicy) {
		var ret string
		return ret
	}
	return *o.PaymentPolicy
}

// GetPaymentPolicyOk returns a tuple with the PaymentPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetPaymentPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentPolicy) {
		return nil, false
	}
	return o.PaymentPolicy, true
}

// HasPaymentPolicy returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) HasPaymentPolicy() bool {
	if o != nil && !IsNil(o.PaymentPolicy) {
		return true
	}

	return false
}

// SetPaymentPolicy gets a reference to the given string and assigns it to the PaymentPolicy field.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) SetPaymentPolicy(v string) {
	o.PaymentPolicy = &v
}

// GetSceneType returns the SceneType field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetSceneType() string {
	if o == nil || IsNil(o.SceneType) {
		var ret string
		return ret
	}
	return *o.SceneType
}

// GetSceneTypeOk returns a tuple with the SceneType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetSceneTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SceneType) {
		return nil, false
	}
	return o.SceneType, true
}

// HasSceneType returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) HasSceneType() bool {
	if o != nil && !IsNil(o.SceneType) {
		return true
	}

	return false
}

// SetSceneType gets a reference to the given string and assigns it to the SceneType field.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) SetSceneType(v string) {
	o.SceneType = &v
}

// GetStandardDesc returns the StandardDesc field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetStandardDesc() string {
	if o == nil || IsNil(o.StandardDesc) {
		var ret string
		return ret
	}
	return *o.StandardDesc
}

// GetStandardDescOk returns a tuple with the StandardDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetStandardDescOk() (*string, bool) {
	if o == nil || IsNil(o.StandardDesc) {
		return nil, false
	}
	return o.StandardDesc, true
}

// HasStandardDesc returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) HasStandardDesc() bool {
	if o != nil && !IsNil(o.StandardDesc) {
		return true
	}

	return false
}

// SetStandardDesc gets a reference to the given string and assigns it to the StandardDesc field.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) SetStandardDesc(v string) {
	o.StandardDesc = &v
}

// GetStandardName returns the StandardName field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetStandardName() string {
	if o == nil || IsNil(o.StandardName) {
		var ret string
		return ret
	}
	return *o.StandardName
}

// GetStandardNameOk returns a tuple with the StandardName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) GetStandardNameOk() (*string, bool) {
	if o == nil || IsNil(o.StandardName) {
		return nil, false
	}
	return o.StandardName, true
}

// HasStandardName returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) HasStandardName() bool {
	if o != nil && !IsNil(o.StandardName) {
		return true
	}

	return false
}

// SetStandardName gets a reference to the given string and assigns it to the StandardName field.
func (o *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) SetStandardName(v string) {
	o.StandardName = &v
}

func (o AlipayEbppInvoiceExpenserulesSceneruleCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceExpenserulesSceneruleCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.AgreementNo) {
		toSerialize["agreement_no"] = o.AgreementNo
	}
	if !IsNil(o.EffectiveEndDate) {
		toSerialize["effective_end_date"] = o.EffectiveEndDate
	}
	if !IsNil(o.EffectiveStartDate) {
		toSerialize["effective_start_date"] = o.EffectiveStartDate
	}
	if !IsNil(o.EmployeeList) {
		toSerialize["employee_list"] = o.EmployeeList
	}
	if !IsNil(o.EmployeeOpenIdList) {
		toSerialize["employee_open_id_list"] = o.EmployeeOpenIdList
	}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterprise_id"] = o.EnterpriseId
	}
	if !IsNil(o.ExpenseCtrlRuleInfoList) {
		toSerialize["expense_ctrl_rule_info_list"] = o.ExpenseCtrlRuleInfoList
	}
	if !IsNil(o.ExpenseType) {
		toSerialize["expense_type"] = o.ExpenseType
	}
	if !IsNil(o.PaymentPolicy) {
		toSerialize["payment_policy"] = o.PaymentPolicy
	}
	if !IsNil(o.SceneType) {
		toSerialize["scene_type"] = o.SceneType
	}
	if !IsNil(o.StandardDesc) {
		toSerialize["standard_desc"] = o.StandardDesc
	}
	if !IsNil(o.StandardName) {
		toSerialize["standard_name"] = o.StandardName
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceExpenserulesSceneruleCreateModel struct {
	value *AlipayEbppInvoiceExpenserulesSceneruleCreateModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceExpenserulesSceneruleCreateModel) Get() *AlipayEbppInvoiceExpenserulesSceneruleCreateModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceExpenserulesSceneruleCreateModel) Set(val *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceExpenserulesSceneruleCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceExpenserulesSceneruleCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceExpenserulesSceneruleCreateModel(val *AlipayEbppInvoiceExpenserulesSceneruleCreateModel) *NullableAlipayEbppInvoiceExpenserulesSceneruleCreateModel {
	return &NullableAlipayEbppInvoiceExpenserulesSceneruleCreateModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceExpenserulesSceneruleCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceExpenserulesSceneruleCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


