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

// AlipayOpenMiniInnerversionOnlineDefaultResponse struct for AlipayOpenMiniInnerversionOnlineDefaultResponse
type AlipayOpenMiniInnerversionOnlineDefaultResponse struct {
	AlipayOpenMiniInnerversionOnlineErrorResponseModel *AlipayOpenMiniInnerversionOnlineErrorResponseModel
	CommonErrorType                                    *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniInnerversionOnlineDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniInnerversionOnlineErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniInnerversionOnlineErrorResponseModel)
	if err == nil {
		jsonAlipayOpenMiniInnerversionOnlineErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniInnerversionOnlineErrorResponseModel)
		if string(jsonAlipayOpenMiniInnerversionOnlineErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniInnerversionOnlineErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniInnerversionOnlineErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniInnerversionOnlineErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniInnerversionOnlineDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniInnerversionOnlineDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniInnerversionOnlineErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniInnerversionOnlineErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenMiniInnerversionOnlineDefaultResponse struct {
	value *AlipayOpenMiniInnerversionOnlineDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniInnerversionOnlineDefaultResponse) Get() *AlipayOpenMiniInnerversionOnlineDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerversionOnlineDefaultResponse) Set(val *AlipayOpenMiniInnerversionOnlineDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerversionOnlineDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerversionOnlineDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerversionOnlineDefaultResponse(val *AlipayOpenMiniInnerversionOnlineDefaultResponse) *NullableAlipayOpenMiniInnerversionOnlineDefaultResponse {
	return &NullableAlipayOpenMiniInnerversionOnlineDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerversionOnlineDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerversionOnlineDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
