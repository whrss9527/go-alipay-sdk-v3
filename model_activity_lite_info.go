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

// checks if the ActivityLiteInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityLiteInfo{}

// ActivityLiteInfo struct for ActivityLiteInfo
type ActivityLiteInfo struct {
	ActivityBaseInfo *ActivityBaseInfo `json:"activity_base_info,omitempty"`
	// 活动 id。
	ActivityId *string `json:"activity_id,omitempty"`
	// 活动名称
	ActivityName *string `json:"activity_name,omitempty"`
	// 活动状态 。 ACTIVE:活动已激活，表示活动已经生效，等到活动开始(publish_start_time)之后用户就可以参与活动。  PAUSE:活动已暂停，表示商户临时暂停该活动，该状态下用户不能参与活动。
	ActivityStatus *string `json:"activity_status,omitempty"`
	// 归属商户PID
	BelongMerchantId *string `json:"belong_merchant_id,omitempty"`
	// 活动创建时间。
	CreateTime *string `json:"create_time,omitempty"`
	// 券发放结束时间。格式为：yyyy-MM-dd HH:mm:ss。
	PublishEndTime *string `json:"publish_end_time,omitempty"`
	// 券发放开始时间。 格式为：yyyy-MM-dd HH:mm:ss
	PublishStartTime          *string                       `json:"publish_start_time,omitempty"`
	VoucherAvailableScopeInfo *VoucherAvailableScopeInfo    `json:"voucher_available_scope_info,omitempty"`
	VoucherCustomerGuideInfo  *VoucherCustomerGuideInfo     `json:"voucher_customer_guide_info,omitempty"`
	VoucherDeductInfo         *VoucherDeductInfo            `json:"voucher_deduct_info,omitempty"`
	VoucherDisplayLiteInfo    *CommonVoucherDisplayLiteInfo `json:"voucher_display_lite_info,omitempty"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo    `json:"voucher_display_pattern_info,omitempty"`
	VoucherSendModeInfo       *VoucherSendModeInfo          `json:"voucher_send_mode_info,omitempty"`
	// 券类型。
	VoucherType            *string                       `json:"voucher_type,omitempty"`
	VoucherUseRuleInfo     *VoucherUseRuleInfo           `json:"voucher_use_rule_info,omitempty"`
	VoucherUseRuleLiteInfo *CommonVoucherUseRuleLiteInfo `json:"voucher_use_rule_lite_info,omitempty"`
}

// NewActivityLiteInfo instantiates a new ActivityLiteInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityLiteInfo() *ActivityLiteInfo {
	this := ActivityLiteInfo{}
	return &this
}

// NewActivityLiteInfoWithDefaults instantiates a new ActivityLiteInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityLiteInfoWithDefaults() *ActivityLiteInfo {
	this := ActivityLiteInfo{}
	return &this
}

// GetActivityBaseInfo returns the ActivityBaseInfo field value if set, zero value otherwise.
func (o *ActivityLiteInfo) GetActivityBaseInfo() ActivityBaseInfo {
	if o == nil || IsNil(o.ActivityBaseInfo) {
		var ret ActivityBaseInfo
		return ret
	}
	return *o.ActivityBaseInfo
}

// GetActivityBaseInfoOk returns a tuple with the ActivityBaseInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLiteInfo) GetActivityBaseInfoOk() (*ActivityBaseInfo, bool) {
	if o == nil || IsNil(o.ActivityBaseInfo) {
		return nil, false
	}
	return o.ActivityBaseInfo, true
}

// HasActivityBaseInfo returns a boolean if a field has been set.
func (o *ActivityLiteInfo) HasActivityBaseInfo() bool {
	if o != nil && !IsNil(o.ActivityBaseInfo) {
		return true
	}

	return false
}

// SetActivityBaseInfo gets a reference to the given ActivityBaseInfo and assigns it to the ActivityBaseInfo field.
func (o *ActivityLiteInfo) SetActivityBaseInfo(v ActivityBaseInfo) {
	o.ActivityBaseInfo = &v
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *ActivityLiteInfo) GetActivityId() string {
	if o == nil || IsNil(o.ActivityId) {
		var ret string
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLiteInfo) GetActivityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *ActivityLiteInfo) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given string and assigns it to the ActivityId field.
func (o *ActivityLiteInfo) SetActivityId(v string) {
	o.ActivityId = &v
}

// GetActivityName returns the ActivityName field value if set, zero value otherwise.
func (o *ActivityLiteInfo) GetActivityName() string {
	if o == nil || IsNil(o.ActivityName) {
		var ret string
		return ret
	}
	return *o.ActivityName
}

// GetActivityNameOk returns a tuple with the ActivityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLiteInfo) GetActivityNameOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityName) {
		return nil, false
	}
	return o.ActivityName, true
}

// HasActivityName returns a boolean if a field has been set.
func (o *ActivityLiteInfo) HasActivityName() bool {
	if o != nil && !IsNil(o.ActivityName) {
		return true
	}

	return false
}

// SetActivityName gets a reference to the given string and assigns it to the ActivityName field.
func (o *ActivityLiteInfo) SetActivityName(v string) {
	o.ActivityName = &v
}

// GetActivityStatus returns the ActivityStatus field value if set, zero value otherwise.
func (o *ActivityLiteInfo) GetActivityStatus() string {
	if o == nil || IsNil(o.ActivityStatus) {
		var ret string
		return ret
	}
	return *o.ActivityStatus
}

// GetActivityStatusOk returns a tuple with the ActivityStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLiteInfo) GetActivityStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityStatus) {
		return nil, false
	}
	return o.ActivityStatus, true
}

// HasActivityStatus returns a boolean if a field has been set.
func (o *ActivityLiteInfo) HasActivityStatus() bool {
	if o != nil && !IsNil(o.ActivityStatus) {
		return true
	}

	return false
}

// SetActivityStatus gets a reference to the given string and assigns it to the ActivityStatus field.
func (o *ActivityLiteInfo) SetActivityStatus(v string) {
	o.ActivityStatus = &v
}

// GetBelongMerchantId returns the BelongMerchantId field value if set, zero value otherwise.
func (o *ActivityLiteInfo) GetBelongMerchantId() string {
	if o == nil || IsNil(o.BelongMerchantId) {
		var ret string
		return ret
	}
	return *o.BelongMerchantId
}

// GetBelongMerchantIdOk returns a tuple with the BelongMerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLiteInfo) GetBelongMerchantIdOk() (*string, bool) {
	if o == nil || IsNil(o.BelongMerchantId) {
		return nil, false
	}
	return o.BelongMerchantId, true
}

// HasBelongMerchantId returns a boolean if a field has been set.
func (o *ActivityLiteInfo) HasBelongMerchantId() bool {
	if o != nil && !IsNil(o.BelongMerchantId) {
		return true
	}

	return false
}

// SetBelongMerchantId gets a reference to the given string and assigns it to the BelongMerchantId field.
func (o *ActivityLiteInfo) SetBelongMerchantId(v string) {
	o.BelongMerchantId = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *ActivityLiteInfo) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLiteInfo) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *ActivityLiteInfo) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *ActivityLiteInfo) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetPublishEndTime returns the PublishEndTime field value if set, zero value otherwise.
func (o *ActivityLiteInfo) GetPublishEndTime() string {
	if o == nil || IsNil(o.PublishEndTime) {
		var ret string
		return ret
	}
	return *o.PublishEndTime
}

// GetPublishEndTimeOk returns a tuple with the PublishEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLiteInfo) GetPublishEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PublishEndTime) {
		return nil, false
	}
	return o.PublishEndTime, true
}

// HasPublishEndTime returns a boolean if a field has been set.
func (o *ActivityLiteInfo) HasPublishEndTime() bool {
	if o != nil && !IsNil(o.PublishEndTime) {
		return true
	}

	return false
}

// SetPublishEndTime gets a reference to the given string and assigns it to the PublishEndTime field.
func (o *ActivityLiteInfo) SetPublishEndTime(v string) {
	o.PublishEndTime = &v
}

// GetPublishStartTime returns the PublishStartTime field value if set, zero value otherwise.
func (o *ActivityLiteInfo) GetPublishStartTime() string {
	if o == nil || IsNil(o.PublishStartTime) {
		var ret string
		return ret
	}
	return *o.PublishStartTime
}

// GetPublishStartTimeOk returns a tuple with the PublishStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLiteInfo) GetPublishStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PublishStartTime) {
		return nil, false
	}
	return o.PublishStartTime, true
}

// HasPublishStartTime returns a boolean if a field has been set.
func (o *ActivityLiteInfo) HasPublishStartTime() bool {
	if o != nil && !IsNil(o.PublishStartTime) {
		return true
	}

	return false
}

// SetPublishStartTime gets a reference to the given string and assigns it to the PublishStartTime field.
func (o *ActivityLiteInfo) SetPublishStartTime(v string) {
	o.PublishStartTime = &v
}

// GetVoucherAvailableScopeInfo returns the VoucherAvailableScopeInfo field value if set, zero value otherwise.
func (o *ActivityLiteInfo) GetVoucherAvailableScopeInfo() VoucherAvailableScopeInfo {
	if o == nil || IsNil(o.VoucherAvailableScopeInfo) {
		var ret VoucherAvailableScopeInfo
		return ret
	}
	return *o.VoucherAvailableScopeInfo
}

// GetVoucherAvailableScopeInfoOk returns a tuple with the VoucherAvailableScopeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLiteInfo) GetVoucherAvailableScopeInfoOk() (*VoucherAvailableScopeInfo, bool) {
	if o == nil || IsNil(o.VoucherAvailableScopeInfo) {
		return nil, false
	}
	return o.VoucherAvailableScopeInfo, true
}

// HasVoucherAvailableScopeInfo returns a boolean if a field has been set.
func (o *ActivityLiteInfo) HasVoucherAvailableScopeInfo() bool {
	if o != nil && !IsNil(o.VoucherAvailableScopeInfo) {
		return true
	}

	return false
}

// SetVoucherAvailableScopeInfo gets a reference to the given VoucherAvailableScopeInfo and assigns it to the VoucherAvailableScopeInfo field.
func (o *ActivityLiteInfo) SetVoucherAvailableScopeInfo(v VoucherAvailableScopeInfo) {
	o.VoucherAvailableScopeInfo = &v
}

// GetVoucherCustomerGuideInfo returns the VoucherCustomerGuideInfo field value if set, zero value otherwise.
func (o *ActivityLiteInfo) GetVoucherCustomerGuideInfo() VoucherCustomerGuideInfo {
	if o == nil || IsNil(o.VoucherCustomerGuideInfo) {
		var ret VoucherCustomerGuideInfo
		return ret
	}
	return *o.VoucherCustomerGuideInfo
}

// GetVoucherCustomerGuideInfoOk returns a tuple with the VoucherCustomerGuideInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLiteInfo) GetVoucherCustomerGuideInfoOk() (*VoucherCustomerGuideInfo, bool) {
	if o == nil || IsNil(o.VoucherCustomerGuideInfo) {
		return nil, false
	}
	return o.VoucherCustomerGuideInfo, true
}

// HasVoucherCustomerGuideInfo returns a boolean if a field has been set.
func (o *ActivityLiteInfo) HasVoucherCustomerGuideInfo() bool {
	if o != nil && !IsNil(o.VoucherCustomerGuideInfo) {
		return true
	}

	return false
}

// SetVoucherCustomerGuideInfo gets a reference to the given VoucherCustomerGuideInfo and assigns it to the VoucherCustomerGuideInfo field.
func (o *ActivityLiteInfo) SetVoucherCustomerGuideInfo(v VoucherCustomerGuideInfo) {
	o.VoucherCustomerGuideInfo = &v
}

// GetVoucherDeductInfo returns the VoucherDeductInfo field value if set, zero value otherwise.
func (o *ActivityLiteInfo) GetVoucherDeductInfo() VoucherDeductInfo {
	if o == nil || IsNil(o.VoucherDeductInfo) {
		var ret VoucherDeductInfo
		return ret
	}
	return *o.VoucherDeductInfo
}

// GetVoucherDeductInfoOk returns a tuple with the VoucherDeductInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLiteInfo) GetVoucherDeductInfoOk() (*VoucherDeductInfo, bool) {
	if o == nil || IsNil(o.VoucherDeductInfo) {
		return nil, false
	}
	return o.VoucherDeductInfo, true
}

// HasVoucherDeductInfo returns a boolean if a field has been set.
func (o *ActivityLiteInfo) HasVoucherDeductInfo() bool {
	if o != nil && !IsNil(o.VoucherDeductInfo) {
		return true
	}

	return false
}

// SetVoucherDeductInfo gets a reference to the given VoucherDeductInfo and assigns it to the VoucherDeductInfo field.
func (o *ActivityLiteInfo) SetVoucherDeductInfo(v VoucherDeductInfo) {
	o.VoucherDeductInfo = &v
}

// GetVoucherDisplayLiteInfo returns the VoucherDisplayLiteInfo field value if set, zero value otherwise.
func (o *ActivityLiteInfo) GetVoucherDisplayLiteInfo() CommonVoucherDisplayLiteInfo {
	if o == nil || IsNil(o.VoucherDisplayLiteInfo) {
		var ret CommonVoucherDisplayLiteInfo
		return ret
	}
	return *o.VoucherDisplayLiteInfo
}

// GetVoucherDisplayLiteInfoOk returns a tuple with the VoucherDisplayLiteInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLiteInfo) GetVoucherDisplayLiteInfoOk() (*CommonVoucherDisplayLiteInfo, bool) {
	if o == nil || IsNil(o.VoucherDisplayLiteInfo) {
		return nil, false
	}
	return o.VoucherDisplayLiteInfo, true
}

// HasVoucherDisplayLiteInfo returns a boolean if a field has been set.
func (o *ActivityLiteInfo) HasVoucherDisplayLiteInfo() bool {
	if o != nil && !IsNil(o.VoucherDisplayLiteInfo) {
		return true
	}

	return false
}

// SetVoucherDisplayLiteInfo gets a reference to the given CommonVoucherDisplayLiteInfo and assigns it to the VoucherDisplayLiteInfo field.
func (o *ActivityLiteInfo) SetVoucherDisplayLiteInfo(v CommonVoucherDisplayLiteInfo) {
	o.VoucherDisplayLiteInfo = &v
}

// GetVoucherDisplayPatternInfo returns the VoucherDisplayPatternInfo field value if set, zero value otherwise.
func (o *ActivityLiteInfo) GetVoucherDisplayPatternInfo() VoucherDisplayPatternInfo {
	if o == nil || IsNil(o.VoucherDisplayPatternInfo) {
		var ret VoucherDisplayPatternInfo
		return ret
	}
	return *o.VoucherDisplayPatternInfo
}

// GetVoucherDisplayPatternInfoOk returns a tuple with the VoucherDisplayPatternInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLiteInfo) GetVoucherDisplayPatternInfoOk() (*VoucherDisplayPatternInfo, bool) {
	if o == nil || IsNil(o.VoucherDisplayPatternInfo) {
		return nil, false
	}
	return o.VoucherDisplayPatternInfo, true
}

// HasVoucherDisplayPatternInfo returns a boolean if a field has been set.
func (o *ActivityLiteInfo) HasVoucherDisplayPatternInfo() bool {
	if o != nil && !IsNil(o.VoucherDisplayPatternInfo) {
		return true
	}

	return false
}

// SetVoucherDisplayPatternInfo gets a reference to the given VoucherDisplayPatternInfo and assigns it to the VoucherDisplayPatternInfo field.
func (o *ActivityLiteInfo) SetVoucherDisplayPatternInfo(v VoucherDisplayPatternInfo) {
	o.VoucherDisplayPatternInfo = &v
}

// GetVoucherSendModeInfo returns the VoucherSendModeInfo field value if set, zero value otherwise.
func (o *ActivityLiteInfo) GetVoucherSendModeInfo() VoucherSendModeInfo {
	if o == nil || IsNil(o.VoucherSendModeInfo) {
		var ret VoucherSendModeInfo
		return ret
	}
	return *o.VoucherSendModeInfo
}

// GetVoucherSendModeInfoOk returns a tuple with the VoucherSendModeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLiteInfo) GetVoucherSendModeInfoOk() (*VoucherSendModeInfo, bool) {
	if o == nil || IsNil(o.VoucherSendModeInfo) {
		return nil, false
	}
	return o.VoucherSendModeInfo, true
}

// HasVoucherSendModeInfo returns a boolean if a field has been set.
func (o *ActivityLiteInfo) HasVoucherSendModeInfo() bool {
	if o != nil && !IsNil(o.VoucherSendModeInfo) {
		return true
	}

	return false
}

// SetVoucherSendModeInfo gets a reference to the given VoucherSendModeInfo and assigns it to the VoucherSendModeInfo field.
func (o *ActivityLiteInfo) SetVoucherSendModeInfo(v VoucherSendModeInfo) {
	o.VoucherSendModeInfo = &v
}

// GetVoucherType returns the VoucherType field value if set, zero value otherwise.
func (o *ActivityLiteInfo) GetVoucherType() string {
	if o == nil || IsNil(o.VoucherType) {
		var ret string
		return ret
	}
	return *o.VoucherType
}

// GetVoucherTypeOk returns a tuple with the VoucherType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLiteInfo) GetVoucherTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VoucherType) {
		return nil, false
	}
	return o.VoucherType, true
}

// HasVoucherType returns a boolean if a field has been set.
func (o *ActivityLiteInfo) HasVoucherType() bool {
	if o != nil && !IsNil(o.VoucherType) {
		return true
	}

	return false
}

// SetVoucherType gets a reference to the given string and assigns it to the VoucherType field.
func (o *ActivityLiteInfo) SetVoucherType(v string) {
	o.VoucherType = &v
}

// GetVoucherUseRuleInfo returns the VoucherUseRuleInfo field value if set, zero value otherwise.
func (o *ActivityLiteInfo) GetVoucherUseRuleInfo() VoucherUseRuleInfo {
	if o == nil || IsNil(o.VoucherUseRuleInfo) {
		var ret VoucherUseRuleInfo
		return ret
	}
	return *o.VoucherUseRuleInfo
}

// GetVoucherUseRuleInfoOk returns a tuple with the VoucherUseRuleInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLiteInfo) GetVoucherUseRuleInfoOk() (*VoucherUseRuleInfo, bool) {
	if o == nil || IsNil(o.VoucherUseRuleInfo) {
		return nil, false
	}
	return o.VoucherUseRuleInfo, true
}

// HasVoucherUseRuleInfo returns a boolean if a field has been set.
func (o *ActivityLiteInfo) HasVoucherUseRuleInfo() bool {
	if o != nil && !IsNil(o.VoucherUseRuleInfo) {
		return true
	}

	return false
}

// SetVoucherUseRuleInfo gets a reference to the given VoucherUseRuleInfo and assigns it to the VoucherUseRuleInfo field.
func (o *ActivityLiteInfo) SetVoucherUseRuleInfo(v VoucherUseRuleInfo) {
	o.VoucherUseRuleInfo = &v
}

// GetVoucherUseRuleLiteInfo returns the VoucherUseRuleLiteInfo field value if set, zero value otherwise.
func (o *ActivityLiteInfo) GetVoucherUseRuleLiteInfo() CommonVoucherUseRuleLiteInfo {
	if o == nil || IsNil(o.VoucherUseRuleLiteInfo) {
		var ret CommonVoucherUseRuleLiteInfo
		return ret
	}
	return *o.VoucherUseRuleLiteInfo
}

// GetVoucherUseRuleLiteInfoOk returns a tuple with the VoucherUseRuleLiteInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLiteInfo) GetVoucherUseRuleLiteInfoOk() (*CommonVoucherUseRuleLiteInfo, bool) {
	if o == nil || IsNil(o.VoucherUseRuleLiteInfo) {
		return nil, false
	}
	return o.VoucherUseRuleLiteInfo, true
}

// HasVoucherUseRuleLiteInfo returns a boolean if a field has been set.
func (o *ActivityLiteInfo) HasVoucherUseRuleLiteInfo() bool {
	if o != nil && !IsNil(o.VoucherUseRuleLiteInfo) {
		return true
	}

	return false
}

// SetVoucherUseRuleLiteInfo gets a reference to the given CommonVoucherUseRuleLiteInfo and assigns it to the VoucherUseRuleLiteInfo field.
func (o *ActivityLiteInfo) SetVoucherUseRuleLiteInfo(v CommonVoucherUseRuleLiteInfo) {
	o.VoucherUseRuleLiteInfo = &v
}

func (o ActivityLiteInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityLiteInfo) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
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
	if !IsNil(o.VoucherDisplayLiteInfo) {
		toSerialize["voucher_display_lite_info"] = o.VoucherDisplayLiteInfo
	}
	if !IsNil(o.VoucherDisplayPatternInfo) {
		toSerialize["voucher_display_pattern_info"] = o.VoucherDisplayPatternInfo
	}
	if !IsNil(o.VoucherSendModeInfo) {
		toSerialize["voucher_send_mode_info"] = o.VoucherSendModeInfo
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

type NullableActivityLiteInfo struct {
	value *ActivityLiteInfo
	isSet bool
}

func (v NullableActivityLiteInfo) Get() *ActivityLiteInfo {
	return v.value
}

func (v *NullableActivityLiteInfo) Set(val *ActivityLiteInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityLiteInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityLiteInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityLiteInfo(val *ActivityLiteInfo) *NullableActivityLiteInfo {
	return &NullableActivityLiteInfo{value: val, isSet: true}
}

func (v NullableActivityLiteInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityLiteInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
