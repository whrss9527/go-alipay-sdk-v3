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


// AlipayOpenSearchAbilityBatchqueryDefaultResponse struct for AlipayOpenSearchAbilityBatchqueryDefaultResponse
type AlipayOpenSearchAbilityBatchqueryDefaultResponse struct {
	AlipayOpenSearchAbilityBatchqueryErrorResponseModel *AlipayOpenSearchAbilityBatchqueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenSearchAbilityBatchqueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenSearchAbilityBatchqueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenSearchAbilityBatchqueryErrorResponseModel);
	if err == nil {
		jsonAlipayOpenSearchAbilityBatchqueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenSearchAbilityBatchqueryErrorResponseModel)
		if string(jsonAlipayOpenSearchAbilityBatchqueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenSearchAbilityBatchqueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenSearchAbilityBatchqueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenSearchAbilityBatchqueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenSearchAbilityBatchqueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenSearchAbilityBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenSearchAbilityBatchqueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenSearchAbilityBatchqueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenSearchAbilityBatchqueryDefaultResponse struct {
	value *AlipayOpenSearchAbilityBatchqueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenSearchAbilityBatchqueryDefaultResponse) Get() *AlipayOpenSearchAbilityBatchqueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenSearchAbilityBatchqueryDefaultResponse) Set(val *AlipayOpenSearchAbilityBatchqueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchAbilityBatchqueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchAbilityBatchqueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchAbilityBatchqueryDefaultResponse(val *AlipayOpenSearchAbilityBatchqueryDefaultResponse) *NullableAlipayOpenSearchAbilityBatchqueryDefaultResponse {
	return &NullableAlipayOpenSearchAbilityBatchqueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchAbilityBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchAbilityBatchqueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

