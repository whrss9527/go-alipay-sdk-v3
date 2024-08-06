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


// AlipayTradeOrderPayDefaultResponse struct for AlipayTradeOrderPayDefaultResponse
type AlipayTradeOrderPayDefaultResponse struct {
	AlipayTradeOrderPayErrorResponseModel *AlipayTradeOrderPayErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayTradeOrderPayDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayTradeOrderPayErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayTradeOrderPayErrorResponseModel);
	if err == nil {
		jsonAlipayTradeOrderPayErrorResponseModel, _ := json.Marshal(dst.AlipayTradeOrderPayErrorResponseModel)
		if string(jsonAlipayTradeOrderPayErrorResponseModel) == "{}" { // empty struct
			dst.AlipayTradeOrderPayErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayTradeOrderPayErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayTradeOrderPayErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayTradeOrderPayDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayTradeOrderPayDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayTradeOrderPayErrorResponseModel != nil {
		return json.Marshal(&src.AlipayTradeOrderPayErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayTradeOrderPayDefaultResponse struct {
	value *AlipayTradeOrderPayDefaultResponse
	isSet bool
}

func (v NullableAlipayTradeOrderPayDefaultResponse) Get() *AlipayTradeOrderPayDefaultResponse {
	return v.value
}

func (v *NullableAlipayTradeOrderPayDefaultResponse) Set(val *AlipayTradeOrderPayDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayTradeOrderPayDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayTradeOrderPayDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayTradeOrderPayDefaultResponse(val *AlipayTradeOrderPayDefaultResponse) *NullableAlipayTradeOrderPayDefaultResponse {
	return &NullableAlipayTradeOrderPayDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayTradeOrderPayDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayTradeOrderPayDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


