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

// AlipayCommerceEcEmployeeInviteQueryDefaultResponse struct for AlipayCommerceEcEmployeeInviteQueryDefaultResponse
type AlipayCommerceEcEmployeeInviteQueryDefaultResponse struct {
	AlipayCommerceEcEmployeeInviteQueryErrorResponseModel *AlipayCommerceEcEmployeeInviteQueryErrorResponseModel
	CommonErrorType                                       *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayCommerceEcEmployeeInviteQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayCommerceEcEmployeeInviteQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayCommerceEcEmployeeInviteQueryErrorResponseModel)
	if err == nil {
		jsonAlipayCommerceEcEmployeeInviteQueryErrorResponseModel, _ := json.Marshal(dst.AlipayCommerceEcEmployeeInviteQueryErrorResponseModel)
		if string(jsonAlipayCommerceEcEmployeeInviteQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayCommerceEcEmployeeInviteQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayCommerceEcEmployeeInviteQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayCommerceEcEmployeeInviteQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayCommerceEcEmployeeInviteQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayCommerceEcEmployeeInviteQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayCommerceEcEmployeeInviteQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayCommerceEcEmployeeInviteQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayCommerceEcEmployeeInviteQueryDefaultResponse struct {
	value *AlipayCommerceEcEmployeeInviteQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayCommerceEcEmployeeInviteQueryDefaultResponse) Get() *AlipayCommerceEcEmployeeInviteQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayCommerceEcEmployeeInviteQueryDefaultResponse) Set(val *AlipayCommerceEcEmployeeInviteQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEcEmployeeInviteQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEcEmployeeInviteQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEcEmployeeInviteQueryDefaultResponse(val *AlipayCommerceEcEmployeeInviteQueryDefaultResponse) *NullableAlipayCommerceEcEmployeeInviteQueryDefaultResponse {
	return &NullableAlipayCommerceEcEmployeeInviteQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayCommerceEcEmployeeInviteQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEcEmployeeInviteQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
