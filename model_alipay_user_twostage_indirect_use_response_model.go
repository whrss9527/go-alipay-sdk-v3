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

// checks if the AlipayUserTwostageIndirectUseResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayUserTwostageIndirectUseResponseModel{}

// AlipayUserTwostageIndirectUseResponseModel struct for AlipayUserTwostageIndirectUseResponseModel
type AlipayUserTwostageIndirectUseResponseModel struct {
	// 蚂蚁统一会员ID
	UserId *string `json:"user_id,omitempty"`
}

// NewAlipayUserTwostageIndirectUseResponseModel instantiates a new AlipayUserTwostageIndirectUseResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayUserTwostageIndirectUseResponseModel() *AlipayUserTwostageIndirectUseResponseModel {
	this := AlipayUserTwostageIndirectUseResponseModel{}
	return &this
}

// NewAlipayUserTwostageIndirectUseResponseModelWithDefaults instantiates a new AlipayUserTwostageIndirectUseResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayUserTwostageIndirectUseResponseModelWithDefaults() *AlipayUserTwostageIndirectUseResponseModel {
	this := AlipayUserTwostageIndirectUseResponseModel{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AlipayUserTwostageIndirectUseResponseModel) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserTwostageIndirectUseResponseModel) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AlipayUserTwostageIndirectUseResponseModel) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AlipayUserTwostageIndirectUseResponseModel) SetUserId(v string) {
	o.UserId = &v
}

func (o AlipayUserTwostageIndirectUseResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayUserTwostageIndirectUseResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableAlipayUserTwostageIndirectUseResponseModel struct {
	value *AlipayUserTwostageIndirectUseResponseModel
	isSet bool
}

func (v NullableAlipayUserTwostageIndirectUseResponseModel) Get() *AlipayUserTwostageIndirectUseResponseModel {
	return v.value
}

func (v *NullableAlipayUserTwostageIndirectUseResponseModel) Set(val *AlipayUserTwostageIndirectUseResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayUserTwostageIndirectUseResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayUserTwostageIndirectUseResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayUserTwostageIndirectUseResponseModel(val *AlipayUserTwostageIndirectUseResponseModel) *NullableAlipayUserTwostageIndirectUseResponseModel {
	return &NullableAlipayUserTwostageIndirectUseResponseModel{value: val, isSet: true}
}

func (v NullableAlipayUserTwostageIndirectUseResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayUserTwostageIndirectUseResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
