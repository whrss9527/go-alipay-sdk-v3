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

// AlipayOpenPublicGroupCrowdQueryDefaultResponse struct for AlipayOpenPublicGroupCrowdQueryDefaultResponse
type AlipayOpenPublicGroupCrowdQueryDefaultResponse struct {
	AlipayOpenPublicGroupCrowdQueryErrorResponseModel *AlipayOpenPublicGroupCrowdQueryErrorResponseModel
	CommonErrorType                                   *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicGroupCrowdQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicGroupCrowdQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicGroupCrowdQueryErrorResponseModel)
	if err == nil {
		jsonAlipayOpenPublicGroupCrowdQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicGroupCrowdQueryErrorResponseModel)
		if string(jsonAlipayOpenPublicGroupCrowdQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicGroupCrowdQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicGroupCrowdQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicGroupCrowdQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicGroupCrowdQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicGroupCrowdQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicGroupCrowdQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicGroupCrowdQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenPublicGroupCrowdQueryDefaultResponse struct {
	value *AlipayOpenPublicGroupCrowdQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicGroupCrowdQueryDefaultResponse) Get() *AlipayOpenPublicGroupCrowdQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicGroupCrowdQueryDefaultResponse) Set(val *AlipayOpenPublicGroupCrowdQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicGroupCrowdQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicGroupCrowdQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicGroupCrowdQueryDefaultResponse(val *AlipayOpenPublicGroupCrowdQueryDefaultResponse) *NullableAlipayOpenPublicGroupCrowdQueryDefaultResponse {
	return &NullableAlipayOpenPublicGroupCrowdQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicGroupCrowdQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicGroupCrowdQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
