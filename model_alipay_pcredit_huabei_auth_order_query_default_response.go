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


// AlipayPcreditHuabeiAuthOrderQueryDefaultResponse struct for AlipayPcreditHuabeiAuthOrderQueryDefaultResponse
type AlipayPcreditHuabeiAuthOrderQueryDefaultResponse struct {
	AlipayPcreditHuabeiAuthOrderQueryErrorResponseModel *AlipayPcreditHuabeiAuthOrderQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayPcreditHuabeiAuthOrderQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayPcreditHuabeiAuthOrderQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayPcreditHuabeiAuthOrderQueryErrorResponseModel);
	if err == nil {
		jsonAlipayPcreditHuabeiAuthOrderQueryErrorResponseModel, _ := json.Marshal(dst.AlipayPcreditHuabeiAuthOrderQueryErrorResponseModel)
		if string(jsonAlipayPcreditHuabeiAuthOrderQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayPcreditHuabeiAuthOrderQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayPcreditHuabeiAuthOrderQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayPcreditHuabeiAuthOrderQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayPcreditHuabeiAuthOrderQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayPcreditHuabeiAuthOrderQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayPcreditHuabeiAuthOrderQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayPcreditHuabeiAuthOrderQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayPcreditHuabeiAuthOrderQueryDefaultResponse struct {
	value *AlipayPcreditHuabeiAuthOrderQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayPcreditHuabeiAuthOrderQueryDefaultResponse) Get() *AlipayPcreditHuabeiAuthOrderQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayPcreditHuabeiAuthOrderQueryDefaultResponse) Set(val *AlipayPcreditHuabeiAuthOrderQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayPcreditHuabeiAuthOrderQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayPcreditHuabeiAuthOrderQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayPcreditHuabeiAuthOrderQueryDefaultResponse(val *AlipayPcreditHuabeiAuthOrderQueryDefaultResponse) *NullableAlipayPcreditHuabeiAuthOrderQueryDefaultResponse {
	return &NullableAlipayPcreditHuabeiAuthOrderQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayPcreditHuabeiAuthOrderQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayPcreditHuabeiAuthOrderQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

