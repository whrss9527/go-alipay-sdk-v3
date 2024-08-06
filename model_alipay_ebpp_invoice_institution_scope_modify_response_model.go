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

// checks if the AlipayEbppInvoiceInstitutionScopeModifyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceInstitutionScopeModifyResponseModel{}

// AlipayEbppInvoiceInstitutionScopeModifyResponseModel struct for AlipayEbppInvoiceInstitutionScopeModifyResponseModel
type AlipayEbppInvoiceInstitutionScopeModifyResponseModel struct {
	// 操作是否成功
	Result *bool `json:"result,omitempty"`
}

// NewAlipayEbppInvoiceInstitutionScopeModifyResponseModel instantiates a new AlipayEbppInvoiceInstitutionScopeModifyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceInstitutionScopeModifyResponseModel() *AlipayEbppInvoiceInstitutionScopeModifyResponseModel {
	this := AlipayEbppInvoiceInstitutionScopeModifyResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceInstitutionScopeModifyResponseModelWithDefaults instantiates a new AlipayEbppInvoiceInstitutionScopeModifyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceInstitutionScopeModifyResponseModelWithDefaults() *AlipayEbppInvoiceInstitutionScopeModifyResponseModel {
	this := AlipayEbppInvoiceInstitutionScopeModifyResponseModel{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInstitutionScopeModifyResponseModel) GetResult() bool {
	if o == nil || IsNil(o.Result) {
		var ret bool
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyResponseModel) GetResultOk() (*bool, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInstitutionScopeModifyResponseModel) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given bool and assigns it to the Result field.
func (o *AlipayEbppInvoiceInstitutionScopeModifyResponseModel) SetResult(v bool) {
	o.Result = &v
}

func (o AlipayEbppInvoiceInstitutionScopeModifyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceInstitutionScopeModifyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceInstitutionScopeModifyResponseModel struct {
	value *AlipayEbppInvoiceInstitutionScopeModifyResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceInstitutionScopeModifyResponseModel) Get() *AlipayEbppInvoiceInstitutionScopeModifyResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceInstitutionScopeModifyResponseModel) Set(val *AlipayEbppInvoiceInstitutionScopeModifyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceInstitutionScopeModifyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceInstitutionScopeModifyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceInstitutionScopeModifyResponseModel(val *AlipayEbppInvoiceInstitutionScopeModifyResponseModel) *NullableAlipayEbppInvoiceInstitutionScopeModifyResponseModel {
	return &NullableAlipayEbppInvoiceInstitutionScopeModifyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceInstitutionScopeModifyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceInstitutionScopeModifyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
