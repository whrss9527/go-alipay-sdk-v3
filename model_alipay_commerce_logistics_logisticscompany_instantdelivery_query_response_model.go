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

// checks if the AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel{}

// AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel struct for AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel
type AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel struct {
	// 即时配送公司列表
	LogisticsCompanies []LogisticsCompanyResult `json:"logistics_companies,omitempty"`
}

// NewAlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel instantiates a new AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel() *AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel {
	this := AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel{}
	return &this
}

// NewAlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModelWithDefaults instantiates a new AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModelWithDefaults() *AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel {
	this := AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel{}
	return &this
}

// GetLogisticsCompanies returns the LogisticsCompanies field value if set, zero value otherwise.
func (o *AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel) GetLogisticsCompanies() []LogisticsCompanyResult {
	if o == nil || IsNil(o.LogisticsCompanies) {
		var ret []LogisticsCompanyResult
		return ret
	}
	return o.LogisticsCompanies
}

// GetLogisticsCompaniesOk returns a tuple with the LogisticsCompanies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel) GetLogisticsCompaniesOk() ([]LogisticsCompanyResult, bool) {
	if o == nil || IsNil(o.LogisticsCompanies) {
		return nil, false
	}
	return o.LogisticsCompanies, true
}

// HasLogisticsCompanies returns a boolean if a field has been set.
func (o *AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel) HasLogisticsCompanies() bool {
	if o != nil && !IsNil(o.LogisticsCompanies) {
		return true
	}

	return false
}

// SetLogisticsCompanies gets a reference to the given []LogisticsCompanyResult and assigns it to the LogisticsCompanies field.
func (o *AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel) SetLogisticsCompanies(v []LogisticsCompanyResult) {
	o.LogisticsCompanies = v
}

func (o AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LogisticsCompanies) {
		toSerialize["logistics_companies"] = o.LogisticsCompanies
	}
	return toSerialize, nil
}

type NullableAlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel struct {
	value *AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel
	isSet bool
}

func (v NullableAlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel) Get() *AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel {
	return v.value
}

func (v *NullableAlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel) Set(val *AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel(val *AlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel) *NullableAlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel {
	return &NullableAlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceLogisticsLogisticscompanyInstantdeliveryQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
