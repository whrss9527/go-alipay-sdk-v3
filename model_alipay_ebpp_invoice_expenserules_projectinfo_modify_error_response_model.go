/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel{}

// AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel struct for AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel
type AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel struct {
	// 错误码
	Code string `json:"code"`
	// 解决方案链接
	Links *string `json:"links,omitempty"`
	// 错误描述
	Message string `json:"message"`
}

type _AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel

// NewAlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel instantiates a new AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel(code string, message string) *AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel {
	this := AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel{}
	this.Code = code
	this.Message = message
	return &this
}

// NewAlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModelWithDefaults instantiates a new AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModelWithDefaults() *AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel {
	this := AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel{}
	return &this
}

// GetCode returns the Code field value
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) SetCode(v string) {
	o.Code = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) GetLinks() string {
	if o == nil || IsNil(o.Links) {
		var ret string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) GetLinksOk() (*string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given string and assigns it to the Links field.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) SetLinks(v string) {
	o.Links = &v
}

// GetMessage returns the Message field value
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) SetMessage(v string) {
	o.Message = v
}

func (o AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel := _AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel)

	if err != nil {
		return err
	}

	*o = AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel(varAlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel)

	return err
}

type NullableAlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel struct {
	value *AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) Get() *AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) Set(val *AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel(val *AlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) *NullableAlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel {
	return &NullableAlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceExpenserulesProjectinfoModifyErrorResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


