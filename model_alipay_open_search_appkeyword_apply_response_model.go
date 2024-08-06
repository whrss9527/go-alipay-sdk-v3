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

// checks if the AlipayOpenSearchAppkeywordApplyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSearchAppkeywordApplyResponseModel{}

// AlipayOpenSearchAppkeywordApplyResponseModel struct for AlipayOpenSearchAppkeywordApplyResponseModel
type AlipayOpenSearchAppkeywordApplyResponseModel struct {
	// 提报之后的后台审核工单的id
	ApplyNo *string `json:"apply_no,omitempty"`
}

// NewAlipayOpenSearchAppkeywordApplyResponseModel instantiates a new AlipayOpenSearchAppkeywordApplyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSearchAppkeywordApplyResponseModel() *AlipayOpenSearchAppkeywordApplyResponseModel {
	this := AlipayOpenSearchAppkeywordApplyResponseModel{}
	return &this
}

// NewAlipayOpenSearchAppkeywordApplyResponseModelWithDefaults instantiates a new AlipayOpenSearchAppkeywordApplyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSearchAppkeywordApplyResponseModelWithDefaults() *AlipayOpenSearchAppkeywordApplyResponseModel {
	this := AlipayOpenSearchAppkeywordApplyResponseModel{}
	return &this
}

// GetApplyNo returns the ApplyNo field value if set, zero value otherwise.
func (o *AlipayOpenSearchAppkeywordApplyResponseModel) GetApplyNo() string {
	if o == nil || IsNil(o.ApplyNo) {
		var ret string
		return ret
	}
	return *o.ApplyNo
}

// GetApplyNoOk returns a tuple with the ApplyNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchAppkeywordApplyResponseModel) GetApplyNoOk() (*string, bool) {
	if o == nil || IsNil(o.ApplyNo) {
		return nil, false
	}
	return o.ApplyNo, true
}

// HasApplyNo returns a boolean if a field has been set.
func (o *AlipayOpenSearchAppkeywordApplyResponseModel) HasApplyNo() bool {
	if o != nil && !IsNil(o.ApplyNo) {
		return true
	}

	return false
}

// SetApplyNo gets a reference to the given string and assigns it to the ApplyNo field.
func (o *AlipayOpenSearchAppkeywordApplyResponseModel) SetApplyNo(v string) {
	o.ApplyNo = &v
}

func (o AlipayOpenSearchAppkeywordApplyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSearchAppkeywordApplyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplyNo) {
		toSerialize["apply_no"] = o.ApplyNo
	}
	return toSerialize, nil
}

type NullableAlipayOpenSearchAppkeywordApplyResponseModel struct {
	value *AlipayOpenSearchAppkeywordApplyResponseModel
	isSet bool
}

func (v NullableAlipayOpenSearchAppkeywordApplyResponseModel) Get() *AlipayOpenSearchAppkeywordApplyResponseModel {
	return v.value
}

func (v *NullableAlipayOpenSearchAppkeywordApplyResponseModel) Set(val *AlipayOpenSearchAppkeywordApplyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchAppkeywordApplyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchAppkeywordApplyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchAppkeywordApplyResponseModel(val *AlipayOpenSearchAppkeywordApplyResponseModel) *NullableAlipayOpenSearchAppkeywordApplyResponseModel {
	return &NullableAlipayOpenSearchAppkeywordApplyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchAppkeywordApplyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchAppkeywordApplyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


