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

// AlipayCommerceEcEnterpriseInfoQueryDefaultResponse struct for AlipayCommerceEcEnterpriseInfoQueryDefaultResponse
type AlipayCommerceEcEnterpriseInfoQueryDefaultResponse struct {
	AlipayCommerceEcEnterpriseInfoQueryErrorResponseModel *AlipayCommerceEcEnterpriseInfoQueryErrorResponseModel
	CommonErrorType                                       *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayCommerceEcEnterpriseInfoQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayCommerceEcEnterpriseInfoQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayCommerceEcEnterpriseInfoQueryErrorResponseModel)
	if err == nil {
		jsonAlipayCommerceEcEnterpriseInfoQueryErrorResponseModel, _ := json.Marshal(dst.AlipayCommerceEcEnterpriseInfoQueryErrorResponseModel)
		if string(jsonAlipayCommerceEcEnterpriseInfoQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayCommerceEcEnterpriseInfoQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayCommerceEcEnterpriseInfoQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayCommerceEcEnterpriseInfoQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayCommerceEcEnterpriseInfoQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayCommerceEcEnterpriseInfoQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayCommerceEcEnterpriseInfoQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayCommerceEcEnterpriseInfoQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayCommerceEcEnterpriseInfoQueryDefaultResponse struct {
	value *AlipayCommerceEcEnterpriseInfoQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayCommerceEcEnterpriseInfoQueryDefaultResponse) Get() *AlipayCommerceEcEnterpriseInfoQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayCommerceEcEnterpriseInfoQueryDefaultResponse) Set(val *AlipayCommerceEcEnterpriseInfoQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEcEnterpriseInfoQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEcEnterpriseInfoQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEcEnterpriseInfoQueryDefaultResponse(val *AlipayCommerceEcEnterpriseInfoQueryDefaultResponse) *NullableAlipayCommerceEcEnterpriseInfoQueryDefaultResponse {
	return &NullableAlipayCommerceEcEnterpriseInfoQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayCommerceEcEnterpriseInfoQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEcEnterpriseInfoQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
