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

// checks if the ActivityAppInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityAppInfo{}

// ActivityAppInfo struct for ActivityAppInfo
type ActivityAppInfo struct {
	// 小程序id
	MiniAppId *string `json:"mini_app_id,omitempty"`
}

// NewActivityAppInfo instantiates a new ActivityAppInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityAppInfo() *ActivityAppInfo {
	this := ActivityAppInfo{}
	return &this
}

// NewActivityAppInfoWithDefaults instantiates a new ActivityAppInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityAppInfoWithDefaults() *ActivityAppInfo {
	this := ActivityAppInfo{}
	return &this
}

// GetMiniAppId returns the MiniAppId field value if set, zero value otherwise.
func (o *ActivityAppInfo) GetMiniAppId() string {
	if o == nil || IsNil(o.MiniAppId) {
		var ret string
		return ret
	}
	return *o.MiniAppId
}

// GetMiniAppIdOk returns a tuple with the MiniAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityAppInfo) GetMiniAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppId) {
		return nil, false
	}
	return o.MiniAppId, true
}

// HasMiniAppId returns a boolean if a field has been set.
func (o *ActivityAppInfo) HasMiniAppId() bool {
	if o != nil && !IsNil(o.MiniAppId) {
		return true
	}

	return false
}

// SetMiniAppId gets a reference to the given string and assigns it to the MiniAppId field.
func (o *ActivityAppInfo) SetMiniAppId(v string) {
	o.MiniAppId = &v
}

func (o ActivityAppInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityAppInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MiniAppId) {
		toSerialize["mini_app_id"] = o.MiniAppId
	}
	return toSerialize, nil
}

type NullableActivityAppInfo struct {
	value *ActivityAppInfo
	isSet bool
}

func (v NullableActivityAppInfo) Get() *ActivityAppInfo {
	return v.value
}

func (v *NullableActivityAppInfo) Set(val *ActivityAppInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityAppInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityAppInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityAppInfo(val *ActivityAppInfo) *NullableActivityAppInfo {
	return &NullableActivityAppInfo{value: val, isSet: true}
}

func (v NullableActivityAppInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityAppInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

