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

// AlipayFundUsertradeBatchqueryDefaultResponse struct for AlipayFundUsertradeBatchqueryDefaultResponse
type AlipayFundUsertradeBatchqueryDefaultResponse struct {
	AlipayFundUsertradeBatchqueryErrorResponseModel *AlipayFundUsertradeBatchqueryErrorResponseModel
	CommonErrorType                                 *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayFundUsertradeBatchqueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayFundUsertradeBatchqueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayFundUsertradeBatchqueryErrorResponseModel)
	if err == nil {
		jsonAlipayFundUsertradeBatchqueryErrorResponseModel, _ := json.Marshal(dst.AlipayFundUsertradeBatchqueryErrorResponseModel)
		if string(jsonAlipayFundUsertradeBatchqueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayFundUsertradeBatchqueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayFundUsertradeBatchqueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayFundUsertradeBatchqueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayFundUsertradeBatchqueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayFundUsertradeBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayFundUsertradeBatchqueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayFundUsertradeBatchqueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayFundUsertradeBatchqueryDefaultResponse struct {
	value *AlipayFundUsertradeBatchqueryDefaultResponse
	isSet bool
}

func (v NullableAlipayFundUsertradeBatchqueryDefaultResponse) Get() *AlipayFundUsertradeBatchqueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayFundUsertradeBatchqueryDefaultResponse) Set(val *AlipayFundUsertradeBatchqueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayFundUsertradeBatchqueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayFundUsertradeBatchqueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayFundUsertradeBatchqueryDefaultResponse(val *AlipayFundUsertradeBatchqueryDefaultResponse) *NullableAlipayFundUsertradeBatchqueryDefaultResponse {
	return &NullableAlipayFundUsertradeBatchqueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayFundUsertradeBatchqueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayFundUsertradeBatchqueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
