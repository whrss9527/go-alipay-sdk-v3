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


// AlipayOpenMiniQrcodeUnbindDefaultResponse struct for AlipayOpenMiniQrcodeUnbindDefaultResponse
type AlipayOpenMiniQrcodeUnbindDefaultResponse struct {
	AlipayOpenMiniQrcodeUnbindErrorResponseModel *AlipayOpenMiniQrcodeUnbindErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniQrcodeUnbindDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniQrcodeUnbindErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniQrcodeUnbindErrorResponseModel);
	if err == nil {
		jsonAlipayOpenMiniQrcodeUnbindErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniQrcodeUnbindErrorResponseModel)
		if string(jsonAlipayOpenMiniQrcodeUnbindErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniQrcodeUnbindErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniQrcodeUnbindErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniQrcodeUnbindErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniQrcodeUnbindDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniQrcodeUnbindDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniQrcodeUnbindErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniQrcodeUnbindErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenMiniQrcodeUnbindDefaultResponse struct {
	value *AlipayOpenMiniQrcodeUnbindDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniQrcodeUnbindDefaultResponse) Get() *AlipayOpenMiniQrcodeUnbindDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniQrcodeUnbindDefaultResponse) Set(val *AlipayOpenMiniQrcodeUnbindDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniQrcodeUnbindDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniQrcodeUnbindDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniQrcodeUnbindDefaultResponse(val *AlipayOpenMiniQrcodeUnbindDefaultResponse) *NullableAlipayOpenMiniQrcodeUnbindDefaultResponse {
	return &NullableAlipayOpenMiniQrcodeUnbindDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniQrcodeUnbindDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniQrcodeUnbindDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

