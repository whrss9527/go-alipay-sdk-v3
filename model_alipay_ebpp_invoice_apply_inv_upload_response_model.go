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

// checks if the AlipayEbppInvoiceApplyInvUploadResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceApplyInvUploadResponseModel{}

// AlipayEbppInvoiceApplyInvUploadResponseModel struct for AlipayEbppInvoiceApplyInvUploadResponseModel
type AlipayEbppInvoiceApplyInvUploadResponseModel struct {
	// 处理是否成功
	IsSuccess *bool `json:"is_success,omitempty"`
}

// NewAlipayEbppInvoiceApplyInvUploadResponseModel instantiates a new AlipayEbppInvoiceApplyInvUploadResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceApplyInvUploadResponseModel() *AlipayEbppInvoiceApplyInvUploadResponseModel {
	this := AlipayEbppInvoiceApplyInvUploadResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceApplyInvUploadResponseModelWithDefaults instantiates a new AlipayEbppInvoiceApplyInvUploadResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceApplyInvUploadResponseModelWithDefaults() *AlipayEbppInvoiceApplyInvUploadResponseModel {
	this := AlipayEbppInvoiceApplyInvUploadResponseModel{}
	return &this
}

// GetIsSuccess returns the IsSuccess field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceApplyInvUploadResponseModel) GetIsSuccess() bool {
	if o == nil || IsNil(o.IsSuccess) {
		var ret bool
		return ret
	}
	return *o.IsSuccess
}

// GetIsSuccessOk returns a tuple with the IsSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceApplyInvUploadResponseModel) GetIsSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSuccess) {
		return nil, false
	}
	return o.IsSuccess, true
}

// HasIsSuccess returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceApplyInvUploadResponseModel) HasIsSuccess() bool {
	if o != nil && !IsNil(o.IsSuccess) {
		return true
	}

	return false
}

// SetIsSuccess gets a reference to the given bool and assigns it to the IsSuccess field.
func (o *AlipayEbppInvoiceApplyInvUploadResponseModel) SetIsSuccess(v bool) {
	o.IsSuccess = &v
}

func (o AlipayEbppInvoiceApplyInvUploadResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceApplyInvUploadResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsSuccess) {
		toSerialize["is_success"] = o.IsSuccess
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceApplyInvUploadResponseModel struct {
	value *AlipayEbppInvoiceApplyInvUploadResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceApplyInvUploadResponseModel) Get() *AlipayEbppInvoiceApplyInvUploadResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceApplyInvUploadResponseModel) Set(val *AlipayEbppInvoiceApplyInvUploadResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceApplyInvUploadResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceApplyInvUploadResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceApplyInvUploadResponseModel(val *AlipayEbppInvoiceApplyInvUploadResponseModel) *NullableAlipayEbppInvoiceApplyInvUploadResponseModel {
	return &NullableAlipayEbppInvoiceApplyInvUploadResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceApplyInvUploadResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceApplyInvUploadResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


