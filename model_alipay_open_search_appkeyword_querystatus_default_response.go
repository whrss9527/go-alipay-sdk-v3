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


// AlipayOpenSearchAppkeywordQuerystatusDefaultResponse struct for AlipayOpenSearchAppkeywordQuerystatusDefaultResponse
type AlipayOpenSearchAppkeywordQuerystatusDefaultResponse struct {
	AlipayOpenSearchAppkeywordQuerystatusErrorResponseModel *AlipayOpenSearchAppkeywordQuerystatusErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenSearchAppkeywordQuerystatusDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenSearchAppkeywordQuerystatusErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenSearchAppkeywordQuerystatusErrorResponseModel);
	if err == nil {
		jsonAlipayOpenSearchAppkeywordQuerystatusErrorResponseModel, _ := json.Marshal(dst.AlipayOpenSearchAppkeywordQuerystatusErrorResponseModel)
		if string(jsonAlipayOpenSearchAppkeywordQuerystatusErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenSearchAppkeywordQuerystatusErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenSearchAppkeywordQuerystatusErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenSearchAppkeywordQuerystatusErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenSearchAppkeywordQuerystatusDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenSearchAppkeywordQuerystatusDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenSearchAppkeywordQuerystatusErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenSearchAppkeywordQuerystatusErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenSearchAppkeywordQuerystatusDefaultResponse struct {
	value *AlipayOpenSearchAppkeywordQuerystatusDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenSearchAppkeywordQuerystatusDefaultResponse) Get() *AlipayOpenSearchAppkeywordQuerystatusDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenSearchAppkeywordQuerystatusDefaultResponse) Set(val *AlipayOpenSearchAppkeywordQuerystatusDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenSearchAppkeywordQuerystatusDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenSearchAppkeywordQuerystatusDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenSearchAppkeywordQuerystatusDefaultResponse(val *AlipayOpenSearchAppkeywordQuerystatusDefaultResponse) *NullableAlipayOpenSearchAppkeywordQuerystatusDefaultResponse {
	return &NullableAlipayOpenSearchAppkeywordQuerystatusDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenSearchAppkeywordQuerystatusDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenSearchAppkeywordQuerystatusDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


