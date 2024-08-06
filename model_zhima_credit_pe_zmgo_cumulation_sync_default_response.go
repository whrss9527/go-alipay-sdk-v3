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

// ZhimaCreditPeZmgoCumulationSyncDefaultResponse struct for ZhimaCreditPeZmgoCumulationSyncDefaultResponse
type ZhimaCreditPeZmgoCumulationSyncDefaultResponse struct {
	CommonErrorType                                   *CommonErrorType
	ZhimaCreditPeZmgoCumulationSyncErrorResponseModel *ZhimaCreditPeZmgoCumulationSyncErrorResponseModel
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ZhimaCreditPeZmgoCumulationSyncDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
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

	// try to unmarshal JSON data into ZhimaCreditPeZmgoCumulationSyncErrorResponseModel
	err = json.Unmarshal(data, &dst.ZhimaCreditPeZmgoCumulationSyncErrorResponseModel)
	if err == nil {
		jsonZhimaCreditPeZmgoCumulationSyncErrorResponseModel, _ := json.Marshal(dst.ZhimaCreditPeZmgoCumulationSyncErrorResponseModel)
		if string(jsonZhimaCreditPeZmgoCumulationSyncErrorResponseModel) == "{}" { // empty struct
			dst.ZhimaCreditPeZmgoCumulationSyncErrorResponseModel = nil
		} else {
			return nil // data stored in dst.ZhimaCreditPeZmgoCumulationSyncErrorResponseModel, return on the first match
		}
	} else {
		dst.ZhimaCreditPeZmgoCumulationSyncErrorResponseModel = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ZhimaCreditPeZmgoCumulationSyncDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ZhimaCreditPeZmgoCumulationSyncDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	if src.ZhimaCreditPeZmgoCumulationSyncErrorResponseModel != nil {
		return json.Marshal(&src.ZhimaCreditPeZmgoCumulationSyncErrorResponseModel)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableZhimaCreditPeZmgoCumulationSyncDefaultResponse struct {
	value *ZhimaCreditPeZmgoCumulationSyncDefaultResponse
	isSet bool
}

func (v NullableZhimaCreditPeZmgoCumulationSyncDefaultResponse) Get() *ZhimaCreditPeZmgoCumulationSyncDefaultResponse {
	return v.value
}

func (v *NullableZhimaCreditPeZmgoCumulationSyncDefaultResponse) Set(val *ZhimaCreditPeZmgoCumulationSyncDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableZhimaCreditPeZmgoCumulationSyncDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableZhimaCreditPeZmgoCumulationSyncDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZhimaCreditPeZmgoCumulationSyncDefaultResponse(val *ZhimaCreditPeZmgoCumulationSyncDefaultResponse) *NullableZhimaCreditPeZmgoCumulationSyncDefaultResponse {
	return &NullableZhimaCreditPeZmgoCumulationSyncDefaultResponse{value: val, isSet: true}
}

func (v NullableZhimaCreditPeZmgoCumulationSyncDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZhimaCreditPeZmgoCumulationSyncDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
