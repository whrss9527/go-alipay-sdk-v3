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

// AlipayOpenPublicLifeLabelCreateDefaultResponse struct for AlipayOpenPublicLifeLabelCreateDefaultResponse
type AlipayOpenPublicLifeLabelCreateDefaultResponse struct {
	AlipayOpenPublicLifeLabelCreateErrorResponseModel *AlipayOpenPublicLifeLabelCreateErrorResponseModel
	CommonErrorType                                   *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicLifeLabelCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicLifeLabelCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicLifeLabelCreateErrorResponseModel)
	if err == nil {
		jsonAlipayOpenPublicLifeLabelCreateErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicLifeLabelCreateErrorResponseModel)
		if string(jsonAlipayOpenPublicLifeLabelCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicLifeLabelCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicLifeLabelCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicLifeLabelCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicLifeLabelCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicLifeLabelCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicLifeLabelCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicLifeLabelCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenPublicLifeLabelCreateDefaultResponse struct {
	value *AlipayOpenPublicLifeLabelCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicLifeLabelCreateDefaultResponse) Get() *AlipayOpenPublicLifeLabelCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicLifeLabelCreateDefaultResponse) Set(val *AlipayOpenPublicLifeLabelCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicLifeLabelCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicLifeLabelCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicLifeLabelCreateDefaultResponse(val *AlipayOpenPublicLifeLabelCreateDefaultResponse) *NullableAlipayOpenPublicLifeLabelCreateDefaultResponse {
	return &NullableAlipayOpenPublicLifeLabelCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicLifeLabelCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicLifeLabelCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
