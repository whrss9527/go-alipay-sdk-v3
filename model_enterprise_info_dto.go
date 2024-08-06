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

// checks if the EnterpriseInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnterpriseInfoDTO{}

// EnterpriseInfoDTO struct for EnterpriseInfoDTO
type EnterpriseInfoDTO struct {
	// 共同账户id
	AccountId *string `json:"account_id,omitempty"`
	// 因公签约状态
	Activate *string `json:"activate,omitempty"`
	// 企业认证等级
	AuthLevel *string `json:"auth_level,omitempty"`
	// 认证状态
	AuthStatus *string `json:"auth_status,omitempty"`
	// 认证时间
	AuthTime *string `json:"auth_time,omitempty"`
	// 企业简称
	EnterpriseAlias *string `json:"enterprise_alias,omitempty"`
	// 统一社会信用码
	EnterpriseCode *string `json:"enterprise_code,omitempty"`
	// 企业邮箱
	EnterpriseEmail *string `json:"enterprise_email,omitempty"`
	// 企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// 企业名称
	EnterpriseName *string `json:"enterprise_name,omitempty"`
	// 创建时间
	GmtCreate *string `json:"gmt_create,omitempty"`
	// 更新时间
	GmtModified *string `json:"gmt_modified,omitempty"`
	// 封闭场景（如班车）的人脸库id
	IotGroupId *string `json:"iot_group_id,omitempty"`
	// 开放场景（如团餐）的人脸库id
	IotLogicGroupId *string `json:"iot_logic_group_id,omitempty"`
	// 外部平台code
	PlatformCode *string `json:"platform_code,omitempty"`
	// 外部平台id
	PlatformOutId *string `json:"platform_out_id,omitempty"`
	// 企业码服务协议签约状态
	SignStatus *string `json:"sign_status,omitempty"`
	// 企业当前状态
	Status *string `json:"status,omitempty"`
}

// NewEnterpriseInfoDTO instantiates a new EnterpriseInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterpriseInfoDTO() *EnterpriseInfoDTO {
	this := EnterpriseInfoDTO{}
	return &this
}

// NewEnterpriseInfoDTOWithDefaults instantiates a new EnterpriseInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterpriseInfoDTOWithDefaults() *EnterpriseInfoDTO {
	this := EnterpriseInfoDTO{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *EnterpriseInfoDTO) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseInfoDTO) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *EnterpriseInfoDTO) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *EnterpriseInfoDTO) SetAccountId(v string) {
	o.AccountId = &v
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *EnterpriseInfoDTO) GetActivate() string {
	if o == nil || IsNil(o.Activate) {
		var ret string
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseInfoDTO) GetActivateOk() (*string, bool) {
	if o == nil || IsNil(o.Activate) {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *EnterpriseInfoDTO) HasActivate() bool {
	if o != nil && !IsNil(o.Activate) {
		return true
	}

	return false
}

// SetActivate gets a reference to the given string and assigns it to the Activate field.
func (o *EnterpriseInfoDTO) SetActivate(v string) {
	o.Activate = &v
}

// GetAuthLevel returns the AuthLevel field value if set, zero value otherwise.
func (o *EnterpriseInfoDTO) GetAuthLevel() string {
	if o == nil || IsNil(o.AuthLevel) {
		var ret string
		return ret
	}
	return *o.AuthLevel
}

// GetAuthLevelOk returns a tuple with the AuthLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseInfoDTO) GetAuthLevelOk() (*string, bool) {
	if o == nil || IsNil(o.AuthLevel) {
		return nil, false
	}
	return o.AuthLevel, true
}

// HasAuthLevel returns a boolean if a field has been set.
func (o *EnterpriseInfoDTO) HasAuthLevel() bool {
	if o != nil && !IsNil(o.AuthLevel) {
		return true
	}

	return false
}

// SetAuthLevel gets a reference to the given string and assigns it to the AuthLevel field.
func (o *EnterpriseInfoDTO) SetAuthLevel(v string) {
	o.AuthLevel = &v
}

// GetAuthStatus returns the AuthStatus field value if set, zero value otherwise.
func (o *EnterpriseInfoDTO) GetAuthStatus() string {
	if o == nil || IsNil(o.AuthStatus) {
		var ret string
		return ret
	}
	return *o.AuthStatus
}

// GetAuthStatusOk returns a tuple with the AuthStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseInfoDTO) GetAuthStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AuthStatus) {
		return nil, false
	}
	return o.AuthStatus, true
}

// HasAuthStatus returns a boolean if a field has been set.
func (o *EnterpriseInfoDTO) HasAuthStatus() bool {
	if o != nil && !IsNil(o.AuthStatus) {
		return true
	}

	return false
}

// SetAuthStatus gets a reference to the given string and assigns it to the AuthStatus field.
func (o *EnterpriseInfoDTO) SetAuthStatus(v string) {
	o.AuthStatus = &v
}

// GetAuthTime returns the AuthTime field value if set, zero value otherwise.
func (o *EnterpriseInfoDTO) GetAuthTime() string {
	if o == nil || IsNil(o.AuthTime) {
		var ret string
		return ret
	}
	return *o.AuthTime
}

// GetAuthTimeOk returns a tuple with the AuthTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseInfoDTO) GetAuthTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthTime) {
		return nil, false
	}
	return o.AuthTime, true
}

// HasAuthTime returns a boolean if a field has been set.
func (o *EnterpriseInfoDTO) HasAuthTime() bool {
	if o != nil && !IsNil(o.AuthTime) {
		return true
	}

	return false
}

// SetAuthTime gets a reference to the given string and assigns it to the AuthTime field.
func (o *EnterpriseInfoDTO) SetAuthTime(v string) {
	o.AuthTime = &v
}

// GetEnterpriseAlias returns the EnterpriseAlias field value if set, zero value otherwise.
func (o *EnterpriseInfoDTO) GetEnterpriseAlias() string {
	if o == nil || IsNil(o.EnterpriseAlias) {
		var ret string
		return ret
	}
	return *o.EnterpriseAlias
}

// GetEnterpriseAliasOk returns a tuple with the EnterpriseAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseInfoDTO) GetEnterpriseAliasOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseAlias) {
		return nil, false
	}
	return o.EnterpriseAlias, true
}

// HasEnterpriseAlias returns a boolean if a field has been set.
func (o *EnterpriseInfoDTO) HasEnterpriseAlias() bool {
	if o != nil && !IsNil(o.EnterpriseAlias) {
		return true
	}

	return false
}

// SetEnterpriseAlias gets a reference to the given string and assigns it to the EnterpriseAlias field.
func (o *EnterpriseInfoDTO) SetEnterpriseAlias(v string) {
	o.EnterpriseAlias = &v
}

// GetEnterpriseCode returns the EnterpriseCode field value if set, zero value otherwise.
func (o *EnterpriseInfoDTO) GetEnterpriseCode() string {
	if o == nil || IsNil(o.EnterpriseCode) {
		var ret string
		return ret
	}
	return *o.EnterpriseCode
}

// GetEnterpriseCodeOk returns a tuple with the EnterpriseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseInfoDTO) GetEnterpriseCodeOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseCode) {
		return nil, false
	}
	return o.EnterpriseCode, true
}

// HasEnterpriseCode returns a boolean if a field has been set.
func (o *EnterpriseInfoDTO) HasEnterpriseCode() bool {
	if o != nil && !IsNil(o.EnterpriseCode) {
		return true
	}

	return false
}

// SetEnterpriseCode gets a reference to the given string and assigns it to the EnterpriseCode field.
func (o *EnterpriseInfoDTO) SetEnterpriseCode(v string) {
	o.EnterpriseCode = &v
}

// GetEnterpriseEmail returns the EnterpriseEmail field value if set, zero value otherwise.
func (o *EnterpriseInfoDTO) GetEnterpriseEmail() string {
	if o == nil || IsNil(o.EnterpriseEmail) {
		var ret string
		return ret
	}
	return *o.EnterpriseEmail
}

// GetEnterpriseEmailOk returns a tuple with the EnterpriseEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseInfoDTO) GetEnterpriseEmailOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseEmail) {
		return nil, false
	}
	return o.EnterpriseEmail, true
}

// HasEnterpriseEmail returns a boolean if a field has been set.
func (o *EnterpriseInfoDTO) HasEnterpriseEmail() bool {
	if o != nil && !IsNil(o.EnterpriseEmail) {
		return true
	}

	return false
}

// SetEnterpriseEmail gets a reference to the given string and assigns it to the EnterpriseEmail field.
func (o *EnterpriseInfoDTO) SetEnterpriseEmail(v string) {
	o.EnterpriseEmail = &v
}

// GetEnterpriseId returns the EnterpriseId field value if set, zero value otherwise.
func (o *EnterpriseInfoDTO) GetEnterpriseId() string {
	if o == nil || IsNil(o.EnterpriseId) {
		var ret string
		return ret
	}
	return *o.EnterpriseId
}

// GetEnterpriseIdOk returns a tuple with the EnterpriseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseInfoDTO) GetEnterpriseIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseId) {
		return nil, false
	}
	return o.EnterpriseId, true
}

// HasEnterpriseId returns a boolean if a field has been set.
func (o *EnterpriseInfoDTO) HasEnterpriseId() bool {
	if o != nil && !IsNil(o.EnterpriseId) {
		return true
	}

	return false
}

// SetEnterpriseId gets a reference to the given string and assigns it to the EnterpriseId field.
func (o *EnterpriseInfoDTO) SetEnterpriseId(v string) {
	o.EnterpriseId = &v
}

// GetEnterpriseName returns the EnterpriseName field value if set, zero value otherwise.
func (o *EnterpriseInfoDTO) GetEnterpriseName() string {
	if o == nil || IsNil(o.EnterpriseName) {
		var ret string
		return ret
	}
	return *o.EnterpriseName
}

// GetEnterpriseNameOk returns a tuple with the EnterpriseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseInfoDTO) GetEnterpriseNameOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseName) {
		return nil, false
	}
	return o.EnterpriseName, true
}

// HasEnterpriseName returns a boolean if a field has been set.
func (o *EnterpriseInfoDTO) HasEnterpriseName() bool {
	if o != nil && !IsNil(o.EnterpriseName) {
		return true
	}

	return false
}

// SetEnterpriseName gets a reference to the given string and assigns it to the EnterpriseName field.
func (o *EnterpriseInfoDTO) SetEnterpriseName(v string) {
	o.EnterpriseName = &v
}

// GetGmtCreate returns the GmtCreate field value if set, zero value otherwise.
func (o *EnterpriseInfoDTO) GetGmtCreate() string {
	if o == nil || IsNil(o.GmtCreate) {
		var ret string
		return ret
	}
	return *o.GmtCreate
}

// GetGmtCreateOk returns a tuple with the GmtCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseInfoDTO) GetGmtCreateOk() (*string, bool) {
	if o == nil || IsNil(o.GmtCreate) {
		return nil, false
	}
	return o.GmtCreate, true
}

// HasGmtCreate returns a boolean if a field has been set.
func (o *EnterpriseInfoDTO) HasGmtCreate() bool {
	if o != nil && !IsNil(o.GmtCreate) {
		return true
	}

	return false
}

// SetGmtCreate gets a reference to the given string and assigns it to the GmtCreate field.
func (o *EnterpriseInfoDTO) SetGmtCreate(v string) {
	o.GmtCreate = &v
}

// GetGmtModified returns the GmtModified field value if set, zero value otherwise.
func (o *EnterpriseInfoDTO) GetGmtModified() string {
	if o == nil || IsNil(o.GmtModified) {
		var ret string
		return ret
	}
	return *o.GmtModified
}

// GetGmtModifiedOk returns a tuple with the GmtModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseInfoDTO) GetGmtModifiedOk() (*string, bool) {
	if o == nil || IsNil(o.GmtModified) {
		return nil, false
	}
	return o.GmtModified, true
}

// HasGmtModified returns a boolean if a field has been set.
func (o *EnterpriseInfoDTO) HasGmtModified() bool {
	if o != nil && !IsNil(o.GmtModified) {
		return true
	}

	return false
}

// SetGmtModified gets a reference to the given string and assigns it to the GmtModified field.
func (o *EnterpriseInfoDTO) SetGmtModified(v string) {
	o.GmtModified = &v
}

// GetIotGroupId returns the IotGroupId field value if set, zero value otherwise.
func (o *EnterpriseInfoDTO) GetIotGroupId() string {
	if o == nil || IsNil(o.IotGroupId) {
		var ret string
		return ret
	}
	return *o.IotGroupId
}

// GetIotGroupIdOk returns a tuple with the IotGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseInfoDTO) GetIotGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.IotGroupId) {
		return nil, false
	}
	return o.IotGroupId, true
}

// HasIotGroupId returns a boolean if a field has been set.
func (o *EnterpriseInfoDTO) HasIotGroupId() bool {
	if o != nil && !IsNil(o.IotGroupId) {
		return true
	}

	return false
}

// SetIotGroupId gets a reference to the given string and assigns it to the IotGroupId field.
func (o *EnterpriseInfoDTO) SetIotGroupId(v string) {
	o.IotGroupId = &v
}

// GetIotLogicGroupId returns the IotLogicGroupId field value if set, zero value otherwise.
func (o *EnterpriseInfoDTO) GetIotLogicGroupId() string {
	if o == nil || IsNil(o.IotLogicGroupId) {
		var ret string
		return ret
	}
	return *o.IotLogicGroupId
}

// GetIotLogicGroupIdOk returns a tuple with the IotLogicGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseInfoDTO) GetIotLogicGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.IotLogicGroupId) {
		return nil, false
	}
	return o.IotLogicGroupId, true
}

// HasIotLogicGroupId returns a boolean if a field has been set.
func (o *EnterpriseInfoDTO) HasIotLogicGroupId() bool {
	if o != nil && !IsNil(o.IotLogicGroupId) {
		return true
	}

	return false
}

// SetIotLogicGroupId gets a reference to the given string and assigns it to the IotLogicGroupId field.
func (o *EnterpriseInfoDTO) SetIotLogicGroupId(v string) {
	o.IotLogicGroupId = &v
}

// GetPlatformCode returns the PlatformCode field value if set, zero value otherwise.
func (o *EnterpriseInfoDTO) GetPlatformCode() string {
	if o == nil || IsNil(o.PlatformCode) {
		var ret string
		return ret
	}
	return *o.PlatformCode
}

// GetPlatformCodeOk returns a tuple with the PlatformCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseInfoDTO) GetPlatformCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PlatformCode) {
		return nil, false
	}
	return o.PlatformCode, true
}

// HasPlatformCode returns a boolean if a field has been set.
func (o *EnterpriseInfoDTO) HasPlatformCode() bool {
	if o != nil && !IsNil(o.PlatformCode) {
		return true
	}

	return false
}

// SetPlatformCode gets a reference to the given string and assigns it to the PlatformCode field.
func (o *EnterpriseInfoDTO) SetPlatformCode(v string) {
	o.PlatformCode = &v
}

// GetPlatformOutId returns the PlatformOutId field value if set, zero value otherwise.
func (o *EnterpriseInfoDTO) GetPlatformOutId() string {
	if o == nil || IsNil(o.PlatformOutId) {
		var ret string
		return ret
	}
	return *o.PlatformOutId
}

// GetPlatformOutIdOk returns a tuple with the PlatformOutId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseInfoDTO) GetPlatformOutIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlatformOutId) {
		return nil, false
	}
	return o.PlatformOutId, true
}

// HasPlatformOutId returns a boolean if a field has been set.
func (o *EnterpriseInfoDTO) HasPlatformOutId() bool {
	if o != nil && !IsNil(o.PlatformOutId) {
		return true
	}

	return false
}

// SetPlatformOutId gets a reference to the given string and assigns it to the PlatformOutId field.
func (o *EnterpriseInfoDTO) SetPlatformOutId(v string) {
	o.PlatformOutId = &v
}

// GetSignStatus returns the SignStatus field value if set, zero value otherwise.
func (o *EnterpriseInfoDTO) GetSignStatus() string {
	if o == nil || IsNil(o.SignStatus) {
		var ret string
		return ret
	}
	return *o.SignStatus
}

// GetSignStatusOk returns a tuple with the SignStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseInfoDTO) GetSignStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SignStatus) {
		return nil, false
	}
	return o.SignStatus, true
}

// HasSignStatus returns a boolean if a field has been set.
func (o *EnterpriseInfoDTO) HasSignStatus() bool {
	if o != nil && !IsNil(o.SignStatus) {
		return true
	}

	return false
}

// SetSignStatus gets a reference to the given string and assigns it to the SignStatus field.
func (o *EnterpriseInfoDTO) SetSignStatus(v string) {
	o.SignStatus = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EnterpriseInfoDTO) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseInfoDTO) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EnterpriseInfoDTO) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EnterpriseInfoDTO) SetStatus(v string) {
	o.Status = &v
}

func (o EnterpriseInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnterpriseInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.Activate) {
		toSerialize["activate"] = o.Activate
	}
	if !IsNil(o.AuthLevel) {
		toSerialize["auth_level"] = o.AuthLevel
	}
	if !IsNil(o.AuthStatus) {
		toSerialize["auth_status"] = o.AuthStatus
	}
	if !IsNil(o.AuthTime) {
		toSerialize["auth_time"] = o.AuthTime
	}
	if !IsNil(o.EnterpriseAlias) {
		toSerialize["enterprise_alias"] = o.EnterpriseAlias
	}
	if !IsNil(o.EnterpriseCode) {
		toSerialize["enterprise_code"] = o.EnterpriseCode
	}
	if !IsNil(o.EnterpriseEmail) {
		toSerialize["enterprise_email"] = o.EnterpriseEmail
	}
	if !IsNil(o.EnterpriseId) {
		toSerialize["enterprise_id"] = o.EnterpriseId
	}
	if !IsNil(o.EnterpriseName) {
		toSerialize["enterprise_name"] = o.EnterpriseName
	}
	if !IsNil(o.GmtCreate) {
		toSerialize["gmt_create"] = o.GmtCreate
	}
	if !IsNil(o.GmtModified) {
		toSerialize["gmt_modified"] = o.GmtModified
	}
	if !IsNil(o.IotGroupId) {
		toSerialize["iot_group_id"] = o.IotGroupId
	}
	if !IsNil(o.IotLogicGroupId) {
		toSerialize["iot_logic_group_id"] = o.IotLogicGroupId
	}
	if !IsNil(o.PlatformCode) {
		toSerialize["platform_code"] = o.PlatformCode
	}
	if !IsNil(o.PlatformOutId) {
		toSerialize["platform_out_id"] = o.PlatformOutId
	}
	if !IsNil(o.SignStatus) {
		toSerialize["sign_status"] = o.SignStatus
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableEnterpriseInfoDTO struct {
	value *EnterpriseInfoDTO
	isSet bool
}

func (v NullableEnterpriseInfoDTO) Get() *EnterpriseInfoDTO {
	return v.value
}

func (v *NullableEnterpriseInfoDTO) Set(val *EnterpriseInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseInfoDTO(val *EnterpriseInfoDTO) *NullableEnterpriseInfoDTO {
	return &NullableEnterpriseInfoDTO{value: val, isSet: true}
}

func (v NullableEnterpriseInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
