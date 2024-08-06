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


// AlipayMobilePublicMessageCustomSendDefaultResponse struct for AlipayMobilePublicMessageCustomSendDefaultResponse
type AlipayMobilePublicMessageCustomSendDefaultResponse struct {
	AlipayMobilePublicMessageCustomSendErrorResponseModel *AlipayMobilePublicMessageCustomSendErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMobilePublicMessageCustomSendDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMobilePublicMessageCustomSendErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMobilePublicMessageCustomSendErrorResponseModel);
	if err == nil {
		jsonAlipayMobilePublicMessageCustomSendErrorResponseModel, _ := json.Marshal(dst.AlipayMobilePublicMessageCustomSendErrorResponseModel)
		if string(jsonAlipayMobilePublicMessageCustomSendErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMobilePublicMessageCustomSendErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMobilePublicMessageCustomSendErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMobilePublicMessageCustomSendErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMobilePublicMessageCustomSendDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMobilePublicMessageCustomSendDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMobilePublicMessageCustomSendErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMobilePublicMessageCustomSendErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMobilePublicMessageCustomSendDefaultResponse struct {
	value *AlipayMobilePublicMessageCustomSendDefaultResponse
	isSet bool
}

func (v NullableAlipayMobilePublicMessageCustomSendDefaultResponse) Get() *AlipayMobilePublicMessageCustomSendDefaultResponse {
	return v.value
}

func (v *NullableAlipayMobilePublicMessageCustomSendDefaultResponse) Set(val *AlipayMobilePublicMessageCustomSendDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMobilePublicMessageCustomSendDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMobilePublicMessageCustomSendDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMobilePublicMessageCustomSendDefaultResponse(val *AlipayMobilePublicMessageCustomSendDefaultResponse) *NullableAlipayMobilePublicMessageCustomSendDefaultResponse {
	return &NullableAlipayMobilePublicMessageCustomSendDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMobilePublicMessageCustomSendDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMobilePublicMessageCustomSendDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


