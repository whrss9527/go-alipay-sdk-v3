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


// AlipayOpenPublicPayeeBindDeleteDefaultResponse struct for AlipayOpenPublicPayeeBindDeleteDefaultResponse
type AlipayOpenPublicPayeeBindDeleteDefaultResponse struct {
	AlipayOpenPublicPayeeBindDeleteErrorResponseModel *AlipayOpenPublicPayeeBindDeleteErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicPayeeBindDeleteDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicPayeeBindDeleteErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicPayeeBindDeleteErrorResponseModel);
	if err == nil {
		jsonAlipayOpenPublicPayeeBindDeleteErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicPayeeBindDeleteErrorResponseModel)
		if string(jsonAlipayOpenPublicPayeeBindDeleteErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicPayeeBindDeleteErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicPayeeBindDeleteErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicPayeeBindDeleteErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicPayeeBindDeleteDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicPayeeBindDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicPayeeBindDeleteErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicPayeeBindDeleteErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenPublicPayeeBindDeleteDefaultResponse struct {
	value *AlipayOpenPublicPayeeBindDeleteDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicPayeeBindDeleteDefaultResponse) Get() *AlipayOpenPublicPayeeBindDeleteDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicPayeeBindDeleteDefaultResponse) Set(val *AlipayOpenPublicPayeeBindDeleteDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicPayeeBindDeleteDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicPayeeBindDeleteDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicPayeeBindDeleteDefaultResponse(val *AlipayOpenPublicPayeeBindDeleteDefaultResponse) *NullableAlipayOpenPublicPayeeBindDeleteDefaultResponse {
	return &NullableAlipayOpenPublicPayeeBindDeleteDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicPayeeBindDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicPayeeBindDeleteDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


