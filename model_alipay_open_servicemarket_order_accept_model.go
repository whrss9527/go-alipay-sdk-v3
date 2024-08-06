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

// checks if the AlipayOpenServicemarketOrderAcceptModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenServicemarketOrderAcceptModel{}

// AlipayOpenServicemarketOrderAcceptModel struct for AlipayOpenServicemarketOrderAcceptModel
type AlipayOpenServicemarketOrderAcceptModel struct {
	// 服务商品订单ID
	CommodityOrderId *string `json:"commodity_order_id,omitempty"`
}

// NewAlipayOpenServicemarketOrderAcceptModel instantiates a new AlipayOpenServicemarketOrderAcceptModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenServicemarketOrderAcceptModel() *AlipayOpenServicemarketOrderAcceptModel {
	this := AlipayOpenServicemarketOrderAcceptModel{}
	return &this
}

// NewAlipayOpenServicemarketOrderAcceptModelWithDefaults instantiates a new AlipayOpenServicemarketOrderAcceptModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenServicemarketOrderAcceptModelWithDefaults() *AlipayOpenServicemarketOrderAcceptModel {
	this := AlipayOpenServicemarketOrderAcceptModel{}
	return &this
}

// GetCommodityOrderId returns the CommodityOrderId field value if set, zero value otherwise.
func (o *AlipayOpenServicemarketOrderAcceptModel) GetCommodityOrderId() string {
	if o == nil || IsNil(o.CommodityOrderId) {
		var ret string
		return ret
	}
	return *o.CommodityOrderId
}

// GetCommodityOrderIdOk returns a tuple with the CommodityOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenServicemarketOrderAcceptModel) GetCommodityOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.CommodityOrderId) {
		return nil, false
	}
	return o.CommodityOrderId, true
}

// HasCommodityOrderId returns a boolean if a field has been set.
func (o *AlipayOpenServicemarketOrderAcceptModel) HasCommodityOrderId() bool {
	if o != nil && !IsNil(o.CommodityOrderId) {
		return true
	}

	return false
}

// SetCommodityOrderId gets a reference to the given string and assigns it to the CommodityOrderId field.
func (o *AlipayOpenServicemarketOrderAcceptModel) SetCommodityOrderId(v string) {
	o.CommodityOrderId = &v
}

func (o AlipayOpenServicemarketOrderAcceptModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenServicemarketOrderAcceptModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommodityOrderId) {
		toSerialize["commodity_order_id"] = o.CommodityOrderId
	}
	return toSerialize, nil
}

type NullableAlipayOpenServicemarketOrderAcceptModel struct {
	value *AlipayOpenServicemarketOrderAcceptModel
	isSet bool
}

func (v NullableAlipayOpenServicemarketOrderAcceptModel) Get() *AlipayOpenServicemarketOrderAcceptModel {
	return v.value
}

func (v *NullableAlipayOpenServicemarketOrderAcceptModel) Set(val *AlipayOpenServicemarketOrderAcceptModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenServicemarketOrderAcceptModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenServicemarketOrderAcceptModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenServicemarketOrderAcceptModel(val *AlipayOpenServicemarketOrderAcceptModel) *NullableAlipayOpenServicemarketOrderAcceptModel {
	return &NullableAlipayOpenServicemarketOrderAcceptModel{value: val, isSet: true}
}

func (v NullableAlipayOpenServicemarketOrderAcceptModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenServicemarketOrderAcceptModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


