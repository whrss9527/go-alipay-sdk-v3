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

// checks if the AlipayMobilePublicShortlinkCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMobilePublicShortlinkCreateResponseModel{}

// AlipayMobilePublicShortlinkCreateResponseModel struct for AlipayMobilePublicShortlinkCreateResponseModel
type AlipayMobilePublicShortlinkCreateResponseModel struct {
	// 结果码
	Code *string `json:"code,omitempty"`
	// 成功
	Msg *string `json:"msg,omitempty"`
	// 短链接url
	Shortlink *string `json:"shortlink,omitempty"`
}

// NewAlipayMobilePublicShortlinkCreateResponseModel instantiates a new AlipayMobilePublicShortlinkCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMobilePublicShortlinkCreateResponseModel() *AlipayMobilePublicShortlinkCreateResponseModel {
	this := AlipayMobilePublicShortlinkCreateResponseModel{}
	return &this
}

// NewAlipayMobilePublicShortlinkCreateResponseModelWithDefaults instantiates a new AlipayMobilePublicShortlinkCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMobilePublicShortlinkCreateResponseModelWithDefaults() *AlipayMobilePublicShortlinkCreateResponseModel {
	this := AlipayMobilePublicShortlinkCreateResponseModel{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AlipayMobilePublicShortlinkCreateResponseModel) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMobilePublicShortlinkCreateResponseModel) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AlipayMobilePublicShortlinkCreateResponseModel) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AlipayMobilePublicShortlinkCreateResponseModel) SetCode(v string) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *AlipayMobilePublicShortlinkCreateResponseModel) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMobilePublicShortlinkCreateResponseModel) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *AlipayMobilePublicShortlinkCreateResponseModel) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *AlipayMobilePublicShortlinkCreateResponseModel) SetMsg(v string) {
	o.Msg = &v
}

// GetShortlink returns the Shortlink field value if set, zero value otherwise.
func (o *AlipayMobilePublicShortlinkCreateResponseModel) GetShortlink() string {
	if o == nil || IsNil(o.Shortlink) {
		var ret string
		return ret
	}
	return *o.Shortlink
}

// GetShortlinkOk returns a tuple with the Shortlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMobilePublicShortlinkCreateResponseModel) GetShortlinkOk() (*string, bool) {
	if o == nil || IsNil(o.Shortlink) {
		return nil, false
	}
	return o.Shortlink, true
}

// HasShortlink returns a boolean if a field has been set.
func (o *AlipayMobilePublicShortlinkCreateResponseModel) HasShortlink() bool {
	if o != nil && !IsNil(o.Shortlink) {
		return true
	}

	return false
}

// SetShortlink gets a reference to the given string and assigns it to the Shortlink field.
func (o *AlipayMobilePublicShortlinkCreateResponseModel) SetShortlink(v string) {
	o.Shortlink = &v
}

func (o AlipayMobilePublicShortlinkCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMobilePublicShortlinkCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !IsNil(o.Shortlink) {
		toSerialize["shortlink"] = o.Shortlink
	}
	return toSerialize, nil
}

type NullableAlipayMobilePublicShortlinkCreateResponseModel struct {
	value *AlipayMobilePublicShortlinkCreateResponseModel
	isSet bool
}

func (v NullableAlipayMobilePublicShortlinkCreateResponseModel) Get() *AlipayMobilePublicShortlinkCreateResponseModel {
	return v.value
}

func (v *NullableAlipayMobilePublicShortlinkCreateResponseModel) Set(val *AlipayMobilePublicShortlinkCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMobilePublicShortlinkCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMobilePublicShortlinkCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMobilePublicShortlinkCreateResponseModel(val *AlipayMobilePublicShortlinkCreateResponseModel) *NullableAlipayMobilePublicShortlinkCreateResponseModel {
	return &NullableAlipayMobilePublicShortlinkCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMobilePublicShortlinkCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMobilePublicShortlinkCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


