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

// AlipayMarketingCardDeleteDefaultResponse struct for AlipayMarketingCardDeleteDefaultResponse
type AlipayMarketingCardDeleteDefaultResponse struct {
	AlipayMarketingCardDeleteErrorResponseModel *AlipayMarketingCardDeleteErrorResponseModel
	CommonErrorType                             *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingCardDeleteDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingCardDeleteErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingCardDeleteErrorResponseModel)
	if err == nil {
		jsonAlipayMarketingCardDeleteErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingCardDeleteErrorResponseModel)
		if string(jsonAlipayMarketingCardDeleteErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingCardDeleteErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingCardDeleteErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingCardDeleteErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingCardDeleteDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingCardDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingCardDeleteErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingCardDeleteErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAlipayMarketingCardDeleteDefaultResponse struct {
	value *AlipayMarketingCardDeleteDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingCardDeleteDefaultResponse) Get() *AlipayMarketingCardDeleteDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingCardDeleteDefaultResponse) Set(val *AlipayMarketingCardDeleteDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingCardDeleteDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingCardDeleteDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingCardDeleteDefaultResponse(val *AlipayMarketingCardDeleteDefaultResponse) *NullableAlipayMarketingCardDeleteDefaultResponse {
	return &NullableAlipayMarketingCardDeleteDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingCardDeleteDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingCardDeleteDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
