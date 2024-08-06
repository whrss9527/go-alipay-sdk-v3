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

// checks if the AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel{}

// AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel struct for AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel
type AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel struct {
	// 删除是否成功
	Result *bool `json:"result,omitempty"`
}

// NewAlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel instantiates a new AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel() *AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel {
	this := AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModelWithDefaults instantiates a new AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModelWithDefaults() *AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel {
	this := AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel) GetResult() bool {
	if o == nil || IsNil(o.Result) {
		var ret bool
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel) GetResultOk() (*bool, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given bool and assigns it to the Result field.
func (o *AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel) SetResult(v bool) {
	o.Result = &v
}

func (o AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel struct {
	value *AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel) Get() *AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel) Set(val *AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel(val *AlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel) *NullableAlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel {
	return &NullableAlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceInstitutionExpenseruleDeleteResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


