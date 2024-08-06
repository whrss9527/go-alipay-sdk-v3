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

// AlipayOpenSearchSubservicekeywordDeleteDefaultResponse struct for AlipayOpenSearchSubservicekeywordDeleteDefaultResponse
type AlipayOpenSearchSubservicekeywordDeleteDefaultResponse struct {
	AlipayOpenSearchSubservicekeywordDeleteErrorResponseModel *AlipayOpenSearchSubservicekeywordDeleteErrorResponseModel
	CommonErrorType                                           *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenSearchSubservicekeywordDeleteDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenSearchSubservicekeywordDeleteErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenSearchSubservicekeywordDeleteErrorResponseModel)
	if err == nil {
		jsonAlipayOpenSearchSubservicekeywordDeleteErrorResponseModel, _ := json.Marshal(dst.AlipayOpenSearchSubservicekeywordDeleteErrorResponseModel)
		if string(jsonAlipayOpenSearchSubservicekeywordDeleteErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenSearchSubservicekeywordDeleteErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenSearchSubservicekeywordDeleteErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenSearchSubservicekeywordDeleteErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenSearchSubservicekeywordDeleteDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenSearchSubservicekeywordDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenSearchSubservicekeywordDeleteErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenSearchSubservicekeywordDeleteErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenSearchSubservicekeywordDeleteDefaultResponse struct {
	value *AlipayOpenSearchSubservicekeywordDeleteDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenSearchSubservicekeywordDeleteDefaultResponse) Get() *AlipayOpenSearchSubservicekeywordDeleteDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenSearchSubservicekeywordDeleteDefaultResponse) Set(val *AlipayOpenSearchSubservicekeywordDeleteDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchSubservicekeywordDeleteDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchSubservicekeywordDeleteDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchSubservicekeywordDeleteDefaultResponse(val *AlipayOpenSearchSubservicekeywordDeleteDefaultResponse) *NullableAlipayOpenSearchSubservicekeywordDeleteDefaultResponse {
	return &NullableAlipayOpenSearchSubservicekeywordDeleteDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchSubservicekeywordDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchSubservicekeywordDeleteDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
