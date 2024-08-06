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

// AlipayIserviceCcmSwTreeDeleteDefaultResponse struct for AlipayIserviceCcmSwTreeDeleteDefaultResponse
type AlipayIserviceCcmSwTreeDeleteDefaultResponse struct {
	AlipayIserviceCcmSwTreeDeleteErrorResponseModel *AlipayIserviceCcmSwTreeDeleteErrorResponseModel
	CommonErrorType                                 *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayIserviceCcmSwTreeDeleteDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayIserviceCcmSwTreeDeleteErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayIserviceCcmSwTreeDeleteErrorResponseModel)
	if err == nil {
		jsonAlipayIserviceCcmSwTreeDeleteErrorResponseModel, _ := json.Marshal(dst.AlipayIserviceCcmSwTreeDeleteErrorResponseModel)
		if string(jsonAlipayIserviceCcmSwTreeDeleteErrorResponseModel) == "{}" { // empty struct
			dst.AlipayIserviceCcmSwTreeDeleteErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayIserviceCcmSwTreeDeleteErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayIserviceCcmSwTreeDeleteErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayIserviceCcmSwTreeDeleteDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayIserviceCcmSwTreeDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayIserviceCcmSwTreeDeleteErrorResponseModel != nil {
		return json.Marshal(&src.AlipayIserviceCcmSwTreeDeleteErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayIserviceCcmSwTreeDeleteDefaultResponse struct {
	value *AlipayIserviceCcmSwTreeDeleteDefaultResponse
	isSet bool
}

func (v NullableAlipayIserviceCcmSwTreeDeleteDefaultResponse) Get() *AlipayIserviceCcmSwTreeDeleteDefaultResponse {
	return v.value
}

func (v *NullableAlipayIserviceCcmSwTreeDeleteDefaultResponse) Set(val *AlipayIserviceCcmSwTreeDeleteDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmSwTreeDeleteDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmSwTreeDeleteDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmSwTreeDeleteDefaultResponse(val *AlipayIserviceCcmSwTreeDeleteDefaultResponse) *NullableAlipayIserviceCcmSwTreeDeleteDefaultResponse {
	return &NullableAlipayIserviceCcmSwTreeDeleteDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmSwTreeDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmSwTreeDeleteDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
