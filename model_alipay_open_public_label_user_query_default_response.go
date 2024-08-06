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

// AlipayOpenPublicLabelUserQueryDefaultResponse struct for AlipayOpenPublicLabelUserQueryDefaultResponse
type AlipayOpenPublicLabelUserQueryDefaultResponse struct {
	AlipayOpenPublicLabelUserQueryErrorResponseModel *AlipayOpenPublicLabelUserQueryErrorResponseModel
	CommonErrorType                                  *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicLabelUserQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicLabelUserQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicLabelUserQueryErrorResponseModel)
	if err == nil {
		jsonAlipayOpenPublicLabelUserQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicLabelUserQueryErrorResponseModel)
		if string(jsonAlipayOpenPublicLabelUserQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicLabelUserQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicLabelUserQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicLabelUserQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicLabelUserQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicLabelUserQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicLabelUserQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicLabelUserQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenPublicLabelUserQueryDefaultResponse struct {
	value *AlipayOpenPublicLabelUserQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicLabelUserQueryDefaultResponse) Get() *AlipayOpenPublicLabelUserQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicLabelUserQueryDefaultResponse) Set(val *AlipayOpenPublicLabelUserQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicLabelUserQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicLabelUserQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicLabelUserQueryDefaultResponse(val *AlipayOpenPublicLabelUserQueryDefaultResponse) *NullableAlipayOpenPublicLabelUserQueryDefaultResponse {
	return &NullableAlipayOpenPublicLabelUserQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicLabelUserQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicLabelUserQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
