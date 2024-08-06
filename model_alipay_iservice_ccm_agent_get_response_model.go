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

// checks if the AlipayIserviceCcmAgentGetResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayIserviceCcmAgentGetResponseModel{}

// AlipayIserviceCcmAgentGetResponseModel struct for AlipayIserviceCcmAgentGetResponseModel
type AlipayIserviceCcmAgentGetResponseModel struct {
	// 热线接入方式： 0：话机 1：电脑耳机；2：webrtc
	AnsweringMode *string `json:"answering_mode,omitempty"`
	// 头像的oss file key
	Avatar *string `json:"avatar,omitempty"`
	// 客服关联的数据权限id列表
	CcsInstanceIds []string `json:"ccs_instance_ids,omitempty"`
	// 客服在线信息
	ChatConfig []AgentChatInfo `json:"chat_config,omitempty"`
	// 创建时间，采用UTC时间，按照ISO8601标准表示，格式为：yyyy-MM-dd'T'HH:mm:ss'Z
	CreateTime *string `json:"create_time,omitempty"`
	// 创建人id
	CreatorId *string `json:"creator_id,omitempty"`
	// 钉钉user_id
	DingtalkUserId *string `json:"dingtalk_user_id,omitempty"`
	// 邮箱
	Email *string `json:"email,omitempty"`
	// 外部系统用户id,比如：金融云的用户id
	ExternalId *string `json:"external_id,omitempty"`
	// 客服热线信息
	HotlineConfig []AgentHotlineInfo `json:"hotline_config,omitempty"`
	// 客服id
	Id *string `json:"id,omitempty"`
	// 工号
	JobNumber *string `json:"job_number,omitempty"`
	// 客服上次登录时间 采用UTC时间，按照ISO8601标准表示，格式为：yyyy-MM-dd'T'HH:mm:ss'Z'
	LastLoginTime *string `json:"last_login_time,omitempty"`
	// 手机号
	Mobile *string `json:"mobile,omitempty"`
	// 客服昵称
	NickName *string `json:"nick_name,omitempty"`
	// 个人简介
	Profile *string `json:"profile,omitempty"`
	// 客服姓名
	RealName *string `json:"real_name,omitempty"`
	// 角色id列表
	RoleIds []RoleId `json:"role_ids,omitempty"`
	// 客服状态：NORMAL，DELETE
	Status *string `json:"status,omitempty"`
	// 用户类型：NORMAL（普通客服），ADMIN（超级管理员，不能被删除）
	Type *string `json:"type,omitempty"`
	// 最后修改时间，采用UTC时间，按照ISO8601标准表示，格式为：yyyy-MM-dd'T'HH:mm:ss'Z'
	UpdateTime *string `json:"update_time,omitempty"`
	// 最后修改人id
	UpdaterId *string `json:"updater_id,omitempty"`
}

// NewAlipayIserviceCcmAgentGetResponseModel instantiates a new AlipayIserviceCcmAgentGetResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayIserviceCcmAgentGetResponseModel() *AlipayIserviceCcmAgentGetResponseModel {
	this := AlipayIserviceCcmAgentGetResponseModel{}
	return &this
}

// NewAlipayIserviceCcmAgentGetResponseModelWithDefaults instantiates a new AlipayIserviceCcmAgentGetResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayIserviceCcmAgentGetResponseModelWithDefaults() *AlipayIserviceCcmAgentGetResponseModel {
	this := AlipayIserviceCcmAgentGetResponseModel{}
	return &this
}

// GetAnsweringMode returns the AnsweringMode field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetAnsweringMode() string {
	if o == nil || IsNil(o.AnsweringMode) {
		var ret string
		return ret
	}
	return *o.AnsweringMode
}

// GetAnsweringModeOk returns a tuple with the AnsweringMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetAnsweringModeOk() (*string, bool) {
	if o == nil || IsNil(o.AnsweringMode) {
		return nil, false
	}
	return o.AnsweringMode, true
}

// HasAnsweringMode returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasAnsweringMode() bool {
	if o != nil && !IsNil(o.AnsweringMode) {
		return true
	}

	return false
}

// SetAnsweringMode gets a reference to the given string and assigns it to the AnsweringMode field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetAnsweringMode(v string) {
	o.AnsweringMode = &v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetAvatar() string {
	if o == nil || IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetAvatar(v string) {
	o.Avatar = &v
}

// GetCcsInstanceIds returns the CcsInstanceIds field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetCcsInstanceIds() []string {
	if o == nil || IsNil(o.CcsInstanceIds) {
		var ret []string
		return ret
	}
	return o.CcsInstanceIds
}

// GetCcsInstanceIdsOk returns a tuple with the CcsInstanceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetCcsInstanceIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CcsInstanceIds) {
		return nil, false
	}
	return o.CcsInstanceIds, true
}

// HasCcsInstanceIds returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasCcsInstanceIds() bool {
	if o != nil && !IsNil(o.CcsInstanceIds) {
		return true
	}

	return false
}

// SetCcsInstanceIds gets a reference to the given []string and assigns it to the CcsInstanceIds field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetCcsInstanceIds(v []string) {
	o.CcsInstanceIds = v
}

// GetChatConfig returns the ChatConfig field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetChatConfig() []AgentChatInfo {
	if o == nil || IsNil(o.ChatConfig) {
		var ret []AgentChatInfo
		return ret
	}
	return o.ChatConfig
}

// GetChatConfigOk returns a tuple with the ChatConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetChatConfigOk() ([]AgentChatInfo, bool) {
	if o == nil || IsNil(o.ChatConfig) {
		return nil, false
	}
	return o.ChatConfig, true
}

// HasChatConfig returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasChatConfig() bool {
	if o != nil && !IsNil(o.ChatConfig) {
		return true
	}

	return false
}

// SetChatConfig gets a reference to the given []AgentChatInfo and assigns it to the ChatConfig field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetChatConfig(v []AgentChatInfo) {
	o.ChatConfig = v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetCreateTime() string {
	if o == nil || IsNil(o.CreateTime) {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetCreateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetDingtalkUserId returns the DingtalkUserId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetDingtalkUserId() string {
	if o == nil || IsNil(o.DingtalkUserId) {
		var ret string
		return ret
	}
	return *o.DingtalkUserId
}

// GetDingtalkUserIdOk returns a tuple with the DingtalkUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetDingtalkUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.DingtalkUserId) {
		return nil, false
	}
	return o.DingtalkUserId, true
}

// HasDingtalkUserId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasDingtalkUserId() bool {
	if o != nil && !IsNil(o.DingtalkUserId) {
		return true
	}

	return false
}

// SetDingtalkUserId gets a reference to the given string and assigns it to the DingtalkUserId field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetDingtalkUserId(v string) {
	o.DingtalkUserId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetEmail(v string) {
	o.Email = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetHotlineConfig returns the HotlineConfig field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetHotlineConfig() []AgentHotlineInfo {
	if o == nil || IsNil(o.HotlineConfig) {
		var ret []AgentHotlineInfo
		return ret
	}
	return o.HotlineConfig
}

// GetHotlineConfigOk returns a tuple with the HotlineConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetHotlineConfigOk() ([]AgentHotlineInfo, bool) {
	if o == nil || IsNil(o.HotlineConfig) {
		return nil, false
	}
	return o.HotlineConfig, true
}

// HasHotlineConfig returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasHotlineConfig() bool {
	if o != nil && !IsNil(o.HotlineConfig) {
		return true
	}

	return false
}

// SetHotlineConfig gets a reference to the given []AgentHotlineInfo and assigns it to the HotlineConfig field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetHotlineConfig(v []AgentHotlineInfo) {
	o.HotlineConfig = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetId(v string) {
	o.Id = &v
}

// GetJobNumber returns the JobNumber field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetJobNumber() string {
	if o == nil || IsNil(o.JobNumber) {
		var ret string
		return ret
	}
	return *o.JobNumber
}

// GetJobNumberOk returns a tuple with the JobNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetJobNumberOk() (*string, bool) {
	if o == nil || IsNil(o.JobNumber) {
		return nil, false
	}
	return o.JobNumber, true
}

// HasJobNumber returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasJobNumber() bool {
	if o != nil && !IsNil(o.JobNumber) {
		return true
	}

	return false
}

// SetJobNumber gets a reference to the given string and assigns it to the JobNumber field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetJobNumber(v string) {
	o.JobNumber = &v
}

// GetLastLoginTime returns the LastLoginTime field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetLastLoginTime() string {
	if o == nil || IsNil(o.LastLoginTime) {
		var ret string
		return ret
	}
	return *o.LastLoginTime
}

// GetLastLoginTimeOk returns a tuple with the LastLoginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetLastLoginTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastLoginTime) {
		return nil, false
	}
	return o.LastLoginTime, true
}

// HasLastLoginTime returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasLastLoginTime() bool {
	if o != nil && !IsNil(o.LastLoginTime) {
		return true
	}

	return false
}

// SetLastLoginTime gets a reference to the given string and assigns it to the LastLoginTime field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetLastLoginTime(v string) {
	o.LastLoginTime = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetMobile() string {
	if o == nil || IsNil(o.Mobile) {
		var ret string
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetMobileOk() (*string, bool) {
	if o == nil || IsNil(o.Mobile) {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasMobile() bool {
	if o != nil && !IsNil(o.Mobile) {
		return true
	}

	return false
}

// SetMobile gets a reference to the given string and assigns it to the Mobile field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetMobile(v string) {
	o.Mobile = &v
}

// GetNickName returns the NickName field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetNickName() string {
	if o == nil || IsNil(o.NickName) {
		var ret string
		return ret
	}
	return *o.NickName
}

// GetNickNameOk returns a tuple with the NickName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetNickNameOk() (*string, bool) {
	if o == nil || IsNil(o.NickName) {
		return nil, false
	}
	return o.NickName, true
}

// HasNickName returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasNickName() bool {
	if o != nil && !IsNil(o.NickName) {
		return true
	}

	return false
}

// SetNickName gets a reference to the given string and assigns it to the NickName field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetNickName(v string) {
	o.NickName = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetProfile() string {
	if o == nil || IsNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetProfileOk() (*string, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetProfile(v string) {
	o.Profile = &v
}

// GetRealName returns the RealName field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetRealName() string {
	if o == nil || IsNil(o.RealName) {
		var ret string
		return ret
	}
	return *o.RealName
}

// GetRealNameOk returns a tuple with the RealName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetRealNameOk() (*string, bool) {
	if o == nil || IsNil(o.RealName) {
		return nil, false
	}
	return o.RealName, true
}

// HasRealName returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasRealName() bool {
	if o != nil && !IsNil(o.RealName) {
		return true
	}

	return false
}

// SetRealName gets a reference to the given string and assigns it to the RealName field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetRealName(v string) {
	o.RealName = &v
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetRoleIds() []RoleId {
	if o == nil || IsNil(o.RoleIds) {
		var ret []RoleId
		return ret
	}
	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetRoleIdsOk() ([]RoleId, bool) {
	if o == nil || IsNil(o.RoleIds) {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasRoleIds() bool {
	if o != nil && !IsNil(o.RoleIds) {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []RoleId and assigns it to the RoleIds field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetRoleIds(v []RoleId) {
	o.RoleIds = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetType(v string) {
	o.Type = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetUpdateTime() string {
	if o == nil || IsNil(o.UpdateTime) {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetUpdateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

// GetUpdaterId returns the UpdaterId field value if set, zero value otherwise.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetUpdaterId() string {
	if o == nil || IsNil(o.UpdaterId) {
		var ret string
		return ret
	}
	return *o.UpdaterId
}

// GetUpdaterIdOk returns a tuple with the UpdaterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) GetUpdaterIdOk() (*string, bool) {
	if o == nil || IsNil(o.UpdaterId) {
		return nil, false
	}
	return o.UpdaterId, true
}

// HasUpdaterId returns a boolean if a field has been set.
func (o *AlipayIserviceCcmAgentGetResponseModel) HasUpdaterId() bool {
	if o != nil && !IsNil(o.UpdaterId) {
		return true
	}

	return false
}

// SetUpdaterId gets a reference to the given string and assigns it to the UpdaterId field.
func (o *AlipayIserviceCcmAgentGetResponseModel) SetUpdaterId(v string) {
	o.UpdaterId = &v
}

func (o AlipayIserviceCcmAgentGetResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayIserviceCcmAgentGetResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnsweringMode) {
		toSerialize["answering_mode"] = o.AnsweringMode
	}
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	if !IsNil(o.CcsInstanceIds) {
		toSerialize["ccs_instance_ids"] = o.CcsInstanceIds
	}
	if !IsNil(o.ChatConfig) {
		toSerialize["chat_config"] = o.ChatConfig
	}
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
	}
	if !IsNil(o.CreatorId) {
		toSerialize["creator_id"] = o.CreatorId
	}
	if !IsNil(o.DingtalkUserId) {
		toSerialize["dingtalk_user_id"] = o.DingtalkUserId
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.ExternalId) {
		toSerialize["external_id"] = o.ExternalId
	}
	if !IsNil(o.HotlineConfig) {
		toSerialize["hotline_config"] = o.HotlineConfig
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.JobNumber) {
		toSerialize["job_number"] = o.JobNumber
	}
	if !IsNil(o.LastLoginTime) {
		toSerialize["last_login_time"] = o.LastLoginTime
	}
	if !IsNil(o.Mobile) {
		toSerialize["mobile"] = o.Mobile
	}
	if !IsNil(o.NickName) {
		toSerialize["nick_name"] = o.NickName
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.RealName) {
		toSerialize["real_name"] = o.RealName
	}
	if !IsNil(o.RoleIds) {
		toSerialize["role_ids"] = o.RoleIds
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["update_time"] = o.UpdateTime
	}
	if !IsNil(o.UpdaterId) {
		toSerialize["updater_id"] = o.UpdaterId
	}
	return toSerialize, nil
}

type NullableAlipayIserviceCcmAgentGetResponseModel struct {
	value *AlipayIserviceCcmAgentGetResponseModel
	isSet bool
}

func (v NullableAlipayIserviceCcmAgentGetResponseModel) Get() *AlipayIserviceCcmAgentGetResponseModel {
	return v.value
}

func (v *NullableAlipayIserviceCcmAgentGetResponseModel) Set(val *AlipayIserviceCcmAgentGetResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmAgentGetResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmAgentGetResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmAgentGetResponseModel(val *AlipayIserviceCcmAgentGetResponseModel) *NullableAlipayIserviceCcmAgentGetResponseModel {
	return &NullableAlipayIserviceCcmAgentGetResponseModel{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmAgentGetResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmAgentGetResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


