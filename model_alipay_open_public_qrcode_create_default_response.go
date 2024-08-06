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


// AlipayOpenPublicQrcodeCreateDefaultResponse struct for AlipayOpenPublicQrcodeCreateDefaultResponse
type AlipayOpenPublicQrcodeCreateDefaultResponse struct {
	AlipayOpenPublicQrcodeCreateErrorResponseModel *AlipayOpenPublicQrcodeCreateErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicQrcodeCreateDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicQrcodeCreateErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicQrcodeCreateErrorResponseModel);
	if err == nil {
		jsonAlipayOpenPublicQrcodeCreateErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicQrcodeCreateErrorResponseModel)
		if string(jsonAlipayOpenPublicQrcodeCreateErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicQrcodeCreateErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicQrcodeCreateErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicQrcodeCreateErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicQrcodeCreateDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicQrcodeCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicQrcodeCreateErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicQrcodeCreateErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenPublicQrcodeCreateDefaultResponse struct {
	value *AlipayOpenPublicQrcodeCreateDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicQrcodeCreateDefaultResponse) Get() *AlipayOpenPublicQrcodeCreateDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicQrcodeCreateDefaultResponse) Set(val *AlipayOpenPublicQrcodeCreateDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicQrcodeCreateDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicQrcodeCreateDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicQrcodeCreateDefaultResponse(val *AlipayOpenPublicQrcodeCreateDefaultResponse) *NullableAlipayOpenPublicQrcodeCreateDefaultResponse {
	return &NullableAlipayOpenPublicQrcodeCreateDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicQrcodeCreateDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicQrcodeCreateDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


