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


// AlipayOpenPublicLabelQueryDefaultResponse struct for AlipayOpenPublicLabelQueryDefaultResponse
type AlipayOpenPublicLabelQueryDefaultResponse struct {
	AlipayOpenPublicLabelQueryErrorResponseModel *AlipayOpenPublicLabelQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicLabelQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicLabelQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicLabelQueryErrorResponseModel);
	if err == nil {
		jsonAlipayOpenPublicLabelQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicLabelQueryErrorResponseModel)
		if string(jsonAlipayOpenPublicLabelQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicLabelQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicLabelQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicLabelQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicLabelQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicLabelQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicLabelQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicLabelQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenPublicLabelQueryDefaultResponse struct {
	value *AlipayOpenPublicLabelQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicLabelQueryDefaultResponse) Get() *AlipayOpenPublicLabelQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicLabelQueryDefaultResponse) Set(val *AlipayOpenPublicLabelQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicLabelQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicLabelQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicLabelQueryDefaultResponse(val *AlipayOpenPublicLabelQueryDefaultResponse) *NullableAlipayOpenPublicLabelQueryDefaultResponse {
	return &NullableAlipayOpenPublicLabelQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicLabelQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicLabelQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


