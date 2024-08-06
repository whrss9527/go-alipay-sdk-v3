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


// AlipayIserviceCcmInstanceGetDefaultResponse struct for AlipayIserviceCcmInstanceGetDefaultResponse
type AlipayIserviceCcmInstanceGetDefaultResponse struct {
	AlipayIserviceCcmInstanceGetErrorResponseModel *AlipayIserviceCcmInstanceGetErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayIserviceCcmInstanceGetDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayIserviceCcmInstanceGetErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayIserviceCcmInstanceGetErrorResponseModel);
	if err == nil {
		jsonAlipayIserviceCcmInstanceGetErrorResponseModel, _ := json.Marshal(dst.AlipayIserviceCcmInstanceGetErrorResponseModel)
		if string(jsonAlipayIserviceCcmInstanceGetErrorResponseModel) == "{}" { // empty struct
			dst.AlipayIserviceCcmInstanceGetErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayIserviceCcmInstanceGetErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayIserviceCcmInstanceGetErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayIserviceCcmInstanceGetDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayIserviceCcmInstanceGetDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayIserviceCcmInstanceGetErrorResponseModel != nil {
		return json.Marshal(&src.AlipayIserviceCcmInstanceGetErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayIserviceCcmInstanceGetDefaultResponse struct {
	value *AlipayIserviceCcmInstanceGetDefaultResponse
	isSet bool
}

func (v NullableAlipayIserviceCcmInstanceGetDefaultResponse) Get() *AlipayIserviceCcmInstanceGetDefaultResponse {
	return v.value
}

func (v *NullableAlipayIserviceCcmInstanceGetDefaultResponse) Set(val *AlipayIserviceCcmInstanceGetDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmInstanceGetDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmInstanceGetDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmInstanceGetDefaultResponse(val *AlipayIserviceCcmInstanceGetDefaultResponse) *NullableAlipayIserviceCcmInstanceGetDefaultResponse {
	return &NullableAlipayIserviceCcmInstanceGetDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmInstanceGetDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmInstanceGetDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


