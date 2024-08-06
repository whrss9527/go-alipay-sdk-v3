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

// checks if the AlipayMerchantAnchorInfoQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMerchantAnchorInfoQueryResponseModel{}

// AlipayMerchantAnchorInfoQueryResponseModel struct for AlipayMerchantAnchorInfoQueryResponseModel
type AlipayMerchantAnchorInfoQueryResponseModel struct {
	// 主播ID+唯一+天猫好房带货+根据uid获取
	AnchorId *string `json:"anchor_id,omitempty"`
}

// NewAlipayMerchantAnchorInfoQueryResponseModel instantiates a new AlipayMerchantAnchorInfoQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMerchantAnchorInfoQueryResponseModel() *AlipayMerchantAnchorInfoQueryResponseModel {
	this := AlipayMerchantAnchorInfoQueryResponseModel{}
	return &this
}

// NewAlipayMerchantAnchorInfoQueryResponseModelWithDefaults instantiates a new AlipayMerchantAnchorInfoQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMerchantAnchorInfoQueryResponseModelWithDefaults() *AlipayMerchantAnchorInfoQueryResponseModel {
	this := AlipayMerchantAnchorInfoQueryResponseModel{}
	return &this
}

// GetAnchorId returns the AnchorId field value if set, zero value otherwise.
func (o *AlipayMerchantAnchorInfoQueryResponseModel) GetAnchorId() string {
	if o == nil || IsNil(o.AnchorId) {
		var ret string
		return ret
	}
	return *o.AnchorId
}

// GetAnchorIdOk returns a tuple with the AnchorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMerchantAnchorInfoQueryResponseModel) GetAnchorIdOk() (*string, bool) {
	if o == nil || IsNil(o.AnchorId) {
		return nil, false
	}
	return o.AnchorId, true
}

// HasAnchorId returns a boolean if a field has been set.
func (o *AlipayMerchantAnchorInfoQueryResponseModel) HasAnchorId() bool {
	if o != nil && !IsNil(o.AnchorId) {
		return true
	}

	return false
}

// SetAnchorId gets a reference to the given string and assigns it to the AnchorId field.
func (o *AlipayMerchantAnchorInfoQueryResponseModel) SetAnchorId(v string) {
	o.AnchorId = &v
}

func (o AlipayMerchantAnchorInfoQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMerchantAnchorInfoQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnchorId) {
		toSerialize["anchor_id"] = o.AnchorId
	}
	return toSerialize, nil
}

type NullableAlipayMerchantAnchorInfoQueryResponseModel struct {
	value *AlipayMerchantAnchorInfoQueryResponseModel
	isSet bool
}

func (v NullableAlipayMerchantAnchorInfoQueryResponseModel) Get() *AlipayMerchantAnchorInfoQueryResponseModel {
	return v.value
}

func (v *NullableAlipayMerchantAnchorInfoQueryResponseModel) Set(val *AlipayMerchantAnchorInfoQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMerchantAnchorInfoQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMerchantAnchorInfoQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMerchantAnchorInfoQueryResponseModel(val *AlipayMerchantAnchorInfoQueryResponseModel) *NullableAlipayMerchantAnchorInfoQueryResponseModel {
	return &NullableAlipayMerchantAnchorInfoQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMerchantAnchorInfoQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMerchantAnchorInfoQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


