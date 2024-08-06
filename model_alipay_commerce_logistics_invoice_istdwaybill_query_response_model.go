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

// checks if the AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel{}

// AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel struct for AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel
type AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel struct {
	// 整体的运单开票状态，0：处理中 1：开票成功 2：开票失败
	Status *int32 `json:"status,omitempty"`
	// 即时配送运单列表
	WaybillInvoices []WaybillInvoiceQueryIstd `json:"waybill_invoices,omitempty"`
}

// NewAlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel instantiates a new AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel() *AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel {
	this := AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel{}
	return &this
}

// NewAlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModelWithDefaults instantiates a new AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModelWithDefaults() *AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel {
	this := AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel) SetStatus(v int32) {
	o.Status = &v
}

// GetWaybillInvoices returns the WaybillInvoices field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel) GetWaybillInvoices() []WaybillInvoiceQueryIstd {
	if o == nil || IsNil(o.WaybillInvoices) {
		var ret []WaybillInvoiceQueryIstd
		return ret
	}
	return o.WaybillInvoices
}

// GetWaybillInvoicesOk returns a tuple with the WaybillInvoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel) GetWaybillInvoicesOk() ([]WaybillInvoiceQueryIstd, bool) {
	if o == nil || IsNil(o.WaybillInvoices) {
		return nil, false
	}
	return o.WaybillInvoices, true
}

// HasWaybillInvoices returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel) HasWaybillInvoices() bool {
	if o != nil && !IsNil(o.WaybillInvoices) {
		return true
	}

	return false
}

// SetWaybillInvoices gets a reference to the given []WaybillInvoiceQueryIstd and assigns it to the WaybillInvoices field.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel) SetWaybillInvoices(v []WaybillInvoiceQueryIstd) {
	o.WaybillInvoices = v
}

func (o AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.WaybillInvoices) {
		toSerialize["waybill_invoices"] = o.WaybillInvoices
	}
	return toSerialize, nil
}

type NullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel struct {
	value *AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel
	isSet bool
}

func (v NullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel) Get() *AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel {
	return v.value
}

func (v *NullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel) Set(val *AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel(val *AlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel) *NullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel {
	return &NullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceLogisticsInvoiceIstdwaybillQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
