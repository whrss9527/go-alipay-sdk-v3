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


// AlipayOpenMiniVersionListQueryDefaultResponse struct for AlipayOpenMiniVersionListQueryDefaultResponse
type AlipayOpenMiniVersionListQueryDefaultResponse struct {
	AlipayOpenMiniVersionListQueryErrorResponseModel *AlipayOpenMiniVersionListQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniVersionListQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniVersionListQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniVersionListQueryErrorResponseModel);
	if err == nil {
		jsonAlipayOpenMiniVersionListQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniVersionListQueryErrorResponseModel)
		if string(jsonAlipayOpenMiniVersionListQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniVersionListQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniVersionListQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniVersionListQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniVersionListQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniVersionListQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniVersionListQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniVersionListQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenMiniVersionListQueryDefaultResponse struct {
	value *AlipayOpenMiniVersionListQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniVersionListQueryDefaultResponse) Get() *AlipayOpenMiniVersionListQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniVersionListQueryDefaultResponse) Set(val *AlipayOpenMiniVersionListQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniVersionListQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniVersionListQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniVersionListQueryDefaultResponse(val *AlipayOpenMiniVersionListQueryDefaultResponse) *NullableAlipayOpenMiniVersionListQueryDefaultResponse {
	return &NullableAlipayOpenMiniVersionListQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniVersionListQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniVersionListQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


