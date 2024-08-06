/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)


// AlipayMerchantIotDeviceBindDefaultResponse struct for AlipayMerchantIotDeviceBindDefaultResponse
type AlipayMerchantIotDeviceBindDefaultResponse struct {
	AlipayMerchantIotDeviceBindErrorResponseModel *AlipayMerchantIotDeviceBindErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMerchantIotDeviceBindDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMerchantIotDeviceBindErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMerchantIotDeviceBindErrorResponseModel);
	if err == nil {
		jsonAlipayMerchantIotDeviceBindErrorResponseModel, _ := json.Marshal(dst.AlipayMerchantIotDeviceBindErrorResponseModel)
		if string(jsonAlipayMerchantIotDeviceBindErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMerchantIotDeviceBindErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMerchantIotDeviceBindErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMerchantIotDeviceBindErrorResponseModel = nil
	}

	// try to unmarshal JSON data into CommonErrorType
	err = json.Unmarshal(data, &dst.CommonErrorType);
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMerchantIotDeviceBindDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMerchantIotDeviceBindDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMerchantIotDeviceBindErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMerchantIotDeviceBindErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMerchantIotDeviceBindDefaultResponse struct {
	value *AlipayMerchantIotDeviceBindDefaultResponse
	isSet bool
}

func (v NullableAlipayMerchantIotDeviceBindDefaultResponse) Get() *AlipayMerchantIotDeviceBindDefaultResponse {
	return v.value
}

func (v *NullableAlipayMerchantIotDeviceBindDefaultResponse) Set(val *AlipayMerchantIotDeviceBindDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMerchantIotDeviceBindDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMerchantIotDeviceBindDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMerchantIotDeviceBindDefaultResponse(val *AlipayMerchantIotDeviceBindDefaultResponse) *NullableAlipayMerchantIotDeviceBindDefaultResponse {
	return &NullableAlipayMerchantIotDeviceBindDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMerchantIotDeviceBindDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMerchantIotDeviceBindDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

