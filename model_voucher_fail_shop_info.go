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

// checks if the VoucherFailShopInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoucherFailShopInfo{}

// VoucherFailShopInfo struct for VoucherFailShopInfo
type VoucherFailShopInfo struct {
	// 请求失败提示信息。
	FailMessage *string `json:"fail_message,omitempty"`
	// 请求失败的原因。
	FailReasons []string `json:"fail_reasons,omitempty"`
	// 请求失败的物理门店 id。
	RealShopId *string `json:"real_shop_id,omitempty"`
	// 支付宝侧蚂蚁店铺 id。
	ShopId *string `json:"shop_id,omitempty"`
}

// NewVoucherFailShopInfo instantiates a new VoucherFailShopInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoucherFailShopInfo() *VoucherFailShopInfo {
	this := VoucherFailShopInfo{}
	return &this
}

// NewVoucherFailShopInfoWithDefaults instantiates a new VoucherFailShopInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoucherFailShopInfoWithDefaults() *VoucherFailShopInfo {
	this := VoucherFailShopInfo{}
	return &this
}

// GetFailMessage returns the FailMessage field value if set, zero value otherwise.
func (o *VoucherFailShopInfo) GetFailMessage() string {
	if o == nil || IsNil(o.FailMessage) {
		var ret string
		return ret
	}
	return *o.FailMessage
}

// GetFailMessageOk returns a tuple with the FailMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherFailShopInfo) GetFailMessageOk() (*string, bool) {
	if o == nil || IsNil(o.FailMessage) {
		return nil, false
	}
	return o.FailMessage, true
}

// HasFailMessage returns a boolean if a field has been set.
func (o *VoucherFailShopInfo) HasFailMessage() bool {
	if o != nil && !IsNil(o.FailMessage) {
		return true
	}

	return false
}

// SetFailMessage gets a reference to the given string and assigns it to the FailMessage field.
func (o *VoucherFailShopInfo) SetFailMessage(v string) {
	o.FailMessage = &v
}

// GetFailReasons returns the FailReasons field value if set, zero value otherwise.
func (o *VoucherFailShopInfo) GetFailReasons() []string {
	if o == nil || IsNil(o.FailReasons) {
		var ret []string
		return ret
	}
	return o.FailReasons
}

// GetFailReasonsOk returns a tuple with the FailReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherFailShopInfo) GetFailReasonsOk() ([]string, bool) {
	if o == nil || IsNil(o.FailReasons) {
		return nil, false
	}
	return o.FailReasons, true
}

// HasFailReasons returns a boolean if a field has been set.
func (o *VoucherFailShopInfo) HasFailReasons() bool {
	if o != nil && !IsNil(o.FailReasons) {
		return true
	}

	return false
}

// SetFailReasons gets a reference to the given []string and assigns it to the FailReasons field.
func (o *VoucherFailShopInfo) SetFailReasons(v []string) {
	o.FailReasons = v
}

// GetRealShopId returns the RealShopId field value if set, zero value otherwise.
func (o *VoucherFailShopInfo) GetRealShopId() string {
	if o == nil || IsNil(o.RealShopId) {
		var ret string
		return ret
	}
	return *o.RealShopId
}

// GetRealShopIdOk returns a tuple with the RealShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherFailShopInfo) GetRealShopIdOk() (*string, bool) {
	if o == nil || IsNil(o.RealShopId) {
		return nil, false
	}
	return o.RealShopId, true
}

// HasRealShopId returns a boolean if a field has been set.
func (o *VoucherFailShopInfo) HasRealShopId() bool {
	if o != nil && !IsNil(o.RealShopId) {
		return true
	}

	return false
}

// SetRealShopId gets a reference to the given string and assigns it to the RealShopId field.
func (o *VoucherFailShopInfo) SetRealShopId(v string) {
	o.RealShopId = &v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *VoucherFailShopInfo) GetShopId() string {
	if o == nil || IsNil(o.ShopId) {
		var ret string
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoucherFailShopInfo) GetShopIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *VoucherFailShopInfo) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given string and assigns it to the ShopId field.
func (o *VoucherFailShopInfo) SetShopId(v string) {
	o.ShopId = &v
}

func (o VoucherFailShopInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoucherFailShopInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FailMessage) {
		toSerialize["fail_message"] = o.FailMessage
	}
	if !IsNil(o.FailReasons) {
		toSerialize["fail_reasons"] = o.FailReasons
	}
	if !IsNil(o.RealShopId) {
		toSerialize["real_shop_id"] = o.RealShopId
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	return toSerialize, nil
}

type NullableVoucherFailShopInfo struct {
	value *VoucherFailShopInfo
	isSet bool
}

func (v NullableVoucherFailShopInfo) Get() *VoucherFailShopInfo {
	return v.value
}

func (v *NullableVoucherFailShopInfo) Set(val *VoucherFailShopInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVoucherFailShopInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVoucherFailShopInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoucherFailShopInfo(val *VoucherFailShopInfo) *NullableVoucherFailShopInfo {
	return &NullableVoucherFailShopInfo{value: val, isSet: true}
}

func (v NullableVoucherFailShopInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoucherFailShopInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


