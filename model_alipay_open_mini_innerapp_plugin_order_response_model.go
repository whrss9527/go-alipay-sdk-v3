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

// checks if the AlipayOpenMiniInnerappPluginOrderResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniInnerappPluginOrderResponseModel{}

// AlipayOpenMiniInnerappPluginOrderResponseModel struct for AlipayOpenMiniInnerappPluginOrderResponseModel
type AlipayOpenMiniInnerappPluginOrderResponseModel struct {
	// 订单号
	OrderId *string `json:"order_id,omitempty"`
}

// NewAlipayOpenMiniInnerappPluginOrderResponseModel instantiates a new AlipayOpenMiniInnerappPluginOrderResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniInnerappPluginOrderResponseModel() *AlipayOpenMiniInnerappPluginOrderResponseModel {
	this := AlipayOpenMiniInnerappPluginOrderResponseModel{}
	return &this
}

// NewAlipayOpenMiniInnerappPluginOrderResponseModelWithDefaults instantiates a new AlipayOpenMiniInnerappPluginOrderResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniInnerappPluginOrderResponseModelWithDefaults() *AlipayOpenMiniInnerappPluginOrderResponseModel {
	this := AlipayOpenMiniInnerappPluginOrderResponseModel{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *AlipayOpenMiniInnerappPluginOrderResponseModel) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniInnerappPluginOrderResponseModel) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *AlipayOpenMiniInnerappPluginOrderResponseModel) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *AlipayOpenMiniInnerappPluginOrderResponseModel) SetOrderId(v string) {
	o.OrderId = &v
}

func (o AlipayOpenMiniInnerappPluginOrderResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniInnerappPluginOrderResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniInnerappPluginOrderResponseModel struct {
	value *AlipayOpenMiniInnerappPluginOrderResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniInnerappPluginOrderResponseModel) Get() *AlipayOpenMiniInnerappPluginOrderResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerappPluginOrderResponseModel) Set(val *AlipayOpenMiniInnerappPluginOrderResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerappPluginOrderResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerappPluginOrderResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerappPluginOrderResponseModel(val *AlipayOpenMiniInnerappPluginOrderResponseModel) *NullableAlipayOpenMiniInnerappPluginOrderResponseModel {
	return &NullableAlipayOpenMiniInnerappPluginOrderResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerappPluginOrderResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerappPluginOrderResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
