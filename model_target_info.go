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

// checks if the TargetInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetInfo{}

// TargetInfo struct for TargetInfo
type TargetInfo struct {
	// 应用id，例如小程序id
	TargetId *string `json:"target_id,omitempty"`
	// 应用类型 小程序传入：APPID 生活号传入：PUBLICID
	TargetType *string `json:"target_type,omitempty"`
}

// NewTargetInfo instantiates a new TargetInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetInfo() *TargetInfo {
	this := TargetInfo{}
	return &this
}

// NewTargetInfoWithDefaults instantiates a new TargetInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetInfoWithDefaults() *TargetInfo {
	this := TargetInfo{}
	return &this
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *TargetInfo) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetInfo) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *TargetInfo) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *TargetInfo) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *TargetInfo) GetTargetType() string {
	if o == nil || IsNil(o.TargetType) {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetInfo) GetTargetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetType) {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *TargetInfo) HasTargetType() bool {
	if o != nil && !IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *TargetInfo) SetTargetType(v string) {
	o.TargetType = &v
}

func (o TargetInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetId) {
		toSerialize["target_id"] = o.TargetId
	}
	if !IsNil(o.TargetType) {
		toSerialize["target_type"] = o.TargetType
	}
	return toSerialize, nil
}

type NullableTargetInfo struct {
	value *TargetInfo
	isSet bool
}

func (v NullableTargetInfo) Get() *TargetInfo {
	return v.value
}

func (v *NullableTargetInfo) Set(val *TargetInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetInfo(val *TargetInfo) *NullableTargetInfo {
	return &NullableTargetInfo{value: val, isSet: true}
}

func (v NullableTargetInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
