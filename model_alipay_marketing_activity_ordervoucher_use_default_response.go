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


// AlipayMarketingActivityOrdervoucherUseDefaultResponse struct for AlipayMarketingActivityOrdervoucherUseDefaultResponse
type AlipayMarketingActivityOrdervoucherUseDefaultResponse struct {
	AlipayMarketingActivityOrdervoucherUseErrorResponseModel *AlipayMarketingActivityOrdervoucherUseErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingActivityOrdervoucherUseDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingActivityOrdervoucherUseErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingActivityOrdervoucherUseErrorResponseModel);
	if err == nil {
		jsonAlipayMarketingActivityOrdervoucherUseErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingActivityOrdervoucherUseErrorResponseModel)
		if string(jsonAlipayMarketingActivityOrdervoucherUseErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingActivityOrdervoucherUseErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingActivityOrdervoucherUseErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingActivityOrdervoucherUseErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingActivityOrdervoucherUseDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingActivityOrdervoucherUseDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingActivityOrdervoucherUseErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingActivityOrdervoucherUseErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMarketingActivityOrdervoucherUseDefaultResponse struct {
	value *AlipayMarketingActivityOrdervoucherUseDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingActivityOrdervoucherUseDefaultResponse) Get() *AlipayMarketingActivityOrdervoucherUseDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingActivityOrdervoucherUseDefaultResponse) Set(val *AlipayMarketingActivityOrdervoucherUseDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityOrdervoucherUseDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityOrdervoucherUseDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityOrdervoucherUseDefaultResponse(val *AlipayMarketingActivityOrdervoucherUseDefaultResponse) *NullableAlipayMarketingActivityOrdervoucherUseDefaultResponse {
	return &NullableAlipayMarketingActivityOrdervoucherUseDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityOrdervoucherUseDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityOrdervoucherUseDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

