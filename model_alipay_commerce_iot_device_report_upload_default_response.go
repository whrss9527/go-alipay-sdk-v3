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


// AlipayCommerceIotDeviceReportUploadDefaultResponse struct for AlipayCommerceIotDeviceReportUploadDefaultResponse
type AlipayCommerceIotDeviceReportUploadDefaultResponse struct {
	AlipayCommerceIotDeviceReportUploadErrorResponseModel *AlipayCommerceIotDeviceReportUploadErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayCommerceIotDeviceReportUploadDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayCommerceIotDeviceReportUploadErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayCommerceIotDeviceReportUploadErrorResponseModel);
	if err == nil {
		jsonAlipayCommerceIotDeviceReportUploadErrorResponseModel, _ := json.Marshal(dst.AlipayCommerceIotDeviceReportUploadErrorResponseModel)
		if string(jsonAlipayCommerceIotDeviceReportUploadErrorResponseModel) == "{}" { // empty struct
			dst.AlipayCommerceIotDeviceReportUploadErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayCommerceIotDeviceReportUploadErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayCommerceIotDeviceReportUploadErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayCommerceIotDeviceReportUploadDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayCommerceIotDeviceReportUploadDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayCommerceIotDeviceReportUploadErrorResponseModel != nil {
		return json.Marshal(&src.AlipayCommerceIotDeviceReportUploadErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayCommerceIotDeviceReportUploadDefaultResponse struct {
	value *AlipayCommerceIotDeviceReportUploadDefaultResponse
	isSet bool
}

func (v NullableAlipayCommerceIotDeviceReportUploadDefaultResponse) Get() *AlipayCommerceIotDeviceReportUploadDefaultResponse {
	return v.value
}

func (v *NullableAlipayCommerceIotDeviceReportUploadDefaultResponse) Set(val *AlipayCommerceIotDeviceReportUploadDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceIotDeviceReportUploadDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceIotDeviceReportUploadDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceIotDeviceReportUploadDefaultResponse(val *AlipayCommerceIotDeviceReportUploadDefaultResponse) *NullableAlipayCommerceIotDeviceReportUploadDefaultResponse {
	return &NullableAlipayCommerceIotDeviceReportUploadDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayCommerceIotDeviceReportUploadDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceIotDeviceReportUploadDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


