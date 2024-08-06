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


// AlipayOpenAgentCancelDefaultResponse struct for AlipayOpenAgentCancelDefaultResponse
type AlipayOpenAgentCancelDefaultResponse struct {
	AlipayOpenAgentCancelErrorResponseModel *AlipayOpenAgentCancelErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenAgentCancelDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenAgentCancelErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenAgentCancelErrorResponseModel);
	if err == nil {
		jsonAlipayOpenAgentCancelErrorResponseModel, _ := json.Marshal(dst.AlipayOpenAgentCancelErrorResponseModel)
		if string(jsonAlipayOpenAgentCancelErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenAgentCancelErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenAgentCancelErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenAgentCancelErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenAgentCancelDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenAgentCancelDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenAgentCancelErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenAgentCancelErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenAgentCancelDefaultResponse struct {
	value *AlipayOpenAgentCancelDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenAgentCancelDefaultResponse) Get() *AlipayOpenAgentCancelDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenAgentCancelDefaultResponse) Set(val *AlipayOpenAgentCancelDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAgentCancelDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAgentCancelDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAgentCancelDefaultResponse(val *AlipayOpenAgentCancelDefaultResponse) *NullableAlipayOpenAgentCancelDefaultResponse {
	return &NullableAlipayOpenAgentCancelDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenAgentCancelDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAgentCancelDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

