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

// checks if the AlipayEbppInvoiceMerchantlistEnterApplyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceMerchantlistEnterApplyResponseModel{}

// AlipayEbppInvoiceMerchantlistEnterApplyResponseModel struct for AlipayEbppInvoiceMerchantlistEnterApplyResponseModel
type AlipayEbppInvoiceMerchantlistEnterApplyResponseModel struct {
	// 工单流水号，支付开票PID模式当前不会返回流水号。
	ProcessId *string `json:"process_id,omitempty"`
}

// NewAlipayEbppInvoiceMerchantlistEnterApplyResponseModel instantiates a new AlipayEbppInvoiceMerchantlistEnterApplyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceMerchantlistEnterApplyResponseModel() *AlipayEbppInvoiceMerchantlistEnterApplyResponseModel {
	this := AlipayEbppInvoiceMerchantlistEnterApplyResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceMerchantlistEnterApplyResponseModelWithDefaults instantiates a new AlipayEbppInvoiceMerchantlistEnterApplyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceMerchantlistEnterApplyResponseModelWithDefaults() *AlipayEbppInvoiceMerchantlistEnterApplyResponseModel {
	this := AlipayEbppInvoiceMerchantlistEnterApplyResponseModel{}
	return &this
}

// GetProcessId returns the ProcessId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceMerchantlistEnterApplyResponseModel) GetProcessId() string {
	if o == nil || IsNil(o.ProcessId) {
		var ret string
		return ret
	}
	return *o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceMerchantlistEnterApplyResponseModel) GetProcessIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessId) {
		return nil, false
	}
	return o.ProcessId, true
}

// HasProcessId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceMerchantlistEnterApplyResponseModel) HasProcessId() bool {
	if o != nil && !IsNil(o.ProcessId) {
		return true
	}

	return false
}

// SetProcessId gets a reference to the given string and assigns it to the ProcessId field.
func (o *AlipayEbppInvoiceMerchantlistEnterApplyResponseModel) SetProcessId(v string) {
	o.ProcessId = &v
}

func (o AlipayEbppInvoiceMerchantlistEnterApplyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceMerchantlistEnterApplyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProcessId) {
		toSerialize["process_id"] = o.ProcessId
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceMerchantlistEnterApplyResponseModel struct {
	value *AlipayEbppInvoiceMerchantlistEnterApplyResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceMerchantlistEnterApplyResponseModel) Get() *AlipayEbppInvoiceMerchantlistEnterApplyResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceMerchantlistEnterApplyResponseModel) Set(val *AlipayEbppInvoiceMerchantlistEnterApplyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceMerchantlistEnterApplyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceMerchantlistEnterApplyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceMerchantlistEnterApplyResponseModel(val *AlipayEbppInvoiceMerchantlistEnterApplyResponseModel) *NullableAlipayEbppInvoiceMerchantlistEnterApplyResponseModel {
	return &NullableAlipayEbppInvoiceMerchantlistEnterApplyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceMerchantlistEnterApplyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceMerchantlistEnterApplyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


