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

// checks if the AlipayMobilePublicMessageCustomSendResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMobilePublicMessageCustomSendResponseModel{}

// AlipayMobilePublicMessageCustomSendResponseModel struct for AlipayMobilePublicMessageCustomSendResponseModel
type AlipayMobilePublicMessageCustomSendResponseModel struct {
	// 结果码
	Code *string `json:"code,omitempty"`
	// 结果描述
	Msg *string `json:"msg,omitempty"`
}

// NewAlipayMobilePublicMessageCustomSendResponseModel instantiates a new AlipayMobilePublicMessageCustomSendResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMobilePublicMessageCustomSendResponseModel() *AlipayMobilePublicMessageCustomSendResponseModel {
	this := AlipayMobilePublicMessageCustomSendResponseModel{}
	return &this
}

// NewAlipayMobilePublicMessageCustomSendResponseModelWithDefaults instantiates a new AlipayMobilePublicMessageCustomSendResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMobilePublicMessageCustomSendResponseModelWithDefaults() *AlipayMobilePublicMessageCustomSendResponseModel {
	this := AlipayMobilePublicMessageCustomSendResponseModel{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AlipayMobilePublicMessageCustomSendResponseModel) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMobilePublicMessageCustomSendResponseModel) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AlipayMobilePublicMessageCustomSendResponseModel) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AlipayMobilePublicMessageCustomSendResponseModel) SetCode(v string) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *AlipayMobilePublicMessageCustomSendResponseModel) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMobilePublicMessageCustomSendResponseModel) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *AlipayMobilePublicMessageCustomSendResponseModel) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *AlipayMobilePublicMessageCustomSendResponseModel) SetMsg(v string) {
	o.Msg = &v
}

func (o AlipayMobilePublicMessageCustomSendResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMobilePublicMessageCustomSendResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	return toSerialize, nil
}

type NullableAlipayMobilePublicMessageCustomSendResponseModel struct {
	value *AlipayMobilePublicMessageCustomSendResponseModel
	isSet bool
}

func (v NullableAlipayMobilePublicMessageCustomSendResponseModel) Get() *AlipayMobilePublicMessageCustomSendResponseModel {
	return v.value
}

func (v *NullableAlipayMobilePublicMessageCustomSendResponseModel) Set(val *AlipayMobilePublicMessageCustomSendResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMobilePublicMessageCustomSendResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMobilePublicMessageCustomSendResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMobilePublicMessageCustomSendResponseModel(val *AlipayMobilePublicMessageCustomSendResponseModel) *NullableAlipayMobilePublicMessageCustomSendResponseModel {
	return &NullableAlipayMobilePublicMessageCustomSendResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMobilePublicMessageCustomSendResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMobilePublicMessageCustomSendResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
