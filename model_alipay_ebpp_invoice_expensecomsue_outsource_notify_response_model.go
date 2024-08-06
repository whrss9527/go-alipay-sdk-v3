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

// checks if the AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel{}

// AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel struct for AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel
type AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel struct {
	// 是否成功
	Success *bool `json:"success,omitempty"`
}

// NewAlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel instantiates a new AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel() *AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel {
	this := AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModelWithDefaults instantiates a new AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModelWithDefaults() *AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel {
	this := AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel) SetSuccess(v bool) {
	o.Success = &v
}

func (o AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel struct {
	value *AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel) Get() *AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel) Set(val *AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel(val *AlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel) *NullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel {
	return &NullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceExpensecomsueOutsourceNotifyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


