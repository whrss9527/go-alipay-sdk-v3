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

// checks if the AlipayEbppInvoiceInfoSendResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceInfoSendResponseModel{}

// AlipayEbppInvoiceInfoSendResponseModel struct for AlipayEbppInvoiceInfoSendResponseModel
type AlipayEbppInvoiceInfoSendResponseModel struct {
	// 发票访问地址，同步红票的情况下不返回该字段，同步蓝票的情况下如果同步单张发票，返回发票详情链接地址，如果同步的是多张发票，返回发票列表地址
	Url *string `json:"url,omitempty"`
}

// NewAlipayEbppInvoiceInfoSendResponseModel instantiates a new AlipayEbppInvoiceInfoSendResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceInfoSendResponseModel() *AlipayEbppInvoiceInfoSendResponseModel {
	this := AlipayEbppInvoiceInfoSendResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceInfoSendResponseModelWithDefaults instantiates a new AlipayEbppInvoiceInfoSendResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceInfoSendResponseModelWithDefaults() *AlipayEbppInvoiceInfoSendResponseModel {
	this := AlipayEbppInvoiceInfoSendResponseModel{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceInfoSendResponseModel) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceInfoSendResponseModel) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceInfoSendResponseModel) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AlipayEbppInvoiceInfoSendResponseModel) SetUrl(v string) {
	o.Url = &v
}

func (o AlipayEbppInvoiceInfoSendResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceInfoSendResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceInfoSendResponseModel struct {
	value *AlipayEbppInvoiceInfoSendResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceInfoSendResponseModel) Get() *AlipayEbppInvoiceInfoSendResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceInfoSendResponseModel) Set(val *AlipayEbppInvoiceInfoSendResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceInfoSendResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceInfoSendResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceInfoSendResponseModel(val *AlipayEbppInvoiceInfoSendResponseModel) *NullableAlipayEbppInvoiceInfoSendResponseModel {
	return &NullableAlipayEbppInvoiceInfoSendResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceInfoSendResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceInfoSendResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
