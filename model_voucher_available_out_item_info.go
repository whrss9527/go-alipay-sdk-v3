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

// checks if the VoucherAvailableOutItemInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherAvailableOutItemInfo{}

// VoucherAvailableOutItemInfo struct for VoucherAvailableOutItemInfo
type VoucherAvailableOutItemInfo struct {
	// 小程序ID
	ItemAppId *string `json:"item_app_id,omitempty"`
	// 外部商品ID
	OutItemId *string `json:"out_item_id,omitempty"`
}

// NewVoucherAvailableOutItemInfo instantiates a new VoucherAvailableOutItemInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherAvailableOutItemInfo() *VoucherAvailableOutItemInfo {
	this := VoucherAvailableOutItemInfo{}
	return &this
}

// NewVoucherAvailableOutItemInfoWithDefaults instantiates a new VoucherAvailableOutItemInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherAvailableOutItemInfoWithDefaults() *VoucherAvailableOutItemInfo {
	this := VoucherAvailableOutItemInfo{}
	return &this
}

// GetItemAppId returns the ItemAppId field value if set, zero value otherwise.
func (o *VoucherAvailableOutItemInfo) GetItemAppId() string {
	if o == nil || IsNil(o.ItemAppId) {
		var ret string
		return ret
	}
	return *o.ItemAppId
}

// GetItemAppIdOk returns a tuple with the ItemAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableOutItemInfo) GetItemAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemAppId) {
		return nil, false
	}
	return o.ItemAppId, true
}

// HasItemAppId returns a boolean if a field has been set.
func (o *VoucherAvailableOutItemInfo) HasItemAppId() bool {
	if o != nil && !IsNil(o.ItemAppId) {
		return true
	}

	return false
}

// SetItemAppId gets a reference to the given string and assigns it to the ItemAppId field.
func (o *VoucherAvailableOutItemInfo) SetItemAppId(v string) {
	o.ItemAppId = &v
}

// GetOutItemId returns the OutItemId field value if set, zero value otherwise.
func (o *VoucherAvailableOutItemInfo) GetOutItemId() string {
	if o == nil || IsNil(o.OutItemId) {
		var ret string
		return ret
	}
	return *o.OutItemId
}

// GetOutItemIdOk returns a tuple with the OutItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherAvailableOutItemInfo) GetOutItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.OutItemId) {
		return nil, false
	}
	return o.OutItemId, true
}

// HasOutItemId returns a boolean if a field has been set.
func (o *VoucherAvailableOutItemInfo) HasOutItemId() bool {
	if o != nil && !IsNil(o.OutItemId) {
		return true
	}

	return false
}

// SetOutItemId gets a reference to the given string and assigns it to the OutItemId field.
func (o *VoucherAvailableOutItemInfo) SetOutItemId(v string) {
	o.OutItemId = &v
}

func (o VoucherAvailableOutItemInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherAvailableOutItemInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemAppId) {
		toSerialize["item_app_id"] = o.ItemAppId
	}
	if !IsNil(o.OutItemId) {
		toSerialize["out_item_id"] = o.OutItemId
	}
	return toSerialize, nil
}

type NullableVoucherAvailableOutItemInfo struct {
	value *VoucherAvailableOutItemInfo
	isSet bool
}

func (v NullableVoucherAvailableOutItemInfo) Get() *VoucherAvailableOutItemInfo {
	return v.value
}

func (v *NullableVoucherAvailableOutItemInfo) Set(val *VoucherAvailableOutItemInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherAvailableOutItemInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherAvailableOutItemInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherAvailableOutItemInfo(val *VoucherAvailableOutItemInfo) *NullableVoucherAvailableOutItemInfo {
	return &NullableVoucherAvailableOutItemInfo{value: val, isSet: true}
}

func (v NullableVoucherAvailableOutItemInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherAvailableOutItemInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


