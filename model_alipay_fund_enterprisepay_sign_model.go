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

// checks if the AlipayFundEnterprisepaySignModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayFundEnterprisepaySignModel{}

// AlipayFundEnterprisepaySignModel struct for AlipayFundEnterprisepaySignModel
type AlipayFundEnterprisepaySignModel struct {
	// 企业简称，传空采用默认规则命名，使用公司名称
	AccountName *string `json:"account_name,omitempty"`
	// 指定签约终端类型 PC-电脑浏览器 MOBILE-手机端
	AppointSignTerminal *string `json:"appoint_sign_terminal,omitempty"`
	// 业务场景，联系支付宝分配
	BizScene *string `json:"biz_scene,omitempty"`
	// 特殊场景下扩展字段
	ExtParams *string `json:"ext_params,omitempty"`
	// 开户账号： 当identity_type是ALIPAY_USER_ID时填支付宝会员ID（2088开头）； 当identity_type是ALIPAY_LOGON_ID 时填支付宝登录号； 当identity_type是OUT_USER_ID时填外部平台的用户uid; 当identity_type是ALIPAY_OPEN_ID时填用户openId
	Identity *string `json:"identity,omitempty"`
	// 开户用户名称，identity_type为ALIPAY_LOGON_ID时必填
	IdentityName *string `json:"identity_name,omitempty"`
	// 账号类型，目前支持如下类型： 1、ALIPAY_USER_ID 支付宝的会员ID 2、ALIPAY_LOGON_ID：支付宝登录号，支持邮箱和手机号格式 3、OUT_USER_ID：外部用户uid
	IdentityType *string `json:"identity_type,omitempty"`
	// 商户侧单号（幂等字段） ，补充说明：幂等逻辑（商户appid+out_biz_no），商户集成时需注意，如换号重复发起，则认为是一笔新的请求
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 外部实体ID，一般表述为外部操作人操作某个外部实体。比如企业码2284号
	OutEntityId *string `json:"out_entity_id,omitempty"`
	// ISV平台定义的共同账户开户来源
	OutSource *string `json:"out_source,omitempty"`
	// 产品码
	ProductCode *string `json:"product_code,omitempty"`
}

// NewAlipayFundEnterprisepaySignModel instantiates a new AlipayFundEnterprisepaySignModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayFundEnterprisepaySignModel() *AlipayFundEnterprisepaySignModel {
	this := AlipayFundEnterprisepaySignModel{}
	return &this
}

// NewAlipayFundEnterprisepaySignModelWithDefaults instantiates a new AlipayFundEnterprisepaySignModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayFundEnterprisepaySignModelWithDefaults() *AlipayFundEnterprisepaySignModel {
	this := AlipayFundEnterprisepaySignModel{}
	return &this
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepaySignModel) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepaySignModel) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepaySignModel) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *AlipayFundEnterprisepaySignModel) SetAccountName(v string) {
	o.AccountName = &v
}

// GetAppointSignTerminal returns the AppointSignTerminal field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepaySignModel) GetAppointSignTerminal() string {
	if o == nil || IsNil(o.AppointSignTerminal) {
		var ret string
		return ret
	}
	return *o.AppointSignTerminal
}

// GetAppointSignTerminalOk returns a tuple with the AppointSignTerminal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepaySignModel) GetAppointSignTerminalOk() (*string, bool) {
	if o == nil || IsNil(o.AppointSignTerminal) {
		return nil, false
	}
	return o.AppointSignTerminal, true
}

// HasAppointSignTerminal returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepaySignModel) HasAppointSignTerminal() bool {
	if o != nil && !IsNil(o.AppointSignTerminal) {
		return true
	}

	return false
}

// SetAppointSignTerminal gets a reference to the given string and assigns it to the AppointSignTerminal field.
func (o *AlipayFundEnterprisepaySignModel) SetAppointSignTerminal(v string) {
	o.AppointSignTerminal = &v
}

// GetBizScene returns the BizScene field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepaySignModel) GetBizScene() string {
	if o == nil || IsNil(o.BizScene) {
		var ret string
		return ret
	}
	return *o.BizScene
}

// GetBizSceneOk returns a tuple with the BizScene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepaySignModel) GetBizSceneOk() (*string, bool) {
	if o == nil || IsNil(o.BizScene) {
		return nil, false
	}
	return o.BizScene, true
}

// HasBizScene returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepaySignModel) HasBizScene() bool {
	if o != nil && !IsNil(o.BizScene) {
		return true
	}

	return false
}

// SetBizScene gets a reference to the given string and assigns it to the BizScene field.
func (o *AlipayFundEnterprisepaySignModel) SetBizScene(v string) {
	o.BizScene = &v
}

// GetExtParams returns the ExtParams field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepaySignModel) GetExtParams() string {
	if o == nil || IsNil(o.ExtParams) {
		var ret string
		return ret
	}
	return *o.ExtParams
}

// GetExtParamsOk returns a tuple with the ExtParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepaySignModel) GetExtParamsOk() (*string, bool) {
	if o == nil || IsNil(o.ExtParams) {
		return nil, false
	}
	return o.ExtParams, true
}

// HasExtParams returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepaySignModel) HasExtParams() bool {
	if o != nil && !IsNil(o.ExtParams) {
		return true
	}

	return false
}

// SetExtParams gets a reference to the given string and assigns it to the ExtParams field.
func (o *AlipayFundEnterprisepaySignModel) SetExtParams(v string) {
	o.ExtParams = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepaySignModel) GetIdentity() string {
	if o == nil || IsNil(o.Identity) {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepaySignModel) GetIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepaySignModel) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *AlipayFundEnterprisepaySignModel) SetIdentity(v string) {
	o.Identity = &v
}

// GetIdentityName returns the IdentityName field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepaySignModel) GetIdentityName() string {
	if o == nil || IsNil(o.IdentityName) {
		var ret string
		return ret
	}
	return *o.IdentityName
}

// GetIdentityNameOk returns a tuple with the IdentityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepaySignModel) GetIdentityNameOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityName) {
		return nil, false
	}
	return o.IdentityName, true
}

// HasIdentityName returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepaySignModel) HasIdentityName() bool {
	if o != nil && !IsNil(o.IdentityName) {
		return true
	}

	return false
}

// SetIdentityName gets a reference to the given string and assigns it to the IdentityName field.
func (o *AlipayFundEnterprisepaySignModel) SetIdentityName(v string) {
	o.IdentityName = &v
}

// GetIdentityType returns the IdentityType field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepaySignModel) GetIdentityType() string {
	if o == nil || IsNil(o.IdentityType) {
		var ret string
		return ret
	}
	return *o.IdentityType
}

// GetIdentityTypeOk returns a tuple with the IdentityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepaySignModel) GetIdentityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityType) {
		return nil, false
	}
	return o.IdentityType, true
}

// HasIdentityType returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepaySignModel) HasIdentityType() bool {
	if o != nil && !IsNil(o.IdentityType) {
		return true
	}

	return false
}

// SetIdentityType gets a reference to the given string and assigns it to the IdentityType field.
func (o *AlipayFundEnterprisepaySignModel) SetIdentityType(v string) {
	o.IdentityType = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepaySignModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepaySignModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepaySignModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayFundEnterprisepaySignModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetOutEntityId returns the OutEntityId field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepaySignModel) GetOutEntityId() string {
	if o == nil || IsNil(o.OutEntityId) {
		var ret string
		return ret
	}
	return *o.OutEntityId
}

// GetOutEntityIdOk returns a tuple with the OutEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepaySignModel) GetOutEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.OutEntityId) {
		return nil, false
	}
	return o.OutEntityId, true
}

// HasOutEntityId returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepaySignModel) HasOutEntityId() bool {
	if o != nil && !IsNil(o.OutEntityId) {
		return true
	}

	return false
}

// SetOutEntityId gets a reference to the given string and assigns it to the OutEntityId field.
func (o *AlipayFundEnterprisepaySignModel) SetOutEntityId(v string) {
	o.OutEntityId = &v
}

// GetOutSource returns the OutSource field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepaySignModel) GetOutSource() string {
	if o == nil || IsNil(o.OutSource) {
		var ret string
		return ret
	}
	return *o.OutSource
}

// GetOutSourceOk returns a tuple with the OutSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepaySignModel) GetOutSourceOk() (*string, bool) {
	if o == nil || IsNil(o.OutSource) {
		return nil, false
	}
	return o.OutSource, true
}

// HasOutSource returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepaySignModel) HasOutSource() bool {
	if o != nil && !IsNil(o.OutSource) {
		return true
	}

	return false
}

// SetOutSource gets a reference to the given string and assigns it to the OutSource field.
func (o *AlipayFundEnterprisepaySignModel) SetOutSource(v string) {
	o.OutSource = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *AlipayFundEnterprisepaySignModel) GetProductCode() string {
	if o == nil || IsNil(o.ProductCode) {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayFundEnterprisepaySignModel) GetProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductCode) {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *AlipayFundEnterprisepaySignModel) HasProductCode() bool {
	if o != nil && !IsNil(o.ProductCode) {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *AlipayFundEnterprisepaySignModel) SetProductCode(v string) {
	o.ProductCode = &v
}

func (o AlipayFundEnterprisepaySignModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayFundEnterprisepaySignModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountName) {
		toSerialize["account_name"] = o.AccountName
	}
	if !IsNil(o.AppointSignTerminal) {
		toSerialize["appoint_sign_terminal"] = o.AppointSignTerminal
	}
	if !IsNil(o.BizScene) {
		toSerialize["biz_scene"] = o.BizScene
	}
	if !IsNil(o.ExtParams) {
		toSerialize["ext_params"] = o.ExtParams
	}
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !IsNil(o.IdentityName) {
		toSerialize["identity_name"] = o.IdentityName
	}
	if !IsNil(o.IdentityType) {
		toSerialize["identity_type"] = o.IdentityType
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.OutEntityId) {
		toSerialize["out_entity_id"] = o.OutEntityId
	}
	if !IsNil(o.OutSource) {
		toSerialize["out_source"] = o.OutSource
	}
	if !IsNil(o.ProductCode) {
		toSerialize["product_code"] = o.ProductCode
	}
	return toSerialize, nil
}

type NullableAlipayFundEnterprisepaySignModel struct {
	value *AlipayFundEnterprisepaySignModel
	isSet bool
}

func (v NullableAlipayFundEnterprisepaySignModel) Get() *AlipayFundEnterprisepaySignModel {
	return v.value
}

func (v *NullableAlipayFundEnterprisepaySignModel) Set(val *AlipayFundEnterprisepaySignModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundEnterprisepaySignModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundEnterprisepaySignModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundEnterprisepaySignModel(val *AlipayFundEnterprisepaySignModel) *NullableAlipayFundEnterprisepaySignModel {
	return &NullableAlipayFundEnterprisepaySignModel{value: val, isSet: true}
}

func (v NullableAlipayFundEnterprisepaySignModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundEnterprisepaySignModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


