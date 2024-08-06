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

// checks if the AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel{}

// AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel struct for AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel
type AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel struct {
	// 开票规则ID
	InvoiceRuleId *string `json:"invoice_rule_id,omitempty"`
}

// NewAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel instantiates a new AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel() *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel {
	this := AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModelWithDefaults instantiates a new AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModelWithDefaults() *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel {
	this := AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel{}
	return &this
}

// GetInvoiceRuleId returns the InvoiceRuleId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel) GetInvoiceRuleId() string {
	if o == nil || IsNil(o.InvoiceRuleId) {
		var ret string
		return ret
	}
	return *o.InvoiceRuleId
}

// GetInvoiceRuleIdOk returns a tuple with the InvoiceRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel) GetInvoiceRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceRuleId) {
		return nil, false
	}
	return o.InvoiceRuleId, true
}

// HasInvoiceRuleId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel) HasInvoiceRuleId() bool {
	if o != nil && !IsNil(o.InvoiceRuleId) {
		return true
	}

	return false
}

// SetInvoiceRuleId gets a reference to the given string and assigns it to the InvoiceRuleId field.
func (o *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel) SetInvoiceRuleId(v string) {
	o.InvoiceRuleId = &v
}

func (o AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoiceRuleId) {
		toSerialize["invoice_rule_id"] = o.InvoiceRuleId
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel struct {
	value *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel) Get() *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel) Set(val *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel(val *AlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel) *NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel {
	return &NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceEnterpriseconsumeEnterpriseopenruleCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
