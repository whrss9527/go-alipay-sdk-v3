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


// ZolozAuthenticationCustomerFtokenQueryDefaultResponse struct for ZolozAuthenticationCustomerFtokenQueryDefaultResponse
type ZolozAuthenticationCustomerFtokenQueryDefaultResponse struct {
	CommonErrorType *CommonErrorType
	ZolozAuthenticationCustomerFtokenQueryErrorResponseModel *ZolozAuthenticationCustomerFtokenQueryErrorResponseModel
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ZolozAuthenticationCustomerFtokenQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
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

	// try to unmarshal JSON data into ZolozAuthenticationCustomerFtokenQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.ZolozAuthenticationCustomerFtokenQueryErrorResponseModel);
	if err == nil {
		jsonZolozAuthenticationCustomerFtokenQueryErrorResponseModel, _ := json.Marshal(dst.ZolozAuthenticationCustomerFtokenQueryErrorResponseModel)
		if string(jsonZolozAuthenticationCustomerFtokenQueryErrorResponseModel) == "{}" { // empty struct
			dst.ZolozAuthenticationCustomerFtokenQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.ZolozAuthenticationCustomerFtokenQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.ZolozAuthenticationCustomerFtokenQueryErrorResponseModel = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ZolozAuthenticationCustomerFtokenQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ZolozAuthenticationCustomerFtokenQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	if src.ZolozAuthenticationCustomerFtokenQueryErrorResponseModel != nil {
		return json.Marshal(&src.ZolozAuthenticationCustomerFtokenQueryErrorResponseModel)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableZolozAuthenticationCustomerFtokenQueryDefaultResponse struct {
	value *ZolozAuthenticationCustomerFtokenQueryDefaultResponse
	isSet bool
}

func (v NullableZolozAuthenticationCustomerFtokenQueryDefaultResponse) Get() *ZolozAuthenticationCustomerFtokenQueryDefaultResponse {
	return v.value
}

func (v *NullableZolozAuthenticationCustomerFtokenQueryDefaultResponse) Set(val *ZolozAuthenticationCustomerFtokenQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableZolozAuthenticationCustomerFtokenQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableZolozAuthenticationCustomerFtokenQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZolozAuthenticationCustomerFtokenQueryDefaultResponse(val *ZolozAuthenticationCustomerFtokenQueryDefaultResponse) *NullableZolozAuthenticationCustomerFtokenQueryDefaultResponse {
	return &NullableZolozAuthenticationCustomerFtokenQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableZolozAuthenticationCustomerFtokenQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZolozAuthenticationCustomerFtokenQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


