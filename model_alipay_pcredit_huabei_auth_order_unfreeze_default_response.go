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

// AlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse struct for AlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse
type AlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse struct {
	AlipayPcreditHuabeiAuthOrderUnfreezeErrorResponseModel *AlipayPcreditHuabeiAuthOrderUnfreezeErrorResponseModel
	CommonErrorType                                        *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayPcreditHuabeiAuthOrderUnfreezeErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayPcreditHuabeiAuthOrderUnfreezeErrorResponseModel)
	if err == nil {
		jsonAlipayPcreditHuabeiAuthOrderUnfreezeErrorResponseModel, _ := json.Marshal(dst.AlipayPcreditHuabeiAuthOrderUnfreezeErrorResponseModel)
		if string(jsonAlipayPcreditHuabeiAuthOrderUnfreezeErrorResponseModel) == "{}" { // empty struct
			dst.AlipayPcreditHuabeiAuthOrderUnfreezeErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayPcreditHuabeiAuthOrderUnfreezeErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayPcreditHuabeiAuthOrderUnfreezeErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayPcreditHuabeiAuthOrderUnfreezeErrorResponseModel != nil {
		return json.Marshal(&src.AlipayPcreditHuabeiAuthOrderUnfreezeErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse struct {
	value *AlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse
	isSet bool
}

func (v NullableAlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse) Get() *AlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse {
	return v.value
}

func (v *NullableAlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse) Set(val *AlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse(val *AlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse) *NullableAlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse {
	return &NullableAlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayPcreditHuabeiAuthOrderUnfreezeDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
