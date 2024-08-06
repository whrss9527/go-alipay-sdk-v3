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


// AlipayMarketingActivityOrdervoucherQueryDefaultResponse struct for AlipayMarketingActivityOrdervoucherQueryDefaultResponse
type AlipayMarketingActivityOrdervoucherQueryDefaultResponse struct {
	AlipayMarketingActivityOrdervoucherQueryErrorResponseModel *AlipayMarketingActivityOrdervoucherQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingActivityOrdervoucherQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingActivityOrdervoucherQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingActivityOrdervoucherQueryErrorResponseModel);
	if err == nil {
		jsonAlipayMarketingActivityOrdervoucherQueryErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingActivityOrdervoucherQueryErrorResponseModel)
		if string(jsonAlipayMarketingActivityOrdervoucherQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingActivityOrdervoucherQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingActivityOrdervoucherQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingActivityOrdervoucherQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingActivityOrdervoucherQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingActivityOrdervoucherQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingActivityOrdervoucherQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingActivityOrdervoucherQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMarketingActivityOrdervoucherQueryDefaultResponse struct {
	value *AlipayMarketingActivityOrdervoucherQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingActivityOrdervoucherQueryDefaultResponse) Get() *AlipayMarketingActivityOrdervoucherQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingActivityOrdervoucherQueryDefaultResponse) Set(val *AlipayMarketingActivityOrdervoucherQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityOrdervoucherQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityOrdervoucherQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityOrdervoucherQueryDefaultResponse(val *AlipayMarketingActivityOrdervoucherQueryDefaultResponse) *NullableAlipayMarketingActivityOrdervoucherQueryDefaultResponse {
	return &NullableAlipayMarketingActivityOrdervoucherQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityOrdervoucherQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityOrdervoucherQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


