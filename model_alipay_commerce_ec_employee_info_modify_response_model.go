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

// checks if the AlipayCommerceEcEmployeeInfoModifyResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceEcEmployeeInfoModifyResponseModel{}

// AlipayCommerceEcEmployeeInfoModifyResponseModel struct for AlipayCommerceEcEmployeeInfoModifyResponseModel
type AlipayCommerceEcEmployeeInfoModifyResponseModel struct {
	// 员工id
	EmployeeId *string `json:"employee_id,omitempty"`
}

// NewAlipayCommerceEcEmployeeInfoModifyResponseModel instantiates a new AlipayCommerceEcEmployeeInfoModifyResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceEcEmployeeInfoModifyResponseModel() *AlipayCommerceEcEmployeeInfoModifyResponseModel {
	this := AlipayCommerceEcEmployeeInfoModifyResponseModel{}
	return &this
}

// NewAlipayCommerceEcEmployeeInfoModifyResponseModelWithDefaults instantiates a new AlipayCommerceEcEmployeeInfoModifyResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceEcEmployeeInfoModifyResponseModelWithDefaults() *AlipayCommerceEcEmployeeInfoModifyResponseModel {
	this := AlipayCommerceEcEmployeeInfoModifyResponseModel{}
	return &this
}

// GetEmployeeId returns the EmployeeId field value if set, zero value otherwise.
func (o *AlipayCommerceEcEmployeeInfoModifyResponseModel) GetEmployeeId() string {
	if o == nil || IsNil(o.EmployeeId) {
		var ret string
		return ret
	}
	return *o.EmployeeId
}

// GetEmployeeIdOk returns a tuple with the EmployeeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceEcEmployeeInfoModifyResponseModel) GetEmployeeIdOk() (*string, bool) {
	if o == nil || IsNil(o.EmployeeId) {
		return nil, false
	}
	return o.EmployeeId, true
}

// HasEmployeeId returns a boolean if a field has been set.
func (o *AlipayCommerceEcEmployeeInfoModifyResponseModel) HasEmployeeId() bool {
	if o != nil && !IsNil(o.EmployeeId) {
		return true
	}

	return false
}

// SetEmployeeId gets a reference to the given string and assigns it to the EmployeeId field.
func (o *AlipayCommerceEcEmployeeInfoModifyResponseModel) SetEmployeeId(v string) {
	o.EmployeeId = &v
}

func (o AlipayCommerceEcEmployeeInfoModifyResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceEcEmployeeInfoModifyResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmployeeId) {
		toSerialize["employee_id"] = o.EmployeeId
	}
	return toSerialize, nil
}

type NullableAlipayCommerceEcEmployeeInfoModifyResponseModel struct {
	value *AlipayCommerceEcEmployeeInfoModifyResponseModel
	isSet bool
}

func (v NullableAlipayCommerceEcEmployeeInfoModifyResponseModel) Get() *AlipayCommerceEcEmployeeInfoModifyResponseModel {
	return v.value
}

func (v *NullableAlipayCommerceEcEmployeeInfoModifyResponseModel) Set(val *AlipayCommerceEcEmployeeInfoModifyResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEcEmployeeInfoModifyResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEcEmployeeInfoModifyResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEcEmployeeInfoModifyResponseModel(val *AlipayCommerceEcEmployeeInfoModifyResponseModel) *NullableAlipayCommerceEcEmployeeInfoModifyResponseModel {
	return &NullableAlipayCommerceEcEmployeeInfoModifyResponseModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceEcEmployeeInfoModifyResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEcEmployeeInfoModifyResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


