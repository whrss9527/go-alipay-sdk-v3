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

// checks if the AlipayOpenServicemarketOrderCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenServicemarketOrderCreateResponseModel{}

// AlipayOpenServicemarketOrderCreateResponseModel struct for AlipayOpenServicemarketOrderCreateResponseModel
type AlipayOpenServicemarketOrderCreateResponseModel struct {
	// 订单号
	CommodityOrderId *string `json:"commodity_order_id,omitempty"`
	// 支付宝分配给开发者的应用ID
	MiniAppId *string `json:"mini_app_id,omitempty"`
}

// NewAlipayOpenServicemarketOrderCreateResponseModel instantiates a new AlipayOpenServicemarketOrderCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenServicemarketOrderCreateResponseModel() *AlipayOpenServicemarketOrderCreateResponseModel {
	this := AlipayOpenServicemarketOrderCreateResponseModel{}
	return &this
}

// NewAlipayOpenServicemarketOrderCreateResponseModelWithDefaults instantiates a new AlipayOpenServicemarketOrderCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenServicemarketOrderCreateResponseModelWithDefaults() *AlipayOpenServicemarketOrderCreateResponseModel {
	this := AlipayOpenServicemarketOrderCreateResponseModel{}
	return &this
}

// GetCommodityOrderId returns the CommodityOrderId field value if set, zero value otherwise.
func (o *AlipayOpenServicemarketOrderCreateResponseModel) GetCommodityOrderId() string {
	if o == nil || IsNil(o.CommodityOrderId) {
		var ret string
		return ret
	}
	return *o.CommodityOrderId
}

// GetCommodityOrderIdOk returns a tuple with the CommodityOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenServicemarketOrderCreateResponseModel) GetCommodityOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.CommodityOrderId) {
		return nil, false
	}
	return o.CommodityOrderId, true
}

// HasCommodityOrderId returns a boolean if a field has been set.
func (o *AlipayOpenServicemarketOrderCreateResponseModel) HasCommodityOrderId() bool {
	if o != nil && !IsNil(o.CommodityOrderId) {
		return true
	}

	return false
}

// SetCommodityOrderId gets a reference to the given string and assigns it to the CommodityOrderId field.
func (o *AlipayOpenServicemarketOrderCreateResponseModel) SetCommodityOrderId(v string) {
	o.CommodityOrderId = &v
}

// GetMiniAppId returns the MiniAppId field value if set, zero value otherwise.
func (o *AlipayOpenServicemarketOrderCreateResponseModel) GetMiniAppId() string {
	if o == nil || IsNil(o.MiniAppId) {
		var ret string
		return ret
	}
	return *o.MiniAppId
}

// GetMiniAppIdOk returns a tuple with the MiniAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenServicemarketOrderCreateResponseModel) GetMiniAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.MiniAppId) {
		return nil, false
	}
	return o.MiniAppId, true
}

// HasMiniAppId returns a boolean if a field has been set.
func (o *AlipayOpenServicemarketOrderCreateResponseModel) HasMiniAppId() bool {
	if o != nil && !IsNil(o.MiniAppId) {
		return true
	}

	return false
}

// SetMiniAppId gets a reference to the given string and assigns it to the MiniAppId field.
func (o *AlipayOpenServicemarketOrderCreateResponseModel) SetMiniAppId(v string) {
	o.MiniAppId = &v
}

func (o AlipayOpenServicemarketOrderCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenServicemarketOrderCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommodityOrderId) {
		toSerialize["commodity_order_id"] = o.CommodityOrderId
	}
	if !IsNil(o.MiniAppId) {
		toSerialize["mini_app_id"] = o.MiniAppId
	}
	return toSerialize, nil
}

type NullableAlipayOpenServicemarketOrderCreateResponseModel struct {
	value *AlipayOpenServicemarketOrderCreateResponseModel
	isSet bool
}

func (v NullableAlipayOpenServicemarketOrderCreateResponseModel) Get() *AlipayOpenServicemarketOrderCreateResponseModel {
	return v.value
}

func (v *NullableAlipayOpenServicemarketOrderCreateResponseModel) Set(val *AlipayOpenServicemarketOrderCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenServicemarketOrderCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenServicemarketOrderCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenServicemarketOrderCreateResponseModel(val *AlipayOpenServicemarketOrderCreateResponseModel) *NullableAlipayOpenServicemarketOrderCreateResponseModel {
	return &NullableAlipayOpenServicemarketOrderCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenServicemarketOrderCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenServicemarketOrderCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

