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

// checks if the AlipayEbppInvoiceSycnModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceSycnModel{}

// AlipayEbppInvoiceSycnModel struct for AlipayEbppInvoiceSycnModel
type AlipayEbppInvoiceSycnModel struct {
	// 同步发票信息模型
	InvoiceInfo []InvoiceModelContent `json:"invoice_info,omitempty"`
	// 商户的品牌名称简称,该字段需要接入前向发票管家申请，  m_short_name+sub_m_short_name具有唯一约束  如：肯德基：KFC
	MShortName *string `json:"m_short_name,omitempty"`
	// 支付宝为商户分配的商户门店简称，该字段需要接入前在发票管家申请  如：肯德基-杭州西湖区文一西路店：KFC-HZ-XH001
	SubMShortName *string `json:"sub_m_short_name,omitempty"`
}

// NewAlipayEbppInvoiceSycnModel instantiates a new AlipayEbppInvoiceSycnModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceSycnModel() *AlipayEbppInvoiceSycnModel {
	this := AlipayEbppInvoiceSycnModel{}
	return &this
}

// NewAlipayEbppInvoiceSycnModelWithDefaults instantiates a new AlipayEbppInvoiceSycnModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceSycnModelWithDefaults() *AlipayEbppInvoiceSycnModel {
	this := AlipayEbppInvoiceSycnModel{}
	return &this
}

// GetInvoiceInfo returns the InvoiceInfo field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceSycnModel) GetInvoiceInfo() []InvoiceModelContent {
	if o == nil || IsNil(o.InvoiceInfo) {
		var ret []InvoiceModelContent
		return ret
	}
	return o.InvoiceInfo
}

// GetInvoiceInfoOk returns a tuple with the InvoiceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceSycnModel) GetInvoiceInfoOk() ([]InvoiceModelContent, bool) {
	if o == nil || IsNil(o.InvoiceInfo) {
		return nil, false
	}
	return o.InvoiceInfo, true
}

// HasInvoiceInfo returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceSycnModel) HasInvoiceInfo() bool {
	if o != nil && !IsNil(o.InvoiceInfo) {
		return true
	}

	return false
}

// SetInvoiceInfo gets a reference to the given []InvoiceModelContent and assigns it to the InvoiceInfo field.
func (o *AlipayEbppInvoiceSycnModel) SetInvoiceInfo(v []InvoiceModelContent) {
	o.InvoiceInfo = v
}

// GetMShortName returns the MShortName field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceSycnModel) GetMShortName() string {
	if o == nil || IsNil(o.MShortName) {
		var ret string
		return ret
	}
	return *o.MShortName
}

// GetMShortNameOk returns a tuple with the MShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceSycnModel) GetMShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.MShortName) {
		return nil, false
	}
	return o.MShortName, true
}

// HasMShortName returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceSycnModel) HasMShortName() bool {
	if o != nil && !IsNil(o.MShortName) {
		return true
	}

	return false
}

// SetMShortName gets a reference to the given string and assigns it to the MShortName field.
func (o *AlipayEbppInvoiceSycnModel) SetMShortName(v string) {
	o.MShortName = &v
}

// GetSubMShortName returns the SubMShortName field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceSycnModel) GetSubMShortName() string {
	if o == nil || IsNil(o.SubMShortName) {
		var ret string
		return ret
	}
	return *o.SubMShortName
}

// GetSubMShortNameOk returns a tuple with the SubMShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceSycnModel) GetSubMShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubMShortName) {
		return nil, false
	}
	return o.SubMShortName, true
}

// HasSubMShortName returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceSycnModel) HasSubMShortName() bool {
	if o != nil && !IsNil(o.SubMShortName) {
		return true
	}

	return false
}

// SetSubMShortName gets a reference to the given string and assigns it to the SubMShortName field.
func (o *AlipayEbppInvoiceSycnModel) SetSubMShortName(v string) {
	o.SubMShortName = &v
}

func (o AlipayEbppInvoiceSycnModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceSycnModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoiceInfo) {
		toSerialize["invoice_info"] = o.InvoiceInfo
	}
	if !IsNil(o.MShortName) {
		toSerialize["m_short_name"] = o.MShortName
	}
	if !IsNil(o.SubMShortName) {
		toSerialize["sub_m_short_name"] = o.SubMShortName
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceSycnModel struct {
	value *AlipayEbppInvoiceSycnModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceSycnModel) Get() *AlipayEbppInvoiceSycnModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceSycnModel) Set(val *AlipayEbppInvoiceSycnModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceSycnModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceSycnModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceSycnModel(val *AlipayEbppInvoiceSycnModel) *NullableAlipayEbppInvoiceSycnModel {
	return &NullableAlipayEbppInvoiceSycnModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceSycnModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceSycnModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
