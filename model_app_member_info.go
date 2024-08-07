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

// checks if the AppMemberInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppMemberInfo{}

// AppMemberInfo struct for AppMemberInfo
type AppMemberInfo struct {
	// 邀请时间，格式为yyyy-MM-dd HH:mm:ss
	GmtInvite *string `json:"gmt_invite,omitempty"`
	// 加入时间，格式为yyyy-MM-dd
	GmtJoin *string `json:"gmt_join,omitempty"`
	// 支付宝登录账号
	LogonId *string `json:"logon_id,omitempty"`
	// 昵称
	NickName *string `json:"nick_name,omitempty"`
	// 支付宝用户openId
	OpenId *string `json:"open_id,omitempty"`
	// 头像地址fileId(完整地址需自行拼接)
	Portrait *string `json:"portrait,omitempty"`
	// 角色类型
	Role *string `json:"role,omitempty"`
	// 成员的状态
	Status *string `json:"status,omitempty"`
	// 支付宝用户id
	UserId *string `json:"user_id,omitempty"`
}

// NewAppMemberInfo instantiates a new AppMemberInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppMemberInfo() *AppMemberInfo {
	this := AppMemberInfo{}
	return &this
}

// NewAppMemberInfoWithDefaults instantiates a new AppMemberInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppMemberInfoWithDefaults() *AppMemberInfo {
	this := AppMemberInfo{}
	return &this
}

// GetGmtInvite returns the GmtInvite field value if set, zero value otherwise.
func (o *AppMemberInfo) GetGmtInvite() string {
	if o == nil || IsNil(o.GmtInvite) {
		var ret string
		return ret
	}
	return *o.GmtInvite
}

// GetGmtInviteOk returns a tuple with the GmtInvite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMemberInfo) GetGmtInviteOk() (*string, bool) {
	if o == nil || IsNil(o.GmtInvite) {
		return nil, false
	}
	return o.GmtInvite, true
}

// HasGmtInvite returns a boolean if a field has been set.
func (o *AppMemberInfo) HasGmtInvite() bool {
	if o != nil && !IsNil(o.GmtInvite) {
		return true
	}

	return false
}

// SetGmtInvite gets a reference to the given string and assigns it to the GmtInvite field.
func (o *AppMemberInfo) SetGmtInvite(v string) {
	o.GmtInvite = &v
}

// GetGmtJoin returns the GmtJoin field value if set, zero value otherwise.
func (o *AppMemberInfo) GetGmtJoin() string {
	if o == nil || IsNil(o.GmtJoin) {
		var ret string
		return ret
	}
	return *o.GmtJoin
}

// GetGmtJoinOk returns a tuple with the GmtJoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMemberInfo) GetGmtJoinOk() (*string, bool) {
	if o == nil || IsNil(o.GmtJoin) {
		return nil, false
	}
	return o.GmtJoin, true
}

// HasGmtJoin returns a boolean if a field has been set.
func (o *AppMemberInfo) HasGmtJoin() bool {
	if o != nil && !IsNil(o.GmtJoin) {
		return true
	}

	return false
}

// SetGmtJoin gets a reference to the given string and assigns it to the GmtJoin field.
func (o *AppMemberInfo) SetGmtJoin(v string) {
	o.GmtJoin = &v
}

// GetLogonId returns the LogonId field value if set, zero value otherwise.
func (o *AppMemberInfo) GetLogonId() string {
	if o == nil || IsNil(o.LogonId) {
		var ret string
		return ret
	}
	return *o.LogonId
}

// GetLogonIdOk returns a tuple with the LogonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMemberInfo) GetLogonIdOk() (*string, bool) {
	if o == nil || IsNil(o.LogonId) {
		return nil, false
	}
	return o.LogonId, true
}

// HasLogonId returns a boolean if a field has been set.
func (o *AppMemberInfo) HasLogonId() bool {
	if o != nil && !IsNil(o.LogonId) {
		return true
	}

	return false
}

// SetLogonId gets a reference to the given string and assigns it to the LogonId field.
func (o *AppMemberInfo) SetLogonId(v string) {
	o.LogonId = &v
}

// GetNickName returns the NickName field value if set, zero value otherwise.
func (o *AppMemberInfo) GetNickName() string {
	if o == nil || IsNil(o.NickName) {
		var ret string
		return ret
	}
	return *o.NickName
}

// GetNickNameOk returns a tuple with the NickName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMemberInfo) GetNickNameOk() (*string, bool) {
	if o == nil || IsNil(o.NickName) {
		return nil, false
	}
	return o.NickName, true
}

// HasNickName returns a boolean if a field has been set.
func (o *AppMemberInfo) HasNickName() bool {
	if o != nil && !IsNil(o.NickName) {
		return true
	}

	return false
}

// SetNickName gets a reference to the given string and assigns it to the NickName field.
func (o *AppMemberInfo) SetNickName(v string) {
	o.NickName = &v
}

// GetOpenId returns the OpenId field value if set, zero value otherwise.
func (o *AppMemberInfo) GetOpenId() string {
	if o == nil || IsNil(o.OpenId) {
		var ret string
		return ret
	}
	return *o.OpenId
}

// GetOpenIdOk returns a tuple with the OpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMemberInfo) GetOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenId) {
		return nil, false
	}
	return o.OpenId, true
}

// HasOpenId returns a boolean if a field has been set.
func (o *AppMemberInfo) HasOpenId() bool {
	if o != nil && !IsNil(o.OpenId) {
		return true
	}

	return false
}

// SetOpenId gets a reference to the given string and assigns it to the OpenId field.
func (o *AppMemberInfo) SetOpenId(v string) {
	o.OpenId = &v
}

// GetPortrait returns the Portrait field value if set, zero value otherwise.
func (o *AppMemberInfo) GetPortrait() string {
	if o == nil || IsNil(o.Portrait) {
		var ret string
		return ret
	}
	return *o.Portrait
}

// GetPortraitOk returns a tuple with the Portrait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMemberInfo) GetPortraitOk() (*string, bool) {
	if o == nil || IsNil(o.Portrait) {
		return nil, false
	}
	return o.Portrait, true
}

// HasPortrait returns a boolean if a field has been set.
func (o *AppMemberInfo) HasPortrait() bool {
	if o != nil && !IsNil(o.Portrait) {
		return true
	}

	return false
}

// SetPortrait gets a reference to the given string and assigns it to the Portrait field.
func (o *AppMemberInfo) SetPortrait(v string) {
	o.Portrait = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *AppMemberInfo) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMemberInfo) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *AppMemberInfo) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *AppMemberInfo) SetRole(v string) {
	o.Role = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AppMemberInfo) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMemberInfo) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AppMemberInfo) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AppMemberInfo) SetStatus(v string) {
	o.Status = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AppMemberInfo) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMemberInfo) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AppMemberInfo) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AppMemberInfo) SetUserId(v string) {
	o.UserId = &v
}

func (o AppMemberInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppMemberInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GmtInvite) {
		toSerialize["gmt_invite"] = o.GmtInvite
	}
	if !IsNil(o.GmtJoin) {
		toSerialize["gmt_join"] = o.GmtJoin
	}
	if !IsNil(o.LogonId) {
		toSerialize["logon_id"] = o.LogonId
	}
	if !IsNil(o.NickName) {
		toSerialize["nick_name"] = o.NickName
	}
	if !IsNil(o.OpenId) {
		toSerialize["open_id"] = o.OpenId
	}
	if !IsNil(o.Portrait) {
		toSerialize["portrait"] = o.Portrait
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableAppMemberInfo struct {
	value *AppMemberInfo
	isSet bool
}

func (v NullableAppMemberInfo) Get() *AppMemberInfo {
	return v.value
}

func (v *NullableAppMemberInfo) Set(val *AppMemberInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAppMemberInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAppMemberInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppMemberInfo(val *AppMemberInfo) *NullableAppMemberInfo {
	return &NullableAppMemberInfo{value: val, isSet: true}
}

func (v NullableAppMemberInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppMemberInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
