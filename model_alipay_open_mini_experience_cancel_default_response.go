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

// AlipayOpenMiniExperienceCancelDefaultResponse struct for AlipayOpenMiniExperienceCancelDefaultResponse
type AlipayOpenMiniExperienceCancelDefaultResponse struct {
	AlipayOpenMiniExperienceCancelErrorResponseModel *AlipayOpenMiniExperienceCancelErrorResponseModel
	CommonErrorType                                  *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniExperienceCancelDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniExperienceCancelErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniExperienceCancelErrorResponseModel)
	if err == nil {
		jsonAlipayOpenMiniExperienceCancelErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniExperienceCancelErrorResponseModel)
		if string(jsonAlipayOpenMiniExperienceCancelErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniExperienceCancelErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniExperienceCancelErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniExperienceCancelErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniExperienceCancelDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniExperienceCancelDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniExperienceCancelErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniExperienceCancelErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenMiniExperienceCancelDefaultResponse struct {
	value *AlipayOpenMiniExperienceCancelDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniExperienceCancelDefaultResponse) Get() *AlipayOpenMiniExperienceCancelDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniExperienceCancelDefaultResponse) Set(val *AlipayOpenMiniExperienceCancelDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniExperienceCancelDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniExperienceCancelDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniExperienceCancelDefaultResponse(val *AlipayOpenMiniExperienceCancelDefaultResponse) *NullableAlipayOpenMiniExperienceCancelDefaultResponse {
	return &NullableAlipayOpenMiniExperienceCancelDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniExperienceCancelDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniExperienceCancelDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
