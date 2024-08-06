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

// checks if the AlipayMobilePublicQrcodeCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMobilePublicQrcodeCreateResponseModel{}

// AlipayMobilePublicQrcodeCreateResponseModel struct for AlipayMobilePublicQrcodeCreateResponseModel
type AlipayMobilePublicQrcodeCreateResponseModel struct {
	// 返回结果码，如200，标识成功
	Code *int32 `json:"code,omitempty"`
	// 图片地址
	CodeImg *string `json:"code_img,omitempty"`
	// 码过期时间，单位：秒
	ExpireSecond *int32 `json:"expire_second,omitempty"`
	// 结果吗描述信息
	Msg *string `json:"msg,omitempty"`
}

// NewAlipayMobilePublicQrcodeCreateResponseModel instantiates a new AlipayMobilePublicQrcodeCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMobilePublicQrcodeCreateResponseModel() *AlipayMobilePublicQrcodeCreateResponseModel {
	this := AlipayMobilePublicQrcodeCreateResponseModel{}
	return &this
}

// NewAlipayMobilePublicQrcodeCreateResponseModelWithDefaults instantiates a new AlipayMobilePublicQrcodeCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMobilePublicQrcodeCreateResponseModelWithDefaults() *AlipayMobilePublicQrcodeCreateResponseModel {
	this := AlipayMobilePublicQrcodeCreateResponseModel{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AlipayMobilePublicQrcodeCreateResponseModel) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMobilePublicQrcodeCreateResponseModel) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AlipayMobilePublicQrcodeCreateResponseModel) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *AlipayMobilePublicQrcodeCreateResponseModel) SetCode(v int32) {
	o.Code = &v
}

// GetCodeImg returns the CodeImg field value if set, zero value otherwise.
func (o *AlipayMobilePublicQrcodeCreateResponseModel) GetCodeImg() string {
	if o == nil || IsNil(o.CodeImg) {
		var ret string
		return ret
	}
	return *o.CodeImg
}

// GetCodeImgOk returns a tuple with the CodeImg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMobilePublicQrcodeCreateResponseModel) GetCodeImgOk() (*string, bool) {
	if o == nil || IsNil(o.CodeImg) {
		return nil, false
	}
	return o.CodeImg, true
}

// HasCodeImg returns a boolean if a field has been set.
func (o *AlipayMobilePublicQrcodeCreateResponseModel) HasCodeImg() bool {
	if o != nil && !IsNil(o.CodeImg) {
		return true
	}

	return false
}

// SetCodeImg gets a reference to the given string and assigns it to the CodeImg field.
func (o *AlipayMobilePublicQrcodeCreateResponseModel) SetCodeImg(v string) {
	o.CodeImg = &v
}

// GetExpireSecond returns the ExpireSecond field value if set, zero value otherwise.
func (o *AlipayMobilePublicQrcodeCreateResponseModel) GetExpireSecond() int32 {
	if o == nil || IsNil(o.ExpireSecond) {
		var ret int32
		return ret
	}
	return *o.ExpireSecond
}

// GetExpireSecondOk returns a tuple with the ExpireSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMobilePublicQrcodeCreateResponseModel) GetExpireSecondOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpireSecond) {
		return nil, false
	}
	return o.ExpireSecond, true
}

// HasExpireSecond returns a boolean if a field has been set.
func (o *AlipayMobilePublicQrcodeCreateResponseModel) HasExpireSecond() bool {
	if o != nil && !IsNil(o.ExpireSecond) {
		return true
	}

	return false
}

// SetExpireSecond gets a reference to the given int32 and assigns it to the ExpireSecond field.
func (o *AlipayMobilePublicQrcodeCreateResponseModel) SetExpireSecond(v int32) {
	o.ExpireSecond = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *AlipayMobilePublicQrcodeCreateResponseModel) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMobilePublicQrcodeCreateResponseModel) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *AlipayMobilePublicQrcodeCreateResponseModel) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *AlipayMobilePublicQrcodeCreateResponseModel) SetMsg(v string) {
	o.Msg = &v
}

func (o AlipayMobilePublicQrcodeCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMobilePublicQrcodeCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.CodeImg) {
		toSerialize["code_img"] = o.CodeImg
	}
	if !IsNil(o.ExpireSecond) {
		toSerialize["expire_second"] = o.ExpireSecond
	}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	return toSerialize, nil
}

type NullableAlipayMobilePublicQrcodeCreateResponseModel struct {
	value *AlipayMobilePublicQrcodeCreateResponseModel
	isSet bool
}

func (v NullableAlipayMobilePublicQrcodeCreateResponseModel) Get() *AlipayMobilePublicQrcodeCreateResponseModel {
	return v.value
}

func (v *NullableAlipayMobilePublicQrcodeCreateResponseModel) Set(val *AlipayMobilePublicQrcodeCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMobilePublicQrcodeCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMobilePublicQrcodeCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMobilePublicQrcodeCreateResponseModel(val *AlipayMobilePublicQrcodeCreateResponseModel) *NullableAlipayMobilePublicQrcodeCreateResponseModel {
	return &NullableAlipayMobilePublicQrcodeCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMobilePublicQrcodeCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMobilePublicQrcodeCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


