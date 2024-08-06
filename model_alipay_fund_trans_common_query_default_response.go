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


// AlipayFundTransCommonQueryDefaultResponse struct for AlipayFundTransCommonQueryDefaultResponse
type AlipayFundTransCommonQueryDefaultResponse struct {
	AlipayFundTransCommonQueryErrorResponseModel *AlipayFundTransCommonQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayFundTransCommonQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayFundTransCommonQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayFundTransCommonQueryErrorResponseModel);
	if err == nil {
		jsonAlipayFundTransCommonQueryErrorResponseModel, _ := json.Marshal(dst.AlipayFundTransCommonQueryErrorResponseModel)
		if string(jsonAlipayFundTransCommonQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayFundTransCommonQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayFundTransCommonQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayFundTransCommonQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayFundTransCommonQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayFundTransCommonQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayFundTransCommonQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayFundTransCommonQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayFundTransCommonQueryDefaultResponse struct {
	value *AlipayFundTransCommonQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayFundTransCommonQueryDefaultResponse) Get() *AlipayFundTransCommonQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayFundTransCommonQueryDefaultResponse) Set(val *AlipayFundTransCommonQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundTransCommonQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundTransCommonQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundTransCommonQueryDefaultResponse(val *AlipayFundTransCommonQueryDefaultResponse) *NullableAlipayFundTransCommonQueryDefaultResponse {
	return &NullableAlipayFundTransCommonQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayFundTransCommonQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundTransCommonQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

