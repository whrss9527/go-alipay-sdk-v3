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

// checks if the AlipayOpenSearchSubservicekeywordApplyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSearchSubservicekeywordApplyResponseModel{}

// AlipayOpenSearchSubservicekeywordApplyResponseModel struct for AlipayOpenSearchSubservicekeywordApplyResponseModel
type AlipayOpenSearchSubservicekeywordApplyResponseModel struct {
	// 审核工单id，唯一，提报服务关键词，提报服务关键词返回该id
	ApplyNo *string `json:"apply_no,omitempty"`
}

// NewAlipayOpenSearchSubservicekeywordApplyResponseModel instantiates a new AlipayOpenSearchSubservicekeywordApplyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSearchSubservicekeywordApplyResponseModel() *AlipayOpenSearchSubservicekeywordApplyResponseModel {
	this := AlipayOpenSearchSubservicekeywordApplyResponseModel{}
	return &this
}

// NewAlipayOpenSearchSubservicekeywordApplyResponseModelWithDefaults instantiates a new AlipayOpenSearchSubservicekeywordApplyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSearchSubservicekeywordApplyResponseModelWithDefaults() *AlipayOpenSearchSubservicekeywordApplyResponseModel {
	this := AlipayOpenSearchSubservicekeywordApplyResponseModel{}
	return &this
}

// GetApplyNo returns the ApplyNo field value if set, zero value otherwise.
func (o *AlipayOpenSearchSubservicekeywordApplyResponseModel) GetApplyNo() string {
	if o == nil || IsNil(o.ApplyNo) {
		var ret string
		return ret
	}
	return *o.ApplyNo
}

// GetApplyNoOk returns a tuple with the ApplyNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchSubservicekeywordApplyResponseModel) GetApplyNoOk() (*string, bool) {
	if o == nil || IsNil(o.ApplyNo) {
		return nil, false
	}
	return o.ApplyNo, true
}

// HasApplyNo returns a boolean if a field has been set.
func (o *AlipayOpenSearchSubservicekeywordApplyResponseModel) HasApplyNo() bool {
	if o != nil && !IsNil(o.ApplyNo) {
		return true
	}

	return false
}

// SetApplyNo gets a reference to the given string and assigns it to the ApplyNo field.
func (o *AlipayOpenSearchSubservicekeywordApplyResponseModel) SetApplyNo(v string) {
	o.ApplyNo = &v
}

func (o AlipayOpenSearchSubservicekeywordApplyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSearchSubservicekeywordApplyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplyNo) {
		toSerialize["apply_no"] = o.ApplyNo
	}
	return toSerialize, nil
}

type NullableAlipayOpenSearchSubservicekeywordApplyResponseModel struct {
	value *AlipayOpenSearchSubservicekeywordApplyResponseModel
	isSet bool
}

func (v NullableAlipayOpenSearchSubservicekeywordApplyResponseModel) Get() *AlipayOpenSearchSubservicekeywordApplyResponseModel {
	return v.value
}

func (v *NullableAlipayOpenSearchSubservicekeywordApplyResponseModel) Set(val *AlipayOpenSearchSubservicekeywordApplyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchSubservicekeywordApplyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchSubservicekeywordApplyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchSubservicekeywordApplyResponseModel(val *AlipayOpenSearchSubservicekeywordApplyResponseModel) *NullableAlipayOpenSearchSubservicekeywordApplyResponseModel {
	return &NullableAlipayOpenSearchSubservicekeywordApplyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchSubservicekeywordApplyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchSubservicekeywordApplyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
