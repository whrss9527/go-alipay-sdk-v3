/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel{}

// AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel struct for AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel
type AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel struct {
	// 错误码
	Code string `json:"code"`
	// 解决方案链接
	Links *string `json:"links,omitempty"`
	// 错误描述
	Message string `json:"message"`
}

type _AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel

// NewAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel instantiates a new AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel(code string, message string) *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel {
	this := AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel{}
	this.Code = code
	this.Message = message
	return &this
}

// NewAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModelWithDefaults instantiates a new AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModelWithDefaults() *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel {
	this := AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel{}
	return &this
}

// GetCode returns the Code field value
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) SetCode(v string) {
	o.Code = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) GetLinks() string {
	if o == nil || IsNil(o.Links) {
		var ret string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) GetLinksOk() (*string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given string and assigns it to the Links field.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) SetLinks(v string) {
	o.Links = &v
}

// GetMessage returns the Message field value
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) SetMessage(v string) {
	o.Message = v
}

func (o AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) UnmarshalJSON(data []byte) (err error) {
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
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel := _AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel)

	if err != nil {
		return err
	}

	*o = AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel(varAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel)

	return err
}

type NullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel struct {
	value *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) Get() *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) Set(val *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel(val *AlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) *NullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel {
	return &NullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeConsumeBatchqueryErrorResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
