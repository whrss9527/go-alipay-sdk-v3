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


// AlipayOpenMiniAppdeployByappidQueryDefaultResponse struct for AlipayOpenMiniAppdeployByappidQueryDefaultResponse
type AlipayOpenMiniAppdeployByappidQueryDefaultResponse struct {
	AlipayOpenMiniAppdeployByappidQueryErrorResponseModel *AlipayOpenMiniAppdeployByappidQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniAppdeployByappidQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniAppdeployByappidQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniAppdeployByappidQueryErrorResponseModel);
	if err == nil {
		jsonAlipayOpenMiniAppdeployByappidQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniAppdeployByappidQueryErrorResponseModel)
		if string(jsonAlipayOpenMiniAppdeployByappidQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniAppdeployByappidQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniAppdeployByappidQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniAppdeployByappidQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniAppdeployByappidQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniAppdeployByappidQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniAppdeployByappidQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniAppdeployByappidQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenMiniAppdeployByappidQueryDefaultResponse struct {
	value *AlipayOpenMiniAppdeployByappidQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniAppdeployByappidQueryDefaultResponse) Get() *AlipayOpenMiniAppdeployByappidQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniAppdeployByappidQueryDefaultResponse) Set(val *AlipayOpenMiniAppdeployByappidQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniAppdeployByappidQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniAppdeployByappidQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniAppdeployByappidQueryDefaultResponse(val *AlipayOpenMiniAppdeployByappidQueryDefaultResponse) *NullableAlipayOpenMiniAppdeployByappidQueryDefaultResponse {
	return &NullableAlipayOpenMiniAppdeployByappidQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniAppdeployByappidQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniAppdeployByappidQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

