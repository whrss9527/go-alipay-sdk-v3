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

// checks if the AlipayUserAgreementPermissionCreateResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayUserAgreementPermissionCreateResponseModel{}

// AlipayUserAgreementPermissionCreateResponseModel struct for AlipayUserAgreementPermissionCreateResponseModel
type AlipayUserAgreementPermissionCreateResponseModel struct {
	// 商户代扣扣款许可
	DeductPermission *string `json:"deduct_permission,omitempty"`
	// 商户代扣扣款许可生效结束时间
	EffectiveTimeEnd *string `json:"effective_time_end,omitempty"`
	// 商户代扣扣款许可生效开始时间
	EffectiveTimeStart *string `json:"effective_time_start,omitempty"`
	// 消息是否发送成功
	IsSuccessSend *bool `json:"is_success_send,omitempty"`
}

// NewAlipayUserAgreementPermissionCreateResponseModel instantiates a new AlipayUserAgreementPermissionCreateResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayUserAgreementPermissionCreateResponseModel() *AlipayUserAgreementPermissionCreateResponseModel {
	this := AlipayUserAgreementPermissionCreateResponseModel{}
	return &this
}

// NewAlipayUserAgreementPermissionCreateResponseModelWithDefaults instantiates a new AlipayUserAgreementPermissionCreateResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayUserAgreementPermissionCreateResponseModelWithDefaults() *AlipayUserAgreementPermissionCreateResponseModel {
	this := AlipayUserAgreementPermissionCreateResponseModel{}
	return &this
}

// GetDeductPermission returns the DeductPermission field value if set, zero value otherwise.
func (o *AlipayUserAgreementPermissionCreateResponseModel) GetDeductPermission() string {
	if o == nil || IsNil(o.DeductPermission) {
		var ret string
		return ret
	}
	return *o.DeductPermission
}

// GetDeductPermissionOk returns a tuple with the DeductPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserAgreementPermissionCreateResponseModel) GetDeductPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.DeductPermission) {
		return nil, false
	}
	return o.DeductPermission, true
}

// HasDeductPermission returns a boolean if a field has been set.
func (o *AlipayUserAgreementPermissionCreateResponseModel) HasDeductPermission() bool {
	if o != nil && !IsNil(o.DeductPermission) {
		return true
	}

	return false
}

// SetDeductPermission gets a reference to the given string and assigns it to the DeductPermission field.
func (o *AlipayUserAgreementPermissionCreateResponseModel) SetDeductPermission(v string) {
	o.DeductPermission = &v
}

// GetEffectiveTimeEnd returns the EffectiveTimeEnd field value if set, zero value otherwise.
func (o *AlipayUserAgreementPermissionCreateResponseModel) GetEffectiveTimeEnd() string {
	if o == nil || IsNil(o.EffectiveTimeEnd) {
		var ret string
		return ret
	}
	return *o.EffectiveTimeEnd
}

// GetEffectiveTimeEndOk returns a tuple with the EffectiveTimeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserAgreementPermissionCreateResponseModel) GetEffectiveTimeEndOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveTimeEnd) {
		return nil, false
	}
	return o.EffectiveTimeEnd, true
}

// HasEffectiveTimeEnd returns a boolean if a field has been set.
func (o *AlipayUserAgreementPermissionCreateResponseModel) HasEffectiveTimeEnd() bool {
	if o != nil && !IsNil(o.EffectiveTimeEnd) {
		return true
	}

	return false
}

// SetEffectiveTimeEnd gets a reference to the given string and assigns it to the EffectiveTimeEnd field.
func (o *AlipayUserAgreementPermissionCreateResponseModel) SetEffectiveTimeEnd(v string) {
	o.EffectiveTimeEnd = &v
}

// GetEffectiveTimeStart returns the EffectiveTimeStart field value if set, zero value otherwise.
func (o *AlipayUserAgreementPermissionCreateResponseModel) GetEffectiveTimeStart() string {
	if o == nil || IsNil(o.EffectiveTimeStart) {
		var ret string
		return ret
	}
	return *o.EffectiveTimeStart
}

// GetEffectiveTimeStartOk returns a tuple with the EffectiveTimeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserAgreementPermissionCreateResponseModel) GetEffectiveTimeStartOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveTimeStart) {
		return nil, false
	}
	return o.EffectiveTimeStart, true
}

// HasEffectiveTimeStart returns a boolean if a field has been set.
func (o *AlipayUserAgreementPermissionCreateResponseModel) HasEffectiveTimeStart() bool {
	if o != nil && !IsNil(o.EffectiveTimeStart) {
		return true
	}

	return false
}

// SetEffectiveTimeStart gets a reference to the given string and assigns it to the EffectiveTimeStart field.
func (o *AlipayUserAgreementPermissionCreateResponseModel) SetEffectiveTimeStart(v string) {
	o.EffectiveTimeStart = &v
}

// GetIsSuccessSend returns the IsSuccessSend field value if set, zero value otherwise.
func (o *AlipayUserAgreementPermissionCreateResponseModel) GetIsSuccessSend() bool {
	if o == nil || IsNil(o.IsSuccessSend) {
		var ret bool
		return ret
	}
	return *o.IsSuccessSend
}

// GetIsSuccessSendOk returns a tuple with the IsSuccessSend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayUserAgreementPermissionCreateResponseModel) GetIsSuccessSendOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSuccessSend) {
		return nil, false
	}
	return o.IsSuccessSend, true
}

// HasIsSuccessSend returns a boolean if a field has been set.
func (o *AlipayUserAgreementPermissionCreateResponseModel) HasIsSuccessSend() bool {
	if o != nil && !IsNil(o.IsSuccessSend) {
		return true
	}

	return false
}

// SetIsSuccessSend gets a reference to the given bool and assigns it to the IsSuccessSend field.
func (o *AlipayUserAgreementPermissionCreateResponseModel) SetIsSuccessSend(v bool) {
	o.IsSuccessSend = &v
}

func (o AlipayUserAgreementPermissionCreateResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayUserAgreementPermissionCreateResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeductPermission) {
		toSerialize["deduct_permission"] = o.DeductPermission
	}
	if !IsNil(o.EffectiveTimeEnd) {
		toSerialize["effective_time_end"] = o.EffectiveTimeEnd
	}
	if !IsNil(o.EffectiveTimeStart) {
		toSerialize["effective_time_start"] = o.EffectiveTimeStart
	}
	if !IsNil(o.IsSuccessSend) {
		toSerialize["is_success_send"] = o.IsSuccessSend
	}
	return toSerialize, nil
}

type NullableAlipayUserAgreementPermissionCreateResponseModel struct {
	value *AlipayUserAgreementPermissionCreateResponseModel
	isSet bool
}

func (v NullableAlipayUserAgreementPermissionCreateResponseModel) Get() *AlipayUserAgreementPermissionCreateResponseModel {
	return v.value
}

func (v *NullableAlipayUserAgreementPermissionCreateResponseModel) Set(val *AlipayUserAgreementPermissionCreateResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayUserAgreementPermissionCreateResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayUserAgreementPermissionCreateResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayUserAgreementPermissionCreateResponseModel(val *AlipayUserAgreementPermissionCreateResponseModel) *NullableAlipayUserAgreementPermissionCreateResponseModel {
	return &NullableAlipayUserAgreementPermissionCreateResponseModel{value: val, isSet: true}
}

func (v NullableAlipayUserAgreementPermissionCreateResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayUserAgreementPermissionCreateResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


