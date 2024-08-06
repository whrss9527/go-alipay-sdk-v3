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

// AlipayOpenMiniInnerversionGrayPublishDefaultResponse struct for AlipayOpenMiniInnerversionGrayPublishDefaultResponse
type AlipayOpenMiniInnerversionGrayPublishDefaultResponse struct {
	AlipayOpenMiniInnerversionGrayPublishErrorResponseModel *AlipayOpenMiniInnerversionGrayPublishErrorResponseModel
	CommonErrorType                                         *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniInnerversionGrayPublishDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniInnerversionGrayPublishErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniInnerversionGrayPublishErrorResponseModel)
	if err == nil {
		jsonAlipayOpenMiniInnerversionGrayPublishErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniInnerversionGrayPublishErrorResponseModel)
		if string(jsonAlipayOpenMiniInnerversionGrayPublishErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniInnerversionGrayPublishErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniInnerversionGrayPublishErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniInnerversionGrayPublishErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniInnerversionGrayPublishDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniInnerversionGrayPublishDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniInnerversionGrayPublishErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniInnerversionGrayPublishErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenMiniInnerversionGrayPublishDefaultResponse struct {
	value *AlipayOpenMiniInnerversionGrayPublishDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniInnerversionGrayPublishDefaultResponse) Get() *AlipayOpenMiniInnerversionGrayPublishDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerversionGrayPublishDefaultResponse) Set(val *AlipayOpenMiniInnerversionGrayPublishDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerversionGrayPublishDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerversionGrayPublishDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerversionGrayPublishDefaultResponse(val *AlipayOpenMiniInnerversionGrayPublishDefaultResponse) *NullableAlipayOpenMiniInnerversionGrayPublishDefaultResponse {
	return &NullableAlipayOpenMiniInnerversionGrayPublishDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerversionGrayPublishDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerversionGrayPublishDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
