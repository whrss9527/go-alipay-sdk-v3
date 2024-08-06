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

// checks if the AlipayMarketingCardBenefitCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayMarketingCardBenefitCreateResponseModel{}

// AlipayMarketingCardBenefitCreateResponseModel struct for AlipayMarketingCardBenefitCreateResponseModel
type AlipayMarketingCardBenefitCreateResponseModel struct {
	// 权益ID
	BenefitId *string `json:"benefit_id,omitempty"`
}

// NewAlipayMarketingCardBenefitCreateResponseModel instantiates a new AlipayMarketingCardBenefitCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayMarketingCardBenefitCreateResponseModel() *AlipayMarketingCardBenefitCreateResponseModel {
	this := AlipayMarketingCardBenefitCreateResponseModel{}
	return &this
}

// NewAlipayMarketingCardBenefitCreateResponseModelWithDefaults instantiates a new AlipayMarketingCardBenefitCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayMarketingCardBenefitCreateResponseModelWithDefaults() *AlipayMarketingCardBenefitCreateResponseModel {
	this := AlipayMarketingCardBenefitCreateResponseModel{}
	return &this
}

// GetBenefitId returns the BenefitId field value if set, zero value otherwise.
func (o *AlipayMarketingCardBenefitCreateResponseModel) GetBenefitId() string {
	if o == nil || IsNil(o.BenefitId) {
		var ret string
		return ret
	}
	return *o.BenefitId
}

// GetBenefitIdOk returns a tuple with the BenefitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayMarketingCardBenefitCreateResponseModel) GetBenefitIdOk() (*string, bool) {
	if o == nil || IsNil(o.BenefitId) {
		return nil, false
	}
	return o.BenefitId, true
}

// HasBenefitId returns a boolean if a field has been set.
func (o *AlipayMarketingCardBenefitCreateResponseModel) HasBenefitId() bool {
	if o != nil && !IsNil(o.BenefitId) {
		return true
	}

	return false
}

// SetBenefitId gets a reference to the given string and assigns it to the BenefitId field.
func (o *AlipayMarketingCardBenefitCreateResponseModel) SetBenefitId(v string) {
	o.BenefitId = &v
}

func (o AlipayMarketingCardBenefitCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayMarketingCardBenefitCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BenefitId) {
		toSerialize["benefit_id"] = o.BenefitId
	}
	return toSerialize, nil
}

type NullableAlipayMarketingCardBenefitCreateResponseModel struct {
	value *AlipayMarketingCardBenefitCreateResponseModel
	isSet bool
}

func (v NullableAlipayMarketingCardBenefitCreateResponseModel) Get() *AlipayMarketingCardBenefitCreateResponseModel {
	return v.value
}

func (v *NullableAlipayMarketingCardBenefitCreateResponseModel) Set(val *AlipayMarketingCardBenefitCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCardBenefitCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCardBenefitCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCardBenefitCreateResponseModel(val *AlipayMarketingCardBenefitCreateResponseModel) *NullableAlipayMarketingCardBenefitCreateResponseModel {
	return &NullableAlipayMarketingCardBenefitCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayMarketingCardBenefitCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCardBenefitCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


