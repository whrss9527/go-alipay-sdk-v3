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


// AlipayOpenSpOpporPageQueryDefaultResponse struct for AlipayOpenSpOpporPageQueryDefaultResponse
type AlipayOpenSpOpporPageQueryDefaultResponse struct {
	AlipayOpenSpOpporPageQueryErrorResponseModel *AlipayOpenSpOpporPageQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenSpOpporPageQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenSpOpporPageQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenSpOpporPageQueryErrorResponseModel);
	if err == nil {
		jsonAlipayOpenSpOpporPageQueryErrorResponseModel, _ := json.Marshal(dst.AlipayOpenSpOpporPageQueryErrorResponseModel)
		if string(jsonAlipayOpenSpOpporPageQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenSpOpporPageQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenSpOpporPageQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenSpOpporPageQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenSpOpporPageQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenSpOpporPageQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenSpOpporPageQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenSpOpporPageQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenSpOpporPageQueryDefaultResponse struct {
	value *AlipayOpenSpOpporPageQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenSpOpporPageQueryDefaultResponse) Get() *AlipayOpenSpOpporPageQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenSpOpporPageQueryDefaultResponse) Set(val *AlipayOpenSpOpporPageQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSpOpporPageQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSpOpporPageQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSpOpporPageQueryDefaultResponse(val *AlipayOpenSpOpporPageQueryDefaultResponse) *NullableAlipayOpenSpOpporPageQueryDefaultResponse {
	return &NullableAlipayOpenSpOpporPageQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenSpOpporPageQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSpOpporPageQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


