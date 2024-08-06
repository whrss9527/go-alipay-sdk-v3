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

// checks if the ConfigInfoBean type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigInfoBean{}

// ConfigInfoBean struct for ConfigInfoBean
type ConfigInfoBean struct {
	AttachmentExplain *AttachmentExplainBean `json:"attachment_explain,omitempty"`
	// 签署完成是否自动归档，默认true，如果false，则归档需要调用签署流程归档接口
	AutoArchive *bool `json:"auto_archive,omitempty"`
	// 是否收集附件（需签署人上传身份证或者其他文件的，需设置为true）
	CollectAttachement *bool `json:"collect_attachement,omitempty"`
	// 物流公司编号（目前仅支持顺丰\"SF\"） 注：避免影响生产订单，联调及测试环境请传入SF_TEST。
	CompanyNo *string `json:"company_no,omitempty"`
	// 合同过期时间：该参数设置的时间若到期，则会触发流程文件过期通知
	ContractValidity *int32 `json:"contract_validity,omitempty"`
	// 签署完成后跳转商户小程序的地址（signModel字段值为2时 选填）
	JumpUrl *string `json:"jump_url,omitempty"`
	// 商户小程序签署地址（signModel字段值为1时 必填）
	MerchantMiniSignUrl *string `json:"merchant_mini_sign_url,omitempty"`
	// https://esign.cn/callback
	NoticeDeveloperUrl *string `json:"notice_developer_url,omitempty"`
	// 平台订单号，可将商户订单与电子合同进行关联。用户通过支付宝官方小程序 合同管家 查看合同时，可快速进入商家的小程序及对应的订单详情页，为商家的小程序提供了流量入口，提高用户的活跃度和留存率；
	PlatformOrderNo *string `json:"platform_order_no,omitempty"`
	// 12或15位物流单号
	SerialNo *string `json:"serial_no,omitempty"`
	// 快递扫码签收方式： 1-商户小程序：扫码后跳转商户小程序进行签收； 0-e签宝小程序：如商户没有支付宝小程序，可使用e签宝小程序完成签收。
	SignModel *int32 `json:"sign_model,omitempty"`
	// 指定签署人需要上传的附件列表，例如：A、B、C三方签署，A需上传附件1，B需指定附件2，C需上传附件1,2,3
	SpecifyAttachmentInfos []SpecifyAttachmentInfo `json:"specify_attachment_infos,omitempty"`
}

// NewConfigInfoBean instantiates a new ConfigInfoBean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigInfoBean() *ConfigInfoBean {
	this := ConfigInfoBean{}
	return &this
}

// NewConfigInfoBeanWithDefaults instantiates a new ConfigInfoBean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigInfoBeanWithDefaults() *ConfigInfoBean {
	this := ConfigInfoBean{}
	return &this
}

// GetAttachmentExplain returns the AttachmentExplain field value if set, zero value otherwise.
func (o *ConfigInfoBean) GetAttachmentExplain() AttachmentExplainBean {
	if o == nil || IsNil(o.AttachmentExplain) {
		var ret AttachmentExplainBean
		return ret
	}
	return *o.AttachmentExplain
}

// GetAttachmentExplainOk returns a tuple with the AttachmentExplain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInfoBean) GetAttachmentExplainOk() (*AttachmentExplainBean, bool) {
	if o == nil || IsNil(o.AttachmentExplain) {
		return nil, false
	}
	return o.AttachmentExplain, true
}

// HasAttachmentExplain returns a boolean if a field has been set.
func (o *ConfigInfoBean) HasAttachmentExplain() bool {
	if o != nil && !IsNil(o.AttachmentExplain) {
		return true
	}

	return false
}

// SetAttachmentExplain gets a reference to the given AttachmentExplainBean and assigns it to the AttachmentExplain field.
func (o *ConfigInfoBean) SetAttachmentExplain(v AttachmentExplainBean) {
	o.AttachmentExplain = &v
}

// GetAutoArchive returns the AutoArchive field value if set, zero value otherwise.
func (o *ConfigInfoBean) GetAutoArchive() bool {
	if o == nil || IsNil(o.AutoArchive) {
		var ret bool
		return ret
	}
	return *o.AutoArchive
}

// GetAutoArchiveOk returns a tuple with the AutoArchive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInfoBean) GetAutoArchiveOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoArchive) {
		return nil, false
	}
	return o.AutoArchive, true
}

// HasAutoArchive returns a boolean if a field has been set.
func (o *ConfigInfoBean) HasAutoArchive() bool {
	if o != nil && !IsNil(o.AutoArchive) {
		return true
	}

	return false
}

// SetAutoArchive gets a reference to the given bool and assigns it to the AutoArchive field.
func (o *ConfigInfoBean) SetAutoArchive(v bool) {
	o.AutoArchive = &v
}

// GetCollectAttachement returns the CollectAttachement field value if set, zero value otherwise.
func (o *ConfigInfoBean) GetCollectAttachement() bool {
	if o == nil || IsNil(o.CollectAttachement) {
		var ret bool
		return ret
	}
	return *o.CollectAttachement
}

// GetCollectAttachementOk returns a tuple with the CollectAttachement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInfoBean) GetCollectAttachementOk() (*bool, bool) {
	if o == nil || IsNil(o.CollectAttachement) {
		return nil, false
	}
	return o.CollectAttachement, true
}

// HasCollectAttachement returns a boolean if a field has been set.
func (o *ConfigInfoBean) HasCollectAttachement() bool {
	if o != nil && !IsNil(o.CollectAttachement) {
		return true
	}

	return false
}

// SetCollectAttachement gets a reference to the given bool and assigns it to the CollectAttachement field.
func (o *ConfigInfoBean) SetCollectAttachement(v bool) {
	o.CollectAttachement = &v
}

// GetCompanyNo returns the CompanyNo field value if set, zero value otherwise.
func (o *ConfigInfoBean) GetCompanyNo() string {
	if o == nil || IsNil(o.CompanyNo) {
		var ret string
		return ret
	}
	return *o.CompanyNo
}

// GetCompanyNoOk returns a tuple with the CompanyNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInfoBean) GetCompanyNoOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyNo) {
		return nil, false
	}
	return o.CompanyNo, true
}

// HasCompanyNo returns a boolean if a field has been set.
func (o *ConfigInfoBean) HasCompanyNo() bool {
	if o != nil && !IsNil(o.CompanyNo) {
		return true
	}

	return false
}

// SetCompanyNo gets a reference to the given string and assigns it to the CompanyNo field.
func (o *ConfigInfoBean) SetCompanyNo(v string) {
	o.CompanyNo = &v
}

// GetContractValidity returns the ContractValidity field value if set, zero value otherwise.
func (o *ConfigInfoBean) GetContractValidity() int32 {
	if o == nil || IsNil(o.ContractValidity) {
		var ret int32
		return ret
	}
	return *o.ContractValidity
}

// GetContractValidityOk returns a tuple with the ContractValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInfoBean) GetContractValidityOk() (*int32, bool) {
	if o == nil || IsNil(o.ContractValidity) {
		return nil, false
	}
	return o.ContractValidity, true
}

// HasContractValidity returns a boolean if a field has been set.
func (o *ConfigInfoBean) HasContractValidity() bool {
	if o != nil && !IsNil(o.ContractValidity) {
		return true
	}

	return false
}

// SetContractValidity gets a reference to the given int32 and assigns it to the ContractValidity field.
func (o *ConfigInfoBean) SetContractValidity(v int32) {
	o.ContractValidity = &v
}

// GetJumpUrl returns the JumpUrl field value if set, zero value otherwise.
func (o *ConfigInfoBean) GetJumpUrl() string {
	if o == nil || IsNil(o.JumpUrl) {
		var ret string
		return ret
	}
	return *o.JumpUrl
}

// GetJumpUrlOk returns a tuple with the JumpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInfoBean) GetJumpUrlOk() (*string, bool) {
	if o == nil || IsNil(o.JumpUrl) {
		return nil, false
	}
	return o.JumpUrl, true
}

// HasJumpUrl returns a boolean if a field has been set.
func (o *ConfigInfoBean) HasJumpUrl() bool {
	if o != nil && !IsNil(o.JumpUrl) {
		return true
	}

	return false
}

// SetJumpUrl gets a reference to the given string and assigns it to the JumpUrl field.
func (o *ConfigInfoBean) SetJumpUrl(v string) {
	o.JumpUrl = &v
}

// GetMerchantMiniSignUrl returns the MerchantMiniSignUrl field value if set, zero value otherwise.
func (o *ConfigInfoBean) GetMerchantMiniSignUrl() string {
	if o == nil || IsNil(o.MerchantMiniSignUrl) {
		var ret string
		return ret
	}
	return *o.MerchantMiniSignUrl
}

// GetMerchantMiniSignUrlOk returns a tuple with the MerchantMiniSignUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInfoBean) GetMerchantMiniSignUrlOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantMiniSignUrl) {
		return nil, false
	}
	return o.MerchantMiniSignUrl, true
}

// HasMerchantMiniSignUrl returns a boolean if a field has been set.
func (o *ConfigInfoBean) HasMerchantMiniSignUrl() bool {
	if o != nil && !IsNil(o.MerchantMiniSignUrl) {
		return true
	}

	return false
}

// SetMerchantMiniSignUrl gets a reference to the given string and assigns it to the MerchantMiniSignUrl field.
func (o *ConfigInfoBean) SetMerchantMiniSignUrl(v string) {
	o.MerchantMiniSignUrl = &v
}

// GetNoticeDeveloperUrl returns the NoticeDeveloperUrl field value if set, zero value otherwise.
func (o *ConfigInfoBean) GetNoticeDeveloperUrl() string {
	if o == nil || IsNil(o.NoticeDeveloperUrl) {
		var ret string
		return ret
	}
	return *o.NoticeDeveloperUrl
}

// GetNoticeDeveloperUrlOk returns a tuple with the NoticeDeveloperUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInfoBean) GetNoticeDeveloperUrlOk() (*string, bool) {
	if o == nil || IsNil(o.NoticeDeveloperUrl) {
		return nil, false
	}
	return o.NoticeDeveloperUrl, true
}

// HasNoticeDeveloperUrl returns a boolean if a field has been set.
func (o *ConfigInfoBean) HasNoticeDeveloperUrl() bool {
	if o != nil && !IsNil(o.NoticeDeveloperUrl) {
		return true
	}

	return false
}

// SetNoticeDeveloperUrl gets a reference to the given string and assigns it to the NoticeDeveloperUrl field.
func (o *ConfigInfoBean) SetNoticeDeveloperUrl(v string) {
	o.NoticeDeveloperUrl = &v
}

// GetPlatformOrderNo returns the PlatformOrderNo field value if set, zero value otherwise.
func (o *ConfigInfoBean) GetPlatformOrderNo() string {
	if o == nil || IsNil(o.PlatformOrderNo) {
		var ret string
		return ret
	}
	return *o.PlatformOrderNo
}

// GetPlatformOrderNoOk returns a tuple with the PlatformOrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInfoBean) GetPlatformOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.PlatformOrderNo) {
		return nil, false
	}
	return o.PlatformOrderNo, true
}

// HasPlatformOrderNo returns a boolean if a field has been set.
func (o *ConfigInfoBean) HasPlatformOrderNo() bool {
	if o != nil && !IsNil(o.PlatformOrderNo) {
		return true
	}

	return false
}

// SetPlatformOrderNo gets a reference to the given string and assigns it to the PlatformOrderNo field.
func (o *ConfigInfoBean) SetPlatformOrderNo(v string) {
	o.PlatformOrderNo = &v
}

// GetSerialNo returns the SerialNo field value if set, zero value otherwise.
func (o *ConfigInfoBean) GetSerialNo() string {
	if o == nil || IsNil(o.SerialNo) {
		var ret string
		return ret
	}
	return *o.SerialNo
}

// GetSerialNoOk returns a tuple with the SerialNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInfoBean) GetSerialNoOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNo) {
		return nil, false
	}
	return o.SerialNo, true
}

// HasSerialNo returns a boolean if a field has been set.
func (o *ConfigInfoBean) HasSerialNo() bool {
	if o != nil && !IsNil(o.SerialNo) {
		return true
	}

	return false
}

// SetSerialNo gets a reference to the given string and assigns it to the SerialNo field.
func (o *ConfigInfoBean) SetSerialNo(v string) {
	o.SerialNo = &v
}

// GetSignModel returns the SignModel field value if set, zero value otherwise.
func (o *ConfigInfoBean) GetSignModel() int32 {
	if o == nil || IsNil(o.SignModel) {
		var ret int32
		return ret
	}
	return *o.SignModel
}

// GetSignModelOk returns a tuple with the SignModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInfoBean) GetSignModelOk() (*int32, bool) {
	if o == nil || IsNil(o.SignModel) {
		return nil, false
	}
	return o.SignModel, true
}

// HasSignModel returns a boolean if a field has been set.
func (o *ConfigInfoBean) HasSignModel() bool {
	if o != nil && !IsNil(o.SignModel) {
		return true
	}

	return false
}

// SetSignModel gets a reference to the given int32 and assigns it to the SignModel field.
func (o *ConfigInfoBean) SetSignModel(v int32) {
	o.SignModel = &v
}

// GetSpecifyAttachmentInfos returns the SpecifyAttachmentInfos field value if set, zero value otherwise.
func (o *ConfigInfoBean) GetSpecifyAttachmentInfos() []SpecifyAttachmentInfo {
	if o == nil || IsNil(o.SpecifyAttachmentInfos) {
		var ret []SpecifyAttachmentInfo
		return ret
	}
	return o.SpecifyAttachmentInfos
}

// GetSpecifyAttachmentInfosOk returns a tuple with the SpecifyAttachmentInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigInfoBean) GetSpecifyAttachmentInfosOk() ([]SpecifyAttachmentInfo, bool) {
	if o == nil || IsNil(o.SpecifyAttachmentInfos) {
		return nil, false
	}
	return o.SpecifyAttachmentInfos, true
}

// HasSpecifyAttachmentInfos returns a boolean if a field has been set.
func (o *ConfigInfoBean) HasSpecifyAttachmentInfos() bool {
	if o != nil && !IsNil(o.SpecifyAttachmentInfos) {
		return true
	}

	return false
}

// SetSpecifyAttachmentInfos gets a reference to the given []SpecifyAttachmentInfo and assigns it to the SpecifyAttachmentInfos field.
func (o *ConfigInfoBean) SetSpecifyAttachmentInfos(v []SpecifyAttachmentInfo) {
	o.SpecifyAttachmentInfos = v
}

func (o ConfigInfoBean) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigInfoBean) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttachmentExplain) {
		toSerialize["attachment_explain"] = o.AttachmentExplain
	}
	if !IsNil(o.AutoArchive) {
		toSerialize["auto_archive"] = o.AutoArchive
	}
	if !IsNil(o.CollectAttachement) {
		toSerialize["collect_attachement"] = o.CollectAttachement
	}
	if !IsNil(o.CompanyNo) {
		toSerialize["company_no"] = o.CompanyNo
	}
	if !IsNil(o.ContractValidity) {
		toSerialize["contract_validity"] = o.ContractValidity
	}
	if !IsNil(o.JumpUrl) {
		toSerialize["jump_url"] = o.JumpUrl
	}
	if !IsNil(o.MerchantMiniSignUrl) {
		toSerialize["merchant_mini_sign_url"] = o.MerchantMiniSignUrl
	}
	if !IsNil(o.NoticeDeveloperUrl) {
		toSerialize["notice_developer_url"] = o.NoticeDeveloperUrl
	}
	if !IsNil(o.PlatformOrderNo) {
		toSerialize["platform_order_no"] = o.PlatformOrderNo
	}
	if !IsNil(o.SerialNo) {
		toSerialize["serial_no"] = o.SerialNo
	}
	if !IsNil(o.SignModel) {
		toSerialize["sign_model"] = o.SignModel
	}
	if !IsNil(o.SpecifyAttachmentInfos) {
		toSerialize["specify_attachment_infos"] = o.SpecifyAttachmentInfos
	}
	return toSerialize, nil
}

type NullableConfigInfoBean struct {
	value *ConfigInfoBean
	isSet bool
}

func (v NullableConfigInfoBean) Get() *ConfigInfoBean {
	return v.value
}

func (v *NullableConfigInfoBean) Set(val *ConfigInfoBean) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigInfoBean) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigInfoBean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigInfoBean(val *ConfigInfoBean) *NullableConfigInfoBean {
	return &NullableConfigInfoBean{value: val, isSet: true}
}

func (v NullableConfigInfoBean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigInfoBean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


