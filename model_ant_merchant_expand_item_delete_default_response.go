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

// AntMerchantExpandItemDeleteDefaultResponse struct for AntMerchantExpandItemDeleteDefaultResponse
type AntMerchantExpandItemDeleteDefaultResponse struct {
	AntMerchantExpandItemDeleteErrorResponseModel *AntMerchantExpandItemDeleteErrorResponseModel
	CommonErrorType                               *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AntMerchantExpandItemDeleteDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AntMerchantExpandItemDeleteErrorResponseModel
	err = json.Unmarshal(data, &dst.AntMerchantExpandItemDeleteErrorResponseModel)
	if err == nil {
		jsonAntMerchantExpandItemDeleteErrorResponseModel, _ := json.Marshal(dst.AntMerchantExpandItemDeleteErrorResponseModel)
		if string(jsonAntMerchantExpandItemDeleteErrorResponseModel) == "{}" { // empty struct
			dst.AntMerchantExpandItemDeleteErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AntMerchantExpandItemDeleteErrorResponseModel, return on the first match
		}
	} else {
		dst.AntMerchantExpandItemDeleteErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AntMerchantExpandItemDeleteDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AntMerchantExpandItemDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AntMerchantExpandItemDeleteErrorResponseModel != nil {
		return json.Marshal(&src.AntMerchantExpandItemDeleteErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAntMerchantExpandItemDeleteDefaultResponse struct {
	value *AntMerchantExpandItemDeleteDefaultResponse
	isSet bool
}

func (v NullableAntMerchantExpandItemDeleteDefaultResponse) Get() *AntMerchantExpandItemDeleteDefaultResponse {
	return v.value
}

func (v *NullableAntMerchantExpandItemDeleteDefaultResponse) Set(val *AntMerchantExpandItemDeleteDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandItemDeleteDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandItemDeleteDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandItemDeleteDefaultResponse(val *AntMerchantExpandItemDeleteDefaultResponse) *NullableAntMerchantExpandItemDeleteDefaultResponse {
	return &NullableAntMerchantExpandItemDeleteDefaultResponse{value: val, isSet: true}
}

func (v NullableAntMerchantExpandItemDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandItemDeleteDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
