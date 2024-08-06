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

// checks if the AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel{}

// AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel struct for AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel
type AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel struct {
	// 开票金额
	InvoiceFee *string `json:"invoice_fee,omitempty"`
	// 整体的运单开票状态，0：处理中 1：开票成功 2：开票失败
	Status *int32 `json:"status,omitempty"`
	// 即时配送运单列表
	WaybillInvoices []WaybillInvoiceIstd `json:"waybill_invoices,omitempty"`
}

// NewAlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel instantiates a new AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel() *AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel {
	this := AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel{}
	return &this
}

// NewAlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModelWithDefaults instantiates a new AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModelWithDefaults() *AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel {
	this := AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel{}
	return &this
}

// GetInvoiceFee returns the InvoiceFee field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) GetInvoiceFee() string {
	if o == nil || IsNil(o.InvoiceFee) {
		var ret string
		return ret
	}
	return *o.InvoiceFee
}

// GetInvoiceFeeOk returns a tuple with the InvoiceFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) GetInvoiceFeeOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceFee) {
		return nil, false
	}
	return o.InvoiceFee, true
}

// HasInvoiceFee returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) HasInvoiceFee() bool {
	if o != nil && !IsNil(o.InvoiceFee) {
		return true
	}

	return false
}

// SetInvoiceFee gets a reference to the given string and assigns it to the InvoiceFee field.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) SetInvoiceFee(v string) {
	o.InvoiceFee = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) SetStatus(v int32) {
	o.Status = &v
}

// GetWaybillInvoices returns the WaybillInvoices field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) GetWaybillInvoices() []WaybillInvoiceIstd {
	if o == nil || IsNil(o.WaybillInvoices) {
		var ret []WaybillInvoiceIstd
		return ret
	}
	return o.WaybillInvoices
}

// GetWaybillInvoicesOk returns a tuple with the WaybillInvoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) GetWaybillInvoicesOk() ([]WaybillInvoiceIstd, bool) {
	if o == nil || IsNil(o.WaybillInvoices) {
		return nil, false
	}
	return o.WaybillInvoices, true
}

// HasWaybillInvoices returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) HasWaybillInvoices() bool {
	if o != nil && !IsNil(o.WaybillInvoices) {
		return true
	}

	return false
}

// SetWaybillInvoices gets a reference to the given []WaybillInvoiceIstd and assigns it to the WaybillInvoices field.
func (o *AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) SetWaybillInvoices(v []WaybillInvoiceIstd) {
	o.WaybillInvoices = v
}

func (o AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoiceFee) {
		toSerialize["invoice_fee"] = o.InvoiceFee
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.WaybillInvoices) {
		toSerialize["waybill_invoices"] = o.WaybillInvoices
	}
	return toSerialize, nil
}

type NullableAlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel struct {
	value *AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel
	isSet bool
}

func (v NullableAlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) Get() *AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel {
	return v.value
}

func (v *NullableAlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) Set(val *AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel(val *AlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) *NullableAlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel {
	return &NullableAlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceLogisticsInvoiceIstdwaybillCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
