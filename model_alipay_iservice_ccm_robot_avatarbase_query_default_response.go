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


// AlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse struct for AlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse
type AlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse struct {
	AlipayIserviceCcmRobotAvatarbaseQueryErrorResponseModel *AlipayIserviceCcmRobotAvatarbaseQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayIserviceCcmRobotAvatarbaseQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayIserviceCcmRobotAvatarbaseQueryErrorResponseModel);
	if err == nil {
		jsonAlipayIserviceCcmRobotAvatarbaseQueryErrorResponseModel, _ := json.Marshal(dst.AlipayIserviceCcmRobotAvatarbaseQueryErrorResponseModel)
		if string(jsonAlipayIserviceCcmRobotAvatarbaseQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayIserviceCcmRobotAvatarbaseQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayIserviceCcmRobotAvatarbaseQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayIserviceCcmRobotAvatarbaseQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayIserviceCcmRobotAvatarbaseQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayIserviceCcmRobotAvatarbaseQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse struct {
	value *AlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse) Get() *AlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse) Set(val *AlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse(val *AlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse) *NullableAlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse {
	return &NullableAlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayIserviceCcmRobotAvatarbaseQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


