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


// ZhimaMerchantSubsidiariesCloseDefaultResponse struct for ZhimaMerchantSubsidiariesCloseDefaultResponse
type ZhimaMerchantSubsidiariesCloseDefaultResponse struct {
	CommonErrorType *CommonErrorType
	ZhimaMerchantSubsidiariesCloseErrorResponseModel *ZhimaMerchantSubsidiariesCloseErrorResponseModel
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ZhimaMerchantSubsidiariesCloseDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
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

	// try to unmarshal JSON data into ZhimaMerchantSubsidiariesCloseErrorResponseModel
	err = json.Unmarshal(data, &dst.ZhimaMerchantSubsidiariesCloseErrorResponseModel);
	if err == nil {
		jsonZhimaMerchantSubsidiariesCloseErrorResponseModel, _ := json.Marshal(dst.ZhimaMerchantSubsidiariesCloseErrorResponseModel)
		if string(jsonZhimaMerchantSubsidiariesCloseErrorResponseModel) == "{}" { // empty struct
			dst.ZhimaMerchantSubsidiariesCloseErrorResponseModel = nil
		} else {
			return nil // data stored in dst.ZhimaMerchantSubsidiariesCloseErrorResponseModel, return on the first match
		}
	} else {
		dst.ZhimaMerchantSubsidiariesCloseErrorResponseModel = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ZhimaMerchantSubsidiariesCloseDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ZhimaMerchantSubsidiariesCloseDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	if src.ZhimaMerchantSubsidiariesCloseErrorResponseModel != nil {
		return json.Marshal(&src.ZhimaMerchantSubsidiariesCloseErrorResponseModel)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableZhimaMerchantSubsidiariesCloseDefaultResponse struct {
	value *ZhimaMerchantSubsidiariesCloseDefaultResponse
	isSet bool
}

func (v NullableZhimaMerchantSubsidiariesCloseDefaultResponse) Get() *ZhimaMerchantSubsidiariesCloseDefaultResponse {
	return v.value
}

func (v *NullableZhimaMerchantSubsidiariesCloseDefaultResponse) Set(val *ZhimaMerchantSubsidiariesCloseDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableZhimaMerchantSubsidiariesCloseDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableZhimaMerchantSubsidiariesCloseDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZhimaMerchantSubsidiariesCloseDefaultResponse(val *ZhimaMerchantSubsidiariesCloseDefaultResponse) *NullableZhimaMerchantSubsidiariesCloseDefaultResponse {
	return &NullableZhimaMerchantSubsidiariesCloseDefaultResponse{value: val, isSet: true}
}

func (v NullableZhimaMerchantSubsidiariesCloseDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZhimaMerchantSubsidiariesCloseDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

