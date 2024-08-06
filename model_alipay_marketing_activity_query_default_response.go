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


// AlipayMarketingActivityQueryDefaultResponse struct for AlipayMarketingActivityQueryDefaultResponse
type AlipayMarketingActivityQueryDefaultResponse struct {
	AlipayMarketingActivityQueryErrorResponseModel *AlipayMarketingActivityQueryErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingActivityQueryDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingActivityQueryErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingActivityQueryErrorResponseModel);
	if err == nil {
		jsonAlipayMarketingActivityQueryErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingActivityQueryErrorResponseModel)
		if string(jsonAlipayMarketingActivityQueryErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingActivityQueryErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingActivityQueryErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingActivityQueryErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingActivityQueryDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingActivityQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingActivityQueryErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingActivityQueryErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMarketingActivityQueryDefaultResponse struct {
	value *AlipayMarketingActivityQueryDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingActivityQueryDefaultResponse) Get() *AlipayMarketingActivityQueryDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingActivityQueryDefaultResponse) Set(val *AlipayMarketingActivityQueryDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityQueryDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityQueryDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityQueryDefaultResponse(val *AlipayMarketingActivityQueryDefaultResponse) *NullableAlipayMarketingActivityQueryDefaultResponse {
	return &NullableAlipayMarketingActivityQueryDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityQueryDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityQueryDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


