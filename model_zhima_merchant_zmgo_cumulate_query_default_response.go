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


// ZhimaMerchantZmgoCumulateQueryDefaultResponse struct for ZhimaMerchantZmgoCumulateQueryDefaultResponse
type ZhimaMerchantZmgoCumulateQueryDefaultResponse struct {
	CommonErrorType *CommonErrorType
	ZhimaMerchantZmgoCumulateQueryErrorResponseModel *ZhimaMerchantZmgoCumulateQueryErrorResponseModel
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ZhimaMerchantZmgoCumulateQueryDefaultResponse) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into ZhimaMerchantZmgoCumulateQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.ZhimaMerchantZmgoCumulateQueryErrorResponseModel);
	if err == nil {
		jsonZhimaMerchantZmgoCumulateQueryErrorResponseModel, _ := json.Marshal(dst.ZhimaMerchantZmgoCumulateQueryErrorResponseModel)
		if string(jsonZhimaMerchantZmgoCumulateQueryErrorResponseModel) == "{}" { // empty struct
			dst.ZhimaMerchantZmgoCumulateQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.ZhimaMerchantZmgoCumulateQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.ZhimaMerchantZmgoCumulateQueryErrorResponseModel = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ZhimaMerchantZmgoCumulateQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ZhimaMerchantZmgoCumulateQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	if src.ZhimaMerchantZmgoCumulateQueryErrorResponseModel != nil {
		return json.Marshal(&src.ZhimaMerchantZmgoCumulateQueryErrorResponseModel)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableZhimaMerchantZmgoCumulateQueryDefaultResponse struct {
	value *ZhimaMerchantZmgoCumulateQueryDefaultResponse
	isSet bool
}

func (v NullableZhimaMerchantZmgoCumulateQueryDefaultResponse) Get() *ZhimaMerchantZmgoCumulateQueryDefaultResponse {
	return v.value
}

func (v *NullableZhimaMerchantZmgoCumulateQueryDefaultResponse) Set(val *ZhimaMerchantZmgoCumulateQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableZhimaMerchantZmgoCumulateQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableZhimaMerchantZmgoCumulateQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZhimaMerchantZmgoCumulateQueryDefaultResponse(val *ZhimaMerchantZmgoCumulateQueryDefaultResponse) *NullableZhimaMerchantZmgoCumulateQueryDefaultResponse {
	return &NullableZhimaMerchantZmgoCumulateQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableZhimaMerchantZmgoCumulateQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZhimaMerchantZmgoCumulateQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


