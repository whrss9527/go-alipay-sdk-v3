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

// checks if the AlipayOpenSearchBoxApplyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenSearchBoxApplyResponseModel{}

// AlipayOpenSearchBoxApplyResponseModel struct for AlipayOpenSearchBoxApplyResponseModel
type AlipayOpenSearchBoxApplyResponseModel struct {
	// 搜索直达配置id
	BoxId *string `json:"box_id,omitempty"`
}

// NewAlipayOpenSearchBoxApplyResponseModel instantiates a new AlipayOpenSearchBoxApplyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenSearchBoxApplyResponseModel() *AlipayOpenSearchBoxApplyResponseModel {
	this := AlipayOpenSearchBoxApplyResponseModel{}
	return &this
}

// NewAlipayOpenSearchBoxApplyResponseModelWithDefaults instantiates a new AlipayOpenSearchBoxApplyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenSearchBoxApplyResponseModelWithDefaults() *AlipayOpenSearchBoxApplyResponseModel {
	this := AlipayOpenSearchBoxApplyResponseModel{}
	return &this
}

// GetBoxId returns the BoxId field value if set, zero value otherwise.
func (o *AlipayOpenSearchBoxApplyResponseModel) GetBoxId() string {
	if o == nil || IsNil(o.BoxId) {
		var ret string
		return ret
	}
	return *o.BoxId
}

// GetBoxIdOk returns a tuple with the BoxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenSearchBoxApplyResponseModel) GetBoxIdOk() (*string, bool) {
	if o == nil || IsNil(o.BoxId) {
		return nil, false
	}
	return o.BoxId, true
}

// HasBoxId returns a boolean if a field has been set.
func (o *AlipayOpenSearchBoxApplyResponseModel) HasBoxId() bool {
	if o != nil && !IsNil(o.BoxId) {
		return true
	}

	return false
}

// SetBoxId gets a reference to the given string and assigns it to the BoxId field.
func (o *AlipayOpenSearchBoxApplyResponseModel) SetBoxId(v string) {
	o.BoxId = &v
}

func (o AlipayOpenSearchBoxApplyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenSearchBoxApplyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BoxId) {
		toSerialize["box_id"] = o.BoxId
	}
	return toSerialize, nil
}

type NullableAlipayOpenSearchBoxApplyResponseModel struct {
	value *AlipayOpenSearchBoxApplyResponseModel
	isSet bool
}

func (v NullableAlipayOpenSearchBoxApplyResponseModel) Get() *AlipayOpenSearchBoxApplyResponseModel {
	return v.value
}

func (v *NullableAlipayOpenSearchBoxApplyResponseModel) Set(val *AlipayOpenSearchBoxApplyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchBoxApplyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchBoxApplyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchBoxApplyResponseModel(val *AlipayOpenSearchBoxApplyResponseModel) *NullableAlipayOpenSearchBoxApplyResponseModel {
	return &NullableAlipayOpenSearchBoxApplyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchBoxApplyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchBoxApplyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
