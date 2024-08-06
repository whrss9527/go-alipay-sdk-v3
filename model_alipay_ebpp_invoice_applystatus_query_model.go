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

// checks if the AlipayEbppInvoiceApplystatusQueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceApplystatusQueryModel{}

// AlipayEbppInvoiceApplystatusQueryModel struct for AlipayEbppInvoiceApplystatusQueryModel
type AlipayEbppInvoiceApplystatusQueryModel struct {
	// 定义商户的一级简称,用于标识商户品牌，对应于商户入驻时填写的\"商户品牌简称\"。 如：肯德基：KFC
	MShortName *string `json:"m_short_name,omitempty"`
	// 待查询订单开票状态列表，各订单号间通过英文逗号分割，不限于支付宝体内交易订单号。如：20200520110046966071,20200520110046966072,20200520110046966073
	OrderNoList []string `json:"order_no_list,omitempty"`
	// 定义商户的二级简称,用于标识商户品牌下的分支机构，如门店，对应于商户入驻时填写的\"商户门店简称\"。 如：肯德基-杭州西湖区文一西路店：KFC-HZ-19003 要求：\"商户品牌简称+商户门店简称\"作为确定商户及其下属机构的唯一标识，不可重复。
	SubMShortName *string `json:"sub_m_short_name,omitempty"`
}

// NewAlipayEbppInvoiceApplystatusQueryModel instantiates a new AlipayEbppInvoiceApplystatusQueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceApplystatusQueryModel() *AlipayEbppInvoiceApplystatusQueryModel {
	this := AlipayEbppInvoiceApplystatusQueryModel{}
	return &this
}

// NewAlipayEbppInvoiceApplystatusQueryModelWithDefaults instantiates a new AlipayEbppInvoiceApplystatusQueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceApplystatusQueryModelWithDefaults() *AlipayEbppInvoiceApplystatusQueryModel {
	this := AlipayEbppInvoiceApplystatusQueryModel{}
	return &this
}

// GetMShortName returns the MShortName field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceApplystatusQueryModel) GetMShortName() string {
	if o == nil || IsNil(o.MShortName) {
		var ret string
		return ret
	}
	return *o.MShortName
}

// GetMShortNameOk returns a tuple with the MShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceApplystatusQueryModel) GetMShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.MShortName) {
		return nil, false
	}
	return o.MShortName, true
}

// HasMShortName returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceApplystatusQueryModel) HasMShortName() bool {
	if o != nil && !IsNil(o.MShortName) {
		return true
	}

	return false
}

// SetMShortName gets a reference to the given string and assigns it to the MShortName field.
func (o *AlipayEbppInvoiceApplystatusQueryModel) SetMShortName(v string) {
	o.MShortName = &v
}

// GetOrderNoList returns the OrderNoList field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceApplystatusQueryModel) GetOrderNoList() []string {
	if o == nil || IsNil(o.OrderNoList) {
		var ret []string
		return ret
	}
	return o.OrderNoList
}

// GetOrderNoListOk returns a tuple with the OrderNoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceApplystatusQueryModel) GetOrderNoListOk() ([]string, bool) {
	if o == nil || IsNil(o.OrderNoList) {
		return nil, false
	}
	return o.OrderNoList, true
}

// HasOrderNoList returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceApplystatusQueryModel) HasOrderNoList() bool {
	if o != nil && !IsNil(o.OrderNoList) {
		return true
	}

	return false
}

// SetOrderNoList gets a reference to the given []string and assigns it to the OrderNoList field.
func (o *AlipayEbppInvoiceApplystatusQueryModel) SetOrderNoList(v []string) {
	o.OrderNoList = v
}

// GetSubMShortName returns the SubMShortName field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceApplystatusQueryModel) GetSubMShortName() string {
	if o == nil || IsNil(o.SubMShortName) {
		var ret string
		return ret
	}
	return *o.SubMShortName
}

// GetSubMShortNameOk returns a tuple with the SubMShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceApplystatusQueryModel) GetSubMShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubMShortName) {
		return nil, false
	}
	return o.SubMShortName, true
}

// HasSubMShortName returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceApplystatusQueryModel) HasSubMShortName() bool {
	if o != nil && !IsNil(o.SubMShortName) {
		return true
	}

	return false
}

// SetSubMShortName gets a reference to the given string and assigns it to the SubMShortName field.
func (o *AlipayEbppInvoiceApplystatusQueryModel) SetSubMShortName(v string) {
	o.SubMShortName = &v
}

func (o AlipayEbppInvoiceApplystatusQueryModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceApplystatusQueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MShortName) {
		toSerialize["m_short_name"] = o.MShortName
	}
	if !IsNil(o.OrderNoList) {
		toSerialize["order_no_list"] = o.OrderNoList
	}
	if !IsNil(o.SubMShortName) {
		toSerialize["sub_m_short_name"] = o.SubMShortName
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceApplystatusQueryModel struct {
	value *AlipayEbppInvoiceApplystatusQueryModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceApplystatusQueryModel) Get() *AlipayEbppInvoiceApplystatusQueryModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceApplystatusQueryModel) Set(val *AlipayEbppInvoiceApplystatusQueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceApplystatusQueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceApplystatusQueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceApplystatusQueryModel(val *AlipayEbppInvoiceApplystatusQueryModel) *NullableAlipayEbppInvoiceApplystatusQueryModel {
	return &NullableAlipayEbppInvoiceApplystatusQueryModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceApplystatusQueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceApplystatusQueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
