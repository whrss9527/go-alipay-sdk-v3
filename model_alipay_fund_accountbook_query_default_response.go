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

// AlipayFundAccountbookQueryDefaultResponse struct for AlipayFundAccountbookQueryDefaultResponse
type AlipayFundAccountbookQueryDefaultResponse struct {
	AlipayFundAccountbookQueryErrorResponseModel *AlipayFundAccountbookQueryErrorResponseModel
	CommonErrorType                              *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayFundAccountbookQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayFundAccountbookQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayFundAccountbookQueryErrorResponseModel)
	if err == nil {
		jsonAlipayFundAccountbookQueryErrorResponseModel, _ := json.Marshal(dst.AlipayFundAccountbookQueryErrorResponseModel)
		if string(jsonAlipayFundAccountbookQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayFundAccountbookQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayFundAccountbookQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayFundAccountbookQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayFundAccountbookQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayFundAccountbookQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayFundAccountbookQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayFundAccountbookQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayFundAccountbookQueryDefaultResponse struct {
	value *AlipayFundAccountbookQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayFundAccountbookQueryDefaultResponse) Get() *AlipayFundAccountbookQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayFundAccountbookQueryDefaultResponse) Set(val *AlipayFundAccountbookQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundAccountbookQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundAccountbookQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundAccountbookQueryDefaultResponse(val *AlipayFundAccountbookQueryDefaultResponse) *NullableAlipayFundAccountbookQueryDefaultResponse {
	return &NullableAlipayFundAccountbookQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayFundAccountbookQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundAccountbookQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
