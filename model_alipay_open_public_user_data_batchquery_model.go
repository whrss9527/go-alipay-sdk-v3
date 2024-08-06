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

// checks if the AlipayOpenPublicUserDataBatchqueryModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayOpenPublicUserDataBatchqueryModel{}

// AlipayOpenPublicUserDataBatchqueryModel struct for AlipayOpenPublicUserDataBatchqueryModel
type AlipayOpenPublicUserDataBatchqueryModel struct {
	// 数据开始日期，时间格式为 \"yyyyMMdd\" 。 
	BeginDate *string `json:"begin_date,omitempty"`
	// 数据结束日期，时间格式为\"yyyyMMdd\"。查询数据开始日期/结束日期时间跨度最大30天。
	EndDate *string `json:"end_date,omitempty"`
}

// NewAlipayOpenPublicUserDataBatchqueryModel instantiates a new AlipayOpenPublicUserDataBatchqueryModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayOpenPublicUserDataBatchqueryModel() *AlipayOpenPublicUserDataBatchqueryModel {
	this := AlipayOpenPublicUserDataBatchqueryModel{}
	return &this
}

// NewAlipayOpenPublicUserDataBatchqueryModelWithDefaults instantiates a new AlipayOpenPublicUserDataBatchqueryModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayOpenPublicUserDataBatchqueryModelWithDefaults() *AlipayOpenPublicUserDataBatchqueryModel {
	this := AlipayOpenPublicUserDataBatchqueryModel{}
	return &this
}

// GetBeginDate returns the BeginDate field value if set, zero value otherwise.
func (o *AlipayOpenPublicUserDataBatchqueryModel) GetBeginDate() string {
	if o == nil || IsNil(o.BeginDate) {
		var ret string
		return ret
	}
	return *o.BeginDate
}

// GetBeginDateOk returns a tuple with the BeginDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicUserDataBatchqueryModel) GetBeginDateOk() (*string, bool) {
	if o == nil || IsNil(o.BeginDate) {
		return nil, false
	}
	return o.BeginDate, true
}

// HasBeginDate returns a boolean if a field has been set.
func (o *AlipayOpenPublicUserDataBatchqueryModel) HasBeginDate() bool {
	if o != nil && !IsNil(o.BeginDate) {
		return true
	}

	return false
}

// SetBeginDate gets a reference to the given string and assigns it to the BeginDate field.
func (o *AlipayOpenPublicUserDataBatchqueryModel) SetBeginDate(v string) {
	o.BeginDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *AlipayOpenPublicUserDataBatchqueryModel) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayOpenPublicUserDataBatchqueryModel) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *AlipayOpenPublicUserDataBatchqueryModel) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *AlipayOpenPublicUserDataBatchqueryModel) SetEndDate(v string) {
	o.EndDate = &v
}

func (o AlipayOpenPublicUserDataBatchqueryModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayOpenPublicUserDataBatchqueryModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BeginDate) {
		toSerialize["begin_date"] = o.BeginDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableAlipayOpenPublicUserDataBatchqueryModel struct {
	value *AlipayOpenPublicUserDataBatchqueryModel
	isSet bool
}

func (v NullableAlipayOpenPublicUserDataBatchqueryModel) Get() *AlipayOpenPublicUserDataBatchqueryModel {
	return v.value
}

func (v *NullableAlipayOpenPublicUserDataBatchqueryModel) Set(val *AlipayOpenPublicUserDataBatchqueryModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicUserDataBatchqueryModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicUserDataBatchqueryModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicUserDataBatchqueryModel(val *AlipayOpenPublicUserDataBatchqueryModel) *NullableAlipayOpenPublicUserDataBatchqueryModel {
	return &NullableAlipayOpenPublicUserDataBatchqueryModel{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicUserDataBatchqueryModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicUserDataBatchqueryModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


