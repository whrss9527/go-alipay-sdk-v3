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


// AntMerchantExpandShopConsultDefaultResponse struct for AntMerchantExpandShopConsultDefaultResponse
type AntMerchantExpandShopConsultDefaultResponse struct {
	AntMerchantExpandShopConsultErrorResponseModel *AntMerchantExpandShopConsultErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AntMerchantExpandShopConsultDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AntMerchantExpandShopConsultErrorResponseModel
	err = json.Unmarshal(data, &dst.AntMerchantExpandShopConsultErrorResponseModel);
	if err == nil {
		jsonAntMerchantExpandShopConsultErrorResponseModel, _ := json.Marshal(dst.AntMerchantExpandShopConsultErrorResponseModel)
		if string(jsonAntMerchantExpandShopConsultErrorResponseModel) == "{}" { // empty struct
			dst.AntMerchantExpandShopConsultErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AntMerchantExpandShopConsultErrorResponseModel, return on the first match
		}
	} else {
		dst.AntMerchantExpandShopConsultErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AntMerchantExpandShopConsultDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AntMerchantExpandShopConsultDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AntMerchantExpandShopConsultErrorResponseModel != nil {
		return json.Marshal(&src.AntMerchantExpandShopConsultErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAntMerchantExpandShopConsultDefaultResponse struct {
	value *AntMerchantExpandShopConsultDefaultResponse
	isSet bool
}

func (v NullableAntMerchantExpandShopConsultDefaultResponse) Get() *AntMerchantExpandShopConsultDefaultResponse {
	return v.value
}

func (v *NullableAntMerchantExpandShopConsultDefaultResponse) Set(val *AntMerchantExpandShopConsultDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAntMerchantExpandShopConsultDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAntMerchantExpandShopConsultDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAntMerchantExpandShopConsultDefaultResponse(val *AntMerchantExpandShopConsultDefaultResponse) *NullableAntMerchantExpandShopConsultDefaultResponse {
	return &NullableAntMerchantExpandShopConsultDefaultResponse{value: val, isSet: true}
}

func (v NullableAntMerchantExpandShopConsultDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAntMerchantExpandShopConsultDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

