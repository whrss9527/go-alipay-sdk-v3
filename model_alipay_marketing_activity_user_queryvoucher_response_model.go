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

// checks if the AlipayMarketingActivityUserQueryvoucherResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityUserQueryvoucherResponseModel{}

// AlipayMarketingActivityUserQueryvoucherResponseModel struct for AlipayMarketingActivityUserQueryvoucherResponseModel
type AlipayMarketingActivityUserQueryvoucherResponseModel struct {
	ActivityBaseInfo *ActivityBaseInfo `json:"activity_base_info,omitempty"`
	// 活动 id
	ActivityId *string `json:"activity_id,omitempty"`
	// 若商家券操作过关联商户订单信息，则该字段返回商家券已关联的商户订单号。
	AssociateTradeNo *string `json:"associate_trade_no,omitempty"`
	// 券可用开始时间。格式为：yyyy-MM-dd HH:mm:ss
	AvailableBeginTime *string `json:"available_begin_time,omitempty"`
	// 券可用结束时间。格式为：yyyy-MM-dd HH:mm:ss
	AvailableEndTime *string `json:"available_end_time,omitempty"`
	// 券归属 pid
	BelongMerchantId *string `json:"belong_merchant_id,omitempty"`
	// 领券时间
	CreateTime                *string                    `json:"create_time,omitempty"`
	UserVoucherBaseInfo       *UserVoucherBaseInfo       `json:"user_voucher_base_info,omitempty"`
	VoucherCustomerGuideInfo  *VoucherCustomerGuideInfo  `json:"voucher_customer_guide_info,omitempty"`
	VoucherDeductInfo         *VoucherDeductInfo         `json:"voucher_deduct_info,omitempty"`
	VoucherDisplayInfo        *CommonVoucherDisplayInfo  `json:"voucher_display_info,omitempty"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo `json:"voucher_display_pattern_info,omitempty"`
	// 对消费者展示的券(商品)名称。
	VoucherName         *string                `json:"voucher_name,omitempty"`
	VoucherSendModeInfo *VoucherSendModeInfo   `json:"voucher_send_mode_info,omitempty"`
	VoucherSendRule     *CommonVoucherSendRule `json:"voucher_send_rule,omitempty"`
	// 券状态。
	VoucherStatus *string `json:"voucher_status,omitempty"`
	// 券类型。
	VoucherType        *string               `json:"voucher_type,omitempty"`
	VoucherUseRule     *CommonVoucherUseRule `json:"voucher_use_rule,omitempty"`
	VoucherUseRuleInfo *VoucherUseRuleInfo   `json:"voucher_use_rule_info,omitempty"`
}

// NewAlipayMarketingActivityUserQueryvoucherResponseModel instantiates a new AlipayMarketingActivityUserQueryvoucherResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityUserQueryvoucherResponseModel() *AlipayMarketingActivityUserQueryvoucherResponseModel {
	this := AlipayMarketingActivityUserQueryvoucherResponseModel{}
	return &this
}

// NewAlipayMarketingActivityUserQueryvoucherResponseModelWithDefaults instantiates a new AlipayMarketingActivityUserQueryvoucherResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityUserQueryvoucherResponseModelWithDefaults() *AlipayMarketingActivityUserQueryvoucherResponseModel {
	this := AlipayMarketingActivityUserQueryvoucherResponseModel{}
	return &this
}

// GetActivityBaseInfo returns the ActivityBaseInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetActivityBaseInfo() ActivityBaseInfo {
	if o == nil || IsNil(o.ActivityBaseInfo) {
		var ret ActivityBaseInfo
		return ret
	}
	return *o.ActivityBaseInfo
}

// GetActivityBaseInfoOk returns a tuple with the ActivityBaseInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetActivityBaseInfoOk() (*ActivityBaseInfo, bool) {
	if o == nil || IsNil(o.ActivityBaseInfo) {
		return nil, false
	}
	return o.ActivityBaseInfo, true
}

// HasActivityBaseInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) HasActivityBaseInfo() bool {
	if o != nil && !IsNil(o.ActivityBaseInfo) {
		return true
	}

	return false
}

// SetActivityBaseInfo gets a reference to the given ActivityBaseInfo and assigns it to the ActivityBaseInfo field.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) SetActivityBaseInfo(v ActivityBaseInfo) {
	o.ActivityBaseInfo = &v
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId) {
		var ret string
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetActivityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given string and assigns it to the ActivityId field.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) SetActivityId(v string) {
	o.ActivityId = &v
}

// GetAssociateTradeNo returns the AssociateTradeNo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetAssociateTradeNo() string {
	if o == nil || IsNil(o.AssociateTradeNo) {
		var ret string
		return ret
	}
	return *o.AssociateTradeNo
}

// GetAssociateTradeNoOk returns a tuple with the AssociateTradeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetAssociateTradeNoOk() (*string, bool) {
	if o == nil || IsNil(o.AssociateTradeNo) {
		return nil, false
	}
	return o.AssociateTradeNo, true
}

// HasAssociateTradeNo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) HasAssociateTradeNo() bool {
	if o != nil && !IsNil(o.AssociateTradeNo) {
		return true
	}

	return false
}

// SetAssociateTradeNo gets a reference to the given string and assigns it to the AssociateTradeNo field.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) SetAssociateTradeNo(v string) {
	o.AssociateTradeNo = &v
}

// GetAvailableBeginTime returns the AvailableBeginTime field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetAvailableBeginTime() string {
	if o == nil || IsNil(o.AvailableBeginTime) {
		var ret string
		return ret
	}
	return *o.AvailableBeginTime
}

// GetAvailableBeginTimeOk returns a tuple with the AvailableBeginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetAvailableBeginTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AvailableBeginTime) {
		return nil, false
	}
	return o.AvailableBeginTime, true
}

// HasAvailableBeginTime returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) HasAvailableBeginTime() bool {
	if o != nil && !IsNil(o.AvailableBeginTime) {
		return true
	}

	return false
}

// SetAvailableBeginTime gets a reference to the given string and assigns it to the AvailableBeginTime field.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) SetAvailableBeginTime(v string) {
	o.AvailableBeginTime = &v
}

// GetAvailableEndTime returns the AvailableEndTime field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetAvailableEndTime() string {
	if o == nil || IsNil(o.AvailableEndTime) {
		var ret string
		return ret
	}
	return *o.AvailableEndTime
}

// GetAvailableEndTimeOk returns a tuple with the AvailableEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetAvailableEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AvailableEndTime) {
		return nil, false
	}
	return o.AvailableEndTime, true
}

// HasAvailableEndTime returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) HasAvailableEndTime() bool {
	if o != nil && !IsNil(o.AvailableEndTime) {
		return true
	}

	return false
}

// SetAvailableEndTime gets a reference to the given string and assigns it to the AvailableEndTime field.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) SetAvailableEndTime(v string) {
	o.AvailableEndTime = &v
}

// GetBelongMerchantId returns the BelongMerchantId field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetBelongMerchantId() string {
	if o == nil || IsNil(o.BelongMerchantId) {
		var ret string
		return ret
	}
	return *o.BelongMerchantId
}

// GetBelongMerchantIdOk returns a tuple with the BelongMerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetBelongMerchantIdOk() (*string, bool) {
	if o == nil || IsNil(o.BelongMerchantId) {
		return nil, false
	}
	return o.BelongMerchantId, true
}

// HasBelongMerchantId returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) HasBelongMerchantId() bool {
	if o != nil && !IsNil(o.BelongMerchantId) {
		return true
	}

	return false
}

// SetBelongMerchantId gets a reference to the given string and assigns it to the BelongMerchantId field.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) SetBelongMerchantId(v string) {
	o.BelongMerchantId = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetUserVoucherBaseInfo returns the UserVoucherBaseInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetUserVoucherBaseInfo() UserVoucherBaseInfo {
	if o == nil || IsNil(o.UserVoucherBaseInfo) {
		var ret UserVoucherBaseInfo
		return ret
	}
	return *o.UserVoucherBaseInfo
}

// GetUserVoucherBaseInfoOk returns a tuple with the UserVoucherBaseInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetUserVoucherBaseInfoOk() (*UserVoucherBaseInfo, bool) {
	if o == nil || IsNil(o.UserVoucherBaseInfo) {
		return nil, false
	}
	return o.UserVoucherBaseInfo, true
}

// HasUserVoucherBaseInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) HasUserVoucherBaseInfo() bool {
	if o != nil && !IsNil(o.UserVoucherBaseInfo) {
		return true
	}

	return false
}

// SetUserVoucherBaseInfo gets a reference to the given UserVoucherBaseInfo and assigns it to the UserVoucherBaseInfo field.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) SetUserVoucherBaseInfo(v UserVoucherBaseInfo) {
	o.UserVoucherBaseInfo = &v
}

// GetVoucherCustomerGuideInfo returns the VoucherCustomerGuideInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherCustomerGuideInfo() VoucherCustomerGuideInfo {
	if o == nil || IsNil(o.VoucherCustomerGuideInfo) {
		var ret VoucherCustomerGuideInfo
		return ret
	}
	return *o.VoucherCustomerGuideInfo
}

// GetVoucherCustomerGuideInfoOk returns a tuple with the VoucherCustomerGuideInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherCustomerGuideInfoOk() (*VoucherCustomerGuideInfo, bool) {
	if o == nil || IsNil(o.VoucherCustomerGuideInfo) {
		return nil, false
	}
	return o.VoucherCustomerGuideInfo, true
}

// HasVoucherCustomerGuideInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) HasVoucherCustomerGuideInfo() bool {
	if o != nil && !IsNil(o.VoucherCustomerGuideInfo) {
		return true
	}

	return false
}

// SetVoucherCustomerGuideInfo gets a reference to the given VoucherCustomerGuideInfo and assigns it to the VoucherCustomerGuideInfo field.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) SetVoucherCustomerGuideInfo(v VoucherCustomerGuideInfo) {
	o.VoucherCustomerGuideInfo = &v
}

// GetVoucherDeductInfo returns the VoucherDeductInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherDeductInfo() VoucherDeductInfo {
	if o == nil || IsNil(o.VoucherDeductInfo) {
		var ret VoucherDeductInfo
		return ret
	}
	return *o.VoucherDeductInfo
}

// GetVoucherDeductInfoOk returns a tuple with the VoucherDeductInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherDeductInfoOk() (*VoucherDeductInfo, bool) {
	if o == nil || IsNil(o.VoucherDeductInfo) {
		return nil, false
	}
	return o.VoucherDeductInfo, true
}

// HasVoucherDeductInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) HasVoucherDeductInfo() bool {
	if o != nil && !IsNil(o.VoucherDeductInfo) {
		return true
	}

	return false
}

// SetVoucherDeductInfo gets a reference to the given VoucherDeductInfo and assigns it to the VoucherDeductInfo field.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) SetVoucherDeductInfo(v VoucherDeductInfo) {
	o.VoucherDeductInfo = &v
}

// GetVoucherDisplayInfo returns the VoucherDisplayInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherDisplayInfo() CommonVoucherDisplayInfo {
	if o == nil || IsNil(o.VoucherDisplayInfo) {
		var ret CommonVoucherDisplayInfo
		return ret
	}
	return *o.VoucherDisplayInfo
}

// GetVoucherDisplayInfoOk returns a tuple with the VoucherDisplayInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherDisplayInfoOk() (*CommonVoucherDisplayInfo, bool) {
	if o == nil || IsNil(o.VoucherDisplayInfo) {
		return nil, false
	}
	return o.VoucherDisplayInfo, true
}

// HasVoucherDisplayInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) HasVoucherDisplayInfo() bool {
	if o != nil && !IsNil(o.VoucherDisplayInfo) {
		return true
	}

	return false
}

// SetVoucherDisplayInfo gets a reference to the given CommonVoucherDisplayInfo and assigns it to the VoucherDisplayInfo field.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) SetVoucherDisplayInfo(v CommonVoucherDisplayInfo) {
	o.VoucherDisplayInfo = &v
}

// GetVoucherDisplayPatternInfo returns the VoucherDisplayPatternInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherDisplayPatternInfo() VoucherDisplayPatternInfo {
	if o == nil || IsNil(o.VoucherDisplayPatternInfo) {
		var ret VoucherDisplayPatternInfo
		return ret
	}
	return *o.VoucherDisplayPatternInfo
}

// GetVoucherDisplayPatternInfoOk returns a tuple with the VoucherDisplayPatternInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherDisplayPatternInfoOk() (*VoucherDisplayPatternInfo, bool) {
	if o == nil || IsNil(o.VoucherDisplayPatternInfo) {
		return nil, false
	}
	return o.VoucherDisplayPatternInfo, true
}

// HasVoucherDisplayPatternInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) HasVoucherDisplayPatternInfo() bool {
	if o != nil && !IsNil(o.VoucherDisplayPatternInfo) {
		return true
	}

	return false
}

// SetVoucherDisplayPatternInfo gets a reference to the given VoucherDisplayPatternInfo and assigns it to the VoucherDisplayPatternInfo field.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) SetVoucherDisplayPatternInfo(v VoucherDisplayPatternInfo) {
	o.VoucherDisplayPatternInfo = &v
}

// GetVoucherName returns the VoucherName field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherName() string {
	if o == nil || IsNil(o.VoucherName) {
		var ret string
		return ret
	}
	return *o.VoucherName
}

// GetVoucherNameOk returns a tuple with the VoucherName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherNameOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherName) {
		return nil, false
	}
	return o.VoucherName, true
}

// HasVoucherName returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) HasVoucherName() bool {
	if o != nil && !IsNil(o.VoucherName) {
		return true
	}

	return false
}

// SetVoucherName gets a reference to the given string and assigns it to the VoucherName field.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) SetVoucherName(v string) {
	o.VoucherName = &v
}

// GetVoucherSendModeInfo returns the VoucherSendModeInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherSendModeInfo() VoucherSendModeInfo {
	if o == nil || IsNil(o.VoucherSendModeInfo) {
		var ret VoucherSendModeInfo
		return ret
	}
	return *o.VoucherSendModeInfo
}

// GetVoucherSendModeInfoOk returns a tuple with the VoucherSendModeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherSendModeInfoOk() (*VoucherSendModeInfo, bool) {
	if o == nil || IsNil(o.VoucherSendModeInfo) {
		return nil, false
	}
	return o.VoucherSendModeInfo, true
}

// HasVoucherSendModeInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) HasVoucherSendModeInfo() bool {
	if o != nil && !IsNil(o.VoucherSendModeInfo) {
		return true
	}

	return false
}

// SetVoucherSendModeInfo gets a reference to the given VoucherSendModeInfo and assigns it to the VoucherSendModeInfo field.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) SetVoucherSendModeInfo(v VoucherSendModeInfo) {
	o.VoucherSendModeInfo = &v
}

// GetVoucherSendRule returns the VoucherSendRule field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherSendRule() CommonVoucherSendRule {
	if o == nil || IsNil(o.VoucherSendRule) {
		var ret CommonVoucherSendRule
		return ret
	}
	return *o.VoucherSendRule
}

// GetVoucherSendRuleOk returns a tuple with the VoucherSendRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherSendRuleOk() (*CommonVoucherSendRule, bool) {
	if o == nil || IsNil(o.VoucherSendRule) {
		return nil, false
	}
	return o.VoucherSendRule, true
}

// HasVoucherSendRule returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) HasVoucherSendRule() bool {
	if o != nil && !IsNil(o.VoucherSendRule) {
		return true
	}

	return false
}

// SetVoucherSendRule gets a reference to the given CommonVoucherSendRule and assigns it to the VoucherSendRule field.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) SetVoucherSendRule(v CommonVoucherSendRule) {
	o.VoucherSendRule = &v
}

// GetVoucherStatus returns the VoucherStatus field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherStatus() string {
	if o == nil || IsNil(o.VoucherStatus) {
		var ret string
		return ret
	}
	return *o.VoucherStatus
}

// GetVoucherStatusOk returns a tuple with the VoucherStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherStatusOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherStatus) {
		return nil, false
	}
	return o.VoucherStatus, true
}

// HasVoucherStatus returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) HasVoucherStatus() bool {
	if o != nil && !IsNil(o.VoucherStatus) {
		return true
	}

	return false
}

// SetVoucherStatus gets a reference to the given string and assigns it to the VoucherStatus field.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) SetVoucherStatus(v string) {
	o.VoucherStatus = &v
}

// GetVoucherType returns the VoucherType field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherType() string {
	if o == nil || IsNil(o.VoucherType) {
		var ret string
		return ret
	}
	return *o.VoucherType
}

// GetVoucherTypeOk returns a tuple with the VoucherType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherType) {
		return nil, false
	}
	return o.VoucherType, true
}

// HasVoucherType returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) HasVoucherType() bool {
	if o != nil && !IsNil(o.VoucherType) {
		return true
	}

	return false
}

// SetVoucherType gets a reference to the given string and assigns it to the VoucherType field.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) SetVoucherType(v string) {
	o.VoucherType = &v
}

// GetVoucherUseRule returns the VoucherUseRule field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherUseRule() CommonVoucherUseRule {
	if o == nil || IsNil(o.VoucherUseRule) {
		var ret CommonVoucherUseRule
		return ret
	}
	return *o.VoucherUseRule
}

// GetVoucherUseRuleOk returns a tuple with the VoucherUseRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherUseRuleOk() (*CommonVoucherUseRule, bool) {
	if o == nil || IsNil(o.VoucherUseRule) {
		return nil, false
	}
	return o.VoucherUseRule, true
}

// HasVoucherUseRule returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) HasVoucherUseRule() bool {
	if o != nil && !IsNil(o.VoucherUseRule) {
		return true
	}

	return false
}

// SetVoucherUseRule gets a reference to the given CommonVoucherUseRule and assigns it to the VoucherUseRule field.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) SetVoucherUseRule(v CommonVoucherUseRule) {
	o.VoucherUseRule = &v
}

// GetVoucherUseRuleInfo returns the VoucherUseRuleInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherUseRuleInfo() VoucherUseRuleInfo {
	if o == nil || IsNil(o.VoucherUseRuleInfo) {
		var ret VoucherUseRuleInfo
		return ret
	}
	return *o.VoucherUseRuleInfo
}

// GetVoucherUseRuleInfoOk returns a tuple with the VoucherUseRuleInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) GetVoucherUseRuleInfoOk() (*VoucherUseRuleInfo, bool) {
	if o == nil || IsNil(o.VoucherUseRuleInfo) {
		return nil, false
	}
	return o.VoucherUseRuleInfo, true
}

// HasVoucherUseRuleInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) HasVoucherUseRuleInfo() bool {
	if o != nil && !IsNil(o.VoucherUseRuleInfo) {
		return true
	}

	return false
}

// SetVoucherUseRuleInfo gets a reference to the given VoucherUseRuleInfo and assigns it to the VoucherUseRuleInfo field.
func (o *AlipayMarketingActivityUserQueryvoucherResponseModel) SetVoucherUseRuleInfo(v VoucherUseRuleInfo) {
	o.VoucherUseRuleInfo = &v
}

func (o AlipayMarketingActivityUserQueryvoucherResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityUserQueryvoucherResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityBaseInfo) {
		toSerialize["activity_base_info"] = o.ActivityBaseInfo
	}
	if !IsNil(o.ActivityId) {
		toSerialize["activity_id"] = o.ActivityId
	}
	if !IsNil(o.AssociateTradeNo) {
		toSerialize["associate_trade_no"] = o.AssociateTradeNo
	}
	if !IsNil(o.AvailableBeginTime) {
		toSerialize["available_begin_time"] = o.AvailableBeginTime
	}
	if !IsNil(o.AvailableEndTime) {
		toSerialize["available_end_time"] = o.AvailableEndTime
	}
	if !IsNil(o.BelongMerchantId) {
		toSerialize["belong_merchant_id"] = o.BelongMerchantId
	}
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
	}
	if !IsNil(o.UserVoucherBaseInfo) {
		toSerialize["user_voucher_base_info"] = o.UserVoucherBaseInfo
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
	if !IsNil(o.VoucherName) {
		toSerialize["voucher_name"] = o.VoucherName
	}
	if !IsNil(o.VoucherSendModeInfo) {
		toSerialize["voucher_send_mode_info"] = o.VoucherSendModeInfo
	}
	if !IsNil(o.VoucherSendRule) {
		toSerialize["voucher_send_rule"] = o.VoucherSendRule
	}
	if !IsNil(o.VoucherStatus) {
		toSerialize["voucher_status"] = o.VoucherStatus
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

type NullableAlipayMarketingActivityUserQueryvoucherResponseModel struct {
	value *AlipayMarketingActivityUserQueryvoucherResponseModel
	isSet bool
}

func (v NullableAlipayMarketingActivityUserQueryvoucherResponseModel) Get() *AlipayMarketingActivityUserQueryvoucherResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityUserQueryvoucherResponseModel) Set(val *AlipayMarketingActivityUserQueryvoucherResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityUserQueryvoucherResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityUserQueryvoucherResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityUserQueryvoucherResponseModel(val *AlipayMarketingActivityUserQueryvoucherResponseModel) *NullableAlipayMarketingActivityUserQueryvoucherResponseModel {
	return &NullableAlipayMarketingActivityUserQueryvoucherResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityUserQueryvoucherResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityUserQueryvoucherResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
