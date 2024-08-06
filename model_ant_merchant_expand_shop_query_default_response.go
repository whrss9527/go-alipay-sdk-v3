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

// AntMerchantExpandShopQueryDefaultResponse struct for AntMerchantExpandShopQueryDefaultResponse
type AntMerchantExpandShopQueryDefaultResponse struct {
	AntMerchantExpandShopQueryErrorResponseModel *AntMerchantExpandShopQueryErrorResponseModel
	CommonErrorType                              *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AntMerchantExpandShopQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AntMerchantExpandShopQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AntMerchantExpandShopQueryErrorResponseModel)
	if err == nil {
		jsonAntMerchantExpandShopQueryErrorResponseModel, _ := json.Marshal(dst.AntMerchantExpandShopQueryErrorResponseModel)
		if string(jsonAntMerchantExpandShopQueryErrorResponseModel) == "{}" { // empty struct
			dst.AntMerchantExpandShopQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AntMerchantExpandShopQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AntMerchantExpandShopQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AntMerchantExpandShopQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AntMerchantExpandShopQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AntMerchantExpandShopQueryErrorResponseModel != nil {
		return json.Marshal(&src.AntMerchantExpandShopQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAntMerchantExpandShopQueryDefaultResponse struct {
	value *AntMerchantExpandShopQueryDefaultResponse
	isSet bool
}

func (v NullableAntMerchantExpandShopQueryDefaultResponse) Get() *AntMerchantExpandShopQueryDefaultResponse {
	return v.value
}

func (v *NullableAntMerchantExpandShopQueryDefaultResponse) Set(val *AntMerchantExpandShopQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandShopQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandShopQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandShopQueryDefaultResponse(val *AntMerchantExpandShopQueryDefaultResponse) *NullableAntMerchantExpandShopQueryDefaultResponse {
	return &NullableAntMerchantExpandShopQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAntMerchantExpandShopQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandShopQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
