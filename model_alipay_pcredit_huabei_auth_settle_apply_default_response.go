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

// AlipayPcreditHuabeiAuthSettleApplyDefaultResponse struct for AlipayPcreditHuabeiAuthSettleApplyDefaultResponse
type AlipayPcreditHuabeiAuthSettleApplyDefaultResponse struct {
	AlipayPcreditHuabeiAuthSettleApplyErrorResponseModel *AlipayPcreditHuabeiAuthSettleApplyErrorResponseModel
	CommonErrorType                                      *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayPcreditHuabeiAuthSettleApplyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayPcreditHuabeiAuthSettleApplyErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayPcreditHuabeiAuthSettleApplyErrorResponseModel)
	if err == nil {
		jsonAlipayPcreditHuabeiAuthSettleApplyErrorResponseModel, _ := json.Marshal(dst.AlipayPcreditHuabeiAuthSettleApplyErrorResponseModel)
		if string(jsonAlipayPcreditHuabeiAuthSettleApplyErrorResponseModel) == "{}" { // empty struct
			dst.AlipayPcreditHuabeiAuthSettleApplyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayPcreditHuabeiAuthSettleApplyErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayPcreditHuabeiAuthSettleApplyErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayPcreditHuabeiAuthSettleApplyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayPcreditHuabeiAuthSettleApplyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayPcreditHuabeiAuthSettleApplyErrorResponseModel != nil {
		return json.Marshal(&src.AlipayPcreditHuabeiAuthSettleApplyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayPcreditHuabeiAuthSettleApplyDefaultResponse struct {
	value *AlipayPcreditHuabeiAuthSettleApplyDefaultResponse
	isSet bool
}

func (v NullableAlipayPcreditHuabeiAuthSettleApplyDefaultResponse) Get() *AlipayPcreditHuabeiAuthSettleApplyDefaultResponse {
	return v.value
}

func (v *NullableAlipayPcreditHuabeiAuthSettleApplyDefaultResponse) Set(val *AlipayPcreditHuabeiAuthSettleApplyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayPcreditHuabeiAuthSettleApplyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayPcreditHuabeiAuthSettleApplyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayPcreditHuabeiAuthSettleApplyDefaultResponse(val *AlipayPcreditHuabeiAuthSettleApplyDefaultResponse) *NullableAlipayPcreditHuabeiAuthSettleApplyDefaultResponse {
	return &NullableAlipayPcreditHuabeiAuthSettleApplyDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayPcreditHuabeiAuthSettleApplyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayPcreditHuabeiAuthSettleApplyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
