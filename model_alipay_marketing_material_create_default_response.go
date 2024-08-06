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

// AlipayMarketingMaterialCreateDefaultResponse struct for AlipayMarketingMaterialCreateDefaultResponse
type AlipayMarketingMaterialCreateDefaultResponse struct {
	AlipayMarketingMaterialCreateErrorResponseModel *AlipayMarketingMaterialCreateErrorResponseModel
	CommonErrorType                                 *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingMaterialCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingMaterialCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingMaterialCreateErrorResponseModel)
	if err == nil {
		jsonAlipayMarketingMaterialCreateErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingMaterialCreateErrorResponseModel)
		if string(jsonAlipayMarketingMaterialCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingMaterialCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingMaterialCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingMaterialCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingMaterialCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingMaterialCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingMaterialCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingMaterialCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayMarketingMaterialCreateDefaultResponse struct {
	value *AlipayMarketingMaterialCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingMaterialCreateDefaultResponse) Get() *AlipayMarketingMaterialCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingMaterialCreateDefaultResponse) Set(val *AlipayMarketingMaterialCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingMaterialCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingMaterialCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingMaterialCreateDefaultResponse(val *AlipayMarketingMaterialCreateDefaultResponse) *NullableAlipayMarketingMaterialCreateDefaultResponse {
	return &NullableAlipayMarketingMaterialCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingMaterialCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingMaterialCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
