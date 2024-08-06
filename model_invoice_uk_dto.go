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

// checks if the InvoiceUkDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceUkDTO{}

// InvoiceUkDTO struct for InvoiceUkDTO
type InvoiceUkDTO struct {
	// 发票代码
	InvoiceCode *string `json:"invoice_code,omitempty"`
	// 发票号码
	InvoiceNo *string `json:"invoice_no,omitempty"`
}

// NewInvoiceUkDTO instantiates a new InvoiceUkDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceUkDTO() *InvoiceUkDTO {
	this := InvoiceUkDTO{}
	return &this
}

// NewInvoiceUkDTOWithDefaults instantiates a new InvoiceUkDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceUkDTOWithDefaults() *InvoiceUkDTO {
	this := InvoiceUkDTO{}
	return &this
}

// GetInvoiceCode returns the InvoiceCode field value if set, zero value otherwise.
func (o *InvoiceUkDTO) GetInvoiceCode() string {
	if o == nil || IsNil(o.InvoiceCode) {
		var ret string
		return ret
	}
	return *o.InvoiceCode
}

// GetInvoiceCodeOk returns a tuple with the InvoiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceUkDTO) GetInvoiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceCode) {
		return nil, false
	}
	return o.InvoiceCode, true
}

// HasInvoiceCode returns a boolean if a field has been set.
func (o *InvoiceUkDTO) HasInvoiceCode() bool {
	if o != nil && !IsNil(o.InvoiceCode) {
		return true
	}

	return false
}

// SetInvoiceCode gets a reference to the given string and assigns it to the InvoiceCode field.
func (o *InvoiceUkDTO) SetInvoiceCode(v string) {
	o.InvoiceCode = &v
}

// GetInvoiceNo returns the InvoiceNo field value if set, zero value otherwise.
func (o *InvoiceUkDTO) GetInvoiceNo() string {
	if o == nil || IsNil(o.InvoiceNo) {
		var ret string
		return ret
	}
	return *o.InvoiceNo
}

// GetInvoiceNoOk returns a tuple with the InvoiceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceUkDTO) GetInvoiceNoOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceNo) {
		return nil, false
	}
	return o.InvoiceNo, true
}

// HasInvoiceNo returns a boolean if a field has been set.
func (o *InvoiceUkDTO) HasInvoiceNo() bool {
	if o != nil && !IsNil(o.InvoiceNo) {
		return true
	}

	return false
}

// SetInvoiceNo gets a reference to the given string and assigns it to the InvoiceNo field.
func (o *InvoiceUkDTO) SetInvoiceNo(v string) {
	o.InvoiceNo = &v
}

func (o InvoiceUkDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceUkDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoiceCode) {
		toSerialize["invoice_code"] = o.InvoiceCode
	}
	if !IsNil(o.InvoiceNo) {
		toSerialize["invoice_no"] = o.InvoiceNo
	}
	return toSerialize, nil
}

type NullableInvoiceUkDTO struct {
	value *InvoiceUkDTO
	isSet bool
}

func (v NullableInvoiceUkDTO) Get() *InvoiceUkDTO {
	return v.value
}

func (v *NullableInvoiceUkDTO) Set(val *InvoiceUkDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceUkDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceUkDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceUkDTO(val *InvoiceUkDTO) *NullableInvoiceUkDTO {
	return &NullableInvoiceUkDTO{value: val, isSet: true}
}

func (v NullableInvoiceUkDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceUkDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
