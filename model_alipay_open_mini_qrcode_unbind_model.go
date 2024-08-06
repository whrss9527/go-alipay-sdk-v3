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

// checks if the AlipayOpenMiniQrcodeUnbindModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniQrcodeUnbindModel{}

// AlipayOpenMiniQrcodeUnbindModel struct for AlipayOpenMiniQrcodeUnbindModel
type AlipayOpenMiniQrcodeUnbindModel struct {
	// 路由规则组，用于唯一标记一条路由规则。调用 https://opendocs.alipay.com/apis/00rkye 接口关联普通二维码后的返回值。
	RouteGroup *string `json:"route_group,omitempty"`
}

// NewAlipayOpenMiniQrcodeUnbindModel instantiates a new AlipayOpenMiniQrcodeUnbindModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniQrcodeUnbindModel() *AlipayOpenMiniQrcodeUnbindModel {
	this := AlipayOpenMiniQrcodeUnbindModel{}
	return &this
}

// NewAlipayOpenMiniQrcodeUnbindModelWithDefaults instantiates a new AlipayOpenMiniQrcodeUnbindModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniQrcodeUnbindModelWithDefaults() *AlipayOpenMiniQrcodeUnbindModel {
	this := AlipayOpenMiniQrcodeUnbindModel{}
	return &this
}

// GetRouteGroup returns the RouteGroup field value if set, zero value otherwise.
func (o *AlipayOpenMiniQrcodeUnbindModel) GetRouteGroup() string {
	if o == nil || IsNil(o.RouteGroup) {
		var ret string
		return ret
	}
	return *o.RouteGroup
}

// GetRouteGroupOk returns a tuple with the RouteGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniQrcodeUnbindModel) GetRouteGroupOk() (*string, bool) {
	if o == nil || IsNil(o.RouteGroup) {
		return nil, false
	}
	return o.RouteGroup, true
}

// HasRouteGroup returns a boolean if a field has been set.
func (o *AlipayOpenMiniQrcodeUnbindModel) HasRouteGroup() bool {
	if o != nil && !IsNil(o.RouteGroup) {
		return true
	}

	return false
}

// SetRouteGroup gets a reference to the given string and assigns it to the RouteGroup field.
func (o *AlipayOpenMiniQrcodeUnbindModel) SetRouteGroup(v string) {
	o.RouteGroup = &v
}

func (o AlipayOpenMiniQrcodeUnbindModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniQrcodeUnbindModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RouteGroup) {
		toSerialize["route_group"] = o.RouteGroup
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniQrcodeUnbindModel struct {
	value *AlipayOpenMiniQrcodeUnbindModel
	isSet bool
}

func (v NullableAlipayOpenMiniQrcodeUnbindModel) Get() *AlipayOpenMiniQrcodeUnbindModel {
	return v.value
}

func (v *NullableAlipayOpenMiniQrcodeUnbindModel) Set(val *AlipayOpenMiniQrcodeUnbindModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniQrcodeUnbindModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniQrcodeUnbindModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniQrcodeUnbindModel(val *AlipayOpenMiniQrcodeUnbindModel) *NullableAlipayOpenMiniQrcodeUnbindModel {
	return &NullableAlipayOpenMiniQrcodeUnbindModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniQrcodeUnbindModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniQrcodeUnbindModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


