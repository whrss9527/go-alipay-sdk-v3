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

// checks if the AlipayOpenSearchAppkeywordQuerystatusResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSearchAppkeywordQuerystatusResponseModel{}

// AlipayOpenSearchAppkeywordQuerystatusResponseModel struct for AlipayOpenSearchAppkeywordQuerystatusResponseModel
type AlipayOpenSearchAppkeywordQuerystatusResponseModel struct {
	// 关键词工单审核状态返回值
	KeyWords []KeyWordInfo `json:"key_words,omitempty"`
}

// NewAlipayOpenSearchAppkeywordQuerystatusResponseModel instantiates a new AlipayOpenSearchAppkeywordQuerystatusResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSearchAppkeywordQuerystatusResponseModel() *AlipayOpenSearchAppkeywordQuerystatusResponseModel {
	this := AlipayOpenSearchAppkeywordQuerystatusResponseModel{}
	return &this
}

// NewAlipayOpenSearchAppkeywordQuerystatusResponseModelWithDefaults instantiates a new AlipayOpenSearchAppkeywordQuerystatusResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSearchAppkeywordQuerystatusResponseModelWithDefaults() *AlipayOpenSearchAppkeywordQuerystatusResponseModel {
	this := AlipayOpenSearchAppkeywordQuerystatusResponseModel{}
	return &this
}

// GetKeyWords returns the KeyWords field value if set, zero value otherwise.
func (o *AlipayOpenSearchAppkeywordQuerystatusResponseModel) GetKeyWords() []KeyWordInfo {
	if o == nil || IsNil(o.KeyWords) {
		var ret []KeyWordInfo
		return ret
	}
	return o.KeyWords
}

// GetKeyWordsOk returns a tuple with the KeyWords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchAppkeywordQuerystatusResponseModel) GetKeyWordsOk() ([]KeyWordInfo, bool) {
	if o == nil || IsNil(o.KeyWords) {
		return nil, false
	}
	return o.KeyWords, true
}

// HasKeyWords returns a boolean if a field has been set.
func (o *AlipayOpenSearchAppkeywordQuerystatusResponseModel) HasKeyWords() bool {
	if o != nil && !IsNil(o.KeyWords) {
		return true
	}

	return false
}

// SetKeyWords gets a reference to the given []KeyWordInfo and assigns it to the KeyWords field.
func (o *AlipayOpenSearchAppkeywordQuerystatusResponseModel) SetKeyWords(v []KeyWordInfo) {
	o.KeyWords = v
}

func (o AlipayOpenSearchAppkeywordQuerystatusResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSearchAppkeywordQuerystatusResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KeyWords) {
		toSerialize["key_words"] = o.KeyWords
	}
	return toSerialize, nil
}

type NullableAlipayOpenSearchAppkeywordQuerystatusResponseModel struct {
	value *AlipayOpenSearchAppkeywordQuerystatusResponseModel
	isSet bool
}

func (v NullableAlipayOpenSearchAppkeywordQuerystatusResponseModel) Get() *AlipayOpenSearchAppkeywordQuerystatusResponseModel {
	return v.value
}

func (v *NullableAlipayOpenSearchAppkeywordQuerystatusResponseModel) Set(val *AlipayOpenSearchAppkeywordQuerystatusResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchAppkeywordQuerystatusResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchAppkeywordQuerystatusResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchAppkeywordQuerystatusResponseModel(val *AlipayOpenSearchAppkeywordQuerystatusResponseModel) *NullableAlipayOpenSearchAppkeywordQuerystatusResponseModel {
	return &NullableAlipayOpenSearchAppkeywordQuerystatusResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchAppkeywordQuerystatusResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchAppkeywordQuerystatusResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
