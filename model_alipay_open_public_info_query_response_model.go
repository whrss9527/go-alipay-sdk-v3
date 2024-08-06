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

// checks if the AlipayOpenPublicInfoQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicInfoQueryResponseModel{}

// AlipayOpenPublicInfoQueryResponseModel struct for AlipayOpenPublicInfoQueryResponseModel
type AlipayOpenPublicInfoQueryResponseModel struct {
	// 生活号名称
	AppName *string `json:"app_name,omitempty"`
	// 最新审核状态描述，如果审核驳回则有相关的驳回理由
	AuditDesc *string `json:"audit_desc,omitempty"`
	// 最新审核状态，对于系统商而言，只有四个状态，AUDITING：审核中，AUDIT_FAILED：审核驳回，AUDIT_SUCCESS：审核通过，AUDIT_NORMAL：无审核状态（当前没有处于审核过程的工单）
	AuditStatus *string `json:"audit_status,omitempty"`
	// 审核状态表，目前支持名称、头像、名称与头像、简介审核状态
	AuditStatusList []PublicAuditStatus `json:"audit_status_list,omitempty"`
	// 背景图片地址
	BackgroundUrl *string `json:"background_url,omitempty"`
	// 生活号简介
	Introduction *string `json:"introduction,omitempty"`
	// 生活号是否上线，T表示上线，F表示未上线
	IsOnline *string `json:"is_online,omitempty"`
	// 生活号是否上架，T表示上架，上架即可在支付宝客户端被搜索到，F表示未上架
	IsRelease *string `json:"is_release,omitempty"`
	// 生活号头像地址
	LogoUrl *string `json:"logo_url,omitempty"`
	// 商家经营类目，详情参见 <a href=\"https://opendocs.alipay.com/open/01n22g\">商家经营类目</a>
	MccCodeDesc *string `json:"mcc_code_desc,omitempty"`
	// 欢迎语
	PublicGreeting *string `json:"public_greeting,omitempty"`
}

// NewAlipayOpenPublicInfoQueryResponseModel instantiates a new AlipayOpenPublicInfoQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicInfoQueryResponseModel() *AlipayOpenPublicInfoQueryResponseModel {
	this := AlipayOpenPublicInfoQueryResponseModel{}
	return &this
}

// NewAlipayOpenPublicInfoQueryResponseModelWithDefaults instantiates a new AlipayOpenPublicInfoQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicInfoQueryResponseModelWithDefaults() *AlipayOpenPublicInfoQueryResponseModel {
	this := AlipayOpenPublicInfoQueryResponseModel{}
	return &this
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *AlipayOpenPublicInfoQueryResponseModel) SetAppName(v string) {
	o.AppName = &v
}

// GetAuditDesc returns the AuditDesc field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetAuditDesc() string {
	if o == nil || IsNil(o.AuditDesc) {
		var ret string
		return ret
	}
	return *o.AuditDesc
}

// GetAuditDescOk returns a tuple with the AuditDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetAuditDescOk() (*string, bool) {
	if o == nil || IsNil(o.AuditDesc) {
		return nil, false
	}
	return o.AuditDesc, true
}

// HasAuditDesc returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) HasAuditDesc() bool {
	if o != nil && !IsNil(o.AuditDesc) {
		return true
	}

	return false
}

// SetAuditDesc gets a reference to the given string and assigns it to the AuditDesc field.
func (o *AlipayOpenPublicInfoQueryResponseModel) SetAuditDesc(v string) {
	o.AuditDesc = &v
}

// GetAuditStatus returns the AuditStatus field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetAuditStatus() string {
	if o == nil || IsNil(o.AuditStatus) {
		var ret string
		return ret
	}
	return *o.AuditStatus
}

// GetAuditStatusOk returns a tuple with the AuditStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetAuditStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AuditStatus) {
		return nil, false
	}
	return o.AuditStatus, true
}

// HasAuditStatus returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) HasAuditStatus() bool {
	if o != nil && !IsNil(o.AuditStatus) {
		return true
	}

	return false
}

// SetAuditStatus gets a reference to the given string and assigns it to the AuditStatus field.
func (o *AlipayOpenPublicInfoQueryResponseModel) SetAuditStatus(v string) {
	o.AuditStatus = &v
}

// GetAuditStatusList returns the AuditStatusList field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetAuditStatusList() []PublicAuditStatus {
	if o == nil || IsNil(o.AuditStatusList) {
		var ret []PublicAuditStatus
		return ret
	}
	return o.AuditStatusList
}

// GetAuditStatusListOk returns a tuple with the AuditStatusList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetAuditStatusListOk() ([]PublicAuditStatus, bool) {
	if o == nil || IsNil(o.AuditStatusList) {
		return nil, false
	}
	return o.AuditStatusList, true
}

// HasAuditStatusList returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) HasAuditStatusList() bool {
	if o != nil && !IsNil(o.AuditStatusList) {
		return true
	}

	return false
}

// SetAuditStatusList gets a reference to the given []PublicAuditStatus and assigns it to the AuditStatusList field.
func (o *AlipayOpenPublicInfoQueryResponseModel) SetAuditStatusList(v []PublicAuditStatus) {
	o.AuditStatusList = v
}

// GetBackgroundUrl returns the BackgroundUrl field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetBackgroundUrl() string {
	if o == nil || IsNil(o.BackgroundUrl) {
		var ret string
		return ret
	}
	return *o.BackgroundUrl
}

// GetBackgroundUrlOk returns a tuple with the BackgroundUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetBackgroundUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BackgroundUrl) {
		return nil, false
	}
	return o.BackgroundUrl, true
}

// HasBackgroundUrl returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) HasBackgroundUrl() bool {
	if o != nil && !IsNil(o.BackgroundUrl) {
		return true
	}

	return false
}

// SetBackgroundUrl gets a reference to the given string and assigns it to the BackgroundUrl field.
func (o *AlipayOpenPublicInfoQueryResponseModel) SetBackgroundUrl(v string) {
	o.BackgroundUrl = &v
}

// GetIntroduction returns the Introduction field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetIntroduction() string {
	if o == nil || IsNil(o.Introduction) {
		var ret string
		return ret
	}
	return *o.Introduction
}

// GetIntroductionOk returns a tuple with the Introduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetIntroductionOk() (*string, bool) {
	if o == nil || IsNil(o.Introduction) {
		return nil, false
	}
	return o.Introduction, true
}

// HasIntroduction returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) HasIntroduction() bool {
	if o != nil && !IsNil(o.Introduction) {
		return true
	}

	return false
}

// SetIntroduction gets a reference to the given string and assigns it to the Introduction field.
func (o *AlipayOpenPublicInfoQueryResponseModel) SetIntroduction(v string) {
	o.Introduction = &v
}

// GetIsOnline returns the IsOnline field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetIsOnline() string {
	if o == nil || IsNil(o.IsOnline) {
		var ret string
		return ret
	}
	return *o.IsOnline
}

// GetIsOnlineOk returns a tuple with the IsOnline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetIsOnlineOk() (*string, bool) {
	if o == nil || IsNil(o.IsOnline) {
		return nil, false
	}
	return o.IsOnline, true
}

// HasIsOnline returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) HasIsOnline() bool {
	if o != nil && !IsNil(o.IsOnline) {
		return true
	}

	return false
}

// SetIsOnline gets a reference to the given string and assigns it to the IsOnline field.
func (o *AlipayOpenPublicInfoQueryResponseModel) SetIsOnline(v string) {
	o.IsOnline = &v
}

// GetIsRelease returns the IsRelease field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetIsRelease() string {
	if o == nil || IsNil(o.IsRelease) {
		var ret string
		return ret
	}
	return *o.IsRelease
}

// GetIsReleaseOk returns a tuple with the IsRelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetIsReleaseOk() (*string, bool) {
	if o == nil || IsNil(o.IsRelease) {
		return nil, false
	}
	return o.IsRelease, true
}

// HasIsRelease returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) HasIsRelease() bool {
	if o != nil && !IsNil(o.IsRelease) {
		return true
	}

	return false
}

// SetIsRelease gets a reference to the given string and assigns it to the IsRelease field.
func (o *AlipayOpenPublicInfoQueryResponseModel) SetIsRelease(v string) {
	o.IsRelease = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) HasLogoUrl() bool {
	if o != nil && !IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *AlipayOpenPublicInfoQueryResponseModel) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetMccCodeDesc returns the MccCodeDesc field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetMccCodeDesc() string {
	if o == nil || IsNil(o.MccCodeDesc) {
		var ret string
		return ret
	}
	return *o.MccCodeDesc
}

// GetMccCodeDescOk returns a tuple with the MccCodeDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetMccCodeDescOk() (*string, bool) {
	if o == nil || IsNil(o.MccCodeDesc) {
		return nil, false
	}
	return o.MccCodeDesc, true
}

// HasMccCodeDesc returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) HasMccCodeDesc() bool {
	if o != nil && !IsNil(o.MccCodeDesc) {
		return true
	}

	return false
}

// SetMccCodeDesc gets a reference to the given string and assigns it to the MccCodeDesc field.
func (o *AlipayOpenPublicInfoQueryResponseModel) SetMccCodeDesc(v string) {
	o.MccCodeDesc = &v
}

// GetPublicGreeting returns the PublicGreeting field value if set, zero value otherwise.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetPublicGreeting() string {
	if o == nil || IsNil(o.PublicGreeting) {
		var ret string
		return ret
	}
	return *o.PublicGreeting
}

// GetPublicGreetingOk returns a tuple with the PublicGreeting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) GetPublicGreetingOk() (*string, bool) {
	if o == nil || IsNil(o.PublicGreeting) {
		return nil, false
	}
	return o.PublicGreeting, true
}

// HasPublicGreeting returns a boolean if a field has been set.
func (o *AlipayOpenPublicInfoQueryResponseModel) HasPublicGreeting() bool {
	if o != nil && !IsNil(o.PublicGreeting) {
		return true
	}

	return false
}

// SetPublicGreeting gets a reference to the given string and assigns it to the PublicGreeting field.
func (o *AlipayOpenPublicInfoQueryResponseModel) SetPublicGreeting(v string) {
	o.PublicGreeting = &v
}

func (o AlipayOpenPublicInfoQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicInfoQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if !IsNil(o.AuditDesc) {
		toSerialize["audit_desc"] = o.AuditDesc
	}
	if !IsNil(o.AuditStatus) {
		toSerialize["audit_status"] = o.AuditStatus
	}
	if !IsNil(o.AuditStatusList) {
		toSerialize["audit_status_list"] = o.AuditStatusList
	}
	if !IsNil(o.BackgroundUrl) {
		toSerialize["background_url"] = o.BackgroundUrl
	}
	if !IsNil(o.Introduction) {
		toSerialize["introduction"] = o.Introduction
	}
	if !IsNil(o.IsOnline) {
		toSerialize["is_online"] = o.IsOnline
	}
	if !IsNil(o.IsRelease) {
		toSerialize["is_release"] = o.IsRelease
	}
	if !IsNil(o.LogoUrl) {
		toSerialize["logo_url"] = o.LogoUrl
	}
	if !IsNil(o.MccCodeDesc) {
		toSerialize["mcc_code_desc"] = o.MccCodeDesc
	}
	if !IsNil(o.PublicGreeting) {
		toSerialize["public_greeting"] = o.PublicGreeting
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicInfoQueryResponseModel struct {
	value *AlipayOpenPublicInfoQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenPublicInfoQueryResponseModel) Get() *AlipayOpenPublicInfoQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenPublicInfoQueryResponseModel) Set(val *AlipayOpenPublicInfoQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicInfoQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicInfoQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicInfoQueryResponseModel(val *AlipayOpenPublicInfoQueryResponseModel) *NullableAlipayOpenPublicInfoQueryResponseModel {
	return &NullableAlipayOpenPublicInfoQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicInfoQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicInfoQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

