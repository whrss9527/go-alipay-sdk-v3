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


// AlipayOpenSpBlueseaactivityQueryDefaultResponse struct for AlipayOpenSpBlueseaactivityQueryDefaultResponse
type AlipayOpenSpBlueseaactivityQueryDefaultResponse struct {
	AlipayOpenSpBlueseaactivityQueryErrorResponseModel *AlipayOpenSpBlueseaactivityQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenSpBlueseaactivityQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenSpBlueseaactivityQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenSpBlueseaactivityQueryErrorResponseModel);
	if err == nil {
		jsonAlipayOpenSpBlueseaactivityQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenSpBlueseaactivityQueryErrorResponseModel)
		if string(jsonAlipayOpenSpBlueseaactivityQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenSpBlueseaactivityQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenSpBlueseaactivityQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenSpBlueseaactivityQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenSpBlueseaactivityQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenSpBlueseaactivityQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenSpBlueseaactivityQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenSpBlueseaactivityQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenSpBlueseaactivityQueryDefaultResponse struct {
	value *AlipayOpenSpBlueseaactivityQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenSpBlueseaactivityQueryDefaultResponse) Get() *AlipayOpenSpBlueseaactivityQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenSpBlueseaactivityQueryDefaultResponse) Set(val *AlipayOpenSpBlueseaactivityQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSpBlueseaactivityQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSpBlueseaactivityQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSpBlueseaactivityQueryDefaultResponse(val *AlipayOpenSpBlueseaactivityQueryDefaultResponse) *NullableAlipayOpenSpBlueseaactivityQueryDefaultResponse {
	return &NullableAlipayOpenSpBlueseaactivityQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenSpBlueseaactivityQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSpBlueseaactivityQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


