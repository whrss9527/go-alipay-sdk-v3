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

// AlipayOpenMiniQrcodeQueryDefaultResponse struct for AlipayOpenMiniQrcodeQueryDefaultResponse
type AlipayOpenMiniQrcodeQueryDefaultResponse struct {
	AlipayOpenMiniQrcodeQueryErrorResponseModel *AlipayOpenMiniQrcodeQueryErrorResponseModel
	CommonErrorType                             *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniQrcodeQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniQrcodeQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniQrcodeQueryErrorResponseModel)
	if err == nil {
		jsonAlipayOpenMiniQrcodeQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniQrcodeQueryErrorResponseModel)
		if string(jsonAlipayOpenMiniQrcodeQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniQrcodeQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniQrcodeQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniQrcodeQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniQrcodeQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniQrcodeQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniQrcodeQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniQrcodeQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenMiniQrcodeQueryDefaultResponse struct {
	value *AlipayOpenMiniQrcodeQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniQrcodeQueryDefaultResponse) Get() *AlipayOpenMiniQrcodeQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniQrcodeQueryDefaultResponse) Set(val *AlipayOpenMiniQrcodeQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniQrcodeQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniQrcodeQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniQrcodeQueryDefaultResponse(val *AlipayOpenMiniQrcodeQueryDefaultResponse) *NullableAlipayOpenMiniQrcodeQueryDefaultResponse {
	return &NullableAlipayOpenMiniQrcodeQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniQrcodeQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniQrcodeQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
