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


// AlipayOpenMiniIsvFastregisterQueryDefaultResponse struct for AlipayOpenMiniIsvFastregisterQueryDefaultResponse
type AlipayOpenMiniIsvFastregisterQueryDefaultResponse struct {
	AlipayOpenMiniIsvFastregisterQueryErrorResponseModel *AlipayOpenMiniIsvFastregisterQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenMiniIsvFastregisterQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenMiniIsvFastregisterQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenMiniIsvFastregisterQueryErrorResponseModel);
	if err == nil {
		jsonAlipayOpenMiniIsvFastregisterQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenMiniIsvFastregisterQueryErrorResponseModel)
		if string(jsonAlipayOpenMiniIsvFastregisterQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenMiniIsvFastregisterQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenMiniIsvFastregisterQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenMiniIsvFastregisterQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenMiniIsvFastregisterQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenMiniIsvFastregisterQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenMiniIsvFastregisterQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenMiniIsvFastregisterQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenMiniIsvFastregisterQueryDefaultResponse struct {
	value *AlipayOpenMiniIsvFastregisterQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenMiniIsvFastregisterQueryDefaultResponse) Get() *AlipayOpenMiniIsvFastregisterQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenMiniIsvFastregisterQueryDefaultResponse) Set(val *AlipayOpenMiniIsvFastregisterQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenMiniIsvFastregisterQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenMiniIsvFastregisterQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenMiniIsvFastregisterQueryDefaultResponse(val *AlipayOpenMiniIsvFastregisterQueryDefaultResponse) *NullableAlipayOpenMiniIsvFastregisterQueryDefaultResponse {
	return &NullableAlipayOpenMiniIsvFastregisterQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenMiniIsvFastregisterQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenMiniIsvFastregisterQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


