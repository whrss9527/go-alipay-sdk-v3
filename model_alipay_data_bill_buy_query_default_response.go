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

// AlipayDataBillBuyQueryDefaultResponse struct for AlipayDataBillBuyQueryDefaultResponse
type AlipayDataBillBuyQueryDefaultResponse struct {
	AlipayDataBillBuyQueryErrorResponseModel *AlipayDataBillBuyQueryErrorResponseModel
	CommonErrorType                          *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayDataBillBuyQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayDataBillBuyQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayDataBillBuyQueryErrorResponseModel)
	if err == nil {
		jsonAlipayDataBillBuyQueryErrorResponseModel, _ := json.Marshal(dst.AlipayDataBillBuyQueryErrorResponseModel)
		if string(jsonAlipayDataBillBuyQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayDataBillBuyQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayDataBillBuyQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayDataBillBuyQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayDataBillBuyQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayDataBillBuyQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayDataBillBuyQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayDataBillBuyQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayDataBillBuyQueryDefaultResponse struct {
	value *AlipayDataBillBuyQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayDataBillBuyQueryDefaultResponse) Get() *AlipayDataBillBuyQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayDataBillBuyQueryDefaultResponse) Set(val *AlipayDataBillBuyQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayDataBillBuyQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayDataBillBuyQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayDataBillBuyQueryDefaultResponse(val *AlipayDataBillBuyQueryDefaultResponse) *NullableAlipayDataBillBuyQueryDefaultResponse {
	return &NullableAlipayDataBillBuyQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayDataBillBuyQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayDataBillBuyQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
