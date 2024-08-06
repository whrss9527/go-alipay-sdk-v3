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


// AlipayOpenAppApiFieldApplyDefaultResponse struct for AlipayOpenAppApiFieldApplyDefaultResponse
type AlipayOpenAppApiFieldApplyDefaultResponse struct {
	AlipayOpenAppApiFieldApplyErrorResponseModel *AlipayOpenAppApiFieldApplyErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenAppApiFieldApplyDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenAppApiFieldApplyErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenAppApiFieldApplyErrorResponseModel);
	if err == nil {
		jsonAlipayOpenAppApiFieldApplyErrorResponseModel, _ := json.Marshal(dst.AlipayOpenAppApiFieldApplyErrorResponseModel)
		if string(jsonAlipayOpenAppApiFieldApplyErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenAppApiFieldApplyErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenAppApiFieldApplyErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenAppApiFieldApplyErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenAppApiFieldApplyDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenAppApiFieldApplyDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenAppApiFieldApplyErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenAppApiFieldApplyErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenAppApiFieldApplyDefaultResponse struct {
	value *AlipayOpenAppApiFieldApplyDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenAppApiFieldApplyDefaultResponse) Get() *AlipayOpenAppApiFieldApplyDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenAppApiFieldApplyDefaultResponse) Set(val *AlipayOpenAppApiFieldApplyDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenAppApiFieldApplyDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenAppApiFieldApplyDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenAppApiFieldApplyDefaultResponse(val *AlipayOpenAppApiFieldApplyDefaultResponse) *NullableAlipayOpenAppApiFieldApplyDefaultResponse {
	return &NullableAlipayOpenAppApiFieldApplyDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenAppApiFieldApplyDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenAppApiFieldApplyDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


