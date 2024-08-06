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


// AlipayOpenPublicDefaultExtensionCreateDefaultResponse struct for AlipayOpenPublicDefaultExtensionCreateDefaultResponse
type AlipayOpenPublicDefaultExtensionCreateDefaultResponse struct {
	AlipayOpenPublicDefaultExtensionCreateErrorResponseModel *AlipayOpenPublicDefaultExtensionCreateErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicDefaultExtensionCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicDefaultExtensionCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicDefaultExtensionCreateErrorResponseModel);
	if err == nil {
		jsonAlipayOpenPublicDefaultExtensionCreateErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicDefaultExtensionCreateErrorResponseModel)
		if string(jsonAlipayOpenPublicDefaultExtensionCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicDefaultExtensionCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicDefaultExtensionCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicDefaultExtensionCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicDefaultExtensionCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicDefaultExtensionCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicDefaultExtensionCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicDefaultExtensionCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenPublicDefaultExtensionCreateDefaultResponse struct {
	value *AlipayOpenPublicDefaultExtensionCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicDefaultExtensionCreateDefaultResponse) Get() *AlipayOpenPublicDefaultExtensionCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicDefaultExtensionCreateDefaultResponse) Set(val *AlipayOpenPublicDefaultExtensionCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicDefaultExtensionCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicDefaultExtensionCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicDefaultExtensionCreateDefaultResponse(val *AlipayOpenPublicDefaultExtensionCreateDefaultResponse) *NullableAlipayOpenPublicDefaultExtensionCreateDefaultResponse {
	return &NullableAlipayOpenPublicDefaultExtensionCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicDefaultExtensionCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicDefaultExtensionCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


