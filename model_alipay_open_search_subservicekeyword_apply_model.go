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

// checks if the AlipayOpenSearchSubservicekeywordApplyModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSearchSubservicekeywordApplyModel{}

// AlipayOpenSearchSubservicekeywordApplyModel struct for AlipayOpenSearchSubservicekeywordApplyModel
type AlipayOpenSearchSubservicekeywordApplyModel struct {
	// 关键词配置id，由支付宝生成，关键词申请通过后会通知接口返回，也可以申请单状态获取
	ConfigId *string `json:"config_id,omitempty"`
	// 服务关键词列表，每批最多传入30个，关键词长度小于12个汉字。 超过数量限制会申请失败，剩余关键词可通过 <a href=\"https://opendocs.alipay.com/mini/062ndt?pathHash=e3e78b68&ref=api&scene=common\">alipay.open.search.appkeywordquota.query</a>查询
	KeyWords []string `json:"key_words,omitempty"`
	// 子服务code，提报服务关键词，alipay.open.app.service.list.query(服务批量查询)这个接口可以获取
	SubServiceCode *string `json:"sub_service_code,omitempty"`
	// 小程序id
	TargetAppid *string `json:"target_appid,omitempty"`
}

// NewAlipayOpenSearchSubservicekeywordApplyModel instantiates a new AlipayOpenSearchSubservicekeywordApplyModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSearchSubservicekeywordApplyModel() *AlipayOpenSearchSubservicekeywordApplyModel {
	this := AlipayOpenSearchSubservicekeywordApplyModel{}
	return &this
}

// NewAlipayOpenSearchSubservicekeywordApplyModelWithDefaults instantiates a new AlipayOpenSearchSubservicekeywordApplyModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSearchSubservicekeywordApplyModelWithDefaults() *AlipayOpenSearchSubservicekeywordApplyModel {
	this := AlipayOpenSearchSubservicekeywordApplyModel{}
	return &this
}

// GetConfigId returns the ConfigId field value if set, zero value otherwise.
func (o *AlipayOpenSearchSubservicekeywordApplyModel) GetConfigId() string {
	if o == nil || IsNil(o.ConfigId) {
		var ret string
		return ret
	}
	return *o.ConfigId
}

// GetConfigIdOk returns a tuple with the ConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchSubservicekeywordApplyModel) GetConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigId) {
		return nil, false
	}
	return o.ConfigId, true
}

// HasConfigId returns a boolean if a field has been set.
func (o *AlipayOpenSearchSubservicekeywordApplyModel) HasConfigId() bool {
	if o != nil && !IsNil(o.ConfigId) {
		return true
	}

	return false
}

// SetConfigId gets a reference to the given string and assigns it to the ConfigId field.
func (o *AlipayOpenSearchSubservicekeywordApplyModel) SetConfigId(v string) {
	o.ConfigId = &v
}

// GetKeyWords returns the KeyWords field value if set, zero value otherwise.
func (o *AlipayOpenSearchSubservicekeywordApplyModel) GetKeyWords() []string {
	if o == nil || IsNil(o.KeyWords) {
		var ret []string
		return ret
	}
	return o.KeyWords
}

// GetKeyWordsOk returns a tuple with the KeyWords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchSubservicekeywordApplyModel) GetKeyWordsOk() ([]string, bool) {
	if o == nil || IsNil(o.KeyWords) {
		return nil, false
	}
	return o.KeyWords, true
}

// HasKeyWords returns a boolean if a field has been set.
func (o *AlipayOpenSearchSubservicekeywordApplyModel) HasKeyWords() bool {
	if o != nil && !IsNil(o.KeyWords) {
		return true
	}

	return false
}

// SetKeyWords gets a reference to the given []string and assigns it to the KeyWords field.
func (o *AlipayOpenSearchSubservicekeywordApplyModel) SetKeyWords(v []string) {
	o.KeyWords = v
}

// GetSubServiceCode returns the SubServiceCode field value if set, zero value otherwise.
func (o *AlipayOpenSearchSubservicekeywordApplyModel) GetSubServiceCode() string {
	if o == nil || IsNil(o.SubServiceCode) {
		var ret string
		return ret
	}
	return *o.SubServiceCode
}

// GetSubServiceCodeOk returns a tuple with the SubServiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchSubservicekeywordApplyModel) GetSubServiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SubServiceCode) {
		return nil, false
	}
	return o.SubServiceCode, true
}

// HasSubServiceCode returns a boolean if a field has been set.
func (o *AlipayOpenSearchSubservicekeywordApplyModel) HasSubServiceCode() bool {
	if o != nil && !IsNil(o.SubServiceCode) {
		return true
	}

	return false
}

// SetSubServiceCode gets a reference to the given string and assigns it to the SubServiceCode field.
func (o *AlipayOpenSearchSubservicekeywordApplyModel) SetSubServiceCode(v string) {
	o.SubServiceCode = &v
}

// GetTargetAppid returns the TargetAppid field value if set, zero value otherwise.
func (o *AlipayOpenSearchSubservicekeywordApplyModel) GetTargetAppid() string {
	if o == nil || IsNil(o.TargetAppid) {
		var ret string
		return ret
	}
	return *o.TargetAppid
}

// GetTargetAppidOk returns a tuple with the TargetAppid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchSubservicekeywordApplyModel) GetTargetAppidOk() (*string, bool) {
	if o == nil || IsNil(o.TargetAppid) {
		return nil, false
	}
	return o.TargetAppid, true
}

// HasTargetAppid returns a boolean if a field has been set.
func (o *AlipayOpenSearchSubservicekeywordApplyModel) HasTargetAppid() bool {
	if o != nil && !IsNil(o.TargetAppid) {
		return true
	}

	return false
}

// SetTargetAppid gets a reference to the given string and assigns it to the TargetAppid field.
func (o *AlipayOpenSearchSubservicekeywordApplyModel) SetTargetAppid(v string) {
	o.TargetAppid = &v
}

func (o AlipayOpenSearchSubservicekeywordApplyModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSearchSubservicekeywordApplyModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConfigId) {
		toSerialize["config_id"] = o.ConfigId
	}
	if !IsNil(o.KeyWords) {
		toSerialize["key_words"] = o.KeyWords
	}
	if !IsNil(o.SubServiceCode) {
		toSerialize["sub_service_code"] = o.SubServiceCode
	}
	if !IsNil(o.TargetAppid) {
		toSerialize["target_appid"] = o.TargetAppid
	}
	return toSerialize, nil
}

type NullableAlipayOpenSearchSubservicekeywordApplyModel struct {
	value *AlipayOpenSearchSubservicekeywordApplyModel
	isSet bool
}

func (v NullableAlipayOpenSearchSubservicekeywordApplyModel) Get() *AlipayOpenSearchSubservicekeywordApplyModel {
	return v.value
}

func (v *NullableAlipayOpenSearchSubservicekeywordApplyModel) Set(val *AlipayOpenSearchSubservicekeywordApplyModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchSubservicekeywordApplyModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchSubservicekeywordApplyModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchSubservicekeywordApplyModel(val *AlipayOpenSearchSubservicekeywordApplyModel) *NullableAlipayOpenSearchSubservicekeywordApplyModel {
	return &NullableAlipayOpenSearchSubservicekeywordApplyModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchSubservicekeywordApplyModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchSubservicekeywordApplyModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

