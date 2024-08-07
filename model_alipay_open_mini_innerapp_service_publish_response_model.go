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

// checks if the AlipayOpenMiniInnerappServicePublishResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniInnerappServicePublishResponseModel{}

// AlipayOpenMiniInnerappServicePublishResponseModel struct for AlipayOpenMiniInnerappServicePublishResponseModel
type AlipayOpenMiniInnerappServicePublishResponseModel struct {
	// 服务ID
	MerchandiseId *string `json:"merchandise_id,omitempty"`
}

// NewAlipayOpenMiniInnerappServicePublishResponseModel instantiates a new AlipayOpenMiniInnerappServicePublishResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniInnerappServicePublishResponseModel() *AlipayOpenMiniInnerappServicePublishResponseModel {
	this := AlipayOpenMiniInnerappServicePublishResponseModel{}
	return &this
}

// NewAlipayOpenMiniInnerappServicePublishResponseModelWithDefaults instantiates a new AlipayOpenMiniInnerappServicePublishResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniInnerappServicePublishResponseModelWithDefaults() *AlipayOpenMiniInnerappServicePublishResponseModel {
	this := AlipayOpenMiniInnerappServicePublishResponseModel{}
	return &this
}

// GetMerchandiseId returns the MerchandiseId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappServicePublishResponseModel) GetMerchandiseId() string {
	if o == nil || IsNil(o.MerchandiseId) {
		var ret string
		return ret
	}
	return *o.MerchandiseId
}

// GetMerchandiseIdOk returns a tuple with the MerchandiseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappServicePublishResponseModel) GetMerchandiseIdOk() (*string, bool) {
	if o == nil || IsNil(o.MerchandiseId) {
		return nil, false
	}
	return o.MerchandiseId, true
}

// HasMerchandiseId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappServicePublishResponseModel) HasMerchandiseId() bool {
	if o != nil && !IsNil(o.MerchandiseId) {
		return true
	}

	return false
}

// SetMerchandiseId gets a reference to the given string and assigns it to the MerchandiseId field.
func (o *AlipayOpenMiniInnerappServicePublishResponseModel) SetMerchandiseId(v string) {
	o.MerchandiseId = &v
}

func (o AlipayOpenMiniInnerappServicePublishResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniInnerappServicePublishResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchandiseId) {
		toSerialize["merchandise_id"] = o.MerchandiseId
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniInnerappServicePublishResponseModel struct {
	value *AlipayOpenMiniInnerappServicePublishResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniInnerappServicePublishResponseModel) Get() *AlipayOpenMiniInnerappServicePublishResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerappServicePublishResponseModel) Set(val *AlipayOpenMiniInnerappServicePublishResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerappServicePublishResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerappServicePublishResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerappServicePublishResponseModel(val *AlipayOpenMiniInnerappServicePublishResponseModel) *NullableAlipayOpenMiniInnerappServicePublishResponseModel {
	return &NullableAlipayOpenMiniInnerappServicePublishResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerappServicePublishResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerappServicePublishResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
