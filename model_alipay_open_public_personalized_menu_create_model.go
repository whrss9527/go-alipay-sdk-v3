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

// checks if the AlipayOpenPublicPersonalizedMenuCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicPersonalizedMenuCreateModel{}

// AlipayOpenPublicPersonalizedMenuCreateModel struct for AlipayOpenPublicPersonalizedMenuCreateModel
type AlipayOpenPublicPersonalizedMenuCreateModel struct {
	// 一级菜单列表。说明： * 如果是文本菜单，最多有4个一级菜单，若开发者在后台打开了\"咨询反馈\"的开关，则只能有3个一级菜单。 * 如果是 ICON 菜单信息，最多有80个一级菜单(忽略二级菜单)。
	Button []ButtonObject `json:"button,omitempty"`
	// 人群分组id。分组创建及管理接入详情参见 <a href=\"https://opendocs.alipay.com/fw/api/106931\">分组管理</a>。 注意：group_id 与 label_rule（标签规则）不能同时为空。
	GroupId *string `json:"group_id,omitempty"`
	// 标签规则，目前限定只能传入1条，在个性化菜单创建成功后，满足该标签规则的用户进入生活号首页，将看到该套菜单。生成标签及打标详情参见<a href=\"https://opendocs.alipay.com/fw/api/106877\">标签管理</a>。 注意：group_id（人群分组id） 与 label_rule 不能同时为空。
	LabelRule []LabelRule `json:"label_rule,omitempty"`
	// 手机客户端类型，枚举支持： *iphone； *android； *wp； 说明：不填为不区分机型。
	MobileClientType *string `json:"mobile_client_type,omitempty"`
	// 菜单类型，不填时默认为 text（文本型菜单）。枚举值如下： * text：文本型菜单。 * icon：表示 icon 型菜单，当传值为\"icon\"时，菜单节点的icon字段必传。
	Type *string `json:"type,omitempty"`
}

// NewAlipayOpenPublicPersonalizedMenuCreateModel instantiates a new AlipayOpenPublicPersonalizedMenuCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicPersonalizedMenuCreateModel() *AlipayOpenPublicPersonalizedMenuCreateModel {
	this := AlipayOpenPublicPersonalizedMenuCreateModel{}
	return &this
}

// NewAlipayOpenPublicPersonalizedMenuCreateModelWithDefaults instantiates a new AlipayOpenPublicPersonalizedMenuCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicPersonalizedMenuCreateModelWithDefaults() *AlipayOpenPublicPersonalizedMenuCreateModel {
	this := AlipayOpenPublicPersonalizedMenuCreateModel{}
	return &this
}

// GetButton returns the Button field value if set, zero value otherwise.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) GetButton() []ButtonObject {
	if o == nil || IsNil(o.Button) {
		var ret []ButtonObject
		return ret
	}
	return o.Button
}

// GetButtonOk returns a tuple with the Button field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) GetButtonOk() ([]ButtonObject, bool) {
	if o == nil || IsNil(o.Button) {
		return nil, false
	}
	return o.Button, true
}

// HasButton returns a boolean if a field has been set.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) HasButton() bool {
	if o != nil && !IsNil(o.Button) {
		return true
	}

	return false
}

// SetButton gets a reference to the given []ButtonObject and assigns it to the Button field.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) SetButton(v []ButtonObject) {
	o.Button = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) SetGroupId(v string) {
	o.GroupId = &v
}

// GetLabelRule returns the LabelRule field value if set, zero value otherwise.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) GetLabelRule() []LabelRule {
	if o == nil || IsNil(o.LabelRule) {
		var ret []LabelRule
		return ret
	}
	return o.LabelRule
}

// GetLabelRuleOk returns a tuple with the LabelRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) GetLabelRuleOk() ([]LabelRule, bool) {
	if o == nil || IsNil(o.LabelRule) {
		return nil, false
	}
	return o.LabelRule, true
}

// HasLabelRule returns a boolean if a field has been set.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) HasLabelRule() bool {
	if o != nil && !IsNil(o.LabelRule) {
		return true
	}

	return false
}

// SetLabelRule gets a reference to the given []LabelRule and assigns it to the LabelRule field.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) SetLabelRule(v []LabelRule) {
	o.LabelRule = v
}

// GetMobileClientType returns the MobileClientType field value if set, zero value otherwise.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) GetMobileClientType() string {
	if o == nil || IsNil(o.MobileClientType) {
		var ret string
		return ret
	}
	return *o.MobileClientType
}

// GetMobileClientTypeOk returns a tuple with the MobileClientType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) GetMobileClientTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MobileClientType) {
		return nil, false
	}
	return o.MobileClientType, true
}

// HasMobileClientType returns a boolean if a field has been set.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) HasMobileClientType() bool {
	if o != nil && !IsNil(o.MobileClientType) {
		return true
	}

	return false
}

// SetMobileClientType gets a reference to the given string and assigns it to the MobileClientType field.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) SetMobileClientType(v string) {
	o.MobileClientType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AlipayOpenPublicPersonalizedMenuCreateModel) SetType(v string) {
	o.Type = &v
}

func (o AlipayOpenPublicPersonalizedMenuCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicPersonalizedMenuCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Button) {
		toSerialize["button"] = o.Button
	}
	if !IsNil(o.GroupId) {
		toSerialize["group_id"] = o.GroupId
	}
	if !IsNil(o.LabelRule) {
		toSerialize["label_rule"] = o.LabelRule
	}
	if !IsNil(o.MobileClientType) {
		toSerialize["mobile_client_type"] = o.MobileClientType
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicPersonalizedMenuCreateModel struct {
	value *AlipayOpenPublicPersonalizedMenuCreateModel
	isSet bool
}

func (v NullableAlipayOpenPublicPersonalizedMenuCreateModel) Get() *AlipayOpenPublicPersonalizedMenuCreateModel {
	return v.value
}

func (v *NullableAlipayOpenPublicPersonalizedMenuCreateModel) Set(val *AlipayOpenPublicPersonalizedMenuCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicPersonalizedMenuCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicPersonalizedMenuCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicPersonalizedMenuCreateModel(val *AlipayOpenPublicPersonalizedMenuCreateModel) *NullableAlipayOpenPublicPersonalizedMenuCreateModel {
	return &NullableAlipayOpenPublicPersonalizedMenuCreateModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicPersonalizedMenuCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicPersonalizedMenuCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
