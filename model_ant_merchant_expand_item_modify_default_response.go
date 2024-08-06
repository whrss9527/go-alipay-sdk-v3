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

// AntMerchantExpandItemModifyDefaultResponse struct for AntMerchantExpandItemModifyDefaultResponse
type AntMerchantExpandItemModifyDefaultResponse struct {
	AntMerchantExpandItemModifyErrorResponseModel *AntMerchantExpandItemModifyErrorResponseModel
	CommonErrorType                               *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AntMerchantExpandItemModifyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AntMerchantExpandItemModifyErrorResponseModel
	err = json.Unmarshal(data, &dst.AntMerchantExpandItemModifyErrorResponseModel)
	if err == nil {
		jsonAntMerchantExpandItemModifyErrorResponseModel, _ := json.Marshal(dst.AntMerchantExpandItemModifyErrorResponseModel)
		if string(jsonAntMerchantExpandItemModifyErrorResponseModel) == "{}" { // empty struct
			dst.AntMerchantExpandItemModifyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AntMerchantExpandItemModifyErrorResponseModel, return on the first match
		}
	} else {
		dst.AntMerchantExpandItemModifyErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AntMerchantExpandItemModifyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AntMerchantExpandItemModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AntMerchantExpandItemModifyErrorResponseModel != nil {
		return json.Marshal(&src.AntMerchantExpandItemModifyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAntMerchantExpandItemModifyDefaultResponse struct {
	value *AntMerchantExpandItemModifyDefaultResponse
	isSet bool
}

func (v NullableAntMerchantExpandItemModifyDefaultResponse) Get() *AntMerchantExpandItemModifyDefaultResponse {
	return v.value
}

func (v *NullableAntMerchantExpandItemModifyDefaultResponse) Set(val *AntMerchantExpandItemModifyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandItemModifyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandItemModifyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandItemModifyDefaultResponse(val *AntMerchantExpandItemModifyDefaultResponse) *NullableAntMerchantExpandItemModifyDefaultResponse {
	return &NullableAntMerchantExpandItemModifyDefaultResponse{value: val, isSet: true}
}

func (v NullableAntMerchantExpandItemModifyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandItemModifyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
