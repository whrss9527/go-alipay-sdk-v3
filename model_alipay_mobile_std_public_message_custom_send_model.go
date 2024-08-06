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

// checks if the AlipayMobileStdPublicMessageCustomSendModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMobileStdPublicMessageCustomSendModel{}

// AlipayMobileStdPublicMessageCustomSendModel struct for AlipayMobileStdPublicMessageCustomSendModel
type AlipayMobileStdPublicMessageCustomSendModel struct {
	// 业务内容，其中包括消息类型msgType、消息体和消息接受人toUserId三大块，具体参见“表1-2 服务窗单发客服消息的biz_content参数说明”。
	BizContent *string `json:"biz_content,omitempty"`
}

// NewAlipayMobileStdPublicMessageCustomSendModel instantiates a new AlipayMobileStdPublicMessageCustomSendModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMobileStdPublicMessageCustomSendModel() *AlipayMobileStdPublicMessageCustomSendModel {
	this := AlipayMobileStdPublicMessageCustomSendModel{}
	return &this
}

// NewAlipayMobileStdPublicMessageCustomSendModelWithDefaults instantiates a new AlipayMobileStdPublicMessageCustomSendModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMobileStdPublicMessageCustomSendModelWithDefaults() *AlipayMobileStdPublicMessageCustomSendModel {
	this := AlipayMobileStdPublicMessageCustomSendModel{}
	return &this
}

// GetBizContent returns the BizContent field value if set, zero value otherwise.
func (o *AlipayMobileStdPublicMessageCustomSendModel) GetBizContent() string {
	if o == nil || IsNil(o.BizContent) {
		var ret string
		return ret
	}
	return *o.BizContent
}

// GetBizContentOk returns a tuple with the BizContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMobileStdPublicMessageCustomSendModel) GetBizContentOk() (*string, bool) {
	if o == nil || IsNil(o.BizContent) {
		return nil, false
	}
	return o.BizContent, true
}

// HasBizContent returns a boolean if a field has been set.
func (o *AlipayMobileStdPublicMessageCustomSendModel) HasBizContent() bool {
	if o != nil && !IsNil(o.BizContent) {
		return true
	}

	return false
}

// SetBizContent gets a reference to the given string and assigns it to the BizContent field.
func (o *AlipayMobileStdPublicMessageCustomSendModel) SetBizContent(v string) {
	o.BizContent = &v
}

func (o AlipayMobileStdPublicMessageCustomSendModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMobileStdPublicMessageCustomSendModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BizContent) {
		toSerialize["biz_content"] = o.BizContent
	}
	return toSerialize, nil
}

type NullableAlipayMobileStdPublicMessageCustomSendModel struct {
	value *AlipayMobileStdPublicMessageCustomSendModel
	isSet bool
}

func (v NullableAlipayMobileStdPublicMessageCustomSendModel) Get() *AlipayMobileStdPublicMessageCustomSendModel {
	return v.value
}

func (v *NullableAlipayMobileStdPublicMessageCustomSendModel) Set(val *AlipayMobileStdPublicMessageCustomSendModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMobileStdPublicMessageCustomSendModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMobileStdPublicMessageCustomSendModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMobileStdPublicMessageCustomSendModel(val *AlipayMobileStdPublicMessageCustomSendModel) *NullableAlipayMobileStdPublicMessageCustomSendModel {
	return &NullableAlipayMobileStdPublicMessageCustomSendModel{value: val, isSet: true}
}

func (v NullableAlipayMobileStdPublicMessageCustomSendModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMobileStdPublicMessageCustomSendModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


