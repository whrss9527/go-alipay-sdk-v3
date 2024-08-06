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

// checks if the AlipayUserCertifyOpenInitializeResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayUserCertifyOpenInitializeResponseModel{}

// AlipayUserCertifyOpenInitializeResponseModel struct for AlipayUserCertifyOpenInitializeResponseModel
type AlipayUserCertifyOpenInitializeResponseModel struct {
	// 本次申请操作的唯一标识，商户需要记录，后续的操作都需要用到
	CertifyId *string `json:"certify_id,omitempty"`
}

// NewAlipayUserCertifyOpenInitializeResponseModel instantiates a new AlipayUserCertifyOpenInitializeResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayUserCertifyOpenInitializeResponseModel() *AlipayUserCertifyOpenInitializeResponseModel {
	this := AlipayUserCertifyOpenInitializeResponseModel{}
	return &this
}

// NewAlipayUserCertifyOpenInitializeResponseModelWithDefaults instantiates a new AlipayUserCertifyOpenInitializeResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayUserCertifyOpenInitializeResponseModelWithDefaults() *AlipayUserCertifyOpenInitializeResponseModel {
	this := AlipayUserCertifyOpenInitializeResponseModel{}
	return &this
}

// GetCertifyId returns the CertifyId field value if set, zero value otherwise.
func (o *AlipayUserCertifyOpenInitializeResponseModel) GetCertifyId() string {
	if o == nil || IsNil(o.CertifyId) {
		var ret string
		return ret
	}
	return *o.CertifyId
}

// GetCertifyIdOk returns a tuple with the CertifyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserCertifyOpenInitializeResponseModel) GetCertifyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CertifyId) {
		return nil, false
	}
	return o.CertifyId, true
}

// HasCertifyId returns a boolean if a field has been set.
func (o *AlipayUserCertifyOpenInitializeResponseModel) HasCertifyId() bool {
	if o != nil && !IsNil(o.CertifyId) {
		return true
	}

	return false
}

// SetCertifyId gets a reference to the given string and assigns it to the CertifyId field.
func (o *AlipayUserCertifyOpenInitializeResponseModel) SetCertifyId(v string) {
	o.CertifyId = &v
}

func (o AlipayUserCertifyOpenInitializeResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayUserCertifyOpenInitializeResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertifyId) {
		toSerialize["certify_id"] = o.CertifyId
	}
	return toSerialize, nil
}

type NullableAlipayUserCertifyOpenInitializeResponseModel struct {
	value *AlipayUserCertifyOpenInitializeResponseModel
	isSet bool
}

func (v NullableAlipayUserCertifyOpenInitializeResponseModel) Get() *AlipayUserCertifyOpenInitializeResponseModel {
	return v.value
}

func (v *NullableAlipayUserCertifyOpenInitializeResponseModel) Set(val *AlipayUserCertifyOpenInitializeResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayUserCertifyOpenInitializeResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayUserCertifyOpenInitializeResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayUserCertifyOpenInitializeResponseModel(val *AlipayUserCertifyOpenInitializeResponseModel) *NullableAlipayUserCertifyOpenInitializeResponseModel {
	return &NullableAlipayUserCertifyOpenInitializeResponseModel{value: val, isSet: true}
}

func (v NullableAlipayUserCertifyOpenInitializeResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayUserCertifyOpenInitializeResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
