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

// checks if the AlipayOpenSearchAppkeywordApplyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSearchAppkeywordApplyModel{}

// AlipayOpenSearchAppkeywordApplyModel struct for AlipayOpenSearchAppkeywordApplyModel
type AlipayOpenSearchAppkeywordApplyModel struct {
	// 关键词配置id，由支付宝生成，关键词申请通过后会通知接口返回，也可以申请单状态获取
	ConfigId *string `json:"config_id,omitempty"`
	// 提报的关键词
	KeyWords []string `json:"key_words,omitempty"`
	// 小程序id
	TargetAppid *string `json:"target_appid,omitempty"`
}

// NewAlipayOpenSearchAppkeywordApplyModel instantiates a new AlipayOpenSearchAppkeywordApplyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSearchAppkeywordApplyModel() *AlipayOpenSearchAppkeywordApplyModel {
	this := AlipayOpenSearchAppkeywordApplyModel{}
	return &this
}

// NewAlipayOpenSearchAppkeywordApplyModelWithDefaults instantiates a new AlipayOpenSearchAppkeywordApplyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSearchAppkeywordApplyModelWithDefaults() *AlipayOpenSearchAppkeywordApplyModel {
	this := AlipayOpenSearchAppkeywordApplyModel{}
	return &this
}

// GetConfigId returns the ConfigId field value if set, zero value otherwise.
func (o *AlipayOpenSearchAppkeywordApplyModel) GetConfigId() string {
	if o == nil || IsNil(o.ConfigId) {
		var ret string
		return ret
	}
	return *o.ConfigId
}

// GetConfigIdOk returns a tuple with the ConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchAppkeywordApplyModel) GetConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigId) {
		return nil, false
	}
	return o.ConfigId, true
}

// HasConfigId returns a boolean if a field has been set.
func (o *AlipayOpenSearchAppkeywordApplyModel) HasConfigId() bool {
	if o != nil && !IsNil(o.ConfigId) {
		return true
	}

	return false
}

// SetConfigId gets a reference to the given string and assigns it to the ConfigId field.
func (o *AlipayOpenSearchAppkeywordApplyModel) SetConfigId(v string) {
	o.ConfigId = &v
}

// GetKeyWords returns the KeyWords field value if set, zero value otherwise.
func (o *AlipayOpenSearchAppkeywordApplyModel) GetKeyWords() []string {
	if o == nil || IsNil(o.KeyWords) {
		var ret []string
		return ret
	}
	return o.KeyWords
}

// GetKeyWordsOk returns a tuple with the KeyWords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchAppkeywordApplyModel) GetKeyWordsOk() ([]string, bool) {
	if o == nil || IsNil(o.KeyWords) {
		return nil, false
	}
	return o.KeyWords, true
}

// HasKeyWords returns a boolean if a field has been set.
func (o *AlipayOpenSearchAppkeywordApplyModel) HasKeyWords() bool {
	if o != nil && !IsNil(o.KeyWords) {
		return true
	}

	return false
}

// SetKeyWords gets a reference to the given []string and assigns it to the KeyWords field.
func (o *AlipayOpenSearchAppkeywordApplyModel) SetKeyWords(v []string) {
	o.KeyWords = v
}

// GetTargetAppid returns the TargetAppid field value if set, zero value otherwise.
func (o *AlipayOpenSearchAppkeywordApplyModel) GetTargetAppid() string {
	if o == nil || IsNil(o.TargetAppid) {
		var ret string
		return ret
	}
	return *o.TargetAppid
}

// GetTargetAppidOk returns a tuple with the TargetAppid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchAppkeywordApplyModel) GetTargetAppidOk() (*string, bool) {
	if o == nil || IsNil(o.TargetAppid) {
		return nil, false
	}
	return o.TargetAppid, true
}

// HasTargetAppid returns a boolean if a field has been set.
func (o *AlipayOpenSearchAppkeywordApplyModel) HasTargetAppid() bool {
	if o != nil && !IsNil(o.TargetAppid) {
		return true
	}

	return false
}

// SetTargetAppid gets a reference to the given string and assigns it to the TargetAppid field.
func (o *AlipayOpenSearchAppkeywordApplyModel) SetTargetAppid(v string) {
	o.TargetAppid = &v
}

func (o AlipayOpenSearchAppkeywordApplyModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSearchAppkeywordApplyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConfigId) {
		toSerialize["config_id"] = o.ConfigId
	}
	if !IsNil(o.KeyWords) {
		toSerialize["key_words"] = o.KeyWords
	}
	if !IsNil(o.TargetAppid) {
		toSerialize["target_appid"] = o.TargetAppid
	}
	return toSerialize, nil
}

type NullableAlipayOpenSearchAppkeywordApplyModel struct {
	value *AlipayOpenSearchAppkeywordApplyModel
	isSet bool
}

func (v NullableAlipayOpenSearchAppkeywordApplyModel) Get() *AlipayOpenSearchAppkeywordApplyModel {
	return v.value
}

func (v *NullableAlipayOpenSearchAppkeywordApplyModel) Set(val *AlipayOpenSearchAppkeywordApplyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchAppkeywordApplyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchAppkeywordApplyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchAppkeywordApplyModel(val *AlipayOpenSearchAppkeywordApplyModel) *NullableAlipayOpenSearchAppkeywordApplyModel {
	return &NullableAlipayOpenSearchAppkeywordApplyModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchAppkeywordApplyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchAppkeywordApplyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
