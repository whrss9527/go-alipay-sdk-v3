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


// AlipayOpenPublicAccountDeleteDefaultResponse struct for AlipayOpenPublicAccountDeleteDefaultResponse
type AlipayOpenPublicAccountDeleteDefaultResponse struct {
	AlipayOpenPublicAccountDeleteErrorResponseModel *AlipayOpenPublicAccountDeleteErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicAccountDeleteDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicAccountDeleteErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicAccountDeleteErrorResponseModel);
	if err == nil {
		jsonAlipayOpenPublicAccountDeleteErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicAccountDeleteErrorResponseModel)
		if string(jsonAlipayOpenPublicAccountDeleteErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicAccountDeleteErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicAccountDeleteErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicAccountDeleteErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicAccountDeleteDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicAccountDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicAccountDeleteErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicAccountDeleteErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenPublicAccountDeleteDefaultResponse struct {
	value *AlipayOpenPublicAccountDeleteDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicAccountDeleteDefaultResponse) Get() *AlipayOpenPublicAccountDeleteDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicAccountDeleteDefaultResponse) Set(val *AlipayOpenPublicAccountDeleteDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicAccountDeleteDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicAccountDeleteDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicAccountDeleteDefaultResponse(val *AlipayOpenPublicAccountDeleteDefaultResponse) *NullableAlipayOpenPublicAccountDeleteDefaultResponse {
	return &NullableAlipayOpenPublicAccountDeleteDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicAccountDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicAccountDeleteDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


