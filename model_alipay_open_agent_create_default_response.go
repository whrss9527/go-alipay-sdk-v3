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

// AlipayOpenAgentCreateDefaultResponse struct for AlipayOpenAgentCreateDefaultResponse
type AlipayOpenAgentCreateDefaultResponse struct {
	AlipayOpenAgentCreateErrorResponseModel *AlipayOpenAgentCreateErrorResponseModel
	CommonErrorType                         *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenAgentCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenAgentCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenAgentCreateErrorResponseModel)
	if err == nil {
		jsonAlipayOpenAgentCreateErrorResponseModel, _ := json.Marshal(dst.AlipayOpenAgentCreateErrorResponseModel)
		if string(jsonAlipayOpenAgentCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenAgentCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenAgentCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenAgentCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenAgentCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenAgentCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenAgentCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenAgentCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenAgentCreateDefaultResponse struct {
	value *AlipayOpenAgentCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenAgentCreateDefaultResponse) Get() *AlipayOpenAgentCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenAgentCreateDefaultResponse) Set(val *AlipayOpenAgentCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAgentCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAgentCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAgentCreateDefaultResponse(val *AlipayOpenAgentCreateDefaultResponse) *NullableAlipayOpenAgentCreateDefaultResponse {
	return &NullableAlipayOpenAgentCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenAgentCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAgentCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
