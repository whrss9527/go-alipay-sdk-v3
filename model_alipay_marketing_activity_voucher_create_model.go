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

// checks if the AlipayMarketingActivityVoucherCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityVoucherCreateModel{}

// AlipayMarketingActivityVoucherCreateModel struct for AlipayMarketingActivityVoucherCreateModel
type AlipayMarketingActivityVoucherCreateModel struct {
	ActivityBaseInfo *ActivityBaseInfo `json:"activity_base_info,omitempty"`
	// 活动名称。
	ActivityName *string `json:"activity_name,omitempty"`
	BelongMerchantInfo *PaymentVoucherBelongMerchantInfo `json:"belong_merchant_info,omitempty"`
	// 商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// 用作幂等控制。  幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。  外部接入方需保证业务单号唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 券发放结束时间。 格式为：yyyy-MM-dd HH:mm:ss 限制： 券发放结束时间 publish_end_time 与 券发放开始时间  publish_start_time 间隔必须小于等于180天
	PublishEndTime *string `json:"publish_end_time,omitempty"`
	// 券发放开始时间。 格式为：yyyy-MM-dd HH:mm:ss
	PublishStartTime *string `json:"publish_start_time,omitempty"`
	VoucherAvailableScopeInfo *VoucherAvailableScopeInfo `json:"voucher_available_scope_info,omitempty"`
	VoucherBudgetInfo *PaymentVoucherBudgetInfo `json:"voucher_budget_info,omitempty"`
	VoucherBudgetSupplyInfo *VoucherBudgetSupplyInfo `json:"voucher_budget_supply_info,omitempty"`
	VoucherCustomerGuideInfo *VoucherCustomerGuideInfo `json:"voucher_customer_guide_info,omitempty"`
	VoucherDeductInfo *VoucherDeductInfo `json:"voucher_deduct_info,omitempty"`
	VoucherDisplayInfo *PaymentVoucherDisplayInfo `json:"voucher_display_info,omitempty"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo `json:"voucher_display_pattern_info,omitempty"`
	VoucherSendModeInfo *VoucherSendModeInfo `json:"voucher_send_mode_info,omitempty"`
	VoucherSendRule *PaymentVoucherSendRule `json:"voucher_send_rule,omitempty"`
	// 券类型
	VoucherType *string `json:"voucher_type,omitempty"`
	VoucherUseRule *PaymentVoucherUseRule `json:"voucher_use_rule,omitempty"`
	VoucherUseRuleInfo *VoucherUseRuleInfo `json:"voucher_use_rule_info,omitempty"`
}

// NewAlipayMarketingActivityVoucherCreateModel instantiates a new AlipayMarketingActivityVoucherCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityVoucherCreateModel() *AlipayMarketingActivityVoucherCreateModel {
	this := AlipayMarketingActivityVoucherCreateModel{}
	return &this
}

// NewAlipayMarketingActivityVoucherCreateModelWithDefaults instantiates a new AlipayMarketingActivityVoucherCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityVoucherCreateModelWithDefaults() *AlipayMarketingActivityVoucherCreateModel {
	this := AlipayMarketingActivityVoucherCreateModel{}
	return &this
}

// GetActivityBaseInfo returns the ActivityBaseInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateModel) GetActivityBaseInfo() ActivityBaseInfo {
	if o == nil || IsNil(o.ActivityBaseInfo) {
		var ret ActivityBaseInfo
		return ret
	}
	return *o.ActivityBaseInfo
}

// GetActivityBaseInfoOk returns a tuple with the ActivityBaseInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) GetActivityBaseInfoOk() (*ActivityBaseInfo, bool) {
	if o == nil || IsNil(o.ActivityBaseInfo) {
		return nil, false
	}
	return o.ActivityBaseInfo, true
}

// HasActivityBaseInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) HasActivityBaseInfo() bool {
	if o != nil && !IsNil(o.ActivityBaseInfo) {
		return true
	}

	return false
}

// SetActivityBaseInfo gets a reference to the given ActivityBaseInfo and assigns it to the ActivityBaseInfo field.
func (o *AlipayMarketingActivityVoucherCreateModel) SetActivityBaseInfo(v ActivityBaseInfo) {
	o.ActivityBaseInfo = &v
}

// GetActivityName returns the ActivityName field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateModel) GetActivityName() string {
	if o == nil || IsNil(o.ActivityName) {
		var ret string
		return ret
	}
	return *o.ActivityName
}

// GetActivityNameOk returns a tuple with the ActivityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) GetActivityNameOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityName) {
		return nil, false
	}
	return o.ActivityName, true
}

// HasActivityName returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) HasActivityName() bool {
	if o != nil && !IsNil(o.ActivityName) {
		return true
	}

	return false
}

// SetActivityName gets a reference to the given string and assigns it to the ActivityName field.
func (o *AlipayMarketingActivityVoucherCreateModel) SetActivityName(v string) {
	o.ActivityName = &v
}

// GetBelongMerchantInfo returns the BelongMerchantInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateModel) GetBelongMerchantInfo() PaymentVoucherBelongMerchantInfo {
	if o == nil || IsNil(o.BelongMerchantInfo) {
		var ret PaymentVoucherBelongMerchantInfo
		return ret
	}
	return *o.BelongMerchantInfo
}

// GetBelongMerchantInfoOk returns a tuple with the BelongMerchantInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) GetBelongMerchantInfoOk() (*PaymentVoucherBelongMerchantInfo, bool) {
	if o == nil || IsNil(o.BelongMerchantInfo) {
		return nil, false
	}
	return o.BelongMerchantInfo, true
}

// HasBelongMerchantInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) HasBelongMerchantInfo() bool {
	if o != nil && !IsNil(o.BelongMerchantInfo) {
		return true
	}

	return false
}

// SetBelongMerchantInfo gets a reference to the given PaymentVoucherBelongMerchantInfo and assigns it to the BelongMerchantInfo field.
func (o *AlipayMarketingActivityVoucherCreateModel) SetBelongMerchantInfo(v PaymentVoucherBelongMerchantInfo) {
	o.BelongMerchantInfo = &v
}

// GetMerchantAccessMode returns the MerchantAccessMode field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateModel) GetMerchantAccessMode() string {
	if o == nil || IsNil(o.MerchantAccessMode) {
		var ret string
		return ret
	}
	return *o.MerchantAccessMode
}

// GetMerchantAccessModeOk returns a tuple with the MerchantAccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) GetMerchantAccessModeOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantAccessMode) {
		return nil, false
	}
	return o.MerchantAccessMode, true
}

// HasMerchantAccessMode returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) HasMerchantAccessMode() bool {
	if o != nil && !IsNil(o.MerchantAccessMode) {
		return true
	}

	return false
}

// SetMerchantAccessMode gets a reference to the given string and assigns it to the MerchantAccessMode field.
func (o *AlipayMarketingActivityVoucherCreateModel) SetMerchantAccessMode(v string) {
	o.MerchantAccessMode = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayMarketingActivityVoucherCreateModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetPublishEndTime returns the PublishEndTime field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateModel) GetPublishEndTime() string {
	if o == nil || IsNil(o.PublishEndTime) {
		var ret string
		return ret
	}
	return *o.PublishEndTime
}

// GetPublishEndTimeOk returns a tuple with the PublishEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) GetPublishEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PublishEndTime) {
		return nil, false
	}
	return o.PublishEndTime, true
}

// HasPublishEndTime returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) HasPublishEndTime() bool {
	if o != nil && !IsNil(o.PublishEndTime) {
		return true
	}

	return false
}

// SetPublishEndTime gets a reference to the given string and assigns it to the PublishEndTime field.
func (o *AlipayMarketingActivityVoucherCreateModel) SetPublishEndTime(v string) {
	o.PublishEndTime = &v
}

// GetPublishStartTime returns the PublishStartTime field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateModel) GetPublishStartTime() string {
	if o == nil || IsNil(o.PublishStartTime) {
		var ret string
		return ret
	}
	return *o.PublishStartTime
}

// GetPublishStartTimeOk returns a tuple with the PublishStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) GetPublishStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PublishStartTime) {
		return nil, false
	}
	return o.PublishStartTime, true
}

// HasPublishStartTime returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) HasPublishStartTime() bool {
	if o != nil && !IsNil(o.PublishStartTime) {
		return true
	}

	return false
}

// SetPublishStartTime gets a reference to the given string and assigns it to the PublishStartTime field.
func (o *AlipayMarketingActivityVoucherCreateModel) SetPublishStartTime(v string) {
	o.PublishStartTime = &v
}

// GetVoucherAvailableScopeInfo returns the VoucherAvailableScopeInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherAvailableScopeInfo() VoucherAvailableScopeInfo {
	if o == nil || IsNil(o.VoucherAvailableScopeInfo) {
		var ret VoucherAvailableScopeInfo
		return ret
	}
	return *o.VoucherAvailableScopeInfo
}

// GetVoucherAvailableScopeInfoOk returns a tuple with the VoucherAvailableScopeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherAvailableScopeInfoOk() (*VoucherAvailableScopeInfo, bool) {
	if o == nil || IsNil(o.VoucherAvailableScopeInfo) {
		return nil, false
	}
	return o.VoucherAvailableScopeInfo, true
}

// HasVoucherAvailableScopeInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) HasVoucherAvailableScopeInfo() bool {
	if o != nil && !IsNil(o.VoucherAvailableScopeInfo) {
		return true
	}

	return false
}

// SetVoucherAvailableScopeInfo gets a reference to the given VoucherAvailableScopeInfo and assigns it to the VoucherAvailableScopeInfo field.
func (o *AlipayMarketingActivityVoucherCreateModel) SetVoucherAvailableScopeInfo(v VoucherAvailableScopeInfo) {
	o.VoucherAvailableScopeInfo = &v
}

// GetVoucherBudgetInfo returns the VoucherBudgetInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherBudgetInfo() PaymentVoucherBudgetInfo {
	if o == nil || IsNil(o.VoucherBudgetInfo) {
		var ret PaymentVoucherBudgetInfo
		return ret
	}
	return *o.VoucherBudgetInfo
}

// GetVoucherBudgetInfoOk returns a tuple with the VoucherBudgetInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherBudgetInfoOk() (*PaymentVoucherBudgetInfo, bool) {
	if o == nil || IsNil(o.VoucherBudgetInfo) {
		return nil, false
	}
	return o.VoucherBudgetInfo, true
}

// HasVoucherBudgetInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) HasVoucherBudgetInfo() bool {
	if o != nil && !IsNil(o.VoucherBudgetInfo) {
		return true
	}

	return false
}

// SetVoucherBudgetInfo gets a reference to the given PaymentVoucherBudgetInfo and assigns it to the VoucherBudgetInfo field.
func (o *AlipayMarketingActivityVoucherCreateModel) SetVoucherBudgetInfo(v PaymentVoucherBudgetInfo) {
	o.VoucherBudgetInfo = &v
}

// GetVoucherBudgetSupplyInfo returns the VoucherBudgetSupplyInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherBudgetSupplyInfo() VoucherBudgetSupplyInfo {
	if o == nil || IsNil(o.VoucherBudgetSupplyInfo) {
		var ret VoucherBudgetSupplyInfo
		return ret
	}
	return *o.VoucherBudgetSupplyInfo
}

// GetVoucherBudgetSupplyInfoOk returns a tuple with the VoucherBudgetSupplyInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherBudgetSupplyInfoOk() (*VoucherBudgetSupplyInfo, bool) {
	if o == nil || IsNil(o.VoucherBudgetSupplyInfo) {
		return nil, false
	}
	return o.VoucherBudgetSupplyInfo, true
}

// HasVoucherBudgetSupplyInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) HasVoucherBudgetSupplyInfo() bool {
	if o != nil && !IsNil(o.VoucherBudgetSupplyInfo) {
		return true
	}

	return false
}

// SetVoucherBudgetSupplyInfo gets a reference to the given VoucherBudgetSupplyInfo and assigns it to the VoucherBudgetSupplyInfo field.
func (o *AlipayMarketingActivityVoucherCreateModel) SetVoucherBudgetSupplyInfo(v VoucherBudgetSupplyInfo) {
	o.VoucherBudgetSupplyInfo = &v
}

// GetVoucherCustomerGuideInfo returns the VoucherCustomerGuideInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherCustomerGuideInfo() VoucherCustomerGuideInfo {
	if o == nil || IsNil(o.VoucherCustomerGuideInfo) {
		var ret VoucherCustomerGuideInfo
		return ret
	}
	return *o.VoucherCustomerGuideInfo
}

// GetVoucherCustomerGuideInfoOk returns a tuple with the VoucherCustomerGuideInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherCustomerGuideInfoOk() (*VoucherCustomerGuideInfo, bool) {
	if o == nil || IsNil(o.VoucherCustomerGuideInfo) {
		return nil, false
	}
	return o.VoucherCustomerGuideInfo, true
}

// HasVoucherCustomerGuideInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) HasVoucherCustomerGuideInfo() bool {
	if o != nil && !IsNil(o.VoucherCustomerGuideInfo) {
		return true
	}

	return false
}

// SetVoucherCustomerGuideInfo gets a reference to the given VoucherCustomerGuideInfo and assigns it to the VoucherCustomerGuideInfo field.
func (o *AlipayMarketingActivityVoucherCreateModel) SetVoucherCustomerGuideInfo(v VoucherCustomerGuideInfo) {
	o.VoucherCustomerGuideInfo = &v
}

// GetVoucherDeductInfo returns the VoucherDeductInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherDeductInfo() VoucherDeductInfo {
	if o == nil || IsNil(o.VoucherDeductInfo) {
		var ret VoucherDeductInfo
		return ret
	}
	return *o.VoucherDeductInfo
}

// GetVoucherDeductInfoOk returns a tuple with the VoucherDeductInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherDeductInfoOk() (*VoucherDeductInfo, bool) {
	if o == nil || IsNil(o.VoucherDeductInfo) {
		return nil, false
	}
	return o.VoucherDeductInfo, true
}

// HasVoucherDeductInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) HasVoucherDeductInfo() bool {
	if o != nil && !IsNil(o.VoucherDeductInfo) {
		return true
	}

	return false
}

// SetVoucherDeductInfo gets a reference to the given VoucherDeductInfo and assigns it to the VoucherDeductInfo field.
func (o *AlipayMarketingActivityVoucherCreateModel) SetVoucherDeductInfo(v VoucherDeductInfo) {
	o.VoucherDeductInfo = &v
}

// GetVoucherDisplayInfo returns the VoucherDisplayInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherDisplayInfo() PaymentVoucherDisplayInfo {
	if o == nil || IsNil(o.VoucherDisplayInfo) {
		var ret PaymentVoucherDisplayInfo
		return ret
	}
	return *o.VoucherDisplayInfo
}

// GetVoucherDisplayInfoOk returns a tuple with the VoucherDisplayInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherDisplayInfoOk() (*PaymentVoucherDisplayInfo, bool) {
	if o == nil || IsNil(o.VoucherDisplayInfo) {
		return nil, false
	}
	return o.VoucherDisplayInfo, true
}

// HasVoucherDisplayInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) HasVoucherDisplayInfo() bool {
	if o != nil && !IsNil(o.VoucherDisplayInfo) {
		return true
	}

	return false
}

// SetVoucherDisplayInfo gets a reference to the given PaymentVoucherDisplayInfo and assigns it to the VoucherDisplayInfo field.
func (o *AlipayMarketingActivityVoucherCreateModel) SetVoucherDisplayInfo(v PaymentVoucherDisplayInfo) {
	o.VoucherDisplayInfo = &v
}

// GetVoucherDisplayPatternInfo returns the VoucherDisplayPatternInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherDisplayPatternInfo() VoucherDisplayPatternInfo {
	if o == nil || IsNil(o.VoucherDisplayPatternInfo) {
		var ret VoucherDisplayPatternInfo
		return ret
	}
	return *o.VoucherDisplayPatternInfo
}

// GetVoucherDisplayPatternInfoOk returns a tuple with the VoucherDisplayPatternInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherDisplayPatternInfoOk() (*VoucherDisplayPatternInfo, bool) {
	if o == nil || IsNil(o.VoucherDisplayPatternInfo) {
		return nil, false
	}
	return o.VoucherDisplayPatternInfo, true
}

// HasVoucherDisplayPatternInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) HasVoucherDisplayPatternInfo() bool {
	if o != nil && !IsNil(o.VoucherDisplayPatternInfo) {
		return true
	}

	return false
}

// SetVoucherDisplayPatternInfo gets a reference to the given VoucherDisplayPatternInfo and assigns it to the VoucherDisplayPatternInfo field.
func (o *AlipayMarketingActivityVoucherCreateModel) SetVoucherDisplayPatternInfo(v VoucherDisplayPatternInfo) {
	o.VoucherDisplayPatternInfo = &v
}

// GetVoucherSendModeInfo returns the VoucherSendModeInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherSendModeInfo() VoucherSendModeInfo {
	if o == nil || IsNil(o.VoucherSendModeInfo) {
		var ret VoucherSendModeInfo
		return ret
	}
	return *o.VoucherSendModeInfo
}

// GetVoucherSendModeInfoOk returns a tuple with the VoucherSendModeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherSendModeInfoOk() (*VoucherSendModeInfo, bool) {
	if o == nil || IsNil(o.VoucherSendModeInfo) {
		return nil, false
	}
	return o.VoucherSendModeInfo, true
}

// HasVoucherSendModeInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) HasVoucherSendModeInfo() bool {
	if o != nil && !IsNil(o.VoucherSendModeInfo) {
		return true
	}

	return false
}

// SetVoucherSendModeInfo gets a reference to the given VoucherSendModeInfo and assigns it to the VoucherSendModeInfo field.
func (o *AlipayMarketingActivityVoucherCreateModel) SetVoucherSendModeInfo(v VoucherSendModeInfo) {
	o.VoucherSendModeInfo = &v
}

// GetVoucherSendRule returns the VoucherSendRule field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherSendRule() PaymentVoucherSendRule {
	if o == nil || IsNil(o.VoucherSendRule) {
		var ret PaymentVoucherSendRule
		return ret
	}
	return *o.VoucherSendRule
}

// GetVoucherSendRuleOk returns a tuple with the VoucherSendRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherSendRuleOk() (*PaymentVoucherSendRule, bool) {
	if o == nil || IsNil(o.VoucherSendRule) {
		return nil, false
	}
	return o.VoucherSendRule, true
}

// HasVoucherSendRule returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) HasVoucherSendRule() bool {
	if o != nil && !IsNil(o.VoucherSendRule) {
		return true
	}

	return false
}

// SetVoucherSendRule gets a reference to the given PaymentVoucherSendRule and assigns it to the VoucherSendRule field.
func (o *AlipayMarketingActivityVoucherCreateModel) SetVoucherSendRule(v PaymentVoucherSendRule) {
	o.VoucherSendRule = &v
}

// GetVoucherType returns the VoucherType field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherType() string {
	if o == nil || IsNil(o.VoucherType) {
		var ret string
		return ret
	}
	return *o.VoucherType
}

// GetVoucherTypeOk returns a tuple with the VoucherType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherType) {
		return nil, false
	}
	return o.VoucherType, true
}

// HasVoucherType returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) HasVoucherType() bool {
	if o != nil && !IsNil(o.VoucherType) {
		return true
	}

	return false
}

// SetVoucherType gets a reference to the given string and assigns it to the VoucherType field.
func (o *AlipayMarketingActivityVoucherCreateModel) SetVoucherType(v string) {
	o.VoucherType = &v
}

// GetVoucherUseRule returns the VoucherUseRule field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherUseRule() PaymentVoucherUseRule {
	if o == nil || IsNil(o.VoucherUseRule) {
		var ret PaymentVoucherUseRule
		return ret
	}
	return *o.VoucherUseRule
}

// GetVoucherUseRuleOk returns a tuple with the VoucherUseRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherUseRuleOk() (*PaymentVoucherUseRule, bool) {
	if o == nil || IsNil(o.VoucherUseRule) {
		return nil, false
	}
	return o.VoucherUseRule, true
}

// HasVoucherUseRule returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) HasVoucherUseRule() bool {
	if o != nil && !IsNil(o.VoucherUseRule) {
		return true
	}

	return false
}

// SetVoucherUseRule gets a reference to the given PaymentVoucherUseRule and assigns it to the VoucherUseRule field.
func (o *AlipayMarketingActivityVoucherCreateModel) SetVoucherUseRule(v PaymentVoucherUseRule) {
	o.VoucherUseRule = &v
}

// GetVoucherUseRuleInfo returns the VoucherUseRuleInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherUseRuleInfo() VoucherUseRuleInfo {
	if o == nil || IsNil(o.VoucherUseRuleInfo) {
		var ret VoucherUseRuleInfo
		return ret
	}
	return *o.VoucherUseRuleInfo
}

// GetVoucherUseRuleInfoOk returns a tuple with the VoucherUseRuleInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) GetVoucherUseRuleInfoOk() (*VoucherUseRuleInfo, bool) {
	if o == nil || IsNil(o.VoucherUseRuleInfo) {
		return nil, false
	}
	return o.VoucherUseRuleInfo, true
}

// HasVoucherUseRuleInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityVoucherCreateModel) HasVoucherUseRuleInfo() bool {
	if o != nil && !IsNil(o.VoucherUseRuleInfo) {
		return true
	}

	return false
}

// SetVoucherUseRuleInfo gets a reference to the given VoucherUseRuleInfo and assigns it to the VoucherUseRuleInfo field.
func (o *AlipayMarketingActivityVoucherCreateModel) SetVoucherUseRuleInfo(v VoucherUseRuleInfo) {
	o.VoucherUseRuleInfo = &v
}

func (o AlipayMarketingActivityVoucherCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityVoucherCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityBaseInfo) {
		toSerialize["activity_base_info"] = o.ActivityBaseInfo
	}
	if !IsNil(o.ActivityName) {
		toSerialize["activity_name"] = o.ActivityName
	}
	if !IsNil(o.BelongMerchantInfo) {
		toSerialize["belong_merchant_info"] = o.BelongMerchantInfo
	}
	if !IsNil(o.MerchantAccessMode) {
		toSerialize["merchant_access_mode"] = o.MerchantAccessMode
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.PublishEndTime) {
		toSerialize["publish_end_time"] = o.PublishEndTime
	}
	if !IsNil(o.PublishStartTime) {
		toSerialize["publish_start_time"] = o.PublishStartTime
	}
	if !IsNil(o.VoucherAvailableScopeInfo) {
		toSerialize["voucher_available_scope_info"] = o.VoucherAvailableScopeInfo
	}
	if !IsNil(o.VoucherBudgetInfo) {
		toSerialize["voucher_budget_info"] = o.VoucherBudgetInfo
	}
	if !IsNil(o.VoucherBudgetSupplyInfo) {
		toSerialize["voucher_budget_supply_info"] = o.VoucherBudgetSupplyInfo
	}
	if !IsNil(o.VoucherCustomerGuideInfo) {
		toSerialize["voucher_customer_guide_info"] = o.VoucherCustomerGuideInfo
	}
	if !IsNil(o.VoucherDeductInfo) {
		toSerialize["voucher_deduct_info"] = o.VoucherDeductInfo
	}
	if !IsNil(o.VoucherDisplayInfo) {
		toSerialize["voucher_display_info"] = o.VoucherDisplayInfo
	}
	if !IsNil(o.VoucherDisplayPatternInfo) {
		toSerialize["voucher_display_pattern_info"] = o.VoucherDisplayPatternInfo
	}
	if !IsNil(o.VoucherSendModeInfo) {
		toSerialize["voucher_send_mode_info"] = o.VoucherSendModeInfo
	}
	if !IsNil(o.VoucherSendRule) {
		toSerialize["voucher_send_rule"] = o.VoucherSendRule
	}
	if !IsNil(o.VoucherType) {
		toSerialize["voucher_type"] = o.VoucherType
	}
	if !IsNil(o.VoucherUseRule) {
		toSerialize["voucher_use_rule"] = o.VoucherUseRule
	}
	if !IsNil(o.VoucherUseRuleInfo) {
		toSerialize["voucher_use_rule_info"] = o.VoucherUseRuleInfo
	}
	return toSerialize, nil
}

type NullableAlipayMarketingActivityVoucherCreateModel struct {
	value *AlipayMarketingActivityVoucherCreateModel
	isSet bool
}

func (v NullableAlipayMarketingActivityVoucherCreateModel) Get() *AlipayMarketingActivityVoucherCreateModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityVoucherCreateModel) Set(val *AlipayMarketingActivityVoucherCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityVoucherCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityVoucherCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityVoucherCreateModel(val *AlipayMarketingActivityVoucherCreateModel) *NullableAlipayMarketingActivityVoucherCreateModel {
	return &NullableAlipayMarketingActivityVoucherCreateModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityVoucherCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityVoucherCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

