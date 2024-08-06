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


// AlipayOpenSearchAppkeywordBatchqueryDefaultResponse struct for AlipayOpenSearchAppkeywordBatchqueryDefaultResponse
type AlipayOpenSearchAppkeywordBatchqueryDefaultResponse struct {
	AlipayOpenSearchAppkeywordBatchqueryErrorResponseModel *AlipayOpenSearchAppkeywordBatchqueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenSearchAppkeywordBatchqueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenSearchAppkeywordBatchqueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenSearchAppkeywordBatchqueryErrorResponseModel);
	if err == nil {
		jsonAlipayOpenSearchAppkeywordBatchqueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenSearchAppkeywordBatchqueryErrorResponseModel)
		if string(jsonAlipayOpenSearchAppkeywordBatchqueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenSearchAppkeywordBatchqueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenSearchAppkeywordBatchqueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenSearchAppkeywordBatchqueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenSearchAppkeywordBatchqueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenSearchAppkeywordBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenSearchAppkeywordBatchqueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenSearchAppkeywordBatchqueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenSearchAppkeywordBatchqueryDefaultResponse struct {
	value *AlipayOpenSearchAppkeywordBatchqueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenSearchAppkeywordBatchqueryDefaultResponse) Get() *AlipayOpenSearchAppkeywordBatchqueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenSearchAppkeywordBatchqueryDefaultResponse) Set(val *AlipayOpenSearchAppkeywordBatchqueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchAppkeywordBatchqueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchAppkeywordBatchqueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchAppkeywordBatchqueryDefaultResponse(val *AlipayOpenSearchAppkeywordBatchqueryDefaultResponse) *NullableAlipayOpenSearchAppkeywordBatchqueryDefaultResponse {
	return &NullableAlipayOpenSearchAppkeywordBatchqueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchAppkeywordBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchAppkeywordBatchqueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


