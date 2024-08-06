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

// AntMerchantExpandApprecommendAvailableQueryDefaultResponse struct for AntMerchantExpandApprecommendAvailableQueryDefaultResponse
type AntMerchantExpandApprecommendAvailableQueryDefaultResponse struct {
	AntMerchantExpandApprecommendAvailableQueryErrorResponseModel *AntMerchantExpandApprecommendAvailableQueryErrorResponseModel
	CommonErrorType                                               *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AntMerchantExpandApprecommendAvailableQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AntMerchantExpandApprecommendAvailableQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AntMerchantExpandApprecommendAvailableQueryErrorResponseModel)
	if err == nil {
		jsonAntMerchantExpandApprecommendAvailableQueryErrorResponseModel, _ := json.Marshal(dst.AntMerchantExpandApprecommendAvailableQueryErrorResponseModel)
		if string(jsonAntMerchantExpandApprecommendAvailableQueryErrorResponseModel) == "{}" { // empty struct
			dst.AntMerchantExpandApprecommendAvailableQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AntMerchantExpandApprecommendAvailableQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AntMerchantExpandApprecommendAvailableQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AntMerchantExpandApprecommendAvailableQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AntMerchantExpandApprecommendAvailableQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AntMerchantExpandApprecommendAvailableQueryErrorResponseModel != nil {
		return json.Marshal(&src.AntMerchantExpandApprecommendAvailableQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAntMerchantExpandApprecommendAvailableQueryDefaultResponse struct {
	value *AntMerchantExpandApprecommendAvailableQueryDefaultResponse
	isSet bool
}

func (v NullableAntMerchantExpandApprecommendAvailableQueryDefaultResponse) Get() *AntMerchantExpandApprecommendAvailableQueryDefaultResponse {
	return v.value
}

func (v *NullableAntMerchantExpandApprecommendAvailableQueryDefaultResponse) Set(val *AntMerchantExpandApprecommendAvailableQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandApprecommendAvailableQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandApprecommendAvailableQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandApprecommendAvailableQueryDefaultResponse(val *AntMerchantExpandApprecommendAvailableQueryDefaultResponse) *NullableAntMerchantExpandApprecommendAvailableQueryDefaultResponse {
	return &NullableAntMerchantExpandApprecommendAvailableQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAntMerchantExpandApprecommendAvailableQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandApprecommendAvailableQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
