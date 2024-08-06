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

// checks if the AlipayEbppInvoiceEcorderOrderQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceEcorderOrderQueryResponseModel{}

// AlipayEbppInvoiceEcorderOrderQueryResponseModel struct for AlipayEbppInvoiceEcorderOrderQueryResponseModel
type AlipayEbppInvoiceEcorderOrderQueryResponseModel struct {
	OrderInfo *EcOrderItem `json:"order_info,omitempty"`
	// 企业码子订单详情列表
	SubOrderList []EcOrderItem `json:"sub_order_list,omitempty"`
}

// NewAlipayEbppInvoiceEcorderOrderQueryResponseModel instantiates a new AlipayEbppInvoiceEcorderOrderQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceEcorderOrderQueryResponseModel() *AlipayEbppInvoiceEcorderOrderQueryResponseModel {
	this := AlipayEbppInvoiceEcorderOrderQueryResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceEcorderOrderQueryResponseModelWithDefaults instantiates a new AlipayEbppInvoiceEcorderOrderQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceEcorderOrderQueryResponseModelWithDefaults() *AlipayEbppInvoiceEcorderOrderQueryResponseModel {
	this := AlipayEbppInvoiceEcorderOrderQueryResponseModel{}
	return &this
}

// GetOrderInfo returns the OrderInfo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEcorderOrderQueryResponseModel) GetOrderInfo() EcOrderItem {
	if o == nil || IsNil(o.OrderInfo) {
		var ret EcOrderItem
		return ret
	}
	return *o.OrderInfo
}

// GetOrderInfoOk returns a tuple with the OrderInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEcorderOrderQueryResponseModel) GetOrderInfoOk() (*EcOrderItem, bool) {
	if o == nil || IsNil(o.OrderInfo) {
		return nil, false
	}
	return o.OrderInfo, true
}

// HasOrderInfo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEcorderOrderQueryResponseModel) HasOrderInfo() bool {
	if o != nil && !IsNil(o.OrderInfo) {
		return true
	}

	return false
}

// SetOrderInfo gets a reference to the given EcOrderItem and assigns it to the OrderInfo field.
func (o *AlipayEbppInvoiceEcorderOrderQueryResponseModel) SetOrderInfo(v EcOrderItem) {
	o.OrderInfo = &v
}

// GetSubOrderList returns the SubOrderList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEcorderOrderQueryResponseModel) GetSubOrderList() []EcOrderItem {
	if o == nil || IsNil(o.SubOrderList) {
		var ret []EcOrderItem
		return ret
	}
	return o.SubOrderList
}

// GetSubOrderListOk returns a tuple with the SubOrderList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEcorderOrderQueryResponseModel) GetSubOrderListOk() ([]EcOrderItem, bool) {
	if o == nil || IsNil(o.SubOrderList) {
		return nil, false
	}
	return o.SubOrderList, true
}

// HasSubOrderList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEcorderOrderQueryResponseModel) HasSubOrderList() bool {
	if o != nil && !IsNil(o.SubOrderList) {
		return true
	}

	return false
}

// SetSubOrderList gets a reference to the given []EcOrderItem and assigns it to the SubOrderList field.
func (o *AlipayEbppInvoiceEcorderOrderQueryResponseModel) SetSubOrderList(v []EcOrderItem) {
	o.SubOrderList = v
}

func (o AlipayEbppInvoiceEcorderOrderQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceEcorderOrderQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderInfo) {
		toSerialize["order_info"] = o.OrderInfo
	}
	if !IsNil(o.SubOrderList) {
		toSerialize["sub_order_list"] = o.SubOrderList
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceEcorderOrderQueryResponseModel struct {
	value *AlipayEbppInvoiceEcorderOrderQueryResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceEcorderOrderQueryResponseModel) Get() *AlipayEbppInvoiceEcorderOrderQueryResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceEcorderOrderQueryResponseModel) Set(val *AlipayEbppInvoiceEcorderOrderQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceEcorderOrderQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceEcorderOrderQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceEcorderOrderQueryResponseModel(val *AlipayEbppInvoiceEcorderOrderQueryResponseModel) *NullableAlipayEbppInvoiceEcorderOrderQueryResponseModel {
	return &NullableAlipayEbppInvoiceEcorderOrderQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceEcorderOrderQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceEcorderOrderQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
