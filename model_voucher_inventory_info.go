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

// checks if the VoucherInventoryInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherInventoryInfo{}

// VoucherInventoryInfo struct for VoucherInventoryInfo
type VoucherInventoryInfo struct {
	// 已发放数量。
	SendCount *int32 `json:"send_count,omitempty"`
	// 已核销数量。
	UseCount *int32 `json:"use_count,omitempty"`
}

// NewVoucherInventoryInfo instantiates a new VoucherInventoryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherInventoryInfo() *VoucherInventoryInfo {
	this := VoucherInventoryInfo{}
	return &this
}

// NewVoucherInventoryInfoWithDefaults instantiates a new VoucherInventoryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherInventoryInfoWithDefaults() *VoucherInventoryInfo {
	this := VoucherInventoryInfo{}
	return &this
}

// GetSendCount returns the SendCount field value if set, zero value otherwise.
func (o *VoucherInventoryInfo) GetSendCount() int32 {
	if o == nil || IsNil(o.SendCount) {
		var ret int32
		return ret
	}
	return *o.SendCount
}

// GetSendCountOk returns a tuple with the SendCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherInventoryInfo) GetSendCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SendCount) {
		return nil, false
	}
	return o.SendCount, true
}

// HasSendCount returns a boolean if a field has been set.
func (o *VoucherInventoryInfo) HasSendCount() bool {
	if o != nil && !IsNil(o.SendCount) {
		return true
	}

	return false
}

// SetSendCount gets a reference to the given int32 and assigns it to the SendCount field.
func (o *VoucherInventoryInfo) SetSendCount(v int32) {
	o.SendCount = &v
}

// GetUseCount returns the UseCount field value if set, zero value otherwise.
func (o *VoucherInventoryInfo) GetUseCount() int32 {
	if o == nil || IsNil(o.UseCount) {
		var ret int32
		return ret
	}
	return *o.UseCount
}

// GetUseCountOk returns a tuple with the UseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherInventoryInfo) GetUseCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UseCount) {
		return nil, false
	}
	return o.UseCount, true
}

// HasUseCount returns a boolean if a field has been set.
func (o *VoucherInventoryInfo) HasUseCount() bool {
	if o != nil && !IsNil(o.UseCount) {
		return true
	}

	return false
}

// SetUseCount gets a reference to the given int32 and assigns it to the UseCount field.
func (o *VoucherInventoryInfo) SetUseCount(v int32) {
	o.UseCount = &v
}

func (o VoucherInventoryInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherInventoryInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SendCount) {
		toSerialize["send_count"] = o.SendCount
	}
	if !IsNil(o.UseCount) {
		toSerialize["use_count"] = o.UseCount
	}
	return toSerialize, nil
}

type NullableVoucherInventoryInfo struct {
	value *VoucherInventoryInfo
	isSet bool
}

func (v NullableVoucherInventoryInfo) Get() *VoucherInventoryInfo {
	return v.value
}

func (v *NullableVoucherInventoryInfo) Set(val *VoucherInventoryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherInventoryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherInventoryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherInventoryInfo(val *VoucherInventoryInfo) *NullableVoucherInventoryInfo {
	return &NullableVoucherInventoryInfo{value: val, isSet: true}
}

func (v NullableVoucherInventoryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherInventoryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


