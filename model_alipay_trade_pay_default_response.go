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


// AlipayTradePayDefaultResponse struct for AlipayTradePayDefaultResponse
type AlipayTradePayDefaultResponse struct {
	AlipayTradePayErrorResponseModel *AlipayTradePayErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayTradePayDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayTradePayErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayTradePayErrorResponseModel);
	if err == nil {
		jsonAlipayTradePayErrorResponseModel, _ := json.Marshal(dst.AlipayTradePayErrorResponseModel)
		if string(jsonAlipayTradePayErrorResponseModel) == "{}" { // empty struct
			dst.AlipayTradePayErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayTradePayErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayTradePayErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayTradePayDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayTradePayDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayTradePayErrorResponseModel != nil {
		return json.Marshal(&src.AlipayTradePayErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayTradePayDefaultResponse struct {
	value *AlipayTradePayDefaultResponse
	isSet bool
}

func (v NullableAlipayTradePayDefaultResponse) Get() *AlipayTradePayDefaultResponse {
	return v.value
}

func (v *NullableAlipayTradePayDefaultResponse) Set(val *AlipayTradePayDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayTradePayDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayTradePayDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayTradePayDefaultResponse(val *AlipayTradePayDefaultResponse) *NullableAlipayTradePayDefaultResponse {
	return &NullableAlipayTradePayDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayTradePayDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayTradePayDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

