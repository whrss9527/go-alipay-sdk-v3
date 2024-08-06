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

// checks if the AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel{}

// AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel struct for AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel
type AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel struct {
	// 费控聚合关系ID
	AggregationId *string `json:"aggregation_id,omitempty"`
}

// NewAlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel instantiates a new AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel() *AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel {
	this := AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel{}
	return &this
}

// NewAlipayEbppInvoiceExpensecontrolAggregationCreateResponseModelWithDefaults instantiates a new AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppInvoiceExpensecontrolAggregationCreateResponseModelWithDefaults() *AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel {
	this := AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel{}
	return &this
}

// GetAggregationId returns the AggregationId field value if set, zero value otherwise.
func (o *AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel) GetAggregationId() string {
	if o == nil || IsNil(o.AggregationId) {
		var ret string
		return ret
	}
	return *o.AggregationId
}

// GetAggregationIdOk returns a tuple with the AggregationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel) GetAggregationIdOk() (*string, bool) {
	if o == nil || IsNil(o.AggregationId) {
		return nil, false
	}
	return o.AggregationId, true
}

// HasAggregationId returns a boolean if a field has been set.
func (o *AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel) HasAggregationId() bool {
	if o != nil && !IsNil(o.AggregationId) {
		return true
	}

	return false
}

// SetAggregationId gets a reference to the given string and assigns it to the AggregationId field.
func (o *AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel) SetAggregationId(v string) {
	o.AggregationId = &v
}

func (o AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AggregationId) {
		toSerialize["aggregation_id"] = o.AggregationId
	}
	return toSerialize, nil
}

type NullableAlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel struct {
	value *AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel
	isSet bool
}

func (v NullableAlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel) Get() *AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel {
	return v.value
}

func (v *NullableAlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel) Set(val *AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel(val *AlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel) *NullableAlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel {
	return &NullableAlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppInvoiceExpensecontrolAggregationCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


