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


// AlipaySecurityProdSssQueryDefaultResponse struct for AlipaySecurityProdSssQueryDefaultResponse
type AlipaySecurityProdSssQueryDefaultResponse struct {
	AlipaySecurityProdSssQueryErrorResponseModel *AlipaySecurityProdSssQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipaySecurityProdSssQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipaySecurityProdSssQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipaySecurityProdSssQueryErrorResponseModel);
	if err == nil {
		jsonAlipaySecurityProdSssQueryErrorResponseModel, _ := json.Marshal(dst.AlipaySecurityProdSssQueryErrorResponseModel)
		if string(jsonAlipaySecurityProdSssQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipaySecurityProdSssQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipaySecurityProdSssQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipaySecurityProdSssQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipaySecurityProdSssQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipaySecurityProdSssQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipaySecurityProdSssQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipaySecurityProdSssQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipaySecurityProdSssQueryDefaultResponse struct {
	value *AlipaySecurityProdSssQueryDefaultResponse
	isSet bool
}

func (v NullableAlipaySecurityProdSssQueryDefaultResponse) Get() *AlipaySecurityProdSssQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipaySecurityProdSssQueryDefaultResponse) Set(val *AlipaySecurityProdSssQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipaySecurityProdSssQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipaySecurityProdSssQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipaySecurityProdSssQueryDefaultResponse(val *AlipaySecurityProdSssQueryDefaultResponse) *NullableAlipaySecurityProdSssQueryDefaultResponse {
	return &NullableAlipaySecurityProdSssQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipaySecurityProdSssQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipaySecurityProdSssQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


