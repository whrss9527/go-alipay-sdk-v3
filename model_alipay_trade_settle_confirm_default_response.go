/*
支付宝开放平台API

支付宝开放平台v3协议文档

API version: 2024-07-05
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alipay

import (
	"encoding/json"
	"fmt"
)

// AlipayTradeSettleConfirmDefaultResponse struct for AlipayTradeSettleConfirmDefaultResponse
type AlipayTradeSettleConfirmDefaultResponse struct {
	AlipayTradeSettleConfirmErrorResponseModel *AlipayTradeSettleConfirmErrorResponseModel
	CommonErrorType                            *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayTradeSettleConfirmDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayTradeSettleConfirmErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayTradeSettleConfirmErrorResponseModel)
	if err == nil {
		jsonAlipayTradeSettleConfirmErrorResponseModel, _ := json.Marshal(dst.AlipayTradeSettleConfirmErrorResponseModel)
		if string(jsonAlipayTradeSettleConfirmErrorResponseModel) == "{}" { // empty struct
			dst.AlipayTradeSettleConfirmErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayTradeSettleConfirmErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayTradeSettleConfirmErrorResponseModel = nil
	}

	// try to unmarshal JSON data into CommonErrorType
	err = json.Unmarshal(data, &dst.CommonErrorType)
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayTradeSettleConfirmDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayTradeSettleConfirmDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayTradeSettleConfirmErrorResponseModel != nil {
		return json.Marshal(&src.AlipayTradeSettleConfirmErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayTradeSettleConfirmDefaultResponse struct {
	value *AlipayTradeSettleConfirmDefaultResponse
	isSet bool
}

func (v NullableAlipayTradeSettleConfirmDefaultResponse) Get() *AlipayTradeSettleConfirmDefaultResponse {
	return v.value
}

func (v *NullableAlipayTradeSettleConfirmDefaultResponse) Set(val *AlipayTradeSettleConfirmDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayTradeSettleConfirmDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayTradeSettleConfirmDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayTradeSettleConfirmDefaultResponse(val *AlipayTradeSettleConfirmDefaultResponse) *NullableAlipayTradeSettleConfirmDefaultResponse {
	return &NullableAlipayTradeSettleConfirmDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayTradeSettleConfirmDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayTradeSettleConfirmDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
