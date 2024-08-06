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

// AlipayOpenMiniCategoryRequireQueryDefaultResponse struct for AlipayOpenMiniCategoryRequireQueryDefaultResponse
type AlipayOpenMiniCategoryRequireQueryDefaultResponse struct {
	AlipayOpenMiniCategoryRequireQueryErrorResponseModel *AlipayOpenMiniCategoryRequireQueryErrorResponseModel
	CommonErrorType                                      *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniCategoryRequireQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniCategoryRequireQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniCategoryRequireQueryErrorResponseModel)
	if err == nil {
		jsonAlipayOpenMiniCategoryRequireQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniCategoryRequireQueryErrorResponseModel)
		if string(jsonAlipayOpenMiniCategoryRequireQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniCategoryRequireQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniCategoryRequireQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniCategoryRequireQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniCategoryRequireQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniCategoryRequireQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniCategoryRequireQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniCategoryRequireQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenMiniCategoryRequireQueryDefaultResponse struct {
	value *AlipayOpenMiniCategoryRequireQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniCategoryRequireQueryDefaultResponse) Get() *AlipayOpenMiniCategoryRequireQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniCategoryRequireQueryDefaultResponse) Set(val *AlipayOpenMiniCategoryRequireQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniCategoryRequireQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniCategoryRequireQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniCategoryRequireQueryDefaultResponse(val *AlipayOpenMiniCategoryRequireQueryDefaultResponse) *NullableAlipayOpenMiniCategoryRequireQueryDefaultResponse {
	return &NullableAlipayOpenMiniCategoryRequireQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniCategoryRequireQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniCategoryRequireQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
