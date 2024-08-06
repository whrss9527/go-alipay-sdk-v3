/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

import (
	"encoding/json"
	"fmt"
)

// AlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse struct for AlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse
type AlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse struct {
	AlipayCommerceTransportChargerChargerbindinfoSyncErrorResponseModel *AlipayCommerceTransportChargerChargerbindinfoSyncErrorResponseModel
	CommonErrorType                                                     *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayCommerceTransportChargerChargerbindinfoSyncErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayCommerceTransportChargerChargerbindinfoSyncErrorResponseModel)
	if err == nil {
		jsonAlipayCommerceTransportChargerChargerbindinfoSyncErrorResponseModel, _ := json.Marshal(dst.AlipayCommerceTransportChargerChargerbindinfoSyncErrorResponseModel)
		if string(jsonAlipayCommerceTransportChargerChargerbindinfoSyncErrorResponseModel) == "{}" { // empty struct
			dst.AlipayCommerceTransportChargerChargerbindinfoSyncErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayCommerceTransportChargerChargerbindinfoSyncErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayCommerceTransportChargerChargerbindinfoSyncErrorResponseModel = nil
	}

	// try to unmarshal JSON data into CommonErrorType
	err = json.Unmarshal(data, &dst.CommonErrorType)
	if err == nil {
		jsonCommonErrorType, _ := json.Marshal(dst.CommonErrorType)
		if string(jsonCommonErrorType) == "{}" { // empty struct
			dst.CommonErrorType = nil
		} else {
			return nil // data stored in dst.CommonErrorType, return on the first match
		}
	} else {
		dst.CommonErrorType = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayCommerceTransportChargerChargerbindinfoSyncErrorResponseModel != nil {
		return json.Marshal(&src.AlipayCommerceTransportChargerChargerbindinfoSyncErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse struct {
	value *AlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse
	isSet bool
}

func (v NullableAlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse) Get() *AlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse {
	return v.value
}

func (v *NullableAlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse) Set(val *AlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse(val *AlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse) *NullableAlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse {
	return &NullableAlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceTransportChargerChargerbindinfoSyncDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
