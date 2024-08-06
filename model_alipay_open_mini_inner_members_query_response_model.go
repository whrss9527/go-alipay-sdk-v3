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

// checks if the AlipayOpenMiniInnerMembersQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniInnerMembersQueryResponseModel{}

// AlipayOpenMiniInnerMembersQueryResponseModel struct for AlipayOpenMiniInnerMembersQueryResponseModel
type AlipayOpenMiniInnerMembersQueryResponseModel struct {
	// 查询结果
	OpenAppOperatorVo []OpenAppOperatorVo `json:"open_app_operator_vo,omitempty"`
}

// NewAlipayOpenMiniInnerMembersQueryResponseModel instantiates a new AlipayOpenMiniInnerMembersQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniInnerMembersQueryResponseModel() *AlipayOpenMiniInnerMembersQueryResponseModel {
	this := AlipayOpenMiniInnerMembersQueryResponseModel{}
	return &this
}

// NewAlipayOpenMiniInnerMembersQueryResponseModelWithDefaults instantiates a new AlipayOpenMiniInnerMembersQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniInnerMembersQueryResponseModelWithDefaults() *AlipayOpenMiniInnerMembersQueryResponseModel {
	this := AlipayOpenMiniInnerMembersQueryResponseModel{}
	return &this
}

// GetOpenAppOperatorVo returns the OpenAppOperatorVo field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerMembersQueryResponseModel) GetOpenAppOperatorVo() []OpenAppOperatorVo {
	if o == nil || IsNil(o.OpenAppOperatorVo) {
		var ret []OpenAppOperatorVo
		return ret
	}
	return o.OpenAppOperatorVo
}

// GetOpenAppOperatorVoOk returns a tuple with the OpenAppOperatorVo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerMembersQueryResponseModel) GetOpenAppOperatorVoOk() ([]OpenAppOperatorVo, bool) {
	if o == nil || IsNil(o.OpenAppOperatorVo) {
		return nil, false
	}
	return o.OpenAppOperatorVo, true
}

// HasOpenAppOperatorVo returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerMembersQueryResponseModel) HasOpenAppOperatorVo() bool {
	if o != nil && !IsNil(o.OpenAppOperatorVo) {
		return true
	}

	return false
}

// SetOpenAppOperatorVo gets a reference to the given []OpenAppOperatorVo and assigns it to the OpenAppOperatorVo field.
func (o *AlipayOpenMiniInnerMembersQueryResponseModel) SetOpenAppOperatorVo(v []OpenAppOperatorVo) {
	o.OpenAppOperatorVo = v
}

func (o AlipayOpenMiniInnerMembersQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniInnerMembersQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OpenAppOperatorVo) {
		toSerialize["open_app_operator_vo"] = o.OpenAppOperatorVo
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniInnerMembersQueryResponseModel struct {
	value *AlipayOpenMiniInnerMembersQueryResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniInnerMembersQueryResponseModel) Get() *AlipayOpenMiniInnerMembersQueryResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerMembersQueryResponseModel) Set(val *AlipayOpenMiniInnerMembersQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerMembersQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerMembersQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerMembersQueryResponseModel(val *AlipayOpenMiniInnerMembersQueryResponseModel) *NullableAlipayOpenMiniInnerMembersQueryResponseModel {
	return &NullableAlipayOpenMiniInnerMembersQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerMembersQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerMembersQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
