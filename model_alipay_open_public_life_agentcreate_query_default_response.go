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

// AlipayOpenPublicLifeAgentcreateQueryDefaultResponse struct for AlipayOpenPublicLifeAgentcreateQueryDefaultResponse
type AlipayOpenPublicLifeAgentcreateQueryDefaultResponse struct {
	AlipayOpenPublicLifeAgentcreateQueryErrorResponseModel *AlipayOpenPublicLifeAgentcreateQueryErrorResponseModel
	CommonErrorType                                        *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicLifeAgentcreateQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicLifeAgentcreateQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicLifeAgentcreateQueryErrorResponseModel)
	if err == nil {
		jsonAlipayOpenPublicLifeAgentcreateQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicLifeAgentcreateQueryErrorResponseModel)
		if string(jsonAlipayOpenPublicLifeAgentcreateQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicLifeAgentcreateQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicLifeAgentcreateQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicLifeAgentcreateQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicLifeAgentcreateQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicLifeAgentcreateQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicLifeAgentcreateQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicLifeAgentcreateQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenPublicLifeAgentcreateQueryDefaultResponse struct {
	value *AlipayOpenPublicLifeAgentcreateQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicLifeAgentcreateQueryDefaultResponse) Get() *AlipayOpenPublicLifeAgentcreateQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicLifeAgentcreateQueryDefaultResponse) Set(val *AlipayOpenPublicLifeAgentcreateQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicLifeAgentcreateQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicLifeAgentcreateQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicLifeAgentcreateQueryDefaultResponse(val *AlipayOpenPublicLifeAgentcreateQueryDefaultResponse) *NullableAlipayOpenPublicLifeAgentcreateQueryDefaultResponse {
	return &NullableAlipayOpenPublicLifeAgentcreateQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicLifeAgentcreateQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicLifeAgentcreateQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
