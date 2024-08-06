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

// AlipayOpenMiniInnerMembersDeleteDefaultResponse struct for AlipayOpenMiniInnerMembersDeleteDefaultResponse
type AlipayOpenMiniInnerMembersDeleteDefaultResponse struct {
	AlipayOpenMiniInnerMembersDeleteErrorResponseModel *AlipayOpenMiniInnerMembersDeleteErrorResponseModel
	CommonErrorType                                    *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniInnerMembersDeleteDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniInnerMembersDeleteErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniInnerMembersDeleteErrorResponseModel)
	if err == nil {
		jsonAlipayOpenMiniInnerMembersDeleteErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniInnerMembersDeleteErrorResponseModel)
		if string(jsonAlipayOpenMiniInnerMembersDeleteErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniInnerMembersDeleteErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniInnerMembersDeleteErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniInnerMembersDeleteErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniInnerMembersDeleteDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniInnerMembersDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniInnerMembersDeleteErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniInnerMembersDeleteErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenMiniInnerMembersDeleteDefaultResponse struct {
	value *AlipayOpenMiniInnerMembersDeleteDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniInnerMembersDeleteDefaultResponse) Get() *AlipayOpenMiniInnerMembersDeleteDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerMembersDeleteDefaultResponse) Set(val *AlipayOpenMiniInnerMembersDeleteDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerMembersDeleteDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerMembersDeleteDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerMembersDeleteDefaultResponse(val *AlipayOpenMiniInnerMembersDeleteDefaultResponse) *NullableAlipayOpenMiniInnerMembersDeleteDefaultResponse {
	return &NullableAlipayOpenMiniInnerMembersDeleteDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerMembersDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerMembersDeleteDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
