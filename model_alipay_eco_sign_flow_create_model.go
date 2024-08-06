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

// checks if the AlipayEcoSignFlowCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEcoSignFlowCreateModel{}

// AlipayEcoSignFlowCreateModel struct for AlipayEcoSignFlowCreateModel
type AlipayEcoSignFlowCreateModel struct {
	// 附件信息
	Attachments []Attachment `json:"attachments,omitempty"`
	// 流程主题
	BusinessScene *string     `json:"business_scene,omitempty"`
	ConfigInfo    *ConfigInfo `json:"config_info,omitempty"`
	// 模板信息
	TemplateInfos []TemplateInfo `json:"template_infos,omitempty"`
}

// NewAlipayEcoSignFlowCreateModel instantiates a new AlipayEcoSignFlowCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEcoSignFlowCreateModel() *AlipayEcoSignFlowCreateModel {
	this := AlipayEcoSignFlowCreateModel{}
	return &this
}

// NewAlipayEcoSignFlowCreateModelWithDefaults instantiates a new AlipayEcoSignFlowCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEcoSignFlowCreateModelWithDefaults() *AlipayEcoSignFlowCreateModel {
	this := AlipayEcoSignFlowCreateModel{}
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *AlipayEcoSignFlowCreateModel) GetAttachments() []Attachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoSignFlowCreateModel) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *AlipayEcoSignFlowCreateModel) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *AlipayEcoSignFlowCreateModel) SetAttachments(v []Attachment) {
	o.Attachments = v
}

// GetBusinessScene returns the BusinessScene field value if set, zero value otherwise.
func (o *AlipayEcoSignFlowCreateModel) GetBusinessScene() string {
	if o == nil || IsNil(o.BusinessScene) {
		var ret string
		return ret
	}
	return *o.BusinessScene
}

// GetBusinessSceneOk returns a tuple with the BusinessScene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoSignFlowCreateModel) GetBusinessSceneOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessScene) {
		return nil, false
	}
	return o.BusinessScene, true
}

// HasBusinessScene returns a boolean if a field has been set.
func (o *AlipayEcoSignFlowCreateModel) HasBusinessScene() bool {
	if o != nil && !IsNil(o.BusinessScene) {
		return true
	}

	return false
}

// SetBusinessScene gets a reference to the given string and assigns it to the BusinessScene field.
func (o *AlipayEcoSignFlowCreateModel) SetBusinessScene(v string) {
	o.BusinessScene = &v
}

// GetConfigInfo returns the ConfigInfo field value if set, zero value otherwise.
func (o *AlipayEcoSignFlowCreateModel) GetConfigInfo() ConfigInfo {
	if o == nil || IsNil(o.ConfigInfo) {
		var ret ConfigInfo
		return ret
	}
	return *o.ConfigInfo
}

// GetConfigInfoOk returns a tuple with the ConfigInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoSignFlowCreateModel) GetConfigInfoOk() (*ConfigInfo, bool) {
	if o == nil || IsNil(o.ConfigInfo) {
		return nil, false
	}
	return o.ConfigInfo, true
}

// HasConfigInfo returns a boolean if a field has been set.
func (o *AlipayEcoSignFlowCreateModel) HasConfigInfo() bool {
	if o != nil && !IsNil(o.ConfigInfo) {
		return true
	}

	return false
}

// SetConfigInfo gets a reference to the given ConfigInfo and assigns it to the ConfigInfo field.
func (o *AlipayEcoSignFlowCreateModel) SetConfigInfo(v ConfigInfo) {
	o.ConfigInfo = &v
}

// GetTemplateInfos returns the TemplateInfos field value if set, zero value otherwise.
func (o *AlipayEcoSignFlowCreateModel) GetTemplateInfos() []TemplateInfo {
	if o == nil || IsNil(o.TemplateInfos) {
		var ret []TemplateInfo
		return ret
	}
	return o.TemplateInfos
}

// GetTemplateInfosOk returns a tuple with the TemplateInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEcoSignFlowCreateModel) GetTemplateInfosOk() ([]TemplateInfo, bool) {
	if o == nil || IsNil(o.TemplateInfos) {
		return nil, false
	}
	return o.TemplateInfos, true
}

// HasTemplateInfos returns a boolean if a field has been set.
func (o *AlipayEcoSignFlowCreateModel) HasTemplateInfos() bool {
	if o != nil && !IsNil(o.TemplateInfos) {
		return true
	}

	return false
}

// SetTemplateInfos gets a reference to the given []TemplateInfo and assigns it to the TemplateInfos field.
func (o *AlipayEcoSignFlowCreateModel) SetTemplateInfos(v []TemplateInfo) {
	o.TemplateInfos = v
}

func (o AlipayEcoSignFlowCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEcoSignFlowCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.BusinessScene) {
		toSerialize["business_scene"] = o.BusinessScene
	}
	if !IsNil(o.ConfigInfo) {
		toSerialize["config_info"] = o.ConfigInfo
	}
	if !IsNil(o.TemplateInfos) {
		toSerialize["template_infos"] = o.TemplateInfos
	}
	return toSerialize, nil
}

type NullableAlipayEcoSignFlowCreateModel struct {
	value *AlipayEcoSignFlowCreateModel
	isSet bool
}

func (v NullableAlipayEcoSignFlowCreateModel) Get() *AlipayEcoSignFlowCreateModel {
	return v.value
}

func (v *NullableAlipayEcoSignFlowCreateModel) Set(val *AlipayEcoSignFlowCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEcoSignFlowCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEcoSignFlowCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEcoSignFlowCreateModel(val *AlipayEcoSignFlowCreateModel) *NullableAlipayEcoSignFlowCreateModel {
	return &NullableAlipayEcoSignFlowCreateModel{value: val, isSet: true}
}

func (v NullableAlipayEcoSignFlowCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEcoSignFlowCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
