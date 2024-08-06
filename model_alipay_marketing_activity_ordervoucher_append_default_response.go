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


// AlipayMarketingActivityOrdervoucherAppendDefaultResponse struct for AlipayMarketingActivityOrdervoucherAppendDefaultResponse
type AlipayMarketingActivityOrdervoucherAppendDefaultResponse struct {
	AlipayMarketingActivityOrdervoucherAppendErrorResponseModel *AlipayMarketingActivityOrdervoucherAppendErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingActivityOrdervoucherAppendDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingActivityOrdervoucherAppendErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingActivityOrdervoucherAppendErrorResponseModel);
	if err == nil {
		jsonAlipayMarketingActivityOrdervoucherAppendErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingActivityOrdervoucherAppendErrorResponseModel)
		if string(jsonAlipayMarketingActivityOrdervoucherAppendErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingActivityOrdervoucherAppendErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingActivityOrdervoucherAppendErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingActivityOrdervoucherAppendErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingActivityOrdervoucherAppendDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingActivityOrdervoucherAppendDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingActivityOrdervoucherAppendErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingActivityOrdervoucherAppendErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMarketingActivityOrdervoucherAppendDefaultResponse struct {
	value *AlipayMarketingActivityOrdervoucherAppendDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingActivityOrdervoucherAppendDefaultResponse) Get() *AlipayMarketingActivityOrdervoucherAppendDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingActivityOrdervoucherAppendDefaultResponse) Set(val *AlipayMarketingActivityOrdervoucherAppendDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityOrdervoucherAppendDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityOrdervoucherAppendDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityOrdervoucherAppendDefaultResponse(val *AlipayMarketingActivityOrdervoucherAppendDefaultResponse) *NullableAlipayMarketingActivityOrdervoucherAppendDefaultResponse {
	return &NullableAlipayMarketingActivityOrdervoucherAppendDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityOrdervoucherAppendDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityOrdervoucherAppendDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


