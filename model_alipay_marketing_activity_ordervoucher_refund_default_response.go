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


// AlipayMarketingActivityOrdervoucherRefundDefaultResponse struct for AlipayMarketingActivityOrdervoucherRefundDefaultResponse
type AlipayMarketingActivityOrdervoucherRefundDefaultResponse struct {
	AlipayMarketingActivityOrdervoucherRefundErrorResponseModel *AlipayMarketingActivityOrdervoucherRefundErrorResponseModel
	CommonErrorType *CommonErrorType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AlipayMarketingActivityOrdervoucherRefundDefaultResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AlipayMarketingActivityOrdervoucherRefundErrorResponseModel
	err = json.Unmarshal(data, &dst.AlipayMarketingActivityOrdervoucherRefundErrorResponseModel);
	if err == nil {
		jsonAlipayMarketingActivityOrdervoucherRefundErrorResponseModel, _ := json.Marshal(dst.AlipayMarketingActivityOrdervoucherRefundErrorResponseModel)
		if string(jsonAlipayMarketingActivityOrdervoucherRefundErrorResponseModel) == "{}" { // empty struct
			dst.AlipayMarketingActivityOrdervoucherRefundErrorResponseModel = nil
		} else {
			return nil // data stored in dst.AlipayMarketingActivityOrdervoucherRefundErrorResponseModel, return on the first match
		}
	} else {
		dst.AlipayMarketingActivityOrdervoucherRefundErrorResponseModel = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AlipayMarketingActivityOrdervoucherRefundDefaultResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AlipayMarketingActivityOrdervoucherRefundDefaultResponse) MarshalJSON() ([]byte, error) {
	if src.AlipayMarketingActivityOrdervoucherRefundErrorResponseModel != nil {
		return json.Marshal(&src.AlipayMarketingActivityOrdervoucherRefundErrorResponseModel)
	}

	if src.CommonErrorType != nil {
		return json.Marshal(&src.CommonErrorType)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableAlipayMarketingActivityOrdervoucherRefundDefaultResponse struct {
	value *AlipayMarketingActivityOrdervoucherRefundDefaultResponse
	isSet bool
}

func (v NullableAlipayMarketingActivityOrdervoucherRefundDefaultResponse) Get() *AlipayMarketingActivityOrdervoucherRefundDefaultResponse {
	return v.value
}

func (v *NullableAlipayMarketingActivityOrdervoucherRefundDefaultResponse) Set(val *AlipayMarketingActivityOrdervoucherRefundDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAlipayMarketingActivityOrdervoucherRefundDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAlipayMarketingActivityOrdervoucherRefundDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlipayMarketingActivityOrdervoucherRefundDefaultResponse(val *AlipayMarketingActivityOrdervoucherRefundDefaultResponse) *NullableAlipayMarketingActivityOrdervoucherRefundDefaultResponse {
	return &NullableAlipayMarketingActivityOrdervoucherRefundDefaultResponse{value: val, isSet: true}
}

func (v NullableAlipayMarketingActivityOrdervoucherRefundDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlipayMarketingActivityOrdervoucherRefundDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


