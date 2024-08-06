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

// checks if the AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel{}

// AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel struct for AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel
type AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel struct {
	// success
	Success *bool `json:"success,omitempty"`
}

// NewAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel instantiates a new AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel() *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel {
	this := AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModelWithDefaults instantiates a new AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModelWithDefaults() *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel {
	this := AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel) SetSuccess(v bool) {
	o.Success = &v
}

func (o AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel struct {
	value *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel) Get() *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel) Set(val *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel(val *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel) *NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel {
	return &NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleModifyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

