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


// AlipayCommerceEcDepartmentCreateDefaultResponse struct for AlipayCommerceEcDepartmentCreateDefaultResponse
type AlipayCommerceEcDepartmentCreateDefaultResponse struct {
	AlipayCommerceEcDepartmentCreateErrorResponseModel *AlipayCommerceEcDepartmentCreateErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayCommerceEcDepartmentCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayCommerceEcDepartmentCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayCommerceEcDepartmentCreateErrorResponseModel);
	if err == nil {
		jsonAlipayCommerceEcDepartmentCreateErrorResponseModel, _ := json.Marshal(dst.AlipayCommerceEcDepartmentCreateErrorResponseModel)
		if string(jsonAlipayCommerceEcDepartmentCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayCommerceEcDepartmentCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayCommerceEcDepartmentCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayCommerceEcDepartmentCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayCommerceEcDepartmentCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayCommerceEcDepartmentCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayCommerceEcDepartmentCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayCommerceEcDepartmentCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayCommerceEcDepartmentCreateDefaultResponse struct {
	value *AlipayCommerceEcDepartmentCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayCommerceEcDepartmentCreateDefaultResponse) Get() *AlipayCommerceEcDepartmentCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayCommerceEcDepartmentCreateDefaultResponse) Set(val *AlipayCommerceEcDepartmentCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayCommerceEcDepartmentCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayCommerceEcDepartmentCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayCommerceEcDepartmentCreateDefaultResponse(val *AlipayCommerceEcDepartmentCreateDefaultResponse) *NullableAlipayCommerceEcDepartmentCreateDefaultResponse {
	return &NullableAlipayCommerceEcDepartmentCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayCommerceEcDepartmentCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayCommerceEcDepartmentCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


