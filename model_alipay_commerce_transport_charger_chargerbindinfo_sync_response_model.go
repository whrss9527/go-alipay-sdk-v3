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

// checks if the AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel{}

// AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel struct for AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel
type AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel struct {
	// 商户同步的充电桩设备编码
	EquipId *string `json:"equip_id,omitempty"`
	// 绑定关系结果接收失败原因
	FailReason *string `json:"fail_reason,omitempty"`
	// 0：成功 1：拒绝
	ResultStatus *string `json:"result_status,omitempty"`
}

// NewAlipayCommerceTransportChargerChargerbindinfoSyncResponseModel instantiates a new AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlipayCommerceTransportChargerChargerbindinfoSyncResponseModel() *AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel {
	this := AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel{}
	return &this
}

// NewAlipayCommerceTransportChargerChargerbindinfoSyncResponseModelWithDefaults instantiates a new AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlipayCommerceTransportChargerChargerbindinfoSyncResponseModelWithDefaults() *AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel {
	this := AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel{}
	return &this
}

// GetEquipId returns the EquipId field value if set, zero value otherwise.
func (o *AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) GetEquipId() string {
	if o == nil || IsNil(o.EquipId) {
		var ret string
		return ret
	}
	return *o.EquipId
}

// GetEquipIdOk returns a tuple with the EquipId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) GetEquipIdOk() (*string, bool) {
	if o == nil || IsNil(o.EquipId) {
		return nil, false
	}
	return o.EquipId, true
}

// HasEquipId returns a boolean if a field has been set.
func (o *AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) HasEquipId() bool {
	if o != nil && !IsNil(o.EquipId) {
		return true
	}

	return false
}

// SetEquipId gets a reference to the given string and assigns it to the EquipId field.
func (o *AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) SetEquipId(v string) {
	o.EquipId = &v
}

// GetFailReason returns the FailReason field value if set, zero value otherwise.
func (o *AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) GetFailReason() string {
	if o == nil || IsNil(o.FailReason) {
		var ret string
		return ret
	}
	return *o.FailReason
}

// GetFailReasonOk returns a tuple with the FailReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) GetFailReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailReason) {
		return nil, false
	}
	return o.FailReason, true
}

// HasFailReason returns a boolean if a field has been set.
func (o *AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) HasFailReason() bool {
	if o != nil && !IsNil(o.FailReason) {
		return true
	}

	return false
}

// SetFailReason gets a reference to the given string and assigns it to the FailReason field.
func (o *AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) SetFailReason(v string) {
	o.FailReason = &v
}

// GetResultStatus returns the ResultStatus field value if set, zero value otherwise.
func (o *AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) GetResultStatus() string {
	if o == nil || IsNil(o.ResultStatus) {
		var ret string
		return ret
	}
	return *o.ResultStatus
}

// GetResultStatusOk returns a tuple with the ResultStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) GetResultStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ResultStatus) {
		return nil, false
	}
	return o.ResultStatus, true
}

// HasResultStatus returns a boolean if a field has been set.
func (o *AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) HasResultStatus() bool {
	if o != nil && !IsNil(o.ResultStatus) {
		return true
	}

	return false
}

// SetResultStatus gets a reference to the given string and assigns it to the ResultStatus field.
func (o *AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) SetResultStatus(v string) {
	o.ResultStatus = &v
}

func (o AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EquipId) {
		toSerialize["equip_id"] = o.EquipId
	}
	if !IsNil(o.FailReason) {
		toSerialize["fail_reason"] = o.FailReason
	}
	if !IsNil(o.ResultStatus) {
		toSerialize["result_status"] = o.ResultStatus
	}
	return toSerialize, nil
}

type NullableAlipayCommerceTransportChargerChargerbindinfoSyncResponseModel struct {
	value *AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel
	isSet bool
}

func (v NullableAlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) Get() *AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel {
	return v.value
}

func (v *NullableAlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) Set(val *AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceTransportChargerChargerbindinfoSyncResponseModel(val *AlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) *NullableAlipayCommerceTransportChargerChargerbindinfoSyncResponseModel {
	return &NullableAlipayCommerceTransportChargerChargerbindinfoSyncResponseModel{value: val, isSet: true}
}

func (v NullableAlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceTransportChargerChargerbindinfoSyncResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


