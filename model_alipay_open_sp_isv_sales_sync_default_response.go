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

// AlipayOpenSpIsvSalesSyncDefaultResponse struct for AlipayOpenSpIsvSalesSyncDefaultResponse
type AlipayOpenSpIsvSalesSyncDefaultResponse struct {
	AlipayOpenSpIsvSalesSyncErrorResponseModel *AlipayOpenSpIsvSalesSyncErrorResponseModel
	CommonErrorType                            *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenSpIsvSalesSyncDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenSpIsvSalesSyncErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenSpIsvSalesSyncErrorResponseModel)
	if err == nil {
		jsonAlipayOpenSpIsvSalesSyncErrorResponseModel, _ := json.Marshal(dst.AlipayOpenSpIsvSalesSyncErrorResponseModel)
		if string(jsonAlipayOpenSpIsvSalesSyncErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenSpIsvSalesSyncErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenSpIsvSalesSyncErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenSpIsvSalesSyncErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenSpIsvSalesSyncDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenSpIsvSalesSyncDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenSpIsvSalesSyncErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenSpIsvSalesSyncErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayOpenSpIsvSalesSyncDefaultResponse struct {
	value *AlipayOpenSpIsvSalesSyncDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenSpIsvSalesSyncDefaultResponse) Get() *AlipayOpenSpIsvSalesSyncDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenSpIsvSalesSyncDefaultResponse) Set(val *AlipayOpenSpIsvSalesSyncDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSpIsvSalesSyncDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSpIsvSalesSyncDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSpIsvSalesSyncDefaultResponse(val *AlipayOpenSpIsvSalesSyncDefaultResponse) *NullableAlipayOpenSpIsvSalesSyncDefaultResponse {
	return &NullableAlipayOpenSpIsvSalesSyncDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenSpIsvSalesSyncDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSpIsvSalesSyncDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
