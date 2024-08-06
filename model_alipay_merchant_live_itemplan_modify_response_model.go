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

// checks if the AlipayMerchantLiveItemplanModifyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMerchantLiveItemplanModifyResponseModel{}

// AlipayMerchantLiveItemplanModifyResponseModel struct for AlipayMerchantLiveItemplanModifyResponseModel
type AlipayMerchantLiveItemplanModifyResponseModel struct {
	// 针对操作结果的说明
	MsgInfo *string `json:"msg_info,omitempty"`
}

// NewAlipayMerchantLiveItemplanModifyResponseModel instantiates a new AlipayMerchantLiveItemplanModifyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMerchantLiveItemplanModifyResponseModel() *AlipayMerchantLiveItemplanModifyResponseModel {
	this := AlipayMerchantLiveItemplanModifyResponseModel{}
	return &this
}

// NewAlipayMerchantLiveItemplanModifyResponseModelWithDefaults instantiates a new AlipayMerchantLiveItemplanModifyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMerchantLiveItemplanModifyResponseModelWithDefaults() *AlipayMerchantLiveItemplanModifyResponseModel {
	this := AlipayMerchantLiveItemplanModifyResponseModel{}
	return &this
}

// GetMsgInfo returns the MsgInfo field value if set, zero value otherwise.
func (o *AlipayMerchantLiveItemplanModifyResponseModel) GetMsgInfo() string {
	if o == nil || IsNil(o.MsgInfo) {
		var ret string
		return ret
	}
	return *o.MsgInfo
}

// GetMsgInfoOk returns a tuple with the MsgInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantLiveItemplanModifyResponseModel) GetMsgInfoOk() (*string, bool) {
	if o == nil || IsNil(o.MsgInfo) {
		return nil, false
	}
	return o.MsgInfo, true
}

// HasMsgInfo returns a boolean if a field has been set.
func (o *AlipayMerchantLiveItemplanModifyResponseModel) HasMsgInfo() bool {
	if o != nil && !IsNil(o.MsgInfo) {
		return true
	}

	return false
}

// SetMsgInfo gets a reference to the given string and assigns it to the MsgInfo field.
func (o *AlipayMerchantLiveItemplanModifyResponseModel) SetMsgInfo(v string) {
	o.MsgInfo = &v
}

func (o AlipayMerchantLiveItemplanModifyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMerchantLiveItemplanModifyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MsgInfo) {
		toSerialize["msg_info"] = o.MsgInfo
	}
	return toSerialize, nil
}

type NullableAlipayMerchantLiveItemplanModifyResponseModel struct {
	value *AlipayMerchantLiveItemplanModifyResponseModel
	isSet bool
}

func (v NullableAlipayMerchantLiveItemplanModifyResponseModel) Get() *AlipayMerchantLiveItemplanModifyResponseModel {
	return v.value
}

func (v *NullableAlipayMerchantLiveItemplanModifyResponseModel) Set(val *AlipayMerchantLiveItemplanModifyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMerchantLiveItemplanModifyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMerchantLiveItemplanModifyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMerchantLiveItemplanModifyResponseModel(val *AlipayMerchantLiveItemplanModifyResponseModel) *NullableAlipayMerchantLiveItemplanModifyResponseModel {
	return &NullableAlipayMerchantLiveItemplanModifyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMerchantLiveItemplanModifyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMerchantLiveItemplanModifyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
