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

// checks if the AlipayOpenMiniIsvCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenMiniIsvCreateResponseModel{}

// AlipayOpenMiniIsvCreateResponseModel struct for AlipayOpenMiniIsvCreateResponseModel
type AlipayOpenMiniIsvCreateResponseModel struct {
	// 小程序代创建订单号
	OrderNo *string `json:"order_no,omitempty"`
}

// NewAlipayOpenMiniIsvCreateResponseModel instantiates a new AlipayOpenMiniIsvCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenMiniIsvCreateResponseModel() *AlipayOpenMiniIsvCreateResponseModel {
	this := AlipayOpenMiniIsvCreateResponseModel{}
	return &this
}

// NewAlipayOpenMiniIsvCreateResponseModelWithDefaults instantiates a new AlipayOpenMiniIsvCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenMiniIsvCreateResponseModelWithDefaults() *AlipayOpenMiniIsvCreateResponseModel {
	this := AlipayOpenMiniIsvCreateResponseModel{}
	return &this
}

// GetOrderNo returns the OrderNo field value if set, zero value otherwise.
func (o *AlipayOpenMiniIsvCreateResponseModel) GetOrderNo() string {
	if o == nil || IsNil(o.OrderNo) {
		var ret string
		return ret
	}
	return *o.OrderNo
}

// GetOrderNoOk returns a tuple with the OrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenMiniIsvCreateResponseModel) GetOrderNoOk() (*string, bool) {
	if o == nil || IsNil(o.OrderNo) {
		return nil, false
	}
	return o.OrderNo, true
}

// HasOrderNo returns a boolean if a field has been set.
func (o *AlipayOpenMiniIsvCreateResponseModel) HasOrderNo() bool {
	if o != nil && !IsNil(o.OrderNo) {
		return true
	}

	return false
}

// SetOrderNo gets a reference to the given string and assigns it to the OrderNo field.
func (o *AlipayOpenMiniIsvCreateResponseModel) SetOrderNo(v string) {
	o.OrderNo = &v
}

func (o AlipayOpenMiniIsvCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenMiniIsvCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderNo) {
		toSerialize["order_no"] = o.OrderNo
	}
	return toSerialize, nil
}

type NullableAlipayOpenMiniIsvCreateResponseModel struct {
	value *AlipayOpenMiniIsvCreateResponseModel
	isSet bool
}

func (v NullableAlipayOpenMiniIsvCreateResponseModel) Get() *AlipayOpenMiniIsvCreateResponseModel {
	return v.value
}

func (v *NullableAlipayOpenMiniIsvCreateResponseModel) Set(val *AlipayOpenMiniIsvCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniIsvCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniIsvCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniIsvCreateResponseModel(val *AlipayOpenMiniIsvCreateResponseModel) *NullableAlipayOpenMiniIsvCreateResponseModel {
	return &NullableAlipayOpenMiniIsvCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniIsvCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniIsvCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
