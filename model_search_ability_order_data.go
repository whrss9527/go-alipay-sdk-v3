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

// checks if the SearchAbilityOrderData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchAbilityOrderData{}

// SearchAbilityOrderData struct for SearchAbilityOrderData
type SearchAbilityOrderData struct {
	// 申请单类型
	AccessType *string `json:"access_type,omitempty"`
	// 小程序名称
	AppName *string `json:"app_name,omitempty"`
	// 小程序状态 ON_LINE、OFF_LINE
	AppStatus *string `json:"app_status,omitempty"`
	// 小程序id
	Appid *string `json:"appid,omitempty"`
	// 申请单id
	ApplyId *string `json:"apply_id,omitempty"`
	// 申请类型 BASE：基础信息， BRAND_BOX：品牌直达，SERVICE_BOX服务直达
	ApplyType *string `json:"apply_type,omitempty"`
	// 申请状态 AGREE、REJECT、EDIT、AUDIT
	AuditStatus *string `json:"audit_status,omitempty"`
	// bizId 为品牌box的boxID
	BizId *string `json:"biz_id,omitempty"`
	// 品牌box的上下架状态 ONLINE OFFLINE
	BoxStatus *string `json:"box_status,omitempty"`
	// 品牌的模板id ONE_WITH_TWO：一拖二、DEFAULT：一拖四
	BrandTemplateId *string `json:"brand_template_id,omitempty"`
	// 二级服务信息
	Children []SearchAbilityOrderInfoOpenApi `json:"children,omitempty"`
	// 二级服务的唯一标识
	DataKey *string `json:"data_key,omitempty"`
	// 更新时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// 服务的唯一标识（优先使用serviceCode可忽略）
	Id *string `json:"id,omitempty"`
	// 是否是老工单
	IsOldData *string `json:"is_old_data,omitempty"`
	// 服务主状态,INITIAL:初始化；VALID：已生效；INVALID：已失效
	MajorStatus *string `json:"major_status,omitempty"`
	// 上架时间
	OnlineTime *string `json:"online_time,omitempty"`
	// 服务的可见性状态
	OpenStatus *bool `json:"open_status,omitempty"`
	// 下架操作者
	Operator *string `json:"operator,omitempty"`
	// 驳回原因
	RejectReason *string `json:"reject_reason,omitempty"`
	// 场景码
	SceneCode *string `json:"scene_code,omitempty"`
	// 场景名称
	SceneName *string `json:"scene_name,omitempty"`
	// 服务码
	ServiceCode *string `json:"service_code,omitempty"`
	// 子功能描述
	SubServiceDesc *string `json:"sub_service_desc,omitempty"`
	// 子功能名称
	SubServiceName *string `json:"sub_service_name,omitempty"`
}

// NewSearchAbilityOrderData instantiates a new SearchAbilityOrderData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchAbilityOrderData() *SearchAbilityOrderData {
	this := SearchAbilityOrderData{}
	return &this
}

// NewSearchAbilityOrderDataWithDefaults instantiates a new SearchAbilityOrderData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchAbilityOrderDataWithDefaults() *SearchAbilityOrderData {
	this := SearchAbilityOrderData{}
	return &this
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetAccessType() string {
	if o == nil || IsNil(o.AccessType) {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetAccessTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *SearchAbilityOrderData) SetAccessType(v string) {
	o.AccessType = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *SearchAbilityOrderData) SetAppName(v string) {
	o.AppName = &v
}

// GetAppStatus returns the AppStatus field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetAppStatus() string {
	if o == nil || IsNil(o.AppStatus) {
		var ret string
		return ret
	}
	return *o.AppStatus
}

// GetAppStatusOk returns a tuple with the AppStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetAppStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AppStatus) {
		return nil, false
	}
	return o.AppStatus, true
}

// HasAppStatus returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasAppStatus() bool {
	if o != nil && !IsNil(o.AppStatus) {
		return true
	}

	return false
}

// SetAppStatus gets a reference to the given string and assigns it to the AppStatus field.
func (o *SearchAbilityOrderData) SetAppStatus(v string) {
	o.AppStatus = &v
}

// GetAppid returns the Appid field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetAppid() string {
	if o == nil || IsNil(o.Appid) {
		var ret string
		return ret
	}
	return *o.Appid
}

// GetAppidOk returns a tuple with the Appid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetAppidOk() (*string, bool) {
	if o == nil || IsNil(o.Appid) {
		return nil, false
	}
	return o.Appid, true
}

// HasAppid returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasAppid() bool {
	if o != nil && !IsNil(o.Appid) {
		return true
	}

	return false
}

// SetAppid gets a reference to the given string and assigns it to the Appid field.
func (o *SearchAbilityOrderData) SetAppid(v string) {
	o.Appid = &v
}

// GetApplyId returns the ApplyId field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetApplyId() string {
	if o == nil || IsNil(o.ApplyId) {
		var ret string
		return ret
	}
	return *o.ApplyId
}

// GetApplyIdOk returns a tuple with the ApplyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetApplyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplyId) {
		return nil, false
	}
	return o.ApplyId, true
}

// HasApplyId returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasApplyId() bool {
	if o != nil && !IsNil(o.ApplyId) {
		return true
	}

	return false
}

// SetApplyId gets a reference to the given string and assigns it to the ApplyId field.
func (o *SearchAbilityOrderData) SetApplyId(v string) {
	o.ApplyId = &v
}

// GetApplyType returns the ApplyType field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetApplyType() string {
	if o == nil || IsNil(o.ApplyType) {
		var ret string
		return ret
	}
	return *o.ApplyType
}

// GetApplyTypeOk returns a tuple with the ApplyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetApplyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ApplyType) {
		return nil, false
	}
	return o.ApplyType, true
}

// HasApplyType returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasApplyType() bool {
	if o != nil && !IsNil(o.ApplyType) {
		return true
	}

	return false
}

// SetApplyType gets a reference to the given string and assigns it to the ApplyType field.
func (o *SearchAbilityOrderData) SetApplyType(v string) {
	o.ApplyType = &v
}

// GetAuditStatus returns the AuditStatus field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetAuditStatus() string {
	if o == nil || IsNil(o.AuditStatus) {
		var ret string
		return ret
	}
	return *o.AuditStatus
}

// GetAuditStatusOk returns a tuple with the AuditStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetAuditStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AuditStatus) {
		return nil, false
	}
	return o.AuditStatus, true
}

// HasAuditStatus returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasAuditStatus() bool {
	if o != nil && !IsNil(o.AuditStatus) {
		return true
	}

	return false
}

// SetAuditStatus gets a reference to the given string and assigns it to the AuditStatus field.
func (o *SearchAbilityOrderData) SetAuditStatus(v string) {
	o.AuditStatus = &v
}

// GetBizId returns the BizId field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetBizId() string {
	if o == nil || IsNil(o.BizId) {
		var ret string
		return ret
	}
	return *o.BizId
}

// GetBizIdOk returns a tuple with the BizId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetBizIdOk() (*string, bool) {
	if o == nil || IsNil(o.BizId) {
		return nil, false
	}
	return o.BizId, true
}

// HasBizId returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasBizId() bool {
	if o != nil && !IsNil(o.BizId) {
		return true
	}

	return false
}

// SetBizId gets a reference to the given string and assigns it to the BizId field.
func (o *SearchAbilityOrderData) SetBizId(v string) {
	o.BizId = &v
}

// GetBoxStatus returns the BoxStatus field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetBoxStatus() string {
	if o == nil || IsNil(o.BoxStatus) {
		var ret string
		return ret
	}
	return *o.BoxStatus
}

// GetBoxStatusOk returns a tuple with the BoxStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetBoxStatusOk() (*string, bool) {
	if o == nil || IsNil(o.BoxStatus) {
		return nil, false
	}
	return o.BoxStatus, true
}

// HasBoxStatus returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasBoxStatus() bool {
	if o != nil && !IsNil(o.BoxStatus) {
		return true
	}

	return false
}

// SetBoxStatus gets a reference to the given string and assigns it to the BoxStatus field.
func (o *SearchAbilityOrderData) SetBoxStatus(v string) {
	o.BoxStatus = &v
}

// GetBrandTemplateId returns the BrandTemplateId field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetBrandTemplateId() string {
	if o == nil || IsNil(o.BrandTemplateId) {
		var ret string
		return ret
	}
	return *o.BrandTemplateId
}

// GetBrandTemplateIdOk returns a tuple with the BrandTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetBrandTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrandTemplateId) {
		return nil, false
	}
	return o.BrandTemplateId, true
}

// HasBrandTemplateId returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasBrandTemplateId() bool {
	if o != nil && !IsNil(o.BrandTemplateId) {
		return true
	}

	return false
}

// SetBrandTemplateId gets a reference to the given string and assigns it to the BrandTemplateId field.
func (o *SearchAbilityOrderData) SetBrandTemplateId(v string) {
	o.BrandTemplateId = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetChildren() []SearchAbilityOrderInfoOpenApi {
	if o == nil || IsNil(o.Children) {
		var ret []SearchAbilityOrderInfoOpenApi
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetChildrenOk() ([]SearchAbilityOrderInfoOpenApi, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []SearchAbilityOrderInfoOpenApi and assigns it to the Children field.
func (o *SearchAbilityOrderData) SetChildren(v []SearchAbilityOrderInfoOpenApi) {
	o.Children = v
}

// GetDataKey returns the DataKey field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetDataKey() string {
	if o == nil || IsNil(o.DataKey) {
		var ret string
		return ret
	}
	return *o.DataKey
}

// GetDataKeyOk returns a tuple with the DataKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetDataKeyOk() (*string, bool) {
	if o == nil || IsNil(o.DataKey) {
		return nil, false
	}
	return o.DataKey, true
}

// HasDataKey returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasDataKey() bool {
	if o != nil && !IsNil(o.DataKey) {
		return true
	}

	return false
}

// SetDataKey gets a reference to the given string and assigns it to the DataKey field.
func (o *SearchAbilityOrderData) SetDataKey(v string) {
	o.DataKey = &v
}

// GetGmtModified returns the GmtModified field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetGmtModified() string {
	if o == nil || IsNil(o.GmtModified) {
		var ret string
		return ret
	}
	return *o.GmtModified
}

// GetGmtModifiedOk returns a tuple with the GmtModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetGmtModifiedOk() (*string, bool) {
	if o == nil || IsNil(o.GmtModified) {
		return nil, false
	}
	return o.GmtModified, true
}

// HasGmtModified returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasGmtModified() bool {
	if o != nil && !IsNil(o.GmtModified) {
		return true
	}

	return false
}

// SetGmtModified gets a reference to the given string and assigns it to the GmtModified field.
func (o *SearchAbilityOrderData) SetGmtModified(v string) {
	o.GmtModified = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SearchAbilityOrderData) SetId(v string) {
	o.Id = &v
}

// GetIsOldData returns the IsOldData field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetIsOldData() string {
	if o == nil || IsNil(o.IsOldData) {
		var ret string
		return ret
	}
	return *o.IsOldData
}

// GetIsOldDataOk returns a tuple with the IsOldData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetIsOldDataOk() (*string, bool) {
	if o == nil || IsNil(o.IsOldData) {
		return nil, false
	}
	return o.IsOldData, true
}

// HasIsOldData returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasIsOldData() bool {
	if o != nil && !IsNil(o.IsOldData) {
		return true
	}

	return false
}

// SetIsOldData gets a reference to the given string and assigns it to the IsOldData field.
func (o *SearchAbilityOrderData) SetIsOldData(v string) {
	o.IsOldData = &v
}

// GetMajorStatus returns the MajorStatus field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetMajorStatus() string {
	if o == nil || IsNil(o.MajorStatus) {
		var ret string
		return ret
	}
	return *o.MajorStatus
}

// GetMajorStatusOk returns a tuple with the MajorStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetMajorStatusOk() (*string, bool) {
	if o == nil || IsNil(o.MajorStatus) {
		return nil, false
	}
	return o.MajorStatus, true
}

// HasMajorStatus returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasMajorStatus() bool {
	if o != nil && !IsNil(o.MajorStatus) {
		return true
	}

	return false
}

// SetMajorStatus gets a reference to the given string and assigns it to the MajorStatus field.
func (o *SearchAbilityOrderData) SetMajorStatus(v string) {
	o.MajorStatus = &v
}

// GetOnlineTime returns the OnlineTime field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetOnlineTime() string {
	if o == nil || IsNil(o.OnlineTime) {
		var ret string
		return ret
	}
	return *o.OnlineTime
}

// GetOnlineTimeOk returns a tuple with the OnlineTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetOnlineTimeOk() (*string, bool) {
	if o == nil || IsNil(o.OnlineTime) {
		return nil, false
	}
	return o.OnlineTime, true
}

// HasOnlineTime returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasOnlineTime() bool {
	if o != nil && !IsNil(o.OnlineTime) {
		return true
	}

	return false
}

// SetOnlineTime gets a reference to the given string and assigns it to the OnlineTime field.
func (o *SearchAbilityOrderData) SetOnlineTime(v string) {
	o.OnlineTime = &v
}

// GetOpenStatus returns the OpenStatus field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetOpenStatus() bool {
	if o == nil || IsNil(o.OpenStatus) {
		var ret bool
		return ret
	}
	return *o.OpenStatus
}

// GetOpenStatusOk returns a tuple with the OpenStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetOpenStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.OpenStatus) {
		return nil, false
	}
	return o.OpenStatus, true
}

// HasOpenStatus returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasOpenStatus() bool {
	if o != nil && !IsNil(o.OpenStatus) {
		return true
	}

	return false
}

// SetOpenStatus gets a reference to the given bool and assigns it to the OpenStatus field.
func (o *SearchAbilityOrderData) SetOpenStatus(v bool) {
	o.OpenStatus = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *SearchAbilityOrderData) SetOperator(v string) {
	o.Operator = &v
}

// GetRejectReason returns the RejectReason field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetRejectReason() string {
	if o == nil || IsNil(o.RejectReason) {
		var ret string
		return ret
	}
	return *o.RejectReason
}

// GetRejectReasonOk returns a tuple with the RejectReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetRejectReasonOk() (*string, bool) {
	if o == nil || IsNil(o.RejectReason) {
		return nil, false
	}
	return o.RejectReason, true
}

// HasRejectReason returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasRejectReason() bool {
	if o != nil && !IsNil(o.RejectReason) {
		return true
	}

	return false
}

// SetRejectReason gets a reference to the given string and assigns it to the RejectReason field.
func (o *SearchAbilityOrderData) SetRejectReason(v string) {
	o.RejectReason = &v
}

// GetSceneCode returns the SceneCode field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetSceneCode() string {
	if o == nil || IsNil(o.SceneCode) {
		var ret string
		return ret
	}
	return *o.SceneCode
}

// GetSceneCodeOk returns a tuple with the SceneCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetSceneCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SceneCode) {
		return nil, false
	}
	return o.SceneCode, true
}

// HasSceneCode returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasSceneCode() bool {
	if o != nil && !IsNil(o.SceneCode) {
		return true
	}

	return false
}

// SetSceneCode gets a reference to the given string and assigns it to the SceneCode field.
func (o *SearchAbilityOrderData) SetSceneCode(v string) {
	o.SceneCode = &v
}

// GetSceneName returns the SceneName field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetSceneName() string {
	if o == nil || IsNil(o.SceneName) {
		var ret string
		return ret
	}
	return *o.SceneName
}

// GetSceneNameOk returns a tuple with the SceneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetSceneNameOk() (*string, bool) {
	if o == nil || IsNil(o.SceneName) {
		return nil, false
	}
	return o.SceneName, true
}

// HasSceneName returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasSceneName() bool {
	if o != nil && !IsNil(o.SceneName) {
		return true
	}

	return false
}

// SetSceneName gets a reference to the given string and assigns it to the SceneName field.
func (o *SearchAbilityOrderData) SetSceneName(v string) {
	o.SceneName = &v
}

// GetServiceCode returns the ServiceCode field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetServiceCode() string {
	if o == nil || IsNil(o.ServiceCode) {
		var ret string
		return ret
	}
	return *o.ServiceCode
}

// GetServiceCodeOk returns a tuple with the ServiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetServiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceCode) {
		return nil, false
	}
	return o.ServiceCode, true
}

// HasServiceCode returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasServiceCode() bool {
	if o != nil && !IsNil(o.ServiceCode) {
		return true
	}

	return false
}

// SetServiceCode gets a reference to the given string and assigns it to the ServiceCode field.
func (o *SearchAbilityOrderData) SetServiceCode(v string) {
	o.ServiceCode = &v
}

// GetSubServiceDesc returns the SubServiceDesc field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetSubServiceDesc() string {
	if o == nil || IsNil(o.SubServiceDesc) {
		var ret string
		return ret
	}
	return *o.SubServiceDesc
}

// GetSubServiceDescOk returns a tuple with the SubServiceDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetSubServiceDescOk() (*string, bool) {
	if o == nil || IsNil(o.SubServiceDesc) {
		return nil, false
	}
	return o.SubServiceDesc, true
}

// HasSubServiceDesc returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasSubServiceDesc() bool {
	if o != nil && !IsNil(o.SubServiceDesc) {
		return true
	}

	return false
}

// SetSubServiceDesc gets a reference to the given string and assigns it to the SubServiceDesc field.
func (o *SearchAbilityOrderData) SetSubServiceDesc(v string) {
	o.SubServiceDesc = &v
}

// GetSubServiceName returns the SubServiceName field value if set, zero value otherwise.
func (o *SearchAbilityOrderData) GetSubServiceName() string {
	if o == nil || IsNil(o.SubServiceName) {
		var ret string
		return ret
	}
	return *o.SubServiceName
}

// GetSubServiceNameOk returns a tuple with the SubServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchAbilityOrderData) GetSubServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubServiceName) {
		return nil, false
	}
	return o.SubServiceName, true
}

// HasSubServiceName returns a boolean if a field has been set.
func (o *SearchAbilityOrderData) HasSubServiceName() bool {
	if o != nil && !IsNil(o.SubServiceName) {
		return true
	}

	return false
}

// SetSubServiceName gets a reference to the given string and assigns it to the SubServiceName field.
func (o *SearchAbilityOrderData) SetSubServiceName(v string) {
	o.SubServiceName = &v
}

func (o SearchAbilityOrderData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchAbilityOrderData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessType) {
		toSerialize["access_type"] = o.AccessType
	}
	if !IsNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if !IsNil(o.AppStatus) {
		toSerialize["app_status"] = o.AppStatus
	}
	if !IsNil(o.Appid) {
		toSerialize["appid"] = o.Appid
	}
	if !IsNil(o.ApplyId) {
		toSerialize["apply_id"] = o.ApplyId
	}
	if !IsNil(o.ApplyType) {
		toSerialize["apply_type"] = o.ApplyType
	}
	if !IsNil(o.AuditStatus) {
		toSerialize["audit_status"] = o.AuditStatus
	}
	if !IsNil(o.BizId) {
		toSerialize["biz_id"] = o.BizId
	}
	if !IsNil(o.BoxStatus) {
		toSerialize["box_status"] = o.BoxStatus
	}
	if !IsNil(o.BrandTemplateId) {
		toSerialize["brand_template_id"] = o.BrandTemplateId
	}
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	if !IsNil(o.DataKey) {
		toSerialize["data_key"] = o.DataKey
	}
	if !IsNil(o.GmtModified) {
		toSerialize["gmt_modified"] = o.GmtModified
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsOldData) {
		toSerialize["is_old_data"] = o.IsOldData
	}
	if !IsNil(o.MajorStatus) {
		toSerialize["major_status"] = o.MajorStatus
	}
	if !IsNil(o.OnlineTime) {
		toSerialize["online_time"] = o.OnlineTime
	}
	if !IsNil(o.OpenStatus) {
		toSerialize["open_status"] = o.OpenStatus
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.RejectReason) {
		toSerialize["reject_reason"] = o.RejectReason
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
	if !IsNil(o.SubServiceDesc) {
		toSerialize["sub_service_desc"] = o.SubServiceDesc
	}
	if !IsNil(o.SubServiceName) {
		toSerialize["sub_service_name"] = o.SubServiceName
	}
	return toSerialize, nil
}

type NullableSearchAbilityOrderData struct {
	value *SearchAbilityOrderData
	isSet bool
}

func (v NullableSearchAbilityOrderData) Get() *SearchAbilityOrderData {
	return v.value
}

func (v *NullableSearchAbilityOrderData) Set(val *SearchAbilityOrderData) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchAbilityOrderData) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchAbilityOrderData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchAbilityOrderData(val *SearchAbilityOrderData) *NullableSearchAbilityOrderData {
	return &NullableSearchAbilityOrderData{value: val, isSet: true}
}

func (v NullableSearchAbilityOrderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchAbilityOrderData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

