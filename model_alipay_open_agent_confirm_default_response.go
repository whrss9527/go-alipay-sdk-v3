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

// AlipayOpenAgentConfirmDefaultResponse struct for AlipayOpenAgentConfirmDefaultResponse
type AlipayOpenAgentConfirmDefaultResponse struct {
	AlipayOpenAgentConfirmErrorResponseModel *AlipayOpenAgentConfirmErrorResponseModel
	CommonErrorType                          *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenAgentConfirmDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenAgentConfirmErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenAgentConfirmErrorResponseModel)
	if err == nil {
		jsonAlipayOpenAgentConfirmErrorResponseModel, _ := json.Marshal(dst.AlipayOpenAgentConfirmErrorResponseModel)
		if string(jsonAlipayOpenAgentConfirmErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenAgentConfirmErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenAgentConfirmErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenAgentConfirmErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenAgentConfirmDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenAgentConfirmDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenAgentConfirmErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenAgentConfirmErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenAgentConfirmDefaultResponse struct {
	value *AlipayOpenAgentConfirmDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenAgentConfirmDefaultResponse) Get() *AlipayOpenAgentConfirmDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenAgentConfirmDefaultResponse) Set(val *AlipayOpenAgentConfirmDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAgentConfirmDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAgentConfirmDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAgentConfirmDefaultResponse(val *AlipayOpenAgentConfirmDefaultResponse) *NullableAlipayOpenAgentConfirmDefaultResponse {
	return &NullableAlipayOpenAgentConfirmDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenAgentConfirmDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAgentConfirmDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
