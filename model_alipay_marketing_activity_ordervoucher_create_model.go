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

// checks if the AlipayMarketingActivityOrdervoucherCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityOrdervoucherCreateModel{}

// AlipayMarketingActivityOrdervoucherCreateModel struct for AlipayMarketingActivityOrdervoucherCreateModel
type AlipayMarketingActivityOrdervoucherCreateModel struct {
	ActivityBaseInfo *ActivityBaseInfo `json:"activity_base_info,omitempty"`
	// 活动名称。不对用户进行展示，仅供商家在后台管理活动使用。
	ActivityName *string `json:"activity_name,omitempty"`
	BelongMerchantInfo *BelongMerchantInfo `json:"belong_merchant_info,omitempty"`
	// 该字段废弃！后续不要使用该字段！商家券业务标签，影响商家券对C端用户的展示形式。
	BizTag *string `json:"biz_tag,omitempty"`
	// 码模式。MERCHANT_UPLOAD：商户上传自定义code，发券时系统随机选取上传的券code发放。MERCHANT_API：发奖时指定券码发奖，此模式无须提前上传券码。
	CodeMode *string `json:"code_mode,omitempty"`
	CustomerGuide *CustomerGuide `json:"customer_guide,omitempty"`
	// 商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// 外部业务单号，用作幂等控制。 幂等作用：参数不变的情况下，再次请求返回与上一次相同的结果。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 券发放结束时间。格式为：yyyy-MM-dd HH:mm:ss
	PublishEndTime *string `json:"publish_end_time,omitempty"`
	// 券发放开始时间。格式为：yyyy-MM-dd HH:mm:ss
	PublishStartTime *string `json:"publish_start_time,omitempty"`
	VoucherAvailableScopeInfo *VoucherAvailableScopeInfo `json:"voucher_available_scope_info,omitempty"`
	VoucherCustomerGuideInfo *VoucherCustomerGuideInfo `json:"voucher_customer_guide_info,omitempty"`
	VoucherDeductInfo *VoucherDeductInfo `json:"voucher_deduct_info,omitempty"`
	VoucherDisplayInfo *VoucherDisplayInfo `json:"voucher_display_info,omitempty"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo `json:"voucher_display_pattern_info,omitempty"`
	VoucherSendModeInfo *VoucherSendModeInfo `json:"voucher_send_mode_info,omitempty"`
	VoucherSendRule *VoucherSendRuleDetail `json:"voucher_send_rule,omitempty"`
	// 券类型。 兑换券仅支持一类券类型：EXCHANGE_VOUCHER: 兑换券；
	VoucherType *string `json:"voucher_type,omitempty"`
	VoucherUseRule *VoucherUseRule `json:"voucher_use_rule,omitempty"`
	VoucherUseRuleInfo *VoucherUseRuleInfo `json:"voucher_use_rule_info,omitempty"`
}

// NewAlipayMarketingActivityOrdervoucherCreateModel instantiates a new AlipayMarketingActivityOrdervoucherCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityOrdervoucherCreateModel() *AlipayMarketingActivityOrdervoucherCreateModel {
	this := AlipayMarketingActivityOrdervoucherCreateModel{}
	return &this
}

// NewAlipayMarketingActivityOrdervoucherCreateModelWithDefaults instantiates a new AlipayMarketingActivityOrdervoucherCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityOrdervoucherCreateModelWithDefaults() *AlipayMarketingActivityOrdervoucherCreateModel {
	this := AlipayMarketingActivityOrdervoucherCreateModel{}
	return &this
}

// GetActivityBaseInfo returns the ActivityBaseInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetActivityBaseInfo() ActivityBaseInfo {
	if o == nil || IsNil(o.ActivityBaseInfo) {
		var ret ActivityBaseInfo
		return ret
	}
	return *o.ActivityBaseInfo
}

// GetActivityBaseInfoOk returns a tuple with the ActivityBaseInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetActivityBaseInfoOk() (*ActivityBaseInfo, bool) {
	if o == nil || IsNil(o.ActivityBaseInfo) {
		return nil, false
	}
	return o.ActivityBaseInfo, true
}

// HasActivityBaseInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasActivityBaseInfo() bool {
	if o != nil && !IsNil(o.ActivityBaseInfo) {
		return true
	}

	return false
}

// SetActivityBaseInfo gets a reference to the given ActivityBaseInfo and assigns it to the ActivityBaseInfo field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetActivityBaseInfo(v ActivityBaseInfo) {
	o.ActivityBaseInfo = &v
}

// GetActivityName returns the ActivityName field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetActivityName() string {
	if o == nil || IsNil(o.ActivityName) {
		var ret string
		return ret
	}
	return *o.ActivityName
}

// GetActivityNameOk returns a tuple with the ActivityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetActivityNameOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityName) {
		return nil, false
	}
	return o.ActivityName, true
}

// HasActivityName returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasActivityName() bool {
	if o != nil && !IsNil(o.ActivityName) {
		return true
	}

	return false
}

// SetActivityName gets a reference to the given string and assigns it to the ActivityName field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetActivityName(v string) {
	o.ActivityName = &v
}

// GetBelongMerchantInfo returns the BelongMerchantInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetBelongMerchantInfo() BelongMerchantInfo {
	if o == nil || IsNil(o.BelongMerchantInfo) {
		var ret BelongMerchantInfo
		return ret
	}
	return *o.BelongMerchantInfo
}

// GetBelongMerchantInfoOk returns a tuple with the BelongMerchantInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetBelongMerchantInfoOk() (*BelongMerchantInfo, bool) {
	if o == nil || IsNil(o.BelongMerchantInfo) {
		return nil, false
	}
	return o.BelongMerchantInfo, true
}

// HasBelongMerchantInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasBelongMerchantInfo() bool {
	if o != nil && !IsNil(o.BelongMerchantInfo) {
		return true
	}

	return false
}

// SetBelongMerchantInfo gets a reference to the given BelongMerchantInfo and assigns it to the BelongMerchantInfo field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetBelongMerchantInfo(v BelongMerchantInfo) {
	o.BelongMerchantInfo = &v
}

// GetBizTag returns the BizTag field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetBizTag() string {
	if o == nil || IsNil(o.BizTag) {
		var ret string
		return ret
	}
	return *o.BizTag
}

// GetBizTagOk returns a tuple with the BizTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetBizTagOk() (*string, bool) {
	if o == nil || IsNil(o.BizTag) {
		return nil, false
	}
	return o.BizTag, true
}

// HasBizTag returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasBizTag() bool {
	if o != nil && !IsNil(o.BizTag) {
		return true
	}

	return false
}

// SetBizTag gets a reference to the given string and assigns it to the BizTag field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetBizTag(v string) {
	o.BizTag = &v
}

// GetCodeMode returns the CodeMode field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetCodeMode() string {
	if o == nil || IsNil(o.CodeMode) {
		var ret string
		return ret
	}
	return *o.CodeMode
}

// GetCodeModeOk returns a tuple with the CodeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetCodeModeOk() (*string, bool) {
	if o == nil || IsNil(o.CodeMode) {
		return nil, false
	}
	return o.CodeMode, true
}

// HasCodeMode returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasCodeMode() bool {
	if o != nil && !IsNil(o.CodeMode) {
		return true
	}

	return false
}

// SetCodeMode gets a reference to the given string and assigns it to the CodeMode field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetCodeMode(v string) {
	o.CodeMode = &v
}

// GetCustomerGuide returns the CustomerGuide field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetCustomerGuide() CustomerGuide {
	if o == nil || IsNil(o.CustomerGuide) {
		var ret CustomerGuide
		return ret
	}
	return *o.CustomerGuide
}

// GetCustomerGuideOk returns a tuple with the CustomerGuide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetCustomerGuideOk() (*CustomerGuide, bool) {
	if o == nil || IsNil(o.CustomerGuide) {
		return nil, false
	}
	return o.CustomerGuide, true
}

// HasCustomerGuide returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasCustomerGuide() bool {
	if o != nil && !IsNil(o.CustomerGuide) {
		return true
	}

	return false
}

// SetCustomerGuide gets a reference to the given CustomerGuide and assigns it to the CustomerGuide field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetCustomerGuide(v CustomerGuide) {
	o.CustomerGuide = &v
}

// GetMerchantAccessMode returns the MerchantAccessMode field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetMerchantAccessMode() string {
	if o == nil || IsNil(o.MerchantAccessMode) {
		var ret string
		return ret
	}
	return *o.MerchantAccessMode
}

// GetMerchantAccessModeOk returns a tuple with the MerchantAccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetMerchantAccessModeOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantAccessMode) {
		return nil, false
	}
	return o.MerchantAccessMode, true
}

// HasMerchantAccessMode returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasMerchantAccessMode() bool {
	if o != nil && !IsNil(o.MerchantAccessMode) {
		return true
	}

	return false
}

// SetMerchantAccessMode gets a reference to the given string and assigns it to the MerchantAccessMode field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetMerchantAccessMode(v string) {
	o.MerchantAccessMode = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetPublishEndTime returns the PublishEndTime field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetPublishEndTime() string {
	if o == nil || IsNil(o.PublishEndTime) {
		var ret string
		return ret
	}
	return *o.PublishEndTime
}

// GetPublishEndTimeOk returns a tuple with the PublishEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetPublishEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PublishEndTime) {
		return nil, false
	}
	return o.PublishEndTime, true
}

// HasPublishEndTime returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasPublishEndTime() bool {
	if o != nil && !IsNil(o.PublishEndTime) {
		return true
	}

	return false
}

// SetPublishEndTime gets a reference to the given string and assigns it to the PublishEndTime field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetPublishEndTime(v string) {
	o.PublishEndTime = &v
}

// GetPublishStartTime returns the PublishStartTime field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetPublishStartTime() string {
	if o == nil || IsNil(o.PublishStartTime) {
		var ret string
		return ret
	}
	return *o.PublishStartTime
}

// GetPublishStartTimeOk returns a tuple with the PublishStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetPublishStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PublishStartTime) {
		return nil, false
	}
	return o.PublishStartTime, true
}

// HasPublishStartTime returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasPublishStartTime() bool {
	if o != nil && !IsNil(o.PublishStartTime) {
		return true
	}

	return false
}

// SetPublishStartTime gets a reference to the given string and assigns it to the PublishStartTime field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetPublishStartTime(v string) {
	o.PublishStartTime = &v
}

// GetVoucherAvailableScopeInfo returns the VoucherAvailableScopeInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherAvailableScopeInfo() VoucherAvailableScopeInfo {
	if o == nil || IsNil(o.VoucherAvailableScopeInfo) {
		var ret VoucherAvailableScopeInfo
		return ret
	}
	return *o.VoucherAvailableScopeInfo
}

// GetVoucherAvailableScopeInfoOk returns a tuple with the VoucherAvailableScopeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherAvailableScopeInfoOk() (*VoucherAvailableScopeInfo, bool) {
	if o == nil || IsNil(o.VoucherAvailableScopeInfo) {
		return nil, false
	}
	return o.VoucherAvailableScopeInfo, true
}

// HasVoucherAvailableScopeInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasVoucherAvailableScopeInfo() bool {
	if o != nil && !IsNil(o.VoucherAvailableScopeInfo) {
		return true
	}

	return false
}

// SetVoucherAvailableScopeInfo gets a reference to the given VoucherAvailableScopeInfo and assigns it to the VoucherAvailableScopeInfo field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetVoucherAvailableScopeInfo(v VoucherAvailableScopeInfo) {
	o.VoucherAvailableScopeInfo = &v
}

// GetVoucherCustomerGuideInfo returns the VoucherCustomerGuideInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherCustomerGuideInfo() VoucherCustomerGuideInfo {
	if o == nil || IsNil(o.VoucherCustomerGuideInfo) {
		var ret VoucherCustomerGuideInfo
		return ret
	}
	return *o.VoucherCustomerGuideInfo
}

// GetVoucherCustomerGuideInfoOk returns a tuple with the VoucherCustomerGuideInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherCustomerGuideInfoOk() (*VoucherCustomerGuideInfo, bool) {
	if o == nil || IsNil(o.VoucherCustomerGuideInfo) {
		return nil, false
	}
	return o.VoucherCustomerGuideInfo, true
}

// HasVoucherCustomerGuideInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasVoucherCustomerGuideInfo() bool {
	if o != nil && !IsNil(o.VoucherCustomerGuideInfo) {
		return true
	}

	return false
}

// SetVoucherCustomerGuideInfo gets a reference to the given VoucherCustomerGuideInfo and assigns it to the VoucherCustomerGuideInfo field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetVoucherCustomerGuideInfo(v VoucherCustomerGuideInfo) {
	o.VoucherCustomerGuideInfo = &v
}

// GetVoucherDeductInfo returns the VoucherDeductInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherDeductInfo() VoucherDeductInfo {
	if o == nil || IsNil(o.VoucherDeductInfo) {
		var ret VoucherDeductInfo
		return ret
	}
	return *o.VoucherDeductInfo
}

// GetVoucherDeductInfoOk returns a tuple with the VoucherDeductInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherDeductInfoOk() (*VoucherDeductInfo, bool) {
	if o == nil || IsNil(o.VoucherDeductInfo) {
		return nil, false
	}
	return o.VoucherDeductInfo, true
}

// HasVoucherDeductInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasVoucherDeductInfo() bool {
	if o != nil && !IsNil(o.VoucherDeductInfo) {
		return true
	}

	return false
}

// SetVoucherDeductInfo gets a reference to the given VoucherDeductInfo and assigns it to the VoucherDeductInfo field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetVoucherDeductInfo(v VoucherDeductInfo) {
	o.VoucherDeductInfo = &v
}

// GetVoucherDisplayInfo returns the VoucherDisplayInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherDisplayInfo() VoucherDisplayInfo {
	if o == nil || IsNil(o.VoucherDisplayInfo) {
		var ret VoucherDisplayInfo
		return ret
	}
	return *o.VoucherDisplayInfo
}

// GetVoucherDisplayInfoOk returns a tuple with the VoucherDisplayInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherDisplayInfoOk() (*VoucherDisplayInfo, bool) {
	if o == nil || IsNil(o.VoucherDisplayInfo) {
		return nil, false
	}
	return o.VoucherDisplayInfo, true
}

// HasVoucherDisplayInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasVoucherDisplayInfo() bool {
	if o != nil && !IsNil(o.VoucherDisplayInfo) {
		return true
	}

	return false
}

// SetVoucherDisplayInfo gets a reference to the given VoucherDisplayInfo and assigns it to the VoucherDisplayInfo field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetVoucherDisplayInfo(v VoucherDisplayInfo) {
	o.VoucherDisplayInfo = &v
}

// GetVoucherDisplayPatternInfo returns the VoucherDisplayPatternInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherDisplayPatternInfo() VoucherDisplayPatternInfo {
	if o == nil || IsNil(o.VoucherDisplayPatternInfo) {
		var ret VoucherDisplayPatternInfo
		return ret
	}
	return *o.VoucherDisplayPatternInfo
}

// GetVoucherDisplayPatternInfoOk returns a tuple with the VoucherDisplayPatternInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherDisplayPatternInfoOk() (*VoucherDisplayPatternInfo, bool) {
	if o == nil || IsNil(o.VoucherDisplayPatternInfo) {
		return nil, false
	}
	return o.VoucherDisplayPatternInfo, true
}

// HasVoucherDisplayPatternInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasVoucherDisplayPatternInfo() bool {
	if o != nil && !IsNil(o.VoucherDisplayPatternInfo) {
		return true
	}

	return false
}

// SetVoucherDisplayPatternInfo gets a reference to the given VoucherDisplayPatternInfo and assigns it to the VoucherDisplayPatternInfo field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetVoucherDisplayPatternInfo(v VoucherDisplayPatternInfo) {
	o.VoucherDisplayPatternInfo = &v
}

// GetVoucherSendModeInfo returns the VoucherSendModeInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherSendModeInfo() VoucherSendModeInfo {
	if o == nil || IsNil(o.VoucherSendModeInfo) {
		var ret VoucherSendModeInfo
		return ret
	}
	return *o.VoucherSendModeInfo
}

// GetVoucherSendModeInfoOk returns a tuple with the VoucherSendModeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherSendModeInfoOk() (*VoucherSendModeInfo, bool) {
	if o == nil || IsNil(o.VoucherSendModeInfo) {
		return nil, false
	}
	return o.VoucherSendModeInfo, true
}

// HasVoucherSendModeInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasVoucherSendModeInfo() bool {
	if o != nil && !IsNil(o.VoucherSendModeInfo) {
		return true
	}

	return false
}

// SetVoucherSendModeInfo gets a reference to the given VoucherSendModeInfo and assigns it to the VoucherSendModeInfo field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetVoucherSendModeInfo(v VoucherSendModeInfo) {
	o.VoucherSendModeInfo = &v
}

// GetVoucherSendRule returns the VoucherSendRule field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherSendRule() VoucherSendRuleDetail {
	if o == nil || IsNil(o.VoucherSendRule) {
		var ret VoucherSendRuleDetail
		return ret
	}
	return *o.VoucherSendRule
}

// GetVoucherSendRuleOk returns a tuple with the VoucherSendRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherSendRuleOk() (*VoucherSendRuleDetail, bool) {
	if o == nil || IsNil(o.VoucherSendRule) {
		return nil, false
	}
	return o.VoucherSendRule, true
}

// HasVoucherSendRule returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasVoucherSendRule() bool {
	if o != nil && !IsNil(o.VoucherSendRule) {
		return true
	}

	return false
}

// SetVoucherSendRule gets a reference to the given VoucherSendRuleDetail and assigns it to the VoucherSendRule field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetVoucherSendRule(v VoucherSendRuleDetail) {
	o.VoucherSendRule = &v
}

// GetVoucherType returns the VoucherType field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherType() string {
	if o == nil || IsNil(o.VoucherType) {
		var ret string
		return ret
	}
	return *o.VoucherType
}

// GetVoucherTypeOk returns a tuple with the VoucherType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherType) {
		return nil, false
	}
	return o.VoucherType, true
}

// HasVoucherType returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasVoucherType() bool {
	if o != nil && !IsNil(o.VoucherType) {
		return true
	}

	return false
}

// SetVoucherType gets a reference to the given string and assigns it to the VoucherType field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetVoucherType(v string) {
	o.VoucherType = &v
}

// GetVoucherUseRule returns the VoucherUseRule field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherUseRule() VoucherUseRule {
	if o == nil || IsNil(o.VoucherUseRule) {
		var ret VoucherUseRule
		return ret
	}
	return *o.VoucherUseRule
}

// GetVoucherUseRuleOk returns a tuple with the VoucherUseRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherUseRuleOk() (*VoucherUseRule, bool) {
	if o == nil || IsNil(o.VoucherUseRule) {
		return nil, false
	}
	return o.VoucherUseRule, true
}

// HasVoucherUseRule returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasVoucherUseRule() bool {
	if o != nil && !IsNil(o.VoucherUseRule) {
		return true
	}

	return false
}

// SetVoucherUseRule gets a reference to the given VoucherUseRule and assigns it to the VoucherUseRule field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetVoucherUseRule(v VoucherUseRule) {
	o.VoucherUseRule = &v
}

// GetVoucherUseRuleInfo returns the VoucherUseRuleInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherUseRuleInfo() VoucherUseRuleInfo {
	if o == nil || IsNil(o.VoucherUseRuleInfo) {
		var ret VoucherUseRuleInfo
		return ret
	}
	return *o.VoucherUseRuleInfo
}

// GetVoucherUseRuleInfoOk returns a tuple with the VoucherUseRuleInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) GetVoucherUseRuleInfoOk() (*VoucherUseRuleInfo, bool) {
	if o == nil || IsNil(o.VoucherUseRuleInfo) {
		return nil, false
	}
	return o.VoucherUseRuleInfo, true
}

// HasVoucherUseRuleInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) HasVoucherUseRuleInfo() bool {
	if o != nil && !IsNil(o.VoucherUseRuleInfo) {
		return true
	}

	return false
}

// SetVoucherUseRuleInfo gets a reference to the given VoucherUseRuleInfo and assigns it to the VoucherUseRuleInfo field.
func (o *AlipayMarketingActivityOrdervoucherCreateModel) SetVoucherUseRuleInfo(v VoucherUseRuleInfo) {
	o.VoucherUseRuleInfo = &v
}

func (o AlipayMarketingActivityOrdervoucherCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityOrdervoucherCreateModel) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.BizTag) {
		toSerialize["biz_tag"] = o.BizTag
	}
	if !IsNil(o.CodeMode) {
		toSerialize["code_mode"] = o.CodeMode
	}
	if !IsNil(o.CustomerGuide) {
		toSerialize["customer_guide"] = o.CustomerGuide
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

type NullableAlipayMarketingActivityOrdervoucherCreateModel struct {
	value *AlipayMarketingActivityOrdervoucherCreateModel
	isSet bool
}

func (v NullableAlipayMarketingActivityOrdervoucherCreateModel) Get() *AlipayMarketingActivityOrdervoucherCreateModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityOrdervoucherCreateModel) Set(val *AlipayMarketingActivityOrdervoucherCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityOrdervoucherCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityOrdervoucherCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityOrdervoucherCreateModel(val *AlipayMarketingActivityOrdervoucherCreateModel) *NullableAlipayMarketingActivityOrdervoucherCreateModel {
	return &NullableAlipayMarketingActivityOrdervoucherCreateModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityOrdervoucherCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityOrdervoucherCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


