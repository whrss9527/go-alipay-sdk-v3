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


// AlipayOpenPublicLifeMsgSendDefaultResponse struct for AlipayOpenPublicLifeMsgSendDefaultResponse
type AlipayOpenPublicLifeMsgSendDefaultResponse struct {
	AlipayOpenPublicLifeMsgSendErrorResponseModel *AlipayOpenPublicLifeMsgSendErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayOpenPublicLifeMsgSendDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayOpenPublicLifeMsgSendErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayOpenPublicLifeMsgSendErrorResponseModel);
	if err == nil {
		jsonAlipayOpenPublicLifeMsgSendErrorResponseModel, _ := json.Marshal(dst.AlipayOpenPublicLifeMsgSendErrorResponseModel)
		if string(jsonAlipayOpenPublicLifeMsgSendErrorResponseModel) == "{}" { // empty struct
			dst.AlipayOpenPublicLifeMsgSendErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayOpenPublicLifeMsgSendErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayOpenPublicLifeMsgSendErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayOpenPublicLifeMsgSendDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayOpenPublicLifeMsgSendDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayOpenPublicLifeMsgSendErrorResponseModel != nil {
		return json.Marshal(&src.AlipayOpenPublicLifeMsgSendErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayOpenPublicLifeMsgSendDefaultResponse struct {
	value *AlipayOpenPublicLifeMsgSendDefaultResponse
	isSet bool
}

func (v NullableAlipayOpenPublicLifeMsgSendDefaultResponse) Get() *AlipayOpenPublicLifeMsgSendDefaultResponse {
	return v.value
}

func (v *NullableAlipayOpenPublicLifeMsgSendDefaultResponse) Set(val *AlipayOpenPublicLifeMsgSendDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayOpenPublicLifeMsgSendDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayOpenPublicLifeMsgSendDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayOpenPublicLifeMsgSendDefaultResponse(val *AlipayOpenPublicLifeMsgSendDefaultResponse) *NullableAlipayOpenPublicLifeMsgSendDefaultResponse {
	return &NullableAlipayOpenPublicLifeMsgSendDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayOpenPublicLifeMsgSendDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayOpenPublicLifeMsgSendDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


