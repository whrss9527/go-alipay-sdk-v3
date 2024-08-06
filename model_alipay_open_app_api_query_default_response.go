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


// AlipayOpenAppApiQueryDefaultResponse struct for AlipayOpenAppApiQueryDefaultResponse
type AlipayOpenAppApiQueryDefaultResponse struct {
	AlipayOpenAppApiQueryErrorResponseModel *AlipayOpenAppApiQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenAppApiQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenAppApiQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenAppApiQueryErrorResponseModel);
	if err == nil {
		jsonAlipayOpenAppApiQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenAppApiQueryErrorResponseModel)
		if string(jsonAlipayOpenAppApiQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenAppApiQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenAppApiQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenAppApiQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenAppApiQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenAppApiQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenAppApiQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenAppApiQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenAppApiQueryDefaultResponse struct {
	value *AlipayOpenAppApiQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenAppApiQueryDefaultResponse) Get() *AlipayOpenAppApiQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenAppApiQueryDefaultResponse) Set(val *AlipayOpenAppApiQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAppApiQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAppApiQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAppApiQueryDefaultResponse(val *AlipayOpenAppApiQueryDefaultResponse) *NullableAlipayOpenAppApiQueryDefaultResponse {
	return &NullableAlipayOpenAppApiQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenAppApiQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAppApiQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


