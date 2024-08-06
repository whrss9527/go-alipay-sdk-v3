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


// AlipayTradeRefundDefaultResponse struct for AlipayTradeRefundDefaultResponse
type AlipayTradeRefundDefaultResponse struct {
	AlipayTradeRefundErrorResponseModel *AlipayTradeRefundErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayTradeRefundDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayTradeRefundErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayTradeRefundErrorResponseModel);
	if err == nil {
		jsonAlipayTradeRefundErrorResponseModel, _ := json.Marshal(dst.AlipayTradeRefundErrorResponseModel)
		if string(jsonAlipayTradeRefundErrorResponseModel) == "{}" { // empty struct
			dst.AlipayTradeRefundErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayTradeRefundErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayTradeRefundErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayTradeRefundDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayTradeRefundDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayTradeRefundErrorResponseModel != nil {
		return json.Marshal(&src.AlipayTradeRefundErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayTradeRefundDefaultResponse struct {
	value *AlipayTradeRefundDefaultResponse
	isSet bool
}

func (v NullableAlipayTradeRefundDefaultResponse) Get() *AlipayTradeRefundDefaultResponse {
	return v.value
}

func (v *NullableAlipayTradeRefundDefaultResponse) Set(val *AlipayTradeRefundDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayTradeRefundDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayTradeRefundDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayTradeRefundDefaultResponse(val *AlipayTradeRefundDefaultResponse) *NullableAlipayTradeRefundDefaultResponse {
	return &NullableAlipayTradeRefundDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayTradeRefundDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayTradeRefundDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


