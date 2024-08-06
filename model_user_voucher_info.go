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

// checks if the UserVoucherInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserVoucherInfo{}

// UserVoucherInfo struct for UserVoucherInfo
type UserVoucherInfo struct {
	ActivityBaseInfo *ActivityBaseInfo `json:"activity_base_info,omitempty"`
	// 活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// 券可用开始时间。
	AvailableBeginTime *string `json:"available_begin_time,omitempty"`
	// 券可用结束时间。
	AvailableEndTime *string `json:"available_end_time,omitempty"`
	// 归属商户PID
	BelongMerchantId *string `json:"belong_merchant_id,omitempty"`
	// 用户领券时间。
	CreateTime                *string                       `json:"create_time,omitempty"`
	UserVoucherBaseInfo       *UserVoucherBaseInfo          `json:"user_voucher_base_info,omitempty"`
	VoucherAvailableScopeInfo *VoucherAvailableScopeInfo    `json:"voucher_available_scope_info,omitempty"`
	VoucherCustomerGuideInfo  *VoucherCustomerGuideInfo     `json:"voucher_customer_guide_info,omitempty"`
	VoucherDeductInfo         *VoucherDeductInfo            `json:"voucher_deduct_info,omitempty"`
	VoucherDisplayLiteInfo    *CommonVoucherDisplayLiteInfo `json:"voucher_display_lite_info,omitempty"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo    `json:"voucher_display_pattern_info,omitempty"`
	// 用户券 id。支付宝为用户优惠券唯一分配的 id。
	VoucherId *string `json:"voucher_id,omitempty"`
	// 对消费者展示的券(商品)名称。
	VoucherName         *string              `json:"voucher_name,omitempty"`
	VoucherSendModeInfo *VoucherSendModeInfo `json:"voucher_send_mode_info,omitempty"`
	// 券状态。
	VoucherStatus *string `json:"voucher_status,omitempty"`
	// 券类型
	VoucherType            *string                       `json:"voucher_type,omitempty"`
	VoucherUseRuleInfo     *VoucherUseRuleInfo           `json:"voucher_use_rule_info,omitempty"`
	VoucherUseRuleLiteInfo *CommonVoucherUseRuleLiteInfo `json:"voucher_use_rule_lite_info,omitempty"`
}

// NewUserVoucherInfo instantiates a new UserVoucherInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserVoucherInfo() *UserVoucherInfo {
	this := UserVoucherInfo{}
	return &this
}

// NewUserVoucherInfoWithDefaults instantiates a new UserVoucherInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserVoucherInfoWithDefaults() *UserVoucherInfo {
	this := UserVoucherInfo{}
	return &this
}

// GetActivityBaseInfo returns the ActivityBaseInfo field value if set, zero value otherwise.
func (o *UserVoucherInfo) GetActivityBaseInfo() ActivityBaseInfo {
	if o == nil || IsNil(o.ActivityBaseInfo) {
		var ret ActivityBaseInfo
		return ret
	}
	return *o.ActivityBaseInfo
}

// GetActivityBaseInfoOk returns a tuple with the ActivityBaseInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVoucherInfo) GetActivityBaseInfoOk() (*ActivityBaseInfo, bool) {
	if o == nil || IsNil(o.ActivityBaseInfo) {
		return nil, false
	}
	return o.ActivityBaseInfo, true
}

// HasActivityBaseInfo returns a boolean if a field has been set.
func (o *UserVoucherInfo) HasActivityBaseInfo() bool {
	if o != nil && !IsNil(o.ActivityBaseInfo) {
		return true
	}

	return false
}

// SetActivityBaseInfo gets a reference to the given ActivityBaseInfo and assigns it to the ActivityBaseInfo field.
func (o *UserVoucherInfo) SetActivityBaseInfo(v ActivityBaseInfo) {
	o.ActivityBaseInfo = &v
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *UserVoucherInfo) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId) {
		var ret string
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVoucherInfo) GetActivityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *UserVoucherInfo) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given string and assigns it to the ActivityId field.
func (o *UserVoucherInfo) SetActivityId(v string) {
	o.ActivityId = &v
}

// GetAvailableBeginTime returns the AvailableBeginTime field value if set, zero value otherwise.
func (o *UserVoucherInfo) GetAvailableBeginTime() string {
	if o == nil || IsNil(o.AvailableBeginTime) {
		var ret string
		return ret
	}
	return *o.AvailableBeginTime
}

// GetAvailableBeginTimeOk returns a tuple with the AvailableBeginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVoucherInfo) GetAvailableBeginTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AvailableBeginTime) {
		return nil, false
	}
	return o.AvailableBeginTime, true
}

// HasAvailableBeginTime returns a boolean if a field has been set.
func (o *UserVoucherInfo) HasAvailableBeginTime() bool {
	if o != nil && !IsNil(o.AvailableBeginTime) {
		return true
	}

	return false
}

// SetAvailableBeginTime gets a reference to the given string and assigns it to the AvailableBeginTime field.
func (o *UserVoucherInfo) SetAvailableBeginTime(v string) {
	o.AvailableBeginTime = &v
}

// GetAvailableEndTime returns the AvailableEndTime field value if set, zero value otherwise.
func (o *UserVoucherInfo) GetAvailableEndTime() string {
	if o == nil || IsNil(o.AvailableEndTime) {
		var ret string
		return ret
	}
	return *o.AvailableEndTime
}

// GetAvailableEndTimeOk returns a tuple with the AvailableEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVoucherInfo) GetAvailableEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AvailableEndTime) {
		return nil, false
	}
	return o.AvailableEndTime, true
}

// HasAvailableEndTime returns a boolean if a field has been set.
func (o *UserVoucherInfo) HasAvailableEndTime() bool {
	if o != nil && !IsNil(o.AvailableEndTime) {
		return true
	}

	return false
}

// SetAvailableEndTime gets a reference to the given string and assigns it to the AvailableEndTime field.
func (o *UserVoucherInfo) SetAvailableEndTime(v string) {
	o.AvailableEndTime = &v
}

// GetBelongMerchantId returns the BelongMerchantId field value if set, zero value otherwise.
func (o *UserVoucherInfo) GetBelongMerchantId() string {
	if o == nil || IsNil(o.BelongMerchantId) {
		var ret string
		return ret
	}
	return *o.BelongMerchantId
}

// GetBelongMerchantIdOk returns a tuple with the BelongMerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVoucherInfo) GetBelongMerchantIdOk() (*string, bool) {
	if o == nil || IsNil(o.BelongMerchantId) {
		return nil, false
	}
	return o.BelongMerchantId, true
}

// HasBelongMerchantId returns a boolean if a field has been set.
func (o *UserVoucherInfo) HasBelongMerchantId() bool {
	if o != nil && !IsNil(o.BelongMerchantId) {
		return true
	}

	return false
}

// SetBelongMerchantId gets a reference to the given string and assigns it to the BelongMerchantId field.
func (o *UserVoucherInfo) SetBelongMerchantId(v string) {
	o.BelongMerchantId = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UserVoucherInfo) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVoucherInfo) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UserVoucherInfo) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *UserVoucherInfo) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetUserVoucherBaseInfo returns the UserVoucherBaseInfo field value if set, zero value otherwise.
func (o *UserVoucherInfo) GetUserVoucherBaseInfo() UserVoucherBaseInfo {
	if o == nil || IsNil(o.UserVoucherBaseInfo) {
		var ret UserVoucherBaseInfo
		return ret
	}
	return *o.UserVoucherBaseInfo
}

// GetUserVoucherBaseInfoOk returns a tuple with the UserVoucherBaseInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVoucherInfo) GetUserVoucherBaseInfoOk() (*UserVoucherBaseInfo, bool) {
	if o == nil || IsNil(o.UserVoucherBaseInfo) {
		return nil, false
	}
	return o.UserVoucherBaseInfo, true
}

// HasUserVoucherBaseInfo returns a boolean if a field has been set.
func (o *UserVoucherInfo) HasUserVoucherBaseInfo() bool {
	if o != nil && !IsNil(o.UserVoucherBaseInfo) {
		return true
	}

	return false
}

// SetUserVoucherBaseInfo gets a reference to the given UserVoucherBaseInfo and assigns it to the UserVoucherBaseInfo field.
func (o *UserVoucherInfo) SetUserVoucherBaseInfo(v UserVoucherBaseInfo) {
	o.UserVoucherBaseInfo = &v
}

// GetVoucherAvailableScopeInfo returns the VoucherAvailableScopeInfo field value if set, zero value otherwise.
func (o *UserVoucherInfo) GetVoucherAvailableScopeInfo() VoucherAvailableScopeInfo {
	if o == nil || IsNil(o.VoucherAvailableScopeInfo) {
		var ret VoucherAvailableScopeInfo
		return ret
	}
	return *o.VoucherAvailableScopeInfo
}

// GetVoucherAvailableScopeInfoOk returns a tuple with the VoucherAvailableScopeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVoucherInfo) GetVoucherAvailableScopeInfoOk() (*VoucherAvailableScopeInfo, bool) {
	if o == nil || IsNil(o.VoucherAvailableScopeInfo) {
		return nil, false
	}
	return o.VoucherAvailableScopeInfo, true
}

// HasVoucherAvailableScopeInfo returns a boolean if a field has been set.
func (o *UserVoucherInfo) HasVoucherAvailableScopeInfo() bool {
	if o != nil && !IsNil(o.VoucherAvailableScopeInfo) {
		return true
	}

	return false
}

// SetVoucherAvailableScopeInfo gets a reference to the given VoucherAvailableScopeInfo and assigns it to the VoucherAvailableScopeInfo field.
func (o *UserVoucherInfo) SetVoucherAvailableScopeInfo(v VoucherAvailableScopeInfo) {
	o.VoucherAvailableScopeInfo = &v
}

// GetVoucherCustomerGuideInfo returns the VoucherCustomerGuideInfo field value if set, zero value otherwise.
func (o *UserVoucherInfo) GetVoucherCustomerGuideInfo() VoucherCustomerGuideInfo {
	if o == nil || IsNil(o.VoucherCustomerGuideInfo) {
		var ret VoucherCustomerGuideInfo
		return ret
	}
	return *o.VoucherCustomerGuideInfo
}

// GetVoucherCustomerGuideInfoOk returns a tuple with the VoucherCustomerGuideInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVoucherInfo) GetVoucherCustomerGuideInfoOk() (*VoucherCustomerGuideInfo, bool) {
	if o == nil || IsNil(o.VoucherCustomerGuideInfo) {
		return nil, false
	}
	return o.VoucherCustomerGuideInfo, true
}

// HasVoucherCustomerGuideInfo returns a boolean if a field has been set.
func (o *UserVoucherInfo) HasVoucherCustomerGuideInfo() bool {
	if o != nil && !IsNil(o.VoucherCustomerGuideInfo) {
		return true
	}

	return false
}

// SetVoucherCustomerGuideInfo gets a reference to the given VoucherCustomerGuideInfo and assigns it to the VoucherCustomerGuideInfo field.
func (o *UserVoucherInfo) SetVoucherCustomerGuideInfo(v VoucherCustomerGuideInfo) {
	o.VoucherCustomerGuideInfo = &v
}

// GetVoucherDeductInfo returns the VoucherDeductInfo field value if set, zero value otherwise.
func (o *UserVoucherInfo) GetVoucherDeductInfo() VoucherDeductInfo {
	if o == nil || IsNil(o.VoucherDeductInfo) {
		var ret VoucherDeductInfo
		return ret
	}
	return *o.VoucherDeductInfo
}

// GetVoucherDeductInfoOk returns a tuple with the VoucherDeductInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVoucherInfo) GetVoucherDeductInfoOk() (*VoucherDeductInfo, bool) {
	if o == nil || IsNil(o.VoucherDeductInfo) {
		return nil, false
	}
	return o.VoucherDeductInfo, true
}

// HasVoucherDeductInfo returns a boolean if a field has been set.
func (o *UserVoucherInfo) HasVoucherDeductInfo() bool {
	if o != nil && !IsNil(o.VoucherDeductInfo) {
		return true
	}

	return false
}

// SetVoucherDeductInfo gets a reference to the given VoucherDeductInfo and assigns it to the VoucherDeductInfo field.
func (o *UserVoucherInfo) SetVoucherDeductInfo(v VoucherDeductInfo) {
	o.VoucherDeductInfo = &v
}

// GetVoucherDisplayLiteInfo returns the VoucherDisplayLiteInfo field value if set, zero value otherwise.
func (o *UserVoucherInfo) GetVoucherDisplayLiteInfo() CommonVoucherDisplayLiteInfo {
	if o == nil || IsNil(o.VoucherDisplayLiteInfo) {
		var ret CommonVoucherDisplayLiteInfo
		return ret
	}
	return *o.VoucherDisplayLiteInfo
}

// GetVoucherDisplayLiteInfoOk returns a tuple with the VoucherDisplayLiteInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVoucherInfo) GetVoucherDisplayLiteInfoOk() (*CommonVoucherDisplayLiteInfo, bool) {
	if o == nil || IsNil(o.VoucherDisplayLiteInfo) {
		return nil, false
	}
	return o.VoucherDisplayLiteInfo, true
}

// HasVoucherDisplayLiteInfo returns a boolean if a field has been set.
func (o *UserVoucherInfo) HasVoucherDisplayLiteInfo() bool {
	if o != nil && !IsNil(o.VoucherDisplayLiteInfo) {
		return true
	}

	return false
}

// SetVoucherDisplayLiteInfo gets a reference to the given CommonVoucherDisplayLiteInfo and assigns it to the VoucherDisplayLiteInfo field.
func (o *UserVoucherInfo) SetVoucherDisplayLiteInfo(v CommonVoucherDisplayLiteInfo) {
	o.VoucherDisplayLiteInfo = &v
}

// GetVoucherDisplayPatternInfo returns the VoucherDisplayPatternInfo field value if set, zero value otherwise.
func (o *UserVoucherInfo) GetVoucherDisplayPatternInfo() VoucherDisplayPatternInfo {
	if o == nil || IsNil(o.VoucherDisplayPatternInfo) {
		var ret VoucherDisplayPatternInfo
		return ret
	}
	return *o.VoucherDisplayPatternInfo
}

// GetVoucherDisplayPatternInfoOk returns a tuple with the VoucherDisplayPatternInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVoucherInfo) GetVoucherDisplayPatternInfoOk() (*VoucherDisplayPatternInfo, bool) {
	if o == nil || IsNil(o.VoucherDisplayPatternInfo) {
		return nil, false
	}
	return o.VoucherDisplayPatternInfo, true
}

// HasVoucherDisplayPatternInfo returns a boolean if a field has been set.
func (o *UserVoucherInfo) HasVoucherDisplayPatternInfo() bool {
	if o != nil && !IsNil(o.VoucherDisplayPatternInfo) {
		return true
	}

	return false
}

// SetVoucherDisplayPatternInfo gets a reference to the given VoucherDisplayPatternInfo and assigns it to the VoucherDisplayPatternInfo field.
func (o *UserVoucherInfo) SetVoucherDisplayPatternInfo(v VoucherDisplayPatternInfo) {
	o.VoucherDisplayPatternInfo = &v
}

// GetVoucherId returns the VoucherId field value if set, zero value otherwise.
func (o *UserVoucherInfo) GetVoucherId() string {
	if o == nil || IsNil(o.VoucherId) {
		var ret string
		return ret
	}
	return *o.VoucherId
}

// GetVoucherIdOk returns a tuple with the VoucherId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVoucherInfo) GetVoucherIdOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherId) {
		return nil, false
	}
	return o.VoucherId, true
}

// HasVoucherId returns a boolean if a field has been set.
func (o *UserVoucherInfo) HasVoucherId() bool {
	if o != nil && !IsNil(o.VoucherId) {
		return true
	}

	return false
}

// SetVoucherId gets a reference to the given string and assigns it to the VoucherId field.
func (o *UserVoucherInfo) SetVoucherId(v string) {
	o.VoucherId = &v
}

// GetVoucherName returns the VoucherName field value if set, zero value otherwise.
func (o *UserVoucherInfo) GetVoucherName() string {
	if o == nil || IsNil(o.VoucherName) {
		var ret string
		return ret
	}
	return *o.VoucherName
}

// GetVoucherNameOk returns a tuple with the VoucherName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVoucherInfo) GetVoucherNameOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherName) {
		return nil, false
	}
	return o.VoucherName, true
}

// HasVoucherName returns a boolean if a field has been set.
func (o *UserVoucherInfo) HasVoucherName() bool {
	if o != nil && !IsNil(o.VoucherName) {
		return true
	}

	return false
}

// SetVoucherName gets a reference to the given string and assigns it to the VoucherName field.
func (o *UserVoucherInfo) SetVoucherName(v string) {
	o.VoucherName = &v
}

// GetVoucherSendModeInfo returns the VoucherSendModeInfo field value if set, zero value otherwise.
func (o *UserVoucherInfo) GetVoucherSendModeInfo() VoucherSendModeInfo {
	if o == nil || IsNil(o.VoucherSendModeInfo) {
		var ret VoucherSendModeInfo
		return ret
	}
	return *o.VoucherSendModeInfo
}

// GetVoucherSendModeInfoOk returns a tuple with the VoucherSendModeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVoucherInfo) GetVoucherSendModeInfoOk() (*VoucherSendModeInfo, bool) {
	if o == nil || IsNil(o.VoucherSendModeInfo) {
		return nil, false
	}
	return o.VoucherSendModeInfo, true
}

// HasVoucherSendModeInfo returns a boolean if a field has been set.
func (o *UserVoucherInfo) HasVoucherSendModeInfo() bool {
	if o != nil && !IsNil(o.VoucherSendModeInfo) {
		return true
	}

	return false
}

// SetVoucherSendModeInfo gets a reference to the given VoucherSendModeInfo and assigns it to the VoucherSendModeInfo field.
func (o *UserVoucherInfo) SetVoucherSendModeInfo(v VoucherSendModeInfo) {
	o.VoucherSendModeInfo = &v
}

// GetVoucherStatus returns the VoucherStatus field value if set, zero value otherwise.
func (o *UserVoucherInfo) GetVoucherStatus() string {
	if o == nil || IsNil(o.VoucherStatus) {
		var ret string
		return ret
	}
	return *o.VoucherStatus
}

// GetVoucherStatusOk returns a tuple with the VoucherStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVoucherInfo) GetVoucherStatusOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherStatus) {
		return nil, false
	}
	return o.VoucherStatus, true
}

// HasVoucherStatus returns a boolean if a field has been set.
func (o *UserVoucherInfo) HasVoucherStatus() bool {
	if o != nil && !IsNil(o.VoucherStatus) {
		return true
	}

	return false
}

// SetVoucherStatus gets a reference to the given string and assigns it to the VoucherStatus field.
func (o *UserVoucherInfo) SetVoucherStatus(v string) {
	o.VoucherStatus = &v
}

// GetVoucherType returns the VoucherType field value if set, zero value otherwise.
func (o *UserVoucherInfo) GetVoucherType() string {
	if o == nil || IsNil(o.VoucherType) {
		var ret string
		return ret
	}
	return *o.VoucherType
}

// GetVoucherTypeOk returns a tuple with the VoucherType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVoucherInfo) GetVoucherTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherType) {
		return nil, false
	}
	return o.VoucherType, true
}

// HasVoucherType returns a boolean if a field has been set.
func (o *UserVoucherInfo) HasVoucherType() bool {
	if o != nil && !IsNil(o.VoucherType) {
		return true
	}

	return false
}

// SetVoucherType gets a reference to the given string and assigns it to the VoucherType field.
func (o *UserVoucherInfo) SetVoucherType(v string) {
	o.VoucherType = &v
}

// GetVoucherUseRuleInfo returns the VoucherUseRuleInfo field value if set, zero value otherwise.
func (o *UserVoucherInfo) GetVoucherUseRuleInfo() VoucherUseRuleInfo {
	if o == nil || IsNil(o.VoucherUseRuleInfo) {
		var ret VoucherUseRuleInfo
		return ret
	}
	return *o.VoucherUseRuleInfo
}

// GetVoucherUseRuleInfoOk returns a tuple with the VoucherUseRuleInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVoucherInfo) GetVoucherUseRuleInfoOk() (*VoucherUseRuleInfo, bool) {
	if o == nil || IsNil(o.VoucherUseRuleInfo) {
		return nil, false
	}
	return o.VoucherUseRuleInfo, true
}

// HasVoucherUseRuleInfo returns a boolean if a field has been set.
func (o *UserVoucherInfo) HasVoucherUseRuleInfo() bool {
	if o != nil && !IsNil(o.VoucherUseRuleInfo) {
		return true
	}

	return false
}

// SetVoucherUseRuleInfo gets a reference to the given VoucherUseRuleInfo and assigns it to the VoucherUseRuleInfo field.
func (o *UserVoucherInfo) SetVoucherUseRuleInfo(v VoucherUseRuleInfo) {
	o.VoucherUseRuleInfo = &v
}

// GetVoucherUseRuleLiteInfo returns the VoucherUseRuleLiteInfo field value if set, zero value otherwise.
func (o *UserVoucherInfo) GetVoucherUseRuleLiteInfo() CommonVoucherUseRuleLiteInfo {
	if o == nil || IsNil(o.VoucherUseRuleLiteInfo) {
		var ret CommonVoucherUseRuleLiteInfo
		return ret
	}
	return *o.VoucherUseRuleLiteInfo
}

// GetVoucherUseRuleLiteInfoOk returns a tuple with the VoucherUseRuleLiteInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVoucherInfo) GetVoucherUseRuleLiteInfoOk() (*CommonVoucherUseRuleLiteInfo, bool) {
	if o == nil || IsNil(o.VoucherUseRuleLiteInfo) {
		return nil, false
	}
	return o.VoucherUseRuleLiteInfo, true
}

// HasVoucherUseRuleLiteInfo returns a boolean if a field has been set.
func (o *UserVoucherInfo) HasVoucherUseRuleLiteInfo() bool {
	if o != nil && !IsNil(o.VoucherUseRuleLiteInfo) {
		return true
	}

	return false
}

// SetVoucherUseRuleLiteInfo gets a reference to the given CommonVoucherUseRuleLiteInfo and assigns it to the VoucherUseRuleLiteInfo field.
func (o *UserVoucherInfo) SetVoucherUseRuleLiteInfo(v CommonVoucherUseRuleLiteInfo) {
	o.VoucherUseRuleLiteInfo = &v
}

func (o UserVoucherInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserVoucherInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityBaseInfo) {
		toSerialize["activity_base_info"] = o.ActivityBaseInfo
	}
	if !IsNil(o.ActivityId) {
		toSerialize["activity_id"] = o.ActivityId
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
	if !IsNil(o.VoucherAvailableScopeInfo) {
		toSerialize["voucher_available_scope_info"] = o.VoucherAvailableScopeInfo
	}
	if !IsNil(o.VoucherCustomerGuideInfo) {
		toSerialize["voucher_customer_guide_info"] = o.VoucherCustomerGuideInfo
	}
	if !IsNil(o.VoucherDeductInfo) {
		toSerialize["voucher_deduct_info"] = o.VoucherDeductInfo
	}
	if !IsNil(o.VoucherDisplayLiteInfo) {
		toSerialize["voucher_display_lite_info"] = o.VoucherDisplayLiteInfo
	}
	if !IsNil(o.VoucherDisplayPatternInfo) {
		toSerialize["voucher_display_pattern_info"] = o.VoucherDisplayPatternInfo
	}
	if !IsNil(o.VoucherId) {
		toSerialize["voucher_id"] = o.VoucherId
	}
	if !IsNil(o.VoucherName) {
		toSerialize["voucher_name"] = o.VoucherName
	}
	if !IsNil(o.VoucherSendModeInfo) {
		toSerialize["voucher_send_mode_info"] = o.VoucherSendModeInfo
	}
	if !IsNil(o.VoucherStatus) {
		toSerialize["voucher_status"] = o.VoucherStatus
	}
	if !IsNil(o.VoucherType) {
		toSerialize["voucher_type"] = o.VoucherType
	}
	if !IsNil(o.VoucherUseRuleInfo) {
		toSerialize["voucher_use_rule_info"] = o.VoucherUseRuleInfo
	}
	if !IsNil(o.VoucherUseRuleLiteInfo) {
		toSerialize["voucher_use_rule_lite_info"] = o.VoucherUseRuleLiteInfo
	}
	return toSerialize, nil
}

type NullableUserVoucherInfo struct {
	value *UserVoucherInfo
	isSet bool
}

func (v NullableUserVoucherInfo) Get() *UserVoucherInfo {
	return v.value
}

func (v *NullableUserVoucherInfo) Set(val *UserVoucherInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUserVoucherInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUserVoucherInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserVoucherInfo(val *UserVoucherInfo) *NullableUserVoucherInfo {
	return &NullableUserVoucherInfo{value: val, isSet: true}
}

func (v NullableUserVoucherInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserVoucherInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
