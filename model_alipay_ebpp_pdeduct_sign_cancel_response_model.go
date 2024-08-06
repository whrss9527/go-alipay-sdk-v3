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

// checks if the AlipayEbppPdeductSignCancelResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayEbppPdeductSignCancelResponseModel{}

// AlipayEbppPdeductSignCancelResponseModel struct for AlipayEbppPdeductSignCancelResponseModel
type AlipayEbppPdeductSignCancelResponseModel struct {
	// 支付宝代扣协议ID
	AgreementId *string `json:"agreement_id,omitempty"`
	// 支付宝协议状态。解约成功则返回success
	AgreementStatus *string `json:"agreement_status,omitempty"`
	// 商户代扣协议ID
	OutAgreementId *string `json:"out_agreement_id,omitempty"`
}

// NewAlipayEbppPdeductSignCancelResponseModel instantiates a new AlipayEbppPdeductSignCancelResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayEbppPdeductSignCancelResponseModel() *AlipayEbppPdeductSignCancelResponseModel {
	this := AlipayEbppPdeductSignCancelResponseModel{}
	return &this
}

// NewAlipayEbppPdeductSignCancelResponseModelWithDefaults instantiates a new AlipayEbppPdeductSignCancelResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayEbppPdeductSignCancelResponseModelWithDefaults() *AlipayEbppPdeductSignCancelResponseModel {
	this := AlipayEbppPdeductSignCancelResponseModel{}
	return &this
}

// GetAgreementId returns the AgreementId field value if set, zero value otherwise.
func (o *AlipayEbppPdeductSignCancelResponseModel) GetAgreementId() string {
	if o == nil || IsNil(o.AgreementId) {
		var ret string
		return ret
	}
	return *o.AgreementId
}

// GetAgreementIdOk returns a tuple with the AgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductSignCancelResponseModel) GetAgreementIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementId) {
		return nil, false
	}
	return o.AgreementId, true
}

// HasAgreementId returns a boolean if a field has been set.
func (o *AlipayEbppPdeductSignCancelResponseModel) HasAgreementId() bool {
	if o != nil && !IsNil(o.AgreementId) {
		return true
	}

	return false
}

// SetAgreementId gets a reference to the given string and assigns it to the AgreementId field.
func (o *AlipayEbppPdeductSignCancelResponseModel) SetAgreementId(v string) {
	o.AgreementId = &v
}

// GetAgreementStatus returns the AgreementStatus field value if set, zero value otherwise.
func (o *AlipayEbppPdeductSignCancelResponseModel) GetAgreementStatus() string {
	if o == nil || IsNil(o.AgreementStatus) {
		var ret string
		return ret
	}
	return *o.AgreementStatus
}

// GetAgreementStatusOk returns a tuple with the AgreementStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductSignCancelResponseModel) GetAgreementStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementStatus) {
		return nil, false
	}
	return o.AgreementStatus, true
}

// HasAgreementStatus returns a boolean if a field has been set.
func (o *AlipayEbppPdeductSignCancelResponseModel) HasAgreementStatus() bool {
	if o != nil && !IsNil(o.AgreementStatus) {
		return true
	}

	return false
}

// SetAgreementStatus gets a reference to the given string and assigns it to the AgreementStatus field.
func (o *AlipayEbppPdeductSignCancelResponseModel) SetAgreementStatus(v string) {
	o.AgreementStatus = &v
}

// GetOutAgreementId returns the OutAgreementId field value if set, zero value otherwise.
func (o *AlipayEbppPdeductSignCancelResponseModel) GetOutAgreementId() string {
	if o == nil || IsNil(o.OutAgreementId) {
		var ret string
		return ret
	}
	return *o.OutAgreementId
}

// GetOutAgreementIdOk returns a tuple with the OutAgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayEbppPdeductSignCancelResponseModel) GetOutAgreementIdOk() (*string, bool) {
	if o == nil || IsNil(o.OutAgreementId) {
		return nil, false
	}
	return o.OutAgreementId, true
}

// HasOutAgreementId returns a boolean if a field has been set.
func (o *AlipayEbppPdeductSignCancelResponseModel) HasOutAgreementId() bool {
	if o != nil && !IsNil(o.OutAgreementId) {
		return true
	}

	return false
}

// SetOutAgreementId gets a reference to the given string and assigns it to the OutAgreementId field.
func (o *AlipayEbppPdeductSignCancelResponseModel) SetOutAgreementId(v string) {
	o.OutAgreementId = &v
}

func (o AlipayEbppPdeductSignCancelResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayEbppPdeductSignCancelResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementId) {
		toSerialize["agreement_id"] = o.AgreementId
	}
	if !IsNil(o.AgreementStatus) {
		toSerialize["agreement_status"] = o.AgreementStatus
	}
	if !IsNil(o.OutAgreementId) {
		toSerialize["out_agreement_id"] = o.OutAgreementId
	}
	return toSerialize, nil
}

type NullableAlipayEbppPdeductSignCancelResponseModel struct {
	value *AlipayEbppPdeductSignCancelResponseModel
	isSet bool
}

func (v NullableAlipayEbppPdeductSignCancelResponseModel) Get() *AlipayEbppPdeductSignCancelResponseModel {
	return v.value
}

func (v *NullableAlipayEbppPdeductSignCancelResponseModel) Set(val *AlipayEbppPdeductSignCancelResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayEbppPdeductSignCancelResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayEbppPdeductSignCancelResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayEbppPdeductSignCancelResponseModel(val *AlipayEbppPdeductSignCancelResponseModel) *NullableAlipayEbppPdeductSignCancelResponseModel {
	return &NullableAlipayEbppPdeductSignCancelResponseModel{value: val, isSet: true}
}

func (v NullableAlipayEbppPdeductSignCancelResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayEbppPdeductSignCancelResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

