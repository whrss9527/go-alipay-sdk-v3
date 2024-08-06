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

// checks if the AlipayEbppInvoiceSycnResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceSycnResponseModel{}

// AlipayEbppInvoiceSycnResponseModel struct for AlipayEbppInvoiceSycnResponseModel
type AlipayEbppInvoiceSycnResponseModel struct {
	// 支付宝发票管家发票详情页链接地址  如果同步的发票是多张时，此链接为当前用户的已开立发票列表地址
	Url *string `json:"url,omitempty"`
}

// NewAlipayEbppInvoiceSycnResponseModel instantiates a new AlipayEbppInvoiceSycnResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceSycnResponseModel() *AlipayEbppInvoiceSycnResponseModel {
	this := AlipayEbppInvoiceSycnResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceSycnResponseModelWithDefaults instantiates a new AlipayEbppInvoiceSycnResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceSycnResponseModelWithDefaults() *AlipayEbppInvoiceSycnResponseModel {
	this := AlipayEbppInvoiceSycnResponseModel{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceSycnResponseModel) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceSycnResponseModel) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceSycnResponseModel) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AlipayEbppInvoiceSycnResponseModel) SetUrl(v string) {
	o.Url = &v
}

func (o AlipayEbppInvoiceSycnResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceSycnResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceSycnResponseModel struct {
	value *AlipayEbppInvoiceSycnResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceSycnResponseModel) Get() *AlipayEbppInvoiceSycnResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceSycnResponseModel) Set(val *AlipayEbppInvoiceSycnResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceSycnResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceSycnResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceSycnResponseModel(val *AlipayEbppInvoiceSycnResponseModel) *NullableAlipayEbppInvoiceSycnResponseModel {
	return &NullableAlipayEbppInvoiceSycnResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceSycnResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceSycnResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
