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


// AlipayMobilePublicShortlinkCreateDefaultResponse struct for AlipayMobilePublicShortlinkCreateDefaultResponse
type AlipayMobilePublicShortlinkCreateDefaultResponse struct {
	AlipayMobilePublicShortlinkCreateErrorResponseModel *AlipayMobilePublicShortlinkCreateErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMobilePublicShortlinkCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMobilePublicShortlinkCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMobilePublicShortlinkCreateErrorResponseModel);
	if err == nil {
		jsonAlipayMobilePublicShortlinkCreateErrorResponseModel, _ := json.Marshal(dst.AlipayMobilePublicShortlinkCreateErrorResponseModel)
		if string(jsonAlipayMobilePublicShortlinkCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMobilePublicShortlinkCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMobilePublicShortlinkCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMobilePublicShortlinkCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMobilePublicShortlinkCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMobilePublicShortlinkCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMobilePublicShortlinkCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMobilePublicShortlinkCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMobilePublicShortlinkCreateDefaultResponse struct {
	value *AlipayMobilePublicShortlinkCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayMobilePublicShortlinkCreateDefaultResponse) Get() *AlipayMobilePublicShortlinkCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayMobilePublicShortlinkCreateDefaultResponse) Set(val *AlipayMobilePublicShortlinkCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMobilePublicShortlinkCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMobilePublicShortlinkCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMobilePublicShortlinkCreateDefaultResponse(val *AlipayMobilePublicShortlinkCreateDefaultResponse) *NullableAlipayMobilePublicShortlinkCreateDefaultResponse {
	return &NullableAlipayMobilePublicShortlinkCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMobilePublicShortlinkCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMobilePublicShortlinkCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


