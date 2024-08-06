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

// checks if the StandardInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardInfo{}

// StandardInfo struct for StandardInfo
type StandardInfo struct {
	AssetShareSourceInfo *AssetShareSourceInfo `json:"asset_share_source_info,omitempty"`
	// 消费模式，不填为默认模式，枚举值：COUPON_ONLY（仅支持点券） COUPON_AND_CAP（支持点券+余额） COUNT（仅支持次卡） DEFAULT（默认模式） 点券：消费时找员工的点券，没有或者用完了不可付； 点券+余额：消费时找员工的点券，没有找员工的余额，没有或者用完了不可付； 次卡：消费时找员工的次卡，没有或者用完了不可付； 默认：有给员工设置员工余额以员工余额为准，用完为止。否则只受规则里的限额和企业账户资金上限管控；
	ConsumeMode *string `json:"consume_mode,omitempty"`
	// 费用类型子类，当制度的费用类型为MEAL时，支持到店、外卖；当费用类型为非MEAL时，与费用类型保持一致
	ExpenseTypeSubCategory *string `json:"expense_type_sub_category,omitempty"`
	// 开票规则id，可通过接口alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.create 创建并得到开票规则ID
	OpenRuleId *string `json:"open_rule_id,omitempty"`
	// 外部使用规则id，制度内使用规则该字段不允许重复
	OuterSourceId *string `json:"outer_source_id,omitempty"`
	// 支付策略 当笔消费金额大于规则可用余额时，用于控制支付策略，该字段缺省时采取因公账户和个人账户组合支付策略
	PaymentPolicy *string `json:"payment_policy,omitempty"`
	// 个人收款码转账是否支持因公付，默认为0。可选值：0（不支持）、1（支持）
	PersonalQrcodeMode *int32 `json:"personal_qrcode_mode,omitempty"`
	// 使用规则条件列表
	StandardConditionInfoList []StandardConditionInfo `json:"standard_condition_info_list,omitempty"`
	// 使用规则描述(敏感词校验)
	StandardDesc *string `json:"standard_desc,omitempty"`
	// 制度ID（创建使用规则时非必填）
	StandardId *string `json:"standard_id,omitempty"`
	// 规则名称
	StandardName *string `json:"standard_name,omitempty"`
}

// NewStandardInfo instantiates a new StandardInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardInfo() *StandardInfo {
	this := StandardInfo{}
	return &this
}

// NewStandardInfoWithDefaults instantiates a new StandardInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardInfoWithDefaults() *StandardInfo {
	this := StandardInfo{}
	return &this
}

// GetAssetShareSourceInfo returns the AssetShareSourceInfo field value if set, zero value otherwise.
func (o *StandardInfo) GetAssetShareSourceInfo() AssetShareSourceInfo {
	if o == nil || IsNil(o.AssetShareSourceInfo) {
		var ret AssetShareSourceInfo
		return ret
	}
	return *o.AssetShareSourceInfo
}

// GetAssetShareSourceInfoOk returns a tuple with the AssetShareSourceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardInfo) GetAssetShareSourceInfoOk() (*AssetShareSourceInfo, bool) {
	if o == nil || IsNil(o.AssetShareSourceInfo) {
		return nil, false
	}
	return o.AssetShareSourceInfo, true
}

// HasAssetShareSourceInfo returns a boolean if a field has been set.
func (o *StandardInfo) HasAssetShareSourceInfo() bool {
	if o != nil && !IsNil(o.AssetShareSourceInfo) {
		return true
	}

	return false
}

// SetAssetShareSourceInfo gets a reference to the given AssetShareSourceInfo and assigns it to the AssetShareSourceInfo field.
func (o *StandardInfo) SetAssetShareSourceInfo(v AssetShareSourceInfo) {
	o.AssetShareSourceInfo = &v
}

// GetConsumeMode returns the ConsumeMode field value if set, zero value otherwise.
func (o *StandardInfo) GetConsumeMode() string {
	if o == nil || IsNil(o.ConsumeMode) {
		var ret string
		return ret
	}
	return *o.ConsumeMode
}

// GetConsumeModeOk returns a tuple with the ConsumeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardInfo) GetConsumeModeOk() (*string, bool) {
	if o == nil || IsNil(o.ConsumeMode) {
		return nil, false
	}
	return o.ConsumeMode, true
}

// HasConsumeMode returns a boolean if a field has been set.
func (o *StandardInfo) HasConsumeMode() bool {
	if o != nil && !IsNil(o.ConsumeMode) {
		return true
	}

	return false
}

// SetConsumeMode gets a reference to the given string and assigns it to the ConsumeMode field.
func (o *StandardInfo) SetConsumeMode(v string) {
	o.ConsumeMode = &v
}

// GetExpenseTypeSubCategory returns the ExpenseTypeSubCategory field value if set, zero value otherwise.
func (o *StandardInfo) GetExpenseTypeSubCategory() string {
	if o == nil || IsNil(o.ExpenseTypeSubCategory) {
		var ret string
		return ret
	}
	return *o.ExpenseTypeSubCategory
}

// GetExpenseTypeSubCategoryOk returns a tuple with the ExpenseTypeSubCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardInfo) GetExpenseTypeSubCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.ExpenseTypeSubCategory) {
		return nil, false
	}
	return o.ExpenseTypeSubCategory, true
}

// HasExpenseTypeSubCategory returns a boolean if a field has been set.
func (o *StandardInfo) HasExpenseTypeSubCategory() bool {
	if o != nil && !IsNil(o.ExpenseTypeSubCategory) {
		return true
	}

	return false
}

// SetExpenseTypeSubCategory gets a reference to the given string and assigns it to the ExpenseTypeSubCategory field.
func (o *StandardInfo) SetExpenseTypeSubCategory(v string) {
	o.ExpenseTypeSubCategory = &v
}

// GetOpenRuleId returns the OpenRuleId field value if set, zero value otherwise.
func (o *StandardInfo) GetOpenRuleId() string {
	if o == nil || IsNil(o.OpenRuleId) {
		var ret string
		return ret
	}
	return *o.OpenRuleId
}

// GetOpenRuleIdOk returns a tuple with the OpenRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardInfo) GetOpenRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenRuleId) {
		return nil, false
	}
	return o.OpenRuleId, true
}

// HasOpenRuleId returns a boolean if a field has been set.
func (o *StandardInfo) HasOpenRuleId() bool {
	if o != nil && !IsNil(o.OpenRuleId) {
		return true
	}

	return false
}

// SetOpenRuleId gets a reference to the given string and assigns it to the OpenRuleId field.
func (o *StandardInfo) SetOpenRuleId(v string) {
	o.OpenRuleId = &v
}

// GetOuterSourceId returns the OuterSourceId field value if set, zero value otherwise.
func (o *StandardInfo) GetOuterSourceId() string {
	if o == nil || IsNil(o.OuterSourceId) {
		var ret string
		return ret
	}
	return *o.OuterSourceId
}

// GetOuterSourceIdOk returns a tuple with the OuterSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardInfo) GetOuterSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.OuterSourceId) {
		return nil, false
	}
	return o.OuterSourceId, true
}

// HasOuterSourceId returns a boolean if a field has been set.
func (o *StandardInfo) HasOuterSourceId() bool {
	if o != nil && !IsNil(o.OuterSourceId) {
		return true
	}

	return false
}

// SetOuterSourceId gets a reference to the given string and assigns it to the OuterSourceId field.
func (o *StandardInfo) SetOuterSourceId(v string) {
	o.OuterSourceId = &v
}

// GetPaymentPolicy returns the PaymentPolicy field value if set, zero value otherwise.
func (o *StandardInfo) GetPaymentPolicy() string {
	if o == nil || IsNil(o.PaymentPolicy) {
		var ret string
		return ret
	}
	return *o.PaymentPolicy
}

// GetPaymentPolicyOk returns a tuple with the PaymentPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardInfo) GetPaymentPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentPolicy) {
		return nil, false
	}
	return o.PaymentPolicy, true
}

// HasPaymentPolicy returns a boolean if a field has been set.
func (o *StandardInfo) HasPaymentPolicy() bool {
	if o != nil && !IsNil(o.PaymentPolicy) {
		return true
	}

	return false
}

// SetPaymentPolicy gets a reference to the given string and assigns it to the PaymentPolicy field.
func (o *StandardInfo) SetPaymentPolicy(v string) {
	o.PaymentPolicy = &v
}

// GetPersonalQrcodeMode returns the PersonalQrcodeMode field value if set, zero value otherwise.
func (o *StandardInfo) GetPersonalQrcodeMode() int32 {
	if o == nil || IsNil(o.PersonalQrcodeMode) {
		var ret int32
		return ret
	}
	return *o.PersonalQrcodeMode
}

// GetPersonalQrcodeModeOk returns a tuple with the PersonalQrcodeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardInfo) GetPersonalQrcodeModeOk() (*int32, bool) {
	if o == nil || IsNil(o.PersonalQrcodeMode) {
		return nil, false
	}
	return o.PersonalQrcodeMode, true
}

// HasPersonalQrcodeMode returns a boolean if a field has been set.
func (o *StandardInfo) HasPersonalQrcodeMode() bool {
	if o != nil && !IsNil(o.PersonalQrcodeMode) {
		return true
	}

	return false
}

// SetPersonalQrcodeMode gets a reference to the given int32 and assigns it to the PersonalQrcodeMode field.
func (o *StandardInfo) SetPersonalQrcodeMode(v int32) {
	o.PersonalQrcodeMode = &v
}

// GetStandardConditionInfoList returns the StandardConditionInfoList field value if set, zero value otherwise.
func (o *StandardInfo) GetStandardConditionInfoList() []StandardConditionInfo {
	if o == nil || IsNil(o.StandardConditionInfoList) {
		var ret []StandardConditionInfo
		return ret
	}
	return o.StandardConditionInfoList
}

// GetStandardConditionInfoListOk returns a tuple with the StandardConditionInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardInfo) GetStandardConditionInfoListOk() ([]StandardConditionInfo, bool) {
	if o == nil || IsNil(o.StandardConditionInfoList) {
		return nil, false
	}
	return o.StandardConditionInfoList, true
}

// HasStandardConditionInfoList returns a boolean if a field has been set.
func (o *StandardInfo) HasStandardConditionInfoList() bool {
	if o != nil && !IsNil(o.StandardConditionInfoList) {
		return true
	}

	return false
}

// SetStandardConditionInfoList gets a reference to the given []StandardConditionInfo and assigns it to the StandardConditionInfoList field.
func (o *StandardInfo) SetStandardConditionInfoList(v []StandardConditionInfo) {
	o.StandardConditionInfoList = v
}

// GetStandardDesc returns the StandardDesc field value if set, zero value otherwise.
func (o *StandardInfo) GetStandardDesc() string {
	if o == nil || IsNil(o.StandardDesc) {
		var ret string
		return ret
	}
	return *o.StandardDesc
}

// GetStandardDescOk returns a tuple with the StandardDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardInfo) GetStandardDescOk() (*string, bool) {
	if o == nil || IsNil(o.StandardDesc) {
		return nil, false
	}
	return o.StandardDesc, true
}

// HasStandardDesc returns a boolean if a field has been set.
func (o *StandardInfo) HasStandardDesc() bool {
	if o != nil && !IsNil(o.StandardDesc) {
		return true
	}

	return false
}

// SetStandardDesc gets a reference to the given string and assigns it to the StandardDesc field.
func (o *StandardInfo) SetStandardDesc(v string) {
	o.StandardDesc = &v
}

// GetStandardId returns the StandardId field value if set, zero value otherwise.
func (o *StandardInfo) GetStandardId() string {
	if o == nil || IsNil(o.StandardId) {
		var ret string
		return ret
	}
	return *o.StandardId
}

// GetStandardIdOk returns a tuple with the StandardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardInfo) GetStandardIdOk() (*string, bool) {
	if o == nil || IsNil(o.StandardId) {
		return nil, false
	}
	return o.StandardId, true
}

// HasStandardId returns a boolean if a field has been set.
func (o *StandardInfo) HasStandardId() bool {
	if o != nil && !IsNil(o.StandardId) {
		return true
	}

	return false
}

// SetStandardId gets a reference to the given string and assigns it to the StandardId field.
func (o *StandardInfo) SetStandardId(v string) {
	o.StandardId = &v
}

// GetStandardName returns the StandardName field value if set, zero value otherwise.
func (o *StandardInfo) GetStandardName() string {
	if o == nil || IsNil(o.StandardName) {
		var ret string
		return ret
	}
	return *o.StandardName
}

// GetStandardNameOk returns a tuple with the StandardName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardInfo) GetStandardNameOk() (*string, bool) {
	if o == nil || IsNil(o.StandardName) {
		return nil, false
	}
	return o.StandardName, true
}

// HasStandardName returns a boolean if a field has been set.
func (o *StandardInfo) HasStandardName() bool {
	if o != nil && !IsNil(o.StandardName) {
		return true
	}

	return false
}

// SetStandardName gets a reference to the given string and assigns it to the StandardName field.
func (o *StandardInfo) SetStandardName(v string) {
	o.StandardName = &v
}

func (o StandardInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandardInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetShareSourceInfo) {
		toSerialize["asset_share_source_info"] = o.AssetShareSourceInfo
	}
	if !IsNil(o.ConsumeMode) {
		toSerialize["consume_mode"] = o.ConsumeMode
	}
	if !IsNil(o.ExpenseTypeSubCategory) {
		toSerialize["expense_type_sub_category"] = o.ExpenseTypeSubCategory
	}
	if !IsNil(o.OpenRuleId) {
		toSerialize["open_rule_id"] = o.OpenRuleId
	}
	if !IsNil(o.OuterSourceId) {
		toSerialize["outer_source_id"] = o.OuterSourceId
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

type NullableStandardInfo struct {
	value *StandardInfo
	isSet bool
}

func (v NullableStandardInfo) Get() *StandardInfo {
	return v.value
}

func (v *NullableStandardInfo) Set(val *StandardInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardInfo(val *StandardInfo) *NullableStandardInfo {
	return &NullableStandardInfo{value: val, isSet: true}
}

func (v NullableStandardInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


