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


// AntMerchantExpandInfoQueryDefaultResponse struct for AntMerchantExpandInfoQueryDefaultResponse
type AntMerchantExpandInfoQueryDefaultResponse struct {
	AntMerchantExpandInfoQueryErrorResponseModel *AntMerchantExpandInfoQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AntMerchantExpandInfoQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AntMerchantExpandInfoQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AntMerchantExpandInfoQueryErrorResponseModel);
	if err == nil {
		jsonAntMerchantExpandInfoQueryErrorResponseModel, _ := json.Marshal(dst.AntMerchantExpandInfoQueryErrorResponseModel)
		if string(jsonAntMerchantExpandInfoQueryErrorResponseModel) == "{}" { // empty struct
			dst.AntMerchantExpandInfoQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AntMerchantExpandInfoQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AntMerchantExpandInfoQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AntMerchantExpandInfoQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AntMerchantExpandInfoQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AntMerchantExpandInfoQueryErrorResponseModel != nil {
		return json.Marshal(&src.AntMerchantExpandInfoQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAntMerchantExpandInfoQueryDefaultResponse struct {
	value *AntMerchantExpandInfoQueryDefaultResponse
	isSet bool
}

func (v NullableAntMerchantExpandInfoQueryDefaultResponse) Get() *AntMerchantExpandInfoQueryDefaultResponse {
	return v.value
}

func (v *NullableAntMerchantExpandInfoQueryDefaultResponse) Set(val *AntMerchantExpandInfoQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandInfoQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandInfoQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandInfoQueryDefaultResponse(val *AntMerchantExpandInfoQueryDefaultResponse) *NullableAntMerchantExpandInfoQueryDefaultResponse {
	return &NullableAntMerchantExpandInfoQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAntMerchantExpandInfoQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandInfoQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

