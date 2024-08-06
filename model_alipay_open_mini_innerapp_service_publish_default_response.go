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


// AlipayOpenMiniInnerappServicePublishDefaultResponse struct for AlipayOpenMiniInnerappServicePublishDefaultResponse
type AlipayOpenMiniInnerappServicePublishDefaultResponse struct {
	AlipayOpenMiniInnerappServicePublishErrorResponseModel *AlipayOpenMiniInnerappServicePublishErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniInnerappServicePublishDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniInnerappServicePublishErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniInnerappServicePublishErrorResponseModel);
	if err == nil {
		jsonAlipayOpenMiniInnerappServicePublishErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniInnerappServicePublishErrorResponseModel)
		if string(jsonAlipayOpenMiniInnerappServicePublishErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniInnerappServicePublishErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniInnerappServicePublishErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniInnerappServicePublishErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniInnerappServicePublishDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniInnerappServicePublishDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniInnerappServicePublishErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniInnerappServicePublishErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenMiniInnerappServicePublishDefaultResponse struct {
	value *AlipayOpenMiniInnerappServicePublishDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniInnerappServicePublishDefaultResponse) Get() *AlipayOpenMiniInnerappServicePublishDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniInnerappServicePublishDefaultResponse) Set(val *AlipayOpenMiniInnerappServicePublishDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniInnerappServicePublishDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniInnerappServicePublishDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniInnerappServicePublishDefaultResponse(val *AlipayOpenMiniInnerappServicePublishDefaultResponse) *NullableAlipayOpenMiniInnerappServicePublishDefaultResponse {
	return &NullableAlipayOpenMiniInnerappServicePublishDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniInnerappServicePublishDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniInnerappServicePublishDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


