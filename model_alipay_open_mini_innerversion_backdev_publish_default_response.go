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


// AlipayOpenMiniInnerversionBackdevPublishDefaultResponse struct for AlipayOpenMiniInnerversionBackdevPublishDefaultResponse
type AlipayOpenMiniInnerversionBackdevPublishDefaultResponse struct {
	AlipayOpenMiniInnerversionBackdevPublishErrorResponseModel *AlipayOpenMiniInnerversionBackdevPublishErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniInnerversionBackdevPublishDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniInnerversionBackdevPublishErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniInnerversionBackdevPublishErrorResponseModel);
	if err == nil {
		jsonAlipayOpenMiniInnerversionBackdevPublishErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniInnerversionBackdevPublishErrorResponseModel)
		if string(jsonAlipayOpenMiniInnerversionBackdevPublishErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniInnerversionBackdevPublishErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniInnerversionBackdevPublishErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniInnerversionBackdevPublishErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniInnerversionBackdevPublishDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniInnerversionBackdevPublishDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniInnerversionBackdevPublishErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniInnerversionBackdevPublishErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenMiniInnerversionBackdevPublishDefaultResponse struct {
	value *AlipayOpenMiniInnerversionBackdevPublishDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniInnerversionBackdevPublishDefaultResponse) Get() *AlipayOpenMiniInnerversionBackdevPublishDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerversionBackdevPublishDefaultResponse) Set(val *AlipayOpenMiniInnerversionBackdevPublishDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerversionBackdevPublishDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerversionBackdevPublishDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerversionBackdevPublishDefaultResponse(val *AlipayOpenMiniInnerversionBackdevPublishDefaultResponse) *NullableAlipayOpenMiniInnerversionBackdevPublishDefaultResponse {
	return &NullableAlipayOpenMiniInnerversionBackdevPublishDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerversionBackdevPublishDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerversionBackdevPublishDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


