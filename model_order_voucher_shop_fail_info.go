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

// checks if the OrderVoucherShopFailInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderVoucherShopFailInfo{}

// OrderVoucherShopFailInfo struct for OrderVoucherShopFailInfo
type OrderVoucherShopFailInfo struct {
	// 请求失败提示信息。
	FailMessage *string `json:"fail_message,omitempty"`
	// 请求失败的原因。
	FailReasons []string `json:"fail_reasons,omitempty"`
	// 支付宝侧蚂蚁店铺 id。
	ShopId *string `json:"shop_id,omitempty"`
}

// NewOrderVoucherShopFailInfo instantiates a new OrderVoucherShopFailInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderVoucherShopFailInfo() *OrderVoucherShopFailInfo {
	this := OrderVoucherShopFailInfo{}
	return &this
}

// NewOrderVoucherShopFailInfoWithDefaults instantiates a new OrderVoucherShopFailInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderVoucherShopFailInfoWithDefaults() *OrderVoucherShopFailInfo {
	this := OrderVoucherShopFailInfo{}
	return &this
}

// GetFailMessage returns the FailMessage field value if set, zero value otherwise.
func (o *OrderVoucherShopFailInfo) GetFailMessage() string {
	if o == nil || IsNil(o.FailMessage) {
		var ret string
		return ret
	}
	return *o.FailMessage
}

// GetFailMessageOk returns a tuple with the FailMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderVoucherShopFailInfo) GetFailMessageOk() (*string, bool) {
	if o == nil || IsNil(o.FailMessage) {
		return nil, false
	}
	return o.FailMessage, true
}

// HasFailMessage returns a boolean if a field has been set.
func (o *OrderVoucherShopFailInfo) HasFailMessage() bool {
	if o != nil && !IsNil(o.FailMessage) {
		return true
	}

	return false
}

// SetFailMessage gets a reference to the given string and assigns it to the FailMessage field.
func (o *OrderVoucherShopFailInfo) SetFailMessage(v string) {
	o.FailMessage = &v
}

// GetFailReasons returns the FailReasons field value if set, zero value otherwise.
func (o *OrderVoucherShopFailInfo) GetFailReasons() []string {
	if o == nil || IsNil(o.FailReasons) {
		var ret []string
		return ret
	}
	return o.FailReasons
}

// GetFailReasonsOk returns a tuple with the FailReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderVoucherShopFailInfo) GetFailReasonsOk() ([]string, bool) {
	if o == nil || IsNil(o.FailReasons) {
		return nil, false
	}
	return o.FailReasons, true
}

// HasFailReasons returns a boolean if a field has been set.
func (o *OrderVoucherShopFailInfo) HasFailReasons() bool {
	if o != nil && !IsNil(o.FailReasons) {
		return true
	}

	return false
}

// SetFailReasons gets a reference to the given []string and assigns it to the FailReasons field.
func (o *OrderVoucherShopFailInfo) SetFailReasons(v []string) {
	o.FailReasons = v
}

// GetShopId returns the ShopId field value if set, zero value otherwise.
func (o *OrderVoucherShopFailInfo) GetShopId() string {
	if o == nil || IsNil(o.ShopId) {
		var ret string
		return ret
	}
	return *o.ShopId
}

// GetShopIdOk returns a tuple with the ShopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderVoucherShopFailInfo) GetShopIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShopId) {
		return nil, false
	}
	return o.ShopId, true
}

// HasShopId returns a boolean if a field has been set.
func (o *OrderVoucherShopFailInfo) HasShopId() bool {
	if o != nil && !IsNil(o.ShopId) {
		return true
	}

	return false
}

// SetShopId gets a reference to the given string and assigns it to the ShopId field.
func (o *OrderVoucherShopFailInfo) SetShopId(v string) {
	o.ShopId = &v
}

func (o OrderVoucherShopFailInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderVoucherShopFailInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FailMessage) {
		toSerialize["fail_message"] = o.FailMessage
	}
	if !IsNil(o.FailReasons) {
		toSerialize["fail_reasons"] = o.FailReasons
	}
	if !IsNil(o.ShopId) {
		toSerialize["shop_id"] = o.ShopId
	}
	return toSerialize, nil
}

type NullableOrderVoucherShopFailInfo struct {
	value *OrderVoucherShopFailInfo
	isSet bool
}

func (v NullableOrderVoucherShopFailInfo) Get() *OrderVoucherShopFailInfo {
	return v.value
}

func (v *NullableOrderVoucherShopFailInfo) Set(val *OrderVoucherShopFailInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderVoucherShopFailInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderVoucherShopFailInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderVoucherShopFailInfo(val *OrderVoucherShopFailInfo) *NullableOrderVoucherShopFailInfo {
	return &NullableOrderVoucherShopFailInfo{value: val, isSet: true}
}

func (v NullableOrderVoucherShopFailInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderVoucherShopFailInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


