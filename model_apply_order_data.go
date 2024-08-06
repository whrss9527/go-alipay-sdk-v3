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

// checks if the ApplyOrderData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplyOrderData{}

// ApplyOrderData struct for ApplyOrderData
type ApplyOrderData struct {
	// 申请类型 BASE：基础信息， BRAND_BOX：品牌直达，SERVICE_BOX服务直达
	AccessType *string `json:"access_type,omitempty"`
	// 小程序名称
	AppName *string `json:"app_name,omitempty"`
	// 小程序id
	Appid *string `json:"appid,omitempty"`
	// 审核原因
	AuditReason *string `json:"audit_reason,omitempty"`
	// 品牌的模板id ONE_WITH_TWO：一拖二、DEFAULT：一拖四
	BrandTemplateId *string `json:"brand_template_id,omitempty"`
	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`
	// 标识老工单
	IsOldData *bool `json:"is_old_data,omitempty"`
	// 是不是最新的工单
	Latest *bool `json:"latest,omitempty"`
	// 服务主状态,INITIAL:初始化；VALID：已生效；INVALID：已失效
	MajorStatus *string `json:"major_status,omitempty"`
	// 申请单ID
	OrderId *string `json:"order_id,omitempty"`
	// 审核提效部分剔除数据的信息
	PartAgreeInfo []SearchPartAgreeInfo `json:"part_agree_info,omitempty"`
	// 场景编码
	SceneCode *string `json:"scene_code,omitempty"`
	// 场景名称
	SceneName *string `json:"scene_name,omitempty"`
	// 服务码
	ServiceCode *string `json:"service_code,omitempty"`
	// 申请状态 AGREE、REJECT、EDIT、AUDIT
	Status *string `json:"status,omitempty"`
	// 子服务名称
	SubServiceName *string `json:"sub_service_name,omitempty"`
}

// NewApplyOrderData instantiates a new ApplyOrderData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyOrderData() *ApplyOrderData {
	this := ApplyOrderData{}
	return &this
}

// NewApplyOrderDataWithDefaults instantiates a new ApplyOrderData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyOrderDataWithDefaults() *ApplyOrderData {
	this := ApplyOrderData{}
	return &this
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *ApplyOrderData) GetAccessType() string {
	if o == nil || IsNil(o.AccessType) {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOrderData) GetAccessTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *ApplyOrderData) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *ApplyOrderData) SetAccessType(v string) {
	o.AccessType = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *ApplyOrderData) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOrderData) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *ApplyOrderData) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *ApplyOrderData) SetAppName(v string) {
	o.AppName = &v
}

// GetAppid returns the Appid field value if set, zero value otherwise.
func (o *ApplyOrderData) GetAppid() string {
	if o == nil || IsNil(o.Appid) {
		var ret string
		return ret
	}
	return *o.Appid
}

// GetAppidOk returns a tuple with the Appid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOrderData) GetAppidOk() (*string, bool) {
	if o == nil || IsNil(o.Appid) {
		return nil, false
	}
	return o.Appid, true
}

// HasAppid returns a boolean if a field has been set.
func (o *ApplyOrderData) HasAppid() bool {
	if o != nil && !IsNil(o.Appid) {
		return true
	}

	return false
}

// SetAppid gets a reference to the given string and assigns it to the Appid field.
func (o *ApplyOrderData) SetAppid(v string) {
	o.Appid = &v
}

// GetAuditReason returns the AuditReason field value if set, zero value otherwise.
func (o *ApplyOrderData) GetAuditReason() string {
	if o == nil || IsNil(o.AuditReason) {
		var ret string
		return ret
	}
	return *o.AuditReason
}

// GetAuditReasonOk returns a tuple with the AuditReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOrderData) GetAuditReasonOk() (*string, bool) {
	if o == nil || IsNil(o.AuditReason) {
		return nil, false
	}
	return o.AuditReason, true
}

// HasAuditReason returns a boolean if a field has been set.
func (o *ApplyOrderData) HasAuditReason() bool {
	if o != nil && !IsNil(o.AuditReason) {
		return true
	}

	return false
}

// SetAuditReason gets a reference to the given string and assigns it to the AuditReason field.
func (o *ApplyOrderData) SetAuditReason(v string) {
	o.AuditReason = &v
}

// GetBrandTemplateId returns the BrandTemplateId field value if set, zero value otherwise.
func (o *ApplyOrderData) GetBrandTemplateId() string {
	if o == nil || IsNil(o.BrandTemplateId) {
		var ret string
		return ret
	}
	return *o.BrandTemplateId
}

// GetBrandTemplateIdOk returns a tuple with the BrandTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOrderData) GetBrandTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrandTemplateId) {
		return nil, false
	}
	return o.BrandTemplateId, true
}

// HasBrandTemplateId returns a boolean if a field has been set.
func (o *ApplyOrderData) HasBrandTemplateId() bool {
	if o != nil && !IsNil(o.BrandTemplateId) {
		return true
	}

	return false
}

// SetBrandTemplateId gets a reference to the given string and assigns it to the BrandTemplateId field.
func (o *ApplyOrderData) SetBrandTemplateId(v string) {
	o.BrandTemplateId = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *ApplyOrderData) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOrderData) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *ApplyOrderData) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *ApplyOrderData) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetIsOldData returns the IsOldData field value if set, zero value otherwise.
func (o *ApplyOrderData) GetIsOldData() bool {
	if o == nil || IsNil(o.IsOldData) {
		var ret bool
		return ret
	}
	return *o.IsOldData
}

// GetIsOldDataOk returns a tuple with the IsOldData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOrderData) GetIsOldDataOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOldData) {
		return nil, false
	}
	return o.IsOldData, true
}

// HasIsOldData returns a boolean if a field has been set.
func (o *ApplyOrderData) HasIsOldData() bool {
	if o != nil && !IsNil(o.IsOldData) {
		return true
	}

	return false
}

// SetIsOldData gets a reference to the given bool and assigns it to the IsOldData field.
func (o *ApplyOrderData) SetIsOldData(v bool) {
	o.IsOldData = &v
}

// GetLatest returns the Latest field value if set, zero value otherwise.
func (o *ApplyOrderData) GetLatest() bool {
	if o == nil || IsNil(o.Latest) {
		var ret bool
		return ret
	}
	return *o.Latest
}

// GetLatestOk returns a tuple with the Latest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOrderData) GetLatestOk() (*bool, bool) {
	if o == nil || IsNil(o.Latest) {
		return nil, false
	}
	return o.Latest, true
}

// HasLatest returns a boolean if a field has been set.
func (o *ApplyOrderData) HasLatest() bool {
	if o != nil && !IsNil(o.Latest) {
		return true
	}

	return false
}

// SetLatest gets a reference to the given bool and assigns it to the Latest field.
func (o *ApplyOrderData) SetLatest(v bool) {
	o.Latest = &v
}

// GetMajorStatus returns the MajorStatus field value if set, zero value otherwise.
func (o *ApplyOrderData) GetMajorStatus() string {
	if o == nil || IsNil(o.MajorStatus) {
		var ret string
		return ret
	}
	return *o.MajorStatus
}

// GetMajorStatusOk returns a tuple with the MajorStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOrderData) GetMajorStatusOk() (*string, bool) {
	if o == nil || IsNil(o.MajorStatus) {
		return nil, false
	}
	return o.MajorStatus, true
}

// HasMajorStatus returns a boolean if a field has been set.
func (o *ApplyOrderData) HasMajorStatus() bool {
	if o != nil && !IsNil(o.MajorStatus) {
		return true
	}

	return false
}

// SetMajorStatus gets a reference to the given string and assigns it to the MajorStatus field.
func (o *ApplyOrderData) SetMajorStatus(v string) {
	o.MajorStatus = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *ApplyOrderData) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOrderData) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *ApplyOrderData) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *ApplyOrderData) SetOrderId(v string) {
	o.OrderId = &v
}

// GetPartAgreeInfo returns the PartAgreeInfo field value if set, zero value otherwise.
func (o *ApplyOrderData) GetPartAgreeInfo() []SearchPartAgreeInfo {
	if o == nil || IsNil(o.PartAgreeInfo) {
		var ret []SearchPartAgreeInfo
		return ret
	}
	return o.PartAgreeInfo
}

// GetPartAgreeInfoOk returns a tuple with the PartAgreeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOrderData) GetPartAgreeInfoOk() ([]SearchPartAgreeInfo, bool) {
	if o == nil || IsNil(o.PartAgreeInfo) {
		return nil, false
	}
	return o.PartAgreeInfo, true
}

// HasPartAgreeInfo returns a boolean if a field has been set.
func (o *ApplyOrderData) HasPartAgreeInfo() bool {
	if o != nil && !IsNil(o.PartAgreeInfo) {
		return true
	}

	return false
}

// SetPartAgreeInfo gets a reference to the given []SearchPartAgreeInfo and assigns it to the PartAgreeInfo field.
func (o *ApplyOrderData) SetPartAgreeInfo(v []SearchPartAgreeInfo) {
	o.PartAgreeInfo = v
}

// GetSceneCode returns the SceneCode field value if set, zero value otherwise.
func (o *ApplyOrderData) GetSceneCode() string {
	if o == nil || IsNil(o.SceneCode) {
		var ret string
		return ret
	}
	return *o.SceneCode
}

// GetSceneCodeOk returns a tuple with the SceneCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOrderData) GetSceneCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SceneCode) {
		return nil, false
	}
	return o.SceneCode, true
}

// HasSceneCode returns a boolean if a field has been set.
func (o *ApplyOrderData) HasSceneCode() bool {
	if o != nil && !IsNil(o.SceneCode) {
		return true
	}

	return false
}

// SetSceneCode gets a reference to the given string and assigns it to the SceneCode field.
func (o *ApplyOrderData) SetSceneCode(v string) {
	o.SceneCode = &v
}

// GetSceneName returns the SceneName field value if set, zero value otherwise.
func (o *ApplyOrderData) GetSceneName() string {
	if o == nil || IsNil(o.SceneName) {
		var ret string
		return ret
	}
	return *o.SceneName
}

// GetSceneNameOk returns a tuple with the SceneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOrderData) GetSceneNameOk() (*string, bool) {
	if o == nil || IsNil(o.SceneName) {
		return nil, false
	}
	return o.SceneName, true
}

// HasSceneName returns a boolean if a field has been set.
func (o *ApplyOrderData) HasSceneName() bool {
	if o != nil && !IsNil(o.SceneName) {
		return true
	}

	return false
}

// SetSceneName gets a reference to the given string and assigns it to the SceneName field.
func (o *ApplyOrderData) SetSceneName(v string) {
	o.SceneName = &v
}

// GetServiceCode returns the ServiceCode field value if set, zero value otherwise.
func (o *ApplyOrderData) GetServiceCode() string {
	if o == nil || IsNil(o.ServiceCode) {
		var ret string
		return ret
	}
	return *o.ServiceCode
}

// GetServiceCodeOk returns a tuple with the ServiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOrderData) GetServiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceCode) {
		return nil, false
	}
	return o.ServiceCode, true
}

// HasServiceCode returns a boolean if a field has been set.
func (o *ApplyOrderData) HasServiceCode() bool {
	if o != nil && !IsNil(o.ServiceCode) {
		return true
	}

	return false
}

// SetServiceCode gets a reference to the given string and assigns it to the ServiceCode field.
func (o *ApplyOrderData) SetServiceCode(v string) {
	o.ServiceCode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplyOrderData) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOrderData) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplyOrderData) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplyOrderData) SetStatus(v string) {
	o.Status = &v
}

// GetSubServiceName returns the SubServiceName field value if set, zero value otherwise.
func (o *ApplyOrderData) GetSubServiceName() string {
	if o == nil || IsNil(o.SubServiceName) {
		var ret string
		return ret
	}
	return *o.SubServiceName
}

// GetSubServiceNameOk returns a tuple with the SubServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOrderData) GetSubServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubServiceName) {
		return nil, false
	}
	return o.SubServiceName, true
}

// HasSubServiceName returns a boolean if a field has been set.
func (o *ApplyOrderData) HasSubServiceName() bool {
	if o != nil && !IsNil(o.SubServiceName) {
		return true
	}

	return false
}

// SetSubServiceName gets a reference to the given string and assigns it to the SubServiceName field.
func (o *ApplyOrderData) SetSubServiceName(v string) {
	o.SubServiceName = &v
}

func (o ApplyOrderData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyOrderData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessType) {
		toSerialize["access_type"] = o.AccessType
	}
	if !IsNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if !IsNil(o.Appid) {
		toSerialize["appid"] = o.Appid
	}
	if !IsNil(o.AuditReason) {
		toSerialize["audit_reason"] = o.AuditReason
	}
	if !IsNil(o.BrandTemplateId) {
		toSerialize["brand_template_id"] = o.BrandTemplateId
	}
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
	}
	if !IsNil(o.IsOldData) {
		toSerialize["is_old_data"] = o.IsOldData
	}
	if !IsNil(o.Latest) {
		toSerialize["latest"] = o.Latest
	}
	if !IsNil(o.MajorStatus) {
		toSerialize["major_status"] = o.MajorStatus
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.PartAgreeInfo) {
		toSerialize["part_agree_info"] = o.PartAgreeInfo
	}
	if !IsNil(o.SceneCode) {
		toSerialize["scene_code"] = o.SceneCode
	}
	if !IsNil(o.SceneName) {
		toSerialize["scene_name"] = o.SceneName
	}
	if !IsNil(o.ServiceCode) {
		toSerialize["service_code"] = o.ServiceCode
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubServiceName) {
		toSerialize["sub_service_name"] = o.SubServiceName
	}
	return toSerialize, nil
}

type NullableApplyOrderData struct {
	value *ApplyOrderData
	isSet bool
}

func (v NullableApplyOrderData) Get() *ApplyOrderData {
	return v.value
}

func (v *NullableApplyOrderData) Set(val *ApplyOrderData) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyOrderData) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyOrderData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyOrderData(val *ApplyOrderData) *NullableApplyOrderData {
	return &NullableApplyOrderData{value: val, isSet: true}
}

func (v NullableApplyOrderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyOrderData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
