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

// checks if the AlipayEbppInvoiceIssueruleCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceIssueruleCreateResponseModel{}

// AlipayEbppInvoiceIssueruleCreateResponseModel struct for AlipayEbppInvoiceIssueruleCreateResponseModel
type AlipayEbppInvoiceIssueruleCreateResponseModel struct {
	// 创建的发放规则ID
	IssueRuleId *string `json:"issue_rule_id,omitempty"`
}

// NewAlipayEbppInvoiceIssueruleCreateResponseModel instantiates a new AlipayEbppInvoiceIssueruleCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceIssueruleCreateResponseModel() *AlipayEbppInvoiceIssueruleCreateResponseModel {
	this := AlipayEbppInvoiceIssueruleCreateResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceIssueruleCreateResponseModelWithDefaults instantiates a new AlipayEbppInvoiceIssueruleCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceIssueruleCreateResponseModelWithDefaults() *AlipayEbppInvoiceIssueruleCreateResponseModel {
	this := AlipayEbppInvoiceIssueruleCreateResponseModel{}
	return &this
}

// GetIssueRuleId returns the IssueRuleId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceIssueruleCreateResponseModel) GetIssueRuleId() string {
	if o == nil || IsNil(o.IssueRuleId) {
		var ret string
		return ret
	}
	return *o.IssueRuleId
}

// GetIssueRuleIdOk returns a tuple with the IssueRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceIssueruleCreateResponseModel) GetIssueRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.IssueRuleId) {
		return nil, false
	}
	return o.IssueRuleId, true
}

// HasIssueRuleId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceIssueruleCreateResponseModel) HasIssueRuleId() bool {
	if o != nil && !IsNil(o.IssueRuleId) {
		return true
	}

	return false
}

// SetIssueRuleId gets a reference to the given string and assigns it to the IssueRuleId field.
func (o *AlipayEbppInvoiceIssueruleCreateResponseModel) SetIssueRuleId(v string) {
	o.IssueRuleId = &v
}

func (o AlipayEbppInvoiceIssueruleCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceIssueruleCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IssueRuleId) {
		toSerialize["issue_rule_id"] = o.IssueRuleId
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceIssueruleCreateResponseModel struct {
	value *AlipayEbppInvoiceIssueruleCreateResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceIssueruleCreateResponseModel) Get() *AlipayEbppInvoiceIssueruleCreateResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceIssueruleCreateResponseModel) Set(val *AlipayEbppInvoiceIssueruleCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceIssueruleCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceIssueruleCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceIssueruleCreateResponseModel(val *AlipayEbppInvoiceIssueruleCreateResponseModel) *NullableAlipayEbppInvoiceIssueruleCreateResponseModel {
	return &NullableAlipayEbppInvoiceIssueruleCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceIssueruleCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceIssueruleCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


