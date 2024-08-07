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

// AlipayCommerceIotSdarttoolMessageQueryDefaultResponse struct for AlipayCommerceIotSdarttoolMessageQueryDefaultResponse
type AlipayCommerceIotSdarttoolMessageQueryDefaultResponse struct {
	AlipayCommerceIotSdarttoolMessageQueryErrorResponseModel *AlipayCommerceIotSdarttoolMessageQueryErrorResponseModel
	CommonErrorType                                          *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayCommerceIotSdarttoolMessageQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayCommerceIotSdarttoolMessageQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayCommerceIotSdarttoolMessageQueryErrorResponseModel)
	if err == nil {
		jsonAlipayCommerceIotSdarttoolMessageQueryErrorResponseModel, _ := json.Marshal(dst.AlipayCommerceIotSdarttoolMessageQueryErrorResponseModel)
		if string(jsonAlipayCommerceIotSdarttoolMessageQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayCommerceIotSdarttoolMessageQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayCommerceIotSdarttoolMessageQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayCommerceIotSdarttoolMessageQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayCommerceIotSdarttoolMessageQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayCommerceIotSdarttoolMessageQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayCommerceIotSdarttoolMessageQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayCommerceIotSdarttoolMessageQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayCommerceIotSdarttoolMessageQueryDefaultResponse struct {
	value *AlipayCommerceIotSdarttoolMessageQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayCommerceIotSdarttoolMessageQueryDefaultResponse) Get() *AlipayCommerceIotSdarttoolMessageQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayCommerceIotSdarttoolMessageQueryDefaultResponse) Set(val *AlipayCommerceIotSdarttoolMessageQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceIotSdarttoolMessageQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceIotSdarttoolMessageQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceIotSdarttoolMessageQueryDefaultResponse(val *AlipayCommerceIotSdarttoolMessageQueryDefaultResponse) *NullableAlipayCommerceIotSdarttoolMessageQueryDefaultResponse {
	return &NullableAlipayCommerceIotSdarttoolMessageQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayCommerceIotSdarttoolMessageQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceIotSdarttoolMessageQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
