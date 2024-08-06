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

// checks if the VoucherAvailableItemInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherAvailableItemInfo{}

// VoucherAvailableItemInfo struct for VoucherAvailableItemInfo
type VoucherAvailableItemInfo struct {
	// 内部商品ID
	ItemId *string `json:"item_id,omitempty"`
	// 外部商品信息
	OutItemInfo []VoucherAvailableOutItemInfo `json:"out_item_info,omitempty"`
}

// NewVoucherAvailableItemInfo instantiates a new VoucherAvailableItemInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherAvailableItemInfo() *VoucherAvailableItemInfo {
	this := VoucherAvailableItemInfo{}
	return &this
}

// NewVoucherAvailableItemInfoWithDefaults instantiates a new VoucherAvailableItemInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherAvailableItemInfoWithDefaults() *VoucherAvailableItemInfo {
	this := VoucherAvailableItemInfo{}
	return &this
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *VoucherAvailableItemInfo) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableItemInfo) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *VoucherAvailableItemInfo) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *VoucherAvailableItemInfo) SetItemId(v string) {
	o.ItemId = &v
}

// GetOutItemInfo returns the OutItemInfo field value if set, zero value otherwise.
func (o *VoucherAvailableItemInfo) GetOutItemInfo() []VoucherAvailableOutItemInfo {
	if o == nil || IsNil(o.OutItemInfo) {
		var ret []VoucherAvailableOutItemInfo
		return ret
	}
	return o.OutItemInfo
}

// GetOutItemInfoOk returns a tuple with the OutItemInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableItemInfo) GetOutItemInfoOk() ([]VoucherAvailableOutItemInfo, bool) {
	if o == nil || IsNil(o.OutItemInfo) {
		return nil, false
	}
	return o.OutItemInfo, true
}

// HasOutItemInfo returns a boolean if a field has been set.
func (o *VoucherAvailableItemInfo) HasOutItemInfo() bool {
	if o != nil && !IsNil(o.OutItemInfo) {
		return true
	}

	return false
}

// SetOutItemInfo gets a reference to the given []VoucherAvailableOutItemInfo and assigns it to the OutItemInfo field.
func (o *VoucherAvailableItemInfo) SetOutItemInfo(v []VoucherAvailableOutItemInfo) {
	o.OutItemInfo = v
}

func (o VoucherAvailableItemInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherAvailableItemInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemId) {
		toSerialize["item_id"] = o.ItemId
	}
	if !IsNil(o.OutItemInfo) {
		toSerialize["out_item_info"] = o.OutItemInfo
	}
	return toSerialize, nil
}

type NullableVoucherAvailableItemInfo struct {
	value *VoucherAvailableItemInfo
	isSet bool
}

func (v NullableVoucherAvailableItemInfo) Get() *VoucherAvailableItemInfo {
	return v.value
}

func (v *NullableVoucherAvailableItemInfo) Set(val *VoucherAvailableItemInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherAvailableItemInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherAvailableItemInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherAvailableItemInfo(val *VoucherAvailableItemInfo) *NullableVoucherAvailableItemInfo {
	return &NullableVoucherAvailableItemInfo{value: val, isSet: true}
}

func (v NullableVoucherAvailableItemInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherAvailableItemInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


