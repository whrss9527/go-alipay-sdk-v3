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

// checks if the AlipayUserAgreementUserverifyQueryResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayUserAgreementUserverifyQueryResponseModel{}

// AlipayUserAgreementUserverifyQueryResponseModel struct for AlipayUserAgreementUserverifyQueryResponseModel
type AlipayUserAgreementUserverifyQueryResponseModel struct {
	// 用户信息是否匹配
	UserMatch *bool `json:"user_match,omitempty"`
}

// NewAlipayUserAgreementUserverifyQueryResponseModel instantiates a new AlipayUserAgreementUserverifyQueryResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayUserAgreementUserverifyQueryResponseModel() *AlipayUserAgreementUserverifyQueryResponseModel {
	this := AlipayUserAgreementUserverifyQueryResponseModel{}
	return &this
}

// NewAlipayUserAgreementUserverifyQueryResponseModelWithDefaults instantiates a new AlipayUserAgreementUserverifyQueryResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayUserAgreementUserverifyQueryResponseModelWithDefaults() *AlipayUserAgreementUserverifyQueryResponseModel {
	this := AlipayUserAgreementUserverifyQueryResponseModel{}
	return &this
}

// GetUserMatch returns the UserMatch field value if set, zero value otherwise.
func (o *AlipayUserAgreementUserverifyQueryResponseModel) GetUserMatch() bool {
	if o == nil || IsNil(o.UserMatch) {
		var ret bool
		return ret
	}
	return *o.UserMatch
}

// GetUserMatchOk returns a tuple with the UserMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserAgreementUserverifyQueryResponseModel) GetUserMatchOk() (*bool, bool) {
	if o == nil || IsNil(o.UserMatch) {
		return nil, false
	}
	return o.UserMatch, true
}

// HasUserMatch returns a boolean if a field has been set.
func (o *AlipayUserAgreementUserverifyQueryResponseModel) HasUserMatch() bool {
	if o != nil && !IsNil(o.UserMatch) {
		return true
	}

	return false
}

// SetUserMatch gets a reference to the given bool and assigns it to the UserMatch field.
func (o *AlipayUserAgreementUserverifyQueryResponseModel) SetUserMatch(v bool) {
	o.UserMatch = &v
}

func (o AlipayUserAgreementUserverifyQueryResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayUserAgreementUserverifyQueryResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserMatch) {
		toSerialize["user_match"] = o.UserMatch
	}
	return toSerialize, nil
}

type NullableAlipayUserAgreementUserverifyQueryResponseModel struct {
	value *AlipayUserAgreementUserverifyQueryResponseModel
	isSet bool
}

func (v NullableAlipayUserAgreementUserverifyQueryResponseModel) Get() *AlipayUserAgreementUserverifyQueryResponseModel {
	return v.value
}

func (v *NullableAlipayUserAgreementUserverifyQueryResponseModel) Set(val *AlipayUserAgreementUserverifyQueryResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayUserAgreementUserverifyQueryResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayUserAgreementUserverifyQueryResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayUserAgreementUserverifyQueryResponseModel(val *AlipayUserAgreementUserverifyQueryResponseModel) *NullableAlipayUserAgreementUserverifyQueryResponseModel {
	return &NullableAlipayUserAgreementUserverifyQueryResponseModel{value: val, isSet: true}
}

func (v NullableAlipayUserAgreementUserverifyQueryResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayUserAgreementUserverifyQueryResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


