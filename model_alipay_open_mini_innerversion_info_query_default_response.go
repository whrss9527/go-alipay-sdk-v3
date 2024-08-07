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

// AlipayOpenMiniInnerversionInfoQueryDefaultResponse struct for AlipayOpenMiniInnerversionInfoQueryDefaultResponse
type AlipayOpenMiniInnerversionInfoQueryDefaultResponse struct {
	AlipayOpenMiniInnerversionInfoQueryErrorResponseModel *AlipayOpenMiniInnerversionInfoQueryErrorResponseModel
	CommonErrorType                                       *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniInnerversionInfoQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniInnerversionInfoQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniInnerversionInfoQueryErrorResponseModel)
	if err == nil {
		jsonAlipayOpenMiniInnerversionInfoQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniInnerversionInfoQueryErrorResponseModel)
		if string(jsonAlipayOpenMiniInnerversionInfoQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniInnerversionInfoQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniInnerversionInfoQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniInnerversionInfoQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniInnerversionInfoQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniInnerversionInfoQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniInnerversionInfoQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniInnerversionInfoQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenMiniInnerversionInfoQueryDefaultResponse struct {
	value *AlipayOpenMiniInnerversionInfoQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniInnerversionInfoQueryDefaultResponse) Get() *AlipayOpenMiniInnerversionInfoQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerversionInfoQueryDefaultResponse) Set(val *AlipayOpenMiniInnerversionInfoQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerversionInfoQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerversionInfoQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerversionInfoQueryDefaultResponse(val *AlipayOpenMiniInnerversionInfoQueryDefaultResponse) *NullableAlipayOpenMiniInnerversionInfoQueryDefaultResponse {
	return &NullableAlipayOpenMiniInnerversionInfoQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerversionInfoQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerversionInfoQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
