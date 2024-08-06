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

// checks if the AlipayMarketingActivityQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingActivityQueryResponseModel{}

// AlipayMarketingActivityQueryResponseModel struct for AlipayMarketingActivityQueryResponseModel
type AlipayMarketingActivityQueryResponseModel struct {
	ActivityBaseInfo *ActivityBaseInfo `json:"activity_base_info,omitempty"`
	// 活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// 活动名称。 不对用户进行展示，仅供商家在后台管理活动使用。
	ActivityName *string `json:"activity_name,omitempty"`
	// 活动状态。活动已激活，表示活动已经生效，等到活动开始(publish_start_time)之后用户就可以参与活动。活动已暂停，表示商户临时暂停该活动，该状态下用户不能参与活动。活动已结束，表示商户主动停止活动或活动到期结束(publish_end_time)不能再进行领取或修改等操作。
	ActivityStatus *string `json:"activity_status,omitempty"`
	// 归属商户PID
	BelongMerchantId *string `json:"belong_merchant_id,omitempty"`
	// 券发放结束时间。 格式为：yyyy-MM-dd HH:mm:ss
	PublishEndTime *string `json:"publish_end_time,omitempty"`
	// 券发放开始时间。 格式为：yyyy-MM-dd HH:mm:ss
	PublishStartTime *string `json:"publish_start_time,omitempty"`
	VoucherAvailableScopeInfo *VoucherAvailableScopeInfo `json:"voucher_available_scope_info,omitempty"`
	VoucherCustomerGuideInfo *VoucherCustomerGuideInfo `json:"voucher_customer_guide_info,omitempty"`
	VoucherDeductInfo *VoucherDeductInfo `json:"voucher_deduct_info,omitempty"`
	VoucherDisplayInfo *CommonVoucherDisplayInfo `json:"voucher_display_info,omitempty"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo `json:"voucher_display_pattern_info,omitempty"`
	VoucherSendModeInfo *VoucherSendModeInfo `json:"voucher_send_mode_info,omitempty"`
	VoucherSendRule *CommonVoucherSendRule `json:"voucher_send_rule,omitempty"`
	// 券类型。
	VoucherType *string `json:"voucher_type,omitempty"`
	VoucherUseRule *CommonVoucherUseRule `json:"voucher_use_rule,omitempty"`
	VoucherUseRuleInfo *VoucherUseRuleInfo `json:"voucher_use_rule_info,omitempty"`
}

// NewAlipayMarketingActivityQueryResponseModel instantiates a new AlipayMarketingActivityQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingActivityQueryResponseModel() *AlipayMarketingActivityQueryResponseModel {
	this := AlipayMarketingActivityQueryResponseModel{}
	return &this
}

// NewAlipayMarketingActivityQueryResponseModelWithDefaults instantiates a new AlipayMarketingActivityQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingActivityQueryResponseModelWithDefaults() *AlipayMarketingActivityQueryResponseModel {
	this := AlipayMarketingActivityQueryResponseModel{}
	return &this
}

// GetActivityBaseInfo returns the ActivityBaseInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityQueryResponseModel) GetActivityBaseInfo() ActivityBaseInfo {
	if o == nil || IsNil(o.ActivityBaseInfo) {
		var ret ActivityBaseInfo
		return ret
	}
	return *o.ActivityBaseInfo
}

// GetActivityBaseInfoOk returns a tuple with the ActivityBaseInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityQueryResponseModel) GetActivityBaseInfoOk() (*ActivityBaseInfo, bool) {
	if o == nil || IsNil(o.ActivityBaseInfo) {
		return nil, false
	}
	return o.ActivityBaseInfo, true
}

// HasActivityBaseInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityQueryResponseModel) HasActivityBaseInfo() bool {
	if o != nil && !IsNil(o.ActivityBaseInfo) {
		return true
	}

	return false
}

// SetActivityBaseInfo gets a reference to the given ActivityBaseInfo and assigns it to the ActivityBaseInfo field.
func (o *AlipayMarketingActivityQueryResponseModel) SetActivityBaseInfo(v ActivityBaseInfo) {
	o.ActivityBaseInfo = &v
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *AlipayMarketingActivityQueryResponseModel) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId) {
		var ret string
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityQueryResponseModel) GetActivityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *AlipayMarketingActivityQueryResponseModel) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given string and assigns it to the ActivityId field.
func (o *AlipayMarketingActivityQueryResponseModel) SetActivityId(v string) {
	o.ActivityId = &v
}

// GetActivityName returns the ActivityName field value if set, zero value otherwise.
func (o *AlipayMarketingActivityQueryResponseModel) GetActivityName() string {
	if o == nil || IsNil(o.ActivityName) {
		var ret string
		return ret
	}
	return *o.ActivityName
}

// GetActivityNameOk returns a tuple with the ActivityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityQueryResponseModel) GetActivityNameOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityName) {
		return nil, false
	}
	return o.ActivityName, true
}

// HasActivityName returns a boolean if a field has been set.
func (o *AlipayMarketingActivityQueryResponseModel) HasActivityName() bool {
	if o != nil && !IsNil(o.ActivityName) {
		return true
	}

	return false
}

// SetActivityName gets a reference to the given string and assigns it to the ActivityName field.
func (o *AlipayMarketingActivityQueryResponseModel) SetActivityName(v string) {
	o.ActivityName = &v
}

// GetActivityStatus returns the ActivityStatus field value if set, zero value otherwise.
func (o *AlipayMarketingActivityQueryResponseModel) GetActivityStatus() string {
	if o == nil || IsNil(o.ActivityStatus) {
		var ret string
		return ret
	}
	return *o.ActivityStatus
}

// GetActivityStatusOk returns a tuple with the ActivityStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityQueryResponseModel) GetActivityStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityStatus) {
		return nil, false
	}
	return o.ActivityStatus, true
}

// HasActivityStatus returns a boolean if a field has been set.
func (o *AlipayMarketingActivityQueryResponseModel) HasActivityStatus() bool {
	if o != nil && !IsNil(o.ActivityStatus) {
		return true
	}

	return false
}

// SetActivityStatus gets a reference to the given string and assigns it to the ActivityStatus field.
func (o *AlipayMarketingActivityQueryResponseModel) SetActivityStatus(v string) {
	o.ActivityStatus = &v
}

// GetBelongMerchantId returns the BelongMerchantId field value if set, zero value otherwise.
func (o *AlipayMarketingActivityQueryResponseModel) GetBelongMerchantId() string {
	if o == nil || IsNil(o.BelongMerchantId) {
		var ret string
		return ret
	}
	return *o.BelongMerchantId
}

// GetBelongMerchantIdOk returns a tuple with the BelongMerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityQueryResponseModel) GetBelongMerchantIdOk() (*string, bool) {
	if o == nil || IsNil(o.BelongMerchantId) {
		return nil, false
	}
	return o.BelongMerchantId, true
}

// HasBelongMerchantId returns a boolean if a field has been set.
func (o *AlipayMarketingActivityQueryResponseModel) HasBelongMerchantId() bool {
	if o != nil && !IsNil(o.BelongMerchantId) {
		return true
	}

	return false
}

// SetBelongMerchantId gets a reference to the given string and assigns it to the BelongMerchantId field.
func (o *AlipayMarketingActivityQueryResponseModel) SetBelongMerchantId(v string) {
	o.BelongMerchantId = &v
}

// GetPublishEndTime returns the PublishEndTime field value if set, zero value otherwise.
func (o *AlipayMarketingActivityQueryResponseModel) GetPublishEndTime() string {
	if o == nil || IsNil(o.PublishEndTime) {
		var ret string
		return ret
	}
	return *o.PublishEndTime
}

// GetPublishEndTimeOk returns a tuple with the PublishEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityQueryResponseModel) GetPublishEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PublishEndTime) {
		return nil, false
	}
	return o.PublishEndTime, true
}

// HasPublishEndTime returns a boolean if a field has been set.
func (o *AlipayMarketingActivityQueryResponseModel) HasPublishEndTime() bool {
	if o != nil && !IsNil(o.PublishEndTime) {
		return true
	}

	return false
}

// SetPublishEndTime gets a reference to the given string and assigns it to the PublishEndTime field.
func (o *AlipayMarketingActivityQueryResponseModel) SetPublishEndTime(v string) {
	o.PublishEndTime = &v
}

// GetPublishStartTime returns the PublishStartTime field value if set, zero value otherwise.
func (o *AlipayMarketingActivityQueryResponseModel) GetPublishStartTime() string {
	if o == nil || IsNil(o.PublishStartTime) {
		var ret string
		return ret
	}
	return *o.PublishStartTime
}

// GetPublishStartTimeOk returns a tuple with the PublishStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityQueryResponseModel) GetPublishStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PublishStartTime) {
		return nil, false
	}
	return o.PublishStartTime, true
}

// HasPublishStartTime returns a boolean if a field has been set.
func (o *AlipayMarketingActivityQueryResponseModel) HasPublishStartTime() bool {
	if o != nil && !IsNil(o.PublishStartTime) {
		return true
	}

	return false
}

// SetPublishStartTime gets a reference to the given string and assigns it to the PublishStartTime field.
func (o *AlipayMarketingActivityQueryResponseModel) SetPublishStartTime(v string) {
	o.PublishStartTime = &v
}

// GetVoucherAvailableScopeInfo returns the VoucherAvailableScopeInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherAvailableScopeInfo() VoucherAvailableScopeInfo {
	if o == nil || IsNil(o.VoucherAvailableScopeInfo) {
		var ret VoucherAvailableScopeInfo
		return ret
	}
	return *o.VoucherAvailableScopeInfo
}

// GetVoucherAvailableScopeInfoOk returns a tuple with the VoucherAvailableScopeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherAvailableScopeInfoOk() (*VoucherAvailableScopeInfo, bool) {
	if o == nil || IsNil(o.VoucherAvailableScopeInfo) {
		return nil, false
	}
	return o.VoucherAvailableScopeInfo, true
}

// HasVoucherAvailableScopeInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityQueryResponseModel) HasVoucherAvailableScopeInfo() bool {
	if o != nil && !IsNil(o.VoucherAvailableScopeInfo) {
		return true
	}

	return false
}

// SetVoucherAvailableScopeInfo gets a reference to the given VoucherAvailableScopeInfo and assigns it to the VoucherAvailableScopeInfo field.
func (o *AlipayMarketingActivityQueryResponseModel) SetVoucherAvailableScopeInfo(v VoucherAvailableScopeInfo) {
	o.VoucherAvailableScopeInfo = &v
}

// GetVoucherCustomerGuideInfo returns the VoucherCustomerGuideInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherCustomerGuideInfo() VoucherCustomerGuideInfo {
	if o == nil || IsNil(o.VoucherCustomerGuideInfo) {
		var ret VoucherCustomerGuideInfo
		return ret
	}
	return *o.VoucherCustomerGuideInfo
}

// GetVoucherCustomerGuideInfoOk returns a tuple with the VoucherCustomerGuideInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherCustomerGuideInfoOk() (*VoucherCustomerGuideInfo, bool) {
	if o == nil || IsNil(o.VoucherCustomerGuideInfo) {
		return nil, false
	}
	return o.VoucherCustomerGuideInfo, true
}

// HasVoucherCustomerGuideInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityQueryResponseModel) HasVoucherCustomerGuideInfo() bool {
	if o != nil && !IsNil(o.VoucherCustomerGuideInfo) {
		return true
	}

	return false
}

// SetVoucherCustomerGuideInfo gets a reference to the given VoucherCustomerGuideInfo and assigns it to the VoucherCustomerGuideInfo field.
func (o *AlipayMarketingActivityQueryResponseModel) SetVoucherCustomerGuideInfo(v VoucherCustomerGuideInfo) {
	o.VoucherCustomerGuideInfo = &v
}

// GetVoucherDeductInfo returns the VoucherDeductInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherDeductInfo() VoucherDeductInfo {
	if o == nil || IsNil(o.VoucherDeductInfo) {
		var ret VoucherDeductInfo
		return ret
	}
	return *o.VoucherDeductInfo
}

// GetVoucherDeductInfoOk returns a tuple with the VoucherDeductInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherDeductInfoOk() (*VoucherDeductInfo, bool) {
	if o == nil || IsNil(o.VoucherDeductInfo) {
		return nil, false
	}
	return o.VoucherDeductInfo, true
}

// HasVoucherDeductInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityQueryResponseModel) HasVoucherDeductInfo() bool {
	if o != nil && !IsNil(o.VoucherDeductInfo) {
		return true
	}

	return false
}

// SetVoucherDeductInfo gets a reference to the given VoucherDeductInfo and assigns it to the VoucherDeductInfo field.
func (o *AlipayMarketingActivityQueryResponseModel) SetVoucherDeductInfo(v VoucherDeductInfo) {
	o.VoucherDeductInfo = &v
}

// GetVoucherDisplayInfo returns the VoucherDisplayInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherDisplayInfo() CommonVoucherDisplayInfo {
	if o == nil || IsNil(o.VoucherDisplayInfo) {
		var ret CommonVoucherDisplayInfo
		return ret
	}
	return *o.VoucherDisplayInfo
}

// GetVoucherDisplayInfoOk returns a tuple with the VoucherDisplayInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherDisplayInfoOk() (*CommonVoucherDisplayInfo, bool) {
	if o == nil || IsNil(o.VoucherDisplayInfo) {
		return nil, false
	}
	return o.VoucherDisplayInfo, true
}

// HasVoucherDisplayInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityQueryResponseModel) HasVoucherDisplayInfo() bool {
	if o != nil && !IsNil(o.VoucherDisplayInfo) {
		return true
	}

	return false
}

// SetVoucherDisplayInfo gets a reference to the given CommonVoucherDisplayInfo and assigns it to the VoucherDisplayInfo field.
func (o *AlipayMarketingActivityQueryResponseModel) SetVoucherDisplayInfo(v CommonVoucherDisplayInfo) {
	o.VoucherDisplayInfo = &v
}

// GetVoucherDisplayPatternInfo returns the VoucherDisplayPatternInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherDisplayPatternInfo() VoucherDisplayPatternInfo {
	if o == nil || IsNil(o.VoucherDisplayPatternInfo) {
		var ret VoucherDisplayPatternInfo
		return ret
	}
	return *o.VoucherDisplayPatternInfo
}

// GetVoucherDisplayPatternInfoOk returns a tuple with the VoucherDisplayPatternInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherDisplayPatternInfoOk() (*VoucherDisplayPatternInfo, bool) {
	if o == nil || IsNil(o.VoucherDisplayPatternInfo) {
		return nil, false
	}
	return o.VoucherDisplayPatternInfo, true
}

// HasVoucherDisplayPatternInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityQueryResponseModel) HasVoucherDisplayPatternInfo() bool {
	if o != nil && !IsNil(o.VoucherDisplayPatternInfo) {
		return true
	}

	return false
}

// SetVoucherDisplayPatternInfo gets a reference to the given VoucherDisplayPatternInfo and assigns it to the VoucherDisplayPatternInfo field.
func (o *AlipayMarketingActivityQueryResponseModel) SetVoucherDisplayPatternInfo(v VoucherDisplayPatternInfo) {
	o.VoucherDisplayPatternInfo = &v
}

// GetVoucherSendModeInfo returns the VoucherSendModeInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherSendModeInfo() VoucherSendModeInfo {
	if o == nil || IsNil(o.VoucherSendModeInfo) {
		var ret VoucherSendModeInfo
		return ret
	}
	return *o.VoucherSendModeInfo
}

// GetVoucherSendModeInfoOk returns a tuple with the VoucherSendModeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherSendModeInfoOk() (*VoucherSendModeInfo, bool) {
	if o == nil || IsNil(o.VoucherSendModeInfo) {
		return nil, false
	}
	return o.VoucherSendModeInfo, true
}

// HasVoucherSendModeInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityQueryResponseModel) HasVoucherSendModeInfo() bool {
	if o != nil && !IsNil(o.VoucherSendModeInfo) {
		return true
	}

	return false
}

// SetVoucherSendModeInfo gets a reference to the given VoucherSendModeInfo and assigns it to the VoucherSendModeInfo field.
func (o *AlipayMarketingActivityQueryResponseModel) SetVoucherSendModeInfo(v VoucherSendModeInfo) {
	o.VoucherSendModeInfo = &v
}

// GetVoucherSendRule returns the VoucherSendRule field value if set, zero value otherwise.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherSendRule() CommonVoucherSendRule {
	if o == nil || IsNil(o.VoucherSendRule) {
		var ret CommonVoucherSendRule
		return ret
	}
	return *o.VoucherSendRule
}

// GetVoucherSendRuleOk returns a tuple with the VoucherSendRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherSendRuleOk() (*CommonVoucherSendRule, bool) {
	if o == nil || IsNil(o.VoucherSendRule) {
		return nil, false
	}
	return o.VoucherSendRule, true
}

// HasVoucherSendRule returns a boolean if a field has been set.
func (o *AlipayMarketingActivityQueryResponseModel) HasVoucherSendRule() bool {
	if o != nil && !IsNil(o.VoucherSendRule) {
		return true
	}

	return false
}

// SetVoucherSendRule gets a reference to the given CommonVoucherSendRule and assigns it to the VoucherSendRule field.
func (o *AlipayMarketingActivityQueryResponseModel) SetVoucherSendRule(v CommonVoucherSendRule) {
	o.VoucherSendRule = &v
}

// GetVoucherType returns the VoucherType field value if set, zero value otherwise.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherType() string {
	if o == nil || IsNil(o.VoucherType) {
		var ret string
		return ret
	}
	return *o.VoucherType
}

// GetVoucherTypeOk returns a tuple with the VoucherType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherType) {
		return nil, false
	}
	return o.VoucherType, true
}

// HasVoucherType returns a boolean if a field has been set.
func (o *AlipayMarketingActivityQueryResponseModel) HasVoucherType() bool {
	if o != nil && !IsNil(o.VoucherType) {
		return true
	}

	return false
}

// SetVoucherType gets a reference to the given string and assigns it to the VoucherType field.
func (o *AlipayMarketingActivityQueryResponseModel) SetVoucherType(v string) {
	o.VoucherType = &v
}

// GetVoucherUseRule returns the VoucherUseRule field value if set, zero value otherwise.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherUseRule() CommonVoucherUseRule {
	if o == nil || IsNil(o.VoucherUseRule) {
		var ret CommonVoucherUseRule
		return ret
	}
	return *o.VoucherUseRule
}

// GetVoucherUseRuleOk returns a tuple with the VoucherUseRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherUseRuleOk() (*CommonVoucherUseRule, bool) {
	if o == nil || IsNil(o.VoucherUseRule) {
		return nil, false
	}
	return o.VoucherUseRule, true
}

// HasVoucherUseRule returns a boolean if a field has been set.
func (o *AlipayMarketingActivityQueryResponseModel) HasVoucherUseRule() bool {
	if o != nil && !IsNil(o.VoucherUseRule) {
		return true
	}

	return false
}

// SetVoucherUseRule gets a reference to the given CommonVoucherUseRule and assigns it to the VoucherUseRule field.
func (o *AlipayMarketingActivityQueryResponseModel) SetVoucherUseRule(v CommonVoucherUseRule) {
	o.VoucherUseRule = &v
}

// GetVoucherUseRuleInfo returns the VoucherUseRuleInfo field value if set, zero value otherwise.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherUseRuleInfo() VoucherUseRuleInfo {
	if o == nil || IsNil(o.VoucherUseRuleInfo) {
		var ret VoucherUseRuleInfo
		return ret
	}
	return *o.VoucherUseRuleInfo
}

// GetVoucherUseRuleInfoOk returns a tuple with the VoucherUseRuleInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingActivityQueryResponseModel) GetVoucherUseRuleInfoOk() (*VoucherUseRuleInfo, bool) {
	if o == nil || IsNil(o.VoucherUseRuleInfo) {
		return nil, false
	}
	return o.VoucherUseRuleInfo, true
}

// HasVoucherUseRuleInfo returns a boolean if a field has been set.
func (o *AlipayMarketingActivityQueryResponseModel) HasVoucherUseRuleInfo() bool {
	if o != nil && !IsNil(o.VoucherUseRuleInfo) {
		return true
	}

	return false
}

// SetVoucherUseRuleInfo gets a reference to the given VoucherUseRuleInfo and assigns it to the VoucherUseRuleInfo field.
func (o *AlipayMarketingActivityQueryResponseModel) SetVoucherUseRuleInfo(v VoucherUseRuleInfo) {
	o.VoucherUseRuleInfo = &v
}

func (o AlipayMarketingActivityQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingActivityQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityBaseInfo) {
		toSerialize["activity_base_info"] = o.ActivityBaseInfo
	}
	if !IsNil(o.ActivityId) {
		toSerialize["activity_id"] = o.ActivityId
	}
	if !IsNil(o.ActivityName) {
		toSerialize["activity_name"] = o.ActivityName
	}
	if !IsNil(o.ActivityStatus) {
		toSerialize["activity_status"] = o.ActivityStatus
	}
	if !IsNil(o.BelongMerchantId) {
		toSerialize["belong_merchant_id"] = o.BelongMerchantId
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

type NullableAlipayMarketingActivityQueryResponseModel struct {
	value *AlipayMarketingActivityQueryResponseModel
	isSet bool
}

func (v NullableAlipayMarketingActivityQueryResponseModel) Get() *AlipayMarketingActivityQueryResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingActivityQueryResponseModel) Set(val *AlipayMarketingActivityQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityQueryResponseModel(val *AlipayMarketingActivityQueryResponseModel) *NullableAlipayMarketingActivityQueryResponseModel {
	return &NullableAlipayMarketingActivityQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

