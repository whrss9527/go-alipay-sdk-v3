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


// AlipayDataDataserviceAdPromotepageDownloadDefaultResponse struct for AlipayDataDataserviceAdPromotepageDownloadDefaultResponse
type AlipayDataDataserviceAdPromotepageDownloadDefaultResponse struct {
	AlipayDataDataserviceAdPromotepageDownloadErrorResponseModel *AlipayDataDataserviceAdPromotepageDownloadErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayDataDataserviceAdPromotepageDownloadDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayDataDataserviceAdPromotepageDownloadErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayDataDataserviceAdPromotepageDownloadErrorResponseModel);
	if err == nil {
		jsonAlipayDataDataserviceAdPromotepageDownloadErrorResponseModel, _ := json.Marshal(dst.AlipayDataDataserviceAdPromotepageDownloadErrorResponseModel)
		if string(jsonAlipayDataDataserviceAdPromotepageDownloadErrorResponseModel) == "{}" { // empty struct
			dst.AlipayDataDataserviceAdPromotepageDownloadErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayDataDataserviceAdPromotepageDownloadErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayDataDataserviceAdPromotepageDownloadErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayDataDataserviceAdPromotepageDownloadDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayDataDataserviceAdPromotepageDownloadDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayDataDataserviceAdPromotepageDownloadErrorResponseModel != nil {
		return json.Marshal(&src.AlipayDataDataserviceAdPromotepageDownloadErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayDataDataserviceAdPromotepageDownloadDefaultResponse struct {
	value *AlipayDataDataserviceAdPromotepageDownloadDefaultResponse
	isSet bool
}

func (v NullableAlipayDataDataserviceAdPromotepageDownloadDefaultResponse) Get() *AlipayDataDataserviceAdPromotepageDownloadDefaultResponse {
	return v.value
}

func (v *NullableAlipayDataDataserviceAdPromotepageDownloadDefaultResponse) Set(val *AlipayDataDataserviceAdPromotepageDownloadDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayDataDataserviceAdPromotepageDownloadDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayDataDataserviceAdPromotepageDownloadDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayDataDataserviceAdPromotepageDownloadDefaultResponse(val *AlipayDataDataserviceAdPromotepageDownloadDefaultResponse) *NullableAlipayDataDataserviceAdPromotepageDownloadDefaultResponse {
	return &NullableAlipayDataDataserviceAdPromotepageDownloadDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayDataDataserviceAdPromotepageDownloadDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayDataDataserviceAdPromotepageDownloadDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


