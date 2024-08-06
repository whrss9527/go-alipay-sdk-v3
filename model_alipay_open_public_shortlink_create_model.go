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

// checks if the AlipayOpenPublicShortlinkCreateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicShortlinkCreateModel{}

// AlipayOpenPublicShortlinkCreateModel struct for AlipayOpenPublicShortlinkCreateModel
type AlipayOpenPublicShortlinkCreateModel struct {
	// 对于场景 ID 的描述，由商户自定义。
	Remark *string `json:"remark,omitempty"`
	// 短链接对应的场景 ID。由商户自定义，仅支持数字、字母及下划线。
	SceneId *string `json:"scene_id,omitempty"`
}

// NewAlipayOpenPublicShortlinkCreateModel instantiates a new AlipayOpenPublicShortlinkCreateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicShortlinkCreateModel() *AlipayOpenPublicShortlinkCreateModel {
	this := AlipayOpenPublicShortlinkCreateModel{}
	return &this
}

// NewAlipayOpenPublicShortlinkCreateModelWithDefaults instantiates a new AlipayOpenPublicShortlinkCreateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicShortlinkCreateModelWithDefaults() *AlipayOpenPublicShortlinkCreateModel {
	this := AlipayOpenPublicShortlinkCreateModel{}
	return &this
}

// GetRemark returns the Remark field value if set, zero value otherwise.
func (o *AlipayOpenPublicShortlinkCreateModel) GetRemark() string {
	if o == nil || IsNil(o.Remark) {
		var ret string
		return ret
	}
	return *o.Remark
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicShortlinkCreateModel) GetRemarkOk() (*string, bool) {
	if o == nil || IsNil(o.Remark) {
		return nil, false
	}
	return o.Remark, true
}

// HasRemark returns a boolean if a field has been set.
func (o *AlipayOpenPublicShortlinkCreateModel) HasRemark() bool {
	if o != nil && !IsNil(o.Remark) {
		return true
	}

	return false
}

// SetRemark gets a reference to the given string and assigns it to the Remark field.
func (o *AlipayOpenPublicShortlinkCreateModel) SetRemark(v string) {
	o.Remark = &v
}

// GetSceneId returns the SceneId field value if set, zero value otherwise.
func (o *AlipayOpenPublicShortlinkCreateModel) GetSceneId() string {
	if o == nil || IsNil(o.SceneId) {
		var ret string
		return ret
	}
	return *o.SceneId
}

// GetSceneIdOk returns a tuple with the SceneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicShortlinkCreateModel) GetSceneIdOk() (*string, bool) {
	if o == nil || IsNil(o.SceneId) {
		return nil, false
	}
	return o.SceneId, true
}

// HasSceneId returns a boolean if a field has been set.
func (o *AlipayOpenPublicShortlinkCreateModel) HasSceneId() bool {
	if o != nil && !IsNil(o.SceneId) {
		return true
	}

	return false
}

// SetSceneId gets a reference to the given string and assigns it to the SceneId field.
func (o *AlipayOpenPublicShortlinkCreateModel) SetSceneId(v string) {
	o.SceneId = &v
}

func (o AlipayOpenPublicShortlinkCreateModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicShortlinkCreateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Remark) {
		toSerialize["remark"] = o.Remark
	}
	if !IsNil(o.SceneId) {
		toSerialize["scene_id"] = o.SceneId
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicShortlinkCreateModel struct {
	value *AlipayOpenPublicShortlinkCreateModel
	isSet bool
}

func (v NullableAlipayOpenPublicShortlinkCreateModel) Get() *AlipayOpenPublicShortlinkCreateModel {
	return v.value
}

func (v *NullableAlipayOpenPublicShortlinkCreateModel) Set(val *AlipayOpenPublicShortlinkCreateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicShortlinkCreateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicShortlinkCreateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicShortlinkCreateModel(val *AlipayOpenPublicShortlinkCreateModel) *NullableAlipayOpenPublicShortlinkCreateModel {
	return &NullableAlipayOpenPublicShortlinkCreateModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicShortlinkCreateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicShortlinkCreateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
