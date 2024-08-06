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

// checks if the AlipayCommerceEcEnterpriseCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceEcEnterpriseCreateModel{}

// AlipayCommerceEcEnterpriseCreateModel struct for AlipayCommerceEcEnterpriseCreateModel
type AlipayCommerceEcEnterpriseCreateModel struct {
	// 场景码
	BizScene *string `json:"biz_scene,omitempty"`
	// 是否设置管理员支付宝为企业出资账户
	CreateFundAccount *bool `json:"create_fund_account,omitempty"`
	// 是否创建企业人脸库，适用于对接团餐刷脸付、门禁刷脸通行等场景
	CreateIotGroup *bool `json:"create_iot_group,omitempty"`
	// 企业简称
	EnterpriseAlias *string `json:"enterprise_alias,omitempty"`
	// 企业名称
	EnterpriseName *string `json:"enterprise_name,omitempty"`
	// 订购【一脸通行开通插件】的小程序appId
	GroupAppId *string `json:"group_app_id,omitempty"`
	// 管理员身份标识
	Identity *string `json:"identity,omitempty"`
	// 管理员姓名
	IdentityName *string `json:"identity_name,omitempty"`
	// 管理员身份openId
	IdentityOpenId *string `json:"identity_open_id,omitempty"`
	// 管理员身份类型
	IdentityType *string `json:"identity_type,omitempty"`
	// 外部业务号
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// 企业码签约后回跳地址
	SignReturnUrl *string `json:"sign_return_url,omitempty"`
	// 管理员签约终端类型
	SignTerminal *string `json:"sign_terminal,omitempty"`
}

// NewAlipayCommerceEcEnterpriseCreateModel instantiates a new AlipayCommerceEcEnterpriseCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceEcEnterpriseCreateModel() *AlipayCommerceEcEnterpriseCreateModel {
	this := AlipayCommerceEcEnterpriseCreateModel{}
	return &this
}

// NewAlipayCommerceEcEnterpriseCreateModelWithDefaults instantiates a new AlipayCommerceEcEnterpriseCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceEcEnterpriseCreateModelWithDefaults() *AlipayCommerceEcEnterpriseCreateModel {
	this := AlipayCommerceEcEnterpriseCreateModel{}
	return &this
}

// GetBizScene returns the BizScene field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetBizScene() string {
	if o == nil || IsNil(o.BizScene) {
		var ret string
		return ret
	}
	return *o.BizScene
}

// GetBizSceneOk returns a tuple with the BizScene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetBizSceneOk() (*string, bool) {
	if o == nil || IsNil(o.BizScene) {
		return nil, false
	}
	return o.BizScene, true
}

// HasBizScene returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) HasBizScene() bool {
	if o != nil && !IsNil(o.BizScene) {
		return true
	}

	return false
}

// SetBizScene gets a reference to the given string and assigns it to the BizScene field.
func (o *AlipayCommerceEcEnterpriseCreateModel) SetBizScene(v string) {
	o.BizScene = &v
}

// GetCreateFundAccount returns the CreateFundAccount field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetCreateFundAccount() bool {
	if o == nil || IsNil(o.CreateFundAccount) {
		var ret bool
		return ret
	}
	return *o.CreateFundAccount
}

// GetCreateFundAccountOk returns a tuple with the CreateFundAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetCreateFundAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateFundAccount) {
		return nil, false
	}
	return o.CreateFundAccount, true
}

// HasCreateFundAccount returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) HasCreateFundAccount() bool {
	if o != nil && !IsNil(o.CreateFundAccount) {
		return true
	}

	return false
}

// SetCreateFundAccount gets a reference to the given bool and assigns it to the CreateFundAccount field.
func (o *AlipayCommerceEcEnterpriseCreateModel) SetCreateFundAccount(v bool) {
	o.CreateFundAccount = &v
}

// GetCreateIotGroup returns the CreateIotGroup field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetCreateIotGroup() bool {
	if o == nil || IsNil(o.CreateIotGroup) {
		var ret bool
		return ret
	}
	return *o.CreateIotGroup
}

// GetCreateIotGroupOk returns a tuple with the CreateIotGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetCreateIotGroupOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateIotGroup) {
		return nil, false
	}
	return o.CreateIotGroup, true
}

// HasCreateIotGroup returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) HasCreateIotGroup() bool {
	if o != nil && !IsNil(o.CreateIotGroup) {
		return true
	}

	return false
}

// SetCreateIotGroup gets a reference to the given bool and assigns it to the CreateIotGroup field.
func (o *AlipayCommerceEcEnterpriseCreateModel) SetCreateIotGroup(v bool) {
	o.CreateIotGroup = &v
}

// GetEnterpriseAlias returns the EnterpriseAlias field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetEnterpriseAlias() string {
	if o == nil || IsNil(o.EnterpriseAlias) {
		var ret string
		return ret
	}
	return *o.EnterpriseAlias
}

// GetEnterpriseAliasOk returns a tuple with the EnterpriseAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetEnterpriseAliasOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseAlias) {
		return nil, false
	}
	return o.EnterpriseAlias, true
}

// HasEnterpriseAlias returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) HasEnterpriseAlias() bool {
	if o != nil && !IsNil(o.EnterpriseAlias) {
		return true
	}

	return false
}

// SetEnterpriseAlias gets a reference to the given string and assigns it to the EnterpriseAlias field.
func (o *AlipayCommerceEcEnterpriseCreateModel) SetEnterpriseAlias(v string) {
	o.EnterpriseAlias = &v
}

// GetEnterpriseName returns the EnterpriseName field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetEnterpriseName() string {
	if o == nil || IsNil(o.EnterpriseName) {
		var ret string
		return ret
	}
	return *o.EnterpriseName
}

// GetEnterpriseNameOk returns a tuple with the EnterpriseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetEnterpriseNameOk() (*string, bool) {
	if o == nil || IsNil(o.EnterpriseName) {
		return nil, false
	}
	return o.EnterpriseName, true
}

// HasEnterpriseName returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) HasEnterpriseName() bool {
	if o != nil && !IsNil(o.EnterpriseName) {
		return true
	}

	return false
}

// SetEnterpriseName gets a reference to the given string and assigns it to the EnterpriseName field.
func (o *AlipayCommerceEcEnterpriseCreateModel) SetEnterpriseName(v string) {
	o.EnterpriseName = &v
}

// GetGroupAppId returns the GroupAppId field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetGroupAppId() string {
	if o == nil || IsNil(o.GroupAppId) {
		var ret string
		return ret
	}
	return *o.GroupAppId
}

// GetGroupAppIdOk returns a tuple with the GroupAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetGroupAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupAppId) {
		return nil, false
	}
	return o.GroupAppId, true
}

// HasGroupAppId returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) HasGroupAppId() bool {
	if o != nil && !IsNil(o.GroupAppId) {
		return true
	}

	return false
}

// SetGroupAppId gets a reference to the given string and assigns it to the GroupAppId field.
func (o *AlipayCommerceEcEnterpriseCreateModel) SetGroupAppId(v string) {
	o.GroupAppId = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetIdentity() string {
	if o == nil || IsNil(o.Identity) {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetIdentityOk() (*string, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *AlipayCommerceEcEnterpriseCreateModel) SetIdentity(v string) {
	o.Identity = &v
}

// GetIdentityName returns the IdentityName field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetIdentityName() string {
	if o == nil || IsNil(o.IdentityName) {
		var ret string
		return ret
	}
	return *o.IdentityName
}

// GetIdentityNameOk returns a tuple with the IdentityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetIdentityNameOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityName) {
		return nil, false
	}
	return o.IdentityName, true
}

// HasIdentityName returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) HasIdentityName() bool {
	if o != nil && !IsNil(o.IdentityName) {
		return true
	}

	return false
}

// SetIdentityName gets a reference to the given string and assigns it to the IdentityName field.
func (o *AlipayCommerceEcEnterpriseCreateModel) SetIdentityName(v string) {
	o.IdentityName = &v
}

// GetIdentityOpenId returns the IdentityOpenId field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetIdentityOpenId() string {
	if o == nil || IsNil(o.IdentityOpenId) {
		var ret string
		return ret
	}
	return *o.IdentityOpenId
}

// GetIdentityOpenIdOk returns a tuple with the IdentityOpenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetIdentityOpenIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityOpenId) {
		return nil, false
	}
	return o.IdentityOpenId, true
}

// HasIdentityOpenId returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) HasIdentityOpenId() bool {
	if o != nil && !IsNil(o.IdentityOpenId) {
		return true
	}

	return false
}

// SetIdentityOpenId gets a reference to the given string and assigns it to the IdentityOpenId field.
func (o *AlipayCommerceEcEnterpriseCreateModel) SetIdentityOpenId(v string) {
	o.IdentityOpenId = &v
}

// GetIdentityType returns the IdentityType field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetIdentityType() string {
	if o == nil || IsNil(o.IdentityType) {
		var ret string
		return ret
	}
	return *o.IdentityType
}

// GetIdentityTypeOk returns a tuple with the IdentityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetIdentityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityType) {
		return nil, false
	}
	return o.IdentityType, true
}

// HasIdentityType returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) HasIdentityType() bool {
	if o != nil && !IsNil(o.IdentityType) {
		return true
	}

	return false
}

// SetIdentityType gets a reference to the given string and assigns it to the IdentityType field.
func (o *AlipayCommerceEcEnterpriseCreateModel) SetIdentityType(v string) {
	o.IdentityType = &v
}

// GetOutBizNo returns the OutBizNo field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetOutBizNo() string {
	if o == nil || IsNil(o.OutBizNo) {
		var ret string
		return ret
	}
	return *o.OutBizNo
}

// GetOutBizNoOk returns a tuple with the OutBizNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetOutBizNoOk() (*string, bool) {
	if o == nil || IsNil(o.OutBizNo) {
		return nil, false
	}
	return o.OutBizNo, true
}

// HasOutBizNo returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) HasOutBizNo() bool {
	if o != nil && !IsNil(o.OutBizNo) {
		return true
	}

	return false
}

// SetOutBizNo gets a reference to the given string and assigns it to the OutBizNo field.
func (o *AlipayCommerceEcEnterpriseCreateModel) SetOutBizNo(v string) {
	o.OutBizNo = &v
}

// GetSignReturnUrl returns the SignReturnUrl field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetSignReturnUrl() string {
	if o == nil || IsNil(o.SignReturnUrl) {
		var ret string
		return ret
	}
	return *o.SignReturnUrl
}

// GetSignReturnUrlOk returns a tuple with the SignReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetSignReturnUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SignReturnUrl) {
		return nil, false
	}
	return o.SignReturnUrl, true
}

// HasSignReturnUrl returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) HasSignReturnUrl() bool {
	if o != nil && !IsNil(o.SignReturnUrl) {
		return true
	}

	return false
}

// SetSignReturnUrl gets a reference to the given string and assigns it to the SignReturnUrl field.
func (o *AlipayCommerceEcEnterpriseCreateModel) SetSignReturnUrl(v string) {
	o.SignReturnUrl = &v
}

// GetSignTerminal returns the SignTerminal field value if set, zero value otherwise.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetSignTerminal() string {
	if o == nil || IsNil(o.SignTerminal) {
		var ret string
		return ret
	}
	return *o.SignTerminal
}

// GetSignTerminalOk returns a tuple with the SignTerminal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) GetSignTerminalOk() (*string, bool) {
	if o == nil || IsNil(o.SignTerminal) {
		return nil, false
	}
	return o.SignTerminal, true
}

// HasSignTerminal returns a boolean if a field has been set.
func (o *AlipayCommerceEcEnterpriseCreateModel) HasSignTerminal() bool {
	if o != nil && !IsNil(o.SignTerminal) {
		return true
	}

	return false
}

// SetSignTerminal gets a reference to the given string and assigns it to the SignTerminal field.
func (o *AlipayCommerceEcEnterpriseCreateModel) SetSignTerminal(v string) {
	o.SignTerminal = &v
}

func (o AlipayCommerceEcEnterpriseCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceEcEnterpriseCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizScene) {
		toSerialize["biz_scene"] = o.BizScene
	}
	if !IsNil(o.CreateFundAccount) {
		toSerialize["create_fund_account"] = o.CreateFundAccount
	}
	if !IsNil(o.CreateIotGroup) {
		toSerialize["create_iot_group"] = o.CreateIotGroup
	}
	if !IsNil(o.EnterpriseAlias) {
		toSerialize["enterprise_alias"] = o.EnterpriseAlias
	}
	if !IsNil(o.EnterpriseName) {
		toSerialize["enterprise_name"] = o.EnterpriseName
	}
	if !IsNil(o.GroupAppId) {
		toSerialize["group_app_id"] = o.GroupAppId
	}
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !IsNil(o.IdentityName) {
		toSerialize["identity_name"] = o.IdentityName
	}
	if !IsNil(o.IdentityOpenId) {
		toSerialize["identity_open_id"] = o.IdentityOpenId
	}
	if !IsNil(o.IdentityType) {
		toSerialize["identity_type"] = o.IdentityType
	}
	if !IsNil(o.OutBizNo) {
		toSerialize["out_biz_no"] = o.OutBizNo
	}
	if !IsNil(o.SignReturnUrl) {
		toSerialize["sign_return_url"] = o.SignReturnUrl
	}
	if !IsNil(o.SignTerminal) {
		toSerialize["sign_terminal"] = o.SignTerminal
	}
	return toSerialize, nil
}

type NullableAlipayCommerceEcEnterpriseCreateModel struct {
	value *AlipayCommerceEcEnterpriseCreateModel
	isSet bool
}

func (v NullableAlipayCommerceEcEnterpriseCreateModel) Get() *AlipayCommerceEcEnterpriseCreateModel {
	return v.value
}

func (v *NullableAlipayCommerceEcEnterpriseCreateModel) Set(val *AlipayCommerceEcEnterpriseCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEcEnterpriseCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEcEnterpriseCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEcEnterpriseCreateModel(val *AlipayCommerceEcEnterpriseCreateModel) *NullableAlipayCommerceEcEnterpriseCreateModel {
	return &NullableAlipayCommerceEcEnterpriseCreateModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceEcEnterpriseCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEcEnterpriseCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
