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


// AlipayOpenSpImageUploadDefaultResponse struct for AlipayOpenSpImageUploadDefaultResponse
type AlipayOpenSpImageUploadDefaultResponse struct {
	AlipayOpenSpImageUploadErrorResponseModel *AlipayOpenSpImageUploadErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenSpImageUploadDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenSpImageUploadErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenSpImageUploadErrorResponseModel);
	if err == nil {
		jsonAlipayOpenSpImageUploadErrorResponseModel, _ := json.Marshal(dst.AlipayOpenSpImageUploadErrorResponseModel)
		if string(jsonAlipayOpenSpImageUploadErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenSpImageUploadErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenSpImageUploadErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenSpImageUploadErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenSpImageUploadDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenSpImageUploadDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenSpImageUploadErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenSpImageUploadErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenSpImageUploadDefaultResponse struct {
	value *AlipayOpenSpImageUploadDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenSpImageUploadDefaultResponse) Get() *AlipayOpenSpImageUploadDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenSpImageUploadDefaultResponse) Set(val *AlipayOpenSpImageUploadDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSpImageUploadDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSpImageUploadDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSpImageUploadDefaultResponse(val *AlipayOpenSpImageUploadDefaultResponse) *NullableAlipayOpenSpImageUploadDefaultResponse {
	return &NullableAlipayOpenSpImageUploadDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenSpImageUploadDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSpImageUploadDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


