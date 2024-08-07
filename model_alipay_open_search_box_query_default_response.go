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

// AlipayOpenSearchBoxQueryDefaultResponse struct for AlipayOpenSearchBoxQueryDefaultResponse
type AlipayOpenSearchBoxQueryDefaultResponse struct {
	AlipayOpenSearchBoxQueryErrorResponseModel *AlipayOpenSearchBoxQueryErrorResponseModel
	CommonErrorType                            *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenSearchBoxQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenSearchBoxQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenSearchBoxQueryErrorResponseModel)
	if err == nil {
		jsonAlipayOpenSearchBoxQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenSearchBoxQueryErrorResponseModel)
		if string(jsonAlipayOpenSearchBoxQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenSearchBoxQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenSearchBoxQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenSearchBoxQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenSearchBoxQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenSearchBoxQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenSearchBoxQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenSearchBoxQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenSearchBoxQueryDefaultResponse struct {
	value *AlipayOpenSearchBoxQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenSearchBoxQueryDefaultResponse) Get() *AlipayOpenSearchBoxQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenSearchBoxQueryDefaultResponse) Set(val *AlipayOpenSearchBoxQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchBoxQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchBoxQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchBoxQueryDefaultResponse(val *AlipayOpenSearchBoxQueryDefaultResponse) *NullableAlipayOpenSearchBoxQueryDefaultResponse {
	return &NullableAlipayOpenSearchBoxQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchBoxQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchBoxQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
